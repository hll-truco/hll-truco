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

p4_b1024 = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3696572.out"),
])

p16_b1024 = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P40AnullIipxxlW256/hllroot1.3704925.out"),
])

p16_b1024_20PTS = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot1.3788070.out"),
])

p16_b1024_20PTS_nomazo = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot1.3848128.out"),
])

p6_b1024_20PTS_nomazo = parse_utils.joint([
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_18.3890487.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_17.3895665.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_16.3902213.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_15.3902217.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_14.3902222.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_13.3902226.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_12.3902230.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_11.3902235.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_10.3902239.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_9.3902246.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_8.3902252.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/backup-wierd-behavior/hllroot_7.3902259.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_16.3903149.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_15.3908456.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_14.3926778.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_13.3952090.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_12.3967094.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_11.3977900.out"),
    # parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_10.3987756.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_9.3992094.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_8.3996251.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_7.4004310.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_6.4008743.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_5.4018203.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_4.4025515.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_3.4031821.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_2.4039920.out"),
    parse_utils.parse("/Users/jp/Downloads/cluster/hll/2p/E1P20AnullIipxxlW256/hllroot_1.4069887.out"),
])

# estimate evolution over time
fig, ax = plt.subplots(1,1, figsize=(10,5))

# legacy
# ax.plot(hll_1[0], hll_1[1], '-', linewidth=1, label='run 1')
# ax.plot(hll_2[0], hll_2[1], '-', linewidth=1, label='run 2')
# ax.plot(hll_3[0], hll_3[1], '-', linewidth=1, label='run 3')
# ax.plot(hll_4[0], hll_4[1], '-', linewidth=1, label='run 4')
# ax.plot(hll_5[0], hll_5[1], '-', linewidth=1, label='run 5')
# ax.plot(hll_6[0], hll_6[1], '-', linewidth=1, label='run 6')

# ax.plot(hll_13_14[0], hll_13_14[1], '-', linewidth=1, label='run 13_14')
# ax.plot(p4_b1024[0], p4_b1024[1], '-', linewidth=1, label='p4 b1024')
ax.plot(p16_b1024[0], p16_b1024[1], '-', linewidth=1, label='VIEJA: prec=16 | base=1024 | pts=20 | mazo | 100 cores')

# 20 pts + totally random chi
# ax.plot(p16_b1024_20PTS[0], p16_b1024_20PTS[1], '-', linewidth=1, label='prec=16 base=1024 pts=20')
# 20 pts + totally random chi + NO-MAZO
# ax.plot(p16_b1024_20PTS_nomazo[0], p16_b1024_20PTS_nomazo[1], '-', linewidth=1, label='prec=16 base=1024 pts=20 no-mazo')

# 20 pts + totally random chi + NO-MAZO (1+78+200 cores)
ax.plot(p6_b1024_20PTS_nomazo[0], p6_b1024_20PTS_nomazo[1], '-', linewidth=1, label='NUEVA: prec=6 | base=1024 | pts=20 | no-mazo | 200 cores')

ax.set_title("HLL")
ax.set_ylabel('Estimated cardinality of infosets at round level')
ax.set_xlabel('Time (sec.)')
ax.legend()

plt.tight_layout()
plt.show()
