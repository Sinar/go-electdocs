package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

// Lookup data; return map tied to Primary keys

// LookupState for Parliament Data
func LookupState() {
	data, err := script.File("testdata/states.csv").Slice()
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
	}
}

// LookupParty for Parliament Data
func LookupParty() {
	data, err := script.File("testdata/party.csv").Slice()
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
		// Enrichment with coalition ..
	}
}

// LookupPAR for Parliament Data
func LookupPAR() {
	data, err := script.File("testdata/par.csv").Slice()
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
		// Annotate against state
	}
}

func LookupResults(state string) {
	data, err := script.File(fmt.Sprintf("testdata/%s.csv", state)).Slice()
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
