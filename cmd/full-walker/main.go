package main

import (
	"fmt"
	"log"
	"time"

	"github.com/filevich/combinatronics"
	"github.com/hll-truco/experiments/utils"
	"github.com/truquito/truco/pdt"
)

var (
	terminals uint64              = 0
	printer   *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
)

// mazo de cartas primas
// var deck = []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}
var deck = []int{10, 20, 11, 21, 14, 24, 15}

// mapa de nivel:branches
// e.g., 0:140
//       1:2
//       2:3
var topology = make(map[uint]int)
var topologyDone = make(map[uint]int)

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

			topology[level] = len(deck) * len(todosMisManojosPosibles) * len(todosSusManojosPosibles)

			for _, opManojoIDs := range todosSusManojosPosibles {
				p, _ = pdt.Parse(string(bs), true)
				utils.SetCartasRonda(p, muestraID, miManojoIDs, opManojoIDs)
				recPlay(p, level+1)
				totalAristasPosibles += 1

				// finalizado uno de los manojos de op en `level`
				if _, ok := topologyDone[level]; !ok {
					topologyDone[level] = 0
				}
				topologyDone[level] += 1

			}
			// finalizado uno de los manojos de mi en `level`
		}
		// finalizado una muestra en `level`
	}

	return totalAristasPosibles
}

func recPlay(p *pdt.Partida, level uint) {
	bs, _ := p.MarshalJSON()

	// para la partida dada, todas las jugadas posibles
	chis := pdt.Chis(p)

	// las juego
	for mix := range chis {
		for aix := range chis[mix] {
			p, _ = pdt.Parse(string(bs), true)
			pkts := chis[mix][aix].Hacer(p)

			// 1. se termino la partida? -> terminals +1
			// 2. se termino la ronda? -> simular todas las aristas del chance node
			// 3. ninguna de las dos -> rec

			if terminoLaPartida := p.Terminada(); terminoLaPartida {
				terminals++
				// utils.PrintEvery(terminals, 100_000)
				printer.Print(fmt.Sprintf("topo: %v, done: %v", topology, topologyDone))
				//
				//
			} else if terminoLaRonda := utils.RondaIsDone(pkts); terminoLaRonda {
				// simular todos repartos de manojos posibles
				todasLasAristasChancePosibles(p, level+1)
			} else {
				// sigue
				recPlay(p, level+1)
			}
		}
	}
}

func main() {
	n := 2
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}
	verbose := true

	// p, err := pdt.NuevaMiniPartida(azules[:n>>1], rojos[:n>>1], verbose)

	p, _ := pdt.NuevaPartida(
		pdt.A5, // 10 pts
		true,   // mini
		azules[:n>>1],
		rojos[:n>>1],
		0, // limiteEnvido
		verbose,
	)

	p.Puntajes[pdt.Azul] = 3
	p.Puntajes[pdt.Rojo] = 3

	log.Println("total aristas nivel 0:", todasLasAristasChancePosibles(p, 0))
	log.Println("terminals:", terminals)
}
