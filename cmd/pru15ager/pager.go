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

func Run() {
	fmt.Println("Runing ..")
	//getCandidates()
	downloadCandidates()
}

func downloadCandidates() {

	//duns := []string{"perlis"}
	duns := []string{"perak"}
	for _, dunName := range duns {
		r, err := script.File("testdata/pru15-" + dunName + ".csv").Slice()
		if err != nil {
			panic(err)
		}
		for _, line := range r {
			spew.Dump(line)
			cols := strings.Split(line, ",")
			if len(cols) != 4 {
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
		fmt.Println("FOUND! at", fmt.Sprintf("%s/%s.html", dataPath, safeName))
	} else {
		fmt.Println("GOTTA DOWNLOAD!!!")
		//n, err := script.Get(baseURL + c.url).WriteFile(fmt.Sprintf("%s/%s.html", dataPath, safeName))
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println("N:", n)
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
