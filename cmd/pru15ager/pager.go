package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"os"
	"regexp"
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
	candidateNames  []string
	matchedFileName string
}

func Run() {
	fmt.Println("Runing ..")
	//getCandidates()
	//downloadCandidates()
	extractCandidates()
}

func extractCandidates() {
	duns := []string{"perlis"}
	//duns := []string{"perak", "pahang"}
	for _, dunName := range duns {
		// Open data from Spreadsheet .. use this instead? so can be written back correctly
		// and can cut + paste?
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
			officialName := cols[1]
			dataPattern := fmt.Sprintf("testdata/%s", code)

			// For matchCandidateName
			mc := matchCandidate{
				code:         code,
				name:         officialName,
				candidateDir: dataPattern,
			}
			matchCandidateName(&mc)
			// LAST STEP; get data from here .. if found; otherwise need manual
			if mc.matchedFileName != "" {
				fmt.Println("FIND AGE in", mc.matchedFileName)
			}
			fmt.Println("================================================================")
			// DEBUG
			//spew.Dump(mc)
			//// For extractCandidateAge
			//c := candidate{
			//	name: officialName,
			//	code: code,
			//}
			//extractCandidatesAge(&c)
			//// After manipulations ..
			//spew.Dump(c)
		}
		// fill it in the file? where it is ordered? match it?
		// load out; and try to match loosely ..
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

func extractCandidatesAge(c *candidate) {
	//ageofCandidate := 2022
	fmt.Println("Get Profile for", c.name)
	fmt.Println(fmt.Sprintf("CODE: %s", c.code))
	safeName := strings.ReplaceAll(c.name, " ", "-")
	dataPath := fmt.Sprintf("testdata/%s", c.code)
	fmt.Println("Open:", safeName, ".html in", dataPath)

	// Dump it out as CSV ..
	//strconv.ParseInt(ageofCandidate, 10, 8)
	//age := string(ageofCandidate)
	c.age = "2022"
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
