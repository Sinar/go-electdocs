package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// Column indices (0-based) for to-review.csv
const (
	colUniqueCode         = 0
	colBallotType         = 2
	colParCode            = 3
	colParName            = 4
	colDunCode            = 5
	colDunName            = 6
	colPollingDistCode    = 7
	colPollingDistName    = 8
	colPollingCentre      = 9
	colVotingChannel      = 20
	colTotalBallots       = 21
	colTotalValidVotes    = 87
	colTotalRejectedVotes = 88
	colTotalUnreturned    = 89
)

// ReviewRow holds parsed data from one row in to-review.csv
type ReviewRow struct {
	LineNum        int
	UniqueCode     string
	BallotType     string
	ParCode        string
	ParName        string
	DunCode        string
	DunName        string
	PolDistCode    string
	PolDistName    string
	PolCentre      string
	VotingChannel  int
	TotalBallots   int // A
	TotalValid     int // B
	TotalRejected  int // C
	TotalUnreturnd int // D
	RawA           string
	RawB           string
	RawC           string
	RawD           string
}

// ResultRow holds parsed data from one pipe-delimited row in a results file
type ResultRow struct {
	Bil        int    // Row number (Bil column)
	Channel    int    // Saluran
	A          int    // Total Ballots Issued
	B          int    // Total Valid Votes
	C          int    // Rejected
	D          int    // Unreturned
	CandVotes  []int  // Individual candidate votes
	RawLine    string // Original pipe line
	Confident  bool   // Whether all numeric fields parsed cleanly
	IsJumlah   bool   // Is this the JUMLAH (total) row
	LineInFile int    // Line number in the source file
}

