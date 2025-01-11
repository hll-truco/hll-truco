import matplotlib.pyplot as plt

# 1%
dist_2p_marked2487 = {"0":915,"1":613,"2":383,"3":267,"4":165,"5":90,"6":39,"7":14,"8":1}
# 10%
dist_2p_marked24870 = {"0":2944,"1":7968,"2":5471,"3":3707,"4":2333,"5":1420,"6":700,"7":252,"8":62,"9":13}
# 50%
dist_2p_marked124366 = {"0":3082,"1":32459,"2":32417,"3":24470,"4":15820,"5":8919,"6":4708,"7":1876,"8":447,"9":168}
# 90%
dist_2p_marked223858 = {"0":3082,"1":39386,"2":60162,"3":53047,"4":36412,"5":18923,"6":8037,"7":3524,"8":975,"9":310}

# 1%
dist_2p_marked2487_mazo_false = {"0":388,"1":407,"2":390,"3":351,"4":313,"5":256,"6":189,"7":105,"8":71,"9":17}
# 10%
dist_2p_marked24870_mazo_false = {"0":2346,"1":4341,"2":4510,"3":3931,"4":3312,"5":2633,"6":1954,"7":1003,"8":566,"9":274}
# 50%
dist_2p_marked124366_mazo_false = {"0":3081,"1":24266,"2":26735,"3":24638,"4":19841,"5":12727,"6":7825,"7":3765,"8":1071,"9":417}
# 90%
dist_2p_marked223858_mazo_false = {"0":3082,"1":39224,"2":55757,"3":49933,"4":37708,"5":21709,"6":10363,"7":4473,"8":1164,"9":445}

fig, ax = plt.subplots(4, 2, figsize=(10, 8))

"""
-mazo=true
"""

# Plot the first distribution (-mazo=true)
ax[0][0].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=2487 ~ 1%)")
ax[0][0].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[0][0].set_ylabel('Unique Hashes Count')
ax[0][0].bar(list(dist_2p_marked2487.keys()), list(dist_2p_marked2487.values()))

# Plot the second distribution (-mazo=true)
ax[1][0].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=24870) ~ 10%")
ax[1][0].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[1][0].set_ylabel('Unique Hashes Count')
ax[1][0].bar(list(dist_2p_marked24870.keys()), list(dist_2p_marked24870.values()))

# Plot the third distribution (-mazo=true)
ax[2][0].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=124366) ~ 50%")
ax[2][0].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[2][0].set_ylabel('Unique Hashes Count')
ax[2][0].bar(list(dist_2p_marked124366.keys()), list(dist_2p_marked124366.values()))

# Plot the fourth distribution (-mazo=true)
ax[3][0].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=223858) ~ 90%")
ax[3][0].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[3][0].set_ylabel('Unique Hashes Count')
ax[3][0].bar(list(dist_2p_marked223858.keys()), list(dist_2p_marked223858.values()))

"""
-mazo=false
"""

# Plot the first distribution (-mazo=true)
ax[0][1].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=2487 ~ 1%)")
ax[0][1].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[0][1].set_ylabel('Unique Hashes Count')
ax[0][1].bar(list(dist_2p_marked2487_mazo_false.keys()), list(dist_2p_marked2487_mazo_false.values()))

# Plot the second distribution (-mazo=true)
ax[1][1].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=24870) ~ 10%")
ax[1][1].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[1][1].set_ylabel('Unique Hashes Count')
ax[1][1].bar(list(dist_2p_marked24870_mazo_false.keys()), list(dist_2p_marked24870_mazo_false.values()))

# Plot the third distribution (-mazo=true)
ax[2][1].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=124366) ~ 50%")
ax[2][1].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[2][1].set_ylabel('Unique Hashes Count')
ax[2][1].bar(list(dist_2p_marked124366_mazo_false.keys()), list(dist_2p_marked124366_mazo_false.values()))

# Plot the fourth distribution (-mazo=true)
ax[3][1].set_title("Distribution of Infoset's Count vs Infoset's depth\n (-deck=14 -marked=223858) ~ 90%")
ax[3][1].set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax[3][1].set_ylabel('Unique Hashes Count')
ax[3][1].bar(list(dist_2p_marked223858_mazo_false.keys()), list(dist_2p_marked223858_mazo_false.values()))

plt.tight_layout()
plt.show()

fig, ax = plt.subplots(figsize=(10, 8))

# Plot all distributions on the same plot
ax.set_title("Distribution of Infoset's Count vs Infoset's depth")
ax.set_xlabel("Distance from the root node at which the infoset's hash is calculated")
ax.set_ylabel('Unique Hashes Count')

# Plot the distributions in the specified order with different colors
ax.bar(list(dist_2p_marked223858.keys()), list(dist_2p_marked223858.values()), color='blue', alpha=0.6, label='-marked=223858 ~ 90%')
ax.bar(list(dist_2p_marked124366.keys()), list(dist_2p_marked124366.values()), color='green', alpha=0.6, label='-marked=124366 ~ 50%')
ax.bar(list(dist_2p_marked24870.keys()), list(dist_2p_marked24870.values()), color='orange', alpha=0.6, label='-marked=24870 ~ 10%')
ax.bar(list(dist_2p_marked2487.keys()), list(dist_2p_marked2487.values()), color='red', alpha=0.6, label='-marked=2487 ~ 1%')

ax.legend()
plt.tight_layout()
plt.show()