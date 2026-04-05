# PHASE 1 REVIEW: UNIQUE CODE Uniqueness Check

## Summary

| Metric | Value |
|--------|-------|
| Total data rows | 4317 |
| Empty UNIQUE CODE values | 0 |
| Distinct non-empty UNIQUE CODE values | 4295 |
| Duplicate UNIQUE CODE values (codes appearing >1 time) | 22 |
| Total rows involved in duplicates | 44 |
| Suffix consistency violations (non-duplicate) | 4 districts |
| Suffix convention inconsistencies (`_1a` vs `_a1`) | 11 rows |

## Result: DUPLICATES FOUND ⚠️

There are **22** UNIQUE CODE values that appear more than once, affecting **44** rows total.
Additionally, **4** districts have suffix consistency violations (same Polling Centre assigned different suffix letters across channels), and **11** rows use a reversed suffix convention (`_1a` instead of `_a1`).

---

## Issue A: Codes WITHOUT Letter Suffixes — Need Suffix Algorithm (3 codes, 6 rows)

These codes have no disambiguation suffix at all and are raw duplicates. The suffix algorithm per AGENTS.md must be applied.

### A1. `P.198_N.19_198/19/29_1` — District `198/19/29`

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 1177 | `P.198_N.19_198/19/29_1` ⚠️ | BALAI RAYA KPG. NUSARAYA | 1 |
| 1178 | `P.198_N.19_198/19/29_1` ⚠️ | BALAI RAYA KPG. BIYA KEMAS | 1 |

**Root cause**: Two different Polling Centres share the same Polling District Code `198/19/29` and both have channel 1. No suffix was assigned.

**Fix**: Apply suffix by first-appearance order:
| Line | Current | Fixed |
|------|---------|-------|
| 1177 | `P.198_N.19_198/19/29_1` | `P.198_N.19_198/19/29_a1` |
| 1178 | `P.198_N.19_198/19/29_1` | `P.198_N.19_198/19/29_b1` |

### A2. `P.198_N.20_198/20/13_1` — District `198/20/13`

All rows in district for context:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 1218 | `P.198_N.20_198/20/13_1` ⚠️ | SEKOLAH KEBANGSAAN PESANG BEGU | 1 |
| 1219 | `P.198_N.20_198/20/13_2` | SEKOLAH KEBANGSAAN PESANG BEGU | 2 |
| 1220 | `P.198_N.20_198/20/13_1` ⚠️ | DEWAN SERBAGUNA KPG. SERUMAH | 1 |

**Root cause**: Two Polling Centres share `198/20/13`, each with channel 1. Line 1219 (channel 2) is unique but belongs to the same centre as line 1218, so it also needs a suffix for consistency.

**Fix**: Apply suffix by first-appearance order, uniformly across all channels for each centre:
| Line | Current | Fixed |
|------|---------|-------|
| 1218 | `P.198_N.20_198/20/13_1` | `P.198_N.20_198/20/13_a1` |
| 1219 | `P.198_N.20_198/20/13_2` | `P.198_N.20_198/20/13_a2` |
| 1220 | `P.198_N.20_198/20/13_1` | `P.198_N.20_198/20/13_b1` |

### A3. `P.208_N.45_208/45/05_1` — District `208/45/05`

All rows in district for context:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 2201 | `P.208_N.45_208/45/05_1` ⚠️ | SEKOLAH JENIS KEBANGSAAN (CINA) SU LOK | 1 |
| 2202 | `P.208_N.45_208/45/05_2` | SEKOLAH JENIS KEBANGSAAN (CINA) SU LOK | 2 |
| 2203 | `P.208_N.45_208/45/05_3` | SEKOLAH JENIS KEBANGSAAN (CINA) SU LOK | 3 |
| 2204 | `P.208_N.45_208/45/05_4` | SEKOLAH JENIS KEBANGSAAN (CINA) SU LOK | 4 |
| 2205 | `P.208_N.45_208/45/05_1` ⚠️ | SEKOLAH JENIS KEBANGSAAN (CINA) SU LEE | 1 |

**Root cause**: Two Polling Centres share `208/45/05`. Only channel 1 collides but the suffix must be applied uniformly across all channels in each centre.

**Fix**: Apply suffix by first-appearance order:
| Line | Current | Fixed |
|------|---------|-------|
| 2201 | `P.208_N.45_208/45/05_1` | `P.208_N.45_208/45/05_a1` |
| 2202 | `P.208_N.45_208/45/05_2` | `P.208_N.45_208/45/05_a2` |
| 2203 | `P.208_N.45_208/45/05_3` | `P.208_N.45_208/45/05_a3` |
| 2204 | `P.208_N.45_208/45/05_4` | `P.208_N.45_208/45/05_a4` |
| 2205 | `P.208_N.45_208/45/05_1` | `P.208_N.45_208/45/05_b1` |

---

## Issue B: Codes WITH Letter Suffixes Still Colliding — Wrong Suffix Assignment (18 codes, 36 rows)

These codes already have disambiguation suffixes (e.g. `_a1`, `_c1`) but two different Polling Centres within the same district were assigned the **same** suffix letter, causing collisions.

