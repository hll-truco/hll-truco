package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/axiomhq/hyperloglog"
	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

// flags/parametros:
var (
	deckSizeFlag = flag.Int("deck", 14, "Deck size")
	absIDFlag    = flag.String("abs", "a1", "Abstractor ID")
	infosetFlag  = flag.String("info", "InfosetRondaBase", "Infoset impl. to use")
	hashIDFlag   = flag.String("hash", "sha1", "Infoset hashing function")
	reportFlag   = flag.Int("report", 60*10, "Delta (in seconds) for printing log msgs")
)

var (
	deck        []int               = nil
	infoBuilder *info.Builder       = nil
	verbose     bool                = true
	terminals   uint64              = 0
	printer     *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
	axiom                           = hyperloglog.New16()
)

var start = time.Now()

var limit = time.Minute

// var limit = 10 * time.Minute

func init() {
	flag.Parse()

	log.Println("deckSize", *deckSizeFlag)
	log.Println("absId", *absIDFlag)
	log.Println("infoset", *infosetFlag)
	log.Println("hash", *hashIDFlag)
	log.Println("report every", *reportFlag)

	deck = utils.Deck(*deckSizeFlag)
	infoBuilder = info.BuilderFactory(*hashIDFlag, *infosetFlag, *absIDFlag)
	printer = utils.NewCronoPrinter(time.Second * time.Duration(*reportFlag))
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
		hashFn := utils.ParseHashFn(*hashIDFlag)
		hash := info.HashBytes(hashFn)
		axiom.Insert(hash)
		// if h.Add(hash) {
		// 	log.Println(time.Since(start), h.M)
		// }

		chis := pdt.Chis(p)
		j := uniformPick(chis)

		pkts, _ := j.Hacer(p)

		if pdt.IsDone(pkts) || p.Terminada() {
			terminals++
		}

		if printer.ShouldPrint() {
			e := axiom.Estimate()
			// mem := utils.GetMemUsage()
			printer.Print(fmt.Sprintf("\n\testimate:%d",
				e))
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

	log.Println("final estimate:", axiom.Estimate())
	log.Println("terminals:", terminals)
	log.Println("finished:", time.Since(start))
}
