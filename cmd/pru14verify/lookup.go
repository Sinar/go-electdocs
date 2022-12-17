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

type par struct {
	name string
	code string
}

type dun struct {
	name string
	code string
}

type state struct {
	name string
}

type party struct {
	name      string
	coalition string
}

// resultsDUN model per candidate of each DUNID
type resultsDUN struct {
	ID          string // DUN_ID/BALLOT_ID
	votesResult        // DM_ID/BALLOT_ID -> votes
}

// Lookup data; return map tied to Primary keys

// LookupState for Parliament Data
func LookupState() map[string]string {
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
		// StateID --> Official Name
	}

	return nil
}

// lookupCoalitionKedah is the unique map for Kedah
func lookupCoalitionKedah() map[string]party {
	// KEDAH LOOKUP Table ..
	//2	PAS	GS
	//37	PKR	PH
	//1	BN	BN
	//20	25	KEY
	//9	PRM
	parties := make(map[string]party, 0)
	parties["1"] = party{
		name:      "BN",
		coalition: "BN",
	}
	parties["2"] = party{
		name:      "PAS",
		coalition: "GS",
	}
	parties["9"] = party{
		name:      "PRM",
		coalition: "",
	}
	parties["37"] = party{
		name:      "PKR",
		coalition: "PH",
	}
	parties["2025"] = party{
		name:      "IND",
		coalition: "KEY",
	}

	return parties
}

// LookupParty for Parliament Data
func LookupParty() map[string]party {
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
		// PartyID --> Official Name
	}
	return nil
}

// LookupPAR for Parliament Data
func LookupPAR() map[string]par {
	data, err := script.File("testdata/par.csv").Slice()
	if err != nil {
		panic(err)
	}
	numCols := 0
	pars := make(map[string]par, 0)
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
		// PAR_ID --> Official Name
		// ID - 0
		// CODE - 1
		// NAME - 4
		pars[cols[0]] = par{
			name: cols[4],
			code: cols[1],
		}
	}
	return pars
}

// LookupDUN for DUN Data
func LookupDUN() map[string]dun {
	data, err := script.File("testdata/dun.csv").Slice()
	if err != nil {
		panic(err)
	}
	numCols := 0
	duns := make(map[string]dun, 0)
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
		// DBEUG
		//spew.Dump(cols)
		// Annotate against state
		// DUN_ID --> Official Name
		// ID - 0
		// CODE - 1
		// NAME - 4
		duns[cols[0]] = dun{
			name: cols[4],
			code: cols[1],
		}
	}
	return duns
}

// LookupResults gets the candidate totals on aggregate; for every PAR
func LookupResults(state string) map[string]candidate {
	data, err := script.File(fmt.Sprintf("testdata/%s.csv", state)).Slice()
	if err != nil {
		panic(err)
	}
	numCols := 0
	candidates := make(map[string]candidate, 200)
	parties := lookupCoalitionKedah()
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
		// 	If Party == 20; append 7 to it
		// COALITION_ID --> LOOKUP
		if cols[3] == "20" {
			cols[3] = fmt.Sprintf("20%s", cols[7])
			fmt.Println("IND_ID: ", cols[3])
		}
		candidate := candidate{
			ID:         fmt.Sprintf("%s/%s", cols[5], cols[10]),
			name:       cols[1],
			gender:     cols[4],
			totalVotes: cols[11],
			party:      parties[cols[3]].name,
			coalition:  parties[cols[3]].coalition,
		}
		candidates[candidate.ID] = candidate
	}
	// DEBUG
	//spew.Dump(candidates)
	return candidates
}
