# Transformation Rules for Sarawak Election Data

This document outlines the rules for transforming the raw CSV data from the `DATA/CSV` directory into the format required for the output CSV files in the `DATA/OUTPUT` directory.

## Input Files

The raw data for each state constituency is provided as a CSV file in the `DATA/CSV` directory. The file naming convention is `N.XX <STATE_CONSTITUENCY_NAME>.csv`, where `XX` is the state constituency code.

## Output Files

The transformed data for each state constituency should be saved as a CSV file in the `DATA/OUTPUT` directory. The file naming convention is `Sarawak-N.XX.csv`, where `XX` is the state constituency code.

## Transformation Rules

The following rules should be applied to each row of the input CSV file to generate the corresponding row in the output CSV file.

*   **`UNIQUE CODE`**: A unique identifier constructed by concatenating the Parliamentary Constituency Code (e.g., `P.192`), the State Constituency Code (e.g., `_N.02_`), the `KODDM` (with slashes replaced by underscores), an underscore, and the `Nombor Tempat Mengundi (saluran)`.
*   **`STATE`**: This is a fixed value: `SARAWAK`.
*   **`BALLOT TYPE`**: Determined by the `NAMADM` field:
    *   `UNDI POS` becomes `POSTAL VOTE`.
    *   `UNDI AWAL` becomes `EARLY VOTE`.
    *   All other values are mapped to `ORDINARY VOTE`.
*   **`PARLIAMENTARY CONSTITUENCY CODE`**: The Parliamentary Constituency Code for the State Constituency (e.g., `P.192`).
*   **`PARLIAMENTARY CONSTITUENCY NAME`**: The Parliamentary Constituency Name for the State Constituency (e.g., `MAS GADING`).
*   **`STATE CONSTITUENCY CODE`**: The State Constituency Code from the file name (e.g., `N.02`).
*   **`STATE CONSTITUENCY NAME`**: The State Constituency Name from the file name (e.g., `TASIK BIRU`).
*   **`POLLING DISTRICT CODE`**: Directly mapped from the `KODDM` field.
*   **`POLLING DISTRICT NAME`**: Directly mapped from the `NAMADM` field.
*   **`POLLING CENTRE`**: Directly mapped from the `Nama Pusat Mengundi` field.
*   **`VOTING CHANNEL NUMBER`**: Directly mapped from the `Nombor Tempat Mengundi (saluran)` field.
*   **`TOTAL BALLOTS ISSUED`**: Directly mapped from the `Jumlah Kertas Undi Yang Patut Berada Di Dalam Peti Undi (A)` field.
*   **Candidate Information**: The candidate names from the source CSV are mapped to specific party slots in the output. The party affiliation needs to be determined for each candidate.
    *   **`BN`**: Barisan Nasional candidates.
    *   **`PH (1)`**: Pakatan Harapan candidates from PKR and PAN (Amanah).
    *   **`PH (2)`**: Pakatan Harapan candidates from DAP.
    *   **`PAS`**: Parti Islam Se-Malaysia candidates.
    *   **`STAR`**: State Reform Party candidates.
    *   **`PBDSB`**: Parti Bangsa Dayak Sarawak Baru candidates.
    *   **`INDEPENDENT 1`**, **`INDEPENDENT 2`**: Independent candidates.
    *   Candidate details like sex and age are not present in the source file and should be left blank.
*   **`TOTAL VALID VOTES`**: Mapped from `Bilangan Undian Oleh Pemilih Bagi Setiap Orang Calon Yang Bertanding (B) :Jumlah Undian Oleh Pemilih`.
*   **`TOTAL REJECTED VOTES`**: Mapped from `Bilangan Kertas Undi Yang Ditolak (C)`.
*   **`TOTAL UNRETURNED BALLOTS`**: Mapped from `Jumlah Kertas Undi Yang Dikeluarkan Kepada Pengundi Tetapi Tidak Dimasukkan Ke Dalam Peti Undi(D)`.
*   **`CHECKER (VALID VOTE)`** and **`CHECKER (TOTAL VOTE ISSUED)`**: These columns are left empty.

## Important Notes

*   **Uncontested Seats**: PLEASE TAKE NOTE THERE WERE TWO UNCONTESTED SEATS - BUKIT KOTA AND BUKIT SARI. NO NEED TO COMPILE ANY RESULTS FOR THESE DUNS.
*   **Party Clashes**: PLEASE TAKE NOTE THERE WERE CLASHES OF PH COMPONENT PARTIES BETWEEN PKR AND DAP.
    *   FOR PH (1) - Put PKR and PAN (Amanah).
    *   FOR PH (2) - Put DAP only.

## Workflow Learnings and Process Summary

### Overview of Process
- **Objective**: Transform raw Sarawak election CSV data (N.01 to N.81) into standardized output format (`Sarawak-N.XX.csv`) following specified rules.
- **Tools Used**: Go scripts for transformation, terminal commands for execution, file operations for copying/editing.
- **Data Sources**: Input CSVs from `DATA/CSV`, party mappings from 2021 election data (Wikipedia), output to `DATA/OUTPUT`.
- **Completion Status**: All 80 available constituencies processed (N.01-N.81 except uncontested N.79 BUKIT KOTA and N.82 BUKIT SARI).

