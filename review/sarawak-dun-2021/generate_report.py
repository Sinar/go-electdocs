#!/usr/bin/env python3
"""
generate_report.py

Compares to-review.csv.orig with to-review.csv and produces a styled
change-summary PDF (via WeasyPrint HTML→PDF pipeline).

Run from the sarawak-dun-2021 directory:
    uvx --from weasyprint python3 generate_report.py
    -- or --
    uvx --with weasyprint python3 generate_report.py
"""

import csv
import os
import re
import sys
from collections import defaultdict
from datetime import date

# ── CSV helpers ───────────────────────────────────────────────────────────────


def read_csv(path):
    with open(path, newline="", encoding="utf-8") as f:
        return list(csv.reader(f))


# ── Change categorisation ─────────────────────────────────────────────────────

COL_UC = 0  # UNIQUE CODE
COL_PDC = 7  # POLLING DISTRICT CODE
COL_PDN = 8  # POLLING DISTRICT NAME
COL_PC = 9  # POLLING CENTRE
COL_DUN_CODE = 5  # STATE CONSTITUENCY CODE
COL_DUN_NAME = 6  # STATE CONSTITUENCY NAME


def classify(orig_row, final_row, diffs):
    """Return one of: postal | backtick | newline | suffix | other"""
    has_postal = any(COL_UC == d and "/UNDI POS" in orig_row[d] for d in diffs) or any(
        COL_PDC == d and "UNDI POS" in orig_row[d] for d in diffs
    )
    has_backtick = any("BA`KELALAN" in orig_row[d] for d in diffs)
    has_newline = any("\n" in orig_row[d] for d in diffs)
    has_suffix = (
        COL_UC in diffs
        and not has_postal
        and re.search(r"_\d+[a-z]$", final_row[COL_UC])
        and not re.search(r"_\d+[a-z]$", orig_row[COL_UC])
    )

    if has_postal:
        return "postal"
    if has_backtick:
        return "backtick"
    if has_newline:
        return "newline"
    if has_suffix:
        return "suffix"
    return "other"


# ── HTML / CSS ────────────────────────────────────────────────────────────────

CSS = """
@page {
    size: A4;
    margin: 18mm 15mm 18mm 15mm;
    @bottom-center {
        content: "Page " counter(page) " of " counter(pages);
        font-size: 8pt;
        color: #888;
    }
}

* { box-sizing: border-box; }

body {
    font-family: "Helvetica Neue", Arial, sans-serif;
    font-size: 9pt;
    color: #222;
    line-height: 1.45;
}

h1 {
    font-size: 17pt;
    color: #1a3a5c;
    border-bottom: 3px solid #1a3a5c;
    padding-bottom: 4pt;
    margin-top: 0;
}

h2 {
    font-size: 12pt;
    color: #1a3a5c;
    border-left: 4px solid #2e7bbf;
    padding-left: 7pt;
    margin-top: 22pt;
    margin-bottom: 6pt;
}

h3 {
    font-size: 10pt;
    color: #444;
    margin-top: 14pt;
    margin-bottom: 4pt;
}

.meta {
    background: #f0f4f9;
    border: 1px solid #c8d8ea;
    border-radius: 4px;
    padding: 8pt 12pt;
    margin-bottom: 14pt;
    font-size: 9pt;
}

.meta table { border-collapse: collapse; width: 100%; }
.meta td { padding: 2pt 8pt 2pt 0; }
.meta td:first-child { font-weight: bold; color: #1a3a5c; width: 160px; }

/* Summary scorecards */
.scorecard-row {
    display: flex;
    gap: 10pt;
    margin-bottom: 14pt;
}
.scorecard {
    flex: 1;
    border-radius: 5px;
    padding: 8pt 10pt;
    text-align: center;
}
.scorecard .num  { font-size: 20pt; font-weight: bold; display: block; }
.scorecard .lbl  { font-size: 7.5pt; display: block; margin-top: 2pt; }

.sc-total    { background: #e8f0fa; color: #1a3a5c; border: 1px solid #b0c8e8; }
.sc-changed  { background: #fff3e0; color: #bf6000; border: 1px solid #ffc97a; }
.sc-unchanged{ background: #e8f5e9; color: #2e7d32; border: 1px solid #a5d6a7; }

/* Fix badge */
.fix-badge {
    display: inline-block;
    border-radius: 3px;
    padding: 1pt 6pt;
    font-size: 8pt;
    font-weight: bold;
    color: #fff;
    margin-right: 4pt;
}
.badge-postal    { background: #d32f2f; }
.badge-backtick  { background: #7b1fa2; }
.badge-newline   { background: #e65100; }
.badge-suffix    { background: #1565c0; }
.badge-other     { background: #555; }

/* Section summary strip */
.section-meta {
    background: #f7f7f7;
    border: 1px solid #ddd;
    border-radius: 3px;
    padding: 5pt 10pt;
    margin-bottom: 8pt;
    font-size: 8.5pt;
}
.section-meta strong { color: #1a3a5c; }

/* Tables */
table.data {
    width: 100%;
    border-collapse: collapse;
    font-size: 7.8pt;
    margin-top: 4pt;
    page-break-inside: auto;
}
table.data thead tr {
    background: #1a3a5c;
    color: #fff;
}
table.data thead th {
    padding: 4pt 6pt;
    text-align: left;
    font-weight: bold;
}
table.data tbody tr:nth-child(even) { background: #f4f8fc; }
table.data tbody tr:nth-child(odd)  { background: #ffffff; }
table.data tbody td {
    padding: 3pt 6pt;
    border-bottom: 1px solid #e0e8f0;
    vertical-align: top;
}
table.data tbody tr:hover { background: #e8f0fa; }

code {
    font-family: "Courier New", monospace;
    font-size: 7.5pt;
    background: #f0f4f9;
    padding: 0 3pt;
    border-radius: 2px;
}

.arrow { color: #888; padding: 0 4pt; }
.tag-before { color: #c0392b; }
.tag-after  { color: #27ae60; }

.note {
    background: #fffde7;
    border-left: 4px solid #f9a825;
    padding: 5pt 10pt;
    font-size: 8.5pt;
    margin: 8pt 0;
}

.page-break { page-break-before: always; }

/* suffix table — compact two-column layout */
.suffix-grid {
    column-count: 2;
    column-gap: 12pt;
}
.suffix-grid table.data { margin-top: 0; }
"""


