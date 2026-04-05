# PHASE-2 REVIEW: Find Missing OR Incorrect DUN + PAR

## Overview

This phase validates that all Parliamentary (PAR) and State (DUN) constituency codes and names in `to-review.csv` match the official Election Commission data from `raw-par.csv` and `raw-dun.csv`.

**Method:** A Go script (`phase2_check.go`) was used to:
1. Load official PAR list from `raw-par.csv` (filtered by `sid=13` for Sarawak)
2. Load official DUN list from `raw-dun.csv` (filtered by `sid=13` for Sarawak)
3. Extract unique PAR and DUN entries from `to-review.csv`
4. Compare codes, names, and DUN-to-PAR mappings

---

## 1. Parliamentary Constituency (PAR) Validation

### Expected Sarawak PARs: P.192 to P.222 (31 constituencies)

| # | PAR Code | Official Name | to-review Name | Status |
|---|----------|--------------|----------------|--------|
| 1 | P.192 | MAS GADING | MAS GADING | ✅ Match |
| 2 | P.193 | SANTUBONG | SANTUBONG | ✅ Match |
| 3 | P.194 | PETRA JAYA | PETRA JAYA | ✅ Match |
| 4 | P.195 | BANDAR KUCHING | BANDAR KUCHING | ✅ Match |
| 5 | P.196 | STAMPIN | STAMPIN | ✅ Match |
| 6 | P.197 | KOTA SAMARAHAN | KOTA SAMARAHAN | ✅ Match |
| 7 | P.198 | PUNCAK BORNEO | PUNCAK BORNEO | ✅ Match |
| 8 | P.199 | SERIAN | SERIAN | ✅ Match |
| 9 | P.200 | BATANG SADONG | BATANG SADONG | ✅ Match |
| 10 | P.201 | BATANG LUPAR | BATANG LUPAR | ✅ Match |
| 11 | P.202 | SRI AMAN | SRI AMAN | ✅ Match |
| 12 | P.203 | LUBOK ANTU | LUBOK ANTU | ✅ Match |
| 13 | P.204 | BETONG | BETONG | ✅ Match |
| 14 | P.205 | SARATOK | SARATOK | ✅ Match |
| 15 | P.206 | TANJONG MANIS | TANJONG MANIS | ✅ Match |
| 16 | P.207 | IGAN | IGAN | ✅ Match |
| 17 | P.208 | SARIKEI | SARIKEI | ✅ Match |
| 18 | P.209 | JULAU | JULAU | ✅ Match |
| 19 | P.210 | KANOWIT | KANOWIT | ✅ Match |
| 20 | P.211 | LANANG | LANANG | ✅ Match |
| 21 | P.212 | SIBU | SIBU | ✅ Match |
| 22 | P.213 | MUKAH | MUKAH | ✅ Match |
| 23 | P.214 | SELANGAU | SELANGAU | ✅ Match |
| 24 | P.215 | KAPIT | KAPIT | ✅ Match |
| 25 | P.216 | HULU RAJANG | HULU RAJANG | ✅ Match |
| 26 | P.217 | BINTULU | BINTULU | ✅ Match |
| 27 | P.218 | SIBUTI | SIBUTI | ✅ Match |
| 28 | P.219 | MIRI | MIRI | ✅ Match |
| 29 | P.220 | BARAM | BARAM | ✅ Match |
| 30 | P.221 | LIMBANG | LIMBANG | ✅ Match |
| 31 | P.222 | LAWAS | LAWAS | ✅ Match |

### PAR Name Mismatches

None found.

### PARs Missing from to-review.csv

None.

### Extra PARs in to-review.csv (not in official data)

None.

**PAR Validation Result: ✅ 0 issues. All 31 PARs match perfectly.**

---

## 2. State Constituency (DUN) Validation

### Expected Sarawak DUNs: N.01 to N.82 (82 constituencies)

