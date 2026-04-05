package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
)

// ColGroup defines a candidate column group in to-review.csv (0-indexed)
type ColGroup struct {
	Name     string
	LabelIdx int
	CandIdx  int
	SexIdx   int
	AgeIdx   int
	VoteIdx  int
}

var colGroups = []ColGroup{
	{"GPS", 12, 13, 14, 15, 16},
	{"PH", 17, 18, 19, 20, 21},
	{"PSB", 22, 23, 24, 25, 26},
	{"PBK", 27, 28, 29, 30, 31},
	{"ASPIRASI", 32, 33, 34, 35, 36},
	{"PBDSB", 37, 38, 39, 40, 41},
	{"SEDAR", 42, 43, 44, 45, 46},
	{"PAS", 47, 48, 49, 50, 51},
	{"INDEPENDENT 1", 52, 53, 54, 55, 56},
	{"INDEPENDENT 2", 57, 58, 59, 60, 61},
}

// pidToColGroup maps raw-candidate party ID to the expected column group in to-review.csv
func pidToColGroup(pid string) string {
	switch pid {
	case "51":
		return "GPS"
	case "37", "81", "3":
		return "PH"
	case "54":
		return "PSB"
	case "95":
		return "PBK"
	case "47":
		return "ASPIRASI"
	case "49":
		return "PBDSB"
	case "70":
		return "SEDAR"
	case "2":
		return "PAS"
	case "20":
		return "INDEPENDENT"
	default:
		return "UNKNOWN(" + pid + ")"
	}
}

// pidToExpectedLabels returns acceptable party labels in to-review.csv for a given pid
func pidToExpectedLabels(pid string) []string {
	switch pid {
	case "51":
		return []string{"PBB", "SUPP", "PRS", "PDP", "SPDP"} // GPS components
	case "37":
		return []string{"PKR"}
	case "81":
		return []string{"PAN", "AMANAH"}
	case "3":
		return []string{"DAP"}
	case "54":
		return []string{"PSB"}
	case "95":
		return []string{"PBK"}
	case "47":
		return []string{"ASPIRASI"}
	case "49":
		return []string{"PBDSB"}
	case "70":
		return []string{"SEDAR"}
	case "2":
		return []string{"PAS"}
	case "20":
		return nil // Independent — ballot symbol, anything goes
	default:
		return nil
	}
}

// RefCandidate represents a candidate from raw-candidate.csv
type RefCandidate struct {
	Name         string
	PartyID      string
	PartyAbbr    string
	Sex          string // L or P
	Status       string // MNG, KLH, HD
	Votes        string
	CandRank     string
	ExpColGroup  string
	NormName     string
	MatchedInRev bool
}

// ReviewCandidate represents a candidate extracted from to-review.csv
type ReviewCandidate struct {
	Name       string
	PartyLabel string
	Sex        string // MALE or FEMALE
	ColGroup   string
	NormName   string
	MatchedRef bool
}

// Issue represents a comparison finding
type Issue struct {
	DUN      string
	Severity string // ERROR, WARNING, INFO
	Type     string
	Detail   string
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "`", "'")
	// Normalize DR. to DR (with or without period)
	// Collapse multiple spaces
	parts := strings.Fields(s)
	return strings.Join(parts, " ")
}

// bigramSimilarity computes Jaccard similarity based on character bigrams
func bigramSimilarity(a, b string) float64 {
	if a == b {
		return 1.0
	}
	if len(a) < 2 || len(b) < 2 {
		if a == b {
			return 1.0
		}
		return 0.0
	}
	bigramsA := make(map[string]int)
	for i := 0; i < len(a)-1; i++ {
		bigramsA[a[i:i+2]]++
	}
	bigramsB := make(map[string]int)
	for i := 0; i < len(b)-1; i++ {
		bigramsB[b[i:i+2]]++
	}
	intersection := 0
	for bg, cntA := range bigramsA {
		if cntB, ok := bigramsB[bg]; ok {
			if cntA < cntB {
				intersection += cntA
			} else {
				intersection += cntB
			}
		}
	}
	totalA := len(a) - 1
	totalB := len(b) - 1
	union := totalA + totalB - intersection
	if union == 0 {
		return 0.0
	}
	return float64(intersection) / float64(union)
}

