#!/usr/bin/env python3
"""
diff_analysis.py

Compares to-review.csv.orig (before any fixes) with to-review.csv (final)
and prints a Markdown table summarising every changed row, grouped by
change type.

Run from the sarawak-dun-2021 directory:
    python3 diff_analysis.py
"""

import csv
import re
from collections import defaultdict


# ── helpers ──────────────────────────────────────────────────────────────────

def read_csv(path):
    with open(path, newline="", encoding="utf-8") as f:
        return list(csv.reader(f))


def categorise(orig_row, final_row, diffs, header):
    """Return a short human-readable description of what changed."""
    labels = []

    # col indices (0-based)
    COL_UC  = 0   # UNIQUE CODE
    COL_PDC = 7   # POLLING DISTRICT CODE
    COL_PC  = 9   # POLLING CENTRE
    COL_DUN = 6   # STATE CONSTITUENCY NAME

    changed_names = [header[d] for d in diffs]

    # 1. Postal DMKOD: /UNDI POS -> /POS in UNIQUE CODE and/or POLLING DISTRICT CODE
    if COL_UC in diffs and "/UNDI POS" in orig_row[COL_UC] and "/POS" in final_row[COL_UC]:
        labels.append("Postal DMKOD: /UNDI POS → /POS in UNIQUE CODE")
    if COL_PDC in diffs and "UNDI POS" in orig_row[COL_PDC] and orig_row[COL_PDC].replace("UNDI POS","POS") == final_row[COL_PDC]:
        labels.append("Postal DMKOD: /UNDI POS → /POS in POLLING DISTRICT CODE")

    # 2. BA`KELALAN backtick fix
    backtick_cols = [d for d in diffs if "BA`KELALAN" in orig_row[d] and "BA'KELALAN" in final_row[d]]
    if backtick_cols:
        labels.append(f"BA`KELALAN → BA'KELALAN in {', '.join(header[d] for d in backtick_cols)}")

    # 3. Polling centre newline removed
    newline_cols = [d for d in diffs if "\n" in orig_row[d] and "\n" not in final_row[d]]
    if newline_cols:
        labels.append(f"Embedded newline removed from POLLING CENTRE: {repr(final_row[COL_PC])}")

    # 4. UNIQUE CODE suffix added/changed (and not already explained by postal fix)
    if COL_UC in diffs and not any("Postal DMKOD" in l for l in labels):
        o_uc = orig_row[COL_UC]
        f_uc = final_row[COL_UC]
        # strip suffix from final to compare base
        base_match = re.match(r'^(.*_\d+)([a-z])$', f_uc)
        if base_match and base_match.group(1) == o_uc:
            labels.append(f"Suffix added: {o_uc} → {f_uc}")
        elif re.match(r'^(.*_\d+)([a-z])$', o_uc) and re.match(r'^(.*_\d+)([a-z])$', f_uc):
            labels.append(f"Suffix changed: {o_uc} → {f_uc}")
        elif re.match(r'^(.*_\d+)([a-z])$', f_uc):
            labels.append(f"Suffix added: {o_uc} → {f_uc}")
        else:
            labels.append(f"UNIQUE CODE changed: {o_uc} → {f_uc}")

    if not labels:
        # fallback
        parts = [f"{header[d]}: {repr(orig_row[d])} → {repr(final_row[d])}" for d in diffs]
        labels.append("; ".join(parts))

    return "; ".join(labels)


# ── main ─────────────────────────────────────────────────────────────────────

