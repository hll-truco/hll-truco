package state

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/hll-truco/hll-truco/hll"
)

type StateCheckpoint struct {
	Gob   string `json:"gob"`
	Total uint64 `json:"total"`
}

type WorkerResult struct {
	NodesVisited uint64 `json:"nodesVisited"`
	GamesPlayed  uint64 `json:"gamesPlayed"`
	Delta        uint64 `json:"delta"`
}

type State struct {
	start   time.Time
	Global  *hll.HyperLogLogExt
	decoder *hll.HyperLogLogExt
	Total   uint64
	// workers' results
	WorkersResults []*WorkerResult
	// multi
	mu *sync.Mutex
}

func NewState(precision uint8) *State {
	return &State{
		start:   time.Now(),
		Global:  GetNewExt(precision),
		decoder: GetNewExt(precision),
		Total:   0,
		// reports
		WorkersResults: make([]*WorkerResult, 0),
		// mutli
		mu: &sync.Mutex{},
	}
}

func GetNewExt(precision uint8) *hll.HyperLogLogExt {
	h1, err := hll.NewExt(precision)
	if err != nil {
		panic(err)
	}
	return h1
}

func (state *State) Merge(data []byte) (bool, error) {
	state.mu.Lock()
	defer state.mu.Unlock()

	if err := state.decoder.GobDecode(data); err != nil {
		return false, err
	}

	bump, err := state.Global.Merge(state.decoder)
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

func (state *State) Estimate() any {
	// estimate := state.Global.Count()
	// estimate := state.Global.CountBig()
	estimate := state.Global.CountBigDynm()
	return estimate
}

func (state *State) Report() {
	state.mu.Lock()
	defer state.mu.Unlock()

	estimate := state.Estimate()

	slog.Info(
		"REPORT",
		"delta", time.Since(state.start).Seconds(),
		"estimate", estimate,
		"total", state.Total)
}

func (state *State) Results() {
	state.mu.Lock()
	defer state.mu.Unlock()

	estimate := state.Estimate()

	slog.Info(
		"RESULTS",
		"finished", time.Since(state.start).Seconds(),
		"finalEstimate", estimate,
		"total", state.Total,
		"reports", state.WorkersResults)
}

func (state *State) Save(filepath string) error {
	data, err := state.Global.GobEncode()
	if err != nil {
		return err
	}
	base64Data := base64.StdEncoding.EncodeToString(data)

	checkpoint := &StateCheckpoint{
		Total: state.Total,
		Gob:   base64Data,
	}

	// Marshal the state to JSON
	bs, err := json.Marshal(checkpoint)
	if err != nil {
		return err
	}

	// Create or open the file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON bs to the file
	_, err = file.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func (state *State) Load(filepath string) error {
	// Create or open the file
	bs, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	checkpoint := &StateCheckpoint{}

	// unmarshal
	err = json.Unmarshal(bs, checkpoint)
	if err != nil {
		return err
	}

	data, err := base64.StdEncoding.DecodeString(checkpoint.Gob)
	if err != nil {
		return err
	}
	state.Merge(data)
	state.Total = checkpoint.Total

	return nil
}
