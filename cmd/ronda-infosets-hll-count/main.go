package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/filevich/truco-cfr/abs"
	"github.com/filevich/truco-cfr/info"
	hll "github.com/hll-truco/experiments"
	"github.com/hll-truco/experiments/utils"
	"github.com/truquito/truco/pdt"
)

// flags/parametros:
var (
	deckSize = flag.Int("deck", 14, "Deck size")
	absID    = flag.String("abs", "a1", "Abstractor ID")
	infoset  = flag.String("info", "InfosetRondaBase", "Infoset impl. to use")
	hashID   = flag.String("hash", "sha1", "Infoset hashing function")
)

var (
	deck        []int               = nil
	a           abs.IAbstraccion    = nil
	infoBuilder info.InfosetBuilder = nil
	verbose     bool                = true
	terminals   uint64              = 0
	printer     *utils.CronoPrinter = nil
	h           *hll.HyperLogLog    = hll.NewHyperLogLog(2)
)

var start = time.Now()

// var limit = time.Millisecond * 500
var limit = time.Second

func init() {
	flag.Parse()

	log.Println("deckSize", *deckSize)
	log.Println("absId", *absID)
	log.Println("infoset", *infoset)
	log.Println("hash", *hashID)

	deck = utils.Deck(*deckSize)
	a = abs.ParseAbstractor(*absID)
	infoBuilder = info.ParseInfosetBuilder(*infoset)
	printer = utils.NewCronoPrinter(0)
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
		info := infoBuilder(p, activePlayer, a, nil)
		hashFn := utils.ParseHashFn(*hashID)
		hash := info.HashBytes(hashFn)
		if h.Add(hash) {
			log.Println(time.Since(start), h.M)
		}

		chis := pdt.Chis(p)
		j := uniformPick(chis)

		pkts, _ := j.Hacer(p)

		if pdt.IsDone(pkts) || p.Terminada() {
			terminals++
			if false {
				mem := utils.GetMemUsage()
				printer.Print(fmt.Sprintf("\n\t%s",
					mem))
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

	log.Println("M_j:", h.M)
	log.Println("terminals:", terminals)
	log.Println("finished:", time.Since(start))
}
