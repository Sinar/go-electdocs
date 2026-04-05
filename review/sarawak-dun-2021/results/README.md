# README

Store results here in markdown format when we need to check the result tallys

---

# CSV ↔ Markdown (PDF Score Sheet) Mapping Rules

## Overview

This document defines the mapping between two data sources for the **Sarawak DUN KE-12 (2021)** election results:

| Source | Description | Location |
|--------|-------------|----------|
| **CSV** (`to-review.csv`) | Structured tabular data — one row per voting channel (saluran) per polling centre. 67 columns. ~3,757 data rows across all 82 DUN constituencies. | `../to-review.csv` |
| **Markdown** (`Sarawak-N.XX.md`) | OCR extraction from official SPR score sheets (Borang SPR 760 Pin. 1/99 — "Helaian Mata"). One file per DUN constituency. | `./Sarawak-N.01.md` … `./Sarawak-N.82.md` |

The CSV is the **authoritative, complete** dataset. The Markdown files are **lossy OCR extractions** that serve as the ground-truth PDF representation for cross-verification.

---

## 1. Document Structure of the Markdown (Score Sheet)

Each Markdown file is an OCR extraction of a multi-page landscape PDF. The structure is:

### 1.1 Page Header (repeated on every page)

| Element | Example | Notes |
|---------|---------|-------|
| Title | `HELAIAN MATA (SCORE SHEET) BAHAGIAN PILIHAN RAYA NEGERI: N.01 OPAR` | Contains `STATE CONSTITUENCY CODE` + `STATE CONSTITUENCY NAME` |
| Form ID | `BORANG SPR 760 Pin. 1/99` | Administrative form code — not in CSV |
| Total Registered Voters | `JUMLAH PEMILIH: 11,436` | **Not in CSV** — unique to the score sheet |
| Page Number | `MUKASURAT: 1/3` | Indicates pagination of PDF |
| Election Name | `PILIHAN RAYA UMUM DUN SARAWAK KE-12` | Usually in footer |
| Print Date | `Tarikh Cetak: 18/12/2021` | **Not in CSV** |

### 1.2 Data Table Columns

The table has the following column structure (left to right):

| Col # | Malay Header | English Meaning | CSV Equivalent |
|-------|-------------|-----------------|----------------|
| 1 | `Bil` | Row number (sequential per polling district) | No direct equivalent |
| 2 | `No. Kod Daerah Mengundi` | Polling District Code | `POLLING DISTRICT CODE` |
| 3 | `Nama Pusat Mengundi` | Polling Centre Name | `POLLING CENTRE` (often truncated in MD) |
| 4 | `Nombor Tempat Mengundi (saluran)` | Voting Channel Number | `VOTING CHANNEL NUMBER` |
| 5 | `Jumlah kertas undi yang patut berada di dalam peti undi (A)` | Total Ballots Expected in Ballot Box | `TOTAL BALLOTS ISSUED` |
| 6…N | Candidate name columns | Per-candidate vote counts | `[PARTY] VOTE` — **see §3 for order mapping** |
| N+1 | `Jumlah undian oleh pemilih (B)` | Total Valid Votes | `TOTAL VALID VOTES` |
| N+2 | `Bilangan kertas undi yang ditolak (C)` | Rejected Ballots | `TOTAL REJECTED VOTES` |
| N+3 | `Kertas undi dikeluarkan kepada pengundi tetapi tidak dimasukkan ke dalam peti undi (D)` | Unreturned Ballots | `TOTAL UNRETURNED BALLOTS` |

### 1.3 Row Organisation

- **Postal votes** appear first, labelled `POS` or `UNDI POS`
- **Early votes** appear next, labelled `AWAL` or `UNDI AWAL`
- **Ordinary votes** follow, labelled `BIASA` or `UNDI BIASA` (label often only on the first ordinary row)
- Multi-saluran polling centres have sub-rows for saluran 2, 3, etc. under the same `Bil` number
- A **JUMLAH (Grand Totals)** row appears at the bottom of the last page

