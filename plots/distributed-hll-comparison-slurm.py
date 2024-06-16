import matplotlib.pyplot as plt
import numpy as np
from common import parse_utils

real = 248_732

# legacy
hll_w1 = parse_utils.parse("logs/hll-dist-http/slurm/legacy/hllroot.3629245-w1d14anull.out")
hll_w2 = parse_utils.parse("logs/hll-dist-http/slurm/legacy/hllroot.3630838-w2d14anull.out")
hll_w4 = parse_utils.parse("logs/hll-dist-http/slurm/legacy/hllroot.3630871-w4d14anull.out")
hll_w8 = parse_utils.parse("logs/hll-dist-http/slurm/legacy/hllroot.3629226-w8d14anull.out")
hll_w64 = parse_utils.parse("logs/hll-dist-http/slurm/legacy/hllroot.3631022-w64d14anull.out")
# big
hll_w1_big = parse_utils.parse("logs/hll-dist-http/slurm/big/hllroot.3631774-w1d14anull.out")
hll_w2_big = parse_utils.parse("logs/hll-dist-http/slurm/big/hllroot.3631909-w2d14anull.out")
hll_w4_big = parse_utils.parse("logs/hll-dist-http/slurm/big/hllroot.3631736-w4d14anull.out")
hll_w8_big = parse_utils.parse("logs/hll-dist-http/slurm/big/hllroot.3631663-w8d14anull.out")
hll_w64_big = parse_utils.parse("logs/hll-dist-http/slurm/big/hllroot.3631597-w64d14anull.out")




# estimate evolution over time

fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
ax.plot(hll_w1[0], hll_w1[1], '-', linewidth=1, label='1 workers')
# ax.plot(hll_w2[0], hll_w2[1], '-', linewidth=1, label='2 workers')
# ax.plot(hll_w4[0], hll_w4[1], '-', linewidth=1, label='4 workers')
ax.plot(hll_w8[0], hll_w8[1], '-', linewidth=1, label='8 workers')
ax.plot(hll_w64[0], hll_w64[1], '-', linewidth=1, label='64 workers')
# big
ax.plot(hll_w1_big[0], hll_w1_big[1], '-', linewidth=1, label='1 workers BIG')
# ax.plot(hll_w2_big[0], hll_w2_big[1], '-', linewidth=1, label='2 workers BIG')
# ax.plot(hll_w4_big[0], hll_w4_big[1], '-', linewidth=1, label='4 workers BIG')
ax.plot(hll_w8_big[0], hll_w8_big[1], '-', linewidth=1, label='8 workers BIG')
ax.plot(hll_w64_big[0], hll_w64_big[1], '-', linewidth=1, label='64 workers BIG')
# real
ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_title("Evolution of Distributed HLL (float64 vs big) on miniTruco-14 for Different Number of Workers")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()



# estimate evolution over time

data = {
    "1": hll_w1[2][-1],
    "2": hll_w2[2][-1],
    "4": hll_w4[2][-1],
    "8": hll_w8[2][-1],
    "64": hll_w64[2][-1],
}

data_big = {
    "1": hll_w1_big[2][-1],
    "2": hll_w2_big[2][-1],
    "4": hll_w4_big[2][-1],
    "8": hll_w8_big[2][-1],
    "64": hll_w64_big[2][-1],
}

# Preparing the data
courses = list(data.keys())
values = list(data.values())
values_alt = list(data_big.values())

# Setting up the plot
fig, ax = plt.subplots(1, 1, figsize=(10, 5))

# Bar width and positions
bar_width = 0.4
index = np.arange(len(courses))

# Plotting the bars
plt.bar(index, values, bar_width, color="maroon", label='float64')
plt.bar(index + bar_width, values_alt, bar_width, color="teal", label='big.Float')

# Adding labels and title
ax.set_title("Total Number of Messages for Different Number of Workers\n(more is better)")
ax.set_ylabel("Estimated cardinality")
ax.set_xlabel("Number of workers")
ax.set_xticks(index + bar_width / 2)
ax.set_xticklabels(courses)
plt.yscale('log', base=2)

# Adding legend
plt.legend()

# Displaying the plot
plt.tight_layout()
plt.show()

