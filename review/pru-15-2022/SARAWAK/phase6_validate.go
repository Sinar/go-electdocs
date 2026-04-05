package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Column indices (0-based) for to-review.csv
const (
	colUniqueCode          = 0
	colParCode             = 3  // PARLIAMENTARY CODE e.g. P.192
	colParName             = 4  // PARLIAMENTARY NAME
	colDunCode             = 5  // STATE CONSTITUENCY CODE e.g. N.02
	colDunName             = 6  // STATE CONSTITUENCY NAME
	colTotalBallotsIssued  = 21 // TOTAL BALLOTS ISSUED
	colBNVote              = 26 // BN VOTE
	colPHVote              = 31 // PH VOTE
	colPNVote              = 36 // PN VOTE
	colGTAVote             = 41 // GTA VOTE
	colGPSVote             = 46 // GPS VOTE
	colGRSVote             = 51 // GRS VOTE
	colWarisanVote         = 56 // WARISAN VOTE
	colOther1Vote          = 61 // OTHER PARTY (1) VOTE
	colOther2Vote          = 66 // OTHER PARTY (2) VOTE
	colOther3Vote          = 71 // OTHER PARTY (3) VOTE
	colIndep1Vote          = 76 // INDEPENDENT 1 VOTE
	colIndep2Vote          = 81 // INDEPENDENT 2 VOTE
	colIndep3Vote          = 86 // INDEPENDENT 3 VOTE
	colTotalValidVotes     = 87 // TOTAL VALID VOTES
	colTotalRejectedVotes  = 88 // TOTAL REJECTED VOTES
	colTotalUnreturnedBall = 89 // TOTAL UNRETURNED BALLOTS

	// Candidate name columns (0-based)
	colBNCandidate      = 23
	colPHCandidate      = 28
	colPNCandidate      = 33
	colGTACandidate     = 38
	colGPSCandidate     = 43
	colGRSCandidate     = 48
	colWarisanCandidate = 53
	colOther1Candidate  = 58
	colOther2Candidate  = 63
	colOther3Candidate  = 68
	colIndep1Candidate  = 73
	colIndep2Candidate  = 78
	colIndep3Candidate  = 83
)

var voteColumns = []int{
	colBNVote, colPHVote, colPNVote, colGTAVote, colGPSVote,
	colGRSVote, colWarisanVote, colOther1Vote, colOther2Vote,
	colOther3Vote, colIndep1Vote, colIndep2Vote, colIndep3Vote,
}

var candidateNameColumns = []int{
	colBNCandidate, colPHCandidate, colPNCandidate, colGTACandidate,
	colGPSCandidate, colGRSCandidate, colWarisanCandidate,
	colOther1Candidate, colOther2Candidate, colOther3Candidate,
	colIndep1Candidate, colIndep2Candidate, colIndep3Candidate,
}

type PARStats struct {
	ParCode string
	ParName string
	A       int // TOTAL BALLOTS ISSUED
	B       int // TOTAL VALID VOTES
	C       int // TOTAL REJECTED VOTES
	D       int // TOTAL UNRETURNED BALLOTS

	// Sum of individual candidate vote columns
	SumCandVotes int

	// Per-candidate vote sums (candidate name → total votes from to-review)
	CandidateVotes map[string]int

	// DUN-level stats
	DUNStats map[string]*DUNStats
}

type DUNStats struct {
	DunCode string
	DunName string
	A       int
	B       int
	C       int
	D       int
	SumCand int
}

type RawCandidate struct {
	KID  int
	Name string
	JU   int // votes for this candidate
	MJ   int // majority (winning margin, same for all cands in constituency)
	UT   int // labelled "ut" in raw — actually equals TOTAL REJECTED VOTES
}

