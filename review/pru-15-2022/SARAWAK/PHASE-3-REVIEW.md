# PHASE 3 REVIEW: Find Missing or Incorrect Candidates

## Overview

Comparison of candidates in `to-review.csv` against official data in `raw-candidates.csv` for PRU-15 (2022) Sarawak parliamentary constituencies (P.192–P.222).

**Method**: A Go script (`phase3_check.go`) was used to:
1. Parse `raw-party-data-clean.csv` to build party ID → abbreviation mapping
2. Extract all Sarawak parlimen candidates from `raw-candidates.csv` (kid 19200–22200, kt="parlimen")
3. Extract all candidates per parliamentary constituency from `to-review.csv` across all 13 party/independent slots
4. Compare candidate names (exact + fuzzy matching), gender, and vote totals
5. Identify missing, extra, or mismatched candidates

---

## Summary Statistics

| Metric | Count |
|--------|-------|
| Parliamentary constituencies checked | 31 (P.192–P.222) |
| Total official candidates (raw-candidates.csv) | 92 |
| Total candidates found in to-review.csv | 92 |
| Candidates matched (exact name) | 91 |
| Candidates matched (fuzzy/similar name) | 1 |
| Candidates MISSING from to-review.csv | **0** |
| EXTRA candidates in to-review.csv (not in raw) | **0** |
| Vote total mismatches | **0** |
| Gender mismatches | **0** |

**Overall verdict: ✅ All 92 official candidates are accounted for in `to-review.csv`. Vote totals match perfectly. One name discrepancy found (see Section 3).**

---

## 1. Candidate Counts per Parliamentary Constituency

All 31 constituencies have the correct number of candidates:

| PAR Code | PAR Name | Raw (Official) | Review (to-review) | Match? |
|----------|----------|---------------|--------------------|---------| 
| P.192 | MAS GADING | 3 | 3 | ✅ |
| P.193 | SANTUBONG | 3 | 3 | ✅ |
| P.194 | PETRA JAYA | 3 | 3 | ✅ |
| P.195 | BANDAR KUCHING | 3 | 3 | ✅ |
| P.196 | STAMPIN | 3 | 3 | ✅ |
| P.197 | KOTA SAMARAHAN | 2 | 2 | ✅ |
| P.198 | PUNCAK BORNEO | 3 | 3 | ✅ |
| P.199 | SERIAN | 4 | 4 | ✅ |
| P.200 | BATANG SADONG | 2 | 2 | ✅ |
| P.201 | BATANG LUPAR | 3 | 3 | ✅ |
| P.202 | SRI AMAN | 4 | 4 | ✅ |
| P.203 | LUBOK ANTU | 4 | 4 | ✅ |
| P.204 | BETONG | 3 | 3 | ✅ |
| P.205 | SARATOK | 3 | 3 | ✅ |
| P.206 | TANJONG MANIS | 2 | 2 | ✅ |
| P.207 | IGAN | 2 | 2 | ✅ |
| P.208 | SARIKEI | 2 | 2 | ✅ |
| P.209 | JULAU | 4 | 4 | ✅ |
| P.210 | KANOWIT | 5 | 5 | ✅ |
| P.211 | LANANG | 4 | 4 | ✅ |
| P.212 | SIBU | 3 | 3 | ✅ |
| P.213 | MUKAH | 2 | 2 | ✅ |
| P.214 | SELANGAU | 3 | 3 | ✅ |
| P.215 | KAPIT | 3 | 3 | ✅ |
| P.216 | HULU RAJANG | 2 | 2 | ✅ |
| P.217 | BINTULU | 3 | 3 | ✅ |
| P.218 | SIBUTI | 3 | 3 | ✅ |
| P.219 | MIRI | 3 | 3 | ✅ |
| P.220 | BARAM | 3 | 3 | ✅ |
| P.221 | LIMBANG | 2 | 2 | ✅ |
| P.222 | LAWAS | 3 | 3 | ✅ |

---

## 2. Name Mismatches (Possible Typos)

Only **one** name discrepancy was found across all 92 candidates:

