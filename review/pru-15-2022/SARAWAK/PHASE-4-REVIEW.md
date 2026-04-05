# Phase 4 Review: Coalition/Party Assignment Consistency

## Objective

Verify that party/coalition assignments in `to-review.csv` are consistent with official data from `raw-candidates.csv` and `raw-party-data-clean.csv` for PRU-15 (2022 Parliamentary Election) Sarawak seats (P.192–P.222, 31 constituencies).

---

## 1. Party Labels Found in `to-review.csv`

### Per-Slot Party Labels

**BN slot (col 23):**

No BN candidates found in any Sarawak parliamentary seat. This is **expected** — BN did not field candidates in Sarawak for PRU-15; GPS contested as the incumbent coalition instead.

**PH slot (col 28):**

| Party Label | Constituencies |
|-------------|----------------|
| DAP | P.192, P.195, P.196, P.199, P.208, P.211, P.212, P.217 |
| PAN | P.193, P.197, P.200, P.201, P.206, P.207 |
| PKR | P.194, P.198, P.202, P.203, P.204, P.205, P.210, P.213, P.214, P.215, P.216, P.218, P.219, P.220, P.221, P.222 |

> **Note on "PAN"**: The label "PAN" is used for Parti Amanah Negara (official abbreviation: AMANAH, pid=45 in raw-party-data-clean.csv). "PAN" is a common informal abbreviation. In raw-candidates.csv, these candidates have pid=31 (PH coalition-level), so the specific component is not distinguishable from raw data alone.

**PN slot (col 33):**

| Party Label | Constituencies |
|-------------|----------------|
| PAS | P.201 |
| PPBM | P.203, P.205, P.217 |

> **Note on "PPBM"**: This is Parti Pribumi Bersatu Malaysia (official abbreviation per raw-party-data-clean.csv: BERSATU, pid=55). "PPBM" is the original registered abbreviation. In raw-candidates.csv, all four PN candidates have pid=27 (PN coalition-level).

**GTA slot (col 38):**

No GTA candidates found in any Sarawak seat. This is expected — Gerakan Tanah Air did not contest in Sarawak.

**GPS slot (col 43):**

| Party Label | Constituencies |
|-------------|----------------|
| PBB | P.193, P.194, P.197, P.198, P.200, P.201, P.204, P.206, P.207, P.213, P.215, P.218, P.221, P.222 (14 seats) |
| PDP | P.192, P.205, P.217, P.220 (4 seats) |
| PRS | P.202, P.203, P.209, P.210, P.214, P.216 (6 seats) |
| SUPP | P.195, P.196, P.199, P.208, P.211, P.212, P.219 (7 seats) |

> GPS fielded candidates in all 31 seats. All are labelled with the correct GPS component party. Total: 14 PBB + 4 PDP + 6 PRS + 7 SUPP = **31 seats** ✅

**GRS slot (col 48):**

No GRS candidates. Expected — GRS is a Sabah coalition and did not contest Sarawak parliamentary seats.

**WARISAN slot (col 53):**

No WARISAN candidates. Expected — Warisan did not contest Sarawak parliamentary seats in PRU-15.

**OTHER PARTY (1) slot (col 58):**

| Party Label | Constituencies |
|-------------|----------------|
| PSB | P.195, P.196, P.198, P.199, P.202, P.203, P.211, P.212, P.219, P.222 (10 seats) |
| PBDS | P.209, P.215, P.218 (3 seats) |
| PBK | P.192 (1 seat) |
| SEDAR | P.194 (1 seat) |

**OTHER PARTY (2) slot (col 63):**

| Party Label | Constituencies |
|-------------|----------------|
| PBM | P.209 (1 seat) |

**OTHER PARTY (3) slot (col 68):**

No entries. ✅

**INDEPENDENT 1 slot (col 73):**

