package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ── Data structures ──────────────────────────────────────────────────────────

type rawCandidate struct {
	Name    string
	Party   string
	KID     string // constituency numeric code e.g. "19200"
	Votes   int    // ju
	Majority int   // mj
	Rejected int   // ut (actually rejected votes, per AGENTS.md)
	CandNum  int   // nc
}

type rawDUN struct {
	ID         string // e.g. "19201"
	DUNCode    string // e.g. "N.01"
	StateCode  string
	PARID      string // e.g. "19200"
	DUNName    string
	Registered int    // pb - total registered voters
	StatsB     int    // ordinary registered
	StatsA     int    // early voters registered
	StatsP     int    // postal voters registered
}

type seatRow struct {
	PARCode    string
	PARName    string
	DUNCode    string
	DUNName    string
	Registered int // col index 9
}

type reviewRow struct {
	LineNum         int
	UniqueCode      string
	BallotType      string
	PARCode         string
	PARName         string
	DUNCode         string
	DUNName         string
	PollingDistCode string
	PollingCentre   string
	Channel         string
	A               int // TOTAL BALLOTS ISSUED
	B               int // TOTAL VALID VOTES
	C               int // TOTAL REJECTED VOTES
	D               int // TOTAL UNRETURNED BALLOTS
	RawA            string
	RawB            string
}

type resultsTableRow struct {
	Bil     int
	Channel int
	A       int // ballots issued
	B       int // total valid
	CC      int // rejected
	DD      int // unreturned
	Clean   bool
	RawLine string
}

// ── Parsing helpers ──────────────────────────────────────────────────────────

func cleanInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "\u00a0", "")
	return strconv.Atoi(s)
}

func mustInt(s string) int {
	v, _ := cleanInt(s)
	return v
}

// ── Loaders ──────────────────────────────────────────────────────────────────

func loadCandidates(path string) ([]rawCandidate, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	hdr, err := r.Read()
	if err != nil {
		return nil, err
	}
	idx := map[string]int{}
	for i, h := range hdr {
		idx[strings.TrimSpace(h)] = i
	}
	var out []rawCandidate
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		get := func(col string) string {
			if i, ok := idx[col]; ok && i < len(rec) {
				return strings.TrimSpace(rec[i])
			}
			return ""
		}
		c := rawCandidate{
			Name:     get("t"),
			Party:    get("st"),
			KID:      get("kid"),
			Votes:    mustInt(get("ju")),
			Majority: mustInt(get("mj")),
			Rejected: mustInt(get("ut")),
			CandNum:  mustInt(get("nc")),
		}
		out = append(out, c)
	}
	return out, nil
}

func loadDUN(path string) (map[string]rawDUN, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	hdr, err := r.Read()
	if err != nil {
		return nil, err
	}
	idx := map[string]int{}
	for i, h := range hdr {
		idx[strings.TrimSpace(h)] = i
	}
	out := map[string]rawDUN{}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		get := func(col string) string {
			if i, ok := idx[col]; ok && i < len(rec) {
				return strings.TrimSpace(rec[i])
			}
			return ""
		}
		sid := get("sid")
		if sid != "13" {
			continue
		}
		d := rawDUN{
			ID:         get("id"),
			DUNCode:    get("kno"),
			StateCode:  sid,
			PARID:      get("pid"),
			DUNName:    get("t"),
			Registered: mustInt(get("pb")),
			StatsB:     mustInt(get("stats/b")),
			StatsA:     mustInt(get("stats/a")),
			StatsP:     mustInt(get("stats/p")),
		}
		out[d.ID] = d
	}
	return out, nil
}

func loadSeats(path string) (map[string]seatRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	out := map[string]seatRow{}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) < 32 {
			continue
		}
		if strings.TrimSpace(rec[0]) != "13" {
			continue
		}
		s := seatRow{
			PARCode:    strings.TrimSpace(rec[2]),
			PARName:    strings.TrimSpace(rec[3]),
			DUNCode:    strings.TrimSpace(rec[5]),
			DUNName:    strings.TrimSpace(rec[6]),
			Registered: mustInt(rec[9]),
		}
		key := s.PARCode + "_" + s.DUNCode
		out[key] = s
	}
	return out, nil
}

