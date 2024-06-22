import matplotlib.pyplot as plt
from common import parse_utils

real = 248_732

# legacy
hll_vanilla = parse_utils.parse("logs/hll-dist-http-32-vs-1024/http-w1-d14-anull-hsha3-1024b.log")
f32 = parse_utils.parse("logs/hll/mini/d14/p16-f32.log")
sp0 = parse_utils.parse("logs/hll/mini/d14/p16-sp0.log")
sp4 = parse_utils.parse("logs/hll/mini/d14/p16-sp4.log")
sp5 = parse_utils.parse("logs/hll/mini/d14/p16-sp5.log")
sp6 = parse_utils.parse("logs/hll/mini/d14/p16-sp6.log")
sp10 = parse_utils.parse("logs/hll/mini/d14/p16-sp10.log")

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
# ax.plot(hll_vanilla[0], hll_vanilla[1], '-', linewidth=1, label='vanilla hll (fixed=32)')
# ax.plot(sp0[0], sp0[1], '-', linewidth=1, label='sp+0')
# ax.plot(sp4[0], sp4[1], '-', linewidth=1, label='sp+4')
# ax.plot(sp10[0], sp10[1], '-', linewidth=1, label='sp+10')
ax.plot(sp5[0], sp5[1], '-', linewidth=1, label='sp+5')
# ax.plot(sp6[0], sp6[1], '-', linewidth=1, label='sp+6')
ax.plot(f32[0], f32[1], '-', linewidth=1, label='fixed 32')

ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, color='black')

ax.set_title("Dynamic HLL (1024 bit sha3) vs vanilla HLL \n(with 16 bit precision)\n applied to for miniTruco-14")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()

