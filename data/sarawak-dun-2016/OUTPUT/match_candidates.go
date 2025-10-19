package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Candidate struct {
	Party         string
	Candidate     string
	Sex           string
	Age           string
	Vote          string
	OrigCandidate string
	OrigSex       string
	OrigAge       string
	OrigVote      string
}

type DUNData map[string][]Candidate

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

	// Simple similarity calculation
	longer := n1
	shorter := n2
	if len(n2) > len(n1) {
		longer = n2
		shorter = n1
	}

	if len(longer) == 0 {
		return 1.0
	}

	// Count matching characters
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

func readSourceOfTruth(filePath string) (DUNData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty CSV file")
	}

	header := records[0]
	dunData := make(DUNData)
	seenCandidates := make(map[string]map[string]bool) // dun_code -> candidate_name -> true

	partiesMap := map[string]string{
		"BN":            "BN",
		"PH (1)":        "PH",
		"PH (2)":        "PH2",
		"PAS":           "PAS",
		"STAR":          "STAR",
		"PBDSB":         "PBDSB",
		"INDEPENDENT 1": "INDEPENDENT 1",
		"INDEPENDENT 2": "INDEPENDENT 2",
	}

	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, val := range record {
			if i < len(header) {
				row[header[i]] = val
			}
		}

		dunCode := row["STATE CONSTITUENCY CODE"]
		if dunCode == "" {
			continue
		}

		if seenCandidates[dunCode] == nil {
			seenCandidates[dunCode] = make(map[string]bool)
		}

		for sourceParty, normalizedParty := range partiesMap {
			candidateKey := sourceParty + " CANDIDATE"
			candidate := strings.TrimSpace(row[candidateKey])

			if candidate != "" && !seenCandidates[dunCode][candidate] {
				seenCandidates[dunCode][candidate] = true

				cand := Candidate{
					Party:     normalizedParty,
					Candidate: candidate,
					Sex:       strings.TrimSpace(row[sourceParty+" CANDIDATE SEX"]),
					Age:       strings.TrimSpace(row[sourceParty+" CANDIDATE AGE"]),
					Vote:      strings.TrimSpace(row[sourceParty+" VOTE"]),
				}

				dunData[dunCode] = append(dunData[dunCode], cand)
			}
		}
	}

	return dunData, nil
}

func readGitHubData(filePath string) (DUNData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("empty CSV file")
	}

	header := records[0]

	// Aggregate data by DUN and candidate
	type CandKey struct {
		DUN       string
		Party     string
		Candidate string
	}

	aggregated := make(map[CandKey]*Candidate)

	partiesMap := map[string]string{
		"BN":            "BN",
		"PH (1)":        "PH",
		"PH (2)":        "PH2",
		"PAS":           "PAS",
		"STAR":          "STAR",
		"PBDSB":         "PBDSB",
		"INDEPENDENT 1": "INDEPENDENT 1",
		"INDEPENDENT 2": "INDEPENDENT 2",
	}

	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, val := range record {
			if i < len(header) {
				row[header[i]] = val
			}
		}

		dunCode := row["STATE CONSTITUENCY CODE"]
		if dunCode == "" {
			continue
		}

		for githubParty, normalizedParty := range partiesMap {
			candidateKey := githubParty + " CANDIDATE"
			voteKey := githubParty + " VOTE"
			candidate := strings.TrimSpace(row[candidateKey])

			if candidate != "" {
				vote := 0
				if voteStr := row[voteKey]; voteStr != "" {
					vote, _ = strconv.Atoi(voteStr)
				}

				sex := strings.TrimSpace(row[githubParty+" CANDIDATE SEX"])
				age := strings.TrimSpace(row[githubParty+" CANDIDATE AGE"])

				key := CandKey{
					DUN:       dunCode,
					Party:     normalizedParty,
					Candidate: candidate,
				}

				if existing, ok := aggregated[key]; ok {
					// Aggregate votes
					existingVote, _ := strconv.Atoi(existing.Vote)
					existing.Vote = strconv.Itoa(existingVote + vote)
					if existing.Sex == "" {
						existing.Sex = sex
					}
					if existing.Age == "" {
						existing.Age = age
					}
				} else {
					aggregated[key] = &Candidate{
						Party:     normalizedParty,
						Candidate: candidate,
						Sex:       sex,
						Age:       age,
						Vote:      strconv.Itoa(vote),
					}
				}
			}
		}
	}

	// Convert to DUNData
	dunData := make(DUNData)
	for key, cand := range aggregated {
		dunData[key.DUN] = append(dunData[key.DUN], *cand)
	}

	return dunData, nil
}

