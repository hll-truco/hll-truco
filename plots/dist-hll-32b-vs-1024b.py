import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

# legacy
hll_32 = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-32b.log")
hll_32_4 = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-32b-case4.log")
hll_1024 = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b.log")
hll_1024_4 = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b-case4.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_32[0], hll_32[1], '-', linewidth=1, label='32bit hll')
ax.plot(hll_32_4[0], hll_32_4[1], '-', linewidth=1, label='32bit hll (4)')
ax.plot(hll_1024[0], hll_1024[1], '-', linewidth=1, label='1024bit hll')
ax.plot(hll_1024_4[0], hll_1024_4[1], '-', linewidth=1, label='1024bit hll (4)')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_title("HLL")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
