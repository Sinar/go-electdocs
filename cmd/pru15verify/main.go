package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"sync"
)

var (
	lookupPAR map[string]par
	lookupDUN map[string]dun
)

func init() {
	// Setup global lookup .. the only allowed use of global :(
	lookupPAR = LookupPAR()
	lookupDUN = LookupDUN()
}

func main() {

	fmt.Println("Welcome to PRU14 Verifier!")
	// DEBUG
	ProcessKedah()
	//ProcessMelaka()
	//ProcessN9()
	//ProcessPahang()
}

func ProcessPahang() {
	//pars := []string{"P078", "P079"}
	pars := []string{
		"P078", "P079", "P080", "P081",
		"P082", "P083", "P084", "P085",
		"P086", "P087", "P088", "P089",
		"P090", "P091",
	}
	AssembleResultsPerPAR("PAHANG", pars)
}

func ProcessN9() {
	pars := []string{
		"P126", "P127", "P128", "P129",
		"P130", "P131", "P132", "P133",
	}
	AssembleResultsPerPAR("N9", pars)
}

func ProcessMelaka() {
	pars := []string{
		"P134", "P135", "P136", "P137", "P138", "P139",
	}
	AssembleResultsPerPAR("MELAKA", pars)
}

func ProcessKedah() {
	// DEBUG
	//LookupState()
	//LookupParty()
	//LookupPAR()
	// DEBUG
	//pars := []string{"P006"}
	pars := []string{
		//"P004", "P005",
		"P006", "P007", "P008", "P009",
		"P010", "P011", "P012", "P013",
		"P014", "P015", "P016", "P017",
		"P018",
	}
	//pars := []string{
	//	"P004", "P005", "P006", "P007", "P008", "P009",
	//	"P010", "P011", "P012", "P013", "P014", "P015",
	//	"P016", "P017", "P018",
	//}

	AssembleResultsPerPAR("KEDAH", pars)
	//SanityAssembleResultsPerPAR("KEDAH", pars)
}

// SanityAssembleResultsPerPAR is stripped down version to ensure col correct ..
func SanityAssembleResultsPerPAR(state string, pars []string) {
	// Extract out candidates; much more diverse in PRU15!!
	candidates := LookupResults(state)
	// DEBUG
	spew.Dump(candidates)
	for _, par := range pars {
		//SanityTestPARSaluran(par)
		currentPAR := par
		//fmt.Println("Processing PAR: ", currentPAR)
		salurans := LoadPARSaluran(currentPAR)
		// DEBUG
		//spew.Dump(salurans)
		fmt.Println("NO ROWS: ", len(salurans))
	}
}

func AssembleResultsPerPAR(state string, pars []string) {
	candidates := LookupResults(state)
	// DEBUG
	//spew.Dump(candidates)

	// For each PAR ID; extract out the first section
	// lookup from Results; which is ordered by SaluranID: 0,1,2
	// attach to the candidateVotes
	// For each pars; extract the number; remove the P
	// Look up the candidates; ordered by the BallotID
	var wg sync.WaitGroup
	prefixData := make([][]string, 0)
	candidateData := make([][]string, 0)
	suffixData := make([][]string, 0)
	for _, par := range pars {
		//SanityTestPARSaluran(par)
		currentPAR := par
		//fmt.Println("Processing PAR: ", currentPAR)
		salurans := LoadPARSaluran(currentPAR)
		wg.Add(3) // MUST match the number of concurrent funcs ..
		go func() {
			prefixData = append(prefixData,
				processPrefix(currentPAR, salurans)...,
			)
			wg.Done()
		}()
		go func() {
			candidateData = append(candidateData,
				processCandidates(currentPAR, salurans, candidates)...,
			)
			wg.Done()
		}()
		go func() {
			suffixData = append(suffixData,
				processSuffix(currentPAR, salurans, candidates)...,
			)
			wg.Done()
		}()
		// Block till done ..
		wg.Wait()
		//break
	}
	// Assemble the sections ..
	outputCSV(fmt.Sprintf("testdata/%s-prefix.csv", state), prefixData)
	outputCSV(fmt.Sprintf("testdata/%s-candidates.csv", state), candidateData)
	outputCSV(fmt.Sprintf("testdata/%s-suffix.csv", state), suffixData)

}