func matchCandidates(sourceData, githubData DUNData) DUNData {
	results := make(DUNData)
	var mu sync.Mutex
	var wg sync.WaitGroup

	dunCodes := make([]string, 0, len(sourceData))
	for dunCode := range sourceData {
		dunCodes = append(dunCodes, dunCode)
	}

	for _, dunCode := range dunCodes {
		wg.Add(1)
		go func(dun string) {
			defer wg.Done()

			sourceCandidates := sourceData[dun]
			githubCandidates := githubData[dun]

			dunResults := make([]Candidate, 0)

			for _, sourceCand := range sourceCandidates {
				matched := false
				var bestMatch *Candidate
				bestScore := 0.0

				// Try to match with GitHub data
				for i := range githubCandidates {
					githubCand := &githubCandidates[i]

					// Must be same party
					if sourceCand.Party != githubCand.Party {
						continue
					}

					// Exact match (ignoring spaces)
					if normalizeName(sourceCand.Candidate) == normalizeName(githubCand.Candidate) {
						resultCand := sourceCand
						resultCand.OrigCandidate = ""
						resultCand.OrigSex = ""
						resultCand.OrigAge = ""
						resultCand.OrigVote = ""

						// Compare votes
						sourceVote := strings.ReplaceAll(sourceCand.Vote, ",", "")
						githubVote := githubCand.Vote
						if sourceVote != githubVote && githubVote != "" && githubVote != "0" {
							resultCand.OrigVote = githubVote
						}

						dunResults = append(dunResults, resultCand)
						matched = true
						break
					}

					// Fuzzy match or subset match
					score := fuzzyMatch(sourceCand.Candidate, githubCand.Candidate)
					if score > bestScore || isSubsetMatch(sourceCand.Candidate, githubCand.Candidate) {
						if isSubsetMatch(sourceCand.Candidate, githubCand.Candidate) {
							score = 0.8
						}
						if score > bestScore {
							bestScore = score
							bestMatch = githubCand
						}
					}
				}

				// If fuzzy match found
				if !matched && bestMatch != nil && bestScore > 0.5 {
					resultCand := sourceCand

					if normalizeName(sourceCand.Candidate) != normalizeName(bestMatch.Candidate) {
						resultCand.OrigCandidate = bestMatch.Candidate
					}
					if sourceCand.Sex != bestMatch.Sex && bestMatch.Sex != "" {
						resultCand.OrigSex = bestMatch.Sex
					}
					if sourceCand.Age != bestMatch.Age && bestMatch.Age != "" {
						resultCand.OrigAge = bestMatch.Age
					}
					sourceVote := strings.ReplaceAll(sourceCand.Vote, ",", "")
					if sourceVote != bestMatch.Vote && bestMatch.Vote != "" && bestMatch.Vote != "0" {
						resultCand.OrigVote = bestMatch.Vote
					}

					dunResults = append(dunResults, resultCand)
					matched = true
				}

				// If still no match, use source only
				if !matched {
					dunResults = append(dunResults, sourceCand)
				}
			}

			mu.Lock()
			results[dun] = dunResults
			mu.Unlock()
		}(dunCode)
	}

	wg.Wait()
	return results
}

func generateOutputLine(dunCode string, candidates []Candidate) []string {
	parts := []string{dunCode}
	totalVote := 0

	for _, cand := range candidates {
		// Party
		parts = append(parts, cand.Party)

		// Candidate name
		if cand.OrigCandidate != "" {
			parts = append(parts, fmt.Sprintf("%s (ORIG - %s)", cand.Candidate, cand.OrigCandidate))
		} else {
			parts = append(parts, cand.Candidate)
		}

		// Sex
		if cand.OrigSex != "" {
			parts = append(parts, fmt.Sprintf("%s (ORIG - %s)", cand.Sex, cand.OrigSex))
		} else {
			parts = append(parts, cand.Sex)
		}

		// Age
		ageStr := cand.Age
		if cand.OrigAge != "" {
			parts = append(parts, fmt.Sprintf("%s (ORIG - %s)", ageStr, cand.OrigAge))
		} else {
			parts = append(parts, ageStr)
		}

		// Vote
		if cand.OrigVote != "" {
			parts = append(parts, fmt.Sprintf("%s (ORIG - %s)", cand.Vote, cand.OrigVote))
		} else {
			parts = append(parts, cand.Vote)
		}

		// Add to total
		voteStr := strings.ReplaceAll(cand.Vote, ",", "")
		if vote, err := strconv.Atoi(voteStr); err == nil {
			totalVote += vote
		}
	}

	// Add total vote
	parts = append(parts, fmt.Sprintf("TOTAL: %s", formatNumber(totalVote)))

	return parts
}

func formatNumber(n int) string {
	s := strconv.Itoa(n)
	var result strings.Builder
	for i, digit := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(digit)
	}
	return result.String()
}

func main() {
	fmt.Println("Reading source of truth...")
	sourceData, err := readSourceOfTruth("/Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/Final-Sarawak-DUN-2016.csv")
	if err != nil {
		fmt.Printf("Error reading source: %v\n", err)
		return
	}

	fmt.Println("Reading GitHub data...")
	githubData, err := readGitHubData("/tmp/github_sarawak_2016.csv")
	if err != nil {
		fmt.Printf("Error reading GitHub data: %v\n", err)
		return
	}

	fmt.Println("Matching candidates using concurrency...")
	results := matchCandidates(sourceData, githubData)

	// Sort DUN codes
	dunCodes := make([]string, 0, len(results))
	for dunCode := range results {
		dunCodes = append(dunCodes, dunCode)
	}
	sort.Slice(dunCodes, func(i, j int) bool {
		return dunCodes[i] < dunCodes[j]
	})

	// Generate output
	fmt.Println("\nGenerating output...")

	// Write CSV
	csvFile, err := os.Create("/Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/matched_output.csv")
	if err != nil {
		fmt.Printf("Error creating CSV: %v\n", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Print table header
	fmt.Println("\n" + strings.Repeat("=", 200))
	fmt.Println("MATCHED RESULTS")
	fmt.Println(strings.Repeat("=", 200))

	for _, dunCode := range dunCodes {
		line := generateOutputLine(dunCode, results[dunCode])
		writer.Write(line)
		fmt.Println(strings.Join(line, " | "))
	}

	fmt.Println("\n" + strings.Repeat("=", 200))
	fmt.Printf("Total DUNs processed: %d\n", len(results))
	fmt.Println("Output saved to: /Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/matched_output.csv")
}