### B1. District `195/10/02` — 3 duplicate codes (6 rows)

All 20 rows in district:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 532 | `P.195_N.10_195/10/02_a1` ⚠️ | SEKOLAH MENENGAH CHUNG HUA NO. 1 | 1 |
| 533 | `P.195_N.10_195/10/02_a2` ⚠️ | SEKOLAH MENENGAH CHUNG HUA NO. 1 | 2 |
| 534 | `P.195_N.10_195/10/02_a3` ⚠️ | SEKOLAH MENENGAH CHUNG HUA NO. 1 | 3 |
| 535 | `P.195_N.10_195/10/02_b1` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 1 |
| 536 | `P.195_N.10_195/10/02_b2` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 2 |
| 537 | `P.195_N.10_195/10/02_b3` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 3 |
| 538 | `P.195_N.10_195/10/02_b4` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 4 |
| 539 | `P.195_N.10_195/10/02_b5` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 5 |
| 540 | `P.195_N.10_195/10/02_b6` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 6 |
| 541 | `P.195_N.10_195/10/02_b7` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 7 |
| 542 | `P.195_N.10_195/10/02_b8` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | 8 |
| 543 | `P.195_N.10_195/10/02_a1` ⚠️ | SEKOLAH KEBANGSAAN TABUAN ULU | 1 |
| 544 | `P.195_N.10_195/10/02_a2` ⚠️ | SEKOLAH KEBANGSAAN TABUAN ULU | 2 |
| 545 | `P.195_N.10_195/10/02_a3` ⚠️ | SEKOLAH KEBANGSAAN TABUAN ULU | 3 |
| 546 | `P.195_N.10_195/10/02_a4` | SEKOLAH KEBANGSAAN TABUAN ULU | 4 |
| 547 | `P.195_N.10_195/10/02_a5` | SEKOLAH KEBANGSAAN TABUAN ULU | 5 |
| 548 | `P.195_N.10_195/10/02_a6` | SEKOLAH KEBANGSAAN TABUAN ULU | 6 |
| 549 | `P.195_N.10_195/10/02_a7` | SEKOLAH KEBANGSAAN TABUAN ULU | 7 |
| 550 | `P.195_N.10_195/10/02_a8` | SEKOLAH KEBANGSAAN TABUAN ULU | 8 |
| 551 | `P.195_N.10_195/10/02_a9` | SEKOLAH KEBANGSAAN TABUAN ULU | 9 |

**Root cause**: Both SEKOLAH MENENGAH CHUNG HUA NO. 1 (3 channels) and SEKOLAH KEBANGSAAN TABUAN ULU (9 channels) were assigned suffix `a`. Channels 1–3 collide.

**Fix**: Reassign TABUAN ULU from `a` → `c` (since `a` = CHUNG HUA NO. 1, `b` = CHUNG HUA NO. 3):
| Line | Current | Fixed |
|------|---------|-------|
| 543 | `P.195_N.10_195/10/02_a1` | `P.195_N.10_195/10/02_c1` |
| 544 | `P.195_N.10_195/10/02_a2` | `P.195_N.10_195/10/02_c2` |
| 545 | `P.195_N.10_195/10/02_a3` | `P.195_N.10_195/10/02_c3` |
| 546 | `P.195_N.10_195/10/02_a4` | `P.195_N.10_195/10/02_c4` |
| 547 | `P.195_N.10_195/10/02_a5` | `P.195_N.10_195/10/02_c5` |
| 548 | `P.195_N.10_195/10/02_a6` | `P.195_N.10_195/10/02_c6` |
| 549 | `P.195_N.10_195/10/02_a7` | `P.195_N.10_195/10/02_c7` |
| 550 | `P.195_N.10_195/10/02_a8` | `P.195_N.10_195/10/02_c8` |
| 551 | `P.195_N.10_195/10/02_a9` | `P.195_N.10_195/10/02_c9` |

### B2. District `219/74/01` — 12 duplicate codes (24 rows)

