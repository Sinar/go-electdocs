package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	ID             string
	PollingCentre  string
	VotingChannel  string
	LineNumber     int
}

func main() {
	file, err := os.Open("Final-Sarawak-DUN-2016.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	reader.LazyQuotes = true

	// Skip header
	_, err = reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Map to track unique IDs and their records
	idMap := make(map[string][]Record)
	lineNumber := 2 // Start from 2 (line 1 is header)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Check if first column exists and is not empty
		if len(record) > 0 && strings.TrimSpace(record[0]) != "" {
			id := strings.TrimSpace(record[0])

			// Get polling centre (column index 9) and voting channel (column index 10)
			pollingCentre := ""
			votingChannel := ""

			if len(record) > 9 {
				pollingCentre = strings.TrimSpace(record[9])
			}
			if len(record) > 10 {
				votingChannel = strings.TrimSpace(record[10])
			}

			rec := Record{
				ID:            id,
				PollingCentre: pollingCentre,
				VotingChannel: votingChannel,
				LineNumber:    lineNumber,
			}

			idMap[id] = append(idMap[id], rec)
		}

		lineNumber++
	}

	// Find duplicates
	duplicates := make(map[string][]Record)
	for id, records := range idMap {
		if len(records) > 1 {
			duplicates[id] = records
		}
	}

	// Report results
	if len(duplicates) == 0 {
		fmt.Println("✓ All IDs in the first column are unique!")
		fmt.Printf("\nTotal unique IDs analyzed: %d\n", len(idMap))
	} else {
		fmt.Printf("✗ Found %d duplicate ID(s):\n\n", len(duplicates))

		// Print table header
		fmt.Printf("%-50s | %-10s | %-70s | %-20s\n", "UNIQUE CODE (ID)", "Line #", "POLLING CENTRE", "VOTING CHANNEL")
		fmt.Println(strings.Repeat("-", 160))

		// Print duplicates
		for id, records := range duplicates {
			fmt.Printf("\nDuplicate ID: %s (appears %d times)\n", id, len(records))
			for _, rec := range records {
				fmt.Printf("  %-50s | %-10d | %-70s | %-20s\n",
					rec.ID,
					rec.LineNumber,
					rec.PollingCentre,
					rec.VotingChannel)
			}
		}

		fmt.Printf("\n\nTotal duplicate IDs: %d\n", len(duplicates))
		fmt.Printf("Total unique IDs: %d\n", len(idMap))
	}
}