func sexRefToReview(s string) string {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "L":
		return "MALE"
	case "P":
		return "FEMALE"
	default:
		return ""
	}
}

func dunSortKey(code string) string {
	// "N.01" -> "N.01", "N.82" -> "N.82"
	// Pad the number part for correct sorting
	parts := strings.SplitN(code, ".", 2)
	if len(parts) == 2 {
		num := parts[1]
		for len(num) < 3 {
			num = "0" + num
		}
		return parts[0] + "." + num
	}
	return code
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// 1. Read party mapping
	partyMap := readPartyMap("raw-party.csv")
	slog.Info("loaded parties", "count", len(partyMap))

	// 2. Read DUN mapping
	dunIDToCode, dunIDToName := readDUNMap("raw-dun.csv")
	slog.Info("loaded DUN mapping", "count", len(dunIDToCode))

	// 3. Read reference candidates
	refByDUN := readRefCandidates("raw-candidate.csv", partyMap, dunIDToCode)
	totalRef := 0
	for _, v := range refByDUN {
		totalRef += len(v)
	}
	slog.Info("loaded reference candidates", "dun_count", len(refByDUN), "total_candidates", totalRef)

	// 4. Read review candidates (first row per DUN)
	revByDUN, revInconsistencies := readReviewCandidates("to-review.csv")
	totalRev := 0
	for _, v := range revByDUN {
		totalRev += len(v)
	}
	slog.Info("loaded review candidates", "dun_count", len(revByDUN), "total_candidates", totalRev)

	// 5. Compare
	issues := compare(refByDUN, revByDUN)

	// Add inconsistency issues
	for _, inc := range revInconsistencies {
		issues = append(issues, inc)
	}

	// 6. Output report
	printReport(issues, refByDUN, revByDUN, dunIDToCode, dunIDToName)
}

func readPartyMap(path string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open party file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		slog.Error("failed to read party CSV", "error", err)
		os.Exit(1)
	}

	m := make(map[string]string)
	for i, rec := range records {
		if i == 0 {
			continue // skip header
		}
		if len(rec) >= 3 {
			m[strings.TrimSpace(rec[0])] = strings.TrimSpace(rec[2])
		}
	}
	return m
}

func readDUNMap(path string) (idToCode map[string]string, idToName map[string]string) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open DUN file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		slog.Error("failed to read DUN CSV", "error", err)
		os.Exit(1)
	}

	idToCode = make(map[string]string)
	idToName = make(map[string]string)
	for _, rec := range records {
		if len(rec) >= 4 {
			id := strings.TrimSpace(rec[0])
			code := strings.TrimSpace(rec[2])
			name := strings.TrimSpace(rec[3])
			idToCode[id] = code
			idToName[id] = name
		}
	}
	return
}

func readRefCandidates(path string, partyMap map[string]string, dunIDToCode map[string]string) map[string][]RefCandidate {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open candidate file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		slog.Error("failed to read candidate CSV", "error", err)
		os.Exit(1)
	}

	m := make(map[string][]RefCandidate)
	for i, rec := range records {
		if i == 0 {
			continue // skip header: id,t,jp,pid,s,kid,kt,i,st,mi,nc,ju,mj,ut
		}
		if len(rec) < 12 {
			slog.Warn("short record in candidates", "line", i+1, "cols", len(rec))
			continue
		}

		name := strings.TrimSpace(rec[1])
		pid := strings.TrimSpace(rec[3])
		sex := strings.TrimSpace(rec[4])
		kid := strings.TrimSpace(rec[5])
		status := strings.TrimSpace(rec[8])
		candRank := strings.TrimSpace(rec[10])
		votes := strings.TrimSpace(rec[11])

		dunCode, ok := dunIDToCode[kid]
		if !ok {
			slog.Warn("unknown DUN ID in candidates", "kid", kid, "candidate", name, "line", i+1)
			continue
		}

		partyAbbr := partyMap[pid]
		if partyAbbr == "" {
			slog.Warn("unknown party ID", "pid", pid, "candidate", name)
			partyAbbr = "PID:" + pid
		}

		colGroup := pidToColGroup(pid)

		m[dunCode] = append(m[dunCode], RefCandidate{
			Name:        name,
			PartyID:     pid,
			PartyAbbr:   partyAbbr,
			Sex:         sex,
			Status:      status,
			Votes:       votes,
			CandRank:    candRank,
			ExpColGroup: colGroup,
			NormName:    normalize(name),
		})
	}

	// Sort candidates within each DUN by candidate rank
	for dun := range m {
		sort.Slice(m[dun], func(i, j int) bool {
			return m[dun][i].CandRank < m[dun][j].CandRank
		})
	}

	return m
}

