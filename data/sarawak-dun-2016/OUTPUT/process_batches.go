package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Batch struct {
	Start int
	End   int
}

func processFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Skip first line (header)
	if scanner.Scan() {
		// header skipped
	}

	// Read all data lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Remove trailing empty lines and summary line from the end
	for len(lines) > 0 {
		lastLine := strings.TrimSpace(lines[len(lines)-1])
		if lastLine == "" || strings.HasPrefix(lastLine, "x,") {
			lines = lines[:len(lines)-1]
		} else {
			break
		}
	}

	return lines, nil
}

func createBatch(start, end int, outputDir string) error {
	var allData []string

	for num := start; num <= end; num++ {
		filename := fmt.Sprintf("Sarawak-N.%02d.csv", num)
		filepath := filepath.Join(outputDir, filename)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			fmt.Printf("Warning: %s not found, skipping...\n", filename)
			continue
		}

		fmt.Printf("Processing %s...\n", filename)

		dataRows, err := processFile(filepath)
		if err != nil {
			return fmt.Errorf("error processing %s: %v", filename, err)
		}

		// Add data rows to combined data
		allData = append(allData, dataRows...)

		// Add empty row between files (except after the last file)
		if num < end {
			allData = append(allData, "")
		}
	}

	// Write combined file
	outputFilename := fmt.Sprintf("combined-N.%02d_N.%02d.csv", start, end)
	outputPath := filepath.Join(outputDir, outputFilename)

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range allData {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	writer.Flush()

	fmt.Printf("Created %s with %d lines\n", outputFilename, len(allData))
	return nil
}

func main() {
	outputDir := "."

	batches := []Batch{
		{1, 10},
		{11, 20},
		{21, 30},
		{31, 40},
		{41, 50},
		{51, 60},
		{61, 70},
		{71, 81},
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(batches))

	for _, batch := range batches {
		wg.Add(1)
		go func(b Batch) {
			defer wg.Done()
			fmt.Printf("\n%s\n", strings.Repeat("=", 60))
			fmt.Printf("Processing batch N.%02d - N.%02d\n", b.Start, b.End)
			fmt.Printf("%s\n", strings.Repeat("=", 60))

			if err := createBatch(b.Start, b.End, outputDir); err != nil {
				errChan <- err
			}
		}(batch)
	}

	wg.Wait()
	close(errChan)

	// Check for errors
	hasErrors := false
	for err := range errChan {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		hasErrors = true
	}

	if !hasErrors {
		fmt.Printf("\n%s\n", strings.Repeat("=", 60))
		fmt.Println("All batches processed successfully!")
		fmt.Printf("%s\n", strings.Repeat("=", 60))
	} else {
		os.Exit(1)
	}
}
