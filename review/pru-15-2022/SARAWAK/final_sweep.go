// final_sweep.go — Full-file validation sweep for to-review.csv
//
// Checks:
//   1. UNIQUE CODE algorithm: rebuilt from component columns, compared to actual col 1
//   2. No duplicate UNIQUE CODEs remain
//   3. Column count consistency (every row = 92 cols)
//   4. VOTING CHANNEL NUMBER matches UNIQUE CODE suffix
//   5. Postal vote format: P.XXX_POS_1, PDC = XXX/POS
//   6. Early vote format: PDC ends /00, PDN = UNDI AWAL
//   7. BALLOT TYPE valid
//   8. STATE = SARAWAK
//   9. CHECK columns = 0
//  10. No comma-formatted numbers in numeric fields
//  11. BA'KELALAN apostrophe check
//  12. ZULHAIDAH name check
//
// Usage:
//   go run final_sweep.go

package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const inputFile = "to-review.csv"

// 0-based column indices
const (
	cUniqueCode    = 0
	cState         = 1
	cBallotType    = 2
	cParCode       = 3
	cParName       = 4
	cStateConstCod = 5
	cStateConstNam = 6
	cPollDistCode  = 7
	cPollDistName  = 8
	cPollCentre    = 9
	cChannel       = 20
	cTotalIssued   = 21
	cTotalValid    = 87
	cTotalReject   = 88
	cTotalUnret    = 89
	cCheckValid    = 90
	cCheckTotal    = 91
	expectedCols   = 92
)

