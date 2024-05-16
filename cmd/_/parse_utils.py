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

def keep(lines :list[dict]) -> tuple[list[int],list[any]]:
    xs = []
    ys = []
    for l in lines:
        match l["msg"]:
            case "REPORT":
                xs += [l["delta"]]
                ys += [l["estimate"]]
            case "RESULTS":
                xs += [l["finished"]]
                ys += [l["finalEstimate"]]
    return xs, ys