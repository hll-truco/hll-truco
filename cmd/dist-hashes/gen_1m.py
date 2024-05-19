# import random
import hashlib
from ks import compute_any_hash

num_samples = 1_000_000

with open("1M_py_sha256_random.log2", "x+") as f:
    for i in range(num_samples):

        # x = random.randint(0, num_samples) # random
        x = i # serie

        h = compute_any_hash(x, hasher=hashlib.sha256)
        f.write(h + "\n")

print("done")