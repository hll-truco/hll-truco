

## Goal

Compare speed/evolution of:
 - Determinisitic
 - HLL
 - Lincoln-Petersen + HLL

## Replicate

1. Deterministic

```bash
go run cmd/count-infosets/ronda/deterministically/main.go -deck=9 -hash=sha160 -info=InfosetRondaBase -abs=null -track=true -report=10
```

2. HLL

```bash
# axiom
go run cmd/count-infosets/ronda/hll-axiom/main.go -hash=sha160 -deck=9 -abs=null -report=10 -limit=1400
# clark
go run cmd/count-infosets/ronda/hll-clarkduvall/main.go -hash=sha160 -deck=9 -abs=null -report=10 -limit=1400
# ours
go run cmd/count-infosets/ronda/hll-hll/main.go -hash=sha160 -deck=9 -abs=null -report=10 -limit=1400 -precision=6
```

3. Lincoln-Petersen + HLL

```bash
% todo
```

