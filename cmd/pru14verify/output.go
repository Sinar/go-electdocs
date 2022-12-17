package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// processPrefix
func processPrefix(par string, salurans []saluran) [][]string {
	// DEBUG
	//fmt.Println("IN processPrefix!!!!")
	//fmt.Println("PAR:", par)
	parID := fmt.Sprintf("%s00", par[1:])
	fmt.Println("RESULTS: ", parID)

	// Prefix FORMAT
	// UNIQUE CODE	STATE	BALLOT TYPE
	// PARLIAMENTARY CODE	PARLIAMENTARY NAME
	// STATE CONSTITUENCY CODE	STATE CONSTITUENCY NAME
	// POLLING DISTRICT CODE	POLLING DISTRICT NAME	POLLING CENTRE
	// VOTING CHANNEL NUMBER	TOTAL BALLOTS ISSUED
	// Output CSV
	data := make([][]string, 0)
	for _, saluran := range salurans {
		// DEBUG
		//fmt.Println("DUN_ID:", saluran.dunID)
		singleRow := make([]string, 12)
		// UNIQUE CODE - <PAR_CODE>_<DUN_CODE>_<DM_ID>_<SALURAN_ID>
		//P.041_P.41/POSTAL VOTE_UNDI POS_1
		//P.041_P.41/EARLY VOTE_041/02/00_1
		//P.041_P.41/EARLY VOTE_041/02/00_2
		//P.041_N.01_041/01/01_1
		//P.041_N.01_041/01/01_2
		parCode := lookupPAR[parID].code
		dunCode := lookupDUN[saluran.dunID].code
		dunName := lookupDUN[saluran.dunID].name
		if saluran.ID == "UNDI POS" {
			dunName = "POSTAL VOTE"
			dunCode = fmt.Sprintf("%s/%s", parCode, dunName)
		}
		// If Early Vote (end with 00)  -
		if saluran.ID[len(saluran.ID)-2:] == "00" {
			// dunCode = P.41/EARLY VOTE
			// dunName = EARLY VOTE
			dunName = "EARLY VOTE"
			dunCode = fmt.Sprintf("%s/%s", parCode, dunName)
		}
		uniqueCode := fmt.Sprintf("%s_%s_%s_%s",
			parCode,
			dunCode,
			saluran.ID,
			saluran.saluranID,
		)
		singleRow[0] = uniqueCode
		// STATE	BALLOT TYPE
		singleRow[1] = "" // just paste manually
		// Translate to English
		ballotType := saluran.voteType
		switch ballotType {
		case "UNDI BIASA":
			ballotType = "NORMAL VOTE"
		case "UNDI AWAL":
			ballotType = "EARLY VOTE"
		case "UNDI POS":
			ballotType = "POSTAL VOTE"
		default:
			ballotType = "UNKNOWN"
		}
		singleRow[2] = ballotType
		// PARLIAMENTARY CODE	PARLIAMENTARY NAME
		singleRow[3] = parCode
		singleRow[4] = lookupPAR[parID].name
		// STATE CONSTITUENCY CODE	STATE CONSTITUENCY NAME
		singleRow[5] = dunCode
		singleRow[6] = dunName
		// DM related
		// POLLING DISTRICT CODE	POLLING DISTRICT NAME	POLLING CENTRE
		// VOTING CHANNEL NUMBER	TOTAL BALLOTS ISSUED
		singleRow[7] = saluran.ID
		singleRow[8] = saluran.daerahMengundi
		singleRow[9] = saluran.votingLocationName
		singleRow[10] = saluran.saluranID
		singleRow[11] = saluran.totalVotesIssued
		data = append(data, singleRow)

	}
	return data
}

