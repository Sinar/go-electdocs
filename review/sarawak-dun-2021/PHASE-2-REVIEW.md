# PHASE-2 REVIEW: Find Missing OR Incorrect DUN

## Summary

| Metric | Value |
|--------|-------|
| Reference DUNs (raw-dun.csv) | 82 |
| Unique DUN codes in to-review.csv | 82 |
| Missing DUNs (in reference but not in review) | 0 |
| Extra DUNs (in review but not in reference) | 0 |
| Name mismatches | **1** |
| Inconsistent naming within to-review.csv | 0 |

## Findings

### ✅ All 82 DUNs Present

All 82 state constituencies (N.01 through N.82) from `raw-dun.csv` are accounted for in `to-review.csv`. There are no missing DUNs and no unexpected extra DUNs.

### ❌ 1 Name Mismatch: N.81 — BA'KELALAN vs BA\`KELALAN

| Code | raw-dun.csv (Reference) | to-review.csv (Review) | Issue |
|------|-------------------------|------------------------|-------|
| N.81 | BA'KELALAN | BA\`KELALAN | Wrong quote character |

**Evidence (hex dump of the differing character):**

- `raw-dun.csv` column 4 for N.81: `42 41 27 4b 45 4c 41 4c 41 4e` → character `0x27` = ASCII apostrophe `'`
- `to-review.csv` column 7 for N.81: `42 41 60 4b 45 4c 41 4c 41 4e` → character `0x60` = ASCII backtick/grave accent `` ` ``

**Verdict:** The review file uses a **backtick** (`` ` ``, U+0060) where the official EC reference data uses a standard **apostrophe** (`'`, U+0027). The official `raw-dun.csv` from the Election Commission is authoritative, so the correct spelling is **BA'KELALAN**. This should be corrected in `to-review.csv`.

### ✅ No Internal Inconsistencies

Each DUN code in `to-review.csv` maps to exactly one DUN name throughout the file. There are no rows where the same code appears with a different name.

## Full DUN Mapping (Reference vs Review)

