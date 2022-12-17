package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func processPart1(par string, salurans []saluran) {
	fmt.Println("IN processPart1!!!!")
	fmt.Println("PAR:", par)
	parID := par[1:]
	fmt.Println("RESULTS: ", parID)

	// UNIQUE CODE,	STATE,	BALLOT TYPE,
	// PARLIAMENTARY CODE,	PARLIAMENTARY NAME
	// STATE CONSTITUENCY CODE,	STATE CONSTITUENCY NAME
	// Output CSV
}

// processCandidates implements the candidates; ensuring consitency per state
func processCandidates(par string, salurans []saluran, candidates map[string]candidate) {
	// DEBUG
	//fmt.Println("IN processCandidates************ ")
	//fmt.Println("PAR:", par)
	parID := par[1:]
	//fmt.Println("RESULTS: ", parID)

	data := make([][]string, 0)
	for _, saluran := range salurans {
		fmt.Println("DUN_ID:", saluran.dunID)
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
			singleRow[mult+2] = "2020" // Age to be replaced later ..
			singleRow[mult+3] = candidates[keyID].gender
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
	outputCSV(par, data)
}

func outputCSV(fileName string, records [][]string) {
	// DEBUG
	//records = [][]string{
	//	{"first_name", "last_name", "username"},
	//	{"Rob", "Pike", "rob"},
	//	{"Ken", "Thompson", "ken"},
	//	{"Robert", "Griesemer", "gri"},
	//}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
