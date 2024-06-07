import matplotlib.pyplot as plt
import numpy as np

# Assuming you have 4 arrays of numbers
y1 = [170,333,520,651,813,1017,1271,1589,1589,1986,1986,1986,2483,2483,5586,3103,3103,3103,3879,3879,3879,3879,3879,4849,4849,4849,4849,4849,10911,6061,6061,6061,6061,6061,6062,13639,7577,7577,7577,7577,7577,7577,7577,7577,9471,9471,9471,9471,9471,9471,9471,9471,9471,9471,21310,11839,11839,]
y2 = [10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10000,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,10001,]

plt.figure(figsize=(10,8))

# Plot each array
plt.plot(range(len(y1)), y1, label='HeapAlloc n=10k dynm')
plt.plot(range(len(y2)), y2, label='HeapAlloc n=10k fixed')

plt.tight_layout()

# Add a legend
plt.legend()

# Show the plot
plt.show()