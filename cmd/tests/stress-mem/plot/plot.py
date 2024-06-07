import matplotlib.pyplot as plt
import numpy as np

# Assuming you have 4 arrays of numbers
y1 = [1830,1986,5586,3879,6061,13639,17048,9471,21310,]
y2 = [5072,9918,15505,19384,30295,37873,47344,47344,59183,]
y3 = [3332,5387,8592,12598,12598,20424,20424,20424,32646,]
# y4 = [23,27,28,30,32,32,33,34,34,]

plt.figure(figsize=(10,6))

# Generate x values
x = range(len(y1))

# Plot each array
plt.plot(x, y1, label='HeapAlloc')
plt.plot(x, y2, label='TotalAlloc')
plt.plot(x, y3, label='Sys')
# plt.plot(x, y4, label='NumGC')

plt.tight_layout()

# Add a legend
plt.legend()

# Show the plot
plt.show()
