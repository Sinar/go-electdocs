# PHASE 6 REVIEW: PAR-level Validation of TOTAL BALLOTS ISSUED vs Candidate Totals

## Objective

Validate that `TOTAL BALLOTS ISSUED` in `to-review.csv` is consistent with vote totals,
and cross-check per-candidate vote sums against official data from `raw-candidates.csv`.

This is a **PRU-15 (2022) parliamentary election** file, so validation is at the **PAR level** (P.192–P.222).

## Method

For each PAR constituency (P.192–P.222), we compute from `to-review.csv`:

- **A**: Sum of `TOTAL BALLOTS ISSUED` (col 22)
- **B**: Sum of `TOTAL VALID VOTES` (col 88)
- **C**: Sum of `TOTAL REJECTED VOTES` (col 89)
- **D**: Sum of `TOTAL UNRETURNED BALLOTS` (col 90)
- **SumCand**: Sum of all individual candidate vote columns (cols 27,32,37,42,47,52,57,62,67,72,77,82,87)
- **RAW_JU**: Sum of `ju` (candidate votes) from `raw-candidates.csv`

### Checks Performed

| Check | Description | Condition |
|-------|-------------|----------|
| **Check-1** | Internal consistency: sum of candidate vote columns equals TOTAL VALID VOTES | `SumCand == B` |
| **Check-2** | Ballot accounting: TOTAL BALLOTS ISSUED = Valid + Rejected + Unreturned | `A == B + C + D` |
| **Check-3** | Cross-check: per-candidate vote sums match official raw-candidates.csv | `to-review candidate total == ju` |

## Summary

- **Total PARs**: 31
- **Check-1 failures** (SumCand ≠ B): **0**
- **Check-2 failures** (A ≠ B+C+D): **0**
- **Check-3 failures** (candidate mismatch vs raw): **0**

## PAR-Level Summary Table