type issue struct {
	Row   int
	Check string
	Msg   string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	f, err := os.Open(inputFile)
	if err != nil {
		slog.Error("open", "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	f.Close()
	if err != nil {
		slog.Error("read", "err", err)
		os.Exit(1)
	}
	slog.Info("read", "rows", len(records))

	// ── Pass 1: Build suffix maps per Polling District Code ─────────────
	// key = PDC, value = ordered list of distinct centres
	type distInfo struct {
		centres     []string
		centreToLtr map[string]string
	}
	districts := make(map[string]*distInfo)
	var distOrder []string

	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < expectedCols {
			continue
		}
		if row[cBallotType] == "POSTAL VOTE" {
			continue
		}
		pdc := strings.TrimSpace(row[cPollDistCode])
		centre := strings.TrimSpace(row[cPollCentre])

		di, ok := districts[pdc]
		if !ok {
			di = &distInfo{centreToLtr: make(map[string]string)}
			districts[pdc] = di
			distOrder = append(distOrder, pdc)
		}
		if _, seen := di.centreToLtr[centre]; !seen {
			ltr := string(rune('a' + len(di.centres)))
			di.centres = append(di.centres, centre)
			di.centreToLtr[centre] = ltr
		}
	}

	// ── Pass 2: Validate every row ──────────────────────────────────────
	var issues []issue
	addIssue := func(row int, check, msg string) {
		issues = append(issues, issue{Row: row, Check: check, Msg: msg})
	}

	codeSeen := make(map[string][]int) // UNIQUE CODE → list of 1-based row numbers

	for i, row := range records {
		if i == 0 {
			continue
		}
		lineNum := i + 1 // 1-based

		// ── Check 1: Column count ───────────────────────────────────
		if len(row) != expectedCols {
			addIssue(lineNum, "COL_COUNT", fmt.Sprintf("expected %d cols, got %d", expectedCols, len(row)))
			continue // can't reliably check other columns
		}

		// ── Check 2: STATE = SARAWAK ────────────────────────────────
		if row[cState] != "SARAWAK" {
			addIssue(lineNum, "STATE", fmt.Sprintf("expected SARAWAK, got %q", row[cState]))
		}

		// ── Check 3: BALLOT TYPE valid ──────────────────────────────
		bt := row[cBallotType]
		if bt != "POSTAL VOTE" && bt != "EARLY VOTE" && bt != "ORDINARY VOTE" {
			addIssue(lineNum, "BALLOT_TYPE", fmt.Sprintf("invalid: %q", bt))
		}

		// ── Check 4: Postal vote format ─────────────────────────────
		if bt == "POSTAL VOTE" {
			par := strings.TrimSpace(row[cParCode])
			parNum := ""
			if strings.HasPrefix(par, "P.") {
				parNum = par[2:]
			}
			wantUC := par + "_POS_1"
			wantSCC := par + "/POS"
			wantPDC := parNum + "/POS"

			if row[cUniqueCode] != wantUC {
				addIssue(lineNum, "POSTAL_UC", fmt.Sprintf("expected %q, got %q", wantUC, row[cUniqueCode]))
			}
			if row[cStateConstCod] != wantSCC {
				addIssue(lineNum, "POSTAL_SCC", fmt.Sprintf("expected %q, got %q", wantSCC, row[cStateConstCod]))
			}
			if row[cPollDistCode] != wantPDC {
				addIssue(lineNum, "POSTAL_PDC", fmt.Sprintf("expected %q, got %q", wantPDC, row[cPollDistCode]))
			}
			if strings.TrimSpace(row[cPollDistName]) != "UNDI POS" {
				addIssue(lineNum, "POSTAL_PDN", fmt.Sprintf("expected UNDI POS, got %q", row[cPollDistName]))
			}
			if strings.TrimSpace(row[cPollCentre]) != "UNDI POS" {
				addIssue(lineNum, "POSTAL_PC", fmt.Sprintf("expected UNDI POS, got %q", row[cPollCentre]))
			}
			if strings.TrimSpace(row[cChannel]) != "1" {
				addIssue(lineNum, "POSTAL_CH", fmt.Sprintf("expected channel 1, got %q", row[cChannel]))
			}
		}

		// ── Check 5: Early vote format ──────────────────────────────
		if bt == "EARLY VOTE" {
			pdc := strings.TrimSpace(row[cPollDistCode])
			if !strings.HasSuffix(pdc, "/00") {
				addIssue(lineNum, "EARLY_PDC", fmt.Sprintf("expected PDC ending /00, got %q", pdc))
			}
			pdn := strings.TrimSpace(row[cPollDistName])
			if pdn != "UNDI AWAL" {
				addIssue(lineNum, "EARLY_PDN", fmt.Sprintf("expected UNDI AWAL, got %q", pdn))
			}
		}

		// ── Check 6: UNIQUE CODE algorithm (non-postal) ─────────────
		if bt != "POSTAL VOTE" {
			par := strings.TrimSpace(row[cParCode])
			dun := strings.TrimSpace(row[cStateConstCod])
			pdc := strings.TrimSpace(row[cPollDistCode])
			ch := strings.TrimSpace(row[cChannel])
			centre := strings.TrimSpace(row[cPollCentre])

			di := districts[pdc]
			var wantUC string
			if di != nil && len(di.centres) > 1 {
				ltr := di.centreToLtr[centre]
				wantUC = fmt.Sprintf("%s_%s_%s_%s%s", par, dun, pdc, ch, ltr)
			} else {
				wantUC = fmt.Sprintf("%s_%s_%s_%s", par, dun, pdc, ch)
			}

			if row[cUniqueCode] != wantUC {
				addIssue(lineNum, "UNIQUE_CODE", fmt.Sprintf("expected %q, got %q", wantUC, row[cUniqueCode]))
			}
		}

		// ── Check 7: VOTING CHANNEL NUMBER matches UC suffix ────────
		if bt != "POSTAL VOTE" {
			uc := row[cUniqueCode]
			ch := strings.TrimSpace(row[cChannel])
			// Extract last segment after final '_'
			lastUnderscore := strings.LastIndex(uc, "_")
			if lastUnderscore >= 0 {
				suffix := uc[lastUnderscore+1:]
				// Channel is the leading digits of the suffix
				chFromUC := ""
				for _, c := range suffix {
					if c >= '0' && c <= '9' {
						chFromUC += string(c)
					} else {
						break
					}
				}
				if chFromUC != ch {
					addIssue(lineNum, "CHANNEL_MATCH", fmt.Sprintf("UC suffix %q → channel %q, but col 21 = %q", suffix, chFromUC, ch))
				}
			}
		}

		// ── Check 8: Comma-formatted numbers ────────────────────────
		numCols := []int{cTotalIssued, cTotalValid, cTotalReject, cTotalUnret}
		numNames := []string{"TOTAL_BALLOTS_ISSUED", "TOTAL_VALID_VOTES", "TOTAL_REJECTED_VOTES", "TOTAL_UNRETURNED_BALLOTS"}
		for j, col := range numCols {
			v := strings.TrimSpace(row[col])
			if v == "" {
				continue
			}
			if strings.Contains(v, ",") || strings.Contains(v, "\"") {
				addIssue(lineNum, "COMMA_NUM", fmt.Sprintf("%s contains comma/quote: %q", numNames[j], v))
			}
		}

		// ── Check 9: CHECK columns = 0 ──────────────────────────────
		for _, col := range []int{cCheckValid, cCheckTotal} {
			v := strings.TrimSpace(row[col])
			if v != "" && v != "0" {
				addIssue(lineNum, "CHECK_COL", fmt.Sprintf("col %d expected 0, got %q", col+1, v))
			}
		}

		// ── Check 10: BA`KELALAN backtick ───────────────────────────
		if strings.Contains(row[cStateConstNam], "`") {
			addIssue(lineNum, "BACKTICK", fmt.Sprintf("backtick found in STATE CONST NAME: %q", row[cStateConstNam]))
		}

		// ── Check 11: ZULBAIDAH (should be ZULHAIDAH) ───────────────
		for _, col := range []int{24, 29, 34, 39, 44, 49, 54, 59, 64, 69, 74, 79, 84} {
			if col < len(row) && strings.Contains(strings.ToUpper(row[col]), "ZULBAIDAH") {
				addIssue(lineNum, "NAME_TYPO", fmt.Sprintf("col %d still has ZULBAIDAH: %q", col+1, row[col]))
			}
		}

		// ── Track for duplicate check ───────────────────────────────
		code := row[cUniqueCode]
		if code != "" {
			codeSeen[code] = append(codeSeen[code], lineNum)
		}
	}

	// ── Check 12: Duplicate UNIQUE CODEs ────────────────────────────────
	dupCount := 0
	for code, lines := range codeSeen {
		if len(lines) > 1 {
			dupCount++
			addIssue(lines[0], "DUPLICATE_UC", fmt.Sprintf("%q appears %d times at lines %v", code, len(lines), lines))
		}
	}

	// ── Check 13: Verify numeric parsability of key fields ──────────────
	for i, row := range records {
		if i == 0 || len(row) < expectedCols {
			continue
		}
		lineNum := i + 1
		ch := strings.TrimSpace(row[cChannel])
		if ch != "" {
			if _, err := strconv.Atoi(ch); err != nil {
				addIssue(lineNum, "CHANNEL_NUM", fmt.Sprintf("non-numeric channel: %q", ch))
			}
		}
		ti := strings.TrimSpace(row[cTotalIssued])
		if ti != "" {
			if _, err := strconv.Atoi(ti); err != nil {
				addIssue(lineNum, "ISSUED_NUM", fmt.Sprintf("non-numeric TOTAL BALLOTS ISSUED: %q", ti))
			}
		}
	}

	// ── Report ──────────────────────────────────────────────────────────
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║         FINAL SWEEP — to-review.csv Validation           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("Total rows (incl header):  %d\n", len(records))
	fmt.Printf("Data rows:                 %d\n", len(records)-1)
	fmt.Printf("Distinct districts:        %d\n", len(districts))
	multiCount := 0
	for _, di := range districts {
		if len(di.centres) > 1 {
			multiCount++
		}
	}
	fmt.Printf("Multi-centre districts:    %d\n", multiCount)
	fmt.Printf("Unique UNIQUE CODEs:       %d\n", len(codeSeen))
	fmt.Printf("Duplicate UNIQUE CODEs:    %d\n", dupCount)
	fmt.Println()

	if len(issues) == 0 {
		fmt.Println("✅ ALL CHECKS PASSED — zero issues found.")
	} else {
		// Group by check
		byCheck := make(map[string][]issue)
		var checkOrder []string
		for _, iss := range issues {
			if _, ok := byCheck[iss.Check]; !ok {
				checkOrder = append(checkOrder, iss.Check)
			}
			byCheck[iss.Check] = append(byCheck[iss.Check], iss)
		}

		fmt.Printf("❌ ISSUES FOUND: %d total across %d check types\n\n", len(issues), len(checkOrder))
		for _, chk := range checkOrder {
			items := byCheck[chk]
			fmt.Printf("── %s (%d) ──\n", chk, len(items))
			limit := 10
			if len(items) < limit {
				limit = len(items)
			}
			for _, it := range items[:limit] {
				fmt.Printf("  Row %4d: %s\n", it.Row, it.Msg)
			}
			if len(items) > 10 {
				fmt.Printf("  ... and %d more\n", len(items)-10)
			}
			fmt.Println()
		}
	}
}
