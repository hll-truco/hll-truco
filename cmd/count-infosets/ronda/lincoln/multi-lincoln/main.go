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

// returns a map of marked elements and the total number of marked elements
func sampleMarked(markedSize int, makePartida PartidaFactory) (map[int](map[string]bool), int) {
	marked := map[int](map[string]bool){}
	total := 0
	currentLevel := 0

	for total < markedSize {
		p := makePartida()
		currentLevel = 0

		for !p.Terminada() && total < markedSize {

			// infoset
			activePlayer := pdt.Rho(p)
			info := infoBuilder.Info(p, activePlayer, nil)
			hashFn := utils.ParseHashFn(*hashIDFlag)
			hash := info.HashBytes(hashFn)
			h := string(hash)

			// use hash

			// check if marked[currentLevel] exists
			if _, ok := marked[currentLevel]; !ok {
				marked[currentLevel] = map[string]bool{}
			}

			// check if hash was marked for its level
			if _, ok := marked[currentLevel][h]; !ok {
				marked[currentLevel][h] = true
				total++
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

	return marked, total
}

// returns captured elements and the number of recaptured elements
func capture(
	captureSize int,
	makePartida PartidaFactory,
	marked map[int](map[string]bool),
) (
	captured map[int](map[string]bool),
	recapturedByLevel map[int]int,
) {
	captured = map[int](map[string]bool){}
	recapturedByLevel = map[int]int{}
	total := 0

	for total < captureSize {
		p := makePartida()
		currentLevel := 0

		for !p.Terminada() && total < captureSize {

			// infoset
			activePlayer := pdt.Rho(p)
			info := infoBuilder.Info(p, activePlayer, nil)
			hashFn := utils.ParseHashFn(*hashIDFlag)
			hash := info.HashBytes(hashFn)
			h := string(hash)

			// use hash.

			// check if captured[currentLevel] exists
			if _, ok := captured[currentLevel]; !ok {
				captured[currentLevel] = map[string]bool{}
			}

			// check if `h` is in `captured[currentLevel]`
			if _, ok := captured[currentLevel][h]; !ok {
				captured[currentLevel][h] = true
				total++

				// check if `currentLevel` exists in `marked`
				if _, ok := marked[currentLevel]; !ok {
					marked[currentLevel] = map[string]bool{}
				}

				// check if `h` is in `marked[currentLevel]`, and thus recaptured
				if _, ok := marked[currentLevel][h]; ok {
					if _, ok := recapturedByLevel[currentLevel]; !ok {
						recapturedByLevel[currentLevel] = 0
					}
					recapturedByLevel[currentLevel]++
				}
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

	return captured, recapturedByLevel
}

func main() {
	n := 2
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
	marked, total := sampleMarked(*markedFlag, makePartida)
	slog.Info(
		"MARKED_DONE",
		"got", total,
		"wanted", *markedFlag,
		"levels_marked", len(marked),
		"finished", time.Since(start).Seconds(),
	)

	// let's capture some elements
	captured, recapturedByLevel := capture(*capturedFlag, makePartida, marked)

	// let's assure that marked has all the keys from recapturedByLevel
	for k := range recapturedByLevel {
		if _, ok := marked[k]; !ok {
			marked[k] = map[string]bool{}
		}
	}

	// let's calculate the recaptured elements using Lincoln-Petersen method
	lincolnEstimatesByLevel := map[int]float64{}
	for level := range recapturedByLevel {
		N := len(marked[level]) * len(captured[level]) / recapturedByLevel[level]
		lincolnEstimatesByLevel[level] = float64(N)
	}

	totalRecaptured := 0.0
	for _, v := range lincolnEstimatesByLevel {
		totalRecaptured += v
	}

	slog.Info(
		"MULTI_LINCOLN_DONE",
		"wanted", *capturedFlag,
		"recapturedByLevel", recapturedByLevel,
		"lincolnEstimatesByLevel", lincolnEstimatesByLevel,
		"N", totalRecaptured,
		"finished", time.Since(start).Seconds(),
	)
}
