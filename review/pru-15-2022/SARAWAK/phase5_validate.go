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

const expectedCols = 92

// Column indices (0-based)
const (
	colUniqueCode    = 0
	colState         = 1
	colBallotType    = 2
	colParCode       = 3
	colParName       = 4
	colStateCode     = 5
	colStateName     = 6
	colPollDistCode  = 7
	colPollDistName  = 8
	colPollCentre    = 9
	colVotingChannel = 20
	colTotalBallots  = 21
	// Party slots: each is 5 cols (party, candidate, gender, age, vote)
	colBNStart     = 22
	colPHStart     = 27
	colPNStart     = 32
	colGTAStart    = 37
	colGPSStart    = 42
	colGRSStart    = 47
	colWarStart    = 52
	colOther1Start = 57
	colOther2Start = 62
	colOther3Start = 67
	colInd1Start   = 72
	colInd2Start   = 77
	colInd3Start   = 82
	// Totals
	colTotalValid      = 87
	colTotalRejected   = 88
	colTotalUnreturned = 89
	colCheckValid      = 90
	colCheckTotal      = 91
)

// Demographics columns (0-based): 10..19
const (
	colDemoStart = 10
	colDemoEnd   = 19
)

var demoNames = []string{
	"BIDAYUH", "BUMIPUTERA SARAWAK", "IBAN", "CINA",
	"BUMIPUTERA SABAH", "MELAYU", "LAIN LAIN", "INDIA",
	"ORANG ULU", "ORANG ASLI",
}

var partySlotStarts = []int{
	colBNStart, colPHStart, colPNStart, colGTAStart,
	colGPSStart, colGRSStart, colWarStart,
	colOther1Start, colOther2Start, colOther3Start,
	colInd1Start, colInd2Start, colInd3Start,
}

var partySlotNames = []string{
	"BN", "PH", "PN", "GTA",
	"GPS", "GRS", "WARISAN",
	"OTHER PARTY (1)", "OTHER PARTY (2)", "OTHER PARTY (3)",
	"INDEPENDENT 1", "INDEPENDENT 2", "INDEPENDENT 3",
}

