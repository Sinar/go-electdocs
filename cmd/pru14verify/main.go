package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {

	fmt.Println("Welcome to PRU14 Verifier!")
	// DEBUG
	//LookupState()
	//LookupParty()
	//LookupPAR()
	// DEBUG
	//pars := []string{"P004"}
	//pars := []string{"P004", "P005"}
	pars := []string{
		"P004", "P005", "P006", "P007", "P008", "P009",
		"P010", "P011", "P012", "P013", "P014", "P015",
		"P016", "P017", "P018",
	}

	AssembleResultsPerPAR(pars)
}

func AssembleResultsPerPAR(pars []string) {
	candidates := LookupResults("KEDAH")
	// DEBUG
	//spew.Dump(candidates)

	// For each PAR ID; extract out the first section
	// lookup from Results; which is ordered by SaluranID: 0,1,2
	// attach to the candidateVotes
	// For each pars; extract the number; remove the P
	// Look up the candidates; ordered by the BallotID
	var wg sync.WaitGroup
	for _, par := range pars {
		//SanityTestPARSaluran(par)
		salurans := LoadPARSaluran(par)
		wg.Add(1)
		go func() {
			processPart1(par, salurans)
			wg.Done()
		}()
		go func() {
			processCandidates(par, salurans, candidates)
			wg.Done()
		}()
		// Block till done ..
		wg.Wait()
		break
	}
}

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

func processCandidates(par string, salurans []saluran, candidates map[string]candidate) {
	fmt.Println("IN processCandidates************ ")
	fmt.Println("PAR:", par)
	parID := par[1:]
	fmt.Println("RESULTS: ", parID)

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
		for i, votes := range saluran.candidateVotes {
			keyID := fmt.Sprintf("%s00/%d", parID, i+1)
			fmt.Println("LOOkUP:", keyID)
			//spew.Dump(candidates[keyID])
			fmt.Println("DM:", saluran.ID, "VOTE FOR:", candidates[keyID].name, "VOTES:", votes)
		}
		// Saluran Enhanced - PART1
		// GROUPED Candidates - PART2
		// Suffix - PART3

		// Map[COALITION] ==> votes; now ordered in the state-wide COLs
	}
	// Write the output of CSV ..
	singleRow := []string{"", "", ""}
	data := [][]string{singleRow}
	outputCSV(par, data)
}

func outputCSV(par string, records [][]string) {
	records = [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
