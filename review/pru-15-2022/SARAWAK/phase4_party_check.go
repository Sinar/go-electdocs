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

// partyInfo holds party metadata from raw-party-data-clean.csv
type partyInfo struct {
	ID   string
	Name string
	Abbr string
}

// rawCandidate holds a candidate from raw-candidates.csv
type rawCandidate struct {
	ID        string
	Name      string
	PartyID   string
	PartyAbbr string // resolved from partyInfo
	KID       string // constituency ID e.g. 19200
	ParCode   string // mapped P.XXX
	NC        string // candidate number
	Votes     string
}

// reviewSlot holds the party label and candidate name from a slot in to-review.csv
type reviewSlot struct {
	SlotName   string // e.g. "BN", "PH", "PN", "GPS", etc.
	PartyLabel string // actual value in the party column e.g. "DAP", "PBB"
	Candidate  string
}

// parAssignment holds candidate-party assignments for a parliamentary constituency
type parAssignment struct {
	ParCode    string
	ParName    string
	Candidates []reviewSlot
}

// coalitionMapping defines which party abbreviations belong to which slot
var coalitionToSlot = map[string]string{
	// GPS components
	"GPS":  "GPS",
	"PBB":  "GPS",
	"SUPP": "GPS",
	"PRS":  "GPS",
	"PDP":  "GPS",
	// PH components
	"PH":     "PH",
	"DAP":    "PH",
	"PKR":    "PH",
	"AMANAH": "PH",
	"PAN":    "PH", // alternate abbreviation for AMANAH used in to-review.csv
	// PN components
	"PN":      "PN",
	"PAS":     "PN",
	"BERSATU": "PN",
	"PPBM":    "PN", // alternate abbreviation for BERSATU
	// BN components
	"BN":   "BN",
	"UMNO": "BN",
	"MCA":  "BN",
	"MIC":  "BN",
	// Others
	"GRS":     "GRS",
	"WARISAN": "WARISAN",
	// Independent
	"BEBAS": "INDEPENDENT",
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// 1. Load party data
	partyMap, err := loadPartyData("raw-party-data-clean.csv")
	if err != nil {
		slog.Error("Failed to load party data", "error", err)
		os.Exit(1)
	}
	slog.Info("Loaded party data", "count", len(partyMap))

	// 2. Load raw candidates for Sarawak parlimen
	rawCandidates, err := loadRawCandidates("raw-candidates.csv", partyMap)
	if err != nil {
		slog.Error("Failed to load raw candidates", "error", err)
		os.Exit(1)
	}
	slog.Info("Loaded raw candidates for Sarawak parlimen", "count", len(rawCandidates))

	// 3. Load to-review.csv party assignments
	reviewAssignments, err := loadReviewAssignments("to-review.csv")
	if err != nil {
		slog.Error("Failed to load review assignments", "error", err)
		os.Exit(1)
	}
	slog.Info("Loaded review assignments", "parCount", len(reviewAssignments))

	// 4. Compare and generate report
	report := compareAssignments(rawCandidates, reviewAssignments, partyMap)

	// 5. Write report
	err = os.WriteFile("PHASE-4-REVIEW.md", []byte(report), 0644)
	if err != nil {
		slog.Error("Failed to write report", "error", err)
		os.Exit(1)
	}
	slog.Info("Report written to PHASE-4-REVIEW.md")
}

func loadPartyData(filename string) (map[string]partyInfo, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read()
	if err != nil {
		return nil, err
	}

	// Find column indices
	idIdx, nameIdx, abbrIdx := -1, -1, -1
	for i, h := range header {
		switch h {
		case "id":
			idIdx = i
		case "n":
			nameIdx = i
		case "a":
			abbrIdx = i
		}
	}
	if idIdx < 0 || abbrIdx < 0 {
		return nil, fmt.Errorf("missing required columns in party data")
	}

	result := make(map[string]partyInfo)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		name := ""
		if nameIdx >= 0 && nameIdx < len(record) {
			name = record[nameIdx]
		}
		result[record[idIdx]] = partyInfo{
			ID:   record[idIdx],
			Name: name,
			Abbr: record[abbrIdx],
		}
	}
	return result, nil
}

