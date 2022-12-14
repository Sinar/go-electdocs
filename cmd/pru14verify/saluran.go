package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// 0 - Bil,
// 1 - No. Kod Daerah Mengundi,
// 2 - Nama Pusat Mengundi,
// 3 - Nombor Tempat Mengundi (saluran),
// 4 - Jumlah kertas undi yang patut berada di dalam peti undi (A),
// 5 - Bilangan undian oleh pemilih bagi setiap orang calon yang bertanding (B) BALLOT_ID 1,
// 6 - Bilangan undian oleh pemilih bagi setiap orang calon yang bertanding (B) BALLOT_ID 2,
// 7 - Bilangan undian oleh pemilih bagi setiap orang calon yang bertanding (B) BALLOT_ID 3,
// 8 - Bilangan undian oleh pemilih bagi setiap orang calon yang bertanding (B) Jumlah undian oleh pemilih,
// 9 - Bilangan kertas undi yang ditolak (C),
// 10 - Jumlah kertas undi yang dikeluarkan kepada pengundi tetapi tidak dimasukkan ke dalam peti undi (D),
// 11 - Nama Daerah Mengundi,
// 12 - Jenis Undi

// saluran is breakdown of each unique saluran
type saluran struct {
	ID                 string // 1 --> Daerah Mengundi ID
	dunID              string
	votingLocationName string // 2
	saluranID          string // 3
	totalVotesIssued   string // 4
	totalVotesCast     string // 8
	totalVotesSpoilt   string // 9
	totalVotesNotCast  string // 10
	daerahMengundi     string // 11
	voteType           string // 12
	candidateVotes     []string
}

// For 3 candidates - it is 5 + 3 + 5 = 13
// Fixed 5 from the front
// Dynamic candidates - x
// Fixed 5 from the end

func NewSaluranRow(prefixSlice, suffixSlice, candidatesSlice []string) saluran {
	// DEBUG
	//fmt.Println("PREFIX: ", prefixSlice)
	//fmt.Println("SUFFIX: ", suffixSlice)
	//fmt.Println("CANDIDATES: ", candidatesSlice)
	//fmt.Println("==============================")

	// Can do validation that prefix + suffix are both 5 in len ..
	if len(prefixSlice) != 5 || len(suffixSlice) != 5 {
		panic("Incorrect Prefix or Suffix length")
	}

	// Process dunID which will be used later for lookup key in SPR data
	dmID := prefixSlice[1]
	dunID := dmID
	if strings.Contains(dmID, "/") {
		dunID = strings.ReplaceAll(dmID, "/", "")[0:5]
	}

	// DEBUG
	//for i, votes := range candidatesSlice {
	//	fmt.Println(dmID, "CANDIDATE:", dunID, "/", i, "got", votes, "in SALURAN", prefixSlice[3])
	//}

	return saluran{
		ID:                 dmID,
		dunID:              dunID,
		votingLocationName: prefixSlice[2],
		saluranID:          prefixSlice[3],
		totalVotesIssued:   prefixSlice[4],
		totalVotesCast:     suffixSlice[0],
		totalVotesSpoilt:   suffixSlice[1],
		totalVotesNotCast:  suffixSlice[2],
		daerahMengundi:     suffixSlice[3],
		voteType:           suffixSlice[4],
		candidateVotes:     candidatesSlice,
	}
}

func SanityTestPARSaluran(par string) {
	// Open the file
	recordFile, err := os.Open(fmt.Sprintf("testdata/%s.csv", par))
	if err != nil {
		fmt.Println("An error encountered ::", err)
		panic(err)
	}

	r := csv.NewReader(recordFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// DEBUG
	//	fmt.Print(records)
	var numCols, numCandidates int
	//var salurans []saluran
	// TODO: Have a map KEY to DMID so can do sanity check
	for _, cols := range records {
		// DEBUZG
		//spew.Dump(cols[0])
		if cols[0] == "Bil" {
			// DEBUG
			//fmt.Println("IN!!!")
			numCols = len(cols)
			numCandidates = len(cols) - 10
			fmt.Println(fmt.Sprintf("PAR: %s NO Candidates: %d", par, numCandidates))

			// Just sanity test so can move on .
			continue
		} else {
			// DEBUG
			//fmt.Println("NOT", cols[0])
			// Verify
			if len(cols) != numCols {
				panic("Incorrect cols!!")
			}
			//x := numCols - 5
			//y := numCols
			// NOw can structure + append
			//currentSaluran := NewSaluranRow(cols[0:5], cols[x:y], cols[5:x])
			//salurans = append(salurans, currentSaluran)
		}
		// DEBUG
		//		spew.Dump(cols)
	}
	// DEBUG
	//	spew.Dump(salurans)

}

func LoadPARSaluran(par string) []saluran {
	// Open the file
	recordFile, err := os.Open(fmt.Sprintf("testdata/%s.csv", par))
	if err != nil {
		fmt.Println("An error encountered ::", err)
		panic(err)
	}

	r := csv.NewReader(recordFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// DEBUG
	//	fmt.Print(records)
	var numCols, numCandidates int
	var salurans []saluran
	// TODO: Have a map KEY to DMID so can do sanity check
	for _, cols := range records {
		// DEBUZG
		//spew.Dump(cols[0])
		if cols[0] == "Bil" {
			// DEBUG
			//fmt.Println("IN!!!")
			numCols = len(cols)
			numCandidates = len(cols) - 10
			fmt.Println(fmt.Sprintf("NO Candidates: %d", numCandidates))
		} else {
			// DEBUG
			//fmt.Println("NOT", cols[0])
			// Verify
			if len(cols) != numCols {
				panic("Incorrect cols!!")
			}
			x := numCols - 5
			y := numCols
			// NOw can structure + append
			currentSaluran := NewSaluranRow(cols[0:5], cols[x:y], cols[5:x])
			salurans = append(salurans, currentSaluran)
		}
		// DEBUG
		//		spew.Dump(cols)
	}
	// DEBUG
	//	spew.Dump(salurans)
	// Do the mapping now and then output it immediately ..
	// DEBUG
	//parID := par[1:]
	//fmt.Println("RESULTS: ", parID)
	//for _, saluran := range salurans {
	//	// Passed in the slice of votes as a Lookup method
	//	// or output method for that row .. as csv?
	//	// Each COALTION will fit into a certain order in the slice;
	//	//	known a-priori per STATE
	//	// If COALITION not there; leave as blank
	//	keyID := fmt.Sprintf("%s/%s", parID, saluran.saluranID)
	//	fmt.Println("LOOkUP:", keyID)
	//}

	// DEBUG
	//s := saluran{
	//	candidateVotes: map[string]string{
	//		"1": "456",
	//		"2": "123",
	//	},
	//}
	//if _, ok := s.candidateVotes["100"]; ok {
	//	fmt.Println("Found KEY 100")
	//} else {
	//	fmt.Println("Did NOT find KEY 100")
	//}
	return salurans
}
