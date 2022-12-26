package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

// LookupResults gets the candidate totals on aggregate; for every PAR
func LookupResults(state string) map[string][]candidate {
	data, err := script.File(fmt.Sprintf("testdata/%s.csv", state)).Slice()
	if err != nil {
		panic(err)
	}
	numCols := 0
	var candidatesPAR map[string][]candidate
	candidatesPAR = make(map[string][]candidate, 0)
	//parties := lookupCoalition()
	for _, l := range data {
		cols := strings.Split(l, ",")
		// DEBUG
		//spew.Dump(cols)
		if cols[0] == "id" || cols[0] == "ID" {
			numCols = len(cols)
			continue
		}
		// Verify
		if len(cols) != numCols {
			panic("Incorrect cols!!" + l)
		}

		// Annotate against state, par, party
		// KEY --> PAR_ID/BALLOT_ID
		// PAR_ID - 5 split --> 01800
		// BALLOT_ID - 10
		// Name - 1
		// Gender - 4 --> convert to long text?
		// Age
		// Party - 3 --> LOOKUP
		// 	If Party == 20; append 7 to it
		// COALITION_ID --> LOOKUP
		parID := cols[5]
		candidate := candidate{
			name: cols[1],
		}
		candidatesPAR[parID] = append(candidatesPAR[parID], candidate)
	}
	// DEBUG
	//spew.Dump(candidatesPAR)
	return candidatesPAR
}

// LookupSaluran gets the candidate totals on aggregate; for every PAR
func LookupSaluran(state string) {
	data, err := script.File(fmt.Sprintf("testdata/Saluran-%s.csv", state)).Slice()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		spew.Dump(row)
		break
	}

	return
}
