import matplotlib.pyplot as plt

import sys
sys.path.append('cmd/_')
import parse_utils

title = "Distributed-HLL evol. with -deck=14 -info=InfosetRondaBase -abs=null"
real = 248_732

hll_hll = parse_utils.parse_structured_log("logs/count-infosets-ronda-hll/hll-d14-anull-irb-l600.log")
hll_dist_http_w2 = parse_utils.parse_structured_log("logs/hll-dist-http/http-w2-d14-anull-hsha3.log")
hll_dist_http_w4 = parse_utils.parse_structured_log("logs/hll-dist-http/http-w4-d14-anull-hsha3.log")
hll_dist_http_w8 = parse_utils.parse_structured_log("logs/hll-dist-http/http-w8-d14-anull-hsha3.log")

hll_hll = parse_utils.keep(hll_hll)
hll_dist_http_w2 = parse_utils.keep(hll_dist_http_w2)
hll_dist_http_w4 = parse_utils.keep(hll_dist_http_w4)
hll_dist_http_w8 = parse_utils.keep(hll_dist_http_w8)

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(hll_hll[0], hll_hll[1], '-', linewidth=0.8, label='hll_hll')
ax.plot(hll_dist_http_w2[0], hll_dist_http_w2[1], '-', linewidth=0.8, label='hll_dist_http_w2')
ax.plot(hll_dist_http_w4[0], hll_dist_http_w4[1], '-', linewidth=0.8, label='hll_dist_http_w4')
ax.plot(hll_dist_http_w8[0], hll_dist_http_w8[1], '-', linewidth=0.8, label='hll_dist_http_w8')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_xlabel('time (s)')
ax.set_ylabel('Estimated cardinality')
ax.set_title(title)
ax.legend()

plt.tight_layout()
plt.show()

