import matplotlib.pyplot as plt

import sys
sys.path.append('cmd/_')
import parse_utils

title = "Distributed-HLL evol. for -deck=14 -info=InfosetRondaBase -abs=null"
real = 248_732

hll_hll = parse_utils.parse_structured_log("logs/count-infosets-ronda-hll/hll-d14-anull-irb-l600.log")
hll_dist_http_w64 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3629053.out")

hll_hll = parse_utils.keep(hll_hll)
hll_dist_http_w64 = parse_utils.keep(hll_dist_http_w64)

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(hll_hll[0], hll_hll[1], '-', linewidth=0.8, label='hll 1 worker')
ax.plot(hll_dist_http_w64[0], hll_dist_http_w64[1], '-', linewidth=0.8, label='hll 64 workers')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_xlabel('time (s)')
ax.set_ylabel('Estimated cardinality')
ax.set_title(title)
ax.legend()

plt.tight_layout()
plt.show()

