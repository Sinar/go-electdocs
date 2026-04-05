package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

// RawCandidate represents a candidate from raw-candidates.csv
type RawCandidate struct {
	ID        string
	Name      string
	Job       string
	PartyID   int
	Gender    string // L=Male, P=Female
	KID       int    // constituency ID
	KT        string // type (parlimen/dun)
	Status    string // KLH=won, MNG=lost, HD=lost deposit
	NC        int    // candidate number
	Votes     int    // ju = total votes
	Majority  int    // mj
	PartyAbbr string // resolved from party map
}

// ReviewCandidate represents a candidate extracted from to-review.csv
type ReviewCandidate struct {
	PartySlot     string // e.g., "BN", "PH", "GPS", etc.
	PartyName     string // actual party name in col (e.g., "DAP", "PBB")
	CandidateName string
	Gender        string
	Age           string
	TotalVotes    int
}

// PartySlotDef defines where to find candidate data in to-review.csv
type PartySlotDef struct {
	SlotName     string
	PartyCol     int // 0-indexed
	CandidateCol int
	GenderCol    int
	AgeCol       int
	VoteCol      int
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// Define party slot columns (0-indexed)
	partySlots := []PartySlotDef{
		{"BN", 22, 23, 24, 25, 26},
		{"PH", 27, 28, 29, 30, 31},
		{"PN", 32, 33, 34, 35, 36},
		{"GTA", 37, 38, 39, 40, 41},
		{"GPS", 42, 43, 44, 45, 46},
		{"GRS", 47, 48, 49, 50, 51},
		{"WARISAN", 52, 53, 54, 55, 56},
		{"OTHER PARTY (1)", 57, 58, 59, 60, 61},
		{"OTHER PARTY (2)", 62, 63, 64, 65, 66},
		{"OTHER PARTY (3)", 67, 68, 69, 70, 71},
		{"INDEPENDENT 1", 72, 73, 74, 75, 76},
		{"INDEPENDENT 2", 77, 78, 79, 80, 81},
		{"INDEPENDENT 3", 82, 83, 84, 85, 86},
	}

	// 1. Parse party data
	partyMap := parsePartyData("raw-party-data-clean.csv")
	slog.Info("loaded party data", "count", len(partyMap))

	// 2. Parse raw candidates for Sarawak parlimen (kid 19200..22200)
	rawCandidates := parseRawCandidates("raw-candidates.csv", partyMap)
	slog.Info("loaded raw candidates", "count", len(rawCandidates))

	// Group raw candidates by PAR code (e.g., "P.192")
	rawByPAR := make(map[string][]RawCandidate)
	for _, rc := range rawCandidates {
		parCode := fmt.Sprintf("P.%d", rc.KID/100)
		rawByPAR[parCode] = append(rawByPAR[parCode], rc)
	}

	// 3. Parse to-review.csv
	reviewCandidates := parseReviewCandidates("to-review.csv", partySlots)
	slog.Info("loaded review candidates", "parCount", len(reviewCandidates))

	// 4. Compare and generate report
	generateReport(rawByPAR, reviewCandidates, partyMap)
}

func parsePartyData(filename string) map[int]string {
	f, err := os.Open(filename)
	if err != nil {
		slog.Error("failed to open party data", "file", filename, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	header, err := r.Read()
	if err != nil {
		slog.Error("failed to read party header", "error", err)
		os.Exit(1)
	}
	slog.Info("party header", "cols", header)

	result := make(map[int]string)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("party data read error", "error", err)
			continue
		}
		id, _ := strconv.Atoi(record[0])
		abbr := strings.TrimSpace(record[2])
		result[id] = abbr
	}
	return result
}