All 51 rows in district:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 3759 | `P.219_N.74_219/74/01_a1` | TADIKA MIRI CHINESE | 1 |
| 3760 | `P.219_N.74_219/74/01_a2` | TADIKA MIRI CHINESE | 2 |
| 3761 | `P.219_N.74_219/74/01_a3` | TADIKA MIRI CHINESE | 3 |
| 3762 | `P.219_N.74_219/74/01_a4` | TADIKA MIRI CHINESE | 4 |
| 3763 | `P.219_N.74_219/74/01_a5` | TADIKA MIRI CHINESE | 5 |
| 3764 | `P.219_N.74_219/74/01_b1` | SEKOLAH RENDAH SRI MAWAR | 1 |
| 3765 | `P.219_N.74_219/74/01_b2` | SEKOLAH RENDAH SRI MAWAR | 2 |
| 3766 | `P.219_N.74_219/74/01_b3` | SEKOLAH RENDAH SRI MAWAR | 3 |
| 3767 | `P.219_N.74_219/74/01_b4` | SEKOLAH RENDAH SRI MAWAR | 4 |
| 3768 | `P.219_N.74_219/74/01_b5` | SEKOLAH RENDAH SRI MAWAR | 5 |
| 3769 | `P.219_N.74_219/74/01_b6` | SEKOLAH RENDAH SRI MAWAR | 6 |
| 3770 | `P.219_N.74_219/74/01_b7` | SEKOLAH RENDAH SRI MAWAR | 7 |
| 3771 | `P.219_N.74_219/74/01_b8` | SEKOLAH RENDAH SRI MAWAR | 8 |
| 3772 | `P.219_N.74_219/74/01_b9` | SEKOLAH RENDAH SRI MAWAR | 9 |
| 3773 | `P.219_N.74_219/74/01_b10` | SEKOLAH RENDAH SRI MAWAR | 10 |
| 3774 | `P.219_N.74_219/74/01_b11` | SEKOLAH RENDAH SRI MAWAR | 11 |
| 3775 | `P.219_N.74_219/74/01_c1` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 1 |
| 3776 | `P.219_N.74_219/74/01_c2` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 2 |
| 3777 | `P.219_N.74_219/74/01_c3` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 3 |
| 3778 | `P.219_N.74_219/74/01_c4` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 4 |
| 3779 | `P.219_N.74_219/74/01_c5` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 5 |
| 3780 | `P.219_N.74_219/74/01_c6` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 6 |
| 3781 | `P.219_N.74_219/74/01_c7` ⚠️ | SEKOLAH KEBANGSAAN ANCHI | 7 |
| 3782 | `P.219_N.74_219/74/01_c8` | SEKOLAH KEBANGSAAN ANCHI | 8 |
| 3783 | `P.219_N.74_219/74/01_d1` ⚠️ | TADIKA PUJUT MIRI | 1 |
| 3784 | `P.219_N.74_219/74/01_d2` ⚠️ | TADIKA PUJUT MIRI | 2 |
| 3785 | `P.219_N.74_219/74/01_d3` ⚠️ | TADIKA PUJUT MIRI | 3 |
| 3786 | `P.219_N.74_219/74/01_d4` ⚠️ | TADIKA PUJUT MIRI | 4 |
| 3787 | `P.219_N.74_219/74/01_d5` ⚠️ | TADIKA PUJUT MIRI | 5 |
| 3788 | `P.219_N.74_219/74/01_c1` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 1 |
| 3789 | `P.219_N.74_219/74/01_c2` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 2 |
| 3790 | `P.219_N.74_219/74/01_c3` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 3 |
| 3791 | `P.219_N.74_219/74/01_c4` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 4 |
| 3792 | `P.219_N.74_219/74/01_c5` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 5 |
| 3793 | `P.219_N.74_219/74/01_c6` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 6 |
| 3794 | `P.219_N.74_219/74/01_c7` ⚠️ | SEKOLAH KEBANGSAAN PUJUT CORNER | 7 |
| 3795 | `P.219_N.74_219/74/01_d1` ⚠️ | SJK (CINA) CHUNG HUA PUJUT | 1 |
| 3796 | `P.219_N.74_219/74/01_d2` ⚠️ | SJK (CINA) CHUNG HUA PUJUT | 2 |
| 3797 | `P.219_N.74_219/74/01_d3` ⚠️ | SJK (CINA) CHUNG HUA PUJUT | 3 |
| 3798 | `P.219_N.74_219/74/01_d4` ⚠️ | SJK (CINA) CHUNG HUA PUJUT | 4 |
| 3799 | `P.219_N.74_219/74/01_d5` ⚠️ | SJK (CINA) CHUNG HUA PUJUT | 5 |
| 3800 | `P.219_N.74_219/74/01_d6` | SJK (CINA) CHUNG HUA PUJUT | 6 |
| 3801 | `P.219_N.74_219/74/01_d7` | SJK (CINA) CHUNG HUA PUJUT | 7 |
| 3802 | `P.219_N.74_219/74/01_d8` | SJK (CINA) CHUNG HUA PUJUT | 8 |
| 3803 | `P.219_N.74_219/74/01_d9` | SJK (CINA) CHUNG HUA PUJUT | 9 |
| 3804 | `P.219_N.74_219/74/01_d10` | SJK (CINA) CHUNG HUA PUJUT | 10 |
| 3805 | `P.219_N.74_219/74/01_e1` | SMK DATO PERMAISURI | 1 |
| 3806 | `P.219_N.74_219/74/01_e2` | SMK DATO PERMAISURI | 2 |
| 3807 | `P.219_N.74_219/74/01_e3` | SMK DATO PERMAISURI | 3 |
| 3808 | `P.219_N.74_219/74/01_e4` | SMK DATO PERMAISURI | 4 |
| 3809 | `P.219_N.74_219/74/01_e5` | SMK DATO PERMAISURI | 5 |

**Root cause**: 7 distinct Polling Centres but only suffix letters a–e were used — two pairs of centres were assigned the same suffix:
- `c` was assigned to both SK ANCHI (8 ch) and SK PUJUT CORNER (7 ch) → channels 1–7 collide
- `d` was assigned to both TADIKA PUJUT MIRI (5 ch) and SJK (CINA) CHUNG HUA PUJUT (10 ch) → channels 1–5 collide

