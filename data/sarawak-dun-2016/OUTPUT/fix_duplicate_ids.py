#!/usr/bin/env uv run
# /// script
# requires-python = ">=3.11"
# dependencies = []
# ///

"""
Fix duplicate UNIQUE CODE IDs in Sarawak election CSV files.
Only modifies the first column (UNIQUE CODE) when duplicates are found.
"""

import csv
import os
from collections import defaultdict
from pathlib import Path


def is_data_row(row):
    """Check if row is a data row (not header, summary, or empty)."""
    if not row or len(row) < 10:
        return False

    unique_code = row[0].strip()
    polling_center = row[9].strip() if len(row) > 9 else ""

    # Skip header
    if unique_code == "UNIQUE CODE":
        return False

    # Skip empty lines
    if not unique_code:
        return False

    # Skip summary lines (no polling center)
    if not polling_center:
        return False

    return True


def find_duplicates(file_path):
    """Find all duplicate IDs and their polling centers."""
    with open(file_path, 'r', encoding='utf-8') as f:
        reader = csv.reader(f)
        rows = list(reader)

    # Map ID -> [(row_index, polling_center)]
    id_map = defaultdict(list)

    for idx, row in enumerate(rows):
        if is_data_row(row):
            unique_id = row[0].strip()
            polling_center = row[9].strip()
            id_map[unique_id].append((idx, polling_center))

    # Find IDs that appear more than once
    duplicates = {}
    for unique_id, occurrences in id_map.items():
        if len(occurrences) > 1:
            duplicates[unique_id] = occurrences

    return rows, duplicates


def fix_duplicates(rows, duplicates):
    """Fix duplicate IDs by adding suffixes based on polling centers."""
    changes = []

    for unique_id, occurrences in duplicates.items():
        # Get unique polling centers in order of first appearance
        seen_centers = {}
        center_to_suffix = {}
        suffix_idx = 0

        for row_idx, polling_center in occurrences:
            if polling_center not in center_to_suffix:
                suffix = chr(ord('a') + suffix_idx)
                center_to_suffix[polling_center] = suffix
                suffix_idx += 1

        # Apply fixes
        for row_idx, polling_center in occurrences:
            suffix = center_to_suffix[polling_center]
            new_id = unique_id + suffix
            old_id = rows[row_idx][0]
            rows[row_idx][0] = new_id

            changes.append({
                'old_id': old_id,
                'new_id': new_id,
                'polling_center': polling_center
            })

    return rows, changes


def process_file(file_path):
    """Process a single file and return changes made."""
    rows, duplicates = find_duplicates(file_path)

    if not duplicates:
        return []

    updated_rows, changes = fix_duplicates(rows, duplicates)

    # Write back to file
    with open(file_path, 'w', encoding='utf-8', newline='') as f:
        writer = csv.writer(f)
        writer.writerows(updated_rows)

    return changes


def main():
    """Process all N.01 to N.81 files."""
    base_dir = Path(__file__).parent

    all_changes = []

    for i in range(1, 82):
        file_name = f"Sarawak-N.{i:02d}.csv"
        file_path = base_dir / file_name

        if not file_path.exists():
            continue

        print(f"Processing {file_name}...")
        changes = process_file(file_path)

        if changes:
            for change in changes:
                all_changes.append({
                    'file': file_name,
                    'duplicated_id': change['old_id'],
                    'polling_center': change['polling_center'],
                    'fixed_id': change['new_id']
                })
            print(f"  Fixed {len(changes)} duplicate IDs")
        else:
            print(f"  No duplicates found")

    # Print summary table
    print("\n" + "="*100)
    print("SUMMARY OF CHANGES")
    print("="*100)
    print(f"{'File Name':<20} {'Duplicated ID':<35} {'Polling Center':<30} {'Fixed ID':<35}")
    print("-"*100)

    for change in all_changes:
        file_name = change['file']
        dup_id = change['duplicated_id'][:33] + "..." if len(change['duplicated_id']) > 35 else change['duplicated_id']
        center = change['polling_center'][:28] + "..." if len(change['polling_center']) > 30 else change['polling_center']
        fixed = change['fixed_id'][:33] + "..." if len(change['fixed_id']) > 35 else change['fixed_id']
        print(f"{file_name:<20} {dup_id:<35} {center:<30} {fixed:<35}")

    print("-"*100)
    print(f"Total changes: {len(all_changes)} across {len(set(c['file'] for c in all_changes))} files")


if __name__ == "__main__":
    main()