func parseRawCandidates(filename string, partyMap map[int]string) []RawCandidate {
	f, err := os.Open(filename)
	if err != nil {
		slog.Error("failed to open raw candidates", "file", filename, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	_, err = r.Read() // skip header
	if err != nil {
		slog.Error("failed to read candidates header", "error", err)
		os.Exit(1)
	}

	var result []RawCandidate
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("candidate read error", "error", err)
			continue
		}

		kid, _ := strconv.Atoi(record[5])
		kt := strings.TrimSpace(record[6])

		// Filter: Sarawak parlimen only (kid 19200..22200)
		if kt != "parlimen" || kid < 19200 || kid > 22200 {
			continue
		}

		pid, _ := strconv.Atoi(record[3])
		nc, _ := strconv.Atoi(record[10])
		ju, _ := strconv.Atoi(record[11])
		mj, _ := strconv.Atoi(record[12])

		rc := RawCandidate{
			ID:        strings.TrimSpace(record[0]),
			Name:      strings.TrimSpace(record[1]),
			Job:       strings.TrimSpace(record[2]),
			PartyID:   pid,
			Gender:    strings.TrimSpace(record[4]),
			KID:       kid,
			KT:        kt,
			Status:    strings.TrimSpace(record[8]),
			NC:        nc,
			Votes:     ju,
			Majority:  mj,
			PartyAbbr: partyMap[pid],
		}
		result = append(result, rc)
	}
	return result
}

// parseReviewCandidates extracts candidate info per PAR from to-review.csv
// Returns map[parCode] -> map[candidateKey] -> ReviewCandidate
// where candidateKey = normalized(candidateName)
func parseReviewCandidates(filename string, slots []PartySlotDef) map[string]map[string]*ReviewCandidate {
	f, err := os.Open(filename)
	if err != nil {
		slog.Error("failed to open review file", "file", filename, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1 // allow variable fields

	header, err := r.Read()
	if err != nil {
		slog.Error("failed to read review header", "error", err)
		os.Exit(1)
	}
	slog.Info("review header columns", "count", len(header))

	result := make(map[string]map[string]*ReviewCandidate)
	lineNum := 1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Warn("review read error", "line", lineNum, "error", err)
			lineNum++
			continue
		}
		lineNum++

		if len(record) < 87 {
			slog.Warn("short row", "line", lineNum, "cols", len(record))
			continue
		}

		parCode := strings.TrimSpace(record[3]) // col 4 = PARLIAMENTARY CODE
		if parCode == "" {
			continue
		}

		if _, ok := result[parCode]; !ok {
			result[parCode] = make(map[string]*ReviewCandidate)
		}

		for _, slot := range slots {
			if slot.CandidateCol >= len(record) {
				continue
			}
			candidateName := strings.TrimSpace(record[slot.CandidateCol])
			if candidateName == "" {
				continue
			}

			partyName := ""
			if slot.PartyCol < len(record) {
				partyName = strings.TrimSpace(record[slot.PartyCol])
			}

			gender := ""
			if slot.GenderCol < len(record) {
				gender = strings.TrimSpace(record[slot.GenderCol])
			}

			age := ""
			if slot.AgeCol < len(record) {
				age = strings.TrimSpace(record[slot.AgeCol])
			}

			votes := 0
			if slot.VoteCol < len(record) {
				v := strings.TrimSpace(record[slot.VoteCol])
				if v != "" {
					votes, _ = strconv.Atoi(v)
				}
			}

			key := normalizeName(candidateName) + "|" + partyName

			if existing, ok := result[parCode][key]; ok {
				existing.TotalVotes += votes
			} else {
				result[parCode][key] = &ReviewCandidate{
					PartySlot:     slot.SlotName,
					PartyName:     partyName,
					CandidateName: candidateName,
					Gender:        gender,
					Age:           age,
					TotalVotes:    votes,
				}
			}
		}
	}

	return result
}

func normalizeName(name string) string {
	name = strings.ToUpper(strings.TrimSpace(name))
	// Collapse multiple spaces
	for strings.Contains(name, "  ") {
		name = strings.ReplaceAll(name, "  ", " ")
	}
	return name
}

type ComparisonResult struct {
	PARCode         string
	PARName         string
	RawCount        int
	ReviewCount     int
	Matched         []MatchInfo
	MissingInReview []RawCandidate
	ExtraInReview   []*ReviewCandidate
	VoteMismatches  []VoteMismatch
}

type MatchInfo struct {
	RawName      string
	ReviewName   string
	RawParty     string
	ReviewParty  string
	ReviewSlot   string
	NameExact    bool
	RawGender    string
	ReviewGender string
	GenderMatch  bool
}

type VoteMismatch struct {
	CandidateName string
	Party         string
	RawVotes      int
	ReviewVotes   int
	Diff          int
}

