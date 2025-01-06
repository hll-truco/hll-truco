import numpy as np

# Your input distributions
dist_2p_marked124366 = {"0":3082,"1":32459,"2":32417,"3":24470,"4":15820,"5":8919,"6":4708,"7":1876,"8":447,"9":168}
dist_2p_marked223858 = {"0":3082,"1":39386,"2":60162,"3":53047,"4":36412,"5":18923,"6":8037,"7":3524,"8":975,"9":310}

# Convert to numpy arrays and ensure same ordering
keys = sorted(dist_2p_marked124366.keys())
p = np.array([dist_2p_marked124366[k] for k in keys], dtype=float)
q = np.array([dist_2p_marked223858[k] for k in keys], dtype=float)

# Normalize to create probability distributions
p = p / np.sum(p)
q = q / np.sum(q)

# Verify that both sum to 1
assert np.isclose(np.sum(p), 1.0)
assert np.isclose(np.sum(1), 1.0)

# Calculate KL divergence
# We add a small epsilon to avoid log(0)
epsilon = 1e-10
kl_div = np.sum(p * np.log((p + epsilon) / (q + epsilon)))
print(f"KL(P||Q) = {kl_div:.6f}")

# we also calculate the KL divergence using scipy to check our result
from scipy.stats import entropy
kl_div_scipy = entropy(p, q)
print(f"KL(P||Q) = {kl_div_scipy:.6f}")


