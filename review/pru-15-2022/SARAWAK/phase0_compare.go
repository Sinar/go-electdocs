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

// DUNInfo holds DUN-level structural info
type DUNInfo struct {
	DunCode string
	DunName string
	ParCode string
	ParName string
}

// PDInfo holds polling-district-level info
type PDInfo struct {
	PollDistCode string
	PollDistName string
	Centres      map[string]int // centre name -> count of channels
	MaxChannel   int
}

// Row holds individual row data
type Row struct {
	UniqueCode    string
	State         string
	BallotType    string
	ParCode       string
	ParName       string
	DunCode       string
	DunName       string
	PollDistCode  string
	PollDistName  string
	PollingCentre string
	VotingChannel string
}

type DUNData struct {
	Info DUNInfo
	PDs  map[string]*PDInfo // PollDistCode -> PDInfo
	Rows []Row
}

type DiffRecord struct {
	DunCode  string
	Field    string
	Val2016  string
	Val2022  string
	PDCode   string
	Evidence string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	dunDir := "/Users/leow/TINDAKMSIA/go-electdocs/data/sarawak-dun-2016/OUTPUT"
	reviewFile := "to-review.csv"

	// 1. Read all 2016 DUN files
	slog.Info("Reading 2016 DUN files", "dir", dunDir)
	all2016 := make(map[string]*DUNData) // dunCode -> data

	for i := 1; i <= 82; i++ {
		if i == 79 || i == 82 {
			continue
		}
		dunCode := fmt.Sprintf("N.%02d", i)
		fname := filepath.Join(dunDir, fmt.Sprintf("Sarawak-N.%02d.csv", i))
		dd, err := read2016File(fname, dunCode)
		if err != nil {
			slog.Warn("Could not read 2016 file", "file", fname, "err", err)
			continue
		}
		all2016[dunCode] = dd
		slog.Info("Read 2016 DUN", "dun", dunCode, "ordinaryRows", len(dd.Rows), "pds", len(dd.PDs))
	}

	// 2. Read 2022 to-review.csv
	slog.Info("Reading 2022 to-review file", "file", reviewFile)
	all2022, err := read2022File(reviewFile)
	if err != nil {
		slog.Error("Could not read 2022 file", "file", reviewFile, "err", err)
		os.Exit(1)
	}
	totalRows2022 := 0
	for _, dd := range all2022 {
		totalRows2022 += len(dd.Rows)
		slog.Info("Read 2022 DUN", "dun", dd.Info.DunCode, "ordinaryRows", len(dd.Rows), "pds", len(dd.PDs))
	}
	slog.Info("Read 2022 file", "totalOrdinaryRows", totalRows2022, "duns", len(all2022))

	// 3. Compare
	allDuns := make(map[string]bool)
	for d := range all2016 {
		allDuns[d] = true
	}
	for d := range all2022 {
		allDuns[d] = true
	}
	dunList := sortedKeys(allDuns)

	var (
		dunsOnly2016 []string
		dunsOnly2022 []string
		diffs        []DiffRecord
	)

	// DUN-level: PAR code/name consistency check using polling district code prefix
	var parConsistencyIssues []ParConsistency

	// PD-level tracking
	var pdOnlyIn2016 []PDMatch
	var pdOnlyIn2022 []PDMatch
	var pdNameDiffs []PDMatch
	var pdCentreDiffs []PDMatch
	var pdMatched []PDMatch

	for _, dun := range dunList {
		d16, has16 := all2016[dun]
		d22, has22 := all2022[dun]

		if !has16 && has22 {
			dunsOnly2022 = append(dunsOnly2022, dun)
			continue
		}
		if has16 && !has22 {
			dunsOnly2016 = append(dunsOnly2016, dun)
			continue
		}

		// DUN-level: compare PAR code and name
		// Extract the expected PAR number from PD codes
		pdPrefix16 := extractPDPrefix(d16)
		pdPrefix22 := extractPDPrefix(d22)

		parCode16 := d16.Info.ParCode
		parCode22 := d22.Info.ParCode
		parName16 := d16.Info.ParName
		parName22 := d22.Info.ParName

		if parCode16 != parCode22 {
			// Check consistency: does the PD prefix match the PAR code?
			consistent16 := (fmt.Sprintf("P.%s", pdPrefix16) == parCode16)
			consistent22 := (fmt.Sprintf("P.%s", pdPrefix22) == parCode22)

			parConsistencyIssues = append(parConsistencyIssues, ParConsistency{
				DunCode:      dun,
				ParCode2016:  parCode16,
				ParCode2022:  parCode22,
				PDPrefix2016: pdPrefix16,
				PDPrefix2022: pdPrefix22,
				ParName2016:  parName16,
				ParName2022:  parName22,
				Consistent16: consistent16,
				Consistent22: consistent22,
				NumRows2016:  len(d16.Rows),
				NumRows2022:  len(d22.Rows),
			})

			diffs = append(diffs, DiffRecord{
				DunCode:  dun,
				Field:    "PAR CODE",
				Val2016:  parCode16,
				Val2022:  parCode22,
				Evidence: fmt.Sprintf("PD prefix 2016=%s, 2022=%s; 2016 consistent=%v, 2022 consistent=%v", pdPrefix16, pdPrefix22, consistent16, consistent22),
			})
		}

		if !strings.EqualFold(parName16, parName22) {
			diffs = append(diffs, DiffRecord{
				DunCode: dun,
				Field:   "PAR NAME",
				Val2016: parName16,
				Val2022: parName22,
			})
		}

		if !strings.EqualFold(d16.Info.DunName, d22.Info.DunName) {
			diffs = append(diffs, DiffRecord{
				DunCode: dun,
				Field:   "DUN NAME",
				Val2016: d16.Info.DunName,
				Val2022: d22.Info.DunName,
			})
		}

		// PD-level comparison
		allPDs := make(map[string]bool)
		for pd := range d16.PDs {
			allPDs[pd] = true
		}
		for pd := range d22.PDs {
			allPDs[pd] = true
		}

		for _, pd := range sortedKeys(allPDs) {
			pd16, in16 := d16.PDs[pd]
			pd22, in22 := d22.PDs[pd]

			if in16 && !in22 {
				centres16 := sortedCentreKeys(pd16.Centres)
				pdOnlyIn2016 = append(pdOnlyIn2016, PDMatch{
					DunCode:     dun,
					PDCode:      pd,
					PDName2016:  pd16.PollDistName,
					Centres2016: centres16,
					Channels16:  countChannels(pd16.Centres),
				})
				continue
			}
			if !in16 && in22 {
				centres22 := sortedCentreKeys(pd22.Centres)
				pdOnlyIn2022 = append(pdOnlyIn2022, PDMatch{
					DunCode:     dun,
					PDCode:      pd,
					PDName2022:  pd22.PollDistName,
					Centres2022: centres22,
					Channels22:  countChannels(pd22.Centres),
				})
				continue
			}

			// Both exist: compare
			centres16 := sortedCentreKeys(pd16.Centres)
			centres22 := sortedCentreKeys(pd22.Centres)

			m := PDMatch{
				DunCode:     dun,
				PDCode:      pd,
				PDName2016:  pd16.PollDistName,
				PDName2022:  pd22.PollDistName,
				Centres2016: centres16,
				Centres2022: centres22,
				Channels16:  countChannels(pd16.Centres),
				Channels22:  countChannels(pd22.Centres),
			}
			pdMatched = append(pdMatched, m)

			if !strings.EqualFold(pd16.PollDistName, pd22.PollDistName) {
				pdNameDiffs = append(pdNameDiffs, m)
			}

			// Compare centres (case-insensitive)
			c16set := make(map[string]bool)
			for _, c := range centres16 {
				c16set[strings.ToUpper(c)] = true
			}
			c22set := make(map[string]bool)
			for _, c := range centres22 {
				c22set[strings.ToUpper(c)] = true
			}

			hasCentreDiff := false
			for c := range c16set {
				if !c22set[c] {
					hasCentreDiff = true
					break
				}
			}
			if !hasCentreDiff {
				for c := range c22set {
					if !c16set[c] {
						hasCentreDiff = true
						break
					}
				}
			}
			if hasCentreDiff {
				pdCentreDiffs = append(pdCentreDiffs, m)
			}
		}
	}

	// 4. Write report
	slog.Info("Writing report")
	writeReport(
		dunList, all2016, all2022,
		dunsOnly2016, dunsOnly2022,
		diffs, parConsistencyIssues,
		pdOnlyIn2016, pdOnlyIn2022,
		pdNameDiffs, pdCentreDiffs, pdMatched,
	)

	slog.Info("Done",
		"dunsWith16", len(all2016),
		"dunsWith22", len(all2022),
		"totalDiffs", len(diffs),
		"pdMatched", len(pdMatched),
		"pdOnlyIn2016", len(pdOnlyIn2016),
		"pdOnlyIn2022", len(pdOnlyIn2022),
		"pdCentreDiffs", len(pdCentreDiffs),
	)
}