**Fix**: Reassign suffixes so each centre gets a unique letter (a–g), preserving first-appearance order:

| Suffix | Polling Centre | Action |
|--------|---------------|--------|
| `a` | TADIKA MIRI CHINESE | Keep ✅ |
| `b` | SEKOLAH RENDAH SRI MAWAR | Keep ✅ |
| `c` | SEKOLAH KEBANGSAAN ANCHI | Keep ✅ |
| `d` | TADIKA PUJUT MIRI | Keep ✅ |
| `e` ← was `c` | SEKOLAH KEBANGSAAN PUJUT CORNER | **Change** |
| `f` ← was `d` | SJK (CINA) CHUNG HUA PUJUT | **Change** |
| `g` ← was `e` | SMK DATO PERMAISURI | **Change** |

Affected rows:
| Line | Current | Fixed |
|------|---------|-------|
| 3788 | `P.219_N.74_219/74/01_c1` | `P.219_N.74_219/74/01_e1` |
| 3789 | `P.219_N.74_219/74/01_c2` | `P.219_N.74_219/74/01_e2` |
| 3790 | `P.219_N.74_219/74/01_c3` | `P.219_N.74_219/74/01_e3` |
| 3791 | `P.219_N.74_219/74/01_c4` | `P.219_N.74_219/74/01_e4` |
| 3792 | `P.219_N.74_219/74/01_c5` | `P.219_N.74_219/74/01_e5` |
| 3793 | `P.219_N.74_219/74/01_c6` | `P.219_N.74_219/74/01_e6` |
| 3794 | `P.219_N.74_219/74/01_c7` | `P.219_N.74_219/74/01_e7` |
| 3795 | `P.219_N.74_219/74/01_d1` | `P.219_N.74_219/74/01_f1` |
| 3796 | `P.219_N.74_219/74/01_d2` | `P.219_N.74_219/74/01_f2` |
| 3797 | `P.219_N.74_219/74/01_d3` | `P.219_N.74_219/74/01_f3` |
| 3798 | `P.219_N.74_219/74/01_d4` | `P.219_N.74_219/74/01_f4` |
| 3799 | `P.219_N.74_219/74/01_d5` | `P.219_N.74_219/74/01_f5` |
| 3800 | `P.219_N.74_219/74/01_d6` | `P.219_N.74_219/74/01_f6` |
| 3801 | `P.219_N.74_219/74/01_d7` | `P.219_N.74_219/74/01_f7` |
| 3802 | `P.219_N.74_219/74/01_d8` | `P.219_N.74_219/74/01_f8` |
| 3803 | `P.219_N.74_219/74/01_d9` | `P.219_N.74_219/74/01_f9` |
| 3804 | `P.219_N.74_219/74/01_d10` | `P.219_N.74_219/74/01_f10` |
| 3805 | `P.219_N.74_219/74/01_e1` | `P.219_N.74_219/74/01_g1` |
| 3806 | `P.219_N.74_219/74/01_e2` | `P.219_N.74_219/74/01_g2` |
| 3807 | `P.219_N.74_219/74/01_e3` | `P.219_N.74_219/74/01_g3` |
| 3808 | `P.219_N.74_219/74/01_e4` | `P.219_N.74_219/74/01_g4` |
| 3809 | `P.219_N.74_219/74/01_e5` | `P.219_N.74_219/74/01_g5` |

### B3. District `220/78/02` — 1 duplicate code (2 rows)

All 15 rows in district:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 4077 | `P.220_N.78_220/78/02_a1` | SEKOLAH KEBANGSAAN SG SETAPANG | 1 |
| 4078 | `P.220_N.78_220/78/02_b1` | RH. LALO SELEJAU | 1 |
| 4079 | `P.220_N.78_220/78/02_c1` | RH. JUGAH SG. BELASOI | 1 |
| 4080 | `P.220_N.78_220/78/02_d1` | RH. BLALANG ANAK ATOM SG. SENGKABANG | 1 |
| 4081 | `P.220_N.78_220/78/02_d2` | RH. BLALANG ANAK ATOM SG. SENGKABANG | 2 |
| 4082 | `P.220_N.78_220/78/02_e1` | RH. LANSAM SG. DABAI | 1 |
| 4083 | `P.220_N.78_220/78/02_f1` | SEKOLAH KEBANGSAAN POYUT | 1 |
| 4084 | `P.220_N.78_220/78/02_f2` | SEKOLAH KEBANGSAAN POYUT | 2 |
| 4085 | `P.220_N.78_220/78/02_g1` | SEKOLAH KEBANGSAAN RUMAH GUDANG | 1 |
| 4086 | `P.220_N.78_220/78/02_h1` ⚠️ | SEKOLAH KEBANGSAAN LONG LENEI | 1 |
| 4087 | `P.220_N.78_220/78/02_h2` | SEKOLAH KEBANGSAAN LONG LENEI | 2 |
| 4088 | `P.220_N.78_220/78/02_i1` | RH. ADANG SG. RIDAN | 1 |
| 4089 | `P.220_N.78_220/78/02_j1` | SEKOLAH KEBANGSAAN SG BRIT | 1 |
| 4090 | `P.220_N.78_220/78/02_j2` | SEKOLAH KEBANGSAAN SG BRIT | 2 |
| 4091 | `P.220_N.78_220/78/02_h1` ⚠️ | RH. RIDAB AK SELAT SG. PASIR | 1 |

