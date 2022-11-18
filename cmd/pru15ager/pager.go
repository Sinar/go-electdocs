package main

import (
	"fmt"
	"github.com/bitfield/script"
	"regexp"
)

func Run() {
	fmt.Println("Runing ..")
	//getCandidates()
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
