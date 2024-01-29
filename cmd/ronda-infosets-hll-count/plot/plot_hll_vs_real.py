import matplotlib.pyplot as plt

title = "HLL estimate for mini-Truco with -deck=7 -info=InfosetRondaBase -hash=sha160"
real = 4_690
hll_evol = [4121,4417,4530,4570,4597,4618,4630,4640,4648,4654,4659,4664,4665,4670,4670,4671,4674,4674,4678,4679,4679,4679,4679,4682,4682,4683,4683,4683,4685,4685,4685,4685,4685,4685,4685,4686,4686,4686,4686,4686,4687,4687,4687,4687,4687,4687,4687,4687,4688,4688,4688,4689,4689,4689,4689,4689,4689,4689,4689,]

fig, ax = plt.subplots(1,1, figsize=(10,5))

ax.plot(list(range(len(hll_evol))), hll_evol, '-', linewidth=0.8, label='hll 1min')

ax.axhline(y=(real), color='r', linestyle='--', linewidth=0.5, alpha=0.5, label='real')

ax.set_xlabel('time (s)')
ax.set_ylabel('Estimated cardinality')
ax.set_title(title)
ax.legend()

plt.tight_layout()
plt.show()
