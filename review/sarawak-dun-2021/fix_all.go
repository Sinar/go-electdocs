// fix_all.go applies three corrections to to-review.csv:
//
//	Fix 1 – Duplicate UNIQUE CODEs: append letter suffixes (a, b, c …) to every
//	         occurrence of each duplicate UNIQUE CODE, grouped by Polling Centre.
//	         Suffix assignment follows first-appearance order of each distinct
//	         Polling Centre.  Only column 1 (UNIQUE CODE) is modified.
//
//	Fix 2 – BA`KELALAN backtick: replace the grave-accent character (U+0060)
//	         with a standard apostrophe (U+0027) wherever BA`KELALAN appears
//	         in any cell (primarily col 7 STATE CONSTITUENCY NAME and
//	         col 9 POLLING DISTRICT NAME for N.81 rows).
//
//	Fix 3 – N.01 postal vote DMKOD: change the polling-district suffix from
//	         the old "/POS" format to the current "/UNDI POS" format used by
//	         all other 81 DUNs, affecting UNIQUE CODE (col 1) and POLLING
//	         DISTRICT CODE (col 8) of the single N.01 POSTAL VOTE row.
//
// Usage (run from the sarawak-dun-2021 directory):
//
//	go run fix_all.go
//
// The input file (to-review.csv) is overwritten in place.
// A backup is written to to-review.csv.bak before any changes are made.
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
	colUniqueCode = 0 // UNIQUE CODE
	colBallotType = 2 // BALLOT TYPE
	colDUNCode    = 5 // STATE CONSTITUENCY CODE
	colPDCode     = 7 // POLLING DISTRICT CODE
	colPollingCtr = 9 // POLLING CENTRE
	expectedCols  = 67
)

