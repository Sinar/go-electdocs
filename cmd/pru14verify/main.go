package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to PRU14 Verifier!")
	// DEBUG
	//LookupState()
	//LookupParty()
	//LookupPAR()
	//LookupResults("KEDAH")
	// DEBUG
	//pars := []string{"P004"}
	//pars := []string{"P004", "P005"}
	pars := []string{
		"P004", "P005", "P006", "P007", "P008", "P009",
		"P010", "P011", "P012", "P013", "P014", "P015",
		"P016", "P017", "P018",
	}

	for _, par := range pars {
		SanityTestPARSaluran(par)
		//LoadPARSaluran(par)
	}
}
