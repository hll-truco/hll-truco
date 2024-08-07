package main

import (
	"flag"
	"fmt"
	"hash"
	"log/slog"
	"os"
	"time"

	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/hll-dist-http/worker"
	"github.com/hll-truco/hll-truco/utils"
	gotruco "github.com/truquito/gotruco"
	"github.com/truquito/gotruco/pdt"
)

// for game-level infosets use `-info=InfosetPartidaXXLarge`

// flags/parametros:
var (
	deckSizeFlag  = flag.Int("deck", 14, "Deck size")
	nFlag         = flag.Int("n", 2, "Number of players <2,4,6>")
	envidoFlag    = flag.Int("e", 1, "Envido limit (default 1)")
	ptsFlag       = flag.Int("p", 40, "Game points limit")
	absIDFlag     = flag.String("abs", "a1", "Abstractor ID")
	infosetFlag   = flag.String("info", "InfosetRondaBase", "Infoset impl. to use")
	hashIDFlag    = flag.String("hash", "sha160", "Infoset hashing function")
	limitFlag     = flag.Int("limit", 60, "Run time limit (in seconds) (default 1m)")
	allowMazoFlag = flag.Bool("mazo", true, "Allow 'Mazo' action in Chi vector")
	reportFlag    = flag.Int("report", 10, "Delta (in seconds) for printing log msgs")
	rootFlag      = flag.String("root", "localhost:8080", "HTTP root server")
	precisionFlag = flag.Int("precision", 16, "HLL precision (defaults to 16)")
	resumeFlag    = flag.String("resume", "", "Full path to bob file to load on start")
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
	printer     *utils.CronoPrinter = nil
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
		"allowMazo", *allowMazoFlag,
		"report", *reportFlag,
		"root", *rootFlag,
		"precision", *precisionFlag,
		"resume", *resumeFlag)

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
	w = worker.NewWorker(
		*rootFlag,
		limit,
		uint8(*precisionFlag),
		*allowMazoFlag)

	if len(*resumeFlag) > 0 {
		w.Load(*resumeFlag)
		slog.Info(
			"LOADED",
			"file", *resumeFlag,
			"estimate", w.H.CountBigDynm())
	}
}

func main() {
	os.Setenv("DECK", fmt.Sprintf("%d", *deckSizeFlag))
	for w.TimeSinceStarted() < w.Limit {
		p := utils.NuevaPartida(pts, azules[:n>>1], rojos[:n>>1], limEnvite, verbose)
		w.RandomWalk(p, infoBuilder, hashFn, printer)
	}

	// a final update + report
	w.SendUpdateRequest()
	w.SendReportRequest()

	slog.Info(
		"RESULTS",
		"trucoVersion", gotruco.VERSION,
		"finalEstimate", w.H.CountBigDynm(),
		"nodesVisited:", w.NodesVisited,
		"gamesPlayed", w.GamesPlayed,
		"finished", w.TimeSinceStarted().Seconds())
}
