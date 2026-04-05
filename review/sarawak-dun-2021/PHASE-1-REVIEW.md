# PHASE 1 REVIEW: UNIQUE CODE Uniqueness Check

## Summary

| Metric | Value |
| --- | --- |
| Total data rows | 3748 |
| Empty UNIQUE CODE rows | 0 |
| Distinct non-empty UNIQUE CODEs | 2406 |
| UNIQUE CODEs appearing exactly once | 1883 |
| **Duplicate UNIQUE CODEs** | **523** |
| **Total rows affected by duplicates** | **1865** |
| Existing suffixed IDs (ending a-z) | 0 |

### ❌ Result: FAIL — Duplicates Found

**523** distinct UNIQUE CODE values are duplicated across **1865** rows.
These need to be disambiguated with letter suffixes (a, b, c, ...) per the suffix rule.

## Existing Suffixed IDs

No existing suffixed IDs found. All suffix assignments need to be applied fresh.

## Duplicate Details

### Top Duplicates by Occurrence Count

| # | UNIQUE CODE | Occurrences | Distinct Polling Centres |
| --- | --- | --- | --- |
| 1 | `P.215_N.61_215/61/04_1` | 18 | 18 |
| 2 | `P.216_N.66_216/66/02_1` | 18 | 18 |
| 3 | `P.216_N.64_216/64/02_1` | 17 | 17 |
| 4 | `P.220_N.76_220/76/04_1` | 16 | 16 |
| 5 | `P.216_N.64_216/64/05_1` | 16 | 16 |
| 6 | `P.216_N.64_216/64/03_1` | 15 | 15 |
| 7 | `P.216_N.64_216/64/04_1` | 14 | 14 |
| 8 | `P.216_N.65_216/65/02_1` | 14 | 14 |
| 9 | `P.215_N.63_215/63/04_1` | 13 | 13 |
| 10 | `P.215_N.62_215/62/02_1` | 12 | 12 |
| 11 | `P.216_N.66_216/66/03_1` | 12 | 12 |
| 12 | `P.215_N.61_215/61/01_1` | 12 | 12 |
| 13 | `P.197_N.16_197/16/01_1` | 12 | 12 |
| 14 | `P.220_N.78_220/78/02_1` | 11 | 11 |
| 15 | `P.220_N.78_220/78/03_1` | 11 | 11 |
| 16 | `P.215_N.62_215/62/04_1` | 11 | 11 |
| 17 | `P.215_N.62_215/62/06_1` | 10 | 10 |
| 18 | `P.199_N.21_199/21/04_1` | 10 | 10 |
| 19 | `P.199_N.22_199/22/01_1` | 10 | 10 |
| 20 | `P.220_N.77_220/77/06_1` | 9 | 9 |
| 21 | `P.216_N.65_216/65/04_1` | 9 | 9 |
| 22 | `P.219_N.75_219/75/01_1` | 9 | 9 |
| 23 | `P.220_N.76_220/76/01_1` | 9 | 9 |
| 24 | `P.196_N.14_196/14/03_1` | 9 | 9 |
| 25 | `P.205_N.39_205/39/08_1` | 9 | 9 |
| 26 | `P.199_N.21_199/21/03_1` | 8 | 8 |
| 27 | `P.205_N.40_205/40/01_1` | 8 | 8 |
| 28 | `P.219_N.75_219/75/01_2` | 8 | 8 |
| 29 | `P.219_N.75_219/75/01_3` | 8 | 8 |
| 30 | `P.199_N.22_199/22/04_1` | 8 | 8 |
| ... | *(and 493 more)* | | |

### Distribution of Duplicate Counts

| Occurrences | Number of UNIQUE CODEs |
| --- | --- |
| 2 | 267 |
| 3 | 85 |
| 4 | 62 |
| 5 | 25 |
| 6 | 27 |
| 7 | 21 |
| 8 | 11 |
| 9 | 6 |
| 10 | 3 |
| 11 | 3 |
| 12 | 4 |
| 13 | 1 |
| 14 | 2 |
| 15 | 1 |
| 16 | 2 |
| 17 | 1 |
| 18 | 2 |

### Recommended Suffix Assignments

Below are all duplicate UNIQUE CODEs with the recommended suffix assignment.
Each unique Polling Centre for a given duplicate ID gets a distinct letter suffix (a, b, c, ...)
based on order of first appearance in the file.

#### 1. `P.192_N.02_192/02/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 36 | `P.192_N.02_192/02/00_1` | BILIK REKREASI IBU PEJABAT POLIS DAERAH BAU | a | `P.192_N.02_192/02/00_1a` |
| 37 | `P.192_N.02_192/02/00_1` | DEWAN SERBAGUNA KEM PLKN PUNCAK PERMAI | b | `P.192_N.02_192/02/00_1b` |

#### 2. `P.192_N.02_192/02/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 54 | `P.192_N.02_192/02/07_1` | BALAI RAYA KPG. SKIAT BARU | a | `P.192_N.02_192/02/07_1a` |
| 55 | `P.192_N.02_192/02/07_1` | SEKOLAH KEBANGSAAN BAU | b | `P.192_N.02_192/02/07_1b` |

#### 3. `P.193_N.03_193/03/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 84 | `P.193_N.03_193/03/00_1` | DEWAN ULS SAMPADI | a | `P.193_N.03_193/03/00_1a` |
| 85 | `P.193_N.03_193/03/00_1` | DEWAN MASYARAKAT LUNDU | b | `P.193_N.03_193/03/00_1b` |

#### 4. `P.193_N.03_193/03/15_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 117 | `P.193_N.03_193/03/15_1` | BALAI RAYA KPG. RAMBUNGAN | a | `P.193_N.03_193/03/15_1a` |
| 118 | `P.193_N.03_193/03/15_1` | DEWAN MASYARAKAT KAMPUNG SUNGAI BELIAN | b | `P.193_N.03_193/03/15_1b` |

#### 5. `P.193_N.04_193/04/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 122 | `P.193_N.04_193/04/01_1` | SEKOLAH KEBANGSAAN TELAGA AIR | a | `P.193_N.04_193/04/01_1a` |
| 125 | `P.193_N.04_193/04/01_1` | DEWAN SELANG LAUT | b | `P.193_N.04_193/04/01_1b` |

#### 6. `P.193_N.04_193/04/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 134 | `P.193_N.04_193/04/06_1` | SEKOLAH KEBANGSAAN BUNTAL | a | `P.193_N.04_193/04/06_1a` |
| 138 | `P.193_N.04_193/04/06_1` | DEWAN SERBAGUNA SG.LUMUT | b | `P.193_N.04_193/04/06_1b` |

#### 7. `P.193_N.04_193/04/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 135 | `P.193_N.04_193/04/06_2` | SEKOLAH KEBANGSAAN BUNTAL | a | `P.193_N.04_193/04/06_2a` |
| 139 | `P.193_N.04_193/04/06_2` | DEWAN SERBAGUNA SG.LUMUT | b | `P.193_N.04_193/04/06_2b` |

#### 8. `P.193_N.04_193/04/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 140 | `P.193_N.04_193/04/07_1` | SEKOLAH KEBANGSAAN RAMPANGI | a | `P.193_N.04_193/04/07_1a` |
| 144 | `P.193_N.04_193/04/07_1` | SEKOLAH KEBANGSAAN BANDAR BARU SAMARIANG | b | `P.193_N.04_193/04/07_1b` |

#### 9. `P.193_N.04_193/04/07_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 141 | `P.193_N.04_193/04/07_2` | SEKOLAH KEBANGSAAN RAMPANGI | a | `P.193_N.04_193/04/07_2a` |
| 145 | `P.193_N.04_193/04/07_2` | SEKOLAH KEBANGSAAN BANDAR BARU SAMARIANG | b | `P.193_N.04_193/04/07_2b` |

#### 10. `P.193_N.04_193/04/07_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 142 | `P.193_N.04_193/04/07_3` | SEKOLAH KEBANGSAAN RAMPANGI | a | `P.193_N.04_193/04/07_3a` |
| 146 | `P.193_N.04_193/04/07_3` | SEKOLAH KEBANGSAAN BANDAR BARU SAMARIANG | b | `P.193_N.04_193/04/07_3b` |

#### 11. `P.193_N.04_193/04/07_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 143 | `P.193_N.04_193/04/07_4` | SEKOLAH KEBANGSAAN RAMPANGI | a | `P.193_N.04_193/04/07_4a` |
| 147 | `P.193_N.04_193/04/07_4` | SEKOLAH KEBANGSAAN BANDAR BARU SAMARIANG | b | `P.193_N.04_193/04/07_4b` |

#### 12. `P.193_N.04_193/04/09_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 152 | `P.193_N.04_193/04/09_1` | SEKOLAH KEBANGSAAN GERSIK | a | `P.193_N.04_193/04/09_1a` |
| 157 | `P.193_N.04_193/04/09_1` | DEWAN SERBAGUNA PANGLIMA SEMAN LAMA | b | `P.193_N.04_193/04/09_1b` |

#### 13. `P.193_N.04_193/04/09_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 153 | `P.193_N.04_193/04/09_2` | SEKOLAH KEBANGSAAN GERSIK | a | `P.193_N.04_193/04/09_2a` |
| 158 | `P.193_N.04_193/04/09_2` | DEWAN SERBAGUNA PANGLIMA SEMAN LAMA | b | `P.193_N.04_193/04/09_2b` |

#### 14. `P.193_N.04_193/04/09_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 154 | `P.193_N.04_193/04/09_3` | SEKOLAH KEBANGSAAN GERSIK | a | `P.193_N.04_193/04/09_3a` |
| 159 | `P.193_N.04_193/04/09_3` | DEWAN SERBAGUNA PANGLIMA SEMAN LAMA | b | `P.193_N.04_193/04/09_3b` |

#### 15. `P.193_N.04_193/04/11_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 162 | `P.193_N.04_193/04/11_1` | SEKOLAH KEBANGSAAN PULO | a | `P.193_N.04_193/04/11_1a` |
| 164 | `P.193_N.04_193/04/11_1` | SEKOLAH AGAMA AL FALAH AL ISLAMIYYAH KPG. PULO | b | `P.193_N.04_193/04/11_1b` |

#### 16. `P.193_N.04_193/04/11_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 163 | `P.193_N.04_193/04/11_2` | SEKOLAH KEBANGSAAN PULO | a | `P.193_N.04_193/04/11_2a` |
| 165 | `P.193_N.04_193/04/11_2` | SEKOLAH AGAMA AL FALAH AL ISLAMIYYAH KPG. PULO | b | `P.193_N.04_193/04/11_2b` |

#### 17. `P.193_N.05_193/05/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 185 | `P.193_N.05_193/05/04_1` | SEKOLAH KEBANGSAAN BAKO | a | `P.193_N.05_193/05/04_1a` |
| 190 | `P.193_N.05_193/05/04_1` | DEWAN HIJRAH 2008 | b | `P.193_N.05_193/05/04_1b` |

#### 18. `P.193_N.05_193/05/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 186 | `P.193_N.05_193/05/04_2` | SEKOLAH KEBANGSAAN BAKO | a | `P.193_N.05_193/05/04_2a` |
| 191 | `P.193_N.05_193/05/04_2` | DEWAN HIJRAH 2008 | b | `P.193_N.05_193/05/04_2b` |

#### 19. `P.193_N.05_193/05/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 192 | `P.193_N.05_193/05/05_1` | SEKOLAH KEBANGSAAN PAJAR SEJINGKAT | a | `P.193_N.05_193/05/05_1a` |
| 195 | `P.193_N.05_193/05/05_1` | TADIKA KEMAS TAMAN SEPAKAT JAYA | b | `P.193_N.05_193/05/05_1b` |

#### 20. `P.193_N.05_193/05/09_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 202 | `P.193_N.05_193/05/09_1` | SEKOLAH KEBANGSAAN TABUAN HILIR | a | `P.193_N.05_193/05/09_1a` |
| 209 | `P.193_N.05_193/05/09_1` | DEWAN DATO SRI WAN JUNAIDI KPG.TABUAN LOT | b | `P.193_N.05_193/05/09_1b` |

#### 21. `P.193_N.05_193/05/09_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 203 | `P.193_N.05_193/05/09_2` | SEKOLAH KEBANGSAAN TABUAN HILIR | a | `P.193_N.05_193/05/09_2a` |
| 210 | `P.193_N.05_193/05/09_2` | DEWAN DATO SRI WAN JUNAIDI KPG.TABUAN LOT | b | `P.193_N.05_193/05/09_2b` |

#### 22. `P.193_N.05_193/05/09_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 204 | `P.193_N.05_193/05/09_3` | SEKOLAH KEBANGSAAN TABUAN HILIR | a | `P.193_N.05_193/05/09_3a` |
| 211 | `P.193_N.05_193/05/09_3` | DEWAN DATO SRI WAN JUNAIDI KPG.TABUAN LOT | b | `P.193_N.05_193/05/09_3b` |

#### 23. `P.194_N.06_194/06/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 220 | `P.194_N.06_194/06/01_1` | SEKOLAH KEBANGSAAN RAKYAT TUPONG | a | `P.194_N.06_194/06/01_1a` |
| 224 | `P.194_N.06_194/06/01_1` | TABIKA KEMAS KPG. TUPONG TENGAH | b | `P.194_N.06_194/06/01_1b` |

#### 24. `P.194_N.06_194/06/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 221 | `P.194_N.06_194/06/01_2` | SEKOLAH KEBANGSAAN RAKYAT TUPONG | a | `P.194_N.06_194/06/01_2a` |
| 225 | `P.194_N.06_194/06/01_2` | TABIKA KEMAS KPG. TUPONG TENGAH | b | `P.194_N.06_194/06/01_2b` |

#### 25. `P.194_N.06_194/06/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 226 | `P.194_N.06_194/06/02_1` | SEKOLAH MENENGAH KEBANGSAAN TUNKU ABDUL RAHMAN | a | `P.194_N.06_194/06/02_1a` |
| 234 | `P.194_N.06_194/06/02_1` | DEWAN SERBAGUNA KPG. PINANG JAWA | b | `P.194_N.06_194/06/02_1b` |

#### 26. `P.194_N.06_194/06/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 227 | `P.194_N.06_194/06/02_2` | SEKOLAH MENENGAH KEBANGSAAN TUNKU ABDUL RAHMAN | a | `P.194_N.06_194/06/02_2a` |
| 235 | `P.194_N.06_194/06/02_2` | DEWAN SERBAGUNA KPG. PINANG JAWA | b | `P.194_N.06_194/06/02_2b` |

#### 27. `P.194_N.06_194/06/03_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 236 | `P.194_N.06_194/06/03_1` | SEKOLAH JENIS KEBANGSAAN SG. TENGAH | a | `P.194_N.06_194/06/03_1a` |
| 238 | `P.194_N.06_194/06/03_1` | DEWAN PERPADUAN KPG. SAGAH | b | `P.194_N.06_194/06/03_1b` |
| 239 | `P.194_N.06_194/06/03_1` | DEWAN SERBAGUNA DATUK AMAR DR SULAIMAN DAUD KPG. KOLONG DUA | c | `P.194_N.06_194/06/03_1c` |
| 241 | `P.194_N.06_194/06/03_1` | SEKOLAH KEBANGSAAN TAN SRI DATUK HAJI MOHAMED | d | `P.194_N.06_194/06/03_1d` |
| 247 | `P.194_N.06_194/06/03_1` | DEWAN SERBAGUNA TAMAN MALIHAH | e | `P.194_N.06_194/06/03_1e` |

#### 28. `P.194_N.06_194/06/03_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 237 | `P.194_N.06_194/06/03_2` | SEKOLAH JENIS KEBANGSAAN SG. TENGAH | a | `P.194_N.06_194/06/03_2a` |
| 240 | `P.194_N.06_194/06/03_2` | DEWAN SERBAGUNA DATUK AMAR DR SULAIMAN DAUD KPG. KOLONG DUA | b | `P.194_N.06_194/06/03_2b` |
| 242 | `P.194_N.06_194/06/03_2` | SEKOLAH KEBANGSAAN TAN SRI DATUK HAJI MOHAMED | c | `P.194_N.06_194/06/03_2c` |
| 248 | `P.194_N.06_194/06/03_2` | DEWAN SERBAGUNA TAMAN MALIHAH | d | `P.194_N.06_194/06/03_2d` |

#### 29. `P.194_N.06_194/06/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 249 | `P.194_N.06_194/06/04_1` | SEKOLAH KEBANGSAAN GITA | a | `P.194_N.06_194/06/04_1a` |
| 251 | `P.194_N.06_194/06/04_1` | FOYER MASJID TAMAN HUSSEIN / RAHMAT | b | `P.194_N.06_194/06/04_1b` |

#### 30. `P.194_N.06_194/06/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 250 | `P.194_N.06_194/06/04_2` | SEKOLAH KEBANGSAAN GITA | a | `P.194_N.06_194/06/04_2a` |
| 252 | `P.194_N.06_194/06/04_2` | FOYER MASJID TAMAN HUSSEIN / RAHMAT | b | `P.194_N.06_194/06/04_2b` |

#### 31. `P.194_N.06_194/06/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 256 | `P.194_N.06_194/06/06_1` | SEKOLAH KEBANGSAAN TAN SRI DATUK DR. HJ. SULAIMAN DAUD | a | `P.194_N.06_194/06/06_1a` |
| 259 | `P.194_N.06_194/06/06_1` | SEKOLAH MENENGAH KEBANGSAAN PETRA JAYA | b | `P.194_N.06_194/06/06_1b` |

#### 32. `P.194_N.06_194/06/08_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 262 | `P.194_N.06_194/06/08_1` | SEKOLAH MENENGAH KEBANGSAAN MATANG HILIR | a | `P.194_N.06_194/06/08_1a` |
| 265 | `P.194_N.06_194/06/08_1` | SEKOLAH JENIS KEBANGSAAN CHUNG HUA BATU 7 | b | `P.194_N.06_194/06/08_1b` |
| 268 | `P.194_N.06_194/06/08_1` | SEKOLAH KEBANGSAAN MATANG | c | `P.194_N.06_194/06/08_1c` |
| 269 | `P.194_N.06_194/06/08_1` | DEWAN PENDIDIKAN DAN SOSIAL KPG. MATANG BATU 10 JLN. MATANG | d | `P.194_N.06_194/06/08_1d` |
| 273 | `P.194_N.06_194/06/08_1` | SEKOLAH KEBANGSAAN MATANG JAYA | e | `P.194_N.06_194/06/08_1e` |
| 279 | `P.194_N.06_194/06/08_1` | SEKOLAH KEBANGSAAN PETRA JAYA | f | `P.194_N.06_194/06/08_1f` |
| 282 | `P.194_N.06_194/06/08_1` | DEWAN SERBAGUNA TAMAN HENG GUAN | g | `P.194_N.06_194/06/08_1g` |

#### 33. `P.194_N.06_194/06/08_2` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 263 | `P.194_N.06_194/06/08_2` | SEKOLAH MENENGAH KEBANGSAAN MATANG HILIR | a | `P.194_N.06_194/06/08_2a` |
| 266 | `P.194_N.06_194/06/08_2` | SEKOLAH JENIS KEBANGSAAN CHUNG HUA BATU 7 | b | `P.194_N.06_194/06/08_2b` |
| 270 | `P.194_N.06_194/06/08_2` | DEWAN PENDIDIKAN DAN SOSIAL KPG. MATANG BATU 10 JLN. MATANG | c | `P.194_N.06_194/06/08_2c` |
| 274 | `P.194_N.06_194/06/08_2` | SEKOLAH KEBANGSAAN MATANG JAYA | d | `P.194_N.06_194/06/08_2d` |
| 280 | `P.194_N.06_194/06/08_2` | SEKOLAH KEBANGSAAN PETRA JAYA | e | `P.194_N.06_194/06/08_2e` |

#### 34. `P.194_N.06_194/06/08_3` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 264 | `P.194_N.06_194/06/08_3` | SEKOLAH MENENGAH KEBANGSAAN MATANG HILIR | a | `P.194_N.06_194/06/08_3a` |
| 267 | `P.194_N.06_194/06/08_3` | SEKOLAH JENIS KEBANGSAAN CHUNG HUA BATU 7 | b | `P.194_N.06_194/06/08_3b` |
| 271 | `P.194_N.06_194/06/08_3` | DEWAN PENDIDIKAN DAN SOSIAL KPG. MATANG BATU 10 JLN. MATANG | c | `P.194_N.06_194/06/08_3c` |
| 275 | `P.194_N.06_194/06/08_3` | SEKOLAH KEBANGSAAN MATANG JAYA | d | `P.194_N.06_194/06/08_3d` |
| 281 | `P.194_N.06_194/06/08_3` | SEKOLAH KEBANGSAAN PETRA JAYA | e | `P.194_N.06_194/06/08_3e` |

#### 35. `P.194_N.06_194/06/08_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 272 | `P.194_N.06_194/06/08_4` | DEWAN PENDIDIKAN DAN SOSIAL KPG. MATANG BATU 10 JLN. MATANG | a | `P.194_N.06_194/06/08_4a` |
| 276 | `P.194_N.06_194/06/08_4` | SEKOLAH KEBANGSAAN MATANG JAYA | b | `P.194_N.06_194/06/08_4b` |

#### 36. `P.194_N.07_194/07/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 288 | `P.194_N.07_194/07/01_1` | SEKOLAH KEBANGSAAN SEMARIANG | a | `P.194_N.07_194/07/01_1a` |
| 295 | `P.194_N.07_194/07/01_1` | SURAU DARUL AMIN SAMARIANG JAYA, FASA 2, JALAN BENTARA | b | `P.194_N.07_194/07/01_1b` |

#### 37. `P.194_N.07_194/07/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 324 | `P.194_N.07_194/07/06_1` | DEWAN MASYARAKAT PETRA JAYA (KPG. TUNKU) | a | `P.194_N.07_194/07/06_1a` |
| 328 | `P.194_N.07_194/07/06_1` | SEKOLAH KEBANGSAAN RANCANGAN PERUMAHAN RAKYAT 'RPR' | b | `P.194_N.07_194/07/06_1b` |

#### 38. `P.194_N.07_194/07/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 325 | `P.194_N.07_194/07/06_2` | DEWAN MASYARAKAT PETRA JAYA (KPG. TUNKU) | a | `P.194_N.07_194/07/06_2a` |
| 329 | `P.194_N.07_194/07/06_2` | SEKOLAH KEBANGSAAN RANCANGAN PERUMAHAN RAKYAT 'RPR' | b | `P.194_N.07_194/07/06_2b` |

#### 39. `P.194_N.07_194/07/06_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 326 | `P.194_N.07_194/07/06_3` | DEWAN MASYARAKAT PETRA JAYA (KPG. TUNKU) | a | `P.194_N.07_194/07/06_3a` |
| 330 | `P.194_N.07_194/07/06_3` | SEKOLAH KEBANGSAAN RANCANGAN PERUMAHAN RAKYAT 'RPR' | b | `P.194_N.07_194/07/06_3b` |

#### 40. `P.194_N.08_194/08/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 341 | `P.194_N.08_194/08/01_1` | TABIKA KEMAS KPG. SG. MAONG | a | `P.194_N.08_194/08/01_1a` |
| 342 | `P.194_N.08_194/08/01_1` | DEWAN SERBAGUNA KPG. SEGEDUP | b | `P.194_N.08_194/08/01_1b` |

#### 41. `P.194_N.08_194/08/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 358 | `P.194_N.08_194/08/07_1` | SEKOLAH JENIS KEBANGSAAN (C) CHUNG HUA NO.4 KUCHING | a | `P.194_N.08_194/08/07_1a` |
| 364 | `P.194_N.08_194/08/07_1` | BANGUNAN LAMA MASJID DARUL IBADAH KPG. KUDEI LAMA | b | `P.194_N.08_194/08/07_1b` |

#### 42. `P.194_N.08_194/08/07_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 359 | `P.194_N.08_194/08/07_2` | SEKOLAH JENIS KEBANGSAAN (C) CHUNG HUA NO.4 KUCHING | a | `P.194_N.08_194/08/07_2a` |
| 365 | `P.194_N.08_194/08/07_2` | BANGUNAN LAMA MASJID DARUL IBADAH KPG. KUDEI LAMA | b | `P.194_N.08_194/08/07_2b` |

#### 43. `P.194_N.08_194/08/07_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 360 | `P.194_N.08_194/08/07_3` | SEKOLAH JENIS KEBANGSAAN (C) CHUNG HUA NO.4 KUCHING | a | `P.194_N.08_194/08/07_3a` |
| 366 | `P.194_N.08_194/08/07_3` | BANGUNAN LAMA MASJID DARUL IBADAH KPG. KUDEI LAMA | b | `P.194_N.08_194/08/07_3b` |

#### 44. `P.195_N.10_195/10/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 427 | `P.195_N.10_195/10/00_1` | RUANG B, DEWAN BADMINTON, KOMPLEKS PERUMAHAN POLIS TABUAN JAYA | a | `P.195_N.10_195/10/00_1a` |
| 429 | `P.195_N.10_195/10/00_1` | BILIK KULIAH SANTUBONG, RNO KUCHING | b | `P.195_N.10_195/10/00_1b` |

#### 45. `P.195_N.10_195/10/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 430 | `P.195_N.10_195/10/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_1a` |
| 436 | `P.195_N.10_195/10/01_1` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_1b` |
| 442 | `P.195_N.10_195/10/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_1c` |

#### 46. `P.195_N.10_195/10/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 431 | `P.195_N.10_195/10/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_2a` |
| 437 | `P.195_N.10_195/10/01_2` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_2b` |
| 443 | `P.195_N.10_195/10/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_2c` |

#### 47. `P.195_N.10_195/10/01_3` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 432 | `P.195_N.10_195/10/01_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_3a` |
| 438 | `P.195_N.10_195/10/01_3` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_3b` |
| 444 | `P.195_N.10_195/10/01_3` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_3c` |

#### 48. `P.195_N.10_195/10/01_4` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 433 | `P.195_N.10_195/10/01_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_4a` |
| 439 | `P.195_N.10_195/10/01_4` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_4b` |
| 445 | `P.195_N.10_195/10/01_4` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_4c` |

#### 49. `P.195_N.10_195/10/01_5` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 434 | `P.195_N.10_195/10/01_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_5a` |
| 440 | `P.195_N.10_195/10/01_5` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_5b` |
| 446 | `P.195_N.10_195/10/01_5` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_5c` |

#### 50. `P.195_N.10_195/10/01_6` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 435 | `P.195_N.10_195/10/01_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PENDING | a | `P.195_N.10_195/10/01_6a` |
| 441 | `P.195_N.10_195/10/01_6` | SEKOLAH MENENGAH KEBANGSAAN PENDING | b | `P.195_N.10_195/10/01_6b` |
| 447 | `P.195_N.10_195/10/01_6` | SEKOLAH JENIS KEBANGSAAN (CINA) BINTAWA | c | `P.195_N.10_195/10/01_6c` |

#### 51. `P.195_N.10_195/10/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 450 | `P.195_N.10_195/10/02_1` | SEKOLAH MENENGAH CHUNG HUA NO. 1 | a | `P.195_N.10_195/10/02_1a` |
| 453 | `P.195_N.10_195/10/02_1` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | b | `P.195_N.10_195/10/02_1b` |
| 460 | `P.195_N.10_195/10/02_1` | SEKOLAH KEBANGSAAN TABUAN ULU | c | `P.195_N.10_195/10/02_1c` |

#### 52. `P.195_N.10_195/10/02_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 451 | `P.195_N.10_195/10/02_2` | SEKOLAH MENENGAH CHUNG HUA NO. 1 | a | `P.195_N.10_195/10/02_2a` |
| 454 | `P.195_N.10_195/10/02_2` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | b | `P.195_N.10_195/10/02_2b` |
| 461 | `P.195_N.10_195/10/02_2` | SEKOLAH KEBANGSAAN TABUAN ULU | c | `P.195_N.10_195/10/02_2c` |

#### 53. `P.195_N.10_195/10/02_3` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 452 | `P.195_N.10_195/10/02_3` | SEKOLAH MENENGAH CHUNG HUA NO. 1 | a | `P.195_N.10_195/10/02_3a` |
| 455 | `P.195_N.10_195/10/02_3` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | b | `P.195_N.10_195/10/02_3b` |
| 462 | `P.195_N.10_195/10/02_3` | SEKOLAH KEBANGSAAN TABUAN ULU | c | `P.195_N.10_195/10/02_3c` |

#### 54. `P.195_N.10_195/10/02_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 456 | `P.195_N.10_195/10/02_4` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | a | `P.195_N.10_195/10/02_4a` |
| 463 | `P.195_N.10_195/10/02_4` | SEKOLAH KEBANGSAAN TABUAN ULU | b | `P.195_N.10_195/10/02_4b` |

#### 55. `P.195_N.10_195/10/02_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 457 | `P.195_N.10_195/10/02_5` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | a | `P.195_N.10_195/10/02_5a` |
| 464 | `P.195_N.10_195/10/02_5` | SEKOLAH KEBANGSAAN TABUAN ULU | b | `P.195_N.10_195/10/02_5b` |

#### 56. `P.195_N.10_195/10/02_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 458 | `P.195_N.10_195/10/02_6` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | a | `P.195_N.10_195/10/02_6a` |
| 465 | `P.195_N.10_195/10/02_6` | SEKOLAH KEBANGSAAN TABUAN ULU | b | `P.195_N.10_195/10/02_6b` |

#### 57. `P.195_N.10_195/10/02_7` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 459 | `P.195_N.10_195/10/02_7` | SEKOLAH MENENGAH CHUNG HUA NO. 3 | a | `P.195_N.10_195/10/02_7a` |
| 466 | `P.195_N.10_195/10/02_7` | SEKOLAH KEBANGSAAN TABUAN ULU | b | `P.195_N.10_195/10/02_7b` |

#### 58. `P.195_N.11_195/11/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 499 | `P.195_N.11_195/11/02_1` | SEKOLAH MENENGAH KEBANGSAAN BATU LINTANG | a | `P.195_N.11_195/11/02_1a` |
| 504 | `P.195_N.11_195/11/02_1` | SEKOLAH KEBANGSAAN JLN. ONG TIANG SWEE | b | `P.195_N.11_195/11/02_1b` |
| 509 | `P.195_N.11_195/11/02_1` | ST. JOHN AMBULANCE HEADQUARTERS BUILDING LORONG LAKSAMANA CHENG HO NO. 8 | c | `P.195_N.11_195/11/02_1c` |

#### 59. `P.195_N.11_195/11/02_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 500 | `P.195_N.11_195/11/02_2` | SEKOLAH MENENGAH KEBANGSAAN BATU LINTANG | a | `P.195_N.11_195/11/02_2a` |
| 505 | `P.195_N.11_195/11/02_2` | SEKOLAH KEBANGSAAN JLN. ONG TIANG SWEE | b | `P.195_N.11_195/11/02_2b` |
| 510 | `P.195_N.11_195/11/02_2` | ST. JOHN AMBULANCE HEADQUARTERS BUILDING LORONG LAKSAMANA CHENG HO NO. 8 | c | `P.195_N.11_195/11/02_2c` |

#### 60. `P.195_N.11_195/11/02_3` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 501 | `P.195_N.11_195/11/02_3` | SEKOLAH MENENGAH KEBANGSAAN BATU LINTANG | a | `P.195_N.11_195/11/02_3a` |
| 506 | `P.195_N.11_195/11/02_3` | SEKOLAH KEBANGSAAN JLN. ONG TIANG SWEE | b | `P.195_N.11_195/11/02_3b` |
| 511 | `P.195_N.11_195/11/02_3` | ST. JOHN AMBULANCE HEADQUARTERS BUILDING LORONG LAKSAMANA CHENG HO NO. 8 | c | `P.195_N.11_195/11/02_3c` |

#### 61. `P.195_N.11_195/11/02_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 502 | `P.195_N.11_195/11/02_4` | SEKOLAH MENENGAH KEBANGSAAN BATU LINTANG | a | `P.195_N.11_195/11/02_4a` |
| 507 | `P.195_N.11_195/11/02_4` | SEKOLAH KEBANGSAAN JLN. ONG TIANG SWEE | b | `P.195_N.11_195/11/02_4b` |

#### 62. `P.195_N.11_195/11/02_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 503 | `P.195_N.11_195/11/02_5` | SEKOLAH MENENGAH KEBANGSAAN BATU LINTANG | a | `P.195_N.11_195/11/02_5a` |
| 508 | `P.195_N.11_195/11/02_5` | SEKOLAH KEBANGSAAN JLN. ONG TIANG SWEE | b | `P.195_N.11_195/11/02_5b` |

#### 63. `P.195_N.11_195/11/08_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 533 | `P.195_N.11_195/11/08_1` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_1a` |
| 539 | `P.195_N.11_195/11/08_1` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_1b` |

#### 64. `P.195_N.11_195/11/08_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 534 | `P.195_N.11_195/11/08_2` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_2a` |
| 540 | `P.195_N.11_195/11/08_2` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_2b` |

#### 65. `P.195_N.11_195/11/08_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 535 | `P.195_N.11_195/11/08_3` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_3a` |
| 541 | `P.195_N.11_195/11/08_3` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_3b` |

#### 66. `P.195_N.11_195/11/08_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 536 | `P.195_N.11_195/11/08_4` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_4a` |
| 542 | `P.195_N.11_195/11/08_4` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_4b` |

#### 67. `P.195_N.11_195/11/08_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 537 | `P.195_N.11_195/11/08_5` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_5a` |
| 543 | `P.195_N.11_195/11/08_5` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_5b` |

#### 68. `P.195_N.11_195/11/08_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 538 | `P.195_N.11_195/11/08_6` | LODGE INTERNATIONAL SCHOOL | a | `P.195_N.11_195/11/08_6a` |
| 544 | `P.195_N.11_195/11/08_6` | SEKOLAH MENENGAH KEBANGSAAN LODGE | b | `P.195_N.11_195/11/08_6b` |

#### 69. `P.196_N.12_196/12/00_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 566 | `P.196_N.12_196/12/00_1` | RUANG A, DEWAN SENTOSA, JLN LAPANGAN TERBANG | a | `P.196_N.12_196/12/00_1a` |
| 567 | `P.196_N.12_196/12/00_1` | DEWAN TAN SRI HAMID BUGO, KEM RIA | b | `P.196_N.12_196/12/00_1b` |
| 568 | `P.196_N.12_196/12/00_1` | DEWAN MAKAN, 11 RAMD, KEM SEMENGGO | c | `P.196_N.12_196/12/00_1c` |
| 570 | `P.196_N.12_196/12/00_1` | DEWAN SERI ANGKASA, PANGKALAN UDARA KUCHING (TUDM) | d | `P.196_N.12_196/12/00_1d` |

#### 70. `P.196_N.12_196/12/00_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 569 | `P.196_N.12_196/12/00_2` | DEWAN MAKAN, 11 RAMD, KEM SEMENGGO | a | `P.196_N.12_196/12/00_2a` |
| 571 | `P.196_N.12_196/12/00_2` | DEWAN SERI ANGKASA, PANGKALAN UDARA KUCHING (TUDM) | b | `P.196_N.12_196/12/00_2b` |

#### 71. `P.196_N.12_196/12/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 573 | `P.196_N.12_196/12/01_1` | BALAI RAYA KPG. STAMPIN | a | `P.196_N.12_196/12/01_1a` |
| 574 | `P.196_N.12_196/12/01_1` | SEKOLAH JENIS KEBANGSAAN CHUNG HUA NO.2 | b | `P.196_N.12_196/12/01_1b` |

#### 72. `P.196_N.12_196/12/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 588 | `P.196_N.12_196/12/03_1` | SEKOLAH KEBANGSAAN SATRIA JAYA | a | `P.196_N.12_196/12/03_1a` |
| 596 | `P.196_N.12_196/12/03_1` | SEKOLAH KEBANGSAAN SG STUTONG | b | `P.196_N.12_196/12/03_1b` |

#### 73. `P.196_N.12_196/12/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 597 | `P.196_N.12_196/12/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SAM HAP HIN | a | `P.196_N.12_196/12/04_1a` |
| 607 | `P.196_N.12_196/12/04_1` | SEKOLAH KEBANGSAAN ST ALBAN | b | `P.196_N.12_196/12/04_1b` |
| 608 | `P.196_N.12_196/12/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAMPIN | c | `P.196_N.12_196/12/04_1c` |

#### 74. `P.196_N.12_196/12/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 598 | `P.196_N.12_196/12/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) SAM HAP HIN | a | `P.196_N.12_196/12/04_2a` |
| 609 | `P.196_N.12_196/12/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAMPIN | b | `P.196_N.12_196/12/04_2b` |

#### 75. `P.196_N.12_196/12/04_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 599 | `P.196_N.12_196/12/04_3` | SEKOLAH JENIS KEBANGSAAN (CINA) SAM HAP HIN | a | `P.196_N.12_196/12/04_3a` |
| 610 | `P.196_N.12_196/12/04_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAMPIN | b | `P.196_N.12_196/12/04_3b` |

#### 76. `P.196_N.13_196/13/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 643 | `P.196_N.13_196/13/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PANGKALAN BARU | a | `P.196_N.13_196/13/03_1a` |
| 644 | `P.196_N.13_196/13/03_1` | SEKOLAH KEBANGSAAN SACRED HEART | b | `P.196_N.13_196/13/03_1b` |
| 647 | `P.196_N.13_196/13/03_1` | DEWAN SERBAGUNA KPG. TEMATU | c | `P.196_N.13_196/13/03_1c` |
| 648 | `P.196_N.13_196/13/03_1` | SEKOLAH KEBANGSAAN JALAN HAJI BAKI | d | `P.196_N.13_196/13/03_1d` |

#### 77. `P.196_N.13_196/13/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 645 | `P.196_N.13_196/13/03_2` | SEKOLAH KEBANGSAAN SACRED HEART | a | `P.196_N.13_196/13/03_2a` |
| 649 | `P.196_N.13_196/13/03_2` | SEKOLAH KEBANGSAAN JALAN HAJI BAKI | b | `P.196_N.13_196/13/03_2b` |

#### 78. `P.196_N.13_196/13/03_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 646 | `P.196_N.13_196/13/03_3` | SEKOLAH KEBANGSAAN SACRED HEART | a | `P.196_N.13_196/13/03_3a` |
| 650 | `P.196_N.13_196/13/03_3` | SEKOLAH KEBANGSAAN JALAN HAJI BAKI | b | `P.196_N.13_196/13/03_3b` |

#### 79. `P.196_N.13_196/13/04_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 653 | `P.196_N.13_196/13/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 8 1/2 | a | `P.196_N.13_196/13/04_1a` |
| 655 | `P.196_N.13_196/13/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BT KITANG | b | `P.196_N.13_196/13/04_1b` |
| 657 | `P.196_N.13_196/13/04_1` | SEKOLAH KEBANGSAAN ST DAVID | c | `P.196_N.13_196/13/04_1c` |
| 659 | `P.196_N.13_196/13/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 10 | d | `P.196_N.13_196/13/04_1d` |
| 668 | `P.196_N.13_196/13/04_1` | SEKOLAH KEBANGSAAN LANDEH | e | `P.196_N.13_196/13/04_1e` |
| 673 | `P.196_N.13_196/13/04_1` | RCBM RECREATIONAL CENTRE BATU 13 | f | `P.196_N.13_196/13/04_1f` |

#### 80. `P.196_N.13_196/13/04_2` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 654 | `P.196_N.13_196/13/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 8 1/2 | a | `P.196_N.13_196/13/04_2a` |
| 656 | `P.196_N.13_196/13/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BT KITANG | b | `P.196_N.13_196/13/04_2b` |
| 658 | `P.196_N.13_196/13/04_2` | SEKOLAH KEBANGSAAN ST DAVID | c | `P.196_N.13_196/13/04_2c` |
| 660 | `P.196_N.13_196/13/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 10 | d | `P.196_N.13_196/13/04_2d` |
| 669 | `P.196_N.13_196/13/04_2` | SEKOLAH KEBANGSAAN LANDEH | e | `P.196_N.13_196/13/04_2e` |

#### 81. `P.196_N.13_196/13/04_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 661 | `P.196_N.13_196/13/04_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 10 | a | `P.196_N.13_196/13/04_3a` |
| 670 | `P.196_N.13_196/13/04_3` | SEKOLAH KEBANGSAAN LANDEH | b | `P.196_N.13_196/13/04_3b` |

#### 82. `P.196_N.13_196/13/04_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 662 | `P.196_N.13_196/13/04_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 10 | a | `P.196_N.13_196/13/04_4a` |
| 671 | `P.196_N.13_196/13/04_4` | SEKOLAH KEBANGSAAN LANDEH | b | `P.196_N.13_196/13/04_4b` |

#### 83. `P.196_N.13_196/13/04_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 663 | `P.196_N.13_196/13/04_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 10 | a | `P.196_N.13_196/13/04_5a` |
| 672 | `P.196_N.13_196/13/04_5` | SEKOLAH KEBANGSAAN LANDEH | b | `P.196_N.13_196/13/04_5b` |

#### 84. `P.196_N.14_196/14/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 677 | `P.196_N.14_196/14/00_1` | RUANG A, DEWAN MINDA, JLN. LAPANGAN TERBANG | a | `P.196_N.14_196/14/00_1a` |
| 678 | `P.196_N.14_196/14/00_1` | DEWAN PAHLAWAN IBU PEJABAT BRIGED SARAWAK | b | `P.196_N.14_196/14/00_1b` |

#### 85. `P.196_N.14_196/14/03_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 685 | `P.196_N.14_196/14/03_1` | SEKOLAH KEBANGSAAN RANTAU PANJANG | a | `P.196_N.14_196/14/03_1a` |
| 686 | `P.196_N.14_196/14/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAPOK | b | `P.196_N.14_196/14/03_1b` |
| 691 | `P.196_N.14_196/14/03_1` | SEKOLAH MENENGAH KEBANGSAAN BATU KAWA | c | `P.196_N.14_196/14/03_1c` |
| 692 | `P.196_N.14_196/14/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | d | `P.196_N.14_196/14/03_1d` |
| 699 | `P.196_N.14_196/14/03_1` | BALAI RAYA KPG. DESA WIRA LOT | e | `P.196_N.14_196/14/03_1e` |
| 701 | `P.196_N.14_196/14/03_1` | SEKOLAH KEBANGSAAN JALAN ARANG | f | `P.196_N.14_196/14/03_1f` |
| 706 | `P.196_N.14_196/14/03_1` | SEKOLAH KEBANGSAAN STAPOK | g | `P.196_N.14_196/14/03_1g` |
| 707 | `P.196_N.14_196/14/03_1` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | h | `P.196_N.14_196/14/03_1h` |
| 714 | `P.196_N.14_196/14/03_1` | DEWAN SERBAGUNA KPG. SINAR BUDI | i | `P.196_N.14_196/14/03_1i` |

#### 86. `P.196_N.14_196/14/03_2` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 687 | `P.196_N.14_196/14/03_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAPOK | a | `P.196_N.14_196/14/03_2a` |
| 693 | `P.196_N.14_196/14/03_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | b | `P.196_N.14_196/14/03_2b` |
| 700 | `P.196_N.14_196/14/03_2` | BALAI RAYA KPG. DESA WIRA LOT | c | `P.196_N.14_196/14/03_2c` |
| 702 | `P.196_N.14_196/14/03_2` | SEKOLAH KEBANGSAAN JALAN ARANG | d | `P.196_N.14_196/14/03_2d` |
| 708 | `P.196_N.14_196/14/03_2` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | e | `P.196_N.14_196/14/03_2e` |
| 715 | `P.196_N.14_196/14/03_2` | DEWAN SERBAGUNA KPG. SINAR BUDI | f | `P.196_N.14_196/14/03_2f` |

#### 87. `P.196_N.14_196/14/03_3` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 688 | `P.196_N.14_196/14/03_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAPOK | a | `P.196_N.14_196/14/03_3a` |
| 694 | `P.196_N.14_196/14/03_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | b | `P.196_N.14_196/14/03_3b` |
| 703 | `P.196_N.14_196/14/03_3` | SEKOLAH KEBANGSAAN JALAN ARANG | c | `P.196_N.14_196/14/03_3c` |
| 709 | `P.196_N.14_196/14/03_3` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | d | `P.196_N.14_196/14/03_3d` |

#### 88. `P.196_N.14_196/14/03_4` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 689 | `P.196_N.14_196/14/03_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAPOK | a | `P.196_N.14_196/14/03_4a` |
| 695 | `P.196_N.14_196/14/03_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | b | `P.196_N.14_196/14/03_4b` |
| 704 | `P.196_N.14_196/14/03_4` | SEKOLAH KEBANGSAAN JALAN ARANG | c | `P.196_N.14_196/14/03_4c` |
| 710 | `P.196_N.14_196/14/03_4` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | d | `P.196_N.14_196/14/03_4d` |

#### 89. `P.196_N.14_196/14/03_5` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 690 | `P.196_N.14_196/14/03_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA STAPOK | a | `P.196_N.14_196/14/03_5a` |
| 696 | `P.196_N.14_196/14/03_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | b | `P.196_N.14_196/14/03_5b` |
| 705 | `P.196_N.14_196/14/03_5` | SEKOLAH KEBANGSAAN JALAN ARANG | c | `P.196_N.14_196/14/03_5c` |
| 711 | `P.196_N.14_196/14/03_5` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | d | `P.196_N.14_196/14/03_5d` |

#### 90. `P.196_N.14_196/14/03_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 697 | `P.196_N.14_196/14/03_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | a | `P.196_N.14_196/14/03_6a` |
| 712 | `P.196_N.14_196/14/03_6` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | b | `P.196_N.14_196/14/03_6b` |

#### 91. `P.196_N.14_196/14/03_7` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 698 | `P.196_N.14_196/14/03_7` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG TAPANG | a | `P.196_N.14_196/14/03_7a` |
| 713 | `P.196_N.14_196/14/03_7` | DEWAN PERPADUAN RPR BATU KAWA FASA 1 | b | `P.196_N.14_196/14/03_7b` |

#### 92. `P.197_N.15_197/15/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 729 | `P.197_N.15_197/15/02_1` | SEKOLAH KEBANGSAAN TANJONG APONG | a | `P.197_N.15_197/15/02_1a` |
| 731 | `P.197_N.15_197/15/02_1` | DEWAN KPG. SRI TAJO | b | `P.197_N.15_197/15/02_1b` |

#### 93. `P.197_N.16_197/16/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 776 | `P.197_N.16_197/16/00_1` | (RUANG B) DEWAN SERBAGUNA IPD KOTA SAMARAHAN | a | `P.197_N.16_197/16/00_1a` |
| 777 | `P.197_N.16_197/16/00_1` | BILIK SEMINAR NO. 25 KOLEJ KENANGA UNIMAS | b | `P.197_N.16_197/16/00_1b` |

#### 94. `P.197_N.16_197/16/01_1` (12 occurrences, 12 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 778 | `P.197_N.16_197/16/01_1` | BILIK GERAKAN ERT KPG. NAKONG | a | `P.197_N.16_197/16/01_1a` |
| 779 | `P.197_N.16_197/16/01_1` | SEKOLAH KEBANGSAAN MANG | b | `P.197_N.16_197/16/01_1b` |
| 781 | `P.197_N.16_197/16/01_1` | SEKOLAH KEBANGSAAN KG MELAYU | c | `P.197_N.16_197/16/01_1c` |
| 784 | `P.197_N.16_197/16/01_1` | DEWAN SERBAGUNA KPG. BANGKA SEMONG | d | `P.197_N.16_197/16/01_1d` |
| 785 | `P.197_N.16_197/16/01_1` | DEWAN KPG. NAIE | e | `P.197_N.16_197/16/01_1e` |
| 786 | `P.197_N.16_197/16/01_1` | SURAU KPG. NAIE BARU | f | `P.197_N.16_197/16/01_1f` |
| 787 | `P.197_N.16_197/16/01_1` | DEWAN KPG. SG. MATA | g | `P.197_N.16_197/16/01_1g` |
| 788 | `P.197_N.16_197/16/01_1` | DEWAN KPG. EMPILA | h | `P.197_N.16_197/16/01_1h` |
| 790 | `P.197_N.16_197/16/01_1` | SEKOLAH KEBANGSAAN NIUP | i | `P.197_N.16_197/16/01_1i` |
| 791 | `P.197_N.16_197/16/01_1` | DEWAN SERBAGUNA KPG. TANJUNG PARANG | j | `P.197_N.16_197/16/01_1j` |
| 792 | `P.197_N.16_197/16/01_1` | (RUANG A) BALAI RAYA KPG. SEMAWANG | k | `P.197_N.16_197/16/01_1k` |
| 793 | `P.197_N.16_197/16/01_1` | SEKOLAH KEBANGSAAN KG TANJUNG TUANG | l | `P.197_N.16_197/16/01_1l` |

#### 95. `P.197_N.16_197/16/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 780 | `P.197_N.16_197/16/01_2` | SEKOLAH KEBANGSAAN MANG | a | `P.197_N.16_197/16/01_2a` |
| 782 | `P.197_N.16_197/16/01_2` | SEKOLAH KEBANGSAAN KG MELAYU | b | `P.197_N.16_197/16/01_2b` |
| 789 | `P.197_N.16_197/16/01_2` | DEWAN KPG. EMPILA | c | `P.197_N.16_197/16/01_2c` |

#### 96. `P.197_N.16_197/16/04_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 801 | `P.197_N.16_197/16/04_1` | SEKOLAH KEBANGSAAN DATO TRAOH MUARA TUANG | a | `P.197_N.16_197/16/04_1a` |
| 805 | `P.197_N.16_197/16/04_1` | DEWAN KPG. MUARA TUANG | b | `P.197_N.16_197/16/04_1b` |
| 807 | `P.197_N.16_197/16/04_1` | SEKOLAH KEBANGSAAN DATO MOHD MUSA | c | `P.197_N.16_197/16/04_1c` |
| 812 | `P.197_N.16_197/16/04_1` | DEWAN KPG. ENDAP | d | `P.197_N.16_197/16/04_1d` |
| 814 | `P.197_N.16_197/16/04_1` | SEKOLAH KEBANGSAAN PINANG | e | `P.197_N.16_197/16/04_1e` |
| 817 | `P.197_N.16_197/16/04_1` | SEKOLAH KEBANGSAAN MERANEK | f | `P.197_N.16_197/16/04_1f` |
| 820 | `P.197_N.16_197/16/04_1` | SEKOLAH KEBANGSAAN KAMPUNG REMBUS | g | `P.197_N.16_197/16/04_1g` |

#### 97. `P.197_N.16_197/16/04_2` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 802 | `P.197_N.16_197/16/04_2` | SEKOLAH KEBANGSAAN DATO TRAOH MUARA TUANG | a | `P.197_N.16_197/16/04_2a` |
| 806 | `P.197_N.16_197/16/04_2` | DEWAN KPG. MUARA TUANG | b | `P.197_N.16_197/16/04_2b` |
| 808 | `P.197_N.16_197/16/04_2` | SEKOLAH KEBANGSAAN DATO MOHD MUSA | c | `P.197_N.16_197/16/04_2c` |
| 813 | `P.197_N.16_197/16/04_2` | DEWAN KPG. ENDAP | d | `P.197_N.16_197/16/04_2d` |
| 815 | `P.197_N.16_197/16/04_2` | SEKOLAH KEBANGSAAN PINANG | e | `P.197_N.16_197/16/04_2e` |
| 818 | `P.197_N.16_197/16/04_2` | SEKOLAH KEBANGSAAN MERANEK | f | `P.197_N.16_197/16/04_2f` |
| 821 | `P.197_N.16_197/16/04_2` | SEKOLAH KEBANGSAAN KAMPUNG REMBUS | g | `P.197_N.16_197/16/04_2g` |

#### 98. `P.197_N.16_197/16/04_3` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 803 | `P.197_N.16_197/16/04_3` | SEKOLAH KEBANGSAAN DATO TRAOH MUARA TUANG | a | `P.197_N.16_197/16/04_3a` |
| 809 | `P.197_N.16_197/16/04_3` | SEKOLAH KEBANGSAAN DATO MOHD MUSA | b | `P.197_N.16_197/16/04_3b` |
| 816 | `P.197_N.16_197/16/04_3` | SEKOLAH KEBANGSAAN PINANG | c | `P.197_N.16_197/16/04_3c` |
| 819 | `P.197_N.16_197/16/04_3` | SEKOLAH KEBANGSAAN MERANEK | d | `P.197_N.16_197/16/04_3d` |

#### 99. `P.197_N.16_197/16/04_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 804 | `P.197_N.16_197/16/04_4` | SEKOLAH KEBANGSAAN DATO TRAOH MUARA TUANG | a | `P.197_N.16_197/16/04_4a` |
| 810 | `P.197_N.16_197/16/04_4` | SEKOLAH KEBANGSAAN DATO MOHD MUSA | b | `P.197_N.16_197/16/04_4b` |

#### 100. `P.197_N.16_197/16/05_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 822 | `P.197_N.16_197/16/05_1` | SEKOLAH KEBANGSAAN TANAH MERAH | a | `P.197_N.16_197/16/05_1a` |
| 823 | `P.197_N.16_197/16/05_1` | BALAI RAYA KPG. SEGENAM | b | `P.197_N.16_197/16/05_1b` |
| 824 | `P.197_N.16_197/16/05_1` | RUMAH PANJANG KPG. RAEH | c | `P.197_N.16_197/16/05_1c` |
| 825 | `P.197_N.16_197/16/05_1` | TADIKA KEMAS KPG. SOH | d | `P.197_N.16_197/16/05_1d` |
| 826 | `P.197_N.16_197/16/05_1` | SEKOLAH KEBANGSAAN GEMANG | e | `P.197_N.16_197/16/05_1e` |
| 828 | `P.197_N.16_197/16/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BATU 29 | f | `P.197_N.16_197/16/05_1f` |

#### 101. `P.197_N.16_197/16/06_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 829 | `P.197_N.16_197/16/06_1` | SEKOLAH KEBANGSAAN LUBOK ANTU REBAN | a | `P.197_N.16_197/16/06_1a` |
| 830 | `P.197_N.16_197/16/06_1` | SEKOLAH KEBANGSAAN PATI | b | `P.197_N.16_197/16/06_1b` |
| 832 | `P.197_N.16_197/16/06_1` | SEKOLAH KEBANGSAAN SAMARAHAN ESTATE | c | `P.197_N.16_197/16/06_1c` |
| 833 | `P.197_N.16_197/16/06_1` | DEWAN KPG. SENIAWAN | d | `P.197_N.16_197/16/06_1d` |
| 834 | `P.197_N.16_197/16/06_1` | SEKOLAH KEBANGSAAN PLAMAN BAKI / MENAUL | e | `P.197_N.16_197/16/06_1e` |
| 836 | `P.197_N.16_197/16/06_1` | BALAI RAYA KPG. TIAN MAWANG | f | `P.197_N.16_197/16/06_1f` |
| 837 | `P.197_N.16_197/16/06_1` | BALAI RAYA KPG. MURUD PLAMAN | g | `P.197_N.16_197/16/06_1g` |

#### 102. `P.197_N.16_197/16/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 831 | `P.197_N.16_197/16/06_2` | SEKOLAH KEBANGSAAN PATI | a | `P.197_N.16_197/16/06_2a` |
| 835 | `P.197_N.16_197/16/06_2` | SEKOLAH KEBANGSAAN PLAMAN BAKI / MENAUL | b | `P.197_N.16_197/16/06_2b` |

#### 103. `P.197_N.16_197/16/07_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 838 | `P.197_N.16_197/16/07_1` | DEWAN KPG. SEBAYOR | a | `P.197_N.16_197/16/07_1a` |
| 839 | `P.197_N.16_197/16/07_1` | BALAI RAYA KPG. MELABAN (RUANG A) | b | `P.197_N.16_197/16/07_1b` |
| 840 | `P.197_N.16_197/16/07_1` | BALAI RAYA KPG. MELABAN (RUANG B) | c | `P.197_N.16_197/16/07_1c` |
| 841 | `P.197_N.16_197/16/07_1` | SEKOLAH KEBANGSAAN PLAIE D C | d | `P.197_N.16_197/16/07_1d` |
| 842 | `P.197_N.16_197/16/07_1` | SEKOLAH KEBANGSAAN ST MICHAEL | e | `P.197_N.16_197/16/07_1e` |

#### 104. `P.197_N.17_197/17/00_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 845 | `P.197_N.17_197/17/00_1` | (RUANG C) DEWAN SERBAGUNA IPD KOTA SAMARAHAN | a | `P.197_N.17_197/17/00_1a` |
| 846 | `P.197_N.17_197/17/00_1` | DEWAN SERBA GUNA KEM MUARA TUANG | b | `P.197_N.17_197/17/00_1b` |
| 851 | `P.197_N.17_197/17/00_1` | DEWAN BUKAVU KEM PENRISSEN | c | `P.197_N.17_197/17/00_1c` |

#### 105. `P.197_N.17_197/17/00_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 847 | `P.197_N.17_197/17/00_2` | DEWAN SERBA GUNA KEM MUARA TUANG | a | `P.197_N.17_197/17/00_2a` |
| 852 | `P.197_N.17_197/17/00_2` | DEWAN BUKAVU KEM PENRISSEN | b | `P.197_N.17_197/17/00_2b` |

#### 106. `P.197_N.17_197/17/00_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 848 | `P.197_N.17_197/17/00_3` | DEWAN SERBA GUNA KEM MUARA TUANG | a | `P.197_N.17_197/17/00_3a` |
| 853 | `P.197_N.17_197/17/00_3` | DEWAN BUKAVU KEM PENRISSEN | b | `P.197_N.17_197/17/00_3b` |

#### 107. `P.197_N.17_197/17/00_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 849 | `P.197_N.17_197/17/00_4` | DEWAN SERBA GUNA KEM MUARA TUANG | a | `P.197_N.17_197/17/00_4a` |
| 854 | `P.197_N.17_197/17/00_4` | DEWAN BUKAVU KEM PENRISSEN | b | `P.197_N.17_197/17/00_4b` |

#### 108. `P.197_N.17_197/17/00_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 850 | `P.197_N.17_197/17/00_5` | DEWAN SERBA GUNA KEM MUARA TUANG | a | `P.197_N.17_197/17/00_5a` |
| 855 | `P.197_N.17_197/17/00_5` | DEWAN BUKAVU KEM PENRISSEN | b | `P.197_N.17_197/17/00_5b` |

#### 109. `P.197_N.17_197/17/01_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 858 | `P.197_N.17_197/17/01_1` | SEKOLAH KEBANGSAAN BINYU | a | `P.197_N.17_197/17/01_1a` |
| 859 | `P.197_N.17_197/17/01_1` | DEWAN KPG. MERDANG GAYAM | b | `P.197_N.17_197/17/01_1b` |
| 860 | `P.197_N.17_197/17/01_1` | SEKOLAH MENENGAH KEBANGSAAN MUARA TUANG | c | `P.197_N.17_197/17/01_1c` |
| 862 | `P.197_N.17_197/17/01_1` | BALAI RAYA SG. EMPIT | d | `P.197_N.17_197/17/01_1d` |
| 863 | `P.197_N.17_197/17/01_1` | SEKOLAH KEBANGSAAN ST MARTIN | e | `P.197_N.17_197/17/01_1e` |
| 867 | `P.197_N.17_197/17/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG JERNANG | f | `P.197_N.17_197/17/01_1f` |

#### 110. `P.197_N.17_197/17/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 861 | `P.197_N.17_197/17/01_2` | SEKOLAH MENENGAH KEBANGSAAN MUARA TUANG | a | `P.197_N.17_197/17/01_2a` |
| 864 | `P.197_N.17_197/17/01_2` | SEKOLAH KEBANGSAAN ST MARTIN | b | `P.197_N.17_197/17/01_2b` |
| 868 | `P.197_N.17_197/17/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG JERNANG | c | `P.197_N.17_197/17/01_2c` |

#### 111. `P.197_N.17_197/17/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 873 | `P.197_N.17_197/17/03_1` | SEKOLAH MENENGAH KEBANGSAAN PENRISSEN | a | `P.197_N.17_197/17/03_1a` |
| 881 | `P.197_N.17_197/17/03_1` | DEWAN SRI LESTARI KPG. SRI ARJUNA | b | `P.197_N.17_197/17/03_1b` |

#### 112. `P.197_N.17_197/17/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 874 | `P.197_N.17_197/17/03_2` | SEKOLAH MENENGAH KEBANGSAAN PENRISSEN | a | `P.197_N.17_197/17/03_2a` |
| 882 | `P.197_N.17_197/17/03_2` | DEWAN SRI LESTARI KPG. SRI ARJUNA | b | `P.197_N.17_197/17/03_2b` |

#### 113. `P.197_N.17_197/17/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 886 | `P.197_N.17_197/17/05_1` | SEKOLAH MENENGAH KEBANGSAAN WIRA PENRISSEN | a | `P.197_N.17_197/17/05_1a` |
| 894 | `P.197_N.17_197/17/05_1` | DEWAN KPG. STAKAN | b | `P.197_N.17_197/17/05_1b` |

#### 114. `P.197_N.17_197/17/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 887 | `P.197_N.17_197/17/05_2` | SEKOLAH MENENGAH KEBANGSAAN WIRA PENRISSEN | a | `P.197_N.17_197/17/05_2a` |
| 895 | `P.197_N.17_197/17/05_2` | DEWAN KPG. STAKAN | b | `P.197_N.17_197/17/05_2b` |

#### 115. `P.198_N.18_198/18/15_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 924 | `P.198_N.18_198/18/15_1` | DEWAN SERBAGUNA KPG SKIO | a | `P.198_N.18_198/18/15_1a` |
| 925 | `P.198_N.18_198/18/15_1` | BALAI RAYA KPG. SOGO | b | `P.198_N.18_198/18/15_1b` |

#### 116. `P.198_N.18_198/18/17_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 927 | `P.198_N.18_198/18/17_1` | BALAI RAYA PANGKALAN TEBANG | a | `P.198_N.18_198/18/17_1a` |
| 928 | `P.198_N.18_198/18/17_1` | BALAI RAYA LEDAN GUMBANG | b | `P.198_N.18_198/18/17_1b` |

#### 117. `P.198_N.19_198/19/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 936 | `P.198_N.19_198/19/00_1` | DEWAN SEKOLAH LATIHAN KONTINJEN, BALAI POLIS LAPANGAN TERBANG ANTARABANGSA KUCHING, JALAN LAPANGAN TERBANG | a | `P.198_N.19_198/19/00_1a` |
| 938 | `P.198_N.19_198/19/00_1` | BANGUNAN TERBUKA BENGKEL PENGANGKUTAN BAHAGIAN LOGISTIK DAN TEKNOLOGI IPD PADAWAN | b | `P.198_N.19_198/19/00_1b` |

#### 118. `P.198_N.19_198/19/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 951 | `P.198_N.19_198/19/03_1` | DEWAN SERBAGUNA KPG. BRA'ANG PAYANG | a | `P.198_N.19_198/19/03_1a` |
| 952 | `P.198_N.19_198/19/03_1` | BALAI RAYA KPG. BRA'ANG SIGANDAR | b | `P.198_N.19_198/19/03_1b` |

#### 119. `P.198_N.19_198/19/16_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 968 | `P.198_N.19_198/19/16_1` | BALAI RAYA KPG. TEMURANG BARU | a | `P.198_N.19_198/19/16_1a` |
| 969 | `P.198_N.19_198/19/16_1` | BALAI RAYA KPG. BIYA JABER | b | `P.198_N.19_198/19/16_1b` |

#### 120. `P.198_N.19_198/19/26_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 988 | `P.198_N.19_198/19/26_1` | BALAI RAYA KPG. BIYA PARANG | a | `P.198_N.19_198/19/26_1a` |
| 989 | `P.198_N.19_198/19/26_1` | BALAI RAYA KPG. SEPIT | b | `P.198_N.19_198/19/26_1b` |

#### 121. `P.198_N.19_198/19/27_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 990 | `P.198_N.19_198/19/27_1` | BALAI RAYA KPG. KIDING | a | `P.198_N.19_198/19/27_1a` |
| 991 | `P.198_N.19_198/19/27_1` | BALAI RAYA KPG. BIYA KAKAS | b | `P.198_N.19_198/19/27_1b` |

#### 122. `P.198_N.19_198/19/29_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 993 | `P.198_N.19_198/19/29_1` | BALAI RAYA KPG. NUSARAYA | a | `P.198_N.19_198/19/29_1a` |
| 994 | `P.198_N.19_198/19/29_1` | BALAI RAYA KPG. BIYA KEMAS | b | `P.198_N.19_198/19/29_1b` |

#### 123. `P.198_N.20_198/20/28_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1048 | `P.198_N.20_198/20/28_1` | BALAI RAYA KPG. ENTAWA SG. BARIE | a | `P.198_N.20_198/20/28_1a` |
| 1049 | `P.198_N.20_198/20/28_1` | SEKOLAH KEBANGSAAN SG KENYAH | b | `P.198_N.20_198/20/28_1b` |

#### 124. `P.198_N.20_198/20/33_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1059 | `P.198_N.20_198/20/33_1` | SEKOLAH KEBANGSAAN TANAH PUTEH | a | `P.198_N.20_198/20/33_1a` |
| 1060 | `P.198_N.20_198/20/33_1` | RH. PANJANG, KAMPUNG MUNGGU KOPI | b | `P.198_N.20_198/20/33_1b` |

#### 125. `P.199_N.21_199/21/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1063 | `P.199_N.21_199/21/01_1` | SEKOLAH KEBANGSAAN ALL SAINTS PLAMAN NYABET | a | `P.199_N.21_199/21/01_1a` |
| 1064 | `P.199_N.21_199/21/01_1` | SEKOLAH KEBANGSAAN ST JOHN TAEE | b | `P.199_N.21_199/21/01_1b` |
| 1067 | `P.199_N.21_199/21/01_1` | SEKOLAH KEBANGSAAN PARUN SUAN | c | `P.199_N.21_199/21/01_1c` |

#### 126. `P.199_N.21_199/21/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1065 | `P.199_N.21_199/21/01_2` | SEKOLAH KEBANGSAAN ST JOHN TAEE | a | `P.199_N.21_199/21/01_2a` |
| 1068 | `P.199_N.21_199/21/01_2` | SEKOLAH KEBANGSAAN PARUN SUAN | b | `P.199_N.21_199/21/01_2b` |

#### 127. `P.199_N.21_199/21/01_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1066 | `P.199_N.21_199/21/01_3` | SEKOLAH KEBANGSAAN ST JOHN TAEE | a | `P.199_N.21_199/21/01_3a` |
| 1069 | `P.199_N.21_199/21/01_3` | SEKOLAH KEBANGSAAN PARUN SUAN | b | `P.199_N.21_199/21/01_3b` |

#### 128. `P.199_N.21_199/21/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1070 | `P.199_N.21_199/21/02_1` | SEKOLAH KEBANGSAAN ST MATHEW LANCHANG | a | `P.199_N.21_199/21/02_1a` |
| 1072 | `P.199_N.21_199/21/02_1` | SEKOLAH KEBANGSAAN ST DOMINIC PICHIN | b | `P.199_N.21_199/21/02_1b` |

#### 129. `P.199_N.21_199/21/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1071 | `P.199_N.21_199/21/02_2` | SEKOLAH KEBANGSAAN ST MATHEW LANCHANG | a | `P.199_N.21_199/21/02_2a` |
| 1073 | `P.199_N.21_199/21/02_2` | SEKOLAH KEBANGSAAN ST DOMINIC PICHIN | b | `P.199_N.21_199/21/02_2b` |

#### 130. `P.199_N.21_199/21/03_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1075 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN TEPOI | a | `P.199_N.21_199/21/03_1a` |
| 1076 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN TEMONG | b | `P.199_N.21_199/21/03_1b` |
| 1079 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN ENTUBUH | c | `P.199_N.21_199/21/03_1c` |
| 1080 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN SEJIJAG | d | `P.199_N.21_199/21/03_1d` |
| 1081 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN TEBEDU | e | `P.199_N.21_199/21/03_1e` |
| 1084 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN SUNGAN | f | `P.199_N.21_199/21/03_1f` |
| 1085 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN TEMA | g | `P.199_N.21_199/21/03_1g` |
| 1087 | `P.199_N.21_199/21/03_1` | SEKOLAH KEBANGSAAN SG SAMERAN | h | `P.199_N.21_199/21/03_1h` |

#### 131. `P.199_N.21_199/21/03_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1077 | `P.199_N.21_199/21/03_2` | SEKOLAH KEBANGSAAN TEMONG | a | `P.199_N.21_199/21/03_2a` |
| 1082 | `P.199_N.21_199/21/03_2` | SEKOLAH KEBANGSAAN TEBEDU | b | `P.199_N.21_199/21/03_2b` |
| 1086 | `P.199_N.21_199/21/03_2` | SEKOLAH KEBANGSAAN TEMA | c | `P.199_N.21_199/21/03_2c` |
| 1088 | `P.199_N.21_199/21/03_2` | SEKOLAH KEBANGSAAN SG SAMERAN | d | `P.199_N.21_199/21/03_2d` |

#### 132. `P.199_N.21_199/21/03_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1078 | `P.199_N.21_199/21/03_3` | SEKOLAH KEBANGSAAN TEMONG | a | `P.199_N.21_199/21/03_3a` |
| 1083 | `P.199_N.21_199/21/03_3` | SEKOLAH KEBANGSAAN TEBEDU | b | `P.199_N.21_199/21/03_3b` |

#### 133. `P.199_N.21_199/21/04_1` (10 occurrences, 10 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1089 | `P.199_N.21_199/21/04_1` | SEKOLAH KEBANGSAAN GAHAT MAWANG | a | `P.199_N.21_199/21/04_1a` |
| 1090 | `P.199_N.21_199/21/04_1` | BANGUNAN KELAS TADIKA BIDAK | b | `P.199_N.21_199/21/04_1b` |
| 1091 | `P.199_N.21_199/21/04_1` | SEKOLAH KEBANGSAAN RETEH | c | `P.199_N.21_199/21/04_1c` |
| 1092 | `P.199_N.21_199/21/04_1` | SEKOLAH KEBANGSAAN MAWANG TAUP | d | `P.199_N.21_199/21/04_1d` |
| 1093 | `P.199_N.21_199/21/04_1` | BALAI RAYA KUJANG SAIN | e | `P.199_N.21_199/21/04_1e` |
| 1094 | `P.199_N.21_199/21/04_1` | BALAI RAYA KPG. KUJANG MAWANG | f | `P.199_N.21_199/21/04_1f` |
| 1095 | `P.199_N.21_199/21/04_1` | SEKOLAH KEBANGSAAN KUJANG MAWANG | g | `P.199_N.21_199/21/04_1g` |
| 1096 | `P.199_N.21_199/21/04_1` | SEKOLAH KEBANGSAAN TESU | h | `P.199_N.21_199/21/04_1h` |
| 1097 | `P.199_N.21_199/21/04_1` | BILIK GERAKAN JKKK KPG. DAHA SEROBAN | i | `P.199_N.21_199/21/04_1i` |
| 1098 | `P.199_N.21_199/21/04_1` | BALAI RAYA KPG. DAHA KISAU | j | `P.199_N.21_199/21/04_1j` |

#### 134. `P.199_N.21_199/21/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1099 | `P.199_N.21_199/21/05_1` | SEKOLAH KEBANGSAAN LOBANG BATU | a | `P.199_N.21_199/21/05_1a` |
| 1102 | `P.199_N.21_199/21/05_1` | BALAI RAYA KPG. SEBINTIN | b | `P.199_N.21_199/21/05_1b` |
| 1103 | `P.199_N.21_199/21/05_1` | SEKOLAH KEBANGSAAN KRUSEN | c | `P.199_N.21_199/21/05_1c` |

#### 135. `P.199_N.21_199/21/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1100 | `P.199_N.21_199/21/05_2` | SEKOLAH KEBANGSAAN LOBANG BATU | a | `P.199_N.21_199/21/05_2a` |
| 1104 | `P.199_N.21_199/21/05_2` | SEKOLAH KEBANGSAAN KRUSEN | b | `P.199_N.21_199/21/05_2b` |

#### 136. `P.199_N.22_199/22/01_1` (10 occurrences, 10 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1107 | `P.199_N.22_199/22/01_1` | SEKOLAH KEBANGSAAN ST ANTHONY KAWAN | a | `P.199_N.22_199/22/01_1a` |
| 1109 | `P.199_N.22_199/22/01_1` | DEWAN SERBAGUNA KPG MAYANG MAWANG | b | `P.199_N.22_199/22/01_1b` |
| 1110 | `P.199_N.22_199/22/01_1` | DEWAN BABUK SALIM | c | `P.199_N.22_199/22/01_1c` |
| 1112 | `P.199_N.22_199/22/01_1` | BALAI RAYA KPG. BUGU | d | `P.199_N.22_199/22/01_1d` |
| 1113 | `P.199_N.22_199/22/01_1` | BALAI RAYA KPG. PRANGKAN MARUNG | e | `P.199_N.22_199/22/01_1e` |
| 1114 | `P.199_N.22_199/22/01_1` | SEKOLAH KEBANGSAAN SG RIMU | f | `P.199_N.22_199/22/01_1f` |
| 1115 | `P.199_N.22_199/22/01_1` | SEKOLAH KEBANGSAAN ST JOHN MENTONG | g | `P.199_N.22_199/22/01_1g` |
| 1117 | `P.199_N.22_199/22/01_1` | SEKOLAH KEBANGSAAN MUBOK BERAWAN | h | `P.199_N.22_199/22/01_1h` |
| 1118 | `P.199_N.22_199/22/01_1` | DEWAN SERBAGUNA KAMPUNG PRIDAN | i | `P.199_N.22_199/22/01_1i` |
| 1120 | `P.199_N.22_199/22/01_1` | BALAI RAYA KPG. SANGGAI MAWANG | j | `P.199_N.22_199/22/01_1j` |

#### 137. `P.199_N.22_199/22/01_2` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1108 | `P.199_N.22_199/22/01_2` | SEKOLAH KEBANGSAAN ST ANTHONY KAWAN | a | `P.199_N.22_199/22/01_2a` |
| 1111 | `P.199_N.22_199/22/01_2` | DEWAN BABUK SALIM | b | `P.199_N.22_199/22/01_2b` |
| 1116 | `P.199_N.22_199/22/01_2` | SEKOLAH KEBANGSAAN ST JOHN MENTONG | c | `P.199_N.22_199/22/01_2c` |
| 1119 | `P.199_N.22_199/22/01_2` | DEWAN SERBAGUNA KAMPUNG PRIDAN | d | `P.199_N.22_199/22/01_2d` |
| 1121 | `P.199_N.22_199/22/01_2` | BALAI RAYA KPG. SANGGAI MAWANG | e | `P.199_N.22_199/22/01_2e` |

#### 138. `P.199_N.22_199/22/02_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1122 | `P.199_N.22_199/22/02_1` | SEKOLAH KEBANGSAAN MENTU TAPU | a | `P.199_N.22_199/22/02_1a` |
| 1124 | `P.199_N.22_199/22/02_1` | SEKOLAH KEBANGSAAN ST MICHAEL MONGKOS | b | `P.199_N.22_199/22/02_1b` |
| 1126 | `P.199_N.22_199/22/02_1` | BALAI RAYA KPG. MUJAT | c | `P.199_N.22_199/22/02_1c` |
| 1128 | `P.199_N.22_199/22/02_1` | MULTIPURPOSE HALL KPG. MONGKOS | d | `P.199_N.22_199/22/02_1d` |
| 1130 | `P.199_N.22_199/22/02_1` | TADIKA KPG. NIBONG | e | `P.199_N.22_199/22/02_1e` |
| 1131 | `P.199_N.22_199/22/02_1` | SEKOLAH KEBANGSAAN ST NORBERT PAUN GAHAT | f | `P.199_N.22_199/22/02_1f` |

#### 139. `P.199_N.22_199/22/02_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1123 | `P.199_N.22_199/22/02_2` | SEKOLAH KEBANGSAAN MENTU TAPU | a | `P.199_N.22_199/22/02_2a` |
| 1125 | `P.199_N.22_199/22/02_2` | SEKOLAH KEBANGSAAN ST MICHAEL MONGKOS | b | `P.199_N.22_199/22/02_2b` |
| 1127 | `P.199_N.22_199/22/02_2` | BALAI RAYA KPG. MUJAT | c | `P.199_N.22_199/22/02_2c` |
| 1129 | `P.199_N.22_199/22/02_2` | MULTIPURPOSE HALL KPG. MONGKOS | d | `P.199_N.22_199/22/02_2d` |

#### 140. `P.199_N.22_199/22/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1132 | `P.199_N.22_199/22/03_1` | SEKOLAH KEBANGSAAN MAPU | a | `P.199_N.22_199/22/03_1a` |
| 1134 | `P.199_N.22_199/22/03_1` | BALAI RAYA KPG. TERBAT LEBAN | b | `P.199_N.22_199/22/03_1b` |
| 1135 | `P.199_N.22_199/22/03_1` | BALAI RAYA BABUK BAREM KPG. BUNAN GEGA | c | `P.199_N.22_199/22/03_1c` |

#### 141. `P.199_N.22_199/22/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1133 | `P.199_N.22_199/22/03_2` | SEKOLAH KEBANGSAAN MAPU | a | `P.199_N.22_199/22/03_2a` |
| 1136 | `P.199_N.22_199/22/03_2` | BALAI RAYA BABUK BAREM KPG. BUNAN GEGA | b | `P.199_N.22_199/22/03_2b` |

#### 142. `P.199_N.22_199/22/04_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1138 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN KRANGAN | a | `P.199_N.22_199/22/04_1a` |
| 1139 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN MERBAU | b | `P.199_N.22_199/22/04_1b` |
| 1141 | `P.199_N.22_199/22/04_1` | BALAI RAYA KPG. SG. BURU | c | `P.199_N.22_199/22/04_1c` |
| 1142 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN KRAIT | d | `P.199_N.22_199/22/04_1d` |
| 1143 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN SUMPAS | e | `P.199_N.22_199/22/04_1e` |
| 1144 | `P.199_N.22_199/22/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SG MENYAN | f | `P.199_N.22_199/22/04_1f` |
| 1146 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN ENTAYAN | g | `P.199_N.22_199/22/04_1g` |
| 1147 | `P.199_N.22_199/22/04_1` | SEKOLAH KEBANGSAAN SEMUKOI | h | `P.199_N.22_199/22/04_1h` |

#### 143. `P.199_N.22_199/22/04_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1140 | `P.199_N.22_199/22/04_2` | SEKOLAH KEBANGSAAN MERBAU | a | `P.199_N.22_199/22/04_2a` |
| 1145 | `P.199_N.22_199/22/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) SG MENYAN | b | `P.199_N.22_199/22/04_2b` |
| 1148 | `P.199_N.22_199/22/04_2` | SEKOLAH KEBANGSAAN SEMUKOI | c | `P.199_N.22_199/22/04_2c` |

#### 144. `P.199_N.23_199/23/01_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1151 | `P.199_N.23_199/23/01_1` | BALAI RAYA KPG. TEBAKANG DAYAK | a | `P.199_N.23_199/23/01_1a` |
| 1153 | `P.199_N.23_199/23/01_1` | SEKOLAH KEBANGSAAN TEBAKANG | b | `P.199_N.23_199/23/01_1b` |
| 1155 | `P.199_N.23_199/23/01_1` | SEKOLAH KEBANGSAAN PANGKALAN SORAH | c | `P.199_N.23_199/23/01_1c` |
| 1156 | `P.199_N.23_199/23/01_1` | DEWAN SERBAGUNA KPG. SORAK DAYAK | d | `P.199_N.23_199/23/01_1d` |
| 1157 | `P.199_N.23_199/23/01_1` | SEKOLAH KEBANGSAAN KORAN | e | `P.199_N.23_199/23/01_1e` |

#### 145. `P.199_N.23_199/23/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1152 | `P.199_N.23_199/23/01_2` | BALAI RAYA KPG. TEBAKANG DAYAK | a | `P.199_N.23_199/23/01_2a` |
| 1154 | `P.199_N.23_199/23/01_2` | SEKOLAH KEBANGSAAN TEBAKANG | b | `P.199_N.23_199/23/01_2b` |
| 1158 | `P.199_N.23_199/23/01_2` | SEKOLAH KEBANGSAAN KORAN | c | `P.199_N.23_199/23/01_2c` |

#### 146. `P.199_N.23_199/23/02_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1160 | `P.199_N.23_199/23/02_1` | SEKOLAH KEBANGSAAN MERAKAI | a | `P.199_N.23_199/23/02_1a` |
| 1161 | `P.199_N.23_199/23/02_1` | BALAI RAYA KPG. KUALA | b | `P.199_N.23_199/23/02_1b` |
| 1162 | `P.199_N.23_199/23/02_1` | SEKOLAH KEBANGSAAN LEBOR REMUN | c | `P.199_N.23_199/23/02_1c` |
| 1165 | `P.199_N.23_199/23/02_1` | SEKOLAH KEBANGSAAN BEDUP | d | `P.199_N.23_199/23/02_1d` |
| 1168 | `P.199_N.23_199/23/02_1` | SEKOLAH KEBANGSAAN MELANSAI | e | `P.199_N.23_199/23/02_1e` |
| 1169 | `P.199_N.23_199/23/02_1` | SEKOLAH JENIS KEBANGSAAN PANGKALAN BEDUP | f | `P.199_N.23_199/23/02_1f` |

#### 147. `P.199_N.23_199/23/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1163 | `P.199_N.23_199/23/02_2` | SEKOLAH KEBANGSAAN LEBOR REMUN | a | `P.199_N.23_199/23/02_2a` |
| 1166 | `P.199_N.23_199/23/02_2` | SEKOLAH KEBANGSAAN BEDUP | b | `P.199_N.23_199/23/02_2b` |

#### 148. `P.199_N.23_199/23/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1164 | `P.199_N.23_199/23/02_3` | SEKOLAH KEBANGSAAN LEBOR REMUN | a | `P.199_N.23_199/23/02_3a` |
| 1167 | `P.199_N.23_199/23/02_3` | SEKOLAH KEBANGSAAN BEDUP | b | `P.199_N.23_199/23/02_3b` |

#### 149. `P.199_N.23_199/23/03_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1170 | `P.199_N.23_199/23/03_1` | SEKOLAH KEBANGSAAN RIIH DASO | a | `P.199_N.23_199/23/03_1a` |
| 1172 | `P.199_N.23_199/23/03_1` | BALAI RAYA KPG. RIIH MAWANG | b | `P.199_N.23_199/23/03_1b` |
| 1173 | `P.199_N.23_199/23/03_1` | SEKOLAH KEBANGSAAN ST PATRICK TANGGA | c | `P.199_N.23_199/23/03_1c` |
| 1176 | `P.199_N.23_199/23/03_1` | DEWAN SERBAGUNA KPG. RASAU | d | `P.199_N.23_199/23/03_1d` |
| 1177 | `P.199_N.23_199/23/03_1` | SEKOLAH KEBANGSAAN ST HENRY SLABI | e | `P.199_N.23_199/23/03_1e` |
| 1179 | `P.199_N.23_199/23/03_1` | BALAI RAYA KPG. SEROBAN | f | `P.199_N.23_199/23/03_1f` |
| 1180 | `P.199_N.23_199/23/03_1` | SEKOLAH KEBANGSAAN SERIAN | g | `P.199_N.23_199/23/03_1g` |

#### 150. `P.199_N.23_199/23/03_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1171 | `P.199_N.23_199/23/03_2` | SEKOLAH KEBANGSAAN RIIH DASO | a | `P.199_N.23_199/23/03_2a` |
| 1174 | `P.199_N.23_199/23/03_2` | SEKOLAH KEBANGSAAN ST PATRICK TANGGA | b | `P.199_N.23_199/23/03_2b` |
| 1178 | `P.199_N.23_199/23/03_2` | SEKOLAH KEBANGSAAN ST HENRY SLABI | c | `P.199_N.23_199/23/03_2c` |
| 1181 | `P.199_N.23_199/23/03_2` | SEKOLAH KEBANGSAAN SERIAN | d | `P.199_N.23_199/23/03_2d` |

#### 151. `P.200_N.25_200/25/08_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1226 | `P.200_N.25_200/25/08_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SIMUNJAN | a | `P.200_N.25_200/25/08_1a` |
| 1230 | `P.200_N.25_200/25/08_1` | PERPUSTAKAAN KPG. SUAL | b | `P.200_N.25_200/25/08_1b` |
| 1231 | `P.200_N.25_200/25/08_1` | SEKOLAH MENENGAH KEBANGSAAN SRI SADONG | c | `P.200_N.25_200/25/08_1c` |

#### 152. `P.200_N.25_200/25/09_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1232 | `P.200_N.25_200/25/09_1` | SEKOLAH KEBANGSAAN SG LINGKAU | a | `P.200_N.25_200/25/09_1a` |
| 1233 | `P.200_N.25_200/25/09_1` | DEWAN KAMPUNG TEMIANG | b | `P.200_N.25_200/25/09_1b` |
| 1234 | `P.200_N.25_200/25/09_1` | SEKOLAH KEBANGSAAN LEPONG EMPLAS | c | `P.200_N.25_200/25/09_1c` |
| 1235 | `P.200_N.25_200/25/09_1` | SEKOLAH KEBANGSAAN GUNONG NGELI | d | `P.200_N.25_200/25/09_1d` |

#### 153. `P.200_N.25_200/25/10_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1236 | `P.200_N.25_200/25/10_1` | SEKOLAH KEBANGSAAN SAGENG | a | `P.200_N.25_200/25/10_1a` |
| 1239 | `P.200_N.25_200/25/10_1` | SEKOLAH KEBANGSAAN ABANG MAN | b | `P.200_N.25_200/25/10_1b` |

#### 154. `P.200_N.25_200/25/10_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1237 | `P.200_N.25_200/25/10_2` | SEKOLAH KEBANGSAAN SAGENG | a | `P.200_N.25_200/25/10_2a` |
| 1240 | `P.200_N.25_200/25/10_2` | SEKOLAH KEBANGSAAN ABANG MAN | b | `P.200_N.25_200/25/10_2b` |

#### 155. `P.200_N.25_200/25/10_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1238 | `P.200_N.25_200/25/10_3` | SEKOLAH KEBANGSAAN SAGENG | a | `P.200_N.25_200/25/10_3a` |
| 1241 | `P.200_N.25_200/25/10_3` | SEKOLAH KEBANGSAAN ABANG MAN | b | `P.200_N.25_200/25/10_3b` |

#### 156. `P.200_N.26_200/26/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1250 | `P.200_N.26_200/26/03_1` | BALAI RAYA KPG. SABANG | a | `P.200_N.26_200/26/03_1a` |
| 1251 | `P.200_N.26_200/26/03_1` | DEWAN KPG. SATEMAN | b | `P.200_N.26_200/26/03_1b` |
| 1252 | `P.200_N.26_200/26/03_1` | SEKOLAH KEBANGSAAN SUNGAI BA | c | `P.200_N.26_200/26/03_1c` |
| 1253 | `P.200_N.26_200/26/03_1` | BANGSAL SEMENTARA KPG. LUBOK SAMSU | d | `P.200_N.26_200/26/03_1d` |

#### 157. `P.200_N.26_200/26/04_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1254 | `P.200_N.26_200/26/04_1` | SEKOLAH KEBANGSAAN LUBOK PUNGGOR | a | `P.200_N.26_200/26/04_1a` |
| 1255 | `P.200_N.26_200/26/04_1` | SEKOLAH KEBANGSAAN LUBOK BUNTIN | b | `P.200_N.26_200/26/04_1b` |
| 1256 | `P.200_N.26_200/26/04_1` | SEKOLAH KEBANGSAAN KG GUMPEY | c | `P.200_N.26_200/26/04_1c` |
| 1257 | `P.200_N.26_200/26/04_1` | DEWAN KAMPUNG SPAOH GEDONG | d | `P.200_N.26_200/26/04_1d` |
| 1258 | `P.200_N.26_200/26/04_1` | DEWAN KPG. BENAT ULU | e | `P.200_N.26_200/26/04_1e` |

#### 158. `P.200_N.26_200/26/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1266 | `P.200_N.26_200/26/06_1` | SEKOLAH KEBANGSAAN TEGELAM | a | `P.200_N.26_200/26/06_1a` |
| 1267 | `P.200_N.26_200/26/06_1` | BALAI RAYA SG. ALIT | b | `P.200_N.26_200/26/06_1b` |

#### 159. `P.200_N.26_200/26/07_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1268 | `P.200_N.26_200/26/07_1` | RH. SAMAU ANAK JELIOH MUNGGU AIR | a | `P.200_N.26_200/26/07_1a` |
| 1269 | `P.200_N.26_200/26/07_1` | SEKOLAH KEBANGSAAN MUNGGU LALLANG | b | `P.200_N.26_200/26/07_1b` |
| 1270 | `P.200_N.26_200/26/07_1` | SEKOLAH KEBANGSAAN KENIONG | c | `P.200_N.26_200/26/07_1c` |
| 1271 | `P.200_N.26_200/26/07_1` | DEWAN ISU JAYA | d | `P.200_N.26_200/26/07_1d` |
| 1272 | `P.200_N.26_200/26/07_1` | SEKOLAH KEBANGSAAN SEMALATONG | e | `P.200_N.26_200/26/07_1e` |

#### 160. `P.201_N.27_201/27/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1276 | `P.201_N.27_201/27/02_1` | SEKOLAH KEBANGSAAN SG LADONG | a | `P.201_N.27_201/27/02_1a` |
| 1279 | `P.201_N.27_201/27/02_1` | BALAI RAYA SG. SEGALI | b | `P.201_N.27_201/27/02_1b` |
| 1280 | `P.201_N.27_201/27/02_1` | SEKOLAH KEBANGSAAN HJ BUJANG SEBANGAN | c | `P.201_N.27_201/27/02_1c` |
| 1282 | `P.201_N.27_201/27/02_1` | DEWAN MASYARAKAT SEBANGAN SAMPAT | d | `P.201_N.27_201/27/02_1d` |

#### 161. `P.201_N.27_201/27/02_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1277 | `P.201_N.27_201/27/02_2` | SEKOLAH KEBANGSAAN SG LADONG | a | `P.201_N.27_201/27/02_2a` |
| 1281 | `P.201_N.27_201/27/02_2` | SEKOLAH KEBANGSAAN HJ BUJANG SEBANGAN | b | `P.201_N.27_201/27/02_2b` |
| 1283 | `P.201_N.27_201/27/02_2` | DEWAN MASYARAKAT SEBANGAN SAMPAT | c | `P.201_N.27_201/27/02_2c` |

#### 162. `P.201_N.27_201/27/03_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1284 | `P.201_N.27_201/27/03_1` | SEKOLAH KEBANGSAAN BAJONG | a | `P.201_N.27_201/27/03_1a` |
| 1285 | `P.201_N.27_201/27/03_1` | SEKOLAH KEBANGSAAN TEBELU | b | `P.201_N.27_201/27/03_1b` |
| 1288 | `P.201_N.27_201/27/03_1` | ASTAKA SEKOLAH KEBANGSAAN KLAIT | c | `P.201_N.27_201/27/03_1c` |
| 1289 | `P.201_N.27_201/27/03_1` | SEKOLAH KEBANGSAAN TUANKU BAGUS | d | `P.201_N.27_201/27/03_1d` |
| 1293 | `P.201_N.27_201/27/03_1` | SURAU KPG. SEBERANG SEBUYAU | e | `P.201_N.27_201/27/03_1e` |
| 1295 | `P.201_N.27_201/27/03_1` | SEKOLAH MENENGAH KEBANGSAAN SEBUYAU | f | `P.201_N.27_201/27/03_1f` |
| 1301 | `P.201_N.27_201/27/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SEBUYAU | g | `P.201_N.27_201/27/03_1g` |

#### 163. `P.201_N.27_201/27/03_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1286 | `P.201_N.27_201/27/03_2` | SEKOLAH KEBANGSAAN TEBELU | a | `P.201_N.27_201/27/03_2a` |
| 1290 | `P.201_N.27_201/27/03_2` | SEKOLAH KEBANGSAAN TUANKU BAGUS | b | `P.201_N.27_201/27/03_2b` |
| 1294 | `P.201_N.27_201/27/03_2` | SURAU KPG. SEBERANG SEBUYAU | c | `P.201_N.27_201/27/03_2c` |
| 1296 | `P.201_N.27_201/27/03_2` | SEKOLAH MENENGAH KEBANGSAAN SEBUYAU | d | `P.201_N.27_201/27/03_2d` |

#### 164. `P.201_N.27_201/27/03_3` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1287 | `P.201_N.27_201/27/03_3` | SEKOLAH KEBANGSAAN TEBELU | a | `P.201_N.27_201/27/03_3a` |
| 1291 | `P.201_N.27_201/27/03_3` | SEKOLAH KEBANGSAAN TUANKU BAGUS | b | `P.201_N.27_201/27/03_3b` |
| 1297 | `P.201_N.27_201/27/03_3` | SEKOLAH MENENGAH KEBANGSAAN SEBUYAU | c | `P.201_N.27_201/27/03_3c` |

#### 165. `P.201_N.27_201/27/03_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1292 | `P.201_N.27_201/27/03_4` | SEKOLAH KEBANGSAAN TUANKU BAGUS | a | `P.201_N.27_201/27/03_4a` |
| 1298 | `P.201_N.27_201/27/03_4` | SEKOLAH MENENGAH KEBANGSAAN SEBUYAU | b | `P.201_N.27_201/27/03_4b` |

#### 166. `P.201_N.27_201/27/05_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1303 | `P.201_N.27_201/27/05_1` | SEKOLAH KEBANGSAAN LUNYING | a | `P.201_N.27_201/27/05_1a` |
| 1304 | `P.201_N.27_201/27/05_1` | DEWAN SEKOLAH KEBANGSAAN BULAN JERAGAM | b | `P.201_N.27_201/27/05_1b` |
| 1305 | `P.201_N.27_201/27/05_1` | RH. BELAYONG SG. NYAMOK | c | `P.201_N.27_201/27/05_1c` |
| 1306 | `P.201_N.27_201/27/05_1` | DEWAN SERBAGUNA SG. RAMA | d | `P.201_N.27_201/27/05_1d` |
| 1307 | `P.201_N.27_201/27/05_1` | SEKOLAH KEBANGSAAN RABA | e | `P.201_N.27_201/27/05_1e` |

#### 167. `P.201_N.27_201/27/06_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1308 | `P.201_N.27_201/27/06_1` | SEKOLAH KEBANGSAAN SUNGAI ARUS LUMUT | a | `P.201_N.27_201/27/06_1a` |
| 1309 | `P.201_N.27_201/27/06_1` | SEKOLAH KEBANGSAAN TUNGKAH MELAYU | b | `P.201_N.27_201/27/06_1b` |
| 1312 | `P.201_N.27_201/27/06_1` | SEKOLAH KEBANGSAAN TUNGKAH DAYAK | c | `P.201_N.27_201/27/06_1c` |
| 1313 | `P.201_N.27_201/27/06_1` | SEKOLAH KEBANGSAAN RAJAU ENSIKA | d | `P.201_N.27_201/27/06_1d` |
| 1314 | `P.201_N.27_201/27/06_1` | SEKOLAH KEBANGSAAN SKITONG/MERANTI | e | `P.201_N.27_201/27/06_1e` |
| 1315 | `P.201_N.27_201/27/06_1` | DEWAN KPG. STIKA | f | `P.201_N.27_201/27/06_1f` |

#### 168. `P.201_N.28_201/28/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1318 | `P.201_N.28_201/28/01_1` | SEKOLAH KEBANGSAAN ST DUNSTAN | a | `P.201_N.28_201/28/01_1a` |
| 1319 | `P.201_N.28_201/28/01_1` | SEKOLAH KEBANGSAAN LELA PAHLAWAN | b | `P.201_N.28_201/28/01_1b` |

#### 169. `P.201_N.28_201/28/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1322 | `P.201_N.28_201/28/02_1` | DEWAN KPG. SG. MULON MELUDAM | a | `P.201_N.28_201/28/02_1a` |
| 1323 | `P.201_N.28_201/28/02_1` | DEWAN WANITA KPG. TRISO | b | `P.201_N.28_201/28/02_1b` |
| 1324 | `P.201_N.28_201/28/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MALUDAM | c | `P.201_N.28_201/28/02_1c` |
| 1327 | `P.201_N.28_201/28/02_1` | SEKOLAH KEBANGSAAN MALUDAM | d | `P.201_N.28_201/28/02_1d` |

#### 170. `P.201_N.28_201/28/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1325 | `P.201_N.28_201/28/02_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MALUDAM | a | `P.201_N.28_201/28/02_2a` |
| 1328 | `P.201_N.28_201/28/02_2` | SEKOLAH KEBANGSAAN MALUDAM | b | `P.201_N.28_201/28/02_2b` |

#### 171. `P.201_N.28_201/28/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1326 | `P.201_N.28_201/28/02_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MALUDAM | a | `P.201_N.28_201/28/02_3a` |
| 1329 | `P.201_N.28_201/28/02_3` | SEKOLAH KEBANGSAAN MALUDAM | b | `P.201_N.28_201/28/02_3b` |

#### 172. `P.201_N.28_201/28/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1332 | `P.201_N.28_201/28/03_1` | BANGUNAN TADIKA KEMAS SEDUKU BARU | a | `P.201_N.28_201/28/03_1a` |
| 1333 | `P.201_N.28_201/28/03_1` | RH. RADIN PUTAT | b | `P.201_N.28_201/28/03_1b` |
| 1334 | `P.201_N.28_201/28/03_1` | RH. PIPIT (SEBAKAK) TG. BIJAT | c | `P.201_N.28_201/28/03_1c` |

#### 173. `P.201_N.28_201/28/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1335 | `P.201_N.28_201/28/04_1` | KLINIK DESA TANJUNG BIJAT | a | `P.201_N.28_201/28/04_1a` |
| 1336 | `P.201_N.28_201/28/04_1` | SEKOLAH MENENGAH KEBANGSAAN LINGGA | b | `P.201_N.28_201/28/04_1b` |
| 1337 | `P.201_N.28_201/28/04_1` | SEKOLAH KEBANGSAAN GRAN/STUMBIN | c | `P.201_N.28_201/28/04_1c` |

#### 174. `P.201_N.29_201/29/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1351 | `P.201_N.29_201/29/01_1` | SEKOLAH KEBANGSAAN BATANG MARO | a | `P.201_N.29_201/29/01_1a` |
| 1352 | `P.201_N.29_201/29/01_1` | SEKOLAH KEBANGSAAN SPINANG | b | `P.201_N.29_201/29/01_1b` |
| 1353 | `P.201_N.29_201/29/01_1` | SEKOLAH MENENGAH KEBANGSAAN BELADIN | c | `P.201_N.29_201/29/01_1c` |
| 1362 | `P.201_N.29_201/29/01_1` | SEKOLAH KEBANGSAAN SEMARANG | d | `P.201_N.29_201/29/01_1d` |

#### 175. `P.201_N.29_201/29/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1354 | `P.201_N.29_201/29/01_2` | SEKOLAH MENENGAH KEBANGSAAN BELADIN | a | `P.201_N.29_201/29/01_2a` |
| 1363 | `P.201_N.29_201/29/01_2` | SEKOLAH KEBANGSAAN SEMARANG | b | `P.201_N.29_201/29/01_2b` |

#### 176. `P.201_N.29_201/29/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1364 | `P.201_N.29_201/29/02_1` | SEKOLAH KEBANGSAAN TAMBAK | a | `P.201_N.29_201/29/02_1a` |
| 1367 | `P.201_N.29_201/29/02_1` | RH. DUAT ANAK AJI BLOK DEBAK | b | `P.201_N.29_201/29/02_1b` |
| 1368 | `P.201_N.29_201/29/02_1` | SEKOLAH KEBANGSAAN MUTON | c | `P.201_N.29_201/29/02_1c` |
| 1370 | `P.201_N.29_201/29/02_1` | DEWAN KAMPUNG KALOK | d | `P.201_N.29_201/29/02_1d` |

#### 177. `P.201_N.29_201/29/02_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1365 | `P.201_N.29_201/29/02_2` | SEKOLAH KEBANGSAAN TAMBAK | a | `P.201_N.29_201/29/02_2a` |
| 1369 | `P.201_N.29_201/29/02_2` | SEKOLAH KEBANGSAAN MUTON | b | `P.201_N.29_201/29/02_2b` |
| 1371 | `P.201_N.29_201/29/02_2` | DEWAN KAMPUNG KALOK | c | `P.201_N.29_201/29/02_2c` |

#### 178. `P.202_N.30_202/30/01_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1381 | `P.202_N.30_202/30/01_1` | BALAI RAYA KPG. KESINDU SIMUNJAN | a | `P.202_N.30_202/30/01_1a` |
| 1382 | `P.202_N.30_202/30/01_1` | RH. SUNTAT GALUNG KPG. SEBANGKOI JAYA SIMUNJAN | b | `P.202_N.30_202/30/01_1b` |
| 1383 | `P.202_N.30_202/30/01_1` | RH. GANTI ANAK JETAI, KAMPUNG SABAL KRUIN, SIMUNJAN | c | `P.202_N.30_202/30/01_1c` |
| 1384 | `P.202_N.30_202/30/01_1` | SEKOLAH KEBANGSAAN SUNGAI PINANG | d | `P.202_N.30_202/30/01_1d` |
| 1385 | `P.202_N.30_202/30/01_1` | SEKOLAH KEBANGSAAN GAWANG EMPILI | e | `P.202_N.30_202/30/01_1e` |
| 1386 | `P.202_N.30_202/30/01_1` | SEKOLAH KEBANGSAAN NYELITAK | f | `P.202_N.30_202/30/01_1f` |
| 1387 | `P.202_N.30_202/30/01_1` | SEKOLAH KEBANGSAAN TELAGUS/JEROK | g | `P.202_N.30_202/30/01_1g` |
| 1388 | `P.202_N.30_202/30/01_1` | SABAL NURSERY CENTRE | h | `P.202_N.30_202/30/01_1h` |

#### 179. `P.202_N.30_202/30/02_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1389 | `P.202_N.30_202/30/02_1` | SEKOLAH KEBANGSAAN SEMADA | a | `P.202_N.30_202/30/02_1a` |
| 1390 | `P.202_N.30_202/30/02_1` | RH. RICKY ANAK LUNSA, KPG. ENSEBANG KUARI | b | `P.202_N.30_202/30/02_1b` |
| 1391 | `P.202_N.30_202/30/02_1` | SEKOLAH KEBANGSAAN BALAI RINGIN | c | `P.202_N.30_202/30/02_1c` |
| 1393 | `P.202_N.30_202/30/02_1` | RH. RENTAB ANAK GASAN MELIKIN SG. PANGGIL | d | `P.202_N.30_202/30/02_1d` |
| 1394 | `P.202_N.30_202/30/02_1` | RH. ANING ANAK MELINTANG ENSEBANG PELAI | e | `P.202_N.30_202/30/02_1e` |
| 1395 | `P.202_N.30_202/30/02_1` | BALAI RAYA TANAH MAWANG | f | `P.202_N.30_202/30/02_1f` |

#### 180. `P.202_N.30_202/30/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1396 | `P.202_N.30_202/30/03_1` | DEWAN SERBAGUNA KPG. KEPIT BUKIT PUNDA SIMUNJAN | a | `P.202_N.30_202/30/03_1a` |
| 1397 | `P.202_N.30_202/30/03_1` | BALAI RAYA, KPG. SPAOH RABA, SIMUNJAN | b | `P.202_N.30_202/30/03_1b` |
| 1398 | `P.202_N.30_202/30/03_1` | DEWAN SRI KUMANG SEKOLAH KEBANGSAAN PADANG PEDALAI | c | `P.202_N.30_202/30/03_1c` |

#### 181. `P.202_N.30_202/30/04_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1399 | `P.202_N.30_202/30/04_1` | RH. ENTERUS KPG. SEGA | a | `P.202_N.30_202/30/04_1a` |
| 1400 | `P.202_N.30_202/30/04_1` | SEKOLAH KEBANGSAAN KEDUMPAI | b | `P.202_N.30_202/30/04_1b` |
| 1401 | `P.202_N.30_202/30/04_1` | RH. UMBANG ANAK KIOT, KAMPUNG PENDAWAN, SIMUNJAN | c | `P.202_N.30_202/30/04_1c` |
| 1402 | `P.202_N.30_202/30/04_1` | SEKOLAH KEBANGSAAN TUBA | d | `P.202_N.30_202/30/04_1d` |
| 1403 | `P.202_N.30_202/30/04_1` | SEKOLAH KEBANGSAAN MENTU | e | `P.202_N.30_202/30/04_1e` |
| 1404 | `P.202_N.30_202/30/04_1` | RH. MENARI IPOH | f | `P.202_N.30_202/30/04_1f` |
| 1405 | `P.202_N.30_202/30/04_1` | RH. MENJAT SEMAWA ILI | g | `P.202_N.30_202/30/04_1g` |
| 1406 | `P.202_N.30_202/30/04_1` | SEKOLAH KEBANGSAAN MUDING | h | `P.202_N.30_202/30/04_1h` |

#### 182. `P.202_N.30_202/30/05_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1407 | `P.202_N.30_202/30/05_1` | RH. NGULU EMPALING A | a | `P.202_N.30_202/30/05_1a` |
| 1408 | `P.202_N.30_202/30/05_1` | RH. JAWAN TEKUYONG B | b | `P.202_N.30_202/30/05_1b` |
| 1409 | `P.202_N.30_202/30/05_1` | SEKOLAH KEBANGSAAN JAONG | c | `P.202_N.30_202/30/05_1c` |
| 1410 | `P.202_N.30_202/30/05_1` | SEKOLAH KEBANGSAAN ABOK | d | `P.202_N.30_202/30/05_1d` |
| 1411 | `P.202_N.30_202/30/05_1` | SEKOLAH KEBANGSAAN APING | e | `P.202_N.30_202/30/05_1e` |
| 1412 | `P.202_N.30_202/30/05_1` | RH. LANGIE ANAK TAWIE SEBEMBAN | f | `P.202_N.30_202/30/05_1f` |
| 1413 | `P.202_N.30_202/30/05_1` | SEKOLAH KEBANGSAAN ST LEO GAYAU | g | `P.202_N.30_202/30/05_1g` |

#### 183. `P.202_N.30_202/30/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1414 | `P.202_N.30_202/30/06_1` | TADIKA KEMAS RAPAK | a | `P.202_N.30_202/30/06_1a` |
| 1415 | `P.202_N.30_202/30/06_1` | RH. HEAROLD SELANTEK ILI PANTU | b | `P.202_N.30_202/30/06_1b` |
| 1416 | `P.202_N.30_202/30/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PANTU | c | `P.202_N.30_202/30/06_1c` |
| 1417 | `P.202_N.30_202/30/06_1` | SEKOLAH KEBANGSAAN PANTU | d | `P.202_N.30_202/30/06_1d` |

#### 184. `P.202_N.30_202/30/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1418 | `P.202_N.30_202/30/07_1` | SEKOLAH KEBANGSAAN KERANGGAS | a | `P.202_N.30_202/30/07_1a` |
| 1419 | `P.202_N.30_202/30/07_1` | PRA-SEKOLAH KPG. SAPAK | b | `P.202_N.30_202/30/07_1b` |

#### 185. `P.202_N.31_202/31/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1421 | `P.202_N.31_202/31/00_1` | SEKOLAH KEBANGSAAN TEMUDOK KEM | a | `P.202_N.31_202/31/00_1a` |
| 1422 | `P.202_N.31_202/31/00_1` | DEWAN MAKAN 13 RAMD KEM PAKIT | b | `P.202_N.31_202/31/00_1b` |

#### 186. `P.202_N.31_202/31/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1424 | `P.202_N.31_202/31/01_1` | BALAI RAYA BANTING | a | `P.202_N.31_202/31/01_1a` |
| 1426 | `P.202_N.31_202/31/01_1` | SEKOLAH KEBANGSAAN ENGKRANJI | b | `P.202_N.31_202/31/01_1b` |
| 1427 | `P.202_N.31_202/31/01_1` | RH. EDWARD MAMUT LANGGIR LINGGA | c | `P.202_N.31_202/31/01_1c` |
| 1428 | `P.202_N.31_202/31/01_1` | RH. JACK ENGKEREPOK LINGGA | d | `P.202_N.31_202/31/01_1d` |

#### 187. `P.202_N.31_202/31/02_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1429 | `P.202_N.31_202/31/02_1` | RH. GUANG, KARA PANTU | a | `P.202_N.31_202/31/02_1a` |
| 1430 | `P.202_N.31_202/31/02_1` | RH. REKIE ANAK SUMPIT, KPG. PUNGGU TENGAH | b | `P.202_N.31_202/31/02_1b` |
| 1431 | `P.202_N.31_202/31/02_1` | RH. FRANCIS DOBLIN ANAK BETOL SELANJAN ANGKONG | c | `P.202_N.31_202/31/02_1c` |
| 1432 | `P.202_N.31_202/31/02_1` | SEKOLAH KEBANGSAAN SELANJAN | d | `P.202_N.31_202/31/02_1d` |
| 1433 | `P.202_N.31_202/31/02_1` | RH. JOHNNY ANAK GURA ENCHIAP BARU | e | `P.202_N.31_202/31/02_1e` |
| 1434 | `P.202_N.31_202/31/02_1` | SEKOLAH KEBANGSAAN ST MARTIN | f | `P.202_N.31_202/31/02_1f` |
| 1435 | `P.202_N.31_202/31/02_1` | RH. TANGGU KUBAU ILI | g | `P.202_N.31_202/31/02_1g` |
| 1436 | `P.202_N.31_202/31/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BANGKONG | h | `P.202_N.31_202/31/02_1h` |

#### 188. `P.202_N.31_202/31/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1437 | `P.202_N.31_202/31/03_1` | DEWAN MASYARAKAT GUA | a | `P.202_N.31_202/31/03_1a` |
| 1438 | `P.202_N.31_202/31/03_1` | RH. JAWAN PANGGIL | b | `P.202_N.31_202/31/03_1b` |
| 1439 | `P.202_N.31_202/31/03_1` | BALAI RAYA RH. ANYAI ENTULANG | c | `P.202_N.31_202/31/03_1c` |
| 1440 | `P.202_N.31_202/31/03_1` | TEMUDOK AGRICULTURE STATION | d | `P.202_N.31_202/31/03_1d` |

#### 189. `P.202_N.31_202/31/04_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1441 | `P.202_N.31_202/31/04_1` | RH. KEDENI, BATU BESAI | a | `P.202_N.31_202/31/04_1a` |
| 1442 | `P.202_N.31_202/31/04_1` | BALAI RAYA RH. BERIKU PAKIT | b | `P.202_N.31_202/31/04_1b` |
| 1443 | `P.202_N.31_202/31/04_1` | SEKOLAH KEBANGSAAN SELEPONG | c | `P.202_N.31_202/31/04_1c` |
| 1444 | `P.202_N.31_202/31/04_1` | SEKOLAH MENENGAH KEBANGSAAN MELUGU | d | `P.202_N.31_202/31/04_1d` |
| 1446 | `P.202_N.31_202/31/04_1` | RH. JUMAT PO AI ENGGAT | e | `P.202_N.31_202/31/04_1e` |

#### 190. `P.202_N.31_202/31/05_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1447 | `P.202_N.31_202/31/05_1` | SEKOLAH KEBANGSAAN BAKONG | a | `P.202_N.31_202/31/05_1a` |
| 1449 | `P.202_N.31_202/31/05_1` | RH. SAWING ANAK BIJU, SUNGAI TAPANG, JUNGKONG BALAU | b | `P.202_N.31_202/31/05_1b` |
| 1450 | `P.202_N.31_202/31/05_1` | RH. RESA ANAK GAYAU EMPELANJAU ASAL BAKONG | c | `P.202_N.31_202/31/05_1c` |
| 1451 | `P.202_N.31_202/31/05_1` | SEKOLAH KEBANGSAAN TG BIJAT | d | `P.202_N.31_202/31/05_1d` |

#### 191. `P.202_N.31_202/31/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1448 | `P.202_N.31_202/31/05_2` | SEKOLAH KEBANGSAAN BAKONG | a | `P.202_N.31_202/31/05_2a` |
| 1452 | `P.202_N.31_202/31/05_2` | SEKOLAH KEBANGSAAN TG BIJAT | b | `P.202_N.31_202/31/05_2b` |

#### 192. `P.202_N.32_202/32/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1457 | `P.202_N.32_202/32/02_1` | SEKOLAH KEBANGSAAN ABANG AING | a | `P.202_N.32_202/32/02_1a` |
| 1466 | `P.202_N.32_202/32/02_1` | SEKOLAH KEBANGSAAN SRI AMAN | b | `P.202_N.32_202/32/02_1b` |
| 1469 | `P.202_N.32_202/32/02_1` | DEWAN SEKOLAH MENENGAH KEBANGSAAN SRI AMAN | c | `P.202_N.32_202/32/02_1c` |

#### 193. `P.202_N.32_202/32/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1458 | `P.202_N.32_202/32/02_2` | SEKOLAH KEBANGSAAN ABANG AING | a | `P.202_N.32_202/32/02_2a` |
| 1467 | `P.202_N.32_202/32/02_2` | SEKOLAH KEBANGSAAN SRI AMAN | b | `P.202_N.32_202/32/02_2b` |

#### 194. `P.202_N.32_202/32/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1459 | `P.202_N.32_202/32/02_3` | SEKOLAH KEBANGSAAN ABANG AING | a | `P.202_N.32_202/32/02_3a` |
| 1468 | `P.202_N.32_202/32/02_3` | SEKOLAH KEBANGSAAN SRI AMAN | b | `P.202_N.32_202/32/02_3b` |

#### 195. `P.202_N.32_202/32/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1470 | `P.202_N.32_202/32/03_1` | RH. THOMAS RAMBA BRANGAN ULU | a | `P.202_N.32_202/32/03_1a` |
| 1471 | `P.202_N.32_202/32/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SIMANGGANG | b | `P.202_N.32_202/32/03_1b` |

#### 196. `P.202_N.32_202/32/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1477 | `P.202_N.32_202/32/05_1` | RH. RENGKANG LEPONG EMPELIAU | a | `P.202_N.32_202/32/05_1a` |
| 1478 | `P.202_N.32_202/32/05_1` | RH. LUTANG, SENGELAU UNDOP | b | `P.202_N.32_202/32/05_1b` |
| 1479 | `P.202_N.32_202/32/05_1` | SEKOLAH KEBANGSAAN NG. KLASSEN | c | `P.202_N.32_202/32/05_1c` |

#### 197. `P.202_N.32_202/32/06_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1480 | `P.202_N.32_202/32/06_1` | RH. THOMAS LAMAN SENGKUANG | a | `P.202_N.32_202/32/06_1a` |
| 1481 | `P.202_N.32_202/32/06_1` | RH. BANDI SEBANGKOI UNDOP | b | `P.202_N.32_202/32/06_1b` |
| 1482 | `P.202_N.32_202/32/06_1` | SEKOLAH KEBANGSAAN BATU LINTANG | c | `P.202_N.32_202/32/06_1c` |
| 1484 | `P.202_N.32_202/32/06_1` | SEKOLAH KEBANGSAAN PAKU | d | `P.202_N.32_202/32/06_1d` |
| 1485 | `P.202_N.32_202/32/06_1` | RH. PILIT KAONG | e | `P.202_N.32_202/32/06_1e` |

#### 198. `P.202_N.32_202/32/07_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1486 | `P.202_N.32_202/32/07_1` | RH. EMPOL SG. TENGGAK | a | `P.202_N.32_202/32/07_1a` |
| 1487 | `P.202_N.32_202/32/07_1` | BALAI RAYA RH LICHANG, KPG ENTAWA, UNDOP, SRI AMAN | b | `P.202_N.32_202/32/07_1b` |
| 1488 | `P.202_N.32_202/32/07_1` | BANGUNAN TABIKA KEMAS, KPG SIGA JALAN PAIP, UNDOP, SRI AMAN | c | `P.202_N.32_202/32/07_1c` |

#### 199. `P.203_N.33_203/33/10_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1506 | `P.203_N.33_203/33/10_1` | SEKOLAH KEBANGSAAN NG MENJUAU | a | `P.203_N.33_203/33/10_1a` |
| 1507 | `P.203_N.33_203/33/10_1` | RH GAIT, KERANGGAS | b | `P.203_N.33_203/33/10_1b` |

#### 200. `P.203_N.34_203/34/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1532 | `P.203_N.34_203/34/04_1` | SEKOLAH KEBANGSAAN NANGA KESIT | a | `P.203_N.34_203/34/04_1a` |
| 1534 | `P.203_N.34_203/34/04_1` | RH. LIMPENG LEPONG MAWANG KESIT | b | `P.203_N.34_203/34/04_1b` |

#### 201. `P.203_N.34_203/34/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1537 | `P.203_N.34_203/34/07_1` | TADIKA KEMAS KAONG ULU | a | `P.203_N.34_203/34/07_1a` |
| 1538 | `P.203_N.34_203/34/07_1` | PUSAT SUMBER NYEMUNGAN SIMPANG | b | `P.203_N.34_203/34/07_1b` |

#### 202. `P.203_N.34_203/34/14_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1549 | `P.203_N.34_203/34/14_1` | TADIKA KEMAS NG. JELA | a | `P.203_N.34_203/34/14_1a` |
| 1550 | `P.203_N.34_203/34/14_1` | NG. TELAUS SPS | b | `P.203_N.34_203/34/14_1b` |

#### 203. `P.203_N.34_203/34/20_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1557 | `P.203_N.34_203/34/20_1` | RH JELI, LUBOK PAYAN | a | `P.203_N.34_203/34/20_1a` |
| 1558 | `P.203_N.34_203/34/20_1` | RH. BAKAR, BAWI PATOH | b | `P.203_N.34_203/34/20_1b` |

#### 204. `P.204_N.36_204/36/08_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1613 | `P.204_N.36_204/36/08_1` | RH. JAWA NG. JELAU | a | `P.204_N.36_204/36/08_1a` |
| 1614 | `P.204_N.36_204/36/08_1` | RH. GUMBANG JELAU ATAS | b | `P.204_N.36_204/36/08_1b` |

#### 205. `P.204_N.36_204/36/17_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1624 | `P.204_N.36_204/36/17_1` | RH. SANG SG. SIBAU | a | `P.204_N.36_204/36/17_1a` |
| 1625 | `P.204_N.36_204/36/17_1` | RH. JOHN RAGAI BUAI MELANJAN | b | `P.204_N.36_204/36/17_1b` |

#### 206. `P.204_N.36_204/36/19_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1627 | `P.204_N.36_204/36/19_1` | RH. PAGAH PELEPOK | a | `P.204_N.36_204/36/19_1a` |
| 1628 | `P.204_N.36_204/36/19_1` | RH. EMPENI BUNGKANG | b | `P.204_N.36_204/36/19_1b` |

#### 207. `P.204_N.37_204/37/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1644 | `P.204_N.37_204/37/05_1` | RH. MANGGI TANJONG | a | `P.204_N.37_204/37/05_1a` |
| 1645 | `P.204_N.37_204/37/05_1` | RH. JAMES JUGOL BELABAK | b | `P.204_N.37_204/37/05_1b` |

#### 208. `P.204_N.37_204/37/11_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1651 | `P.204_N.37_204/37/11_1` | SEKOLAH KEBANGSAAN NG. GAYAU | a | `P.204_N.37_204/37/11_1a` |
| 1652 | `P.204_N.37_204/37/11_1` | RH. DANGGAT BABU TENGAH | b | `P.204_N.37_204/37/11_1b` |

#### 209. `P.204_N.37_204/37/12_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1653 | `P.204_N.37_204/37/12_1` | RH. STANG BABU ULU | a | `P.204_N.37_204/37/12_1a` |
| 1654 | `P.204_N.37_204/37/12_1` | RH. ACHON NG. TOT | b | `P.204_N.37_204/37/12_1b` |

#### 210. `P.205_N.38_205/38/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1685 | `P.205_N.38_205/38/01_1` | SEKOLAH KEBANGSAAN SG NYIAR | a | `P.205_N.38_205/38/01_1a` |
| 1686 | `P.205_N.38_205/38/01_1` | RH. BUNDAN SG. METONG ROBAN | b | `P.205_N.38_205/38/01_1b` |
| 1687 | `P.205_N.38_205/38/01_1` | RH. SITI JUBAIDAH ANAK LANCHANG, SG. ENGKABANG | c | `P.205_N.38_205/38/01_1c` |
| 1688 | `P.205_N.38_205/38/01_1` | SEKOLAH MENENGAH KEBANGSAAN KALAKA | d | `P.205_N.38_205/38/01_1d` |

#### 211. `P.205_N.38_205/38/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1689 | `P.205_N.38_205/38/02_1` | SEKOLAH KEBANGSAAN LUBOK NIBONG | a | `P.205_N.38_205/38/02_1a` |
| 1690 | `P.205_N.38_205/38/02_1` | SEKOLAH KEBANGSAAN KG PALOH | b | `P.205_N.38_205/38/02_1b` |
| 1691 | `P.205_N.38_205/38/02_1` | SEKOLAH KEBANGSAAN PERPAT | c | `P.205_N.38_205/38/02_1c` |

#### 212. `P.205_N.38_205/38/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1692 | `P.205_N.38_205/38/03_1` | SEKOLAH KEBANGSAAN HAJI BOLLAH | a | `P.205_N.38_205/38/03_1a` |
| 1693 | `P.205_N.38_205/38/03_1` | SEKOLAH KEBANGSAAN KG KUPANG | b | `P.205_N.38_205/38/03_1b` |
| 1695 | `P.205_N.38_205/38/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) MIN SYN SARATOK | c | `P.205_N.38_205/38/03_1c` |

#### 213. `P.205_N.38_205/38/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1694 | `P.205_N.38_205/38/03_2` | SEKOLAH KEBANGSAAN KG KUPANG | a | `P.205_N.38_205/38/03_2a` |
| 1696 | `P.205_N.38_205/38/03_2` | SEKOLAH JENIS KEBANGSAAN (CINA) MIN SYN SARATOK | b | `P.205_N.38_205/38/03_2b` |

#### 214. `P.205_N.39_205/39/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1708 | `P.205_N.39_205/39/01_1` | RH. JIMBUN KOP IBUS | a | `P.205_N.39_205/39/01_1a` |
| 1709 | `P.205_N.39_205/39/01_1` | SEKOLAH KEBANGSAAN MUDONG | b | `P.205_N.39_205/39/01_1b` |
| 1710 | `P.205_N.39_205/39/01_1` | SEKOLAH KEBANGSAAN SUPOK | c | `P.205_N.39_205/39/01_1c` |
| 1711 | `P.205_N.39_205/39/01_1` | RH. MAWAT, BRATONG | d | `P.205_N.39_205/39/01_1d` |

#### 215. `P.205_N.39_205/39/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1712 | `P.205_N.39_205/39/02_1` | SEKOLAH KEBANGSAAN BERAYANG | a | `P.205_N.39_205/39/02_1a` |
| 1714 | `P.205_N.39_205/39/02_1` | SEKOLAH KEBANGSAAN LICHOK | b | `P.205_N.39_205/39/02_1b` |

#### 216. `P.205_N.39_205/39/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1713 | `P.205_N.39_205/39/02_2` | SEKOLAH KEBANGSAAN BERAYANG | a | `P.205_N.39_205/39/02_2a` |
| 1715 | `P.205_N.39_205/39/02_2` | SEKOLAH KEBANGSAAN LICHOK | b | `P.205_N.39_205/39/02_2b` |

#### 217. `P.205_N.39_205/39/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1716 | `P.205_N.39_205/39/03_1` | SEKOLAH KEBANGSAAN ULU SEBETAN | a | `P.205_N.39_205/39/03_1a` |
| 1717 | `P.205_N.39_205/39/03_1` | SEKOLAH KEBANGSAAN KALAKA CENTRAL | b | `P.205_N.39_205/39/03_1b` |
| 1719 | `P.205_N.39_205/39/03_1` | SEKOLAH KEBANGSAAN ENGKUDU | c | `P.205_N.39_205/39/03_1c` |
| 1720 | `P.205_N.39_205/39/03_1` | SEKOLAH KEBANGSAAN SUNGAI KLAMPAI | d | `P.205_N.39_205/39/03_1d` |

#### 218. `P.205_N.39_205/39/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1721 | `P.205_N.39_205/39/04_1` | SEKOLAH KEBANGSAAN TEMUDOK AWIK | a | `P.205_N.39_205/39/04_1a` |
| 1722 | `P.205_N.39_205/39/04_1` | SEKOLAH KEBANGSAAN ST PETER SARATOK | b | `P.205_N.39_205/39/04_1b` |
| 1725 | `P.205_N.39_205/39/04_1` | SEKOLAH KEBANGSAAN NG MALONG | c | `P.205_N.39_205/39/04_1c` |

#### 219. `P.205_N.39_205/39/05_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1726 | `P.205_N.39_205/39/05_1` | SEKOLAH KEBANGSAAN ULU AWIK | a | `P.205_N.39_205/39/05_1a` |
| 1727 | `P.205_N.39_205/39/05_1` | RH. TANGAN ULU RISAU | b | `P.205_N.39_205/39/05_1b` |
| 1728 | `P.205_N.39_205/39/05_1` | SEKOLAH KEBANGSAAN NG ATOI | c | `P.205_N.39_205/39/05_1c` |
| 1729 | `P.205_N.39_205/39/05_1` | RH. ASON, TEMBAWAI KAPOK | d | `P.205_N.39_205/39/05_1d` |
| 1730 | `P.205_N.39_205/39/05_1` | SEKOLAH KEBANGSAAN LUBOK KEPAYANG | e | `P.205_N.39_205/39/05_1e` |
| 1731 | `P.205_N.39_205/39/05_1` | RH. AYOM NG. LUAU | f | `P.205_N.39_205/39/05_1f` |

#### 220. `P.205_N.39_205/39/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1732 | `P.205_N.39_205/39/06_1` | DEWAN SERBAGUNA ST. JEROME KAKI WONG SARATOK | a | `P.205_N.39_205/39/06_1a` |
| 1733 | `P.205_N.39_205/39/06_1` | RH. MAWAN KAWIT BUNSI SARATOK | b | `P.205_N.39_205/39/06_1b` |
| 1734 | `P.205_N.39_205/39/06_1` | SEKOLAH KEBANGSAAN NG ASSAM | c | `P.205_N.39_205/39/06_1c` |
| 1735 | `P.205_N.39_205/39/06_1` | SEKOLAH KEBANGSAAN MENDAS | d | `P.205_N.39_205/39/06_1d` |

#### 221. `P.205_N.39_205/39/07_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1736 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN KABO HILIR | a | `P.205_N.39_205/39/07_1a` |
| 1737 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN WONG BESI | b | `P.205_N.39_205/39/07_1b` |
| 1738 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN PRAJA | c | `P.205_N.39_205/39/07_1c` |
| 1739 | `P.205_N.39_205/39/07_1` | RH. ILLAI KRANGAN RUSA ILI | d | `P.205_N.39_205/39/07_1d` |
| 1740 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN BABANG | e | `P.205_N.39_205/39/07_1e` |
| 1741 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN LEMPA | f | `P.205_N.39_205/39/07_1f` |
| 1742 | `P.205_N.39_205/39/07_1` | SEKOLAH KEBANGSAAN ULU KABO | g | `P.205_N.39_205/39/07_1g` |

#### 222. `P.205_N.39_205/39/08_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1743 | `P.205_N.39_205/39/08_1` | SEKOLAH KEBANGSAAN ULU BUDU | a | `P.205_N.39_205/39/08_1a` |
| 1744 | `P.205_N.39_205/39/08_1` | RH. UDENG ULU BUDU SARATOK | b | `P.205_N.39_205/39/08_1b` |
| 1745 | `P.205_N.39_205/39/08_1` | SEKOLAH KEBANGSAAN NG BUDU | c | `P.205_N.39_205/39/08_1c` |
| 1746 | `P.205_N.39_205/39/08_1` | RH. LEMAIE ENGKALA, SARATOK | d | `P.205_N.39_205/39/08_1d` |
| 1747 | `P.205_N.39_205/39/08_1` | RH. JABU CHUNDI, NG. SENULAU BUDU | e | `P.205_N.39_205/39/08_1e` |
| 1748 | `P.205_N.39_205/39/08_1` | RH. SUBING ULU KRIAN | f | `P.205_N.39_205/39/08_1f` |
| 1749 | `P.205_N.39_205/39/08_1` | RH. NGITAR AWAS KRIAN | g | `P.205_N.39_205/39/08_1g` |
| 1750 | `P.205_N.39_205/39/08_1` | SEKOLAH KEBANGSAAN NG GRENJANG | h | `P.205_N.39_205/39/08_1h` |
| 1751 | `P.205_N.39_205/39/08_1` | RH. ENDIT NG. BATANG | i | `P.205_N.39_205/39/08_1i` |

#### 223. `P.205_N.39_205/39/09_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1752 | `P.205_N.39_205/39/09_1` | SEKOLAH KEBANGSAAN NG ABU | a | `P.205_N.39_205/39/09_1a` |
| 1753 | `P.205_N.39_205/39/09_1` | RH. EMPAWIE BAJAU SARATOK | b | `P.205_N.39_205/39/09_1b` |

#### 224. `P.205_N.39_205/39/10_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1754 | `P.205_N.39_205/39/10_1` | SEKOLAH KEBANGSAAN KLUA | a | `P.205_N.39_205/39/10_1a` |
| 1755 | `P.205_N.39_205/39/10_1` | SEKOLAH KEBANGSAAN ORANG KAYA TEMENGGONG TANDUK | b | `P.205_N.39_205/39/10_1b` |

#### 225. `P.205_N.40_205/40/01_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1758 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN HJ JUNID | a | `P.205_N.40_205/40/01_1a` |
| 1760 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN KPG ALIT | b | `P.205_N.40_205/40/01_1b` |
| 1761 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN ENGKABANG | c | `P.205_N.40_205/40/01_1c` |
| 1762 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN ST MICHAEL PLASSU | d | `P.205_N.40_205/40/01_1d` |
| 1763 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN SG PASIR | e | `P.205_N.40_205/40/01_1e` |
| 1764 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN TO' EMAN | f | `P.205_N.40_205/40/01_1f` |
| 1766 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN KG EMPLAM | g | `P.205_N.40_205/40/01_1g` |
| 1767 | `P.205_N.40_205/40/01_1` | SEKOLAH KEBANGSAAN ABG. MOH SESSANG | h | `P.205_N.40_205/40/01_1h` |

#### 226. `P.205_N.40_205/40/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1759 | `P.205_N.40_205/40/01_2` | SEKOLAH KEBANGSAAN HJ JUNID | a | `P.205_N.40_205/40/01_2a` |
| 1765 | `P.205_N.40_205/40/01_2` | SEKOLAH KEBANGSAAN TO' EMAN | b | `P.205_N.40_205/40/01_2b` |
| 1768 | `P.205_N.40_205/40/01_2` | SEKOLAH KEBANGSAAN ABG. MOH SESSANG | c | `P.205_N.40_205/40/01_2c` |

#### 227. `P.205_N.40_205/40/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1769 | `P.205_N.40_205/40/02_1` | SEKOLAH KEBANGSAAN ULU ROBAN | a | `P.205_N.40_205/40/02_1a` |
| 1770 | `P.205_N.40_205/40/02_1` | SEKOLAH KEBANGSAAN ST PAUL ROBAN | b | `P.205_N.40_205/40/02_1b` |
| 1772 | `P.205_N.40_205/40/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA ROBAN | c | `P.205_N.40_205/40/02_1c` |

#### 228. `P.205_N.40_205/40/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1771 | `P.205_N.40_205/40/02_2` | SEKOLAH KEBANGSAAN ST PAUL ROBAN | a | `P.205_N.40_205/40/02_2a` |
| 1773 | `P.205_N.40_205/40/02_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA ROBAN | b | `P.205_N.40_205/40/02_2b` |

#### 229. `P.205_N.40_205/40/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1774 | `P.205_N.40_205/40/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KABONG | a | `P.205_N.40_205/40/03_1a` |
| 1776 | `P.205_N.40_205/40/03_1` | SEKOLAH KEBANGSAAN ABANG LEMAN | b | `P.205_N.40_205/40/03_1b` |

#### 230. `P.205_N.40_205/40/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1775 | `P.205_N.40_205/40/03_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KABONG | a | `P.205_N.40_205/40/03_2a` |
| 1777 | `P.205_N.40_205/40/03_2` | SEKOLAH KEBANGSAAN ABANG LEMAN | b | `P.205_N.40_205/40/03_2b` |

#### 231. `P.206_N.41_206/41/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1788 | `P.206_N.41_206/41/02_1` | SEKOLAH KEBANGSAAN STALON | a | `P.206_N.41_206/41/02_1a` |
| 1789 | `P.206_N.41_206/41/02_1` | SEKOLAH KEBANGSAAN ABANG GESA | b | `P.206_N.41_206/41/02_1b` |

#### 232. `P.206_N.41_206/41/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1794 | `P.206_N.41_206/41/03_1` | SEKOLAH KEBANGSAAN BAYANG | a | `P.206_N.41_206/41/03_1a` |
| 1795 | `P.206_N.41_206/41/03_1` | SEKOLAH KEBANGSAAN ABANG GALAU | b | `P.206_N.41_206/41/03_1b` |
| 1798 | `P.206_N.41_206/41/03_1` | SEKOLAH KEBANGSAAN ABANG BUYUK | c | `P.206_N.41_206/41/03_1c` |

#### 233. `P.206_N.41_206/41/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1796 | `P.206_N.41_206/41/03_2` | SEKOLAH KEBANGSAAN ABANG GALAU | a | `P.206_N.41_206/41/03_2a` |
| 1799 | `P.206_N.41_206/41/03_2` | SEKOLAH KEBANGSAAN ABANG BUYUK | b | `P.206_N.41_206/41/03_2b` |

#### 234. `P.206_N.41_206/41/04_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1800 | `P.206_N.41_206/41/04_1` | SEKOLAH KEBANGSAAN JAWA | a | `P.206_N.41_206/41/04_1a` |
| 1801 | `P.206_N.41_206/41/04_1` | SEKOLAH KEBANGSAAN ADIN | b | `P.206_N.41_206/41/04_1b` |
| 1804 | `P.206_N.41_206/41/04_1` | SEKOLAH KEBANGSAAN BUKIT KINYAU | c | `P.206_N.41_206/41/04_1c` |
| 1805 | `P.206_N.41_206/41/04_1` | SEKOLAH KEBANGSAAN SG SENTEBU | d | `P.206_N.41_206/41/04_1d` |
| 1806 | `P.206_N.41_206/41/04_1` | SEKOLAH KEBANGSAAN ST. ANDREW | e | `P.206_N.41_206/41/04_1e` |

#### 235. `P.206_N.41_206/41/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1807 | `P.206_N.41_206/41/05_1` | SEKOLAH KEBANGSAAN MUARA PAYANG | a | `P.206_N.41_206/41/05_1a` |
| 1808 | `P.206_N.41_206/41/05_1` | SEKOLAH KEBANGSAAN ABANG HAJI MATAHIR | b | `P.206_N.41_206/41/05_1b` |
| 1812 | `P.206_N.41_206/41/05_1` | SEKOLAH KEBANGSAAN AGAMA SARIKEI | c | `P.206_N.41_206/41/05_1c` |

#### 236. `P.206_N.41_206/41/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1809 | `P.206_N.41_206/41/05_2` | SEKOLAH KEBANGSAAN ABANG HAJI MATAHIR | a | `P.206_N.41_206/41/05_2a` |
| 1813 | `P.206_N.41_206/41/05_2` | SEKOLAH KEBANGSAAN AGAMA SARIKEI | b | `P.206_N.41_206/41/05_2b` |

#### 237. `P.206_N.41_206/41/05_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1810 | `P.206_N.41_206/41/05_3` | SEKOLAH KEBANGSAAN ABANG HAJI MATAHIR | a | `P.206_N.41_206/41/05_3a` |
| 1814 | `P.206_N.41_206/41/05_3` | SEKOLAH KEBANGSAAN AGAMA SARIKEI | b | `P.206_N.41_206/41/05_3b` |

#### 238. `P.206_N.41_206/41/05_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1811 | `P.206_N.41_206/41/05_4` | SEKOLAH KEBANGSAAN ABANG HAJI MATAHIR | a | `P.206_N.41_206/41/05_4a` |
| 1815 | `P.206_N.41_206/41/05_4` | SEKOLAH KEBANGSAAN AGAMA SARIKEI | b | `P.206_N.41_206/41/05_4b` |

#### 239. `P.206_N.42_206/42/01_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1819 | `P.206_N.42_206/42/01_1` | SEKOLAH KEBANGSAAN SEBAKO | a | `P.206_N.42_206/42/01_1a` |
| 1820 | `P.206_N.42_206/42/01_1` | PUSAT KOMUNITI KPG. SEDI | b | `P.206_N.42_206/42/01_1b` |
| 1821 | `P.206_N.42_206/42/01_1` | SEKOLAH KEBANGSAAN ORANG KAYA MUDA | c | `P.206_N.42_206/42/01_1c` |
| 1823 | `P.206_N.42_206/42/01_1` | DEWAN SERBAGUNA, BERANGAN | d | `P.206_N.42_206/42/01_1d` |
| 1824 | `P.206_N.42_206/42/01_1` | TADIKA KEMAS KPG. KEDANG | e | `P.206_N.42_206/42/01_1e` |

#### 240. `P.206_N.42_206/42/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1825 | `P.206_N.42_206/42/02_1` | SEKOLAH KEBANGSAAN MUPONG | a | `P.206_N.42_206/42/02_1a` |
| 1826 | `P.206_N.42_206/42/02_1` | SEKOLAH KEBANGSAAN MUPONG ULIN | b | `P.206_N.42_206/42/02_1b` |

#### 241. `P.206_N.42_206/42/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1827 | `P.206_N.42_206/42/03_1` | SEKOLAH KEBANGSAAN KG SERDENG | a | `P.206_N.42_206/42/03_1a` |
| 1828 | `P.206_N.42_206/42/03_1` | SEKOLAH KEBANGSAAN TELOK GELAM | b | `P.206_N.42_206/42/03_1b` |
| 1829 | `P.206_N.42_206/42/03_1` | SEKOLAH KEBANGSAAN BAKERKONG | c | `P.206_N.42_206/42/03_1c` |
| 1830 | `P.206_N.42_206/42/03_1` | SEKOLAH KEBANGSAAN SEMOP | d | `P.206_N.42_206/42/03_1d` |

#### 242. `P.206_N.42_206/42/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1832 | `P.206_N.42_206/42/04_1` | SEKOLAH KEBANGSAAN KG KUT | a | `P.206_N.42_206/42/04_1a` |
| 1833 | `P.206_N.42_206/42/04_1` | SEKOLAH KEBANGSAAN MOHAMAD REDEH SAAI | b | `P.206_N.42_206/42/04_1b` |

#### 243. `P.206_N.42_206/42/05_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1834 | `P.206_N.42_206/42/05_1` | SEKOLAH KEBANGSAAN KG PENIBONG | a | `P.206_N.42_206/42/05_1a` |
| 1835 | `P.206_N.42_206/42/05_1` | SEKOLAH KEBANGSAAN KG BETANAK | b | `P.206_N.42_206/42/05_1b` |
| 1837 | `P.206_N.42_206/42/05_1` | SEKOLAH KEBANGSAAN SALAH KECHIL | c | `P.206_N.42_206/42/05_1c` |
| 1838 | `P.206_N.42_206/42/05_1` | SEKOLAH KEBANGSAAN KG PENIPAH | d | `P.206_N.42_206/42/05_1d` |
| 1840 | `P.206_N.42_206/42/05_1` | SEKOLAH KEBANGSAAN KG TEKAJONG | e | `P.206_N.42_206/42/05_1e` |

#### 244. `P.206_N.42_206/42/05_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1836 | `P.206_N.42_206/42/05_2` | SEKOLAH KEBANGSAAN KG BETANAK | a | `P.206_N.42_206/42/05_2a` |
| 1839 | `P.206_N.42_206/42/05_2` | SEKOLAH KEBANGSAAN KG PENIPAH | b | `P.206_N.42_206/42/05_2b` |
| 1841 | `P.206_N.42_206/42/05_2` | SEKOLAH KEBANGSAAN KG TEKAJONG | c | `P.206_N.42_206/42/05_2c` |

#### 245. `P.206_N.42_206/42/07_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1845 | `P.206_N.42_206/42/07_1` | SEKOLAH KEBANGSAAN SG SELIDAP | a | `P.206_N.42_206/42/07_1a` |
| 1846 | `P.206_N.42_206/42/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) MING SHING | b | `P.206_N.42_206/42/07_1b` |
| 1847 | `P.206_N.42_206/42/07_1` | SEKOLAH KEBANGSAAN SG SIAN | c | `P.206_N.42_206/42/07_1c` |
| 1848 | `P.206_N.42_206/42/07_1` | SEKOLAH KEBANGSAAN TANJUNG BUNDUNG | d | `P.206_N.42_206/42/07_1d` |

#### 246. `P.207_N.43_207/43/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1856 | `P.207_N.43_207/43/01_1` | RH. MAGERAT AK. RENTAP, SG. STUBAH | a | `P.207_N.43_207/43/01_1a` |
| 1857 | `P.207_N.43_207/43/01_1` | SEKOLAH KEBANGSAAN NANGA SEMAH | b | `P.207_N.43_207/43/01_1b` |
| 1858 | `P.207_N.43_207/43/01_1` | SEKOLAH KEBANGSAAN SABENA | c | `P.207_N.43_207/43/01_1c` |
| 1859 | `P.207_N.43_207/43/01_1` | SEKOLAH KEBANGSAAN KPG PENASU DARO | d | `P.207_N.43_207/43/01_1d` |

#### 247. `P.207_N.43_207/43/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1860 | `P.207_N.43_207/43/02_1` | SEKOLAH KEBANGSAAN CAMPORAN | a | `P.207_N.43_207/43/02_1a` |
| 1863 | `P.207_N.43_207/43/02_1` | SEKOLAH KEBANGSAAN KG TEBAANG | b | `P.207_N.43_207/43/02_1b` |

#### 248. `P.207_N.43_207/43/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1861 | `P.207_N.43_207/43/02_2` | SEKOLAH KEBANGSAAN CAMPORAN | a | `P.207_N.43_207/43/02_2a` |
| 1864 | `P.207_N.43_207/43/02_2` | SEKOLAH KEBANGSAAN KG TEBAANG | b | `P.207_N.43_207/43/02_2b` |

#### 249. `P.207_N.43_207/43/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1865 | `P.207_N.43_207/43/03_1` | SEKOLAH KEBANGSAAN NANGAR | a | `P.207_N.43_207/43/03_1a` |
| 1867 | `P.207_N.43_207/43/03_1` | SEKOLAH KEBANGSAAN HIJRAH BADONG | b | `P.207_N.43_207/43/03_1b` |
| 1869 | `P.207_N.43_207/43/03_1` | TADIKA KIE MING DARO | c | `P.207_N.43_207/43/03_1c` |

#### 250. `P.207_N.43_207/43/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1866 | `P.207_N.43_207/43/03_2` | SEKOLAH KEBANGSAAN NANGAR | a | `P.207_N.43_207/43/03_2a` |
| 1868 | `P.207_N.43_207/43/03_2` | SEKOLAH KEBANGSAAN HIJRAH BADONG | b | `P.207_N.43_207/43/03_2b` |
| 1870 | `P.207_N.43_207/43/03_2` | TADIKA KIE MING DARO | c | `P.207_N.43_207/43/03_2c` |

#### 251. `P.207_N.43_207/43/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1872 | `P.207_N.43_207/43/04_1` | SEKOLAH KEBANGSAAN ULU DARO | a | `P.207_N.43_207/43/04_1a` |
| 1874 | `P.207_N.43_207/43/04_1` | SEKOLAH KEBANGSAAN KG PANGTRAY | b | `P.207_N.43_207/43/04_1b` |

#### 252. `P.207_N.43_207/43/05_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1875 | `P.207_N.43_207/43/05_1` | SEKOLAH KEBANGSAAN SG  LENGAN | a | `P.207_N.43_207/43/05_1a` |
| 1876 | `P.207_N.43_207/43/05_1` | SEKOLAH KEBANGSAAN SINGAT | b | `P.207_N.43_207/43/05_1b` |
| 1877 | `P.207_N.43_207/43/05_1` | SEKOLAH KEBANGSAAN KG SAWAI | c | `P.207_N.43_207/43/05_1c` |
| 1878 | `P.207_N.43_207/43/05_1` | SEKOLAH KEBANGSAAN SUNGEI PASSIN | d | `P.207_N.43_207/43/05_1d` |
| 1879 | `P.207_N.43_207/43/05_1` | SEKOLAH KEBANGSAAN BATANG LASSA | e | `P.207_N.43_207/43/05_1e` |

#### 253. `P.207_N.44_207/44/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1882 | `P.207_N.44_207/44/01_1` | SEKOLAH KEBANGSAAN BRUAN MAPAL | a | `P.207_N.44_207/44/01_1a` |
| 1883 | `P.207_N.44_207/44/01_1` | SEKOLAH KEBANGSAAN KUALA MATU | b | `P.207_N.44_207/44/01_1b` |
| 1886 | `P.207_N.44_207/44/01_1` | SEKOLAH KEBANGSAAN KG PERGAU | c | `P.207_N.44_207/44/01_1c` |

#### 254. `P.207_N.44_207/44/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1889 | `P.207_N.44_207/44/04_1` | BANGUNAN PUSAT SUMBER JKKK KAMPUNG TRENG | a | `P.207_N.44_207/44/04_1a` |
| 1890 | `P.207_N.44_207/44/04_1` | SEKOLAH KEBANGSAAN O K SELAIR | b | `P.207_N.44_207/44/04_1b` |

#### 255. `P.207_N.44_207/44/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1894 | `P.207_N.44_207/44/06_1` | SEKOLAH KEBANGSAAN SEKAAN KECHIL | a | `P.207_N.44_207/44/06_1a` |
| 1895 | `P.207_N.44_207/44/06_1` | DEWAN KPG. BAWANG | b | `P.207_N.44_207/44/06_1b` |

#### 256. `P.208_N.45_208/45/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1911 | `P.208_N.45_208/45/01_1` | SEKOLAH KEBANGSAAN SELANGAN | a | `P.208_N.45_208/45/01_1a` |
| 1914 | `P.208_N.45_208/45/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SIUNG HUA | b | `P.208_N.45_208/45/01_1b` |
| 1915 | `P.208_N.45_208/45/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) BULAT | c | `P.208_N.45_208/45/01_1c` |

#### 257. `P.208_N.45_208/45/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1912 | `P.208_N.45_208/45/01_2` | SEKOLAH KEBANGSAAN SELANGAN | a | `P.208_N.45_208/45/01_2a` |
| 1916 | `P.208_N.45_208/45/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) BULAT | b | `P.208_N.45_208/45/01_2b` |

#### 258. `P.208_N.45_208/45/01_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1913 | `P.208_N.45_208/45/01_3` | SEKOLAH KEBANGSAAN SELANGAN | a | `P.208_N.45_208/45/01_3a` |
| 1917 | `P.208_N.45_208/45/01_3` | SEKOLAH JENIS KEBANGSAAN (CINA) BULAT | b | `P.208_N.45_208/45/01_3b` |

#### 259. `P.208_N.45_208/45/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1940 | `P.208_N.45_208/45/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SU LOK | a | `P.208_N.45_208/45/05_1a` |
| 1943 | `P.208_N.45_208/45/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SU LEE | b | `P.208_N.45_208/45/05_1b` |

#### 260. `P.208_N.45_208/45/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1944 | `P.208_N.45_208/45/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SZE LU | a | `P.208_N.45_208/45/06_1a` |
| 1945 | `P.208_N.45_208/45/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TIONG HO | b | `P.208_N.45_208/45/06_1b` |

#### 261. `P.208_N.45_208/45/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1947 | `P.208_N.45_208/45/07_1` | SEKOLAH MENENGAH KEBANGSAAN SG PAOH | a | `P.208_N.45_208/45/07_1a` |
| 1950 | `P.208_N.45_208/45/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) ST MARTIN | b | `P.208_N.45_208/45/07_1b` |

#### 262. `P.208_N.45_208/45/07_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1948 | `P.208_N.45_208/45/07_2` | SEKOLAH MENENGAH KEBANGSAAN SG PAOH | a | `P.208_N.45_208/45/07_2a` |
| 1951 | `P.208_N.45_208/45/07_2` | SEKOLAH JENIS KEBANGSAAN (CINA) ST MARTIN | b | `P.208_N.45_208/45/07_2b` |

#### 263. `P.208_N.45_208/45/07_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1949 | `P.208_N.45_208/45/07_3` | SEKOLAH MENENGAH KEBANGSAAN SG PAOH | a | `P.208_N.45_208/45/07_3a` |
| 1952 | `P.208_N.45_208/45/07_3` | SEKOLAH JENIS KEBANGSAAN (CINA) ST MARTIN | b | `P.208_N.45_208/45/07_3b` |

#### 264. `P.208_N.46_208/46/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1967 | `P.208_N.46_208/46/02_1` | SEKOLAH KEBANGSAAN TULAI | a | `P.208_N.46_208/46/02_1a` |
| 1968 | `P.208_N.46_208/46/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TONG HUA | b | `P.208_N.46_208/46/02_1b` |

#### 265. `P.208_N.46_208/46/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1969 | `P.208_N.46_208/46/03_1` | SEKOLAH KEBANGSAAN SG KAWI | a | `P.208_N.46_208/46/03_1a` |
| 1970 | `P.208_N.46_208/46/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KUNG CHENG | b | `P.208_N.46_208/46/03_1b` |
| 1971 | `P.208_N.46_208/46/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) YONG KWONG | c | `P.208_N.46_208/46/03_1c` |

#### 266. `P.208_N.46_208/46/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1973 | `P.208_N.46_208/46/04_1` | SEKOLAH KEBANGSAAN GEMUAN | a | `P.208_N.46_208/46/04_1a` |
| 1974 | `P.208_N.46_208/46/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) NAN CHIEW | b | `P.208_N.46_208/46/04_1b` |
| 1975 | `P.208_N.46_208/46/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) MING TEE | c | `P.208_N.46_208/46/04_1c` |

#### 267. `P.208_N.46_208/46/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1978 | `P.208_N.46_208/46/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KAI CHUNG | a | `P.208_N.46_208/46/05_1a` |
| 1981 | `P.208_N.46_208/46/05_1` | SEKOLAH MENENGAH KEBANGSAAN KAI CHUNG | b | `P.208_N.46_208/46/05_1b` |
| 1985 | `P.208_N.46_208/46/05_1` | SEKOLAH KEBANGSAAN BANDAR BINTANGOR | c | `P.208_N.46_208/46/05_1c` |

#### 268. `P.208_N.46_208/46/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1979 | `P.208_N.46_208/46/05_2` | SEKOLAH JENIS KEBANGSAAN (CINA) KAI CHUNG | a | `P.208_N.46_208/46/05_2a` |
| 1982 | `P.208_N.46_208/46/05_2` | SEKOLAH MENENGAH KEBANGSAAN KAI CHUNG | b | `P.208_N.46_208/46/05_2b` |

#### 269. `P.208_N.46_208/46/05_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1980 | `P.208_N.46_208/46/05_3` | SEKOLAH JENIS KEBANGSAAN (CINA) KAI CHUNG | a | `P.208_N.46_208/46/05_3a` |
| 1983 | `P.208_N.46_208/46/05_3` | SEKOLAH MENENGAH KEBANGSAAN KAI CHUNG | b | `P.208_N.46_208/46/05_3b` |

#### 270. `P.208_N.46_208/46/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1987 | `P.208_N.46_208/46/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) MING LU | a | `P.208_N.46_208/46/07_1a` |
| 1988 | `P.208_N.46_208/46/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TUNG KWONG | b | `P.208_N.46_208/46/07_1b` |

#### 271. `P.208_N.46_208/46/08_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1989 | `P.208_N.46_208/46/08_1` | SEKOLAH KEBANGSAAN SG MADOR | a | `P.208_N.46_208/46/08_1a` |
| 1991 | `P.208_N.46_208/46/08_1` | TADIKA KEMAS SG. RAYAH | b | `P.208_N.46_208/46/08_1b` |
| 1992 | `P.208_N.46_208/46/08_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KAI SING | c | `P.208_N.46_208/46/08_1c` |

#### 272. `P.208_N.46_208/46/09_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1993 | `P.208_N.46_208/46/09_1` | SEKOLAH KEBANGSAAN ULU BINATANG | a | `P.208_N.46_208/46/09_1a` |
| 1995 | `P.208_N.46_208/46/09_1` | SEKOLAH KEBANGSAAN ULU STRAS | b | `P.208_N.46_208/46/09_1b` |
| 1996 | `P.208_N.46_208/46/09_1` | SEKOLAH KEBANGSAAN NANGA STRAS | c | `P.208_N.46_208/46/09_1c` |
| 1999 | `P.208_N.46_208/46/09_1` | SEKOLAH JENIS KEBANGSAAN (CINA) MIN DAIK | d | `P.208_N.46_208/46/09_1d` |

#### 273. `P.208_N.46_208/46/09_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 1994 | `P.208_N.46_208/46/09_2` | SEKOLAH KEBANGSAAN ULU BINATANG | a | `P.208_N.46_208/46/09_2a` |
| 1997 | `P.208_N.46_208/46/09_2` | SEKOLAH KEBANGSAAN NANGA STRAS | b | `P.208_N.46_208/46/09_2b` |
| 2000 | `P.208_N.46_208/46/09_2` | SEKOLAH JENIS KEBANGSAAN (CINA) MIN DAIK | c | `P.208_N.46_208/46/09_2c` |

#### 274. `P.208_N.46_208/46/10_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2001 | `P.208_N.46_208/46/10_1` | BALAI RAYA KPG. SELEMAS | a | `P.208_N.46_208/46/10_1a` |
| 2002 | `P.208_N.46_208/46/10_1` | TADIKA METHODIST BINTANGOR TOWN | b | `P.208_N.46_208/46/10_1b` |

#### 275. `P.208_N.46_208/46/11_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2003 | `P.208_N.46_208/46/11_1` | RH REGINA SG. SELIDAP | a | `P.208_N.46_208/46/11_1a` |
| 2004 | `P.208_N.46_208/46/11_1` | SEKOLAH KEBANGSAAN TANAH PUTIH | b | `P.208_N.46_208/46/11_1b` |
| 2005 | `P.208_N.46_208/46/11_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SU TAK | c | `P.208_N.46_208/46/11_1c` |
| 2006 | `P.208_N.46_208/46/11_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SU MING | d | `P.208_N.46_208/46/11_1d` |

#### 276. `P.208_N.46_208/46/12_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2007 | `P.208_N.46_208/46/12_1` | SEKOLAH KEBANGSAAN SG PETAI | a | `P.208_N.46_208/46/12_1a` |
| 2009 | `P.208_N.46_208/46/12_1` | SEKOLAH KEBANGSAAN RENTAP | b | `P.208_N.46_208/46/12_1b` |
| 2010 | `P.208_N.46_208/46/12_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HUA KEE | c | `P.208_N.46_208/46/12_1c` |
| 2011 | `P.208_N.46_208/46/12_1` | SEKOLAH JENIS KEBANGSAAN (CINA) NAM KIEW | d | `P.208_N.46_208/46/12_1d` |
| 2012 | `P.208_N.46_208/46/12_1` | SEKOLAH KEBANGSAAN BUKIT NIBONG | e | `P.208_N.46_208/46/12_1e` |

#### 277. `P.208_N.46_208/46/12_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2008 | `P.208_N.46_208/46/12_2` | SEKOLAH KEBANGSAAN SG PETAI | a | `P.208_N.46_208/46/12_2a` |
| 2013 | `P.208_N.46_208/46/12_2` | SEKOLAH KEBANGSAAN BUKIT NIBONG | b | `P.208_N.46_208/46/12_2b` |

#### 278. `P.209_N.47_209/47/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2015 | `P.209_N.47_209/47/01_1` | SEKOLAH KEBANGSAAN BALONG | a | `P.209_N.47_209/47/01_1a` |
| 2016 | `P.209_N.47_209/47/01_1` | SEKOLAH KEBANGSAAN UDIN | b | `P.209_N.47_209/47/01_1b` |
| 2017 | `P.209_N.47_209/47/01_1` | RH. ANDREW MANGKA SG. RIBONG | c | `P.209_N.47_209/47/01_1c` |
| 2018 | `P.209_N.47_209/47/01_1` | RH. TIRAI SEBANGKOI | d | `P.209_N.47_209/47/01_1d` |

#### 279. `P.209_N.47_209/47/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2019 | `P.209_N.47_209/47/02_1` | SEKOLAH KEBANGSAAN PENGHULU ANDIN | a | `P.209_N.47_209/47/02_1a` |
| 2021 | `P.209_N.47_209/47/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SU DOK | b | `P.209_N.47_209/47/02_1b` |

#### 280. `P.209_N.47_209/47/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2022 | `P.209_N.47_209/47/03_1` | RH. BUDUM  NG. TUBAI BUAH | a | `P.209_N.47_209/47/03_1a` |
| 2023 | `P.209_N.47_209/47/03_1` | SEKOLAH KEBANGSAAN ULU PEDANUM | b | `P.209_N.47_209/47/03_1b` |
| 2024 | `P.209_N.47_209/47/03_1` | SEKOLAH KEBANGSAAN NANGA BUKU | c | `P.209_N.47_209/47/03_1c` |
| 2025 | `P.209_N.47_209/47/03_1` | SEKOLAH KEBANGSAAN NANGA BABAI | d | `P.209_N.47_209/47/03_1d` |

#### 281. `P.209_N.47_209/47/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2028 | `P.209_N.47_209/47/04_1` | SEKOLAH KEBANGSAAN NANGA PAKAN | a | `P.209_N.47_209/47/04_1a` |
| 2034 | `P.209_N.47_209/47/04_1` | SEKOLAH KEBANGSAAN  ULU MANDING | b | `P.209_N.47_209/47/04_1b` |
| 2035 | `P.209_N.47_209/47/04_1` | SEKOLAH KEBANGSAAN SUNGAI SUGAI | c | `P.209_N.47_209/47/04_1c` |

#### 282. `P.209_N.47_209/47/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2029 | `P.209_N.47_209/47/04_2` | SEKOLAH KEBANGSAAN NANGA PAKAN | a | `P.209_N.47_209/47/04_2a` |
| 2036 | `P.209_N.47_209/47/04_2` | SEKOLAH KEBANGSAAN SUNGAI SUGAI | b | `P.209_N.47_209/47/04_2b` |

#### 283. `P.209_N.47_209/47/04_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2030 | `P.209_N.47_209/47/04_3` | SEKOLAH KEBANGSAAN NANGA PAKAN | a | `P.209_N.47_209/47/04_3a` |
| 2037 | `P.209_N.47_209/47/04_3` | SEKOLAH KEBANGSAAN SUNGAI SUGAI | b | `P.209_N.47_209/47/04_3b` |

#### 284. `P.209_N.47_209/47/05_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2038 | `P.209_N.47_209/47/05_1` | RH. JULIUS DUAT MAUR @ JULIUS EDWARD | a | `P.209_N.47_209/47/05_1a` |
| 2039 | `P.209_N.47_209/47/05_1` | SEKOLAH KEBANGSAAN NANGA KEDUP | b | `P.209_N.47_209/47/05_1b` |
| 2040 | `P.209_N.47_209/47/05_1` | RH. MANDAU SG. MARAM PAKAN | c | `P.209_N.47_209/47/05_1c` |
| 2041 | `P.209_N.47_209/47/05_1` | SEKOLAH KEBANGSAAN NANGA WAK | d | `P.209_N.47_209/47/05_1d` |

#### 285. `P.209_N.47_209/47/06_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2043 | `P.209_N.47_209/47/06_1` | SEKOLAH KEBANGSAAN NANGA KOTA | a | `P.209_N.47_209/47/06_1a` |
| 2044 | `P.209_N.47_209/47/06_1` | SEKOLAH KEBANGSAAN NANGA DAYU | b | `P.209_N.47_209/47/06_1b` |
| 2045 | `P.209_N.47_209/47/06_1` | SEKOLAH KEBANGSAAN NANGA KARA | c | `P.209_N.47_209/47/06_1c` |
| 2046 | `P.209_N.47_209/47/06_1` | SEKOLAH KEBANGSAAN NANGA SEMAWANG | d | `P.209_N.47_209/47/06_1d` |
| 2047 | `P.209_N.47_209/47/06_1` | RH. KO ANAK CHANGGAN SG. BALONG | e | `P.209_N.47_209/47/06_1e` |

#### 286. `P.209_N.48_209/48/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2050 | `P.209_N.48_209/48/01_1` | SEKOLAH KEBANGSAAN NANGA LASI | a | `P.209_N.48_209/48/01_1a` |
| 2054 | `P.209_N.48_209/48/01_1` | SEKOLAH KEBANGSAAN NG SERAU | b | `P.209_N.48_209/48/01_1b` |

#### 287. `P.209_N.48_209/48/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2055 | `P.209_N.48_209/48/02_1` | RH. LIBA NG. BUA | a | `P.209_N.48_209/48/02_1a` |
| 2056 | `P.209_N.48_209/48/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) YUK KUNG | b | `P.209_N.48_209/48/02_1b` |

#### 288. `P.209_N.48_209/48/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2061 | `P.209_N.48_209/48/03_1` | BALAI RAYA RH. LIDOH | a | `P.209_N.48_209/48/03_1a` |
| 2063 | `P.209_N.48_209/48/03_1` | SEKOLAH KEBANGSAAN ST ALPHONSUS | b | `P.209_N.48_209/48/03_1b` |

#### 289. `P.209_N.48_209/48/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2062 | `P.209_N.48_209/48/03_2` | BALAI RAYA RH. LIDOH | a | `P.209_N.48_209/48/03_2a` |
| 2064 | `P.209_N.48_209/48/03_2` | SEKOLAH KEBANGSAAN ST ALPHONSUS | b | `P.209_N.48_209/48/03_2b` |

#### 290. `P.209_N.48_209/48/04_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2065 | `P.209_N.48_209/48/04_1` | SEKOLAH KEBANGSAAN NANGA KELANGAS | a | `P.209_N.48_209/48/04_1a` |
| 2066 | `P.209_N.48_209/48/04_1` | SEKOLAH KEBANGSAAN NANGA MERURUN | b | `P.209_N.48_209/48/04_1b` |
| 2067 | `P.209_N.48_209/48/04_1` | TADIKA KEMAS NG. BILAT/LIJAN | c | `P.209_N.48_209/48/04_1c` |
| 2069 | `P.209_N.48_209/48/04_1` | SEKOLAH KEBANGSAAN NANGA MELUAN | d | `P.209_N.48_209/48/04_1d` |

#### 291. `P.209_N.48_209/48/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2070 | `P.209_N.48_209/48/05_1` | SEKOLAH KEBANGSAAN NANGA SENGAIH | a | `P.209_N.48_209/48/05_1a` |
| 2072 | `P.209_N.48_209/48/05_1` | SEKOLAH KEBANGSAAN NANGA ENGKAMOP | b | `P.209_N.48_209/48/05_1b` |

#### 292. `P.209_N.48_209/48/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2071 | `P.209_N.48_209/48/05_2` | SEKOLAH KEBANGSAAN NANGA SENGAIH | a | `P.209_N.48_209/48/05_2a` |
| 2073 | `P.209_N.48_209/48/05_2` | SEKOLAH KEBANGSAAN NANGA ENGKAMOP | b | `P.209_N.48_209/48/05_2b` |

#### 293. `P.209_N.48_209/48/06_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2075 | `P.209_N.48_209/48/06_1` | SEKOLAH KEBANGSAAN NANGA ENTAIH | a | `P.209_N.48_209/48/06_1a` |
| 2077 | `P.209_N.48_209/48/06_1` | SEKOLAH KEBANGSAAN ULU ENTABAI | b | `P.209_N.48_209/48/06_1b` |
| 2079 | `P.209_N.48_209/48/06_1` | SEKOLAH KEBANGSAAN ULU ENTAIH | c | `P.209_N.48_209/48/06_1c` |

#### 294. `P.209_N.48_209/48/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2076 | `P.209_N.48_209/48/06_2` | SEKOLAH KEBANGSAAN NANGA ENTAIH | a | `P.209_N.48_209/48/06_2a` |
| 2078 | `P.209_N.48_209/48/06_2` | SEKOLAH KEBANGSAAN ULU ENTABAI | b | `P.209_N.48_209/48/06_2b` |

#### 295. `P.209_N.48_209/48/07_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2080 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN NANGA JAMBU | a | `P.209_N.48_209/48/07_1a` |
| 2081 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN NG ENSIRING | b | `P.209_N.48_209/48/07_1b` |
| 2082 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN TAPANG PUNGGU | c | `P.209_N.48_209/48/07_1c` |
| 2083 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN NANGA JU | d | `P.209_N.48_209/48/07_1d` |
| 2084 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN NANGA MAONG | e | `P.209_N.48_209/48/07_1e` |
| 2085 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN LUBOK ASSAM | f | `P.209_N.48_209/48/07_1f` |
| 2086 | `P.209_N.48_209/48/07_1` | SEKOLAH KEBANGSAAN NANGA ENTABAI | g | `P.209_N.48_209/48/07_1g` |

#### 296. `P.210_N.49_210/49/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2091 | `P.210_N.49_210/49/02_1` | SEKOLAH KEBANGSAAN NANGA NIROK NGEMAH | a | `P.210_N.49_210/49/02_1a` |
| 2092 | `P.210_N.49_210/49/02_1` | PEJABAT PERTANIAN NG. NGEMAH | b | `P.210_N.49_210/49/02_1b` |
| 2093 | `P.210_N.49_210/49/02_1` | SEKOLAH KEBANGSAAN NANGA NGUNGUN | c | `P.210_N.49_210/49/02_1c` |

#### 297. `P.210_N.49_210/49/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2097 | `P.210_N.49_210/49/03_1` | SEKOLAH KEBANGSAAN NANGA NGEMAH | a | `P.210_N.49_210/49/03_1a` |
| 2098 | `P.210_N.49_210/49/03_1` | SEKOLAH KEBANGSAAN RANTAU DILANG | b | `P.210_N.49_210/49/03_1b` |
| 2100 | `P.210_N.49_210/49/03_1` | SEKOLAH KEBANGSAAN SENGAYAN | c | `P.210_N.49_210/49/03_1c` |
| 2101 | `P.210_N.49_210/49/03_1` | SEKOLAH KEBANGSAAN NANGA DAP | d | `P.210_N.49_210/49/03_1d` |

#### 298. `P.210_N.49_210/49/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2099 | `P.210_N.49_210/49/03_2` | SEKOLAH KEBANGSAAN RANTAU DILANG | a | `P.210_N.49_210/49/03_2a` |
| 2102 | `P.210_N.49_210/49/03_2` | SEKOLAH KEBANGSAAN NANGA DAP | b | `P.210_N.49_210/49/03_2b` |

#### 299. `P.210_N.49_210/49/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2103 | `P.210_N.49_210/49/04_1` | DEWAN DATUK AARON | a | `P.210_N.49_210/49/04_1a` |
| 2105 | `P.210_N.49_210/49/04_1` | SEKOLAH KEBANGSAAN  NANGA TADA | b | `P.210_N.49_210/49/04_1b` |

#### 300. `P.210_N.49_210/49/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2104 | `P.210_N.49_210/49/04_2` | DEWAN DATUK AARON | a | `P.210_N.49_210/49/04_2a` |
| 2106 | `P.210_N.49_210/49/04_2` | SEKOLAH KEBANGSAAN  NANGA TADA | b | `P.210_N.49_210/49/04_2b` |

#### 301. `P.210_N.49_210/49/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2109 | `P.210_N.49_210/49/05_1` | SEKOLAH KEBANGSAAN ULU BAWAN | a | `P.210_N.49_210/49/05_1a` |
| 2111 | `P.210_N.49_210/49/05_1` | SEKOLAH KEBANGSAAN SUNGAI TUAH | b | `P.210_N.49_210/49/05_1b` |
| 2112 | `P.210_N.49_210/49/05_1` | RH. STEPHEN CHENDANG | c | `P.210_N.49_210/49/05_1c` |

#### 302. `P.210_N.49_210/49/06_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2113 | `P.210_N.49_210/49/06_1` | SEKOLAH KEBANGSAAN NANGA PEDAI | a | `P.210_N.49_210/49/06_1a` |
| 2115 | `P.210_N.49_210/49/06_1` | SEKOLAH KEBANGSAAN NANGA JIH | b | `P.210_N.49_210/49/06_1b` |
| 2116 | `P.210_N.49_210/49/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SING SHING | c | `P.210_N.49_210/49/06_1c` |

#### 303. `P.210_N.50_210/50/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2119 | `P.210_N.50_210/50/01_1` | SEKOLAH JENIS KEBANGSAAN  (CINA) CHIH MONG | a | `P.210_N.50_210/50/01_1a` |
| 2120 | `P.210_N.50_210/50/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SHING HUA | b | `P.210_N.50_210/50/01_1b` |

#### 304. `P.210_N.50_210/50/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2121 | `P.210_N.50_210/50/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) YEE TING | a | `P.210_N.50_210/50/02_1a` |
| 2126 | `P.210_N.50_210/50/02_1` | TADIKA TAMAN MUHIBBAH | b | `P.210_N.50_210/50/02_1b` |

#### 305. `P.210_N.50_210/50/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2127 | `P.210_N.50_210/50/03_1` | SEKOLAH MENENGAH KEBANGSAAN KANOWIT | a | `P.210_N.50_210/50/03_1a` |
| 2129 | `P.210_N.50_210/50/03_1` | SEKOLAH KEBANGSAAN ULU RANAN | b | `P.210_N.50_210/50/03_1b` |
| 2130 | `P.210_N.50_210/50/03_1` | SEKOLAH KEBANGSAAN ULU MAJAU | c | `P.210_N.50_210/50/03_1c` |
| 2131 | `P.210_N.50_210/50/03_1` | SEKOLAH KEBANGSAAN BATU LUKING | d | `P.210_N.50_210/50/03_1d` |

#### 306. `P.210_N.50_210/50/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2132 | `P.210_N.50_210/50/04_1` | RH. ALI RANTAU ENSURAI POI | a | `P.210_N.50_210/50/04_1a` |
| 2133 | `P.210_N.50_210/50/04_1` | SEKOLAH KEBANGSAAN NANGA POI | b | `P.210_N.50_210/50/04_1b` |

#### 307. `P.210_N.50_210/50/05_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2134 | `P.210_N.50_210/50/05_1` | RH. BALAI ULU SG. POI | a | `P.210_N.50_210/50/05_1a` |
| 2135 | `P.210_N.50_210/50/05_1` | SEKOLAH KEBANGSAAN NANGA MENALUN | b | `P.210_N.50_210/50/05_1b` |
| 2136 | `P.210_N.50_210/50/05_1` | RH. GAMANG ULU MENUAN POI | c | `P.210_N.50_210/50/05_1c` |
| 2137 | `P.210_N.50_210/50/05_1` | SEKOLAH KEBANGSAAN ULU POI | d | `P.210_N.50_210/50/05_1d` |

#### 308. `P.210_N.50_210/50/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2138 | `P.210_N.50_210/50/06_1` | SEKOLAH KEBANGSAAN NG JAGOI | a | `P.210_N.50_210/50/06_1a` |
| 2139 | `P.210_N.50_210/50/06_1` | RH. STEPHEN JOK, NG. GEREMAI | b | `P.210_N.50_210/50/06_1b` |
| 2140 | `P.210_N.50_210/50/06_1` | SEKOLAH KEBANGSAAN NANGA LIPUS | c | `P.210_N.50_210/50/06_1c` |
| 2141 | `P.210_N.50_210/50/06_1` | SEKOLAH KEBANGSAAN RANTAU KEMIDING | d | `P.210_N.50_210/50/06_1d` |

#### 309. `P.210_N.50_210/50/07_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2142 | `P.210_N.50_210/50/07_1` | RH. LAJANG ULU MACHAN KANOWIT | a | `P.210_N.50_210/50/07_1a` |
| 2143 | `P.210_N.50_210/50/07_1` | SEKOLAH KEBANGSAAN ULU MACHAN | b | `P.210_N.50_210/50/07_1b` |
| 2144 | `P.210_N.50_210/50/07_1` | SEKOLAH KEBANGSAAN NANGA MACHAN | c | `P.210_N.50_210/50/07_1c` |
| 2148 | `P.210_N.50_210/50/07_1` | RH. MULING ANAK NYAPANG, MAONG KANOWIT | d | `P.210_N.50_210/50/07_1d` |
| 2149 | `P.210_N.50_210/50/07_1` | RH. JELANI NANGA LESIH | e | `P.210_N.50_210/50/07_1e` |
| 2150 | `P.210_N.50_210/50/07_1` | RH. SAWANG SIMPANG MACHAN | f | `P.210_N.50_210/50/07_1f` |

#### 310. `P.211_N.51_211/51/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2165 | `P.211_N.51_211/51/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HANG KWONG | a | `P.211_N.51_211/51/03_1a` |
| 2166 | `P.211_N.51_211/51/03_1` | SEKOLAH KEBANGSAAN NG SALIM | b | `P.211_N.51_211/51/03_1b` |
| 2167 | `P.211_N.51_211/51/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SUNG SANG | c | `P.211_N.51_211/51/03_1c` |

#### 311. `P.211_N.52_211/52/00_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2219 | `P.211_N.52_211/52/00_1` | DEWAN SERBAGUNA BATALION 10 JALAN UPPER LANANG, SIBU | a | `P.211_N.52_211/52/00_1a` |
| 2221 | `P.211_N.52_211/52/00_1` | DEWAN BADMINTON, BALAI POLIS SIBUJAYA | b | `P.211_N.52_211/52/00_1b` |
| 2222 | `P.211_N.52_211/52/00_1` | DEWAN SERBAGUNA KEM PLKN JUNACO PARK, SIBU | c | `P.211_N.52_211/52/00_1c` |
| 2224 | `P.211_N.52_211/52/00_1` | DEWAN SERBAGUNA KEM PLKN BUMIMAS, SIBU | d | `P.211_N.52_211/52/00_1d` |

#### 312. `P.211_N.52_211/52/00_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2220 | `P.211_N.52_211/52/00_2` | DEWAN SERBAGUNA BATALION 10 JALAN UPPER LANANG, SIBU | a | `P.211_N.52_211/52/00_2a` |
| 2223 | `P.211_N.52_211/52/00_2` | DEWAN SERBAGUNA KEM PLKN JUNACO PARK, SIBU | b | `P.211_N.52_211/52/00_2b` |

#### 313. `P.211_N.52_211/52/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2225 | `P.211_N.52_211/52/01_1` | SEKOLAH KEBANGSAAN BATU WONG | a | `P.211_N.52_211/52/01_1a` |
| 2228 | `P.211_N.52_211/52/01_1` | SEKOLAH KEBANGSAAN NANGA ASSAN | b | `P.211_N.52_211/52/01_1b` |
| 2231 | `P.211_N.52_211/52/01_1` | SEKOLAH KEBANGSAAN TANJUNG LATAP | c | `P.211_N.52_211/52/01_1c` |
| 2232 | `P.211_N.52_211/52/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SING MING | d | `P.211_N.52_211/52/01_1d` |

#### 314. `P.211_N.52_211/52/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2226 | `P.211_N.52_211/52/01_2` | SEKOLAH KEBANGSAAN BATU WONG | a | `P.211_N.52_211/52/01_2a` |
| 2229 | `P.211_N.52_211/52/01_2` | SEKOLAH KEBANGSAAN NANGA ASSAN | b | `P.211_N.52_211/52/01_2b` |

#### 315. `P.211_N.52_211/52/01_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2227 | `P.211_N.52_211/52/01_3` | SEKOLAH KEBANGSAAN BATU WONG | a | `P.211_N.52_211/52/01_3a` |
| 2230 | `P.211_N.52_211/52/01_3` | SEKOLAH KEBANGSAAN NANGA ASSAN | b | `P.211_N.52_211/52/01_3b` |

#### 316. `P.211_N.52_211/52/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2233 | `P.211_N.52_211/52/02_1` | SEKOLAH KEBANGSAAN SG NAMAN | a | `P.211_N.52_211/52/02_1a` |
| 2234 | `P.211_N.52_211/52/02_1` | SEKOLAH KEBANGSAAN ULU SG NAMAN | b | `P.211_N.52_211/52/02_1b` |
| 2236 | `P.211_N.52_211/52/02_1` | SEKOLAH KEBANGSAAN ULU DURIN KIBA | c | `P.211_N.52_211/52/02_1c` |

#### 317. `P.211_N.52_211/52/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2237 | `P.211_N.52_211/52/03_1` | SEKOLAH KEBANGSAAN SG DURIN | a | `P.211_N.52_211/52/03_1a` |
| 2240 | `P.211_N.52_211/52/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HING UNG | b | `P.211_N.52_211/52/03_1b` |
| 2241 | `P.211_N.52_211/52/03_1` | SEKOLAH KEBANGSAAN NANGA PAK | c | `P.211_N.52_211/52/03_1c` |
| 2242 | `P.211_N.52_211/52/03_1` | SEKOLAH KEBANGSAAN SG NIBONG | d | `P.211_N.52_211/52/03_1d` |

#### 318. `P.211_N.52_211/52/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2238 | `P.211_N.52_211/52/03_2` | SEKOLAH KEBANGSAAN SG DURIN | a | `P.211_N.52_211/52/03_2a` |
| 2243 | `P.211_N.52_211/52/03_2` | SEKOLAH KEBANGSAAN SG NIBONG | b | `P.211_N.52_211/52/03_2b` |

#### 319. `P.211_N.52_211/52/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2244 | `P.211_N.52_211/52/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) YONG SHING | a | `P.211_N.52_211/52/04_1a` |
| 2246 | `P.211_N.52_211/52/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KWONG KOK | b | `P.211_N.52_211/52/04_1b` |

#### 320. `P.211_N.52_211/52/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2245 | `P.211_N.52_211/52/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) YONG SHING | a | `P.211_N.52_211/52/04_2a` |
| 2247 | `P.211_N.52_211/52/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) KWONG KOK | b | `P.211_N.52_211/52/04_2b` |

#### 321. `P.211_N.52_211/52/08_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2269 | `P.211_N.52_211/52/08_1` | SEKOLAH KEBANGSAAN SG MENYAN | a | `P.211_N.52_211/52/08_1a` |
| 2271 | `P.211_N.52_211/52/08_1` | SEKOLAH KEBANGSAAN ULU SG SENGAN | b | `P.211_N.52_211/52/08_1b` |
| 2272 | `P.211_N.52_211/52/08_1` | SEKOLAH JENIS KEBANGSAAN (CINA) SANG MING | c | `P.211_N.52_211/52/08_1c` |

#### 322. `P.211_N.52_211/52/08_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2270 | `P.211_N.52_211/52/08_2` | SEKOLAH KEBANGSAAN SG MENYAN | a | `P.211_N.52_211/52/08_2a` |
| 2273 | `P.211_N.52_211/52/08_2` | SEKOLAH JENIS KEBANGSAAN (CINA) SANG MING | b | `P.211_N.52_211/52/08_2b` |

#### 323. `P.211_N.52_211/52/09_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2274 | `P.211_N.52_211/52/09_1` | SEKOLAH KEBANGSAAN ULU SUNGAI SALIM | a | `P.211_N.52_211/52/09_1a` |
| 2276 | `P.211_N.52_211/52/09_1` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN HUA | b | `P.211_N.52_211/52/09_1b` |

#### 324. `P.211_N.52_211/52/09_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2275 | `P.211_N.52_211/52/09_2` | SEKOLAH KEBANGSAAN ULU SUNGAI SALIM | a | `P.211_N.52_211/52/09_2a` |
| 2277 | `P.211_N.52_211/52/09_2` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN HUA | b | `P.211_N.52_211/52/09_2b` |

#### 325. `P.212_N.53_212/53/01_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2302 | `P.212_N.53_212/53/01_1` | SEKOLAH KEBANGSAAN SG RASAU | a | `P.212_N.53_212/53/01_1a` |
| 2304 | `P.212_N.53_212/53/01_1` | RH. LIMBANG SG. SEBEDIL | b | `P.212_N.53_212/53/01_1b` |
| 2305 | `P.212_N.53_212/53/01_1` | SEKOLAH KEBANGSAAN NANGA TUTUS | c | `P.212_N.53_212/53/01_1c` |
| 2306 | `P.212_N.53_212/53/01_1` | SEKOLAH KEBANGSAAN SG PASSAI | d | `P.212_N.53_212/53/01_1d` |
| 2309 | `P.212_N.53_212/53/01_1` | SEKOLAH KEBANGSAAN SG PINANG | e | `P.212_N.53_212/53/01_1e` |
| 2310 | `P.212_N.53_212/53/01_1` | SEKOLAH KEBANGSAAN KG BUNGAN | f | `P.212_N.53_212/53/01_1f` |

#### 326. `P.212_N.53_212/53/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2303 | `P.212_N.53_212/53/01_2` | SEKOLAH KEBANGSAAN SG RASAU | a | `P.212_N.53_212/53/01_2a` |
| 2307 | `P.212_N.53_212/53/01_2` | SEKOLAH KEBANGSAAN SG PASSAI | b | `P.212_N.53_212/53/01_2b` |
| 2311 | `P.212_N.53_212/53/01_2` | SEKOLAH KEBANGSAAN KG BUNGAN | c | `P.212_N.53_212/53/01_2c` |

#### 327. `P.212_N.53_212/53/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2312 | `P.212_N.53_212/53/02_1` | SEKOLAH KEBANGSAAN RANTAU PANJANG | a | `P.212_N.53_212/53/02_1a` |
| 2314 | `P.212_N.53_212/53/02_1` | SEKOLAH KEBANGSAAN TANJUNG PENASU | b | `P.212_N.53_212/53/02_1b` |
| 2317 | `P.212_N.53_212/53/02_1` | SEKOLAH KEBANGSAAN SG AUP | c | `P.212_N.53_212/53/02_1c` |

#### 328. `P.212_N.53_212/53/02_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2313 | `P.212_N.53_212/53/02_2` | SEKOLAH KEBANGSAAN RANTAU PANJANG | a | `P.212_N.53_212/53/02_2a` |
| 2315 | `P.212_N.53_212/53/02_2` | SEKOLAH KEBANGSAAN TANJUNG PENASU | b | `P.212_N.53_212/53/02_2b` |
| 2318 | `P.212_N.53_212/53/02_2` | SEKOLAH KEBANGSAAN SG AUP | c | `P.212_N.53_212/53/02_2c` |

#### 329. `P.212_N.53_212/53/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2316 | `P.212_N.53_212/53/02_3` | SEKOLAH KEBANGSAAN TANJUNG PENASU | a | `P.212_N.53_212/53/02_3a` |
| 2319 | `P.212_N.53_212/53/02_3` | SEKOLAH KEBANGSAAN SG AUP | b | `P.212_N.53_212/53/02_3b` |

#### 330. `P.212_N.53_212/53/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2320 | `P.212_N.53_212/53/03_1` | DEWAN MASYARAKAT BAWANG ASSAN | a | `P.212_N.53_212/53/03_1a` |
| 2322 | `P.212_N.53_212/53/03_1` | PERPUSTAKAAN DESA BAWANG ASSAN | b | `P.212_N.53_212/53/03_1b` |
| 2324 | `P.212_N.53_212/53/03_1` | SEKOLAH KEBANGSAAN TANJUNG BEKAKAP | c | `P.212_N.53_212/53/03_1c` |
| 2327 | `P.212_N.53_212/53/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) BOI ING | d | `P.212_N.53_212/53/03_1d` |

#### 331. `P.212_N.53_212/53/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2321 | `P.212_N.53_212/53/03_2` | DEWAN MASYARAKAT BAWANG ASSAN | a | `P.212_N.53_212/53/03_2a` |
| 2323 | `P.212_N.53_212/53/03_2` | PERPUSTAKAAN DESA BAWANG ASSAN | b | `P.212_N.53_212/53/03_2b` |
| 2325 | `P.212_N.53_212/53/03_2` | SEKOLAH KEBANGSAAN TANJUNG BEKAKAP | c | `P.212_N.53_212/53/03_2c` |

#### 332. `P.212_N.53_212/53/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2328 | `P.212_N.53_212/53/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHAO SU | a | `P.212_N.53_212/53/04_1a` |
| 2329 | `P.212_N.53_212/53/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KAI NANG | b | `P.212_N.53_212/53/04_1b` |

#### 333. `P.212_N.53_212/53/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2333 | `P.212_N.53_212/53/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) DO NANG | a | `P.212_N.53_212/53/07_1a` |
| 2334 | `P.212_N.53_212/53/07_1` | SEKOLAH MENENGAH KEBANGSAAN KWONG HUA MIDDLE | b | `P.212_N.53_212/53/07_1b` |

#### 334. `P.212_N.53_212/53/10_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2339 | `P.212_N.53_212/53/10_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TING SING | a | `P.212_N.53_212/53/10_1a` |
| 2343 | `P.212_N.53_212/53/10_1` | SEKOLAH JENIS KEBANGSAAN (CINA) ING GUONG | b | `P.212_N.53_212/53/10_1b` |

#### 335. `P.212_N.53_212/53/10_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2340 | `P.212_N.53_212/53/10_2` | SEKOLAH JENIS KEBANGSAAN (CINA) TING SING | a | `P.212_N.53_212/53/10_2a` |
| 2344 | `P.212_N.53_212/53/10_2` | SEKOLAH JENIS KEBANGSAAN (CINA) ING GUONG | b | `P.212_N.53_212/53/10_2b` |

#### 336. `P.212_N.55_212/55/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2431 | `P.212_N.55_212/55/00_1` | DEWAN KEM RASCOM SIBU | a | `P.212_N.55_212/55/00_1a` |
| 2434 | `P.212_N.55_212/55/00_1` | DEWAN KEM OYA SIBU | b | `P.212_N.55_212/55/00_1b` |

#### 337. `P.212_N.55_212/55/00_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2432 | `P.212_N.55_212/55/00_2` | DEWAN KEM RASCOM SIBU | a | `P.212_N.55_212/55/00_2a` |
| 2435 | `P.212_N.55_212/55/00_2` | DEWAN KEM OYA SIBU | b | `P.212_N.55_212/55/00_2b` |

#### 338. `P.212_N.55_212/55/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2446 | `P.212_N.55_212/55/04_1` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_1a` |
| 2452 | `P.212_N.55_212/55/04_1` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_1b` |

#### 339. `P.212_N.55_212/55/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2447 | `P.212_N.55_212/55/04_2` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_2a` |
| 2453 | `P.212_N.55_212/55/04_2` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_2b` |

#### 340. `P.212_N.55_212/55/04_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2448 | `P.212_N.55_212/55/04_3` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_3a` |
| 2454 | `P.212_N.55_212/55/04_3` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_3b` |

#### 341. `P.212_N.55_212/55/04_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2449 | `P.212_N.55_212/55/04_4` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_4a` |
| 2455 | `P.212_N.55_212/55/04_4` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_4b` |

#### 342. `P.212_N.55_212/55/04_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2450 | `P.212_N.55_212/55/04_5` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_5a` |
| 2456 | `P.212_N.55_212/55/04_5` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_5b` |

#### 343. `P.212_N.55_212/55/04_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2451 | `P.212_N.55_212/55/04_6` | SEKOLAH KEBANGSAAN SIBU BANDARAN NO 3 | a | `P.212_N.55_212/55/04_6a` |
| 2457 | `P.212_N.55_212/55/04_6` | SEKOLAH KEBANGSAAN ABANG ALI | b | `P.212_N.55_212/55/04_6b` |

#### 344. `P.212_N.55_212/55/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2458 | `P.212_N.55_212/55/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) KIANG HIN | a | `P.212_N.55_212/55/05_1a` |
| 2460 | `P.212_N.55_212/55/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) GUONG MING | b | `P.212_N.55_212/55/05_1b` |

#### 345. `P.212_N.55_212/55/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2461 | `P.212_N.55_212/55/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN CHIN | a | `P.212_N.55_212/55/06_1a` |
| 2465 | `P.212_N.55_212/55/06_1` | SEKOLAH KEBANGSAAN ULU SG MERAH | b | `P.212_N.55_212/55/06_1b` |
| 2473 | `P.212_N.55_212/55/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) NANG SANG | c | `P.212_N.55_212/55/06_1c` |
| 2474 | `P.212_N.55_212/55/06_1` | SEKOLAH JENIS KEBANGSAAN (CINA) DUNG SANG | d | `P.212_N.55_212/55/06_1d` |

#### 346. `P.212_N.55_212/55/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2462 | `P.212_N.55_212/55/06_2` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN CHIN | a | `P.212_N.55_212/55/06_2a` |
| 2466 | `P.212_N.55_212/55/06_2` | SEKOLAH KEBANGSAAN ULU SG MERAH | b | `P.212_N.55_212/55/06_2b` |

#### 347. `P.212_N.55_212/55/06_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2463 | `P.212_N.55_212/55/06_3` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN CHIN | a | `P.212_N.55_212/55/06_3a` |
| 2467 | `P.212_N.55_212/55/06_3` | SEKOLAH KEBANGSAAN ULU SG MERAH | b | `P.212_N.55_212/55/06_3b` |

#### 348. `P.212_N.55_212/55/06_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2464 | `P.212_N.55_212/55/06_4` | SEKOLAH JENIS KEBANGSAAN (CINA) THIAN CHIN | a | `P.212_N.55_212/55/06_4a` |
| 2468 | `P.212_N.55_212/55/06_4` | SEKOLAH KEBANGSAAN ULU SG MERAH | b | `P.212_N.55_212/55/06_4b` |

#### 349. `P.213_N.57_213/57/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2526 | `P.213_N.57_213/57/05_1` | SEKOLAH KEBANGSAAN DIJIH | a | `P.213_N.57_213/57/05_1a` |
| 2527 | `P.213_N.57_213/57/05_1` | SEKOLAH KEBANGSAAN LUBOK BEMBAN | b | `P.213_N.57_213/57/05_1b` |

#### 350. `P.213_N.58_213/58/08_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2562 | `P.213_N.58_213/58/08_1` | SEKOLAH KEBANGSAAN KAMPUNG TEH | a | `P.213_N.58_213/58/08_1a` |
| 2564 | `P.213_N.58_213/58/08_1` | SEKOLAH KEBANGSAAN KAMPUNG SAU | b | `P.213_N.58_213/58/08_1b` |
| 2565 | `P.213_N.58_213/58/08_1` | DEWAN SERBAGUNA KPG. SESOK BARU | c | `P.213_N.58_213/58/08_1c` |

#### 351. `P.214_N.59_214/59/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2578 | `P.214_N.59_214/59/01_1` | SEKOLAH KEBANGSAAN ULU SUNGAI SIONG | a | `P.214_N.59_214/59/01_1a` |
| 2579 | `P.214_N.59_214/59/01_1` | SEKOLAH MENENGAH KEBANGSAAN LUAR BANDAR NO 1 | b | `P.214_N.59_214/59/01_1b` |
| 2582 | `P.214_N.59_214/59/01_1` | SEKOLAH KEBANGSAAN SG. SIONG TENGAH | c | `P.214_N.59_214/59/01_1c` |
| 2583 | `P.214_N.59_214/59/01_1` | SEKOLAH KEBANGSAAN SG SIONG | d | `P.214_N.59_214/59/01_1d` |

#### 352. `P.214_N.59_214/59/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2584 | `P.214_N.59_214/59/02_1` | RH. JANGIT ANAK LANTIENG | a | `P.214_N.59_214/59/02_1a` |
| 2585 | `P.214_N.59_214/59/02_1` | SEKOLAH KEBANGSAAN BATU 36 | b | `P.214_N.59_214/59/02_1b` |
| 2587 | `P.214_N.59_214/59/02_1` | SEKOLAH KEBANGSAAN ST MARK | c | `P.214_N.59_214/59/02_1c` |
| 2588 | `P.214_N.59_214/59/02_1` | SEKOLAH KEBANGSAAN NANGA TAJAM | d | `P.214_N.59_214/59/02_1d` |

#### 353. `P.214_N.59_214/59/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2586 | `P.214_N.59_214/59/02_2` | SEKOLAH KEBANGSAAN BATU 36 | a | `P.214_N.59_214/59/02_2a` |
| 2589 | `P.214_N.59_214/59/02_2` | SEKOLAH KEBANGSAAN NANGA TAJAM | b | `P.214_N.59_214/59/02_2b` |

#### 354. `P.214_N.59_214/59/03_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2590 | `P.214_N.59_214/59/03_1` | SEKOLAH KEBANGSAAN PENGHULU IMBAN | a | `P.214_N.59_214/59/03_1a` |
| 2593 | `P.214_N.59_214/59/03_1` | RH. SARIKEI | b | `P.214_N.59_214/59/03_1b` |
| 2595 | `P.214_N.59_214/59/03_1` | RH. LANTING ANAK BUDA | c | `P.214_N.59_214/59/03_1c` |
| 2596 | `P.214_N.59_214/59/03_1` | DEWAN MASYARAKAT SEKUAU | d | `P.214_N.59_214/59/03_1d` |
| 2598 | `P.214_N.59_214/59/03_1` | SEKOLAH KEBANGSAAN SUNGAI PAKOH | e | `P.214_N.59_214/59/03_1e` |
| 2599 | `P.214_N.59_214/59/03_1` | RH. TUJOH BEJAIT ULU OYA | f | `P.214_N.59_214/59/03_1f` |

#### 355. `P.214_N.59_214/59/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2591 | `P.214_N.59_214/59/03_2` | SEKOLAH KEBANGSAAN PENGHULU IMBAN | a | `P.214_N.59_214/59/03_2a` |
| 2594 | `P.214_N.59_214/59/03_2` | RH. SARIKEI | b | `P.214_N.59_214/59/03_2b` |
| 2597 | `P.214_N.59_214/59/03_2` | DEWAN MASYARAKAT SEKUAU | c | `P.214_N.59_214/59/03_2c` |

#### 356. `P.214_N.59_214/59/04_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2600 | `P.214_N.59_214/59/04_1` | SEKOLAH KEBANGSAAN SG ARAU | a | `P.214_N.59_214/59/04_1a` |
| 2601 | `P.214_N.59_214/59/04_1` | RH. TEGONG ULU SELANGAU | b | `P.214_N.59_214/59/04_1b` |
| 2602 | `P.214_N.59_214/59/04_1` | SEKOLAH KEBANGSAAN NANGA KUA | c | `P.214_N.59_214/59/04_1c` |
| 2604 | `P.214_N.59_214/59/04_1` | RH. PENGARAH NG. SELABI | d | `P.214_N.59_214/59/04_1d` |
| 2605 | `P.214_N.59_214/59/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TONG AH | e | `P.214_N.59_214/59/04_1e` |
| 2612 | `P.214_N.59_214/59/04_1` | RH. PENGUANG ANAK NYANGAU | f | `P.214_N.59_214/59/04_1f` |
| 2613 | `P.214_N.59_214/59/04_1` | RH. JUDAN AK JARAU SG. KUA BATU 43 JLN. SIBU/BINTULU | g | `P.214_N.59_214/59/04_1g` |

#### 357. `P.214_N.59_214/59/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2603 | `P.214_N.59_214/59/04_2` | SEKOLAH KEBANGSAAN NANGA KUA | a | `P.214_N.59_214/59/04_2a` |
| 2606 | `P.214_N.59_214/59/04_2` | SEKOLAH JENIS KEBANGSAAN (CINA) TONG AH | b | `P.214_N.59_214/59/04_2b` |

#### 358. `P.214_N.59_214/59/05_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2614 | `P.214_N.59_214/59/05_1` | SEKOLAH KEBANGSAAN SG BULOH | a | `P.214_N.59_214/59/05_1a` |
| 2616 | `P.214_N.59_214/59/05_1` | SEKOLAH KEBANGSAAN SG SEPIRING/SG TEPUS | b | `P.214_N.59_214/59/05_1b` |
| 2618 | `P.214_N.59_214/59/05_1` | SEKOLAH KEBANGSAAN KUALA LEMAI | c | `P.214_N.59_214/59/05_1c` |
| 2619 | `P.214_N.59_214/59/05_1` | SEKOLAH MENENGAH KEBANGSAAN ULU BALINGIAN | d | `P.214_N.59_214/59/05_1d` |
| 2621 | `P.214_N.59_214/59/05_1` | SEKOLAH KEBANGSAAN SG KEMENA | e | `P.214_N.59_214/59/05_1e` |
| 2622 | `P.214_N.59_214/59/05_1` | SEKOLAH KEBANGSAAN KUALA PELUGAU | f | `P.214_N.59_214/59/05_1f` |

#### 359. `P.214_N.59_214/59/05_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2615 | `P.214_N.59_214/59/05_2` | SEKOLAH KEBANGSAAN SG BULOH | a | `P.214_N.59_214/59/05_2a` |
| 2617 | `P.214_N.59_214/59/05_2` | SEKOLAH KEBANGSAAN SG SEPIRING/SG TEPUS | b | `P.214_N.59_214/59/05_2b` |
| 2620 | `P.214_N.59_214/59/05_2` | SEKOLAH MENENGAH KEBANGSAAN ULU BALINGIAN | c | `P.214_N.59_214/59/05_2c` |

#### 360. `P.214_N.60_214/60/02_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2629 | `P.214_N.60_214/60/02_1` | SEKOLAH KEBANGSAAN SG ANAK | a | `P.214_N.60_214/60/02_1a` |
| 2630 | `P.214_N.60_214/60/02_1` | RH. JIMIN NG. BAWANG | b | `P.214_N.60_214/60/02_1b` |
| 2631 | `P.214_N.60_214/60/02_1` | SEKOLAH KEBANGSAAN SG TAU | c | `P.214_N.60_214/60/02_1c` |
| 2632 | `P.214_N.60_214/60/02_1` | SEKOLAH KEBANGSAAN SG BAWANG | d | `P.214_N.60_214/60/02_1d` |
| 2633 | `P.214_N.60_214/60/02_1` | SEKOLAH KEBANGSAAN ULU SG ARIP | e | `P.214_N.60_214/60/02_1e` |
| 2635 | `P.214_N.60_214/60/02_1` | SEKOLAH KEBANGSAAN IBAN UNION | f | `P.214_N.60_214/60/02_1f` |

#### 361. `P.214_N.60_214/60/17_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2655 | `P.214_N.60_214/60/17_1` | RH. EDWIN NG. KULAU KAKUS | a | `P.214_N.60_214/60/17_1a` |
| 2656 | `P.214_N.60_214/60/17_1` | RH. BATOK BILONG, LONG BEYAK | b | `P.214_N.60_214/60/17_1b` |

#### 362. `P.214_N.60_214/60/25_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2664 | `P.214_N.60_214/60/25_1` | TADIKA KEMAS RH. MAT SEPARAI | a | `P.214_N.60_214/60/25_1a` |
| 2665 | `P.214_N.60_214/60/25_1` | SEKOLAH KEBANGSAAN KELAWIT | b | `P.214_N.60_214/60/25_1b` |

#### 363. `P.215_N.61_215/61/01_1` (12 occurrences, 12 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2669 | `P.215_N.61_215/61/01_1` | RH. BUNYAU, PULAU PISANG, BATANG RAJANG | a | `P.215_N.61_215/61/01_1a` |
| 2670 | `P.215_N.61_215/61/01_1` | RH. MENGGA, NG. SENUANG, BATANG RAJANG | b | `P.215_N.61_215/61/01_1b` |
| 2671 | `P.215_N.61_215/61/01_1` | RH. JELANI, KERANGAN BANGAT, BATANG RAJANG | c | `P.215_N.61_215/61/01_1c` |
| 2672 | `P.215_N.61_215/61/01_1` | SEKOLAH KEBANGSAAN NG ENCHEREMIN | d | `P.215_N.61_215/61/01_1d` |
| 2673 | `P.215_N.61_215/61/01_1` | RH. SELIONG, ULU ENCHEREMIN | e | `P.215_N.61_215/61/01_1e` |
| 2674 | `P.215_N.61_215/61/01_1` | RH. RABAR, NG. STAPANG ULU, BATANG RAJANG | f | `P.215_N.61_215/61/01_1f` |
| 2675 | `P.215_N.61_215/61/01_1` | RH. JENGENG, NG. MELA, BATANG RAJANG | g | `P.215_N.61_215/61/01_1g` |
| 2676 | `P.215_N.61_215/61/01_1` | SEKOLAH KEBANGSAAN NG PELAGUS | h | `P.215_N.61_215/61/01_1h` |
| 2677 | `P.215_N.61_215/61/01_1` | RH. ACHONG, NG. BENIN, BATANG RAJANG | i | `P.215_N.61_215/61/01_1i` |
| 2678 | `P.215_N.61_215/61/01_1` | SEKOLAH KEBANGSAAN ULU PELAGUS | j | `P.215_N.61_215/61/01_1j` |
| 2679 | `P.215_N.61_215/61/01_1` | RH. ANTAU, NG. ARA, SG. PELAGUS | k | `P.215_N.61_215/61/01_1k` |
| 2680 | `P.215_N.61_215/61/01_1` | RH. GEORGE, NG. SEBUNUT, SG. PELAGUS | l | `P.215_N.61_215/61/01_1l` |

#### 364. `P.215_N.61_215/61/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2683 | `P.215_N.61_215/61/03_1` | SEKOLAH KEBANGSAAN NG PERARAN | a | `P.215_N.61_215/61/03_1a` |
| 2685 | `P.215_N.61_215/61/03_1` | RH. BUDIT, NG. MEKEY, BATANG RAJANG | b | `P.215_N.61_215/61/03_1b` |
| 2686 | `P.215_N.61_215/61/03_1` | RH. JELANIE, NG. SAMA, BATANG RAJANG | c | `P.215_N.61_215/61/03_1c` |
| 2687 | `P.215_N.61_215/61/03_1` | RH. SABANG, NG. BUYA, BATANG RAJANG | d | `P.215_N.61_215/61/03_1d` |

#### 365. `P.215_N.61_215/61/04_1` (18 occurrences, 18 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2688 | `P.215_N.61_215/61/04_1` | RH. JOHNNY, NG. AUN, SG. SUT | a | `P.215_N.61_215/61/04_1a` |
| 2689 | `P.215_N.61_215/61/04_1` | RH. SANGGAU, KERANGAN PULAU, SG. SUT | b | `P.215_N.61_215/61/04_1b` |
| 2690 | `P.215_N.61_215/61/04_1` | RH. BELIKAU @ IKAU, NG. SEKERANGAN, SG. SUT | c | `P.215_N.61_215/61/04_1c` |
| 2691 | `P.215_N.61_215/61/04_1` | RH. JUGAH, EMPERAN LIMAU, SG. SUT | d | `P.215_N.61_215/61/04_1d` |
| 2692 | `P.215_N.61_215/61/04_1` | RH. BARAOH, TEMBAWAI SANDONG BAROH, SG. SUT | e | `P.215_N.61_215/61/04_1e` |
| 2693 | `P.215_N.61_215/61/04_1` | RH. EMAK, NG. ANTAROH, SG. SUT | f | `P.215_N.61_215/61/04_1f` |
| 2694 | `P.215_N.61_215/61/04_1` | RH. AI, ENTAWAU ULU, SG. BENA, SUT | g | `P.215_N.61_215/61/04_1g` |
| 2695 | `P.215_N.61_215/61/04_1` | RH. MAMOT, ENTAWAU ILI, SG. BENA, SUT | h | `P.215_N.61_215/61/04_1h` |
| 2696 | `P.215_N.61_215/61/04_1` | RH. TAING, NG. BENA, SG. SUT | i | `P.215_N.61_215/61/04_1i` |
| 2697 | `P.215_N.61_215/61/04_1` | SEKOLAH KEBANGSAAN NG BENA | j | `P.215_N.61_215/61/04_1j` |
| 2698 | `P.215_N.61_215/61/04_1` | RH. NYAWAI, NG. LUN ILI, SG. BENA | k | `P.215_N.61_215/61/04_1k` |
| 2699 | `P.215_N.61_215/61/04_1` | RH. NAGA, NG. MERATING, SG. BENA | l | `P.215_N.61_215/61/04_1l` |
| 2700 | `P.215_N.61_215/61/04_1` | SEKOLAH KEBANGSAAN NG KEBIAW | m | `P.215_N.61_215/61/04_1m` |
| 2701 | `P.215_N.61_215/61/04_1` | SEKOLAH KEBANGSAAN LEPONG BALLEH | n | `P.215_N.61_215/61/04_1n` |
| 2703 | `P.215_N.61_215/61/04_1` | RH. GAWAN, EMPERAN PISANG | o | `P.215_N.61_215/61/04_1o` |
| 2704 | `P.215_N.61_215/61/04_1` | SEKOLAH KEBANGSAAN NG MUJONG | p | `P.215_N.61_215/61/04_1p` |
| 2705 | `P.215_N.61_215/61/04_1` | SEKOLAH KEBANGSAAN NG BAWAI | q | `P.215_N.61_215/61/04_1q` |
| 2707 | `P.215_N.61_215/61/04_1` | RH. BENGAU, NG. SUT, BATANG BALLEH | r | `P.215_N.61_215/61/04_1r` |

#### 366. `P.215_N.61_215/61/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2702 | `P.215_N.61_215/61/04_2` | SEKOLAH KEBANGSAAN LEPONG BALLEH | a | `P.215_N.61_215/61/04_2a` |
| 2706 | `P.215_N.61_215/61/04_2` | SEKOLAH KEBANGSAAN NG BAWAI | b | `P.215_N.61_215/61/04_2b` |

#### 367. `P.215_N.62_215/62/01_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2710 | `P.215_N.62_215/62/01_1` | RH. BUNI, SG. LAJAN | a | `P.215_N.62_215/62/01_1a` |
| 2711 | `P.215_N.62_215/62/01_1` | RH. EDWARD, SG. LAJAN | b | `P.215_N.62_215/62/01_1b` |
| 2712 | `P.215_N.62_215/62/01_1` | RH. DARLIN, SG. LIJAU | c | `P.215_N.62_215/62/01_1c` |
| 2713 | `P.215_N.62_215/62/01_1` | RH. TIMOTHY ASSON, NG. MANAP | d | `P.215_N.62_215/62/01_1d` |
| 2714 | `P.215_N.62_215/62/01_1` | RH. MUNI, ULU SG. MANAP | e | `P.215_N.62_215/62/01_1e` |
| 2715 | `P.215_N.62_215/62/01_1` | RH. TANANG, NG. EMBUAU | f | `P.215_N.62_215/62/01_1f` |
| 2716 | `P.215_N.62_215/62/01_1` | RH. NGITAR, LUBOK RERONG, SG. SONG | g | `P.215_N.62_215/62/01_1g` |
| 2717 | `P.215_N.62_215/62/01_1` | RH. SUGAI, NG. SEBETONG, SG. SONG | h | `P.215_N.62_215/62/01_1h` |

#### 368. `P.215_N.62_215/62/02_1` (12 occurrences, 12 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2718 | `P.215_N.62_215/62/02_1` | RH. KENNEDY, NG. BEKAKAP, ULU SG. IRAN | a | `P.215_N.62_215/62/02_1a` |
| 2719 | `P.215_N.62_215/62/02_1` | RH. RUNI, NG. SELEPONG, ULU SG. IRAN | b | `P.215_N.62_215/62/02_1b` |
| 2720 | `P.215_N.62_215/62/02_1` | SEKOLAH KEBANGSAAN NG DALAI | c | `P.215_N.62_215/62/02_1c` |
| 2721 | `P.215_N.62_215/62/02_1` | RH. MUSIN, NG. NANSANG, SG. IRAN | d | `P.215_N.62_215/62/02_1d` |
| 2722 | `P.215_N.62_215/62/02_1` | RH. GAWAN, NG. SEBIRAH, SG. IRAN | e | `P.215_N.62_215/62/02_1e` |
| 2723 | `P.215_N.62_215/62/02_1` | RH. CHIRRY, EMPERAN MUNTI, ULU SG. IRAN | f | `P.215_N.62_215/62/02_1f` |
| 2724 | `P.215_N.62_215/62/02_1` | RH. JAMBA, NG. SANTU, NG. IRAN | g | `P.215_N.62_215/62/02_1g` |
| 2725 | `P.215_N.62_215/62/02_1` | SEKOLAH KEBANGSAAN NG TEMALAT | h | `P.215_N.62_215/62/02_1h` |
| 2726 | `P.215_N.62_215/62/02_1` | RH. NOBERT, NG. SESAWA, BATANG RAJANG, SONG | i | `P.215_N.62_215/62/02_1i` |
| 2727 | `P.215_N.62_215/62/02_1` | RH. SERING, NG. TEMIANG, BATANG RAJANG, SONG | j | `P.215_N.62_215/62/02_1j` |
| 2728 | `P.215_N.62_215/62/02_1` | RH. AUSTIN EKAU, NG. SIPAN, BATANG RAJANG, SONG | k | `P.215_N.62_215/62/02_1k` |
| 2729 | `P.215_N.62_215/62/02_1` | SEKOLAH KEBANGSAAN NG BEGUANG | l | `P.215_N.62_215/62/02_1l` |

#### 369. `P.215_N.62_215/62/03_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2731 | `P.215_N.62_215/62/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HIN HUA | a | `P.215_N.62_215/62/03_1a` |
| 2734 | `P.215_N.62_215/62/03_1` | MADRASAH, KPG. GELAM, SONG | b | `P.215_N.62_215/62/03_1b` |

#### 370. `P.215_N.62_215/62/04_1` (11 occurrences, 11 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2735 | `P.215_N.62_215/62/04_1` | RH. JOSLEE, NG. MATALAU, SG. MUSAH | a | `P.215_N.62_215/62/04_1a` |
| 2736 | `P.215_N.62_215/62/04_1` | RH. DEMANG, EMPERAN MENUANG, SG. MUSAH | b | `P.215_N.62_215/62/04_1b` |
| 2737 | `P.215_N.62_215/62/04_1` | RH. BARO, NG. SEMULONG, SG. MUSAH | c | `P.215_N.62_215/62/04_1c` |
| 2738 | `P.215_N.62_215/62/04_1` | RH. JAPOK, NG. SENYARO, NG. MUSAH | d | `P.215_N.62_215/62/04_1d` |
| 2739 | `P.215_N.62_215/62/04_1` | RH. GELANA, NG. TEKALIT | e | `P.215_N.62_215/62/04_1e` |
| 2740 | `P.215_N.62_215/62/04_1` | RH. JALA, NG. MIAW, KATIBAS | f | `P.215_N.62_215/62/04_1f` |
| 2741 | `P.215_N.62_215/62/04_1` | RH. SERIT, ULU SG. TAKAN | g | `P.215_N.62_215/62/04_1g` |
| 2742 | `P.215_N.62_215/62/04_1` | RH. ZACHIUS NYALU, NG. TAKAN, KATIBAS | h | `P.215_N.62_215/62/04_1h` |
| 2743 | `P.215_N.62_215/62/04_1` | RH. CECELIA BUNSU, NG. KEBIAU | i | `P.215_N.62_215/62/04_1i` |
| 2744 | `P.215_N.62_215/62/04_1` | RH. WIL, NG. NYIMOH, KATIBAS | j | `P.215_N.62_215/62/04_1j` |
| 2745 | `P.215_N.62_215/62/04_1` | RH. DELOK, SG. ENGKABAU | k | `P.215_N.62_215/62/04_1k` |

#### 371. `P.215_N.62_215/62/05_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2746 | `P.215_N.62_215/62/05_1` | RH. API, NG. TERUSA, KATIBAS | a | `P.215_N.62_215/62/05_1a` |
| 2747 | `P.215_N.62_215/62/05_1` | SEKOLAH KEBANGSAAN LUBOK IPOH | b | `P.215_N.62_215/62/05_1b` |
| 2748 | `P.215_N.62_215/62/05_1` | RH. SEBUN, NG. MASAK, KATIBAS | c | `P.215_N.62_215/62/05_1c` |
| 2749 | `P.215_N.62_215/62/05_1` | RH. GENDANG, KARANGAN RANGKANG, ULU KATIBAS | d | `P.215_N.62_215/62/05_1d` |
| 2750 | `P.215_N.62_215/62/05_1` | RH. NUGU, NG. SESIBAU, ULU KATIBAS | e | `P.215_N.62_215/62/05_1e` |
| 2751 | `P.215_N.62_215/62/05_1` | RH. JINGGONG, NG. ANCHAU, ULU KATIBAS | f | `P.215_N.62_215/62/05_1f` |

#### 372. `P.215_N.62_215/62/06_1` (10 occurrences, 10 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2752 | `P.215_N.62_215/62/06_1` | RH. KANA, KERANGAN PANJANG, ULU SG. BANGKIT | a | `P.215_N.62_215/62/06_1a` |
| 2753 | `P.215_N.62_215/62/06_1` | RH. AKANG, BATU PIKUL, SG. BANGKIT | b | `P.215_N.62_215/62/06_1b` |
| 2754 | `P.215_N.62_215/62/06_1` | RH. LAYANG, SG. AYAT, ULU BANGKIT | c | `P.215_N.62_215/62/06_1c` |
| 2755 | `P.215_N.62_215/62/06_1` | RH. LASIN, NG. NANSANG, SG. BANGKIT | d | `P.215_N.62_215/62/06_1d` |
| 2756 | `P.215_N.62_215/62/06_1` | RH. BADA, NG. MELUAN, SG. BANGKIT | e | `P.215_N.62_215/62/06_1e` |
| 2757 | `P.215_N.62_215/62/06_1` | SEKOLAH KEBANGSAAN NG BANGKIT | f | `P.215_N.62_215/62/06_1f` |
| 2758 | `P.215_N.62_215/62/06_1` | RH. RIBUT, NG. SERAU, KATIBAS | g | `P.215_N.62_215/62/06_1g` |
| 2759 | `P.215_N.62_215/62/06_1` | RH. DAGUM, NG. MAKUT, KATIBAS | h | `P.215_N.62_215/62/06_1h` |
| 2760 | `P.215_N.62_215/62/06_1` | RH. BANGAU, NG. ENTUAT | i | `P.215_N.62_215/62/06_1i` |
| 2761 | `P.215_N.62_215/62/06_1` | RH. IJAU, NG. MUKIH, KATIBAS | j | `P.215_N.62_215/62/06_1j` |

#### 373. `P.215_N.62_215/62/07_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2762 | `P.215_N.62_215/62/07_1` | RH. GILBERT NYADANG, NG. SEPUNGGOK, TEKALIT | a | `P.215_N.62_215/62/07_1a` |
| 2763 | `P.215_N.62_215/62/07_1` | RH. MELAYU, NG. LATONG, TEKALIT | b | `P.215_N.62_215/62/07_1b` |
| 2764 | `P.215_N.62_215/62/07_1` | RH. SIBAT, NG. SEGENDANG, ULU SG. JANAN, TEKALIT | c | `P.215_N.62_215/62/07_1c` |
| 2765 | `P.215_N.62_215/62/07_1` | SEKOLAH KEBANGSAAN NG JANAN | d | `P.215_N.62_215/62/07_1d` |
| 2767 | `P.215_N.62_215/62/07_1` | RH. CHANGAI, NG. TENGANGAI, TEKALIT | e | `P.215_N.62_215/62/07_1e` |
| 2768 | `P.215_N.62_215/62/07_1` | RH. ENDAH, NG. SEPAYANG, SG. TEKALIT | f | `P.215_N.62_215/62/07_1f` |
| 2769 | `P.215_N.62_215/62/07_1` | SEKOLAH KEBANGSAAN NG NANSANG | g | `P.215_N.62_215/62/07_1g` |
| 2770 | `P.215_N.62_215/62/07_1` | RH. BARAIN, NG. SEBUNGKANG, SG. TEKALIT | h | `P.215_N.62_215/62/07_1h` |

#### 374. `P.215_N.63_215/63/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2773 | `P.215_N.63_215/63/01_1` | SEKOLAH KEBANGSAAN SUNGAI KAPIT | a | `P.215_N.63_215/63/01_1a` |
| 2776 | `P.215_N.63_215/63/01_1` | RH. BUNDONG, SEBABAI ILI, JALAN BUKIT GORAM | b | `P.215_N.63_215/63/01_1b` |

#### 375. `P.215_N.63_215/63/03_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2782 | `P.215_N.63_215/63/03_1` | SEKOLAH KEBANGSAAN SG MENUAN | a | `P.215_N.63_215/63/03_1a` |
| 2783 | `P.215_N.63_215/63/03_1` | RH. SLI, NG. SEBELU, SG. MENUAN | b | `P.215_N.63_215/63/03_1b` |
| 2784 | `P.215_N.63_215/63/03_1` | RH. MANGGAN, NG. BELINYU, SG. MENUAN | c | `P.215_N.63_215/63/03_1c` |
| 2785 | `P.215_N.63_215/63/03_1` | RH. LUMPOH, NG. METAH, SG. MENUAN | d | `P.215_N.63_215/63/03_1d` |
| 2786 | `P.215_N.63_215/63/03_1` | SEKOLAH KEBANGSAAN LP MENUAN | e | `P.215_N.63_215/63/03_1e` |
| 2787 | `P.215_N.63_215/63/03_1` | RH. JAMBON, NG. ENSILAI, BATANG RAJANG | f | `P.215_N.63_215/63/03_1f` |
| 2788 | `P.215_N.63_215/63/03_1` | RH. DINGAI, SG. GOH ILI, BATANG RAJANG | g | `P.215_N.63_215/63/03_1g` |

#### 376. `P.215_N.63_215/63/04_1` (13 occurrences, 13 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2789 | `P.215_N.63_215/63/04_1` | SEKOLAH KEBANGSAAN ULU YONG | a | `P.215_N.63_215/63/04_1a` |
| 2790 | `P.215_N.63_215/63/04_1` | RH. MANOK MANCHAL, LUBOK MAWANG, SG. YONG | b | `P.215_N.63_215/63/04_1b` |
| 2791 | `P.215_N.63_215/63/04_1` | RH. EMPAWI, NG. SEKUKUT, SG. YONG | c | `P.215_N.63_215/63/04_1c` |
| 2792 | `P.215_N.63_215/63/04_1` | RH. NGELAI, NG. TISA, SG. YONG | d | `P.215_N.63_215/63/04_1d` |
| 2793 | `P.215_N.63_215/63/04_1` | SEKOLAH KEBANGSAAN NG TRUSA | e | `P.215_N.63_215/63/04_1e` |
| 2794 | `P.215_N.63_215/63/04_1` | RH. NUGA, SEMUJAN ILI, BELAWAI | f | `P.215_N.63_215/63/04_1f` |
| 2795 | `P.215_N.63_215/63/04_1` | RH. SELIONG, SEKERANGAN ATAS, SG. BELAWAI | g | `P.215_N.63_215/63/04_1g` |
| 2796 | `P.215_N.63_215/63/04_1` | RH. BENA, NG. SEMA, SG. YONG | h | `P.215_N.63_215/63/04_1h` |
| 2797 | `P.215_N.63_215/63/04_1` | RH. EKAU, NG. SEMAWANG, BATANG RAJANG | i | `P.215_N.63_215/63/04_1i` |
| 2798 | `P.215_N.63_215/63/04_1` | RH. KAYAN, NG. DIA, BATANG RAJANG | j | `P.215_N.63_215/63/04_1j` |
| 2799 | `P.215_N.63_215/63/04_1` | SEKOLAH KEBANGSAAN NG SEGENOK | k | `P.215_N.63_215/63/04_1k` |
| 2800 | `P.215_N.63_215/63/04_1` | RH. BELI, NG. ENSURAI, SG. IBAU | l | `P.215_N.63_215/63/04_1l` |
| 2801 | `P.215_N.63_215/63/04_1` | SEKOLAH KEBANGSAAN NG IBAU | m | `P.215_N.63_215/63/04_1m` |

#### 377. `P.215_N.63_215/63/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2807 | `P.215_N.63_215/63/06_1` | SEKOLAH KEBANGSAAN METHODIST | a | `P.215_N.63_215/63/06_1a` |
| 2812 | `P.215_N.63_215/63/06_1` | RH. JUIN, RANTAU TAPANG, SG. SERANAU | b | `P.215_N.63_215/63/06_1b` |
| 2813 | `P.215_N.63_215/63/06_1` | RH. BARNABAS, KAMPUNG SERIAN, BATANG RAJANG | c | `P.215_N.63_215/63/06_1c` |
| 2814 | `P.215_N.63_215/63/06_1` | RH. AYU, NG. TULIE BARUH | d | `P.215_N.63_215/63/06_1d` |

#### 378. `P.215_N.63_215/63/08_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2816 | `P.215_N.63_215/63/08_1` | RH. BAJA, NG. SELIRIK, SG. ENTANGAI | a | `P.215_N.63_215/63/08_1a` |
| 2817 | `P.215_N.63_215/63/08_1` | SEKOLAH KEBANGSAAN ULU MELIPIS | b | `P.215_N.63_215/63/08_1b` |
| 2818 | `P.215_N.63_215/63/08_1` | RH. UNTAT, EMPERAN BEMBAM, SG. MELIPIS | c | `P.215_N.63_215/63/08_1c` |
| 2819 | `P.215_N.63_215/63/08_1` | RH. ROBERT, NG. SEBATU, BATANG RAJANG | d | `P.215_N.63_215/63/08_1d` |

#### 379. `P.216_N.64_216/64/01_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2821 | `P.216_N.64_216/64/01_1` | RH. LAWAN NG. RAMONG | a | `P.216_N.64_216/64/01_1a` |
| 2822 | `P.216_N.64_216/64/01_1` | RH. MENGGA NG. SEPULAU GAAT | b | `P.216_N.64_216/64/01_1b` |
| 2823 | `P.216_N.64_216/64/01_1` | SEKOLAH KEBANGSAAN NG BALANG | c | `P.216_N.64_216/64/01_1c` |
| 2824 | `P.216_N.64_216/64/01_1` | RH. NADING NG. SELENTANG | d | `P.216_N.64_216/64/01_1d` |
| 2825 | `P.216_N.64_216/64/01_1` | RH. SAGEN, NG. AJAN SG. GAAT | e | `P.216_N.64_216/64/01_1e` |
| 2826 | `P.216_N.64_216/64/01_1` | RH. MELINTANG NG. SEBIRO SG. GAAT | f | `P.216_N.64_216/64/01_1f` |

#### 380. `P.216_N.64_216/64/02_1` (17 occurrences, 17 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2827 | `P.216_N.64_216/64/02_1` | RH. JOS SG. MUJONG | a | `P.216_N.64_216/64/02_1a` |
| 2828 | `P.216_N.64_216/64/02_1` | RH. SENGIANG SG. MUJONG | b | `P.216_N.64_216/64/02_1b` |
| 2829 | `P.216_N.64_216/64/02_1` | SEKOLAH KEBANGSAAN BEBANGAN | c | `P.216_N.64_216/64/02_1c` |
| 2830 | `P.216_N.64_216/64/02_1` | RH. UMBAR (BUNSU) SG. TIAU | d | `P.216_N.64_216/64/02_1d` |
| 2831 | `P.216_N.64_216/64/02_1` | RH. JAMES SAKA | e | `P.216_N.64_216/64/02_1e` |
| 2832 | `P.216_N.64_216/64/02_1` | RH. SENGIANG ANAK USAT, WONG PANTU | f | `P.216_N.64_216/64/02_1f` |
| 2833 | `P.216_N.64_216/64/02_1` | RH. PETER RANTAU BIDAI | g | `P.216_N.64_216/64/02_1g` |
| 2834 | `P.216_N.64_216/64/02_1` | RH. GINDAL PULAU SIBAU | h | `P.216_N.64_216/64/02_1h` |
| 2835 | `P.216_N.64_216/64/02_1` | RH. GAWAN NG. AMAN | i | `P.216_N.64_216/64/02_1i` |
| 2836 | `P.216_N.64_216/64/02_1` | SEKOLAH KEBANGSAAN MUJONG TENGAH | j | `P.216_N.64_216/64/02_1j` |
| 2837 | `P.216_N.64_216/64/02_1` | SEKOLAH KEBANGSAAN LUBOK BAYA | k | `P.216_N.64_216/64/02_1k` |
| 2838 | `P.216_N.64_216/64/02_1` | SEKOLAH KEBANGSAAN NG OYAN | l | `P.216_N.64_216/64/02_1l` |
| 2839 | `P.216_N.64_216/64/02_1` | RH. BANGKONG NG. SEBOLA | m | `P.216_N.64_216/64/02_1m` |
| 2840 | `P.216_N.64_216/64/02_1` | RH. ASUN, LEPONG MUJONG | n | `P.216_N.64_216/64/02_1n` |
| 2841 | `P.216_N.64_216/64/02_1` | RH. MUNGKO, RANTAU LELANGAI | o | `P.216_N.64_216/64/02_1o` |
| 2842 | `P.216_N.64_216/64/02_1` | SEKOLAH KEBANGSAAN OYAN TENGAH | p | `P.216_N.64_216/64/02_1p` |
| 2843 | `P.216_N.64_216/64/02_1` | RH. UNTING NG SEKEROH | q | `P.216_N.64_216/64/02_1q` |

#### 381. `P.216_N.64_216/64/03_1` (15 occurrences, 15 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2844 | `P.216_N.64_216/64/03_1` | SEKOLAH KEBANGSAAN SG TUNOH | a | `P.216_N.64_216/64/03_1a` |
| 2845 | `P.216_N.64_216/64/03_1` | RH. KILLAU SG. MELINAU | b | `P.216_N.64_216/64/03_1b` |
| 2846 | `P.216_N.64_216/64/03_1` | RH. BARANG NG. PANG MELINAU | c | `P.216_N.64_216/64/03_1c` |
| 2847 | `P.216_N.64_216/64/03_1` | RH. ANDING SG. MELINAU | d | `P.216_N.64_216/64/03_1d` |
| 2848 | `P.216_N.64_216/64/03_1` | RH. JALA SG. MELINAU KAPIT | e | `P.216_N.64_216/64/03_1e` |
| 2849 | `P.216_N.64_216/64/03_1` | RH. KALAT ANAK MANJAH, SUNGAI MAJAU | f | `P.216_N.64_216/64/03_1f` |
| 2850 | `P.216_N.64_216/64/03_1` | SEKOLAH KEBANGSAAN LBK MAWANG | g | `P.216_N.64_216/64/03_1g` |
| 2851 | `P.216_N.64_216/64/03_1` | RH. MANSAI SG. MAJAU | h | `P.216_N.64_216/64/03_1h` |
| 2852 | `P.216_N.64_216/64/03_1` | RH. LUGAT NG. MAJAU | i | `P.216_N.64_216/64/03_1i` |
| 2853 | `P.216_N.64_216/64/03_1` | RH. PIOH SG. MAJAU | j | `P.216_N.64_216/64/03_1j` |
| 2854 | `P.216_N.64_216/64/03_1` | RH. BERAUH NG. SEBELANDA, SG.PAKU | k | `P.216_N.64_216/64/03_1k` |
| 2855 | `P.216_N.64_216/64/03_1` | RH. INGUH MIUT SG. PAKU | l | `P.216_N.64_216/64/03_1l` |
| 2856 | `P.216_N.64_216/64/03_1` | RH. JAMBA (GARIT) SG. PAKU | m | `P.216_N.64_216/64/03_1m` |
| 2857 | `P.216_N.64_216/64/03_1` | RH. JANTING NG. SEPINANG | n | `P.216_N.64_216/64/03_1n` |
| 2858 | `P.216_N.64_216/64/03_1` | RH. NISSON (BUNYAU) SG. PAKU | o | `P.216_N.64_216/64/03_1o` |

#### 382. `P.216_N.64_216/64/04_1` (14 occurrences, 14 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2859 | `P.216_N.64_216/64/04_1` | RH. KACHIN SG. STAPANG, GAAT | a | `P.216_N.64_216/64/04_1a` |
| 2860 | `P.216_N.64_216/64/04_1` | RH. SAIT NG. PURO GAAT | b | `P.216_N.64_216/64/04_1b` |
| 2861 | `P.216_N.64_216/64/04_1` | RH. NYAMOK SG. GAAT | c | `P.216_N.64_216/64/04_1c` |
| 2862 | `P.216_N.64_216/64/04_1` | SEKOLAH KEBANGSAAN LEPONG GAAT | d | `P.216_N.64_216/64/04_1d` |
| 2863 | `P.216_N.64_216/64/04_1` | RH. TAJAI (MAMAT) NG. SEMIRAH GAAT | e | `P.216_N.64_216/64/04_1e` |
| 2864 | `P.216_N.64_216/64/04_1` | RH. JOHN KATIL | f | `P.216_N.64_216/64/04_1f` |
| 2865 | `P.216_N.64_216/64/04_1` | RH. BIDOK NG. SEBETONG | g | `P.216_N.64_216/64/04_1g` |
| 2866 | `P.216_N.64_216/64/04_1` | RH. GON, NG. SERIAN BALEH | h | `P.216_N.64_216/64/04_1h` |
| 2867 | `P.216_N.64_216/64/04_1` | RH. WONG ENSONG NG. SEPATA, BALEH | i | `P.216_N.64_216/64/04_1i` |
| 2868 | `P.216_N.64_216/64/04_1` | RH. LAMAU ANAK JENGGUT, TELOK BUING | j | `P.216_N.64_216/64/04_1j` |
| 2869 | `P.216_N.64_216/64/04_1` | RH. PININ ANAK TUJANG NANGA USUN BALEH | k | `P.216_N.64_216/64/04_1k` |
| 2870 | `P.216_N.64_216/64/04_1` | RH. GARE ANAK TIMBANG SG. KAIN | l | `P.216_N.64_216/64/04_1l` |
| 2871 | `P.216_N.64_216/64/04_1` | RH. ENGSONG SG. KAIN | m | `P.216_N.64_216/64/04_1m` |
| 2872 | `P.216_N.64_216/64/04_1` | SEKOLAH KEBANGSAAN NG KAIN | n | `P.216_N.64_216/64/04_1n` |

#### 383. `P.216_N.64_216/64/05_1` (16 occurrences, 16 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2873 | `P.216_N.64_216/64/05_1` | RH. BAJANG ANAK LIANG SG. TELIAI ENTULOH | a | `P.216_N.64_216/64/05_1a` |
| 2874 | `P.216_N.64_216/64/05_1` | RH. JACK SG. ENTULOH | b | `P.216_N.64_216/64/05_1b` |
| 2875 | `P.216_N.64_216/64/05_1` | RH. RICHARD (UNYAT) SG. MERIRAI | c | `P.216_N.64_216/64/05_1c` |
| 2876 | `P.216_N.64_216/64/05_1` | RH. GOYANG APAT SG. MERIRAI | d | `P.216_N.64_216/64/05_1d` |
| 2877 | `P.216_N.64_216/64/05_1` | RH. BANGAU UNDI KERANGAN LAIH | e | `P.216_N.64_216/64/05_1e` |
| 2878 | `P.216_N.64_216/64/05_1` | RH. JANTAI ANAK SIBA BATANG BALEH | f | `P.216_N.64_216/64/05_1f` |
| 2879 | `P.216_N.64_216/64/05_1` | RH. BULLY KERANGAN BESAI BALEH | g | `P.216_N.64_216/64/05_1g` |
| 2880 | `P.216_N.64_216/64/05_1` | RH. SEBUANG NG. MERAMA BALEH | h | `P.216_N.64_216/64/05_1h` |
| 2881 | `P.216_N.64_216/64/05_1` | TADIKA KEMAS, ANTAWAU | i | `P.216_N.64_216/64/05_1i` |
| 2882 | `P.216_N.64_216/64/05_1` | RH. TAJAI SANGGAU NG. SEBIRO BALEH | j | `P.216_N.64_216/64/05_1j` |
| 2883 | `P.216_N.64_216/64/05_1` | SEKOLAH KEBANGSAAN NG SEMPILI | k | `P.216_N.64_216/64/05_1k` |
| 2884 | `P.216_N.64_216/64/05_1` | RH. SAMON ANAK CHEPAU | l | `P.216_N.64_216/64/05_1l` |
| 2885 | `P.216_N.64_216/64/05_1` | RH. ITU SG. PULANG BALEH | m | `P.216_N.64_216/64/05_1m` |
| 2886 | `P.216_N.64_216/64/05_1` | RH. BANSA ANAK LANGGA SEPULAU, BALEH | n | `P.216_N.64_216/64/05_1n` |
| 2887 | `P.216_N.64_216/64/05_1` | RH. JAMIT NG. SEPANGGIL BALEH | o | `P.216_N.64_216/64/05_1o` |
| 2888 | `P.216_N.64_216/64/05_1` | RH. SANA ANAK RUMPANG NG. GAAT | p | `P.216_N.64_216/64/05_1p` |

#### 384. `P.216_N.65_216/65/02_1` (14 occurrences, 14 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2894 | `P.216_N.65_216/65/02_1` | SEKOLAH KEBANGSAAN NG MERIT | a | `P.216_N.65_216/65/02_1a` |
| 2896 | `P.216_N.65_216/65/02_1` | RH. DANGGAT NG. MUJAN | b | `P.216_N.65_216/65/02_1b` |
| 2897 | `P.216_N.65_216/65/02_1` | RH. BERUNDANG NG. ENSAWI | c | `P.216_N.65_216/65/02_1c` |
| 2898 | `P.216_N.65_216/65/02_1` | RH. GUYANG NG. JELI | d | `P.216_N.65_216/65/02_1d` |
| 2899 | `P.216_N.65_216/65/02_1` | RH. ENTILI NG. MUSA | e | `P.216_N.65_216/65/02_1e` |
| 2900 | `P.216_N.65_216/65/02_1` | RH. EMPANG NG. BILAT ULU PAKU | f | `P.216_N.65_216/65/02_1f` |
| 2901 | `P.216_N.65_216/65/02_1` | RH. GENDANG SG. PILA | g | `P.216_N.65_216/65/02_1g` |
| 2902 | `P.216_N.65_216/65/02_1` | RH. AJI NG. IBUN | h | `P.216_N.65_216/65/02_1h` |
| 2903 | `P.216_N.65_216/65/02_1` | RH. MERO LUBOK DABAI | i | `P.216_N.65_216/65/02_1i` |
| 2904 | `P.216_N.65_216/65/02_1` | RH. RANYING ULU METAH | j | `P.216_N.65_216/65/02_1j` |
| 2905 | `P.216_N.65_216/65/02_1` | RH. PILLAI NG. METAH | k | `P.216_N.65_216/65/02_1k` |
| 2906 | `P.216_N.65_216/65/02_1` | RH. SETEPAN NG. PILA | l | `P.216_N.65_216/65/02_1l` |
| 2907 | `P.216_N.65_216/65/02_1` | SEKOLAH KEBANGSAAN NG METAH | m | `P.216_N.65_216/65/02_1m` |
| 2908 | `P.216_N.65_216/65/02_1` | RH. JAMIT NG. LATAP PILLA | n | `P.216_N.65_216/65/02_1n` |

#### 385. `P.216_N.65_216/65/03_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2909 | `P.216_N.65_216/65/03_1` | SEKOLAH KEBANGSAAN PUNAN BA | a | `P.216_N.65_216/65/03_1a` |
| 2911 | `P.216_N.65_216/65/03_1` | UMA PUNAN BIAU | b | `P.216_N.65_216/65/03_1b` |
| 2912 | `P.216_N.65_216/65/03_1` | UMA PUNAN SAMA | c | `P.216_N.65_216/65/03_1c` |
| 2913 | `P.216_N.65_216/65/03_1` | UMA TANJONG LONG PAWAH | d | `P.216_N.65_216/65/03_1d` |
| 2914 | `P.216_N.65_216/65/03_1` | SEKOLAH KEBANGSAAN AIRPORT | e | `P.216_N.65_216/65/03_1e` |
| 2916 | `P.216_N.65_216/65/03_1` | UMA SEKAPAN PIIT | f | `P.216_N.65_216/65/03_1f` |
| 2917 | `P.216_N.65_216/65/03_1` | UMA LONG AMO | g | `P.216_N.65_216/65/03_1g` |

#### 386. `P.216_N.65_216/65/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2910 | `P.216_N.65_216/65/03_2` | SEKOLAH KEBANGSAAN PUNAN BA | a | `P.216_N.65_216/65/03_2a` |
| 2915 | `P.216_N.65_216/65/03_2` | SEKOLAH KEBANGSAAN AIRPORT | b | `P.216_N.65_216/65/03_2b` |

#### 387. `P.216_N.65_216/65/04_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2918 | `P.216_N.65_216/65/04_1` | SEKOLAH KEBANGSAAN ABUN MATU | a | `P.216_N.65_216/65/04_1a` |
| 2921 | `P.216_N.65_216/65/04_1` | UMA KAHEI LONG MAKERO | b | `P.216_N.65_216/65/04_1b` |
| 2922 | `P.216_N.65_216/65/04_1` | UMA AGING LONG DAAH | c | `P.216_N.65_216/65/04_1c` |
| 2923 | `P.216_N.65_216/65/04_1` | UMA LAHANAN LONG SENUANG | d | `P.216_N.65_216/65/04_1d` |
| 2924 | `P.216_N.65_216/65/04_1` | UMA KEJAMAN NEH LONG LITEN | e | `P.216_N.65_216/65/04_1e` |
| 2925 | `P.216_N.65_216/65/04_1` | UMA APAN LONG MENJAWAH | f | `P.216_N.65_216/65/04_1f` |
| 2926 | `P.216_N.65_216/65/04_1` | UMA NYAVING LONG MENJAWAH | g | `P.216_N.65_216/65/04_1g` |
| 2927 | `P.216_N.65_216/65/04_1` | SEKOLAH KEBANGSAAN LONG SEGAHAM | h | `P.216_N.65_216/65/04_1h` |
| 2928 | `P.216_N.65_216/65/04_1` | DEWAN UMA AGING, NAHA JALEI, LONG KEBUHO | i | `P.216_N.65_216/65/04_1i` |

#### 388. `P.216_N.65_216/65/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2932 | `P.216_N.65_216/65/07_1` | SEKOLAH KEBANGSAAN UMA SAMBOP | a | `P.216_N.65_216/65/07_1a` |
| 2933 | `P.216_N.65_216/65/07_1` | LONG NANYAN | b | `P.216_N.65_216/65/07_1b` |

#### 389. `P.216_N.66_216/66/01_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2935 | `P.216_N.66_216/66/01_1` | SEKOLAH KEBANGSAAN METALUN | a | `P.216_N.66_216/66/01_1a` |
| 2936 | `P.216_N.66_216/66/01_1` | RH. PENAN LONG PERAN | b | `P.216_N.66_216/66/01_1b` |
| 2937 | `P.216_N.66_216/66/01_1` | DEWAN SERBAGUNA UMA PENAN LONG WAT | c | `P.216_N.66_216/66/01_1c` |
| 2938 | `P.216_N.66_216/66/01_1` | RH. PENAN LONG LIDEM | d | `P.216_N.66_216/66/01_1d` |
| 2939 | `P.216_N.66_216/66/01_1` | SEKOLAH KEBANGSAAN LUSONG LAKU | e | `P.216_N.66_216/66/01_1e` |
| 2940 | `P.216_N.66_216/66/01_1` | RH. PENAN LONG TANYIT | f | `P.216_N.66_216/66/01_1f` |

#### 390. `P.216_N.66_216/66/02_1` (18 occurrences, 18 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2941 | `P.216_N.66_216/66/02_1` | SEKOLAH KEBANGSAAN LONG URUN | a | `P.216_N.66_216/66/02_1a` |
| 2943 | `P.216_N.66_216/66/02_1` | RH. SEPING LONG BALA | b | `P.216_N.66_216/66/02_1b` |
| 2944 | `P.216_N.66_216/66/02_1` | UMA SEPING LONG KOYAN | c | `P.216_N.66_216/66/02_1c` |
| 2945 | `P.216_N.66_216/66/02_1` | UMA KULIT DAERAH KECIL SUNGAI ASAP | d | `P.216_N.66_216/66/02_1d` |
| 2947 | `P.216_N.66_216/66/02_1` | UMA BELOR DAERAH KECIL SUNGAI ASAP | e | `P.216_N.66_216/66/02_1e` |
| 2948 | `P.216_N.66_216/66/02_1` | UMA DARO DAERAH KECIL SUNGAI ASAP | f | `P.216_N.66_216/66/02_1f` |
| 2949 | `P.216_N.66_216/66/02_1` | UMA KELEP DAERAH KECIL SUNGAI ASAP | g | `P.216_N.66_216/66/02_1g` |
| 2950 | `P.216_N.66_216/66/02_1` | UMA LAHANAN DAERAH KECIL SUNGAI ASAP | h | `P.216_N.66_216/66/02_1h` |
| 2951 | `P.216_N.66_216/66/02_1` | UMA NYAVING DAERAH KECIL SUNGAI ASAP | i | `P.216_N.66_216/66/02_1i` |
| 2952 | `P.216_N.66_216/66/02_1` | UMA BAWANG DAERAH KECIL SUNGAI ASAP | j | `P.216_N.66_216/66/02_1j` |
| 2953 | `P.216_N.66_216/66/02_1` | UMA BALUI LIKO DAERAH KECIL SUNGAI ASAP | k | `P.216_N.66_216/66/02_1k` |
| 2954 | `P.216_N.66_216/66/02_1` | UMA BAKAH DAERAH KECIL SUNGAI ASAP | l | `P.216_N.66_216/66/02_1l` |
| 2956 | `P.216_N.66_216/66/02_1` | UMA BALUI UKAP DAERAH KECIL SUNGAI ASAP | m | `P.216_N.66_216/66/02_1m` |
| 2957 | `P.216_N.66_216/66/02_1` | UMA LESONG DAERAH KECIL SUNGAI ASAP | n | `P.216_N.66_216/66/02_1n` |
| 2958 | `P.216_N.66_216/66/02_1` | UMA UKIT DAERAH KECIL SUNGAI ASAP | o | `P.216_N.66_216/66/02_1o` |
| 2959 | `P.216_N.66_216/66/02_1` | UMA JUMAN DAERAH KECIL SUNGAI ASAP | p | `P.216_N.66_216/66/02_1p` |
| 2960 | `P.216_N.66_216/66/02_1` | UMA PENAN TALUN DAERAH KECIL SUNGAI ASAP | q | `P.216_N.66_216/66/02_1q` |
| 2961 | `P.216_N.66_216/66/02_1` | UMA BADENG DAERAH KECIL SUNGAI ASAP | r | `P.216_N.66_216/66/02_1r` |

#### 391. `P.216_N.66_216/66/02_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2942 | `P.216_N.66_216/66/02_2` | SEKOLAH KEBANGSAAN LONG URUN | a | `P.216_N.66_216/66/02_2a` |
| 2946 | `P.216_N.66_216/66/02_2` | UMA KULIT DAERAH KECIL SUNGAI ASAP | b | `P.216_N.66_216/66/02_2b` |
| 2955 | `P.216_N.66_216/66/02_2` | UMA BAKAH DAERAH KECIL SUNGAI ASAP | c | `P.216_N.66_216/66/02_2c` |
| 2962 | `P.216_N.66_216/66/02_2` | UMA BADENG DAERAH KECIL SUNGAI ASAP | d | `P.216_N.66_216/66/02_2d` |

#### 392. `P.216_N.66_216/66/03_1` (12 occurrences, 12 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2964 | `P.216_N.66_216/66/03_1` | SEKOLAH KEBANGSAAN SG BUKIT BALAI | a | `P.216_N.66_216/66/03_1a` |
| 2966 | `P.216_N.66_216/66/03_1` | RH. JOSHUA BTG TUBAU | b | `P.216_N.66_216/66/03_1b` |
| 2967 | `P.216_N.66_216/66/03_1` | RH. ALIP, LONG UNAN TUBAU | c | `P.216_N.66_216/66/03_1c` |
| 2968 | `P.216_N.66_216/66/03_1` | RH. LASAH TUBA SG. DUSAN JELALONG | d | `P.216_N.66_216/66/03_1d` |
| 2969 | `P.216_N.66_216/66/03_1` | RH. PATRICK KEBING PO | e | `P.216_N.66_216/66/03_1e` |
| 2970 | `P.216_N.66_216/66/03_1` | RH. AUGUSTINE LATENG SAGING TUBAU | f | `P.216_N.66_216/66/03_1f` |
| 2971 | `P.216_N.66_216/66/03_1` | TADIKA KEMAS RH. BALRULLY | g | `P.216_N.66_216/66/03_1g` |
| 2973 | `P.216_N.66_216/66/03_1` | RH. LICHONG SG. SEBUTIN | h | `P.216_N.66_216/66/03_1h` |
| 2974 | `P.216_N.66_216/66/03_1` | RH. JERANDING ULU JELALONG | i | `P.216_N.66_216/66/03_1i` |
| 2975 | `P.216_N.66_216/66/03_1` | RH. JULAIHI NG. SAU JELALONG | j | `P.216_N.66_216/66/03_1j` |
| 2976 | `P.216_N.66_216/66/03_1` | SEKOLAH KEBANGSAAN KUALA KEBULU | k | `P.216_N.66_216/66/03_1k` |
| 2977 | `P.216_N.66_216/66/03_1` | RH. MALEK SG. KEBULU JELALONG | l | `P.216_N.66_216/66/03_1l` |

#### 393. `P.216_N.66_216/66/03_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2965 | `P.216_N.66_216/66/03_2` | SEKOLAH KEBANGSAAN SG BUKIT BALAI | a | `P.216_N.66_216/66/03_2a` |
| 2972 | `P.216_N.66_216/66/03_2` | TADIKA KEMAS RH. BALRULLY | b | `P.216_N.66_216/66/03_2b` |

#### 394. `P.217_N.67_217/67/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2980 | `P.217_N.67_217/67/01_1` | SEKOLAH KEBANGSAAN KUALA ANNAU | a | `P.217_N.67_217/67/01_1a` |
| 2981 | `P.217_N.67_217/67/01_1` | SEKOLAH KEBANGSAAN KUALA TATAU | b | `P.217_N.67_217/67/01_1b` |
| 2983 | `P.217_N.67_217/67/01_1` | SEKOLAH KEBANGSAAN KUALA SERUPAI | c | `P.217_N.67_217/67/01_1c` |
| 2984 | `P.217_N.67_217/67/01_1` | SEKOLAH KEBANGSAAN SG. SETULAN | d | `P.217_N.67_217/67/01_1d` |

#### 395. `P.217_N.67_217/67/03_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2990 | `P.217_N.67_217/67/03_1` | SEKOLAH KEBANGSAAN KAMPUNG JEPAK | a | `P.217_N.67_217/67/03_1a` |
| 2997 | `P.217_N.67_217/67/03_1` | RH. JOSHUA MANIT AK BUYU | b | `P.217_N.67_217/67/03_1b` |
| 2999 | `P.217_N.67_217/67/03_1` | SEKOLAH KEBANGSAAN SG SETIAM | c | `P.217_N.67_217/67/03_1c` |
| 3000 | `P.217_N.67_217/67/03_1` | SEKOLAH KEBANGSAAN SG SELAD | d | `P.217_N.67_217/67/03_1d` |

#### 396. `P.217_N.67_217/67/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 2991 | `P.217_N.67_217/67/03_2` | SEKOLAH KEBANGSAAN KAMPUNG JEPAK | a | `P.217_N.67_217/67/03_2a` |
| 2998 | `P.217_N.67_217/67/03_2` | RH. JOSHUA MANIT AK BUYU | b | `P.217_N.67_217/67/03_2b` |
| 3001 | `P.217_N.67_217/67/03_2` | SEKOLAH KEBANGSAAN SG SELAD | c | `P.217_N.67_217/67/03_2c` |

#### 397. `P.217_N.67_217/67/04_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3002 | `P.217_N.67_217/67/04_1` | RH. ATAN SG. SILAS | a | `P.217_N.67_217/67/04_1a` |
| 3003 | `P.217_N.67_217/67/04_1` | SEKOLAH KEBANGSAAN ULU SEGAN | b | `P.217_N.67_217/67/04_1b` |
| 3005 | `P.217_N.67_217/67/04_1` | TABIKA KEMAS KUALA SEGAN | c | `P.217_N.67_217/67/04_1c` |

#### 398. `P.217_N.68_217/68/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3019 | `P.217_N.68_217/68/01_1` | TADIKA PIMPIN | a | `P.217_N.68_217/68/01_1a` |
| 3023 | `P.217_N.68_217/68/01_1` | BANGUNAN PERSEKUTUAN PERKUMPULAN WANITA SARAWAK DAERAH BINTULU ( W.I
) | b | `P.217_N.68_217/68/01_1b` |
| 3024 | `P.217_N.68_217/68/01_1` | BANGUNAN PERSATUAN BULAN SABIT MERAH BINTULU | c | `P.217_N.68_217/68/01_1c` |

#### 399. `P.217_N.68_217/68/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3020 | `P.217_N.68_217/68/01_2` | TADIKA PIMPIN | a | `P.217_N.68_217/68/01_2a` |
| 3025 | `P.217_N.68_217/68/01_2` | BANGUNAN PERSATUAN BULAN SABIT MERAH BINTULU | b | `P.217_N.68_217/68/01_2b` |

#### 400. `P.217_N.68_217/68/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3043 | `P.217_N.68_217/68/04_1` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_1a` |
| 3056 | `P.217_N.68_217/68/04_1` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_1b` |

#### 401. `P.217_N.68_217/68/04_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3044 | `P.217_N.68_217/68/04_2` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_2a` |
| 3057 | `P.217_N.68_217/68/04_2` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_2b` |

#### 402. `P.217_N.68_217/68/04_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3045 | `P.217_N.68_217/68/04_3` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_3a` |
| 3058 | `P.217_N.68_217/68/04_3` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_3b` |

#### 403. `P.217_N.68_217/68/04_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3046 | `P.217_N.68_217/68/04_4` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_4a` |
| 3059 | `P.217_N.68_217/68/04_4` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_4b` |

#### 404. `P.217_N.68_217/68/04_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3047 | `P.217_N.68_217/68/04_5` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_5a` |
| 3060 | `P.217_N.68_217/68/04_5` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_5b` |

#### 405. `P.217_N.68_217/68/04_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3048 | `P.217_N.68_217/68/04_6` | SEKOLAH MENENGAH KEBANGSAAN BINTULU | a | `P.217_N.68_217/68/04_6a` |
| 3061 | `P.217_N.68_217/68/04_6` | SEKOLAH MENENGAH KEBANGSAAN BANDAR BINTULU | b | `P.217_N.68_217/68/04_6b` |

#### 406. `P.217_N.69_217/69/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3068 | `P.217_N.69_217/69/01_1` | RH. LAPI SG. BINAI | a | `P.217_N.69_217/69/01_1a` |
| 3069 | `P.217_N.69_217/69/01_1` | SEKOLAH KEBANGSAAN SEBAUH | b | `P.217_N.69_217/69/01_1b` |
| 3071 | `P.217_N.69_217/69/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG SAN | c | `P.217_N.69_217/69/01_1c` |

#### 407. `P.217_N.69_217/69/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3070 | `P.217_N.69_217/69/01_2` | SEKOLAH KEBANGSAAN SEBAUH | a | `P.217_N.69_217/69/01_2a` |
| 3072 | `P.217_N.69_217/69/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG SAN | b | `P.217_N.69_217/69/01_2b` |

#### 408. `P.217_N.69_217/69/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3075 | `P.217_N.69_217/69/02_1` | RH. RO TEBAN BARU | a | `P.217_N.69_217/69/02_1a` |
| 3076 | `P.217_N.69_217/69/02_1` | SEKOLAH KEBANGSAAN SG. SENGIAN LABANG | b | `P.217_N.69_217/69/02_1b` |
| 3077 | `P.217_N.69_217/69/02_1` | BALAI RAYA KAMPUNG LABANG | c | `P.217_N.69_217/69/02_1c` |
| 3079 | `P.217_N.69_217/69/02_1` | SEKOLAH KEBANGSAAN PANDAN | d | `P.217_N.69_217/69/02_1d` |

#### 409. `P.217_N.69_217/69/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3078 | `P.217_N.69_217/69/02_2` | BALAI RAYA KAMPUNG LABANG | a | `P.217_N.69_217/69/02_2a` |
| 3080 | `P.217_N.69_217/69/02_2` | SEKOLAH KEBANGSAAN PANDAN | b | `P.217_N.69_217/69/02_2b` |

#### 410. `P.217_N.69_217/69/05_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3089 | `P.217_N.69_217/69/05_1` | RH. JIMBAI SG. GERONG | a | `P.217_N.69_217/69/05_1a` |
| 3090 | `P.217_N.69_217/69/05_1` | RH. TABOR AK LASAH SG. SEBAUH | b | `P.217_N.69_217/69/05_1b` |
| 3091 | `P.217_N.69_217/69/05_1` | RH. RAYMOND PLEN SG. GELAM | c | `P.217_N.69_217/69/05_1c` |

#### 411. `P.217_N.69_217/69/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3093 | `P.217_N.69_217/69/06_1` | RH. ROBERT SG. SEBUNGAN | a | `P.217_N.69_217/69/06_1a` |
| 3095 | `P.217_N.69_217/69/06_1` | RH. AUGUSTINE LAMAU | b | `P.217_N.69_217/69/06_1b` |
| 3096 | `P.217_N.69_217/69/06_1` | RH. JENANG SG. GUSI KELABAT | c | `P.217_N.69_217/69/06_1c` |
| 3097 | `P.217_N.69_217/69/06_1` | RH. NOMPANG SG. SUJAN | d | `P.217_N.69_217/69/06_1d` |

#### 412. `P.217_N.69_217/69/07_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3098 | `P.217_N.69_217/69/07_1` | SEKOLAH KEBANGSAAN KUALA SIGU | a | `P.217_N.69_217/69/07_1a` |
| 3099 | `P.217_N.69_217/69/07_1` | RH. NICHOLAS SANDUM SG. SIGU | b | `P.217_N.69_217/69/07_1b` |
| 3100 | `P.217_N.69_217/69/07_1` | SEKOLAH KEBANGSAAN KUALA BINYO | c | `P.217_N.69_217/69/07_1c` |
| 3102 | `P.217_N.69_217/69/07_1` | RH. NYIPA AJONG SG. BINYO | d | `P.217_N.69_217/69/07_1d` |
| 3103 | `P.217_N.69_217/69/07_1` | SEKOLAH KEBANGSAAN BUKIT MAWANG | e | `P.217_N.69_217/69/07_1e` |
| 3104 | `P.217_N.69_217/69/07_1` | SEKOLAH KEBANGSAAN SG. GENAAN | f | `P.217_N.69_217/69/07_1f` |

#### 413. `P.217_N.70_217/70/01_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3112 | `P.217_N.70_217/70/01_1` | SEKOLAH KEBANGSAAN SG KEM BATU 18 | a | `P.217_N.70_217/70/01_1a` |
| 3117 | `P.217_N.70_217/70/01_1` | PRA SEKOLAH, SK. SG. TISANG | b | `P.217_N.70_217/70/01_1b` |
| 3123 | `P.217_N.70_217/70/01_1` | BLOK A & BLOK B, SEKOLAH JENIS KEBANGSAAN (CINA) SEBIEW CHINESE | c | `P.217_N.70_217/70/01_1c` |

#### 414. `P.217_N.70_217/70/01_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3113 | `P.217_N.70_217/70/01_2` | SEKOLAH KEBANGSAAN SG KEM BATU 18 | a | `P.217_N.70_217/70/01_2a` |
| 3118 | `P.217_N.70_217/70/01_2` | PRA SEKOLAH, SK. SG. TISANG | b | `P.217_N.70_217/70/01_2b` |
| 3124 | `P.217_N.70_217/70/01_2` | BLOK A & BLOK B, SEKOLAH JENIS KEBANGSAAN (CINA) SEBIEW CHINESE | c | `P.217_N.70_217/70/01_2c` |

#### 415. `P.217_N.70_217/70/01_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3114 | `P.217_N.70_217/70/01_3` | SEKOLAH KEBANGSAAN SG KEM BATU 18 | a | `P.217_N.70_217/70/01_3a` |
| 3119 | `P.217_N.70_217/70/01_3` | PRA SEKOLAH, SK. SG. TISANG | b | `P.217_N.70_217/70/01_3b` |

#### 416. `P.217_N.70_217/70/01_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3115 | `P.217_N.70_217/70/01_4` | SEKOLAH KEBANGSAAN SG KEM BATU 18 | a | `P.217_N.70_217/70/01_4a` |
| 3120 | `P.217_N.70_217/70/01_4` | PRA SEKOLAH, SK. SG. TISANG | b | `P.217_N.70_217/70/01_4b` |

#### 417. `P.217_N.70_217/70/01_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3116 | `P.217_N.70_217/70/01_5` | SEKOLAH KEBANGSAAN SG KEM BATU 18 | a | `P.217_N.70_217/70/01_5a` |
| 3121 | `P.217_N.70_217/70/01_5` | PRA SEKOLAH, SK. SG. TISANG | b | `P.217_N.70_217/70/01_5b` |

#### 418. `P.217_N.70_217/70/03_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3130 | `P.217_N.70_217/70/03_1` | BILIK MESYUARAT SURAU KPG.  KUALA SUAI | a | `P.217_N.70_217/70/03_1a` |
| 3131 | `P.217_N.70_217/70/03_1` | SEKOLAH KEBANGSAAN KPG. TEGAGENG | b | `P.217_N.70_217/70/03_1b` |
| 3133 | `P.217_N.70_217/70/03_1` | SEKOLAH KEBANGSAAN KPG. IRAN | c | `P.217_N.70_217/70/03_1c` |
| 3135 | `P.217_N.70_217/70/03_1` | SEKOLAH KEBANGSAAN SG. SEBATU | d | `P.217_N.70_217/70/03_1d` |
| 3137 | `P.217_N.70_217/70/03_1` | SEKOLAH KEBANGSAAN KUALA NYALAU | e | `P.217_N.70_217/70/03_1e` |

#### 419. `P.217_N.70_217/70/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3132 | `P.217_N.70_217/70/03_2` | SEKOLAH KEBANGSAAN KPG. TEGAGENG | a | `P.217_N.70_217/70/03_2a` |
| 3134 | `P.217_N.70_217/70/03_2` | SEKOLAH KEBANGSAAN KPG. IRAN | b | `P.217_N.70_217/70/03_2b` |
| 3136 | `P.217_N.70_217/70/03_2` | SEKOLAH KEBANGSAAN SG. SEBATU | c | `P.217_N.70_217/70/03_2c` |

#### 420. `P.218_N.71_218/71/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3159 | `P.218_N.71_218/71/02_1` | SEKOLAH KEBANGSAAN RANCANGAN SEPUPOK | a | `P.218_N.71_218/71/02_1a` |
| 3161 | `P.218_N.71_218/71/02_1` | DEWAN MASYARAKAT SEPUPOK | b | `P.218_N.71_218/71/02_1b` |
| 3162 | `P.218_N.71_218/71/02_1` | SEKOLAH KEBANGSAAN KPG. TARIKAN | c | `P.218_N.71_218/71/02_1c` |
| 3163 | `P.218_N.71_218/71/02_1` | SEKOLAH KEBANGSAAN KITA | d | `P.218_N.71_218/71/02_1d` |

#### 421. `P.218_N.71_218/71/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3160 | `P.218_N.71_218/71/02_2` | SEKOLAH KEBANGSAAN RANCANGAN SEPUPOK | a | `P.218_N.71_218/71/02_2a` |
| 3164 | `P.218_N.71_218/71/02_2` | SEKOLAH KEBANGSAAN KITA | b | `P.218_N.71_218/71/02_2b` |

#### 422. `P.218_N.71_218/71/03_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3165 | `P.218_N.71_218/71/03_1` | SEKOLAH KEBANGSAAN RUMAH MENTALI | a | `P.218_N.71_218/71/03_1a` |
| 3166 | `P.218_N.71_218/71/03_1` | SEKOLAH KEBANGSAAN SG. SAEH | b | `P.218_N.71_218/71/03_1b` |
| 3168 | `P.218_N.71_218/71/03_1` | BALAI RAYA RH. DAUD | c | `P.218_N.71_218/71/03_1c` |
| 3169 | `P.218_N.71_218/71/03_1` | SEKOLAH KEBANGSAAN TANJUNG BELIPAT | d | `P.218_N.71_218/71/03_1d` |
| 3170 | `P.218_N.71_218/71/03_1` | SEKOLAH KEBANGSAAN SG. TANGAP | e | `P.218_N.71_218/71/03_1e` |

#### 423. `P.218_N.71_218/71/04_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3171 | `P.218_N.71_218/71/04_1` | BALAI RAYA KUALA SIBUTI | a | `P.218_N.71_218/71/04_1a` |
| 3172 | `P.218_N.71_218/71/04_1` | SEKOLAH KEBANGSAAN BELIAU ISA | b | `P.218_N.71_218/71/04_1b` |

#### 424. `P.218_N.71_218/71/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3174 | `P.218_N.71_218/71/05_1` | SEKOLAH KEBANGSAAN BELIAU AHAD | a | `P.218_N.71_218/71/05_1a` |
| 3176 | `P.218_N.71_218/71/05_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SG LUMUT | b | `P.218_N.71_218/71/05_1b` |

#### 425. `P.218_N.71_218/71/06_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3177 | `P.218_N.71_218/71/06_1` | SEKOLAH KEBANGSAAN KPG. ANGUS | a | `P.218_N.71_218/71/06_1a` |
| 3178 | `P.218_N.71_218/71/06_1` | SEKOLAH KEBANGSAAN KPG. BUNGAI | b | `P.218_N.71_218/71/06_1b` |
| 3179 | `P.218_N.71_218/71/06_1` | DEWAN KAMPUNG MENJELIN | c | `P.218_N.71_218/71/06_1c` |
| 3180 | `P.218_N.71_218/71/06_1` | BALAI RAYA KAMPUNG SELANYAU | d | `P.218_N.71_218/71/06_1d` |

#### 426. `P.218_N.71_218/71/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3181 | `P.218_N.71_218/71/07_1` | SEKOLAH KEBANGSAAN KPG. BULAU | a | `P.218_N.71_218/71/07_1a` |
| 3182 | `P.218_N.71_218/71/07_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA SIBUTI | b | `P.218_N.71_218/71/07_1b` |

#### 427. `P.218_N.71_218/71/08_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3185 | `P.218_N.71_218/71/08_1` | SEKOLAH KEBANGSAAN RUMAH TINGGI | a | `P.218_N.71_218/71/08_1a` |
| 3186 | `P.218_N.71_218/71/08_1` | SEKOLAH KEBANGSAAN SG. BAKAS | b | `P.218_N.71_218/71/08_1b` |
| 3187 | `P.218_N.71_218/71/08_1` | SEKOLAH KEBANGSAAN RUMAH ESSAU | c | `P.218_N.71_218/71/08_1c` |
| 3188 | `P.218_N.71_218/71/08_1` | SEKOLAH KEBANGSAAN KELAPA SAWIT NO. 4 | d | `P.218_N.71_218/71/08_1d` |

#### 428. `P.218_N.72_218/72/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3194 | `P.218_N.72_218/72/00_1` | DEWAN MAKAN PASUKAN,KEM SRI MIRI | a | `P.218_N.72_218/72/00_1a` |
| 3196 | `P.218_N.72_218/72/00_1` | RUANG B,BANGUNAN RUAI KENYALANG IPD MIRI | b | `P.218_N.72_218/72/00_1b` |

#### 429. `P.218_N.72_218/72/01_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3197 | `P.218_N.72_218/72/01_1` | SEKOLAH MENENGAH KEBANGSAAN LUAR BANDAR MIRI | a | `P.218_N.72_218/72/01_1a` |
| 3199 | `P.218_N.72_218/72/01_1` | RH. DUAT SG. KELINTANG | b | `P.218_N.72_218/72/01_1b` |
| 3200 | `P.218_N.72_218/72/01_1` | RH. JAMES BIRI | c | `P.218_N.72_218/72/01_1c` |
| 3201 | `P.218_N.72_218/72/01_1` | DEWAN PUSAT PEMBANGUNAN KEMAHIRAN SARAWAK (PPKS) | d | `P.218_N.72_218/72/01_1d` |
| 3202 | `P.218_N.72_218/72/01_1` | SEKOLAH KEBANGSAAN KELAPA SAWIT NO. 2 | e | `P.218_N.72_218/72/01_1e` |
| 3204 | `P.218_N.72_218/72/01_1` | RH. LAGAN | f | `P.218_N.72_218/72/01_1f` |

#### 430. `P.218_N.72_218/72/01_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3198 | `P.218_N.72_218/72/01_2` | SEKOLAH MENENGAH KEBANGSAAN LUAR BANDAR MIRI | a | `P.218_N.72_218/72/01_2a` |
| 3203 | `P.218_N.72_218/72/01_2` | SEKOLAH KEBANGSAAN KELAPA SAWIT NO. 2 | b | `P.218_N.72_218/72/01_2b` |

#### 431. `P.218_N.72_218/72/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3205 | `P.218_N.72_218/72/02_1` | SEKOLAH KEBANGSAAN KELURU TENGAH | a | `P.218_N.72_218/72/02_1a` |
| 3208 | `P.218_N.72_218/72/02_1` | SEKOLAH KEBANGSAAN TAWAKAL SATAP | b | `P.218_N.72_218/72/02_1b` |

#### 432. `P.218_N.72_218/72/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3206 | `P.218_N.72_218/72/02_2` | SEKOLAH KEBANGSAAN KELURU TENGAH | a | `P.218_N.72_218/72/02_2a` |
| 3209 | `P.218_N.72_218/72/02_2` | SEKOLAH KEBANGSAAN TAWAKAL SATAP | b | `P.218_N.72_218/72/02_2b` |

#### 433. `P.218_N.72_218/72/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3210 | `P.218_N.72_218/72/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA BAKAM | a | `P.218_N.72_218/72/03_1a` |
| 3211 | `P.218_N.72_218/72/03_1` | SEKOLAH KEBANGSAAN KG. BAKAM | b | `P.218_N.72_218/72/03_1b` |
| 3215 | `P.218_N.72_218/72/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) TUKAU | c | `P.218_N.72_218/72/03_1c` |

#### 434. `P.218_N.72_218/72/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3223 | `P.218_N.72_218/72/05_1` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_1a` |
| 3231 | `P.218_N.72_218/72/05_1` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_1b` |

#### 435. `P.218_N.72_218/72/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3224 | `P.218_N.72_218/72/05_2` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_2a` |
| 3232 | `P.218_N.72_218/72/05_2` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_2b` |

#### 436. `P.218_N.72_218/72/05_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3225 | `P.218_N.72_218/72/05_3` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_3a` |
| 3233 | `P.218_N.72_218/72/05_3` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_3b` |

#### 437. `P.218_N.72_218/72/05_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3226 | `P.218_N.72_218/72/05_4` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_4a` |
| 3234 | `P.218_N.72_218/72/05_4` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_4b` |

#### 438. `P.218_N.72_218/72/05_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3227 | `P.218_N.72_218/72/05_5` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_5a` |
| 3235 | `P.218_N.72_218/72/05_5` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_5b` |

#### 439. `P.218_N.72_218/72/05_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3228 | `P.218_N.72_218/72/05_6` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_6a` |
| 3236 | `P.218_N.72_218/72/05_6` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_6b` |

#### 440. `P.218_N.72_218/72/05_7` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3229 | `P.218_N.72_218/72/05_7` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_7a` |
| 3237 | `P.218_N.72_218/72/05_7` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_7b` |

#### 441. `P.218_N.72_218/72/05_8` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3230 | `P.218_N.72_218/72/05_8` | SEKOLAH MENENGAH KEBANGSAAN RIAM | a | `P.218_N.72_218/72/05_8a` |
| 3238 | `P.218_N.72_218/72/05_8` | SEKOLAH MENENGAH KEBANGSAAN TAMAN TUNKU | b | `P.218_N.72_218/72/05_8b` |

#### 442. `P.219_N.73_219/73/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3246 | `P.219_N.73_219/73/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_1a` |
| 3252 | `P.219_N.73_219/73/02_1` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_1b` |

#### 443. `P.219_N.73_219/73/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3247 | `P.219_N.73_219/73/02_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_2a` |
| 3253 | `P.219_N.73_219/73/02_2` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_2b` |

#### 444. `P.219_N.73_219/73/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3248 | `P.219_N.73_219/73/02_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_3a` |
| 3254 | `P.219_N.73_219/73/02_3` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_3b` |

#### 445. `P.219_N.73_219/73/02_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3249 | `P.219_N.73_219/73/02_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_4a` |
| 3255 | `P.219_N.73_219/73/02_4` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_4b` |

#### 446. `P.219_N.73_219/73/02_5` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3250 | `P.219_N.73_219/73/02_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_5a` |
| 3256 | `P.219_N.73_219/73/02_5` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_5b` |

#### 447. `P.219_N.73_219/73/02_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3251 | `P.219_N.73_219/73/02_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA MIRI | a | `P.219_N.73_219/73/02_6a` |
| 3257 | `P.219_N.73_219/73/02_6` | SEKOLAH KEBANGSAAN ST JOSEPH | b | `P.219_N.73_219/73/02_6b` |

#### 448. `P.219_N.73_219/73/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3265 | `P.219_N.73_219/73/05_1` | SEKOLAH KEBANGSAAN SOUTH | a | `P.219_N.73_219/73/05_1a` |
| 3269 | `P.219_N.73_219/73/05_1` | SEKOLAH KEBANGSAAN PULAU MELAYU | b | `P.219_N.73_219/73/05_1b` |

#### 449. `P.219_N.73_219/73/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3266 | `P.219_N.73_219/73/05_2` | SEKOLAH KEBANGSAAN SOUTH | a | `P.219_N.73_219/73/05_2a` |
| 3270 | `P.219_N.73_219/73/05_2` | SEKOLAH KEBANGSAAN PULAU MELAYU | b | `P.219_N.73_219/73/05_2b` |

#### 450. `P.219_N.73_219/73/07_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3274 | `P.219_N.73_219/73/07_1` | SEKOLAH JENIS KEBANGSAAN  (CINA) CHUNG HUA LUTONG | a | `P.219_N.73_219/73/07_1a` |
| 3278 | `P.219_N.73_219/73/07_1` | SEKOLAH MENENGAH KEBANGSAAN LUTONG | b | `P.219_N.73_219/73/07_1b` |
| 3282 | `P.219_N.73_219/73/07_1` | SEKOLAH KEBANGSAAN LUTONG | c | `P.219_N.73_219/73/07_1c` |

#### 451. `P.219_N.73_219/73/07_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3275 | `P.219_N.73_219/73/07_2` | SEKOLAH JENIS KEBANGSAAN  (CINA) CHUNG HUA LUTONG | a | `P.219_N.73_219/73/07_2a` |
| 3279 | `P.219_N.73_219/73/07_2` | SEKOLAH MENENGAH KEBANGSAAN LUTONG | b | `P.219_N.73_219/73/07_2b` |
| 3283 | `P.219_N.73_219/73/07_2` | SEKOLAH KEBANGSAAN LUTONG | c | `P.219_N.73_219/73/07_2c` |

#### 452. `P.219_N.73_219/73/07_3` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3276 | `P.219_N.73_219/73/07_3` | SEKOLAH JENIS KEBANGSAAN  (CINA) CHUNG HUA LUTONG | a | `P.219_N.73_219/73/07_3a` |
| 3280 | `P.219_N.73_219/73/07_3` | SEKOLAH MENENGAH KEBANGSAAN LUTONG | b | `P.219_N.73_219/73/07_3b` |
| 3284 | `P.219_N.73_219/73/07_3` | SEKOLAH KEBANGSAAN LUTONG | c | `P.219_N.73_219/73/07_3c` |

#### 453. `P.219_N.73_219/73/07_4` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3277 | `P.219_N.73_219/73/07_4` | SEKOLAH JENIS KEBANGSAAN  (CINA) CHUNG HUA LUTONG | a | `P.219_N.73_219/73/07_4a` |
| 3281 | `P.219_N.73_219/73/07_4` | SEKOLAH MENENGAH KEBANGSAAN LUTONG | b | `P.219_N.73_219/73/07_4b` |
| 3285 | `P.219_N.73_219/73/07_4` | SEKOLAH KEBANGSAAN LUTONG | c | `P.219_N.73_219/73/07_4c` |

#### 454. `P.219_N.74_219/74/01_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3299 | `P.219_N.74_219/74/01_1` | TADIKA MIRI CHINESE | a | `P.219_N.74_219/74/01_1a` |
| 3304 | `P.219_N.74_219/74/01_1` | SEKOLAH RENDAH SRI MAWAR | b | `P.219_N.74_219/74/01_1b` |
| 3311 | `P.219_N.74_219/74/01_1` | SEKOLAH KEBANGSAAN ANCHI | c | `P.219_N.74_219/74/01_1c` |
| 3318 | `P.219_N.74_219/74/01_1` | TADIKA PUJUT MIRI | d | `P.219_N.74_219/74/01_1d` |
| 3322 | `P.219_N.74_219/74/01_1` | SEKOLAH KEBANGSAAN PUJUT CORNER | e | `P.219_N.74_219/74/01_1e` |
| 3327 | `P.219_N.74_219/74/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | f | `P.219_N.74_219/74/01_1f` |
| 3334 | `P.219_N.74_219/74/01_1` | SEKOLAH MENENGAH KEBANGSAAN DATO PERMAISURI | g | `P.219_N.74_219/74/01_1g` |

#### 455. `P.219_N.74_219/74/01_2` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3300 | `P.219_N.74_219/74/01_2` | TADIKA MIRI CHINESE | a | `P.219_N.74_219/74/01_2a` |
| 3305 | `P.219_N.74_219/74/01_2` | SEKOLAH RENDAH SRI MAWAR | b | `P.219_N.74_219/74/01_2b` |
| 3312 | `P.219_N.74_219/74/01_2` | SEKOLAH KEBANGSAAN ANCHI | c | `P.219_N.74_219/74/01_2c` |
| 3319 | `P.219_N.74_219/74/01_2` | TADIKA PUJUT MIRI | d | `P.219_N.74_219/74/01_2d` |
| 3323 | `P.219_N.74_219/74/01_2` | SEKOLAH KEBANGSAAN PUJUT CORNER | e | `P.219_N.74_219/74/01_2e` |
| 3328 | `P.219_N.74_219/74/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | f | `P.219_N.74_219/74/01_2f` |
| 3335 | `P.219_N.74_219/74/01_2` | SEKOLAH MENENGAH KEBANGSAAN DATO PERMAISURI | g | `P.219_N.74_219/74/01_2g` |

#### 456. `P.219_N.74_219/74/01_3` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3301 | `P.219_N.74_219/74/01_3` | TADIKA MIRI CHINESE | a | `P.219_N.74_219/74/01_3a` |
| 3306 | `P.219_N.74_219/74/01_3` | SEKOLAH RENDAH SRI MAWAR | b | `P.219_N.74_219/74/01_3b` |
| 3313 | `P.219_N.74_219/74/01_3` | SEKOLAH KEBANGSAAN ANCHI | c | `P.219_N.74_219/74/01_3c` |
| 3320 | `P.219_N.74_219/74/01_3` | TADIKA PUJUT MIRI | d | `P.219_N.74_219/74/01_3d` |
| 3324 | `P.219_N.74_219/74/01_3` | SEKOLAH KEBANGSAAN PUJUT CORNER | e | `P.219_N.74_219/74/01_3e` |
| 3329 | `P.219_N.74_219/74/01_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | f | `P.219_N.74_219/74/01_3f` |
| 3336 | `P.219_N.74_219/74/01_3` | SEKOLAH MENENGAH KEBANGSAAN DATO PERMAISURI | g | `P.219_N.74_219/74/01_3g` |

#### 457. `P.219_N.74_219/74/01_4` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3302 | `P.219_N.74_219/74/01_4` | TADIKA MIRI CHINESE | a | `P.219_N.74_219/74/01_4a` |
| 3307 | `P.219_N.74_219/74/01_4` | SEKOLAH RENDAH SRI MAWAR | b | `P.219_N.74_219/74/01_4b` |
| 3314 | `P.219_N.74_219/74/01_4` | SEKOLAH KEBANGSAAN ANCHI | c | `P.219_N.74_219/74/01_4c` |
| 3321 | `P.219_N.74_219/74/01_4` | TADIKA PUJUT MIRI | d | `P.219_N.74_219/74/01_4d` |
| 3325 | `P.219_N.74_219/74/01_4` | SEKOLAH KEBANGSAAN PUJUT CORNER | e | `P.219_N.74_219/74/01_4e` |
| 3330 | `P.219_N.74_219/74/01_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | f | `P.219_N.74_219/74/01_4f` |
| 3337 | `P.219_N.74_219/74/01_4` | SEKOLAH MENENGAH KEBANGSAAN DATO PERMAISURI | g | `P.219_N.74_219/74/01_4g` |

#### 458. `P.219_N.74_219/74/01_5` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3303 | `P.219_N.74_219/74/01_5` | TADIKA MIRI CHINESE | a | `P.219_N.74_219/74/01_5a` |
| 3308 | `P.219_N.74_219/74/01_5` | SEKOLAH RENDAH SRI MAWAR | b | `P.219_N.74_219/74/01_5b` |
| 3315 | `P.219_N.74_219/74/01_5` | SEKOLAH KEBANGSAAN ANCHI | c | `P.219_N.74_219/74/01_5c` |
| 3326 | `P.219_N.74_219/74/01_5` | SEKOLAH KEBANGSAAN PUJUT CORNER | d | `P.219_N.74_219/74/01_5d` |
| 3331 | `P.219_N.74_219/74/01_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | e | `P.219_N.74_219/74/01_5e` |
| 3338 | `P.219_N.74_219/74/01_5` | SEKOLAH MENENGAH KEBANGSAAN DATO PERMAISURI | f | `P.219_N.74_219/74/01_5f` |

#### 459. `P.219_N.74_219/74/01_6` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3309 | `P.219_N.74_219/74/01_6` | SEKOLAH RENDAH SRI MAWAR | a | `P.219_N.74_219/74/01_6a` |
| 3316 | `P.219_N.74_219/74/01_6` | SEKOLAH KEBANGSAAN ANCHI | b | `P.219_N.74_219/74/01_6b` |
| 3332 | `P.219_N.74_219/74/01_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | c | `P.219_N.74_219/74/01_6c` |

#### 460. `P.219_N.74_219/74/01_7` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3310 | `P.219_N.74_219/74/01_7` | SEKOLAH RENDAH SRI MAWAR | a | `P.219_N.74_219/74/01_7a` |
| 3317 | `P.219_N.74_219/74/01_7` | SEKOLAH KEBANGSAAN ANCHI | b | `P.219_N.74_219/74/01_7b` |
| 3333 | `P.219_N.74_219/74/01_7` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA PUJUT | c | `P.219_N.74_219/74/01_7c` |

#### 461. `P.219_N.74_219/74/02_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3339 | `P.219_N.74_219/74/02_1` | KOLEJ VOKASIONAL MIRI | a | `P.219_N.74_219/74/02_1a` |
| 3344 | `P.219_N.74_219/74/02_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | b | `P.219_N.74_219/74/02_1b` |
| 3350 | `P.219_N.74_219/74/02_1` | SEKOLAH KEBANGSAAN MIRI | c | `P.219_N.74_219/74/02_1c` |
| 3354 | `P.219_N.74_219/74/02_1` | SEKOLAH MENENGAH PEI MIN | d | `P.219_N.74_219/74/02_1d` |

#### 462. `P.219_N.74_219/74/02_2` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3340 | `P.219_N.74_219/74/02_2` | KOLEJ VOKASIONAL MIRI | a | `P.219_N.74_219/74/02_2a` |
| 3345 | `P.219_N.74_219/74/02_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | b | `P.219_N.74_219/74/02_2b` |
| 3351 | `P.219_N.74_219/74/02_2` | SEKOLAH KEBANGSAAN MIRI | c | `P.219_N.74_219/74/02_2c` |
| 3355 | `P.219_N.74_219/74/02_2` | SEKOLAH MENENGAH PEI MIN | d | `P.219_N.74_219/74/02_2d` |

#### 463. `P.219_N.74_219/74/02_3` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3341 | `P.219_N.74_219/74/02_3` | KOLEJ VOKASIONAL MIRI | a | `P.219_N.74_219/74/02_3a` |
| 3346 | `P.219_N.74_219/74/02_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | b | `P.219_N.74_219/74/02_3b` |
| 3352 | `P.219_N.74_219/74/02_3` | SEKOLAH KEBANGSAAN MIRI | c | `P.219_N.74_219/74/02_3c` |
| 3356 | `P.219_N.74_219/74/02_3` | SEKOLAH MENENGAH PEI MIN | d | `P.219_N.74_219/74/02_3d` |

#### 464. `P.219_N.74_219/74/02_4` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3342 | `P.219_N.74_219/74/02_4` | KOLEJ VOKASIONAL MIRI | a | `P.219_N.74_219/74/02_4a` |
| 3347 | `P.219_N.74_219/74/02_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | b | `P.219_N.74_219/74/02_4b` |
| 3353 | `P.219_N.74_219/74/02_4` | SEKOLAH KEBANGSAAN MIRI | c | `P.219_N.74_219/74/02_4c` |
| 3357 | `P.219_N.74_219/74/02_4` | SEKOLAH MENENGAH PEI MIN | d | `P.219_N.74_219/74/02_4d` |

#### 465. `P.219_N.74_219/74/02_5` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3343 | `P.219_N.74_219/74/02_5` | KOLEJ VOKASIONAL MIRI | a | `P.219_N.74_219/74/02_5a` |
| 3348 | `P.219_N.74_219/74/02_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | b | `P.219_N.74_219/74/02_5b` |
| 3358 | `P.219_N.74_219/74/02_5` | SEKOLAH MENENGAH PEI MIN | c | `P.219_N.74_219/74/02_5c` |

#### 466. `P.219_N.74_219/74/02_6` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3349 | `P.219_N.74_219/74/02_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA KROKOP | a | `P.219_N.74_219/74/02_6a` |
| 3359 | `P.219_N.74_219/74/02_6` | SEKOLAH MENENGAH PEI MIN | b | `P.219_N.74_219/74/02_6b` |

#### 467. `P.219_N.75_219/75/01_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3363 | `P.219_N.75_219/75/01_1` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_1a` |
| 3370 | `P.219_N.75_219/75/01_1` | SEKOLAH KEBANGSAAN SENADIN | b | `P.219_N.75_219/75/01_1b` |
| 3374 | `P.219_N.75_219/75/01_1` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | c | `P.219_N.75_219/75/01_1c` |
| 3381 | `P.219_N.75_219/75/01_1` | DEWAN KG PANGKALAN LUTONG | d | `P.219_N.75_219/75/01_1d` |
| 3384 | `P.219_N.75_219/75/01_1` | SEKOLAH KEBANGSAAN MERBAU | e | `P.219_N.75_219/75/01_1e` |
| 3397 | `P.219_N.75_219/75/01_1` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | f | `P.219_N.75_219/75/01_1f` |
| 3404 | `P.219_N.75_219/75/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | g | `P.219_N.75_219/75/01_1g` |
| 3410 | `P.219_N.75_219/75/01_1` | SEKOLAH KEBANGSAAN KUALA BARAM | h | `P.219_N.75_219/75/01_1h` |
| 3411 | `P.219_N.75_219/75/01_1` | SEKOLAH KEBANGSAAN KUALA BARAM II | i | `P.219_N.75_219/75/01_1i` |

#### 468. `P.219_N.75_219/75/01_2` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3364 | `P.219_N.75_219/75/01_2` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_2a` |
| 3371 | `P.219_N.75_219/75/01_2` | SEKOLAH KEBANGSAAN SENADIN | b | `P.219_N.75_219/75/01_2b` |
| 3375 | `P.219_N.75_219/75/01_2` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | c | `P.219_N.75_219/75/01_2c` |
| 3382 | `P.219_N.75_219/75/01_2` | DEWAN KG PANGKALAN LUTONG | d | `P.219_N.75_219/75/01_2d` |
| 3385 | `P.219_N.75_219/75/01_2` | SEKOLAH KEBANGSAAN MERBAU | e | `P.219_N.75_219/75/01_2e` |
| 3398 | `P.219_N.75_219/75/01_2` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | f | `P.219_N.75_219/75/01_2f` |
| 3405 | `P.219_N.75_219/75/01_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | g | `P.219_N.75_219/75/01_2g` |
| 3412 | `P.219_N.75_219/75/01_2` | SEKOLAH KEBANGSAAN KUALA BARAM II | h | `P.219_N.75_219/75/01_2h` |

#### 469. `P.219_N.75_219/75/01_3` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3365 | `P.219_N.75_219/75/01_3` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_3a` |
| 3372 | `P.219_N.75_219/75/01_3` | SEKOLAH KEBANGSAAN SENADIN | b | `P.219_N.75_219/75/01_3b` |
| 3376 | `P.219_N.75_219/75/01_3` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | c | `P.219_N.75_219/75/01_3c` |
| 3383 | `P.219_N.75_219/75/01_3` | DEWAN KG PANGKALAN LUTONG | d | `P.219_N.75_219/75/01_3d` |
| 3386 | `P.219_N.75_219/75/01_3` | SEKOLAH KEBANGSAAN MERBAU | e | `P.219_N.75_219/75/01_3e` |
| 3399 | `P.219_N.75_219/75/01_3` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | f | `P.219_N.75_219/75/01_3f` |
| 3406 | `P.219_N.75_219/75/01_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | g | `P.219_N.75_219/75/01_3g` |
| 3413 | `P.219_N.75_219/75/01_3` | SEKOLAH KEBANGSAAN KUALA BARAM II | h | `P.219_N.75_219/75/01_3h` |

#### 470. `P.219_N.75_219/75/01_4` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3366 | `P.219_N.75_219/75/01_4` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_4a` |
| 3373 | `P.219_N.75_219/75/01_4` | SEKOLAH KEBANGSAAN SENADIN | b | `P.219_N.75_219/75/01_4b` |
| 3377 | `P.219_N.75_219/75/01_4` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | c | `P.219_N.75_219/75/01_4c` |
| 3387 | `P.219_N.75_219/75/01_4` | SEKOLAH KEBANGSAAN MERBAU | d | `P.219_N.75_219/75/01_4d` |
| 3400 | `P.219_N.75_219/75/01_4` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | e | `P.219_N.75_219/75/01_4e` |
| 3407 | `P.219_N.75_219/75/01_4` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | f | `P.219_N.75_219/75/01_4f` |
| 3414 | `P.219_N.75_219/75/01_4` | SEKOLAH KEBANGSAAN KUALA BARAM II | g | `P.219_N.75_219/75/01_4g` |

#### 471. `P.219_N.75_219/75/01_5` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3367 | `P.219_N.75_219/75/01_5` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_5a` |
| 3378 | `P.219_N.75_219/75/01_5` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | b | `P.219_N.75_219/75/01_5b` |
| 3388 | `P.219_N.75_219/75/01_5` | SEKOLAH KEBANGSAAN MERBAU | c | `P.219_N.75_219/75/01_5c` |
| 3401 | `P.219_N.75_219/75/01_5` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | d | `P.219_N.75_219/75/01_5d` |
| 3408 | `P.219_N.75_219/75/01_5` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | e | `P.219_N.75_219/75/01_5e` |

#### 472. `P.219_N.75_219/75/01_6` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3368 | `P.219_N.75_219/75/01_6` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_6a` |
| 3379 | `P.219_N.75_219/75/01_6` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | b | `P.219_N.75_219/75/01_6b` |
| 3389 | `P.219_N.75_219/75/01_6` | SEKOLAH KEBANGSAAN MERBAU | c | `P.219_N.75_219/75/01_6c` |
| 3402 | `P.219_N.75_219/75/01_6` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | d | `P.219_N.75_219/75/01_6d` |
| 3409 | `P.219_N.75_219/75/01_6` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA TUDAN | e | `P.219_N.75_219/75/01_6e` |

#### 473. `P.219_N.75_219/75/01_7` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3369 | `P.219_N.75_219/75/01_7` | SEKOLAH KEBANGSAAN TUDAN | a | `P.219_N.75_219/75/01_7a` |
| 3380 | `P.219_N.75_219/75/01_7` | SEKOLAH MENENGAH KEBANGSAAN PUJUT | b | `P.219_N.75_219/75/01_7b` |
| 3390 | `P.219_N.75_219/75/01_7` | SEKOLAH KEBANGSAAN MERBAU | c | `P.219_N.75_219/75/01_7c` |
| 3403 | `P.219_N.75_219/75/01_7` | SEKOLAH MENENGAH KEBANGSAAN MERBAU | d | `P.219_N.75_219/75/01_7d` |

#### 474. `P.219_N.75_219/75/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3415 | `P.219_N.75_219/75/02_1` | SEKOLAH KEBANGSAAN AGAMA | a | `P.219_N.75_219/75/02_1a` |
| 3420 | `P.219_N.75_219/75/02_1` | SEKOLAH MENENGAH KEBANGSAAN BARU | b | `P.219_N.75_219/75/02_1b` |
| 3424 | `P.219_N.75_219/75/02_1` | RH DOK SG. TENIKU MIRI | c | `P.219_N.75_219/75/02_1c` |

#### 475. `P.219_N.75_219/75/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3416 | `P.219_N.75_219/75/02_2` | SEKOLAH KEBANGSAAN AGAMA | a | `P.219_N.75_219/75/02_2a` |
| 3421 | `P.219_N.75_219/75/02_2` | SEKOLAH MENENGAH KEBANGSAAN BARU | b | `P.219_N.75_219/75/02_2b` |

#### 476. `P.219_N.75_219/75/02_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3417 | `P.219_N.75_219/75/02_3` | SEKOLAH KEBANGSAAN AGAMA | a | `P.219_N.75_219/75/02_3a` |
| 3422 | `P.219_N.75_219/75/02_3` | SEKOLAH MENENGAH KEBANGSAAN BARU | b | `P.219_N.75_219/75/02_3b` |

#### 477. `P.219_N.75_219/75/02_4` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3418 | `P.219_N.75_219/75/02_4` | SEKOLAH KEBANGSAAN AGAMA | a | `P.219_N.75_219/75/02_4a` |
| 3423 | `P.219_N.75_219/75/02_4` | SEKOLAH MENENGAH KEBANGSAAN BARU | b | `P.219_N.75_219/75/02_4b` |

#### 478. `P.220_N.76_220/76/01_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3440 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN LEPONG AJAI | a | `P.220_N.76_220/76/01_1a` |
| 3441 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN SG. LIAM | b | `P.220_N.76_220/76/01_1b` |
| 3442 | `P.220_N.76_220/76/01_1` | RH. WILSON AK JUNA, SG. NAKAT, BUKIT SONG, BAKONG | c | `P.220_N.76_220/76/01_1c` |
| 3443 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN SG. ENTULANG | d | `P.220_N.76_220/76/01_1d` |
| 3444 | `P.220_N.76_220/76/01_1` | RH. ENGKAS AK ENTARI, SG. MALLANG SANIN, BAKONG | e | `P.220_N.76_220/76/01_1e` |
| 3445 | `P.220_N.76_220/76/01_1` | RH. JOSHUA AK DUNGKONG, SG. MALLANG ULU, BAKONG | f | `P.220_N.76_220/76/01_1f` |
| 3446 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN SG. BURI | g | `P.220_N.76_220/76/01_1g` |
| 3447 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN SG ARANG | h | `P.220_N.76_220/76/01_1h` |
| 3448 | `P.220_N.76_220/76/01_1` | SEKOLAH KEBANGSAAN SG. BIAR | i | `P.220_N.76_220/76/01_1i` |

#### 479. `P.220_N.76_220/76/02_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3449 | `P.220_N.76_220/76/02_1` | RH. BUDIN ASSAM PAYA | a | `P.220_N.76_220/76/02_1a` |
| 3450 | `P.220_N.76_220/76/02_1` | SEKOLAH KEBANGSAAN PENGELAYAN | b | `P.220_N.76_220/76/02_1b` |
| 3451 | `P.220_N.76_220/76/02_1` | RH. NGELINGKONG TERAJA | c | `P.220_N.76_220/76/02_1c` |
| 3452 | `P.220_N.76_220/76/02_1` | RH. CHABOP LBK. AMAM | d | `P.220_N.76_220/76/02_1d` |
| 3453 | `P.220_N.76_220/76/02_1` | SEKOLAH KEBANGSAAN BENAWA | e | `P.220_N.76_220/76/02_1e` |
| 3454 | `P.220_N.76_220/76/02_1` | SEKOLAH KEBANGSAAN DATO SHARIF HAMID | f | `P.220_N.76_220/76/02_1f` |
| 3457 | `P.220_N.76_220/76/02_1` | KAMPUNG LONG MARO TINJAR | g | `P.220_N.76_220/76/02_1g` |

#### 480. `P.220_N.76_220/76/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3458 | `P.220_N.76_220/76/03_1` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA | a | `P.220_N.76_220/76/03_1a` |
| 3461 | `P.220_N.76_220/76/03_1` | SEKOLAH JENIS KEBANGSAAN SUNGAI JAONG MARUDI | b | `P.220_N.76_220/76/03_1b` |
| 3465 | `P.220_N.76_220/76/03_1` | SEKOLAH KEBANGSAAN GOOD SHEPHERD | c | `P.220_N.76_220/76/03_1c` |

#### 481. `P.220_N.76_220/76/03_2` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3459 | `P.220_N.76_220/76/03_2` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA | a | `P.220_N.76_220/76/03_2a` |
| 3462 | `P.220_N.76_220/76/03_2` | SEKOLAH JENIS KEBANGSAAN SUNGAI JAONG MARUDI | b | `P.220_N.76_220/76/03_2b` |
| 3466 | `P.220_N.76_220/76/03_2` | SEKOLAH KEBANGSAAN GOOD SHEPHERD | c | `P.220_N.76_220/76/03_2c` |

#### 482. `P.220_N.76_220/76/03_3` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3460 | `P.220_N.76_220/76/03_3` | SEKOLAH JENIS KEBANGSAAN (CINA) CHUNG HUA | a | `P.220_N.76_220/76/03_3a` |
| 3463 | `P.220_N.76_220/76/03_3` | SEKOLAH JENIS KEBANGSAAN SUNGAI JAONG MARUDI | b | `P.220_N.76_220/76/03_3b` |

#### 483. `P.220_N.76_220/76/04_1` (16 occurrences, 16 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3467 | `P.220_N.76_220/76/04_1` | RH. SALWAN SG. LAI BAKONG | a | `P.220_N.76_220/76/04_1a` |
| 3468 | `P.220_N.76_220/76/04_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HUA KWONG | b | `P.220_N.76_220/76/04_1b` |
| 3470 | `P.220_N.76_220/76/04_1` | DEWAN SERBAGUNA KPG. MELAYU, BELURU | c | `P.220_N.76_220/76/04_1c` |
| 3471 | `P.220_N.76_220/76/04_1` | RH. MORGAN, SG. NIPA, BAKONG | d | `P.220_N.76_220/76/04_1d` |
| 3472 | `P.220_N.76_220/76/04_1` | RH. LINGGIE, SG. URONG BAKONG | e | `P.220_N.76_220/76/04_1e` |
| 3473 | `P.220_N.76_220/76/04_1` | SEKOLAH KEBANGSAAN SG. SELEPIN | f | `P.220_N.76_220/76/04_1f` |
| 3474 | `P.220_N.76_220/76/04_1` | RH. SABA AK BAGUS, SELEPIN ATAS | g | `P.220_N.76_220/76/04_1g` |
| 3475 | `P.220_N.76_220/76/04_1` | SEKOLAH KEBANGSAAN SG. LAONG | h | `P.220_N.76_220/76/04_1h` |
| 3476 | `P.220_N.76_220/76/04_1` | RH. TAN SG. TEMAM | i | `P.220_N.76_220/76/04_1i` |
| 3477 | `P.220_N.76_220/76/04_1` | SEKOLAH KEBANGSAAN SG. KELABIT | j | `P.220_N.76_220/76/04_1j` |
| 3478 | `P.220_N.76_220/76/04_1` | RH. BANTAN SG. LUTONG ATAS | k | `P.220_N.76_220/76/04_1k` |
| 3479 | `P.220_N.76_220/76/04_1` | RH. DRAHMAM, SG. LUTONG ATAS | l | `P.220_N.76_220/76/04_1l` |
| 3480 | `P.220_N.76_220/76/04_1` | RH. JANTING SG. LUTONG BAWAH, BAKONG | m | `P.220_N.76_220/76/04_1m` |
| 3481 | `P.220_N.76_220/76/04_1` | RH. ALI SG. SEBUKUT, BAKONG | n | `P.220_N.76_220/76/04_1n` |
| 3482 | `P.220_N.76_220/76/04_1` | SEKOLAH KEBANGSAAN SG. BAKAS | o | `P.220_N.76_220/76/04_1o` |
| 3483 | `P.220_N.76_220/76/04_1` | RH. PENGIRAN DAMIT SG. MANGKA BAKONG | p | `P.220_N.76_220/76/04_1p` |

#### 484. `P.220_N.76_220/76/05_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3484 | `P.220_N.76_220/76/05_1` | SEKOLAH KEBANGSAAN PENGARAH ENTERI | a | `P.220_N.76_220/76/05_1a` |
| 3486 | `P.220_N.76_220/76/05_1` | RH. RIGGIE ANAK BELULOK SG. NAT ULU TERU | b | `P.220_N.76_220/76/05_1b` |
| 3487 | `P.220_N.76_220/76/05_1` | SEKOLAH KEBANGSAAN SG BONG | c | `P.220_N.76_220/76/05_1c` |
| 3489 | `P.220_N.76_220/76/05_1` | SEKOLAH KEBANGSAAN SUNGAI BAIN | d | `P.220_N.76_220/76/05_1d` |
| 3490 | `P.220_N.76_220/76/05_1` | RH. BANYAH AK BANYI, SG. SEBUBU | e | `P.220_N.76_220/76/05_1e` |
| 3491 | `P.220_N.76_220/76/05_1` | SEKOLAH KEBANGSAAN SUNGAI SEBATU | f | `P.220_N.76_220/76/05_1f` |
| 3492 | `P.220_N.76_220/76/05_1` | SEKOLAH KEBANGSAAN SG PEKING | g | `P.220_N.76_220/76/05_1g` |

#### 485. `P.220_N.76_220/76/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3485 | `P.220_N.76_220/76/05_2` | SEKOLAH KEBANGSAAN PENGARAH ENTERI | a | `P.220_N.76_220/76/05_2a` |
| 3488 | `P.220_N.76_220/76/05_2` | SEKOLAH KEBANGSAAN SG BONG | b | `P.220_N.76_220/76/05_2b` |

#### 486. `P.220_N.76_220/76/06_1` (8 occurrences, 8 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3493 | `P.220_N.76_220/76/06_1` | RH. JAMU NG. SERIDAN | a | `P.220_N.76_220/76/06_1a` |
| 3494 | `P.220_N.76_220/76/06_1` | RH. NYIGA SG. NIRU | b | `P.220_N.76_220/76/06_1b` |
| 3495 | `P.220_N.76_220/76/06_1` | RH. BALANG LG. TUYUT | c | `P.220_N.76_220/76/06_1c` |
| 3496 | `P.220_N.76_220/76/06_1` | RH. HILLARY JUNGANG NG. AJOI | d | `P.220_N.76_220/76/06_1d` |
| 3497 | `P.220_N.76_220/76/06_1` | RH. KAJAN SIGEH LG. TERU, TINJAR | e | `P.220_N.76_220/76/06_1e` |
| 3498 | `P.220_N.76_220/76/06_1` | RH. MERAN SURANG, LOGAN BUNUT | f | `P.220_N.76_220/76/06_1f` |
| 3499 | `P.220_N.76_220/76/06_1` | RH. MUSIN SEBATANG BOK | g | `P.220_N.76_220/76/06_1g` |
| 3500 | `P.220_N.76_220/76/06_1` | SEKOLAH KEBANGSAAN LONG SEPILING | h | `P.220_N.76_220/76/06_1h` |

#### 487. `P.220_N.76_220/76/07_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3501 | `P.220_N.76_220/76/07_1` | RH. MALINA NG. TISAM | a | `P.220_N.76_220/76/07_1a` |
| 3502 | `P.220_N.76_220/76/07_1` | SEKOLAH KEBANGSAAN LONG JEGAN | b | `P.220_N.76_220/76/07_1b` |
| 3503 | `P.220_N.76_220/76/07_1` | RH. JARAU LG. TABING | c | `P.220_N.76_220/76/07_1c` |
| 3504 | `P.220_N.76_220/76/07_1` | SEKOLAH KEBANGSAAN LONG TERAN KANAN | d | `P.220_N.76_220/76/07_1d` |
| 3506 | `P.220_N.76_220/76/07_1` | SEKOLAH KEBANGSAAN SG. SEPUTI | e | `P.220_N.76_220/76/07_1e` |
| 3507 | `P.220_N.76_220/76/07_1` | SEKOLAH MENENGAH KEBANGSAAN TINJAR | f | `P.220_N.76_220/76/07_1f` |

#### 488. `P.220_N.76_220/76/07_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3505 | `P.220_N.76_220/76/07_2` | SEKOLAH KEBANGSAAN LONG TERAN KANAN | a | `P.220_N.76_220/76/07_2a` |
| 3508 | `P.220_N.76_220/76/07_2` | SEKOLAH MENENGAH KEBANGSAAN TINJAR | b | `P.220_N.76_220/76/07_2b` |

#### 489. `P.220_N.77_220/77/01_1` (4 occurrences, 4 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3511 | `P.220_N.77_220/77/01_1` | SEKOLAH KEBANGSAAN LONG MIRI | a | `P.220_N.77_220/77/01_1a` |
| 3512 | `P.220_N.77_220/77/01_1` | SEKOLAH KEBANGSAAN LONG PELUTAN | b | `P.220_N.77_220/77/01_1b` |
| 3513 | `P.220_N.77_220/77/01_1` | SEKOLAH KEBANGSAAN UMA BAWANG | c | `P.220_N.77_220/77/01_1c` |
| 3514 | `P.220_N.77_220/77/01_1` | SEKOLAH KEBANGSAAN LONG PILLAH | d | `P.220_N.77_220/77/01_1d` |

#### 490. `P.220_N.77_220/77/02_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3516 | `P.220_N.77_220/77/02_1` | DEWAN SERBAGUNA KAMPUNG APAU NYARING | a | `P.220_N.77_220/77/02_1a` |
| 3517 | `P.220_N.77_220/77/02_1` | RH. TK BITANG SAKAI LG. LIAW | b | `P.220_N.77_220/77/02_1b` |
| 3518 | `P.220_N.77_220/77/02_1` | SEKOLAH KEBANGSAAN LONG ATON | c | `P.220_N.77_220/77/02_1c` |
| 3519 | `P.220_N.77_220/77/02_1` | SEKOLAH KEBANGSAAN LONG SOBENG | d | `P.220_N.77_220/77/02_1d` |
| 3520 | `P.220_N.77_220/77/02_1` | SEKOLAH KEBANGSAAN LONG LOYANG | e | `P.220_N.77_220/77/02_1e` |
| 3522 | `P.220_N.77_220/77/02_1` | RH. LG. AYA | f | `P.220_N.77_220/77/02_1f` |
| 3523 | `P.220_N.77_220/77/02_1` | KPG. LG. TEBANYI TINJAR | g | `P.220_N.77_220/77/02_1g` |

#### 491. `P.220_N.77_220/77/03_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3524 | `P.220_N.77_220/77/03_1` | SEKOLAH KEBANGSAAN LONG KESSEH | a | `P.220_N.77_220/77/03_1a` |
| 3525 | `P.220_N.77_220/77/03_1` | SEKOLAH KEBANGSAAN LONG NAAH | b | `P.220_N.77_220/77/03_1b` |
| 3526 | `P.220_N.77_220/77/03_1` | RH. TK MADANG WAN | c | `P.220_N.77_220/77/03_1c` |
| 3527 | `P.220_N.77_220/77/03_1` | SEKOLAH KEBANGSAAN LONG LUTENG | d | `P.220_N.77_220/77/03_1d` |
| 3529 | `P.220_N.77_220/77/03_1` | RH. TK IBAU WAN LG. TEBANGAN | e | `P.220_N.77_220/77/03_1e` |

#### 492. `P.220_N.77_220/77/04_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3530 | `P.220_N.77_220/77/04_1` | SEKOLAH KEBANGSAAN ST PIUS LONG SAN | a | `P.220_N.77_220/77/04_1a` |
| 3532 | `P.220_N.77_220/77/04_1` | SEKOLAH KEBANGSAAN LONG PALAI | b | `P.220_N.77_220/77/04_1b` |
| 3533 | `P.220_N.77_220/77/04_1` | SEKOLAH KEBANGSAAN LONG ANAP | c | `P.220_N.77_220/77/04_1c` |
| 3534 | `P.220_N.77_220/77/04_1` | SEKOLAH KEBANGSAAN LONG APU | d | `P.220_N.77_220/77/04_1d` |
| 3535 | `P.220_N.77_220/77/04_1` | RH. TK MATTHEW BELULOK LALO, KPG LEPO GAH | e | `P.220_N.77_220/77/04_1e` |
| 3536 | `P.220_N.77_220/77/04_1` | RH. TK LONG TAP | f | `P.220_N.77_220/77/04_1f` |

#### 493. `P.220_N.77_220/77/05_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3537 | `P.220_N.77_220/77/05_1` | RH. TK JACOB LAWAI NAWAN SG. DUA BARAM | a | `P.220_N.77_220/77/05_1a` |
| 3538 | `P.220_N.77_220/77/05_1` | SEKOLAH KEBANGSAAN LONG LAPUT | b | `P.220_N.77_220/77/05_1b` |
| 3540 | `P.220_N.77_220/77/05_1` | SEKOLAH MENENGAH KEBANGSAAN LONG LAMA | c | `P.220_N.77_220/77/05_1c` |
| 3542 | `P.220_N.77_220/77/05_1` | RH. TK ANTHONY NGAU KPG. UMA AKEH | d | `P.220_N.77_220/77/05_1d` |
| 3543 | `P.220_N.77_220/77/05_1` | SEKOLAH KEBANGSAAN MOREK (LONG BANYOK) | e | `P.220_N.77_220/77/05_1e` |
| 3544 | `P.220_N.77_220/77/05_1` | SEKOLAH KEBANGSAAN LONG IKANG | f | `P.220_N.77_220/77/05_1f` |

#### 494. `P.220_N.77_220/77/05_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3539 | `P.220_N.77_220/77/05_2` | SEKOLAH KEBANGSAAN LONG LAPUT | a | `P.220_N.77_220/77/05_2a` |
| 3541 | `P.220_N.77_220/77/05_2` | SEKOLAH MENENGAH KEBANGSAAN LONG LAMA | b | `P.220_N.77_220/77/05_2b` |

#### 495. `P.220_N.77_220/77/06_1` (9 occurrences, 9 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3545 | `P.220_N.77_220/77/06_1` | SEKOLAH KEBANGSAAN LONG KEVOK | a | `P.220_N.77_220/77/06_1a` |
| 3546 | `P.220_N.77_220/77/06_1` | SEKOLAH KEBANGSAAN LONG BEDIAN | b | `P.220_N.77_220/77/06_1b` |
| 3548 | `P.220_N.77_220/77/06_1` | TADIKA KAMPUNG LONG BELUK | c | `P.220_N.77_220/77/06_1c` |
| 3549 | `P.220_N.77_220/77/06_1` | SEKOLAH KEBANGSAAN LONG BEMANG | d | `P.220_N.77_220/77/06_1d` |
| 3552 | `P.220_N.77_220/77/06_1` | SEKOLAH KEBANGSAAN LONG ATIP | e | `P.220_N.77_220/77/06_1e` |
| 3553 | `P.220_N.77_220/77/06_1` | SEKOLAH KEBANGSAAN LONG WAT | f | `P.220_N.77_220/77/06_1f` |
| 3554 | `P.220_N.77_220/77/06_1` | DEWAN SERBAGUNA KAMPUNG LONG LATEI | g | `P.220_N.77_220/77/06_1g` |
| 3555 | `P.220_N.77_220/77/06_1` | RH. TK KAMPUNG LONG WIN | h | `P.220_N.77_220/77/06_1h` |
| 3556 | `P.220_N.77_220/77/06_1` | BALAI RAYA KPG BA' KABENG | i | `P.220_N.77_220/77/06_1i` |

#### 496. `P.220_N.77_220/77/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3547 | `P.220_N.77_220/77/06_2` | SEKOLAH KEBANGSAAN LONG BEDIAN | a | `P.220_N.77_220/77/06_2a` |
| 3550 | `P.220_N.77_220/77/06_2` | SEKOLAH KEBANGSAAN LONG BEMANG | b | `P.220_N.77_220/77/06_2b` |

#### 497. `P.220_N.78_220/78/01_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3558 | `P.220_N.78_220/78/01_1` | SEKOLAH KEBANGSAAN LUBOK NIBONG | a | `P.220_N.78_220/78/01_1a` |
| 3559 | `P.220_N.78_220/78/01_1` | SEKOLAH JENIS KEBANGSAAN (CINA) HUA NAM | b | `P.220_N.78_220/78/01_1b` |

#### 498. `P.220_N.78_220/78/02_1` (11 occurrences, 11 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3560 | `P.220_N.78_220/78/02_1` | SEKOLAH KEBANGSAAN SG SETAPANG | a | `P.220_N.78_220/78/02_1a` |
| 3561 | `P.220_N.78_220/78/02_1` | RH. LALO SELEJAU | b | `P.220_N.78_220/78/02_1b` |
| 3562 | `P.220_N.78_220/78/02_1` | RH. JUGAH SG. BELASOI | c | `P.220_N.78_220/78/02_1c` |
| 3563 | `P.220_N.78_220/78/02_1` | RH. BLALANG ANAK ATOM SG. SENGKABANG | d | `P.220_N.78_220/78/02_1d` |
| 3564 | `P.220_N.78_220/78/02_1` | RH. LANSAM SG. DABAI | e | `P.220_N.78_220/78/02_1e` |
| 3565 | `P.220_N.78_220/78/02_1` | SEKOLAH KEBANGSAAN POYUT | f | `P.220_N.78_220/78/02_1f` |
| 3566 | `P.220_N.78_220/78/02_1` | SEKOLAH KEBANGSAAN RUMAH GUDANG | g | `P.220_N.78_220/78/02_1g` |
| 3567 | `P.220_N.78_220/78/02_1` | SEKOLAH KEBANGSAAN LONG LENEI | h | `P.220_N.78_220/78/02_1h` |
| 3568 | `P.220_N.78_220/78/02_1` | RH. ADANG SG. RIDAN | i | `P.220_N.78_220/78/02_1i` |
| 3569 | `P.220_N.78_220/78/02_1` | SEKOLAH KEBANGSAAN SG BRIT | j | `P.220_N.78_220/78/02_1j` |
| 3570 | `P.220_N.78_220/78/02_1` | RH. RIDAB AK SELAT SG. PASIR | k | `P.220_N.78_220/78/02_1k` |

#### 499. `P.220_N.78_220/78/03_1` (11 occurrences, 11 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3571 | `P.220_N.78_220/78/03_1` | RH. TK TEARIE ABENG | a | `P.220_N.78_220/78/03_1a` |
| 3572 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG SAIT | b | `P.220_N.78_220/78/03_1b` |
| 3573 | `P.220_N.78_220/78/03_1` | RH. TK JOHN JAU WAN LG. SEMIANG | c | `P.220_N.78_220/78/03_1c` |
| 3574 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG TUNGAN | d | `P.220_N.78_220/78/03_1d` |
| 3575 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG MOH | e | `P.220_N.78_220/78/03_1e` |
| 3576 | `P.220_N.78_220/78/03_1` | RH. TK LUCAS NGAU LG. SELAAN TEPUAN | f | `P.220_N.78_220/78/03_1f` |
| 3577 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG MEKABAR | g | `P.220_N.78_220/78/03_1g` |
| 3578 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG JEKITAN | h | `P.220_N.78_220/78/03_1h` |
| 3579 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG JEEH | i | `P.220_N.78_220/78/03_1i` |
| 3580 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LONG LAMEI | j | `P.220_N.78_220/78/03_1j` |
| 3581 | `P.220_N.78_220/78/03_1` | SEKOLAH KEBANGSAAN LIO MATO | k | `P.220_N.78_220/78/03_1k` |

#### 500. `P.220_N.78_220/78/04_1` (7 occurrences, 7 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3582 | `P.220_N.78_220/78/04_1` | SEKOLAH KEBANGSAAN BATU BUNGAN | a | `P.220_N.78_220/78/04_1a` |
| 3583 | `P.220_N.78_220/78/04_1` | SEKOLAH KEBANGSAAN LONG SERIDAN | b | `P.220_N.78_220/78/04_1b` |
| 3584 | `P.220_N.78_220/78/04_1` | SEKOLAH KEBANGSAAN PENGHULU BAYA MALLANG | c | `P.220_N.78_220/78/04_1c` |
| 3585 | `P.220_N.78_220/78/04_1` | SEKOLAH KEBANGSAAN LONG PANAI | d | `P.220_N.78_220/78/04_1d` |
| 3587 | `P.220_N.78_220/78/04_1` | RH. TK CHRISTHOPHER PUSU YUN LG UKOK | e | `P.220_N.78_220/78/04_1e` |
| 3588 | `P.220_N.78_220/78/04_1` | RH. TK ASONG JABAN | f | `P.220_N.78_220/78/04_1f` |
| 3589 | `P.220_N.78_220/78/04_1` | SEKOLAH KEBANGSAAN KUALA TUTOH | g | `P.220_N.78_220/78/04_1g` |

#### 501. `P.220_N.78_220/78/05_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3590 | `P.220_N.78_220/78/05_1` | RH. TK TULOI BAYO LG. PELUAN | a | `P.220_N.78_220/78/05_1a` |
| 3591 | `P.220_N.78_220/78/05_1` | SEKOLAH KEBANGSAAN LONG BANGA | b | `P.220_N.78_220/78/05_1b` |

#### 502. `P.220_N.78_220/78/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3594 | `P.220_N.78_220/78/07_1` | RH. TK BUJANG PA`UKAT | a | `P.220_N.78_220/78/07_1a` |
| 3595 | `P.220_N.78_220/78/07_1` | RH. TK MAREN PU`UN PA`LUNGAN | b | `P.220_N.78_220/78/07_1b` |

#### 503. `P.220_N.78_220/78/08_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3596 | `P.220_N.78_220/78/08_1` | RH. TK BARANG LUGUN | a | `P.220_N.78_220/78/08_1a` |
| 3597 | `P.220_N.78_220/78/08_1` | SEKOLAH MENENGAH KEBANGSAAN BARIO | b | `P.220_N.78_220/78/08_1b` |

#### 504. `P.220_N.78_220/78/09_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3598 | `P.220_N.78_220/78/09_1` | SEKOLAH KEBANGSAAN PA ' DALLIH | a | `P.220_N.78_220/78/09_1a` |
| 3599 | `P.220_N.78_220/78/09_1` | RH. TK MAREN LUGUN REMUDU | b | `P.220_N.78_220/78/09_1b` |

#### 505. `P.221_N.79_221/79/00_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3602 | `P.221_N.79_221/79/00_1` | RUANG A, BALAI POLIS LIMBANG | a | `P.221_N.79_221/79/00_1a` |
| 3603 | `P.221_N.79_221/79/00_1` | BANGUNAN KANTIN KOMPENI 'C' PGA LIMBANG | b | `P.221_N.79_221/79/00_1b` |

#### 506. `P.221_N.79_221/79/02_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3605 | `P.221_N.79_221/79/02_1` | SEKOLAH KEBANGSAAN LIMPAKI | a | `P.221_N.79_221/79/02_1a` |
| 3606 | `P.221_N.79_221/79/02_1` | DEWAN KPG. PATIAMBUN | b | `P.221_N.79_221/79/02_1b` |
| 3608 | `P.221_N.79_221/79/02_1` | TADIKA KEMAS LIMBANG | c | `P.221_N.79_221/79/02_1c` |

#### 507. `P.221_N.79_221/79/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3607 | `P.221_N.79_221/79/02_2` | DEWAN KPG. PATIAMBUN | a | `P.221_N.79_221/79/02_2a` |
| 3609 | `P.221_N.79_221/79/02_2` | TADIKA KEMAS LIMBANG | b | `P.221_N.79_221/79/02_2b` |

#### 508. `P.221_N.79_221/79/03_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3611 | `P.221_N.79_221/79/03_1` | SEKOLAH KEBANGSAAN KAMPUNG PAHLAWAN | a | `P.221_N.79_221/79/03_1a` |
| 3614 | `P.221_N.79_221/79/03_1` | SEKOLAH KEBANGSAAN KUBONG | b | `P.221_N.79_221/79/03_1b` |
| 3615 | `P.221_N.79_221/79/03_1` | SEKOLAH KEBANGSAAN BATU EMPAT | c | `P.221_N.79_221/79/03_1c` |

#### 509. `P.221_N.79_221/79/06_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3625 | `P.221_N.79_221/79/06_1` | SEKOLAH KEBANGSAAN SUNGAI POYAN | a | `P.221_N.79_221/79/06_1a` |
| 3628 | `P.221_N.79_221/79/06_1` | DEWAN PIBG, SEKOLAH KEBANGSAAN MELAYU PUSAT | b | `P.221_N.79_221/79/06_1b` |

#### 510. `P.221_N.79_221/79/06_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3626 | `P.221_N.79_221/79/06_2` | SEKOLAH KEBANGSAAN SUNGAI POYAN | a | `P.221_N.79_221/79/06_2a` |
| 3629 | `P.221_N.79_221/79/06_2` | DEWAN PIBG, SEKOLAH KEBANGSAAN MELAYU PUSAT | b | `P.221_N.79_221/79/06_2b` |

#### 511. `P.221_N.79_221/79/07_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3630 | `P.221_N.79_221/79/07_1` | DEWAN SURAU KPG. LIMPAONG | a | `P.221_N.79_221/79/07_1a` |
| 3632 | `P.221_N.79_221/79/07_1` | TADIKA KEMAS SERI PELIMBU, KPG. LIMPAONG | b | `P.221_N.79_221/79/07_1b` |

#### 512. `P.221_N.79_221/79/08_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3633 | `P.221_N.79_221/79/08_1` | SEKOLAH KEBANGSAAN GADONG | a | `P.221_N.79_221/79/08_1a` |
| 3634 | `P.221_N.79_221/79/08_1` | MASJID KPG. PENDAM (BAHAGIAN BAWAH) | b | `P.221_N.79_221/79/08_1b` |

#### 513. `P.221_N.79_221/79/09_1` (3 occurrences, 3 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3635 | `P.221_N.79_221/79/09_1` | SEKOLAH KEBANGSAAN TELAHAK | a | `P.221_N.79_221/79/09_1a` |
| 3636 | `P.221_N.79_221/79/09_1` | SURAU KPG. IPAI | b | `P.221_N.79_221/79/09_1b` |
| 3637 | `P.221_N.79_221/79/09_1` | SEKOLAH KEBANGSAAN MERITAM | c | `P.221_N.79_221/79/09_1c` |

#### 514. `P.221_N.79_221/79/10_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3639 | `P.221_N.79_221/79/10_1` | DEWAN INSPIRASI, SEKOLAH MENENGAH KEBANGSAAN KUBONG | a | `P.221_N.79_221/79/10_1a` |
| 3641 | `P.221_N.79_221/79/10_1` | SEKOLAH KEBANGSAAN R.
C. KUBONG | b | `P.221_N.79_221/79/10_1b` |
| 3642 | `P.221_N.79_221/79/10_1` | BALAI RAYA KPG. BATU BAKARANG | c | `P.221_N.79_221/79/10_1c` |
| 3644 | `P.221_N.79_221/79/10_1` | SEKOLAH KEBANGSAAN BUKIT LUBA | d | `P.221_N.79_221/79/10_1d` |
| 3645 | `P.221_N.79_221/79/10_1` | SEKOLAH KEBANGSAAN MERAMBUT | e | `P.221_N.79_221/79/10_1e` |
| 3646 | `P.221_N.79_221/79/10_1` | SURAU (HALAMAN) KPG. PANGKALAN REJAB | f | `P.221_N.79_221/79/10_1f` |

#### 515. `P.221_N.79_221/79/10_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3640 | `P.221_N.79_221/79/10_2` | DEWAN INSPIRASI, SEKOLAH MENENGAH KEBANGSAAN KUBONG | a | `P.221_N.79_221/79/10_2a` |
| 3643 | `P.221_N.79_221/79/10_2` | BALAI RAYA KPG. BATU BAKARANG | b | `P.221_N.79_221/79/10_2b` |

#### 516. `P.221_N.80_221/80/01_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3649 | `P.221_N.80_221/80/01_1` | SEKOLAH KEBANGSAAN BATU DANAU | a | `P.221_N.80_221/80/01_1a` |
| 3651 | `P.221_N.80_221/80/01_1` | DEWAN KPG. BIDANG | b | `P.221_N.80_221/80/01_1b` |
| 3652 | `P.221_N.80_221/80/01_1` | DEWAN PENGKALAN MADANG | c | `P.221_N.80_221/80/01_1c` |
| 3653 | `P.221_N.80_221/80/01_1` | SEKOLAH KEBANGSAAN PANGKALAN JAWA | d | `P.221_N.80_221/80/01_1d` |
| 3654 | `P.221_N.80_221/80/01_1` | SEKOLAH KEBANGSAAN KUALA AWANG | e | `P.221_N.80_221/80/01_1e` |

#### 517. `P.221_N.80_221/80/02_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3655 | `P.221_N.80_221/80/02_1` | SEKOLAH KEBANGSAAN KUALA PENGANAN | a | `P.221_N.80_221/80/02_1a` |
| 3657 | `P.221_N.80_221/80/02_1` | SEKOLAH KEBANGSAAN NANGA MERIT | b | `P.221_N.80_221/80/02_1b` |
| 3658 | `P.221_N.80_221/80/02_1` | RH. LUTA ENGKASING LUBAI | c | `P.221_N.80_221/80/02_1c` |
| 3659 | `P.221_N.80_221/80/02_1` | SEKOLAH KEBANGSAAN MENUANG | d | `P.221_N.80_221/80/02_1d` |
| 3661 | `P.221_N.80_221/80/02_1` | RH. DAMPIN AK LAYAN TERIMAH | e | `P.221_N.80_221/80/02_1e` |
| 3662 | `P.221_N.80_221/80/02_1` | RH. SLI AK MINGAN MENGARI, JALAN MEDAMIT | f | `P.221_N.80_221/80/02_1f` |

#### 518. `P.221_N.80_221/80/02_2` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3656 | `P.221_N.80_221/80/02_2` | SEKOLAH KEBANGSAAN KUALA PENGANAN | a | `P.221_N.80_221/80/02_2a` |
| 3660 | `P.221_N.80_221/80/02_2` | SEKOLAH KEBANGSAAN MENUANG | b | `P.221_N.80_221/80/02_2b` |

#### 519. `P.221_N.80_221/80/03_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3663 | `P.221_N.80_221/80/03_1` | RH. LIBAN, TERUSAN LUBAI | a | `P.221_N.80_221/80/03_1a` |
| 3664 | `P.221_N.80_221/80/03_1` | DEWAN KPG. LUBOK LASAS | b | `P.221_N.80_221/80/03_1b` |
| 3665 | `P.221_N.80_221/80/03_1` | DEWAN KPG. UKONG | c | `P.221_N.80_221/80/03_1c` |
| 3666 | `P.221_N.80_221/80/03_1` | DEWAN SERBAGUNA KPG. BULOH BALUI | d | `P.221_N.80_221/80/03_1d` |
| 3667 | `P.221_N.80_221/80/03_1` | SEKOLAH KEBANGSAAN TANJONG | e | `P.221_N.80_221/80/03_1e` |
| 3668 | `P.221_N.80_221/80/03_1` | DEWAN KPG. RANGGU | f | `P.221_N.80_221/80/03_1f` |

#### 520. `P.221_N.80_221/80/04_1` (6 occurrences, 6 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3669 | `P.221_N.80_221/80/04_1` | SEKOLAH KEBANGSAAN LONG NAPIR | a | `P.221_N.80_221/80/04_1a` |
| 3670 | `P.221_N.80_221/80/04_1` | RH. LAWAI, SELIDONG | b | `P.221_N.80_221/80/04_1b` |
| 3671 | `P.221_N.80_221/80/04_1` | RH. BRAIN, MERUYU | c | `P.221_N.80_221/80/04_1c` |
| 3672 | `P.221_N.80_221/80/04_1` | RH. SING, KARANGAN MA MEDAMIT | d | `P.221_N.80_221/80/04_1d` |
| 3673 | `P.221_N.80_221/80/04_1` | SEKOLAH KEBANGSAAN MELABAN | e | `P.221_N.80_221/80/04_1e` |
| 3674 | `P.221_N.80_221/80/04_1` | SEKOLAH KEBANGSAAN KUALA MEDALAM | f | `P.221_N.80_221/80/04_1f` |

#### 521. `P.221_N.80_221/80/05_1` (5 occurrences, 5 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3675 | `P.221_N.80_221/80/05_1` | RH. EKOM SEPANGAH | a | `P.221_N.80_221/80/05_1a` |
| 3676 | `P.221_N.80_221/80/05_1` | DEWAN MASYARAKAT MEDAMIT | b | `P.221_N.80_221/80/05_1b` |
| 3677 | `P.221_N.80_221/80/05_1` | RH. AKOH KARANGAN ENSURAI | c | `P.221_N.80_221/80/05_1c` |
| 3678 | `P.221_N.80_221/80/05_1` | RH. TAN ANAK KAYAN, RANTAU KIRAN | d | `P.221_N.80_221/80/05_1d` |
| 3679 | `P.221_N.80_221/80/05_1` | SEKOLAH KEBANGSAAN BUKIT BATU | e | `P.221_N.80_221/80/05_1e` |

#### 522. `P.222_N.81_222/81/02_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3683 | `P.222_N.81_222/81/02_1` | DEWAN KPG. PURU SIA | a | `P.222_N.81_222/81/02_1a` |
| 3684 | `P.222_N.81_222/81/02_1` | SEKOLAH KEBANGSAAN PURU SIA | b | `P.222_N.81_222/81/02_1b` |

#### 523. `P.222_N.82_222/82/15_1` (2 occurrences, 2 centres)

| Line | Current UNIQUE CODE | POLLING CENTRE | Recommended Suffix | New UNIQUE CODE |
| --- | --- | --- | --- | --- |
| 3736 | `P.222_N.82_222/82/15_1` | DEWAN KPG. SUALAI | a | `P.222_N.82_222/82/15_1a` |
| 3737 | `P.222_N.82_222/82/15_1` | DEWAN MASYARAKAT KPG. SIANG-SIANG LAUT | b | `P.222_N.82_222/82/15_1b` |

## Recommendations

1. **Apply suffix disambiguation**: All 523 duplicate UNIQUE CODEs must be suffixed with letters (a, b, c, ...) based on the order of first appearance of each distinct Polling Centre.
2. **Only modify column 1 (UNIQUE CODE)**: No other columns should change.
3. **Pattern explanation**: The duplicates arise because multiple Polling Centres share the same Polling District Code and Voting Channel Number. The current ID scheme (ParliCode_DUNCode_DMKOD_Channel) does not distinguish between different Polling Centres within the same polling district/channel.
4. **Scope**: This affects all 82 DUNs — it is a systematic issue, not limited to specific constituencies.