| # | DUN Code | Official Name | to-review Name | Parent PAR | Status |
|---|----------|--------------|----------------|------------|--------|
| 1 | N.01 | OPAR | OPAR | P.192 | ✅ Match |
| 2 | N.02 | TASIK BIRU | TASIK BIRU | P.192 | ✅ Match |
| 3 | N.03 | TANJONG DATU | TANJONG DATU | P.193 | ✅ Match |
| 4 | N.04 | PANTAI DAMAI | PANTAI DAMAI | P.193 | ✅ Match |
| 5 | N.05 | DEMAK LAUT | DEMAK LAUT | P.193 | ✅ Match |
| 6 | N.06 | TUPONG | TUPONG | P.194 | ✅ Match |
| 7 | N.07 | SAMARIANG | SAMARIANG | P.194 | ✅ Match |
| 8 | N.08 | SATOK | SATOK | P.194 | ✅ Match |
| 9 | N.09 | PADUNGAN | PADUNGAN | P.195 | ✅ Match |
| 10 | N.10 | PENDING | PENDING | P.195 | ✅ Match |
| 11 | N.11 | BATU LINTANG | BATU LINTANG | P.195 | ✅ Match |
| 12 | N.12 | KOTA SENTOSA | KOTA SENTOSA | P.196 | ✅ Match |
| 13 | N.13 | BATU KITANG | BATU KITANG | P.196 | ✅ Match |
| 14 | N.14 | BATU KAWAH | BATU KAWAH | P.196 | ✅ Match |
| 15 | N.15 | ASAJAYA | ASAJAYA | P.197 | ✅ Match |
| 16 | N.16 | MUARA TUANG | MUARA TUANG | P.197 | ✅ Match |
| 17 | N.17 | STAKAN | STAKAN | P.197 | ✅ Match |
| 18 | N.18 | SEREMBU | SEREMBU | P.198 | ✅ Match |
| 19 | N.19 | MAMBONG | MAMBONG | P.198 | ✅ Match |
| 20 | N.20 | TARAT | TARAT | P.198 | ✅ Match |
| 21 | N.21 | TEBEDU | TEBEDU | P.199 | ✅ Match |
| 22 | N.22 | KEDUP | KEDUP | P.199 | ✅ Match |
| 23 | N.23 | BUKIT SEMUJA | BUKIT SEMUJA | P.199 | ✅ Match |
| 24 | N.24 | SADONG JAYA | SADONG JAYA | P.200 | ✅ Match |
| 25 | N.25 | SIMUNJAN | SIMUNJAN | P.200 | ✅ Match |
| 26 | N.26 | GEDONG | GEDONG | P.200 | ✅ Match |
| 27 | N.27 | SEBUYAU | SEBUYAU | P.201 | ✅ Match |
| 28 | N.28 | LINGGA | LINGGA | P.201 | ✅ Match |
| 29 | N.29 | BETING MARO | BETING MARO | P.201 | ✅ Match |
| 30 | N.30 | BALAI RINGIN | BALAI RINGIN | P.202 | ✅ Match |
| 31 | N.31 | BUKIT BEGUNAN | BUKIT BEGUNAN | P.202 | ✅ Match |
| 32 | N.32 | SIMANGGANG | SIMANGGANG | P.202 | ✅ Match |
| 33 | N.33 | ENGKILILI | ENGKILILI | P.203 | ✅ Match |
| 34 | N.34 | BATANG AI | BATANG AI | P.203 | ✅ Match |
| 35 | N.35 | SARIBAS | SARIBAS | P.204 | ✅ Match |
| 36 | N.36 | LAYAR | LAYAR | P.204 | ✅ Match |
| 37 | N.37 | BUKIT SABAN | BUKIT SABAN | P.204 | ✅ Match |
| 38 | N.38 | KALAKA | KALAKA | P.205 | ✅ Match |
| 39 | N.39 | KRIAN | KRIAN | P.205 | ✅ Match |
| 40 | N.40 | KABONG | KABONG | P.205 | ✅ Match |
| 41 | N.41 | KUALA RAJANG | KUALA RAJANG | P.206 | ✅ Match |
| 42 | N.42 | SEMOP | SEMOP | P.206 | ✅ Match |
| 43 | N.43 | DARO | DARO | P.207 | ✅ Match |
| 44 | N.44 | JEMORENG | JEMORENG | P.207 | ✅ Match |
| 45 | N.45 | REPOK | REPOK | P.208 | ✅ Match |
| 46 | N.46 | MERADONG | MERADONG | P.208 | ✅ Match |
| 47 | N.47 | PAKAN | PAKAN | P.209 | ✅ Match |
| 48 | N.48 | MELUAN | MELUAN | P.209 | ✅ Match |
| 49 | N.49 | NGEMAH | NGEMAH | P.210 | ✅ Match |
| 50 | N.50 | MACHAN | MACHAN | P.210 | ✅ Match |
| 51 | N.51 | BUKIT ASSEK | BUKIT ASSEK | P.211 | ✅ Match |
| 52 | N.52 | DUDONG | DUDONG | P.211 | ✅ Match |
| 53 | N.53 | BAWANG ASSAN | BAWANG ASSAN | P.212 | ✅ Match |
| 54 | N.54 | PELAWAN | PELAWAN | P.212 | ✅ Match |
| 55 | N.55 | NANGKA | NANGKA | P.212 | ✅ Match |
| 56 | N.56 | DALAT | DALAT | P.213 | ✅ Match |
| 57 | N.57 | TELLIAN | TELLIAN | P.213 | ✅ Match |
| 58 | N.58 | BALINGIAN | BALINGIAN | P.213 | ✅ Match |
| 59 | N.59 | TAMIN | TAMIN | P.214 | ✅ Match |
| 60 | N.60 | KAKUS | KAKUS | P.214 | ✅ Match |
| 61 | N.61 | PELAGUS | PELAGUS | P.215 | ✅ Match |
| 62 | N.62 | KATIBAS | KATIBAS | P.215 | ✅ Match |
| 63 | N.63 | BUKIT GORAM | BUKIT GORAM | P.215 | ✅ Match |
| 64 | N.64 | BALEH | BALEH | P.216 | ✅ Match |
| 65 | N.65 | BELAGA | BELAGA | P.216 | ✅ Match |
| 66 | N.66 | MURUM | MURUM | P.216 | ✅ Match |
| 67 | N.67 | JEPAK | JEPAK | P.217 | ✅ Match |
| 68 | N.68 | TANJONG BATU | TANJONG BATU | P.217 | ✅ Match |
| 69 | N.69 | KEMENA | KEMENA | P.217 | ✅ Match |
| 70 | N.70 | SAMALAJU | SAMALAJU | P.217 | ✅ Match |
| 71 | N.71 | BEKENU | BEKENU | P.218 | ✅ Match |
| 72 | N.72 | LAMBIR | LAMBIR | P.218 | ✅ Match |
| 73 | N.73 | PIASAU | PIASAU | P.219 | ✅ Match |
| 74 | N.74 | PUJUT | PUJUT | P.219 | ✅ Match |
| 75 | N.75 | SENADIN | SENADIN | P.219 | ✅ Match |
| 76 | N.76 | MARUDI | MARUDI | P.220 | ✅ Match |
| 77 | N.77 | TELANG USAN | TELANG USAN | P.220 | ✅ Match |
| 78 | N.78 | MULU | MULU | P.220 | ✅ Match |
| 79 | N.79 | BUKIT KOTA | BUKIT KOTA | P.221 | ✅ Match |
| 80 | N.80 | BATU DANAU | BATU DANAU | P.221 | ✅ Match |
| 81 | N.81 | BA'KELALAN | BA\`KELALAN | ⚠️ NAME MISMATCH |
| 82 | N.82 | BUKIT SARI | BUKIT SARI | P.222 | ✅ Match |

