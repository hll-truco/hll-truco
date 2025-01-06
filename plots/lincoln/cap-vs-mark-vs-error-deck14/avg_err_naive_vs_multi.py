import json

# Path to the JSON file
base = '/Users/jp/Workspace/facu/hll-truco/hll-truco/plots/lincoln/cap-vs-mark-vs-error-deck14/data'
naive_lincoln_path = '/naive_lincoln.json'
multi_lincoln_path = '/multi_lincoln.json'

with open(base + naive_lincoln_path, 'r') as file: data_naive = json.load(file)
with open(base + multi_lincoln_path, 'r') as file: data_multi = json.load(file)

correct_N = 248732

def avg_error(data):
    total_error = 0
    for entry in data:
        relative_error = abs(entry["N"] - correct_N) / correct_N * 100
        total_error += relative_error
    return total_error / len(data)

avg_err_naive = avg_error(data_naive)
avg_err_multi = avg_error(data_multi)

print(f"{avg_err_naive=} {avg_err_multi=}")

"""
got:
    avg_err_naive=35.35208760652209
    avg_err_multi=23.174931129476573
"""
