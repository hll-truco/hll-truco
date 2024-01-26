import argparse
import os
import re
import json

import sys
sys.path.append('cmd/python-common')
import parse_utils

parser = argparse.ArgumentParser(description='List all .out files in a directory')
parser.add_argument('-d', '--directory', type=str, required=True, help='Directory to search for .out files')
parser.add_argument('-o', '--output', type=str, default=None, help='Path to output JSON file')
args = parser.parse_args()

data = {}

for root, dirs, files in os.walk(args.directory):
    for file in files:
        if file.endswith('.out'):
            file_path = os.path.join(root, file)
            with open(file_path, 'r') as f:
                lines = f.readlines()
                done = []
                count = []
                first_commit = None
                last_commit = None
                for line in lines:
                    match_done = re.search(r'\tdone: map\[0:(\d+)', line)
                    match_count = re.search(r'\tcount:\s(\d+)', line)
                    match_date = re.search(r'^\d{4}/\d{2}/\d{2} \d{,2}:\d{,2}:\d{,2}', line)
                    if match_done: done.append(int(match_done.group(1)))
                    elif match_count: count.append(int(match_count.group(1)))
                    elif match_date:
                        last_commit = match_date.group(0)
                        if first_commit is None: first_commit = last_commit
                data[file] = {
                    "done": done,
                    "count": count,
                    "first_commit": first_commit,
                    "last_commit": last_commit,
                }

# post proc

# progress
for file,d in data.items():
    current = d['done'][-1]
    goal = 480480
    first_commit = parse_utils.parse_date(d['first_commit'])
    last_commit = parse_utils.parse_date(d['last_commit'])
    prog, delta, eta, eta_total = parse_utils.progress(current, goal, first_commit, last_commit)
    d["prog"] = prog
    d["delta"] = delta.total_seconds()
    d["eta"] = eta.total_seconds()
    d["eta_total"] = eta_total.total_seconds()
    print("{:<0}\t{:>20}\t{:>20}\t{:>20}".format(
        file,
        f"progress:{round(prog*100)}%",
        f"delta:{str(delta)[:str(delta).rfind(':')]}",
        f"eta_total:{str(eta_total)[:str(eta_total).rfind(':')]}"))

output_path = os.path.join(args.directory, 'result.json') \
    if args.output is None else args.output

with open(output_path, 'w') as f:
    json.dump(data, f, indent=2)

print(f"\nResult saved to {output_path}")
