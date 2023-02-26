package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
)

// FuzzyDownloadCandidatePerPAR go from raw PAR to PAR_ID
func FuzzyDownloadCandidatePerPAR(state string, pars []string) {
	// Load all Results ..
	candidatesPAR := LookupResults(state)

	// For each PAR
	for _, par := range pars {
		// TODO: Only do something if folder does NOT exist?
		// Derive PAR_ID
		parID := fmt.Sprintf("%s00", par[1:])
		fmt.Println("PAR:", parID)
		candidatesInPAR := candidatesPAR[parID]
		// Fuzzy download below ...
		spew.Dump(candidatesInPAR)
	}
}

func fuzzySearch(parID, fullName string) {
	// For each PAR; use the components to fuzzy match possibility via Google ..
	possibleCandidatesQuery := ""
	script.Get(possibleCandidatesQuery)

}
