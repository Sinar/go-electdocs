#!/usr/bin/env python3
"""
verify_ballots.py  —  Three-part integrity check for Sarawak DUN 2021 data.

PART 1  Corruption guard
        Compares to-review.csv.orig vs to-review.csv row-by-row for
        TOTAL BALLOTS ISSUED.  Confirms the fix scripts (UNIQUE CODE
        suffixes, postal DMKOD, BA'KELALAN, newlines) did not alter any
        vote or ballot figure.

PART 2  Valid-vote cross-check against raw-candidate.csv
        For every DUN, sums the 'ju' (votes received) field across all
        candidates in raw-candidate.csv and compares the result against
        the sum of TOTAL VALID VOTES from to-review.csv.
        raw-candidate.csv comes directly from the EC API and is the most
        machine-readable authoritative source; this check is not affected
        by OCR quality.

PART 3  Registered-voter sanity check
        Parses JUMLAH PEMILIH (registered voters) from each
        results/Sarawak-N.XX.md header and confirms that
        TOTAL BALLOTS ISSUED (CSV) <= registered voters per DUN.

Usage (from the sarawak-dun-2021 directory):
    python3 verify_ballots.py
"""

import csv
import os
import re
from collections import defaultdict

# ---------------------------------------------------------------------------
# File paths
# ---------------------------------------------------------------------------
CSV_FINAL = "to-review.csv"
CSV_ORIG = "to-review.csv.orig"
RAW_CAND = "raw-candidate.csv"
RAW_DUN = "raw-dun.csv"
RESULTS_DIR = "results"

# ---------------------------------------------------------------------------
# Column indices in to-review.csv  (0-based)
# ---------------------------------------------------------------------------
COL_DUN_CODE = 5  # STATE CONSTITUENCY CODE  e.g. N.01
COL_DUN_NAME = 6  # STATE CONSTITUENCY NAME  e.g. OPAR
COL_BALLOTS = 11  # TOTAL BALLOTS ISSUED
COL_VALID = 62  # TOTAL VALID VOTES
COL_REJECTED = 63  # TOTAL REJECTED VOTES
COL_UNRETURNED = 64  # TOTAL UNRETURNED BALLOTS

# ---------------------------------------------------------------------------
# Column indices in raw-candidate.csv  (0-based)
# header: id,t,jp,pid,s,kid,kt,i,st,mi,nc,ju,mj,ut
# ---------------------------------------------------------------------------
RC_NAME = 1  # candidate name
RC_KID = 5  # DUN numeric ID (maps to raw-dun col 0)
RC_JU = 11  # votes received

# ---------------------------------------------------------------------------
# Column indices in raw-dun.csv  (no header, 0-based)
# ---------------------------------------------------------------------------
RD_ID = 0  # DUN numeric ID
RD_CODE = 2  # DUN code  e.g. N.01
RD_NAME = 3  # DUN name  e.g. OPAR
RD_REG = 5  # registered voters

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------


def safe_int(s):
    """Parse integer from possibly comma-formatted string; None on failure."""
    cleaned = re.sub(r"[,\s]", "", str(s).strip())
    return int(cleaned) if cleaned.lstrip("-").isdigit() and cleaned else None


def dun_sort_key(code):
    m = re.search(r"(\d+)", str(code))
    return int(m.group(1)) if m else 0


SEP = "=" * 110
HSEP = "-" * 110


def section(title):
    print(f"\n{SEP}")
    print(f"  {title}")
    print(f"{SEP}\n")


# ---------------------------------------------------------------------------
# CSV aggregation helpers
# ---------------------------------------------------------------------------


def aggregate_review_csv(path):
    """
    Read a to-review.csv (or .orig) and return:
        { dun_code: { name, ballots, valid, rejected, unreturned, rows } }
    """
    agg = defaultdict(
        lambda: dict(name="", ballots=0, valid=0, rejected=0, unreturned=0, rows=0)
    )
    with open(path, newline="", encoding="utf-8") as fh:
        reader = csv.reader(fh)
        next(reader)  # skip header
        for row in reader:
            if len(row) <= COL_BALLOTS:
                continue
            code = row[COL_DUN_CODE].strip()
            if not code:
                continue
            d = agg[code]
            d["name"] = row[COL_DUN_NAME].strip()
            d["rows"] += 1
            for key, col in (
                ("ballots", COL_BALLOTS),
                ("valid", COL_VALID),
                ("rejected", COL_REJECTED),
                ("unreturned", COL_UNRETURNED),
            ):
                if len(row) > col:
                    v = safe_int(row[col])
                    if v is not None:
                        d[key] += v
    return agg


# ---------------------------------------------------------------------------
# raw-dun.csv helpers
# ---------------------------------------------------------------------------


