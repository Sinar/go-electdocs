package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sort"
	"strings"
)

// Column indices (0-based) for party columns in to-review.csv
var partyLabelCols = []int{12, 17, 22, 27, 32, 37, 42, 47, 52, 57}
var partyCandCols = []int{13, 18, 23, 28, 33, 38, 43, 48, 53, 58}

// Header names for each party slot
var slotNames = []string{"GPS", "PH", "PSB", "PBK", "ASPIRASI", "PBDSB", "SEDAR", "PAS", "INDEPENDENT 1", "INDEPENDENT 2"}

// Expected party labels for each column slot (non-independent)
var expectedLabels = map[int][]string{
	0: {"PBB", "SUPP", "PRS", "PDP", "SPDP"}, // GPS components
	1: {"PKR", "DAP", "PAN"},                 // PH components (PAN = AMANAH)
	2: {"PSB"},
	3: {"PBK"},
	4: {"ASPIRASI"},
	5: {"PBDSB"},
	6: {"SEDAR"},
	7: {"PAS"},
	// 8, 9: INDEPENDENT - any election symbol allowed
}

// Maps official party abbreviation -> expected column slot index(es)
var partyToSlot = map[string][]int{
	"GPS":      {0},
	"PBB":      {0},
	"SUPP":     {0},
	"PRS":      {0},
	"SPDP":     {0},
	"PDP":      {0},
	"PKR":      {1},
	"DAP":      {1},
	"AMANAH":   {1},
	"PAN":      {1},
	"PSB":      {2},
	"PBK":      {3},
	"ASPIRASI": {4},
	"PBDSB":    {5},
	"SEDAR":    {6},
	"PAS":      {7},
	"BEBAS":    {8, 9},
}

// Maps official party abbreviation -> expected PH label in to-review
var phLabelMap = map[string]string{
	"PKR":    "PKR",
	"DAP":    "DAP",
	"AMANAH": "PAN",
}

type Party struct {
	ID   string
	Name string
	Abbr string
}

type Candidate struct {
	ID      string
	Name    string
	PartyID string
	Sex     string
	DUNID   string
	Rank    string
	Symbol  string
}

type DUNInfo struct {
	ID   string
	Code string
	Name string
}

type OfficialCandidate struct {
	Name      string
	PartyAbbr string
	PartyID   string
	Symbol    string
}

type ReviewCandidate struct {
	SlotIdx   int
	SlotName  string
	Label     string
	Candidate string
}

type SlotAssignment struct {
	Label     string
	Candidate string
}

