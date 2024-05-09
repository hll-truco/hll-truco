package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/filevich/combinatronics"
	"github.com/filevich/truco-cfr/abs"
	"github.com/filevich/truco-cfr/info"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/truco/pdt"
)

// flags/parametros:
var (
	deckSize = flag.Int("deck", 14, "Deck size")
	report   = flag.Int("report", 60*10, "Delta (in seconds) for printing log msgs")
	track    = flag.Bool("track", true, "Should I count infosets?")
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
	infosets    map[string]bool     = map[string]bool{}
	printer     *utils.CronoPrinter = nil
	topology                        = make(map[uint]int)
	dones                           = make(map[uint]int)
)

func init() {
	flag.Parse()

	log.Println("deckSize", *deckSize)
	log.Println("track", *track)
	log.Println("absId", *absID)
	log.Println("infoset", *infoset)
	log.Println("hash", *hashID)
	log.Println("report", *report)

	deck = utils.Deck(*deckSize)
	a = abs.ParseAbstractor(*absID)
	infoBuilder = info.ParseInfosetBuilder(*infoset)
	printer = utils.NewCronoPrinter(time.Second * time.Duration(*report))
}

func todasLasAristasChancePosibles(p *pdt.Partida, level uint) uint64 {

	var totalAristasPosibles uint64 = 0

	bs, _ := p.MarshalJSON()

	// todas las muestras posibles
	for _, muestraID := range deck {
		resto := utils.CopyWithoutThese(deck, muestraID)
		// todos mis manojos posibles
		todosMisManojosPosibles := combinatronics.Combs(resto, 3)
		for _, miManojoIDs := range todosMisManojosPosibles {
			resto2 := utils.CopyWithoutThese(resto, miManojoIDs...)
			// todos sus manojos posibles
			todosSusManojosPosibles := combinatronics.Combs(resto2, 3)

			for _, opManojoIDs := range todosSusManojosPosibles {

				topology[level] = len(deck) * len(todosMisManojosPosibles) * len(todosSusManojosPosibles)

				// ini
				p, _ = pdt.Parse(string(bs), verbose)
				utils.SetCartasRonda(p, muestraID, miManojoIDs, opManojoIDs)
				recPlay(p, level+1)
				// fin

				totalAristasPosibles += 1
				_, ok := dones[level]
				if !ok {
					dones[level] = 0
				}
				dones[level] += 1

			}
			// finalizado uno de los manojos de mi en `level`
		}
	}

	// termine con todas las de este level.
	// la borro
	delete(dones, level)
	delete(topology, level)

	return totalAristasPosibles
}

func recPlay(p *pdt.Partida, level uint) {
	bs, _ := p.MarshalJSON()

	// para la partida dada, todas las jugadas posibles
	chis := pdt.Chis(p)

	// largo?
	t := 0
	for mix := range chis {
		t += len(chis[mix])
	}
	topology[level] = t
	dones[level] = 0

	// las juego
	for mix := range chis {
		for aix := range chis[mix] {

			// ini
			p, _ = pdt.Parse(string(bs), verbose)

			// infoset?
			if *track {
				activePlayer := pdt.Rho(p)
				// info := info.NewInfosetRondaBase(p, activePlayer, a, nil)
				info := infoBuilder(p, activePlayer, a, nil)
				hashFn := utils.ParseHashFn(*hashID)
				infosets[info.Hash(hashFn)] = true
			}

			pkts, _ := chis[mix][aix].Hacer(p)

			if pdt.IsDone(pkts) || p.Terminada() {
				terminals++
				mem := utils.GetMemUsage()
				printer.Print(fmt.Sprintf("\n\ttopo: %v\n\tdone: %v\n\t%s\n\tcount: %d",
					topology,
					dones,
					mem,
					len(infosets)))
			} else {
				// sigue
				recPlay(p, level+1)
			}

			// termine con una arista
			dones[level] += 1
		}
	}

	// termine con todas
	delete(topology, level)
	delete(dones, level)
}

func main() {
	start := time.Now()
	n := 2
	limEnvite := 1
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}

	// p, err := pdt.NuevaMiniPartida(azules[:n>>1], rojos[:n>>1], verbose)

	rand.Seed(time.Now().UnixNano())

	p, _ := pdt.NuevaPartida(
		pdt.A40, // <----- no importa poque la condicion de parada es Ronda
		true,
		deck,
		azules[:n>>1],
		rojos[:n>>1],
		limEnvite, // limiteEnvido
		verbose)

	log.Println("total aristas nivel 0:", todasLasAristasChancePosibles(p, 0))
	log.Println("terminals:", terminals)
	log.Println("infosets:", len(infosets))
	log.Println("finished:", time.Since(start))
}

/*

bench (M2)
for i in {1..10}; do go run cmd/ronda-walker/*.go; done
NO-verbose:
x=617ms, s=2.05

verbose, usando `empiezaNuevaRonda`
x=647ms, s=5.26
(5% más lento que no-verbose)

verbose, usando `pdt.IsDone`
x=648, s=3.87
(5% más lento que no-verbose)

*/
