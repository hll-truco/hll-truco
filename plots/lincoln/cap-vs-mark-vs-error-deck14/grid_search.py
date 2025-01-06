import subprocess

total = 248732

makred_percentages = [1, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90]

params = []

for marked_percentage in makred_percentages:
    for captured_percentage in makred_percentages:
        if captured_percentage < marked_percentage: continue
        p = marked_percentage / 100
        q = captured_percentage / 100
        marked = round(total * p)
        captured = round(total * q)
        # print(f"{marked_percentage=} {captured_percentage} -> {marked=} {captured=}")
        params += [((marked, captured))]

for i, (marked,captured) in enumerate(params):
    print(f"using {marked=} {captured=} ({i+1}/{len(params)})")
    # execute this shell command `go run cmd/count-infosets/ronda/lincoln/main.go -hash=sha160 -deck=14 -abs=null -report=1 -limit=1400 -marked=2487 -captured=24870`
    command = [
        "go",
        # "run", "cmd/count-infosets/ronda/lincoln/main.go",
        "run", "cmd/count-infosets/ronda/lincoln/multi-lincoln/main.go",
        "-hash=sha160", "-deck=14", "-abs=null", "-report=1",
        f"-marked={marked}", f"-captured={captured}"
    ]
    subprocess.run(command)
    print("-"*80)