func loadRawCandidates(filename string, partyMap map[string]partyInfo) (map[string][]rawCandidate, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read()
	if err != nil {
		return nil, err
	}

	// Find column indices
	colIdx := make(map[string]int)
	for i, h := range header {
		colIdx[h] = i
	}

	result := make(map[string][]rawCandidate) // keyed by P.XXX
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		kt := record[colIdx["kt"]]
		if kt != "parlimen" {
			continue
		}

		kid := record[colIdx["kid"]]
		kidNum := 0
		fmt.Sscanf(kid, "%d", &kidNum)

		// Sarawak parlimen: P.192 (19200) to P.222 (22200)
		if kidNum < 19200 || kidNum > 22200 {
			continue
		}

		parCode := fmt.Sprintf("P.%d", kidNum/100)
		pid := record[colIdx["pid"]]
		abbr := "UNKNOWN"
		if p, ok := partyMap[pid]; ok {
			abbr = p.Abbr
		}

		c := rawCandidate{
			ID:        record[colIdx["id"]],
			Name:      record[colIdx["t"]],
			PartyID:   pid,
			PartyAbbr: abbr,
			KID:       kid,
			ParCode:   parCode,
			NC:        record[colIdx["nc"]],
			Votes:     record[colIdx["ju"]],
		}
		result[parCode] = append(result[parCode], c)
	}

	// Sort candidates by nc within each par
	for k := range result {
		sort.Slice(result[k], func(i, j int) bool {
			return result[k][i].NC < result[k][j].NC
		})
	}

	return result, nil
}

func loadReviewAssignments(filename string) (map[string]*parAssignment, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	header, err := r.Read()
	if err != nil {
		return nil, err
	}

	// Column indices (0-based): party label columns and candidate columns
	// Col 23(idx22)=BN, Col 24(idx23)=BN CANDIDATE
	// Col 28(idx27)=PH, Col 29(idx28)=PH CANDIDATE
	// Col 33(idx32)=PN, Col 34(idx33)=PN CANDIDATE
	// Col 38(idx37)=GTA, Col 39(idx38)=GTA CANDIDATE
	// Col 43(idx42)=GPS, Col 44(idx43)=GPS CANDIDATE
	// Col 48(idx47)=GRS, Col 49(idx48)=GRS CANDIDATE
	// Col 53(idx52)=WARISAN, Col 54(idx53)=WARISAN CANDIDATE
	// Col 58(idx57)=OTHER PARTY (1), Col 59(idx58)=OTHER PARTY (1) CANDIDATE
	// Col 63(idx62)=OTHER PARTY (2), Col 64(idx63)=OTHER PARTY (2) CANDIDATE
	// Col 68(idx67)=OTHER PARTY (3), Col 69(idx68)=OTHER PARTY (3) CANDIDATE
	// Col 73(idx72)=INDEPENDENT 1, Col 74(idx73)=INDEPENDENT 1 CANDIDATE
	// Col 78(idx77)=INDEPENDENT 2, Col 79(idx78)=INDEPENDENT 2 CANDIDATE
	// Col 83(idx82)=INDEPENDENT 3, Col 84(idx83)=INDEPENDENT 3 CANDIDATE

	type slotDef struct {
		slotName    string
		partyColIdx int
		candColIdx  int
	}

	// Verify header matches expectations
	_ = header
	slots := []slotDef{
		{"BN", 22, 23},
		{"PH", 27, 28},
		{"PN", 32, 33},
		{"GTA", 37, 38},
		{"GPS", 42, 43},
		{"GRS", 47, 48},
		{"WARISAN", 52, 53},
		{"OTHER PARTY (1)", 57, 58},
		{"OTHER PARTY (2)", 62, 63},
		{"OTHER PARTY (3)", 67, 68},
		{"INDEPENDENT 1", 72, 73},
		{"INDEPENDENT 2", 77, 78},
		{"INDEPENDENT 3", 82, 83},
	}

	result := make(map[string]*parAssignment)
	parNameMap := make(map[string]string)

	lineNum := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("CSV read error", "line", lineNum, "error", err)
			lineNum++
			continue
		}
		lineNum++

		if len(record) < 84 {
			continue
		}

		parCode := strings.TrimSpace(record[3]) // col 4 = PARLIAMENTARY CODE
		parName := strings.TrimSpace(record[4]) // col 5 = PARLIAMENTARY NAME

		if parCode == "" {
			continue
		}

		if _, ok := parNameMap[parCode]; !ok {
			parNameMap[parCode] = parName
		}

		if _, ok := result[parCode]; !ok {
			result[parCode] = &parAssignment{
				ParCode: parCode,
				ParName: parName,
			}
		}

		// We only need the first data row per PAR to get party labels + candidate names
		// But let's collect from the first row that has non-numeric party labels
		// (some rows have vote numbers in these columns; we want the label rows)
		for _, slot := range slots {
			partyLabel := strings.TrimSpace(record[slot.partyColIdx])
			candidate := strings.TrimSpace(record[slot.candColIdx])

			// Skip if party label is empty or looks numeric (it's a vote count)
			if partyLabel == "" || isNumeric(partyLabel) {
				continue
			}

			// Skip if this is a header-like row with MALE/FEMALE
			if partyLabel == "MALE" || partyLabel == "FEMALE" {
				continue
			}

			// Check if we already have this slot for this PAR
			found := false
			for _, existing := range result[parCode].Candidates {
				if existing.SlotName == slot.slotName {
					found = true
					break
				}
			}
			if !found && partyLabel != "" {
				result[parCode].Candidates = append(result[parCode].Candidates, reviewSlot{
					SlotName:   slot.slotName,
					PartyLabel: partyLabel,
					Candidate:  candidate,
				})
			}
		}
	}

	return result, nil
}

