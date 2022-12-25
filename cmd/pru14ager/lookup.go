package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
)

// LookupResults gets the candidate totals on aggregate; for every PAR
func LookupResults(state string) {
	data, err := script.File(fmt.Sprintf("testdata/%s.csv", state)).Slice()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		spew.Dump(row)
		break
	}

	return
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