**Root cause**: SK LONG LENEI and RH. RIDAB AK SELAT SG. PASIR both assigned suffix `h`. Channel 1 collides.

**Fix**: Reassign RH. RIDAB AK SELAT SG. PASIR from `h` → `k` (next available after `j`):
| Line | Current | Fixed |
|------|---------|-------|
| 4091 | `P.220_N.78_220/78/02_h1` | `P.220_N.78_220/78/02_k1` |

### B4. District `221/80/02` — 2 duplicate codes (4 rows)

All 9 rows in district:

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 4205 | `P.221_N.80_221/80/02_a1` ⚠️ | SEKOLAH KEBANGSAAN KUALA PENGANAN | 1 |
| 4206 | `P.221_N.80_221/80/02_a2` | SEKOLAH KEBANGSAAN KUALA PENGANAN | 2 |
| 4207 | `P.221_N.80_221/80/02_b1` ⚠️ | SEKOLAH KEBANGSAAN NANGA MERIT | 1 |
| 4208 | `P.221_N.80_221/80/02_c1` | RH. LUTA ENGKASING LUBAI | 1 |
| 4209 | `P.221_N.80_221/80/02_d1` | SEKOLAH KEBANGSAAN MENUANG | 1 |
| 4210 | `P.221_N.80_221/80/02_d2` | SEKOLAH KEBANGSAAN MENUANG | 2 |
| 4211 | `P.221_N.80_221/80/02_d3` | SEKOLAH KEBANGSAAN MENUANG | 3 |
| 4212 | `P.221_N.80_221/80/02_a1` ⚠️ | RH. DAMPIN AK LAYAN TERIMAH | 1 |
| 4213 | `P.221_N.80_221/80/02_b1` ⚠️ | RH. SLI AK MINGAN MENGARI, JALAN MEDAMIT | 1 |

**Root cause**: Two new centres were assigned suffix letters already used by earlier centres:
- `a` used by both SK KUALA PENGANAN and RH. DAMPIN AK LAYAN TERIMAH
- `b` used by both SK NANGA MERIT and RH. SLI AK MINGAN MENGARI

**Fix**: Reassign the later centres to next available letters (`e`, `f`):
| Line | Current | Fixed |
|------|---------|-------|
| 4212 | `P.221_N.80_221/80/02_a1` | `P.221_N.80_221/80/02_e1` |
| 4213 | `P.221_N.80_221/80/02_b1` | `P.221_N.80_221/80/02_f1` |

---

## Issue C: Genuine Duplicate Row — Same Code, Same Centre, Different Data (1 code, 2 rows)

### C1. District `213/58/08` — `P.213_N.58_213/58/08_a1`

| Line | UNIQUE CODE | Polling Centre | Channel | Ballots Issued | Valid Votes | Rejected | Unreturned |
|------|-------------|----------------|---------|----------------|-------------|----------|------------|
| 2921 | `P.213_N.58_213/58/08_a1` | SEKOLAH KEBANGSAAN KAMPUNG TEH | 1 | 270 | 269 | 1 | 0 |
| 2922 | `P.213_N.58_213/58/08_a1` | SEKOLAH KEBANGSAAN KAMPUNG TEH | 1 | 288 | 285 | 3 | 0 |

Candidate vote comparison:

| Field | Line 2921 | Line 2922 |
|-------|-----------|-----------|
| PH (PKR) ABDUL JALIL vote | 39 | 38 |
| GPS (PBB) HANIFAH HAJAR TAIB vote | 230 | 247 |
| Total valid votes | 269 | 285 |

**Root cause**: This is NOT a suffix collision — both rows have the same UNIQUE CODE, same Polling Centre, same channel, same district, same candidates. But the **vote data is different**. This is a genuine data duplication where one row is likely erroneous or they represent two different salurans that were both labeled channel 1.

**Recommended action**: This requires investigation against the raw PDF source data. Either:
1. One row is a data entry error and should be removed, OR
2. These are actually two channels (1 and 2) at the same centre, and the channel number needs correcting on one row

**⚠️ This cannot be resolved by suffix assignment alone. Needs PDF verification.**

---

## Issue D: Suffix Consistency Violations — Same Centre Gets Different Suffixes (4 districts, no duplicate codes)

These districts have no duplicate UNIQUE CODEs, but the same Polling Centre was assigned different suffix letters for different channels, violating the rule that suffix assignment must be uniform per centre within a district.

### D1. District `198/20/33` — SK TANAH PUTEH ⚠️

| Line | UNIQUE CODE | Polling Centre | Channel | Ballots Issued | Valid Votes | Rejected | Unreturned |
|------|-------------|----------------|---------|----------------|-------------|----------|------------|
| 1258 | `P.198_N.20_198/20/33_a1` | SEKOLAH KEBANGSAAN TANAH PUTEH | 1 | 169 | 168 | 1 | 0 |
| 1259 | `P.198_N.20_198/20/33_a2` | SEKOLAH KEBANGSAAN TANAH PUTEH | 2 | 159 | 157 | 2 | 0 |
| 1260 | `P.198_N.20_198/20/33_b1` | SEKOLAH KEBANGSAAN TANAH PUTEH | 1 | 208 | 208 | 0 | 0 |