| Party Label | Constituencies |
|-------------|----------------|
| INDEPENDENT - KEY | P.193, P.199, P.210 |
| INDEPENDENT - TREE | P.202, P.220 |
| INDEPENDENT - SAMPAN | P.204 |
| INDEPENDENT - AEROPLANE | P.209, P.214 |
| INDEPENDENT - DONKEY | P.211 |

**INDEPENDENT 2 slot (col 78):**

| Party Label | Constituencies |
|-------------|----------------|
| INDEPENDENT - BOOK | P.210 |

**INDEPENDENT 3 slot (col 83):**

| Party Label | Constituencies |
|-------------|----------------|
| INDEPENDENT - CHAIR | P.210 |

> **Note on Independent labels**: Independent candidates in Malaysian elections are assigned a symbol (key, tree, sampan, etc.) instead of a party logo. These symbols are correctly reflected in the independent labels.

---

## 2. Official Candidates vs to-review.csv Mapping

### Method

1. Built party ID → abbreviation map from `raw-party-data-clean.csv`
2. Extracted all Sarawak parlimen candidates from `raw-candidates.csv` (kid 19200–22200, kt="parlimen")
3. Determined expected slot by mapping party IDs to coalitions
4. Matched candidates by name to slots in `to-review.csv`
5. Compared expected vs actual slot placement

### Detailed Per-Constituency Mapping

