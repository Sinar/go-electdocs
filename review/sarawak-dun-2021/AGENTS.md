# AGENTS

## OBJECTIVE

To thoroughly review the to-review.csv and see if any of the data is not matching, has typo, does not add up and is placed under wrong location.  Any analysis should be backed by strong evidence.

The raw data is from the official Election Commission website so can be assumed to be the authoritative version. These files are prefixed with raw-*.csv

If needed, we can breakdown the original PDFs to Markdown and use those data to reconfirm anything suspicous. Ask for help if this step is needed

## REVIEW PROCESS

We will start the review of the file to-review.csv in phases from easiest to hardest. The number of columns should be the same and aligned logically; you will NEVER find one column having candidate name and another with the DMKOD 

Do each phase one by one; at the end of the phase write out the review of the phase in the file PHASE-<NUMBER>-REVIEW.md

### PHASE-0: Check against PRN DUN Sarawak 2016 results

The files in the folder /Users/leow/TINDAKMSIA/go-electdocs/data/sarawak-dun-2016/OUTPUT with the pattern Sarawak-N.<DUN>.csv should be compared against to-review.csv.  Go compare DUN by DUN (there should be 81) so easy to extract; note DUN N.82 in 2016 did not compete so no data.

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


### PHASE-2: Find Missing OR Incorrect DUN

Check against the list in raw-dun.csv. Flag out incorrect DUN name and if not all DUNs are matched.  There should not be any unaccounted for.

### PHASE-3: Find Missing OR Incorrect Candidates

Check against the list in raw-candidates.csv.  Flag out incorrect name and if not all candidates are matched.  There should not be any unaccounted for.


### PHASE-4: Check consistency of coallition

In the Sarawak state assembly is made out of many small parties, a few major co-oaltion + possibly few independents.  Point out if anything not found in the official data.
 
Use the raw-party.csv file; analyze the header to extract the semantic meaning and map as per necessary

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
- Reference implementation: `fix_unique_codes.go`.

## TOOLING

If needed to do more advanced processing, write Golang code using stdlib as much as possible; use the Golang scripts library if needed; use slog for strucgtured logging for debugging.  Do NOT use anything else  (e.g. Python) unless it can be justified; even then should be self-contained (e.g. using uvx for Python) as standalone script

## TRACKING OF REVIEW PROGRESS

Can place summary + phases of review after every step here so the next agent can know where to continue. Store details + evidence in own file PHASE-<NUMBER>-<DESC>.md insude this folder

### PHASE-0: Check against PRN DUN Sarawak 2016 results — ✅ COMPLETED
- **Tool**: `phase0_check.go`
- **Result**: Compared columns 1–11 of `to-review.csv` against 80 individual 2016 `Sarawak-N.XX.csv` files (N.79 and N.82 had no 2016 files — both won uncontested).
- **Matched rows** (by UNIQUE CODE): 1,090 out of 3,748 (2021) / ~3,127 (2016)
- **Column differences found**: 1,037 (all in matched rows)
  - 1,012 in POLLING CENTRE (col 10): abbreviation expansions (`SEK. KEB.` → `SEKOLAH KEBANGSAAN`, etc.) — expected normalization
  - 23 in STATE CONSTITUENCY NAME (col 7): BA`KELALAN backtick issue (N.81 only) — already flagged in Phase-2
  - 2 in POLLING DISTRICT NAME (col 9): BA`KELALAN backtick issue (N.81 only)
- **Unmatched rows**: 2,658 only-in-2021 + 2,037 only-in-2016, explained by:
  - **RC-1**: 27 DUNs (N.11–N.37, N.58) have wrong P-codes in 2016 UNIQUE CODEs — a known 2016 dataset error
  - **RC-2**: Multiple DUNs have suffixed IDs in 2016 (`_1a`, `_1b`) vs duplicate IDs in 2021 — tracked in Phase-1
  - **RC-3**: All 82 DUNs had wrong `/UNDI POS` in POLLING DISTRICT CODE and UNIQUE CODE — fixed to `/POS` (correct per rule). `UNDI POS` belongs only in POLLING DISTRICT NAME and POLLING CENTRE columns.