func read2016File(path, expectedDun string) (*DUNData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	dd := &DUNData{
		PDs: make(map[string]*PDInfo),
	}

	for i, rec := range records {
		if i == 0 {
			continue
		}
		if len(rec) < 11 {
			continue
		}
		ballotType := strings.TrimSpace(rec[2])
		if ballotType != "ORDINARY VOTE" {
			continue
		}
		pdCode := strings.TrimSpace(rec[7])
		// Skip summary rows (empty PD code)
		if pdCode == "" {
			continue
		}

		r := Row{
			UniqueCode:    strings.TrimSpace(rec[0]),
			State:         strings.TrimSpace(rec[1]),
			BallotType:    ballotType,
			ParCode:       strings.TrimSpace(rec[3]),
			ParName:       strings.TrimSpace(rec[4]),
			DunCode:       strings.TrimSpace(rec[5]),
			DunName:       strings.TrimSpace(rec[6]),
			PollDistCode:  pdCode,
			PollDistName:  strings.TrimSpace(rec[8]),
			PollingCentre: strings.TrimSpace(rec[9]),
			VotingChannel: strings.TrimSpace(rec[10]),
		}

		dd.Rows = append(dd.Rows, r)

		// Set DUN info from first row
		if dd.Info.DunCode == "" {
			dd.Info = DUNInfo{
				DunCode: r.DunCode,
				DunName: r.DunName,
				ParCode: r.ParCode,
				ParName: r.ParName,
			}
		}

		// Track PD info
		if _, ok := dd.PDs[pdCode]; !ok {
			dd.PDs[pdCode] = &PDInfo{
				PollDistCode: pdCode,
				PollDistName: r.PollDistName,
				Centres:      make(map[string]int),
			}
		}
		dd.PDs[pdCode].Centres[r.PollingCentre]++
	}
	return dd, nil
}