def load_raw_dun(path):
    """
    Returns two dicts:
        id_to_code  { '19201': 'N.01', … }
        code_to_reg { 'N.01': 11436,   … }
    """
    id_to_code = {}
    code_to_reg = {}
    with open(path, newline="", encoding="utf-8") as fh:
        for row in csv.reader(fh):
            if len(row) <= RD_REG:
                continue
            dun_id = row[RD_ID].strip()
            dun_code = row[RD_CODE].strip()
            dun_reg = safe_int(row[RD_REG])
            if dun_id and dun_code:
                id_to_code[dun_id] = dun_code
            if dun_code and dun_reg is not None:
                code_to_reg[dun_code] = dun_reg
    return id_to_code, code_to_reg


# ---------------------------------------------------------------------------
# raw-candidate.csv helpers
# ---------------------------------------------------------------------------


def load_raw_candidate_votes(path, id_to_code):
    """
    Returns { dun_code: total_votes } by summing 'ju' across all
    candidates belonging to that DUN.
    """
    votes = defaultdict(int)
    skipped = 0
    with open(path, newline="", encoding="utf-8") as fh:
        reader = csv.reader(fh)
        next(reader)  # skip header
        for row in reader:
            if len(row) <= RC_JU:
                continue
            kid = row[RC_KID].strip()
            code = id_to_code.get(kid)
            if code is None:
                skipped += 1
                continue
            ju = safe_int(row[RC_JU])
            if ju is not None:
                votes[code] += ju
    return dict(votes), skipped


# ---------------------------------------------------------------------------
# Markdown: JUMLAH PEMILIH (registered voters)
# ---------------------------------------------------------------------------


def extract_registered_voters(dun_code):
    """Parse 'JUMLAH PEMILIH: X,XXX' from the DUN's markdown file header."""
    m = re.search(r"(\d+)", dun_code)
    if not m:
        return None
    fname = f"Sarawak-N.{int(m.group(1)):02d}.md"
    path = os.path.join(RESULTS_DIR, fname)
    if not os.path.exists(path):
        return None
    with open(path, encoding="utf-8") as fh:
        for line in fh:
            hit = re.search(r"JUMLAH\s+PEMILIH\s*:\s*([\d,]+)", line, re.IGNORECASE)
            if hit:
                return safe_int(hit.group(1))
    return None


# ---------------------------------------------------------------------------
# PART 1 — Corruption check
# ---------------------------------------------------------------------------


def part1_corruption_check():
    section("PART 1 — Corruption Check: to-review.csv.orig  vs  to-review.csv")
    print(
        "  Confirms fix scripts did NOT alter TOTAL BALLOTS ISSUED, TOTAL VALID VOTES,"
    )
    print("  TOTAL REJECTED VOTES, or TOTAL UNRETURNED BALLOTS.\n")

    if not os.path.exists(CSV_ORIG):
        print(f"  ⚠️   {CSV_ORIG} not found — skipping corruption check.\n")
        return

    orig = aggregate_review_csv(CSV_ORIG)
    final = aggregate_review_csv(CSV_FINAL)
    codes = sorted(set(orig) | set(final), key=dun_sort_key)

    issues = []
    for code in codes:
        o = orig.get(code, {})
        f = final.get(code, {})
        for field in ("ballots", "valid", "rejected", "unreturned"):
            ov = o.get(field, 0)
            fv = f.get(field, 0)
            if ov != fv:
                issues.append((code, field, ov, fv))

    if issues:
        print(f"  ❌  DATA CORRUPTION DETECTED — {len(issues)} field(s) changed:\n")
        print(f"  {'DUN':<8} {'Field':<22} {'Orig':>12} {'Final':>12} {'Diff':>10}")
        print(f"  {HSEP[2:70]}")
        for code, field, ov, fv in issues:
            print(f"  {code:<8} {field:<22} {ov:>12,} {fv:>12,} {fv - ov:>+10,}")
    else:
        total_orig = sum(v["ballots"] for v in orig.values())
        total_final = sum(v["ballots"] for v in final.values())
        print(f"  ✅  PASS — no corruption detected across all {len(codes)} DUNs.\n")
        print(f"  {'Field':<30} {'Orig total':>14} {'Final total':>14}")
        print(f"  {HSEP[2:60]}")
        for field, label in (
            ("ballots", "TOTAL BALLOTS ISSUED"),
            ("valid", "TOTAL VALID VOTES"),
            ("rejected", "TOTAL REJECTED VOTES"),
            ("unreturned", "TOTAL UNRETURNED BALLOTS"),
        ):
            ot = sum(v.get(field, 0) for v in orig.values())
            ft = sum(v.get(field, 0) for v in final.values())
            match = "✅" if ot == ft else "❌"
            print(f"  {match} {label:<28} {ot:>14,} {ft:>14,}")


# ---------------------------------------------------------------------------
# PART 2 — Valid-vote cross-check via raw-candidate.csv
# ---------------------------------------------------------------------------


