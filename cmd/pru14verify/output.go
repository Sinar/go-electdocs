package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func processPrefix(par string, salurans []saluran) {
	fmt.Println("IN processPrefix!!!!")
	fmt.Println("PAR:", par)
	parID := par[1:]
	fmt.Println("RESULTS: ", parID)

	// UNIQUE CODE,	STATE,	BALLOT TYPE,
	// PARLIAMENTARY CODE,	PARLIAMENTARY NAME
	// STATE CONSTITUENCY CODE,	STATE CONSTITUENCY NAME
	// Output CSV
}

// processSuffix has the end portion: Independents + TOTAL VALID VOTES	TOTAL REJECTED VOTES	TOTAL UNRETURNED BALLOTS
func processSuffix(par string, salurans []saluran, candidates map[string]candidate) [][]string {
	fmt.Println("IN processPrefix!!!!")
	fmt.Println("PAR:", par)
	parID := par[1:]
	fmt.Println("RESULTS: ", parID)

	// UNIQUE CODE,	STATE,	BALLOT TYPE,
	// PARLIAMENTARY CODE,	PARLIAMENTARY NAME
	// STATE CONSTITUENCY CODE,	STATE CONSTITUENCY NAME
	// Output CSV
	data := make([][]string, 0)
	for _, saluran := range salurans {
		fmt.Println("DUN_ID:", saluran.dunID)
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
			singleRow[mult+3] = "2020" // Age to be replaced later ..
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