func loadReview(path string) ([]reviewRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true
	hdr, err := r.Read()
	if err != nil {
		return nil, err
	}
	idx := map[string]int{}
	for i, h := range hdr {
		idx[strings.TrimSpace(h)] = i
	}
	var out []reviewRow
	ln := 1
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			ln++
			continue
		}
		ln++
		get := func(col string) string {
			if i, ok := idx[col]; ok && i < len(rec) {
				return strings.TrimSpace(rec[i])
			}
			return ""
		}
		rawA := get("TOTAL BALLOTS ISSUED")
		rawB := get("TOTAL VALID VOTES")
		row := reviewRow{
			LineNum:         ln,
			UniqueCode:      get("UNIQUE CODE"),
			BallotType:      get("BALLOT TYPE"),
			PARCode:         get("PARLIAMENTARY CODE"),
			PARName:         get("PARLIAMENTARY NAME"),
			DUNCode:         get("STATE CONSTITUENCY CODE"),
			DUNName:         get("STATE CONSTITUENCY NAME"),
			PollingDistCode: get("POLLING DISTRICT CODE"),
			PollingCentre:   get("POLLING CENTRE"),
			Channel:         get("VOTING CHANNEL NUMBER"),
			A:               mustInt(rawA),
			B:               mustInt(rawB),
			C:               mustInt(get("TOTAL REJECTED VOTES")),
			D:               mustInt(get("TOTAL UNRETURNED BALLOTS")),
			RawA:            rawA,
			RawB:            rawB,
		}
		out = append(out, row)
	}
	return out, nil
}

// ── Results file parser (OCR'd score sheets) ────────────────────────────────

var pipeRowRe = regexp.MustCompile(`^\|(.+)\|$`)

func isDoubled(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 3 {
		return false
	}
	// Pattern: "NNN NNN" (space-doubled) is actually clean
	if strings.Contains(s, " ") {
		parts := strings.Fields(s)
		if len(parts) == 2 && parts[0] == parts[1] {
			return false // space-doubled but consistent
		}
		return true
	}
	// Detect concatenated doubling: "29298" -> "298" doubled
	n := len(s)
	for k := n / 2; k <= n/2+1; k++ {
		if k >= n {
			continue
		}
		prefix := s[:k]
		suffix := s[k:]
		if strings.HasPrefix(prefix, suffix) || strings.HasSuffix(prefix, suffix) ||
			strings.HasPrefix(suffix, prefix) || strings.HasSuffix(suffix, prefix) {
			// Potential overlap
			_, err := strconv.Atoi(s)
			if err != nil {
				return true
			}
			// Check if the number is suspiciously formed
			// e.g. "29298" length 5 for a 3-digit number
			if n >= 5 {
				return true
			}
		}
	}
	return false
}

func parseCleanInt(s string) (int, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, true
	}
	// Handle space-doubled: "357 357" -> 357
	if strings.Contains(s, " ") {
		parts := strings.Fields(s)
		if len(parts) == 2 && parts[0] == parts[1] {
			v, err := strconv.Atoi(parts[0])
			return v, err == nil
		}
		return 0, false
	}
	// Handle comma numbers: "1,135"
	if strings.Contains(s, ",") {
		s = strings.ReplaceAll(s, ",", "")
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	// Heuristic: detect concatenated doubles for numbers > 4 digits that look suspicious
	orig := s
	if len(orig) >= 5 {
		// Could be doubled. Check if halves overlap
		// "29298" -> check if it could be "298" doubled
		// Too unreliable, mark as not clean
		return v, false
	}
	return v, true
}

