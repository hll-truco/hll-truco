
n = [
    100_000,
    1_000_000,
    2_000_000,
    10_000_000,
    100_000_000,
]

mac = [25.8,153,283,1140,11270]
sys = [28,188,369,1361,11549]
mac_gc = [12,90,137,1030,6650]
sys_gc = [28,188,369,1361,11549]

import matplotlib.pyplot as plt

plt.plot(n, mac, label='Mac', marker='o')
plt.plot(n, sys, label='Sys', marker='o')
plt.plot(n, mac_gc, label='Mac GC', marker='o')
plt.plot(n, sys_gc, label='Sys GC', marker='o')

plt.xlabel('Number of Elements')
plt.ylabel('Memory Consumption (MB)')
plt.title('Memory Consumption Comparison')
plt.legend()
# plt.grid(True)
# plt.xscale('log')
# plt.yscale('log')

plt.show()

# rule is:
# 100M infosets = 11549 MiB ~ 11.6 GiB
# 100M infosets =  7000 MiB ~    7 GiB (con `GOMEMLIMIT=7GiB`)
# por eso, pensar que:
# 100M ~ 8 GiB
# Luego, si tenemos 503 GiB max
# Los multiples de 8 son 496, 504
# Tomando (496 / 8) * 100M info = 6200M info
# Haciendo el split con un ratio 100:120 tenemos que:
# r = 120/100 = 1.2
# 6200(x) + 6200(x*1.2) = 6200
# (x + x*1.2) = 1
# x(2.2) = 1
# x = .4545
# 
# Luego,
# mark = 6200(.4545) = 2817
# capture = 6200(.4545*1.2) = 3381
# 
# Notar que 2817 + 3381 = 6198
# Redoneando +2
# 2818 + 3382 = 6200
# Luego,
# 3382 / 2818 ~ 1.2001
# 
