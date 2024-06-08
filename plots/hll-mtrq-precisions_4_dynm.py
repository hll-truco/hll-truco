import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# legacy
hll_1024_prec_4 = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre4.log")
hll_1024_prec_4_dynm_max0 = parse("logs/hll-dist-http-32-vs-1024/local-d14-anull-hsha3-1024b-case4-pre4-dynm-max+0.log")
hll_1024_prec_4_dynm_max1 = parse("logs/hll-dist-http-32-vs-1024/local-d14-anull-hsha3-1024b-case4-pre4-dynm-max+1.log")
hll_1024_prec_4_dynm_max2 = parse("logs/hll-dist-http-32-vs-1024/local-d14-anull-hsha3-1024b-case4-pre4-dynm-max+2.log")
hll_1024_prec_4_dynm_max3 = parse("logs/hll-dist-http-32-vs-1024/local-d14-anull-hsha3-1024b-case4-pre4-dynm-max+3.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_1024_prec_4[0], hll_1024_prec_4[1], '-', linewidth=1, label='4 bits precision vanilla (fixed 32)')
ax.plot(hll_1024_prec_4_dynm_max0[0], hll_1024_prec_4_dynm_max0[1], '-', linewidth=1, label='4 bits precision + dynm (max+0)')
ax.plot(hll_1024_prec_4_dynm_max1[0], hll_1024_prec_4_dynm_max1[1], '-', linewidth=1, label='4 bits precision + dynm (max+1)')
ax.plot(hll_1024_prec_4_dynm_max2[0], hll_1024_prec_4_dynm_max2[1], '-', linewidth=1, label='4 bits precision + dynm (max+2)')
ax.plot(hll_1024_prec_4_dynm_max3[0], hll_1024_prec_4_dynm_max3[1], '-', linewidth=1, label='4 bits precision + dynm (max+3)')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("HLL (1024 bit sha3) 4 bit precision comparison for miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()