package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func main() {

	fmt.Println("Welcome to PRU14 Verifier!")
	// DEBUG
	LookupState()
	LookupParty()
	LookupPAR()
	LookupResults("KEDAH")
	// DEBUG
	//EnrichPARSaluran()

}

func EnrichPARSaluran() {
	pars := []string{"P004", "P005"}
	for _, par := range pars {
		data, err := script.File(fmt.Sprintf("testdata/%s.csv", par)).Slice()
		if err != nil {
			panic(err)
		}
		numCols := 0
		for _, l := range data {
			cols := strings.Split(l, ",")
			if cols[0] == "id" || cols[0] == "ID" {
				numCols = len(cols)
			} else {
				// Verify
				if len(cols) != numCols {
					panic("Incorrect cols!!" + l)
				}
			}
			spew.Dump(cols)
			// Annotate against state, par, party
		}
	}
}