### DUN Name Mismatches

| DUN Code | Official Name | to-review Name | Character (hex) | Rows Affected | Notes |
|----------|--------------|----------------|-----------------|---------------|-------|
| N.81 | `BA'KELALAN` | `BA`KELALAN` | Official: apostrophe `'` (0x27) → to-review: backtick `` ` `` (0x60) | **32 rows** | Wrong character in column 7 (STATE CONSTITUENCY NAME). The backtick (grave accent, 0x60) should be replaced with an apostrophe (0x27) to match the official EC data. |

**Evidence:**
- `raw-dun.csv` line for N.81: `22281,N.81,13,22200,BA'KELALAN,...` — uses apostrophe `'` (0x27)
- `to-review.csv` hex dump of column 7: `42 41 60 4b 45 4c 41 4c 41 4e` — byte 0x60 is backtick `` ` ``
- All 32 data rows for N.81 in `to-review.csv` use the backtick character

### DUNs Missing from to-review.csv

None. All 82 DUNs (N.01 through N.82) are present.

### Extra DUNs in to-review.csv (not in official data)

None. No unexpected DUN codes found.

**DUN Validation Result: ⚠️ 1 issue — character typo in N.81 BA'KELALAN (32 rows affected).**

---

## 3. DUN-PAR Mapping Validation

Checking that each DUN in `to-review.csv` maps to the correct parent PAR as per `raw-dun.csv`.

The mapping is verified by linking `raw-dun.csv` field `pid` (parent ID) to `raw-par.csv` field `id`, then comparing against the PAR code found in each DUN's rows in `to-review.csv`.

| PAR Code | PAR Name | DUN Codes (Official) | DUN Codes (to-review) | Status |
|----------|----------|---------------------|----------------------|--------|
| P.192 | MAS GADING | N.01, N.02 | N.01, N.02 | ✅ |
| P.193 | SANTUBONG | N.03, N.04, N.05 | N.03, N.04, N.05 | ✅ |
| P.194 | PETRA JAYA | N.06, N.07, N.08 | N.06, N.07, N.08 | ✅ |
| P.195 | BANDAR KUCHING | N.09, N.10, N.11 | N.09, N.10, N.11 | ✅ |
| P.196 | STAMPIN | N.12, N.13, N.14 | N.12, N.13, N.14 | ✅ |
| P.197 | KOTA SAMARAHAN | N.15, N.16, N.17 | N.15, N.16, N.17 | ✅ |
| P.198 | PUNCAK BORNEO | N.18, N.19, N.20 | N.18, N.19, N.20 | ✅ |
| P.199 | SERIAN | N.21, N.22, N.23 | N.21, N.22, N.23 | ✅ |
| P.200 | BATANG SADONG | N.24, N.25, N.26 | N.24, N.25, N.26 | ✅ |
| P.201 | BATANG LUPAR | N.27, N.28, N.29 | N.27, N.28, N.29 | ✅ |
| P.202 | SRI AMAN | N.30, N.31, N.32 | N.30, N.31, N.32 | ✅ |
| P.203 | LUBOK ANTU | N.33, N.34 | N.33, N.34 | ✅ |
| P.204 | BETONG | N.35, N.36, N.37 | N.35, N.36, N.37 | ✅ |
| P.205 | SARATOK | N.38, N.39, N.40 | N.38, N.39, N.40 | ✅ |
| P.206 | TANJONG MANIS | N.41, N.42 | N.41, N.42 | ✅ |
| P.207 | IGAN | N.43, N.44 | N.43, N.44 | ✅ |
| P.208 | SARIKEI | N.45, N.46 | N.45, N.46 | ✅ |
| P.209 | JULAU | N.47, N.48 | N.47, N.48 | ✅ |
| P.210 | KANOWIT | N.49, N.50 | N.49, N.50 | ✅ |
| P.211 | LANANG | N.51, N.52 | N.51, N.52 | ✅ |
| P.212 | SIBU | N.53, N.54, N.55 | N.53, N.54, N.55 | ✅ |
| P.213 | MUKAH | N.56, N.57, N.58 | N.56, N.57, N.58 | ✅ |
| P.214 | SELANGAU | N.59, N.60 | N.59, N.60 | ✅ |
| P.215 | KAPIT | N.61, N.62, N.63 | N.61, N.62, N.63 | ✅ |
| P.216 | HULU RAJANG | N.64, N.65, N.66 | N.64, N.65, N.66 | ✅ |
| P.217 | BINTULU | N.67, N.68, N.69, N.70 | N.67, N.68, N.69, N.70 | ✅ |
| P.218 | SIBUTI | N.71, N.72 | N.71, N.72 | ✅ |
| P.219 | MIRI | N.73, N.74, N.75 | N.73, N.74, N.75 | ✅ |
| P.220 | BARAM | N.76, N.77, N.78 | N.76, N.77, N.78 | ✅ |
| P.221 | LIMBANG | N.79, N.80 | N.79, N.80 | ✅ |
| P.222 | LAWAS | N.81, N.82 | N.81, N.82 | ✅ |

**DUN-PAR Mapping Validation Result: ✅ 0 issues. All 82 DUNs map to their correct parent PAR.**

---

## 4. Summary

| Check | Count | Result |
|-------|-------|--------|
| PAR codes (P.192–P.222) present | 31/31 | ✅ All present |
| PAR names match official data | 31/31 | ✅ All match |
| DUN codes (N.01–N.82) present | 82/82 | ✅ All present |
| DUN names match official data | 81/82 | ⚠️ 1 mismatch |
| DUN-to-PAR mapping correct | 82/82 | ✅ All correct |

### Issues Found: 1

| # | Severity | DUN | Issue | Rows Affected | Recommended Fix |
|---|----------|-----|-------|---------------|-----------------|
| 1 | ⚠️ Minor (typo) | N.81 | Backtick `` ` `` (0x60) used instead of apostrophe `'` (0x27) in `BA'KELALAN` | 32 rows (col 7) | Replace `BA`KELALAN` → `BA'KELALAN` in column 7 of all N.81 rows |

### Conclusion

The PAR and DUN structure in `to-review.csv` is **complete and correctly mapped**. All 31 parliamentary constituencies and all 82 state constituencies are accounted for with correct codes and correct parent relationships. The only issue is a single-character typo (backtick vs apostrophe) in the DUN name for N.81 BA'KELALAN, affecting 32 data rows. This is a cosmetic/encoding issue that should be corrected but does not affect data integrity.