---

## 2. CSV Column Reference

The CSV has **67 columns** in this order:

| # | Column Name | Description |
|---|-------------|-------------|
| 1 | `UNIQUE CODE` | Synthetic key: `P.XXX_N.XX_XXX/XX/XX_N` (parliament_dun_district_saluran) |
| 2 | `STATE` | Always `SARAWAK` |
| 3 | `BALLOT TYPE` | `POSTAL VOTE` / `EARLY VOTE` / `ORDINARY VOTE` |
| 4 | `PARLIAMENTARY CONSTITUENCY CODE` | e.g. `P.192` |
| 5 | `PARLIAMENTARY CONSTITUENCY NAME` | e.g. `MAS GADING` |
| 6 | `STATE CONSTITUENCY CODE` | e.g. `N.01` → maps to filename `Sarawak-N.01.md` |
| 7 | `STATE CONSTITUENCY NAME` | e.g. `OPAR` |
| 8 | `POLLING DISTRICT CODE` | e.g. `192/01/01` |
| 9 | `POLLING DISTRICT NAME` | e.g. `SEBIRIS` |
| 10 | `POLLING CENTRE` | e.g. `SEKOLAH KEBANGSAAN SEBIRIS` |
| 11 | `VOTING CHANNEL NUMBER` | e.g. `1`, `2`, `3` |
| 12 | `TOTAL BALLOTS ISSUED` | Total ballots for this saluran |
| 13–17 | `GPS` / `GPS CANDIDATE` / `GPS CANDIDATE SEX` / `GPS CANDIDATE AGE` / `GPS VOTE` | GPS coalition candidate block |
| 18–22 | `PH` / ... / `PH VOTE` | PH coalition candidate block |
| 23–27 | `PSB` / ... / `PSB VOTE` | PSB candidate block |
| 28–32 | `PBK` / ... / `PBK VOTE` | PBK candidate block |
| 33–37 | `ASPIRASI` / ... / `ASPIRASI VOTE` | ASPIRASI candidate block |
| 38–42 | `PBDSB` / ... / `PBDSB VOTE` | PBDSB candidate block |
| 43–47 | `SEDAR` / ... / `SEDAR VOTE` | SEDAR candidate block |
| 48–52 | `PAS` / ... / `PAS VOTE` | PAS candidate block |
| 53–57 | `INDEPENDENT 1` / ... / `INDEPENDENT 1 VOTE` | Independent 1 candidate block |
| 58–62 | `INDEPENDENT 2` / ... / `INDEPENDENT 2 VOTE` | Independent 2 candidate block |
| 63 | `TOTAL VALID VOTES` | Sum of all candidate votes |
| 64 | `TOTAL REJECTED VOTES` | Rejected/spoilt ballots |
| 65 | `TOTAL UNRETURNED BALLOTS` | Ballots issued but not placed in box |
| 66 | `CHECKER (VALID VOTE)` | Validation flag (always `1` if data is correct) |
| 67 | `CHECKER (TOTAL VOTE ISSUED)` | Validation flag (always `1` if data is correct) |

Each candidate block has 5 columns: `[COALITION]`, `[COALITION] CANDIDATE`, `[COALITION] CANDIDATE SEX`, `[COALITION] CANDIDATE AGE`, `[COALITION] VOTE`. Only constituencies where that party/coalition contested will have non-empty values.

---

## 3. Critical Mapping Rules

### 3.1 Candidate Column Order is DIFFERENT

> ⚠️ **This is the most important rule.** The candidate columns appear in **different orders** between the two sources.

| Source | Ordering Principle |
|--------|--------------------|
| **CSV** | Fixed coalition grouping: GPS → PH → PSB → PBK → ASPIRASI → PBDSB → SEDAR → PAS → IND1 → IND2 |
| **Markdown** | **Ballot paper order** — varies per constituency, determined by the SPR's candidate numbering |

