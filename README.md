
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
- e.g.: `go run cmd/count-infosets-ronda-deterministically/main.go -deck=7 
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

`adler32`, `sha160`, `sha256`, `sha512`

#### supported abstractions

`a1`, `b`, `a2`, `a3`, `null`