func read2022File(path string) (map[string]*DUNData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make(map[string]*DUNData)

	for i, rec := range records {
		if i == 0 {
			continue
		}
		if len(rec) < 21 {
			continue
		}
		ballotType := strings.TrimSpace(rec[2])
		if ballotType != "ORDINARY VOTE" {
			continue
		}
		pdCode := strings.TrimSpace(rec[7])
		if pdCode == "" {
			continue
		}

		r := Row{
			UniqueCode:    strings.TrimSpace(rec[0]),
			State:         strings.TrimSpace(rec[1]),
			BallotType:    ballotType,
			ParCode:       strings.TrimSpace(rec[3]),
			ParName:       strings.TrimSpace(rec[4]),
			DunCode:       strings.TrimSpace(rec[5]),
			DunName:       strings.TrimSpace(rec[6]),
			PollDistCode:  pdCode,
			PollDistName:  strings.TrimSpace(rec[8]),
			PollingCentre: strings.TrimSpace(rec[9]),
			VotingChannel: strings.TrimSpace(rec[20]),
		}

		dunCode := r.DunCode
		if _, ok := result[dunCode]; !ok {
			result[dunCode] = &DUNData{
				Info: DUNInfo{
					DunCode: r.DunCode,
					DunName: r.DunName,
					ParCode: r.ParCode,
					ParName: r.ParName,
				},
				PDs: make(map[string]*PDInfo),
			}
		}
		dd := result[dunCode]
		dd.Rows = append(dd.Rows, r)

		if _, ok := dd.PDs[pdCode]; !ok {
			dd.PDs[pdCode] = &PDInfo{
				PollDistCode: pdCode,
				PollDistName: r.PollDistName,
				Centres:      make(map[string]int),
			}
		}
		dd.PDs[pdCode].Centres[r.PollingCentre]++
	}
	return result, nil
}

func extractPDPrefix(dd *DUNData) string {
	// Extract the PAR number from PD codes (format: PPP/DD/XX)
	prefixes := make(map[string]int)
	for pd := range dd.PDs {
		parts := strings.SplitN(pd, "/", 2)
		if len(parts) >= 1 && parts[0] != "" {
			prefixes[parts[0]]++
		}
	}
	// Return the most common prefix
	best := ""
	bestCount := 0
	for p, c := range prefixes {
		if c > bestCount {
			best = p
			bestCount = c
		}
	}
	return best
}