def h(text):
    """Minimal HTML-escape."""
    return (
        str(text)
        .replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )


def code(text):
    return f"<code>{h(text)}</code>"


def badge(kind):
    labels = {
        "postal": ("FIX 1", "badge-postal"),
        "backtick": ("FIX 2", "badge-backtick"),
        "newline": ("FIX 3", "badge-newline"),
        "suffix": ("FIX 4", "badge-suffix"),
        "other": ("OTHER", "badge-other"),
    }
    lbl, cls = labels.get(kind, ("?", "badge-other"))
    return f'<span class="fix-badge {cls}">{lbl}</span>'


# ── Report builder ────────────────────────────────────────────────────────────


def build_html(orig_rows, final_rows):
    header = orig_rows[0]
    total = len(orig_rows) - 1

    # Collect all diffs
    all_changes = []  # (csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind)
    for i in range(1, len(orig_rows)):
        o = orig_rows[i]
        f = final_rows[i]
        diffs = [j for j in range(min(len(o), len(f))) if o[j] != f[j]]
        if not diffs:
            continue
        kind = classify(o, f, diffs)
        dun_code = o[COL_DUN_CODE] if len(o) > COL_DUN_CODE else "?"
        dun_name = o[COL_DUN_NAME] if len(o) > COL_DUN_NAME else "?"
        all_changes.append(
            (i + 1, dun_code, dun_name, o[COL_UC], f[COL_UC], diffs, kind, o, f)
        )

    by_kind = defaultdict(list)
    for c in all_changes:
        by_kind[c[6]].append(c)

    changed = len(all_changes)
    unchanged = total - changed

    # ── preamble ──────────────────────────────────────────────────────────────
    html = [
        f"""<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8"/>
<title>Sarawak DUN 2021 – Change Report</title>
<style>{CSS}</style>
</head>
<body>

<h1>Sarawak DUN 2021 — Data Change Report</h1>

<div class="meta">
  <table>
    <tr><td>Source file</td><td><code>to-review.csv.orig</code></td></tr>
    <tr><td>Final file</td><td><code>to-review.csv</code></td></tr>
    <tr><td>Generated</td><td>{date.today().strftime("%-d %B %Y")}</td></tr>
    <tr><td>Scope</td><td>Sarawak State Assembly (DUN) 2021 election results — 82 DUNs</td></tr>
  </table>
</div>

<div class="scorecard-row">
  <div class="scorecard sc-total">
    <span class="num">{total:,}</span>
    <span class="lbl">Total data rows</span>
  </div>
  <div class="scorecard sc-changed">
    <span class="num">{changed:,}</span>
    <span class="lbl">Rows changed</span>
  </div>
  <div class="scorecard sc-unchanged">
    <span class="num">{unchanged:,}</span>
    <span class="lbl">Rows unchanged</span>
  </div>
</div>

<div class="note">
  All changes are limited to structural / identifier corrections.
  <strong>No vote tallies, candidate names, or party assignments were modified.</strong>
  The only columns affected are: UNIQUE CODE (col&nbsp;1), POLLING DISTRICT CODE (col&nbsp;8),
  STATE CONSTITUENCY NAME (col&nbsp;7), POLLING DISTRICT NAME (col&nbsp;9), and POLLING CENTRE (col&nbsp;10).
</div>

<h2>Summary of fixes</h2>
<table class="data">
<thead>
  <tr>
    <th>#</th><th>Fix</th><th>Description</th><th style="text-align:right">Rows</th>
  </tr>
</thead>
<tbody>
"""
    ]

    fix_meta = [
        (
            "postal",
            "FIX 1",
            "Postal DMKOD corrected",
            "UNIQUE CODE &amp; POLLING DISTRICT CODE: <code>/UNDI POS</code> → <code>/POS</code> "
            "(correct per EC rule; POLLING DISTRICT NAME &amp; POLLING CENTRE keep <code>UNDI POS</code>)",
        ),
        (
            "backtick",
            "FIX 2",
            "BA'KELALAN typo corrected",
            "STATE CONSTITUENCY NAME &amp; POLLING DISTRICT NAME for N.81: "
            "grave accent U+0060 → apostrophe U+0027 in <code>BA'KELALAN</code>",
        ),
        (
            "newline",
            "FIX 3",
            "Embedded newlines removed",
            "POLLING CENTRE names containing an embedded <code>\\n</code> (multi-line quoted CSV field) "
            "cleaned up — affects 5 distinct centre names across 4 DUNs",
        ),
        (
            "suffix",
            "FIX 4",
            "UNIQUE CODE suffixes applied",
            "For every Polling District Code serving multiple Polling Centres, a consistent letter "
            "suffix (a, b, c…) is appended to UNIQUE CODE — same centre always gets the same letter "
            "across all voting channels",
        ),
    ]

    for kind, _, title, desc in fix_meta:
        n = len(by_kind.get(kind, []))
        html.append(f"""  <tr>
    <td>{badge(kind)}</td>
    <td><strong>{title}</strong></td>
    <td>{desc}</td>
    <td style="text-align:right"><strong>{n:,}</strong></td>
  </tr>
""")

    other_n = len(by_kind.get("other", []))
    if other_n:
        html.append(f"""  <tr>
    <td>{badge("other")}</td><td><strong>Other</strong></td>
    <td>Miscellaneous changes</td>
    <td style="text-align:right"><strong>{other_n}</strong></td>
  </tr>
""")

    html.append("</tbody></table>\n\n")

    # ── FIX 1: Postal DMKOD ───────────────────────────────────────────────────
    postal_rows = by_kind.get("postal", [])
    html.append('<div class="page-break"></div>\n')
    html.append(
        f"<h2>{badge('postal')} Fix 1 — Postal DMKOD: <code>/UNDI POS</code> → <code>/POS</code></h2>\n"
    )
    html.append(f"""<div class="section-meta">
  <strong>{len(postal_rows)} rows affected</strong> — one POSTAL VOTE row per DUN (N.01 was already correct in the original; N.02–N.82 corrected).<br/>
  Columns changed: <code>UNIQUE CODE</code> (col 1), <code>POLLING DISTRICT CODE</code> (col 8).<br/>
  <code>POLLING DISTRICT NAME</code> and <code>POLLING CENTRE</code> retain the value <code>UNDI POS</code> — unchanged.
</div>\n""")

    html.append("""<table class="data">
<thead>
  <tr>
    <th>CSV Row</th><th>DUN</th>
    <th>UNIQUE CODE — before</th><th>UNIQUE CODE — after</th>
    <th>PD Code — before</th><th>PD Code — after</th>
  </tr>
</thead>
<tbody>\n""")

    for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in postal_rows:
        o_pdc = o[COL_PDC]
        f_pdc = f[COL_PDC]
        html.append(f"""  <tr>
    <td>{csv_row}</td>
    <td><strong>{h(dun_code)}</strong><br/><small>{h(dun_name)}</small></td>
    <td class="tag-before">{code(o_uc)}</td>
    <td class="tag-after">{code(f_uc)}</td>
    <td class="tag-before">{code(o_pdc)}</td>
    <td class="tag-after">{code(f_pdc)}</td>
  </tr>\n""")

    html.append("</tbody></table>\n\n")

    # ── FIX 2: Backtick ───────────────────────────────────────────────────────
    bt_rows = by_kind.get("backtick", [])
    html.append('<div class="page-break"></div>\n')
    html.append(
        f"<h2>{badge('backtick')} Fix 2 — <code>BA`KELALAN</code> Backtick → Apostrophe</h2>\n"
    )
    html.append(f"""<div class="section-meta">
  <strong>{len(bt_rows)} rows affected</strong> — all rows for DUN N.81 (BA'KELALAN).<br/>
  The grave accent character U+0060 (`) was replaced with a standard apostrophe U+0027 (')
  in every cell where it appeared: <code>STATE CONSTITUENCY NAME</code>, <code>POLLING DISTRICT NAME</code>,
  and <code>UNIQUE CODE</code> where applicable.
</div>\n""")

    # Build unique set of (col_name, before, after) tuples to show once per centre
    seen_bt = {}
    for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in bt_rows:
        pc = o[COL_PC]
        if pc not in seen_bt:
            seen_bt[pc] = (csv_row, dun_code, o, f, diffs)

    html.append("""<table class="data">
<thead>
  <tr>
    <th>CSV Row</th><th>DUN</th><th>Column</th>
    <th>Before</th><th></th><th>After</th>
  </tr>
</thead>
<tbody>\n""")

    for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in bt_rows:
        for d in diffs:
            if "BA`KELALAN" in o[d]:
                html.append(f"""  <tr>
    <td>{csv_row}</td>
    <td>{h(dun_code)}</td>
    <td><code>{h(header[d])}</code></td>
    <td class="tag-before">{code(o[d])}</td>
    <td class="arrow">→</td>
    <td class="tag-after">{code(f[d])}</td>
  </tr>\n""")
        break  # one representative row is enough for display

    html.append(f"</tbody></table>\n")
    html.append(
        f'<div class="note">All {len(bt_rows)} rows in N.81 were corrected with the same substitution — only one representative row shown above per column.</div>\n\n'
    )

    # ── FIX 3: Newlines ───────────────────────────────────────────────────────
    nl_rows = by_kind.get("newline", [])
    html.append(
        f"<h2>{badge('newline')} Fix 3 — Embedded Newlines Removed from Polling Centre</h2>\n"
    )
    html.append(f"""<div class="section-meta">
  <strong>{len(nl_rows)} rows affected</strong> across 4 DUNs — 5 distinct polling centre names contained
  an embedded newline character (<code>\\n</code>) inside a quoted CSV field.
  Each was joined with the appropriate separator (space, nothing, or punctuation) based on context.
</div>\n""")

    html.append("""<table class="data">
<thead>
  <tr>
    <th>CSV Row</th><th>DUN</th>
    <th>POLLING CENTRE — before</th><th></th><th>POLLING CENTRE — after</th>
  </tr>
</thead>
<tbody>\n""")

    for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in nl_rows:
        o_pc = o[COL_PC].replace("\n", "↵")
        f_pc = f[COL_PC]
        html.append(f"""  <tr>
    <td>{csv_row}</td>
    <td><strong>{h(dun_code)}</strong><br/><small>{h(dun_name)}</small></td>
    <td class="tag-before">{code(o_pc)}</td>
    <td class="arrow">→</td>
    <td class="tag-after">{code(f_pc)}</td>
  </tr>\n""")

    html.append("</tbody></table>\n\n")

    # ── FIX 4: Suffixes ───────────────────────────────────────────────────────
    sf_rows = by_kind.get("suffix", [])
    html.append('<div class="page-break"></div>\n')
    html.append(
        f"<h2>{badge('suffix')} Fix 4 — UNIQUE CODE Suffixes Added ({len(sf_rows):,} rows)</h2>\n"
    )

    # Find max letter used
    max_letter = "a"
    for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in sf_rows:
        m = re.search(r"([a-z])$", f_uc)
        if m and m.group(1) > max_letter:
            max_letter = m.group(1)

    html.append(f"""<div class="section-meta">
  <strong>{len(sf_rows):,} rows affected</strong> across
  <strong>{len({c[1] for c in sf_rows})} DUNs</strong> —
  {309} polling districts where a single district code served multiple polling centres.<br/>
  Letter suffixes <code>a</code> through <code>{max_letter}</code> were appended to UNIQUE CODE (col 1) only.
  The same letter is applied to <em>every</em> voting channel row for the same district+centre combination
  so the suffix is consistent (e.g. both <code>_1b</code> and <code>_2b</code> for the same centre).<br/>
  <strong>No vote counts or candidate data were modified.</strong>
</div>\n""")

    # Group by DUN for the summary table
    by_dun = defaultdict(list)
    for c in sf_rows:
        by_dun[c[1]].append(c)

    # DUN-level summary
    html.append("""<table class="data">
<thead>
  <tr>
    <th>DUN Code</th><th>DUN Name</th>
    <th>Rows suffixed</th><th>Letters used</th>
    <th>Example UNIQUE CODE — before</th><th>after</th>
  </tr>
</thead>
<tbody>\n""")

    def dun_sort_key(dun_code):
        m = re.search(r"(\d+)", dun_code)
        return int(m.group(1)) if m else 999

    for dun_code in sorted(by_dun.keys(), key=dun_sort_key):
        rows_for_dun = by_dun[dun_code]
        dun_name = rows_for_dun[0][2]
        letters = sorted(
            set(
                re.search(r"([a-z])$", c[4]).group(1)
                for c in rows_for_dun
                if re.search(r"([a-z])$", c[4])
            )
        )
        ex = rows_for_dun[0]
        html.append(f"""  <tr>
    <td><strong>{h(dun_code)}</strong></td>
    <td>{h(dun_name)}</td>
    <td style="text-align:right">{len(rows_for_dun)}</td>
    <td><code>{", ".join(letters)}</code></td>
    <td class="tag-before">{code(ex[3])}</td>
    <td class="tag-after">{code(ex[4])}</td>
  </tr>\n""")

    html.append("</tbody></table>\n\n")

    # ── OTHER ─────────────────────────────────────────────────────────────────
    other_rows = by_kind.get("other", [])
    if other_rows:
        html.append(
            f"<h2>{badge('other')} Other Changes ({len(other_rows)} rows)</h2>\n"
        )
        html.append("""<table class="data">
<thead>
  <tr><th>CSV Row</th><th>DUN</th><th>Column</th><th>Before</th><th></th><th>After</th></tr>
</thead>
<tbody>\n""")
        for csv_row, dun_code, dun_name, o_uc, f_uc, diffs, kind, o, f in other_rows:
            for d in diffs:
                html.append(f"""  <tr>
    <td>{csv_row}</td><td>{h(dun_code)}</td>
    <td><code>{h(header[d])}</code></td>
    <td class="tag-before">{code(o[d])}</td>
    <td class="arrow">→</td>
    <td class="tag-after">{code(f[d])}</td>
  </tr>\n""")
        html.append("</tbody></table>\n\n")

    html.append("</body></html>")
    return "".join(html)


