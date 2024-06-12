import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# legacy
dynm_max_sqrt = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre10-dynm-max+sqrt.log")
dynm_max_4 = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre10-dynm-max+4.log")
# dynm_max_3_5 = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre10-dynm-max+3.5.log")
hll_vanilla = parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4-pre10.log")
m_m667 = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre10-m+powM.667.log")
hll_experimental = parse("logs/optimal-m-search/local-d14-anull-hsha3-1024b-case4-pre10-dynm-experimental.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(dynm_max_sqrt[0], dynm_max_sqrt[1], '-', linewidth=1, label='ours: dynamic base=(max+√(precision))')
ax.plot(dynm_max_4[0], dynm_max_4[1], '-', linewidth=1, label='ours: dynamic base=(max+4)')
# ax.plot(dynm_max_3_5[0], dynm_max_3_5[1], '-', linewidth=1, label='ours: dynamic base=(max+3.5)')
ax.plot(hll_vanilla[0], hll_vanilla[1], '-', linewidth=1, label='vanilla hll: fixed base=32')
ax.plot(m_m667[0], m_m667[1], '-', linewidth=1, label='m+m^2.3')
# ax.plot(hll_experimental[0], hll_experimental[1], '-', marker='o', linewidth=1, label='experimental')
ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("Dynamic HLL (1024 bit sha3) vs vanilla HLL \nusing 10 bit precision\n applied to miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
