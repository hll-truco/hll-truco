import matplotlib.pyplot as plt
from common import parse_utils

title = "HLL estimate for mini-Truco with -deck=14 -info=InfosetRondaBase -abs=null -hash=sha3"
real = 248_732

ours_max = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm.log")
dynm_m2 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_max2.log")
dynm_m3 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_max3.log")
fixed_20 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f20.log")
fixed_21 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f21.log")
fixed_22 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f22.log")
fixed_24 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f24.log")
fixed_27 = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f27.log")
fixed_32 = parse_utils.parse("logs/count-infosets-ronda-hll/hll-d14-anull-irb-l600.log")
# fixed_32_bis = parse_utils.parse("logs/count-infosets-ronda-hll/ours-d14-anull-irb-l600_dynm_f32.log")

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(ours_max[0], ours_max[1], '-', linewidth=0.8, label='hll ours (sha3 + max))')
# ax.plot(dynm_m2[0], dynm_m2[1], '-', linewidth=0.8, label='hll ours (sha3 + m+2)')
ax.plot(dynm_m3[0], dynm_m3[1], '-', linewidth=0.8, label='hll ours (sha3 + m+3)')
ax.plot(fixed_20[0], fixed_20[1], '-', linewidth=0.8, label='hll ours (sha3 + f20)')
ax.plot(fixed_21[0], fixed_21[1], '-', linewidth=0.8, label='hll ours (sha3 + f21)')
ax.plot(fixed_22[0], fixed_22[1], '-', linewidth=0.8, label='hll ours (sha3 + f22)')
ax.plot(fixed_24[0], fixed_24[1], '-', linewidth=0.8, label='hll ours (sha3 + f24)')
ax.plot(fixed_27[0], fixed_27[1], '-', linewidth=0.8, label='hll ours (sha3 + f27))')
ax.plot(fixed_32[0], fixed_32[1], '-', linewidth=0.8, label='hll ours (sha3 + f32)')
# ax.plot(fixed_32_bis[0], fixed_32_bis[1], '-', linewidth=0.8, label='hll ours (sha3 + f32) bis')

ax.axhline(y=(real), color='r', linestyle='--', linewidth=0.5, alpha=0.5)

ax.set_xlabel('time (s)')
ax.set_ylabel('Estimated cardinality')
ax.set_title(title)
ax.legend()

plt.tight_layout()
plt.show()

