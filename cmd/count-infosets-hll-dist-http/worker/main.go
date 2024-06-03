package main

import (
	"flag"
	"hash"
	"log/slog"
	"os"
	"time"

	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/hll-dist-http/worker"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

// flags/parametros:
var (
	// deckSizeFlag = flag.Int("deck", 7, "Deck size")
	nFlag       = flag.Int("n", 2, "Number of players <2,4,6>")
	envidoFlag  = flag.Int("e", 1, "Envido limit (default 1)")
	ptsFlag     = flag.Int("p", 40, "Game points limit")
	absIDFlag   = flag.String("abs", "a1", "Abstractor ID")
	infosetFlag = flag.String("info", "InfosetRondaBase", "Infoset impl. to use")
	hashIDFlag  = flag.String("hash", "sha160", "Infoset hashing function")
	limitFlag   = flag.Int("limit", 60, "Run time limit (in seconds) (default 1m)")
	reportFlag  = flag.Int("report", 10, "Delta (in seconds) for printing log msgs")
	rootFlag    = flag.String("root", "localhost:8080", "HTTP root server")
)

var (
	// gameplay params
	n         int
	limEnvite int
	pts       pdt.Puntuacion
	azules    = []string{"Alice", "Ariana", "Annie"}
	rojos     = []string{"Bob", "Ben", "Bill"}
	// hll params
	infoBuilder *info.Builder       = nil
	verbose     bool                = true
	printer     *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
	hashFn      hash.Hash           = nil
	// worker
	w *worker.Worker = nil
)

func init() {
	flag.Parse()

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// dump args
	slog.Info(
		"START",
		"n", *nFlag,
		"envido", *envidoFlag,
		"pts", *ptsFlag,
		"absId", *absIDFlag,
		"infoset", *infosetFlag,
		"hash", *hashIDFlag,
		"limit", *limitFlag,
		"report", *reportFlag,
		"root", *rootFlag)

	// gameplay vars
	n = *nFlag
	limEnvite = *envidoFlag
	pts = pdt.Puntuacion(*ptsFlag)

	// hll
	// hardcode "sha160" to avoid panics; we will use hashFn instead anyways
	infoBuilder = info.BuilderFactory("sha160", *infosetFlag, *absIDFlag)
	hashFn = worker.ParseHashFn(*hashIDFlag)
	printer = utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))
	limit := time.Second * time.Duration(*limitFlag)

	// worker
	w = worker.NewWorker(*rootFlag, limit)
}

func main() {
	for w.TimeSinceStarted() < w.Limit {
		p, _ := pdt.NuevaPartida(
			pts,
			azules[:n>>1],
			rojos[:n>>1],
			limEnvite,
			verbose)
		// p, _ := pdt.NuevaPartida(
		// 	pts,
		// 	true,           // isMini
		// 	utils.Deck(14), // decksize
		// 	azules[:n>>1],
		// 	rojos[:n>>1],
		// 	limEnvite,
		// 	verbose)
		w.RandomWalk(p, infoBuilder, hashFn, printer)
	}

	// a final update + report
	w.SendUpdateRequest()
	w.SendReportRequest()

	slog.Info(
		"RESULTS",
		"finalEstimate", w.H.CountBig(),
		"nodesVisited:", w.NodesVisited,
		"gamesPlayed", w.GamesPlayed,
		"finished", w.TimeSinceStarted().Seconds())
}

/*

go run cmd/count-infosets-hll-dist-http/worker/*.go n", -n=2 -e=1 -p=20 -hash=sha3 -abs=null -report=5 -limit=20 -root=http://localhost:8080 && curl -X GET http://localhost:8080/exit

*/
