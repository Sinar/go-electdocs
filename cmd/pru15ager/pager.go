package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type candidate struct {
	name string
	url  string
	par  string
	dun  string
	code string
	age  string // COuld be date pattern or word with tahun ..
}

type matchCandidate struct {
	code            string
	name            string
	candidateDir    string
	matchedFileName string
	candidateAge    []string
}

func Run() {
	fmt.Println("Runing ..")
	//getCandidates()
	//downloadCandidates()
	extractCandidates()
}

func extractCandidates() {
	//duns := []string{"perlis"}
	duns := []string{"perak", "pahang"}
	for _, dunName := range duns {
		// FInal data
		var finalData []string
		// Open the data from scaper
		r, err := script.File("testdata/pru15-" + dunName + "-final.csv").Slice()
		if err != nil {
			panic(err)
		}
		for _, line := range r {
			cols := strings.Split(line, ",")
			// Schema for Final Candidate Data
			// DUN_ID,NAME,GENDER,AGE
			if len(cols) != 4 {
				// DEBUG
				spew.Dump(line)
				spew.Dump(cols)
				panic(fmt.Errorf("WRONG! Must have 4 cols! Got %d", len(cols)))
			}
			code := fmt.Sprintf("%05s", cols[0])
			dataPattern := fmt.Sprintf("testdata/%s", code)
			officialName := cols[1]
			jantina := cols[2]

			// For matchCandidateName
			mc := matchCandidate{
				code:         code,
				name:         officialName,
				candidateDir: dataPattern,
			}
			matchCandidateName(&mc)
			// LAST STEP; get data from here .. if found; otherwise need manual
			var age string
			if mc.matchedFileName != "" {
				fmt.Println("FIND AGE in", mc.matchedFileName)
				age = extractCandidatesAge(&mc)
				fmt.Println("FINAL_AGE:", age)
				// DEBUG
				//if len(mc.candidateAge) > 0 {
				//	spew.Dump(mc)
				//}
			}
			finalData = append(finalData, fmt.Sprintf("%s,%s,%s,%s", code, officialName, jantina, age))
			fmt.Println("================================================================")
		}
		// DEBUG
		//for _, candidateData := range finalData {
		//	fmt.Println(candidateData)
		//}
		werr := os.WriteFile("testdata/pru15-"+dunName+"-age.csv", []byte(strings.Join(finalData, "\n")), 0666)
		if werr != nil {
			panic(werr)
		}
	}

}

func matchCandidateName(mc *matchCandidate) {
	// Look for all files in dataPath
	candidateFiles, err := script.ListFiles(mc.candidateDir).Slice()
	if err != nil {
		panic(err)
	}
	//spew.Dump(candidateFiles)
	candidateFilePath := ""
	safeName := strings.ReplaceAll(mc.name, " ", "-")
	fmt.Println("<<<<<<<<<<<<", mc.name, "in", mc.candidateDir, ">>>>>>>>>>")
	for _, dataFilePath := range candidateFiles {
		// DEBUG
		//fmt.Println("Look for:", safeName, "  in", dataFilePath)
		if strings.Contains(dataFilePath, safeName) {
			candidateFilePath = dataFilePath
			break
		}

		for _, namePart := range strings.Split(safeName, "-") {
			//spew.Dump(namePart)
			if strings.Contains(dataFilePath, namePart) {
				candidateFilePath = dataFilePath
				break
			}
		}
	}
	// check if candidates components are there
	// favor exact match; otherwise pick the most numbers ..
	if candidateFilePath == "" {
		// Should NOT get here .. means need to DEBUG!
		fmt.Println("DEBUG: Look for ", mc.name, "in:")
		for _, dataFilePath := range candidateFiles {
			fmt.Println(dataFilePath)
		}
		return
	}
	mc.matchedFileName = candidateFilePath
	fmt.Println("FOUND_CHOSEN_PATH:", candidateFilePath)

	return
}

