package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	toReviewPath = "to-review.csv"
	data2016Dir  = "/Users/leow/TINDAKMSIA/go-electdocs/data/sarawak-dun-2016/OUTPUT"
	numCompCols  = 11 // columns 0-10 (UNIQUE CODE through VOTING CHANNEL NUMBER)
)

var colNames = []string{
	"UNIQUE CODE",
	"STATE",
	"BALLOT TYPE",
	"PARLIAMENTARY CONSTITUENCY CODE",
	"PARLIAMENTARY CONSTITUENCY NAME",
	"STATE CONSTITUENCY CODE",
	"STATE CONSTITUENCY NAME",
	"POLLING DISTRICT CODE",
	"POLLING DISTRICT NAME",
	"POLLING CENTRE",
	"VOTING CHANNEL NUMBER",
}

// Row holds the first 11 columns of a CSV row
type Row struct {
	Cols [11]string
}

// Diff records a single column difference between 2016 and 2021
type Diff struct {
	UniqueCode string
	ColIndex   int
	ColName    string
	Val2016    string
	Val2021    string
}

// DUNResult holds comparison results for a single DUN
type DUNResult struct {
	DUNCode       string
	DUNName2021   string
	DUNName2016   string
	OnlyIn2021    []Row
	OnlyIn2016    []Row
	Diffs         []Diff
	MatchedCount  int
	Total2021     int
	Total2016     int
	File2016Found bool
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// Step 1: Read 2021 to-review.csv
	slog.Info("reading 2021 to-review.csv", "path", toReviewPath)
	rows2021, err := readCSV(toReviewPath)
	if err != nil {
		slog.Error("failed to read 2021 CSV", "error", err)
		os.Exit(1)
	}
	slog.Info("loaded 2021 data", "rows", len(rows2021))

	// Step 2: Group 2021 rows by STATE CONSTITUENCY CODE (col 5)
	grouped2021 := make(map[string][]Row)
	dunNames2021 := make(map[string]string)
	for _, r := range rows2021 {
		dunCode := r.Cols[5]
		grouped2021[dunCode] = append(grouped2021[dunCode], r)
		if dunNames2021[dunCode] == "" {
			dunNames2021[dunCode] = r.Cols[6]
		}
	}
	slog.Info("grouped 2021 data by DUN", "dunCount", len(grouped2021))

	// Step 3: Process each DUN from N.01 to N.82
	var results []DUNResult
	totalMatched := 0
	totalDiffs := 0
	totalOnlyIn2021 := 0
	totalOnlyIn2016 := 0
	dunsWithDiffs := 0
	dunsNoFile := 0

	for i := 1; i <= 82; i++ {
		dunCode := fmt.Sprintf("N.%02d", i)
		dunName := dunNames2021[dunCode]

		result := DUNResult{
			DUNCode:     dunCode,
			DUNName2021: dunName,
			Total2021:   len(grouped2021[dunCode]),
		}

		// Read 2016 file
		fname := fmt.Sprintf("Sarawak-%s.csv", dunCode)
		fpath := filepath.Join(data2016Dir, fname)

		rows2016, err := readCSV(fpath)
		if err != nil {
			if os.IsNotExist(err) {
				slog.Warn("2016 file not found", "dun", dunCode, "path", fpath)
				result.File2016Found = false
				// All 2021 rows are "new"
				result.OnlyIn2021 = grouped2021[dunCode]
				results = append(results, result)
				totalOnlyIn2021 += len(grouped2021[dunCode])
				dunsNoFile++
				continue
			}
			slog.Error("failed to read 2016 file", "dun", dunCode, "error", err)
			os.Exit(1)
		}
		result.File2016Found = true
		result.Total2016 = len(rows2016)

		if len(rows2016) > 0 {
			result.DUNName2016 = rows2016[0].Cols[6]
		}

		slog.Info("comparing DUN", "dun", dunCode, "rows2021", result.Total2021, "rows2016", result.Total2016)

		// Build maps by UNIQUE CODE (col 0)
		map2016 := make(map[string]Row)
		for _, r := range rows2016 {
			map2016[r.Cols[0]] = r
		}
		map2021 := make(map[string]Row)
		for _, r := range grouped2021[dunCode] {
			map2021[r.Cols[0]] = r
		}

		// Find matches, only-in-2021, only-in-2016
		for _, r21 := range grouped2021[dunCode] {
			uid := r21.Cols[0]
			r16, found := map2016[uid]
			if !found {
				result.OnlyIn2021 = append(result.OnlyIn2021, r21)
				continue
			}
			result.MatchedCount++
			// Compare cols 1..10 (skip col 0 which is the key)
			for c := 1; c < numCompCols; c++ {
				v21 := strings.TrimSpace(r21.Cols[c])
				v16 := strings.TrimSpace(r16.Cols[c])
				if v21 != v16 {
					result.Diffs = append(result.Diffs, Diff{
						UniqueCode: uid,
						ColIndex:   c,
						ColName:    colNames[c],
						Val2016:    v16,
						Val2021:    v21,
					})
				}
			}
		}

		for _, r16 := range rows2016 {
			uid := r16.Cols[0]
			if _, found := map2021[uid]; !found {
				result.OnlyIn2016 = append(result.OnlyIn2016, r16)
			}
		}

		// Sort for deterministic output
		sort.Slice(result.OnlyIn2021, func(i, j int) bool {
			return result.OnlyIn2021[i].Cols[0] < result.OnlyIn2021[j].Cols[0]
		})
		sort.Slice(result.OnlyIn2016, func(i, j int) bool {
			return result.OnlyIn2016[i].Cols[0] < result.OnlyIn2016[j].Cols[0]
		})

		totalMatched += result.MatchedCount
		totalDiffs += len(result.Diffs)
		totalOnlyIn2021 += len(result.OnlyIn2021)
		totalOnlyIn2016 += len(result.OnlyIn2016)

		if len(result.Diffs) > 0 || len(result.OnlyIn2021) > 0 || len(result.OnlyIn2016) > 0 {
			dunsWithDiffs++
		}

		results = append(results, result)
	}

	// Step 4: Output summary report
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("PHASE-0: COMPARISON OF 2016 vs 2021 SARAWAK DUN DATA (COLUMNS 1-11)")
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println()

	fmt.Println("## OVERALL SUMMARY")
	fmt.Printf("Total DUNs processed:          %d\n", len(results))
	fmt.Printf("DUNs with 2016 file found:     %d\n", len(results)-dunsNoFile)
	fmt.Printf("DUNs with NO 2016 file:        %d\n", dunsNoFile)
	fmt.Printf("Total matched rows (by ID):    %d\n", totalMatched)
	fmt.Printf("Total column differences:      %d\n", totalDiffs)
	fmt.Printf("Total rows only in 2021:       %d\n", totalOnlyIn2021)
	fmt.Printf("Total rows only in 2016:       %d\n", totalOnlyIn2016)
	fmt.Printf("DUNs with any difference:      %d\n", dunsWithDiffs)
	fmt.Println()

	// Step 5: Per-DUN report
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("## PER-DUN DETAIL")
	fmt.Println("=" + strings.Repeat("=", 79))

	for _, res := range results {
		hasDiffs := len(res.Diffs) > 0 || len(res.OnlyIn2021) > 0 || len(res.OnlyIn2016) > 0 || !res.File2016Found

		if !hasDiffs {
			continue
		}

		fmt.Println()
		fmt.Println("-" + strings.Repeat("-", 79))
		dunLabel := res.DUNName2021
		if dunLabel == "" {
			dunLabel = "(not in 2021)"
		}
		fmt.Printf("### %s - %s\n", res.DUNCode, dunLabel)

		if !res.File2016Found {
			fmt.Printf("  ** NO 2016 FILE (did not compete in 2016) **\n")
			fmt.Printf("  2021 rows: %d (all new)\n", res.Total2021)
			continue
		}

		fmt.Printf("  2021 rows: %d | 2016 rows: %d | Matched: %d\n", res.Total2021, res.Total2016, res.MatchedCount)

		if res.DUNName2016 != "" && res.DUNName2016 != res.DUNName2021 {
			fmt.Printf("  ** DUN NAME CHANGE: 2016=\"%s\" -> 2021=\"%s\" **\n", res.DUNName2016, res.DUNName2021)
		}

		if len(res.OnlyIn2021) > 0 {
			fmt.Printf("\n  ROWS ONLY IN 2021 (%d):\n", len(res.OnlyIn2021))
			for _, r := range res.OnlyIn2021 {
				fmt.Printf("    [NEW] %s | %s | %s | %s\n", r.Cols[0], r.Cols[2], r.Cols[8], r.Cols[9])
			}
		}

		if len(res.OnlyIn2016) > 0 {
			fmt.Printf("\n  ROWS ONLY IN 2016 (%d):\n", len(res.OnlyIn2016))
			for _, r := range res.OnlyIn2016 {
				fmt.Printf("    [REMOVED] %s | %s | %s | %s\n", r.Cols[0], r.Cols[2], r.Cols[8], r.Cols[9])
			}
		}

		if len(res.Diffs) > 0 {
			fmt.Printf("\n  COLUMN DIFFERENCES (%d):\n", len(res.Diffs))
			// Group by unique code for readability
			diffsByUID := make(map[string][]Diff)
			var uidOrder []string
			for _, d := range res.Diffs {
				if _, exists := diffsByUID[d.UniqueCode]; !exists {
					uidOrder = append(uidOrder, d.UniqueCode)
				}
				diffsByUID[d.UniqueCode] = append(diffsByUID[d.UniqueCode], d)
			}
			for _, uid := range uidOrder {
				diffs := diffsByUID[uid]
				fmt.Printf("    ID: %s\n", uid)
				for _, d := range diffs {
					fmt.Printf("      Col[%d] %s: \"%s\" (2016) -> \"%s\" (2021)\n", d.ColIndex, d.ColName, d.Val2016, d.Val2021)
				}
			}
		}
	}

	// Step 6: Summary table for quick overview
	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("## DIFFERENCE SUMMARY TABLE")
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Printf("%-8s %-25s %6s %6s %7s %8s %8s %6s\n",
		"DUN", "NAME", "R2021", "R2016", "MATCH", "NEW(21)", "REM(16)", "DIFFS")
	fmt.Println(strings.Repeat("-", 80))

	for _, res := range results {
		hasDiffs := len(res.Diffs) > 0 || len(res.OnlyIn2021) > 0 || len(res.OnlyIn2016) > 0 || !res.File2016Found
		if !hasDiffs {
			continue
		}
		name := res.DUNName2021
		if len(name) > 25 {
			name = name[:25]
		}
		file16 := ""
		if !res.File2016Found {
			file16 = "N/A"
		} else {
			file16 = fmt.Sprintf("%d", res.Total2016)
		}
		fmt.Printf("%-8s %-25s %6d %6s %7d %8d %8d %6d\n",
			res.DUNCode, name, res.Total2021, file16, res.MatchedCount,
			len(res.OnlyIn2021), len(res.OnlyIn2016), len(res.Diffs))
	}

	// Step 7: Aggregate differences by column
	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("## DIFFERENCES BY COLUMN")
	fmt.Println("=" + strings.Repeat("=", 79))

	colDiffCounts := make(map[int]int)
	colDiffExamples := make(map[int][]Diff)
	for _, res := range results {
		for _, d := range res.Diffs {
			colDiffCounts[d.ColIndex]++
			if len(colDiffExamples[d.ColIndex]) < 5 {
				colDiffExamples[d.ColIndex] = append(colDiffExamples[d.ColIndex], d)
			}
		}
	}

	for c := 1; c < numCompCols; c++ {
		cnt := colDiffCounts[c]
		if cnt == 0 {
			continue
		}
		fmt.Printf("\nColumn %d (%s): %d differences\n", c, colNames[c], cnt)
		for _, ex := range colDiffExamples[c] {
			fmt.Printf("  Example: %s -> \"%s\" (2016) vs \"%s\" (2021)\n", ex.UniqueCode, ex.Val2016, ex.Val2021)
		}
	}

	// Step 8: DUNs with NO differences
	fmt.Println()
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("## DUNs WITH NO DIFFERENCES (identical columns 1-11)")
	fmt.Println("=" + strings.Repeat("=", 79))
	identicalCount := 0
	for _, res := range results {
		if res.File2016Found && len(res.Diffs) == 0 && len(res.OnlyIn2021) == 0 && len(res.OnlyIn2016) == 0 {
			identicalCount++
		}
	}
	fmt.Printf("Count: %d out of %d DUNs with 2016 data\n", identicalCount, len(results)-dunsNoFile)

	if identicalCount > 0 {
		fmt.Print("DUNs: ")
		first := true
		for _, res := range results {
			if res.File2016Found && len(res.Diffs) == 0 && len(res.OnlyIn2021) == 0 && len(res.OnlyIn2016) == 0 {
				if !first {
					fmt.Print(", ")
				}
				fmt.Printf("%s(%s)", res.DUNCode, res.DUNName2021)
				first = false
			}
		}
		fmt.Println()
	}

	slog.Info("phase-0 comparison complete",
		"totalDUNs", len(results),
		"dunsWithDiffs", dunsWithDiffs,
		"totalMatched", totalMatched,
		"totalDiffs", totalDiffs,
		"totalOnlyIn2021", totalOnlyIn2021,
		"totalOnlyIn2016", totalOnlyIn2016,
	)
}

// readCSV reads a CSV file and returns rows as Row structs (first 11 cols).
// Skips the header row.
func readCSV(path string) ([]Row, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // variable number of fields

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing CSV %s: %w", path, err)
	}

	if len(records) < 2 {
		return nil, nil // empty file or header only
	}

	var rows []Row
	for i := 1; i < len(records); i++ {
		rec := records[i]
		var r Row
		for c := 0; c < numCompCols && c < len(rec); c++ {
			r.Cols[c] = strings.TrimSpace(rec[c])
		}
		// Skip completely empty rows
		if r.Cols[0] == "" {
			continue
		}
		rows = append(rows, r)
	}
	return rows, nil
}