| Code | raw-dun.csv Name | to-review.csv Name | Status |
|------|------------------|--------------------|--------|
| N.01 | OPAR | OPAR | ✅ Match |
| N.02 | TASIK BIRU | TASIK BIRU | ✅ Match |
| N.03 | TANJONG DATU | TANJONG DATU | ✅ Match |
| N.04 | PANTAI DAMAI | PANTAI DAMAI | ✅ Match |
| N.05 | DEMAK LAUT | DEMAK LAUT | ✅ Match |
| N.06 | TUPONG | TUPONG | ✅ Match |
| N.07 | SAMARIANG | SAMARIANG | ✅ Match |
| N.08 | SATOK | SATOK | ✅ Match |
| N.09 | PADUNGAN | PADUNGAN | ✅ Match |
| N.10 | PENDING | PENDING | ✅ Match |
| N.11 | BATU LINTANG | BATU LINTANG | ✅ Match |
| N.12 | KOTA SENTOSA | KOTA SENTOSA | ✅ Match |
| N.13 | BATU KITANG | BATU KITANG | ✅ Match |
| N.14 | BATU KAWAH | BATU KAWAH | ✅ Match |
| N.15 | ASAJAYA | ASAJAYA | ✅ Match |
| N.16 | MUARA TUANG | MUARA TUANG | ✅ Match |
| N.17 | STAKAN | STAKAN | ✅ Match |
| N.18 | SEREMBU | SEREMBU | ✅ Match |
| N.19 | MAMBONG | MAMBONG | ✅ Match |
| N.20 | TARAT | TARAT | ✅ Match |
| N.21 | TEBEDU | TEBEDU | ✅ Match |
| N.22 | KEDUP | KEDUP | ✅ Match |
| N.23 | BUKIT SEMUJA | BUKIT SEMUJA | ✅ Match |
| N.24 | SADONG JAYA | SADONG JAYA | ✅ Match |
| N.25 | SIMUNJAN | SIMUNJAN | ✅ Match |
| N.26 | GEDONG | GEDONG | ✅ Match |
| N.27 | SEBUYAU | SEBUYAU | ✅ Match |
| N.28 | LINGGA | LINGGA | ✅ Match |
| N.29 | BETING MARO | BETING MARO | ✅ Match |
| N.30 | BALAI RINGIN | BALAI RINGIN | ✅ Match |
| N.31 | BUKIT BEGUNAN | BUKIT BEGUNAN | ✅ Match |
| N.32 | SIMANGGANG | SIMANGGANG | ✅ Match |
| N.33 | ENGKILILI | ENGKILILI | ✅ Match |
| N.34 | BATANG AI | BATANG AI | ✅ Match |
| N.35 | SARIBAS | SARIBAS | ✅ Match |
| N.36 | LAYAR | LAYAR | ✅ Match |
| N.37 | BUKIT SABAN | BUKIT SABAN | ✅ Match |
| N.38 | KALAKA | KALAKA | ✅ Match |
| N.39 | KRIAN | KRIAN | ✅ Match |
| N.40 | KABONG | KABONG | ✅ Match |
| N.41 | KUALA RAJANG | KUALA RAJANG | ✅ Match |
| N.42 | SEMOP | SEMOP | ✅ Match |
| N.43 | DARO | DARO | ✅ Match |
| N.44 | JEMORENG | JEMORENG | ✅ Match |
| N.45 | REPOK | REPOK | ✅ Match |
| N.46 | MERADONG | MERADONG | ✅ Match |
| N.47 | PAKAN | PAKAN | ✅ Match |
| N.48 | MELUAN | MELUAN | ✅ Match |
| N.49 | NGEMAH | NGEMAH | ✅ Match |
| N.50 | MACHAN | MACHAN | ✅ Match |
| N.51 | BUKIT ASSEK | BUKIT ASSEK | ✅ Match |
| N.52 | DUDONG | DUDONG | ✅ Match |
| N.53 | BAWANG ASSAN | BAWANG ASSAN | ✅ Match |
| N.54 | PELAWAN | PELAWAN | ✅ Match |
| N.55 | NANGKA | NANGKA | ✅ Match |
| N.56 | DALAT | DALAT | ✅ Match |
| N.57 | TELLIAN | TELLIAN | ✅ Match |
| N.58 | BALINGIAN | BALINGIAN | ✅ Match |
| N.59 | TAMIN | TAMIN | ✅ Match |
| N.60 | KAKUS | KAKUS | ✅ Match |
| N.61 | PELAGUS | PELAGUS | ✅ Match |
| N.62 | KATIBAS | KATIBAS | ✅ Match |
| N.63 | BUKIT GORAM | BUKIT GORAM | ✅ Match |
| N.64 | BALEH | BALEH | ✅ Match |
| N.65 | BELAGA | BELAGA | ✅ Match |
| N.66 | MURUM | MURUM | ✅ Match |
| N.67 | JEPAK | JEPAK | ✅ Match |
| N.68 | TANJONG BATU | TANJONG BATU | ✅ Match |
| N.69 | KEMENA | KEMENA | ✅ Match |
| N.70 | SAMALAJU | SAMALAJU | ✅ Match |
| N.71 | BEKENU | BEKENU | ✅ Match |
| N.72 | LAMBIR | LAMBIR | ✅ Match |
| N.73 | PIASAU | PIASAU | ✅ Match |
| N.74 | PUJUT | PUJUT | ✅ Match |
| N.75 | SENADIN | SENADIN | ✅ Match |
| N.76 | MARUDI | MARUDI | ✅ Match |
| N.77 | TELANG USAN | TELANG USAN | ✅ Match |
| N.78 | MULU | MULU | ✅ Match |
| N.79 | BUKIT KOTA | BUKIT KOTA | ✅ Match |
| N.80 | BATU DANAU | BATU DANAU | ✅ Match |
| N.81 | BA'KELALAN | BA\`KELALAN | ❌ Mismatch |
| N.82 | BUKIT SARI | BUKIT SARI | ✅ Match |

## Methodology

- **Tool**: `phase2_check.go` (Go program using stdlib only, `log/slog` for structured logging)
- **Reference**: `raw-dun.csv` — 82 rows, no header; columns 3 and 4 (0-indexed: 2 and 3) contain DUN code and name
- **Review file**: `to-review.csv` — 3757 data rows + 1 header; columns 6 and 7 (0-indexed: 5 and 6) contain STATE CONSTITUENCY CODE and STATE CONSTITUENCY NAME
- Character-level verification performed via `xxd` hex dump for the mismatch

## Action Items

1. **Fix N.81 name**: Replace backtick (`` ` ``, 0x60) with apostrophe (`'`, 0x27) in all occurrences of `BA`KELALAN` in `to-review.csv` → should be `BA'KELALAN`
