// fix_unique_codes.go rebuilds ALL UNIQUE CODEs in to-review.csv from their component columns,
// applying the duplicate-disambiguation suffix algorithm per the spec.
//
// Suffix format: _CHANNELletter  (e.g. _1a, _2a, _1b, _2b)
//
// Algorithm:
//   1. Group all rows by Polling District Code (col 8).
//   2. For each district, count distinct Polling Centres (col 10).
//   3. If only 1 centre → UNIQUE CODE = {PAR}_{DUN}_{PDC}_{CHANNEL}  (no suffix)
//   4. If >1 centre:
//      a. Build PollingCentre → letter map in order of first appearance in the file.
//      b. Every row in that district gets: {PAR}_{DUN}_{PDC}_{CHANNEL}{letter}
//   5. Postal vote rows use the format: {PAR}_POS_1  (already correct, preserved).
//
// Usage:
//   go run fix_unique_codes.go

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
)

// Column indices (0-based)
const (
	colUniqueCode    = 0
	colBallotType    = 2
	colParCode       = 3  // e.g. "P.192"
	colStateConstCod = 5  // e.g. "N.02" or "P.192/POS"
	colPollDistCode  = 7  // e.g. "192/02/00" or "192/POS"
	colPollCentre    = 9  // e.g. "SEKOLAH KEBANGSAAN BAU"
	colChannel       = 20 // VOTING CHANNEL NUMBER e.g. "1"
)

