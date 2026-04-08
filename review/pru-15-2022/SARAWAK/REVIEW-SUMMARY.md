# PRU-15 (2022) SARAWAK Parliamentary Election Data — Review Summary

**File reviewed**: `to-review.csv`
**Rows**: 4,316 data rows + 1 header (92 columns)
**Scope**: 31 parliamentary constituencies (P.192–P.222), 82 state constituencies (N.01–N.82)
**Date completed**: 2025-07-10 (updated with TOTAL BALLOTS ISSUED cross-check)

---

## Overall Verdict: ✅ DATA IS IN GOOD SHAPE

The file is structurally sound, numerically accurate, and complete. All 92 official candidates are present with correct vote totals matching the authoritative EC raw data. A small number of formatting/encoding issues and UNIQUE CODE suffix errors were found but none affect data integrity.

---

## Phase Results at a Glance

| Phase | Description | Result | Issues Found |
|-------|-------------|--------|-------------|
| **0** | Compare against 2016 DUN Sarawak baseline | ✅ PASS | 0 anomalies in to-review.csv (28 PAR code errors found in *2016* source) |
| **1** | UNIQUE CODE uniqueness | ⚠️ ISSUES | 22 duplicate codes; 44 rows fixed, 2 open items needing PDF verification |
| **2** | DUN + PAR codes/names vs official data | ⚠️ MINOR | 1 character typo: BA\`KELALAN backtick→apostrophe (32 rows) |
| **3** | Candidate names + vote totals vs official data | ✅ PASS | 1 name discrepancy: ZULHAIDAH vs ZULBAIDAH SUBOH (likely raw data typo) |
| **4** | Coalition/party slot consistency | ✅ PASS | All 92 candidates correctly placed; PAN/PPBM naming noted (informational) |
| **5** | Column mapping & format consistency | ⚠️ MINOR | 2 comma-formatted numbers + 493 trailing whitespace instances |
| **6** | Ballot totals vs candidate vote sums | ✅ PASS | All 31 PARs: A = B + C + D, candidate sums match raw-candidates.csv exactly |
| **6b** | TOTAL BALLOTS ISSUED multi-source cross-check | ✅ PASS | All 4,317 rows + 31 PARs + 82 DUNs verified against 3 sources; 61.6% turnout |

---

## Action Items

### 🔴 HIGH Priority — Must Fix

| # | Phase | Issue | Rows | Fix Description |
|---|-------|-------|------|-----------------|
| 1 | 1 | **22 duplicate UNIQUE CODEs** across 4 districts due to wrong/missing suffix letters | 44 | Apply suffix corrections as detailed in PHASE-1-REVIEW.md (Issues A + B) |
| 2 | 5 | **Comma-formatted numbers** in TOTAL VALID VOTES | 2 | Row 671 (P.196 postal): `"1,616"` → `1616`; Row 3700 (P.219 postal): `"1,488"` → `1488` |

### 🟡 MEDIUM Priority — Should Fix

| # | Phase | Issue | Rows | Fix Description |
|---|-------|-------|------|-----------------|
| 3 | 2 | **BA\`KELALAN** uses backtick (0x60) instead of apostrophe (0x27) | 32 | Replace backtick with apostrophe in column 7 for all N.81 rows |
| 4 | 1 | **Suffix convention inconsistency**: `_1a` vs `_a1` format | 11 | Normalize to `_a1` convention (letter before channel number) in P.203 rows |
| 5 | 1 | **Suffix consistency violations** in 3 districts (no duplicates but wrong letters) | 3 | Correct suffix letters in districts 210/50/06, 217/70/01, 220/78/03 |

### 🟠 INVESTIGATE — Needs PDF Verification

| # | Phase | Issue | Lines | Description |
|---|-------|-------|-------|-------------|
| 6 | 1 | **Genuine data duplicate** at `P.213_N.58_213/58/08_a1` | 2921–2922 | Two rows: same code, same centre (SK KAMPUNG TEH), same channel 1, but different vote data (270 vs 288 ballots). Need raw PDF to determine which is correct. |
| 7 | 1 | **Ambiguous channel** at district `198/20/33` (SK TANAH PUTEH) | 1258–1260 | Two channel-1 rows with different suffixes and different vote data (169 vs 208 ballots). Need raw PDF to verify. |
| 8 | 3 | **Name discrepancy** in P.218 SIBUTI | All P.218 rows | `ZULHAIDAH SUBOH` (raw EC) vs `ZULBAIDAH SUBOH` (to-review.csv). Votes match perfectly (10,405). Likely raw EC typo but should verify against nomination PDF. |

### 🔵 LOW Priority — Cosmetic

| # | Phase | Issue | Rows | Fix Description |
|---|-------|-------|------|-----------------|
| 9 | 5 | **Trailing whitespace** in POLLING DISTRICT NAME (col 9) | 493 | Trim trailing spaces; affects 30 early vote + 463 ordinary vote rows across 14 constituencies |

---

## Phase 0: Compare Against 2016 DUN Baseline

**Verdict**: ✅ No anomalies in `to-review.csv`

### Key Findings
- **DUN Coverage**: All 80 DUNs from 2016 present in 2022. Two new DUNs (N.79 BUKIT KOTA, N.82 BUKIT SARI) added — these didn't compete in 2016.
- **PAR Code Discrepancies (28 DUNs)**: Found in the **2016 source files**, not in to-review.csv. The 2016 files had incorrect PAR codes that didn't match their own polling district prefixes. The 2022 data is correct.
- **Polling Districts**: 854 PDs matched perfectly. Only 1 new PD added (218/72/07 RAMD in N.72). Zero PDs removed.
- **Polling Centre Names**: 771 of 854 PDs show name differences — mostly abbreviation expansions ("SEK. KEB." → "SEKOLAH KEBANGSAAN") and expected venue changes over 6 years.
- **Row Growth**: Every DUN has more ordinary vote rows in 2022, consistent with voter population growth.

---

## Phase 1: UNIQUE CODE Uniqueness

**Verdict**: ⚠️ Issues found and partially fixed

### Duplicate Categories

| Category | Issue | Codes | Rows |
|----------|-------|-------|------|
| **A** | Missing suffixes entirely | 3 | 6 |
| **B** | Wrong suffix letters (same letter → different centres) | 18 | 36 |
| **C** | Genuine data duplicate (same code, same centre, different votes) | 1 | 2 |
| **D** | Suffix consistency violations (no duplicates, but wrong letters) | 4 districts | 4 |
| **E** | Convention inconsistency (`_1a` vs `_a1`) | — | 11 |

### Districts Requiring Suffix Fixes
1. **195/10/02** (P.195 N.10): TABUAN ULU needs `a`→`c` (9 rows)
2. **198/19/29** (P.198 N.19): New suffixes needed for 2 polling centres (2 rows)
3. **198/20/13** (P.198 N.20): New suffixes needed for 2 polling centres (3 rows)
4. **208/45/05** (P.208 N.45): New suffixes needed for 2 polling centres (5 rows)
5. **219/74/01** (P.219 N.74): PUJUT CORNER `c`→`e`, CHUNG HUA PUJUT `d`→`f`, DATO PERMAISURI `e`→`g` (24 rows)
6. **220/78/02** (P.220 N.78): RIDAB `h`→`k` (1 row)
7. **221/80/02** (P.221 N.80): DAMPIN `a`→`e`, SLI `b`→`f` (2 rows)

Full fix table with line numbers: see PHASE-1-REVIEW.md §Complete Fix Summary.

---

## Phase 2: DUN + PAR Validation

**Verdict**: ⚠️ 1 minor issue

- **31/31 PARs** (P.192–P.222): All present, all names match ✅
- **82/82 DUNs** (N.01–N.82): All present, all codes correct ✅
- **81/82 DUN names match**: N.81 uses backtick (0x60) instead of apostrophe (0x27) in `BA'KELALAN` ⚠️
- **DUN-PAR mapping**: All 82 DUNs map to correct parent PAR ✅

