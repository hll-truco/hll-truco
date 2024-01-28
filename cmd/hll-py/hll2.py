import math
import hashlib

class HyperLogLog:
    def __init__(self, b):
        self.b = b
        self.m = 1 << b
        self.M = [0] * self.m

    def add(self, x):
        try:
            hash_binary = bin(int(hashlib.sha256(str(x).encode('utf-8')).hexdigest(), 16))[2:]
            i = int(hash_binary[-self.b:], 2)
            w = hash_binary[:-self.b][::-1]
            rho = w.index("1") + 1
            self.M[i] = max(self.M[i], rho)
        except Exception as e:
            print(e)
            import sys
            sys.exit(0)

    def count(self):
        alpha = self.alpha(self.m)
        Z = alpha * float(self.m ** 2) / sum([2.0 ** (-x) for x in self.M])
        if Z <= 2.5 * self.m:
            V = self.M.count(0)
            if V != 0:
                return round(self.m * math.log(float(self.m) / V)), self.M
            else:
                return round(Z), self.M
        elif Z <= (1 << 32) / 30.0:
            return round(Z), self.M
        else:
            return round(-1 * (1 << 32) * math.log(1 - Z / (1 << 32))), self.M

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
