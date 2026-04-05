// fix_postal_votes.go fixes the 31 POSTAL VOTE rows in to-review.csv.
//
// Problem: postal vote rows were generated with incorrect fields:
//   Col 1 (UNIQUE CODE)             : P.192_P.192/POSTAL VOTE_UNDI POS_1
//   Col 6 (STATE CONSTITUENCY CODE) : P.192/POSTAL VOTE
//   Col 8 (POLLING DISTRICT CODE)   : UNDI POS
//
// Correct pattern (per AGENTS.md Phase-5 spec):
//   Col 1 (UNIQUE CODE)             : P.192_POS_1
//   Col 6 (STATE CONSTITUENCY CODE) : P.192/POS
//   Col 8 (POLLING DISTRICT CODE)   : 192/POS
//
// All other columns are left unchanged.
//
// Usage:
//   go run fix_postal_votes.go
//
// The script rewrites to-review.csv in-place.

package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	inputFile  = "to-review.csv"
	outputFile = "to-review.csv"
	backupFile = "to-review.csv.bak"
)

// Column indices (0-based)
const (
	colUniqueCode          = 0
	colBallotType          = 2
	colParCode             = 3 // e.g. "P.192"
	colStateConstCode      = 5 // STATE CONSTITUENCY CODE
	colPollingDistrictCode = 7 // POLLING DISTRICT CODE
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// ── 1. Read input ────────────────────────────────────────────────────────
	f, err := os.Open(inputFile)
	if err != nil {
		slog.Error("cannot open input", "file", inputFile, "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1 // tolerate rows with varying field counts during read
	records, err := r.ReadAll()
	f.Close()
	if err != nil {
		slog.Error("cannot read CSV", "err", err)
		os.Exit(1)
	}
	slog.Info("read CSV", "total_rows", len(records))

	// ── 2. Backup original ───────────────────────────────────────────────────
	if err := copyFile(inputFile, backupFile); err != nil {
		slog.Error("cannot create backup", "err", err)
		os.Exit(1)
	}
	slog.Info("backup created", "file", backupFile)

	// ── 3. Process rows ───────────────────────────────────────────────────────
	fixed := 0
	problems := 0

	for i, row := range records {
		if i == 0 {
			// header row — skip
			continue
		}

		if len(row) <= colPollingDistrictCode {
			slog.Warn("row too short, skipping", "row", i+1, "cols", len(row))
			continue
		}

		if row[colBallotType] != "POSTAL VOTE" {
			continue
		}

		parCode := strings.TrimSpace(row[colParCode]) // e.g. "P.192"
		if !strings.HasPrefix(parCode, "P.") {
			slog.Warn("unexpected PAR code format on postal vote row",
				"row", i+1, "par_code", parCode)
			problems++
			continue
		}

		// Extract the numeric part: "P.192" → "192"
		parNum := parCode[2:] // strip "P."

		// Compute corrected values
		wantUniqueCode := parCode + "_POS_1" // e.g. "P.192_POS_1"
		wantStateCode := parCode + "/POS"    // e.g. "P.192/POS"
		wantDistCode := parNum + "/POS"      // e.g. "192/POS"

		// Report before/after
		slog.Debug("fixing postal row",
			"row", i+1,
			"par", parCode,
			"unique_code_old", row[colUniqueCode],
			"unique_code_new", wantUniqueCode,
			"state_code_old", row[colStateConstCode],
			"state_code_new", wantStateCode,
			"dist_code_old", row[colPollingDistrictCode],
			"dist_code_new", wantDistCode,
		)

		records[i][colUniqueCode] = wantUniqueCode
		records[i][colStateConstCode] = wantStateCode
		records[i][colPollingDistrictCode] = wantDistCode
		fixed++
	}

	slog.Info("processing complete", "postal_rows_fixed", fixed, "problems", problems)

	if problems > 0 {
		slog.Warn("some rows had unexpected format — review WARN logs above")
	}

	// ── 4. Write output ───────────────────────────────────────────────────────
	out, err := os.Create(outputFile)
	if err != nil {
		slog.Error("cannot create output file", "err", err)
		os.Exit(1)
	}
	w := csv.NewWriter(out)
	if err := w.WriteAll(records); err != nil {
		out.Close()
		slog.Error("cannot write CSV", "err", err)
		os.Exit(1)
	}
	w.Flush()
	out.Close()
	if err := w.Error(); err != nil {
		slog.Error("CSV flush error", "err", err)
		os.Exit(1)
	}

	slog.Info("output written", "file", outputFile)

	// ── 5. Verify ─────────────────────────────────────────────────────────────
	fmt.Printf("\n=== Verification — first 5 fixed postal vote rows ===\n")
	shown := 0
	for i, row := range records {
		if i == 0 || len(row) <= colPollingDistrictCode {
			continue
		}
		if row[colBallotType] != "POSTAL VOTE" {
			continue
		}
		fmt.Printf("Row %4d  UNIQUE_CODE=%-25s  STATE_CODE=%-12s  DIST_CODE=%s\n",
			i+1,
			row[colUniqueCode],
			row[colStateConstCode],
			row[colPollingDistrictCode],
		)
		shown++
		if shown >= 5 {
			break
		}
	}
}

// copyFile duplicates src → dst.
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read %s: %w", src, err)
	}
	if err := os.WriteFile(dst, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", dst, err)
	}
	return nil
}
