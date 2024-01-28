import math
import hashlib

class HyperLogLog:
    def __init__(self):
        self.M = 0

    def add(self, x):
        hash_binary = bin(int(hashlib.sha256(str(x).encode('utf-8')).hexdigest(), 16))[2:]
        w = hash_binary[::-1]
        rho = w.index("1") + 1
        self.M = max(self.M, rho)

def count(buckets):
    m = len(buckets)
    alpha = _alpha(m)
    Z = alpha * float(m ** 2) / sum([2.0 ** (-x) for x in buckets])
    if Z <= 2.5 * m:
        V = m.count(0)
        if V != 0:
            return round(m * math.log(float(m) / V))
        else:
            return round(Z)
    elif Z <= (1 << 32) / 30.0:
        return round(Z)
    else:
        return round(-1 * (1 << 32) * math.log(1 - Z / (1 << 32)))

def _alpha(m):
    if m == 16:
        return 0.673
    elif m == 32:
        return 0.697
    elif m == 64:
        return 0.709
    else:
        return 0.7213 / (1 + 1.079 / m)

import datetime

# Example usage
for bits_buckets in range(1,14+1):
    m = 2 ** bits_buckets
    start = datetime.datetime.now()
    T = 10000000 - 10
    delta = round(T / m)

    buckets = []

    for i in range(m):
        _from, _to = delta*i, delta*(i+1)
        # print(f"from {_from} to {_to}")
        hll = HyperLogLog()
        for i in range(_from,_to+1): hll.add(i)
        buckets += [hll.M]

    c = count(buckets)
    delta = datetime.datetime.now() - start
    print(f"Estimated cardinality for b={bits_buckets} is {c} ({delta}) {buckets}")
