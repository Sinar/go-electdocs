# PHASE 5 REVIEW: Column Mapping Consistency

**File reviewed**: `to-review.csv`
**Date**: 2026-04-06
**Methodology**: Automated validation via `phase5_validate.go` (Go stdlib + slog)

---

## Summary

| Metric | Value |
|--------|-------|
| Header columns | 92 (expected 92) ✅ |
| Data rows | 4,317 |
| Parliamentary constituencies | 31 (P.192–P.222) |
| Postal vote rows | 31 (1 per constituency) ✅ |
| Early vote rows | 103 |
| Ordinary vote rows | 4,183 |

---

## Check Results

| # | Check | Violations | Status |
|---|-------|-----------|--------|
| 1 | Column count — every row must have 92 columns | 0 | ✅ PASS |
| 2 | STATE (col 2) must always be `SARAWAK` | 0 | ✅ PASS |
| 3 | BALLOT TYPE (col 3) must be `POSTAL VOTE`, `EARLY VOTE`, or `ORDINARY VOTE` | 0 | ✅ PASS |
| 4 | Postal vote format rules (PDC, PDN, PC, state code, channel) | 0 | ✅ PASS |
| 5 | At most 1 postal vote row per parliamentary constituency | 0 | ✅ PASS |
| 6 | Early vote format rules (PDC ends `/00`, PDN = `UNDI AWAL`) | 0 | ✅ PASS |
| 7 | VOTING CHANNEL NUMBER (col 21) matches UNIQUE CODE suffix | 0 | ✅ PASS |
| 8 | VOTING CHANNEL NUMBER is numeric | 0 | ✅ PASS |
| 9 | Gender/Age consistency per candidate across rows | 0 | ✅ PASS |
| 10 | CHECK ON VALID VOTES (col 91) correctness | 0 | ✅ PASS |
| 11 | CHECK ON TOTAL VOTES ISSUED (col 92) correctness | 0 | ✅ PASS |
| 12 | UNIQUE CODE construction consistency | 0 | ✅ PASS |
| 13 | Comma-formatted numbers in numeric fields | **2** | ❌ FAIL |
| 14 | Ordinary vote rows must not use UNDI POS/UNDI AWAL names or /POS, /00 codes | 0 | ✅ PASS |
| 15 | Party name consistency per constituency per slot | 0 | ✅ PASS |
| 16 | Candidate name consistency per constituency per slot | 0 | ✅ PASS |
| 17 | No empty TOTAL BALLOTS ISSUED | 0 | ✅ PASS |
| 18 | Trailing/leading whitespace in key fields | **493** | ⚠️ WARNING |
| 19 | Demographics columns — mixed fill within a single row | 0 | ✅ PASS |
| 20 | Party slot internal consistency (party ↔ candidate ↔ vote) | 0 | ✅ PASS |
| 21 | PAR CODE format (`P.XXX`) and STATE CODE format (`N.XX`) | 0 | ✅ PASS |

**Total: 21 checks · 19 passed · 2 with issues · 495 total violations/warnings**

---

## Detailed Findings

### Finding 1: Comma-Formatted Numbers in TOTAL VALID VOTES (2 rows) ❌

Two postal vote rows store `TOTAL VALID VOTES` (col 88) as a comma-formatted quoted string instead of a plain integer. The underlying data is **correct** — both rows pass the CHECK formulas — but the formatting breaks naive CSV parsers that don't handle quoted fields.

| Row | UNIQUE CODE | TOTAL VALID VOTES (raw) | Expected | Candidate Votes | Verified |
|-----|-------------|------------------------|----------|----------------|----------|
| 671 | `P.196_P.196/POSTAL VOTE_UNDI POS_1` | `"1,616"` | `1616` | 433 + 1101 + 82 = 1616 | ✅ |
| 3700 | `P.219_P.219/POSTAL VOTE_UNDI POS_1` | `"1,488"` | `1488` | 356 + 1062 + 70 = 1488 | ✅ |