**Issue**: Same centre has suffix `a` (ch 1–2) AND suffix `b` (ch 1). Lines 1258 and 1260 are both channel 1 at the same centre but with **different vote data** (169 vs 208 ballots). This is similar to Issue C1 — a potential genuine data duplication or mislabeled channel number.

**⚠️ Recommended action**: Investigate against raw PDF. Either:
1. Line 1260 is actually at a different polling centre that was mislabeled, OR
2. Line 1260 is actually channel 3 (mislabeled as channel 1), in which case it should be `_a3`, OR
3. One of lines 1258/1260 is erroneous

### D2. District `210/50/06` — SK NG JAGOI

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 2421 | `P.210_N.50_210/50/06_a1` | SEKOLAH KEBANGSAAN NG JAGOI | 1 |
| 2422 | `P.210_N.50_210/50/06_b2` | SEKOLAH KEBANGSAAN NG JAGOI | 2 |

**Issue**: Same centre got `a` for channel 1 and `b` for channel 2. Should both be `a`.

**Fix**:
| Line | Current | Fixed |
|------|---------|-------|
| 2422 | `P.210_N.50_210/50/06_b2` | `P.210_N.50_210/50/06_a2` |

### D3. District `217/70/01` — BLOK A & BLOK B, SJK (CINA) SEBIEW CHINESE

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 3544 | `P.217_N.70_217/70/01_b1` | BLOK A & BLOK B, SJK (CINA) SEBIEW CHINESE | 1 |
| 3545 | `P.217_N.70_217/70/01_c2` | BLOK A & BLOK B, SJK (CINA) SEBIEW CHINESE | 2 |

**Issue**: Same centre got `b` for channel 1 and `c` for channel 2. Should both be `b`.

**Fix**:
| Line | Current | Fixed |
|------|---------|-------|
| 3545 | `P.217_N.70_217/70/01_c2` | `P.217_N.70_217/70/01_b2` |

### D4. District `220/78/03` — SK LONG TUNGAN

| Line | UNIQUE CODE | Polling Centre | Channel |
|------|-------------|----------------|---------|
| 4097 | `P.220_N.78_220/78/03_d1` | SEKOLAH KEBANGSAAN LONG TUNGAN | 1 |
| 4098 | `P.220_N.78_220/78/03_e2` | SEKOLAH KEBANGSAAN LONG TUNGAN | 2 |

**Issue**: Same centre got `d` for channel 1 and `e` for channel 2. Should both be `d`.

**Fix**:
| Line | Current | Fixed |
|------|---------|-------|
| 4098 | `P.220_N.78_220/78/03_e2` | `P.220_N.78_220/78/03_d2` |

---

## Issue E: Suffix Convention Inconsistency — `_1a` vs `_a1` Format (11 rows)

Most of the file uses the `_<letter><channel>` format (e.g. `_a1`, `_b2`) where the suffix letter precedes the channel number.

However, **11 rows** in P.203 (N.33 and N.34) use the reversed `_<channel><letter>` format (e.g. `_1a`, `_1b`), matching the AGENTS.md example convention.

| Line | UNIQUE CODE | Polling District Code | Polling Centre | Channel |
|------|-------------|-----------------------|----------------|---------|
| 1746 | `P.203_N.33_203/33/10_1a` | 203/33/10 | SEKOLAH KEBANGSAAN NG MENJUAU | 1 |
| 1747 | `P.203_N.33_203/33/10_1b` | 203/33/10 | RH GAIT, NG SEBEMBAN | 1 |
| 1770 | `P.203_N.34_203/34/04_1a` | 203/34/04 | SEKOLAH KEBANGSAAN NANGA KESIT | 1 |
| 1771 | `P.203_N.34_203/34/04_2a` | 203/34/04 | SEKOLAH KEBANGSAAN NANGA KESIT | 2 |
| 1772 | `P.203_N.34_203/34/04_1b` | 203/34/04 | RH. LIMPENG LEPONG MAWANG KESIT | 1 |
| 1775 | `P.203_N.34_203/34/07_1a` | 203/34/07 | TADIKA KEMAS KAONG ULU | 1 |
| 1776 | `P.203_N.34_203/34/07_1b` | 203/34/07 | PUSAT SUMBER NYEMUNGAN SIMPANG | 1 |
| 1787 | `P.203_N.34_203/34/14_1a` | 203/34/14 | TADIKA KEMAS NG. JELA | 1 |
| 1788 | `P.203_N.34_203/34/14_1b` | 203/34/14 | NG. TELAUS SPS | 1 |
| 1795 | `P.203_N.34_203/34/20_1a` | 203/34/20 | RH JELI, NG SEMPILI | 1 |
| 1796 | `P.203_N.34_203/34/20_1b` | 203/34/20 | RH. BAKAR, NG SEMPILI | 1 |

