package main

import (
	"fmt"
	"github.com/bitfield/script"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

// results data
// ID, - 0
// BALLOT_NAME, - 1
// JOB, - 2
// PID, - 3
// JANTINA, - 4
// DUN_ID, - 5
// TYPE (PAR/DUN), - 6
// IND_SYMBOL_ID, - 7
// STATUS (WIN/LOSE/DEPOSIT), - 8
// LAST_UPDATED, - 9
// BALLOT_ID, - 10
// TOTAL_VOTES, - 11
// MAJORITY, - 12
// WINNER_BALLOT_ID (?) - 13

// candidatesMap - DUN_ID/BALLOT_ID?

type candidate struct {
	ID         string // DUN_ID/BALLOT_ID
	name       string
	gender     string
	totalVotes string
	party      string
	coalition  string
}

// resultsDUN model per candidate of each DUNID
type resultsDUN struct {
	ID          string // DUN_ID
	votesResult        // BALLOT_ID -> CANDIDATE_ID
}

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
