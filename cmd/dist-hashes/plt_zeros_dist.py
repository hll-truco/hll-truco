import matplotlib.pyplot as plt
import json

def load_dist(filepath:str) -> dict:
    with open(filepath) as f:
        d = json.loads(f.read())
        d = {int(k):v for k,v in d.items()}
        return d

sha256 = load_dist("sha256dist-1B.json")
sha512 = load_dist("sha512dist-1B.json")
sha3_1024 = load_dist("sha3-1024-dist-1B.json")

sha3_1024_4B = load_dist("sha3-1024-dist-4B.json")

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
ax0.set_title('Dist. of consecutive zeros in sha256 hashes of the first 1B integers')

# Plot the line plots on the right subplot
ax1.plot([sha256[k] for k in sorted(sha256.keys())[:n]], label="sha256")
ax1.plot([sha512[k] for k in sorted(sha512.keys())[:n]], label="sha512")
ax1.plot([sha3_1024[k] for k in sorted(sha3_1024.keys())[:n]], label="sha3_1024")
# ax1.plot([sha3_1024_4B[k] for k in sorted(sha3_1024_4B.keys())[:n]], label="sha3_1024_4B")
ax1.set_yscale('log', base=10)
ax1.set_xlabel('Number of consecutive zeros')
ax1.set_ylabel('Freq.')
ax1.set_title('Dist. of consecutive zeros in multiple hashes of the first 1B integers')
ax1.legend()

# Adjust the spacing between subplots
# plt.subplots_adjust(wspace=0.4)
plt.tight_layout()

# Display the plot
plt.show()