**To map candidate votes correctly, you MUST match by candidate name, not by column position.**

Example from N.01 OPAR:

| MD Position | MD Candidate | CSV Party | CSV Position |
|:-----------:|-------------|-----------|:------------:|
| 1st | BAYANG AK TERON | SEDAR | 6th |
| 2nd | BILLY ANAK SUJANG | GPS (SUPP) | 1st |
| 3rd | RANUM ANAK MINA | PSB | 3rd |
| 4th | MENENG ANAK BIRIS | PH (PKR) | 2nd |
| 5th | FREEDY ANAK MISID | PBK | 4th |
| 6th | CIKGU SAINI ANAK KAKONG | PBDSB | 5th |

### 3.2 Ballot Type Mapping

| CSV (English) | Markdown (Malay) |
|---------------|-----------------|
| `POSTAL VOTE` | `POS` or `UNDI POS` |
| `EARLY VOTE` | `AWAL` or `UNDI AWAL` |
| `ORDINARY VOTE` | `BIASA` or `UNDI BIASA` |

### 3.3 Row Matching Key

To match a CSV row to a Markdown table row, use this composite key:

```
STATE CONSTITUENCY CODE → selects the .md file (e.g. N.01 → Sarawak-N.01.md)
POLLING DISTRICT NAME   → find the Bil group in the markdown table
VOTING CHANNEL NUMBER   → find the saluran sub-row within that group
```

Alternatively, the `POLLING DISTRICT CODE` (e.g. `192/01/06`) appears in the markdown's second column when the OCR captured it.

### 3.4 Polling Centre Name Matching

Polling centre names are **frequently truncated or garbled** in the markdown. Common patterns:

| Issue | Example |
|-------|---------|
| Name truncated to first word | `SEKOLAH` instead of `SEKOLAH KEBANGSAAN SEBIRIS` |
| Name split across sub-rows | `SEKOLAH KEBANGSAAN` on row 1, `SEBIRIS` on row 2 |
| `(CINA)` qualifier dropped | `SJK CHUNG HUA` instead of `SJK(C) CHUNG HUA` |
| Parenthetical suffix lost | `SEKOLAH MENENGAH KEBANGSAAN BANDAR KUCHING` instead of `...NO.1 (BLOK K - KIRI)` |
| OCR garbling | `SE CHUNG KOLAH J HUA ENIS NO. 5` instead of `SJK(C) CHUNG HUA NO. 5` |

**Matching strategy**: Use fuzzy/substring matching on centre names, and prefer matching by polling district code + saluran number as the primary key.

---

## 4. Fields Present in One Source Only

### 4.1 In CSV Only (not in Markdown)

| Field(s) | Notes |
|----------|-------|
| `UNIQUE CODE` | Synthetic row key |
| `STATE` | Always `SARAWAK` |
| `PARLIAMENTARY CONSTITUENCY CODE` / `NAME` | e.g. `P.192` / `MAS GADING` |
| Party / Coalition names | `SUPP`, `PKR`, `DAP`, `PBB`, `PSB`, `PBK`, etc. |
| Candidate sex | `MALE` / `FEMALE` |
| Candidate age | Numeric age at time of election |
| Election symbol (for independents) | e.g. `RUMAH`, `POKOK` |
| Empty party slots | Structural columns for parties not contesting that seat |
| `CHECKER (VALID VOTE)` | Validation flag (derived) |
| `CHECKER (TOTAL VOTE ISSUED)` | Validation flag (derived) |

### 4.2 In Markdown Only (not in CSV)

