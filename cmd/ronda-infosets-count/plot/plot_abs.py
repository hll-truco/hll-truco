# plot data info
info = {
    "infosetc_14d_a2_base_sha1.3249891.out": {
        "label": "a2",
    },
    "infosetc_14d_null_base_sha1.3249893.out": {
        "label": "None",
    },
    "infosetc_14d_b_base_sha1.3250781.out": {
        "label": "b",
    },
    "infosetc_14d_a3_base_sha1.3249892.out": {
        "label": "a3",
    },
    "infosetc_14d_a1_base_sha1.3249890.out": {
        "label": "a1",
    }
}

# example expected data structure
data = {
    "infosetc_14d_a2_base_sha1.3249891.out": {
        "done": [984,1906,2759,],
        "count": [9052,12568,13929,]
    },
}

# fetch the data
# `rsync -avz -e 'ssh -p 10022' 'juan.filevich@cluster.uy:~/batches/out/infosetc_*.out' /tmp/infosetc`
# parse it
# `python cmd/ronda-infosets-count/plot/parse_abs.py -d /tmp/infosetc`
# read it
with open('/tmp/infosetc/result.json', 'r') as f:
    import json
    data = json.loads(f.read())

# show only
show_only = [
    # "infosetc_14d_b_base_sha1.3250781.out",
    # "infosetc_14d_a1_base_sha1.3249890.out"
]

if len(show_only): data = {k:v for k,v in data.items() if k in show_only}

import matplotlib.pyplot as plt

fig, axs = plt.subplots(1, 2, figsize=(10, 5))
fig.suptitle("ronda-infosets-count 1c 2p 1' d=14 i=InfosetRondaBase h=sha1")

axs[0].set_title('root chance node edges completed')
for file,d in data.items():
    ys = d['done']
    l = info[file]['label']
    axs[0].plot(range(len(ys)), ys, label=l)
axs[0].legend()

axs[1].set_title('infosets count')
for file,d in data.items():
    ys = d['count']
    l = info[file]['label']
    axs[1].plot(range(len(ys)), ys, label=l)
axs[1].legend()

plt.tight_layout()
plt.show()

# progress
import matplotlib.pyplot as plt

# Plot
fig, ax = plt.subplots(1, 1, figsize=(10, 5))

labels = [k for k in data.keys()]
deltas = [data[k]['delta'] for k in labels]
etas = [data[k]['eta'] for k in labels]
labels = [k[:k.index('.')] for k in labels]

ax.barh(labels, deltas, label='done')
ax.barh(labels, etas, left=deltas, label='100%')

# limit
import datetime
for i in range(1,10+1):
    ax.axvline(
        x=datetime.timedelta(days=i).total_seconds(),
        linewidth=0.5,
        alpha=0.2,
        color='black')

ax.axvline(
    x=datetime.timedelta(days=5).total_seconds(),
    linewidth=0.85,
    alpha=1,
    label='cluster lim',
    color='red')

max_secs = max([d+e for d,e in zip(deltas,etas)])
max_delta = str(datetime.timedelta(seconds=max_secs))

ax.axvline(
    x=max_secs,
    linewidth=0.85,
    label=f"{max_delta[:max_delta.rfind(':')]}",
    alpha=1,
    color='purple')

ax.set_xlabel('time (s)')
ax.set_title('Progress')
ax.legend()

plt.tight_layout()
plt.show()
