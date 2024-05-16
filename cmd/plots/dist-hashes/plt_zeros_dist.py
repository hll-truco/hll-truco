import matplotlib.pyplot as plt
import json

def load_dist(filepath:str) -> dict:
    with open(filepath) as f:
        d = json.loads(f.read())
        d = {int(k):v for k,v in d.items()}
        return d

sha256 = load_dist("cmd/dist-hashes/sha256dist-1B.json")
sha512 = load_dist("cmd/dist-hashes/sha512dist-1B.json")
blake2b_512 = load_dist("cmd/dist-hashes/blake2b512dist-1B.json")
md5 = load_dist("cmd/dist-hashes/md5dist-1B.json")
sha3_1024 = load_dist("cmd/dist-hashes/sha3-1024-dist-1B.json")
# newsha3_1024 = load_dist("cmd/dist-hashes/newsha3-1024-dist-1B.json")

sha3_1024_4B = load_dist("cmd/dist-hashes/sha3-1024-dist-4B.json")

# Extract keys and values from the dictionary
n = 30
keys = list(sorted(sha256.keys()))[:n]
values = [sha256[k] for k in keys][:n]

# Create a figure with two subplots
fig, (ax0, ax1) = plt.subplots(1, 2, figsize=(12, 5))

# Plot the histogram on the left subplot
ax0.bar(keys, values, label="bar")
ax0.set_yscale('log', base=10)
ax0.set_xlabel('Number of consecutive zeros')
ax0.set_ylabel('Freq.')
ax0.set_title('Dist. of Consecutive Zeros in SHA-256 Hashes\nfor the First One Billion Integers')

def double_plot(ax, data, label):
    line_plot = ax1.plot(data, alpha=0.3)
    line_color = line_plot[0].get_color()
    ax1.plot(data, label=label, marker='o', linestyle='None', color=line_color, markersize=3)

double_plot(ax1, [md5[k] for k in sorted(md5.keys())[:n]], "md5")
double_plot(ax1, [sha256[k] for k in sorted(sha256.keys())[:n]], "sha256")
double_plot(ax1, [sha512[k] for k in sorted(sha512.keys())[:n]], "sha512")
double_plot(ax1, [blake2b_512[k] for k in sorted(blake2b_512.keys())[:n]], "blake2b512")
double_plot(ax1, [sha3_1024[k] for k in sorted(sha3_1024.keys())[:n]], "sha3-shake256_1024")
# double_plot(ax1, [newsha3_1024[k] for k in sorted(newsha3_1024.keys())[:n]], "newsha3_1024")

# Plot the line plots on the right subplot
# ax1.plot([sha256[k] for k in sorted(sha256.keys())[:n]], label="sha256", marker='o')
# ax1.plot([sha512[k] for k in sorted(sha512.keys())[:n]], label="sha512")
# ax1.plot([blake2b_512[k] for k in sorted(blake2b_512.keys())[:n]], label="blake2b_512")
# ax1.plot([sha3_1024[k] for k in sorted(sha3_1024.keys())[:n]], label="sha3_1024")
# ax1.plot([newsha3_1024[k] for k in sorted(newsha3_1024.keys())[:n]], label="newsha3_1024")
# ax1.plot([md5[k] for k in sorted(md5.keys())[:n]], label="md5")

# ax1.plot([sha3_1024_4B[k] for k in sorted(sha3_1024_4B.keys())[:n]], label="sha3_1024_4B")
ax1.set_yscale('log', base=10)
ax1.set_xlabel('Number of consecutive zeros')
ax1.set_ylabel('Freq.')
ax1.set_title('Dist. of Consecutive Zeros in Multiple Hash Functions\nfor the First One Billion Integers')
ax1.legend()

# Adjust the spacing between subplots
# plt.subplots_adjust(wspace=0.4)
plt.tight_layout()

# Display the plot
plt.show()