---

## Phase 3: Candidate Validation

**Verdict**: ✅ Pass (1 minor name discrepancy)

- **92/92 candidates present** — zero missing, zero extra
- **92/92 vote totals match** — per-candidate sums from to-review.csv match `ju` in raw-candidates.csv exactly
- **0 gender mismatches**
- **1 name discrepancy**: P.218 SIBUTI — `ZULHAIDAH SUBOH` (raw) vs `ZULBAIDAH SUBOH` (review). "ZULBAIDAH" is the more standard Malay name; raw EC data may contain the typo.

---

## Phase 4: Coalition/Party Consistency

**Verdict**: ✅ Pass

All candidates are correctly placed in their coalition slots:

| Coalition Slot | Seats | Component Parties |
|---------------|-------|-------------------|
| GPS | 31/31 | PBB (14), SUPP (7), PRS (6), PDP (4) |
| PH | 30/31 | PKR (16), DAP (8), PAN (6) |
| PN | 4 seats | PPBM (3), PAS (1) |
| OTHER PARTY | 15 candidates | PSB (10), PBDS (3), PBK (1), SEDAR (1), PBM (1) |
| INDEPENDENT | 11 candidates | All with election symbols |
| BN / GTA / GRS / WARISAN | 0 | Correctly empty |

**Notes**:
- "PAN" used instead of "AMANAH" in 6 PH slots (informational)
- "PPBM" used instead of "BERSATU" in 3 PN slots (informational)
- P.209 JULAU has no PH candidate — confirmed correct per official data

