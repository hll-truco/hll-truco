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

parse = lambda f: keep(parse_structured_log(f))

def joint(XYZs:list[list[tuple[int,int,int]]]) -> list[tuple[int,int,int]]:
    X, Y, Z = [], [], []
    delta_x = 0
    delta_z = 0
    for _X, _Y, _Z in XYZs:
        X += [x+delta_x for x in _X]
        Y += _Y
        Z += [z+delta_z for z in _Z]
        delta_x += _X[-1]
        delta_z += _Z[-1]
    return X, Y, Z