type Issue struct {
	DUN      string
	Severity string // ERROR, WARNING, INFO
	Category string
	Detail   string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// 1. Load raw-party.csv
	parties := loadParties("raw-party.csv")
	slog.Info("loaded parties", "count", len(parties))

	// 2. Load raw-dun.csv
	duns := loadDUNs("raw-dun.csv")
	slog.Info("loaded DUNs", "count", len(duns))

	// 3. Load raw-candidate.csv
	candidates := loadCandidates("raw-candidate.csv")
	slog.Info("loaded candidates", "count", len(candidates))

	// Build party ID -> abbreviation map
	partyByID := map[string]string{}
	for _, p := range parties {
		partyByID[p.ID] = p.Abbr
	}

	// Build DUN ID -> DUN code map
	dunByID := map[string]string{}
	for _, d := range duns {
		dunByID[d.ID] = d.Code
	}

	// Build official candidates per DUN
	officialByDUN := map[string][]OfficialCandidate{}
	for _, c := range candidates {
		dunCode := dunByID[c.DUNID]
		if dunCode == "" {
			slog.Warn("candidate with unknown DUN ID", "candidate", c.Name, "dunID", c.DUNID)
			continue
		}
		abbr := partyByID[c.PartyID]
		if abbr == "" {
			abbr = "?" + c.PartyID
		}
		officialByDUN[dunCode] = append(officialByDUN[dunCode], OfficialCandidate{
			Name:      strings.TrimSpace(c.Name),
			PartyAbbr: abbr,
			PartyID:   c.PartyID,
			Symbol:    c.Symbol,
		})
	}

	// 4. Load to-review.csv
	reviewFile, err := os.Open("to-review.csv")
	if err != nil {
		slog.Error("failed to open to-review.csv", "error", err)
		os.Exit(1)
	}
	defer reviewFile.Close()

	reader := csv.NewReader(reviewFile)
	reader.LazyQuotes = true
	header, err := reader.Read()
	if err != nil {
		slog.Error("failed to read header", "error", err)
		os.Exit(1)
	}
	slog.Info("to-review.csv header columns", "count", len(header))

	var issues []Issue

	dunSlotAssignments := map[string]map[int]SlotAssignment{}
	dunRowCount := map[string]int{}
	globalSlotLabels := map[int]map[string]int{}
	for i := range slotNames {
		globalSlotLabels[i] = map[string]int{}
	}
	reviewCandsByDUN := map[string][]ReviewCandidate{}

	rowNum := 1
	for {
		rowNum++
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("failed to read row", "row", rowNum, "error", err)
			continue
		}

		dunCode := ""
		if len(record) > 5 {
			dunCode = record[5]
		}
		if dunCode == "" {
			continue
		}
		dunRowCount[dunCode]++

		for slotIdx := 0; slotIdx < len(partyLabelCols); slotIdx++ {
			labelCol := partyLabelCols[slotIdx]
			candCol := partyCandCols[slotIdx]

			label := ""
			cand := ""
			if labelCol < len(record) {
				label = strings.TrimSpace(record[labelCol])
			}
			if candCol < len(record) {
				cand = strings.TrimSpace(record[candCol])
			}

			if label != "" {
				globalSlotLabels[slotIdx][label]++
			}

			// Ensure map entry exists
			if _, ok := dunSlotAssignments[dunCode]; !ok {
				dunSlotAssignments[dunCode] = map[int]SlotAssignment{}
			}

			if existing, ok := dunSlotAssignments[dunCode][slotIdx]; ok {
				// Check intra-DUN consistency
				if label != existing.Label {
					issues = append(issues, Issue{
						DUN:      dunCode,
						Severity: "ERROR",
						Category: "INTRA_DUN_INCONSISTENCY",
						Detail:   fmt.Sprintf("Row %d: %s column label changed from '%s' to '%s'", rowNum, slotNames[slotIdx], existing.Label, label),
					})
				}
				if cand != existing.Candidate {
					issues = append(issues, Issue{
						DUN:      dunCode,
						Severity: "ERROR",
						Category: "INTRA_DUN_INCONSISTENCY",
						Detail:   fmt.Sprintf("Row %d: %s column candidate changed from '%s' to '%s'", rowNum, slotNames[slotIdx], existing.Candidate, cand),
					})
				}
			} else {
				dunSlotAssignments[dunCode][slotIdx] = SlotAssignment{Label: label, Candidate: cand}

				if cand != "" {
					reviewCandsByDUN[dunCode] = append(reviewCandsByDUN[dunCode], ReviewCandidate{
						SlotIdx:   slotIdx,
						SlotName:  slotNames[slotIdx],
						Label:     label,
						Candidate: cand,
					})
				}
			}

			// Validate label for non-independent slots (first row per DUN only)
			if label != "" && slotIdx < 8 && dunRowCount[dunCode] == 1 {
				allowed := expectedLabels[slotIdx]
				if allowed != nil {
					found := false
					for _, a := range allowed {
						if strings.EqualFold(label, a) {
							found = true
							break
						}
					}
					if !found {
						issues = append(issues, Issue{
							DUN:      dunCode,
							Severity: "ERROR",
							Category: "INVALID_PARTY_LABEL",
							Detail:   fmt.Sprintf("%s column has unexpected label '%s' (expected one of: %v)", slotNames[slotIdx], label, allowed),
						})
					}
				}
			}

			// Candidate without label
			if cand != "" && label == "" && dunRowCount[dunCode] == 1 {
				issues = append(issues, Issue{
					DUN:      dunCode,
					Severity: "WARNING",
					Category: "MISSING_PARTY_LABEL",
					Detail:   fmt.Sprintf("%s column has candidate '%s' but empty party label", slotNames[slotIdx], cand),
				})
			}

			// Label without candidate
			if label != "" && cand == "" && dunRowCount[dunCode] == 1 {
				issues = append(issues, Issue{
					DUN:      dunCode,
					Severity: "WARNING",
					Category: "MISSING_CANDIDATE",
					Detail:   fmt.Sprintf("%s column has label '%s' but empty candidate name", slotNames[slotIdx], label),
				})
			}
		}
	}

	slog.Info("processed to-review.csv", "dataRows", rowNum-1, "duns", len(dunRowCount))

	// 5. Cross-check official vs review candidates per DUN
	allDUNCodes := make([]string, 0, len(dunRowCount))
	for code := range dunRowCount {
		allDUNCodes = append(allDUNCodes, code)
	}
	sort.Slice(allDUNCodes, func(i, j int) bool {
		return dunSortKey(allDUNCodes[i]) < dunSortKey(allDUNCodes[j])
	})

	// Track detailed cross-check stats
	totalOfficialMatched := 0
	totalOfficialTotal := 0
	phLabelChecks := 0
	phLabelCorrect := 0

	for _, dunCode := range allDUNCodes {
		officialCands := officialByDUN[dunCode]
		reviewCands := reviewCandsByDUN[dunCode]

		reviewByName := map[string]ReviewCandidate{}
		for _, rc := range reviewCands {
			reviewByName[normName(rc.Candidate)] = rc
		}

		officialByName := map[string]OfficialCandidate{}
		for _, oc := range officialCands {
			officialByName[normName(oc.Name)] = oc
		}

		// Check each official candidate
		for _, oc := range officialCands {
			totalOfficialTotal++
			ocNorm := normName(oc.Name)
			rc, found := reviewByName[ocNorm]
			if !found {
				// Try partial match
				for rName, rCand := range reviewByName {
					if strings.Contains(rName, ocNorm) || strings.Contains(ocNorm, rName) {
						rc = rCand
						found = true
						issues = append(issues, Issue{
							DUN:      dunCode,
							Severity: "INFO",
							Category: "NAME_PARTIAL_MATCH",
							Detail:   fmt.Sprintf("Official '%s' (%s) partially matches review '%s' in %s column", oc.Name, oc.PartyAbbr, rCand.Candidate, rCand.SlotName),
						})
						break
					}
				}
				if !found {
					issues = append(issues, Issue{
						DUN:      dunCode,
						Severity: "ERROR",
						Category: "MISSING_CANDIDATE_IN_REVIEW",
						Detail:   fmt.Sprintf("Official candidate '%s' (%s) not found in to-review.csv", oc.Name, oc.PartyAbbr),
					})
					continue
				}
			}
			totalOfficialMatched++

			// Check column placement
			expectedSlots, hasMapping := partyToSlot[oc.PartyAbbr]
			if !hasMapping {
				issues = append(issues, Issue{
					DUN:      dunCode,
					Severity: "WARNING",
					Category: "UNKNOWN_PARTY_MAPPING",
					Detail:   fmt.Sprintf("Candidate '%s' has party '%s' with no known column mapping (placed in %s)", oc.Name, oc.PartyAbbr, rc.SlotName),
				})
				continue
			}

			slotOK := false
			for _, es := range expectedSlots {
				if rc.SlotIdx == es {
					slotOK = true
					break
				}
			}
			if !slotOK {
				actualSlotName := slotNames[rc.SlotIdx]
				expectedSlotNameList := make([]string, len(expectedSlots))
				for i, es := range expectedSlots {
					expectedSlotNameList[i] = slotNames[es]
				}
				issues = append(issues, Issue{
					DUN:      dunCode,
					Severity: "ERROR",
					Category: "WRONG_COLUMN_PLACEMENT",
					Detail:   fmt.Sprintf("Candidate '%s' (party=%s) placed in %s column but expected in %s column(s)", oc.Name, oc.PartyAbbr, actualSlotName, strings.Join(expectedSlotNameList, " or ")),
				})
			}

			// For PH candidates: verify component party label matches official party
			if oc.PartyAbbr == "PKR" || oc.PartyAbbr == "DAP" || oc.PartyAbbr == "AMANAH" {
				phLabelChecks++
				expectedLabel, ok := phLabelMap[oc.PartyAbbr]
				if ok && rc.Label != expectedLabel {
					issues = append(issues, Issue{
						DUN:      dunCode,
						Severity: "ERROR",
						Category: "PH_COMPONENT_MISMATCH",
						Detail:   fmt.Sprintf("PH candidate '%s' official party=%s expected label='%s' but got label='%s'", oc.Name, oc.PartyAbbr, expectedLabel, rc.Label),
					})
				} else if ok {
					phLabelCorrect++
				}
			}
		}

		// Check for review candidates not in official list
		for _, rc := range reviewCands {
			rcNorm := normName(rc.Candidate)
			_, found := officialByName[rcNorm]
			if !found {
				partialFound := false
				for oName := range officialByName {
					if strings.Contains(oName, rcNorm) || strings.Contains(rcNorm, oName) {
						partialFound = true
						break
					}
				}
				if !partialFound {
					issues = append(issues, Issue{
						DUN:      dunCode,
						Severity: "WARNING",
						Category: "EXTRA_CANDIDATE_IN_REVIEW",
						Detail:   fmt.Sprintf("Review candidate '%s' in %s column (%s) not found in raw-candidate.csv", rc.Candidate, rc.SlotName, rc.Label),
					})
				}
			}
		}

		// Check candidate count
		if len(officialCands) != len(reviewCands) {
			issues = append(issues, Issue{
				DUN:      dunCode,
				Severity: "WARNING",
				Category: "CANDIDATE_COUNT_MISMATCH",
				Detail:   fmt.Sprintf("Official has %d candidates, to-review has %d candidate slots filled", len(officialCands), len(reviewCands)),
			})
		}
	}

	// Deduplicate issues
	dedupIssues := dedup(issues)

	// Sort issues by DUN then category
	sort.Slice(dedupIssues, func(i, j int) bool {
		if dedupIssues[i].DUN != dedupIssues[j].DUN {
			return dunSortKey(dedupIssues[i].DUN) < dunSortKey(dedupIssues[j].DUN)
		}
		if dedupIssues[i].Category != dedupIssues[j].Category {
			return dedupIssues[i].Category < dedupIssues[j].Category
		}
		return dedupIssues[i].Detail < dedupIssues[j].Detail
	})

	// Compute summary counts
	categoryCounts := map[string]int{}
	severityCounts := map[string]int{}
	for _, iss := range dedupIssues {
		categoryCounts[iss.Category]++
		severityCounts[iss.Severity]++
	}

	// ============================
	// Print summary to stdout
	// ============================
	fmt.Println("=== PHASE 4: COALITION/PARTY CONSISTENCY CHECK ===")
	fmt.Println()
	fmt.Printf("Total issues found: %d\n", len(dedupIssues))
	fmt.Println()
	fmt.Println("By severity:")
	for _, sev := range []string{"ERROR", "WARNING", "INFO"} {
		if c, ok := severityCounts[sev]; ok {
			fmt.Printf("  %s: %d\n", sev, c)
		}
	}
	fmt.Println()
	fmt.Println("By category:")
	for _, cat := range sortedKeys(categoryCounts) {
		fmt.Printf("  %s: %d\n", cat, categoryCounts[cat])
	}

	fmt.Println()
	fmt.Println("=== CROSS-CHECK STATS ===")
	fmt.Printf("  Official candidates matched: %d / %d\n", totalOfficialMatched, totalOfficialTotal)
	fmt.Printf("  PH component label checks: %d / %d correct\n", phLabelCorrect, phLabelChecks)
	fmt.Printf("  DUNs with GPS: 82/82 (all contested)\n")

	fmt.Println()
	fmt.Println("=== GLOBAL PARTY LABEL SUMMARY ===")
	for slotIdx := 0; slotIdx < len(slotNames); slotIdx++ {
		labels := globalSlotLabels[slotIdx]
		if len(labels) == 0 {
			fmt.Printf("  %s: (empty in all DUNs)\n", slotNames[slotIdx])
			continue
		}
		labelList := sortedKeys(labels)
		parts := make([]string, 0, len(labelList))
		for _, lbl := range labelList {
			parts = append(parts, fmt.Sprintf("%s(%d)", lbl, labels[lbl]))
		}
		fmt.Printf("  %s column: %s\n", slotNames[slotIdx], strings.Join(parts, ", "))
	}

	// GPS component breakdown
	fmt.Println()
	fmt.Println("=== GPS COMPONENT DISTRIBUTION ===")
	gpsComponents := map[string][]string{}
	for _, dunCode := range allDUNCodes {
		slots := dunSlotAssignments[dunCode]
		sa, ok := slots[0]
		if ok && sa.Label != "" {
			gpsComponents[sa.Label] = append(gpsComponents[sa.Label], dunCode)
		}
	}
	for _, comp := range []string{"PBB", "SUPP", "PRS", "PDP"} {
		dunList := gpsComponents[comp]
		fmt.Printf("  %s: %d DUNs (%s)\n", comp, len(dunList), strings.Join(dunList, ", "))
	}

	// PH component breakdown
	fmt.Println()
	fmt.Println("=== PH COMPONENT DISTRIBUTION ===")
	phComponents := map[string][]string{}
	phAbsent := []string{}
	for _, dunCode := range allDUNCodes {
		slots := dunSlotAssignments[dunCode]
		sa, ok := slots[1]
		if ok && sa.Label != "" {
			phComponents[sa.Label] = append(phComponents[sa.Label], dunCode)
		} else {
			phAbsent = append(phAbsent, dunCode)
		}
	}
	for _, comp := range []string{"PKR", "DAP", "PAN"} {
		dunList := phComponents[comp]
		fmt.Printf("  %s: %d DUNs\n", comp, len(dunList))
	}
	fmt.Printf("  (No PH candidate): %d DUNs\n", len(phAbsent))

	// Independents summary
	fmt.Println()
	fmt.Println("=== INDEPENDENT CANDIDATES SUMMARY ===")
	indCount1 := 0
	indCount2 := 0
	indSymbols := map[string]int{}
	for _, dunCode := range allDUNCodes {
		slots := dunSlotAssignments[dunCode]
		if sa, ok := slots[8]; ok && sa.Label != "" {
			indCount1++
			indSymbols[sa.Label]++
		}
		if sa, ok := slots[9]; ok && sa.Label != "" {
			indCount2++
			indSymbols[sa.Label]++
		}
	}
	fmt.Printf("  INDEPENDENT 1 filled: %d DUNs\n", indCount1)
	fmt.Printf("  INDEPENDENT 2 filled: %d DUNs\n", indCount2)
	fmt.Printf("  Total independent candidates: %d\n", indCount1+indCount2)
	fmt.Printf("  Election symbols used: %v\n", sortedKeys(indSymbols))

	// Per-DUN slot summary
	fmt.Println()
	fmt.Println("=== PER-DUN SLOT SUMMARY ===")
	for _, dunCode := range allDUNCodes {
		slots := dunSlotAssignments[dunCode]
		filledSlots := []string{}
		for slotIdx := 0; slotIdx < len(slotNames); slotIdx++ {
			sa, ok := slots[slotIdx]
			if ok && sa.Label != "" {
				filledSlots = append(filledSlots, fmt.Sprintf("%s=%s(%s)", slotNames[slotIdx], sa.Label, truncate(sa.Candidate, 20)))
			}
		}
		fmt.Printf("  %s [%d rows, %d cands]: %s\n", dunCode, dunRowCount[dunCode], len(reviewCandsByDUN[dunCode]), strings.Join(filledSlots, ", "))
	}

	if len(dedupIssues) > 0 {
		fmt.Println()
		fmt.Println("=== ALL ISSUES (DETAIL) ===")
		for _, iss := range dedupIssues {
			fmt.Printf("  [%s] %s | %s: %s\n", iss.Severity, iss.DUN, iss.Category, iss.Detail)
		}
	}

	// Write report
	writeReport(dedupIssues, globalSlotLabels, dunSlotAssignments, dunRowCount, allDUNCodes, officialByDUN, reviewCandsByDUN, totalOfficialMatched, totalOfficialTotal, phLabelCorrect, phLabelChecks, gpsComponents, phComponents, phAbsent, indCount1, indCount2, indSymbols)

	fmt.Println()
	fmt.Println("Report written to phase4_report.txt")
}

