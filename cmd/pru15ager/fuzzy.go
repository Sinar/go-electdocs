package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"os"
	"regexp"
	"strings"
)

// fuzzyCandidate represents the RAW Calons
type fuzzyCandidate struct {
	url               string
	candidateRawName  string
	candidateRawAge   string
	candidateRawParty string
}

// FuzzyDownloadCandidatePerPAR go from raw PAR to PAR_ID
func FuzzyDownloadCandidatePerPAR(state string, pars []string) {
	// Load all Results ..
	candidatesPAR := LookupResults(state)
	// For each PAR
	for _, par := range pars {
		// Open par raw HTML file and cache the calons ..
		calons, err := script.Exec(fmt.Sprintf("grep -i calon testdata/%s.html", par)).Match("/calon").Slice()
		if err != nil {
			panic(err)
		}
		// DEBUG
		//spew.Dump(calons)
		cacheCandidates(par, calons)
		// Derive PAR_ID
		parID := fmt.Sprintf("%s00", par[1:])
		fmt.Println("PAR:", parID)
		candidatesInPAR := candidatesPAR[parID]
		// DEBUG
		//spew.Dump(candidatesInPAR)
		processCandidates(parID, &candidatesInPAR)
		// Final output below ..
		// TOOD: OUtput ..
		// For each mapKey; dump it all out!
		//spew.Dump(candidateData)
		// WriteTo CSV: <state>-candidates.csv
		// parID/ballotID, Name, MatchedName, Age..
		//outputCSV(fmt.Sprintf("testdata/%s-candidates.csv", state), candidateData)

	}
}

func processCandidates(parID string, candidates *[]candidate) {
	// Load all files in directory .. intp fuzzyCandidate
	// Get filename .. is calon ..
	// for each candidate .. pick closest leveschutein ,.
	//	extract age
	//	extract party ..

	// for each candidate; leveschite distance of calon ..
	// try exact match first ..
	for _, c := range *candidates {
		fmt.Println("NAME:", c.name) // NAME is the final lookup key ..
		// Get Age + Party ..
	}
}

// cacheCandidates will download if file not there ..
func cacheCandidates(par string, calons []string) {
	fmt.Println("PAR:", par)
	baseURL := "https://pru.sinarharian.com.my"
	// PAR == P123 --> 123
	code := par[1:]
	parID := fmt.Sprintf("%s00", code)
	dataPath := fmt.Sprintf("testdata/%s", parID)
	/*
		([]string) (len=3 cap=4) {
		 (string) (len=153) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_0\" class=\"calon\" href=\"/calon/6151/shamsul-iskandar-mohd-akin\">",
		 (string) (len=141) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_1\" class=\"calon\" href=\"/calon/858/mohd-ali-rustam\">",
		 (string) (len=143) "<a id=\"ContentPlaceHolder1_uscKeputusanParlimen_rptParlimen_rptParlimenCalon_0_lnkToCalon_2\" class=\"calon\" href=\"/calon/2820/md-khalid-kassim\">"
		}
	*/
	// Create data folder
	merr := os.MkdirAll(dataPath, 0755)
	if merr != nil {
		panic(merr)
	}
	// Go through each candidate .. with these pattern
	rexp := regexp.MustCompile("^<a.*href=\"(.+/(.+?))\".*$")
	replaceTemplate := "$1,$2"
	for _, calon := range calons {
		// Pattern assumes start of A tag .. trim any stray ..
		calon = strings.TrimSpace(calon)
		s := rexp.ReplaceAllString(calon, replaceTemplate)
		// DEBUG
		//fmt.Println("LINE: ", s)
		// Extract out from template ..
		/*
			(string) (len=56) "/calon/288/mas-ermieyati-samsudin,mas-ermieyati-samsudin"
			(string) (len=37) "/calon/6219/nasir-othman,nasir-othman"
			(string) (len=43) "/calon/6217/sabirin-ja`afar,sabirin-ja`afar"
			(string) (len=49) "/calon/6066/mohd-redzuan-yusof,mohd-redzuan-yusof"
		*/
		c := strings.Split(s, ",")
		if len(c) != 2 {
			fmt.Println("UNEXPECTED calon! Skipping ..")
			spew.Dump(c)
			continue
		}
		candidateURL := baseURL + c[0]
		safeName := c[1]
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
	}
}