| PAR | Name | A (Issued) | B (Valid) | C (Rejected) | D (Unreturned) | B+C+D | SumCand | RAW_JU | Chk1 | Chk2 | Chk3 |
|-----|------|------------|-----------|--------------|----------------|-------|---------|--------|------|------|------|
| P.192 | MAS GADING | 31892 | 31379 | 419 | 94 | 31892 | 31379 | 31379 | ✅ | ✅ | ✅ |
| P.193 | SANTUBONG | 52713 | 51809 | 650 | 254 | 52713 | 51809 | 51809 | ✅ | ✅ | ✅ |
| P.194 | PETRA JAYA | 70074 | 69163 | 551 | 360 | 70074 | 69163 | 69163 | ✅ | ✅ | ✅ |
| P.195 | BANDAR KUCHING | 64068 | 63575 | 301 | 192 | 64068 | 63575 | 63575 | ✅ | ✅ | ✅ |
| P.196 | STAMPIN | 74547 | 73753 | 546 | 248 | 74547 | 73753 | 73753 | ✅ | ✅ | ✅ |
| P.197 | KOTA SAMARAHAN | 56238 | 55111 | 892 | 235 | 56238 | 55111 | 55111 | ✅ | ✅ | ✅ |
| P.198 | PUNCAK BORNEO | 52049 | 51154 | 640 | 255 | 52049 | 51154 | 51154 | ✅ | ✅ | ✅ |
| P.199 | SERIAN | 40620 | 39974 | 511 | 135 | 40620 | 39974 | 39974 | ✅ | ✅ | ✅ |
| P.200 | BATANG SADONG | 22924 | 22443 | 379 | 102 | 22924 | 22443 | 22443 | ✅ | ✅ | ✅ |
| P.201 | BATANG LUPAR | 28118 | 27559 | 458 | 101 | 28118 | 27559 | 27559 | ✅ | ✅ | ✅ |
| P.202 | SRI AMAN | 32449 | 31917 | 435 | 97 | 32449 | 31917 | 31917 | ✅ | ✅ | ✅ |
| P.203 | LUBOK ANTU | 19537 | 19294 | 200 | 43 | 19537 | 19294 | 19294 | ✅ | ✅ | ✅ |
| P.204 | BETONG | 27239 | 26713 | 428 | 98 | 27239 | 26713 | 26713 | ✅ | ✅ | ✅ |
| P.205 | SARATOK | 31293 | 30841 | 376 | 76 | 31293 | 30841 | 30841 | ✅ | ✅ | ✅ |
| P.206 | TANJONG MANIS | 19395 | 19040 | 238 | 117 | 19395 | 19040 | 19040 | ✅ | ✅ | ✅ |
| P.207 | IGAN | 17286 | 16986 | 234 | 66 | 17286 | 16986 | 16986 | ✅ | ✅ | ✅ |
| P.208 | SARIKEI | 37025 | 36463 | 436 | 126 | 37025 | 36463 | 36463 | ✅ | ✅ | ✅ |
| P.209 | JULAU | 22927 | 22537 | 331 | 59 | 22927 | 22537 | 22537 | ✅ | ✅ | ✅ |
| P.210 | KANOWIT | 18410 | 18043 | 306 | 61 | 18410 | 18043 | 18043 | ✅ | ✅ | ✅ |
| P.211 | LANANG | 53972 | 52946 | 865 | 161 | 53972 | 52946 | 52946 | ✅ | ✅ | ✅ |
| P.212 | SIBU | 67409 | 65942 | 1230 | 237 | 67409 | 65942 | 65942 | ✅ | ✅ | ✅ |
| P.213 | MUKAH | 28167 | 27780 | 316 | 71 | 28167 | 27780 | 27780 | ✅ | ✅ | ✅ |
| P.214 | SELANGAU | 29280 | 28796 | 428 | 56 | 29280 | 28796 | 28796 | ✅ | ✅ | ✅ |
| P.215 | KAPIT | 22545 | 21999 | 382 | 164 | 22545 | 21999 | 21999 | ✅ | ✅ | ✅ |
| P.216 | HULU RAJANG | 23807 | 23407 | 311 | 89 | 23807 | 23407 | 23407 | ✅ | ✅ | ✅ |
| P.217 | BINTULU | 71256 | 70392 | 760 | 104 | 71256 | 70392 | 70392 | ✅ | ✅ | ✅ |
| P.218 | SIBUTI | 34371 | 33916 | 368 | 87 | 34371 | 33916 | 33916 | ✅ | ✅ | ✅ |
| P.219 | MIRI | 78996 | 78148 | 620 | 228 | 78996 | 78148 | 78148 | ✅ | ✅ | ✅ |
| P.220 | BARAM | 30218 | 29783 | 332 | 103 | 30218 | 29783 | 29783 | ✅ | ✅ | ✅ |
| P.221 | LIMBANG | 20130 | 19796 | 228 | 106 | 20130 | 19796 | 19796 | ✅ | ✅ | ✅ |
| P.222 | LAWAS | 18495 | 18208 | 198 | 89 | 18495 | 18208 | 18208 | ✅ | ✅ | ✅ |

## Check-1 Detail: SumCand vs TOTAL VALID VOTES

**No failures.** All 31 PARs have sum of candidate vote columns equal to TOTAL VALID VOTES. ✅

## Check-2 Detail: A vs B + C + D

**No failures.** All 31 PARs satisfy `TOTAL BALLOTS ISSUED = TOTAL VALID VOTES + TOTAL REJECTED VOTES + TOTAL UNRETURNED BALLOTS`. ✅

## Check-3 Detail: Candidate Vote Comparison vs raw-candidates.csv

**No failures.** All per-candidate vote totals in `to-review.csv` match `raw-candidates.csv`. ✅

## Candidate Name Spelling Differences

Cases where `to-review.csv` and `raw-candidates.csv` have different spellings for the same candidate (matched by fuzzy logic, votes may match):

### P.218 SIBUTI