type violation struct {
	Row     int
	Code    string
	Message string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	f, err := os.Open("to-review.csv")
	if err != nil {
		slog.Error("cannot open file", "err", err)
		os.Exit(1)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1 // allow variable fields
	reader.LazyQuotes = true

	allRows, err := reader.ReadAll()
	if err != nil {
		slog.Error("cannot read CSV", "err", err)
		os.Exit(1)
	}

	if len(allRows) < 2 {
		slog.Error("file has fewer than 2 rows")
		os.Exit(1)
	}

	header := allRows[0]
	dataRows := allRows[1:]

	slog.Info("file loaded", "headerCols", len(header), "dataRows", len(dataRows))

	// Violation trackers
	var colCountViolations []violation
	var stateViolations []violation
	var ballotTypeViolations []violation
	var postalViolations []violation
	var earlyVoteViolations []violation
	var channelViolations []violation
	var genderAgeViolations []violation
	var checkValidViolations []violation
	var checkTotalViolations []violation
	var uniqueCodeViolations []violation
	var quotedNumberViolations []violation
	var trailingSpaceViolations []violation
	var demoViolations []violation

	// For gender/age consistency tracking per par constituency per party slot
	type candidateGAInfo struct {
		candidate    string
		hasGenderSet bool
		hasAgeSet    bool
		missingGRows []int
		missingARows []int
		filledGRows  []int
		filledARows  []int
	}
	type gaKey struct {
		parCode  string
		slotIdx  int
		slotName string
	}
	gaTracker := make(map[gaKey]*candidateGAInfo)

	// Track postal votes per parliamentary constituency
	postalCountByPar := make(map[string][]int) // parCode -> row numbers

	headerColCount := len(header)

	// Track trailing whitespace by column across all rows
	type wsKey struct {
		col     int
		trimmed string
		raw     string
	}
	wsExamples := make(map[int][]violation) // colIndex -> example violations

	// Track demographics pattern per ballot type
	// Postal votes generally have empty demographics; ordinary/early have values
	demoFilledByBallotType := make(map[string]int) // "POSTAL VOTE:filled" etc
	demoEmptyByBallotType := make(map[string]int)

	// Track unique par codes seen
	parCodes := make(map[string]bool)

	for i, row := range dataRows {
		rowNum := i + 2 // 1-indexed, +1 for header
		uid := ""
		if len(row) > 0 {
			uid = row[colUniqueCode]
		}

		// ============================================================
		// CHECK 1: Column count
		// ============================================================
		if len(row) != expectedCols {
			colCountViolations = append(colCountViolations, violation{
				Row:     rowNum,
				Code:    uid,
				Message: fmt.Sprintf("expected %d columns, got %d", expectedCols, len(row)),
			})
			if len(row) < expectedCols {
				continue
			}
		}

		// ============================================================
		// CHECK: Comma-formatted numbers in numeric fields
		// ============================================================
		numericCols := []int{colTotalBallots, colTotalValid, colTotalRejected, colTotalUnreturned, colCheckValid, colCheckTotal}
		for _, ci := range numericCols {
			if ci < len(row) && strings.Contains(row[ci], ",") {
				quotedNumberViolations = append(quotedNumberViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("col %d (%s) has comma-formatted number: '%s'", ci+1, header[ci], row[ci]),
				})
			}
		}
		// Also check vote columns in party slots
		for si, start := range partySlotStarts {
			voteCol := start + 4
			if voteCol < len(row) && strings.Contains(row[voteCol], ",") {
				quotedNumberViolations = append(quotedNumberViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("%s vote (col %d) has comma-formatted number: '%s'", partySlotNames[si], voteCol+1, row[voteCol]),
				})
			}
		}

		// ============================================================
		// CHECK 2: STATE must be SARAWAK
		// ============================================================
		if strings.TrimSpace(row[colState]) != "SARAWAK" {
			stateViolations = append(stateViolations, violation{
				Row:     rowNum,
				Code:    uid,
				Message: fmt.Sprintf("STATE = '%s'", row[colState]),
			})
		}

		// ============================================================
		// CHECK 3: BALLOT TYPE must be one of 3 values
		// ============================================================
		bt := strings.TrimSpace(row[colBallotType])
		validBallot := bt == "POSTAL VOTE" || bt == "EARLY VOTE" || bt == "ORDINARY VOTE"
		if !validBallot {
			ballotTypeViolations = append(ballotTypeViolations, violation{
				Row:     rowNum,
				Code:    uid,
				Message: fmt.Sprintf("BALLOT TYPE = '%s'", row[colBallotType]),
			})
		}

		parCode := strings.TrimSpace(row[colParCode])
		parCodes[parCode] = true

		// ============================================================
		// CHECK 4: Postal vote rules
		// ============================================================
		if bt == "POSTAL VOTE" {
			postalCountByPar[parCode] = append(postalCountByPar[parCode], rowNum)

			pollDistCode := strings.TrimSpace(row[colPollDistCode])
			pollDistName := strings.TrimSpace(row[colPollDistName])
			pollCentre := strings.TrimSpace(row[colPollCentre])
			stateCode := strings.TrimSpace(row[colStateCode])
			stateName := strings.TrimSpace(row[colStateName])

			// For parliamentary-level postal votes: POLLING DISTRICT CODE = UNDI POS
			if pollDistCode != "UNDI POS" && !strings.HasSuffix(pollDistCode, "/POS") {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POLLING DISTRICT CODE = '%s' (expected 'UNDI POS' or ending '/POS')", pollDistCode),
				})
			}

			if pollDistName != "UNDI POS" {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POLLING DISTRICT NAME = '%s' (expected 'UNDI POS')", pollDistName),
				})
			}

			if pollCentre != "UNDI POS" {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POLLING CENTRE = '%s' (expected 'UNDI POS')", pollCentre),
				})
			}

			// STATE CONSTITUENCY CODE for parliamentary postal = P.XXX/POSTAL VOTE
			expectedSC := parCode + "/POSTAL VOTE"
			if stateCode != expectedSC {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("STATE CONSTITUENCY CODE = '%s' (expected '%s')", stateCode, expectedSC),
				})
			}

			// STATE CONSTITUENCY NAME for postal = UNDI POS
			if stateName != "UNDI POS" {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("STATE CONSTITUENCY NAME = '%s' (expected 'UNDI POS')", stateName),
				})
			}

			// VOTING CHANNEL should be 1 for postal
			ch := strings.TrimSpace(row[colVotingChannel])
			if ch != "1" {
				postalViolations = append(postalViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("VOTING CHANNEL = '%s' (expected '1' for postal)", ch),
				})
			}
		}

		// ============================================================
		// CHECK 5: Early vote rules
		// ============================================================
		if bt == "EARLY VOTE" {
			pollDistCode := strings.TrimSpace(row[colPollDistCode])
			pollDistName := strings.TrimSpace(row[colPollDistName])

			if !strings.HasSuffix(pollDistCode, "/00") {
				earlyVoteViolations = append(earlyVoteViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POLLING DISTRICT CODE = '%s' (expected suffix '/00')", pollDistCode),
				})
			}

			if pollDistName != "UNDI AWAL" {
				earlyVoteViolations = append(earlyVoteViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POLLING DISTRICT NAME = '%s' (expected 'UNDI AWAL')", pollDistName),
				})
			}
		}

		// ============================================================
		// CHECK 6: VOTING CHANNEL NUMBER must match UNIQUE CODE suffix
		// ============================================================
		channelStr := strings.TrimSpace(row[colVotingChannel])
		uidTrimmed := strings.TrimSpace(uid)

		if uidTrimmed != "" && channelStr != "" {
			lastUnderscore := strings.LastIndex(uidTrimmed, "_")
			if lastUnderscore >= 0 {
				suffix := uidTrimmed[lastUnderscore+1:]
				// For suffix-disambiguated codes like "a1", "b2", extract the numeric part
				re := regexp.MustCompile(`^[a-z]*(\d+)$`)
				matches := re.FindStringSubmatch(suffix)
				if len(matches) == 2 {
					expectedChannel := matches[1]
					if expectedChannel != channelStr {
						channelViolations = append(channelViolations, violation{
							Row:     rowNum,
							Code:    uid,
							Message: fmt.Sprintf("VOTING CHANNEL = '%s' but UNIQUE CODE suffix = '%s' (extracted channel '%s')", channelStr, suffix, expectedChannel),
						})
					}
				} else {
					// Special case for postal: suffix might be "UNDI POS" part
					// P.192_P.192/POSTAL VOTE_UNDI POS_1 -> last _ -> "1" which is fine
					// But if parsing fails, flag it
					if suffix != "POS" {
						channelViolations = append(channelViolations, violation{
							Row:     rowNum,
							Code:    uid,
							Message: fmt.Sprintf("Cannot parse channel from UNIQUE CODE suffix '%s'", suffix),
						})
					}
				}
			}
		}

		// ============================================================
		// CHECK 7: Gender/Age consistency per candidate per party slot
		// ============================================================
		for si, start := range partySlotStarts {
			candCol := start + 1
			genderCol := start + 2
			ageCol := start + 3

			if candCol >= len(row) {
				continue
			}

			cand := strings.TrimSpace(row[candCol])
			if cand == "" {
				continue
			}

			gender := ""
			age := ""
			if genderCol < len(row) {
				gender = strings.TrimSpace(row[genderCol])
			}
			if ageCol < len(row) {
				age = strings.TrimSpace(row[ageCol])
			}

			// Validate gender values when present
			if gender != "" && gender != "MALE" && gender != "FEMALE" {
				genderAgeViolations = append(genderAgeViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("%s slot: invalid GENDER value '%s' for candidate '%s'", partySlotNames[si], gender, cand),
				})
			}

			// Validate age is numeric when present
			if age != "" {
				if _, err := strconv.Atoi(age); err != nil {
					genderAgeViolations = append(genderAgeViolations, violation{
						Row:     rowNum,
						Code:    uid,
						Message: fmt.Sprintf("%s slot: non-numeric AGE '%s' for candidate '%s'", partySlotNames[si], age, cand),
					})
				}
			}

			key := gaKey{parCode: parCode, slotIdx: si, slotName: partySlotNames[si]}
			info, ok := gaTracker[key]
			if !ok {
				info = &candidateGAInfo{candidate: cand}
				gaTracker[key] = info
			}

			if gender != "" {
				info.hasGenderSet = true
				info.filledGRows = append(info.filledGRows, rowNum)
			} else {
				info.missingGRows = append(info.missingGRows, rowNum)
			}
			if age != "" {
				info.hasAgeSet = true
				info.filledARows = append(info.filledARows, rowNum)
			} else {
				info.missingARows = append(info.missingARows, rowNum)
			}
		}

		// ============================================================
		// CHECK 8: CHECK columns validation
		// ============================================================
		totalCandVotes := 0
		candVoteParseOK := true
		for _, start := range partySlotStarts {
			voteCol := start + 4
			if voteCol >= len(row) {
				continue
			}
			vStr := strings.TrimSpace(row[voteCol])
			if vStr == "" {
				continue
			}
			vStr = strings.ReplaceAll(vStr, ",", "")
			v, err := strconv.Atoi(vStr)
			if err != nil {
				candVoteParseOK = false
				break
			}
			totalCandVotes += v
		}

		totalValid := parseIntSafe(row[colTotalValid])
		totalRejected := parseIntSafe(row[colTotalRejected])
		totalUnreturned := parseIntSafe(row[colTotalUnreturned])
		totalBallots := parseIntSafe(row[colTotalBallots])
		checkValid := parseIntSafe(row[colCheckValid])
		checkTotal := parseIntSafe(row[colCheckTotal])

		// CHECK ON VALID VOTES = sum of candidate votes - TOTAL VALID VOTES (should be 0)
		if candVoteParseOK {
			computedCheckValid := totalCandVotes - totalValid
			if computedCheckValid != checkValid {
				checkValidViolations = append(checkValidViolations, violation{
					Row:  rowNum,
					Code: uid,
					Message: fmt.Sprintf("CHECK ON VALID VOTES: stored=%d, computed=%d (candVotes=%d - totalValid=%d)",
						checkValid, computedCheckValid, totalCandVotes, totalValid),
				})
			} else if checkValid != 0 {
				checkValidViolations = append(checkValidViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("CHECK ON VALID VOTES is non-zero: %d (correctly computed but still a data issue)", checkValid),
				})
			}
		}

		// CHECK ON TOTAL = BALLOTS - VALID - REJECTED - UNRETURNED (should be 0)
		computedCheckTotal := totalBallots - totalValid - totalRejected - totalUnreturned
		if computedCheckTotal != checkTotal {
			checkTotalViolations = append(checkTotalViolations, violation{
				Row:  rowNum,
				Code: uid,
				Message: fmt.Sprintf("CHECK ON TOTAL: stored=%d, computed=%d (ballots=%d - valid=%d - rejected=%d - unreturned=%d)",
					checkTotal, computedCheckTotal, totalBallots, totalValid, totalRejected, totalUnreturned),
			})
		} else if checkTotal != 0 {
			checkTotalViolations = append(checkTotalViolations, violation{
				Row:     rowNum,
				Code:    uid,
				Message: fmt.Sprintf("CHECK ON TOTAL VOTES ISSUED is non-zero: %d (correctly computed but still a data issue)", checkTotal),
			})
		}

		// ============================================================
		// CHECK 9: UNIQUE CODE construction consistency
		// ============================================================
		if bt == "ORDINARY VOTE" || bt == "EARLY VOTE" {
			stateCodeVal := strings.TrimSpace(row[colStateCode])
			pollDistCodeVal := strings.TrimSpace(row[colPollDistCode])

			if uidTrimmed != "" && stateCodeVal != "" && pollDistCodeVal != "" {
				expectedPrefix := parCode + "_" + stateCodeVal + "_" + pollDistCodeVal + "_"
				if !strings.HasPrefix(uidTrimmed, expectedPrefix) {
					re := regexp.MustCompile(`^(.+)_([a-z]?\d+)$`)
					matches := re.FindStringSubmatch(uidTrimmed)
					if len(matches) == 3 {
						uidBase := matches[1]
						expectedBase := parCode + "_" + stateCodeVal + "_" + pollDistCodeVal
						if uidBase != expectedBase {
							uniqueCodeViolations = append(uniqueCodeViolations, violation{
								Row:  rowNum,
								Code: uid,
								Message: fmt.Sprintf("UNIQUE CODE base '%s' doesn't match expected '%s' (parCode=%s, stateCode=%s, PDC=%s)",
									uidBase, expectedBase, parCode, stateCodeVal, pollDistCodeVal),
							})
						}
					} else {
						uniqueCodeViolations = append(uniqueCodeViolations, violation{
							Row:     rowNum,
							Code:    uid,
							Message: fmt.Sprintf("UNIQUE CODE format unrecognized (expected prefix '%s')", expectedPrefix),
						})
					}
				}
			}
		} else if bt == "POSTAL VOTE" {
			expectedPostalPrefix := parCode + "_" + parCode + "/POSTAL VOTE_UNDI POS_"
			if !strings.HasPrefix(uidTrimmed, expectedPostalPrefix) {
				uniqueCodeViolations = append(uniqueCodeViolations, violation{
					Row:     rowNum,
					Code:    uid,
					Message: fmt.Sprintf("POSTAL VOTE UNIQUE CODE doesn't match expected prefix '%s'", expectedPostalPrefix),
				})
			}
		}

		// ============================================================
		// CHECK: Trailing/leading whitespace in key fields
		// ============================================================
		wsCheckCols := []int{
			colUniqueCode, colState, colBallotType, colParCode, colParName,
			colStateCode, colStateName, colPollDistCode, colPollDistName, colPollCentre,
			colVotingChannel,
		}
		for _, ci := range wsCheckCols {
			if ci < len(row) {
				raw := row[ci]
				trimmed := strings.TrimSpace(raw)
				if raw != trimmed && trimmed != "" {
					v := violation{
						Row:     rowNum,
						Code:    uid,
						Message: fmt.Sprintf("col %d (%s): has whitespace — raw='%s' trimmed='%s'", ci+1, header[ci], raw, trimmed),
					}
					wsExamples[ci] = append(wsExamples[ci], v)
				}
			}
		}

		// ============================================================
		// CHECK: Demographics consistency
		// ============================================================
		hasDemoFilled := false
		hasDemoEmpty := false
		for di := colDemoStart; di <= colDemoEnd; di++ {
			if di < len(row) {
				val := strings.TrimSpace(row[di])
				if val != "" {
					hasDemoFilled = true
				} else {
					hasDemoEmpty = true
				}
			}
		}
		btKey := bt
		if hasDemoFilled {
			demoFilledByBallotType[btKey]++
		}
		if !hasDemoFilled {
			demoEmptyByBallotType[btKey]++
		}
		// Mixed within a row: some demo cols filled, some empty — possible issue
		// Skip this for postal votes (expected empty) and if demographics are percentages (all-or-nothing)
		if bt != "POSTAL VOTE" && hasDemoFilled && hasDemoEmpty {
			// Count how many are filled vs empty
			filled := 0
			empty := 0
			for di := colDemoStart; di <= colDemoEnd; di++ {
				if di < len(row) {
					if strings.TrimSpace(row[di]) != "" {
						filled++
					} else {
						empty++
					}
				}
			}
			// Only flag if it's a mix (not all filled, not all empty)
			// Some demographics might legitimately be 0.00% — so check if it's truly missing
			if filled > 0 && empty > 0 {
				// Check: are the "empty" ones actually 0.00%? If so that's fine.
				// The pattern is percentage strings like "42.95%"
				// Flag only truly empty (not "0.00%")
				trulyEmpty := 0
				for di := colDemoStart; di <= colDemoEnd; di++ {
					if di < len(row) && strings.TrimSpace(row[di]) == "" {
						trulyEmpty++
					}
				}
				if trulyEmpty > 0 && filled > 0 {
					demoViolations = append(demoViolations, violation{
						Row:     rowNum,
						Code:    uid,
						Message: fmt.Sprintf("Mixed demographics: %d filled, %d empty (btType=%s)", filled, trulyEmpty, bt),
					})
				}
			}
		}
	}

	// ============================================================
	// Consolidate trailing whitespace violations
	// ============================================================
	for ci, examples := range wsExamples {
		// Aggregate: count per column and show sample
		sampleRows := examples
		if len(sampleRows) > 5 {
			sampleRows = sampleRows[:5]
		}
		for _, ex := range sampleRows {
			trailingSpaceViolations = append(trailingSpaceViolations, violation{
				Row:     ex.Row,
				Code:    ex.Code,
				Message: fmt.Sprintf("[col %d, %d total occurrences] %s", ci+1, len(examples), ex.Message),
			})
		}
	}

	// ============================================================
	// Post-processing: Gender/Age inconsistencies
	// ============================================================
	for key, info := range gaTracker {
		if info.hasGenderSet && len(info.missingGRows) > 0 {
			exampleFilled := info.filledGRows
			exampleMissing := info.missingGRows
			if len(exampleFilled) > 3 {
				exampleFilled = exampleFilled[:3]
			}
			if len(exampleMissing) > 3 {
				exampleMissing = exampleMissing[:3]
			}
			genderAgeViolations = append(genderAgeViolations, violation{
				Row:  info.missingGRows[0],
				Code: key.parCode,
				Message: fmt.Sprintf("%s slot — candidate '%s': GENDER filled in %d rows (e.g. %v) but missing in %d rows (e.g. %v)",
					key.slotName, info.candidate, len(info.filledGRows), exampleFilled, len(info.missingGRows), exampleMissing),
			})
		}
		if info.hasAgeSet && len(info.missingARows) > 0 {
			exampleFilled := info.filledARows
			exampleMissing := info.missingARows
			if len(exampleFilled) > 3 {
				exampleFilled = exampleFilled[:3]
			}
			if len(exampleMissing) > 3 {
				exampleMissing = exampleMissing[:3]
			}
			genderAgeViolations = append(genderAgeViolations, violation{
				Row:  info.missingARows[0],
				Code: key.parCode,
				Message: fmt.Sprintf("%s slot — candidate '%s': AGE filled in %d rows (e.g. %v) but missing in %d rows (e.g. %v)",
					key.slotName, info.candidate, len(info.filledARows), exampleFilled, len(info.missingARows), exampleMissing),
			})
		}
	}

	// ============================================================
	// Postal vote count check (at most 1 per par constituency)
	// ============================================================
	var postalCountViolations []violation
	for parCode, rows := range postalCountByPar {
		if len(rows) > 1 {
			postalCountViolations = append(postalCountViolations, violation{
				Row:     rows[0],
				Code:    parCode,
				Message: fmt.Sprintf("Multiple postal vote rows (%d) for %s: rows %v", len(rows), parCode, rows),
			})
		}
	}

	// ============================================================
	// Check: Ordinary vote POLLING DISTRICT NAME should NOT be UNDI POS or UNDI AWAL
	// ============================================================
	var ordinaryVoteNameViolations []violation
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		bt := strings.TrimSpace(row[colBallotType])
		if bt == "ORDINARY VOTE" {
			pdn := strings.TrimSpace(row[colPollDistName])
			if pdn == "UNDI POS" || pdn == "UNDI AWAL" {
				ordinaryVoteNameViolations = append(ordinaryVoteNameViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("ORDINARY VOTE has POLLING DISTRICT NAME = '%s'", pdn),
				})
			}
			pdc := strings.TrimSpace(row[colPollDistCode])
			if strings.HasSuffix(pdc, "/POS") || strings.HasSuffix(pdc, "/00") {
				ordinaryVoteNameViolations = append(ordinaryVoteNameViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("ORDINARY VOTE has POLLING DISTRICT CODE = '%s' (ends with /POS or /00)", pdc),
				})
			}
		}
	}

	// ============================================================
	// Party name consistency within parliamentary constituency per slot
	// ============================================================
	type partySlotKey struct {
		parCode string
		slotIdx int
	}
	partyNameTracker := make(map[partySlotKey]map[string][]int) // key -> partyName -> rowNums
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		pc := strings.TrimSpace(row[colParCode])
		for si, start := range partySlotStarts {
			partyName := strings.TrimSpace(row[start])
			if partyName == "" {
				continue
			}
			key := partySlotKey{parCode: pc, slotIdx: si}
			if partyNameTracker[key] == nil {
				partyNameTracker[key] = make(map[string][]int)
			}
			partyNameTracker[key][partyName] = append(partyNameTracker[key][partyName], rowNum)
		}
	}
	var partyNameViolations []violation
	for key, names := range partyNameTracker {
		if len(names) > 1 {
			msg := fmt.Sprintf("Constituency %s, slot %s: multiple party names: ", key.parCode, partySlotNames[key.slotIdx])
			for name, rows := range names {
				exRows := rows
				if len(exRows) > 3 {
					exRows = exRows[:3]
				}
				msg += fmt.Sprintf("'%s' (%d rows, e.g. %v) ", name, len(rows), exRows)
			}
			partyNameViolations = append(partyNameViolations, violation{
				Row:     0,
				Code:    key.parCode,
				Message: msg,
			})
		}
	}

	// ============================================================
	// Candidate name consistency within par constituency per slot
	// ============================================================
	type candSlotKey struct {
		parCode string
		slotIdx int
	}
	candNameTracker := make(map[candSlotKey]map[string][]int)
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		pc := strings.TrimSpace(row[colParCode])
		for si, start := range partySlotStarts {
			candName := strings.TrimSpace(row[start+1])
			if candName == "" {
				continue
			}
			key := candSlotKey{parCode: pc, slotIdx: si}
			if candNameTracker[key] == nil {
				candNameTracker[key] = make(map[string][]int)
			}
			candNameTracker[key][candName] = append(candNameTracker[key][candName], rowNum)
		}
	}
	var candNameViolations []violation
	for key, names := range candNameTracker {
		if len(names) > 1 {
			msg := fmt.Sprintf("Constituency %s, slot %s: multiple candidate names: ", key.parCode, partySlotNames[key.slotIdx])
			for name, rows := range names {
				exRows := rows
				if len(exRows) > 3 {
					exRows = exRows[:3]
				}
				msg += fmt.Sprintf("'%s' (%d rows, e.g. %v) ", name, len(rows), exRows)
			}
			candNameViolations = append(candNameViolations, violation{
				Row:     0,
				Code:    key.parCode,
				Message: msg,
			})
		}
	}

	// ============================================================
	// Check: Empty TOTAL BALLOTS ISSUED
	// ============================================================
	var emptyRowViolations []violation
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		tb := strings.TrimSpace(row[colTotalBallots])
		if tb == "" {
			emptyRowViolations = append(emptyRowViolations, violation{
				Row:     rowNum,
				Code:    row[colUniqueCode],
				Message: "TOTAL BALLOTS ISSUED is empty",
			})
		}
	}

	// ============================================================
	// Check: VOTING CHANNEL NUMBER should be numeric
	// ============================================================
	var channelNumericViolations []violation
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		ch := strings.TrimSpace(row[colVotingChannel])
		if ch == "" {
			channelNumericViolations = append(channelNumericViolations, violation{
				Row:     rowNum,
				Code:    row[colUniqueCode],
				Message: "VOTING CHANNEL NUMBER is empty",
			})
		} else {
			_, err := strconv.Atoi(ch)
			if err != nil {
				channelNumericViolations = append(channelNumericViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("VOTING CHANNEL NUMBER is not numeric: '%s'", ch),
				})
			}
		}
	}

	// ============================================================
	// Check: Party slot has vote but no candidate (or vice versa)
	// ============================================================
	var slotConsistencyViolations []violation
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		for si, start := range partySlotStarts {
			partyVal := strings.TrimSpace(row[start])
			candVal := strings.TrimSpace(row[start+1])
			voteStr := strings.TrimSpace(row[start+4])

			hasParty := partyVal != ""
			hasCand := candVal != ""
			hasVote := voteStr != ""

			// If there's a vote but no candidate
			if hasVote && !hasCand {
				slotConsistencyViolations = append(slotConsistencyViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("%s slot: has vote '%s' but no candidate", partySlotNames[si], voteStr),
				})
			}
			// If there's a candidate but no party
			if hasCand && !hasParty {
				slotConsistencyViolations = append(slotConsistencyViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("%s slot: has candidate '%s' but no party label", partySlotNames[si], candVal),
				})
			}
			// If there's a party but no candidate
			if hasParty && !hasCand {
				slotConsistencyViolations = append(slotConsistencyViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("%s slot: has party '%s' but no candidate", partySlotNames[si], partyVal),
				})
			}
		}
	}

	// ============================================================
	// Check: PAR CODE format P.XXX
	// ============================================================
	var parCodeFormatViolations []violation
	parCodeRe := regexp.MustCompile(`^P\.\d+$`)
	for i, row := range dataRows {
		if len(row) < expectedCols {
			continue
		}
		rowNum := i + 2
		pc := strings.TrimSpace(row[colParCode])
		if !parCodeRe.MatchString(pc) {
			parCodeFormatViolations = append(parCodeFormatViolations, violation{
				Row:     rowNum,
				Code:    row[colUniqueCode],
				Message: fmt.Sprintf("PARLIAMENTARY CODE '%s' doesn't match P.XXX format", pc),
			})
		}
		// STATE CODE for non-postal should be N.XX
		bt := strings.TrimSpace(row[colBallotType])
		if bt != "POSTAL VOTE" {
			sc := strings.TrimSpace(row[colStateCode])
			scRe := regexp.MustCompile(`^N\.\d+$`)
			if !scRe.MatchString(sc) {
				parCodeFormatViolations = append(parCodeFormatViolations, violation{
					Row:     rowNum,
					Code:    row[colUniqueCode],
					Message: fmt.Sprintf("STATE CONSTITUENCY CODE '%s' doesn't match N.XX format (ballot=%s)", sc, bt),
				})
			}
		}
	}

	// ============================================================
	// GENERATE REPORT
	// ============================================================
	var sb strings.Builder
	sb.WriteString("# PHASE-5-REVIEW: Column Mapping Consistency\n\n")
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **File**: `to-review.csv`\n"))
	sb.WriteString(fmt.Sprintf("- **Header columns**: %d (expected %d)\n", headerColCount, expectedCols))
	sb.WriteString(fmt.Sprintf("- **Data rows**: %d\n", len(dataRows)))
	sb.WriteString(fmt.Sprintf("- **Parliamentary constituencies**: %d\n", len(parCodes)))
	sb.WriteString(fmt.Sprintf("- **Postal vote rows**: %d (across %d constituencies)\n", countPostalRows(postalCountByPar), len(postalCountByPar)))

	// Ballot type breakdown
	btCounts := map[string]int{}
	for _, row := range dataRows {
		if len(row) >= expectedCols {
			btCounts[strings.TrimSpace(row[colBallotType])]++
		}
	}
	sb.WriteString("\n### Ballot Type Breakdown\n\n")
	sb.WriteString("| Ballot Type | Count |\n")
	sb.WriteString("|-------------|-------|\n")
	for _, bt := range []string{"POSTAL VOTE", "EARLY VOTE", "ORDINARY VOTE"} {
		sb.WriteString(fmt.Sprintf("| %s | %d |\n", bt, btCounts[bt]))
	}

	// Demographics summary
	sb.WriteString("\n### Demographics Consistency\n\n")
	sb.WriteString("| Ballot Type | Rows with Demographics | Rows without Demographics |\n")
	sb.WriteString("|-------------|----------------------|-------------------------|\n")
	for _, bt := range []string{"POSTAL VOTE", "EARLY VOTE", "ORDINARY VOTE"} {
		sb.WriteString(fmt.Sprintf("| %s | %d | %d |\n", bt, demoFilledByBallotType[bt], demoEmptyByBallotType[bt]))
	}

	sb.WriteString("\n## Check Results\n\n")
	sb.WriteString("| # | Check | Violations | Status |\n")
	sb.WriteString("|---|-------|-----------|--------|\n")

	checks := []struct {
		name       string
		violations []violation
	}{
		{"Column count (all rows must have 92 columns)", colCountViolations},
		{"STATE must be 'SARAWAK'", stateViolations},
		{"BALLOT TYPE must be valid", ballotTypeViolations},
		{"Postal vote format rules", postalViolations},
		{"Postal vote count (max 1 per constituency)", postalCountViolations},
		{"Early vote format rules", earlyVoteViolations},
		{"VOTING CHANNEL matches UNIQUE CODE suffix", channelViolations},
		{"VOTING CHANNEL is numeric", channelNumericViolations},
		{"Gender/Age consistency per candidate", genderAgeViolations},
		{"CHECK ON VALID VOTES correctness", checkValidViolations},
		{"CHECK ON TOTAL VOTES ISSUED correctness", checkTotalViolations},
		{"UNIQUE CODE construction consistency", uniqueCodeViolations},
		{"Comma-formatted numbers in numeric fields", quotedNumberViolations},
		{"Ordinary vote naming/code crosscheck", ordinaryVoteNameViolations},
		{"Party name consistency per constituency/slot", partyNameViolations},
		{"Candidate name consistency per constituency/slot", candNameViolations},
		{"Empty TOTAL BALLOTS ISSUED", emptyRowViolations},
		{"Trailing/leading whitespace in key fields", trailingSpaceViolations},
		{"Demographics mixed fill within row", demoViolations},
		{"Party slot internal consistency (party/cand/vote)", slotConsistencyViolations},
		{"PAR CODE and STATE CODE format", parCodeFormatViolations},
	}

	totalViolations := 0
	passCount := 0
	failCount := 0
	for i, c := range checks {
		status := "✅ PASS"
		if len(c.violations) > 0 {
			status = "❌ FAIL"
			failCount++
		} else {
			passCount++
		}
		totalViolations += len(c.violations)
		sb.WriteString(fmt.Sprintf("| %d | %s | %d | %s |\n", i+1, c.name, len(c.violations), status))
	}

	sb.WriteString(fmt.Sprintf("\n**Total checks: %d | Passed: %d | Failed: %d | Total violations: %d**\n", len(checks), passCount, failCount, totalViolations))

	// ============================================================
	// Detailed findings
	// ============================================================
	sb.WriteString("\n---\n\n## Detailed Findings\n\n")

	for i, c := range checks {
		sb.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, c.name))
		if len(c.violations) == 0 {
			sb.WriteString("No violations found. ✅\n\n")
			continue
		}

		sb.WriteString(fmt.Sprintf("**%d violation(s) found.**\n\n", len(c.violations)))

		maxShow := 50
		if len(c.violations) > maxShow {
			sb.WriteString(fmt.Sprintf("_(Showing first %d of %d)_\n\n", maxShow, len(c.violations)))
		}

		sb.WriteString("| Row | UNIQUE CODE / Key | Details |\n")
		sb.WriteString("|-----|-------------------|--------|\n")
		shown := len(c.violations)
		if shown > maxShow {
			shown = maxShow
		}
		for j := 0; j < shown; j++ {
			v := c.violations[j]
			code := v.Code
			if len(code) > 55 {
				code = code[:52] + "..."
			}
			// Escape pipe chars in message
			msg := strings.ReplaceAll(v.Message, "|", "\\|")
			sb.WriteString(fmt.Sprintf("| %d | `%s` | %s |\n", v.Row, code, msg))
		}
		sb.WriteString("\n")
	}

	// ============================================================
	// Whitespace summary (separate section)
	// ============================================================
	if len(wsExamples) > 0 {
		sb.WriteString("---\n\n## Trailing/Leading Whitespace Summary\n\n")
		sb.WriteString("| Column # | Column Name | Occurrences | Sample Raw Value |\n")
		sb.WriteString("|----------|-------------|-------------|------------------|\n")
		// Sort by column index
		var colIdxs []int
		for ci := range wsExamples {
			colIdxs = append(colIdxs, ci)
		}
		sort.Ints(colIdxs)
		for _, ci := range colIdxs {
			exs := wsExamples[ci]
			colName := ""
			if ci < len(header) {
				colName = header[ci]
			}
			sampleRaw := ""
			if len(exs) > 0 {
				// Extract raw value from the message
				sampleRaw = fmt.Sprintf("'%s' (row %d)", exs[0].Code, exs[0].Row)
			}
			sb.WriteString(fmt.Sprintf("| %d | %s | %d | %s |\n", ci+1, colName, len(exs), sampleRaw))
		}
		sb.WriteString("\n")
	}

	// ============================================================
	// Recommendations
	// ============================================================
	sb.WriteString("\n---\n\n## Recommendations\n\n")

	recNum := 1
	if len(quotedNumberViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix comma-formatted numbers**: %d field(s) contain numbers like `\"1,616\"` instead of `1616`. These should be cleaned to plain integers to avoid CSV parsing issues.\n\n", recNum, len(quotedNumberViolations)))
		recNum++
	}
	if len(colCountViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix column count mismatches**: %d row(s) have incorrect column counts. This could indicate unquoted commas or structural issues.\n\n", recNum, len(colCountViolations)))
		recNum++
	}
	if len(checkValidViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Investigate CHECK ON VALID VOTES issues**: %d violation(s) found. Verify candidate vote sums against TOTAL VALID VOTES.\n\n", recNum, len(checkValidViolations)))
		recNum++
	}
	if len(checkTotalViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Investigate CHECK ON TOTAL VOTES ISSUED issues**: %d violation(s) found. Verify TOTAL BALLOTS = VALID + REJECTED + UNRETURNED.\n\n", recNum, len(checkTotalViolations)))
		recNum++
	}
	if len(genderAgeViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Normalize Gender/Age fields**: %d inconsistency(ies) found where some rows have gender/age for a candidate but others don't.\n\n", recNum, len(genderAgeViolations)))
		recNum++
	}
	if len(postalViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix postal vote format**: %d postal vote format issue(s) found.\n\n", recNum, len(postalViolations)))
		recNum++
	}
	if len(earlyVoteViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix early vote format**: %d early vote issue(s) found.\n\n", recNum, len(earlyVoteViolations)))
		recNum++
	}
	if len(channelViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix VOTING CHANNEL / UNIQUE CODE mismatches**: %d mismatch(es) found.\n\n", recNum, len(channelViolations)))
		recNum++
	}
	if len(uniqueCodeViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix UNIQUE CODE construction**: %d code(s) don't match the expected construction pattern.\n\n", recNum, len(uniqueCodeViolations)))
		recNum++
	}
	if len(partyNameViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Party name inconsistencies**: %d constituency/slot combination(s) have multiple party names.\n\n", recNum, len(partyNameViolations)))
		recNum++
	}
	if len(candNameViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Candidate name inconsistencies**: %d constituency/slot combination(s) have multiple candidate names.\n\n", recNum, len(candNameViolations)))
		recNum++
	}
	if len(ordinaryVoteNameViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Ordinary vote naming issues**: %d ordinary vote row(s) have incorrect POLLING DISTRICT CODE/NAME.\n\n", recNum, len(ordinaryVoteNameViolations)))
		recNum++
	}
	if len(trailingSpaceViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Trim trailing/leading whitespace**: %d instance(s) of trailing/leading whitespace found in key fields. While mostly cosmetic, this can cause matching issues.\n\n", recNum, len(trailingSpaceViolations)))
		recNum++
	}
	if len(demoViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix demographics inconsistencies**: %d row(s) have partial demographics (some filled, some empty within the same row).\n\n", recNum, len(demoViolations)))
		recNum++
	}
	if len(slotConsistencyViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix party slot consistency**: %d row(s) have a party/candidate/vote mismatch within a party slot.\n\n", recNum, len(slotConsistencyViolations)))
		recNum++
	}
	if len(parCodeFormatViolations) > 0 {
		sb.WriteString(fmt.Sprintf("%d. **Fix PAR/STATE CODE format**: %d row(s) have codes that don't match the expected format.\n\n", recNum, len(parCodeFormatViolations)))
		recNum++
	}

	if recNum == 1 {
		sb.WriteString("No issues found — all checks passed! 🎉\n")
	}

	// Write report
	err = os.WriteFile("PHASE-5-REVIEW.md", []byte(sb.String()), 0644)
	if err != nil {
		slog.Error("cannot write report", "err", err)
		os.Exit(1)
	}

	slog.Info("report written", "file", "PHASE-5-REVIEW.md")

	// Print summary to stdout
	fmt.Println("=== Phase 5 Validation Summary ===")
	for _, c := range checks {
		status := "PASS"
		if len(c.violations) > 0 {
			status = fmt.Sprintf("FAIL (%d)", len(c.violations))
		}
		fmt.Printf("  %-55s %s\n", c.name, status)
	}
	fmt.Printf("\n  Total: %d checks, %d passed, %d failed, %d total violations\n", len(checks), passCount, failCount, totalViolations)
}

func parseIntSafe(s string) int {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func countPostalRows(m map[string][]int) int {
	total := 0
	for _, rows := range m {
		total += len(rows)
	}
	return total
}
