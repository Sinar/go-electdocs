package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/xuri/excelize/v2"
)

type row map[string]string

func main() {
	excelPath := "XLSX/N.02 TASIK BIRU.xlsx"
	csvPath := "CSV/N.02 TASIK BIRU.csv"
	sheetName := "" // leave empty to auto-pick first sheet

	if err := run(excelPath, csvPath, sheetName); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func run(excelPath, csvPath, sheetName string) error {
	if !fileExists(excelPath) {
		return fmt.Errorf("excel file missing: %s", excelPath)
	}
	if !fileExists(csvPath) {
		return fmt.Errorf("csv file missing: %s", csvPath)
	}

	exRows, exCols, err := loadExcel(excelPath, sheetName)
	if err != nil {
		return fmt.Errorf("load excel: %w", err)
	}
	csvRows, csvCols, err := loadCSV(csvPath)
	if err != nil {
		return fmt.Errorf("load csv: %w", err)
	}

	fmt.Printf("Excel: %d rows, %d columns\n", len(exRows), len(exCols))
	fmt.Printf("CSV:   %d rows, %d columns\n", len(csvRows), len(csvCols))

	// Column set comparison
	missingInCSV := diffStrings(exCols, csvCols)
	missingInExcel := diffStrings(csvCols, exCols)
	if len(missingInCSV) == 0 && len(missingInExcel) == 0 {
		fmt.Println("Column sets match.")
	} else {
		if len(missingInCSV) > 0 {
			fmt.Printf("Columns present in Excel but missing in CSV: %v\n", missingInCSV)
		}
		if len(missingInExcel) > 0 {
			fmt.Printf("Columns present in CSV but missing in Excel: %v\n", missingInExcel)
		}
	}

	// Normalize columns intersection
	common := intersectStrings(exCols, csvCols)
	if len(common) == 0 {
		return errors.New("no overlapping columns; cannot compare rows meaningfully")
	}

	// Build canonical serialized representation for order-agnostic comparison
	exSet := map[string]int{}
	for _, r := range exRows {
		exSet[serializeRow(r, common)]++
	}
	csvSet := map[string]int{}
	for _, r := range csvRows {
		csvSet[serializeRow(r, common)]++
	}

	var onlyInExcel, onlyInCSV int
	for sig, cnt := range exSet {
		if csvSet[sig] < cnt {
			onlyInExcel += cnt - csvSet[sig]
		}
	}
	for sig, cnt := range csvSet {
		if exSet[sig] < cnt {
			onlyInCSV += cnt - exSet[sig]
		}
	}

	if onlyInExcel == 0 && onlyInCSV == 0 {
		fmt.Println("All rows match (order-agnostic, counting duplicates).")
	} else {
		fmt.Printf("Row differences (order-agnostic): only in Excel=%d, only in CSV=%d\n", onlyInExcel, onlyInCSV)
		showSampleDiff(exSet, csvSet)
	}

	// Order-sensitive + cell diff (only if row counts same and we assume aligned)
	if len(exRows) == len(csvRows) {
		diffCount := 0
		maxReport := 20
		for i := 0; i < len(exRows) && diffCount < maxReport; i++ {
			diffs := cellDiff(exRows[i], csvRows[i], common)
			if len(diffs) > 0 {
				if diffCount == 0 {
					fmt.Println("First order-sensitive cell differences (row index = 0-based):")
				}
				fmt.Printf("Row %d:\n", i)
				for _, d := range diffs {
					fmt.Printf("  %s: Excel=%q CSV=%q\n", d.col, d.a, d.b)
				}
				diffCount++
			}
		}
		if diffCount == 0 {
			fmt.Println("No cell differences under an order-sensitive assumption.")
		} else if diffCount == maxReport {
			fmt.Println("... truncated additional differences ...")
		}
	} else {
		fmt.Println("Skipping order-sensitive cell-by-cell diff (row counts differ).")
	}

	return nil
}

func loadExcel(path, sheet string) ([]row, []string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	if sheet == "" {
		sheets := f.GetSheetList()
		if len(sheets) == 0 {
			return nil, nil, errors.New("no sheets in workbook")
		}
		sheet = sheets[0]
	}

	rows, err := f.GetRows(sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		return nil, nil, err
	}
	if len(rows) == 0 {
		return nil, nil, errors.New("empty sheet")
	}

	// Debug: optionally dump raw rows as strings for troubleshooting
	if os.Getenv("DUMP_RAW") == "1" {
		fmt.Println("=== RAW EXCEL ROWS (unprocessed) ===")
		for i, rr := range rows {
			joined := strings.Join(rr, "||")
			fmt.Printf("Row %d (len=%d): %s\n", i, len(rr), joined)
		}
		fmt.Println("=== END RAW EXCEL ROWS ===")
	}

	// Heuristic: find the first row that looks like a header (>=2 non-empty cells).
	headerIdx := 0
	for i, r := range rows {
		if countNonEmpty(r) >= 2 {
			headerIdx = i
			break
		}
	}

	// Build header by optionally merging the header row with the next row
	// when the file uses multi-row headers (common in exported sheets).
	rawHeader := make([]string, 0)
	rawHeader = append(rawHeader, rows[headerIdx]...)
	// If the next row also contains meaningful labels, merge empty header cells from it.
	if headerIdx+1 < len(rows) && countNonEmpty(rows[headerIdx+1]) >= 2 {
		next := rows[headerIdx+1]
		// Ensure rawHeader is long enough
		if len(next) > len(rawHeader) {
			tmp := make([]string, len(next))
			copy(tmp, rawHeader)
			rawHeader = tmp
		}
		for i := 0; i < len(next); i++ {
			if strings.TrimSpace(rawHeader[i]) == "" {
				rawHeader[i] = next[i]
			} else if strings.TrimSpace(next[i]) != "" {
				// If both rows have text, join them with a space to preserve info
				rawHeader[i] = strings.TrimSpace(rawHeader[i]) + " " + strings.TrimSpace(next[i])
			}
		}
	}

	if os.Getenv("DUMP_RAW") == "1" {
		fmt.Printf("Detected header at row %d (0-based), raw header cols=%d:\n", headerIdx, len(rawHeader))
		for i, v := range rawHeader {
			fmt.Printf("  col[%d]=%q\n", i, v)
		}
	}

	headers := normalizeHeaders(rawHeader)
	var data []row
	// Data starts after the header row; if we merged with next row then skip an extra row.
	dataStart := headerIdx + 1
	if headerIdx+1 < len(rows) && countNonEmpty(rows[headerIdx+1]) >= 2 {
		dataStart = headerIdx + 2
	}
	for _, raw := range rows[dataStart:] {
		if isEmptyRow(raw) {
			continue
		}
		r := row{}
		for i, h := range headers {
			var v string
			if i < len(raw) {
				v = strings.TrimSpace(raw[i])
			}
			r[h] = canonicalizeValue(v)
		}
		data = append(data, r)
	}
	return data, headers, nil
}

func countNonEmpty(cells []string) int {
	n := 0
	for _, c := range cells {
		if strings.TrimSpace(c) != "" {
			n++
		}
	}
	return n
}

func loadCSV(path string) ([]row, []string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	rawHeaders, err := r.Read()
	if err != nil {
		return nil, nil, err
	}
	headers := normalizeHeaders(rawHeaders)

	var data []row
	for {
		rec, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		if isEmptyRow(rec) {
			continue
		}
		rowMap := row{}
		for i, h := range headers {
			var v string
			if i < len(rec) {
				v = strings.TrimSpace(rec[i])
			}
			rowMap[h] = canonicalizeValue(v)
		}
		data = append(data, rowMap)
	}
	return data, headers, nil
}

func normalizeHeaders(h []string) []string {
	out := make([]string, len(h))
	seen := map[string]int{}
	for i, v := range h {
		base := strings.TrimSpace(strings.ToLower(v))
		base = strings.ReplaceAll(base, " ", "_")
		if base == "" {
			base = fmt.Sprintf("col_%d", i)
		}
		if seen[base] > 0 {
			base = fmt.Sprintf("%s_%d", base, seen[base])
		}
		seen[base]++
		out[i] = base
	}
	return out
}

func canonicalizeValue(s string) string {
	// Could add numeric normalization here if desired
	return s
}

func isEmptyRow(cells []string) bool {
	for _, c := range cells {
		if strings.TrimSpace(c) != "" {
			return false
		}
	}
	return true
}

func diffStrings(a, b []string) []string {
	setB := map[string]struct{}{}
	for _, v := range b {
		setB[v] = struct{}{}
	}
	var out []string
	for _, v := range a {
		if _, ok := setB[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

func intersectStrings(a, b []string) []string {
	setB := map[string]struct{}{}
	for _, v := range b {
		setB[v] = struct{}{}
	}
	var out []string
	for _, v := range a {
		if _, ok := setB[v]; ok {
			out = append(out, v)
		}
	}
	sort.Strings(out)
	return out
}

func serializeRow(r row, cols []string) string {
	var b strings.Builder
	for _, c := range cols {
		b.WriteString(r[c])
		b.WriteRune('\x1f') // unit separator
	}
	return b.String()
}

type diff struct {
	col string
	a   string
	b   string
}

func cellDiff(a, b row, cols []string) []diff {
	var out []diff
	for _, c := range cols {
		if a[c] != b[c] {
			out = append(out, diff{col: c, a: a[c], b: b[c]})
		}
	}
	return out
}

func showSampleDiff(exSet, csvSet map[string]int) {
	fmt.Println("Sample differing signatures (truncated):")
	count := 0
	for sig, exCnt := range exSet {
		csvCnt := csvSet[sig]
		if exCnt > csvCnt {
			fmt.Printf("  RowSig hash=%x onlyInExcel count=%d\n", hash32(sig), exCnt-csvCnt)
			count++
			if count >= 5 {
				break
			}
		}
	}
	count = 0
	for sig, csvCnt := range csvSet {
		exCnt := exSet[sig]
		if csvCnt > exCnt {
			fmt.Printf("  RowSig hash=%x onlyInCSV count=%d\n", hash32(sig), csvCnt-exCnt)
			count++
			if count >= 5 {
				break
			}
		}
	}
}

func hash32(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func fileExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && !st.IsDir()
}