### Key Challenges and Solutions
- **File Naming Inconsistencies**: Some files have apostrophes (e.g., BA'KELALAN vs. BA_KELALAN). Solution: Verify file names in directory listing and adjust `getName` function accordingly.
- **Party Mapping Accuracy**: Candidate parties not in source CSV; required external research. Solution: Use Wikipedia 2021 election results for mappings, hardcode in Go scripts.
- **Large Dataset Handling**: 80 constituencies. Solution: Process in small batches (e.g., 10-12 constituencies per script) for reliability and error isolation.
- **Uncontested Seats**: N.79 and N.82 skipped as per rules; no input files present.
- **Script Errors**: Main redeclared (multiple Go files); resolved by running scripts individually from project root.

### Party Mapping Methodology
- **Source**: 2021 Sarawak state election Wikipedia page for candidate lists and affiliations.
- **Mapping Rules**:
  - BN: GPS components (PBB, SUPP, PRS, PDP).
  - PH(1): PKR, PAN (Amanah).
  - PH(2): DAP.
  - PAS: As listed.
  - STAR: As listed.
  - PBDSB: As listed.
  - INDEPENDENT 1/2: PSB, GAS parties (ASPIRASI, SEDAR, PBK), independents.
- **Implementation**: Hardcoded in `getParty` function per constituency; default to INDEPENDENT 1 if unmatched.

### Batch Processing Strategy
- **Batch Sizes**: 10-12 constituencies per Go script (e.g., N.11-20, N.21-32) to ensure slow, steady progress.
- **Script Structure**: Each batch script copies previous, updates ns array, adds party cases, updates getName/getPCode/getPName.
- **Execution**: Run via `go run DATA/transform_nXX_to_nYY.go` from project root.
- **Parallel Potential**: Data self-contained; future agents can run multiple scripts simultaneously if needed.

### Completed Constituencies
- **Total Processed**: 80 files (N.01-N.81, skipping N.79).
- **Batches**:
  - N.01-N.10: Pre-existing.
  - N.11-N.20: Initial request.
  - N.21-N.32: Batch 1.
  - N.33-N.48: Batch 2.
  - N.49-N.60: Batch 3.
  - N.61-N.70: Batch 4.
  - N.71-N.81: Batch 5.
- **Output Verification**: All `Sarawak-N.XX.csv` files present in `DATA/OUTPUT`; format matches rules.

### Recommendations for Future Agents
- **Verify Inputs**: Always list `DATA/CSV` to confirm file names and availability.
- **External Data**: Use Wikipedia or official sources for party mappings; update if election year changes.
- **Error Handling**: Test scripts individually; check for file path issues.
- **Scalability**: For larger datasets, consider automating party mapping via API or database.
- **Documentation**: Update this file with new learnings after each major task.
- **Backup**: Preserve original scripts for reference.

## Duplicate ID Resolution

### Problem Identification
After initial transformation, duplicate UNIQUE CODE values were identified in column 1 across multiple output files. These duplicates violated the uniqueness constraint required for proper data integrity.

### Root Cause
Some combinations of Parliamentary Constituency Code, State Constituency Code, KODDM, and Voting Channel Number were not unique across different Polling Centers within the same file, resulting in identical UNIQUE CODE values.

### Solution: Suffix Assignment Algorithm
A Python script (`fix_duplicate_ids.py`) was created to resolve duplicates by appending alphabetical suffixes based on Polling Center:

**Rule**: When duplicate IDs exist in a file:
1. Group all occurrences of the duplicate ID by their Polling Center (column 10)
2. Assign suffixes (a, b, c, d...) to ALL occurrences of the duplicate ID
3. Same Polling Center gets the same suffix letter
4. Suffix assignment is based on the order of first appearance of each unique Polling Center
5. ONLY column 1 (UNIQUE CODE) is modified; all other columns remain unchanged

**Example**:
```
Original duplicate IDs:
- P.220_N.77_220/77/01_1 at SK KAMPONG TANJONG ASSAM → P.220_N.77_220/77/01_1a
- P.220_N.77_220/77/01_1 at SK NYABOR → P.220_N.77_220/77/01_1b
- P.220_N.77_220/77/01_1 at SJK CHUNG HUA NYABOR → P.220_N.77_220/77/01_1c
```

### Implementation Details
- **Script**: `/Users/leow/GOMOD/go-electdocs/data/sarawak-dun-2016/OUTPUT/fix_duplicate_ids.py`
- **Execution**: `uv run fix_duplicate_ids.py` (standalone with contained dependencies)
- **Files Processed**: All N.01-N.81 files (80 files total; N.79 doesn't exist)
- **Results**:
  - 1,527 duplicate IDs fixed across 69 files
  - 12 files had no duplicates (N.01, N.03, N.09, N.18, N.24, N.35, N.54, N.56, N.78, and others)

### Key Learnings
- **Data Integrity**: Same Polling Center with same channel number cannot exist (they are always different centers)
- **All Suffixes Required**: When duplicates exist, ALL occurrences get suffixes starting from 'a' (not just subsequent ones)
- **Column-Specific Modification**: Only modify column 1; preserve all other data including spacing and formatting
- **Ignore Non-Data Rows**: Skip header rows, summary rows (no polling center), and empty lines when identifying duplicates
- **Standalone Scripts**: Use `uv` for Python scripts to ensure portability and contained dependencies

### Verification
After execution, all UNIQUE CODE values in column 1 are guaranteed unique within each file while maintaining the relationship to their respective Polling Centers.