| PAR | Raw (Official) Name | to-review.csv Name | Raw Party | Review Party | Review Slot |
|-----|---------------------|-------------------|-----------|--------------|-------------|
| P.218 (SIBUTI) | **ZULHAIDAH** SUBOH | **ZULBAIDAH** SUBOH | PH (pid=31) | PKR | PH |

### Analysis of the discrepancy

- **Raw data**: `ZULHAIDAH SUBOH` (from `raw-candidates.csv`, id=1210, pid=31)
- **to-review.csv**: `ZULBAIDAH SUBOH` (in PKR/PH slot)
- **Difference**: Characters at position 4–5 differ — `HAID` vs `BAID`
- **"ZULBAIDAH"** is a well-known Malay feminine name (variant of Zubaidah/Zulbaidah). "ZULHAIDAH" appears to be less standard.
- **Verdict**: The name `ZULBAIDAH` in `to-review.csv` is likely correct; the raw EC data may contain a typo (`H` instead of `B`). However, since the raw EC data is the authoritative source, this discrepancy should be flagged. **Recommend verifying against the original PDF nomination form.**
- **Vote totals match**: raw `ju`=10405, review sum=10405 ✅

---

## 3. Gender Comparison

**No gender mismatches found.** All candidates where gender data is available in both sources match correctly.

Mapping used: `L` (raw) = `MALE` (review), `P` (raw) = `FEMALE` (review).

---

## 4. Candidates Missing from to-review.csv

**None.** All 92 official candidates from `raw-candidates.csv` are present in `to-review.csv`.

---

## 5. Extra Candidates in to-review.csv

**None.** There are no candidates in `to-review.csv` that are absent from `raw-candidates.csv`.

---

## 6. Vote Total Comparison

**All 92 candidates have matching vote totals.** The sum of per-row votes in `to-review.csv` for each candidate matches the `ju` field in `raw-candidates.csv` exactly.

This confirms that:
- No votes are missing or duplicated across polling station rows
- Candidate-to-row attribution is correct throughout the file

---

## 7. Party Mapping Analysis

### Coalition vs Component Party Names

`raw-candidates.csv` uses **coalition-level** party IDs, while `to-review.csv` uses **component party** names. This is by design and is correct — the review file provides more granular party information.

#### Raw party IDs used in Sarawak parlimen constituencies:

| Party ID (pid) | Abbreviation | Full Name | Count |
|----------------|-------------|-----------|-------|
| 32 | GPS | Gabungan Parti Sarawak | 31 |
| 31 | PH | Pakatan Harapan | 22 |
| 20 | BEBAS | Independent | 11 |
| 35 | PSB | Parti Sarawak Bersatu | 10 |
| 3 | DAP | Democratic Action Party | 8 |
| 27 | PN | Perikatan Nasional | 4 |
| 30 | PBDS | Parti Bansa Dayak Sarawak | 3 |
| 58 | PBK | Party Bumi Kenyalang | 1 |
| 42 | PBM | Parti Bangsa Malaysia | 1 |
| 38 | SEDAR | Parti Sedar Rakyat Sarawak | 1 |

#### How to-review.csv maps these to party slots:

| Raw Coalition (pid) | to-review.csv Party Name | to-review.csv Slot | Notes |
|---------------------|--------------------------|-------------------|-------|
| GPS (32) | PBB | GPS | Parti Pesaka Bumiputra Bersatu |
| GPS (32) | PRS | GPS | Parti Rakyat Sarawak |
| GPS (32) | PDP | GPS | Parti Demokratik Progresif |
| GPS (32) | SUPP | GPS | Sarawak United Peoples' Party |
| PH (31) | PKR | PH | Parti Keadilan Rakyat |
| PH (31) | PAN | PH | Parti Amanah Negara |
| DAP (3) | DAP | PH | Democratic Action Party (own pid=3, not under PH pid=31) |
| PN (27) | PAS | PN | Parti Islam Se-Malaysia |
| PN (27) | PPBM | PN | Parti Pribumi Bersatu Malaysia |
| PSB (35) | PSB | OTHER PARTY (1) | Parti Sarawak Bersatu |
| PBDS (30) | PBDS | OTHER PARTY (1) | Parti Bansa Dayak Sarawak |
| PBK (58) | PBK | OTHER PARTY (1) | Party Bumi Kenyalang |
| PBM (42) | PBM | OTHER PARTY (2) | Parti Bangsa Malaysia |
| SEDAR (38) | SEDAR | OTHER PARTY (1) | Parti Sedar Rakyat Sarawak |
| BEBAS (20) | INDEPENDENT - {SYMBOL} | INDEPENDENT 1/2/3 | Symbol-based identifier |

