import random
from ks import num_samples, compute_sha256_hash

with open("1M_py_sha256_random.log2", "x+") as f:
    for i in range(num_samples):

        # x = random.randint(0, num_samples) # random
        x = i # serie

        h = compute_sha256_hash(x)
        f.write(h + "\n")

print("done")