func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func normalizeCandidate(name string) string {
	name = strings.ToUpper(strings.TrimSpace(name))
	// Remove common prefixes/suffixes for matching
	name = strings.ReplaceAll(name, "  ", " ")
	return name
}

// commonParticles are name particles that should be ignored in fuzzy matching
var commonParticles = map[string]bool{
	"BIN":   true,
	"BINTI": true,
	"ANAK":  true,
	"AK.":   true,
	"AK":    true,
	"A/K":   true,
	"A/L":   true,
	"HJ":    true,
	"HJ.":   true,
	"HJH":   true,
	"HAJI":  true,
	"HAJAH": true,
	"DATO":  true,
	"DATO'": true,
	"DATUK": true,
	"DR":    true,
	"DR.":   true,
	"IR.":   true,
	"YANG":  true,
	"SRI":   true,
	"SERI":  true,
	"@":     true,
}

func candidateMatch(raw, review string) bool {
	r := normalizeCandidate(raw)
	v := normalizeCandidate(review)
	if r == v {
		return true
	}
	// Check if one contains the other (only if both are reasonably long)
	if len(r) > 5 && len(v) > 5 {
		if strings.Contains(r, v) || strings.Contains(v, r) {
			return true
		}
	}
	// Check significant name-part overlap (excluding common particles)
	rParts := significantParts(strings.Fields(r))
	vParts := significantParts(strings.Fields(v))
	if len(rParts) == 0 || len(vParts) == 0 {
		return false
	}
	matchCount := 0
	for _, rp := range rParts {
		for _, vp := range vParts {
			if rp == vp && len(rp) > 2 {
				matchCount++
			}
		}
	}
	// Require at least 2 significant parts to match, or all significant parts if fewer than 2
	minSignificant := len(rParts)
	if len(vParts) < minSignificant {
		minSignificant = len(vParts)
	}
	threshold := 2
	if minSignificant < threshold {
		threshold = minSignificant
	}
	return matchCount >= threshold
}

func significantParts(parts []string) []string {
	var result []string
	for _, p := range parts {
		if !commonParticles[p] {
			result = append(result, p)
		}
	}
	return result
}

