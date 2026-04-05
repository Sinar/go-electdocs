package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
)

type DUN struct {
	Code string
	Name string
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	// 1. Read raw-dun.csv (reference list, no header)
	refDUNs, err := readRawDUN("raw-dun.csv")
	if err != nil {
		slog.Error("failed to read raw-dun.csv", "error", err)
		os.Exit(1)
	}
	slog.Info("loaded reference DUNs", "count", len(refDUNs))

	// 2. Read to-review.csv and extract unique DUN code+name pairs (columns 6 and 7, 1-indexed)
	reviewDUNs, err := readReviewDUN("to-review.csv")
	if err != nil {
		slog.Error("failed to read to-review.csv", "error", err)
		os.Exit(1)
	}
	slog.Info("loaded review DUNs", "unique_pairs", len(reviewDUNs))

	// Build maps: code -> name
	refMap := make(map[string]string) // code -> name from raw-dun.csv
	revMap := make(map[string]string) // code -> name from to-review.csv
	refCodes := make([]string, 0)
	revCodes := make([]string, 0)

	for _, d := range refDUNs {
		refMap[d.Code] = d.Name
		refCodes = append(refCodes, d.Code)
	}
	for _, d := range reviewDUNs {
		if _, exists := revMap[d.Code]; !exists {
			revCodes = append(revCodes, d.Code)
		}
		revMap[d.Code] = d.Name
	}

	sort.Strings(refCodes)
	sort.Strings(revCodes)

	// 3. Compare
	// a) DUNs in raw-dun.csv but NOT in to-review.csv
	var missingInReview []DUN
	for _, code := range refCodes {
		if _, found := revMap[code]; !found {
			missingInReview = append(missingInReview, DUN{Code: code, Name: refMap[code]})
		}
	}

	// b) DUNs in to-review.csv but NOT in raw-dun.csv
	var extraInReview []DUN
	for _, code := range revCodes {
		if _, found := refMap[code]; !found {
			extraInReview = append(extraInReview, DUN{Code: code, Name: revMap[code]})
		}
	}

	// c) Name mismatches (same code, different name)
	type Mismatch struct {
		Code    string
		RefName string
		RevName string
	}
	var mismatches []Mismatch
	for _, code := range refCodes {
		if revName, found := revMap[code]; found {
			if strings.TrimSpace(refMap[code]) != strings.TrimSpace(revName) {
				mismatches = append(mismatches, Mismatch{
					Code:    code,
					RefName: refMap[code],
					RevName: revName,
				})
			}
		}
	}

	// Also check for multiple different names for same code in to-review.csv
	reviewMulti, err := readReviewDUNMulti("to-review.csv")
	if err != nil {
		slog.Error("failed to re-read to-review.csv for multi-name check", "error", err)
		os.Exit(1)
	}
	type MultiName struct {
		Code  string
		Names []string
	}
	var multiNames []MultiName
	for code, names := range reviewMulti {
		if len(names) > 1 {
			nameList := make([]string, 0, len(names))
			for n := range names {
				nameList = append(nameList, n)
			}
			sort.Strings(nameList)
			multiNames = append(multiNames, MultiName{Code: code, Names: nameList})
		}
	}
	sort.Slice(multiNames, func(i, j int) bool { return multiNames[i].Code < multiNames[j].Code })

	// 4. Report
	fmt.Println("=== PHASE-2: DUN Code & Name Verification ===")
	fmt.Println()
	fmt.Printf("Reference DUNs (raw-dun.csv):    %d\n", len(refDUNs))
	fmt.Printf("Review DUN pairs (to-review.csv): %d unique codes\n", len(revMap))
	fmt.Println()

	if len(missingInReview) == 0 {
		fmt.Println("✅ No missing DUNs — all reference DUNs found in to-review.csv")
	} else {
		fmt.Printf("❌ %d DUN(s) in raw-dun.csv but MISSING from to-review.csv:\n", len(missingInReview))
		for _, d := range missingInReview {
			fmt.Printf("   - %s (%s)\n", d.Code, d.Name)
			slog.Warn("missing DUN in review", "code", d.Code, "name", d.Name)
		}
	}
	fmt.Println()

	if len(extraInReview) == 0 {
		fmt.Println("✅ No extra DUNs — to-review.csv has no codes absent from raw-dun.csv")
	} else {
		fmt.Printf("❌ %d DUN(s) in to-review.csv but NOT in raw-dun.csv:\n", len(extraInReview))
		for _, d := range extraInReview {
			fmt.Printf("   - %s (%s)\n", d.Code, d.Name)
			slog.Warn("extra DUN in review", "code", d.Code, "name", d.Name)
		}
	}
	fmt.Println()

	if len(mismatches) == 0 {
		fmt.Println("✅ No name mismatches — all matching codes have identical names")
	} else {
		fmt.Printf("❌ %d DUN name mismatch(es):\n", len(mismatches))
		fmt.Println("   Code       | raw-dun.csv Name       | to-review.csv Name")
		fmt.Println("   -----------|------------------------|-----------------------")
		for _, m := range mismatches {
			fmt.Printf("   %-10s | %-22s | %s\n", m.Code, m.RefName, m.RevName)
			slog.Warn("name mismatch", "code", m.Code, "ref_name", m.RefName, "rev_name", m.RevName)
		}
	}
	fmt.Println()

	if len(multiNames) == 0 {
		fmt.Println("✅ No inconsistent naming within to-review.csv — each code maps to exactly one name")
	} else {
		fmt.Printf("⚠️  %d DUN code(s) have multiple names within to-review.csv:\n", len(multiNames))
		for _, mn := range multiNames {
			fmt.Printf("   - %s: %v\n", mn.Code, mn.Names)
			slog.Warn("multiple names for code in review", "code", mn.Code, "names", mn.Names)
		}
	}

	fmt.Println()
	fmt.Println("=== Full DUN mapping (reference vs review) ===")
	fmt.Println()
	fmt.Println("| Code  | raw-dun.csv Name       | to-review.csv Name     | Status |")
	fmt.Println("|-------|------------------------|------------------------|--------|")
	for _, code := range refCodes {
		refName := refMap[code]
		revName, found := revMap[code]
		status := "✅ Match"
		if !found {
			revName = "(MISSING)"
			status = "❌ Missing"
		} else if strings.TrimSpace(refName) != strings.TrimSpace(revName) {
			status = "❌ Mismatch"
		}
		fmt.Printf("| %-5s | %-22s | %-22s | %s |\n", code, refName, revName, status)
	}

	slog.Info("phase-2 check complete",
		"ref_count", len(refDUNs),
		"review_unique_codes", len(revMap),
		"missing", len(missingInReview),
		"extra", len(extraInReview),
		"mismatches", len(mismatches),
		"multi_names", len(multiNames),
	)
}

