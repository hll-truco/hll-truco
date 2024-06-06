import matplotlib.pyplot as plt

import sys
sys.path.append('cmd/_')
import parse_utils

real = 248_732

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# legacy
hll_1024_prec_16 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b.log")
hll_1024_prec_10 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre10.log")
hll_1024_prec_6 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre6.log")
hll_1024_prec_5 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre5.log")
hll_1024_prec_4 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre4.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_1024_prec_16[0], hll_1024_prec_16[1], '-', linewidth=1, label='16 bits precision')
ax.plot(hll_1024_prec_10[0], hll_1024_prec_10[1], '-', linewidth=1, label='10 bits precision')
ax.plot(hll_1024_prec_6[0], hll_1024_prec_6[1], '-', linewidth=1, label='6 bits precision')
ax.plot(hll_1024_prec_5[0], hll_1024_prec_5[1], '-', linewidth=1, label='5 bits precision')
ax.plot(hll_1024_prec_4[0], hll_1024_prec_4[1], '-', linewidth=1, label='4 bits precision')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("HLL (1024 bit sha3) precision comparison for miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