#### Independent candidate symbols:

| PAR | Candidate | Symbol in to-review.csv | Slot |
|-----|-----------|------------------------|------|
| P.193 | AFFENDI BIN JEMAN | INDEPENDENT - KEY | INDEPENDENT 1 |
| P.199 | DR ALIM IMPIRA | INDEPENDENT - KEY | INDEPENDENT 1 |
| P.202 | MASIR ANAK KUJAT | INDEPENDENT - TREE | INDEPENDENT 1 |
| P.204 | HASBIE BIN SATAR | INDEPENDENT - SAMPAN | INDEPENDENT 1 |
| P.209 | ELLY LAWAI NGALAI | INDEPENDENT - AEROPLANE | INDEPENDENT 1 |
| P.210 | DR. ELLI LUHAT | INDEPENDENT - KEY | INDEPENDENT 1 |
| P.210 | GEORGE CHEN | INDEPENDENT - BOOK | INDEPENDENT 2 |
| P.210 | MICHAEL ANAK LIAS | INDEPENDENT - CHAIR | INDEPENDENT 3 |
| P.211 | DATO WONG TIING KIONG | INDEPENDENT - DONKEY | INDEPENDENT 1 |
| P.214 | HENRY JOSEPH USAU | INDEPENDENT - AEROPLANE | INDEPENDENT 1 |
| P.220 | WILFREDENTIKA | INDEPENDENT - TREE | INDEPENDENT 1 |

All 11 BEBAS (independent) candidates from `raw-candidates.csv` (pid=20) are correctly placed in INDEPENDENT slots in `to-review.csv`, each with their unique election symbol identifier. Where multiple independents contest the same seat (P.210 has 3), they are spread across INDEPENDENT 1/2/3 slots.

---

## 8. Key Observations

### ✅ Strengths
1. **100% candidate coverage** — All 92 official candidates are present
2. **100% vote total accuracy** — Every candidate's summed votes match official figures exactly
3. **Correct gender data** — No gender mismatches found
4. **Correct slot placement** — Coalition candidates are correctly placed in their respective slots (GPS components in GPS slot, PH components in PH slot, etc.)

### ⚠️ Issues Found
1. **One name discrepancy** (P.218 SIBUTI): `ZULHAIDAH SUBOH` (raw) vs `ZULBAIDAH SUBOH` (to-review). Likely a typo in the raw EC data. Recommend verification against nomination form PDF.

### ℹ️ Notes for Future Phases
- The `to-review.csv` uses component party names (PBB, PKR, DAP, PAS, etc.) rather than coalition names (GPS, PH, PN). This is more informative but differs from `raw-candidates.csv` which uses coalition-level party IDs.
- Independent candidates are labelled with their election symbols (KEY, TREE, SAMPAN, etc.) — this is useful for disambiguation where multiple independents contest the same seat.
- **DAP has its own pid=3** in `raw-candidates.csv`, separate from the PH coalition pid=31. All 8 DAP candidates (pid=3) are correctly placed in the PH slot in `to-review.csv` with party name "DAP". The remaining 22 PH candidates (pid=31) appear under component names PKR or PAN. This means `raw-candidates.csv` treats DAP as a standalone party while PKR/PAN are grouped under PH — a quirk of the EC data model, but `to-review.csv` handles it correctly by placing all of them in the PH slot.

---

## Conclusion

**Phase 3 passes with one minor finding.** All 92 candidates are present, correctly placed in their party slots, and have accurate vote totals. The single name discrepancy (`ZULHAIDAH` vs `ZULBAIDAH` SUBOH in P.218) should be verified against the original nomination PDF but is likely a raw data typo.