def part2_valid_votes_crosscheck():
    section("PART 2 — Valid-Vote Cross-Check: to-review.csv  vs  raw-candidate.csv")
    print("  For each DUN, compares SUM(TOTAL VALID VOTES) from to-review.csv against")
    print("  SUM(ju) from raw-candidate.csv (official EC candidate vote totals).")
    print(
        "  This check is unaffected by OCR quality — raw-candidate.csv is structured EC data.\n"
    )

    if not os.path.exists(RAW_DUN):
        print(f"  ⚠️   {RAW_DUN} not found — cannot map candidate DUN IDs.\n")
        return
    if not os.path.exists(RAW_CAND):
        print(f"  ⚠️   {RAW_CAND} not found.\n")
        return

    id_to_code, _ = load_raw_dun(RAW_DUN)
    rc_votes, skipped = load_raw_candidate_votes(RAW_CAND, id_to_code)
    csv_agg = aggregate_review_csv(CSV_FINAL)

    all_codes = sorted(set(csv_agg) | set(rc_votes), key=dun_sort_key)

    print(
        f"  {'DUN':<8} {'NAME':<22} {'CSV Valid Votes':>16} {'EC Cand Votes':>15} {'Diff':>8}  Status"
    )
    print(f"  {HSEP[2:]}")

    mismatches = []
    exact = 0
    near = []  # |diff| <= 10
    possible = []  # |diff| <= 100

    for code in all_codes:
        csv_v = csv_agg.get(code, {}).get("valid", 0)
        ec_v = rc_votes.get(code, None)
        name = csv_agg.get(code, {}).get("name", rc_votes.get(code, "?"))

        if ec_v is None:
            status = "⚠️  NO RC DATA"
            note = ""
        elif csv_v == ec_v:
            status = "✅ MATCH"
            exact += 1
            note = ""
        elif abs(csv_v - ec_v) <= 10:
            status = "🔶 NEAR"
            near.append((code, name, csv_v, ec_v))
            note = f"diff={csv_v - ec_v:+d}"
        elif abs(csv_v - ec_v) <= 100:
            status = "🟡 POSSIBLE"
            possible.append((code, name, csv_v, ec_v))
            note = f"diff={csv_v - ec_v:+d}"
        else:
            status = "❌ MISMATCH"
            mismatches.append((code, name, csv_v, ec_v))
            note = f"diff={csv_v - ec_v:+d}"

        ec_s = f"{ec_v:>15,}" if ec_v is not None else f"{'N/A':>15}"
        print(
            f"  {code:<8} {name:<22} {csv_v:>16,} {ec_s} "
            f"{(csv_v - ec_v):>+8,}  {status}"
            if ec_v is not None
            else f"  {code:<8} {name:<22} {csv_v:>16,} {'N/A':>15} {'':>8}  {status}"
        )

    # Totals
    total_csv = sum(csv_agg[c]["valid"] for c in all_codes if c in csv_agg)
    total_ec = sum(rc_votes[c] for c in all_codes if c in rc_votes)
    print(f"  {HSEP[2:]}")
    print(
        f"  {'GRAND TOTAL':<30} {total_csv:>16,} {total_ec:>15,} {total_csv - total_ec:>+8,}"
    )

    # Summary
    print(f"\n  {'─' * 60}")
    print(f"  Summary")
    print(f"  {'─' * 60}")
    print(f"  ✅  Exact match          : {exact}")
    print(f"  🔶  Near match  (≤ 10)  : {len(near)}")
    print(f"  🟡  Possible    (≤ 100) : {len(possible)}")
    print(f"  ❌  Mismatch    (> 100) : {len(mismatches)}")

    if skipped:
        print(f"\n  ℹ️   {skipped} candidate row(s) skipped (kid not in raw-dun.csv)")

    if mismatches:
        print(f"\n  ❌  MISMATCHES REQUIRING INVESTIGATION:")
        print(
            f"  {'DUN':<8} {'NAME':<22} {'CSV Valid':>12} {'EC Votes':>12} {'Diff':>10}"
        )
        print(f"  {HSEP[2:70]}")
        for code, name, cv, ev in mismatches:
            print(f"  {code:<8} {name:<22} {cv:>12,} {ev:>12,} {cv - ev:>+10,}")
    else:
        print(
            f"\n  ✅  No mismatches — all DUN valid-vote totals match raw-candidate.csv."
        )

    if near:
        print(f"\n  🔶  NEAR-MATCHES (|diff| ≤ 10) — likely rounding, spot-check:")
        for code, name, cv, ev in near:
            print(f"  {code:<8} {name:<22}  CSV={cv:,}  EC={ev:,}  diff={cv - ev:+d}")

    if possible:
        print(f"\n  🟡  POSSIBLE MATCHES (|diff| ≤ 100):")
        for code, name, cv, ev in possible:
            print(f"  {code:<8} {name:<22}  CSV={cv:,}  EC={ev:,}  diff={cv - ev:+d}")


