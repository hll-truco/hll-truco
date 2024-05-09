
### cmds

#### `cmd/ronda-infosets-hll-count`

- for a given time limit (`l`) approximates the number of information sets **AT 
  ROUND LEVEL** using axiomhq's HLL.
- it outputs the estimated number of infosets + total terminal nodes visited
- e.g.: `go run cmd/ronda-infosets-hll-count/main.go -deck=7 -hash=sha160 -limit=120 -report=10`

#### `cmd/stressmem-test`

- a memory benchmark / stress-test / experiment comparing dynamic vs fixed 
  slices and its memory consumption
- program takes parameter `-n` (int) which represent the targeted memory 
  consumption / buffer size in MiB (e.g., `-n-500` ~ 500 MiB buffer)
- e.g.: `go run cmd/stressmem-test/main.go -n 500`
- for `-n=500` + dynamic slice activity monitor shows a memory consumption of 
  about `1 GiB` but go's runtime shows a `HeapAlloc` usage of `520 MiB` (this 
  is expected)
- but then if we cap the mem pool size to 600MiB using env var `GOMEMLIMIT` as
  in `GOMEMLIMIT=600000000 go run cmd/stressmem-test/main.go -n 500` now 
  activity monitor shows a memory consumption of about `564 MiB`

### args

#### supported hashes

- adler32
- sha160
- sha256
- sha512