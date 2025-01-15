package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"time"

	// "github.com/axiomhq/hyperloglog"
	"github.com/clarkduvall/hyperloglog"
	"github.com/filevich/truco-ai/info"
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

	// actual size for -deck=14 is 248_732
	// we will mark 1% of the total population: 248_732 * 0.01 = 2_487
	// we will capture 10% of the total population: 248_732 * 0.1 = 24_873
	markedFlag    = flag.Int("marked", 2_487, "Number of marked elements")
	capturedFlag  = flag.Int("captured", 248732, "Number of captured elements")
	allowMazoFlag = flag.Bool("mazo", true, "Allow mazo?")
)

var (
	deck        []int               = nil //lint:ignore U1000 <your reason here>
	infoBuilder *info.Builder       = nil
	verbose     bool                = true
	terminals   uint64              = 0
	printer     *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
	// axiom            = hyperloglog.New16()
	// axiom            = hyperloglog.New16NoSparse()
	// h, _                = hyperloglog.New(16)
	// h, _                = hyperloglog.New(4)
	h, _                = hyperloglog.NewPlus(16)
	start time.Time     = time.Now()
	limit time.Duration = 0
)

// type fakeHash32 []byte
// func (f fakeHash32) Sum32() uint32 {
// 	return binary.LittleEndian.Uint32(f[:4])
// }

type fakeHash32 []byte

func (f fakeHash32) Sum64() uint64 {
	return binary.LittleEndian.Uint64(f[:8])
}

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
		"markedFlag", *markedFlag,
		"capturedFlag", *capturedFlag,
		"allowMazoFlag", *allowMazoFlag,
	)

	deck = utils.Deck(*deckSizeFlag)
	infoBuilder = info.BuilderFactory(*hashIDFlag, *infosetFlag, *absIDFlag)
	limit = time.Second * time.Duration(*limitFlag)
	printer = utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))

	// start timer
	start = time.Now()
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

type PartidaFactory func() *pdt.Partida

// returns a map of marked elements and a map of level distribution
func sampleMarked(markedSize int, makePartida PartidaFactory) (map[string]bool, map[int]int) {
	marked := map[string]bool{}
	levelDist := map[int]int{}
	currentLevel := 0

	for len(marked) < markedSize {
		p := makePartida()
		currentLevel = 0

		for !p.Terminada() && len(marked) < markedSize {

			// infoset
			activePlayer := pdt.Rho(p)
			info := infoBuilder.Info(p, activePlayer, nil)
			hashFn := utils.ParseHashFn(*hashIDFlag)
			hash := info.HashBytes(hashFn)
			h := string(hash)

			// use hash
			// check if `h` is in `marked`
			if _, ok := marked[h]; !ok {
				marked[h] = true
				if _, ok := levelDist[currentLevel]; !ok {
					levelDist[currentLevel] = 0
				}
				levelDist[currentLevel]++
			}

			// apply a random move
			chis := pdt.MetaChis(p, *allowMazoFlag)
			j := uniformPick(chis)
			pkts := j.Hacer(p)
			if pdt.IsDone(pkts, true) || p.Terminada() {
				currentLevel = 0
			} else {
				currentLevel++
			}
		}
	}

	return marked, levelDist
}

func capture(
	captureSize int,
	makePartida PartidaFactory,
	marked map[string]bool,
) (
	map[string]bool, // a map of captured elements
	int, // the number of recaptured elements
	map[int]int, // a map of level distribution
) {

	captured := map[string]bool{}
	recaptured := 0
	levelDist := map[int]int{}
	currentLevel := 0

	for len(captured) < captureSize {
		p := makePartida()
		currentLevel = 0

		for !p.Terminada() && len(captured) < captureSize {

			// infoset
			activePlayer := pdt.Rho(p)
			info := infoBuilder.Info(p, activePlayer, nil)
			hashFn := utils.ParseHashFn(*hashIDFlag)
			hash := info.HashBytes(hashFn)
			h := string(hash)

			// use hash
			// check if `h` is in `captured`
			if _, ok := captured[h]; !ok {
				captured[h] = true

				// check if `h` is in `marked`, and thus recaptured
				if _, ok := marked[h]; ok {
					recaptured++
				}

				if _, ok := levelDist[currentLevel]; !ok {
					levelDist[currentLevel] = 0
				}
				levelDist[currentLevel]++
			}

			// apply a random move
			chis := pdt.MetaChis(p, *allowMazoFlag)
			j := uniformPick(chis)
			pkts := j.Hacer(p)
			if pdt.IsDone(pkts, true) || p.Terminada() {
				currentLevel = 0
			} else {
				currentLevel++
			}
		}
	}

	return captured, recaptured, levelDist
}

func main() {
	n := 4
	limEnvite := 1
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}

	os.Setenv("DECK", fmt.Sprintf("%d", *deckSizeFlag))

	makePartida := func() *pdt.Partida {
		return utils.NuevaPartida(
			pdt.A40,
			azules[:n>>1],
			rojos[:n>>1],
			limEnvite,
			verbose,
		)
	}

	// let's mark some elements
	marked, levelDist := sampleMarked(*markedFlag, makePartida)
	slog.Info(
		"MARKED_DONE",
		"got", len(marked),
		"wanted", *markedFlag,
		"marked", len(marked),
		"levelDist", levelDist,
		"finished", time.Since(start).Seconds(),
	)

	// let's capture some elements
	captured, recaptured, levelDist := capture(*capturedFlag, makePartida, marked)

	slog.Info(
		"CAPTURE_DONE",
		"wanted", *capturedFlag,
		"marked", len(marked),
		"captured", len(captured),
		"recaptured", recaptured,
		"levelDist", levelDist,
		"finished", time.Since(start).Seconds(),
	)

	// calculate N using int64
	// N := len(marked) * len(captured) / recaptured

	// calculate N using `big`
	precision := uint(4096) // 4096 bits for a max value of 10^1233
	N_big := utils.EstimatePopulation(len(marked), len(captured), recaptured, precision)

	slog.Info("RESULTS", "N_big", N_big)
}