func compareAssignments(rawCandidates map[string][]rawCandidate, reviewAssignments map[string]*parAssignment, partyMap map[string]partyInfo) string {
	var sb strings.Builder

	sb.WriteString("# Phase 4 Review: Coalition/Party Assignment Consistency\n\n")
	sb.WriteString("## Objective\n\n")
	sb.WriteString("Verify that party/coalition assignments in `to-review.csv` are consistent with official data from `raw-candidates.csv` and `raw-party-data-clean.csv`.\n\n")

	// Collect all PAR codes
	parCodes := make([]string, 0)
	parSet := make(map[string]bool)
	for k := range rawCandidates {
		if !parSet[k] {
			parCodes = append(parCodes, k)
			parSet[k] = true
		}
	}
	for k := range reviewAssignments {
		if !parSet[k] {
			parCodes = append(parCodes, k)
			parSet[k] = true
		}
	}
	sort.Slice(parCodes, func(i, j int) bool {
		// Sort by numeric part
		var ni, nj int
		fmt.Sscanf(parCodes[i], "P.%d", &ni)
		fmt.Sscanf(parCodes[j], "P.%d", &nj)
		return ni < nj
	})

	// ========= Section 1: Summary of party labels found in to-review.csv =========
	sb.WriteString("## 1. Party Labels Found in `to-review.csv`\n\n")
	sb.WriteString("### Per-Slot Party Labels\n\n")

	slotLabels := make(map[string]map[string][]string) // slot -> label -> []parCodes
	for _, pc := range parCodes {
		ra, ok := reviewAssignments[pc]
		if !ok {
			continue
		}
		for _, c := range ra.Candidates {
			if slotLabels[c.SlotName] == nil {
				slotLabels[c.SlotName] = make(map[string][]string)
			}
			slotLabels[c.SlotName][c.PartyLabel] = append(slotLabels[c.SlotName][c.PartyLabel], pc)
		}
	}

	slotOrder := []string{"BN", "PH", "PN", "GTA", "GPS", "GRS", "WARISAN",
		"OTHER PARTY (1)", "OTHER PARTY (2)", "OTHER PARTY (3)",
		"INDEPENDENT 1", "INDEPENDENT 2", "INDEPENDENT 3"}
	for _, slot := range slotOrder {
		labels, ok := slotLabels[slot]
		if !ok {
			continue
		}
		sb.WriteString(fmt.Sprintf("**%s slot:**\n\n", slot))
		sb.WriteString("| Party Label | Constituencies |\n")
		sb.WriteString("|-------------|----------------|\n")

		labelList := make([]string, 0, len(labels))
		for l := range labels {
			labelList = append(labelList, l)
		}
		sort.Strings(labelList)
		for _, l := range labelList {
			pcs := labels[l]
			sort.Strings(pcs)
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", l, strings.Join(pcs, ", ")))
		}
		sb.WriteString("\n")
	}

	// ========= Section 2: Official candidates vs to-review mapping =========
	sb.WriteString("## 2. Official Candidates vs to-review.csv Mapping\n\n")
	sb.WriteString("For each constituency, compare the official candidate list (from `raw-candidates.csv`) with the slot assignments in `to-review.csv`.\n\n")

	type mismatch struct {
		ParCode       string
		ParName       string
		CandidateName string
		OfficialParty string
		OfficialPID   string
		ExpectedSlot  string
		ActualSlot    string
		ActualLabel   string
		Issue         string
	}

	var mismatches []mismatch
	var allGood []string
	var candidateNameMismatches []mismatch

	// Build full detail table
	sb.WriteString("### Detailed Per-Constituency Mapping\n\n")
	sb.WriteString("| PAR | Candidate (Official) | Official Party (pid) | Expected Slot | Actual Slot | Actual Label | Status |\n")
	sb.WriteString("|-----|---------------------|---------------------|---------------|-------------|--------------|--------|\n")

	for _, pc := range parCodes {
		rawCands, rawOK := rawCandidates[pc]
		ra, revOK := reviewAssignments[pc]

		if !rawOK {
			sb.WriteString(fmt.Sprintf("| %s | — | — | — | — | — | ⚠️ No official candidates found |\n", pc))
			continue
		}
		if !revOK {
			sb.WriteString(fmt.Sprintf("| %s | — | — | — | — | — | ⚠️ Not in to-review.csv |\n", pc))
			continue
		}

		parName := ra.ParName
		parAllGood := true

		for _, rc := range rawCands {
			expectedSlot := getExpectedSlot(rc.PartyAbbr, rc.PartyID)
			// Find matching slot in review
			actualSlot := ""
			actualLabel := ""
			matched := false

			for _, rs := range ra.Candidates {
				if candidateMatch(rc.Name, rs.Candidate) {
					actualSlot = rs.SlotName
					actualLabel = rs.PartyLabel
					matched = true
					break
				}
			}

			status := "✅"
			if !matched {
				status = "⚠️ Not found in to-review"
				parAllGood = false
				mismatches = append(mismatches, mismatch{
					ParCode:       pc,
					ParName:       parName,
					CandidateName: rc.Name,
					OfficialParty: rc.PartyAbbr,
					OfficialPID:   rc.PartyID,
					ExpectedSlot:  expectedSlot,
					ActualSlot:    "NOT FOUND",
					ActualLabel:   "",
					Issue:         "Candidate not found in to-review.csv",
				})
			} else {
				// Check if actual slot matches expected
				slotMatch := checkSlotMatch(expectedSlot, actualSlot, actualLabel, rc.PartyAbbr)
				if !slotMatch {
					status = fmt.Sprintf("❌ Expected %s slot", expectedSlot)
					parAllGood = false
					mismatches = append(mismatches, mismatch{
						ParCode:       pc,
						ParName:       parName,
						CandidateName: rc.Name,
						OfficialParty: rc.PartyAbbr,
						OfficialPID:   rc.PartyID,
						ExpectedSlot:  expectedSlot,
						ActualSlot:    actualSlot,
						ActualLabel:   actualLabel,
						Issue:         fmt.Sprintf("In %s slot (label=%s) but expected %s", actualSlot, actualLabel, expectedSlot),
					})
				}
			}

			sb.WriteString(fmt.Sprintf("| %s | %s | %s (pid=%s) | %s | %s | %s | %s |\n",
				pc, rc.Name, rc.PartyAbbr, rc.PartyID, expectedSlot, actualSlot, actualLabel, status))
		}

		// Also check for candidates in review that aren't in raw
		for _, rs := range ra.Candidates {
			if rs.Candidate == "" {
				continue
			}
			foundInRaw := false
			for _, rc := range rawCands {
				if candidateMatch(rc.Name, rs.Candidate) {
					foundInRaw = true
					break
				}
			}
			if !foundInRaw {
				status := "⚠️ Not in official data"
				parAllGood = false
				sb.WriteString(fmt.Sprintf("| %s | — | — | — | %s | %s | %s (candidate: %s) |\n",
					pc, rs.SlotName, rs.PartyLabel, status, rs.Candidate))
				candidateNameMismatches = append(candidateNameMismatches, mismatch{
					ParCode:       pc,
					ParName:       parName,
					CandidateName: rs.Candidate,
					ActualSlot:    rs.SlotName,
					ActualLabel:   rs.PartyLabel,
					Issue:         "Candidate in to-review.csv not found in official data",
				})
			}
		}

		if parAllGood {
			allGood = append(allGood, pc)
		}
	}

	// ========= Section 3: Coalition mapping validation =========
	sb.WriteString("\n## 3. Coalition Mapping Validation\n\n")
	sb.WriteString("### Expected Coalition → Slot Mapping (PRU-15 2022)\n\n")
	sb.WriteString("| Coalition/Slot | Component Parties | Notes |\n")
	sb.WriteString("|----------------|-------------------|-------|\n")
	sb.WriteString("| GPS | PBB, SUPP, PRS, PDP | Gabungan Parti Sarawak |\n")
	sb.WriteString("| PH | DAP, PKR, AMANAH/PAN | Pakatan Harapan |\n")
	sb.WriteString("| PN | PAS, BERSATU/PPBM | Perikatan Nasional |\n")
	sb.WriteString("| BN | UMNO, MCA, MIC | Barisan Nasional |\n")
	sb.WriteString("| GRS | GRS | Gabungan Rakyat Sabah |\n")
	sb.WriteString("| WARISAN | WARISAN | Parti Warisan Sabah |\n")
	sb.WriteString("| OTHER PARTY | PSB, PBK, SEDAR, ASPIRASI, PBDS, PBM, etc. | Minor parties |\n")
	sb.WriteString("| INDEPENDENT | BEBAS | Independent candidates |\n\n")

	// Check label-to-slot consistency
	sb.WriteString("### Label-to-Slot Consistency Check\n\n")
	sb.WriteString("Verifying that party labels appear in the correct coalition slot:\n\n")

	type labelSlotPair struct {
		Label string
		Slot  string
		PARs  []string
	}
	var wrongSlotPairs []labelSlotPair

	for slot, labels := range slotLabels {
		for label, pars := range labels {
			expected := getExpectedSlotForLabel(label)
			if expected == "" {
				continue
			}
			// Check if the slot mapping is genuinely wrong
			// "INDEPENDENT - KEY" is OK in any INDEPENDENT slot
			if strings.HasPrefix(label, "INDEPENDENT") && strings.HasPrefix(slot, "INDEPENDENT") {
				continue
			}
			// "OTHER PARTY (1)", "OTHER PARTY (2)", etc. all satisfy "OTHER PARTY"
			if expected == "OTHER PARTY" && strings.HasPrefix(slot, "OTHER PARTY") {
				continue
			}
			if expected != slot {
				wrongSlotPairs = append(wrongSlotPairs, labelSlotPair{
					Label: label,
					Slot:  slot,
					PARs:  pars,
				})
			}
		}
	}

	if len(wrongSlotPairs) > 0 {
		sb.WriteString("| Party Label | Found In Slot | Expected Slot | Constituencies |\n")
		sb.WriteString("|-------------|---------------|---------------|----------------|\n")
		sort.Slice(wrongSlotPairs, func(i, j int) bool {
			return wrongSlotPairs[i].Label < wrongSlotPairs[j].Label
		})
		for _, wsp := range wrongSlotPairs {
			expected := getExpectedSlotForLabel(wsp.Label)
			sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
				wsp.Label, wsp.Slot, expected, strings.Join(wsp.PARs, ", ")))
		}
		sb.WriteString("\n")
	} else {
		sb.WriteString("✅ All party labels appear in the correct coalition slots.\n\n")
	}

	// ========= Section 4: Mismatches Summary =========
	sb.WriteString("## 4. Mismatches Summary\n\n")

	if len(mismatches) == 0 && len(candidateNameMismatches) == 0 {
		sb.WriteString("✅ **No mismatches found.** All candidates and party assignments are consistent.\n\n")
	} else {
		if len(mismatches) > 0 {
			sb.WriteString("### Slot Assignment Mismatches\n\n")
			sb.WriteString("| PAR | PAR Name | Candidate | Official Party (pid) | Expected Slot | Actual Slot | Actual Label | Issue |\n")
			sb.WriteString("|-----|----------|-----------|---------------------|---------------|-------------|--------------|-------|\n")
			for _, m := range mismatches {
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s (pid=%s) | %s | %s | %s | %s |\n",
					m.ParCode, m.ParName, m.CandidateName, m.OfficialParty, m.OfficialPID,
					m.ExpectedSlot, m.ActualSlot, m.ActualLabel, m.Issue))
			}
			sb.WriteString("\n")
		}

		if len(candidateNameMismatches) > 0 {
			sb.WriteString("### Candidates in to-review.csv Not Found in Official Data\n\n")
			sb.WriteString("These may be name spelling differences or genuinely unmatched candidates.\n\n")
			sb.WriteString("| PAR | PAR Name | Candidate (to-review) | Slot | Party Label | Issue |\n")
			sb.WriteString("|-----|----------|----------------------|------|-------------|-------|\n")
			for _, m := range candidateNameMismatches {
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
					m.ParCode, m.ParName, m.CandidateName, m.ActualSlot, m.ActualLabel, m.Issue))
			}
			sb.WriteString("\n")
		}
	}

	// ========= Section 5: Constituencies with no issues =========
	sb.WriteString("## 5. Constituencies With No Issues\n\n")
	if len(allGood) > 0 {
		sb.WriteString(fmt.Sprintf("%d out of %d constituencies have fully consistent party assignments:\n\n", len(allGood), len(parCodes)))
		sb.WriteString(strings.Join(allGood, ", "))
		sb.WriteString("\n\n")
	} else {
		sb.WriteString("No constituencies were fully clean — all have at least one potential issue to review.\n\n")
	}

	// ========= Section 6: Special observations =========
	sb.WriteString("## 6. Special Observations\n\n")

	// Check if any GPS component appears outside GPS slot
	sb.WriteString("### GPS Component Party Placement\n\n")
	gpsComponents := []string{"PBB", "SUPP", "PRS", "PDP"}
	gpsIssues := false
	for _, slot := range slotOrder {
		if slot == "GPS" {
			continue
		}
		labels, ok := slotLabels[slot]
		if !ok {
			continue
		}
		for _, comp := range gpsComponents {
			if pars, found := labels[comp]; found {
				sb.WriteString(fmt.Sprintf("- ⚠️ **%s** found in **%s** slot (not GPS): %s\n", comp, slot, strings.Join(pars, ", ")))
				gpsIssues = true
			}
		}
	}
	if !gpsIssues {
		sb.WriteString("✅ All GPS component parties (PBB, SUPP, PRS, PDP) are correctly placed in the GPS slot.\n\n")
	} else {
		sb.WriteString("\n")
	}

	// Check PH components
	sb.WriteString("### PH Component Party Placement\n\n")
	phComponents := []string{"DAP", "PKR", "PAN", "AMANAH"}
	phIssues := false
	for _, slot := range slotOrder {
		if slot == "PH" {
			continue
		}
		labels, ok := slotLabels[slot]
		if !ok {
			continue
		}
		for _, comp := range phComponents {
			if pars, found := labels[comp]; found {
				sb.WriteString(fmt.Sprintf("- ⚠️ **%s** found in **%s** slot (not PH): %s\n", comp, slot, strings.Join(pars, ", ")))
				phIssues = true
			}
		}
	}
	if !phIssues {
		sb.WriteString("✅ All PH component parties (DAP, PKR, PAN/AMANAH) are correctly placed in the PH slot.\n\n")
	} else {
		sb.WriteString("\n")
	}

	// Check PN components
	sb.WriteString("### PN Component Party Placement\n\n")
	pnComponents := []string{"PAS", "BERSATU", "PPBM"}
	pnIssues := false
	for _, slot := range slotOrder {
		if slot == "PN" {
			continue
		}
		labels, ok := slotLabels[slot]
		if !ok {
			continue
		}
		for _, comp := range pnComponents {
			if pars, found := labels[comp]; found {
				sb.WriteString(fmt.Sprintf("- ⚠️ **%s** found in **%s** slot (not PN): %s\n", comp, slot, strings.Join(pars, ", ")))
				pnIssues = true
			}
		}
	}
	if !pnIssues {
		sb.WriteString("✅ All PN component parties (PAS, BERSATU/PPBM) are correctly placed in the PN slot.\n\n")
	} else {
		sb.WriteString("\n")
	}

	// Check BN slot contents
	sb.WriteString("### BN Slot Usage\n\n")
	if bnLabels, ok := slotLabels["BN"]; ok {
		for label, pars := range bnLabels {
			if label != "BN" && label != "UMNO" && label != "MCA" && label != "MIC" {
				sb.WriteString(fmt.Sprintf("- ⚠️ Unexpected label **%s** in BN slot: %s\n", label, strings.Join(pars, ", ")))
			}
		}
		sb.WriteString("\n")
	} else {
		sb.WriteString("No candidates found in BN slot for Sarawak (expected — BN did not contest directly in most Sarawak seats in PRU-15).\n\n")
	}

	// Check for INDEPENDENT label patterns
	sb.WriteString("### Independent Candidate Labels\n\n")
	for _, slot := range slotOrder {
		if !strings.HasPrefix(slot, "INDEPENDENT") {
			continue
		}
		labels, ok := slotLabels[slot]
		if !ok {
			continue
		}
		for label, pars := range labels {
			sb.WriteString(fmt.Sprintf("- **%s** in %s slot: %s\n", label, slot, strings.Join(pars, ", ")))
		}
	}
	sb.WriteString("\n")

	// ========= Section 7: Party ID to Abbreviation Cross-Reference =========
	sb.WriteString("## 7. Party IDs Used in Sarawak Parlimen (raw-candidates.csv)\n\n")
	sb.WriteString("| Party ID | Abbreviation | Full Name | Expected Slot | Count |\n")
	sb.WriteString("|----------|-------------|-----------|---------------|-------|\n")

	pidCount := make(map[string]int)
	for _, cands := range rawCandidates {
		for _, c := range cands {
			pidCount[c.PartyID]++
		}
	}
	pidList := make([]string, 0, len(pidCount))
	for k := range pidCount {
		pidList = append(pidList, k)
	}
	sort.Slice(pidList, func(i, j int) bool {
		var ni, nj int
		fmt.Sscanf(pidList[i], "%d", &ni)
		fmt.Sscanf(pidList[j], "%d", &nj)
		return ni < nj
	})
	for _, pid := range pidList {
		p := partyMap[pid]
		slot := getExpectedSlot(p.Abbr, pid)
		sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %d |\n",
			pid, p.Abbr, p.Name, slot, pidCount[pid]))
	}
	sb.WriteString("\n")

	// ========= Section 8: Recommendations =========
	sb.WriteString("## 8. Recommendations\n\n")

	if len(mismatches) == 0 && len(candidateNameMismatches) == 0 && len(wrongSlotPairs) == 0 {
		sb.WriteString("No action required. All party/coalition assignments appear consistent.\n\n")
	} else {
		sb.WriteString("The following items should be investigated:\n\n")
		if len(mismatches) > 0 {
			sb.WriteString(fmt.Sprintf("1. **%d candidate slot assignment issues** — Review the mismatches table above to verify correct coalition placement.\n", len(mismatches)))
		}
		if len(candidateNameMismatches) > 0 {
			sb.WriteString(fmt.Sprintf("2. **%d candidates in to-review not matched to official data** — These may be spelling differences requiring manual verification.\n", len(candidateNameMismatches)))
		}
		if len(wrongSlotPairs) > 0 {
			sb.WriteString(fmt.Sprintf("3. **%d party labels in wrong coalition slots** — Party labels should be moved to the correct slot.\n", len(wrongSlotPairs)))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func getExpectedSlot(partyAbbr string, partyID string) string {
	abbr := strings.ToUpper(partyAbbr)
	if slot, ok := coalitionToSlot[abbr]; ok {
		return slot
	}

	// Check by party ID
	switch partyID {
	case "1", "17", "18", "19":
		return "BN"
	case "3", "23", "31", "45":
		return "PH"
	case "2", "27", "55":
		return "PN"
	case "5", "11", "22", "24", "32":
		return "GPS"
	case "63":
		return "GRS"
	case "54":
		return "WARISAN"
	case "20":
		return "INDEPENDENT"
	}

	return "OTHER PARTY"
}

func getExpectedSlotForLabel(label string) string {
	upper := strings.ToUpper(label)
	if strings.HasPrefix(upper, "INDEPENDENT") {
		return "" // any INDEPENDENT slot is fine
	}
	if slot, ok := coalitionToSlot[upper]; ok {
		return slot
	}
	// Known OTHER PARTY labels
	otherParties := map[string]bool{
		"PSB": true, "PBK": true, "SEDAR": true, "ASPIRASI": true,
		"PBDS": true, "PBM": true, "PEJUANG": true, "MUDA": true,
		"PUTRA": true, "PCM": true, "PSM": true,
	}
	if otherParties[upper] {
		return "OTHER PARTY"
	}
	return ""
}

func checkSlotMatch(expectedSlot, actualSlot, actualLabel, officialAbbr string) bool {
	// Normalize expected and actual
	if expectedSlot == actualSlot {
		return true
	}

	// GPS component parties in GPS slot is correct
	if expectedSlot == "GPS" && actualSlot == "GPS" {
		return true
	}

	// PH component parties in PH slot is correct
	if expectedSlot == "PH" && actualSlot == "PH" {
		return true
	}

	// PN component parties in PN slot is correct
	if expectedSlot == "PN" && actualSlot == "PN" {
		return true
	}

	// BN component parties in BN slot is correct
	if expectedSlot == "BN" && actualSlot == "BN" {
		return true
	}

	// OTHER PARTY can be in any of the OTHER PARTY slots
	if expectedSlot == "OTHER PARTY" && strings.HasPrefix(actualSlot, "OTHER PARTY") {
		return true
	}

	// INDEPENDENT can be in any INDEPENDENT slot
	if expectedSlot == "INDEPENDENT" && strings.HasPrefix(actualSlot, "INDEPENDENT") {
		return true
	}

	return false
}
