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

// votesResult is Map of BallotID to voteCount
type votesResult map[string]string // KEY --> DM_ID/BALLOT_ID /coalition?

// candidate final value ..
type candidate struct {
	ID         string // PAR_ID/BALLOT_ID
	name       string
	gender     string
	totalVotes string
	party      string
	coalition  string
}

// resultsDUN model per candidate of each DUNID
type resultsDUN struct {
	ID          string // DUN_ID/BALLOT_ID
	votesResult        // DM_ID/BALLOT_ID -> votes
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
	candidates := make(map[string]candidate, 200)

	for _, l := range data {
		cols := strings.Split(l, ",")
		// DEBUG
		//spew.Dump(cols)
		if cols[0] == "id" || cols[0] == "ID" {
			numCols = len(cols)
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
		// COALITION_ID --> LOOKUP
		candidate := candidate{
			ID:         fmt.Sprintf("%s/%s", cols[5], cols[10]),
			name:       cols[1],
			gender:     cols[4],
			totalVotes: cols[11],
			party:      cols[3],
			coalition:  cols[3],
		}
		candidates[candidate.ID] = candidate
	}
	spew.Dump(candidates)

}