func readRawDUN(path string) ([]DUN, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1 // variable fields
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	var duns []DUN
	for i, rec := range records {
		if len(rec) < 4 {
			slog.Warn("skipping short row in raw-dun.csv", "row", i+1, "fields", len(rec))
			continue
		}
		code := strings.TrimSpace(rec[2])
		name := strings.TrimSpace(rec[3])
		if code == "" {
			continue
		}
		duns = append(duns, DUN{Code: code, Name: name})
		slog.Debug("raw-dun entry", "row", i+1, "code", code, "name", name)
	}
	return duns, nil
}

func readReviewDUN(path string) ([]DUN, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	seen := make(map[string]bool)
	var duns []DUN

	for i, rec := range records {
		if i == 0 {
			// skip header
			continue
		}
		if len(rec) < 7 {
			slog.Warn("skipping short row in to-review.csv", "row", i+1, "fields", len(rec))
			continue
		}
		code := strings.TrimSpace(rec[5]) // column 6, 0-indexed = 5
		name := strings.TrimSpace(rec[6]) // column 7, 0-indexed = 6
		key := code + "|" + name
		if !seen[key] {
			seen[key] = true
			duns = append(duns, DUN{Code: code, Name: name})
			slog.Debug("review DUN entry", "code", code, "name", name)
		}
	}
	return duns, nil
}

// readReviewDUNMulti returns code -> set of names found in to-review.csv
func readReviewDUNMulti(path string) (map[string]map[string]bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	result := make(map[string]map[string]bool)
	for i, rec := range records {
		if i == 0 {
			continue
		}
		if len(rec) < 7 {
			continue
		}
		code := strings.TrimSpace(rec[5])
		name := strings.TrimSpace(rec[6])
		if result[code] == nil {
			result[code] = make(map[string]bool)
		}
		result[code][name] = true
	}
	return result, nil
}