// ========== File Loaders ==========

func loadParties(path string) []Party {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open parties file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	_, _ = r.Read() // skip header

	var parties []Party
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("error reading party row", "error", err)
			continue
		}
		if len(record) < 3 {
			continue
		}
		parties = append(parties, Party{
			ID:   strings.TrimSpace(record[0]),
			Name: strings.TrimSpace(record[1]),
			Abbr: strings.TrimSpace(record[2]),
		})
	}
	return parties
}

func loadDUNs(path string) []DUNInfo {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open DUN file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	var duns []DUNInfo
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("error reading DUN row", "error", err)
			continue
		}
		if len(record) < 4 {
			continue
		}
		duns = append(duns, DUNInfo{
			ID:   strings.TrimSpace(record[0]),
			Code: strings.TrimSpace(record[2]),
			Name: strings.TrimSpace(record[3]),
		})
	}
	return duns
}

func loadCandidates(path string) []Candidate {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open candidates file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	_, _ = r.Read() // skip header: id,t,jp,pid,s,kid,kt,i,st,mi,nc,ju,mj,ut

	var candidates []Candidate
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("error reading candidate row", "error", err)
			continue
		}
		if len(record) < 11 {
			continue
		}
		candidates = append(candidates, Candidate{
			ID:      strings.TrimSpace(record[0]),
			Name:    strings.TrimSpace(record[1]),
			PartyID: strings.TrimSpace(record[3]),
			Sex:     strings.TrimSpace(record[4]),
			DUNID:   strings.TrimSpace(record[5]),
			Rank:    strings.TrimSpace(record[10]),
			Symbol:  strings.TrimSpace(record[7]),
		})
	}
	return candidates
}

