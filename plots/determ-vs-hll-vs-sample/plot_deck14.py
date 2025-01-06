import json
from datetime import datetime
import matplotlib.pyplot as plt

base = '/Users/jp/Workspace/facu/hll-truco/hll-truco/plots/determ-vs-hll-vs-sample/data/deck14'
file_path_axion = base + '/axiom.data'
file_path_clark = base + '/clark.data'
file_path_hll_p6 = base + '/hll-p6.data'
file_path_hll_p10 = base + '/hll-p10.data'

xs_axiom, ys_axiom = [], []
xs_clark, ys_clark = [], []
xs_hll_p6, ys_hll_p6 = [], []
xs_hll_p10, ys_hll_p10 = [], []

def load_data(file_path):
    xs, ys = [], []
    with open(file_path, 'r') as file:
        for line in file:
            data = json.loads(line)
            xs += [data['time']]
            if data['msg'] == 'START': ys += [0]
            elif data['msg'] == 'REPORT': ys += [data['estimate']]
            elif data['msg'] == 'RESULTS': ys += [data['finalEstimate']]
    return xs, ys

def load_data_hll(file_path):
    xs, ys = [], []
    with open(file_path, 'r') as file:
        for line in file:
            data = json.loads(line)
            if data['msg'] not in ['START', 'REPORT', 'RESULTS']: continue
            xs += [data['time']]
            if data['msg'] == 'START': ys += [0]
            elif data['msg'] == 'REPORT': ys += [float(data['estimate'])]
            elif data['msg'] == 'RESULTS': ys += [float(data['finalEstimate'])]
    return xs, ys

def convert_to_seconds(xs):
    start_time = datetime.fromisoformat(xs[0])
    return [(datetime.fromisoformat(x) - start_time).total_seconds() for x in xs]

def cap_at(xs, ys, max_secods):
    xs = [x for x in xs if x <= max_secods]
    return xs, ys[:len(xs)]

xs_axiom, ys_axiom = load_data(file_path_axion)
xs_clark, ys_clark = load_data(file_path_clark)
xs_hll_p6, ys_hll_p6 = load_data_hll(file_path_hll_p6)
xs_hll_p10, ys_hll_p10 = load_data_hll(file_path_hll_p10)

xs_axiom   = convert_to_seconds(xs_axiom)
xs_clark   = convert_to_seconds(xs_clark)
xs_hll_p6 = convert_to_seconds(xs_hll_p6)
xs_hll_p10 = convert_to_seconds(xs_hll_p10)

n = 800
xs_axiom, ys_axiom = cap_at(xs_axiom, ys_axiom, n)
xs_clark, ys_clark = cap_at(xs_clark, ys_clark, n)
xs_hll_p6, ys_hll_p6 = cap_at(xs_hll_p6, ys_hll_p6, n)
xs_hll_p10, ys_hll_p10 = cap_at(xs_hll_p10, ys_hll_p10, n)

# Create the plot
plt.figure(figsize=(10, 5))
plt.plot(xs_axiom, ys_axiom, label='Axiom')
plt.plot(xs_clark, ys_clark, label='Clark')
plt.plot(xs_hll_p6, ys_hll_p6, label='HLL precision 6')
plt.plot(xs_hll_p10, ys_hll_p10, label='HLL precision 10')

# horizontal line with the value of the last element
plt.axhline(y=248_732, color='r', linestyle='--', label='Real value')

# Label the axes
plt.xlabel('Time (seconds)')
plt.ylabel('Infosets count')

# Title of the plot
plt.title('Time vs Infoset count (single thread)')

# Show the plot
plt.legend()
plt.show()

