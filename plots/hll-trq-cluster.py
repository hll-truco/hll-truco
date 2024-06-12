import matplotlib.pyplot as plt
from common import parse_utils

parse = lambda f: parse_utils.keep(parse_utils.parse_structured_log(f))

# n:2 envido:1 pts:40 absId:null infoset:InfosetPartidaXXLarge
hll_prec4 = parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3666028.out")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_prec4[0], hll_prec4[1], '-', linewidth=1, label='hll')

ax.set_title("HLL (1024 bits; 4 bit precision) estimate on infosets \nfor 2-players 40-points imprefect-recall Uruguayan Truco")
ax.set_ylabel('Estimated cardinality of infosets')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