// processSuffix has the end portion: Independents + TOTAL VALID VOTES	TOTAL REJECTED VOTES	TOTAL UNRETURNED BALLOTS
func processSuffix(par string, salurans []saluran, candidates map[string]candidate) [][]string {
	// DEBUG
	//fmt.Println("IN processPrefix!!!!")
	//fmt.Println("PAR:", par)
	parID := par[1:]
	//fmt.Println("RESULTS: ", parID)

	// UNIQUE CODE,	STATE,	BALLOT TYPE,
	// PARLIAMENTARY CODE,	PARLIAMENTARY NAME
	// STATE CONSTITUENCY CODE,	STATE CONSTITUENCY NAME
	// Output CSV
	data := make([][]string, 0)
	for _, saluran := range salurans {
		// DEBUG
		//fmt.Println("DUN_ID:", saluran.dunID)
		// Loop through candidates, fill in the know order
		// One Candidate: PARTY, NAME, AGE, GENDER, VOTES
		singleRow := make([]string, 15)
		for i, votes := range saluran.candidateVotes {
			keyID := fmt.Sprintf("%s00/%d", parID, i+1)
			if candidates[keyID].party != "IND" {
				fmt.Println("Skipping non-IND")
				continue
			}
			singleRow[0] = candidates[keyID].party
			singleRow[1] = candidates[keyID].name
			singleRow[2] = candidates[keyID].gender
			singleRow[3] = "2020" // Age to be replaced later ..
			singleRow[4] = votes
			// If more than 1 IND ..
		}
		// Everything else..
		// TOTAL VALID VOTES
		// TOTAL REJECTED VOTES
		// TOTAL UNRETURNED BALLOTS
		singleRow[10] = saluran.totalVotesCast
		singleRow[11] = saluran.totalVotesSpoilt
		singleRow[12] = saluran.totalVotesNotCast
		data = append(data, singleRow)
	}
	return data
}

// processCandidates implements the candidates; ensuring consitency per state
func processCandidates(par string, salurans []saluran, candidates map[string]candidate) [][]string {
	// DEBUG
	//fmt.Println("IN processCandidates************ ")
	//fmt.Println("PAR:", par)
	parID := par[1:]
	//fmt.Println("RESULTS: ", parID)

	data := make([][]string, 0)
	for _, saluran := range salurans {
		// DEBUG
		//fmt.Println("DUN_ID:", saluran.dunID)
		// DUN_ID == "UNDI POS"
		// Passed in the slice of votes as a Lookup method
		// or output method for that row .. as csv?
		// Each COALTION will fit into a certain order in the slice;
		//	known a-priori per STATE
		// If COALITION not there; leave as blank
		//keyID := fmt.Sprintf("%s/%s", parID, saluran.saluranID)
		//fmt.Println("LOOkUP:", keyID)

		// Loop through candidates, fill in the know order
		// One Candidate: PARTY, NAME, AGE, GENDER, VOTES
		singleRow := make([]string, 20)
		for i, votes := range saluran.candidateVotes {
			keyID := fmt.Sprintf("%s00/%d", parID, i+1)
			// DEBUG
			//fmt.Println("LOOkUP:", keyID)
			//spew.Dump(candidates[keyID])
			//fmt.Println("DM:", saluran.ID, "VOTE FOR:", candidates[keyID].name, "VOTES:", votes)
			if candidates[keyID].party == "IND" {
				fmt.Println("Skipping IND")
				continue
			}

			var mult int
			switch candidates[keyID].coalition {
			case "BN":
				mult = 0
			case "PH":
				mult = 5
			case "GS":
				mult = 10
			default:
				mult = 15
			}
			singleRow[mult] = candidates[keyID].party
			singleRow[mult+1] = candidates[keyID].name
			singleRow[mult+2] = candidates[keyID].gender
			singleRow[mult+3] = "" // Age to be replaced later ..
			singleRow[mult+4] = votes
		}
		// Saluran Enhanced - PART1
		// GROUPED Candidates - PART2
		// Suffix - PART3
		//spew.Dump(singleRow)
		data = append(data, singleRow)
		// Map[COALITION] ==> votes; now ordered in the state-wide COLs
	}
	// Column for Candidates:
	// BN | PH | GS | OTHERS | IND_KEY
	// Write the output of CSV ..
	outputCSV(fmt.Sprintf("testdata/%s-candidates.csv", par), data)

	return data
}

func outputCSV(fileName string, records [][]string) {
	// DEBUG
	//records = [][]string{
	//	{"first_name", "last_name", "username"},
	//	{"Rob", "Pike", "rob"},
	//	{"Ken", "Thompson", "ken"},
	//	{"Robert", "Griesemer", "gri"},
	//}
	var iow io.Writer
	var oerr error
	if fileName == "DEBUG" {
		iow = os.Stdout
	} else {
		iow, oerr = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0744)
		if oerr != nil {
			panic(oerr)
		}
	}
	w := csv.NewWriter(iow)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
