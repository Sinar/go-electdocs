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
	pars := []string{"P004"}
	//pars := []string{"P004", "P005"}
	for _, par := range pars {
		LoadPARSaluran(par)
	}
}
