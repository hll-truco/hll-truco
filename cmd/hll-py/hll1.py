import math
import hashlib

class HyperLogLog:
    def __init__(self, b):
        self.b = b
        self.buckets = [0] * (2 ** b)

    def add(self, x):
        h = hashlib.sha256(str(x).encode('utf-8')).digest()
        # 4 * 8bit = 32 bit
        # el tipo agarra los primeros 4 bytes
        i = int.from_bytes(h[:4], byteorder='big')
        w = int.from_bytes(h[4:], byteorder='big')
        j = i >> (32 - self.b)
        self.buckets[j] = max(self.buckets[j], self.rho(w))

    def count(self):
        m = len(self.buckets)
        alpha = self.alpha(m)
        Z = alpha * float(m ** 2) / sum([2.0 ** (-x) for x in self.buckets])
        if Z <= 2.5 * m:
            V = self.buckets.count(0)
            if V != 0:
                return round(m * math.log(float(m) / V)), 0
            else:
                return round(Z), 0
        elif Z <= (1 << 32) / 30.0:
            return round(Z), 0
        else:
            return round(-1 * (1 << 32) * math.log(1 - Z / (1 << 32))), 0

    def rho(self, w):
        i = 0
        while (w & (1 << i)) == 0:
            i += 1
        return i + 1

    def alpha(self, m):
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
for b in range(1,14+1):
    start = datetime.datetime.now()
    hll = HyperLogLog(b)
    for i in range(10,10000000):
        hll.add(i)
    c,M = hll.count()
    delta = datetime.datetime.now() - start
    print(f"Estimated cardinality for b={b} is {c} {M} ({delta})")