| Field | Notes |
|-------|-------|
| `JUMLAH PEMILIH` | Total registered voters for the constituency |
| `BORANG SPR 760 Pin. 1/99` | Form reference number |
| `Tarikh Cetak` | Print date (usually `18/12/2021`) |
| `PILIHAN RAYA UMUM DUN SARAWAK KE-12` | Full election name |
| `Bil` numbering | Sequential row numbers per polling district |
| `JUMLAH` (grand totals row) | Aggregated totals across all saluran — can be cross-checked against CSV sums |
| `MUKASURAT` page numbers | Page X/Y pagination info |

---

## 5. OCR Quality & Known Artifacts

### 5.1 Data Completeness

Based on analysis of N.01, N.10, N.41, and N.82:

| Category | Typical Range | Description |
|----------|:------------:|-------------|
| Full match rows | 44–65% | All numeric values present and correct |
| Partial data rows | 12–15% | Some cells missing/empty but row exists |
| Completely missing rows | 22–29% | No corresponding row found in markdown |

### 5.2 Common OCR Artifacts

| Artifact | Example | Decoded Value | How to Handle |
|----------|---------|:-------------:|---------------|
| **Doubled numbers** | `18187` | `187` | The PDF has overlapping text layers. Strip the duplicated prefix: if the string has odd length, the first `(len-1)/2` characters are a duplicate. |
| **Doubled with space** | `114 114` | `114` | Two copies separated by space — take either one. |
| **Doubled Bil numbers** | `15 15` | `15` | Same deduplication rule applies to Bil and saluran. |
| **Empty table cells** | `\|\|` | Missing | OCR failed to capture the value. Cannot be trusted as zero. |
| **Garbled candidate headers** | `DATO' SRI ABANG TALIF @ LEN BIN ADITAJAYA SALLEH` | Two candidate names merged | Column headers from adjacent candidates merged into one string. |

### 5.3 Page Boundary Data Loss

**The most common cause of missing rows is PDF page boundaries.** Data rows at the top and bottom of each page are frequently lost because:

- Row data overlaps with the repeated column headers on the next page
- OCR merges the last row of one page with the header of the next
- Rows that straddle two pages get dropped entirely

---

## 6. Validation Rules for Agents

These rules can be used by future agents to cross-check the CSV against the Markdown:

### Rule 1: Constituency Identification
```
CSV.STATE_CONSTITUENCY_CODE  ==  heading "N.XX" in Sarawak-N.XX.md
CSV.STATE_CONSTITUENCY_NAME  ==  heading name after "N.XX" in Sarawak-N.XX.md
```

### Rule 2: Vote Total Integrity
```
Sum of all [PARTY] VOTE columns == TOTAL VALID VOTES  (per CSV row)
TOTAL BALLOTS ISSUED >= TOTAL VALID VOTES + TOTAL REJECTED VOTES + TOTAL UNRETURNED BALLOTS
```
(The CHECKER columns in CSV confirm this: both should equal `1`.)

### Rule 3: Candidate Name Verification
```
For each non-empty [PARTY] CANDIDATE in the CSV:
  → The candidate name MUST appear as a column header in the corresponding .md file
  → Match is case-insensitive and should tolerate OCR line-break artifacts
  → The markdown has NO party names — only candidate names
```

### Rule 4: Per-Saluran Vote Cross-Check
```
For a given CSV row:
  1. Open Sarawak-N.{STATE_CONSTITUENCY_CODE number}.md
  2. Find the Bil group matching POLLING_DISTRICT_NAME (or POLLING_DISTRICT_CODE)
  3. Find the sub-row matching VOTING_CHANNEL_NUMBER
  4. For each candidate, match by name to find the correct MD column
  5. Compare: CSV.[PARTY]_VOTE == MD.candidate_column_value (after OCR deduplication)
  6. Compare: CSV.TOTAL_VALID_VOTES == MD.column_(B)
  7. Compare: CSV.TOTAL_REJECTED_VOTES == MD.column_(C)
  8. Compare: CSV.TOTAL_UNRETURNED_BALLOTS == MD.column_(D)
  9. Compare: CSV.TOTAL_BALLOTS_ISSUED == MD.column_(A)
```

