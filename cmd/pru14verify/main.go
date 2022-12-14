package main

import (
	"fmt"
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
	//LookupResults("KEDAH")
	// For each PAR ID; extract out the first section
	// lookup from Results; which is ordered by SaluranID: 0,1,2
	// attach to the candidateVotes
	// For each pars; extract the number; remove the P
	// Look up the candidates; ordered by the BallotID
	for _, par := range pars {
		//SanityTestPARSaluran(par)
		salurans := LoadPARSaluran(par)
		fmt.Println("PAR:", par)
		parID := par[1:]
		fmt.Println("RESULTS: ", parID)
		for _, saluran := range salurans {
			// Passed in the slice of votes as a Lookup method
			// or output method for that row .. as csv?
			// Each COALTION will fit into a certain order in the slice;
			//	known a-priori per STATE
			// If COALITION not there; leave as blank
			keyID := fmt.Sprintf("%s/%s", parID, saluran.saluranID)
			fmt.Println("LOOkUP:", keyID)
		}

		break
	}
}
