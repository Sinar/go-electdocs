package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
)

type rowInfo struct {
	lineNum       int // 1-based line number in the CSV file (header=1)
	uniqueCode    string
	pollingCentre string
}

type duplicateGroup struct {
	uniqueCode string
	rows       []rowInfo
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	csvPath := "to-review.csv"
	outputPath := "PHASE-1-REVIEW.md"

	f, err := os.Open(csvPath)
	if err != nil {
		slog.Error("failed to open CSV", "path", csvPath, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // allow variable field counts

	allRecords, err := reader.ReadAll()
	if err != nil {
		slog.Error("failed to read CSV", "error", err)
		os.Exit(1)
	}

	if len(allRecords) < 2 {
		slog.Error("CSV has no data rows")
		os.Exit(1)
	}

	header := allRecords[0]
	slog.Info("CSV loaded", "totalRows", len(allRecords)-1, "columns", len(header))

	// Verify column positions
	if header[0] != "UNIQUE CODE" {
		slog.Warn("Column 1 is not 'UNIQUE CODE'", "found", header[0])
	}
	if len(header) >= 10 && header[9] != "POLLING CENTRE" {
		slog.Warn("Column 10 is not 'POLLING CENTRE'", "found", header[9])
	}

	// Collect all rows by unique code
	codeToRows := make(map[string][]rowInfo)
	totalRows := 0
	emptyCodeCount := 0

	for i := 1; i < len(allRecords); i++ {
		row := allRecords[i]
		lineNum := i + 1 // 1-based, header is line 1
		totalRows++

		uniqueCode := ""
		if len(row) > 0 {
			uniqueCode = strings.TrimSpace(row[0])
		}
		pollingCentre := ""
		if len(row) > 9 {
			pollingCentre = strings.TrimSpace(row[9])
		}

		if uniqueCode == "" {
			emptyCodeCount++
			continue
		}

		codeToRows[uniqueCode] = append(codeToRows[uniqueCode], rowInfo{
			lineNum:       lineNum,
			uniqueCode:    uniqueCode,
			pollingCentre: pollingCentre,
		})
	}

	slog.Info("Scan complete",
		"totalDataRows", totalRows,
		"emptyCodeRows", emptyCodeCount,
		"distinctCodes", len(codeToRows),
	)

	// Find duplicates (codes with more than one row)
	var duplicates []duplicateGroup
	totalDupRows := 0
	for code, rows := range codeToRows {
		if len(rows) > 1 {
			duplicates = append(duplicates, duplicateGroup{
				uniqueCode: code,
				rows:       rows,
			})
			totalDupRows += len(rows)
		}
	}

	// Sort duplicates by unique code for deterministic output
	sort.Slice(duplicates, func(i, j int) bool {
		return duplicates[i].uniqueCode < duplicates[j].uniqueCode
	})

	slog.Info("Duplicates found",
		"duplicateCodeCount", len(duplicates),
		"totalAffectedRows", totalDupRows,
	)

	// Check for existing suffixed IDs (ending in a-z)
	var existingSuffixed []rowInfo
	for code, rows := range codeToRows {
		if len(code) > 0 {
			lastChar := code[len(code)-1]
			if lastChar >= 'a' && lastChar <= 'z' {
				for _, r := range rows {
					existingSuffixed = append(existingSuffixed, r)
				}
			}
		}
	}
	sort.Slice(existingSuffixed, func(i, j int) bool {
		return existingSuffixed[i].lineNum < existingSuffixed[j].lineNum
	})

	slog.Info("Existing suffixed IDs", "count", len(existingSuffixed))

	// For each duplicate group, compute what the suffixes SHOULD be
	type suffixRecommendation struct {
		uniqueCode     string
		occurrences    int
		centreOrder    []string          // polling centres in order of first appearance
		centreToSuffix map[string]string // polling centre -> suffix letter
		rowDetails     []rowInfo
	}

	var recommendations []suffixRecommendation
	for _, dg := range duplicates {
		rec := suffixRecommendation{
			uniqueCode:     dg.uniqueCode,
			occurrences:    len(dg.rows),
			centreToSuffix: make(map[string]string),
		}

		// Determine order of first appearance of each unique polling centre
		seen := make(map[string]bool)
		for _, r := range dg.rows {
			if !seen[r.pollingCentre] {
				seen[r.pollingCentre] = true
				rec.centreOrder = append(rec.centreOrder, r.pollingCentre)
			}
		}

		// Assign suffix letters
		for idx, centre := range rec.centreOrder {
			if idx < 26 {
				rec.centreToSuffix[centre] = string(rune('a' + idx))
			} else {
				rec.centreToSuffix[centre] = fmt.Sprintf("%c%c", 'a'+idx/26-1, 'a'+idx%26)
			}
		}

		rec.rowDetails = dg.rows
		recommendations = append(recommendations, rec)
	}

	// Count unique codes that appear exactly once
	uniqueCount := 0
	for _, rows := range codeToRows {
		if len(rows) == 1 {
			uniqueCount++
		}
	}

	// ---- Generate the Markdown report ----
	var sb strings.Builder

	sb.WriteString("# PHASE 1 REVIEW: UNIQUE CODE Uniqueness Check\n\n")
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("| Metric | Value |\n"))
	sb.WriteString(fmt.Sprintf("| --- | --- |\n"))
	sb.WriteString(fmt.Sprintf("| Total data rows | %d |\n", totalRows))
	sb.WriteString(fmt.Sprintf("| Empty UNIQUE CODE rows | %d |\n", emptyCodeCount))
	sb.WriteString(fmt.Sprintf("| Distinct non-empty UNIQUE CODEs | %d |\n", len(codeToRows)))
	sb.WriteString(fmt.Sprintf("| UNIQUE CODEs appearing exactly once | %d |\n", uniqueCount))
	sb.WriteString(fmt.Sprintf("| **Duplicate UNIQUE CODEs** | **%d** |\n", len(duplicates)))
	sb.WriteString(fmt.Sprintf("| **Total rows affected by duplicates** | **%d** |\n", totalDupRows))
	sb.WriteString(fmt.Sprintf("| Existing suffixed IDs (ending a-z) | %d |\n", len(existingSuffixed)))
	sb.WriteString("\n")

	if emptyCodeCount > 0 {
		sb.WriteString(fmt.Sprintf("> **Note**: %d row(s) with empty UNIQUE CODE were skipped (allowed per rules).\n\n", emptyCodeCount))
	}

	// Verdict
	if len(duplicates) == 0 && len(existingSuffixed) == 0 {
		sb.WriteString("### ✅ Result: PASS\n\n")
		sb.WriteString("All non-empty UNIQUE CODEs are unique. No action needed.\n\n")
	} else {
		sb.WriteString("### ❌ Result: FAIL — Duplicates Found\n\n")
		sb.WriteString(fmt.Sprintf("**%d** distinct UNIQUE CODE values are duplicated across **%d** rows.\n", len(duplicates), totalDupRows))
		sb.WriteString("These need to be disambiguated with letter suffixes (a, b, c, ...) per the suffix rule.\n\n")
	}

	// Existing suffixed IDs
	if len(existingSuffixed) > 0 {
		sb.WriteString("## Existing Suffixed IDs\n\n")
		sb.WriteString("The following IDs already end with a lowercase letter suffix:\n\n")
		sb.WriteString("| Line | UNIQUE CODE | POLLING CENTRE |\n")
		sb.WriteString("| --- | --- | --- |\n")
		for _, r := range existingSuffixed {
			sb.WriteString(fmt.Sprintf("| %d | `%s` | %s |\n", r.lineNum, r.uniqueCode, r.pollingCentre))
		}
		sb.WriteString("\n")
	} else {
		sb.WriteString("## Existing Suffixed IDs\n\n")
		sb.WriteString("No existing suffixed IDs found. All suffix assignments need to be applied fresh.\n\n")
	}

	// Duplicate details
	if len(duplicates) > 0 {
		sb.WriteString("## Duplicate Details\n\n")

		// Overview table: top 30 worst offenders
		sb.WriteString("### Top Duplicates by Occurrence Count\n\n")
		sortedByCount := make([]duplicateGroup, len(duplicates))
		copy(sortedByCount, duplicates)
		sort.Slice(sortedByCount, func(i, j int) bool {
			return len(sortedByCount[i].rows) > len(sortedByCount[j].rows)
		})

		sb.WriteString("| # | UNIQUE CODE | Occurrences | Distinct Polling Centres |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		limit := 30
		if len(sortedByCount) < limit {
			limit = len(sortedByCount)
		}
		for i := 0; i < limit; i++ {
			dg := sortedByCount[i]
			centres := make(map[string]bool)
			for _, r := range dg.rows {
				centres[r.pollingCentre] = true
			}
			sb.WriteString(fmt.Sprintf("| %d | `%s` | %d | %d |\n", i+1, dg.uniqueCode, len(dg.rows), len(centres)))
		}
		if len(sortedByCount) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* | | |\n", len(sortedByCount)-limit))
		}
		sb.WriteString("\n")

		// Distribution of duplicate counts
		sb.WriteString("### Distribution of Duplicate Counts\n\n")
		countDist := make(map[int]int)
		for _, dg := range duplicates {
			countDist[len(dg.rows)]++
		}
		var counts []int
		for c := range countDist {
			counts = append(counts, c)
		}
		sort.Ints(counts)
		sb.WriteString("| Occurrences | Number of UNIQUE CODEs |\n")
		sb.WriteString("| --- | --- |\n")
		for _, c := range counts {
			sb.WriteString(fmt.Sprintf("| %d | %d |\n", c, countDist[c]))
		}
		sb.WriteString("\n")

		// Full detail for each duplicate - grouped by DUN for readability
		sb.WriteString("### Recommended Suffix Assignments\n\n")
		sb.WriteString("Below are all duplicate UNIQUE CODEs with the recommended suffix assignment.\n")
		sb.WriteString("Each unique Polling Centre for a given duplicate ID gets a distinct letter suffix (a, b, c, ...)\n")
		sb.WriteString("based on order of first appearance in the file.\n\n")

		for i, rec := range recommendations {
			sb.WriteString(fmt.Sprintf("#### %d. `%s` (%d occurrences, %d centres)\n\n",
				i+1, rec.uniqueCode, rec.occurrences, len(rec.centreOrder)))
			sb.WriteString("| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |\n")
			sb.WriteString("| --- | --- | --- | --- | --- |\n")
			for _, r := range rec.rowDetails {
				suffix := rec.centreToSuffix[r.pollingCentre]
				newCode := rec.uniqueCode + suffix
				sb.WriteString(fmt.Sprintf("| %d | `%s` | %s | %s | `%s` |\n",
					r.lineNum, r.uniqueCode, r.pollingCentre, suffix, newCode))
			}
			sb.WriteString("\n")
		}
	}

	// Recommendations section
	sb.WriteString("## Recommendations\n\n")
	if len(duplicates) > 0 {
		sb.WriteString(fmt.Sprintf("1. **Apply suffix disambiguation**: All %d duplicate UNIQUE CODEs must be suffixed ", len(duplicates)))
		sb.WriteString("with letters (a, b, c, ...) based on the order of first appearance of each distinct Polling Centre.\n")
		sb.WriteString("2. **Only modify column 1 (UNIQUE CODE)**: No other columns should change.\n")
		sb.WriteString("3. **Pattern explanation**: The duplicates arise because multiple Polling Centres share the same ")
		sb.WriteString("Polling District Code and Voting Channel Number. The current ID scheme (ParliCode_DUNCode_DMKOD_Channel) ")
		sb.WriteString("does not distinguish between different Polling Centres within the same polling district/channel.\n")
		sb.WriteString("4. **Scope**: This affects all 82 DUNs — it is a systematic issue, not limited to specific constituencies.\n")
	} else {
		sb.WriteString("No action needed. All UNIQUE CODEs are already unique.\n")
	}
	sb.WriteString("\n")

	// Write report
	err = os.WriteFile(outputPath, []byte(sb.String()), 0644)
	if err != nil {
		slog.Error("failed to write report", "path", outputPath, "error", err)
		os.Exit(1)
	}

	slog.Info("Report written", "path", outputPath)

	// Print summary to stdout
	fmt.Println("=== PHASE 1: UNIQUE CODE Uniqueness Check ===")
	fmt.Printf("Total data rows:          %d\n", totalRows)
	fmt.Printf("Empty UNIQUE CODE rows:   %d\n", emptyCodeCount)
	fmt.Printf("Distinct UNIQUE CODEs:    %d\n", len(codeToRows))
	fmt.Printf("Unique (appear once):     %d\n", uniqueCount)
	fmt.Printf("Duplicate UNIQUE CODEs:   %d\n", len(duplicates))
	fmt.Printf("Rows affected:            %d\n", totalDupRows)
	fmt.Printf("Existing suffixed IDs:    %d\n", len(existingSuffixed))
	if len(duplicates) > 0 {
		fmt.Println("\n❌ FAIL: Duplicates found. See PHASE-1-REVIEW.md for details.")
	} else {
		fmt.Println("\n✅ PASS: All UNIQUE CODEs are unique.")
	}
}