| raw-candidates name | to-review name | raw ju | to-review votes | Votes match? |
|---------------------|----------------|--------|-----------------|-------------|
| ZULHAIDAH SUBOH | ZULBAIDAH SUBOH | 10405 | 10405 | ✅ |

## Cross-Check: RAW_JU (sum of all candidate ju) vs B (TOTAL VALID VOTES)

In a single-vote-per-ballot election, the sum of all candidates' votes should equal TOTAL VALID VOTES.

| PAR | Name | RAW_JU | B (Valid) | Diff | Status |
|-----|------|--------|-----------|------|--------|
| P.192 | MAS GADING | 31379 | 31379 | +0 | ✅ |
| P.193 | SANTUBONG | 51809 | 51809 | +0 | ✅ |
| P.194 | PETRA JAYA | 69163 | 69163 | +0 | ✅ |
| P.195 | BANDAR KUCHING | 63575 | 63575 | +0 | ✅ |
| P.196 | STAMPIN | 73753 | 73753 | +0 | ✅ |
| P.197 | KOTA SAMARAHAN | 55111 | 55111 | +0 | ✅ |
| P.198 | PUNCAK BORNEO | 51154 | 51154 | +0 | ✅ |
| P.199 | SERIAN | 39974 | 39974 | +0 | ✅ |
| P.200 | BATANG SADONG | 22443 | 22443 | +0 | ✅ |
| P.201 | BATANG LUPAR | 27559 | 27559 | +0 | ✅ |
| P.202 | SRI AMAN | 31917 | 31917 | +0 | ✅ |
| P.203 | LUBOK ANTU | 19294 | 19294 | +0 | ✅ |
| P.204 | BETONG | 26713 | 26713 | +0 | ✅ |
| P.205 | SARATOK | 30841 | 30841 | +0 | ✅ |
| P.206 | TANJONG MANIS | 19040 | 19040 | +0 | ✅ |
| P.207 | IGAN | 16986 | 16986 | +0 | ✅ |
| P.208 | SARIKEI | 36463 | 36463 | +0 | ✅ |
| P.209 | JULAU | 22537 | 22537 | +0 | ✅ |
| P.210 | KANOWIT | 18043 | 18043 | +0 | ✅ |
| P.211 | LANANG | 52946 | 52946 | +0 | ✅ |
| P.212 | SIBU | 65942 | 65942 | +0 | ✅ |
| P.213 | MUKAH | 27780 | 27780 | +0 | ✅ |
| P.214 | SELANGAU | 28796 | 28796 | +0 | ✅ |
| P.215 | KAPIT | 21999 | 21999 | +0 | ✅ |
| P.216 | HULU RAJANG | 23407 | 23407 | +0 | ✅ |
| P.217 | BINTULU | 70392 | 70392 | +0 | ✅ |
| P.218 | SIBUTI | 33916 | 33916 | +0 | ✅ |
| P.219 | MIRI | 78148 | 78148 | +0 | ✅ |
| P.220 | BARAM | 29783 | 29783 | +0 | ✅ |
| P.221 | LIMBANG | 19796 | 19796 | +0 | ✅ |
| P.222 | LAWAS | 18208 | 18208 | +0 | ✅ |

**Failures**: 0 / 31

## Discovery: raw-candidates.csv `ut` Field Is Actually TOTAL REJECTED VOTES

The `ut` column in `raw-candidates.csv` is labelled suggestively as "unreturned ballots",
but empirical comparison shows it matches **TOTAL REJECTED VOTES (C)**, not TOTAL UNRETURNED BALLOTS (D).