### Rule 5: Polling Centre Name Verification
```
CSV.POLLING_CENTRE should be a substring-match or fuzzy-match of the
Markdown's "Nama Pusat Mengundi" field (which may be split across sub-rows).
Expect:
  - Truncation (only first 1-2 words present in MD)
  - Missing qualifiers like "(CINA)", "(MIS)", block identifiers
  - OCR garbling of multi-word names
Use POLLING_DISTRICT_CODE + VOTING_CHANNEL_NUMBER as the reliable join key.
```

### Rule 6: Grand Totals Cross-Check
```
The JUMLAH row at the bottom of the .md file contains aggregated totals.
These should equal the SUM of the corresponding CSV column across all rows
for that constituency:
  - MD.JUMLAH.(A)  == SUM(CSV.TOTAL_BALLOTS_ISSUED)
  - MD.JUMLAH.(B)  == SUM(CSV.TOTAL_VALID_VOTES)
  - MD.JUMLAH.(C)  == SUM(CSV.TOTAL_REJECTED_VOTES)
  - MD.JUMLAH.(D)  == SUM(CSV.TOTAL_UNRETURNED_BALLOTS)
  - MD.JUMLAH.candidate_col == SUM(CSV.[PARTY]_VOTE) for that candidate
Note: The JUMLAH row is often heavily garbled by OCR.
```

### Rule 7: Registered Voters Sanity Check
```
MD.JUMLAH_PEMILIH (total registered voters) >= SUM(CSV.TOTAL_BALLOTS_ISSUED)
This is a sanity check — more ballots than registered voters indicates an error.
```

---

## 7. File Naming Convention

| Pattern | Example | Description |
|---------|---------|-------------|
| `Sarawak-N.{XX}.md` | `Sarawak-N.01.md` | Markdown extraction for DUN constituency N.XX |
| `INDEX.csv` | — | Lookup table: `dun_code,dun_name,output_file` |
| `../to-review.csv` | — | The authoritative CSV with all election data |

The `{XX}` in the filename corresponds to `STATE CONSTITUENCY CODE` column value `N.{XX}` in the CSV. Zero-padded to 2 digits.

---

## 8. Quick Reference: Mapping Cheat Sheet

```
CSV Row                              Markdown Location
─────────────────────────────────    ──────────────────────────────────
STATE CONSTITUENCY CODE (N.01)   →   File: Sarawak-N.01.md
                                     Heading: "N.01 OPAR"

BALLOT TYPE = "POSTAL VOTE"      →   Row labelled "POS" / "UNDI POS"
BALLOT TYPE = "EARLY VOTE"       →   Row labelled "AWAL" / "UNDI AWAL"
BALLOT TYPE = "ORDINARY VOTE"    →   Rows after AWAL, labelled "BIASA"

POLLING DISTRICT CODE            →   Column 2 ("No. Kod Daerah Mengundi")
POLLING DISTRICT NAME            →   Text label in Bil group
POLLING CENTRE                   →   Column 3 (often truncated)
VOTING CHANNEL NUMBER            →   Column 4 ("saluran")

TOTAL BALLOTS ISSUED             →   Column (A)
[PARTY] VOTE                     →   Candidate column matched BY NAME
TOTAL VALID VOTES                →   Column (B)
TOTAL REJECTED VOTES             →   Column (C)
TOTAL UNRETURNED BALLOTS         →   Column (D)

GPS CANDIDATE / PH CANDIDATE … →   Column headers (names only, no party)
CANDIDATE SEX / AGE              →   NOT in markdown
Party / Coalition names          →   NOT in markdown
CHECKER columns                  →   NOT in markdown
JUMLAH PEMILIH (reg. voters)     →   In markdown header, NOT in CSV
JUMLAH (grand totals)            →   Last row of markdown, NOT in CSV
```
