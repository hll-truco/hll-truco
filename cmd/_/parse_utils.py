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
        list[int], # deltas: time in seconds
        list[float], # estimates: hll estimates
        list[int] # total_msgs: totals
    ]:
    
    deltas, estimates, total_msgs = [], [], []
    for l in lines:
        match l["msg"]:
            case "REPORT":
                deltas += [float(l.get("delta", 0))]
                estimates += [float(l.get("estimate", 0))]
                total_msgs += [float(l.get("total", 0))]
            case "RESULTS":
                deltas += [float(l.get("finished", 0))]
                estimates += [float(l.get("finalEstimate", 0))]
                total_msgs += [float(l.get("total", 0))]

    if normalize:
        for i in range(len(deltas)):
            deltas[i] -= deltas[0]

    return deltas, estimates, total_msgs