**Note**: These codes are all unique — there are no duplicate UNIQUE CODE values among them. However, the format inconsistency means the VOTING CHANNEL NUMBER extraction algorithm fails on these rows (regex extracts empty string for the channel since there are no trailing digits after letters).

**Recommended action**: Normalize to the dominant `_<letter><channel>` convention used in the rest of the file:
- `_1a` → `_a1`
- `_1b` → `_b1`
- `_2a` → `_a2`

---

## VOTING CHANNEL NUMBER vs UNIQUE CODE Cross-Check

The 11 `_<channel><letter>` convention rows from Issue E are the only mismatches. All other 4306 rows have the trailing number in UNIQUE CODE matching the VOTING CHANNEL NUMBER column. ✅

---

## Complete Fix Summary

### Fix statistics

| Category | Codes | Rows to modify | Type |
|----------|-------|----------------|------|
| A: No suffix, need new suffix | 3 | 11 (3 districts) | Add suffixes to all rows in affected districts |
| B: Wrong suffix causing collision | 18 | 35 (4 districts) | Reassign suffix letters |
| C: Genuine data duplicate | 1 | 2 | **Needs PDF investigation** |
| D: Suffix inconsistency (no dup) | 4 | 3 | Correct suffix letter |
| E: Convention `_1a` vs `_a1` | — | 11 | Normalize format |
| **Total** | **22** | **62** | |

### Rows that MUST be changed to resolve duplicate UNIQUE CODE violations (Issues A + B)

After these fixes, all 22 duplicate code violations will be resolved (except Issue C which requires manual investigation):

