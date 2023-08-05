package main

import (
	"fmt"
	"github.com/bitfield/script"
	"regexp"
	"strconv"
	"strings"
)

// ExtractAgeParty prints out DUN; find info for them ..
func ExtractAgeParty(state string) {
	var endID int
	// State will fix the PAR
	var mapCandidate map[string][]string

	switch state {
	case "N9":
		endID = 36
		mapCandidate = LookUpN9Candidate()
	case "PENANG":

	case "SELANGOR":

	default:
		panic("INVALID STATE!!")

	}

	dunID := ""
	for i := 1; i <= endID; i++ {
		dunID = fmt.Sprintf("N%02d", i)
		fmt.Println("DUNID: ", dunID)
		// DEBUG
		//spew.Dump(mapCandidate[dunID])
		for _, candidateURL := range mapCandidate[dunID] {
			u := strings.Split(candidateURL, "/")
			shortname := u[len(u)-1]
			// DEBUG
			//fmt.Println("SHORT_NAME: ", shortname)
			// If file does not exist ..
			candidatePath := fmt.Sprintf("testdata/%s-%s", state, shortname)
			if script.IfExists(candidatePath).Error() != nil {
				// has error; means the file does not exist!
				fmt.Println("GOTTA DOWNLOAD!!!", candidateURL, "INTO", candidatePath)
				n, err := script.Get(candidateURL).WriteFile(candidatePath)
				if err != nil {
					panic(err)
				}
				fmt.Println("N:", n)
			} else {
				// DEBUG
				//fmt.Println("FOUND! at", candidatePath)
			}

			// Now can apply fuzzy ..
			candidateRawParty, party := findCandidateRawParty(candidatePath)
			fmt.Println(fmt.Sprintf("CANDIDATE: %s PARTY: %s RAW: %s", shortname, party, candidateRawParty))
		}
		if i > 1 {
			break
		}
	}
}

func findCandidateRawParty(filePath string) (candidateRawParty, party string) {
	rexp := regexp.MustCompile("<span .+>(.+)</span>.*$")
	replaceTemplate := "$1"
	// Look for party / parti?
	// Should it look for top level component parties?
	partyMatches, perr := script.File(filePath).Match("PartiKomponen").ReplaceRegexp(rexp, replaceTemplate).Slice()
	if perr != nil {
		panic(perr)
	}
	candidateRawParty = strings.TrimSpace(partyMatches[0])
	party = candidateRawParty
	// Re-map exceptions
	switch party {
	case "AMANAH":
		party = "PAN"
	case "PPBM/BERSATU":
		party = "PPBM"
	}
	// DEBUG
	//spew.Dump(party)
	return candidateRawParty, party
}

func findCandidateRawAge(filePath string) (candidateRawAge, age, url string) {
	age = "2023" // PRU15 DUN is on 2023 ..
	// Used multopleplaces ..
	replaceTemplate := "$1"

	// Extract metadata from the content
	reURL := regexp.MustCompile("^.+content=\"(.+)\".+$")
	urlMatches, uerr := script.File(filePath).Match("og:url").ReplaceRegexp(reURL, replaceTemplate).Slice()
	if uerr != nil {
		panic(uerr)
	}

	if len(urlMatches) == 0 {
		panic(urlMatches)
	} else if len(urlMatches) > 0 {
		// DEBUG
		fmt.Println(">>>>>>>>>>>>>>>>>>>>> URL:", urlMatches[0])
		for _, urlMatch := range urlMatches {
			// Only take URL that has http!
			if strings.Contains(urlMatch, "http") {
				url = urlMatch
			}
		}
	}
	// If find DOB; extract and leave first!
	// DOB pattern "ContentPlaceHolder1_lblDob"
	// Pattern DD/M/YYYY e.g. 18/1/1967
	reDOB := regexp.MustCompile("^.+\\d+/\\d+/(\\d+).+$")
	dobMatches, derr := script.File(filePath).Match("ContentPlaceHolder1_lblDob").ReplaceRegexp(reDOB, replaceTemplate).Slice()
	if derr != nil {
		panic(derr)
	}
	if len(dobMatches) > 0 {
		// Not needed quite useless
		//mc.candidateRawAge = append(mc.candidateRawAge, dobMatches...)
		candidateRawAge = dobMatches[0]
		// DEBUG
		//fmt.Println("DATA_DOB:")
		//spew.Dump(dobMatches)
		year, cerr := strconv.Atoi(dobMatches[0])
		if cerr != nil {
			panic(cerr)
		}
		// DEBUG
		//fmt.Println("YEAR_BIRTH:", year)
		//fmt.Println("DEBUG_AGE:", 2022-year)
		age = strconv.Itoa(2022 - year)
		return candidateRawAge, age, url
	}
	// If cannot find DOB; try a more generic search; add the findings?
	re := regexp.MustCompile("^.+(\\d{2})\\s+.+tahun.*$")
	matches, err := script.File(filePath).MatchRegexp(re).ReplaceRegexp(re, replaceTemplate).Slice()
	if err != nil {
		panic(err)
	}
	//if len(matches) > 0 {
	//	mc.candidateRawAge = append(mc.candidateRawAge, matches...)
	//	// Check should at least be 21
	//}
	if len(matches) > 0 {
		// NOTE: Below are unsure
		fmt.Println("FUZZY search for: ", filePath, "SOURCE:", url)
		for _, match := range matches {
			if match == "" {
				continue
			}
			candidateRawAge = match
			// DEBUG
			//fmt.Println("DATA_MATCHES:")
			//spew.Dump(matches)
			possibleAge, cerr := strconv.Atoi(match)
			if cerr != nil {
				panic(cerr)
			}
			if age != "2022" {
				fmt.Println("POTENTIAL CONFLICT:", url, "PREV:", age, "NEW:", possibleAge)
			}
			if possibleAge < 21 {
				fmt.Println("IMPOSSIBLE: AGE MUST >21", possibleAge)
				age = ""
			} else if possibleAge > 100 {
				fmt.Println("IMPOSSIBLE: AGE MUST <100", possibleAge)
				age = ""
			} else {
				age = strconv.Itoa(possibleAge)
			}
		}
		return candidateRawAge, age, url
	}
	// Default is zero value .. if find nothing ...
	return "", "", url
}
