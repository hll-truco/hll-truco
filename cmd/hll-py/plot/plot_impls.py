import matplotlib.pyplot as plt

data = {
    1: 7860608,
    2: 12707390,
    3: 11374086,
    4: 9635043,
    5: 9847343,
    6: 9854796,
    7: 9933916,
    8: 10246900,
    9: 10483361,
    10: 10373610,
    11: 10123535,
    12: 10017492,
    13: 10077054,
    14: 10097461
}

data2 = {
    1: 6288486,
    2: 8471593,
    3: 7755059,
    4: 8759130,
    5: 10596787,
    6: 11322823,
    7: 10507271,
    8: 9973279,
    9: 9963834,
    10: 9981375,
    11: 9895565,
    12: 10035686,
    13: 10098659,
    14: 10025792,
}

data3 = {
    1: 7860608,
    2: 9530542,
    3: 10663206,
    4: 13896697,
    5: 10925519,
    6: 10133550,
    7: 10295981,
    8: 10075386,
    9: 9916920,
    10: 10142008,
    11: 10563907,
    12: 10134355,
    13: 9966376,
    14: 10007551,
}

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(list(data.keys()), list(data.values()), '-o', label='original')
ax.plot(list(data2.keys()), list(data2.values()), '-o', label='jp orig')
ax.plot(list(data3.keys()), list(data3.values()), '-o', label='jp parallel')

ax.set_xticks(list(data.keys()), [str((2**x)) for x in data.keys()])
ax.axhline(y=(10000000-10), color='r', linestyle='--')
ax.set_xlabel('m=2^b')
ax.set_ylabel('Estimated cardinality for 1M distinct elements')
ax.set_title('Estimated cardinality for different values of b')
ax.legend()

plt.tight_layout()
plt.show()