| PAR | Name | C (Rejected) | D (Unreturned) | raw ut | ut==C? | ut==D? |
|-----|------|--------------|----------------|--------|--------|--------|
| P.192 | MAS GADING | 419 | 94 | 419 | ✅ | ❌ |
| P.193 | SANTUBONG | 650 | 254 | 650 | ✅ | ❌ |
| P.194 | PETRA JAYA | 551 | 360 | 551 | ✅ | ❌ |
| P.195 | BANDAR KUCHING | 301 | 192 | 301 | ✅ | ❌ |
| P.196 | STAMPIN | 546 | 248 | 546 | ✅ | ❌ |
| P.197 | KOTA SAMARAHAN | 892 | 235 | 892 | ✅ | ❌ |
| P.198 | PUNCAK BORNEO | 640 | 255 | 640 | ✅ | ❌ |
| P.199 | SERIAN | 511 | 135 | 511 | ✅ | ❌ |
| P.200 | BATANG SADONG | 379 | 102 | 379 | ✅ | ❌ |
| P.201 | BATANG LUPAR | 458 | 101 | 458 | ✅ | ❌ |
| P.202 | SRI AMAN | 435 | 97 | 435 | ✅ | ❌ |
| P.203 | LUBOK ANTU | 200 | 43 | 200 | ✅ | ❌ |
| P.204 | BETONG | 428 | 98 | 428 | ✅ | ❌ |
| P.205 | SARATOK | 376 | 76 | 376 | ✅ | ❌ |
| P.206 | TANJONG MANIS | 238 | 117 | 238 | ✅ | ❌ |
| P.207 | IGAN | 234 | 66 | 234 | ✅ | ❌ |
| P.208 | SARIKEI | 436 | 126 | 436 | ✅ | ❌ |
| P.209 | JULAU | 331 | 59 | 331 | ✅ | ❌ |
| P.210 | KANOWIT | 306 | 61 | 306 | ✅ | ❌ |
| P.211 | LANANG | 865 | 161 | 865 | ✅ | ❌ |
| P.212 | SIBU | 1230 | 237 | 1230 | ✅ | ❌ |
| P.213 | MUKAH | 316 | 71 | 316 | ✅ | ❌ |
| P.214 | SELANGAU | 428 | 56 | 428 | ✅ | ❌ |
| P.215 | KAPIT | 382 | 164 | 382 | ✅ | ❌ |
| P.216 | HULU RAJANG | 311 | 89 | 311 | ✅ | ❌ |
| P.217 | BINTULU | 760 | 104 | 760 | ✅ | ❌ |
| P.218 | SIBUTI | 368 | 87 | 368 | ✅ | ❌ |
| P.219 | MIRI | 620 | 228 | 620 | ✅ | ❌ |
| P.220 | BARAM | 332 | 103 | 332 | ✅ | ❌ |
| P.221 | LIMBANG | 228 | 106 | 228 | ✅ | ❌ |
| P.222 | LAWAS | 198 | 89 | 198 | ✅ | ❌ |

**ut matches C (Rejected)**: 31 / 31
**ut matches D (Unreturned)**: 0 / 31

**Conclusion**: The `ut` field in `raw-candidates.csv` contains **TOTAL REJECTED VOTES**, not unreturned ballots. This is a mislabelled column in the raw data source.

## DUN-Level Internal Consistency: A == B + C + D

**No DUN-level failures.** All DUNs across all 31 PARs satisfy `A == B + C + D`. ✅

## Grand Totals

| Metric | Value |
|--------|-------|
| Total Ballots Issued (A) | 1197450 |
| Total Valid Votes (B) | 1178867 |
| Total Rejected Votes (C) | 14369 |
| Total Unreturned Ballots (D) | 4214 |
| B + C + D | 1197450 |
| A − (B+C+D) | 0 |
| Sum of Candidate Votes (SumCand) | 1178867 |
| Sum of raw ju (RAW_JU) | 1178867 |
| SumCand − B | 0 |
| RAW_JU − B | 0 |

## Recommendations

1. **Name Spellings**: Some candidate names differ between `to-review.csv` and `raw-candidates.csv`. While votes match, the spelling in `to-review.csv` should be verified against the original PDF source to determine the correct spelling.
2. **raw-candidates.csv `ut` field**: This field contains TOTAL REJECTED VOTES, not unreturned ballots as the header implies. Future processing should treat `ut` as rejected votes.