// PARSummary aggregates to-review.csv data per PAR
type PARSummary struct {
	ParCode  string
	ParName  string
	SumA     int
	SumB     int
	SumC     int
	SumD     int
	RowCount int
	Rows     []ReviewRow
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// --- Load to-review.csv ---
	reviewRows, err := loadReviewCSV("to-review.csv")
	if err != nil {
		slog.Error("failed to load to-review.csv", "err", err)
		os.Exit(1)
	}
	slog.Info("loaded to-review.csv", "rows", len(reviewRows))

	// Group by PAR, preserving insertion order
	parMap := make(map[string]*PARSummary)
	var parOrder []string
	for i := range reviewRows {
		r := &reviewRows[i]
		ps, ok := parMap[r.ParCode]
		if !ok {
			ps = &PARSummary{ParCode: r.ParCode, ParName: r.ParName}
			parMap[r.ParCode] = ps
			parOrder = append(parOrder, r.ParCode)
		}
		ps.SumA += r.TotalBallots
		ps.SumB += r.TotalValid
		ps.SumC += r.TotalRejected
		ps.SumD += r.TotalUnreturnd
		ps.RowCount++
		ps.Rows = append(ps.Rows, *r)
	}

	fmt.Println(strings.Repeat("=", 90))
	fmt.Println("BALLOT CHECK: to-review.csv vs results/Sarawak-P.{192-222}.csv")
	fmt.Println(strings.Repeat("=", 90))

	// ================================================================
	// PART 1: Internal consistency check (A = B + C + D)
	// ================================================================
	fmt.Println()
	printHeader("PART 1: Internal Consistency — A = B + C + D (per row)")

	mismatchCount := 0
	for _, r := range reviewRows {
		expected := r.TotalValid + r.TotalRejected + r.TotalUnreturnd
		if r.TotalBallots != expected {
			mismatchCount++
			diff := r.TotalBallots - expected
			fmt.Printf("  ❌ row %d | %s | A=%d  B+C+D=%d (B=%d C=%d D=%d)  diff=%+d\n",
				r.LineNum, r.UniqueCode, r.TotalBallots, expected,
				r.TotalValid, r.TotalRejected, r.TotalUnreturnd, diff)
			fmt.Printf("     raw: A=%q  B=%q  C=%q  D=%q\n", r.RawA, r.RawB, r.RawC, r.RawD)
			fmt.Printf("     PAR=%s  DUN=%s  Centre=%s  Channel=%d\n",
				r.ParCode, r.DunCode, r.PolCentre, r.VotingChannel)
		}
	}
	if mismatchCount == 0 {
		fmt.Println("  ✅ All", len(reviewRows), "rows satisfy A = B + C + D")
	} else {
		fmt.Printf("  ❌ %d / %d rows fail A = B + C + D\n", mismatchCount, len(reviewRows))
	}

	// ================================================================
	// PART 2 & 3: Parse results files, compare by Bil + Channel
	// ================================================================
	fmt.Println()
	printHeader("PART 2 & 3: Results File Extraction + Row-by-Row Comparison (by Bil+Channel)")
	fmt.Println()
	fmt.Println("  NOTE: Results files are OCR'd PDFs. Many rows have doubled/garbled numbers.")
	fmt.Println("  Only rows where ALL numeric fields parsed cleanly are compared.")
	fmt.Println("  Mismatches may indicate errors in to-review.csv OR OCR noise.")
	fmt.Println()

	totalCompared := 0
	totalMatched := 0
	totalMismatched := 0
	totalConfident := 0
	totalPipeRows := 0
	parCompareResults := make(map[string]string)

	for pNum := 192; pNum <= 222; pNum++ {
		parCode := fmt.Sprintf("P.%d", pNum)
		resultsFile := fmt.Sprintf("results/Sarawak-%s.csv", parCode)

		ps := parMap[parCode]
		if ps == nil {
			fmt.Printf("  [%s] ⚠  No rows in to-review.csv\n", parCode)
			continue
		}

		resultRows, err := loadResultsFile(resultsFile)
		if err != nil {
			fmt.Printf("  [%s %-20s] ⚠  Could not load: %v\n", parCode, ps.ParName, err)
			parCompareResults[parCode] = "NO FILE"
			continue
		}

		// Separate confident rows from noisy ones, skip JUMLAH
		var confident []ResultRow
		noisy := 0
		for _, rr := range resultRows {
			if rr.IsJumlah {
				continue
			}
			if rr.Confident {
				confident = append(confident, rr)
			} else {
				noisy++
			}
		}
		totalPipeRows += len(resultRows)
		totalConfident += len(confident)

		// Build a lookup from (Bil, Channel) → ResultRow for confident rows
		type bilChan struct {
			bil int
			ch  int
		}
		resultLookup := make(map[bilChan]ResultRow)
		for _, rr := range confident {
			if rr.Bil > 0 && rr.Channel >= 0 {
				key := bilChan{rr.Bil, rr.Channel}
				// First clean occurrence wins
				if _, exists := resultLookup[key]; !exists {
					resultLookup[key] = rr
				}
			}
		}

		// Try to match review rows to results rows.
		// Strategy: assign a sequential Bil to review rows within this PAR,
		// grouping rows by their polling district. Each new district increments Bil.
		// Within a district, channel number identifies the row.
		reviewBils := assignBilNumbers(ps.Rows)

		matched := 0
		mismatched := 0
		compared := 0

		var mismatches []string

		for i, rv := range ps.Rows {
			bil := reviewBils[i]
			key := bilChan{bil, rv.VotingChannel}
			rr, found := resultLookup[key]
			if !found {
				continue
			}
			compared++
			if rv.TotalBallots == rr.A {
				matched++
			} else {
				mismatched++
				detail := fmt.Sprintf("    Bil=%d Ch=%d | review_line=%d %s | review_A=%d  results_A=%d  diff=%+d",
					bil, rv.VotingChannel, rv.LineNum, rv.UniqueCode, rv.TotalBallots, rr.A, rv.TotalBallots-rr.A)
				if rv.TotalValid != rr.B || rv.TotalRejected != rr.C || rv.TotalUnreturnd != rr.D {
					detail += fmt.Sprintf("\n      B: %d vs %d  C: %d vs %d  D: %d vs %d",
						rv.TotalValid, rr.B, rv.TotalRejected, rr.C, rv.TotalUnreturnd, rr.D)
				}
				detail += fmt.Sprintf("\n      results line %d: %s", rr.LineInFile, truncate(rr.RawLine, 110))
				mismatches = append(mismatches, detail)
			}
		}

		totalCompared += compared
		totalMatched += matched
		totalMismatched += mismatched

		// Print PAR summary line
		status := "✅"
		statusDetail := fmt.Sprintf("%d/%d matched", matched, compared)
		if mismatched > 0 {
			status = "❌"
			statusDetail = fmt.Sprintf("%d/%d matched, %d MISMATCH", matched, compared, mismatched)
		} else if compared == 0 {
			status = "⚠ "
			statusDetail = "no confident rows matched by Bil+Channel"
		}
		parCompareResults[parCode] = statusDetail

		fmt.Printf("  [%s %-20s] %s  review=%d  confident=%d  noisy=%d  compared=%d  %s\n",
			parCode, ps.ParName, status, len(ps.Rows), len(confident), noisy, compared, statusDetail)

		// Show first few mismatches per PAR (limit detail noise)
		showMax := 5
		for j, m := range mismatches {
			if j >= showMax {
				fmt.Printf("    ... and %d more mismatches\n", len(mismatches)-showMax)
				break
			}
			fmt.Println(m)
		}
	}

	fmt.Println()
	fmt.Printf("  SUMMARY: pipe_rows_parsed=%d  confident=%d  compared=%d  matched=%d  mismatched=%d\n",
		totalPipeRows, totalConfident, totalCompared, totalMatched, totalMismatched)

	// ================================================================
	// PART 4: PAR-level totals from to-review.csv
	// ================================================================
	fmt.Println()
	printHeader("PART 4: PAR-Level Aggregate Totals (from to-review.csv)")
	fmt.Println()

	fmt.Printf("  %-7s %-22s %5s %8s %8s %6s %5s %8s  %-5s  %s\n",
		"PAR", "NAME", "ROWS", "A", "B", "C", "D", "B+C+D", "OK?", "vs Results")
	fmt.Println("  " + strings.Repeat("-", 115))

	grandA, grandB, grandC, grandD := 0, 0, 0, 0
	parFailCount := 0

	for pNum := 192; pNum <= 222; pNum++ {
		parCode := fmt.Sprintf("P.%d", pNum)
		ps := parMap[parCode]
		if ps == nil {
			fmt.Printf("  %-7s %-22s  NO DATA\n", parCode, "???")
			continue
		}

		bcd := ps.SumB + ps.SumC + ps.SumD
		ok := "✅"
		if ps.SumA != bcd {
			ok = fmt.Sprintf("❌ %+d", ps.SumA-bcd)
			parFailCount++
		}

		cmp := parCompareResults[parCode]
		if cmp == "" {
			cmp = "-"
		}

		fmt.Printf("  %-7s %-22s %5d %8d %8d %6d %5d %8d  %-5s  %s\n",
			ps.ParCode, truncate(ps.ParName, 22), ps.RowCount,
			ps.SumA, ps.SumB, ps.SumC, ps.SumD, bcd, ok, cmp)

		grandA += ps.SumA
		grandB += ps.SumB
		grandC += ps.SumC
		grandD += ps.SumD
	}

	fmt.Println("  " + strings.Repeat("-", 115))
	grandBCD := grandB + grandC + grandD
	grandOK := "✅"
	if grandA != grandBCD {
		grandOK = fmt.Sprintf("❌ %+d", grandA-grandBCD)
	}
	fmt.Printf("  %-7s %-22s %5d %8d %8d %6d %5d %8d  %s\n",
		"TOTAL", "", len(reviewRows), grandA, grandB, grandC, grandD, grandBCD, grandOK)

	fmt.Println()
	if parFailCount == 0 {
		fmt.Println("  ✅ All PARs satisfy aggregate A = B + C + D")
	} else {
		fmt.Printf("  ❌ %d PAR(s) fail aggregate A = B + C + D\n", parFailCount)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 90))
	fmt.Println("DONE")
}

