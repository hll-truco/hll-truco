package main

import (
	"encoding/base64"
	"flag"
	"hash"
	"log/slog"
	"math/rand"
	"os"
	"time"

	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/hll"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

// flags/parametros:
var (
	deckSizeFlag = flag.Int("deck", 7, "Deck size")
	absIDFlag    = flag.String("abs", "a1", "Abstractor ID")
	infosetFlag  = flag.String("info", "InfosetRondaBase", "Infoset impl. to use")
	hashIDFlag   = flag.String("hash", "sha160", "Infoset hashing function")
	limitFlag    = flag.Int("limit", 60, "Run time limit (in seconds) (default 1m)")
	reportFlag   = flag.Int("report", 10, "Delta (in seconds) for printing log msgs")
	rootFlag     = flag.String("root", "localhost:8080", "HTTP root server")
)

var (
	deck        []int               = nil
	infoBuilder *info.Builder       = nil
	verbose     bool                = true
	terminals   uint64              = 0
	printer     *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
	h, _                            = hll.NewExt(16)
	start       time.Time           = time.Now()
	limit       time.Duration       = 0
	hashFn      hash.Hash           = nil
	totalLocal  uint64              = 0
)

func init() {
	flag.Parse()

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info(
		"START",
		"deckSize", *deckSizeFlag,
		"absId", *absIDFlag,
		"infoset", *infosetFlag,
		"hash", *hashIDFlag,
		"limitFlag", *limitFlag,
		"reportFlag", *reportFlag,
		"rootFlag", *rootFlag)

	deck = utils.Deck(*deckSizeFlag)
	// hardcode "sha160" to avoid panics; we will use hashFn anyways
	infoBuilder = info.BuilderFactory("sha160", *infosetFlag, *absIDFlag)
	limit = time.Second * time.Duration(*limitFlag)
	printer = utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))

	if *hashIDFlag == "sha3" {
		hashFn = hll.NewSha3Hash(128)
		slog.Warn("USING_SHA3SHAKE", "size", 128)
	} else {
		hashFn = utils.ParseHashFn(*hashIDFlag)
		slog.Warn("USING_FIXED_HASH", "hash", *hashIDFlag)
	}

	// start timer
	start = time.Now()
}

func update() {
	data, err := h.GobEncode()
	if err != nil {
		panic(err)
	}
	base64Data := base64.StdEncoding.EncodeToString(data)
	sendUpdateRequest(*rootFlag, base64Data)
}

func uniformPick(chis [][]pdt.IJugada) pdt.IJugada {
	// hago un flatten del vector chis
	n := len(chis) * 15
	flatten := make([]pdt.IJugada, 0, n)

	for _, chi := range chis {
		flatten = append(flatten, chi...)
	}

	// elijo una jugada al azar
	rfix := rand.Intn(len(flatten))

	return flatten[rfix]
}

func randomWalk(p *pdt.Partida) {
	for {
		if p.Terminada() || time.Since(start) > limit {
			return
		}

		// infoset
		activePlayer := pdt.Rho(p)
		info := infoBuilder.Info(p, activePlayer, nil)

		hash := info.HashBytes(hashFn)
		h.Add(hash)

		chis := pdt.Chis(p)
		j := uniformPick(chis)

		pkts, _ := j.Hacer(p)

		if pdt.IsDone(pkts) || p.Terminada() {
			terminals++
		}

		if printer.ShouldPrint() {
			e := h.Count()
			delta := printer.Check().Seconds()
			slog.Info("REPORT", "delta", delta, "estimate", e)

			if e > totalLocal {
				totalLocal = e
				update()
			}
		}
	}
}

func main() {
	n := 2
	limEnvite := 1
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}

	rand.Seed(time.Now().UnixNano())

	for {
		if time.Since(start) > limit {
			break
		}
		p, _ := pdt.NuevaPartida(
			pdt.A40, // <----- no importa poque la condicion de parada es Ronda
			true,
			deck,
			azules[:n>>1],
			rojos[:n>>1],
			limEnvite, // limiteEnvido
			verbose)
		randomWalk(p)
		// termino la partida o se acabó el tiempo
	}

	// a final update
	update()

	slog.Info(
		"RESULTS",
		"finalEstimate", h.Count(),
		"terminals:", terminals,
		"finished", time.Since(start).Seconds())
}