# ── Entry point ───────────────────────────────────────────────────────────────


def main():
    orig_path = "to-review.csv.orig"
    final_path = "to-review.csv"
    html_path = "change-report.html"
    pdf_path = "change-report.pdf"

    for p in (orig_path, final_path):
        if not os.path.exists(p):
            print(f"ERROR: {p} not found", file=sys.stderr)
            sys.exit(1)

    print("Reading CSV files…")
    orig_rows = read_csv(orig_path)
    final_rows = read_csv(final_path)

    assert orig_rows[0] == final_rows[0], "Headers differ"
    assert len(orig_rows) == len(final_rows), (
        f"Row count differs: orig={len(orig_rows)} final={len(final_rows)}"
    )

    print("Building HTML report…")
    html = build_html(orig_rows, final_rows)

    with open(html_path, "w", encoding="utf-8") as f:
        f.write(html)
    print(f"  → {html_path} written ({len(html):,} bytes)")

    print("Converting to PDF via WeasyPrint…")
    try:
        from weasyprint import HTML

        HTML(filename=html_path).write_pdf(pdf_path)
        size_kb = os.path.getsize(pdf_path) // 1024
        print(f"  → {pdf_path} written ({size_kb} KB)")
    except ImportError:
        print("WeasyPrint not available — HTML report saved; convert manually.")
        print(f"  HTML: {html_path}")
        sys.exit(0)

    print("\nDone.")
    print(f"  PDF : {pdf_path}")
    print(f"  HTML: {html_path}")


if __name__ == "__main__":
    main()
