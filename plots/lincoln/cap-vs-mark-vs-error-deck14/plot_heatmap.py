
possible_values = [2487, 12437, 24873, 49746, 74620, 99493, 124366, 149239, 174112, 198986, 223859]
correct_N = 248732

# load `data.json` from the same directory
import json
with open('/Users/jp/Workspace/facu/hll-truco/hll-truco/plots/lincoln/cap-vs-mark-vs-error-deck14/data/naive_lincoln_mazo=false.json', 'r') as f:
    data = json.load(f)

# make a color map, where `marked` is used on the x axis and `captured` is used on the y axis
# the color is determined by the error rate between `correct_N` and `N`

import numpy as np
import matplotlib.pyplot as plt

# Create meshgrid of possible values
marked_values = possible_values
captured_values = possible_values

# Initialize matrix with NaN values
error_matrix = np.full((len(captured_values), len(marked_values)), np.nan)

# Fill in the matrix with relative errors where we have data
for entry in data:
    marked_idx = marked_values.index(entry["marked"])
    captured_idx = captured_values.index(entry["captured"])
    relative_error = abs(entry["N"] - correct_N) / correct_N * 100
    error_matrix[captured_idx, marked_idx] = relative_error
    print(f"{captured_idx=} {marked_idx=} {relative_error=}")

# Create formatted tick labels with percentages
def format_label(value):
    percentage = (value / correct_N) * 100
    return f'{value}\n({round(percentage)}%)'

# Create the plot
plt.figure(figsize=(12, 8))
heatmap = plt.imshow(error_matrix, cmap='RdYlGn_r')
plt.gca().invert_yaxis()
plt.colorbar(heatmap, label='Relative Error (%)')

# Set labels
plt.xlabel('Marked Infosets')
plt.ylabel('Captured Infosets')
plt.title(f"Relative Error in Population Estimation using Lincoln-Petersen under a population of {correct_N}")

# Apply to both axes
plt.xticks(range(len(marked_values)), [format_label(x) for x in marked_values], rotation=45)
plt.yticks(range(len(captured_values)), [format_label(x) for x in captured_values])

# Add grid
plt.grid(False)

# Adjust layout to prevent label cutoff
plt.tight_layout()

for i in range(len(captured_values)):
    for j in range(len(marked_values)):
        if not np.isnan(error_matrix[i, j]):
            plt.text(j, i, f'{error_matrix[i, j]:.1f}%', ha='center', va='center')

plt.show()