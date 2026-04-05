// fix_postal_dmkod.go corrects the POLLING DISTRICT CODE and UNIQUE CODE for
// all postal vote rows in to-review.csv.
//
// The correct rule for postal votes is:
//
//	UNIQUE CODE          → P.193_N.04_193/04/POS_1   (suffix /POS_N)
//	POLLING DISTRICT CODE → 193/04/POS               (suffix /POS)
//	POLLING DISTRICT NAME → UNDI POS                 (unchanged)
//	POLLING CENTRE        → UNDI POS                 (unchanged)
//
// All 82 postal vote rows currently have "/UNDI POS" in both UNIQUE CODE and
// POLLING DISTRICT CODE. This program replaces "/UNDI POS" with "/POS" in
// those two columns only, leaving all other columns untouched.
//
// Usage (run from the sarawak-dun-2021 directory):
//
//	go run fix_postal_dmkod.go
//
// The input file (to-review.csv) is overwritten in place.
// A backup is written to to-review.csv.bak3 before any changes are made.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// Column indices (0-based)
const (
	colUniqueCode   = 0 // UNIQUE CODE
	colBallotType   = 2 // BALLOT TYPE
	colPDCode       = 7 // POLLING DISTRICT CODE
	colPDName       = 8 // POLLING DISTRICT NAME
	colPollingCentr = 9 // POLLING CENTRE
)

func main() {
	slog.Info("fix_postal_dmkod.go starting")

	const inputPath = "to-review.csv"
	const backupPath = "to-review.csv.bak3"

	// ── Read ──────────────────────────────────────────────────────────────────
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

	// ── Backup ────────────────────────────────────────────────────────────────
	if err := copyFile(inputPath, backupPath); err != nil {
		slog.Error("cannot create backup", "err", err)
		os.Exit(1)
	}
	slog.Info("backup created", "path", backupPath)

	// ── Apply fix ─────────────────────────────────────────────────────────────
	// For every POSTAL VOTE row, replace "/UNDI POS" with "/POS" in:
	//   col 1  UNIQUE CODE           e.g. P.192_N.01_192/01/UNDI POS_1 → P.192_N.01_192/01/POS_1
	//   col 8  POLLING DISTRICT CODE e.g. 192/01/UNDI POS              → 192/01/POS
	// POLLING DISTRICT NAME (col 9) and POLLING CENTRE (col 10) keep "UNDI POS".

	rowsFixed := 0
	rowsSkipped := 0 // postal rows that didn't need changing (already correct)

	for i, row := range data {
		if len(row) <= colPollingCentr {
			continue
		}
		if row[colBallotType] != "POSTAL VOTE" {
			continue
		}

		oldUC := row[colUniqueCode]
		oldPD := row[colPDCode]

		newUC := strings.Replace(oldUC, "/UNDI POS_", "/POS_", 1)
		newPD := strings.Replace(oldPD, "/UNDI POS", "/POS", 1)

		if newUC == oldUC && newPD == oldPD {
			// Already correct (no "/UNDI POS" found) — sanity-log it.
			slog.Debug("postal row already correct, no change needed",
				"data_row", i+1,
				"unique_code", oldUC,
				"pd_code", oldPD,
			)
			rowsSkipped++
			continue
		}

		data[i][colUniqueCode] = newUC
		data[i][colPDCode] = newPD
		rowsFixed++

		slog.Info("fixed postal row",
			"data_row", i+1,
			"old_unique_code", oldUC,
			"new_unique_code", newUC,
			"old_pd_code", oldPD,
			"new_pd_code", newPD,
		)
	}

	slog.Info("fix pass complete",
		"rows_fixed", rowsFixed,
		"rows_already_correct", rowsSkipped,
	)

	// ── Post-fix validation ───────────────────────────────────────────────────
	problems := 0
	for _, row := range data {
		if len(row) <= colPollingCentr {
			continue
		}
		if row[colBallotType] != "POSTAL VOTE" {
			continue
		}

		uc := row[colUniqueCode]
		pd := row[colPDCode]
		pdName := row[colPDName]
		pc := row[colPollingCentr]

		// UNIQUE CODE must contain "/POS_", not "/UNDI POS_"
		if strings.Contains(uc, "/UNDI POS_") {
			slog.Error("validation: UNIQUE CODE still contains /UNDI POS_", "value", uc)
			problems++
		} else if !strings.Contains(uc, "/POS_") {
			slog.Warn("validation: UNIQUE CODE does not contain /POS_", "value", uc)
			problems++
		}

		// POLLING DISTRICT CODE must end with "/POS", not "/UNDI POS"
		if strings.Contains(pd, "UNDI POS") {
			slog.Error("validation: POLLING DISTRICT CODE still contains UNDI POS", "value", pd)
			problems++
		} else if !strings.HasSuffix(pd, "/POS") {
			slog.Warn("validation: POLLING DISTRICT CODE does not end with /POS", "value", pd)
			problems++
		}

		// POLLING DISTRICT NAME and POLLING CENTRE must still say "UNDI POS"
		if pdName != "UNDI POS" {
			slog.Warn("validation: POLLING DISTRICT NAME is not UNDI POS", "value", pdName)
			problems++
		}
		if pc != "UNDI POS" {
			slog.Warn("validation: POLLING CENTRE is not UNDI POS", "value", pc)
			problems++
		}
	}

	if problems == 0 {
		slog.Info("validation: all postal vote rows are correctly formatted")
	} else {
		slog.Error("validation: problems found", "count", problems)
	}

	// ── Write output ──────────────────────────────────────────────────────────
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
			slog.Error("failed to write data row", "err", err)
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

	fmt.Printf("\n=== Postal DMKOD Fix Summary ===\n")
	fmt.Printf("Postal rows fixed:           %d\n", rowsFixed)
	fmt.Printf("Postal rows already correct: %d\n", rowsSkipped)
	fmt.Printf("Validation problems:         %d\n", problems)
	fmt.Printf("Output: %s  (backup: %s)\n", inputPath, backupPath)
	if problems > 0 {
		os.Exit(1)
	}
}

// copyFile copies src to dst, creating dst if it does not exist.
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
