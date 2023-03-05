package main

import (
	"fmt"
	"github.com/bitfield/script"
	googlesearch "github.com/rocketlaunchr/google-search"
	"golang.org/x/time/rate"
	"net/url"
	"os"
	"strings"
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
	// Setup data folder if not already existing ..
	dataPath := fmt.Sprintf("testdata/%s", parID)
	// Create data folder
	merr := os.MkdirAll(dataPath, 0755)
	if merr != nil {
		panic(merr)
	}
	// Will search only if no lock already ..
	// For each PAR; use the components to fuzzy match possibility via Google ..
	// Limit to 5 per sec; 10 burstable ..
	googlesearch.RateLimit = rate.NewLimiter(1.0, 5)
	sopt := googlesearch.SearchOptions{
		CountryCode: "my",
		Limit:       2,
		OverLimit:   true,
	}
	q := url.QueryEscape(fullName + " \"pru14\"" + " site:pru.sinarharian.com.my")
	fmt.Println("SEARCH_QUERY: ", q)
	res, err := googlesearch.Search(nil, q, sopt)
	if err != nil {
		panic(err)
	}
	if len(res) == 0 {
		fmt.Println("No result ... :(")
		return
	}
	// DEBUG
	//spew.Dump(res)
	// Got at least 2 result; just take the first one ..
	candidateURL := res[0].URL
	if !strings.Contains(candidateURL, "calon") {
		// If no calon; then try backup ..
		fmt.Println("URL did not have calon! ", candidateURL)
		candidateURL = res[0].URL
	}
	safeName := strings.ReplaceAll(strings.ToLower(fullName), " ", "-")
	// Cannot have something looking like UNIX filepath ..
	if strings.Contains(safeName, "/") {
		safeName = strings.ReplaceAll(safeName, "/", "")
	}
	fmt.Println("Safe file as ", safeName)
	candidatePath := fmt.Sprintf("%s/%s.html", dataPath, safeName)
	// DEBUG
	//spew.Dump(s)
	// Get file and save it ..
	// If the file exists; ignore it again??
	// testdata/<PAR>/...html
	if script.IfExists(candidatePath).Error() != nil {
		// has error; means the file does not exist!
		fmt.Println("GOTTA DOWNLOAD!!!", candidateURL, "INTO", candidatePath)
		n, err := script.Get(candidateURL).WriteFile(candidatePath)
		if err != nil {
			panic(err)
		}
		fmt.Println("N:", n)
	} else {
		fmt.Println("FOUND! at", candidatePath)
	}
	// OK done ..
	return
}

// ExactCandidatePerPAR no need to match as we know the file exactly; just a transform ..
func ExactCandidatePerPAR(state string, pars []string) {
	// Load all Results ..
	candidatesPAR := LookupResults(state)
	// DEBUG
	//spew.Dump(candidatesPAR)
	// maybe no need
	//mapCandidate = make(map[string][]candidate, len(pars))
	candidateData := make([][]string, 0)
	// Header ..
	row := make([]string, 6)
	row[0] = "ID"
	row[1] = "NAME"
	row[2] = "AGE"
	row[3] = "MATCH_NAME"
	row[4] = "RAW_AGE"
	row[5] = "MATCH_URL"
	candidateData = append(candidateData, row)
	// For each PAR
	for _, par := range pars {
		// Derive PAR_ID
		parID := fmt.Sprintf("%s00", par[1:])
		fmt.Println("PAR:", parID)
		// Data folder is function of the parID
		dataPath := fmt.Sprintf("testdata/%s", parID)

		// Each Row has Column
		// PAR_ID, Name, Age, MatchedName, RawAge, MatchedURL
		for i, c := range candidatesPAR[parID] {
			safeName := strings.ReplaceAll(strings.ToLower(c.name), " ", "-")
			// Cannot have something looking like UNIX filepath ..
			if strings.Contains(safeName, "/") {
				safeName = strings.ReplaceAll(safeName, "/", "")
			}
			// matchURL only if the file exists ..
			candidatePath := fmt.Sprintf("%s/%s.html", dataPath, safeName)
			fmt.Println("CHECK:", candidatePath)
			if script.IfExists(candidatePath).Error() != nil {
				fmt.Println("SKIP:", candidatePath)
				// If the calon file does not exist; ensure this is empty ..
				candidatePath = ""
			}
			row := make([]string, 6)
			row[0] = fmt.Sprintf("%s/%d", parID, i+1)
			row[1] = c.name
			row[2] = c.age
			row[3] = safeName
			row[4] = c.matchRawAge
			row[5] = candidatePath
			candidateData = append(candidateData, row)
		}
		//break
	}

	// For each mapKey; dump it all out!
	//spew.Dump(candidateData)
	// WriteTo CSV: <state>-candidates.csv
	// parID/ballotID, Name, MatchedName, Age..
	outputCSV(fmt.Sprintf("testdata/%s-candidates.csv", state), candidateData)
}
