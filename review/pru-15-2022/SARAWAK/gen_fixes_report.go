// gen_fixes_report.go — Compare original vs fixed to-review.csv and produce FIXES.md
//
// Usage:
//   go run gen_fixes_report.go
//
// Reads:
//   - Original file from git: `git show 6fb5fde:review/pru-15-2022/SARAWAK/to-review.csv`
//   - Current file: to-review.csv
//
// Writes:
//   - FIXES.md

package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

const (
	currentFile = "to-review.csv"
	outputFile  = "FIXES.md"
	gitRef      = "6fb5fde:review/pru-15-2022/SARAWAK/to-review.csv"
)

// Column indices (0-based)
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
	cTotalValid    = 87
)

// Candidate name columns (0-based): each party slot's candidate name column
var candidateNameCols = []int{23, 28, 33, 38, 43, 48, 53, 58, 63, 68, 73, 78, 83}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// ── 1. Read original from git ───────────────────────────────────────
	slog.Info("reading original from git", "ref", gitRef)
	cmd := exec.Command("git", "show", gitRef)
	origBytes, err := cmd.Output()
	if err != nil {
		slog.Error("git show failed", "err", err)
		os.Exit(1)
	}
	origRecords := parseCSV(string(origBytes), "original")
	slog.Info("original", "rows", len(origRecords))

	// ── 2. Read current file ────────────────────────────────────────────
	slog.Info("reading current file", "file", currentFile)
	curBytes, err := os.ReadFile(currentFile)
	if err != nil {
		slog.Error("read current", "err", err)
		os.Exit(1)
	}
	curRecords := parseCSV(string(curBytes), "current")
	slog.Info("current", "rows", len(curRecords))

	// ── 3. Compare row by row ───────────────────────────────────────────
	type fix struct {
		Row     int    // 1-based row in current file (including header)
		OldUC   string // original UNIQUE CODE
		NewUC   string // current UNIQUE CODE
		Changes []string
	}

	var fixes []fix

	// Use the shorter length
	maxRows := len(origRecords)
	if len(curRecords) < maxRows {
		maxRows = len(curRecords)
	}

	header := curRecords[0]

	for i := 1; i < maxRows; i++ {
		orig := origRecords[i]
		cur := curRecords[i]
		lineNum := i + 1

		var changes []string

		origLen := len(orig)
		curLen := len(cur)
		maxCols := origLen
		if curLen > maxCols {
			maxCols = curLen
		}

		for j := 0; j < maxCols; j++ {
			var ov, cv string
			if j < origLen {
				ov = orig[j]
			}
			if j < curLen {
				cv = cur[j]
			}
			if ov != cv {
				colName := fmt.Sprintf("Col %d", j+1)
				if j < len(header) {
					colName = header[j]
				}
				changes = append(changes, fmt.Sprintf("**%s**: `%s` → `%s`", colName, truncate(ov, 60), truncate(cv, 60)))
			}
		}

		if len(changes) > 0 {
			oldUC := ""
			newUC := ""
			if len(orig) > 0 {
				oldUC = orig[cUniqueCode]
			}
			if len(cur) > 0 {
				newUC = cur[cUniqueCode]
			}
			fixes = append(fixes, fix{
				Row:     lineNum,
				OldUC:   oldUC,
				NewUC:   newUC,
				Changes: changes,
			})
		}
	}

	// Check for extra rows in current
	if len(curRecords) > len(origRecords) {
		for i := len(origRecords); i < len(curRecords); i++ {
			fixes = append(fixes, fix{
				Row:     i + 1,
				NewUC:   curRecords[i][cUniqueCode],
				Changes: []string{"**NEW ROW**: added in current file"},
			})
		}
	}

	slog.Info("comparison done", "total_fixes", len(fixes))

	// ── 4. Categorize fixes ─────────────────────────────────────────────
	type category struct {
		Name  string
		Fixes []fix
	}

	var catPostal, catUniqueCode, catBakelalan, catCandidate, catCommaNum, catPollingCentre, catChannelFix []fix

	for _, f := range fixes {
		isPostal := false
		isUC := false
		isBak := false
		isCand := false
		isComma := false
		isPC := false
		isCh := false

		for _, c := range f.Changes {
			switch {
			case strings.Contains(c, "UNIQUE CODE") && !strings.Contains(c, "only"):
				isUC = true
			case strings.Contains(c, "BA`KELALAN") || strings.Contains(c, "BA'KELALAN"):
				isBak = true
			case strings.Contains(c, "ZULBAIDAH") || strings.Contains(c, "ZULHAIDAH"):
				isCand = true
			case strings.Contains(c, "1,616") || strings.Contains(c, "1,488") || strings.Contains(c, "\"1,"):
				isComma = true
			case strings.Contains(c, "POLLING DISTRICT CODE") && strings.Contains(c, "POS"):
				isPostal = true
			case strings.Contains(c, "STATE CONSTITUENCY CODE") && strings.Contains(c, "POS"):
				isPostal = true
			case strings.Contains(c, "POLLING CENTRE") && strings.Contains(c, "MUNGGU KOPI"):
				isPC = true
			case strings.Contains(c, "VOTING CHANNEL NUMBER"):
				isCh = true
			}
		}

		if isPostal {
			catPostal = append(catPostal, f)
		}
		if isUC {
			catUniqueCode = append(catUniqueCode, f)
		}
		if isBak {
			catBakelalan = append(catBakelalan, f)
		}
		if isCand {
			catCandidate = append(catCandidate, f)
		}
		if isComma {
			catCommaNum = append(catCommaNum, f)
		}
		if isPC {
			catPollingCentre = append(catPollingCentre, f)
		}
		if isCh {
			catChannelFix = append(catChannelFix, f)
		}
	}

	// ── 5. Write FIXES.md ───────────────────────────────────────────────
	out, err := os.Create(outputFile)
	if err != nil {
		slog.Error("create output", "err", err)
		os.Exit(1)
	}
	defer out.Close()

	w := func(format string, args ...any) {
		fmt.Fprintf(out, format+"\n", args...)
	}

	w("# FIXES REPORT — to-review.csv (PRU-15 2022 Sarawak)")
	w("")
	w("**Original**: `git show 6fb5fde` (pre-review)")
	w("**Fixed**: current `to-review.csv`")
	w("**Total rows changed**: %d out of %d data rows", len(fixes), len(curRecords)-1)
	w("")
	w("---")
	w("")
	w("## Summary of Fix Categories")
	w("")
	w("| # | Category | Rows Affected | Description |")
	w("|---|----------|--------------|-------------|")
	w("| 1 | Postal Vote Format | %d | UNIQUE CODE, STATE CONSTITUENCY CODE, POLLING DISTRICT CODE corrected to `P.XXX_POS_1` / `XXX/POS` pattern |", len(catPostal))
	w("| 2 | UNIQUE CODE Suffix | %d | Rebuilt all codes with correct `_CHANNELletter` suffix algorithm per Polling District Code |", len(catUniqueCode))
	w("| 3 | BA'KELALAN Typo | %d | Backtick (0x60) replaced with apostrophe (0x27) in STATE CONSTITUENCY NAME for N.81 |", len(catBakelalan))
	w("| 4 | Candidate Name | %d | `ZULBAIDAH SUBOH` → `ZULHAIDAH SUBOH` in P.218 SIBUTI (verified against official PDF) |", len(catCandidate))
	w("| 5 | Comma-Formatted Number | %d | Removed thousands separator from TOTAL VALID VOTES (`\"1,616\"` → `1616`, `\"1,488\"` → `1488`) |", len(catCommaNum))
	w("| 6 | Polling Centre Correction | %d | Corrected centre name from PDF verification (district 198/20/33) |", len(catPollingCentre))
	w("| 7 | Channel Number Correction | %d | Corrected VOTING CHANNEL NUMBER from PDF verification (district 213/58/08) |", len(catChannelFix))
	w("| | **TOTAL** | **%d** | |", len(fixes))
	w("")
	w("---")
	w("")

	// ── Detailed table ──────────────────────────────────────────────────
	w("## Detailed Row-by-Row Changes")
	w("")
	w("| Row | Current UNIQUE CODE | Changes |")
	w("|-----|-------------------|---------|")

	for _, f := range fixes {
		uc := f.NewUC
		if uc == "" {
			uc = f.OldUC
		}
		changeSummary := strings.Join(f.Changes, "<br>")
		w("| %d | `%s` | %s |", f.Row, truncate(uc, 45), changeSummary)
	}

	w("")
	w("---")
	w("")

	// ── Per-category details ────────────────────────────────────────────
	w("## Fix Details by Category")
	w("")

	// Category 1: Postal
	w("### 1. Postal Vote Format (%d rows)", len(catPostal))
	w("")
	w("All 31 parliamentary postal vote rows were corrected:")
	w("")
	w("| Field | Before | After |")
	w("|-------|--------|-------|")
	w("| UNIQUE CODE | `P.XXX_P.XXX/POSTAL VOTE_UNDI POS_1` | `P.XXX_POS_1` |")
	w("| STATE CONSTITUENCY CODE | `P.XXX/POSTAL VOTE` | `P.XXX/POS` |")
	w("| POLLING DISTRICT CODE | `UNDI POS` | `XXX/POS` |")
	w("")

	// Category 2: UNIQUE CODE
	w("### 2. UNIQUE CODE Suffix Algorithm (%d rows)", len(catUniqueCode))
	w("")
	w("All UNIQUE CODEs were rebuilt from component columns using the correct suffix algorithm:")
	w("")
	w("- **Format**: `P.XXX_N.YY_PDC_CHANNELletter` (e.g., `_1a`, `_2b`)")
	w("- Suffix letter assigned per Polling District Code based on first-appearance order of Polling Centres")
	w("- Same Polling Centre → same letter across all channels")
	w("- Single-centre districts get no suffix")
	w("")
	w("Sample changes:")
	w("")
	w("| Row | Before | After | Reason |")
	w("|-----|--------|-------|--------|")
	shown := 0
	for _, f := range catUniqueCode {
		if shown >= 15 {
			w("| ... | ... | ... | (%d more rows) |", len(catUniqueCode)-15)
			break
		}
		w("| %d | `%s` | `%s` | Suffix rebuild |", f.Row, truncate(f.OldUC, 40), truncate(f.NewUC, 40))
		shown++
	}
	w("")

	// Category 3: BA'KELALAN
	w("### 3. BA'KELALAN Character Fix (%d rows)", len(catBakelalan))
	w("")
	w("All N.81 rows: backtick `` ` `` (0x60) → apostrophe `'` (0x27) in column 7 (STATE CONSTITUENCY NAME).")
	w("")

	// Category 4: Candidate name
	w("### 4. Candidate Name Correction (%d rows)", len(catCandidate))
	w("")
	w("P.218 SIBUTI, PH/PKR candidate: `ZULBAIDAH SUBOH` → `ZULHAIDAH SUBOH`")
	w("")
	w("**Evidence**: Official PDF election results list the candidate as **ZULHAIDAH SUBOH**.")
	w("")

	// Category 5: Comma numbers
	w("### 5. Comma-Formatted Numbers (%d rows)", len(catCommaNum))
	w("")
	w("| Row | UNIQUE CODE | Before | After |")
	w("|-----|-------------|--------|-------|")
	for _, f := range catCommaNum {
		for _, c := range f.Changes {
			if strings.Contains(c, "1,616") || strings.Contains(c, "1,488") || strings.Contains(c, "TOTAL VALID") {
				w("| %d | `%s` | %s |", f.Row, f.NewUC, c)
			}
		}
	}
	w("")

	// Category 6: Polling Centre
	w("### 6. Polling Centre Correction (%d rows)", len(catPollingCentre))
	w("")
	w("District `198/20/33`, row 1260: Polling Centre corrected after PDF verification.")
	w("")
	w("- **Before**: `SEKOLAH KEBANGSAAN TANAH PUTEH`")
	w("- **After**: `RH. PANJANG, KPG. MUNGGU KOPI`")
	w("")
	w("This was a data entry error — the third row in this district belongs to a different polling centre.")
	w("")

	// Category 7: Channel fix
	w("### 7. Channel Number Correction (%d rows)", len(catChannelFix))
	w("")
	w("District `213/58/08`, row 2922: VOTING CHANNEL NUMBER corrected after PDF verification.")
	w("")
	w("- **Before**: channel `1` (caused duplicate with row 2921)")
	w("- **After**: channel `2` (confirmed from official PDF)")
	w("")

	w("---")
	w("")
	w("*Report generated by `gen_fixes_report.go`*")

	slog.Info("report written", "file", outputFile, "rows_changed", len(fixes))
}

func parseCSV(data string, label string) [][]string {
	r := csv.NewReader(strings.NewReader(data))
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		slog.Error("parse CSV", "label", label, "err", err)
		os.Exit(1)
	}
	return records
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