func readReviewCandidates(path string) (map[string][]ReviewCandidate, []Issue) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("failed to open review file", "path", path, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1 // allow variable fields (quoted fields with commas)
	r.LazyQuotes = true

	allRecords, err := r.ReadAll()
	if err != nil {
		slog.Error("failed to read review CSV", "error", err)
		os.Exit(1)
	}

	if len(allRecords) == 0 {
		slog.Error("review CSV is empty")
		os.Exit(1)
	}

	// Skip header
	dataRecords := allRecords[1:]
	slog.Info("review data rows", "count", len(dataRecords))

	m := make(map[string][]ReviewCandidate)
	firstRowForDUN := make(map[string]int) // DUN -> line number of first occurrence
	var inconsistencies []Issue

	// Track per-DUN candidate signatures for consistency checking
	type candSig struct {
		name       string
		partyLabel string
		colGroup   string
	}
	dunSigs := make(map[string][]candSig) // DUN -> signatures from first row

	for lineIdx, rec := range dataRecords {
		if len(rec) < 63 {
			slog.Warn("short row in review", "line", lineIdx+2, "cols", len(rec))
			continue
		}

		dunCode := strings.TrimSpace(rec[5])
		if dunCode == "" {
			continue
		}

		// Extract candidates from this row
		var rowCands []candSig
		for _, cg := range colGroups {
			if cg.CandIdx >= len(rec) {
				continue
			}
			candName := strings.TrimSpace(rec[cg.CandIdx])
			if candName == "" {
				continue
			}
			partyLabel := ""
			if cg.LabelIdx < len(rec) {
				partyLabel = strings.TrimSpace(rec[cg.LabelIdx])
			}
			rowCands = append(rowCands, candSig{
				name:       candName,
				partyLabel: partyLabel,
				colGroup:   cg.Name,
			})
		}

		if _, seen := firstRowForDUN[dunCode]; !seen {
			// First row for this DUN: record as canonical
			firstRowForDUN[dunCode] = lineIdx + 2

			var candidates []ReviewCandidate
			for _, cg := range colGroups {
				if cg.CandIdx >= len(rec) {
					continue
				}
				candName := strings.TrimSpace(rec[cg.CandIdx])
				if candName == "" {
					continue
				}
				partyLabel := ""
				if cg.LabelIdx < len(rec) {
					partyLabel = strings.TrimSpace(rec[cg.LabelIdx])
				}
				sex := ""
				if cg.SexIdx < len(rec) {
					sex = strings.TrimSpace(rec[cg.SexIdx])
				}

				candidates = append(candidates, ReviewCandidate{
					Name:       candName,
					PartyLabel: partyLabel,
					Sex:        sex,
					ColGroup:   cg.Name,
					NormName:   normalize(candName),
				})
			}
			m[dunCode] = candidates
			dunSigs[dunCode] = rowCands
		} else {
			// Check consistency with first row
			firstSigs := dunSigs[dunCode]
			if len(rowCands) != len(firstSigs) {
				inconsistencies = append(inconsistencies, Issue{
					DUN:      dunCode,
					Severity: "ERROR",
					Type:     "INCONSISTENT_ROW",
					Detail: fmt.Sprintf("Row %d has %d candidates but first row (line %d) has %d candidates",
						lineIdx+2, len(rowCands), firstRowForDUN[dunCode], len(firstSigs)),
				})
			} else {
				for ci := 0; ci < len(firstSigs); ci++ {
					if ci < len(rowCands) {
						if normalize(rowCands[ci].name) != normalize(firstSigs[ci].name) {
							inconsistencies = append(inconsistencies, Issue{
								DUN:      dunCode,
								Severity: "ERROR",
								Type:     "INCONSISTENT_CANDIDATE",
								Detail: fmt.Sprintf("Row %d col %s: %q vs first row: %q",
									lineIdx+2, firstSigs[ci].colGroup, rowCands[ci].name, firstSigs[ci].name),
							})
						}
					}
				}
			}
		}
	}

	return m, inconsistencies
}

