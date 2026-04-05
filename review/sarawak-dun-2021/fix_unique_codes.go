// fix_unique_codes.go correctly re-applies the suffix rule for duplicate UNIQUE
// CODEs in to-review.csv.
//
// THE CORRECT RULE
// ────────────────
// For each Polling District Code (col 8):
//  1. Collect all distinct Polling Centres (col 10) in first-appearance order
//     across the file.
//  2. If the district has 2+ distinct Polling Centres, assign letters a, b, c …
//     to those centres.
//  3. Stamp that letter on EVERY row for that district+centre combination —
//     across ALL channel numbers — not just the rows that happened to form
//     duplicates.
//
// WHY THIS REPLACES THE PREVIOUS FIX
// ────────────────────────────────────
// The previous fix (fix_all.go) treated each channel number independently:
// _1 duplicates formed one group, _2 duplicates another, etc.  Letters were
// assigned separately per group, so the same Polling Centre could receive
// different letters in different groups (or no letter at all for unique
// channels), producing inconsistencies such as _1a alongside _2 for the same
// centre.
//
// WHAT THIS PROGRAM DOES
// ──────────────────────
//
//	Step 1  Strip any trailing letter suffix added by the previous fix.
//	        A suffix looks like the last character being [a-z] after a digit,
//	        e.g. "P.215_N.61_215/61/04_1a" → "P.215_N.61_215/61/04_1".
//	Step 2  Build a per-district centre→letter map from the stripped data,
//	        visiting rows in file order.
//	Step 3  Apply the letter to EVERY row whose district has 2+ centres.
//	Step 4  Validate that all resulting UNIQUE CODEs are distinct.
//	Step 5  Write the corrected file; backup saved as to-review.csv.bak4.
//
// Usage (from the sarawak-dun-2021 directory):
//
//	go run fix_unique_codes.go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

// Column indices (0-based).
const (
	colUniqueCode = 0 // UNIQUE CODE
	colPDCode     = 7 // POLLING DISTRICT CODE
	colPCentre    = 9 // POLLING CENTRE
	expectedCols  = 67
)

// suffixRe matches a UNIQUE CODE that already carries a single trailing letter
// suffix appended immediately after the channel number, e.g.:
//
//	"P.215_N.61_215/61/04_1a"   →  base="P.215_N.61_215/61/04_1"  letter="a"
//	"P.193_N.04_193/04/01_12b"  →  base="P.193_N.04_193/04/01_12" letter="b"
//
// Codes without a suffix (ending in just digits) do not match.
var suffixRe = regexp.MustCompile(`^(.*_\d+)([a-z])$`)

