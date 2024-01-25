import argparse
import os
import re
import json

parser = argparse.ArgumentParser(description='List all .out files in a directory')
parser.add_argument('-d', '--directory', type=str, required=True, help='Directory to search for .out files')
parser.add_argument('-o', '--output', type=str, default=None, help='Path to output JSON file')
args = parser.parse_args()

result_dict = {}

for root, dirs, files in os.walk(args.directory):
    for file in files:
        if file.endswith('.out'):
            print(file)
            file_path = os.path.join(root, file)
            with open(file_path, 'r') as f:
                lines = f.readlines()
                done = []
                count = []
                for line in lines:
                    match_done = re.search(r'\tdone: map\[0:(\d+)', line)
                    match_count = re.search(r'\tcount:\s(\d+)', line)
                    if match_done:
                        done.append(int(match_done.group(1)))
                    elif match_count:
                        count.append(int(match_count.group(1)))

                result_dict[file] = {
                    "done": done,
                    "count": count
                }

if args.output is None:
    output_path = os.path.join(args.directory, 'result.json')
else:
    output_path = args.output

with open(output_path, 'w') as f:
    json.dump(result_dict, f, indent=2)

print(f"\nResult saved to {output_path}")
