# AGENTS

## OBJECTIVE

To thoroughly review the to-review.csv and see if any of the data is not matching, has typo, does not add up and is placed under wrong location.  Any analysis should be backed by strong evidence.

The raw data is from the official Election Commission website so can be assumed to be the authoritative version. These files are prefixed with raw-*.csv

If needed, we can breakdown the original PDFs to Markdown and use those data to reconfirm anything suspicous. Ask for help if this step is needed

## REVIEW PROCESS

We will start the review of the file to-review.csv in phases from easiest to hardest. The number of columns should be the same and aligned logically; you will NEVER find one column having candidate name and another with the DMKOD 

Do each phase one by one; at the end of the phase write out the review of the phase in the file PHASE-<NUMBER>-REVIEW.md

### PHASE-0: Check against PRN DUN Sarawak 2016 results

The files in the folder /Users/leow/TINDAKMSIA/go-electdocs/data/sarawak-dun-2016/OUTPUT with the pattern Sarawak-N.<DUN>.csv should be compared against to-review.csv.  Go compare DUN by DUN (there should be 81) so easy to extract; note DUN N.79 + N.82 in 2016 did not compete so no data.

The expected changes might be the candidates and votes but actually most should be the same.  Compare against the fields that should be the same (expected a few changes as there might be new areas) but most are the same: 

     1  UNIQUE CODE 
     2	STATE
     3	BALLOT TYPE
     4	PARLIAMENTARY CONSTITUENCY CODE
     5	PARLIAMENTARY CONSTITUENCY NAME
     6	STATE CONSTITUENCY CODE
     7	STATE CONSTITUENCY NAME
     8	POLLING DISTRICT CODE
     9	POLLING DISTRICT NAME
    10	POLLING CENTRE
    11	VOTING CHANNEL NUMBER

Highlight differences in table and justify with evidence. 


### PHASE-1: Ensure ID field is unique

The first column of the file to be reviewed is UNIQUE CODE.

This field is unique and if there are non-unique value (other than empty space); point them out if NOT following the rules

*Rule**: When duplicate IDs exist in a file:
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


### PHASE-2: Find Missing OR Incorrect DUN + PAR

Check against the list in raw-dun.csv + raw-par.csv. Flag out incorrect DUN name and if not all DUNs are matched.  Do the same for PAR.  There should not be any unaccounted for.

### PHASE-3: Find Missing OR Incorrect Candidates

Check against the list in raw-candidates.csv.  Flag out incorrect name and if not all candidates are matched.  There should not be any unaccounted for.


### PHASE-4: Check consistency of coallition

In the Sarawak state assembly is made out of many small parties, a few major co-oaltion + possibly few independents.  Point out if anything not found in the official data.
 
Use the raw-party-data-clean.csv file; analyze the header to extract the semantic meaning and map as per necessary

- **Source**: 2021 Sarawak state election Wikipedia page for candidate lists and affiliations.
- **Mapping Rules**:
  - BN: GPS components (PBB, SUPP, PRS, PDP).
  - PH(1): PKR, PAN (Amanah).
  - PH(2): DAP.
  - PAS: As listed.
  - STAR: As listed.
  - PBDSB: As listed.
  - INDEPENDENT 1/2: PSB, GAS parties (ASPIRASI, SEDAR, PBK), independents.


### PHASE-5: Check consistency of column mappings

Analyze whole file by columns; point out inconsistencies
Example: the number of columns MUST always match.  

**Rule**
- AGE + SEX columns are the most likely missing; leave it empty if not filled.  BUt if for candidate some is filled and some not; flag it! 
- Postal vote can be 0 or 1 ONLY an dhave DMKOD of POS. Example: P.221_N.80_221/80/POS_1 <> POSTAL VOTE <> UNDI POS
- Early vote can be 0 or more and have DMKOD of 00. Example: P.221_N.80_221/80/00_1 <> EARLY VOTE <> UNDI AWAL
- VOTING CHANNEL NUMBER column will be used as part of the unique ID first column.  Example: P.221_N.80_221/80/02_2 ==> _2 should match VOTING CHANNEL NUMBER

**Postal Votes**

See below for the pattern for Postal vote rows; which will be first row (if present):

UNIQUE CODE - P.193_N.04_193/04/POS_1,
BALLOT TYPE - POSTAL VOTE
POLLING DISTRICT CODE - 193/04/POS
POLLING DISTRICT NAME - UNDI POS 
POLLING CENTRE - UNDI POS

**Mapping**

Below are some of the important columns (not all) that you will encounter.
It will be important to map data from the source PDF when needed.

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

### PHASE-6 REVIEW: DUN-level Validation of TOTAL BALLOTS ISSUED vs Candidate Totals

#### Objective

Validate that `TOTAL BALLOTS ISSUED` in `to-review.csv` remains correct after the analysis/fix process by cross-checking DUN-level aggregates against official vote totals from `raw-candidate.csv` (grouped by DUN).

---

#### Method

For each DUN (`N.01` to `N.82`), compute:

1. **A (Ballots Issued)**  
   Sum of column `TOTAL BALLOTS ISSUED` from `to-review.csv`.

2. **RAW_JU (Official Candidate Votes)**  
   Sum of `ju` from `raw-candidate.csv`, grouped by DUN (`kid` mapped to `N.XX`).

3. **B (Total Valid Votes)**  
   Sum of `TOTAL VALID VOTES` from `to-review.csv`.

4. **C (Total Rejected Votes)**  
   Sum of `TOTAL REJECTED VOTES` from `to-review.csv`.

5. **D (Total Unreturned Ballots)**  
   Sum of `TOTAL UNRETURNED BALLOTS` from `to-review.csv`.

Then verify:

- **Check-1:** `RAW_JU == B`  
- **Check-2:** `A == B + C + D`  
- Equivalent identity: **`A - RAW_JU == C + D`**


## LESSONS LEARNED — DO NOT REPEAT THESE MISTAKES

### 1. Postal Vote DMKOD format
- **POLLING DISTRICT CODE** (col 8) and the KODDM segment inside **UNIQUE CODE** (col 1) use the suffix `/POS` — e.g. `193/04/POS`.
- **POLLING DISTRICT NAME** (col 9) and **POLLING CENTRE** (col 10) contain the text `UNDI POS`.
- These are different columns. Do **not** put `UNDI POS` into the POLLING DISTRICT CODE or UNIQUE CODE.

### 2. UNIQUE CODE suffix algorithm (duplicate disambiguation)
- Suffixes are assigned **per Polling District Code**, not per channel number.
- Build one `PollingCentre → letter` map for each district (first-appearance order in file).
- Stamp that letter on **every row** for that district+centre — all channels uniformly.
- **Wrong**: treating `_1` duplicates and `_2` duplicates as independent groups produces inconsistent letters (same centre gets `a` for channel 1 but `b` for channel 2, or no letter for unique channels).

## TOOLING

If needed to do more advanced processing, write Golang code using stdlib as much as possible; use the Golang scripts library if needed; use slog for strucgtured logging for debugging.  Do NOT use anything else  (e.g. Python) unless it can be justified; even then should be self-contained (e.g. using uvx for Python) as standalone script

## TRACKING OF REVIEW PROGRESS

Can place summary + phases of review after every step here so the next agent can know where to continue. Store details + evidence in own file PHASE-<NUMBER>-<DESC>.md insude this folder