| PAR | Candidate (Official) | Official Party (pid) | Expected Slot | Actual Slot | Actual Label | Status |
|-----|---------------------|---------------------|---------------|-------------|--------------|--------|
| P.192 | RYAN SIM MIN LEONG | PBK (pid=58) | OTHER PARTY | OTHER PARTY (1) | PBK | ✅ |
| P.192 | LIDANG DISEN | GPS (pid=32) | GPS | GPS | PDP | ✅ |
| P.192 | MORDI ANAK BIMOL | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.193 | NANCY SHUKRI | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.193 | AFFENDI BIN JEMAN | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - KEY | ✅ |
| P.193 | MOHAMAD ZEN PELI | PH (pid=31) | PH | PH | PAN | ✅ |
| P.194 | FADILLAH BIN YUSOF | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.194 | SOPIAN JULAIHI | PH (pid=31) | PH | PH | PKR | ✅ |
| P.194 | OTHMAN BIN ABDILLAH | SEDAR (pid=38) | OTHER PARTY | OTHER PARTY (1) | SEDAR | ✅ |
| P.195 | TAY TZE KOK | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.195 | VOON LEE SHAN | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.195 | KELVIN YII LEE WUEN | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.196 | CHONG CHIENG JEN | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.196 | LO KHERE CHIANG | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.196 | LUE CHENG HING | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.197 | DATUK HJH RUBIAH BINTI WANG | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.197 | ABANG HALIL | PH (pid=31) | PH | PH | PAN | ✅ |
| P.198 | WILLIE ANAK MONGIN | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.198 | DIOG ANAK DIOS | PH (pid=31) | PH | PH | PKR | ✅ |
| P.198 | IANA ANAK AKAM | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.199 | DR ALIM IMPIRA | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - KEY | ✅ |
| P.199 | DATO' SRI RICHARD RIOT ANAK JAEM | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.199 | ELSIY A/K TINGANG | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.199 | LEARRY ANAK JABUL | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.200 | RODIYAH BINTI SAPIEE | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.200 | CIKGU LAHAJI | PH (pid=31) | PH | PH | PAN | ✅ |
| P.201 | MOHAMAD SHAFIZAN | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.201 | WEL @ MAXWEL ROJIS | PH (pid=31) | PH | PH | PAN | ✅ |
| P.201 | HAMDAN SANI | PN (pid=27) | PN | PN | PAS | ✅ |
| P.202 | DORIS SOPHIA ANAK BRODI | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.202 | WILSON ANAK ENTABANG | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.202 | MASIR ANAK KUJAT | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - TREE | ✅ |
| P.202 | NAGA LIBAU @ TAY | PH (pid=31) | PH | PH | PKR | ✅ |
| P.203 | JUGAH MUYANG | PN (pid=27) | PN | PN | PPBM | ✅ |
| P.203 | JOHNICHAL RAYONG NGIPA | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.203 | ROY ANGAU ANAK GINGKOI | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.203 | LANGGA LIAS | PH (pid=31) | PH | PH | PKR | ✅ |
| P.204 | PATRICK KAMIS @ HJ KAMENG | PH (pid=31) | PH | PH | PKR | ✅ |
| P.204 | HASBIE BIN SATAR | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - SAMPAN | ✅ |
| P.204 | DR. RICHARD RAPU | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.205 | GIENDAM JONATHAN TAIT | GPS (pid=32) | GPS | GPS | PDP | ✅ |
| P.205 | DATUK ALI BIJU | PN (pid=27) | PN | PN | PPBM | ✅ |
| P.205 | IBIL JAYA | PH (pid=31) | PH | PH | PKR | ✅ |
| P.206 | IR. YUSUF ABD WAHAB | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.206 | USTAZAH ZAINAB | PH (pid=31) | PH | PH | PAN | ✅ |
| P.207 | AHMAD JOHNIE BIN ZAWAWI | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.207 | HUD ANDRI | PH (pid=31) | PH | PH | PAN | ✅ |
| P.208 | HUANG TIONG SII | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.208 | RODERICK WONG SIEW LEAD | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.209 | SUSAN ANAK GEORGE | PBDS (pid=30) | OTHER PARTY | OTHER PARTY (1) | PBDS | ✅ |
| P.209 | JOSEPH SALANG | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.209 | LARRY SOON @ LARRY SNG WEI SHIEN | PBM (pid=42) | OTHER PARTY | OTHER PARTY (2) | PBM | ✅ |
| P.209 | ELLY LAWAI NGALAI | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - AEROPLANE | ✅ |
| P.210 | JOSEPH NYAMBONG | PH (pid=31) | PH | PH | PKR | ✅ |
| P.210 | DR. ELLI LUHAT | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - KEY | ✅ |
| P.210 | AARON AGO ANAK DAGANG | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.210 | GEORGE CHEN | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 2 | INDEPENDENT - BOOK | ✅ |
| P.210 | MICHAEL ANAK LIAS | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 3 | INDEPENDENT - CHAIR | ✅ |
| P.211 | DATO WONG TIING KIONG | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - DONKEY | ✅ |
| P.211 | PRISCILLA LAU | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.211 | ALICE LAU KIONG YIENG | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.211 | WONG CHING YONG | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.212 | OSCAR LING CHAI YEW | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.212 | CLARENCE TING ING HORH | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.212 | WONG SOON KOH | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.213 | HANIFAH HAJAR TAIB | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.213 | ABDUL JALIL | PH (pid=31) | PH | PH | PKR | ✅ |
| P.214 | UMPANG ANAK SABANG | PH (pid=31) | PH | PH | PKR | ✅ |
| P.214 | EDWIN ANAK BANTA | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.214 | HENRY JOSEPH USAU | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - AEROPLANE | ✅ |
| P.215 | ALEXANDER NANTA LINGGI | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.215 | PANGKAS AK. UNGGANG | PH (pid=31) | PH | PH | PKR | ✅ |
| P.215 | ROBERT SAWENG | PBDS (pid=30) | OTHER PARTY | OTHER PARTY (1) | PBDS | ✅ |
| P.216 | ABUN SUI ANYIT | PH (pid=31) | PH | PH | PKR | ✅ |
| P.216 | UGAK ANAK KUMBONG | GPS (pid=32) | GPS | GPS | PRS | ✅ |
| P.217 | DUKE ANAK JANTENG | PN (pid=27) | PN | PN | PPBM | ✅ |
| P.217 | DATO SRI TIONG KING SING | GPS (pid=32) | GPS | GPS | PDP | ✅ |
| P.217 | CHIEW CHAN YEW | DAP (pid=3) | PH | PH | DAP | ✅ |
| P.218 | LUKANISMAN BIN AWANG SAUNI | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.218 | ZULHAIDAH SUBOH | PH (pid=31) | PH | PH | PKR | ⚠️ Name mismatch |
| P.218 | BOBBY ANAK WILLIAM | PBDS (pid=30) | OTHER PARTY | OTHER PARTY (1) | PBDS | ✅ |
| P.219 | JEFFERY PHANG | GPS (pid=32) | GPS | GPS | SUPP | ✅ |
| P.219 | LAWRENCE LAI | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.219 | CHIEW CHOON MAN | PH (pid=31) | PH | PH | PKR | ✅ |
| P.220 | ANYI NGAU | GPS (pid=32) | GPS | GPS | PDP | ✅ |
| P.220 | ROLAND ENGAN | PH (pid=31) | PH | PH | PKR | ✅ |
| P.220 | WILFREDENTIKA | BEBAS (pid=20) | INDEPENDENT | INDEPENDENT 1 | INDEPENDENT - TREE | ✅ |
| P.221 | RACHA BALANG | PH (pid=31) | PH | PH | PKR | ✅ |
| P.221 | HASBI BIN HABIBOLLAH | GPS (pid=32) | GPS | GPS | PBB | ✅ |
| P.222 | HJ JAPAR BIN HJ SUYUT | PH (pid=31) | PH | PH | PKR | ✅ |
| P.222 | BARU BIAN | PSB (pid=35) | OTHER PARTY | OTHER PARTY (1) | PSB | ✅ |
| P.222 | HENRY SUM AGONG | GPS (pid=32) | GPS | GPS | PBB | ✅ |

