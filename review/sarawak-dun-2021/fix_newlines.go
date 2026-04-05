// fix_newlines.go removes embedded newline characters from POLLING CENTRE
// (column 10) in to-review.csv.
//
// There are exactly 5 distinct values that contain a '\n'.  Each is fixed with
// an explicit, hand-verified replacement so the join character (nothing, space,
// or stripped punctuation) is correct for that particular name:
//
//	RUANG D, DEWAN BADMINTON\n, KOMPLEKS POLIS TABUAN JAYA
//	  → RUANG D, DEWAN BADMINTON, KOMPLEKS POLIS TABUAN JAYA
//
//	SEKOLAH KEBANGSAAN BANDARAN SIBU NO.\n2
//	  → SEKOLAH KEBANGSAAN BANDARAN SIBU NO.2
//
//	BANGUNAN PERSEKUTUAN PERKUMPULAN WANITA SARAWAK DAERAH BINTULU ( W.I\n)
//	  → BANGUNAN PERSEKUTUAN PERKUMPULAN WANITA SARAWAK DAERAH BINTULU ( W.I)
//
//	SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH\nAS-SYIBYAN)
//	  → SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH AS-SYIBYAN)
//
//	SEKOLAH KEBANGSAAN R.\nC. KUBONG
//	  → SEKOLAH KEBANGSAAN R.C. KUBONG
//
// Usage (run from the sarawak-dun-2021 directory):
//
//	go run fix_newlines.go
//
// The input file (to-review.csv) is overwritten in place.
// A backup is written to to-review.csv.bak2 before any changes are made.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// colPollingCentre is the 0-based index of the POLLING CENTRE column.
const colPollingCentre = 9

// fixes maps the exact current (broken) value to the correct replacement.
// Keys use a real newline character (\n) — same as what is stored in the CSV.
var fixes = map[string]string{
	"RUANG D, DEWAN BADMINTON\n, KOMPLEKS POLIS TABUAN JAYA":                  "RUANG D, DEWAN BADMINTON, KOMPLEKS POLIS TABUAN JAYA",
	"SEKOLAH KEBANGSAAN BANDARAN SIBU NO.\n2":                                 "SEKOLAH KEBANGSAAN BANDARAN SIBU NO.2",
	"BANGUNAN PERSEKUTUAN PERKUMPULAN WANITA SARAWAK DAERAH BINTULU ( W.I\n)": "BANGUNAN PERSEKUTUAN PERKUMPULAN WANITA SARAWAK DAERAH BINTULU ( W.I)",
	"SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH\nAS-SYIBYAN)":                 "SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH AS-SYIBYAN)",
	"SEKOLAH KEBANGSAAN R.\nC. KUBONG":                                        "SEKOLAH KEBANGSAAN R.C. KUBONG",
}

func main() {
	slog.Info("fix_newlines.go starting")

	const inputPath = "to-review.csv"
	const backupPath = "to-review.csv.bak2"

	// ── Read ──────────────────────────────────────────────────────────────────
	f, err := os.Open(inputPath)
	if err != nil {
		slog.Error("cannot open input file", "path", inputPath, "err", err)
		os.Exit(1)
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	f.Close()
	if err != nil {
		slog.Error("failed to read CSV", "err", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		slog.Error("CSV has no data rows")
		os.Exit(1)
	}

	header := records[0]
	data := records[1:]
	slog.Info("CSV loaded", "header_cols", len(header), "data_rows", len(data))

	// ── Backup ────────────────────────────────────────────────────────────────
	if err := copyFile(inputPath, backupPath); err != nil {
		slog.Error("cannot create backup", "src", inputPath, "dst", backupPath, "err", err)
		os.Exit(1)
	}
	slog.Info("backup created", "path", backupPath)

	// ── Apply fixes ───────────────────────────────────────────────────────────
	rowsFixed := 0
	unknownNewlines := 0

	for i, row := range data {
		if len(row) <= colPollingCentre {
			continue
		}
		cell := row[colPollingCentre]
		if !strings.Contains(cell, "\n") {
			continue
		}

		if replacement, known := fixes[cell]; known {
			slog.Info("fixing polling centre newline",
				"data_row", i+1,
				"before", strings.ReplaceAll(cell, "\n", `\n`),
				"after", replacement,
			)
			data[i][colPollingCentre] = replacement
			rowsFixed++
		} else {
			// Unexpected — log and leave unchanged so nothing is silently mangled.
			slog.Warn("unknown newline in polling centre — not fixed",
				"data_row", i+1,
				"value", strings.ReplaceAll(cell, "\n", `\n`),
			)
			unknownNewlines++
		}
	}

	slog.Info("fix pass complete",
		"rows_fixed", rowsFixed,
		"unknown_newlines_skipped", unknownNewlines,
	)

	// ── Post-fix validation ───────────────────────────────────────────────────
	remaining := 0
	for _, row := range data {
		for _, cell := range row {
			if strings.Contains(cell, "\n") {
				remaining++
			}
		}
	}
	if remaining == 0 {
		slog.Info("validation: no remaining newlines in any cell")
	} else {
		slog.Warn("validation: newlines still present in some cells", "count", remaining)
	}

	if unknownNewlines > 0 {
		slog.Warn("there were unknown newline cases — review the log and update the fixes map")
	}

	// ── Write output ──────────────────────────────────────────────────────────
	out, err := os.Create(inputPath)
	if err != nil {
		slog.Error("cannot create output file", "path", inputPath, "err", err)
		os.Exit(1)
	}
	w := csv.NewWriter(out)

	if err := w.Write(header); err != nil {
		out.Close()
		slog.Error("failed to write header", "err", err)
		os.Exit(1)
	}
	for _, row := range data {
		if err := w.Write(row); err != nil {
			out.Close()
			slog.Error("failed to write data row", "err", err)
			os.Exit(1)
		}
	}
	w.Flush()
	out.Close()
	if err := w.Error(); err != nil {
		slog.Error("CSV writer flush error", "err", err)
		os.Exit(1)
	}

	slog.Info("to-review.csv written successfully", "total_rows", len(data))

	fmt.Printf("\n=== Newline Fix Summary ===\n")
	fmt.Printf("Rows fixed:               %d\n", rowsFixed)
	fmt.Printf("Unknown cases skipped:    %d\n", unknownNewlines)
	fmt.Printf("Remaining newline cells:  %d\n", remaining)
	fmt.Printf("Output: %s  (backup: %s)\n", inputPath, backupPath)
}

// copyFile copies src to dst, creating dst if it does not exist.
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open src: %w", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create dst: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy: %w", err)
	}
	return out.Sync()
}
