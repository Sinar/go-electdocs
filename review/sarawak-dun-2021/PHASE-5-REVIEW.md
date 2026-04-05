# PHASE 5 REVIEW: Column Consistency and Mapping Rules

## Executive Summary

Total data rows analyzed: **3748**

| Rule | Description | Result | Violations |
| --- | --- | --- | --- |
| 1 | Column count (expected 67) | ✅ PASS | 0 |
| 2 | STATE = SARAWAK | ✅ PASS | 0 |
| 3 | BALLOT TYPE valid values | ✅ PASS | 0 |
| 4 | Postal vote rules | ✅ PASS | 0 issues + 0 multi-postal DUNs |
| 5 | Early vote rules | ✅ PASS | 0 |
| 6 | VOTING CHANNEL NUMBER consistency | ✅ PASS | 0 |
| 7 | SEX/AGE consistency | ✅ PASS | 0 inconsistencies, 0 invalid SEX, 0 invalid AGE |
| 8 | Numeric columns | ✅ PASS | 0 |
| 9 | CHECKER columns | ✅ PASS | 0 |
| 10 | UNIQUE CODE structure | ✅ PASS | 0 |
| — | PDC in UNIQUE CODE vs col 8 | ✅ PASS | 0 |

### ✅ Overall: PASS — All column consistency rules satisfied

---

## Rule 1: Column Count

✅ **PASS**: All rows have exactly 67 columns.

---

## Rule 2: STATE Column

✅ **PASS**: All rows have STATE = 'SARAWAK'.

---

## Rule 3: BALLOT TYPE

✅ **PASS**: All rows have valid BALLOT TYPE values.

---

## Rule 4: Postal Vote Rules

Total postal vote rows: **82**

✅ **PASS**: All postal vote rows follow the expected rules.

### Postal Vote DMKOD Format Distribution

| Format | Count | DUNs |
| --- | --- | --- |
| `POS` | 1 | N.01 |
| `UNDI POS` | 81 | N.02, N.03, N.04, N.05, N.06 ... (81 total) |

**Note**: N.01 (OPAR) uses the older `/POS` format while all other DUNs use `/UNDI POS`. This was already identified in Phase-0.

---

## Rule 5: Early Vote Rules

✅ **PASS**: All early vote rows follow the expected rules.

---

## Rule 6: VOTING CHANNEL NUMBER Consistency

✅ **PASS**: VOTING CHANNEL NUMBER matches UNIQUE CODE suffix for all rows.

---

## Rule 7: SEX and AGE Column Consistency

✅ **PASS**: All SEX/AGE values are consistent within each DUN/candidate.

---

## Rule 8: Numeric Columns

✅ **PASS**: All expected numeric columns contain valid integers or are empty.

---

## Rule 9: CHECKER Columns

✅ **PASS**: All CHECKER columns are correct.

### CHECKER (VALID VOTE) — Sum of candidate votes vs TOTAL VALID VOTES

✅ Sum of candidate votes equals TOTAL VALID VOTES for all rows.

### CHECKER (TOTAL VOTE ISSUED) — TOTAL BALLOTS ISSUED vs Valid+Rejected+Unreturned

✅ TOTAL BALLOTS ISSUED = Valid + Rejected + Unreturned for all rows.

---

## Rule 10: UNIQUE CODE Structure

Expected pattern: `P.XXX_N.YY_XXX/YY/ZZ_C`

✅ **PASS**: All UNIQUE CODEs match the expected pattern and are internally consistent.

---

## Additional: POLLING DISTRICT CODE in UNIQUE CODE vs Column 8

✅ **PASS**: POLLING DISTRICT CODE embedded in UNIQUE CODE matches column 8 for all rows.

---

## Appendix: Data Distribution

### Ballot Type Distribution

| Ballot Type | Count |
| --- | --- |
| POSTAL VOTE | 82 |
| EARLY VOTE | 111 |
| ORDINARY VOTE | 3555 |

### Party Column Usage

Shows how many rows have a non-empty candidate name for each party group.

| Party Group | Rows with Candidate | Unique Candidates | Party Labels Used |
| --- | --- | --- | --- |
| GPS | 3748 | 82 | PBB, PDP, PRS, SUPP |
| PH | 2950 | 62 | DAP, PAN, PKR |
| PSB | 3239 | 70 | PSB |
| PBK | 3368 | 73 | PBK |
| ASPIRASI | 884 | 15 | ASPIRASI |
| PBDSB | 567 | 11 | PBDSB |
| SEDAR | 221 | 5 | SEDAR |
| PAS | 30 | 1 | PAS |
| INDEPENDENT 1 | 1165 | 26 | BUKU, CANGKUL, JAM, KAPAL TERBANG, KERUSI, KUDA, KUNCI, PEN, PERAHU, POKOK, RUMAH, UDANG |
| INDEPENDENT 2 | 195 | 4 | GAJAH, KERUSI, KUNCI, POKOK |

