# Review for PRU 15 - SARAWAK

Raw data from SPR website

Extract PDF to Markdown using nutrient.io

```
Review @SKILLS.md to  convert the pdf at ./PDF/PRU15-SARAWAK from PDF to markdown (MD); store it into @results  with pattern filename Sarawak-N.<DUN>.csv


Source PDF can be found here: /Users/leow/TINDAKMSIA/go-electdocs/PDF/

Store output in: /Users/leow/TINDAKMSIA/go-electdocs/review/pru-15-2022/SARAWAK/results

```

## Review Progress

### PHASE-0: Compare against 2016 DUN Sarawak results — ✅ COMPLETED

**File:** [PHASE-0-REVIEW.md](PHASE-0-REVIEW.md)

**Summary:** Compared ORDINARY VOTE rows in `to-review.csv` (PRU-15, 2022) against 80 DUN files from the 2016 Sarawak state election.

**Key Findings:**
- **DUN Coverage:** All 80 DUNs from 2016 present in 2022. Two additional DUNs (N.79 BUKIT KOTA, N.82 BUKIT SARI) in 2022 were not contested in 2016.
- **PAR Code Discrepancies (28 DUNs):** The 2016 DUN source files have incorrect PAR codes that don't match their own polling district code prefixes. The 2022 `to-review.csv` has the correct PAR codes. This is a 2016 data quality issue, not a 2022 problem.
- **Polling Districts:** 854 PDs matched perfectly between 2016 and 2022. 0 PDs only in 2016, 1 new PD only in 2022 (N.72 `218/72/07 RAMD`).
- **DUN Name:** 1 trivial difference — apostrophe vs backtick in BA'KELALAN / BA`KELALAN.
- **Polling Centre Names:** Most differences are abbreviation expansions (e.g. "SEK. KEB." → "SEKOLAH KEBANGSAAN") or venue changes between elections.
- **Row Counts:** Every DUN has more ordinary vote rows in 2022 (voter growth → more channels).
- **Overall:** No anomalous structural issues found in `to-review.csv`.

**Script:** `phase0_compare.go` (Go, stdlib only, slog logging)

### PHASE-1: Ensure ID field is unique — ✅ COMPLETED

**File:** [PHASE-1-REVIEW.md](PHASE-1-REVIEW.md)

### PHASE-2: Find Missing OR Incorrect DUN + PAR — ✅ COMPLETED

**File:** [PHASE-2-REVIEW.md](PHASE-2-REVIEW.md)

### PHASE-3: Find Missing OR Incorrect Candidates — ⬜ NOT STARTED
### PHASE-4: Check consistency of coalition — ✅ COMPLETED

**File:** [PHASE-4-REVIEW.md](PHASE-4-REVIEW.md)

**Summary:** Verified party/coalition assignments in `to-review.csv` against official data from `raw-candidates.csv` and `raw-party-data-clean.csv` for all 31 Sarawak parliamentary seats (P.192–P.222).

**Key Findings:**
- **Slot Placement:** All 92 candidates across 31 constituencies are placed in the correct coalition slot. Zero mismatches.
- **GPS Components:** PBB (14), SUPP (7), PRS (6), PDP (4) — all correctly in GPS slot, totalling 31 seats.
- **PH Components:** PKR (16), DAP (8), PAN/AMANAH (6) — all correctly in PH slot, covering 30 of 31 seats.
- **PN Components:** PPBM (3), PAS (1) — all correctly in PN slot, covering 4 seats.
- **Other Parties:** PSB (10), PBDS (3), PBK (1), SEDAR (1), PBM (1) — all correctly in OTHER PARTY slots.
- **Independents:** 11 BEBAS candidates with election symbols (KEY, TREE, AEROPLANE, etc.) — all in INDEPENDENT slots.
- **Empty Slots:** BN, GTA, GRS, WARISAN correctly empty (these coalitions did not contest in Sarawak).
- **Name Spelling Issue:** P.218 SIBUTI — `ZULBAIDAH SUBOH` in to-review.csv should be `ZULHAIDAH SUBOH` per official EC data.
- **Naming Conventions:** "PAN" used instead of "AMANAH", "PPBM" used instead of "BERSATU" — internally consistent but differ from `raw-party-data-clean.csv` abbreviations. Not errors, but noted for standardization.

**Overall Assessment: PASS** — with 1 minor name spelling correction needed.

**Script:** `phase4_party_check.go` (Go, stdlib only, slog logging)
### PHASE-5: Check consistency of column mappings — ⬜ NOT STARTED
### PHASE-6: PAR-level Validation of TOTAL BALLOTS ISSUED vs Candidate Totals — ✅ COMPLETED

**File:** [PHASE-6-REVIEW.md](PHASE-6-REVIEW.md)

**Summary:** Validated ballot accounting and candidate vote totals across all 31 Sarawak parliamentary constituencies (P.192–P.222) by cross-checking `to-review.csv` against official `raw-candidates.csv`.

**Key Findings:**
- **Check-1 (SumCand == B):** ✅ All 31 PARs pass. Sum of individual candidate vote columns equals TOTAL VALID VOTES in every constituency.
- **Check-2 (A == B+C+D):** ✅ All 31 PARs pass. TOTAL BALLOTS ISSUED = TOTAL VALID VOTES + TOTAL REJECTED VOTES + TOTAL UNRETURNED BALLOTS for every constituency.
- **Check-3 (per-candidate vs raw):** ✅ All 31 PARs pass. Every candidate's aggregated votes in `to-review.csv` match their `ju` value in `raw-candidates.csv`.
- **RAW_JU vs B cross-check:** ✅ All 31 PARs pass. Sum of all candidate votes from raw data equals TOTAL VALID VOTES.
- **DUN-level consistency:** ✅ All DUNs across all 31 PARs satisfy `A == B + C + D`.
- **Grand totals:** A=1,197,450 | B=1,178,867 | C=14,369 | D=4,214 | A−(B+C+D)=0 | SumCand−B=0 | RAW_JU−B=0
- **Name Spelling Issue:** P.218 SIBUTI — `ZULBAIDAH SUBOH` in to-review.csv vs `ZULHAIDAH SUBOH` in raw-candidates.csv (votes match at 10,405). Same issue flagged in Phase 4.
- **Discovery:** The `ut` field in `raw-candidates.csv` is mislabelled — it contains TOTAL REJECTED VOTES (matches C in all 31/31 PARs), not unreturned ballots as the header implies.

**Overall Assessment: PASS** — all numerical checks pass; 1 known name spelling difference (P.218); 1 raw data labelling issue discovered.

**Script:** `phase6_validate.go` (Go, stdlib only, slog logging)

