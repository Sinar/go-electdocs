# PHASE 3 REVIEW: Find Missing OR Incorrect Candidates

## Objective

Compare all candidates in `to-review.csv` against the authoritative reference in `raw-candidate.csv`, checking:
1. Are all 349 reference candidates present in `to-review.csv`?
2. Are there extra candidates in `to-review.csv` not in the reference?
3. Are candidate names spelled correctly?
4. Is each candidate placed in the correct party column?
5. Are sex fields consistent between reference and review?

## Methodology

- **Tool**: `phase3_check.go` (Go, stdlib only, `encoding/csv` + `slog`)
- **Reference**: `raw-candidate.csv` (349 candidates across 82 DUNs)
- **Review**: `to-review.csv` (3,748 data rows, 67 columns; candidates extracted from first row per DUN)
- **Supporting files**: `raw-party.csv` (party ID → abbreviation), `raw-dun.csv` (DUN ID → DUN code)

### Mapping Logic

| raw-candidate.csv `pid` | Party | to-review.csv Column Group | Expected Label(s) |
|---|---|---|---|
| 51 | GPS | GPS (cols 13–17) | PBB, SUPP, PRS, PDP |
| 37 | PKR | PH (cols 18–22) | PKR |
| 81 | AMANAH | PH (cols 18–22) | PAN |
| 3 | DAP | PH (cols 18–22) | DAP |
| 54 | PSB | PSB (cols 23–27) | PSB |
| 95 | PBK | PBK (cols 28–32) | PBK |
| 47 | ASPIRASI | ASPIRASI (cols 33–37) | ASPIRASI |
| 49 | PBDSB | PBDSB (cols 38–42) | PBDSB |
| 70 | SEDAR | SEDAR (cols 43–47) | SEDAR |
| 2 | PAS | PAS (cols 48–52) | PAS |
| 20 | BEBAS | INDEPENDENT 1/2 (cols 53–62) | Ballot symbol (varies) |

### Matching Strategy