- **Columns with ZERO differences** (among matched rows): STATE, BALLOT TYPE, PARLIAMENTARY CONSTITUENCY CODE/NAME, STATE CONSTITUENCY CODE, POLLING DISTRICT CODE, VOTING CHANNEL NUMBER
- **Actionable issues for 2021 data**:
  1. BA`KELALAN backtick → apostrophe (⚠️ Medium, already in Phase-2)
  2. ~~N.01 postal format~~ — `/POS` in POLLING DISTRICT CODE is **correct** for ALL DUNs. All 82 postal rows were fixed from `/UNDI POS` → `/POS`.
- **Details**: See `PHASE-0-REVIEW.md`

### PHASE-1: Ensure ID field is unique — ❌ FAIL (duplicates found)
- **Tool**: `phase1_check.go`
- **Result**: 523 duplicate UNIQUE CODE values found across 1865 rows (out of 3748 total data rows). No empty UNIQUE CODEs. No existing suffixed IDs.
- **Root cause**: Multiple Polling Centres share the same Polling District Code + Voting Channel Number, so the current ID scheme (`ParliCode_DUNCode_DMKOD_Channel`) does not disambiguate them.
- **Distribution**: Most duplicates appear 2× (267 codes), but some go up to 18× (e.g. `P.215_N.61_215/61/04_1`, `P.216_N.66_216/66/02_1`).
- **Fix needed**: For each Polling District Code, assign one letter per distinct Polling Centre (first-appearance order across the whole file). Apply that letter to **ALL rows** for that district+centre combination — every channel number, not just channels that were duplicates. Only column 1 (UNIQUE CODE) is modified.
- **⚠️ WRONG approach**: Treating each channel (`_1`, `_2`, …) as an independent duplicate group causes the same Polling Centre to receive different letters across channels (e.g. `_1a` but `_2` or `_2b` for the same centre). See `fix_unique_codes.go` for the correct district-level implementation.
- **Scope**: Systematic issue affecting all 82 DUNs.
- **Details**: See `PHASE-1-REVIEW.md`

### PHASE-2: Find Missing OR Incorrect DUN — ✅ COMPLETED
- **Tool**: `phase2_check.go`
- **Result**: All 82 DUNs (N.01–N.82) present in `to-review.csv`; no missing, no extra.
- **Issue found**: 1 name mismatch — N.81 uses backtick (`` ` ``, 0x60) instead of apostrophe (`'`, 0x27): `BA`KELALAN` should be `BA'KELALAN`
- **Details**: See `PHASE-2-REVIEW.md`

### PHASE-3: Find Missing OR Incorrect Candidates — ✅ PASS (all candidates matched)
- **Tool**: `phase3_check.go`
- **Result**: All 349 candidates from `raw-candidate.csv` matched against `to-review.csv` with **0 issues** across all 82 DUNs.
- **Checks performed**:
  1. Candidate name matching (normalized: trim, uppercase, collapse spaces, backtick→apostrophe) ✅
  2. Fuzzy matching fallback (bigram Jaccard ≥ 0.70) for unmatched names — none needed ✅
  3. Party column assignment (pid→column group: GPS, PH, PSB, PBK, ASPIRASI, PBDSB, SEDAR, PAS, INDEPENDENT) ✅
  4. Party sub-label validation (GPS→PBB/SUPP/PRS/PDP; PH→PKR/DAP/PAN; independents→ballot symbols) ✅
  5. Sex field consistency (L→MALE, P→FEMALE) ✅
  6. Intra-DUN candidate consistency across all 3,748 rows ✅
- **Candidate distribution**: GPS=82, PBK=73, PSB=70, BEBAS=30, PKR=28, DAP=26, ASPIRASI=15, PBDSB=11, AMANAH=8, SEDAR=5, PAS=1
- **Notable observations** (not issues):
  - AMANAH (pid=81) uses abbreviation "PAN" in to-review.csv — consistent across all 8 candidates
  - Independent candidates use ballot symbols as party labels (KAPAL TERBANG, RUMAH, KUNCI, UDANG, JAM, KERUSI, KUDA, GAJAH, POKOK, etc.)
  - 4 DUNs have 2 independent candidates (N.41, N.47, N.52, N.60); correctly placed in INDEPENDENT 1 + INDEPENDENT 2
  - N.52 Dudong has maximum 8 candidates; N.04/N.36/N.56/N.59 have minimum 2 candidates
- **Details**: See `PHASE-3-REVIEW.md`