func compare(refByDUN map[string][]RefCandidate, revByDUN map[string][]ReviewCandidate) []Issue {
	var issues []Issue

	// Gather all DUN codes
	allDUNs := make(map[string]bool)
	for k := range refByDUN {
		allDUNs[k] = true
	}
	for k := range revByDUN {
		allDUNs[k] = true
	}
	dunCodes := make([]string, 0, len(allDUNs))
	for k := range allDUNs {
		dunCodes = append(dunCodes, k)
	}
	sort.Slice(dunCodes, func(i, j int) bool {
		return dunSortKey(dunCodes[i]) < dunSortKey(dunCodes[j])
	})

	for _, dun := range dunCodes {
		refCands, hasRef := refByDUN[dun]
		revCands, hasRev := revByDUN[dun]

		if !hasRef {
			issues = append(issues, Issue{
				DUN:      dun,
				Severity: "ERROR",
				Type:     "DUN_ONLY_IN_REVIEW",
				Detail:   "DUN found in to-review.csv but not in raw-candidate.csv",
			})
			continue
		}
		if !hasRev {
			issues = append(issues, Issue{
				DUN:      dun,
				Severity: "ERROR",
				Type:     "DUN_ONLY_IN_REFERENCE",
				Detail:   "DUN found in raw-candidate.csv but not in to-review.csv",
			})
			continue
		}

		// Count comparison
		if len(refCands) != len(revCands) {
			issues = append(issues, Issue{
				DUN:      dun,
				Severity: "WARNING",
				Type:     "CANDIDATE_COUNT_MISMATCH",
				Detail:   fmt.Sprintf("Reference has %d candidates, review has %d", len(refCands), len(revCands)),
			})
		}

		// Build lookup maps for review candidates (normalized name -> index)
		revByNorm := make(map[string]int) // norm name -> index in revCands
		for i := range revCands {
			revByNorm[revCands[i].NormName] = i
		}

		// Phase A: Exact normalized match
		for ri := range refCands {
			ref := &refCands[ri]
			if idx, ok := revByNorm[ref.NormName]; ok {
				rev := &revCands[idx]
				ref.MatchedInRev = true
				rev.MatchedRef = true

				// Check original spelling difference
				if ref.Name != rev.Name {
					issues = append(issues, Issue{
						DUN:      dun,
						Severity: "INFO",
						Type:     "NAME_SPELLING_DIFF",
						Detail: fmt.Sprintf("Ref: %q vs Review: %q (col %s) [normalized match]",
							ref.Name, rev.Name, rev.ColGroup),
					})
				}

				// Check party column assignment
				issues = append(issues, checkPartyAssignment(dun, ref, rev)...)

				// Check sex
				issues = append(issues, checkSex(dun, ref, rev)...)

				delete(revByNorm, ref.NormName)
			}
		}

		// Phase B: Fuzzy match for unmatched reference candidates
		for ri := range refCands {
			ref := &refCands[ri]
			if ref.MatchedInRev {
				continue
			}

			bestIdx := -1
			bestSim := 0.0
			bestNorm := ""
			for normName, idx := range revByNorm {
				if revCands[idx].MatchedRef {
					continue
				}
				sim := bigramSimilarity(ref.NormName, normName)
				if sim > bestSim {
					bestSim = sim
					bestIdx = idx
					bestNorm = normName
				}
			}

			if bestSim >= 0.70 && bestIdx >= 0 {
				rev := &revCands[bestIdx]
				ref.MatchedInRev = true
				rev.MatchedRef = true

				issues = append(issues, Issue{
					DUN:      dun,
					Severity: "WARNING",
					Type:     "NAME_MISMATCH",
					Detail: fmt.Sprintf("Ref: %q (%s/%s) ~ Review: %q (col %s, label %s) [similarity=%.2f]",
						ref.Name, ref.PartyAbbr, ref.ExpColGroup,
						rev.Name, rev.ColGroup, rev.PartyLabel, bestSim),
				})

				issues = append(issues, checkPartyAssignment(dun, ref, rev)...)
				issues = append(issues, checkSex(dun, ref, rev)...)

				delete(revByNorm, bestNorm)
			}
		}

		// Report unmatched reference candidates
		for _, ref := range refCands {
			if !ref.MatchedInRev {
				issues = append(issues, Issue{
					DUN:      dun,
					Severity: "ERROR",
					Type:     "MISSING_IN_REVIEW",
					Detail: fmt.Sprintf("Reference candidate %q (%s / pid=%s / %s) NOT found in to-review.csv",
						ref.Name, ref.PartyAbbr, ref.PartyID, ref.ExpColGroup),
				})
			}
		}

		// Report unmatched review candidates
		for _, rev := range revCands {
			if !rev.MatchedRef {
				issues = append(issues, Issue{
					DUN:      dun,
					Severity: "ERROR",
					Type:     "EXTRA_IN_REVIEW",
					Detail: fmt.Sprintf("Review candidate %q (col %s, label %s) NOT found in raw-candidate.csv",
						rev.Name, rev.ColGroup, rev.PartyLabel),
				})
			}
		}
	}

	return issues
}