**Verification for P.196 (STAMPIN):**
- TOTAL BALLOTS ISSUED = 1918
- TOTAL VALID VOTES = 1,616 → 1616
- TOTAL REJECTED VOTES = 55
- TOTAL UNRETURNED BALLOTS = 247
- Check: 1918 − 1616 − 55 − 247 = **0** ✅

**Verification for P.219 (MIRI):**
- TOTAL BALLOTS ISSUED = 1770
- TOTAL VALID VOTES = 1,488 → 1488
- TOTAL REJECTED VOTES = 57
- TOTAL UNRETURNED BALLOTS = 225
- Check: 1770 − 1488 − 57 − 225 = **0** ✅

**Root cause**: These are the only two rows where `TOTAL VALID VOTES ≥ 1,000` in a single postal-vote row. The number was likely formatted with a thousands separator before export.

---

### Finding 2: Trailing Whitespace in POLLING DISTRICT NAME (493 rows) ⚠️

493 data rows have a trailing space in column 9 (`POLLING DISTRICT NAME`). No other key column is affected. This is cosmetic but can cause string-matching failures in downstream tools.

**Breakdown by type:**

| Category | Affected Rows | Example |
|----------|--------------|---------|
| Early vote rows (`UNDI AWAL ` with trailing space) | 30 | Row 4: `UNDI AWAL ` → should be `UNDI AWAL` |
| Ordinary vote rows (various district names with trailing space) | 463 | Row 54: `TONDONG ` → should be `TONDONG` |

**Affected constituencies (ordinary vote trailing spaces):**

| PAR Code | Constituency | Affected Rows |
|----------|-------------|--------------|
| P.193 | SANTUBONG | 149 |
| P.208 | SIBU | 111 |
| P.209 | LANANG | 80 |
| P.192 | MAS GADING | 52 |
| P.205 | SARIKEI | 26 |
| P.210 | KANOWIT | 23 |
| P.211 | JULAU | 21 |
| P.207 | MUKAH | 8 |
| P.212 | KAPIT | 7 |
| P.195 | BANDAR KUCHING | 5 |
| P.194 | PETRA JAYA | 5 |
| P.200 | BETONG | 3 |
| P.197 | KOTA SAMARAHAN | 2 |
| P.204 | TANJONG MANIS | 1 |

