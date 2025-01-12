
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

plt.plot(n, mac, label='Mac')
plt.plot(n, sys, label='Sys')
plt.plot(n, mac_gc, label='Mac GC')
plt.plot(n, sys_gc, label='Sys GC')

plt.xlabel('Number of Elements')
plt.ylabel('Memory Consumption (MB)')
plt.title('Memory Consumption Comparison')
plt.legend()
# plt.grid(True)
# plt.xscale('log')
# plt.yscale('log')

plt.show()