// ---------------------------------------------------------------------------
// to-review.csv loading
// ---------------------------------------------------------------------------

func loadReviewCSV(path string) ([]ReviewRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) < 2 {
		return nil, fmt.Errorf("fewer than 2 rows")
	}

	header := records[0]
	slog.Debug("to-review.csv header", "cols", len(header))

	var rows []ReviewRow
	for i := 1; i < len(records); i++ {
		rec := records[i]
		if len(rec) <= colTotalUnreturned {
			slog.Warn("short row skipped", "line", i+1, "cols", len(rec))
			continue
		}

		rawA := strings.TrimSpace(rec[colTotalBallots])
		rawB := strings.TrimSpace(rec[colTotalValidVotes])
		rawC := strings.TrimSpace(rec[colTotalRejectedVotes])
		rawD := strings.TrimSpace(rec[colTotalUnreturned])

		rows = append(rows, ReviewRow{
			LineNum:        i + 1,
			UniqueCode:     strings.TrimSpace(rec[colUniqueCode]),
			BallotType:     strings.TrimSpace(rec[colBallotType]),
			ParCode:        strings.TrimSpace(rec[colParCode]),
			ParName:        strings.TrimSpace(rec[colParName]),
			DunCode:        strings.TrimSpace(rec[colDunCode]),
			DunName:        strings.TrimSpace(rec[colDunName]),
			PolDistCode:    strings.TrimSpace(rec[colPollingDistCode]),
			PolDistName:    strings.TrimSpace(rec[colPollingDistName]),
			PolCentre:      strings.TrimSpace(rec[colPollingCentre]),
			VotingChannel:  parseCleanInt(strings.TrimSpace(rec[colVotingChannel])),
			TotalBallots:   parseCleanInt(rawA),
			TotalValid:     parseCleanInt(rawB),
			TotalRejected:  parseCleanInt(rawC),
			TotalUnreturnd: parseCleanInt(rawD),
			RawA:           rawA,
			RawB:           rawB,
			RawC:           rawC,
			RawD:           rawD,
		})
	}
	return rows, nil
}