**Note**: The early vote check (Check #6) passed because the validation script trims whitespace before comparing. The 30 early vote rows with `UNDI AWAL ` (trailing space) are functionally correct but inconsistent with the 73 early vote rows that have `UNDI AWAL` (no trailing space).

---

## Checks That Passed — Key Details

### Column Count (Check 1)
All 4,317 data rows have exactly 92 columns. The Go `encoding/csv` reader correctly handled quoted fields containing commas (e.g., polling centres like `"SURAU DARUL AMIN SAMARIANG JAYA, FASA 2, JALAN BENTARA"`).

### Postal Vote Format (Check 4)
All 31 postal vote rows follow the parliamentary-level pattern correctly:

| Field | Expected Pattern | All Match? |
|-------|-----------------|-----------|
| UNIQUE CODE | `P.XXX_P.XXX/POSTAL VOTE_UNDI POS_1` | ✅ |
| STATE CONSTITUENCY CODE | `P.XXX/POSTAL VOTE` | ✅ |
| STATE CONSTITUENCY NAME | `UNDI POS` | ✅ |
| POLLING DISTRICT CODE | `UNDI POS` | ✅ |
| POLLING DISTRICT NAME | `UNDI POS` | ✅ |
| POLLING CENTRE | `UNDI POS` | ✅ |
| VOTING CHANNEL NUMBER | `1` | ✅ |

### Early Vote Format (Check 6)
All 103 early vote rows have:
- POLLING DISTRICT CODE ending with `/00` ✅
- POLLING DISTRICT NAME = `UNDI AWAL` (after trimming) ✅

### VOTING CHANNEL ↔ UNIQUE CODE (Check 7)
Every row's VOTING CHANNEL NUMBER matches the numeric suffix extracted from the UNIQUE CODE:
- Simple: `P.221_N.80_221/80/02_2` → channel `2` ✅
- Disambiguated: `P.192_N.02_192/02/00_a1` → channel `1` ✅
- Postal: `P.192_P.192/POSTAL VOTE_UNDI POS_1` → channel `1` ✅

### Gender/Age Consistency (Check 9)
For every candidate in every party slot, gender and age are either:
- **Consistently filled** across all rows for that candidate, OR
- **Consistently empty** across all rows

No mixed fill was found. All gender values are either `MALE` or `FEMALE`. All age values are numeric integers.

### CHECK Columns (Checks 10–11)
- **CHECK ON VALID VOTES** (col 91): `sum(candidate votes) − TOTAL VALID VOTES = 0` for all 4,317 rows ✅
- **CHECK ON TOTAL VOTES ISSUED** (col 92): `TOTAL BALLOTS − TOTAL VALID − TOTAL REJECTED − TOTAL UNRETURNED = 0` for all 4,317 rows ✅

(Note: The comma-formatted numbers in rows 671 and 3700 were handled by stripping commas before arithmetic — the underlying values are correct.)

### UNIQUE CODE Construction (Check 12)
Every non-postal row matches the pattern: `{PAR_CODE}_{STATE_CODE}_{POLLING_DISTRICT_CODE}_{[letter]channel}`

Every postal row matches: `{PAR_CODE}_{PAR_CODE}/POSTAL VOTE_UNDI POS_1`

### Party/Candidate Consistency (Checks 15–16)
Within each parliamentary constituency:
- Each party slot uses exactly **one** party name across all rows ✅
- Each party slot uses exactly **one** candidate name across all rows ✅

### Party Slot Internal Consistency (Check 20)
No row has:
- A vote without a candidate name
- A candidate without a party label
- A party label without a candidate name

### Demographics (Check 19)
Demographics columns (11–20) are filled on an all-or-nothing basis within each row:
- **Postal vote rows**: Always empty (31/31) — expected, as postal votes don't have district-level demographics
- **Non-postal rows with demographics**: 103 rows (2 early + 101 ordinary)
- **Non-postal rows without demographics**: 4,183 rows (101 early + 4,082 ordinary)

Demographics are stored as percentage strings (e.g., `42.95%`). The 103 rows with demographics are concentrated in specific DUNs and appear to come from a supplementary data source. No row has a partial mix (some demographic fields filled, others empty).

---

## Recommendations

### 1. Fix Comma-Formatted Numbers (Priority: HIGH)

**Rows 671 and 3700**: Replace `"1,616"` with `1616` and `"1,488"` with `1488` in column 88 (TOTAL VALID VOTES). These quoted comma-formatted numbers will break any CSV parser that doesn't handle quoted fields, and are inconsistent with every other numeric value in the file.

**Fix**: In `to-review.csv`, rows for `P.196` and `P.219` postal votes — change TOTAL VALID VOTES from the comma-formatted to plain integer format.

### 2. Trim Trailing Whitespace in POLLING DISTRICT NAME (Priority: LOW)

493 rows in column 9 have a single trailing space. While functionally harmless (most comparisons trim first), this is inconsistent with the remaining 3,824 rows. A simple trim operation on column 9 would clean this up.

Specifically:
- 30 early vote rows: `UNDI AWAL ` → `UNDI AWAL`
- 463 ordinary vote rows: various names like `TONDONG ` → `TONDONG`, `GROGO ` → `GROGO`, etc.

### 3. No Action Needed on Demographics

The sparse fill of demographics columns (103/4,286 non-postal rows = 2.4%) appears intentional — the data comes from a supplementary source and is only available for certain DUNs. Since it's consistently all-or-nothing per row, this is not a data quality issue.

---

## Conclusion

The file is in **very good shape**. Out of 21 systematic checks across 4,317 rows:

- **19 checks pass completely** with zero violations
- **1 check** found 2 comma-formatted numbers (cosmetic/format issue — underlying data is correct)
- **1 check** found 493 trailing spaces in one column (cosmetic)

No structural, logical, or data integrity issues were found. The UNIQUE CODE construction, ballot type mappings, party/candidate assignments, vote arithmetic (CHECK columns), and channel numbering are all internally consistent.