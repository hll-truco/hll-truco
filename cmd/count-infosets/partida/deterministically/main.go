package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/filevich/combinatronics"
	"github.com/filevich/truco-ai/info"
	"github.com/hll-truco/hll-truco/utils"
	"github.com/truquito/gotruco/pdt"
)

var (
	verbose   bool                = false
	terminals uint64              = 0
	infosets  map[string]bool     = map[string]bool{}
	printer   *utils.CronoPrinter = utils.NewCronoPrinter(time.Second * 10)
)

// mazo de cartas primas
// var deck = []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}
// var deck = []int{10, 20, 11, 21, 14, 24, 15}

var deck = []int{20, 0, 26, 36, 12, 16, 5}

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

			// infoset?
			activePlayer := pdt.Rho(p)
			// info := info.NewInfosetRondaBase(p, activePlayer, a, nil)
			infoBuilder := info.BuilderFactory("sha160", "InfosetRondaBase", "a3")
			i := infoBuilder.Info(p, activePlayer, nil)
			infosets[i.Hash(infoBuilder.Hash)] = true

			pkts := chis[mix][aix].Hacer(p)

			if terminoLaPartida := p.Terminada(); terminoLaPartida {
				terminals++
				if printer.ShouldPrint() {
					slog.Info(
						"REPORT",
						"topo", topology,
						"dones", dones,
						"count", len(infosets),
						"mem", utils.GetMemUsage())
					printer.Check()
				}
			} else if pdt.IsDone(pkts, true) {
				// simular todos repartos de manojos posibles
				todasLasAristasChancePosibles(p, level+1)
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

func init() {
	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info(
		"START")
}

func main() {
	start := time.Now()
	n := 2
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}

	os.Setenv("DECK", fmt.Sprintf("%d", deck))
	p := utils.NuevaPartida(pdt.A40, azules[:n>>1], rojos[:n>>1], 0, verbose)

	p.Puntajes[pdt.Azul] = 9
	p.Puntajes[pdt.Rojo] = 9

	slog.Info(
		"RESULTS",
		"totalAristasNivel0", todasLasAristasChancePosibles(p, 0),
		"terminals:", terminals,
		"infosets:", len(infosets),
		"finished", time.Since(start))
}
