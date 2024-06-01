package state

import (
	"log/slog"
	"sync"
	"time"

	"github.com/hll-truco/hll-truco/hll"
)

type WorkerResult struct {
	NodesVisited uint64 `json:"nodesVisited"`
	GamesPlayed  uint64 `json:"gamesPlayed"`
	Delta        uint64 `json:"delta"`
}

type State struct {
	start  time.Time
	Global *hll.HyperLogLogExt
	Total  uint64
	// workers' results
	WorkersResults []*WorkerResult
	// multi
	mu *sync.Mutex
}

func NewState() *State {
	return &State{
		start:  time.Now(),
		Global: GetNewExt(),
		Total:  0,
		// reports
		WorkersResults: make([]*WorkerResult, 0),
		// mutli
		mu: &sync.Mutex{},
	}
}

func GetNewExt() *hll.HyperLogLogExt {
	h1, err := hll.NewExt(16)
	if err != nil {
		panic(err)
	}
	return h1
}

func (state *State) Merge(other *hll.HyperLogLogExt) (bool, error) {
	state.mu.Lock()
	defer state.mu.Unlock()

	bump, err := state.Global.Merge(other)
	if err != nil {
		return false, err
	}

	state.Total++

	return bump, nil
}

func (state *State) AddWorkerResult(r *WorkerResult) {
	state.mu.Lock()
	defer state.mu.Unlock()
	state.WorkersResults = append(state.WorkersResults, r)
}

func (state *State) Report(delta float64) {
	state.mu.Lock()
	defer state.mu.Unlock()

	estimate := state.Global.Count()
	slog.Info(
		"REPORT",
		"delta", delta,
		"estimate", estimate,
		"total", state.Total)
}

func (state *State) Results() {
	state.mu.Lock()
	defer state.mu.Unlock()

	estimate := state.Global.Count()
	slog.Info(
		"RESULTS",
		"delta", time.Since(state.start).Seconds(),
		"reports", state.WorkersResults,
		"estimate", estimate,
		"total", state.Total)
}
