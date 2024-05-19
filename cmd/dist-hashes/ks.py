import math
import numpy as np
from scipy.stats import kstest
import hashlib
import random
from typing import Callable

def compute_any_hash(n:int, hasher:Callable[[bytes],object]) -> str:
    """Compute the "hasher" hash of an integer n."""
    n_str = str(n).encode('utf-8')
    hash_obj = hasher(n_str)
    hash_hex = hash_obj.hexdigest()
    return hash_hex

def get_zero_bit_positions(hash_hex:str, hash_bits_size:int=256) -> list[int]:
    """Get the positions of zero bits in the hexadecimal hash string."""
     # SHA-256 hash is 256 bits
    hash_bin = bin(int(hash_hex, 16))[2:].zfill(hash_bits_size)
    zero_positions = [i for i, bit in enumerate(hash_bin) if bit == '0']
    return zero_positions

def generate_random_n_bit_hash(n:int=256) -> str:
    """Generate a random n-bit hash represented as a hexadecimal string."""
    random_bits = ''.join(random.choice('01') for _ in range(n))  # Generate n random bits
    hex_length = math.ceil(n / 4)
    hash_hex = hex(int(random_bits, 2))[2:].zfill(hex_length)  # Convert to hexadecimal and pad with zeros if necessary
    return hash_hex

def extract_zero_positions_from_hashes(
        num_hashes     :int,
        hash_func      :Callable[[],str],
        hash_bits_size :int=256
    ) -> np.ndarray:

    """Extract normalized zero bit positions from a number of hashes."""
    all_zero_positions = []
    for _ in range(num_hashes):
        hash_hex = hash_func()
        zero_positions = get_zero_bit_positions(hash_hex, hash_bits_size)
        all_zero_positions.extend(zero_positions)
    return np.array(all_zero_positions) / float(hash_bits_size - 1)  # Normalize positions to [0, 1]

def read_hashes(file:str) -> list[str]:
    res = []
    with open(file) as f:
        for _,l in enumerate(f.readlines()):
            res += [l.strip()]
    return res

if __name__ == "__main__":

    num_samples    = 1_000_000
    hasher         = hashlib.sha512
    hash_bits_size = 512

    # randomly generated sha256
    # computed = [
    #     compute_any_hash(random.randint(0, num_samples), hasher=hashlib.sha256)
    #     for _ in range(num_samples)
    # ]

    # sequentially generated sha256
    # computed = [
    #     compute_any_hash(i, hasher)
    #     for i in range(num_samples)
    # ]

    # sequentially generated sha512
    computed = [
        compute_any_hash(i, hasher)
        for i in range(num_samples)
    ]

    # read sequential sha256 genereted with `gen_1m.py`
    # computed = read_hashes(file="1M_py_sha256_random.log")

    it = iter(computed)
    next_computed_hash = lambda: next(it)

    actual_zero_positions_normalized = extract_zero_positions_from_hashes(
                                            num_samples,
                                            next_computed_hash,
                                            hash_bits_size)

    # Generate simulated 256-bit hashes and extract zero positions
    simulated_zero_positions_normalized = extract_zero_positions_from_hashes(
                                            num_samples,
                                            lambda: generate_random_n_bit_hash(n=hash_bits_size),
                                            hash_bits_size)

    # Perform the 2-sample K-S test
    ks_statistic, p_value = kstest(
                                actual_zero_positions_normalized,
                                simulated_zero_positions_normalized)

    # Output the results
    print(f"KS Statistic: {ks_statistic}")
    print(f"P-Value: {p_value}")

    # Interpret the result
    if p_value > 0.05:
        print("The distributions of zero positions are similar (fail to reject H0).")
    else:
        print("The distributions of zero positions are different (reject H0).")

    # plot

    import matplotlib.pyplot as plt

    # Histogram of zero bit positions
    plt.hist(actual_zero_positions_normalized, bins=256, density=True, alpha=0.75, color='blue')
    plt.title('Distribution of Zero Bit Positions in One Billion SHA-256 Hashes')
    plt.xlabel('Normalized Bit Position')
    plt.ylabel('Density')
    plt.show()

    # sha256
    # KS Statistic: 6.998626494669757e-05
    # P-Value: 0.9125131694719305
    # The distributions of zero positions are similar (fail to reject H0).

    # random-sha256
    # KS Statistic: 6.575667544847508e-05
    # P-Value: 0.9447675180662787

    # random-sha256
    # KS Statistic: 7.9673563030902e-05
    # P-Value: 0.8112103087681888

    # sequential-sha256
    # KS Statistic: 6.470931814273051e-05
    # P-Value: 0.9514827707727992

    # KS Statistic: 0.00014691937639882413
    # P-Value: 0.12616106744044997
    # The distributions of zero positions are similar (fail to reject H0).

    # KS Statistic: 9.152486883529098e-05
    # P-Value: 0.6570888929927498
    # The distributions of zero positions are similar (fail to reject H0).

    # parametric (dynamic) hash length
    # KS Statistic: 6.590226494451734e-05
    # P-Value: 0.9438012713093233
    # The distributions of zero positions are similar (fail to reject H0).