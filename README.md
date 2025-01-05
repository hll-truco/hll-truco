
## cmds

#### `cmd/count-infosets-ronda-hll-axiom`

- for a given time limit (`l`) approximates the number of information sets **AT 
  ROUND LEVEL** using axiomhq's HLL.
- it outputs the estimated number of infosets + total terminal nodes visited
- e.g.: `go run cmd/count-infosets-ronda-hll-axiom/main.go -deck=7 -hash=sha160 -limit
  =120 -report=10`

#### `cmd/stress-mem`

- a memory benchmark / stress / experiment comparing dynamic vs fixed 
  slices and its memory consumption
- program takes parameter `-n` (int) which represent the targeted memory 
  consumption / buffer size in MiB (e.g., `-n-500` ~ 500 MiB buffer)
- e.g.: `go run cmd/stress-mem/main.go -n 500`
- for `-n=500` + dynamic slice activity monitor shows a memory consumption of 
  about `1 GiB` but go's runtime shows a `HeapAlloc` usage of `520 MiB` (this 
  is expected)
- but then if we cap the mem pool size to 600MiB using env var `GOMEMLIMIT` as
  in `GOMEMLIMIT=600000000 go run cmd/stress-mem/main.go -n 500` now 
  activity monitor shows a memory consumption of about `564 MiB`

#### `cmd/count-infosets-ronda-deterministically`

- deterministically count the number of infoset for a given deck size **at round
  level**.
- e.g.: `go run cmd/count-infosets/ronda/deterministically/main.go -deck=7 
  -hash=sha160 -info=InfosetRondaBase -abs=a1 -track=true -report=10`

#### `cmd/count-infosets-partida-deterministically`

- deterministically count the number of infoset for a given deck size **at game
  level**.
- WARNING: this can take a **LOT** of time
- e.g.: `go run cmd/count-infosets-partida-deterministically/main.go`

#### `cmd/hll-py-example`

- hyperloglog python experimental implementation
- run as `python cmd/hll-py-example/hll1.py`

#### `cmd/hll-axiom-example`

- hyperloglog example in go using the axiom lib
- run as `go run cmd/hll-axiom-example/main.go`

#### `cmd/check-hay-flor`

- for a given set of cards (i.e., a deck) it returns weather there's at least
  one flor (i.e., a combination of 3 cards + 1 muestra so that it's a flor)
- run as `go run cmd/check-hay-flor/main.go`

#### `cmd/count-root-edges-deterministically`

- for a given deck size (harcoded) it returns the total possible number of edges
  comming out of the root chance node.
- it should obey the equation $...$
- run as `go run cmd/count-root-edges-deterministically/main.go`

#### `cmd/count-infosets-hll-http-dist/`

- root server: `go run cmd/count-infosets-hll-http-dist/root/cmd/main.go -port=8080 | tee logs/hll-dist-http/http-w2-d14-anull-hsha3_6.log`
- single worker `go run cmd/count-infosets-hll-http-dist/worker/*.go -hash=sha3 -deck=14 -abs=null -report=10 -limit=600 -root=http://localhost:8080`
- single worker + exit: `curl -X GET http://localhost:8080/exit`
- multiple workers: `for i in {1..16}; do go run cmd/count-infosets-hll-http-dist/worker/*.go -hash=sha3 -deck=14 -abs=null -report=10 -limit=600 -root=http://localhost:8080 &>/dev/null & done`


## notes

#### supported hashes

`adler32`, `sha160`, `sha256`, `sha512`, `sha3`

#### supported abstractions

`a1`, `b`, `a2`, `a3`, `null`

### supported infosets

`InfosetRondaBase`, `InfosetRondaLarge`, `InfosetRondaXLarge`, `InfosetRondaXXLarge`, `InfosetPartidaXXLarge`

## Cluster

### Upload

`rsync -avz --copy-links --progress -e 'ssh -p 10022 -i ~/.ssh/id_rsa' --exclude '__pycache__/' --exclude '*.out' --exclude '*.log' --exclude '.git/' ~/Workspace/facu/hll-truco 'cluster.uy:Workspace/facu'`

### Download

`rsync -avz -e 'ssh -p 10022 -i ~/.ssh/id_rsa' 'cluster.uy:batches/out/hll-http/hllroot.3629053.out' logs/hll-dist-http/slurm`

## Build (local or remote) for x86-linux (i.e., slurm-cluster)

`GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/http/worker cmd/count-infosets-hll-dist-http/worker/main.go`

`GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/http/root cmd/count-infosets-hll-dist-http/root/main.go`

## run

Make sure both the logs dirs `--output` and `--error` exist before execution!

```bash
mkdir -p ~/batches/out/hll-http/2p/E1P40AnullIipxxlW256 && sbatch -J hllroot1 --time=3-00:00:00 --begin=now+0 ~/Workspace/facu/hll-truco/hll-truco/sbatch/http/root.sbatch 10 16
```

the positional args are:

1. Report (in secs)
2. HLL precision

```bash
jobname=hllworkers1 && \
rootname=hllroot1 && \
rootid=$(squeue -u $(whoami) --name=${rootname} -h -o "%i") && \
roothost=$(squeue -u $(whoami) --name=${rootname} --states=R -h -o "%N:%k") && \
sbatch \
-J ${jobname} \
--time=3-00:00:00 \
--dependency=after:${rootid} \
~/Workspace/facu/hll-truco/hll-truco/sbatch/http/workers.sbatch \
2 1 40 null InfosetPartidaXXLarge sha3 252000 10 http://${roothost} 16 true
```

the positional args are:

1. Number of players `<2,4,6>`
2. Envido limit (default 1)
3. Game points limit
4. Abstractor ID `<a1, b, a2, a3, null>`
5. Infoset impl. to use
6. Infoset hashing function
7. Run time limit (in seconds) (600 = 10min, 259.200 = 3 days)
8. Delta (in seconds) for printing log msgs
9. HTTP root server
10. HLL precision
11. Allow 'Mazo' action in Chi vector (`true` by default)