func main() {
	slog.Info("fix_all.go starting")

	const inputPath = "to-review.csv"
	const backupPath = "to-review.csv.bak"

	// ── Read ──────────────────────────────────────────────────────────────────
	f, err := os.Open(inputPath)
	if err != nil {
		slog.Error("cannot open input file", "path", inputPath, "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1 // allow variable (we validate below)
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

	// Validate column count
	for i, row := range data {
		if len(row) != expectedCols {
			slog.Warn("unexpected column count", "data_row", i+1, "cols", len(row))
		}
	}

	// ── Backup ────────────────────────────────────────────────────────────────
	if err := copyFile(inputPath, backupPath); err != nil {
		slog.Error("cannot create backup", "src", inputPath, "dst", backupPath, "err", err)
		os.Exit(1)
	}
	slog.Info("backup created", "path", backupPath)

	// ══════════════════════════════════════════════════════════════════════════
	// FIX 2 – BA`KELALAN backtick → apostrophe
	// Applied before Fix 3 and Fix 1 so that UNIQUE CODE values are already
	// normalised when we scan for duplicates.
	// ══════════════════════════════════════════════════════════════════════════
	fix2Cells := 0
	for i := range data {
		for j := range data[i] {
			if strings.Contains(data[i][j], "BA`KELALAN") {
				before := data[i][j]
				data[i][j] = strings.ReplaceAll(data[i][j], "BA`KELALAN", "BA'KELALAN")
				if data[i][j] != before {
					fix2Cells++
					slog.Debug("fix2 backtick replaced",
						"data_row", i+1, "col", j+1,
						"before", before, "after", data[i][j])
				}
			}
		}
	}
	slog.Info("Fix 2 complete (BA'KELALAN backtick)", "cells_fixed", fix2Cells)

	// ══════════════════════════════════════════════════════════════════════════
	// FIX 3 – N.01 postal vote DMKOD: 192/01/POS → 192/01/UNDI POS
	// Affects UNIQUE CODE (col 1) and POLLING DISTRICT CODE (col 8).
	// ══════════════════════════════════════════════════════════════════════════
	fix3Rows := 0
	for i := range data {
		row := data[i]
		if len(row) <= colPDCode {
			continue
		}
		if row[colDUNCode] == "N.01" && row[colBallotType] == "POSTAL VOTE" {
			oldUC := row[colUniqueCode]
			oldPD := row[colPDCode]

			// UNIQUE CODE: …/POS_  →  …/UNDI POS_
			row[colUniqueCode] = strings.Replace(row[colUniqueCode], "/POS_", "/UNDI POS_", 1)
			// POLLING DISTRICT CODE: …/POS  →  …/UNDI POS
			row[colPDCode] = strings.Replace(row[colPDCode], "/POS", "/UNDI POS", 1)

			if row[colUniqueCode] != oldUC || row[colPDCode] != oldPD {
				fix3Rows++
				slog.Info("fix3 N.01 postal DMKOD corrected",
					"data_row", i+1,
					"old_unique_code", oldUC, "new_unique_code", row[colUniqueCode],
					"old_pd_code", oldPD, "new_pd_code", row[colPDCode])
			}
		}
	}
	slog.Info("Fix 3 complete (N.01 postal DMKOD)", "rows_fixed", fix3Rows)

	// ══════════════════════════════════════════════════════════════════════════
	// FIX 1 – Duplicate UNIQUE CODEs with letter suffixes
	//
	// Algorithm:
	//   Pass 1: tally how many times each UNIQUE CODE appears → duplicates set.
	//   Pass 2: for each duplicate code, maintain a map[pollingCentre]letter.
	//           Walk rows in file order; when we encounter a duplicate UNIQUE
	//           CODE, look up (or assign) the letter for its Polling Centre and
	//           append it to UNIQUE CODE.
	//
	// Rules enforced:
	//   • Same Polling Centre always gets the same letter within a duplicate group.
	//   • Letters are assigned a, b, c … in first-appearance order of centres.
	//   • Only col 1 (UNIQUE CODE) is modified.
	// ══════════════════════════════════════════════════════════════════════════

	// Pass 1: count occurrences
	codeFreq := make(map[string]int, len(data))
	for _, row := range data {
		if len(row) == 0 {
			continue
		}
		code := row[colUniqueCode]
		if code != "" {
			codeFreq[code]++
		}
	}

	// Build the set of duplicate codes
	isDup := make(map[string]bool)
	dupCount := 0
	for code, freq := range codeFreq {
		if freq > 1 {
			isDup[code] = true
			dupCount++
		}
	}
	slog.Info("Fix 1 duplicate analysis", "distinct_duplicate_codes", dupCount)

	// Per-code state: map[uniqueCode] → map[pollingCentre]suffix_letter
	type codeState struct {
		centreToLetter map[string]string
		nextIdx        int // 0→'a', 1→'b', …
	}
	states := make(map[string]*codeState, dupCount)
	for code := range isDup {
		states[code] = &codeState{centreToLetter: make(map[string]string)}
	}

	// Pass 2: assign and apply suffixes
	fix1Rows := 0
	for i := range data {
		row := data[i]
		if len(row) <= colPollingCtr {
			continue
		}
		code := row[colUniqueCode]
		if code == "" {
			continue
		}
		state, dup := states[code]
		if !dup {
			continue
		}

		centre := row[colPollingCtr]
		letter, seen := state.centreToLetter[centre]
		if !seen {
			if state.nextIdx > 25 {
				slog.Error("more than 26 distinct polling centres for one UNIQUE CODE",
					"code", code)
				os.Exit(1)
			}
			letter = string(rune('a' + state.nextIdx))
			state.centreToLetter[centre] = letter
			state.nextIdx++
			slog.Debug("fix1 new suffix assigned",
				"code", code, "centre", centre, "suffix", letter)
		}

		data[i][colUniqueCode] = code + letter
		fix1Rows++
	}
	slog.Info("Fix 1 complete (UNIQUE CODE suffixes)", "rows_modified", fix1Rows)

	// ── Post-fix validation ───────────────────────────────────────────────────
	finalCodes := make(map[string]int, len(data))
	for _, row := range data {
		if len(row) > 0 && row[colUniqueCode] != "" {
			finalCodes[row[colUniqueCode]]++
		}
	}
	remainingDups := 0
	for _, freq := range finalCodes {
		if freq > 1 {
			remainingDups++
		}
	}
	if remainingDups > 0 {
		slog.Error("UNIQUE CODE still has duplicates after Fix 1!", "count", remainingDups)
	} else {
		slog.Info("validation: all UNIQUE CODEs are now distinct")
	}

	// Verify N.01 postal row
	n01PostalFixed := false
	for _, row := range data {
		if len(row) > colPDCode && row[colDUNCode] == "N.01" && row[colBallotType] == "POSTAL VOTE" {
			if strings.Contains(row[colUniqueCode], "UNDI POS") &&
				strings.Contains(row[colPDCode], "UNDI POS") {
				n01PostalFixed = true
			}
		}
	}
	if n01PostalFixed {
		slog.Info("validation: N.01 postal vote DMKOD correctly uses UNDI POS")
	} else {
		slog.Warn("validation: N.01 postal vote DMKOD may not be fixed")
	}

	// Verify no remaining BA`KELALAN
	backtickRemaining := 0
	for _, row := range data {
		for _, cell := range row {
			if strings.Contains(cell, "BA`KELALAN") {
				backtickRemaining++
			}
		}
	}
	if backtickRemaining == 0 {
		slog.Info("validation: no remaining BA`KELALAN backtick instances")
	} else {
		slog.Warn("validation: BA`KELALAN backtick still present", "cells", backtickRemaining)
	}

	// ── Write output ──────────────────────────────────────────────────────────
	out, err := os.Create(inputPath)
	if err != nil {
		slog.Error("cannot create output file", "path", inputPath, "err", err)
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

	slog.Info("to-review.csv written successfully",
		"total_rows", len(data),
		"fix1_rows_modified", fix1Rows,
		"fix2_cells_modified", fix2Cells,
		"fix3_rows_modified", fix3Rows)

	fmt.Printf("\n=== Fix Summary ===\n")
	fmt.Printf("Fix 1 (UNIQUE CODE suffixes):  %d rows modified\n", fix1Rows)
	fmt.Printf("Fix 2 (BA'KELALAN backtick):   %d cells modified\n", fix2Cells)
	fmt.Printf("Fix 3 (N.01 postal DMKOD):     %d rows modified\n", fix3Rows)
	fmt.Printf("Output: %s  (backup: %s)\n", inputPath, backupPath)
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
