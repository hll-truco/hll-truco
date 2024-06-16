import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

# legacy
# hll_1 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636010.out")
# hll_2 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636274.out")
hll_3 = parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3636397.out")
hll_4 = parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3646786.out")
hll_5 = parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3648398.out")
hll_6 = parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3650313.out")

hll_13_14 = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3669198.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3681497.out")
])

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
# ax.plot(hll_1[0], hll_1[1], '-', linewidth=1, label='run 1')
# ax.plot(hll_2[0], hll_2[1], '-', linewidth=1, label='run 2')
ax.plot(hll_3[0], hll_3[1], '-', linewidth=1, label='run 3')
ax.plot(hll_4[0], hll_4[1], '-', linewidth=1, label='run 4')
ax.plot(hll_5[0], hll_5[1], '-', linewidth=1, label='run 5')
ax.plot(hll_6[0], hll_6[1], '-', linewidth=1, label='run 6')

ax.plot(hll_13_14[0], hll_13_14[1], '-', linewidth=1, label='run 13_14')

ax.set_title("HLL")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