### PHASE-4: Check consistency of coalition — ✅ PASS (0 issues found)
- **Tool**: `phase4_check.go`
- **Result**: All 349 candidates' party/coalition assignments verified with **0 issues** across all 82 DUNs.
- **Checks performed**:
  1. Party label validation — all labels in columns 13/18/23/28/33/38/43/48/53/58 match expected values for their column slot ✅
  2. Candidate-to-column cross-check — all 349 official candidates found in correct party column ✅
  3. PH component label verification — 62/62 PH candidates have correct component label (PKR/DAP/PAN) matching official party ID ✅
  4. Intra-DUN consistency — party labels and candidate names identical across all rows per DUN ✅
  5. Candidate count per DUN matches between raw-candidate.csv and to-review.csv ✅
  6. No extra or missing candidates in either dataset ✅
- **GPS component distribution**: PBB=47, SUPP=18, PRS=11, PDP=6 (total=82, all seats contested)
- **PH component distribution**: PKR=28, DAP=26, PAN(AMANAH)=8; 20 DUNs with no PH candidate
- **Other parties**: PSB=70, PBK=73, ASPIRASI=15, PBDSB=11, SEDAR=5, PAS=1
- **Independents**: 30 BEBAS candidates across 26 DUNs (4 DUNs with 2 independents: N.41, N.47, N.52, N.60); all use election symbols as party labels (KUNCI, RUMAH, KAPAL TERBANG, POKOK, KERUSI, etc.)
- **Notable observations** (not issues):
  - AMANAH (pid=81) consistently labelled "PAN" in to-review.csv — correct official Malay abbreviation
  - GPS component party (PBB/SUPP/PRS/PDP) cannot be cross-checked against raw-candidate.csv (all GPS candidates have pid=51), but all labels are valid GPS component names per raw-party.csv
  - STAR (pid=15) listed in AGENTS.md mapping rules but fielded 0 candidates in 2021 — correctly absent from to-review.csv
  - PDP label used instead of SPDP — reflects 2016 party name change
- **Details**: See `PHASE-4-REVIEW.md`

### PHASE-5: Check consistency of column mappings — ✅ PASS (all rules satisfied)
- **Tool**: `phase5_check.go`
- **Result**: All 10 rules pass with **0 violations** across 3,748 data rows (67 columns each).
- **Rules checked**:
  1. Column count = 67 for every row ✅
  2. STATE = "SARAWAK" for every row ✅
  3. BALLOT TYPE ∈ {POSTAL VOTE, EARLY VOTE, ORDINARY VOTE} ✅
  4. Postal vote rules (DMKOD suffix, POLLING DISTRICT NAME, POLLING CENTRE, UNIQUE CODE contain POS, ≤1 per DUN) ✅
  5. Early vote rules (DMKOD suffix /00, POLLING DISTRICT NAME = UNDI AWAL, UNIQUE CODE contains /00_) ✅
  6. VOTING CHANNEL NUMBER matches last segment of UNIQUE CODE ✅
  7. SEX/AGE consistency per DUN/candidate (no mixed filled/empty for same candidate) ✅, all SEX values MALE/FEMALE/empty ✅, all AGE values numeric/empty ✅
  8. All expected numeric columns (votes, totals) contain valid integers or empty ✅
  9. CHECKER (VALID VOTE) correct: sum of candidate votes = TOTAL VALID VOTES for all rows ✅; CHECKER (TOTAL VOTE ISSUED) correct: TOTAL BALLOTS ISSUED = Valid+Rejected+Unreturned for all rows ✅
  10. UNIQUE CODE structure matches `P.XXX_N.YY_XXX/YY/ZZ_C` pattern, internal P/DUN codes consistent with columns 4 and 6 ✅; POLLING DISTRICT CODE in UC matches column 8 ✅
- **Notable observations** (not violations):
  1. ~~N.01 postal anomaly~~ — `/POS` in POLLING DISTRICT CODE is correct for ALL 82 DUNs. `UNDI POS` belongs only in POLLING DISTRICT NAME and POLLING CENTRE. All 82 postal rows corrected (`fix_postal_dmkod.go`).
  2. 9 records had multi-line polling centre names (embedded newlines in quoted CSV fields) — cleaned up (`fix_newlines.go`).
  3. GPS column populated for all 3,748 rows (all 82 seats contested); sub-labels: PBB, PDP, PRS, SUPP
  4. Independent candidates use election symbols as party labels (BUKU, CANGKUL, JAM, KAPAL TERBANG, etc.)
- **Ballot type distribution**: 82 POSTAL VOTE, 111 EARLY VOTE, 3,555 ORDINARY VOTE
- **Details**: See `PHASE-5-REVIEW.md`