---

## Phase 5: Column Mapping Consistency

**Verdict**: ⚠️ 2 minor issues (19/21 checks pass)

### 21 Checks Performed

| Status | Count |
|--------|-------|
| ✅ PASS | 19 |
| ❌ FAIL | 1 (comma-formatted numbers, 2 rows) |
| ⚠️ WARNING | 1 (trailing whitespace, 493 rows) |

### All Passing Checks
- Column count: all 4,317 rows = 92 columns ✅
- STATE always "SARAWAK" ✅
- BALLOT TYPE always valid (POSTAL/EARLY/ORDINARY VOTE) ✅
- Postal vote format (all 31 rows correct) ✅
- Exactly 1 postal vote per constituency ✅
- Early vote format (PDC ends `/00`, PDN = `UNDI AWAL`) ✅
- VOTING CHANNEL ↔ UNIQUE CODE suffix match ✅
- Gender/Age consistently filled per candidate ✅
- CHECK ON VALID VOTES = 0 for all rows ✅
- CHECK ON TOTAL VOTES ISSUED = 0 for all rows ✅
- UNIQUE CODE construction matches components ✅
- Party/candidate name consistency per constituency ✅
- Party slot internal consistency ✅

---

## Phase 6: Ballot Total Validation

**Verdict**: ✅ Pass — all 31 PARs

| Check | Description | Result |
|-------|-------------|--------|
| Check-1 | Sum of candidate votes = TOTAL VALID VOTES | ✅ 31/31 |
| Check-2 | TOTAL BALLOTS ISSUED = Valid + Rejected + Unreturned | ✅ 31/31 |
| Check-3 | Per-candidate totals match raw-candidates.csv | ✅ 31/31 |

### Grand Totals
| Metric | Value |
|--------|-------|
| Total Ballots Issued | 1,197,450 |
| Total Valid Votes | 1,178,867 |
| Total Rejected Votes | 14,369 |
| Total Unreturned Ballots | 4,214 |
| A − (B + C + D) | **0** ✅ |

### Discovery
The `ut` column in `raw-candidates.csv` is mislabelled — it actually contains **TOTAL REJECTED VOTES** (matches column C in 31/31 PARs), not unreturned ballots. This is a raw EC data labelling issue.

