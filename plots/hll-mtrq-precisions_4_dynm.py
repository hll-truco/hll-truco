import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

# legacy
hll_1024_prec_4 = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre4.log")
hll_1024_prec_4_dynm_max0 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-minlarger+0.log")
hll_1024_prec_4_dynm_max1 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-minlarger+1.log")
hll_1024_prec_4_dynm_max2 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-minlarger+2.log")
hll_1024_prec_4_dynm_max3 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-minlarger+3.log")
hll_1024_prec_4_dynm_max4 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-minlarger+4.log")
hll_1024_prec_4_dynm_f1024 = parse_utils.parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre4-fixed-1024.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_1024_prec_4[0], hll_1024_prec_4[1], '-', linewidth=1, label='4 bits precision vanilla (base=32)')
ax.plot(hll_1024_prec_4_dynm_max0[0], hll_1024_prec_4_dynm_max0[1], '-', linewidth=1, label='4 bits precision + base=(max+0)')
ax.plot(hll_1024_prec_4_dynm_max1[0], hll_1024_prec_4_dynm_max1[1], '-', linewidth=1, label='4 bits precision + base=(max+1) ★!')
ax.plot(hll_1024_prec_4_dynm_max2[0], hll_1024_prec_4_dynm_max2[1], '-', linewidth=1, label='4 bits precision + base=(max+2)')
ax.plot(hll_1024_prec_4_dynm_max3[0], hll_1024_prec_4_dynm_max3[1], '-', linewidth=1, label='4 bits precision + base=(max+3)')
ax.plot(hll_1024_prec_4_dynm_max4[0], hll_1024_prec_4_dynm_max4[1], '-', linewidth=1, label='4 bits precision + base=(max+4)')
# ax.plot(hll_1024_prec_4_dynm_f1024[0], hll_1024_prec_4_dynm_f1024[1], '-', linewidth=1, label='4 bits precision + base=(f1024)')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("Dynamic HLL (1024 bit sha3) vs vanilla HLL (using 4 bit precision)\n for miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()

rel_err = lambda y_hat: (y_hat[1][-1]-real)/real

errors = {
    "m+0": rel_err(hll_1024_prec_4_dynm_max0),
    "m+1": rel_err(hll_1024_prec_4_dynm_max1),
    "m+2": rel_err(hll_1024_prec_4_dynm_max2),
    "m+3": rel_err(hll_1024_prec_4_dynm_max3),
    # "m+4": rel_err(min_larger_4),
    "vanilla": rel_err(hll_1024_prec_4),
}

import json
print(json.dumps(errors, indent=2))
