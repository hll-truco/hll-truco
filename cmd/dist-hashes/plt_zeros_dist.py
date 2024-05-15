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

# sha3_1024_4B = load_dist("sha3-1024-dist-4B.json")

# Extract keys and values from the dictionary
keys = list(sha256.keys())
values = list(sha256.values())
n = 30

# Plot the histogram
# plt.bar(keys, values)

plt.plot([sha256[k] for k in sorted(sha256.keys())[:n]], label="sha256")
plt.plot([sha512[k] for k in sorted(sha512.keys())[:n]], label="sha512")
plt.plot([sha3_1024[k] for k in sorted(sha3_1024.keys())[:n]], label="sha3_1024")

# plt.plot([sha3_1024_4B[k] for k in sorted(sha3_1024_4B.keys())[:n]], label="sha3_1024_4B")

plt.legend()
plt.yscale('log', base=10)

plt.xlabel('Keys')
plt.ylabel('Values')
plt.title('Histogram from Dictionary')

# Display the plot
plt.show()