# FIXES REPORT

Here is the full summary of all changes:

---

## Diff Summary: `to-review.csv.orig` → `to-review.csv`

| | |
|---|---|
| **Total data rows** | 3,748 |
| **Changed rows** | 2,320 |
| **Unchanged rows** | 1,428 |

---

### Fix 1 — Postal DMKOD: `/UNDI POS` → `/POS` in UNIQUE CODE & POLLING DISTRICT CODE

**81 rows** — one per DUN, N.02 through N.82 (N.01 was already correct in the original)

| CSV Row | DUN | UNIQUE CODE before | UNIQUE CODE after | PD Code before | PD Code after |
|---------|-----|--------------------|-------------------|----------------|---------------|
| 35 | N.02 | `P.192_N.02_192/02/UNDI POS_1` | `P.192_N.02_192/02/POS_1` | `192/02/UNDI POS` | `192/02/POS` |
| 83 | N.03 | `P.193_N.03_193/03/UNDI POS_1` | `P.193_N.03_193/03/POS_1` | `193/03/UNDI POS` | `193/03/POS` |
| 120–3709 | N.04 … N.82 | *(same pattern for all remaining DUNs)* | | | |

---

### Fix 2 — `BA\`KELALAN` backtick → apostrophe

**28 rows** — all rows for **N.81** only

| Columns fixed | Before | After |
|---|---|---|
| STATE CONSTITUENCY NAME (col 7) | `BA`KELALAN` (U+0060) | `BA'KELALAN` (U+0027) |
| POLLING DISTRICT NAME (col 9) | `BA`KELALAN` | `BA'KELALAN` |
| UNIQUE CODE (col 1) | `…/81/BA`KELALAN/…` *(where present)* | `…/81/BA'KELALAN/…` |

---

### Fix 3 — Embedded newline removed from POLLING CENTRE

**10 rows** across 4 DUNs

| CSV Row | DUN | POLLING CENTRE before | POLLING CENTRE after |
|---------|-----|-----------------------|----------------------|
| 121 | N.04 | `RUANG D, DEWAN BADMINTON`↵`, KOMPLEKS POLIS TABUAN JAYA` | `RUANG D, DEWAN BADMINTON, KOMPLEKS POLIS TABUAN JAYA` |
| 2263–2268 | N.52 | `SEKOLAH KEBANGSAAN BANDARAN SIBU NO.`↵`2` *(6 rows)* | `SEKOLAH KEBANGSAAN BANDARAN SIBU NO.2` |
| 3023 | N.68 | `BANGUNAN PERSEKUTUAN … ( W.I`↵`)` | `BANGUNAN PERSEKUTUAN … ( W.I)` |
| 3264 | N.73 | `SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH`↵`AS-SYIBYAN)` | `SEKOLAH RENDAH AGAMA RAKYAT MIRI (MADRASAH AS-SYIBYAN)` |
| 3641 | N.79 | `SEKOLAH KEBANGSAAN R.`↵`C. KUBONG` | `SEKOLAH KEBANGSAAN R.C. KUBONG` |

---

### Fix 4 — UNIQUE CODE suffix added (district-level centre disambiguation)

**2,201 rows** — letter suffixes `a`–`r` appended to UNIQUE CODE only (col 1). All other columns unchanged. Affects 309 polling districts across 55 DUNs where a single district code serves multiple polling centres.

| DUN | Rows suffixed | Example before | Example after |
|-----|:---:|---|---|
| N.02 | 5 | `P.192_N.02_192/02/00_1` | `P.192_N.02_192/02/00_1a` |
| N.03 | 5 | `P.193_N.03_193/03/00_1` | `P.193_N.03_193/03/00_1a` |
| N.04 | 34 | `P.193_N.04_193/04/01_1` | `P.193_N.04_193/04/01_1a` |
| N.05 | 21 | `P.193_N.05_193/05/04_1` | `P.193_N.05_193/05/04_1a` |
| N.06 | 58 | `P.194_N.06_194/06/01_1` | `P.194_N.06_194/06/01_1a` |
| N.07–N.82 | varies | *(same pattern)* | *(suffix a–r depending on centre count)* |
| **Max centres** | 18 rows @ N.61 | `P.215_N.61_215/61/04_1` | `P.215_N.61_215/61/04_1a` … `_1r` |

> The suffix letter for each polling centre is consistent across **all channel numbers** (`_1`, `_2`, `_3`…) within the same polling district. No vote counts, candidate names, or any other columns were modified.

---

### Totals by fix type

| Fix | Description | Rows changed |
|-----|-------------|:---:|
| 1 | Postal DMKOD `/UNDI POS` → `/POS` in UNIQUE CODE + POLLING DISTRICT CODE | **81** |
| 2 | `BA\`KELALAN` backtick → apostrophe (N.81) | **28** |
| 3 | Embedded `\n` removed from POLLING CENTRE | **10** |
| 4 | UNIQUE CODE letter suffix added (centre disambiguation) | **2,201** |
| | **Total** | **2,320
