package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to PRU14 Ager..")
	Run()
}

func Run() {

	// Experiment to run took too long ..
	// checkPAR ..
	//checkPAR()
	// Get all PAR so can be passed on to each PAR extraction ..
	//examplePAR()
	// Above not working; bypass by just getting raw data directly :(

	//ProcessKedah()
	//ProcessMelaka()
	//ProcessN9()

	// Using fuxxy method as direct html gone :(
	ProcessPahang()
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
	ExactCandidatePerPAR("PAHANG", pars)
	ExtractCandidateAgePerPAR("PAHANG")
}

func ProcessN9() {
	pars := []string{
		"P126", "P127", "P128", "P129",
		"P130", "P131", "P132", "P133",
	}
	//DownloadCandidatePerPAR("N9", pars)
	//spew.Dump(pars)
	//ExtractCandidatePerPAR("N9", pars)
	fmt.Println("N9_LEN:", len(pars))
	// Extracted out final ..
	//ExtractCandidateAgePerPAR("N9")
}

func ProcessMelaka() {
	pars := []string{
		"P134", "P135", "P136", "P137", "P138", "P139",
	}
	//spew.Dump(pars)
	//LookupResults("MELAKA")
	//LookupSaluran("MELAKA")
	//AssembleResultsPerPAR("MELAKA", pars)
	// Step #1
	//DownloadCandidatePerPAR("MELAKA", pars)
	// Step #2:
	//ExtractCandidatePerPAR("MELAKA", pars)
	fmt.Println("MELAKA_LEN:", len(pars))
	// Extracted out ,,
	//ExtractCandidateAgePerPAR("MELAKA")
}

func ProcessKedah() {
	// DEBUG
	//LookupState()
	//LookupParty()
	//LookupPAR()
	// DEBUG
	//pars := []string{"P004"}
	//pars := []string{"P004", "P005"}
	pars := []string{
		"P004", "P005", "P006", "P007", "P008", "P009",
		"P010", "P011", "P012", "P013", "P014", "P015",
		"P016", "P017", "P018",
	}
	//DownloadCandidatePerPAR("KEDAH", pars)
	//spew.Dump(pars)
	//ExtractCandidatePerPAR("KEDAH", pars)
	fmt.Println("KEDAH_LEN:", len(pars))
	// Extracted out and corrected
	//ExtractCandidateAgePerPAR("KEDAH")
}
