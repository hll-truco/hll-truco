package state

import (
	"log/slog"

	"github.com/hll-truco/hll-truco/hll"
)

type State struct {
	Global  *hll.HyperLogLogExt
	Decoder *hll.HyperLogLogExt
	Total   uint64
}

func NewState() *State {
	return &State{
		Global:  GetNewExt(),
		Decoder: GetNewExt(),
		Total:   0,
	}
}

func GetNewExt() *hll.HyperLogLogExt {
	h1, err := hll.NewExt(16)
	if err != nil {
		panic(err)
	}
	return h1
}

func (state *State) Report(delta float64) {
	estimate := state.Global.Count()
	slog.Info(
		"REPORT",
		"delta", delta,
		"estimate", estimate,
		"total", state.Total)
}