func generateReport(rawByPAR map[string][]RawCandidate, reviewByPAR map[string]map[string]*ReviewCandidate, partyMap map[int]string) {
	// Get sorted PAR codes
	var parCodes []string
	allPARs := make(map[string]bool)
	for k := range rawByPAR {
		allPARs[k] = true
	}
	for k := range reviewByPAR {
		allPARs[k] = true
	}
	for k := range allPARs {
		parCodes = append(parCodes, k)
	}
	sort.Slice(parCodes, func(i, j int) bool {
		ni := extractPARNum(parCodes[i])
		nj := extractPARNum(parCodes[j])
		return ni < nj
	})

	var results []ComparisonResult
	totalMatched := 0
	totalMissing := 0
	totalExtra := 0
	totalNameMismatch := 0
	totalVoteMismatch := 0
	totalGenderMismatch := 0

	// Build a coalition mapping for matching: raw party abbreviation -> likely slot
	// This helps us match raw candidates to review candidates

	for _, parCode := range parCodes {
		rawCands := rawByPAR[parCode]
		reviewCands := reviewByPAR[parCode]

		cr := ComparisonResult{
			PARCode:     parCode,
			RawCount:    len(rawCands),
			ReviewCount: len(reviewCands),
		}

		if len(rawCands) > 0 {
			// Try to get PAR name from review data (we'd need it from a different source)
			// For now, leave empty
		}

		// Track which review candidates have been matched
		matchedReview := make(map[string]bool)

		for _, raw := range rawCands {
			rawNorm := normalizeName(raw.Name)
			found := false

			// Try exact match first (by normalized name)
			for key, rev := range reviewCands {
				revNorm := normalizeName(rev.CandidateName)
				if rawNorm == revNorm {
					found = true
					matchedReview[key] = true

					rawGender := raw.Gender
					revGender := mapGender(rev.Gender)

					mi := MatchInfo{
						RawName:      raw.Name,
						ReviewName:   rev.CandidateName,
						RawParty:     raw.PartyAbbr,
						ReviewParty:  rev.PartyName,
						ReviewSlot:   rev.PartySlot,
						NameExact:    true,
						RawGender:    rawGender,
						ReviewGender: revGender,
						GenderMatch:  rawGender == revGender || rev.Gender == "",
					}
					cr.Matched = append(cr.Matched, mi)

					if !mi.GenderMatch {
						totalGenderMismatch++
					}

					// Check votes
					if raw.Votes != rev.TotalVotes {
						cr.VoteMismatches = append(cr.VoteMismatches, VoteMismatch{
							CandidateName: raw.Name,
							Party:         raw.PartyAbbr,
							RawVotes:      raw.Votes,
							ReviewVotes:   rev.TotalVotes,
							Diff:          rev.TotalVotes - raw.Votes,
						})
						totalVoteMismatch++
					}

					break
				}
			}

			if found {
				totalMatched++
				continue
			}

			// Try fuzzy match: check if one name contains the other, or
			// if they differ by small edit distance
			bestKey := ""
			bestScore := 0.0
			for key, rev := range reviewCands {
				if matchedReview[key] {
					continue
				}
				revNorm := normalizeName(rev.CandidateName)
				score := similarityScore(rawNorm, revNorm)
				if score > bestScore {
					bestScore = score
					bestKey = key
				}
			}

			if bestScore >= 0.70 && bestKey != "" {
				rev := reviewCands[bestKey]
				matchedReview[bestKey] = true
				found = true
				totalMatched++
				totalNameMismatch++

				rawGender := raw.Gender
				revGender := mapGender(rev.Gender)

				mi := MatchInfo{
					RawName:      raw.Name,
					ReviewName:   rev.CandidateName,
					RawParty:     raw.PartyAbbr,
					ReviewParty:  rev.PartyName,
					ReviewSlot:   rev.PartySlot,
					NameExact:    false,
					RawGender:    rawGender,
					ReviewGender: revGender,
					GenderMatch:  rawGender == revGender || rev.Gender == "",
				}
				cr.Matched = append(cr.Matched, mi)

				if !mi.GenderMatch {
					totalGenderMismatch++
				}

				// Check votes
				if raw.Votes != rev.TotalVotes {
					cr.VoteMismatches = append(cr.VoteMismatches, VoteMismatch{
						CandidateName: raw.Name,
						Party:         raw.PartyAbbr,
						RawVotes:      raw.Votes,
						ReviewVotes:   rev.TotalVotes,
						Diff:          rev.TotalVotes - raw.Votes,
					})
					totalVoteMismatch++
				}
			}

			if !found {
				cr.MissingInReview = append(cr.MissingInReview, raw)
				totalMissing++
			}
		}

		// Find extra candidates in review not in raw
		for key, rev := range reviewCands {
			if !matchedReview[key] {
				cr.ExtraInReview = append(cr.ExtraInReview, rev)
				totalExtra++
			}
		}

		results = append(results, cr)
	}

	// Generate markdown report
	var sb strings.Builder

	sb.WriteString("# PHASE 3 REVIEW: Find Missing or Incorrect Candidates\n\n")
	sb.WriteString("## Overview\n\n")
	sb.WriteString("Comparison of candidates in `to-review.csv` against official data in `raw-candidates.csv` for PRU-15 (2022) Sarawak parliamentary constituencies (P.192–P.222).\n\n")

	sb.WriteString("### Summary Statistics\n\n")
	sb.WriteString(fmt.Sprintf("| Metric | Count |\n"))
	sb.WriteString(fmt.Sprintf("|--------|-------|\n"))
	sb.WriteString(fmt.Sprintf("| Parliamentary constituencies checked | %d |\n", len(parCodes)))
	sb.WriteString(fmt.Sprintf("| Total raw (official) candidates | %d |\n", countTotalRaw(rawByPAR)))
	sb.WriteString(fmt.Sprintf("| Total review candidates found | %d |\n", countTotalReview(reviewByPAR)))
	sb.WriteString(fmt.Sprintf("| Candidates matched (exact name) | %d |\n", totalMatched-totalNameMismatch))
	sb.WriteString(fmt.Sprintf("| Candidates matched (fuzzy/similar name) | %d |\n", totalNameMismatch))
	sb.WriteString(fmt.Sprintf("| Candidates MISSING from to-review.csv | %d |\n", totalMissing))
	sb.WriteString(fmt.Sprintf("| EXTRA candidates in to-review.csv (not in raw) | %d |\n", totalExtra))
	sb.WriteString(fmt.Sprintf("| Vote total mismatches | %d |\n", totalVoteMismatch))
	sb.WriteString(fmt.Sprintf("| Gender mismatches | %d |\n", totalGenderMismatch))
	sb.WriteString("\n")

	// Candidate count per PAR
	sb.WriteString("## 1. Candidate Counts per Parliamentary Constituency\n\n")
	sb.WriteString("| PAR Code | Raw (Official) | Review (to-review) | Match? |\n")
	sb.WriteString("|----------|---------------|--------------------|---------|\n")
	for _, cr := range results {
		match := "✅"
		if cr.RawCount != cr.ReviewCount {
			match = fmt.Sprintf("❌ (diff: %+d)", cr.ReviewCount-cr.RawCount)
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %s |\n", cr.PARCode, cr.RawCount, cr.ReviewCount, match))
	}
	sb.WriteString("\n")

	// Name mismatches
	sb.WriteString("## 2. Name Mismatches (Fuzzy Matches / Possible Typos)\n\n")
	anyNameMismatch := false
	for _, cr := range results {
		for _, mi := range cr.Matched {
			if !mi.NameExact {
				if !anyNameMismatch {
					sb.WriteString("| PAR | Raw (Official) Name | Review Name | Raw Party | Review Party | Slot |\n")
					sb.WriteString("|-----|---------------------|-------------|-----------|--------------|------|\n")
					anyNameMismatch = true
				}
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
					cr.PARCode, mi.RawName, mi.ReviewName, mi.RawParty, mi.ReviewParty, mi.ReviewSlot))
			}
		}
	}
	if !anyNameMismatch {
		sb.WriteString("**No name mismatches found.** All candidate names matched exactly.\n")
	}
	sb.WriteString("\n")

	// Gender mismatches
	sb.WriteString("## 3. Gender Mismatches\n\n")
	anyGenderMismatch := false
	for _, cr := range results {
		for _, mi := range cr.Matched {
			if !mi.GenderMatch {
				if !anyGenderMismatch {
					sb.WriteString("| PAR | Candidate | Raw Gender | Review Gender | Party |\n")
					sb.WriteString("|-----|-----------|-----------|---------------|-------|\n")
					anyGenderMismatch = true
				}
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s |\n",
					cr.PARCode, mi.RawName, mi.RawGender, mi.ReviewGender, mi.RawParty))
			}
		}
	}
	if !anyGenderMismatch {
		sb.WriteString("**No gender mismatches found.**\n")
	}
	sb.WriteString("\n")

	// Missing candidates
	sb.WriteString("## 4. Candidates MISSING from to-review.csv\n\n")
	anyMissing := false
	for _, cr := range results {
		if len(cr.MissingInReview) > 0 {
			if !anyMissing {
				sb.WriteString("| PAR | Candidate Name | Party (raw) | Party Abbr | Gender | Votes | Status |\n")
				sb.WriteString("|-----|---------------|-------------|------------|--------|-------|--------|\n")
				anyMissing = true
			}
			for _, m := range cr.MissingInReview {
				sb.WriteString(fmt.Sprintf("| %s | %s | pid=%d | %s | %s | %d | %s |\n",
					cr.PARCode, m.Name, m.PartyID, m.PartyAbbr, m.Gender, m.Votes, m.Status))
			}
		}
	}
	if !anyMissing {
		sb.WriteString("**No missing candidates.** All official candidates are present in to-review.csv.\n")
	}
	sb.WriteString("\n")

	// Extra candidates
	sb.WriteString("## 5. EXTRA Candidates in to-review.csv (Not in Official Data)\n\n")
	anyExtra := false
	for _, cr := range results {
		if len(cr.ExtraInReview) > 0 {
			if !anyExtra {
				sb.WriteString("| PAR | Candidate Name | Party Name | Slot | Total Votes |\n")
				sb.WriteString("|-----|---------------|-----------|------|-------------|\n")
				anyExtra = true
			}
			for _, e := range cr.ExtraInReview {
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %d |\n",
					cr.PARCode, e.CandidateName, e.PartyName, e.PartySlot, e.TotalVotes))
			}
		}
	}
	if !anyExtra {
		sb.WriteString("**No extra candidates found.**\n")
	}
	sb.WriteString("\n")

	// Vote total mismatches
	sb.WriteString("## 6. Vote Total Mismatches\n\n")
	sb.WriteString("Comparison of total votes per candidate: sum of votes across all rows in `to-review.csv` vs `ju` in `raw-candidates.csv`.\n\n")
	anyVoteMismatch := false
	for _, cr := range results {
		if len(cr.VoteMismatches) > 0 {
			if !anyVoteMismatch {
				sb.WriteString("| PAR | Candidate | Party | Raw Votes (ju) | Review Sum | Difference |\n")
				sb.WriteString("|-----|-----------|-------|---------------|------------|------------|\n")
				anyVoteMismatch = true
			}
			for _, vm := range cr.VoteMismatches {
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %d | %d | %+d |\n",
					cr.PARCode, vm.CandidateName, vm.Party, vm.RawVotes, vm.ReviewVotes, vm.Diff))
			}
		}
	}
	if !anyVoteMismatch {
		sb.WriteString("**All vote totals match perfectly.**\n")
	}
	sb.WriteString("\n")

	// Party mapping analysis
	sb.WriteString("## 7. Party Slot Mapping Analysis\n\n")
	sb.WriteString("How raw party IDs map to to-review.csv party slots:\n\n")
	sb.WriteString("| PAR | Candidate | Raw Party (pid→abbr) | Review Party Name | Review Slot |\n")
	sb.WriteString("|-----|-----------|---------------------|-------------------|-------------|\n")

	// Collect unique mappings
	type partyMapping struct {
		rawParty    string
		reviewParty string
		reviewSlot  string
	}
	uniqueMappings := make(map[string]bool)
	var mappingRows []string

	for _, cr := range results {
		for _, mi := range cr.Matched {
			key := fmt.Sprintf("%s->%s(%s)", mi.RawParty, mi.ReviewParty, mi.ReviewSlot)
			if !uniqueMappings[key] {
				uniqueMappings[key] = true
				mappingRows = append(mappingRows, fmt.Sprintf("| %s | %s | %s | %s | %s |",
					cr.PARCode, mi.RawName, mi.RawParty, mi.ReviewParty, mi.ReviewSlot))
			}
		}
	}
	for _, row := range mappingRows {
		sb.WriteString(row + "\n")
	}
	sb.WriteString("\n")

	// Detailed per-constituency comparison
	sb.WriteString("## 8. Detailed Per-Constituency Comparison\n\n")
	for _, cr := range results {
		hasIssue := len(cr.MissingInReview) > 0 || len(cr.ExtraInReview) > 0 || len(cr.VoteMismatches) > 0
		anyFuzzy := false
		for _, mi := range cr.Matched {
			if !mi.NameExact || !mi.GenderMatch {
				anyFuzzy = true
				break
			}
		}
		if !hasIssue && !anyFuzzy {
			continue
		}

		sb.WriteString(fmt.Sprintf("### %s\n\n", cr.PARCode))

		if len(cr.Matched) > 0 {
			sb.WriteString("**Matched candidates:**\n\n")
			sb.WriteString("| # | Raw Name | Review Name | Raw Party | Review Party | Slot | Name Match | Votes Match |\n")
			sb.WriteString("|---|----------|-------------|-----------|--------------|------|------------|-------------|\n")
			for i, mi := range cr.Matched {
				nameMatch := "✅ exact"
				if !mi.NameExact {
					nameMatch = "⚠️ fuzzy"
				}
				voteMatch := "✅"
				for _, vm := range cr.VoteMismatches {
					if normalizeName(vm.CandidateName) == normalizeName(mi.RawName) {
						voteMatch = fmt.Sprintf("❌ raw=%d rev=%d", vm.RawVotes, vm.ReviewVotes)
						break
					}
				}
				sb.WriteString(fmt.Sprintf("| %d | %s | %s | %s | %s | %s | %s | %s |\n",
					i+1, mi.RawName, mi.ReviewName, mi.RawParty, mi.ReviewParty, mi.ReviewSlot, nameMatch, voteMatch))
			}
			sb.WriteString("\n")
		}

		if len(cr.MissingInReview) > 0 {
			sb.WriteString("**Missing from review:**\n")
			for _, m := range cr.MissingInReview {
				sb.WriteString(fmt.Sprintf("- ❌ %s (%s, pid=%d, votes=%d)\n", m.Name, m.PartyAbbr, m.PartyID, m.Votes))
			}
			sb.WriteString("\n")
		}

		if len(cr.ExtraInReview) > 0 {
			sb.WriteString("**Extra in review (not in raw):**\n")
			for _, e := range cr.ExtraInReview {
				sb.WriteString(fmt.Sprintf("- ⚠️ %s (%s in slot %s, votes=%d)\n", e.CandidateName, e.PartyName, e.PartySlot, e.TotalVotes))
			}
			sb.WriteString("\n")
		}
	}

	// Write report
	err := os.WriteFile("PHASE-3-REVIEW.md", []byte(sb.String()), 0644)
	if err != nil {
		slog.Error("failed to write report", "error", err)
		os.Exit(1)
	}
	fmt.Println("Report written to PHASE-3-REVIEW.md")
}

