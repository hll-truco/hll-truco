import matplotlib.pyplot as plt

import sys
sys.path.append('cmd/_')
import parse_utils

real = 248_732

hll_w1 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3629245-w1d14anull.out")
hll_w2 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3630838-w2d14anull.out")
hll_w4 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3630871-w4d14anull.out")
hll_w8 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3629226-w8d14anull.out")
hll_w64 = parse_utils.parse_structured_log("logs/hll-dist-http/slurm/hllroot.3631022-w64d14anull.out")

hll_w1 = parse_utils.keep(hll_w1)
hll_w2 = parse_utils.keep(hll_w2)
hll_w4 = parse_utils.keep(hll_w4)
hll_w8 = parse_utils.keep(hll_w8)
hll_w64 = parse_utils.keep(hll_w64)



# estimate evolution over time

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(hll_w1[0], hll_w1[1], '-', linewidth=0.8, label='hll 1 workers')
ax.plot(hll_w2[0], hll_w2[1], '-', linewidth=0.8, label='hll 2 workers')
ax.plot(hll_w4[0], hll_w4[1], '-', linewidth=0.8, label='hll 4 workers')
ax.plot(hll_w8[0], hll_w8[1], '-', linewidth=0.8, label='hll 8 workers')
ax.plot(hll_w64[0], hll_w64[1], '-', linewidth=0.8, label='hll 64 workers')
ax.axhline(y=(real), linestyle='--', linewidth=0.5, alpha=0.5, label=f"real {real:,}")

ax.set_title("Distributed-HLL evol. on miniTruco for different num. of workers")
ax.set_ylabel('Estimated cardinality')
ax.set_xlabel('time (s)')
ax.legend()

plt.tight_layout()
plt.show()



# estimate evolution over time

data = {
    "w1": hll_w1[2][-1],
    "w2": hll_w2[2][-1],
    "w4": hll_w4[2][-1],
    "w8": hll_w8[2][-1],
    "w64": hll_w64[2][-1],
}
courses = list(data.keys())
values = list(data.values())

fig, ax = plt.subplots(1,1, figsize=(10,5))
plt.bar(courses, values, color="maroon", width=0.4)

ax.set_title("Total num. of messages for diff. num. of workers")
ax.set_ylabel("Estimated cardinality")
ax.set_xlabel("time (s)")
# ax.legend()

plt.yscale('log', base=2)
plt.tight_layout()
plt.show()