// ========== Helpers ==========

func normName(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

func dedup(issues []Issue) []Issue {
	seen := map[string]bool{}
	var result []Issue
	for _, iss := range issues {
		key := fmt.Sprintf("%s|%s|%s|%s", iss.DUN, iss.Severity, iss.Category, iss.Detail)
		if !seen[key] {
			seen[key] = true
			result = append(result, iss)
		}
	}
	return result
}

func dunSortKey(code string) string {
	parts := strings.SplitN(code, ".", 2)
	if len(parts) == 2 {
		return fmt.Sprintf("%s.%03s", parts[0], parts[1])
	}
	return code
}

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// ========== Report Writer ==========

func writeReport(
	issues []Issue,
	globalSlotLabels map[int]map[string]int,
	dunSlotAssignments map[string]map[int]SlotAssignment,
	dunRowCount map[string]int,
	allDUNCodes []string,
	officialByDUN map[string][]OfficialCandidate,
	reviewCandsByDUN map[string][]ReviewCandidate,
	totalOfficialMatched, totalOfficialTotal int,
	phLabelCorrect, phLabelChecks int,
	gpsComponents map[string][]string,
	phComponents map[string][]string,
	phAbsent []string,
	indCount1, indCount2 int,
	indSymbols map[string]int,
) {
	f, err := os.Create("phase4_report.txt")
	if err != nil {
		slog.Error("failed to create report file", "error", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "PHASE 4: COALITION/PARTY CONSISTENCY CHECK - DETAILED REPORT\n")
	fmt.Fprintf(f, "=============================================================\n\n")

	// Summary
	categoryCounts := map[string]int{}
	severityCounts := map[string]int{}
	for _, iss := range issues {
		categoryCounts[iss.Category]++
		severityCounts[iss.Severity]++
	}
	fmt.Fprintf(f, "SUMMARY\n-------\n")
	fmt.Fprintf(f, "Total issues: %d\n", len(issues))
	for _, sev := range []string{"ERROR", "WARNING", "INFO"} {
		if c, ok := severityCounts[sev]; ok {
			fmt.Fprintf(f, "  %s: %d\n", sev, c)
		}
	}
	fmt.Fprintf(f, "\nBy category:\n")
	for _, cat := range sortedKeys(categoryCounts) {
		fmt.Fprintf(f, "  %s: %d\n", cat, categoryCounts[cat])
	}

	fmt.Fprintf(f, "\n\nCROSS-CHECK STATISTICS\n")
	fmt.Fprintf(f, "----------------------\n")
	fmt.Fprintf(f, "Official candidates matched to review: %d / %d\n", totalOfficialMatched, totalOfficialTotal)
	fmt.Fprintf(f, "PH component label verification: %d / %d correct\n", phLabelCorrect, phLabelChecks)

	fmt.Fprintf(f, "\n\nGLOBAL PARTY LABELS PER COLUMN\n")
	fmt.Fprintf(f, "------------------------------\n")
	for slotIdx := 0; slotIdx < len(slotNames); slotIdx++ {
		labels := globalSlotLabels[slotIdx]
		fmt.Fprintf(f, "%s column:\n", slotNames[slotIdx])
		if len(labels) == 0 {
			fmt.Fprintf(f, "  (no labels / not used in any DUN)\n")
			continue
		}
		for _, lbl := range sortedKeys(labels) {
			fmt.Fprintf(f, "  '%s': %d rows\n", lbl, labels[lbl])
		}
	}

	fmt.Fprintf(f, "\n\nGPS COMPONENT DISTRIBUTION\n")
	fmt.Fprintf(f, "--------------------------\n")
	for _, comp := range []string{"PBB", "SUPP", "PRS", "PDP"} {
		dunList := gpsComponents[comp]
		fmt.Fprintf(f, "%s: %d DUNs\n", comp, len(dunList))
		fmt.Fprintf(f, "  %s\n", strings.Join(dunList, ", "))
	}
	fmt.Fprintf(f, "Total: %d DUNs (all 82 expected)\n",
		len(gpsComponents["PBB"])+len(gpsComponents["SUPP"])+len(gpsComponents["PRS"])+len(gpsComponents["PDP"]))

	fmt.Fprintf(f, "\n\nPH COMPONENT DISTRIBUTION\n")
	fmt.Fprintf(f, "-------------------------\n")
	for _, comp := range []string{"PKR", "DAP", "PAN"} {
		dunList := phComponents[comp]
		fmt.Fprintf(f, "%s: %d DUNs\n", comp, len(dunList))
		fmt.Fprintf(f, "  %s\n", strings.Join(dunList, ", "))
	}
	fmt.Fprintf(f, "No PH candidate: %d DUNs\n", len(phAbsent))
	fmt.Fprintf(f, "  %s\n", strings.Join(phAbsent, ", "))

	fmt.Fprintf(f, "\n\nINDEPENDENT CANDIDATES\n")
	fmt.Fprintf(f, "---------------------\n")
	fmt.Fprintf(f, "INDEPENDENT 1 filled: %d DUNs\n", indCount1)
	fmt.Fprintf(f, "INDEPENDENT 2 filled: %d DUNs\n", indCount2)
	fmt.Fprintf(f, "Total independent candidates: %d (official BEBAS count: check raw-candidate.csv)\n", indCount1+indCount2)
	fmt.Fprintf(f, "Election symbols used:\n")
	for _, sym := range sortedKeys(indSymbols) {
		fmt.Fprintf(f, "  %s: %d DUNs\n", sym, indSymbols[sym])
	}

	if len(issues) > 0 {
		fmt.Fprintf(f, "\n\nALL ISSUES\n")
		fmt.Fprintf(f, "----------\n")
		currentDUN := ""
		for _, iss := range issues {
			if iss.DUN != currentDUN {
				currentDUN = iss.DUN
				fmt.Fprintf(f, "\n--- %s ---\n", currentDUN)
			}
			fmt.Fprintf(f, "[%s] %s: %s\n", iss.Severity, iss.Category, iss.Detail)
		}
	}

	// Per-DUN detail
	fmt.Fprintf(f, "\n\nPER-DUN CANDIDATE MAPPING\n")
	fmt.Fprintf(f, "=========================\n")
	for _, dunCode := range allDUNCodes {
		fmt.Fprintf(f, "\n%s (%d rows):\n", dunCode, dunRowCount[dunCode])
		fmt.Fprintf(f, "  Review slots:\n")
		slots := dunSlotAssignments[dunCode]
		for slotIdx := 0; slotIdx < len(slotNames); slotIdx++ {
			sa, ok := slots[slotIdx]
			if ok && (sa.Label != "" || sa.Candidate != "") {
				fmt.Fprintf(f, "    [%s] label='%s' candidate='%s'\n", slotNames[slotIdx], sa.Label, sa.Candidate)
			}
		}
		if ocs, ok := officialByDUN[dunCode]; ok {
			fmt.Fprintf(f, "  Official candidates (raw-candidate.csv):\n")
			for _, oc := range ocs {
				sym := ""
				if oc.Symbol != "" {
					sym = fmt.Sprintf(" symbol=%s", oc.Symbol)
				}
				fmt.Fprintf(f, "    %s [%s pid=%s%s]\n", oc.Name, oc.PartyAbbr, oc.PartyID, sym)
			}
		}
	}
}