// assignBilNumbers assigns a Bil number to each review row within a PAR.
// Bil increments each time we see a new polling district code.
// The first row (typically POSTAL VOTE) gets Bil=1, next distinct district gets Bil=2, etc.
func assignBilNumbers(rows []ReviewRow) []int {
	bils := make([]int, len(rows))
	bil := 0
	lastDistCode := ""
	for i, r := range rows {
		dc := r.PolDistCode
		if dc != lastDistCode {
			bil++
			lastDistCode = dc
		}
		bils[i] = bil
	}
	return bils
}

// ---------------------------------------------------------------------------
// Results file parsing
// ---------------------------------------------------------------------------

func loadResultsFile(path string) ([]ResultRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var rows []ResultRow
	lineNum := 0
	lastBil := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			continue
		}
		// Skip separator rows (|---|---|...)
		if strings.Contains(line, "---") {
			continue
		}

		parts := splitPipe(line)
		if len(parts) < 7 {
			continue
		}

		// Detect JUMLAH row
		isJumlah := false
		for _, p := range parts {
			if strings.Contains(strings.ToUpper(p), "JUMLAH") {
				isJumlah = true
				break
			}
		}

		// Skip rows that are entirely empty or mostly blank
		nonEmptyCount := 0
		for _, p := range parts {
			if strings.TrimSpace(p) != "" {
				nonEmptyCount++
			}
		}
		if nonEmptyCount < 4 && !isJumlah {
			continue
		}

		// Parse fields — handle "N N" space-doubled OCR artifacts by taking first token
		bilRaw := parts[0]
		channelRaw := parts[3]
		aRaw := parts[4]

		nParts := len(parts)
		bRaw := parts[nParts-3]
		cRaw := parts[nParts-2]
		dRaw := parts[nParts-1]

		bil := parseOCRInt(bilRaw)
		channel := parseOCRInt(channelRaw)
		a := parseOCRInt(aRaw)
		b := parseOCRInt(bRaw)
		c := parseOCRInt(cRaw)
		d := parseOCRInt(dRaw)

		// Inherit Bil from previous row if blank (continuation row for same station)
		if bil == 0 && bilRaw == "" {
			bil = lastBil
		}
		if bil > 0 {
			lastBil = bil
		}

		// Parse candidate votes (between A and B)
		var candVotes []int
		for j := 5; j < nParts-3; j++ {
			candVotes = append(candVotes, parseOCRInt(parts[j]))
		}

		// Determine confidence: ALL numeric fields must be "clean" (no doubling artifacts)
		confident := true

		// Check A field
		if isDoubledOrGarbled(aRaw) {
			confident = false
		}
		// Check B field
		if isDoubledOrGarbled(bRaw) {
			confident = false
		}
		// Check C, D
		if isDoubledOrGarbled(cRaw) || isDoubledOrGarbled(dRaw) {
			confident = false
		}
		// Check channel
		if isDoubledOrGarbled(channelRaw) {
			confident = false
		}
		// Check Bil
		if isDoubledOrGarbled(bilRaw) {
			confident = false
		}
		// Check each candidate vote
		for j := 5; j < nParts-3; j++ {
			if isDoubledOrGarbled(parts[j]) {
				confident = false
				break
			}
		}

		// A must be > 0 for a confident data row (non-JUMLAH)
		if a == 0 && !isJumlah {
			confident = false
		}

		// Channel should be reasonable (0-20, or POS-like)
		if channel < 0 || channel > 20 {
			confident = false
		}

		// B should be <= A
		if a > 0 && b > a {
			confident = false
		}

		// Sum of candidate votes should equal B (if we have them)
		if confident && len(candVotes) > 0 && b > 0 {
			sumCand := 0
			for _, cv := range candVotes {
				sumCand += cv
			}
			// Allow sum=0 if some candidate fields are empty in OCR, but flag if sum > 0 and doesn't match
			if sumCand > 0 && sumCand != b {
				// Could be OCR noise in individual candidate fields but A/B/C/D are OK
				// Don't lose confidence for this — we mainly care about A/B/C/D
			}
		}

		// A = B + C + D should hold for confident rows
		if confident && a > 0 && a != b+c+d {
			confident = false
		}

		rr := ResultRow{
			Bil:        bil,
			Channel:    channel,
			A:          a,
			B:          b,
			C:          c,
			D:          d,
			CandVotes:  candVotes,
			RawLine:    line,
			Confident:  confident,
			IsJumlah:   isJumlah,
			LineInFile: lineNum,
		}
		rows = append(rows, rr)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return rows, nil
}