func parseIntField(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	// Strip commas (e.g. "1,616" → "1616")
	s = strings.ReplaceAll(s, ",", "")
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func normName(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	// Collapse multiple spaces
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}

func kidToPar(kid int) string {
	pNum := kid / 100
	return fmt.Sprintf("P.%d", pNum)
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	return r.ReadAll()
}

// fuzzyMatch tries progressively looser matching strategies between two names.
// Returns true if they should be considered the same candidate.
func fuzzyMatch(a, b string) bool {
	if a == b {
		return true
	}
	// Strip periods
	ac := strings.ReplaceAll(a, ".", "")
	bc := strings.ReplaceAll(b, ".", "")
	if ac == bc {
		return true
	}
	// Strip hyphens and apostrophes too
	for _, ch := range []string{"-", "'", "`"} {
		ac = strings.ReplaceAll(ac, ch, "")
		bc = strings.ReplaceAll(bc, ch, "")
	}
	if ac == bc {
		return true
	}
	// Check if one contains the other (for minor prefix/suffix differences)
	if len(a) > 5 && len(b) > 5 {
		if strings.Contains(a, b) || strings.Contains(b, a) {
			return true
		}
	}
	// Levenshtein-like: if only 1-2 characters differ in similarly-lengthed names
	if len(ac) == len(bc) && len(ac) > 8 {
		diffs := 0
		for i := range ac {
			if ac[i] != bc[i] {
				diffs++
			}
		}
		if diffs <= 2 {
			return true
		}
	}
	return false
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// --- Read to-review.csv ---
	slog.Info("reading to-review.csv")
	reviewRows, err := readCSV("to-review.csv")
	if err != nil {
		slog.Error("failed to read to-review.csv", "err", err)
		os.Exit(1)
	}
	slog.Info("to-review.csv loaded", "rows", len(reviewRows)-1)

	// --- Read raw-candidates.csv ---
	slog.Info("reading raw-candidates.csv")
	rawRows, err := readCSV("raw-candidates.csv")
	if err != nil {
		slog.Error("failed to read raw-candidates.csv", "err", err)
		os.Exit(1)
	}
	slog.Info("raw-candidates.csv loaded", "rows", len(rawRows)-1)

	// Parse raw candidates (filter to Sarawak PARs: kid 19200-22200)
	rawCandsByPar := make(map[string][]RawCandidate) // P.XXX → []RawCandidate
	for i := 1; i < len(rawRows); i++ {
		row := rawRows[i]
		if len(row) < 14 {
			continue
		}
		kid := parseIntField(row[5])
		if kid < 19200 || kid > 22200 {
			continue
		}
		parCode := kidToPar(kid)
		rc := RawCandidate{
			KID:  kid,
			Name: normName(row[1]),
			JU:   parseIntField(row[11]),
			MJ:   parseIntField(row[12]),
			UT:   parseIntField(row[13]),
		}
		rawCandsByPar[parCode] = append(rawCandsByPar[parCode], rc)
	}
	slog.Info("parsed raw candidates", "pars", len(rawCandsByPar))

	// Build PAR stats from to-review.csv
	parStats := make(map[string]*PARStats)
	var parOrder []string

	for i := 1; i < len(reviewRows); i++ {
		row := reviewRows[i]
		if len(row) < 90 {
			slog.Warn("short row", "line", i+1, "cols", len(row))
			continue
		}

		parCode := strings.TrimSpace(row[colParCode])
		parName := strings.TrimSpace(row[colParName])
		dunCode := strings.TrimSpace(row[colDunCode])
		dunName := strings.TrimSpace(row[colDunName])

		ps, ok := parStats[parCode]
		if !ok {
			ps = &PARStats{
				ParCode:        parCode,
				ParName:        parName,
				CandidateVotes: make(map[string]int),
				DUNStats:       make(map[string]*DUNStats),
			}
			parStats[parCode] = ps
			parOrder = append(parOrder, parCode)
		}

		a := parseIntField(row[colTotalBallotsIssued])
		b := parseIntField(row[colTotalValidVotes])
		c := parseIntField(row[colTotalRejectedVotes])
		d := parseIntField(row[colTotalUnreturnedBall])

		ps.A += a
		ps.B += b
		ps.C += c
		ps.D += d

		// Sum individual candidate votes
		rowCandSum := 0
		for idx, vc := range voteColumns {
			var v int
			if vc < len(row) {
				v = parseIntField(row[vc])
			}
			rowCandSum += v

			// Also accumulate per-candidate name
			nameCol := candidateNameColumns[idx]
			if nameCol < len(row) {
				name := normName(row[nameCol])
				if name != "" {
					ps.CandidateVotes[name] += v
				}
			}
		}
		ps.SumCandVotes += rowCandSum

		// DUN-level
		ds, ok := ps.DUNStats[dunCode]
		if !ok {
			ds = &DUNStats{DunCode: dunCode, DunName: dunName}
			ps.DUNStats[dunCode] = ds
		}
		ds.A += a
		ds.B += b
		ds.C += c
		ds.D += d
		ds.SumCand += rowCandSum
	}

	// Sort parOrder
	sort.Slice(parOrder, func(i, j int) bool {
		ni := parseIntField(strings.TrimPrefix(parOrder[i], "P."))
		nj := parseIntField(strings.TrimPrefix(parOrder[j], "P."))
		return ni < nj
	})

	// --- Perform Checks ---
	type NameMismatch struct {
		RawName    string
		ReviewName string
		RawJU      int
		ReviewJU   int
	}

	type CheckResult struct {
		ParCode      string
		ParName      string
		A, B, C, D   int
		SumCandVotes int
		RawJU        int
		RawUT        int // labelled ut in raw; actually = rejected votes
		Check1Pass   bool
		Check2Pass   bool
		Check3Pass   bool
		Check3Detail string
		NameIssues   []NameMismatch
		FailingDUNs  []string
	}

	var results []CheckResult
	totalCheck1Fail := 0
	totalCheck2Fail := 0
	totalCheck3Fail := 0

	for _, parCode := range parOrder {
		ps := parStats[parCode]

		rawJU := 0
		rawUT := 0
		rawCands := rawCandsByPar[parCode]
		for _, rc := range rawCands {
			rawJU += rc.JU
			rawUT = rc.UT // same for all candidates in a PAR
		}

		cr := CheckResult{
			ParCode:      parCode,
			ParName:      ps.ParName,
			A:            ps.A,
			B:            ps.B,
			C:            ps.C,
			D:            ps.D,
			SumCandVotes: ps.SumCandVotes,
			RawJU:        rawJU,
			RawUT:        rawUT,
		}

		// Check-1: Sum of candidate vote columns == TOTAL VALID VOTES
		cr.Check1Pass = (ps.SumCandVotes == ps.B)
		if !cr.Check1Pass {
			totalCheck1Fail++
		}

		// Check-2: A == B + C + D
		cr.Check2Pass = (ps.A == ps.B+ps.C+ps.D)
		if !cr.Check2Pass {
			totalCheck2Fail++
		}

		// Check-3: Per-candidate comparison with raw-candidates.csv
		var mismatches []string
		matchedRawIdx := make(map[int]bool)
		matchedReviewName := make(map[string]bool)

		for ri, rc := range rawCands {
			// First try exact match
			reviewVotes, found := ps.CandidateVotes[rc.Name]
			matchedName := rc.Name
			if found {
				matchedRawIdx[ri] = true
				matchedReviewName[rc.Name] = true
				if reviewVotes != rc.JU {
					mismatches = append(mismatches,
						fmt.Sprintf("`%s`: to-review=%d, raw=%d, diff=%+d",
							rc.Name, reviewVotes, rc.JU, reviewVotes-rc.JU))
				}
				continue
			}

			// Try fuzzy match
			bestMatch := ""
			for rn := range ps.CandidateVotes {
				if matchedReviewName[rn] {
					continue
				}
				if fuzzyMatch(rn, rc.Name) {
					bestMatch = rn
					break
				}
			}

			if bestMatch != "" {
				matchedRawIdx[ri] = true
				matchedReviewName[bestMatch] = true
				reviewVotes = ps.CandidateVotes[bestMatch]
				matchedName = bestMatch

				// Record name mismatch for reporting
				if bestMatch != rc.Name {
					cr.NameIssues = append(cr.NameIssues, NameMismatch{
						RawName:    rc.Name,
						ReviewName: bestMatch,
						RawJU:      rc.JU,
						ReviewJU:   reviewVotes,
					})
				}

				if reviewVotes != rc.JU {
					mismatches = append(mismatches,
						fmt.Sprintf("`%s` (matched as `%s`): to-review=%d, raw=%d, diff=%+d",
							rc.Name, matchedName, reviewVotes, rc.JU, reviewVotes-rc.JU))
				}
			} else {
				var names []string
				for rn := range ps.CandidateVotes {
					names = append(names, rn)
				}
				sort.Strings(names)
				mismatches = append(mismatches,
					fmt.Sprintf("RAW candidate `%s` (ju=%d) NOT FOUND in to-review. Available: %v",
						rc.Name, rc.JU, names))
			}
		}

		// Check for candidates in to-review not in raw
		for name, votes := range ps.CandidateVotes {
			if !matchedReviewName[name] && votes > 0 {
				mismatches = append(mismatches,
					fmt.Sprintf("to-review candidate `%s` (votes=%d) NOT FOUND in raw-candidates",
						name, votes))
			}
		}

		cr.Check3Pass = len(mismatches) == 0
		if !cr.Check3Pass {
			totalCheck3Fail++
			cr.Check3Detail = strings.Join(mismatches, "; ")
		}

		// DUN-level check: A == B + C + D for each DUN
		var dunCodes []string
		for dc := range ps.DUNStats {
			dunCodes = append(dunCodes, dc)
		}
		sort.Strings(dunCodes)
		for _, dc := range dunCodes {
			ds := ps.DUNStats[dc]
			if ds.A != ds.B+ds.C+ds.D {
				cr.FailingDUNs = append(cr.FailingDUNs,
					fmt.Sprintf("%s (%s): A=%d, B=%d, C=%d, D=%d, B+C+D=%d, diff=%d",
						dc, ds.DunName, ds.A, ds.B, ds.C, ds.D, ds.B+ds.C+ds.D, ds.A-(ds.B+ds.C+ds.D)))
			}
		}

		results = append(results, cr)
	}

	// --- Generate PHASE-6-REVIEW.md ---
	slog.Info("generating PHASE-6-REVIEW.md",
		"check1_fail", totalCheck1Fail,
		"check2_fail", totalCheck2Fail,
		"check3_fail", totalCheck3Fail)

	var sb strings.Builder
	sb.WriteString("# PHASE 6 REVIEW: PAR-level Validation of TOTAL BALLOTS ISSUED vs Candidate Totals\n\n")
	sb.WriteString("## Objective\n\n")
	sb.WriteString("Validate that `TOTAL BALLOTS ISSUED` in `to-review.csv` is consistent with vote totals,\n")
	sb.WriteString("and cross-check per-candidate vote sums against official data from `raw-candidates.csv`.\n\n")
	sb.WriteString("This is a **PRU-15 (2022) parliamentary election** file, so validation is at the **PAR level** (P.192–P.222).\n\n")

	sb.WriteString("## Method\n\n")
	sb.WriteString("For each PAR constituency (P.192–P.222), we compute from `to-review.csv`:\n\n")
	sb.WriteString("- **A**: Sum of `TOTAL BALLOTS ISSUED` (col 22)\n")
	sb.WriteString("- **B**: Sum of `TOTAL VALID VOTES` (col 88)\n")
	sb.WriteString("- **C**: Sum of `TOTAL REJECTED VOTES` (col 89)\n")
	sb.WriteString("- **D**: Sum of `TOTAL UNRETURNED BALLOTS` (col 90)\n")
	sb.WriteString("- **SumCand**: Sum of all individual candidate vote columns (cols 27,32,37,42,47,52,57,62,67,72,77,82,87)\n")
	sb.WriteString("- **RAW_JU**: Sum of `ju` (candidate votes) from `raw-candidates.csv`\n\n")
	sb.WriteString("### Checks Performed\n\n")
	sb.WriteString("| Check | Description | Condition |\n")
	sb.WriteString("|-------|-------------|----------|\n")
	sb.WriteString("| **Check-1** | Internal consistency: sum of candidate vote columns equals TOTAL VALID VOTES | `SumCand == B` |\n")
	sb.WriteString("| **Check-2** | Ballot accounting: TOTAL BALLOTS ISSUED = Valid + Rejected + Unreturned | `A == B + C + D` |\n")
	sb.WriteString("| **Check-3** | Cross-check: per-candidate vote sums match official raw-candidates.csv | `to-review candidate total == ju` |\n\n")

	// Summary
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **Total PARs**: %d\n", len(results)))
	sb.WriteString(fmt.Sprintf("- **Check-1 failures** (SumCand ≠ B): **%d**\n", totalCheck1Fail))
	sb.WriteString(fmt.Sprintf("- **Check-2 failures** (A ≠ B+C+D): **%d**\n", totalCheck2Fail))
	sb.WriteString(fmt.Sprintf("- **Check-3 failures** (candidate mismatch vs raw): **%d**\n\n", totalCheck3Fail))

	// Summary table
	sb.WriteString("## PAR-Level Summary Table\n\n")
	sb.WriteString("| PAR | Name | A (Issued) | B (Valid) | C (Rejected) | D (Unreturned) | B+C+D | SumCand | RAW_JU | Chk1 | Chk2 | Chk3 |\n")
	sb.WriteString("|-----|------|------------|-----------|--------------|----------------|-------|---------|--------|------|------|------|\n")

	for _, cr := range results {
		chk1 := "✅"
		if !cr.Check1Pass {
			chk1 = "❌"
		}
		chk2 := "✅"
		if !cr.Check2Pass {
			chk2 = "❌"
		}
		chk3 := "✅"
		if !cr.Check3Pass {
			chk3 = "❌"
		}
		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %d | %d | %d | %d | %d | %s | %s | %s |\n",
			cr.ParCode, cr.ParName, cr.A, cr.B, cr.C, cr.D,
			cr.B+cr.C+cr.D, cr.SumCandVotes, cr.RawJU,
			chk1, chk2, chk3))
	}

	// --- Check-1 Failures Detail ---
	sb.WriteString("\n## Check-1 Detail: SumCand vs TOTAL VALID VOTES\n\n")
	if totalCheck1Fail == 0 {
		sb.WriteString("**No failures.** All 31 PARs have sum of candidate vote columns equal to TOTAL VALID VOTES. ✅\n\n")
	} else {
		sb.WriteString("The following PARs have a mismatch between the sum of individual candidate vote columns and TOTAL VALID VOTES:\n\n")
		sb.WriteString("| PAR | Name | SumCand | B (Valid) | Diff |\n")
		sb.WriteString("|-----|------|---------|-----------|------|\n")
		for _, cr := range results {
			if !cr.Check1Pass {
				sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %+d |\n",
					cr.ParCode, cr.ParName, cr.SumCandVotes, cr.B, cr.SumCandVotes-cr.B))
			}
		}
		sb.WriteString("\n")
	}

	// --- Check-2 Failures Detail ---
	sb.WriteString("## Check-2 Detail: A vs B + C + D\n\n")
	if totalCheck2Fail == 0 {
		sb.WriteString("**No failures.** All 31 PARs satisfy `TOTAL BALLOTS ISSUED = TOTAL VALID VOTES + TOTAL REJECTED VOTES + TOTAL UNRETURNED BALLOTS`. ✅\n\n")
	} else {
		sb.WriteString("The following PARs have a ballot accounting mismatch:\n\n")
		sb.WriteString("| PAR | Name | A (Issued) | B+C+D | Diff (A − B−C−D) |\n")
		sb.WriteString("|-----|------|------------|-------|-------------------|\n")
		for _, cr := range results {
			if !cr.Check2Pass {
				sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %+d |\n",
					cr.ParCode, cr.ParName, cr.A, cr.B+cr.C+cr.D, cr.A-(cr.B+cr.C+cr.D)))
			}
		}
		sb.WriteString("\n")

		sb.WriteString("### DUN-Level Breakdown for Failing PARs\n\n")
		for _, cr := range results {
			if !cr.Check2Pass {
				sb.WriteString(fmt.Sprintf("#### %s %s\n\n", cr.ParCode, cr.ParName))
				ps := parStats[cr.ParCode]
				var dunCodes []string
				for dc := range ps.DUNStats {
					dunCodes = append(dunCodes, dc)
				}
				sort.Strings(dunCodes)
				sb.WriteString("| DUN | Name | A | B | C | D | B+C+D | Diff | Status |\n")
				sb.WriteString("|-----|------|---|---|---|---|-------|------|--------|\n")
				for _, dc := range dunCodes {
					ds := ps.DUNStats[dc]
					status := "✅"
					diff := ds.A - (ds.B + ds.C + ds.D)
					if diff != 0 {
						status = "❌"
					}
					sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %d | %d | %d | %+d | %s |\n",
						dc, ds.DunName, ds.A, ds.B, ds.C, ds.D,
						ds.B+ds.C+ds.D, diff, status))
				}
				sb.WriteString("\n")
			}
		}
	}

	// --- Check-3 Failures Detail ---
	sb.WriteString("## Check-3 Detail: Candidate Vote Comparison vs raw-candidates.csv\n\n")
	if totalCheck3Fail == 0 {
		sb.WriteString("**No failures.** All per-candidate vote totals in `to-review.csv` match `raw-candidates.csv`. ✅\n\n")
	} else {
		sb.WriteString("The following PARs have candidate-level vote mismatches or name issues:\n\n")
		for _, cr := range results {
			if !cr.Check3Pass {
				sb.WriteString(fmt.Sprintf("### %s %s\n\n", cr.ParCode, cr.ParName))
				sb.WriteString(fmt.Sprintf("- **SumCand (to-review)**: %d\n", cr.SumCandVotes))
				sb.WriteString(fmt.Sprintf("- **RAW_JU (raw-candidates)**: %d\n", cr.RawJU))
				sb.WriteString(fmt.Sprintf("- **Diff**: %+d\n\n", cr.SumCandVotes-cr.RawJU))

				rawCands := rawCandsByPar[cr.ParCode]
				ps := parStats[cr.ParCode]

				sb.WriteString("| Candidate (raw) | raw ju | to-review votes | Diff | Status |\n")
				sb.WriteString("|-----------------|--------|-----------------|------|--------|\n")

				matchedReview := make(map[string]bool)
				for _, rc := range rawCands {
					rv, found := ps.CandidateVotes[rc.Name]
					matchedName := rc.Name
					if !found {
						for rn := range ps.CandidateVotes {
							if fuzzyMatch(rn, rc.Name) {
								rv = ps.CandidateVotes[rn]
								matchedName = rn
								found = true
								break
							}
						}
					}
					if found {
						matchedReview[matchedName] = true
						status := "✅"
						diff := rv - rc.JU
						if diff != 0 {
							status = "❌"
						}
						nameNote := ""
						if matchedName != rc.Name {
							nameNote = fmt.Sprintf(" ⚠️ matched as `%s`", matchedName)
						}
						sb.WriteString(fmt.Sprintf("| %s | %d | %d | %+d | %s%s |\n",
							rc.Name, rc.JU, rv, diff, status, nameNote))
					} else {
						sb.WriteString(fmt.Sprintf("| %s | %d | NOT FOUND | — | ❌ |\n",
							rc.Name, rc.JU))
					}
				}
				sb.WriteString("\n")
				sb.WriteString("**Issues**: " + cr.Check3Detail + "\n\n")
			}
		}
	}

	// --- Name spelling differences (even for passing Check-3 due to fuzzy match) ---
	sb.WriteString("## Candidate Name Spelling Differences\n\n")
	sb.WriteString("Cases where `to-review.csv` and `raw-candidates.csv` have different spellings for the same candidate (matched by fuzzy logic, votes may match):\n\n")
	anyNameIssues := false
	for _, cr := range results {
		if len(cr.NameIssues) > 0 {
			anyNameIssues = true
			sb.WriteString(fmt.Sprintf("### %s %s\n\n", cr.ParCode, cr.ParName))
			sb.WriteString("| raw-candidates name | to-review name | raw ju | to-review votes | Votes match? |\n")
			sb.WriteString("|---------------------|----------------|--------|-----------------|-------------|\n")
			for _, ni := range cr.NameIssues {
				match := "✅"
				if ni.RawJU != ni.ReviewJU {
					match = "❌"
				}
				sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %s |\n",
					ni.RawName, ni.ReviewName, ni.RawJU, ni.ReviewJU, match))
			}
			sb.WriteString("\n")
		}
	}
	if !anyNameIssues {
		sb.WriteString("**No name spelling differences found.** All candidate names match exactly. ✅\n\n")
	}

	// --- Cross-check: RAW_JU aggregate vs B ---
	sb.WriteString("## Cross-Check: RAW_JU (sum of all candidate ju) vs B (TOTAL VALID VOTES)\n\n")
	sb.WriteString("In a single-vote-per-ballot election, the sum of all candidates' votes should equal TOTAL VALID VOTES.\n\n")
	sb.WriteString("| PAR | Name | RAW_JU | B (Valid) | Diff | Status |\n")
	sb.WriteString("|-----|------|--------|-----------|------|--------|\n")
	rawVsBFail := 0
	for _, cr := range results {
		status := "✅"
		diff := cr.RawJU - cr.B
		if diff != 0 {
			status = "❌"
			rawVsBFail++
		}
		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %+d | %s |\n",
			cr.ParCode, cr.ParName, cr.RawJU, cr.B, diff, status))
	}
	sb.WriteString(fmt.Sprintf("\n**Failures**: %d / %d\n\n", rawVsBFail, len(results)))

	// --- Unreturned Ballots / raw ut field analysis ---
	sb.WriteString("## Discovery: raw-candidates.csv `ut` Field Is Actually TOTAL REJECTED VOTES\n\n")
	sb.WriteString("The `ut` column in `raw-candidates.csv` is labelled suggestively as \"unreturned ballots\",\n")
	sb.WriteString("but empirical comparison shows it matches **TOTAL REJECTED VOTES (C)**, not TOTAL UNRETURNED BALLOTS (D).\n\n")
	sb.WriteString("| PAR | Name | C (Rejected) | D (Unreturned) | raw ut | ut==C? | ut==D? |\n")
	sb.WriteString("|-----|------|--------------|----------------|--------|--------|--------|\n")
	utMatchesC := 0
	utMatchesD := 0
	for _, cr := range results {
		eqC := "✅"
		if cr.RawUT != cr.C {
			eqC = "❌"
		} else {
			utMatchesC++
		}
		eqD := "✅"
		if cr.RawUT != cr.D {
			eqD = "❌"
		} else {
			utMatchesD++
		}
		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %d | %s | %s |\n",
			cr.ParCode, cr.ParName, cr.C, cr.D, cr.RawUT, eqC, eqD))
	}
	sb.WriteString(fmt.Sprintf("\n**ut matches C (Rejected)**: %d / %d\n", utMatchesC, len(results)))
	sb.WriteString(fmt.Sprintf("**ut matches D (Unreturned)**: %d / %d\n\n", utMatchesD, len(results)))
	if utMatchesC == len(results) {
		sb.WriteString("**Conclusion**: The `ut` field in `raw-candidates.csv` contains **TOTAL REJECTED VOTES**, not unreturned ballots. This is a mislabelled column in the raw data source.\n\n")
	} else if utMatchesC > utMatchesD {
		sb.WriteString(fmt.Sprintf("**Conclusion**: The `ut` field predominantly matches TOTAL REJECTED VOTES (%d/%d). Investigate mismatches.\n\n", utMatchesC, len(results)))
	}

	// --- DUN-level internal consistency (all DUNs) ---
	sb.WriteString("## DUN-Level Internal Consistency: A == B + C + D\n\n")
	dunFailCount := 0
	var dunFailDetails []string
	for _, cr := range results {
		if len(cr.FailingDUNs) > 0 {
			for _, fd := range cr.FailingDUNs {
				dunFailCount++
				dunFailDetails = append(dunFailDetails, fmt.Sprintf("- **%s** %s: %s", cr.ParCode, cr.ParName, fd))
			}
		}
	}
	if dunFailCount == 0 {
		sb.WriteString("**No DUN-level failures.** All DUNs across all 31 PARs satisfy `A == B + C + D`. ✅\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("**%d DUN(s) have A ≠ B+C+D**:\n\n", dunFailCount))
		for _, d := range dunFailDetails {
			sb.WriteString(d + "\n")
		}
		sb.WriteString("\n")
	}

	// --- Grand Totals ---
	sb.WriteString("## Grand Totals\n\n")
	grandA, grandB, grandC, grandD, grandSumCand, grandRawJU := 0, 0, 0, 0, 0, 0
	for _, cr := range results {
		grandA += cr.A
		grandB += cr.B
		grandC += cr.C
		grandD += cr.D
		grandSumCand += cr.SumCandVotes
		grandRawJU += cr.RawJU
	}
	sb.WriteString("| Metric | Value |\n")
	sb.WriteString("|--------|-------|\n")
	sb.WriteString(fmt.Sprintf("| Total Ballots Issued (A) | %d |\n", grandA))
	sb.WriteString(fmt.Sprintf("| Total Valid Votes (B) | %d |\n", grandB))
	sb.WriteString(fmt.Sprintf("| Total Rejected Votes (C) | %d |\n", grandC))
	sb.WriteString(fmt.Sprintf("| Total Unreturned Ballots (D) | %d |\n", grandD))
	sb.WriteString(fmt.Sprintf("| B + C + D | %d |\n", grandB+grandC+grandD))
	sb.WriteString(fmt.Sprintf("| A − (B+C+D) | %d |\n", grandA-(grandB+grandC+grandD)))
	sb.WriteString(fmt.Sprintf("| Sum of Candidate Votes (SumCand) | %d |\n", grandSumCand))
	sb.WriteString(fmt.Sprintf("| Sum of raw ju (RAW_JU) | %d |\n", grandRawJU))
	sb.WriteString(fmt.Sprintf("| SumCand − B | %d |\n", grandSumCand-grandB))
	sb.WriteString(fmt.Sprintf("| RAW_JU − B | %d |\n", grandRawJU-grandB))
	sb.WriteString("\n")

	// --- Recommendations ---
	sb.WriteString("## Recommendations\n\n")

	recNum := 1
	if totalCheck1Fail > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Check-1**: %d PAR(s) have sum of candidate votes ≠ TOTAL VALID VOTES. Review these rows to ensure no vote data is missing or duplicated.\n", recNum, totalCheck1Fail))
		recNum++
	}
	if totalCheck2Fail > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Check-2**: %d PAR(s) have ballot accounting errors (A ≠ B+C+D). Investigate the DUN-level breakdown above to identify specific polling stations.\n", recNum, totalCheck2Fail))
		recNum++
	}
	if totalCheck3Fail > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Check-3**: %d PAR(s) have per-candidate issues vs official data. See details above for name mismatches and vote discrepancies.\n", recNum, totalCheck3Fail))
		recNum++
	}
	if rawVsBFail > 0 {
		sb.WriteString(fmt.Sprintf("%d. **RAW_JU vs B**: %d PAR(s) have aggregate raw candidate votes ≠ TOTAL VALID VOTES.\n", recNum, rawVsBFail))
		recNum++
	}

	// Always note the name issue and ut discovery
	hasNameIssues := false
	for _, cr := range results {
		if len(cr.NameIssues) > 0 {
			hasNameIssues = true
			break
		}
	}
	if hasNameIssues {
		sb.WriteString(fmt.Sprintf("%d. **Name Spellings**: Some candidate names differ between `to-review.csv` and `raw-candidates.csv`. While votes match, the spelling in `to-review.csv` should be verified against the original PDF source to determine the correct spelling.\n", recNum))
		recNum++
	}
	if utMatchesC == len(results) {
		sb.WriteString(fmt.Sprintf("%d. **raw-candidates.csv `ut` field**: This field contains TOTAL REJECTED VOTES, not unreturned ballots as the header implies. Future processing should treat `ut` as rejected votes.\n", recNum))
		recNum++
	}

	if totalCheck1Fail == 0 && totalCheck2Fail == 0 && totalCheck3Fail == 0 && rawVsBFail == 0 && !hasNameIssues {
		sb.WriteString("All checks passed. The data in `to-review.csv` is internally consistent and matches the official `raw-candidates.csv` data. ✅\n")
	}

	// Write output
	err = os.WriteFile("PHASE-6-REVIEW.md", []byte(sb.String()), 0644)
	if err != nil {
		slog.Error("failed to write PHASE-6-REVIEW.md", "err", err)
		os.Exit(1)
	}

	slog.Info("PHASE-6-REVIEW.md written successfully",
		"check1_fail", totalCheck1Fail,
		"check2_fail", totalCheck2Fail,
		"check3_fail", totalCheck3Fail,
		"rawVsB_fail", rawVsBFail,
		"ut_matches_C", utMatchesC,
		"ut_matches_D", utMatchesD,
	)
}