func checkPartyAssignment(dun string, ref *RefCandidate, rev *ReviewCandidate) []Issue {
	var issues []Issue
	expectedColGroup := ref.ExpColGroup
	actualColGroup := rev.ColGroup

	// For INDEPENDENT, allow INDEPENDENT 1 or INDEPENDENT 2
	if expectedColGroup == "INDEPENDENT" {
		if !strings.HasPrefix(actualColGroup, "INDEPENDENT") {
			issues = append(issues, Issue{
				DUN:      dun,
				Severity: "ERROR",
				Type:     "WRONG_COLUMN",
				Detail: fmt.Sprintf("Candidate %q (BEBAS/pid=%s): expected INDEPENDENT column, found %s (label %s)",
					ref.Name, ref.PartyID, actualColGroup, rev.PartyLabel),
			})
		}
		// Independent label is ballot symbol — no further label check
		return issues
	}

	if expectedColGroup != actualColGroup {
		issues = append(issues, Issue{
			DUN:      dun,
			Severity: "ERROR",
			Type:     "WRONG_COLUMN",
			Detail: fmt.Sprintf("Candidate %q (%s/pid=%s): expected col %s, found col %s (label %s)",
				ref.Name, ref.PartyAbbr, ref.PartyID, expectedColGroup, actualColGroup, rev.PartyLabel),
		})
		return issues // Don't check label if column is wrong
	}

	// Check party label consistency
	expectedLabels := pidToExpectedLabels(ref.PartyID)
	if expectedLabels != nil && len(expectedLabels) > 0 {
		found := false
		for _, el := range expectedLabels {
			if strings.EqualFold(rev.PartyLabel, el) {
				found = true
				break
			}
		}
		if !found {
			issues = append(issues, Issue{
				DUN:      dun,
				Severity: "WARNING",
				Type:     "PARTY_LABEL_UNEXPECTED",
				Detail: fmt.Sprintf("Candidate %q (pid=%s/%s): party label %q not in expected set %v",
					ref.Name, ref.PartyID, ref.PartyAbbr, rev.PartyLabel, expectedLabels),
			})
		}
	}

	return issues
}

