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
    eta = 1 * delta / prog
    prog = round(prog, 3)
    return prog, delta, eta, delta+eta