**Result: 92 of 92 candidate slot placements are correct. 1 name spelling discrepancy found.**

---

## 3. Coalition Mapping Validation

### Expected Coalition → Slot Mapping (PRU-15 2022)

| Coalition/Slot | Component Parties | Notes |
|----------------|-------------------|-------|
| GPS | PBB, SUPP, PRS, PDP | Gabungan Parti Sarawak — incumbent Sarawak coalition |
| PH | DAP, PKR, AMANAH/PAN | Pakatan Harapan |
| PN | PAS, BERSATU/PPBM | Perikatan Nasional |
| BN | UMNO, MCA, MIC | Barisan Nasional — **did not contest in Sarawak** |
| GRS | — | Gabungan Rakyat Sabah — **did not contest in Sarawak** |
| WARISAN | — | Parti Warisan Sabah — **did not contest in Sarawak** |
| OTHER PARTY | PSB, PBK, SEDAR, PBDS, PBM | Minor Sarawak-based parties |
| INDEPENDENT | BEBAS | Independent candidates (assigned election symbols) |

### Label-to-Slot Consistency Check

✅ **All party labels appear in their correct coalition slots.** Specifically:

- All GPS component labels (PBB, SUPP, PRS, PDP) appear **only** in the GPS slot
- All PH component labels (DAP, PKR, PAN) appear **only** in the PH slot
- All PN component labels (PAS, PPBM) appear **only** in the PN slot
- All minor party labels (PSB, PBDS, PBK, SEDAR, PBM) appear **only** in OTHER PARTY slots
- All independent labels appear **only** in INDEPENDENT slots
- BN, GTA, GRS, WARISAN slots are correctly empty for Sarawak

---

## 4. Issues Found

### 4.1 Candidate Name Spelling Discrepancy

| PAR | PAR Name | Official Name (raw-candidates.csv) | Name in to-review.csv | Party | Slot |
|-----|----------|-----------------------------------|-----------------------|-------|------|
| P.218 | SIBUTI | **ZULHAIDAH** SUBOH | **ZULBAIDAH** SUBOH | PKR | PH |

**Evidence:**
- `raw-candidates.csv`: `ZULHAIDAH SUBOH` (pid=31, kid=21800)
- `to-review.csv`: `ZULBAIDAH SUBOH` (PH slot, label=PKR)
- Difference: **ZULHAIDAH** vs **ZULBAIDAH** (letter H vs B at position 4)
- The raw-candidates.csv data is from the official Election Commission, which is the authoritative source
- **Severity**: Minor typo; the candidate is correctly placed in the PH/PKR slot