| # | Line | Current UNIQUE CODE | Fixed UNIQUE CODE | Reason |
|---|------|--------------------|--------------------|--------|
| 1 | 543 | `P.195_N.10_195/10/02_a1` | `P.195_N.10_195/10/02_c1` | B1: TABUAN ULU `a`→`c` |
| 2 | 544 | `P.195_N.10_195/10/02_a2` | `P.195_N.10_195/10/02_c2` | B1: TABUAN ULU `a`→`c` |
| 3 | 545 | `P.195_N.10_195/10/02_a3` | `P.195_N.10_195/10/02_c3` | B1: TABUAN ULU `a`→`c` |
| 4 | 546 | `P.195_N.10_195/10/02_a4` | `P.195_N.10_195/10/02_c4` | B1: TABUAN ULU `a`→`c` |
| 5 | 547 | `P.195_N.10_195/10/02_a5` | `P.195_N.10_195/10/02_c5` | B1: TABUAN ULU `a`→`c` |
| 6 | 548 | `P.195_N.10_195/10/02_a6` | `P.195_N.10_195/10/02_c6` | B1: TABUAN ULU `a`→`c` |
| 7 | 549 | `P.195_N.10_195/10/02_a7` | `P.195_N.10_195/10/02_c7` | B1: TABUAN ULU `a`→`c` |
| 8 | 550 | `P.195_N.10_195/10/02_a8` | `P.195_N.10_195/10/02_c8` | B1: TABUAN ULU `a`→`c` |
| 9 | 551 | `P.195_N.10_195/10/02_a9` | `P.195_N.10_195/10/02_c9` | B1: TABUAN ULU `a`→`c` |
| 10 | 1177 | `P.198_N.19_198/19/29_1` | `P.198_N.19_198/19/29_a1` | A1: new suffix |
| 11 | 1178 | `P.198_N.19_198/19/29_1` | `P.198_N.19_198/19/29_b1` | A1: new suffix |
| 12 | 1218 | `P.198_N.20_198/20/13_1` | `P.198_N.20_198/20/13_a1` | A2: new suffix |
| 13 | 1219 | `P.198_N.20_198/20/13_2` | `P.198_N.20_198/20/13_a2` | A2: uniformity |
| 14 | 1220 | `P.198_N.20_198/20/13_1` | `P.198_N.20_198/20/13_b1` | A2: new suffix |
| 15 | 2201 | `P.208_N.45_208/45/05_1` | `P.208_N.45_208/45/05_a1` | A3: new suffix |
| 16 | 2202 | `P.208_N.45_208/45/05_2` | `P.208_N.45_208/45/05_a2` | A3: uniformity |
| 17 | 2203 | `P.208_N.45_208/45/05_3` | `P.208_N.45_208/45/05_a3` | A3: uniformity |
| 18 | 2204 | `P.208_N.45_208/45/05_4` | `P.208_N.45_208/45/05_a4` | A3: uniformity |
| 19 | 2205 | `P.208_N.45_208/45/05_1` | `P.208_N.45_208/45/05_b1` | A3: new suffix |
| 20 | 3788 | `P.219_N.74_219/74/01_c1` | `P.219_N.74_219/74/01_e1` | B2: PUJUT CORNER `c`→`e` |
| 21 | 3789 | `P.219_N.74_219/74/01_c2` | `P.219_N.74_219/74/01_e2` | B2: PUJUT CORNER `c`→`e` |
| 22 | 3790 | `P.219_N.74_219/74/01_c3` | `P.219_N.74_219/74/01_e3` | B2: PUJUT CORNER `c`→`e` |
| 23 | 3791 | `P.219_N.74_219/74/01_c4` | `P.219_N.74_219/74/01_e4` | B2: PUJUT CORNER `c`→`e` |
| 24 | 3792 | `P.219_N.74_219/74/01_c5` | `P.219_N.74_219/74/01_e5` | B2: PUJUT CORNER `c`→`e` |
| 25 | 3793 | `P.219_N.74_219/74/01_c6` | `P.219_N.74_219/74/01_e6` | B2: PUJUT CORNER `c`→`e` |
| 26 | 3794 | `P.219_N.74_219/74/01_c7` | `P.219_N.74_219/74/01_e7` | B2: PUJUT CORNER `c`→`e` |
| 27 | 3795 | `P.219_N.74_219/74/01_d1` | `P.219_N.74_219/74/01_f1` | B2: CHUNG HUA PUJUT `d`→`f` |
| 28 | 3796 | `P.219_N.74_219/74/01_d2` | `P.219_N.74_219/74/01_f2` | B2: CHUNG HUA PUJUT `d`→`f` |
| 29 | 3797 | `P.219_N.74_219/74/01_d3` | `P.219_N.74_219/74/01_f3` | B2: CHUNG HUA PUJUT `d`→`f` |
| 30 | 3798 | `P.219_N.74_219/74/01_d4` | `P.219_N.74_219/74/01_f4` | B2: CHUNG HUA PUJUT `d`→`f` |
| 31 | 3799 | `P.219_N.74_219/74/01_d5` | `P.219_N.74_219/74/01_f5` | B2: CHUNG HUA PUJUT `d`→`f` |
| 32 | 3800 | `P.219_N.74_219/74/01_d6` | `P.219_N.74_219/74/01_f6` | B2: CHUNG HUA PUJUT `d`→`f` |
| 33 | 3801 | `P.219_N.74_219/74/01_d7` | `P.219_N.74_219/74/01_f7` | B2: CHUNG HUA PUJUT `d`→`f` |
| 34 | 3802 | `P.219_N.74_219/74/01_d8` | `P.219_N.74_219/74/01_f8` | B2: CHUNG HUA PUJUT `d`→`f` |
| 35 | 3803 | `P.219_N.74_219/74/01_d9` | `P.219_N.74_219/74/01_f9` | B2: CHUNG HUA PUJUT `d`→`f` |
| 36 | 3804 | `P.219_N.74_219/74/01_d10` | `P.219_N.74_219/74/01_f10` | B2: CHUNG HUA PUJUT `d`→`f` |
| 37 | 3805 | `P.219_N.74_219/74/01_e1` | `P.219_N.74_219/74/01_g1` | B2: DATO PERMAISURI `e`→`g` |
| 38 | 3806 | `P.219_N.74_219/74/01_e2` | `P.219_N.74_219/74/01_g2` | B2: DATO PERMAISURI `e`→`g` |
| 39 | 3807 | `P.219_N.74_219/74/01_e3` | `P.219_N.74_219/74/01_g3` | B2: DATO PERMAISURI `e`→`g` |
| 40 | 3808 | `P.219_N.74_219/74/01_e4` | `P.219_N.74_219/74/01_g4` | B2: DATO PERMAISURI `e`→`g` |
| 41 | 3809 | `P.219_N.74_219/74/01_e5` | `P.219_N.74_219/74/01_g5` | B2: DATO PERMAISURI `e`→`g` |
| 42 | 4091 | `P.220_N.78_220/78/02_h1` | `P.220_N.78_220/78/02_k1` | B3: RIDAB `h`→`k` |
| 43 | 4212 | `P.221_N.80_221/80/02_a1` | `P.221_N.80_221/80/02_e1` | B4: DAMPIN `a`→`e` |
| 44 | 4213 | `P.221_N.80_221/80/02_b1` | `P.221_N.80_221/80/02_f1` | B4: SLI `b`→`f` |

### Open items requiring manual investigation

| # | Line(s) | Issue | Action needed |
|---|---------|-------|---------------|
| 1 | 2921–2922 | `P.213_N.58_213/58/08_a1` genuine duplicate: same centre, same channel, different vote data (270 vs 288 ballots) | Verify against raw PDF |
| 2 | 1258–1260 | District `198/20/33`: SK TANAH PUTEH has `_a1`(ch1), `_a2`(ch2), `_b1`(ch1) — two channel 1 rows at same centre with different vote data (169 vs 208 ballots) | Verify against raw PDF |

### Normalization items (non-blocking, no duplicates)

| # | Lines | Issue | Fix |
|---|-------|-------|-----|
| 1 | 2422 | District `210/50/06`: SK NG JAGOI `_b2` → `_a2` | Suffix consistency |
| 2 | 3545 | District `217/70/01`: SJK SEBIEW CHINESE `_c2` → `_b2` | Suffix consistency |
| 3 | 4098 | District `220/78/03`: SK LONG TUNGAN `_e2` → `_d2` | Suffix consistency |
| 4 | 11 rows in P.203 | `_1a`/`_1b`/`_2a` format → `_a1`/`_b1`/`_a2` | Convention normalization |