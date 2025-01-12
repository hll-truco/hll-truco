import json
from scipy.stats import pearsonr, spearmanr

# Open the file and parse it as JSON
with open('plots/lincoln/kl-div-evol/kl_div_gridsearch_correlation_-mazo=false.json', 'r') as file:
    data = json.load(file)

# filter data
data = [entry for entry in data if entry["marked"] == entry["captured"]]

# kls
kls = [d["kl"] for d in data]

# calc rel. errors
correct_N = 248732
calc_rel_err = lambda x: abs(x - correct_N) / correct_N
errs = [calc_rel_err(d["N"]) for d in data]

# Print the parsed data
for n,kl,e in zip([d["N"] for d in data], kls, errs):
    print(f"{n=} {kl=} {e=}")

# Calculate Pearson correlation
pearson_corr, _ = pearsonr(kls, errs)
print(f"Pearson correlation: {pearson_corr}")

# Calculate Spearman correlation
spearman_corr, _ = spearmanr(kls, errs)
print(f"Spearman correlation: {spearman_corr}")


"""
n=22329 kl=0.2093290374872201 e=0.9102286798642716
n=40823 kl=0.16715206482322517 e=0.8358755608446038
n=64103 kl=0.12569616040456694 e=0.7422808484634064
n=100132 kl=0.07192752823278833 e=0.5974301658009423
n=126390 kl=0.0449069213196529 e=0.49186272775517426
n=148531 kl=0.030442933934064352 e=0.40284724120740395
n=169446 kl=0.02154050042590266 e=0.3187607545470627
n=189129 kl=0.0160694165416218 e=0.23962739012270234
n=205881 kl=0.012508693131902563 e=0.17227779296592316
n=221280 kl=0.007722595926590106 e=0.11036778540758728
n=236161 kl=0.002445360326837599 e=0.05054034060756155
Pearson correlation: 0.9448586069361133
Spearman correlation: 1.0
"""