// Helper functions

func extractPARNum(parCode string) int {
	s := strings.TrimPrefix(parCode, "P.")
	n, _ := strconv.Atoi(s)
	return n
}

func countTotalRaw(m map[string][]RawCandidate) int {
	total := 0
	for _, v := range m {
		total += len(v)
	}
	return total
}

func countTotalReview(m map[string]map[string]*ReviewCandidate) int {
	total := 0
	for _, v := range m {
		total += len(v)
	}
	return total
}

func mapGender(reviewGender string) string {
	switch strings.ToUpper(strings.TrimSpace(reviewGender)) {
	case "MALE":
		return "L"
	case "FEMALE":
		return "P"
	default:
		return strings.TrimSpace(reviewGender)
	}
}

// similarityScore computes a rough similarity between two strings
// Uses longest common subsequence ratio
func similarityScore(a, b string) float64 {
	if a == b {
		return 1.0
	}
	if len(a) == 0 || len(b) == 0 {
		return 0.0
	}

	// Try containment first
	if strings.Contains(a, b) || strings.Contains(b, a) {
		shorter := len(a)
		if len(b) < shorter {
			shorter = len(b)
		}
		longer := len(a)
		if len(b) > longer {
			longer = len(b)
		}
		return float64(shorter) / float64(longer)
	}

	// LCS-based similarity
	lcsLen := lcs(a, b)
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	return float64(lcsLen) / float64(maxLen)
}

func lcs(a, b string) int {
	m := len(a)
	n := len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