def main():
    orig_rows  = read_csv("to-review.csv.orig")
    final_rows = read_csv("to-review.csv")
    header = orig_rows[0]

    assert header == final_rows[0], "Headers differ — cannot compare"
    assert len(orig_rows) == len(final_rows), (
        f"Row count differs: orig={len(orig_rows)} final={len(final_rows)}"
    )

    total_data = len(orig_rows) - 1

    # Collect all diffs
    changes = []  # (csv_row_num, dun, orig_uc, final_uc, diffs, description)

    COL_UC  = 0
    COL_DUN = 5   # STATE CONSTITUENCY CODE

    for i in range(1, len(orig_rows)):
        o = orig_rows[i]
        f = final_rows[i]
        diffs = [j for j in range(min(len(o), len(f))) if o[j] != f[j]]
        if diffs:
            desc = categorise(o, f, diffs, header)
            dun  = o[COL_DUN] if len(o) > COL_DUN else "?"
            changes.append((i + 1, dun, o[COL_UC], f[COL_UC], diffs, desc))

    # ── Group by change type ──────────────────────────────────────────────────
    # We bucket changes so we can summarise compactly.
    BUCKET_POSTAL   = "postal_dmkod"
    BUCKET_BACKTICK = "backtick"
    BUCKET_NEWLINE  = "newline"
    BUCKET_SUFFIX   = "suffix"
    BUCKET_OTHER    = "other"

    def bucket(desc):
        d = desc.lower()
        if "postal dmkod" in d:    return BUCKET_POSTAL
        if "ba`kelalan"  in d:     return BUCKET_BACKTICK
        if "newline"     in d:     return BUCKET_NEWLINE
        if "suffix"      in d:     return BUCKET_SUFFIX
        return BUCKET_OTHER

    buckets = defaultdict(list)
    for c in changes:
        buckets[bucket(c[5])].append(c)

    # ── Print summary ─────────────────────────────────────────────────────────
    print(f"## Diff Summary: to-review.csv.orig → to-review.csv\n")
    print(f"- Total data rows : {total_data}")
    print(f"- Changed rows    : {len(changes)}")
    print(f"- Unchanged rows  : {total_data - len(changes)}\n")

    bucket_order = [
        (BUCKET_POSTAL,   "Fix 1 — Postal DMKOD (`/UNDI POS` → `/POS` in UNIQUE CODE & POLLING DISTRICT CODE)"),
        (BUCKET_BACKTICK, "Fix 2 — BA\`KELALAN backtick → apostrophe (STATE CONSTITUENCY NAME / POLLING DISTRICT NAME)"),
        (BUCKET_NEWLINE,  "Fix 3 — Embedded newline removed from POLLING CENTRE"),
        (BUCKET_SUFFIX,   "Fix 4 — UNIQUE CODE suffix added/corrected (district-level centre disambiguation)"),
        (BUCKET_OTHER,    "Other changes"),
    ]

    for bkey, btitle in bucket_order:
        rows = buckets.get(bkey, [])
        if not rows:
            continue

        print(f"### {btitle}\n")
        print(f"**{len(rows)} row(s) affected**\n")

        if bkey == BUCKET_SUFFIX:
            # For suffix changes, group by DUN and show concise info
            by_dun = defaultdict(list)
            for (csv_row, dun, o_uc, f_uc, diffs, desc) in rows:
                by_dun[dun].append((csv_row, o_uc, f_uc, desc))

            print("| DUN | Rows affected | Example: UNIQUE CODE before → after | Description |")
            print("|-----|--------------|--------------------------------------|-------------|")
            for dun in sorted(by_dun.keys(), key=lambda x: int(x.replace("N.", "")) if x.startswith("N.") else 999):
                dun_rows = by_dun[dun]
                ex_csv, ex_o, ex_f, ex_desc = dun_rows[0]
                # shorten for table
                desc_short = ex_desc if len(ex_desc) < 80 else ex_desc[:77] + "…"
                print(f"| {dun} | {len(dun_rows)} | `{ex_o}` → `{ex_f}` | {desc_short} |")

        elif bkey == BUCKET_POSTAL:
            print("| CSV Row | DUN | UNIQUE CODE before | UNIQUE CODE after | POLLING DISTRICT CODE before | after |")
            print("|---------|-----|--------------------|-------------------|------------------------------|-------|")
            for (csv_row, dun, o_uc, f_uc, diffs, desc) in rows:
                o_pdc = orig_rows[csv_row - 1][7]
                f_pdc = final_rows[csv_row - 1][7]
                print(f"| {csv_row} | {dun} | `{o_uc}` | `{f_uc}` | `{o_pdc}` | `{f_pdc}` |")

        elif bkey == BUCKET_BACKTICK:
            # Show unique DUNs / columns affected
            col_set = set()
            for (csv_row, dun, o_uc, f_uc, diffs, desc) in rows:
                for d in diffs:
                    col_set.add(header[d])
            print(f"- DUN affected: **N.81** (`BA'KELALAN`)")
            print(f"- Columns fixed: {', '.join(sorted(col_set))}")
            print(f"- All {len(rows)} rows in N.81 corrected (grave accent U+0060 → apostrophe U+0027)\n")

        elif bkey == BUCKET_NEWLINE:
            print("| CSV Row | DUN | POLLING CENTRE before | POLLING CENTRE after |")
            print("|---------|-----|-----------------------|----------------------|")
            for (csv_row, dun, o_uc, f_uc, diffs, desc) in rows:
                o_pc = orig_rows[csv_row - 1][9].replace("\n", "\\n")
                f_pc = final_rows[csv_row - 1][9]
                print(f"| {csv_row} | {dun} | `{o_pc}` | `{f_pc}` |")

        else:
            print("| CSV Row | DUN | Description |")
            print("|---------|-----|-------------|")
            for (csv_row, dun, o_uc, f_uc, diffs, desc) in rows:
                print(f"| {csv_row} | {dun} | {desc} |")

        print()


if __name__ == "__main__":
    main()
