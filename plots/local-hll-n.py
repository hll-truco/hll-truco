import matplotlib.pyplot as plt
from common import parse_utils

# legacy
# hll_bin_p16_r1s_big = parse_utils.parse("logs/hll-local-n/n1e9-bin-p16-r1s_big.log")
hll_n1e8_bin_p16_r1s_ui64 = parse_utils.parse("logs/hll-local-n/n1e8-bin-p16-r1s_ui64.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
# ax.plot(hll_bin_p16_r1s_big[0], hll_bin_p16_r1s_big[1], '-', linewidth=1, label='bin p16 r1s big')
ax.plot(hll_n1e8_bin_p16_r1s_ui64[0], hll_n1e8_bin_p16_r1s_ui64[1], '-', linewidth=1, label='hll-')

# ax.axhline(y=(1e9), linestyle='--', linewidth=0.5, alpha=0.5)

ax.set_title("HLL")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
