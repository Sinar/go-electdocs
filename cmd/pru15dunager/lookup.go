package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Candidate struct {
	KodDUN          string `json:"Kod DUN"`
	NamaDUN         string `json:"Nama DUN"`
	BilanganPemilih int    `json:"Bilangan Pemilih"`
	NoCalon         int    `json:"No Calon"`
	Nama            string `json:"Nama"`
	Jantina         string `json:"Jantina"`
	Umur            int    `json:"Umur"`
	Parti           string `json:"Parti"`
}

// LoadRawData ..

// Lookup Functions

func MapPARDUNs(state string) map[string]string {
	pd := make(map[string]string, 0)

	// Iterate through ..
	// Action if NOT exist; otherwise skip?
	return pd
}

// LookupCandidatesByDUN extracts Candidate based on PAR, DUN NAME
func LookupCandidatesByDUN(state, par, dun string) {
	// Example: https://pru.sinarharian.com.my/prn15/negeri-sembilan/tampin/repah

}

// LoadTSVCandidates load from the raw data from Bernama
func LoadTSVCandidates(state string) {
	fmt.Println("Parsing data for State:", state)
	filePath := fmt.Sprintf("testdata/%s.csv", state) // Update the file path with the actual file name

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	candidates := make([]Candidate, 0)

	for i, row := range records {
		if i == 0 {
			// Skip the header row
			continue
		}

		candidate := Candidate{
			KodDUN:          strings.TrimSpace(row[0]),
			NamaDUN:         strings.TrimSpace(row[1]),
			BilanganPemilih: parseInt(strings.ReplaceAll(row[2], ",", "")),
			NoCalon:         parseInt(row[3]),
			Nama:            strings.TrimSpace(strings.ToUpper(row[4])),
			Jantina:         strings.TrimSpace(row[5]),
			Umur:            parseInt(row[6]),
			Parti:           strings.TrimSpace(row[7]),
		}

		candidates = append(candidates, candidate)
	}

	jsonData, err := json.MarshalIndent(candidates, "", "    ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func parseInt(s string) int {

	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	return num
}
