package worker

import (
	"encoding/base64"
	"hash"
	"log/slog"
	"math/big"
	"strings"
	"time"

	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/hll"
	"github.com/hll-truco/hll-truco/hll-dist-http/root/state"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

type Worker struct {
	// data
	H            *hll.HyperLogLogExt
	NodesVisited uint64
	GamesPlayed  uint64
	TotalLocal   *big.Float
	//
	start       time.Time
	Limit       time.Duration
	rootBaseURL string
	allowMazo   bool
}

func NewWorker(
	rootBaseURL string,
	limit time.Duration,
	precision uint8,
	allowMazo bool,
) *Worker {
	h, _ := hll.NewExt(precision)
	if !strings.HasPrefix(rootBaseURL, "http://") {
		rootBaseURL = "http://" + rootBaseURL
	}
	return &Worker{
		H:            h,
		NodesVisited: 0,
		GamesPlayed:  0,
		TotalLocal:   nil,
		start:        time.Now(),
		Limit:        limit,
		allowMazo:    allowMazo,
		rootBaseURL:  rootBaseURL,
	}
}

func (w *Worker) TimeSinceStarted() time.Duration {
	return time.Since(w.start)
}

func (w *Worker) incGamesPlayed() {
	w.GamesPlayed++
	w.NodesVisited++
}

func (w *Worker) checkpoint(printer *utils.CronoPrinter) {
	if printer.ShouldPrint() {
		e := w.H.CountBigDynm()
		delta := printer.Check().Seconds()
		slog.Info("REPORT", "delta", delta, "estimate", e)

		if w.TotalLocal == nil || e.Cmp(w.TotalLocal) == 1 {
			w.TotalLocal = e
			w.SendUpdateRequest()
		}
	}
}

func (w *Worker) TimesUp() bool {
	return w.TimeSinceStarted() > w.Limit
}

func (w *Worker) getInfosetHashBytes(
	p *pdt.Partida,
	infoBuilder *info.Builder,
	hashFn hash.Hash,
) []byte {
	activePlayer := pdt.Rho(p)
	info := infoBuilder.Info(p, activePlayer, nil)
	hash := info.HashBytes(hashFn)
	return hash
}

func (w *Worker) randomAction(p *pdt.Partida) {
	chis := pdt.MetaChis(p, w.allowMazo)
	j := UniformPick(chis)
	j.Hacer(p)
}

func (w *Worker) RandomWalk(
	p *pdt.Partida,
	infoBuilder *info.Builder,
	hashFn hash.Hash,
	printer *utils.CronoPrinter,
) {
	defer w.incGamesPlayed()
	for !p.Terminada() && !w.TimesUp() {
		hash := w.getInfosetHashBytes(p, infoBuilder, hashFn)
		w.H.Add(hash)
		w.randomAction(p)
		w.NodesVisited++
		w.checkpoint(printer)
	}
}

// http

func (w *Worker) SendUpdateRequest() {
	data, err := w.H.GobEncode()
	if err != nil {
		panic(err)
	}
	base64Data := base64.StdEncoding.EncodeToString(data)
	// send
	url := w.rootBaseURL + "/update"
	update := UpdateRequest{Gob: base64Data}
	sendPOSTJsonData(url, update)
}

func (w *Worker) SendReportRequest() {
	data := state.WorkerResult{
		NodesVisited: w.NodesVisited,
		GamesPlayed:  w.GamesPlayed,
		Delta:        uint64(time.Since(w.start).Seconds()),
	}

	// send
	url := w.rootBaseURL + "/report"
	sendPOSTJsonData(url, data)
}