1. **Normalization**: Trim whitespace, uppercase, collapse multiple spaces, replace backtick (`` ` ``) with apostrophe (`'`)
2. **Phase A**: Exact normalized name match
3. **Phase B**: Fuzzy match (bigram Jaccard similarity ≥ 0.70) for any unmatched candidates
4. **Consistency check**: Verify all rows within a DUN have the same candidates (detect intra-DUN inconsistencies)

## Results

### ✅ PASS — All candidates matched with zero issues

```
==========================================================
  PHASE 3: CANDIDATE COMPARISON REPORT
==========================================================

  Reference DUNs:          82
  Review DUNs:             82
  Reference candidates:    349
  Review candidates:       349
  Total issues found:      0

  Total clean DUNs: 82 / 82
```

### Summary Statistics

| Metric | Value |
|---|---|
| Reference DUNs | 82 |
| Review DUNs | 82 |
| Reference candidates | 349 |
| Review candidates (unique per DUN) | 349 |
| Missing in review | 0 |
| Extra in review | 0 |
| Name mismatches (fuzzy) | 0 |
| Name spelling differences | 0 |
| Wrong column assignments | 0 |
| Party label issues | 0 |
| Sex mismatches | 0 |
| Intra-DUN inconsistencies | 0 |

### Candidate Distribution by Party

| Party (pid) | Abbreviation | Column | Candidates |
|---|---|---|---|
| 51 | GPS | GPS | 82 |
| 95 | PBK | PBK | 73 |
| 54 | PSB | PSB | 70 |
| 20 | BEBAS | INDEPENDENT 1/2 | 30 |
| 37 | PKR | PH | 28 |
| 3 | DAP | PH | 26 |
| 47 | ASPIRASI | ASPIRASI | 15 |
| 49 | PBDSB | PBDSB | 11 |
| 81 | AMANAH | PH | 8 |
| 70 | SEDAR | SEDAR | 5 |
| 2 | PAS | PAS | 1 |
| **Total** | | | **349** |

### GPS Component Party Labels in to-review.csv

All 82 GPS candidates (pid=51) are correctly placed in the GPS column. The party label sub-column shows the specific GPS component party:

| GPS Component | Count |
|---|---|
| PBB | Most seats |
| SUPP | Urban/Chinese-majority seats |
| PRS | Dayak-majority seats |
| PDP | Select seats (e.g., N.02 Tasik Biru, N.52 Dudong) |

### PH Component Party Labels in to-review.csv

All 62 PH candidates (PKR=28, DAP=26, AMANAH=8) are correctly placed in the PH column with accurate sub-labels:

| PH Component | pid | Label in to-review.csv | Count |
|---|---|---|---|
| PKR | 37 | PKR | 28 |
| DAP | 3 | DAP | 26 |
| AMANAH | 81 | PAN | 8 |

> **Note**: AMANAH (Parti Amanah Negara) uses the abbreviation "PAN" in `to-review.csv`, which is an accepted alternative abbreviation for the party. This is consistent across all 8 AMANAH candidates.

### Independent Candidates

All 30 independent candidates (pid=20, BEBAS) are correctly placed in INDEPENDENT 1 or INDEPENDENT 2 columns. The party label column shows the candidate's ballot symbol (e.g., KAPAL TERBANG, RUMAH, KUNCI, UDANG, JAM, KERUSI, KUDA, GAJAH, POKOK, etc.), which is the expected format for independents.

4 DUNs have 2 independent candidates each (using both INDEPENDENT 1 and INDEPENDENT 2):
- N.41 Kuala Rajang
- N.47 Pakan
- N.52 Dudong
- N.60 Kakus

### Edge Cases Verified

| Case | Candidate | DUN | Result |
|---|---|---|---|
| Trailing whitespace in raw data | `ROBERTSON MAWA ` (trailing space) | N.62 Katibas | ✅ Matched after trim |
| Apostrophe in name | `DATO' IDRIS BUANG` | N.16 Muara Tuang | ✅ Exact match |
| Backtick in name | `DATO` SRI ABANG ADITAJAYA` | N.41 Kuala Rajang | ✅ Matched (backtick in both files) |
| Single-word name | `AGNES` | N.81 Ba'Kelalan | ✅ Exact match |
| Single-word name | `CARDOCK` | N.52 Dudong | ✅ Exact match |
| 8 candidates (max) | N.52 Dudong | N.52 | ✅ All 8 matched |
| 2 candidates (min) | N.04 Pantai Damai, N.36 Layar, N.56 Dalat, N.59 Tamin | Multiple | ✅ All matched |
| Only PAS candidate | `ARIF PAIJO` (pid=2) | N.29 Beting Maro | ✅ Correct PAS column |

### Sex Field Validation

The reference file uses `L` (Lelaki/Male) and `P` (Perempuan/Female). The review file uses `MALE` and `FEMALE`. All 349 mappings are consistent — zero sex mismatches detected.

### Intra-DUN Consistency

All rows within each DUN have identical candidate names in corresponding columns. Zero inconsistencies detected across 3,748 data rows.

## Conclusion

**PHASE 3: ✅ PASS**

All 349 candidates from the authoritative `raw-candidate.csv` are present in `to-review.csv` with:
- **Correct names** (zero spelling differences after normalization)
- **Correct party column assignments** (GPS, PH, PSB, PBK, ASPIRASI, PBDSB, SEDAR, PAS, INDEPENDENT 1/2)
- **Correct party sub-labels** (PBB/SUPP/PRS/PDP for GPS; PKR/DAP/PAN for PH; ballot symbols for independents)
- **Correct sex fields** (L→MALE, P→FEMALE)
- **No extra or missing candidates**
- **Full intra-DUN consistency** across all 3,748 rows

No corrective action required.