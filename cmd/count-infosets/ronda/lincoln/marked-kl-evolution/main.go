package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"time"

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
	reportFlag   = flag.Int("report", 10, "Delta (in seconds) for printing log msgs")
	calcKLEvery  = flag.Int("calc_kl_every", 2487, "Calculate KL divergence every `calcKLEvery` elements")

	// actual size for -deck=14 is 248_732
	// we will mark 1% of the total population: 248_732 * 0.01 = 2_487
	// we will capture 10% of the total population: 248_732 * 0.1 = 24_873
	markedFlag    = flag.Int("marked", 2_487, "Number of marked elements")
	allowMazoFlag = flag.Bool("mazo", true, "Allow mazo?")
)

var (
	deck        []int               = nil //lint:ignore U1000 <your reason here>
	infoBuilder *info.Builder       = nil
	verbose     bool                = true
	printer     *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
	start       time.Time           = time.Now()
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
		"reportFlag", *reportFlag,
		"markedFlag", *markedFlag,
		"calcKLEvery", *calcKLEvery,
		"allowMazoFlag", *allowMazoFlag,
	)

	deck = utils.Deck(*deckSizeFlag)
	infoBuilder = info.BuilderFactory(*hashIDFlag, *infosetFlag, *absIDFlag)
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

func copyMap(originalMap map[int]int, destMap map[int]int) {
	for key, value := range originalMap {
		destMap[key] = value
	}
}

func equalizeMaps(p map[int]int, q map[int]int) {
	for key := range p {
		if _, ok := q[key]; !ok {
			q[key] = 0
		}
	}
	for key := range q {
		if _, ok := p[key]; !ok {
			p[key] = 0
		}
	}
}

func KL(p map[int]int, q map[int]int) float64 {
	// equalize maps
	equalizeMaps(p, q)

	// Calculate total counts for normalization
	pTotal := 0
	qTotal := 0
	for _, count := range p {
		pTotal += count
	}
	for _, count := range q {
		qTotal += count
	}

	// Initialize KL divergence
	kl := 0.0

	// Calculate KL divergence for each key in p
	for key, pCount := range p {
		// Skip if p(x) = 0 since 0 * log(0/q) = 0
		if pCount == 0 {
			continue
		}

		// Convert counts to probabilities
		pProb := float64(pCount) / float64(pTotal)

		// Get corresponding q count (0 if key doesn't exist in q)
		qCount := q[key]

		// Handle the case where q(x) = 0
		if qCount == 0 {
			// KL divergence is undefined when q(x) = 0 and p(x) > 0
			// Return positive infinity
			return math.Inf(1)
		}

		qProb := float64(qCount) / float64(qTotal)

		// Add to KL divergence: p(x) * log(p(x)/q(x))
		kl += pProb * math.Log(pProb/qProb)
	}

	return kl
}

func checkKL(currentLevelDist map[int]int, prevLevelDist map[int]int) {
	// base case
	if len(prevLevelDist) == 0 {
		copyMap(currentLevelDist, prevLevelDist)
		return
	}
	// now, we can calculate the KL divergence between currentLevelDist and prevLevelDist
	kl := KL(prevLevelDist, currentLevelDist)
	copyMap(currentLevelDist, prevLevelDist)
	slog.Info("KL", "kl", kl)
}

// returns a map of marked elements and a map of level distribution
func sampleMarked(markedSize int, makePartida PartidaFactory, calcKLEvery int) (map[string]bool, map[int]int) {
	marked := map[string]bool{}
	currentLevelDist := map[int]int{}
	prevLevelDist := map[int]int{}
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
				if _, ok := currentLevelDist[currentLevel]; !ok {
					currentLevelDist[currentLevel] = 0
				}
				currentLevelDist[currentLevel]++

				// calculate KL divergence every `calcKLEvery` elements
				if len(marked)%calcKLEvery == 0 {
					checkKL(currentLevelDist, prevLevelDist)
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

	return marked, currentLevelDist
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
	slog.Info("MARKED_START", "calcKLEvery", *calcKLEvery)
	marked, currentLevelDist := sampleMarked(*markedFlag, makePartida, *calcKLEvery)
	slog.Info(
		"MARKED_DONE",
		"got", len(marked),
		"wanted", *markedFlag,
		"marked", len(marked),
		"currentLevelDist", currentLevelDist,
		"finished", time.Since(start).Seconds(),
	)
}