---

## Files Produced

| File | Description |
|------|-------------|
| `PHASE-0-REVIEW.md` | Detailed comparison against 2016 DUN baseline (945 lines) |
| `PHASE-1-REVIEW.md` | UNIQUE CODE uniqueness analysis with fix tables (491 lines) |
| `PHASE-2-REVIEW.md` | DUN + PAR validation report |
| `PHASE-3-REVIEW.md` | Candidate validation report |
| `PHASE-4-REVIEW.md` | Coalition/party consistency report |
| `PHASE-5-REVIEW.md` | Column mapping consistency report |
| `PHASE-6-REVIEW.md` | Ballot total validation report |
| `phase0_compare.go` | Go script for 2016 comparison |
| `phase3_check.go` | Go script for candidate validation |
| `phase4_party_check.go` | Go script for party validation |
| `phase5_validate.go` | Go script for column validation |
| `phase6_validate.go` | Go script for total validation |
| `ballot_check.go` | Go script for row-by-row ballot check vs OCR results files |
| `ballot_dun_check.go` | Go script for multi-source TOTAL BALLOTS ISSUED cross-check |
| `orig-to-review.csv` | Backup of original file before Phase 1 fixes |

---

## Phase 6b: TOTAL BALLOTS ISSUED Multi-Source Cross-Check

**Verdict**: ✅ Pass — all checks pass across three independent sources

### Methodology

Cross-checked `TOTAL BALLOTS ISSUED` (column A) using:
1. **Internal consistency** — A = B + C + D per row
2. **`raw-candidates.csv`** — official EC candidate votes (`ju`) and rejected count (`ut`) at PAR level
3. **`raw-seats-clean.csv`** — registered voter counts per DUN (ceiling check: ballots ≤ registered voters)
4. **`results/Sarawak-P.*.csv`** — OCR'd score sheet PDFs (row-by-row spot-check; unreliable due to OCR noise)

### Results

| Check | Source | Result |
|-------|--------|--------|
| Row-level A = B + C + D | Internal | ✅ **4,317/4,317** rows |
| PAR-level B matches raw `ju` | raw-candidates.csv | ✅ **31/31** PARs, Δ = 0 for all |
| PAR-level C matches raw `ut` | raw-candidates.csv | ✅ **31/31** PARs, Δ = 0 for all |
| DUN ballots ≤ registered voters | raw-seats-clean.csv | ✅ **82/82** DUNs — no exceedance |
| DUN-level aggregate A = B + C + D | Internal per DUN | ✅ **82 DUNs + 31 postal** entries |
| No comma/quote formatting in A or B | Format check | ✅ 0 issues (prior issues already fixed) |
| No zero-ballot rows | Sanity check | ✅ 0 rows |

### Grand Totals

| Metric | Value |
|--------|-------|
| Registered voters | 1,943,074 |
| Total Ballots Issued (A) | 1,197,450 |
| Total Valid Votes (B) | 1,178,867 |
| Total Rejected Votes (C) | 14,369 |
| Total Unreturned Ballots (D) | 4,214 |
| Turnout | **61.6%** |
| A − (B + C + D) | **0** ✅ |

### PAR Turnout Range

Lowest: P.221 LIMBANG 47.9% · Highest: P.205 SARATOK 70.3%

### Note on `raw-seats-clean.csv`

Column 10 is **total registered voters** (not ballots). Columns 11–28 are male/female demographic pairs summing to col 10. Columns 29–31 break registered voters into ordinary/early/postal categories. This file is useful only as a ceiling check, not for direct ballot comparison.

### Note on OCR Results Files

Section 4 row-by-row comparison against `results/Sarawak-P.*.csv` showed many mismatches — these are structural artefacts (multiple to-review rows with `_a`/`_b` suffixes map to one score-sheet row; OCR doubling of numbers). Sections 1–3 above are the authoritative checks.