func main() {
	slog.Info("fix_unique_codes.go starting")

	const inputPath = "to-review.csv"
	const backupPath = "to-review.csv.bak4"

	// ── 1. Read ───────────────────────────────────────────────────────────────
	f, err := os.Open(inputPath)
	if err != nil {
		slog.Error("cannot open input file", "path", inputPath, "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	f.Close()
	if err != nil {
		slog.Error("failed to read CSV", "err", err)
		os.Exit(1)
	}
	if len(records) < 2 {
		slog.Error("CSV has no data rows")
		os.Exit(1)
	}

	header := records[0]
	data := records[1:]
	slog.Info("CSV loaded", "header_cols", len(header), "data_rows", len(data))

	for i, row := range data {
		if len(row) != expectedCols {
			slog.Warn("unexpected column count", "data_row", i+1, "cols", len(row))
		}
	}

	// ── 2. Backup ─────────────────────────────────────────────────────────────
	if err := copyFile(inputPath, backupPath); err != nil {
		slog.Error("cannot create backup", "err", err)
		os.Exit(1)
	}
	slog.Info("backup created", "path", backupPath)

	// ── 3. Strip existing letter suffixes ─────────────────────────────────────
	stripped := 0
	for i, row := range data {
		if len(row) <= colUniqueCode {
			continue
		}
		uc := row[colUniqueCode]
		if m := suffixRe.FindStringSubmatch(uc); m != nil {
			data[i][colUniqueCode] = m[1] // drop the trailing letter (m[2])
			stripped++
			slog.Debug("stripped suffix",
				"data_row", i+1,
				"before", uc,
				"after", data[i][colUniqueCode],
			)
		}
	}
	slog.Info("Step 1 complete: stripped existing suffixes", "rows_stripped", stripped)

	// ── 4. Build district → centre → letter mapping ───────────────────────────
	//
	// Pass A: for each Polling District Code, collect distinct Polling Centres
	//         in first-appearance order.
	type districtInfo struct {
		centreOrder []string       // centres in first-appearance order
		centreIndex map[string]int // centre → position (0-based)
	}

	districtMap := make(map[string]*districtInfo)

	for _, row := range data {
		if len(row) <= colPCentre {
			continue
		}
		pd := row[colPDCode]
		pc := row[colPCentre]
		if pd == "" {
			continue
		}

		info, exists := districtMap[pd]
		if !exists {
			info = &districtInfo{
				centreIndex: make(map[string]int),
			}
			districtMap[pd] = info
		}
		if _, seen := info.centreIndex[pc]; !seen {
			info.centreIndex[pc] = len(info.centreOrder)
			info.centreOrder = append(info.centreOrder, pc)
		}
	}

	// Count multi-centre districts
	multiCentreDistricts := 0
	for _, info := range districtMap {
		if len(info.centreOrder) > 1 {
			multiCentreDistricts++
		}
	}
	slog.Info("Step 2 complete: district mapping built",
		"total_districts", len(districtMap),
		"multi_centre_districts", multiCentreDistricts,
	)

	// Verify no district has more than 26 distinct centres (letter limit).
	for pd, info := range districtMap {
		if len(info.centreOrder) > 26 {
			slog.Error("district has more than 26 distinct polling centres — cannot suffix with a-z",
				"district", pd,
				"centre_count", len(info.centreOrder),
			)
			os.Exit(1)
		}
	}

	// ── 5. Apply suffix to all rows in multi-centre districts ─────────────────
	rowsModified := 0
	rowsUnchanged := 0

	for i, row := range data {
		if len(row) <= colPCentre {
			continue
		}
		pd := row[colPDCode]
		pc := row[colPCentre]

		info, exists := districtMap[pd]
		if !exists || len(info.centreOrder) < 2 {
			// Single-centre district: no suffix needed.
			rowsUnchanged++
			continue
		}

		idx, found := info.centreIndex[pc]
		if !found {
			// Shouldn't happen, but be safe.
			slog.Warn("polling centre not found in district map",
				"data_row", i+1, "district", pd, "centre", pc)
			rowsUnchanged++
			continue
		}

		letter := string(rune('a' + idx))
		oldUC := data[i][colUniqueCode]
		data[i][colUniqueCode] = oldUC + letter
		rowsModified++

		slog.Debug("applied suffix",
			"data_row", i+1,
			"district", pd,
			"centre", pc,
			"suffix", letter,
			"unique_code", data[i][colUniqueCode],
		)
	}

	slog.Info("Step 3 complete: suffixes applied",
		"rows_modified", rowsModified,
		"rows_unchanged", rowsUnchanged,
	)

	// ── 6. Validate uniqueness ────────────────────────────────────────────────
	codeCount := make(map[string]int, len(data))
	for _, row := range data {
		if len(row) > colUniqueCode && row[colUniqueCode] != "" {
			codeCount[row[colUniqueCode]]++
		}
	}

	duplicatesRemaining := 0
	for code, count := range codeCount {
		if count > 1 {
			slog.Error("UNIQUE CODE still duplicate after fix",
				"code", code, "count", count)
			duplicatesRemaining++
		}
	}

	if duplicatesRemaining == 0 {
		slog.Info("validation PASSED: all UNIQUE CODEs are distinct",
			"distinct_codes", len(codeCount))
	} else {
		slog.Error("validation FAILED: duplicates remain", "count", duplicatesRemaining)
	}

	// Validate that rows NOT modified (single-centre districts) still have
	// no duplicates among themselves — just a double-check.
	slog.Info("Step 4 complete: validation done",
		"duplicate_codes_remaining", duplicatesRemaining,
	)

	// ── 7. Report suffix mapping for multi-centre districts ───────────────────
	// Log a few examples so the output is easy to audit.
	logged := 0
	for pd, info := range districtMap {
		if len(info.centreOrder) < 2 {
			continue
		}
		for idx, centre := range info.centreOrder {
			letter := string(rune('a' + idx))
			slog.Debug("district centre mapping",
				"district", pd, "centre", centre, "letter", letter)
		}
		logged++
		if logged >= 10 {
			slog.Debug("(further district mappings omitted from log)")
			break
		}
	}

	// ── 8. Write output ───────────────────────────────────────────────────────
	out, err := os.Create(inputPath)
	if err != nil {
		slog.Error("cannot create output file", "err", err)
		os.Exit(1)
	}
	w := csv.NewWriter(out)

	if err := w.Write(header); err != nil {
		out.Close()
		slog.Error("failed to write header", "err", err)
		os.Exit(1)
	}
	for _, row := range data {
		if err := w.Write(row); err != nil {
			out.Close()
			slog.Error("failed to write row", "err", err)
			os.Exit(1)
		}
	}
	w.Flush()
	out.Close()
	if err := w.Error(); err != nil {
		slog.Error("CSV writer flush error", "err", err)
		os.Exit(1)
	}

	slog.Info("to-review.csv written successfully", "total_rows", len(data))

	fmt.Printf("\n=== Unique Code Fix Summary ===\n")
	fmt.Printf("Existing suffixes stripped:   %d rows\n", stripped)
	fmt.Printf("Multi-centre districts:       %d\n", multiCentreDistricts)
	fmt.Printf("Rows with suffix applied:     %d\n", rowsModified)
	fmt.Printf("Rows unchanged (no suffix):   %d\n", rowsUnchanged)
	fmt.Printf("Duplicate codes remaining:    %d\n", duplicatesRemaining)
	fmt.Printf("Output: %s  (backup: %s)\n", inputPath, backupPath)

	if duplicatesRemaining > 0 {
		os.Exit(1)
	}
}

// copyFile copies src to dst.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open src: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create dst: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	return out.Sync()
}

// letterFor returns the lowercase letter (a, b, c, …) for the given 0-based index.
// Provided as a named helper so the intent is clear at call sites.
func letterFor(idx int) string {
	return strings.ToLower(string(rune('a' + idx)))
}