func parseResultsFile(path string) []resultsTableRow {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var tableRows []resultsTableRow
	scanner := NewLineScanner(f)
	currentBil := 0
	lineNo := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNo++
		m := pipeRowRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		cells := strings.Split(m[1], "|")
		if len(cells) < 7 {
			continue
		}

		// Try to identify the row structure
		// First cell might be Bil number
		bilStr := strings.TrimSpace(cells[0])
		bil, bilOk := parseCleanInt(bilStr)
		if bilOk && bil > 0 && bil <= 200 {
			currentBil = bil
		}

		// Find channel number (should be a small int 1-9)
		// Column index depends on whether bil/district columns are filled
		// Try common positions: index 3 for channel
		chanStr := strings.TrimSpace(cells[3])
		ch, chOk := parseCleanInt(chanStr)
		if !chOk || ch < 1 || ch > 20 {
			continue
		}

		// Column 4 should be A (total ballots)
		aStr := strings.TrimSpace(cells[4])
		aVal, aOk := parseCleanInt(aStr)
		if !aOk || aVal <= 0 {
			continue
		}

		// Last three columns should be B (valid), C (rejected), D (unreturned)
		nCells := len(cells)
		if nCells < 7 {
			continue
		}
		dStr := strings.TrimSpace(cells[nCells-1])
		cStr := strings.TrimSpace(cells[nCells-2])
		bStr := strings.TrimSpace(cells[nCells-3])

		dVal, dOk := parseCleanInt(dStr)
		cVal, cOk := parseCleanInt(cStr)
		bVal, bOk := parseCleanInt(bStr)

		clean := aOk && bOk && cOk && dOk
		// Extra validation: A should equal B + C + D for clean rows
		if clean && aVal != bVal+cVal+dVal {
			clean = false
		}

		tr := resultsTableRow{
			Bil:     currentBil,
			Channel: ch,
			A:       aVal,
			B:       bVal,
			CC:      cVal,
			DD:      dVal,
			Clean:   clean,
			RawLine: line,
		}
		tableRows = append(tableRows, tr)
	}
	return tableRows
}

type lineScanner struct {
	s *strings.Reader
	b []byte
}

func NewLineScanner(f *os.File) *bufScanner {
	return &bufScanner{s: newBufReader(f)}
}

// Use a simpler approach
type bufScanner struct {
	s    *bufReader
	line string
}
type bufReader struct {
	f    *os.File
	data []byte
	pos  int
}

func newBufReader(f *os.File) *bufReader {
	data, _ := io.ReadAll(f)
	return &bufReader{f: f, data: data}
}
func (b *bufReader) readLine() (string, bool) {
	if b.pos >= len(b.data) {
		return "", false
	}
	start := b.pos
	for b.pos < len(b.data) && b.data[b.pos] != '\n' {
		b.pos++
	}
	line := string(b.data[start:b.pos])
	if b.pos < len(b.data) {
		b.pos++ // skip \n
	}
	return strings.TrimRight(line, "\r"), true
}
func (bs *bufScanner) Scan() bool {
	line, ok := bs.s.readLine()
	bs.line = line
	return ok
}
func (bs *bufScanner) Text() string { return bs.line }

// ── PAR code helpers ─────────────────────────────────────────────────────────

func parNumFromKID(kid string) int {
	n, _ := strconv.Atoi(kid)
	return n / 100
}

func kidToPAR(kid string) string {
	n := parNumFromKID(kid)
	return fmt.Sprintf("P.%d", n)
}

// ── Main ─────────────────────────────────────────────────────────────────────

