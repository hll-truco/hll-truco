import matplotlib.pyplot as plt

data_orig = {
    1: 1706,
    2: 2256,
    3: 2583,
    4: 2934,
    5: 3546,
    6: 4223,
    7: 4637,
    8: 4647,
    9: 4736,
    10: 4817,
    11: 4664,
    12: 4761,
    13: 4658,
    14: 4724
}

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(list(data_orig.keys()), list(data_orig.values()), '-', linewidth=0.8, label='original')

ax.set_xticks(list(data_orig.keys()), [str((2**x)) for x in data_orig.keys()])
ax.axhline(y=(4_690), color='r', linestyle='--', linewidth=0.5, alpha=0.5)

ax.set_xlabel('m=2^b')
ax.set_ylabel('Estimated cardinality')
ax.set_title('Estimated cardinality for 4,690 different')
ax.legend()

plt.tight_layout()
plt.show()
