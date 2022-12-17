package main

import (
	"fmt"
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

	AssembleResultsPerPAR("KEDAH", pars)
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
