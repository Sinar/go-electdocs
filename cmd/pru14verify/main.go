package main

import (
	"fmt"
	"sync"
)

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

	AssembleResultsPerPAR(pars)
}

func AssembleResultsPerPAR(pars []string) {
	candidates := LookupResults("KEDAH")
	// DEBUG
	//spew.Dump(candidates)

	// For each PAR ID; extract out the first section
	// lookup from Results; which is ordered by SaluranID: 0,1,2
	// attach to the candidateVotes
	// For each pars; extract the number; remove the P
	// Look up the candidates; ordered by the BallotID
	var wg sync.WaitGroup
	for _, par := range pars {
		//SanityTestPARSaluran(par)
		currentPAR := par
		//fmt.Println("Processing PAR: ", currentPAR)
		salurans := LoadPARSaluran(currentPAR)
		wg.Add(2) // MUST match the number of concurrent funcs ..
		go func() {
			//processPart1(currentPAR, salurans)
			wg.Done()
		}()
		go func() {
			processCandidates(currentPAR, salurans, candidates)
			wg.Done()
		}()
		// Block till done ..
		wg.Wait()
		break
	}
}