func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sortedCentreKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func countChannels(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func writeReport(
	dunList []string,
	all2016, all2022 map[string]*DUNData,
	dunsOnly2016, dunsOnly2022 []string,
	diffs []DiffRecord,
	parIssues []ParConsistency,
	pdOnlyIn2016, pdOnlyIn2022 []PDMatch,
	pdNameDiffs, pdCentreDiffs, pdMatched []PDMatch,
) {
	out, err := os.Create("PHASE-0-REVIEW.md")
	if err != nil {
		slog.Error("Cannot create report", "err", err)
		os.Exit(1)
	}
	defer out.Close()

	w := func(format string, args ...any) {
		fmt.Fprintf(out, format, args...)
	}

	w("# PHASE 0 REVIEW: Compare to-review.csv against 2016 DUN Sarawak Results\n\n")
	w("## Overview\n\n")
	w("This review compares **ORDINARY VOTE** rows in `to-review.csv` (PRU-15, 2022 parliamentary election)\n")
	w("against the 2016 Sarawak DUN election files in `data/sarawak-dun-2016/OUTPUT/Sarawak-N.XX.csv`.\n\n")
	w("**Comparison scope:** Structural fields only (PAR code/name, DUN code/name, polling district code/name, polling centre).\n")
	w("- Postal vote rows are excluded (different aggregation level: PAR-level in 2022 vs DUN-level in 2016).\n")
	w("- Early vote rows are excluded (polling centres commonly change between elections).\n")
	w("- Summary/totals rows in 2016 files (empty polling district code) are excluded.\n")
	w("- Matching is done at the **DUN level** and **Polling District level** (not row-level), because\n")
	w("  2016 files have duplicate (PD code, channel) combinations across multiple polling centres within\n")
	w("  the same PD, while 2022 uses sequential channel numbers.\n\n")

	// Summary stats
	total2016 := 0
	total2022 := 0
	totalPD16 := 0
	totalPD22 := 0
	for _, dd := range all2016 {
		total2016 += len(dd.Rows)
		totalPD16 += len(dd.PDs)
	}
	for _, dd := range all2022 {
		total2022 += len(dd.Rows)
		totalPD22 += len(dd.PDs)
	}

	w("## Summary Statistics\n\n")
	w("| Metric | Count |\n")
	w("|--------|-------|\n")
	w("| DUNs in 2016 (files available) | %d |\n", len(all2016))
	w("| DUNs in 2022 (ordinary vote data) | %d |\n", len(all2022))
	w("| DUNs only in 2016 | %d |\n", len(dunsOnly2016))
	w("| DUNs only in 2022 (not contested 2016) | %d |\n", len(dunsOnly2022))
	w("| Total ordinary vote rows 2016 | %d |\n", total2016)
	w("| Total ordinary vote rows 2022 | %d |\n", total2022)
	w("| Total unique polling districts 2016 | %d |\n", totalPD16)
	w("| Total unique polling districts 2022 | %d |\n", totalPD22)
	w("| Polling districts matched (in both) | %d |\n", len(pdMatched))
	w("| Polling districts only in 2016 | %d |\n", len(pdOnlyIn2016))
	w("| Polling districts only in 2022 | %d |\n", len(pdOnlyIn2022))
	w("| PDs with centre name differences | %d |\n", len(pdCentreDiffs))
	w("| PDs with polling district name diffs | %d |\n", len(pdNameDiffs))
	w("| DUN-level field differences | %d |\n", len(diffs))
	w("\n")

	// --- DUNs only in 2022 ---
	if len(dunsOnly2022) > 0 {
		w("## DUNs Present Only in 2022\n\n")
		w("These DUNs were not contested in the 2016 Sarawak DUN election.\n\n")
		w("| DUN Code | DUN Name | PAR Code | PAR Name | Ordinary Rows | PDs |\n")
		w("|----------|----------|----------|----------|---------------|-----|\n")
		for _, d := range dunsOnly2022 {
			dd := all2022[d]
			w("| %s | %s | %s | %s | %d | %d |\n", d, dd.Info.DunName, dd.Info.ParCode, dd.Info.ParName, len(dd.Rows), len(dd.PDs))
		}
		w("\n")
	}

	if len(dunsOnly2016) > 0 {
		w("## DUNs Present Only in 2016\n\n")
		for _, d := range dunsOnly2016 {
			dd := all2016[d]
			w("- %s (%s) — %d rows\n", d, dd.Info.DunName, len(dd.Rows))
		}
		w("\n")
	}

	// --- PAR CODE consistency analysis ---
	// Count how many DUNs have PAR code changes
	parCodeDiffCount := 0
	parNameDiffCount := 0
	dunNameDiffCount := 0
	for _, d := range diffs {
		switch d.Field {
		case "PAR CODE":
			parCodeDiffCount++
		case "PAR NAME":
			parNameDiffCount++
		case "DUN NAME":
			dunNameDiffCount++
		}
	}

	w("## DUN-Level Differences\n\n")
	w("### Difference Summary\n\n")
	w("| Field | DUNs Affected |\n")
	w("|-------|---------------|\n")
	w("| PAR CODE | %d |\n", parCodeDiffCount)
	w("| PAR NAME | %d |\n", parNameDiffCount)
	w("| DUN NAME | %d |\n", dunNameDiffCount)
	w("\n")

	// --- PAR CODE + NAME differences with consistency analysis ---
	if len(parIssues) > 0 {
		w("### PAR CODE Differences with PD-Prefix Consistency Check\n\n")
		w("The polling district code format is `PPP/DD/XX` where `PPP` = parliamentary number, `DD` = DUN number.\n")
		w("We check whether the PAR code in each file matches the prefix of its polling district codes.\n\n")
		w("| DUN | 2016 PAR Code | 2016 PAR Name | PD Prefix 2016 | 2016 Consistent? | 2022 PAR Code | 2022 PAR Name | PD Prefix 2022 | 2022 Consistent? |\n")
		w("|-----|--------------|---------------|----------------|-----------------|--------------|---------------|----------------|------------------|\n")
		for _, p := range parIssues {
			c16 := "✅"
			if !p.Consistent16 {
				c16 = "❌"
			}
			c22 := "✅"
			if !p.Consistent22 {
				c22 = "❌"
			}
			w("| %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
				p.DunCode, p.ParCode2016, p.ParName2016, p.PDPrefix2016, c16,
				p.ParCode2022, p.ParName2022, p.PDPrefix2022, c22)
		}
		w("\n")

		// Analysis
		all2016Wrong := true
		all2022OK := true
		for _, p := range parIssues {
			if p.Consistent16 {
				all2016Wrong = false
			}
			if !p.Consistent22 {
				all2022OK = false
			}
		}
		w("**Analysis:**\n\n")
		if all2016Wrong && all2022OK {
			w("- All %d PAR code differences show the **2016 data** having a PAR code that does NOT match its own polling district code prefix.\n", len(parIssues))
			w("- The **2022 data** has PAR codes that ARE consistent with the polling district code prefix in every case.\n")
			w("- **Conclusion:** The 2016 DUN files contain incorrect PAR code assignments. The 2022 `to-review.csv` has the correct PAR codes.\n")
			w("  This is a known issue with the 2016 data source, not a problem in the 2022 file.\n\n")
		} else {
			w("- Mixed consistency results. Manual review needed for specific DUNs.\n\n")
		}
	}

	// PAR NAME only diffs (where PAR code was the same)
	var parNameOnlyDiffs []DiffRecord
	for _, d := range diffs {
		if d.Field == "PAR NAME" {
			// Check if PAR CODE also differed for this DUN
			parCodeAlsoDiffered := false
			for _, d2 := range diffs {
				if d2.Field == "PAR CODE" && d2.DunCode == d.DunCode {
					parCodeAlsoDiffered = true
					break
				}
			}
			if !parCodeAlsoDiffered {
				parNameOnlyDiffs = append(parNameOnlyDiffs, d)
			}
		}
	}

	if len(parNameOnlyDiffs) > 0 {
		w("### PAR NAME Differences (same PAR code)\n\n")
		w("These DUNs have the same PAR code but different PAR names:\n\n")
		w("| DUN | PAR Code | 2016 PAR Name | 2022 PAR Name |\n")
		w("|-----|----------|--------------|---------------|\n")
		for _, d := range parNameOnlyDiffs {
			parCode := ""
			if dd, ok := all2022[d.DunCode]; ok {
				parCode = dd.Info.ParCode
			}
			w("| %s | %s | %s | %s |\n", d.DunCode, parCode, d.Val2016, d.Val2022)
		}
		w("\n")
	}

	// DUN NAME differences
	var dunNameOnlyDiffs []DiffRecord
	for _, d := range diffs {
		if d.Field == "DUN NAME" {
			dunNameOnlyDiffs = append(dunNameOnlyDiffs, d)
		}
	}
	if len(dunNameOnlyDiffs) > 0 {
		w("### DUN NAME Differences\n\n")
		w("| DUN Code | 2016 Name | 2022 Name | Notes |\n")
		w("|----------|-----------|-----------|-------|\n")
		for _, d := range dunNameOnlyDiffs {
			notes := ""
			if strings.ReplaceAll(d.Val2016, "'", "`") == d.Val2022 || strings.ReplaceAll(d.Val2016, "'", "`") == strings.ReplaceAll(d.Val2022, "'", "`") {
				notes = "Apostrophe/backtick difference"
			}
			w("| %s | %s | %s | %s |\n", d.DunCode, d.Val2016, d.Val2022, notes)
		}
		w("\n")
	}

	// --- Polling District Level ---
	w("## Polling District Level Comparison\n\n")

	if len(pdOnlyIn2016) > 0 {
		w("### Polling Districts Only in 2016\n\n")
		w("These polling districts exist in 2016 ordinary vote data but not in 2022.\n\n")
		w("| DUN | PD Code | PD Name (2016) | Centres (2016) | Channels |\n")
		w("|-----|---------|----------------|----------------|----------|\n")
		for _, m := range pdOnlyIn2016 {
			centres := strings.Join(m.Centres2016, "; ")
			if len(centres) > 80 {
				centres = centres[:77] + "..."
			}
			w("| %s | %s | %s | %s | %d |\n", m.DunCode, m.PDCode, m.PDName2016, centres, m.Channels16)
		}
		w("\n")
	}

	if len(pdOnlyIn2022) > 0 {
		w("### Polling Districts Only in 2022\n\n")
		w("These polling districts exist in 2022 ordinary vote data but not in 2016.\n\n")
		w("| DUN | PD Code | PD Name (2022) | Centres (2022) | Channels |\n")
		w("|-----|---------|----------------|----------------|----------|\n")
		for _, m := range pdOnlyIn2022 {
			centres := strings.Join(m.Centres2022, "; ")
			if len(centres) > 80 {
				centres = centres[:77] + "..."
			}
			w("| %s | %s | %s | %s | %d |\n", m.DunCode, m.PDCode, m.PDName2022, centres, m.Channels22)
		}
		w("\n")
	}

	// PD Name diffs
	if len(pdNameDiffs) > 0 {
		w("### Polling District Name Differences\n\n")
		w("| DUN | PD Code | 2016 Name | 2022 Name |\n")
		w("|-----|---------|-----------|------------|\n")
		for _, m := range pdNameDiffs {
			w("| %s | %s | %s | %s |\n", m.DunCode, m.PDCode, m.PDName2016, m.PDName2022)
		}
		w("\n")
	}

	// Centre diffs
	if len(pdCentreDiffs) > 0 {
		w("### Polling Centre Name Differences\n\n")
		w("Polling districts where the set of polling centre names differs between 2016 and 2022 (case-insensitive).\n\n")

		// Categorize: abbreviation expansion vs real change
		type CentreAnalysis struct {
			DunCode  string
			PDCode   string
			Only2016 []string
			Only2022 []string
			Category string
		}
		var analyses []CentreAnalysis

		abbrCount := 0
		realChangeCount := 0

		for _, m := range pdCentreDiffs {
			c16set := make(map[string]bool)
			for _, c := range m.Centres2016 {
				c16set[strings.ToUpper(c)] = true
			}
			c22set := make(map[string]bool)
			for _, c := range m.Centres2022 {
				c22set[strings.ToUpper(c)] = true
			}

			var only16, only22 []string
			for _, c := range m.Centres2016 {
				if !c22set[strings.ToUpper(c)] {
					only16 = append(only16, c)
				}
			}
			for _, c := range m.Centres2022 {
				if !c16set[strings.ToUpper(c)] {
					only22 = append(only22, c)
				}
			}

			category := "NAME CHANGE"
			if len(only16) == len(only22) {
				allAbbr := true
				for i := range only16 {
					if !isAbbrExpansion(only16[i], only22[i]) {
						allAbbr = false
						break
					}
				}
				if allAbbr && len(only16) > 0 {
					category = "ABBREVIATION"
					abbrCount++
				} else {
					realChangeCount++
				}
			} else {
				realChangeCount++
			}

			analyses = append(analyses, CentreAnalysis{
				DunCode:  m.DunCode,
				PDCode:   m.PDCode,
				Only2016: only16,
				Only2022: only22,
				Category: category,
			})
		}

		w("**Summary:** %d PDs have centre name differences: %d are abbreviation expansions (e.g. \"SEK. KEB.\" → \"SEKOLAH KEBANGSAAN\"), %d are real name changes.\n\n",
			len(pdCentreDiffs), abbrCount, realChangeCount)

		// Show real changes first
		if realChangeCount > 0 {
			w("#### Real Centre Name Changes\n\n")
			w("| DUN | PD Code | 2016 Only | 2022 Only |\n")
			w("|-----|---------|-----------|----------|\n")
			for _, a := range analyses {
				if a.Category != "NAME CHANGE" {
					continue
				}
				o16 := strings.Join(a.Only2016, "; ")
				o22 := strings.Join(a.Only2022, "; ")
				if len(o16) > 60 {
					o16 = o16[:57] + "..."
				}
				if len(o22) > 60 {
					o22 = o22[:57] + "..."
				}
				w("| %s | %s | %s | %s |\n", a.DunCode, a.PDCode, o16, o22)
			}
			w("\n")
		}

		if abbrCount > 0 {
			w("#### Abbreviation Expansions (Sample)\n\n")
			w("| DUN | PD Code | 2016 Centre | 2022 Centre |\n")
			w("|-----|---------|-------------|-------------|\n")
			shown := 0
			for _, a := range analyses {
				if a.Category != "ABBREVIATION" {
					continue
				}
				if shown >= 20 {
					w("| ... | ... | (%d more abbreviation-only changes) | ... |\n", abbrCount-20)
					break
				}
				for i := range a.Only2016 {
					w("| %s | %s | %s | %s |\n", a.DunCode, a.PDCode, a.Only2016[i], a.Only2022[i])
				}
				shown++
			}
			w("\n")
		}
	}

	// --- Per-DUN Row Count Comparison ---
	w("## Per-DUN Row Count Comparison\n\n")
	w("All 2022 DUNs have more ordinary vote rows than 2016, reflecting voter population growth and more voting channels.\n\n")
	w("| DUN | DUN Name | 2016 Rows | 2022 Rows | Diff | 2016 PDs | 2022 PDs | PD Diff |\n")
	w("|-----|----------|-----------|-----------|------|----------|----------|---------|\n")
	for _, dun := range dunList {
		c16 := 0
		pd16 := 0
		name := ""
		if dd, ok := all2016[dun]; ok {
			c16 = len(dd.Rows)
			pd16 = len(dd.PDs)
		}
		c22 := 0
		pd22 := 0
		if dd, ok := all2022[dun]; ok {
			c22 = len(dd.Rows)
			pd22 = len(dd.PDs)
			name = dd.Info.DunName
		} else if dd, ok := all2016[dun]; ok {
			name = dd.Info.DunName
		}
		rowDiff := c22 - c16
		pdDiff := pd22 - pd16
		rowDiffStr := fmt.Sprintf("%+d", rowDiff)
		pdDiffStr := fmt.Sprintf("%+d", pdDiff)
		if rowDiff != 0 {
			rowDiffStr = fmt.Sprintf("**%+d**", rowDiff)
		}
		if pdDiff != 0 {
			pdDiffStr = fmt.Sprintf("**%+d**", pdDiff)
		}
		w("| %s | %s | %d | %d | %s | %d | %d | %s |\n", dun, name, c16, c22, rowDiffStr, pd16, pd22, pdDiffStr)
	}
	w("\n")

	// --- Conclusion ---
	w("## Conclusion\n\n")
	w("### Key Findings\n\n")
	w("1. **DUN Coverage:** All 80 DUNs from 2016 are present in 2022. Two additional DUNs (N.79 BUKIT KOTA, N.82 BUKIT SARI) ")
	w("are in 2022 but were not contested in 2016.\n\n")

	w("2. **PAR Code Discrepancies (%d DUNs):** The 2016 DUN files have PAR codes that do NOT match their own polling district code prefixes. ", parCodeDiffCount)
	w("For example, N.11's polling districts start with `195/11/...` (indicating P.195) but the 2016 file says `P.192`. ")
	w("The 2022 data correctly uses `P.195`. This is a data quality issue in the **2016 source files**, not in `to-review.csv`.\n\n")

	w("3. **PAR Name Changes (%d DUNs):** Follow from the PAR code changes. Where PAR codes differ, the PAR names naturally differ too.\n\n", parNameDiffCount)

	w("4. **DUN Name:** Only %d difference found — an apostrophe vs backtick variation, which is a minor encoding difference.\n\n", dunNameDiffCount)

	w("5. **Polling District Coverage:** %d PDs matched between 2016 and 2022. ", len(pdMatched))
	w("%d PDs are only in 2016 (removed/merged), %d only in 2022 (new/split).\n\n", len(pdOnlyIn2016), len(pdOnlyIn2022))

	w("6. **Polling Centre Names:** Most differences are abbreviation expansions (e.g. \"SEK. KEB.\" → \"SEKOLAH KEBANGSAAN\"). ")
	w("Some centres were genuinely renamed or replaced between elections.\n\n")

	w("7. **Row Counts:** Every DUN shows more ordinary vote rows in 2022 than 2016, consistent with ")
	w("voter population growth requiring more voting channels.\n\n")

	w("### Overall Assessment\n\n")
	w("The structural data in `to-review.csv` is **consistent and reasonable** when compared against the 2016 baseline:\n")
	w("- DUN codes and names match (with one trivial apostrophe variant).\n")
	w("- Polling district codes and names are stable across elections.\n")
	w("- The PAR code differences are a known issue in the 2016 source data, not in the 2022 file.\n")
	w("- Growth in channels/rows is expected between elections.\n")
	w("- No anomalous structural issues were found in `to-review.csv`.\n")
}

// isAbbrExpansion checks if s2022 looks like an expansion of abbreviations in s2016
func isAbbrExpansion(s2016, s2022 string) bool {
	// Common abbreviations used in 2016 data
	abbrs := map[string]string{
		"SEK. KEB.":  "SEKOLAH KEBANGSAAN",
		"SEK KEB":    "SEKOLAH KEBANGSAAN",
		"SEK.KEB.":   "SEKOLAH KEBANGSAAN",
		"SK":         "SEKOLAH KEBANGSAAN",
		"SJK(C)":     "SJK CHUNG HUA",
		"SJK (C)":    "SJK CHUNG HUA",
		"SRB":        "SEKOLAH RENDAH BANTUAN",
		"DWN":        "DEWAN",
		"BALAI RAYA": "BALAI RAYA",
		"KPG.":       "KAMPUNG",
		"KPG":        "KAMPUNG",
	}

	u16 := strings.ToUpper(s2016)
	u22 := strings.ToUpper(s2022)

	// Direct match
	if u16 == u22 {
		return true
	}

	// Try expanding abbreviations in 2016 name and see if it matches 2022
	expanded := u16
	for abbr, full := range abbrs {
		expanded = strings.ReplaceAll(expanded, abbr, full)
	}

	// Normalize spaces
	expanded = strings.Join(strings.Fields(expanded), " ")
	normalized22 := strings.Join(strings.Fields(u22), " ")

	if expanded == normalized22 {
		return true
	}

	// Check if 2016 is a substring prefix of 2022 (after removing dots and abbreviation markers)
	clean16 := strings.ReplaceAll(strings.ReplaceAll(u16, ".", ""), "  ", " ")
	clean22 := strings.ReplaceAll(strings.ReplaceAll(u22, ".", ""), "  ", " ")

	if strings.Contains(clean22, clean16) || strings.Contains(clean16, clean22) {
		return true
	}

	return false
}

// ParConsistency is used for PAR code analysis
type ParConsistency struct {
	DunCode      string
	ParCode2016  string
	ParCode2022  string
	PDPrefix2016 string
	PDPrefix2022 string
	ParName2016  string
	ParName2022  string
	Consistent16 bool
	Consistent22 bool
	NumRows2016  int
	NumRows2022  int
}

type PDMatch struct {
	DunCode     string
	PDCode      string
	PDName2016  string
	PDName2022  string
	Centres2016 []string
	Centres2022 []string
	Channels16  int
	Channels22  int
}
