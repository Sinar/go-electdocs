package main

import (
	"fmt"
	googlesearch "github.com/rocketlaunchr/google-search"
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
		//spew.Dump(candidatesInPAR)
		for _, candidate := range candidatesInPAR {
			fuzzySearch(parID, candidate.name)
		}
	}
}

func fuzzySearch(parID, fullName string) {
	// For each PAR; use the components to fuzzy match possibility via Google ..

	fmt.Println(googlesearch.Search(nil, "cars for sale in Toronto, Canada"))

	//baseURL := "https://www.google.com.sg/search"
	//q := url.QueryEscape(fullName + " \"pru14\"" + " site:pru.sinarharian.com.my")
	//possibleCandidatesQuery := fmt.Sprintf("%s?q=%s", baseURL, q)
	//fmt.Println("GET: ", possibleCandidatesQuery)
	//_, err := script.Get(possibleCandidatesQuery).Stdout()
	//if err != nil {
	//	panic(err)
	//}
	//script.Get(possibleCandidatesQuery).FilterLine(func(s string) string {
	//	// Data
	//	fmt.Println(s)
	//	// Look for A tag .. with calon in it ..
	//	if strings.Contains(s, "calon") {
	//		fmt.Println("LOOK HERE: ", s)
	//	}
	//	return ""
	//})

}
