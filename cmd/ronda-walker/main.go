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
	verbose   bool                = false
	terminals uint64              = 0
	printer   *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 60)
)

// full
// var deck = []int{
// 	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
// 	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
// }

// 14
var deck = []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}

// 12
// var deck = []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27}

// var deck = []int{10, 20, 11, 21, 14, 24, 15}

// var deck = []int{20, 0, 26, 36, 12, 16, 5}

// mapa de nivel:branches
// e.g., 0:140
//       1:2
//       2:3

var topology = make(map[uint]int)
var dones = make(map[uint]int)

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
			_, empiezaNuevaRonda := chis[mix][aix].Hacer(p)

			// empiezaNuevaRondaCalc := utils.RondaIsDone(pkts)
			// if empiezaNuevaRondaCalc != empiezaNuevaRonda {
			// 	p, _ = pdt.Parse(string(bs), verbose)
			// 	_, empiezaNuevaRonda := chis[mix][aix].Hacer(p)
			// 	panic(fmt.Sprintf("no da lo mismo %v %v\n", empiezaNuevaRondaCalc, empiezaNuevaRonda))
			// }

			// 1. se termino la partida? -> terminals +1
			// 2. se termino la ronda? -> simular todas las aristas del chance node
			// 3. ninguna de las dos -> rec

			if empiezaNuevaRonda || p.Terminada() {
				terminals++
				printer.Print(fmt.Sprintf("\n\ttopo: %v\n\tdone: %v", topology, dones))
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

	p, _ := pdt.NuevaPartida(
		pdt.A40, // <----- no importa poque la condicion de parada es Ronda
		true,    // mini
		azules[:n>>1],
		rojos[:n>>1],
		limEnvite, // limiteEnvido
		verbose)

	log.Println("total aristas nivel 0:", todasLasAristasChancePosibles(p, 0))
	log.Println("terminals:", terminals)
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
