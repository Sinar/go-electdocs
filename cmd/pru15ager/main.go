package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Welcome to ageism ..")
	// RunClassic()
	Run()

}

func Run() {
	ProcessKedah()
	// ProcessPahang() // Copied from pru14-ager; maybe future
}

func ProcessKedah() {

	//pars := []string{"P004", "P005", "P006"}
	pars := []string{
		"P004", "P005",
		"P006", "P007", "P008", "P009",
		"P010", "P011", "P012", "P013",
		"P014", "P015", "P016", "P017",
		"P018",
	}
	//pars := []string{"P018"}
	fmt.Println("KEDAH_LEN:", len(pars))
	// For each PAR; extract the calon page + download it ..
	// Fuzzy search + levishtein match name
	// Look up the candidate Party if found ..
	// Below one time only; no lock for now .. careful of Google rate limit + block ..
	FuzzyDownloadCandidatePerPAR("KEDAH", pars)
	// Below is fast so can run multiple times?
	// Below should not be run after manual correction ..
	//ExactCandidatePerPAR("KEDAH", pars)
	// Below after manual check should be frozen
	//ExtractCandidateAgePerPAR("KEDAH")
	//spew.Dump(pars)

}

func ProcessPahang() {
	//pars := []string{"P078", "P079"}
	pars := []string{
		"P078", "P079", "P080", "P081",
		"P082", "P083", "P084", "P085",
		"P086", "P087", "P088",
		"P089", "P090", "P091",
	}
	//pars := []string{"P089", "P090", "P091"}
	fmt.Println("PAHANG_LEN:", len(pars))
	// Below one time only; no lock for now .. careful of Google rate limit + block ..
	//FuzzyDownloadCandidatePerPAR("PAHANG", pars)
	// Below is fast so can run multiple times?
	// Below should not be run after manual correction ..
	//ExactCandidatePerPAR("PAHANG", pars)
	// Below after manual check should be frozen
	//ExtractCandidateAgePerPAR("PAHANG")
	spew.Dump(pars)
}
