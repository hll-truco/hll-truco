import matplotlib.pyplot as plt

import sys
sys.path.append('cmd/_')
import parse_utils

title = "HLL estimate for mini-Truco with -deck=14 -info=InfosetRondaBase -abs=null"
real = 248_732

hll_axiom = parse_utils.parse_structured_log("cmd/plots/count-infosets-ronda-hll/data/axiom-d14-anull-irb-l600.log")
hll_duvall_pp = parse_utils.parse_structured_log("cmd/plots/count-infosets-ronda-hll/data/clarkduvall-++-d14-anull-irb-l600.log")
hll_duvall = parse_utils.parse_structured_log("cmd/plots/count-infosets-ronda-hll/data/clarkduvall-d14-anull-irb-l600.log")
hll_hll = parse_utils.parse_structured_log("cmd/plots/count-infosets-ronda-hll/data/hll-d14-anull-irb-l600.log")

hll_axiom = parse_utils.keep(hll_axiom)
hll_duvall = parse_utils.keep(hll_duvall)
hll_duvall_pp = parse_utils.keep(hll_duvall_pp)
hll_hll = parse_utils.keep(hll_hll)

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(hll_axiom[0], hll_axiom[1], '-', linewidth=0.8, label='hll_axiom')
ax.plot(hll_duvall[0], hll_duvall[1], '-', linewidth=0.8, label='hll_duvall')
ax.plot(hll_duvall_pp[0], hll_duvall_pp[1], '-', linewidth=0.8, label='hll_duvall_++')
ax.plot(hll_hll[0], hll_hll[1], '-', linewidth=0.8, label='hll_hll')

ax.axhline(y=(real), color='r', linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_xlabel('time (s)')
ax.set_ylabel('Estimated cardinality')
ax.set_title(title)
ax.legend()

plt.tight_layout()
plt.show()