# ---------------------------------------------------------------------------
# PART 3 — Registered-voter sanity check
# ---------------------------------------------------------------------------


def part3_registered_voters_sanity():
    section("PART 3 — Registered-Voter Sanity Check")
    print("  Verifies TOTAL BALLOTS ISSUED (CSV) <= JUMLAH PEMILIH (registered voters)")
    print(
        "  for each DUN.  Registered voters are parsed from results/Sarawak-N.XX.md headers.\n"
    )

    if not os.path.exists(RAW_DUN):
        print(
            f"  ⚠️   {RAW_DUN} not found — using markdown-only registered voter source.\n"
        )
        _, code_to_reg = {}, {}
    else:
        _, code_to_reg = load_raw_dun(RAW_DUN)

    csv_agg = aggregate_review_csv(CSV_FINAL)
    codes = sorted(csv_agg.keys(), key=dun_sort_key)

    print(
        f"  {'DUN':<8} {'NAME':<22} {'CSV Ballots':>12} {'Reg.Voters':>12} "
        f"{'Turnout':>8}  Status"
    )
    print(f"  {HSEP[2:]}")

    failures = []
    no_reg_data = []
    total_ballots = 0
    total_reg = 0

    for code in codes:
        d = csv_agg[code]
        ballots = d["ballots"]
        total_ballots += ballots

        # Prefer raw-dun.csv (structured), fall back to markdown OCR
        reg = code_to_reg.get(code)
        if reg is None:
            reg = extract_registered_voters(code)

        if reg is None:
            status = "⚠️  NO REG DATA"
            turnout_s = "N/A"
            no_reg_data.append(code)
        elif ballots > reg:
            status = f"❌ FAIL  ballots({ballots:,}) > registered({reg:,})"
            turnout_s = f"{ballots / reg * 100:.1f}%"
            failures.append((code, d["name"], ballots, reg))
        else:
            turnout_s = f"{ballots / reg * 100:.1f}%"
            if ballots / reg < 0.30:
                status = "⚠️  LOW TURNOUT (<30%)"
            elif ballots / reg > 0.95:
                status = "⚠️  HIGH TURNOUT (>95%)"
            else:
                status = "✅ OK"
            total_reg += reg

        reg_s = f"{reg:>12,}" if reg else f"{'N/A':>12}"
        print(
            f"  {code:<8} {d['name']:<22} {ballots:>12,} "
            f"{reg_s} {turnout_s:>8}  {status}"
        )

    print(f"  {HSEP[2:]}")
    if total_reg:
        overall_turnout = total_ballots / total_reg * 100
        print(
            f"  {'GRAND TOTAL':<30} {total_ballots:>12,} {total_reg:>12,} "
            f"{overall_turnout:>7.1f}%"
        )

    # Summary
    print(f"\n  {'─' * 60}")
    print(f"  Summary")
    print(f"  {'─' * 60}")
    if failures:
        print(
            f"\n  ❌  SANITY FAILURES ({len(failures)}) — ballots exceed registered voters:"
        )
        print(f"  {'DUN':<8} {'NAME':<22} {'Ballots':>12} {'Registered':>12}")
        print(f"  {HSEP[2:60]}")
        for code, name, b, r in failures:
            print(f"  {code:<8} {name:<22} {b:>12,} {r:>12,}")
    else:
        print(
            f"  ✅  PASS — no DUN has TOTAL BALLOTS ISSUED exceeding registered voters."
        )

    if no_reg_data:
        print(
            f"\n  ⚠️   Registered voter data unavailable for {len(no_reg_data)} DUN(s): "
            f"{', '.join(no_reg_data)}"
        )

    if total_reg:
        print(f"\n  Overall turnout  : {total_ballots / total_reg * 100:.1f}%")
        print(f"  Total ballots    : {total_ballots:,}")
        print(f"  Total registered : {total_reg:,}")


# ---------------------------------------------------------------------------
# Entry point
# ---------------------------------------------------------------------------


def main():
    print()
    print("╔══════════════════════════════════════════════════════════════╗")
    print("║  Sarawak DUN 2021 — TOTAL BALLOTS ISSUED Integrity Checks   ║")
    print("╚══════════════════════════════════════════════════════════════╝")

    for path in (CSV_FINAL,):
        if not os.path.exists(path):
            print(f"\nERROR: required file not found: {path}")
            raise SystemExit(1)

    part1_corruption_check()
    part2_valid_votes_crosscheck()
    part3_registered_voters_sanity()

    print(f"\n{SEP}")
    print("  Done.")
    print(f"{SEP}\n")


if __name__ == "__main__":
    main()