// splitPipe splits "|a|b|c|" into ["a","b","c"]
func splitPipe(line string) []string {
	line = strings.TrimPrefix(line, "|")
	line = strings.TrimSuffix(line, "|")
	parts := strings.Split(line, "|")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// ---------------------------------------------------------------------------
// Number parsing — OCR-aware
// ---------------------------------------------------------------------------

// parseCleanInt strips commas and quotes, then parses int. Returns 0 on failure.
func parseCleanInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "\"", "")
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// parseOCRInt handles OCR artifacts:
//   - "357 357" → 357 (space-doubled, take first token)
//   - "1 1" → 1
//   - "7 7" → 7
//   - "138" → 138 (clean)
//   - "" → 0
//   - "POS" → 0
//
// It does NOT attempt to fix concatenated doubles like "29298" — those are flagged as garbled.
func parseOCRInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	// If it contains spaces, it might be "N N" doubled: split and take first numeric token
	if strings.Contains(s, " ") {
		tokens := strings.Fields(s)
		// Check if all tokens are the same number (space-doubled)
		if len(tokens) == 2 && tokens[0] == tokens[1] {
			return parseCleanInt(tokens[0])
		}
		// Otherwise just try the first token that looks numeric
		for _, t := range tokens {
			v := parseCleanInt(t)
			if v > 0 {
				return v
			}
			// Handle "0" explicitly
			if t == "0" {
				return 0
			}
		}
		return 0
	}

	return parseCleanInt(s)
}

// isDoubledOrGarbled detects OCR artifacts in a numeric field.
// Returns true if the field appears corrupted.
func isDoubledOrGarbled(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	// Space-doubled like "357 357" or "7 7" — these are parseable, NOT garbled
	if strings.Contains(s, " ") {
		tokens := strings.Fields(s)
		if len(tokens) == 2 && tokens[0] == tokens[1] {
			return false // clean doubled display
		}
		// Two different tokens in a numeric field is suspicious
		allNumeric := true
		for _, t := range tokens {
			if !isNumericStr(t) {
				allNumeric = false
			}
		}
		if allNumeric && len(tokens) == 2 {
			// Two different numbers in one cell — garbled
			return true
		}
		// Non-numeric content (like "15 15" where tokens match is handled above)
		return false
	}

	// Pure numeric string: check for concatenated doubling
	if !isNumericStr(s) {
		return false // not a number at all, so not "doubled"
	}

	n := len(s)
	if n <= 3 {
		return false // 1-3 digits can't really be doubled
	}

	// Check for prefix-overlap pattern:
	// "29298" → prefix "29" matches start of suffix "298" → doubled
	// "32323" → prefix "323" matches start of suffix "23" — wait no...
	// Actually "32323": try split at 2 → prefix="32", suffix="323", "32"=="32"[0:2] of "323"? Yes.
	// "31316": split at 2 → "31" vs "316", "31" matches "31" in "316" → doubled
	// "12127": split at 2 → "12" vs "127", "12" matches "12" in "127" → doubled
	// "21219": split at 2 → "21" vs "219", "21" matches → doubled
	// "39393": split at 2 → "39" vs "393", "39" matches → doubled

	for splitAt := 1; splitAt < n; splitAt++ {
		prefix := s[:splitAt]
		suffix := s[splitAt:]

		if len(suffix) < len(prefix) {
			continue
		}

		// Does the suffix start with the prefix? (overlap pattern)
		if strings.HasPrefix(suffix, prefix) {
			// The "real" number would be the suffix
			suffixVal, err := strconv.Atoi(suffix)
			if err != nil {
				continue
			}
			fullVal, _ := strconv.Atoi(s)

			// Sanity: the full value should be significantly larger than the real value
			// and the real value should be a reasonable ballot count
			if suffixVal > 0 && fullVal > suffixVal*5 {
				return true
			}
		}
	}

	// Check repeating-digit patterns for 5-digit numbers
	// e.g., "39393" where d0==d2 && d1==d3 (interleaved doubling)
	if n == 5 {
		if s[0] == s[2] && s[1] == s[3] {
			return true
		}
	}

	return false
}

// isNumericStr checks if string is purely digits
func isNumericStr(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 4 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

func printHeader(title string) {
	fmt.Println(strings.Repeat("-", 90))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 90))
}
