import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

# legacy
hll_vanilla = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre6.log")
dynm_max_sqrt = parse_utils.parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre6-dynm-max+sqrt.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_vanilla[0], hll_vanilla[1], '-', linewidth=1, label='vanilla hll: fixed base=32')
ax.plot(dynm_max_sqrt[0], dynm_max_sqrt[1], '-', linewidth=1, label='ours: dynamic base=(max+√(precision))')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("Dynamic HLL (1024 bit sha3) vs vanilla HLL \nusing 6 bit precision\n applied to miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