func extractCandidatesAge(mc *matchCandidate) (age string) {
	age = "2022"
	// If find DOB; extract and leave first!
	// DOB pattern "ContentPlaceHolder1_lblDob"
	// Pattern DD/M/YYYY e.g. 18/1/1967
	reDOB := regexp.MustCompile("^.+\\d+/\\d+/(\\d+).+$")
	replaceTemplate := "$1"
	dobMatches, derr := script.File(mc.matchedFileName).Match("ContentPlaceHolder1_lblDob").ReplaceRegexp(reDOB, replaceTemplate).Slice()
	if derr != nil {
		panic(derr)
	}
	if len(dobMatches) > 0 {
		// Not needed quite useless
		//mc.candidateAge = append(mc.candidateAge, dobMatches...)
		// DEBUG
		//spew.Dump(dobMatches)
		year, cerr := strconv.Atoi(dobMatches[0])
		if cerr != nil {
			panic(cerr)
		}
		fmt.Println("YEAR_BIRTH:", year)
		age = strconv.Itoa(2022 - year)
		// DEBUG
		//fmt.Println("DEBUG_AGE:", 2022-year)
		return age
	}
	// If cannot find DOB; try a more generic search; add the findings?
	re := regexp.MustCompile("^.+(\\d{2})\\s+.+tahun.*$")
	matches, err := script.File(mc.matchedFileName).MatchRegexp(re).ReplaceRegexp(re, replaceTemplate).Slice()
	if err != nil {
		panic(err)
	}
	//if len(matches) > 0 {
	//	mc.candidateAge = append(mc.candidateAge, matches...)
	//	// Check should at least be 21
	//}
	if len(matches) > 0 {
		// Not needed quite useless
		//mc.candidateAge = append(mc.candidateAge, dobMatches...)
		// DEBUG
		//spew.Dump(matches)
		possibleAge, cerr := strconv.Atoi(matches[0])
		if cerr != nil {
			panic(cerr)
		}
		if possibleAge < 21 {
			fmt.Println("IMPOSSIBLE: AGE MUST >21", possibleAge)
		} else if possibleAge > 100 {
			fmt.Println("IMPOSSIBLE: AGE MUST <100", possibleAge)
		} else {
			age = strconv.Itoa(possibleAge)
		}

		return age
	}

	return age
}

func downloadCandidates() {

	//duns := []string{"perlis"}
	duns := []string{"perak", "pahang"}
	for _, dunName := range duns {
		r, err := script.File("testdata/pru15-" + dunName + ".csv").Slice()
		if err != nil {
			panic(err)
		}
		for _, line := range r {
			cols := strings.Split(line, ",")
			if len(cols) != 4 {
				// DEBUG
				spew.Dump(line)
				spew.Dump(cols)
				panic(fmt.Errorf("WRONG! Must have 4 cols! Got %d", len(cols)))
			}
			processDUNCandidatesProfile(candidate{
				name: cols[3],
				url:  cols[0],
				par:  fmt.Sprintf("%03s", cols[1]),
				dun:  fmt.Sprintf("%02s", cols[2]),
				code: fmt.Sprintf("%03s%02s", cols[1], cols[2]),
			})
		}
	}

}

func processDUNCandidatesProfile(c candidate) {
	baseURL := "https://pru.sinarharian.com.my"
	fmt.Println("Get Profile for", c.name)
	fmt.Println("Call URL:", baseURL+c.url)
	fmt.Println(fmt.Sprintf("CODE: %s", c.code))
	safeName := strings.ReplaceAll(c.name, " ", "-")
	dataPath := fmt.Sprintf("testdata/%s", c.code)
	merr := os.MkdirAll(dataPath, 0755)
	if merr != nil {
		panic(merr)
	}
	if script.IfExists(fmt.Sprintf("%s/%s.html", dataPath, safeName)).Error() != nil {
		// has error; means the file does not exist!
		fmt.Println("GOTTA DOWNLOAD!!!")
		n, err := script.Get(baseURL + c.url).WriteFile(fmt.Sprintf("%s/%s.html", dataPath, safeName))
		if err != nil {
			panic(err)
		}
		fmt.Println("N:", n)
	} else {
		fmt.Println("FOUND! at", fmt.Sprintf("%s/%s.html", dataPath, safeName))
	}
}

func getCandidates() {
	fmt.Println("RAW_PERLIS: ")
	fmt.Println("===============")
	processDUNGetCandidates("perlis")
	fmt.Println("RAW_PERAK: ")
	fmt.Println("===============")
	processDUNGetCandidates("perak")
	fmt.Println("RAW_PAHANG: ")
	fmt.Println("===============")
	processDUNGetCandidates("pahang")
}

func processDUNGetCandidates(dunName string) {
	rexp := regexp.MustCompile("^<a.*href=\"(.+)\".*P(\\d+)-N(\\d+) (.+)</a>$")
	replaceTemplate := "$1,$2,$3,$4"
	// DEBUG
	//n, err := script.File("testdata/pru15-"+dunName+".txt").ReplaceRegexp(rexp, replaceTemplate).Stdout()
	n, err := script.File("testdata/pru15-"+dunName+".txt").ReplaceRegexp(rexp, replaceTemplate).WriteFile("testdata/pru15-" + dunName + ".csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("N: ", n)
}