### 4.2 Naming Convention Observations (Informational — Not Errors)

These are not errors but are worth documenting for consistency:

| Convention in to-review.csv | Official/Alternate | Notes |
|----------------------------|-------------------|-------|
| **PAN** | AMANAH (pid=45) | "PAN" is a common informal abbreviation for Parti Amanah Negara. Raw-candidates.csv uses pid=31 (PH coalition-level) for these candidates, so the component assignment cannot be verified from raw data alone. PAN is used consistently across all 6 AMANAH constituencies. |
| **PPBM** | BERSATU (pid=55) | "PPBM" is the party's registered abbreviation (Parti Pribumi Bersatu Malaysia). `raw-party-data-clean.csv` uses "BERSATU" (pid=55). Raw-candidates.csv uses pid=27 (PN coalition-level). PPBM is used consistently for all 3 BERSATU constituencies. |
| **PAS** (P.201 PN slot) | PAS (pid=2) | Hamdan Sani is labelled PAS under the PN slot. Raw-candidates.csv gives pid=27 (PN). Wikipedia's Batang Lupar page confirms the candidate as simply "PN". However, PAS contested this same seat (P.201) in GE-14 2018 under its own flag. The PAS label is plausible but cannot be confirmed solely from raw EC data. |

> **Summary**: The component party labels within coalition slots (PAN vs AMANAH, PPBM vs BERSATU, PAS within PN) are used **consistently** within to-review.csv and are reasonable assignments. However, since `raw-candidates.csv` uses coalition-level IDs (pid=31 for PH, pid=27 for PN), the specific component party assignments must have been sourced from supplementary information.

---

## 5. Party IDs Used in Sarawak Parlimen (raw-candidates.csv)

| Party ID | Abbreviation | Full Name | Expected Slot | Candidates |
|----------|-------------|-----------|---------------|------------|
| 3 | DAP | Democratic Action Party | PH | 8 |
| 20 | BEBAS | Independent | INDEPENDENT | 11 |
| 27 | PN | Perikatan Nasional | PN | 4 |
| 30 | PBDS | Parti Bansa Dayak Sarawak | OTHER PARTY | 3 |
| 31 | PH | Pakatan Harapan | PH | 22 |
| 32 | GPS | Gabungan Parti Sarawak | GPS | 31 |
| 35 | PSB | Parti Sarawak Bersatu | OTHER PARTY | 10 |
| 38 | SEDAR | Parti Sedar Rakyat Sarawak | OTHER PARTY | 1 |
| 42 | PBM | Parti Bangsa Malaysia | OTHER PARTY | 1 |
| 58 | PBK | Party Bumi Kenyalang | OTHER PARTY | 1 |

**Total candidates: 92** (across 31 constituencies)

> Note: `raw-candidates.csv` uses **coalition-level IDs** for GPS (32), PH (31), and PN (27) in most cases. DAP (3) is the exception — 8 DAP candidates are recorded with the component party ID rather than the PH coalition ID. All other PH candidates use pid=31.

---

## 6. Constituency Coverage Summary

| Metric | Count |
|--------|-------|
| Total Sarawak parliamentary seats | 31 (P.192–P.222) |
| Constituencies with all assignments correct | **30** |
| Constituencies with minor issues | **1** (P.218 — name spelling only) |
| Total candidates verified | **92** |
| Slot placement mismatches | **0** |
| Name spelling discrepancies | **1** |

### Constituencies with no issues (30/31):

P.192, P.193, P.194, P.195, P.196, P.197, P.198, P.199, P.200, P.201, P.202, P.203, P.204, P.205, P.206, P.207, P.208, P.209, P.210, P.211, P.212, P.213, P.214, P.215, P.216, P.217, P.219, P.220, P.221, P.222

---

## 7. Special Observations