// districtInfo collects rows and centre ordering for one Polling District Code.
type districtInfo struct {
	rowIndices   []int
	centreOrder  []string          // distinct centres in first-appearance order
	centreToChar map[string]string // centre → suffix letter
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// ── 1. Read ─────────────────────────────────────────────────────────────
	f, err := os.Open(inputFile)
	if err != nil {
		slog.Error("open", "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	f.Close()
	if err != nil {
		slog.Error("read", "err", err)
		os.Exit(1)
	}
	slog.Info("read", "rows", len(records))

	// ── 2. Group non-postal rows by Polling District Code ───────────────────
	// key = Polling District Code (col 8)
	districts := make(map[string]*districtInfo)
	// We need to preserve insertion order for deterministic output.
	var districtOrder []string

	for i, row := range records {
		if i == 0 {
			continue // header
		}
		if len(row) <= colChannel {
			continue
		}
		if row[colBallotType] == "POSTAL VOTE" {
			continue // postal rows already handled
		}

		pdc := strings.TrimSpace(row[colPollDistCode])
		centre := strings.TrimSpace(row[colPollCentre])

		di, exists := districts[pdc]
		if !exists {
			di = &districtInfo{
				centreToChar: make(map[string]string),
			}
			districts[pdc] = di
			districtOrder = append(districtOrder, pdc)
		}
		di.rowIndices = append(di.rowIndices, i)

		// Track centre first-appearance order
		if _, seen := di.centreToChar[centre]; !seen {
			letter := string(rune('a' + len(di.centreOrder)))
			di.centreOrder = append(di.centreOrder, centre)
			di.centreToChar[centre] = letter
		}
	}

	slog.Info("districts collected", "count", len(districts))

	// ── 3. Rebuild UNIQUE CODEs ─────────────────────────────────────────────
	changed := 0
	multiCentreDistricts := 0
	remainingDuplicates := 0

	for _, pdc := range districtOrder {
		di := districts[pdc]
		needsSuffix := len(di.centreOrder) > 1

		if needsSuffix {
			multiCentreDistricts++
		}

		for _, idx := range di.rowIndices {
			row := records[idx]
			par := strings.TrimSpace(row[colParCode])
			dun := strings.TrimSpace(row[colStateConstCod])
			channel := strings.TrimSpace(row[colChannel])
			centre := strings.TrimSpace(row[colPollCentre])

			var newCode string
			if needsSuffix {
				letter := di.centreToChar[centre]
				// Format: P.XXX_N.YY_PDC_CHANNELletter
				newCode = fmt.Sprintf("%s_%s_%s_%s%s", par, dun, pdc, channel, letter)
			} else {
				// Format: P.XXX_N.YY_PDC_CHANNEL (no suffix)
				newCode = fmt.Sprintf("%s_%s_%s_%s", par, dun, pdc, channel)
			}

			oldCode := row[colUniqueCode]
			if oldCode != newCode {
				slog.Debug("change",
					"row", idx+1,
					"old", oldCode,
					"new", newCode,
					"pdc", pdc,
					"centre", centre,
					"channel", channel,
				)
				records[idx][colUniqueCode] = newCode
				changed++
			}
		}
	}

	// ── 4. Check for remaining duplicates ───────────────────────────────────
	seen := make(map[string][]int)
	for i, row := range records {
		if i == 0 {
			continue
		}
		code := row[colUniqueCode]
		if code == "" {
			continue
		}
		seen[code] = append(seen[code], i+1) // 1-based line numbers
	}
	for code, lines := range seen {
		if len(lines) > 1 {
			remainingDuplicates++
			slog.Warn("remaining duplicate",
				"code", code,
				"lines", lines,
				"count", len(lines),
			)
		}
	}

	slog.Info("rebuild complete",
		"changed", changed,
		"multi_centre_districts", multiCentreDistricts,
		"remaining_duplicates", remainingDuplicates,
	)

	// ── 5. Write ────────────────────────────────────────────────────────────
	out, err := os.Create(outputFile)
	if err != nil {
		slog.Error("create output", "err", err)
		os.Exit(1)
	}
	w := csv.NewWriter(out)
	if err := w.WriteAll(records); err != nil {
		out.Close()
		slog.Error("write", "err", err)
		os.Exit(1)
	}
	w.Flush()
	out.Close()
	if err := w.Error(); err != nil {
		slog.Error("flush", "err", err)
		os.Exit(1)
	}
	slog.Info("written", "file", outputFile)

	// ── 6. Print summary ────────────────────────────────────────────────────
	fmt.Println()
	fmt.Println("=== Summary ===")
	fmt.Printf("Total data rows:            %d\n", len(records)-1)
	fmt.Printf("UNIQUE CODEs changed:       %d\n", changed)
	fmt.Printf("Multi-centre districts:     %d\n", multiCentreDistricts)
	fmt.Printf("Remaining duplicate codes:  %d\n", remainingDuplicates)

	if remainingDuplicates > 0 {
		fmt.Println()
		fmt.Println("=== Remaining Duplicates (need manual investigation) ===")
		for code, lines := range seen {
			if len(lines) > 1 {
				fmt.Printf("  %s  → lines %v\n", code, lines)
			}
		}
	}

	// ── 7. Show sample of multi-centre districts ────────────────────────────
	fmt.Println()
	fmt.Println("=== Sample: first 5 multi-centre districts ===")
	shown := 0
	for _, pdc := range districtOrder {
		di := districts[pdc]
		if len(di.centreOrder) <= 1 {
			continue
		}
		fmt.Printf("\n  District %s  (%d centres, %d rows)\n", pdc, len(di.centreOrder), len(di.rowIndices))
		for _, c := range di.centreOrder {
			fmt.Printf("    '%s' → %s\n", di.centreToChar[c], c)
		}
		// Show first few rows
		limit := 6
		if limit > len(di.rowIndices) {
			limit = len(di.rowIndices)
		}
		for _, idx := range di.rowIndices[:limit] {
			fmt.Printf("    row %4d: %s\n", idx+1, records[idx][colUniqueCode])
		}
		if len(di.rowIndices) > limit {
			fmt.Printf("    ... and %d more rows\n", len(di.rowIndices)-limit)
		}
		shown++
		if shown >= 5 {
			break
		}
	}
}
