package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// partyGroup defines a block of 5 columns for each party/candidate
type partyGroup struct {
	name     string
	partyCol int // 0-based index for the party column
	candCol  int
	sexCol   int
	ageCol   int
	voteCol  int
}

var partyGroups = []partyGroup{
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

// violation records
type violation struct {
	row     int // 1-based line in CSV (header=1)
	col     int // 1-based column number
	colName string
	value   string
	msg     string
}

type sexAgeInconsistency struct {
	dun       string // e.g. N.01
	party     string
	candidate string
	field     string // SEX or AGE
	filledCnt int
	emptyCnt  int
	examples  []string // sample row numbers
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	csvPath := "to-review.csv"

	f, err := os.Open(csvPath)
	if err != nil {
		slog.Error("failed to open CSV", "path", csvPath, "error", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // allow variable

	allRecords, err := reader.ReadAll()
	if err != nil {
		slog.Error("failed to read CSV", "error", err)
		os.Exit(1)
	}

	if len(allRecords) < 2 {
		slog.Error("CSV has no data rows")
		os.Exit(1)
	}

	header := allRecords[0]
	slog.Info("CSV loaded", "totalRows", len(allRecords)-1, "headerCols", len(header))

	expectedCols := 67
	dataRows := allRecords[1:]

	// ---------- RULE 1: Column count ----------
	var colCountViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2 // 1-based, header is row 1
		if len(row) != expectedCols {
			colCountViolations = append(colCountViolations, violation{
				row:   rowNum,
				msg:   fmt.Sprintf("Expected %d columns, got %d", expectedCols, len(row)),
				value: safeGet(row, 0),
			})
		}
	}
	slog.Info("Rule 1 (column count) done", "violations", len(colCountViolations))

	// ---------- RULE 2: STATE = SARAWAK ----------
	var stateViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 2 {
			continue
		}
		state := strings.TrimSpace(row[1])
		if state != "SARAWAK" {
			stateViolations = append(stateViolations, violation{
				row:     rowNum,
				col:     2,
				colName: "STATE",
				value:   state,
				msg:     fmt.Sprintf("Expected 'SARAWAK', got '%s'", state),
			})
		}
	}
	slog.Info("Rule 2 (STATE) done", "violations", len(stateViolations))

	// ---------- RULE 3: BALLOT TYPE ----------
	validBallotTypes := map[string]bool{
		"POSTAL VOTE":   true,
		"EARLY VOTE":    true,
		"ORDINARY VOTE": true,
	}
	var ballotTypeViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 3 {
			continue
		}
		bt := strings.TrimSpace(row[2])
		if !validBallotTypes[bt] {
			ballotTypeViolations = append(ballotTypeViolations, violation{
				row:     rowNum,
				col:     3,
				colName: "BALLOT TYPE",
				value:   bt,
				msg:     fmt.Sprintf("Invalid ballot type: '%s'", bt),
			})
		}
	}
	slog.Info("Rule 3 (BALLOT TYPE) done", "violations", len(ballotTypeViolations))

	// ---------- RULE 4: Postal vote rules ----------
	type postalViolation struct {
		row     int
		subRule string
		detail  string
	}
	var postalViolations []postalViolation
	// Also count postal rows per DUN
	postalPerDUN := make(map[string][]int) // DUN code -> row numbers

	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		bt := strings.TrimSpace(row[2])
		if bt != "POSTAL VOTE" {
			continue
		}
		dunCode := strings.TrimSpace(row[5])
		postalPerDUN[dunCode] = append(postalPerDUN[dunCode], rowNum)

		// Check POLLING DISTRICT CODE ends with /POS or /UNDI POS
		pdc := strings.TrimSpace(row[7])
		if !strings.HasSuffix(pdc, "/POS") && !strings.HasSuffix(pdc, "/UNDI POS") {
			postalViolations = append(postalViolations, postalViolation{
				row:     rowNum,
				subRule: "POLLING DISTRICT CODE suffix",
				detail:  fmt.Sprintf("Expected ending '/POS' or '/UNDI POS', got '%s'", pdc),
			})
		}

		// Check POLLING DISTRICT NAME = UNDI POS
		pdn := strings.TrimSpace(row[8])
		if pdn != "UNDI POS" {
			postalViolations = append(postalViolations, postalViolation{
				row:     rowNum,
				subRule: "POLLING DISTRICT NAME",
				detail:  fmt.Sprintf("Expected 'UNDI POS', got '%s'", pdn),
			})
		}

		// Check POLLING CENTRE = UNDI POS
		pc := strings.TrimSpace(row[9])
		if pc != "UNDI POS" {
			postalViolations = append(postalViolations, postalViolation{
				row:     rowNum,
				subRule: "POLLING CENTRE",
				detail:  fmt.Sprintf("Expected 'UNDI POS', got '%s'", pc),
			})
		}

		// Check UNIQUE CODE contains POS
		uc := strings.TrimSpace(row[0])
		if !strings.Contains(uc, "POS") {
			postalViolations = append(postalViolations, postalViolation{
				row:     rowNum,
				subRule: "UNIQUE CODE contains POS",
				detail:  fmt.Sprintf("UNIQUE CODE '%s' does not contain 'POS'", uc),
			})
		}
	}
	// Check >1 postal rows per DUN
	var multiPostalDUNs []struct {
		dun  string
		rows []int
	}
	for dun, rows := range postalPerDUN {
		if len(rows) > 1 {
			multiPostalDUNs = append(multiPostalDUNs, struct {
				dun  string
				rows []int
			}{dun, rows})
		}
	}
	sort.Slice(multiPostalDUNs, func(i, j int) bool {
		return multiPostalDUNs[i].dun < multiPostalDUNs[j].dun
	})
	slog.Info("Rule 4 (postal vote) done", "violations", len(postalViolations), "multiPostalDUNs", len(multiPostalDUNs))

	// ---------- RULE 5: Early vote rules ----------
	type earlyViolation struct {
		row     int
		subRule string
		detail  string
	}
	var earlyViolations []earlyViolation

	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		bt := strings.TrimSpace(row[2])
		if bt != "EARLY VOTE" {
			continue
		}

		// Check POLLING DISTRICT CODE ends with /00
		pdc := strings.TrimSpace(row[7])
		if !strings.HasSuffix(pdc, "/00") {
			earlyViolations = append(earlyViolations, earlyViolation{
				row:     rowNum,
				subRule: "POLLING DISTRICT CODE suffix",
				detail:  fmt.Sprintf("Expected ending '/00', got '%s'", pdc),
			})
		}

		// Check POLLING DISTRICT NAME = UNDI AWAL
		pdn := strings.TrimSpace(row[8])
		if pdn != "UNDI AWAL" {
			earlyViolations = append(earlyViolations, earlyViolation{
				row:     rowNum,
				subRule: "POLLING DISTRICT NAME",
				detail:  fmt.Sprintf("Expected 'UNDI AWAL', got '%s'", pdn),
			})
		}

		// Check UNIQUE CODE contains /00_
		uc := strings.TrimSpace(row[0])
		if !strings.Contains(uc, "/00_") {
			earlyViolations = append(earlyViolations, earlyViolation{
				row:     rowNum,
				subRule: "UNIQUE CODE contains /00_",
				detail:  fmt.Sprintf("UNIQUE CODE '%s' does not contain '/00_'", uc),
			})
		}
	}
	slog.Info("Rule 5 (early vote) done", "violations", len(earlyViolations))

	// ---------- RULE 6: VOTING CHANNEL NUMBER consistency ----------
	var channelViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		uc := strings.TrimSpace(row[0])
		vcn := strings.TrimSpace(row[10])
		if uc == "" || vcn == "" {
			continue
		}
		// Extract the part after the last underscore in UNIQUE CODE
		lastUnderscore := strings.LastIndex(uc, "_")
		if lastUnderscore >= 0 && lastUnderscore < len(uc)-1 {
			ucChannel := uc[lastUnderscore+1:]
			if ucChannel != vcn {
				channelViolations = append(channelViolations, violation{
					row:     rowNum,
					col:     11,
					colName: "VOTING CHANNEL NUMBER",
					value:   vcn,
					msg:     fmt.Sprintf("Channel from UNIQUE CODE '%s' is '%s', but VOTING CHANNEL NUMBER is '%s'", uc, ucChannel, vcn),
				})
			}
		}
	}
	slog.Info("Rule 6 (channel number) done", "violations", len(channelViolations))

	// ---------- RULE 7: SEX + AGE consistency per DUN/candidate ----------
	// For each DUN + party group + candidate, track whether SEX/AGE is filled/empty
	type candKey struct {
		dun       string
		party     string
		candidate string
	}
	type fillStats struct {
		sexFilled  int
		sexEmpty   int
		ageFilled  int
		ageEmpty   int
		sexExRows  []int
		ageExRows  []int
		invalidSex []violation
		invalidAge []violation
	}
	candFillMap := make(map[candKey]*fillStats)

	var sexValueViolations []violation
	var ageValueViolations []violation

	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 62 {
			continue
		}
		dunCode := strings.TrimSpace(safeGet(row, 5))

		for _, pg := range partyGroups {
			if pg.candCol >= len(row) || pg.sexCol >= len(row) || pg.ageCol >= len(row) {
				continue
			}
			candName := strings.TrimSpace(row[pg.candCol])
			if candName == "" {
				continue
			}
			sexVal := strings.TrimSpace(row[pg.sexCol])
			ageVal := strings.TrimSpace(row[pg.ageCol])

			key := candKey{dun: dunCode, party: pg.name, candidate: candName}
			stats, ok := candFillMap[key]
			if !ok {
				stats = &fillStats{}
				candFillMap[key] = stats
			}

			if sexVal != "" {
				stats.sexFilled++
				// Validate sex value
				if sexVal != "MALE" && sexVal != "FEMALE" {
					v := violation{
						row:     rowNum,
						col:     pg.sexCol + 1,
						colName: pg.name + " CANDIDATE SEX",
						value:   sexVal,
						msg:     fmt.Sprintf("Invalid SEX value '%s' for candidate '%s'", sexVal, candName),
					}
					stats.invalidSex = append(stats.invalidSex, v)
					sexValueViolations = append(sexValueViolations, v)
				}
			} else {
				stats.sexEmpty++
				if len(stats.sexExRows) < 5 {
					stats.sexExRows = append(stats.sexExRows, rowNum)
				}
			}

			if ageVal != "" {
				stats.ageFilled++
				// Validate age is numeric
				if _, err := strconv.Atoi(ageVal); err != nil {
					v := violation{
						row:     rowNum,
						col:     pg.ageCol + 1,
						colName: pg.name + " CANDIDATE AGE",
						value:   ageVal,
						msg:     fmt.Sprintf("Non-numeric AGE value '%s' for candidate '%s'", ageVal, candName),
					}
					stats.invalidAge = append(stats.invalidAge, v)
					ageValueViolations = append(ageValueViolations, v)
				}
			} else {
				stats.ageEmpty++
				if len(stats.ageExRows) < 5 {
					stats.ageExRows = append(stats.ageExRows, rowNum)
				}
			}
		}
	}

	// Find inconsistencies: same candidate in same DUN has mix of filled/empty
	var sexAgeIssues []sexAgeInconsistency
	for key, stats := range candFillMap {
		if stats.sexFilled > 0 && stats.sexEmpty > 0 {
			sexAgeIssues = append(sexAgeIssues, sexAgeInconsistency{
				dun:       key.dun,
				party:     key.party,
				candidate: key.candidate,
				field:     "SEX",
				filledCnt: stats.sexFilled,
				emptyCnt:  stats.sexEmpty,
				examples:  intsToStrings(stats.sexExRows),
			})
		}
		if stats.ageFilled > 0 && stats.ageEmpty > 0 {
			sexAgeIssues = append(sexAgeIssues, sexAgeInconsistency{
				dun:       key.dun,
				party:     key.party,
				candidate: key.candidate,
				field:     "AGE",
				filledCnt: stats.ageFilled,
				emptyCnt:  stats.ageEmpty,
				examples:  intsToStrings(stats.ageExRows),
			})
		}
	}
	sort.Slice(sexAgeIssues, func(i, j int) bool {
		if sexAgeIssues[i].dun != sexAgeIssues[j].dun {
			return dunSortKey(sexAgeIssues[i].dun) < dunSortKey(sexAgeIssues[j].dun)
		}
		if sexAgeIssues[i].party != sexAgeIssues[j].party {
			return sexAgeIssues[i].party < sexAgeIssues[j].party
		}
		return sexAgeIssues[i].field < sexAgeIssues[j].field
	})
	slog.Info("Rule 7 (SEX/AGE) done", "inconsistencies", len(sexAgeIssues), "invalidSex", len(sexValueViolations), "invalidAge", len(ageValueViolations))

	// ---------- RULE 8: Numeric columns ----------
	numericColIndices := []int{11, 16, 21, 26, 31, 36, 41, 46, 51, 56, 61, 62, 63, 64} // 0-based
	numericColNames := []string{
		"TOTAL BALLOTS ISSUED", "GPS VOTE", "PH VOTE", "PSB VOTE", "PBK VOTE",
		"ASPIRASI VOTE", "PBDSB VOTE", "SEDAR VOTE", "PAS VOTE",
		"INDEPENDENT 1 VOTE", "INDEPENDENT 2 VOTE",
		"TOTAL VALID VOTES", "TOTAL REJECTED VOTES", "TOTAL UNRETURNED BALLOTS",
	}
	var numericViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		for k, colIdx := range numericColIndices {
			if colIdx >= len(row) {
				continue
			}
			val := strings.TrimSpace(row[colIdx])
			if val == "" {
				// Empty is acceptable for vote columns where no candidate
				continue
			}
			if _, err := strconv.Atoi(val); err != nil {
				numericViolations = append(numericViolations, violation{
					row:     rowNum,
					col:     colIdx + 1,
					colName: numericColNames[k],
					value:   val,
					msg:     fmt.Sprintf("Non-numeric value '%s' in %s", val, numericColNames[k]),
				})
			}
		}
	}
	slog.Info("Rule 8 (numeric) done", "violations", len(numericViolations))

	// ---------- RULE 9: CHECKER columns ----------
	type checkerViolation struct {
		row       int
		checkType string // VALID VOTE or TOTAL VOTE ISSUED
		expected  string
		actual    string
		detail    string
	}
	var checkerViolations []checkerViolation

	voteColIndices := []int{16, 21, 26, 31, 36, 41, 46, 51, 56, 61} // 0-based vote columns

	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 67 {
			continue
		}

		// CHECKER (VALID VOTE) = col 66 (index 65)
		// Sum of individual votes should = TOTAL VALID VOTES
		totalValidStr := strings.TrimSpace(row[62])
		checkerValidStr := strings.TrimSpace(row[65])
		totalValid, errTV := strconv.Atoi(totalValidStr)

		if errTV == nil {
			sumVotes := 0
			allParsed := true
			for _, vi := range voteColIndices {
				if vi >= len(row) {
					continue
				}
				v := strings.TrimSpace(row[vi])
				if v == "" {
					continue
				}
				n, err := strconv.Atoi(v)
				if err != nil {
					allParsed = false
					break
				}
				sumVotes += n
			}
			if allParsed {
				expectedChecker := "1"
				if sumVotes != totalValid {
					expectedChecker = "0"
				}
				if checkerValidStr != "" && checkerValidStr != expectedChecker {
					checkerViolations = append(checkerViolations, checkerViolation{
						row:       rowNum,
						checkType: "VALID VOTE",
						expected:  expectedChecker,
						actual:    checkerValidStr,
						detail:    fmt.Sprintf("Sum of candidate votes=%d, TOTAL VALID VOTES=%d, checker=%s", sumVotes, totalValid, checkerValidStr),
					})
				}
				// Also flag when sum != total (regardless of checker value)
				if sumVotes != totalValid {
					checkerViolations = append(checkerViolations, checkerViolation{
						row:       rowNum,
						checkType: "VALID VOTE MISMATCH",
						expected:  strconv.Itoa(sumVotes),
						actual:    totalValidStr,
						detail:    fmt.Sprintf("Sum of candidate votes (%d) != TOTAL VALID VOTES (%d)", sumVotes, totalValid),
					})
				}
			}
		}

		// CHECKER (TOTAL VOTE ISSUED) = col 67 (index 66)
		// TOTAL BALLOTS ISSUED = TOTAL VALID VOTES + TOTAL REJECTED VOTES + TOTAL UNRETURNED BALLOTS
		totalBallotsStr := strings.TrimSpace(row[11])
		totalRejectedStr := strings.TrimSpace(row[63])
		totalUnreturnedStr := strings.TrimSpace(row[64])
		checkerTotalStr := strings.TrimSpace(row[66])

		totalBallots, errTB := strconv.Atoi(totalBallotsStr)
		totalRejected, errTR := strconv.Atoi(totalRejectedStr)
		totalUnreturned, errTU := strconv.Atoi(totalUnreturnedStr)

		if errTB == nil && errTV == nil && errTR == nil && errTU == nil {
			computedTotal := totalValid + totalRejected + totalUnreturned
			expectedChecker := "1"
			if computedTotal != totalBallots {
				expectedChecker = "0"
			}
			if checkerTotalStr != "" && checkerTotalStr != expectedChecker {
				checkerViolations = append(checkerViolations, checkerViolation{
					row:       rowNum,
					checkType: "TOTAL VOTE ISSUED",
					expected:  expectedChecker,
					actual:    checkerTotalStr,
					detail:    fmt.Sprintf("TOTAL BALLOTS ISSUED=%d, Valid+Rejected+Unreturned=%d+%d+%d=%d, checker=%s", totalBallots, totalValid, totalRejected, totalUnreturned, computedTotal, checkerTotalStr),
				})
			}
			if computedTotal != totalBallots {
				checkerViolations = append(checkerViolations, checkerViolation{
					row:       rowNum,
					checkType: "TOTAL VOTE ISSUED MISMATCH",
					expected:  strconv.Itoa(totalBallots),
					actual:    strconv.Itoa(computedTotal),
					detail:    fmt.Sprintf("TOTAL BALLOTS ISSUED (%d) != Valid(%d)+Rejected(%d)+Unreturned(%d) = %d", totalBallots, totalValid, totalRejected, totalUnreturned, computedTotal),
				})
			}
		}
	}
	slog.Info("Rule 9 (CHECKER) done", "violations", len(checkerViolations))

	// ---------- RULE 10: UNIQUE CODE structure ----------
	ucPattern := regexp.MustCompile(`^P\.(\d+)_N\.(\d+)_(\d+)/(\d+)/(.+)_(\d+)$`)
	var ucViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		uc := strings.TrimSpace(row[0])
		if uc == "" {
			ucViolations = append(ucViolations, violation{
				row: rowNum,
				msg: "Empty UNIQUE CODE",
			})
			continue
		}

		matches := ucPattern.FindStringSubmatch(uc)
		if matches == nil {
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("UNIQUE CODE '%s' does not match pattern P.XXX_N.YY_XXX/YY/ZZ_C", uc),
			})
			continue
		}

		// matches: [full, parlNum, dunNum, parlNum2, dunNum2, suffix, channel]
		parlNum := matches[1]
		dunNum := matches[2]
		parlNum2 := matches[3]
		dunNum2 := matches[4]
		channel := matches[6]

		// Validate P code matches
		parlCode := strings.TrimSpace(row[3]) // e.g. P.192
		expectedParlNum := strings.TrimPrefix(parlCode, "P.")
		if parlNum != expectedParlNum {
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("P-code in UNIQUE CODE (%s) doesn't match PARLIAMENTARY CONSTITUENCY CODE (%s)", parlNum, parlCode),
			})
		}

		// Validate DUN code matches
		dunCodeCol := strings.TrimSpace(row[5]) // e.g. N.01
		expectedDunNum := strings.TrimPrefix(dunCodeCol, "N.")
		if dunNum != expectedDunNum {
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("DUN code in UNIQUE CODE (%s) doesn't match STATE CONSTITUENCY CODE (%s)", dunNum, dunCodeCol),
			})
		}

		// Validate inner P and DUN codes match outer
		if parlNum != parlNum2 {
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("P-code mismatch within UNIQUE CODE: P.%s vs %s/", parlNum, parlNum2),
			})
		}
		if dunNum != dunNum2 {
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("DUN code mismatch within UNIQUE CODE: N.%s vs /%s/", dunNum, dunNum2),
			})
		}

		// Validate channel matches VOTING CHANNEL NUMBER
		vcn := strings.TrimSpace(row[10])
		if channel != vcn {
			// already caught in rule 6, but flag here too for completeness
			ucViolations = append(ucViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("Channel in UNIQUE CODE (%s) doesn't match VOTING CHANNEL NUMBER (%s)", channel, vcn),
			})
		}
	}
	slog.Info("Rule 10 (UNIQUE CODE structure) done", "violations", len(ucViolations))

	// ---------- Additional: Check POLLING DISTRICT CODE in UNIQUE CODE ----------
	// The UNIQUE CODE should embed the POLLING DISTRICT CODE with slashes
	var pdcUCViolations []violation
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		uc := strings.TrimSpace(row[0])
		pdc := strings.TrimSpace(row[7])
		if uc == "" || pdc == "" {
			continue
		}
		// The UNIQUE CODE format is P.XXX_N.YY_PDC_Channel
		// PDC is between the second underscore and the last underscore
		parts := strings.SplitN(uc, "_", 3)
		if len(parts) < 3 {
			continue
		}
		remainder := parts[2] // e.g. "192/01/POS_1" or "192/01/01_1"
		lastUS := strings.LastIndex(remainder, "_")
		if lastUS < 0 {
			continue
		}
		ucPDC := remainder[:lastUS] // e.g. "192/01/POS" or "192/01/01"
		if ucPDC != pdc {
			pdcUCViolations = append(pdcUCViolations, violation{
				row:   rowNum,
				col:   1,
				value: uc,
				msg:   fmt.Sprintf("POLLING DISTRICT CODE in UNIQUE CODE ('%s') doesn't match col 8 ('%s')", ucPDC, pdc),
			})
		}
	}
	slog.Info("Additional: PDC in UNIQUE CODE check done", "violations", len(pdcUCViolations))

	// ---------- Additional: Verify postal DMKOD consistency (POS vs UNDI POS) ----------
	type postalFormatInfo struct {
		dun    string
		format string // "POS" or "UNDI POS"
		row    int
		uc     string
	}
	var postalFormats []postalFormatInfo
	for i, row := range dataRows {
		rowNum := i + 2
		if len(row) < 11 {
			continue
		}
		bt := strings.TrimSpace(row[2])
		if bt != "POSTAL VOTE" {
			continue
		}
		pdc := strings.TrimSpace(row[7])
		dunCode := strings.TrimSpace(row[5])
		uc := strings.TrimSpace(row[0])
		format := "unknown"
		if strings.HasSuffix(pdc, "/POS") {
			format = "POS"
		} else if strings.HasSuffix(pdc, "/UNDI POS") {
			format = "UNDI POS"
		}
		postalFormats = append(postalFormats, postalFormatInfo{
			dun:    dunCode,
			format: format,
			row:    rowNum,
			uc:     uc,
		})
	}

	posFormatCount := make(map[string]int)
	for _, pf := range postalFormats {
		posFormatCount[pf.format]++
	}
	slog.Info("Postal format distribution", "counts", posFormatCount)

	// ========== GENERATE REPORT ==========
	var sb strings.Builder

	sb.WriteString("# PHASE 5 REVIEW: Column Consistency and Mapping Rules\n\n")
	sb.WriteString("## Executive Summary\n\n")

	totalViolations := len(colCountViolations) + len(stateViolations) + len(ballotTypeViolations) +
		len(postalViolations) + len(multiPostalDUNs) + len(earlyViolations) +
		len(channelViolations) + len(sexAgeIssues) + len(sexValueViolations) +
		len(ageValueViolations) + len(numericViolations) + len(checkerViolations) +
		len(ucViolations) + len(pdcUCViolations)

	sb.WriteString(fmt.Sprintf("Total data rows analyzed: **%d**\n\n", len(dataRows)))

	sb.WriteString("| Rule | Description | Result | Violations |\n")
	sb.WriteString("| --- | --- | --- | --- |\n")
	sb.WriteString(fmt.Sprintf("| 1 | Column count (expected %d) | %s | %d |\n", expectedCols, passOrFail(len(colCountViolations)), len(colCountViolations)))
	sb.WriteString(fmt.Sprintf("| 2 | STATE = SARAWAK | %s | %d |\n", passOrFail(len(stateViolations)), len(stateViolations)))
	sb.WriteString(fmt.Sprintf("| 3 | BALLOT TYPE valid values | %s | %d |\n", passOrFail(len(ballotTypeViolations)), len(ballotTypeViolations)))
	sb.WriteString(fmt.Sprintf("| 4 | Postal vote rules | %s | %d issues + %d multi-postal DUNs |\n", passOrFail(len(postalViolations)+len(multiPostalDUNs)), len(postalViolations), len(multiPostalDUNs)))
	sb.WriteString(fmt.Sprintf("| 5 | Early vote rules | %s | %d |\n", passOrFail(len(earlyViolations)), len(earlyViolations)))
	sb.WriteString(fmt.Sprintf("| 6 | VOTING CHANNEL NUMBER consistency | %s | %d |\n", passOrFail(len(channelViolations)), len(channelViolations)))
	sb.WriteString(fmt.Sprintf("| 7 | SEX/AGE consistency | %s | %d inconsistencies, %d invalid SEX, %d invalid AGE |\n",
		passOrFail(len(sexAgeIssues)+len(sexValueViolations)+len(ageValueViolations)), len(sexAgeIssues), len(sexValueViolations), len(ageValueViolations)))
	sb.WriteString(fmt.Sprintf("| 8 | Numeric columns | %s | %d |\n", passOrFail(len(numericViolations)), len(numericViolations)))
	sb.WriteString(fmt.Sprintf("| 9 | CHECKER columns | %s | %d |\n", passOrFail(len(checkerViolations)), len(checkerViolations)))
	sb.WriteString(fmt.Sprintf("| 10 | UNIQUE CODE structure | %s | %d |\n", passOrFail(len(ucViolations)), len(ucViolations)))
	sb.WriteString(fmt.Sprintf("| — | PDC in UNIQUE CODE vs col 8 | %s | %d |\n", passOrFail(len(pdcUCViolations)), len(pdcUCViolations)))
	sb.WriteString("\n")

	if totalViolations == 0 {
		sb.WriteString("### ✅ Overall: PASS — All column consistency rules satisfied\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("### ⚠️ Overall: Issues found — %d total violation(s) across all rules\n\n", totalViolations))
	}

	// ---------- Detail sections ----------

	// RULE 1
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 1: Column Count\n\n")
	if len(colCountViolations) == 0 {
		sb.WriteString("✅ **PASS**: All rows have exactly 67 columns.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d row(s) have incorrect column counts.\n\n", len(colCountViolations)))
		sb.WriteString("| Row | UNIQUE CODE (col 1) | Actual Columns | Expected |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		limit := min(len(colCountViolations), 50)
		for _, v := range colCountViolations[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | `%s` | %s | %d |\n", v.row, truncate(v.value, 50), extractColCount(v.msg), expectedCols))
		}
		if len(colCountViolations) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* | | |\n", len(colCountViolations)-limit))
		}
		sb.WriteString("\n")

		// Show distribution of column counts
		colCountDist := make(map[string]int)
		for _, v := range colCountViolations {
			colCountDist[extractColCount(v.msg)]++
		}
		sb.WriteString("**Column count distribution (non-67 rows):**\n\n")
		sb.WriteString("| Column Count | Occurrences |\n")
		sb.WriteString("| --- | --- |\n")
		var sortedCounts []string
		for c := range colCountDist {
			sortedCounts = append(sortedCounts, c)
		}
		sort.Strings(sortedCounts)
		for _, c := range sortedCounts {
			sb.WriteString(fmt.Sprintf("| %s | %d |\n", c, colCountDist[c]))
		}
		sb.WriteString("\n")

		sb.WriteString("**Note**: Column count discrepancies are typically caused by unescaped commas in polling centre names ")
		sb.WriteString("(e.g., embedded commas in school names or addresses). The Go `encoding/csv` reader handles quoted fields properly, ")
		sb.WriteString("so these are counted using proper CSV parsing. If the Go CSV reader parsed all rows without error and this ")
		sb.WriteString("rule shows violations, it means the actual CSV field counts differ from 67.\n\n")
	}

	// RULE 2
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 2: STATE Column\n\n")
	if len(stateViolations) == 0 {
		sb.WriteString("✅ **PASS**: All rows have STATE = 'SARAWAK'.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d row(s) have incorrect STATE values.\n\n", len(stateViolations)))
		sb.WriteString("| Row | STATE Value | UNIQUE CODE |\n")
		sb.WriteString("| --- | --- | --- |\n")
		for _, v := range stateViolations {
			sb.WriteString(fmt.Sprintf("| %d | `%s` | *(see data)* |\n", v.row, truncate(v.value, 60)))
		}
		sb.WriteString("\n")
		sb.WriteString("**Analysis**: Rows where column 2 is not 'SARAWAK' likely indicate shifted/misaligned data, ")
		sb.WriteString("possibly caused by unescaped special characters (commas, quotes, newlines) in preceding fields.\n\n")
	}

	// RULE 3
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 3: BALLOT TYPE\n\n")
	if len(ballotTypeViolations) == 0 {
		sb.WriteString("✅ **PASS**: All rows have valid BALLOT TYPE values.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d row(s) have invalid BALLOT TYPE values.\n\n", len(ballotTypeViolations)))
		sb.WriteString("| Row | BALLOT TYPE Value |\n")
		sb.WriteString("| --- | --- |\n")
		limit := min(len(ballotTypeViolations), 30)
		for _, v := range ballotTypeViolations[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | `%s` |\n", v.row, truncate(v.value, 60)))
		}
		if len(ballotTypeViolations) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* |\n", len(ballotTypeViolations)-limit))
		}
		sb.WriteString("\n")
	}

	// RULE 4
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 4: Postal Vote Rules\n\n")
	sb.WriteString(fmt.Sprintf("Total postal vote rows: **%d**\n\n", len(postalFormats)))

	if len(postalViolations) == 0 && len(multiPostalDUNs) == 0 {
		sb.WriteString("✅ **PASS**: All postal vote rows follow the expected rules.\n\n")
	} else {
		if len(postalViolations) > 0 {
			sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d postal vote sub-rule violation(s) found.\n\n", len(postalViolations)))
			sb.WriteString("| Row | Sub-Rule | Detail |\n")
			sb.WriteString("| --- | --- | --- |\n")
			for _, v := range postalViolations {
				sb.WriteString(fmt.Sprintf("| %d | %s | %s |\n", v.row, v.subRule, v.detail))
			}
			sb.WriteString("\n")
		}
		if len(multiPostalDUNs) > 0 {
			sb.WriteString(fmt.Sprintf("⚠️ **WARNING**: %d DUN(s) have more than 1 postal vote row.\n\n", len(multiPostalDUNs)))
			sb.WriteString("| DUN | Postal Rows | Row Numbers |\n")
			sb.WriteString("| --- | --- | --- |\n")
			for _, d := range multiPostalDUNs {
				rowStrs := make([]string, len(d.rows))
				for i, r := range d.rows {
					rowStrs[i] = strconv.Itoa(r)
				}
				sb.WriteString(fmt.Sprintf("| %s | %d | %s |\n", d.dun, len(d.rows), strings.Join(rowStrs, ", ")))
			}
			sb.WriteString("\n")
		}
	}

	// Postal format distribution
	sb.WriteString("### Postal Vote DMKOD Format Distribution\n\n")
	sb.WriteString("| Format | Count | DUNs |\n")
	sb.WriteString("| --- | --- | --- |\n")
	posFmtDUNs := make(map[string][]string)
	for _, pf := range postalFormats {
		posFmtDUNs[pf.format] = append(posFmtDUNs[pf.format], pf.dun)
	}
	for fmt2, duns := range posFmtDUNs {
		sort.Strings(duns)
		display := strings.Join(duns, ", ")
		if len(display) > 100 {
			display = fmt.Sprintf("%s ... (%d total)", strings.Join(duns[:5], ", "), len(duns))
		}
		sb.WriteString(fmt.Sprintf("| `%s` | %d | %s |\n", fmt2, len(duns), display))
	}
	sb.WriteString("\n")
	sb.WriteString("**Note**: N.01 (OPAR) uses the older `/POS` format while all other DUNs use `/UNDI POS`. ")
	sb.WriteString("This was already identified in Phase-0.\n\n")

	// RULE 5
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 5: Early Vote Rules\n\n")
	if len(earlyViolations) == 0 {
		sb.WriteString("✅ **PASS**: All early vote rows follow the expected rules.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d early vote rule violation(s) found.\n\n", len(earlyViolations)))
		sb.WriteString("| Row | Sub-Rule | Detail |\n")
		sb.WriteString("| --- | --- | --- |\n")
		for _, v := range earlyViolations {
			sb.WriteString(fmt.Sprintf("| %d | %s | %s |\n", v.row, v.subRule, v.detail))
		}
		sb.WriteString("\n")
	}

	// RULE 6
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 6: VOTING CHANNEL NUMBER Consistency\n\n")
	if len(channelViolations) == 0 {
		sb.WriteString("✅ **PASS**: VOTING CHANNEL NUMBER matches UNIQUE CODE suffix for all rows.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d row(s) have mismatched channel numbers.\n\n", len(channelViolations)))
		sb.WriteString("| Row | UNIQUE CODE | Channel in UC | VOTING CHANNEL NUMBER (col 11) |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		limit := min(len(channelViolations), 50)
		for _, v := range channelViolations[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | `%s` | *(see msg)* | `%s` |\n", v.row, truncate(v.msg, 100), v.value))
		}
		if len(channelViolations) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* | | |\n", len(channelViolations)-limit))
		}
		sb.WriteString("\n")
	}

	// RULE 7
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 7: SEX and AGE Column Consistency\n\n")

	if len(sexAgeIssues) == 0 && len(sexValueViolations) == 0 && len(ageValueViolations) == 0 {
		sb.WriteString("✅ **PASS**: All SEX/AGE values are consistent within each DUN/candidate.\n\n")
	} else {
		if len(sexAgeIssues) > 0 {
			sb.WriteString(fmt.Sprintf("⚠️ **WARNING**: %d SEX/AGE inconsistencies found (same candidate has both filled and empty values within same DUN).\n\n", len(sexAgeIssues)))
			sb.WriteString("| DUN | Party | Candidate | Field | Filled Rows | Empty Rows | Example Empty Row(s) |\n")
			sb.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
			limit := min(len(sexAgeIssues), 100)
			for _, iss := range sexAgeIssues[:limit] {
				sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %d | %d | %s |\n",
					iss.dun, iss.party, truncate(iss.candidate, 30), iss.field,
					iss.filledCnt, iss.emptyCnt, strings.Join(iss.examples, ", ")))
			}
			if len(sexAgeIssues) > limit {
				sb.WriteString(fmt.Sprintf("| ... | | | | | | *(and %d more)* |\n", len(sexAgeIssues)-limit))
			}
			sb.WriteString("\n")
		}
		if len(sexValueViolations) > 0 {
			sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d invalid SEX values found (not MALE/FEMALE/empty).\n\n", len(sexValueViolations)))
			sb.WriteString("| Row | Column | Value | Candidate |\n")
			sb.WriteString("| --- | --- | --- | --- |\n")
			limit := min(len(sexValueViolations), 30)
			for _, v := range sexValueViolations[:limit] {
				sb.WriteString(fmt.Sprintf("| %d | %s | `%s` | %s |\n", v.row, v.colName, v.value, extractCandidate(v.msg)))
			}
			sb.WriteString("\n")
		}
		if len(ageValueViolations) > 0 {
			sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d non-numeric AGE values found.\n\n", len(ageValueViolations)))
			sb.WriteString("| Row | Column | Value | Candidate |\n")
			sb.WriteString("| --- | --- | --- | --- |\n")
			limit := min(len(ageValueViolations), 30)
			for _, v := range ageValueViolations[:limit] {
				sb.WriteString(fmt.Sprintf("| %d | %s | `%s` | %s |\n", v.row, v.colName, v.value, extractCandidate(v.msg)))
			}
			sb.WriteString("\n")
		}
	}

	// RULE 8
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 8: Numeric Columns\n\n")
	if len(numericViolations) == 0 {
		sb.WriteString("✅ **PASS**: All expected numeric columns contain valid integers or are empty.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d non-numeric value(s) found in expected numeric columns.\n\n", len(numericViolations)))
		sb.WriteString("| Row | Column | Column Name | Value |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		limit := min(len(numericViolations), 50)
		for _, v := range numericViolations[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | %d | %s | `%s` |\n", v.row, v.col, v.colName, truncate(v.value, 40)))
		}
		if len(numericViolations) > limit {
			sb.WriteString(fmt.Sprintf("| ... | | | *(and %d more)* |\n", len(numericViolations)-limit))
		}
		sb.WriteString("\n")
	}

	// RULE 9
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 9: CHECKER Columns\n\n")

	// Separate by check type
	var validVoteMismatch []checkerViolation
	var validVoteCheckerWrong []checkerViolation
	var totalVoteMismatch []checkerViolation
	var totalVoteCheckerWrong []checkerViolation
	for _, cv := range checkerViolations {
		switch cv.checkType {
		case "VALID VOTE MISMATCH":
			validVoteMismatch = append(validVoteMismatch, cv)
		case "VALID VOTE":
			validVoteCheckerWrong = append(validVoteCheckerWrong, cv)
		case "TOTAL VOTE ISSUED MISMATCH":
			totalVoteMismatch = append(totalVoteMismatch, cv)
		case "TOTAL VOTE ISSUED":
			totalVoteCheckerWrong = append(totalVoteCheckerWrong, cv)
		}
	}

	if len(checkerViolations) == 0 {
		sb.WriteString("✅ **PASS**: All CHECKER columns are correct.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("⚠️ **Issues found**: %d total checker violation(s).\n\n", len(checkerViolations)))
	}

	sb.WriteString("### CHECKER (VALID VOTE) — Sum of candidate votes vs TOTAL VALID VOTES\n\n")
	if len(validVoteMismatch) == 0 {
		sb.WriteString("✅ Sum of candidate votes equals TOTAL VALID VOTES for all rows.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ %d row(s) where sum of candidate votes ≠ TOTAL VALID VOTES.\n\n", len(validVoteMismatch)))
		sb.WriteString("| Row | Detail |\n")
		sb.WriteString("| --- | --- |\n")
		limit := min(len(validVoteMismatch), 30)
		for _, cv := range validVoteMismatch[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | %s |\n", cv.row, cv.detail))
		}
		if len(validVoteMismatch) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* |\n", len(validVoteMismatch)-limit))
		}
		sb.WriteString("\n")
	}

	if len(validVoteCheckerWrong) > 0 {
		sb.WriteString(fmt.Sprintf("⚠️ %d row(s) where CHECKER (VALID VOTE) flag is incorrect.\n\n", len(validVoteCheckerWrong)))
		sb.WriteString("| Row | Expected | Actual | Detail |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		for _, cv := range validVoteCheckerWrong {
			sb.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n", cv.row, cv.expected, cv.actual, cv.detail))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("### CHECKER (TOTAL VOTE ISSUED) — TOTAL BALLOTS ISSUED vs Valid+Rejected+Unreturned\n\n")
	if len(totalVoteMismatch) == 0 {
		sb.WriteString("✅ TOTAL BALLOTS ISSUED = Valid + Rejected + Unreturned for all rows.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ %d row(s) where TOTAL BALLOTS ISSUED ≠ Valid+Rejected+Unreturned.\n\n", len(totalVoteMismatch)))
		sb.WriteString("| Row | Detail |\n")
		sb.WriteString("| --- | --- |\n")
		limit := min(len(totalVoteMismatch), 30)
		for _, cv := range totalVoteMismatch[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | %s |\n", cv.row, cv.detail))
		}
		if len(totalVoteMismatch) > limit {
			sb.WriteString(fmt.Sprintf("| ... | *(and %d more)* |\n", len(totalVoteMismatch)-limit))
		}
		sb.WriteString("\n")
	}

	if len(totalVoteCheckerWrong) > 0 {
		sb.WriteString(fmt.Sprintf("⚠️ %d row(s) where CHECKER (TOTAL VOTE ISSUED) flag is incorrect.\n\n", len(totalVoteCheckerWrong)))
		sb.WriteString("| Row | Expected | Actual | Detail |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		for _, cv := range totalVoteCheckerWrong {
			sb.WriteString(fmt.Sprintf("| %d | %s | %s | %s |\n", cv.row, cv.expected, cv.actual, cv.detail))
		}
		sb.WriteString("\n")
	}

	// RULE 10
	sb.WriteString("---\n\n")
	sb.WriteString("## Rule 10: UNIQUE CODE Structure\n\n")
	sb.WriteString("Expected pattern: `P.XXX_N.YY_XXX/YY/ZZ_C`\n\n")
	if len(ucViolations) == 0 {
		sb.WriteString("✅ **PASS**: All UNIQUE CODEs match the expected pattern and are internally consistent.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d UNIQUE CODE structural violation(s) found.\n\n", len(ucViolations)))

		// Categorize violations
		var patternFails []violation
		var pCodeMismatch []violation
		var dunCodeMismatch []violation
		var internalMismatch []violation
		var channelMismatch []violation
		var emptyUC []violation
		for _, v := range ucViolations {
			switch {
			case strings.Contains(v.msg, "does not match pattern"):
				patternFails = append(patternFails, v)
			case strings.Contains(v.msg, "P-code in UNIQUE CODE"):
				pCodeMismatch = append(pCodeMismatch, v)
			case strings.Contains(v.msg, "DUN code in UNIQUE CODE"):
				dunCodeMismatch = append(dunCodeMismatch, v)
			case strings.Contains(v.msg, "mismatch within UNIQUE CODE"):
				internalMismatch = append(internalMismatch, v)
			case strings.Contains(v.msg, "Channel in UNIQUE CODE"):
				channelMismatch = append(channelMismatch, v)
			case strings.Contains(v.msg, "Empty UNIQUE CODE"):
				emptyUC = append(emptyUC, v)
			}
		}

		if len(emptyUC) > 0 {
			sb.WriteString(fmt.Sprintf("### Empty UNIQUE CODEs: %d\n\n", len(emptyUC)))
			sb.WriteString("| Row |\n| --- |\n")
			for _, v := range emptyUC {
				sb.WriteString(fmt.Sprintf("| %d |\n", v.row))
			}
			sb.WriteString("\n")
		}

		if len(patternFails) > 0 {
			sb.WriteString(fmt.Sprintf("### Pattern Mismatch: %d\n\n", len(patternFails)))
			sb.WriteString("| Row | UNIQUE CODE | Issue |\n")
			sb.WriteString("| --- | --- | --- |\n")
			limit := min(len(patternFails), 30)
			for _, v := range patternFails[:limit] {
				sb.WriteString(fmt.Sprintf("| %d | `%s` | %s |\n", v.row, truncate(v.value, 50), v.msg))
			}
			if len(patternFails) > limit {
				sb.WriteString(fmt.Sprintf("| ... | | *(and %d more)* |\n", len(patternFails)-limit))
			}
			sb.WriteString("\n")
		}

		if len(pCodeMismatch) > 0 {
			sb.WriteString(fmt.Sprintf("### P-Code Mismatch (UC vs col 4): %d\n\n", len(pCodeMismatch)))
			sb.WriteString("| Row | Issue |\n| --- | --- |\n")
			for _, v := range pCodeMismatch {
				sb.WriteString(fmt.Sprintf("| %d | %s |\n", v.row, v.msg))
			}
			sb.WriteString("\n")
		}

		if len(dunCodeMismatch) > 0 {
			sb.WriteString(fmt.Sprintf("### DUN Code Mismatch (UC vs col 6): %d\n\n", len(dunCodeMismatch)))
			sb.WriteString("| Row | Issue |\n| --- | --- |\n")
			for _, v := range dunCodeMismatch {
				sb.WriteString(fmt.Sprintf("| %d | %s |\n", v.row, v.msg))
			}
			sb.WriteString("\n")
		}

		if len(internalMismatch) > 0 {
			sb.WriteString(fmt.Sprintf("### Internal Code Mismatch within UC: %d\n\n", len(internalMismatch)))
			sb.WriteString("| Row | Issue |\n| --- | --- |\n")
			for _, v := range internalMismatch {
				sb.WriteString(fmt.Sprintf("| %d | %s |\n", v.row, v.msg))
			}
			sb.WriteString("\n")
		}

		if len(channelMismatch) > 0 {
			sb.WriteString(fmt.Sprintf("### Channel Mismatch (UC vs col 11): %d\n\n", len(channelMismatch)))
			sb.WriteString("| Row | Issue |\n| --- | --- |\n")
			limit := min(len(channelMismatch), 20)
			for _, v := range channelMismatch[:limit] {
				sb.WriteString(fmt.Sprintf("| %d | %s |\n", v.row, v.msg))
			}
			sb.WriteString("\n")
		}
	}

	// PDC consistency
	sb.WriteString("---\n\n")
	sb.WriteString("## Additional: POLLING DISTRICT CODE in UNIQUE CODE vs Column 8\n\n")
	if len(pdcUCViolations) == 0 {
		sb.WriteString("✅ **PASS**: POLLING DISTRICT CODE embedded in UNIQUE CODE matches column 8 for all rows.\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("❌ **FAIL**: %d mismatch(es) found.\n\n", len(pdcUCViolations)))
		sb.WriteString("| Row | UNIQUE CODE | PDC in UC | PDC in Col 8 |\n")
		sb.WriteString("| --- | --- | --- | --- |\n")
		limit := min(len(pdcUCViolations), 30)
		for _, v := range pdcUCViolations[:limit] {
			sb.WriteString(fmt.Sprintf("| %d | `%s` | *(see msg)* | %s |\n", v.row, truncate(v.value, 50), v.msg))
		}
		if len(pdcUCViolations) > limit {
			sb.WriteString(fmt.Sprintf("| ... | | | *(and %d more)* |\n", len(pdcUCViolations)-limit))
		}
		sb.WriteString("\n")
	}

	// ========== Ballot type distribution ==========
	sb.WriteString("---\n\n")
	sb.WriteString("## Appendix: Data Distribution\n\n")
	sb.WriteString("### Ballot Type Distribution\n\n")
	btDist := make(map[string]int)
	for _, row := range dataRows {
		if len(row) < 3 {
			btDist["<insufficient columns>"]++
			continue
		}
		btDist[strings.TrimSpace(row[2])]++
	}
	sb.WriteString("| Ballot Type | Count |\n")
	sb.WriteString("| --- | --- |\n")
	for _, bt := range []string{"POSTAL VOTE", "EARLY VOTE", "ORDINARY VOTE"} {
		if c, ok := btDist[bt]; ok {
			sb.WriteString(fmt.Sprintf("| %s | %d |\n", bt, c))
		}
	}
	// Show any extras
	for bt, c := range btDist {
		if bt != "POSTAL VOTE" && bt != "EARLY VOTE" && bt != "ORDINARY VOTE" {
			sb.WriteString(fmt.Sprintf("| `%s` (INVALID) | %d |\n", truncate(bt, 40), c))
		}
	}
	sb.WriteString("\n")

	// Party column usage
	sb.WriteString("### Party Column Usage\n\n")
	sb.WriteString("Shows how many rows have a non-empty candidate name for each party group.\n\n")
	sb.WriteString("| Party Group | Rows with Candidate | Unique Candidates | Party Labels Used |\n")
	sb.WriteString("| --- | --- | --- | --- |\n")
	for _, pg := range partyGroups {
		candCount := 0
		uniqueCands := make(map[string]bool)
		partyLabels := make(map[string]bool)
		for _, row := range dataRows {
			if pg.candCol >= len(row) {
				continue
			}
			cand := strings.TrimSpace(row[pg.candCol])
			if cand != "" {
				candCount++
				uniqueCands[cand] = true
			}
			if pg.partyCol < len(row) {
				pLabel := strings.TrimSpace(row[pg.partyCol])
				if pLabel != "" {
					partyLabels[pLabel] = true
				}
			}
		}
		labels := make([]string, 0, len(partyLabels))
		for l := range partyLabels {
			labels = append(labels, l)
		}
		sort.Strings(labels)
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %s |\n", pg.name, candCount, len(uniqueCands), strings.Join(labels, ", ")))
	}
	sb.WriteString("\n")

	// Write report
	outputPath := "PHASE-5-REVIEW.md"
	err = os.WriteFile(outputPath, []byte(sb.String()), 0644)
	if err != nil {
		slog.Error("failed to write report", "path", outputPath, "error", err)
		os.Exit(1)
	}
	slog.Info("Report written", "path", outputPath)

	// Print summary
	fmt.Println("=== PHASE 5: Column Consistency and Mapping Rules ===")
	fmt.Printf("Total data rows:            %d\n", len(dataRows))
	fmt.Printf("Rule 1  (column count):     %d violations\n", len(colCountViolations))
	fmt.Printf("Rule 2  (STATE):            %d violations\n", len(stateViolations))
	fmt.Printf("Rule 3  (BALLOT TYPE):      %d violations\n", len(ballotTypeViolations))
	fmt.Printf("Rule 4  (postal vote):      %d issues + %d multi-postal DUNs\n", len(postalViolations), len(multiPostalDUNs))
	fmt.Printf("Rule 5  (early vote):       %d violations\n", len(earlyViolations))
	fmt.Printf("Rule 6  (channel number):   %d violations\n", len(channelViolations))
	fmt.Printf("Rule 7  (SEX/AGE):          %d inconsistencies, %d invalid SEX, %d invalid AGE\n", len(sexAgeIssues), len(sexValueViolations), len(ageValueViolations))
	fmt.Printf("Rule 8  (numeric cols):     %d violations\n", len(numericViolations))
	fmt.Printf("Rule 9  (CHECKER cols):     %d violations\n", len(checkerViolations))
	fmt.Printf("Rule 10 (UNIQUE CODE):      %d violations\n", len(ucViolations))
	fmt.Printf("Extra   (PDC in UC):        %d violations\n", len(pdcUCViolations))
	fmt.Printf("\nTotal violations:           %d\n", totalViolations)
	if totalViolations > 0 {
		fmt.Println("\n⚠️  Issues found. See PHASE-5-REVIEW.md for details.")
	} else {
		fmt.Println("\n✅ All checks passed.")
	}
}

// ---- Helpers ----

func safeGet(row []string, idx int) string {
	if idx < len(row) {
		return row[idx]
	}
	return ""
}

func passOrFail(count int) string {
	if count == 0 {
		return "✅ PASS"
	}
	return "❌ FAIL"
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func extractColCount(msg string) string {
	// Extract "got XX" from message
	parts := strings.Split(msg, "got ")
	if len(parts) == 2 {
		return strings.TrimSpace(parts[1])
	}
	return msg
}

func extractCandidate(msg string) string {
	// Extract candidate name from "... for candidate 'XXX'"
	idx := strings.Index(msg, "for candidate '")
	if idx >= 0 {
		rest := msg[idx+len("for candidate '"):]
		end := strings.Index(rest, "'")
		if end >= 0 {
			return rest[:end]
		}
	}
	return ""
}

func intsToStrings(ints []int) []string {
	result := make([]string, len(ints))
	for i, v := range ints {
		result[i] = strconv.Itoa(v)
	}
	return result
}

func dunSortKey(dun string) int {
	s := strings.TrimPrefix(dun, "N.")
	n, _ := strconv.Atoi(s)
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
