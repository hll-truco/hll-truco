import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# legacy
hll_vanilla = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b.log")
dynm_max4 = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre16-dynm-max+4.log")
min_larger = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-dynm-minlarger.log")
min_larger_1 = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-dynm-minlarger+1.log")
min_larger_2 = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-dynm-minlarger+2.log")
min_larger_3 = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-dynm-minlarger+3.log")
min_larger_4 = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-dynm-minlarger+4.log")
fixed_1024 = parse("logs/hll-min-larger-experiments/local-d14-anull-hsha3-1024b-pre16-1024.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
# ax.plot(min_larger[0], min_larger[1], '-', linewidth=1, label='ours: dynamic base=(min larger)')
# ax.plot(min_larger_1[0], min_larger_1[1], '-', linewidth=1, label='ours: dynamic base=(min larger+1)')
# ax.plot(min_larger_2[0], min_larger_2[1], '-', linewidth=1, label='ours: dynamic base=(min larger+2)')
# ax.plot(min_larger_3[0], min_larger_3[1], '-', linewidth=1, label='ours: dynamic base=(min larger+3)')
# ax.plot(dynm_max4[0], dynm_max4[1], '-', linewidth=1, label='ours: dynamic base=(max+√(precision))')
ax.plot(min_larger_4[0], min_larger_4[1], '-', linewidth=1, label='ours: dynamic base=(min larger+4)')
# ax.plot(hll_vanilla[0], hll_vanilla[1], '-', linewidth=1, label='vanilla hll: fixed base=32')
ax.plot(fixed_1024[0], fixed_1024[1], '-', linewidth=1, label='hll: fixed base=1024')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("Dynamic HLL (1024 bit sha3) vs vanilla HLL \n(with 16 bit precision)\n applied to for miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()

rel_err = lambda y_hat: (y_hat[1][-1]-real)/real

errors = {
    "m+0": rel_err(min_larger),
    "m+1": rel_err(min_larger_1),
    "m+2": rel_err(min_larger_2),
    "m+3": rel_err(min_larger_3),
    "m+4": rel_err(min_larger_4),
    "f1024": rel_err(fixed_1024),
    "vanilla": rel_err(hll_vanilla),
}
import json
print(json.dumps(errors, indent=2))
