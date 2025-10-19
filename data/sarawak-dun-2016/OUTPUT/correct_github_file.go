package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func normalizeName(name string) string {
	name = strings.TrimSpace(strings.ToUpper(name))
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(name, " ")
}

func fuzzyMatch(name1, name2 string) float64 {
	n1 := normalizeName(name1)
	n2 := normalizeName(name2)
	if n1 == n2 {
		return 1.0
	}
	longer := n1
	shorter := n2
	if len(n2) > len(n1) {
		longer = n2
		shorter = n1
	}
	if len(longer) == 0 {
		return 1.0
	}
	matches := 0
	for i := 0; i < len(shorter); i++ {
		if strings.Contains(longer, string(shorter[i])) {
			matches++
		}
	}
	return float64(matches) / float64(len(longer))
}

func isSubsetMatch(name1, name2 string) bool {
	n1 := normalizeName(name1)
	n2 := normalizeName(name2)
	return strings.Contains(n1, n2) || strings.Contains(n2, n1)
}

func main() {
	// Read our source file to get correct candidate names by DUN
	fmt.Println("Reading our source file (correct names)...")
	sourceFile, err := os.Open("/Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/Final-Sarawak-DUN-2016.csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer sourceFile.Close()

	sourceReader := csv.NewReader(sourceFile)
	sourceRecords, err := sourceReader.ReadAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Build map of DUN -> Party -> Correct Name
	correctNames := make(map[string]map[string]string)

	sourceHeader := sourceRecords[0]
	sourceColIdx := make(map[string]int)
	for i, col := range sourceHeader {
		sourceColIdx[col] = i
	}

	dunIdx := sourceColIdx["STATE CONSTITUENCY CODE"]

	partiesMap := map[string]string{
		"BN CANDIDATE":            "BN",
		"PH (1) CANDIDATE":        "PH (1)",
		"PH (2) CANDIDATE":        "PH (2)",
		"PAS CANDIDATE":           "PAS",
		"STAR CANDIDATE":          "STAR",
		"PBDSB CANDIDATE":         "PBDSB",
		"INDEPENDENT 1 CANDIDATE": "INDEPENDENT 1",
		"INDEPENDENT 2 CANDIDATE": "INDEPENDENT 2",
	}

	seenDUNCandidates := make(map[string]map[string]bool)

	for i := 1; i < len(sourceRecords); i++ {
		dunCode := sourceRecords[i][dunIdx]
		if dunCode == "" {
			continue
		}

		if correctNames[dunCode] == nil {
			correctNames[dunCode] = make(map[string]string)
			seenDUNCandidates[dunCode] = make(map[string]bool)
		}

		for candCol, party := range partiesMap {
			if colIdx, exists := sourceColIdx[candCol]; exists {
				candidateName := strings.TrimSpace(sourceRecords[i][colIdx])
				if candidateName != "" {
					key := party + "|" + candidateName
					if !seenDUNCandidates[dunCode][key] {
						correctNames[dunCode][party] = candidateName
						seenDUNCandidates[dunCode][key] = true
					}
				}
			}
		}
	}

	fmt.Printf("Loaded correct names for %d DUNs\n", len(correctNames))

	// Read GitHub file
	fmt.Println("Reading GitHub file...")
	githubFile, err := os.Open("/tmp/github_sarawak_2016.csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer githubFile.Close()

	githubReader := csv.NewReader(githubFile)
	githubRecords, err := githubReader.ReadAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	githubHeader := githubRecords[0]
	githubColIdx := make(map[string]int)
	for i, col := range githubHeader {
		githubColIdx[col] = i
	}

	githubDunIdx := githubColIdx["STATE CONSTITUENCY CODE"]

	githubPartyCols := map[string]string{
		"BN CANDIDATE":            "BN",
		"PH (1) CANDIDATE":        "PH (1)",
		"PH (2) CANDIDATE":        "PH (2)",
		"PAS CANDIDATE":           "PAS",
		"STAR CANDIDATE":          "STAR",
		"PBDSB CANDIDATE":         "PBDSB",
		"INDEPENDENT 1 CANDIDATE": "INDEPENDENT 1",
		"INDEPENDENT 2 CANDIDATE": "INDEPENDENT 2",
	}

	updateCount := 0

	// Update GitHub file candidate names
	fmt.Println("Updating candidate names in GitHub file...")
	for i := 1; i < len(githubRecords); i++ {
		dunCode := githubRecords[i][githubDunIdx]
		if dunCode == "" {
			continue
		}

		if correctNamesForDUN, exists := correctNames[dunCode]; exists {
			for candCol, party := range githubPartyCols {
				if colIdx, hasCol := githubColIdx[candCol]; hasCol {
					githubName := strings.TrimSpace(githubRecords[i][colIdx])
					if githubName != "" {
						if correctName, hasCorrectName := correctNamesForDUN[party]; hasCorrectName {
							// Check if names are different (fuzzy match)
							if normalizeName(githubName) != normalizeName(correctName) {
								// Verify it's the same person with fuzzy matching
								score := fuzzyMatch(githubName, correctName)
								if score > 0.5 || isSubsetMatch(githubName, correctName) {
									oldName := githubRecords[i][colIdx]
									githubRecords[i][colIdx] = correctName
									updateCount++
									fmt.Printf("DUN %s, %s: %s → %s\n", dunCode, party, oldName, correctName)
								}
							}
						}
					}
				}
			}
		}
	}

	// Write corrected GitHub file
	fmt.Println("\nWriting corrected GitHub file...")
	outputFile, err := os.Create("/Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/SARAWAK_2016_DUN_RESULTS_CORRECTED.csv")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range githubRecords {
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing: %v\n", err)
			return
		}
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("✓ Updated %d candidate names in GitHub file\n", updateCount)
	fmt.Println("✓ Output: SARAWAK_2016_DUN_RESULTS_CORRECTED.csv")
	fmt.Println(strings.Repeat("=", 80))
}
