package root

import (
	"log/slog"

	"github.com/hll-truco/hll-truco/hll"
)

type State struct {
	global  *hll.HyperLogLogExt
	decoder *hll.HyperLogLogExt
	total   uint64
}

func NewState() *State {
	return &State{
		global:  getNewExt(),
		decoder: getNewExt(),
		total:   0,
	}
}

func getNewExt() *hll.HyperLogLogExt {
	h1, err := hll.NewExt(16)
	if err != nil {
		panic(err)
	}
	return h1
}

func (state *State) Report(delta float64) {
	estimate := state.global.Count()
	slog.Info(
		"REPORT",
		"delta", delta,
		"estimate", estimate,
		"total", state.total)
}