### 7.1 GPS Component Party Distribution

GPS contested all 31 Sarawak parliamentary seats with component parties distributed as follows:

| GPS Component | Seats | Percentage |
|---------------|-------|------------|
| PBB | 14 | 45.2% |
| SUPP | 7 | 22.6% |
| PRS | 6 | 19.4% |
| PDP | 4 | 12.9% |
| **Total** | **31** | **100%** |

This distribution is consistent with the known GPS internal seat allocation where PBB holds the largest share.

### 7.2 PH Component Party Distribution

PH contested 30 of 31 seats (all except P.209 Julau):

| PH Component | Seats | Percentage |
|--------------|-------|------------|
| PKR | 16 | 53.3% |
| DAP | 8 | 26.7% |
| PAN/AMANAH | 6 | 20.0% |
| **Total** | **30** | **100%** |

### 7.3 Independent Candidate Symbols

| Symbol | Count | Constituencies |
|--------|-------|----------------|
| KEY | 3 | P.193, P.199, P.210 |
| TREE | 2 | P.202, P.220 |
| AEROPLANE | 2 | P.209, P.214 |
| SAMPAN | 1 | P.204 |
| DONKEY | 1 | P.211 |
| BOOK | 1 | P.210 |
| CHAIR | 1 | P.210 |
| **Total** | **11** | |

P.210 (Kanowit) had the most independent candidates with 3 (KEY, BOOK, CHAIR).

### 7.4 Parties Not Contesting in Sarawak PRU-15

The following coalition/party slots in to-review.csv are correctly empty across all 31 seats:
- **BN** — Did not field direct candidates (GPS is Sarawak's BN-equivalent)
- **GTA** — Gerakan Tanah Air, Peninsular Malaysia focus
- **GRS** — Gabungan Rakyat Sabah, Sabah-only coalition
- **WARISAN** — Parti Warisan Sabah, Sabah-only party

---

## 8. Recommendations

### Must Fix

1. **P.218 SIBUTI — Candidate name spelling**: Change `ZULBAIDAH SUBOH` → `ZULHAIDAH SUBOH` in the PH CANDIDATE column (col 29) to match the official EC data in `raw-candidates.csv`.

### No Action Required

2. **All 92 candidates are in the correct coalition/party slots** — no slot reassignment needed.
3. **All party labels are consistent** within their respective coalition slots — no label corrections needed.
4. **BN/GTA/GRS/WARISAN slots correctly empty** — accurately reflects that these coalitions did not contest in Sarawak.

### For Future Reference

5. **PAN vs AMANAH**: Consider standardizing to "AMANAH" to match `raw-party-data-clean.csv` (pid=45, a="AMANAH"). Currently "PAN" is used consistently but differs from the official abbreviation.
6. **PPBM vs BERSATU**: Consider standardizing to "BERSATU" to match `raw-party-data-clean.csv` (pid=55, a="BERSATU"). Currently "PPBM" is used consistently but differs from the official abbreviation.
7. **Component party verification**: Since `raw-candidates.csv` uses coalition-level IDs (pid=31 for PH, pid=27 for PN), the specific component party labels (PAN, PKR, DAP within PH; PAS, PPBM within PN) cannot be verified from raw EC data alone. An external authoritative source (e.g., official nomination papers) would be needed for full verification.

---

## 9. Conclusion

**Phase 4 review finds the coalition/party assignments in `to-review.csv` are highly consistent and accurate.** Out of 92 candidates across 31 constituencies:

- ✅ **0 slot placement errors** — every candidate is in the correct coalition slot
- ✅ **All party labels are valid** and correctly mapped to their coalition
- ⚠️ **1 minor name spelling issue** in P.218 (ZULBAIDAH → ZULHAIDAH)
- ℹ️ **Naming conventions** (PAN/AMANAH, PPBM/BERSATU) are internally consistent but differ from `raw-party-data-clean.csv` abbreviations

**Overall assessment: PASS** — with 1 minor correction needed.