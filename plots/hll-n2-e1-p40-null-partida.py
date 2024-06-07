import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# legacy
hll_1 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636010.out")
hll_2 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636274.out")
hll_3 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636397.out")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_1[0], hll_1[1], '-', linewidth=1, label='run 1')
ax.plot(hll_2[0], hll_2[1], '-', linewidth=1, label='run 2')
ax.plot(hll_3[0], hll_3[1], '-', linewidth=1, label='run 3')

ax.set_title("HLL")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