func checkSex(dun string, ref *RefCandidate, rev *ReviewCandidate) []Issue {
	var issues []Issue
	expectedSex := sexRefToReview(ref.Sex)
	actualSex := strings.TrimSpace(rev.Sex)

	if expectedSex != "" && actualSex != "" && expectedSex != actualSex {
		issues = append(issues, Issue{
			DUN:      dun,
			Severity: "WARNING",
			Type:     "SEX_MISMATCH",
			Detail: fmt.Sprintf("Candidate %q: ref sex=%s (%s) vs review sex=%s",
				ref.Name, ref.Sex, expectedSex, actualSex),
		})
	}
	return issues
}

func printReport(issues []Issue, refByDUN map[string][]RefCandidate, revByDUN map[string][]ReviewCandidate, dunIDToCode map[string]string, dunIDToName map[string]string) {
	// Build reverse map: code -> name
	codeToName := make(map[string]string)
	for id, code := range dunIDToCode {
		codeToName[code] = dunIDToName[id]
	}

	// Sort issues by DUN then type
	sort.Slice(issues, func(i, j int) bool {
		if issues[i].DUN != issues[j].DUN {
			return dunSortKey(issues[i].DUN) < dunSortKey(issues[j].DUN)
		}
		// Sort by severity: ERROR > WARNING > INFO
		sevOrder := map[string]int{"ERROR": 0, "WARNING": 1, "INFO": 2}
		if sevOrder[issues[i].Severity] != sevOrder[issues[j].Severity] {
			return sevOrder[issues[i].Severity] < sevOrder[issues[j].Severity]
		}
		return issues[i].Type < issues[j].Type
	})

	// Count by type and severity
	typeCounts := make(map[string]int)
	sevCounts := make(map[string]int)
	dunIssues := make(map[string][]Issue)
	for _, iss := range issues {
		typeCounts[iss.Type]++
		sevCounts[iss.Severity]++
		dunIssues[iss.DUN] = append(dunIssues[iss.DUN], iss)
	}

	// All DUN codes sorted
	allDUNs := make(map[string]bool)
	for k := range refByDUN {
		allDUNs[k] = true
	}
	for k := range revByDUN {
		allDUNs[k] = true
	}
	dunCodes := make([]string, 0, len(allDUNs))
	for k := range allDUNs {
		dunCodes = append(dunCodes, k)
	}
	sort.Slice(dunCodes, func(i, j int) bool {
		return dunSortKey(dunCodes[i]) < dunSortKey(dunCodes[j])
	})

	fmt.Println("==========================================================")
	fmt.Println("  PHASE 3: CANDIDATE COMPARISON REPORT")
	fmt.Println("  raw-candidate.csv vs to-review.csv")
	fmt.Println("==========================================================")
	fmt.Println()

	// --- SUMMARY ---
	fmt.Println("## SUMMARY")
	fmt.Println()
	totalRefCands := 0
	for _, v := range refByDUN {
		totalRefCands += len(v)
	}
	totalRevCands := 0
	for _, v := range revByDUN {
		totalRevCands += len(v)
	}
	fmt.Printf("  Reference DUNs:          %d\n", len(refByDUN))
	fmt.Printf("  Review DUNs:             %d\n", len(revByDUN))
	fmt.Printf("  Reference candidates:    %d\n", totalRefCands)
	fmt.Printf("  Review candidates:       %d\n", totalRevCands)
	fmt.Printf("  Total issues found:      %d\n", len(issues))
	fmt.Println()

	fmt.Println("  By severity:")
	for _, sev := range []string{"ERROR", "WARNING", "INFO"} {
		if sevCounts[sev] > 0 {
			fmt.Printf("    %-10s %d\n", sev, sevCounts[sev])
		}
	}
	fmt.Println()

	fmt.Println("  By type:")
	typeNames := make([]string, 0, len(typeCounts))
	for t := range typeCounts {
		typeNames = append(typeNames, t)
	}
	sort.Strings(typeNames)
	for _, t := range typeNames {
		fmt.Printf("    %-30s %d\n", t, typeCounts[t])
	}
	fmt.Println()

	// --- DUN-BY-DUN CANDIDATE COMPARISON ---
	fmt.Println("==========================================================")
	fmt.Println("  DUN-BY-DUN CANDIDATE MAPPING")
	fmt.Println("==========================================================")

	for _, dun := range dunCodes {
		name := codeToName[dun]
		refs := refByDUN[dun]
		revs := revByDUN[dun]

		fmt.Printf("\n--- %s %s (ref=%d, rev=%d) ---\n", dun, name, len(refs), len(revs))

		// Print reference candidates
		fmt.Println("  Reference (raw-candidate.csv):")
		for _, ref := range refs {
			marker := "✓"
			if !ref.MatchedInRev {
				marker = "✗"
			}
			fmt.Printf("    %s %-40s  %s(pid=%s)  %s  %s  votes=%s\n",
				marker, ref.Name, ref.PartyAbbr, ref.PartyID, ref.ExpColGroup, ref.Sex, ref.Votes)
		}

		// Print review candidates
		fmt.Println("  Review (to-review.csv):")
		for _, rev := range revs {
			marker := "✓"
			if !rev.MatchedRef {
				marker = "✗"
			}
			fmt.Printf("    %s %-40s  label=%-10s  col=%-15s  sex=%s\n",
				marker, rev.Name, rev.PartyLabel, rev.ColGroup, rev.Sex)
		}

		// Print issues for this DUN
		if dunIss, ok := dunIssues[dun]; ok {
			fmt.Println("  Issues:")
			for _, iss := range dunIss {
				fmt.Printf("    [%s/%s] %s\n", iss.Severity, iss.Type, iss.Detail)
			}
		} else {
			fmt.Println("  ✓ No issues")
		}
	}

	// --- ISSUES GROUPED BY TYPE ---
	fmt.Println()
	fmt.Println("==========================================================")
	fmt.Println("  ALL ISSUES GROUPED BY TYPE")
	fmt.Println("==========================================================")

	for _, t := range typeNames {
		fmt.Printf("\n--- %s (%d) ---\n", t, typeCounts[t])
		for _, iss := range issues {
			if iss.Type == t {
				fmt.Printf("  [%s] %s: %s\n", iss.Severity, iss.DUN, iss.Detail)
			}
		}
	}

	// --- CLEAN DUNs ---
	fmt.Println()
	fmt.Println("==========================================================")
	fmt.Println("  DUNs WITH NO ISSUES")
	fmt.Println("==========================================================")
	cleanCount := 0
	for _, dun := range dunCodes {
		if _, has := dunIssues[dun]; !has {
			name := codeToName[dun]
			fmt.Printf("  %s %s (%d candidates)\n", dun, name, len(refByDUN[dun]))
			cleanCount++
		}
	}
	fmt.Printf("\n  Total clean DUNs: %d / %d\n", cleanCount, len(dunCodes))

	// --- DUNs WITH ISSUES ---
	fmt.Println()
	fmt.Println("==========================================================")
	fmt.Println("  DUNs WITH ISSUES (summary)")
	fmt.Println("==========================================================")
	for _, dun := range dunCodes {
		if dunIss, has := dunIssues[dun]; has {
			name := codeToName[dun]
			errCount := 0
			warnCount := 0
			infoCount := 0
			for _, iss := range dunIss {
				switch iss.Severity {
				case "ERROR":
					errCount++
				case "WARNING":
					warnCount++
				case "INFO":
					infoCount++
				}
			}
			fmt.Printf("  %s %-20s  errors=%d  warnings=%d  info=%d\n",
				dun, name, errCount, warnCount, infoCount)
		}
	}
}
