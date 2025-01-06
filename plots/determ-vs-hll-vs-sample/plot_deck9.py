import json
from datetime import datetime
import matplotlib.pyplot as plt

base = '/Users/jp/Workspace/facu/hll-truco/hll-truco/plots/determ-vs-hll-vs-sample/data/deck9'
file_path_det = base + '/deterministic.data'
file_path_axion = base + '/axiom.data'
file_path_clark = base + '/clark.data'
file_path_hll_p10 = base + '/hll-p10.data'

xs_det, ys_det = [], []
xs_axiom, ys_axiom = [], []
xs_clark, ys_clark = [], []
xs_hll_p10, ys_hll_p10 = [], []

with open(file_path_det, 'r') as file:
    for line in file:
        data = json.loads(line)
        xs_det += [data['time']]
        if data['msg'] == 'START': ys_det += [0]
        elif data['msg'] == 'REPORT': ys_det += [data['count']]
        elif data['msg'] == 'RESULTS': ys_det += [data['infosets:']]

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

xs_axiom, ys_axiom = load_data(file_path_axion)
xs_clark, ys_clark = load_data(file_path_clark)
xs_hll_p10, ys_hll_p10 = load_data_hll(file_path_hll_p10)

xs_det     = convert_to_seconds(xs_det)
xs_axiom   = convert_to_seconds(xs_axiom)
xs_clark   = convert_to_seconds(xs_clark)
xs_hll_p10 = convert_to_seconds(xs_hll_p10)

# Create the plot
plt.figure(figsize=(10, 5))
plt.plot(xs_det, ys_det, label='Deterministic')
plt.plot(xs_axiom, ys_axiom, label='Axiom')
plt.plot(xs_clark, ys_clark, label='Clark')
plt.plot(xs_hll_p10, ys_hll_p10, label='HLL')

# horizontal line with the value of the last element
plt.axhline(y=ys_det[-1], color='r', linestyle='--', label='Real value')

# Label the axes
plt.xlabel('Time (seconds)')
plt.ylabel('Infosets count')

# Title of the plot
plt.title('Time vs Infoset count (single thread)')

# Show the plot
plt.legend()
plt.show()