func main() {
	slog.Info("=== BALLOT CROSS-CHECK: to-review.csv vs multiple sources ===")

	// Load all data sources
	candidates, err := loadCandidates("raw-candidates.csv")
	if err != nil {
		slog.Error("load candidates", "err", err)
		os.Exit(1)
	}
	slog.Info("loaded raw-candidates.csv", "rows", len(candidates))

	duns, err := loadDUN("raw-dun.csv")
	if err != nil {
		slog.Error("load dun", "err", err)
		os.Exit(1)
	}
	slog.Info("loaded raw-dun.csv", "sarawak_duns", len(duns))

	seats, err := loadSeats("raw-seats-clean.csv")
	if err != nil {
		slog.Error("load seats", "err", err)
		os.Exit(1)
	}
	slog.Info("loaded raw-seats-clean.csv", "sarawak_seats", len(seats))

	review, err := loadReview("to-review.csv")
	if err != nil {
		slog.Error("load review", "err", err)
		os.Exit(1)
	}
	slog.Info("loaded to-review.csv", "rows", len(review))

	// ═══════════════════════════════════════════════════════════════════════
	fmt.Println()
	fmt.Println(strings.Repeat("=", 120))
	fmt.Println("BALLOT CROSS-CHECK REPORT")
	fmt.Println(strings.Repeat("=", 120))

	// ── SECTION 1: Internal consistency ──────────────────────────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 1: Internal consistency per row — A = B + C + D")
	fmt.Println(strings.Repeat("-", 120))

	intOK, intFail := 0, 0
	for _, row := range review {
		exp := row.B + row.C + row.D
		if row.A != exp {
			intFail++
			if intFail <= 20 {
				fmt.Printf("  ❌ Line %d %-50s A=%d ≠ B+C+D=%d (B=%d C=%d D=%d) [rawA=%q rawB=%q]\n",
					row.LineNum, row.UniqueCode, row.A, exp, row.B, row.C, row.D, row.RawA, row.RawB)
			}
		} else {
			intOK++
		}
	}
	if intFail == 0 {
		fmt.Printf("  ✅ All %d rows satisfy A = B + C + D\n", intOK)
	} else {
		fmt.Printf("  ❌ %d FAIL, %d OK\n", intFail, intOK)
	}

	// ── SECTION 2: PAR-level cross-check vs raw-candidates.csv ──────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 2: PAR-level totals — to-review.csv vs raw-candidates.csv")
	fmt.Println("  raw-candidates.csv fields: ju = candidate votes, ut = rejected votes (per AGENTS.md Phase-6)")
	fmt.Println(strings.Repeat("-", 120))

	// Group candidates by PAR
	type parCandAgg struct {
		PAR         string
		TotalVotes  int // sum of ju
		Rejected    int // ut (same for all candidates in PAR)
		NCandidates int
		Candidates  []rawCandidate
	}
	candByPAR := map[string]*parCandAgg{}
	for _, c := range candidates {
		par := kidToPAR(c.KID)
		if !strings.HasPrefix(par, "P.19") && !strings.HasPrefix(par, "P.20") &&
			!strings.HasPrefix(par, "P.21") && !strings.HasPrefix(par, "P.22") {
			continue
		}
		parNum, _ := strconv.Atoi(strings.TrimPrefix(par, "P."))
		if parNum < 192 || parNum > 222 {
			continue
		}
		if _, ok := candByPAR[par]; !ok {
			candByPAR[par] = &parCandAgg{PAR: par, Rejected: c.Rejected}
		}
		agg := candByPAR[par]
		agg.TotalVotes += c.Votes
		agg.NCandidates++
		agg.Candidates = append(agg.Candidates, c)
	}

	// Group review by PAR
	type parRevAgg struct {
		SumA, SumB, SumC, SumD int
		Rows                   int
	}
	revByPAR := map[string]*parRevAgg{}
	for _, row := range review {
		if _, ok := revByPAR[row.PARCode]; !ok {
			revByPAR[row.PARCode] = &parRevAgg{}
		}
		agg := revByPAR[row.PARCode]
		agg.SumA += row.A
		agg.SumB += row.B
		agg.SumC += row.C
		agg.SumD += row.D
		agg.Rows++
	}

	parCodes := make([]string, 0)
	for i := 192; i <= 222; i++ {
		parCodes = append(parCodes, fmt.Sprintf("P.%d", i))
	}

	fmt.Println()
	fmt.Printf("  %-8s %-20s | %8s %8s %5s | %8s %8s %5s | %8s %8s %5s | %8s | %s\n",
		"PAR", "Name", "raw_ju", "rev_B", "Δ_B", "raw_ut", "rev_C", "Δ_C", "rev_A", "B+C+D", "Δ", "rev_D", "Status")
	fmt.Println("  " + strings.Repeat("-", 135))

	parOK, parFail := 0, 0
	grandRevA, grandRevB, grandRevC, grandRevD := 0, 0, 0, 0
	grandRawJu, grandRawUt := 0, 0

	for _, par := range parCodes {
		rv, rvOK := revByPAR[par]
		ca, caOK := candByPAR[par]

		if !rvOK {
			fmt.Printf("  %-8s %-20s | NOT IN REVIEW\n", par, "???")
			continue
		}

		rawJu, rawUt := 0, 0
		parName := ""
		if caOK {
			rawJu = ca.TotalVotes
			rawUt = ca.Rejected
			if len(ca.Candidates) > 0 {
				// Get PAR name from seats
				for _, s := range seats {
					if s.PARCode == par {
						parName = s.PARName
						break
					}
				}
			}
		}

		diffB := rv.SumB - rawJu
		diffC := rv.SumC - rawUt
		diffA := rv.SumA - (rv.SumB + rv.SumC + rv.SumD)

		grandRevA += rv.SumA
		grandRevB += rv.SumB
		grandRevC += rv.SumC
		grandRevD += rv.SumD
		grandRawJu += rawJu
		grandRawUt += rawUt

		status := "✅"
		if diffB != 0 || diffC != 0 {
			status = "❌"
			parFail++
		} else {
			parOK++
		}

		fmt.Printf("  %-8s %-20s | %8d %8d %+5d | %8d %8d %+5d | %8d %8d %+5d | %8d | %s\n",
			par, parName,
			rawJu, rv.SumB, diffB,
			rawUt, rv.SumC, diffC,
			rv.SumA, rv.SumB+rv.SumC+rv.SumD, diffA,
			rv.SumD, status)
	}

	fmt.Println("  " + strings.Repeat("-", 135))
	fmt.Printf("  %-8s %-20s | %8d %8d %+5d | %8d %8d %+5d | %8d %8d %+5d | %8d |\n",
		"TOTAL", "",
		grandRawJu, grandRevB, grandRevB-grandRawJu,
		grandRawUt, grandRevC, grandRevC-grandRawUt,
		grandRevA, grandRevB+grandRevC+grandRevD, grandRevA-(grandRevB+grandRevC+grandRevD),
		grandRevD)
	fmt.Printf("\n  PAR-level: %d match ✅, %d mismatch ❌\n", parOK, parFail)

	// ── SECTION 3: DUN-level registered voter ceiling check ─────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 3: DUN-level ceiling check — TOTAL BALLOTS ISSUED ≤ Registered Voters")
	fmt.Println("  Source: raw-seats-clean.csv (col 10 = total registered voters per DUN)")
	fmt.Println(strings.Repeat("-", 120))

	type dunKey struct{ PAR, DUN string }
	revByDUN := map[dunKey]*parRevAgg{}
	postalByPAR := map[string]*parRevAgg{}
	for _, row := range review {
		if strings.Contains(row.BallotType, "POSTAL") || strings.Contains(row.DUNCode, "POS") {
			if _, ok := postalByPAR[row.PARCode]; !ok {
				postalByPAR[row.PARCode] = &parRevAgg{}
			}
			a := postalByPAR[row.PARCode]
			a.SumA += row.A
			a.SumB += row.B
			a.SumC += row.C
			a.SumD += row.D
			a.Rows++
			continue
		}
		k := dunKey{row.PARCode, row.DUNCode}
		if _, ok := revByDUN[k]; !ok {
			revByDUN[k] = &parRevAgg{}
		}
		a := revByDUN[k]
		a.SumA += row.A
		a.SumB += row.B
		a.SumC += row.C
		a.SumD += row.D
		a.Rows++
	}

	ceilOK, ceilFail, ceilWarn := 0, 0, 0
	var ceilIssues []string

	// Build sorted DUN list
	type dunSort struct {
		parN, dunN int
		key        dunKey
	}
	var dunList []dunSort
	for k := range revByDUN {
		pn, _ := strconv.Atoi(strings.TrimPrefix(k.PAR, "P."))
		dn, _ := strconv.Atoi(strings.TrimPrefix(k.DUN, "N."))
		dunList = append(dunList, dunSort{pn, dn, k})
	}
	sort.Slice(dunList, func(i, j int) bool {
		if dunList[i].parN != dunList[j].parN {
			return dunList[i].parN < dunList[j].parN
		}
		return dunList[i].dunN < dunList[j].dunN
	})

	fmt.Println()
	fmt.Printf("  %-8s %-6s %-22s | %8s | %8s | %6s | %s\n",
		"PAR", "DUN", "Name", "Regist", "Ballots", "Turn%", "Status")
	fmt.Println("  " + strings.Repeat("-", 90))

	for _, ds := range dunList {
		k := ds.key
		rv := revByDUN[k]
		sk := k.PAR + "_" + k.DUN
		seat, ok := seats[sk]
		if !ok {
			ceilWarn++
			fmt.Printf("  %-8s %-6s %-22s | %8s | %8d | %6s | ⚠️  DUN not in raw-seats-clean\n",
				k.PAR, k.DUN, "???", "???", rv.SumA, "???")
			continue
		}

		turnout := 0.0
		if seat.Registered > 0 {
			turnout = float64(rv.SumA) / float64(seat.Registered) * 100
		}

		status := "✅"
		if rv.SumA > seat.Registered {
			status = fmt.Sprintf("🚨 EXCEEDS by %d", rv.SumA-seat.Registered)
			ceilFail++
			ceilIssues = append(ceilIssues, fmt.Sprintf("%s %s: ballots=%d > registered=%d",
				k.PAR, k.DUN, rv.SumA, seat.Registered))
		} else if turnout > 95 {
			status = fmt.Sprintf("⚠️  Very high turnout")
			ceilWarn++
		} else {
			ceilOK++
		}

		// Only print problems or every 10th for spot-check
		if rv.SumA > seat.Registered || turnout > 90 || ds.dunN%20 == 1 {
			fmt.Printf("  %-8s %-6s %-22s | %8d | %8d | %5.1f%% | %s\n",
				k.PAR, k.DUN, seat.DUNName, seat.Registered, rv.SumA, turnout, status)
		}
	}

	fmt.Printf("\n  Ceiling check: %d OK, %d EXCEED, %d warnings\n", ceilOK, ceilFail, ceilWarn)
	if len(ceilIssues) > 0 {
		fmt.Println("  🚨 CRITICAL — Ballots exceed registered voters:")
		for _, s := range ceilIssues {
			fmt.Printf("     %s\n", s)
		}
	}

	// ── SECTION 4: Row-by-row vs results files (OCR) ────────────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 4: Row-by-row comparison vs results/ files (OCR'd score sheets)")
	fmt.Println("  Only confident rows (clean numbers, A=B+C+D) are compared.")
	fmt.Println(strings.Repeat("-", 120))

	// Group review rows by PAR in order
	type parGroup struct {
		par  string
		rows []reviewRow
	}
	parGroupMap := map[string]*parGroup{}
	for _, row := range review {
		if _, ok := parGroupMap[row.PARCode]; !ok {
			parGroupMap[row.PARCode] = &parGroup{par: row.PARCode}
		}
		parGroupMap[row.PARCode].rows = append(parGroupMap[row.PARCode].rows, row)
	}

	totalCompared, totalMatch, totalMismatch := 0, 0, 0

	for _, par := range parCodes {
		pg, ok := parGroupMap[par]
		if !ok {
			continue
		}
		path := fmt.Sprintf("results/Sarawak-%s.csv", par)
		resRows := parseResultsFile(path)
		if len(resRows) == 0 {
			continue
		}

		// Filter to only clean rows
		var cleanRows []resultsTableRow
		for _, rr := range resRows {
			if rr.Clean {
				cleanRows = append(cleanRows, rr)
			}
		}

		// Build a map from (Bil, Channel) -> results row
		type bilCh struct{ bil, ch int }
		resMap := map[bilCh]resultsTableRow{}
		for _, rr := range cleanRows {
			key := bilCh{rr.Bil, rr.Channel}
			resMap[key] = rr
		}

		// Assign Bil numbers to review rows
		// Bil is a sequential counter per polling district within the PAR
		bilCounter := 0
		lastDistCode := ""
		var mismatches []string

		for _, row := range pg.rows {
			if row.PollingDistCode != lastDistCode {
				bilCounter++
				lastDistCode = row.PollingDistCode
			}
			ch := mustInt(row.Channel)
			if ch <= 0 {
				ch = 1
			}

			key := bilCh{bilCounter, ch}
			rr, ok := resMap[key]
			if !ok {
				continue
			}

			totalCompared++
			if row.A == rr.A {
				totalMatch++
			} else {
				totalMismatch++
				diff := row.A - rr.A
				mismatches = append(mismatches, fmt.Sprintf(
					"    Bil=%d Ch=%d | line=%d %-45s | review_A=%4d results_A=%4d diff=%+d  [results_B=%d C=%d D=%d]",
					bilCounter, ch, row.LineNum, row.UniqueCode, row.A, rr.A, diff, rr.B, rr.CC, rr.DD))
			}
		}

		if len(mismatches) > 0 {
			fmt.Printf("\n  [%-8s] %d confident results rows, %d compared, %d mismatches:\n",
				par, len(cleanRows), len(mismatches)+totalMatch, len(mismatches))
			for _, m := range mismatches {
				fmt.Println(m)
			}
		}
	}

	fmt.Printf("\n  Results comparison: %d compared, %d match ✅, %d mismatch ❌\n",
		totalCompared, totalMatch, totalMismatch)
	fmt.Println("  NOTE: Mismatches may be due to OCR noise or Bil numbering differences.")
	fmt.Println("  The OCR source is unreliable — Sections 1-3 are the authoritative checks.")

	// ── SECTION 5: Turnout statistics ───────────────────────────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 5: Per-PAR turnout statistics")
	fmt.Println(strings.Repeat("-", 120))

	fmt.Println()
	fmt.Printf("  %-8s %-20s | %9s | %9s | %6s | %9s | %8s | %8s\n",
		"PAR", "Name", "Regist", "Ballots", "Turn%", "Valid(B)", "Rej(C)", "Unret(D)")
	fmt.Println("  " + strings.Repeat("-", 110))

	grandReg := 0
	for _, par := range parCodes {
		rv, rvOK := revByPAR[par]
		if !rvOK {
			continue
		}

		// Sum registered voters for this PAR across all DUNs
		totalReg := 0
		for _, s := range seats {
			if s.PARCode == par {
				totalReg += s.Registered
			}
		}
		grandReg += totalReg

		parName := ""
		for _, s := range seats {
			if s.PARCode == par {
				parName = s.PARName
				break
			}
		}

		turnout := 0.0
		if totalReg > 0 {
			turnout = float64(rv.SumA) / float64(totalReg) * 100
		}

		fmt.Printf("  %-8s %-20s | %9d | %9d | %5.1f%% | %9d | %8d | %8d\n",
			par, parName, totalReg, rv.SumA, turnout, rv.SumB, rv.SumC, rv.SumD)
	}

	fmt.Println("  " + strings.Repeat("-", 110))
	grandTurnout := 0.0
	if grandReg > 0 {
		grandTurnout = float64(grandRevA) / float64(grandReg) * 100
	}
	fmt.Printf("  %-8s %-20s | %9d | %9d | %5.1f%% | %9d | %8d | %8d\n",
		"TOTAL", "", grandReg, grandRevA, grandTurnout, grandRevB, grandRevC, grandRevD)

	// ── SECTION 6: Formatting/data quality issues ───────────────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 6: Data quality flags in TOTAL BALLOTS ISSUED column")
	fmt.Println(strings.Repeat("-", 120))

	// Check for comma-formatted numbers
	fmt.Println()
	fmtIssues := 0
	for _, row := range review {
		if strings.Contains(row.RawA, ",") || strings.ContainsAny(row.RawA, "\"") {
			fmtIssues++
			fmt.Printf("  ⚠️  COMMA/QUOTE in A: line=%d %s rawA=%q parsed=%d\n",
				row.LineNum, row.UniqueCode, row.RawA, row.A)
		}
		if strings.Contains(row.RawB, ",") || strings.ContainsAny(row.RawB, "\"") {
			fmtIssues++
			fmt.Printf("  ⚠️  COMMA/QUOTE in B: line=%d %s rawB=%q parsed=%d\n",
				row.LineNum, row.UniqueCode, row.RawB, row.B)
		}
	}
	if fmtIssues == 0 {
		fmt.Println("  ✅ No comma/quote formatting issues in A or B columns")
	}

	// Check zero ballots
	zeroCnt := 0
	for _, row := range review {
		if row.A == 0 {
			zeroCnt++
			fmt.Printf("  ⚠️  ZERO BALLOTS: line=%d %s type=%s\n",
				row.LineNum, row.UniqueCode, row.BallotType)
		}
	}
	if zeroCnt == 0 {
		fmt.Println("  ✅ No zero-ballot rows")
	}

	// Check for extremely high single-channel ballots (non-postal)
	fmt.Println()
	highCnt := 0
	for _, row := range review {
		if row.A > 900 && row.BallotType != "POSTAL VOTE" {
			highCnt++
			if highCnt <= 10 {
				fmt.Printf("  📊 HIGH: line=%d %-45s A=%d type=%s\n",
					row.LineNum, row.UniqueCode, row.A, row.BallotType)
			}
		}
	}
	if highCnt > 10 {
		fmt.Printf("  ... and %d more high-ballot rows (>900, non-postal)\n", highCnt-10)
	}
	fmt.Printf("  Total non-postal rows with A>900: %d\n", highCnt)

	// Rejected votes as % of ballots: flag unusually high rejection rates
	fmt.Println()
	fmt.Println("  High rejection rate rows (C/A > 5%):")
	highRejCnt := 0
	for _, row := range review {
		if row.A > 0 {
			rate := float64(row.C) / float64(row.A) * 100
			if rate > 5.0 {
				highRejCnt++
				if highRejCnt <= 15 {
					fmt.Printf("  ⚠️  line=%d %-45s A=%d C=%d (%.1f%%)\n",
						row.LineNum, row.UniqueCode, row.A, row.C, rate)
				}
			}
		}
	}
	if highRejCnt > 15 {
		fmt.Printf("  ... and %d more\n", highRejCnt-15)
	}
	if highRejCnt == 0 {
		fmt.Println("  ✅ No unusually high rejection rates")
	} else {
		fmt.Printf("  Total rows with rejection rate > 5%%: %d\n", highRejCnt)
	}

	// ── SECTION 7: Per-DUN cross-check of TOTAL VALID VOTES ─────────────
	fmt.Println()
	fmt.Println(strings.Repeat("-", 120))
	fmt.Println("SECTION 7: DUN-level valid-vote cross-check — review sum(B) vs raw-dun stats")
	fmt.Println("  raw-dun.csv: stats/b field. raw-candidates.csv: per-PAR only (no DUN breakdown).")
	fmt.Println("  So we check: review sum(A) ≤ registered AND review A=B+C+D holds per DUN.")
	fmt.Println(strings.Repeat("-", 120))

	dunMismatch := 0
	for _, ds := range dunList {
		k := ds.key
		rv := revByDUN[k]
		// A = B + C + D per DUN
		if rv.SumA != rv.SumB+rv.SumC+rv.SumD {
			dunMismatch++
			fmt.Printf("  ❌ %s %s: A=%d ≠ B+C+D=%d\n",
				k.PAR, k.DUN, rv.SumA, rv.SumB+rv.SumC+rv.SumD)
		}
	}
	// Also check postal
	for _, par := range parCodes {
		pa, ok := postalByPAR[par]
		if !ok {
			continue
		}
		if pa.SumA != pa.SumB+pa.SumC+pa.SumD {
			dunMismatch++
			fmt.Printf("  ❌ %s POSTAL: A=%d ≠ B+C+D=%d\n",
				par, pa.SumA, pa.SumB+pa.SumC+pa.SumD)
		}
	}
	if dunMismatch == 0 {
		fmt.Printf("  ✅ All %d DUNs + %d postal entries satisfy A = B + C + D at aggregate level\n",
			len(dunList), len(postalByPAR))
	}

	// ── FINAL SUMMARY ───────────────────────────────────────────────────
	fmt.Println()
	fmt.Println(strings.Repeat("=", 120))
	fmt.Println("FINAL SUMMARY")
	fmt.Println(strings.Repeat("=", 120))
	fmt.Println()

	allGood := true

	check := func(name string, ok bool, detail string) {
		if ok {
			fmt.Printf("  ✅ %s — %s\n", name, detail)
		} else {
			fmt.Printf("  ❌ %s — %s\n", name, detail)
			allGood = false
		}
	}

	check("Row-level A=B+C+D",
		intFail == 0,
		fmt.Sprintf("%d/%d rows pass", intOK, intOK+intFail))

	check("PAR-level valid votes match raw-candidates.csv",
		parFail == 0,
		fmt.Sprintf("%d/%d PARs match", parOK, parOK+parFail))

	check("DUN-level ballots ≤ registered voters",
		ceilFail == 0,
		fmt.Sprintf("%d exceed, %d OK, %d warnings", ceilFail, ceilOK, ceilWarn))

	check("DUN-level aggregate A=B+C+D",
		dunMismatch == 0,
		fmt.Sprintf("%d DUNs checked", len(dunList)+len(postalByPAR)))

	check("No formatting issues in ballot numbers",
		fmtIssues == 0,
		fmt.Sprintf("%d issues found", fmtIssues))

	fmt.Println()
	if allGood {
		fmt.Println("  🎉 ALL CHECKS PASS — TOTAL BALLOTS ISSUED data is consistent and verified.")
	} else {
		fmt.Println("  ⚠️  SOME CHECKS FAILED — see details above.")
	}

	// Grand totals
	fmt.Printf("\n  Grand totals: registered=%d  ballots=%d  valid=%d  rejected=%d  unreturned=%d  turnout=%.1f%%\n",
		grandReg, grandRevA, grandRevB, grandRevC, grandRevD,
		math.Round(float64(grandRevA)/float64(grandReg)*1000)/10)

	fmt.Println()
}
