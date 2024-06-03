import json
import datetime

parse_date = lambda s: datetime.datetime.strptime(s, '%Y/%m/%d %H:%M:%S')

def progress(
    current,
    goal,
    start:datetime.datetime,
    last_commit:datetime.datetime,
):
    prog = current / goal
    delta = last_commit - start
    eta_total = delta.total_seconds() / prog
    eta_total = datetime.timedelta(seconds=round(eta_total))
    prog = round(prog, 3)
    return prog, delta, eta_total-delta, eta_total

def parse_structured_log(
    logfile: str,       
) -> list[any]:
    with open(logfile) as f:
        return [
            json.loads(line)
            for line in f if line and line.startswith('{')
        ]

def keep(lines :list[dict], normalize=True) -> tuple[
        list[int], # xs: time in seconds
        list[float], # ys: hll estimates
        list[int] # zs: totals
    ]:
    
    xs, ys, zs = [], [], []
    for l in lines:
        match l["msg"]:
            case "REPORT":
                xs += [float(l["delta"])]
                ys += [float(l["estimate"])]
                zs += [float(l["total"])]
            case "RESULTS":
                xs += [float(l["finished"])]
                ys += [float(l["finalEstimate"])]
                zs += [float(l["total"])]

    if normalize:
        for i in range(len(xs)):
            xs[i] -= xs[0]

    return xs, ys, zs