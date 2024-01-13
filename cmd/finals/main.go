package main

import (
	"fmt"

	"github.com/truquito/truco/pdt"
)

var terminals uint = 0

func print(x uint) {
	if x%100_000 == 0 {
		fmt.Println(x)
	}
}

func rec_play(p *pdt.Partida) {
	bs, _ := p.MarshalJSON()

	// para la partida dada, todas las jugadas posibles
	chis := pdt.Chis(p)

	// las juego
	for mix := range chis {
		for aix := range chis[mix] {
			p, _ = pdt.Parse(string(bs), true)
			pkts2 := chis[mix][aix].Hacer(p)
			if pdt.IsDone(pkts2) {
				terminals++
				print(terminals)
			} else {
				rec_play(p)
			}
		}
	}
}

func main() {
	n := 2
	azules := []string{"Alice", "Ariana", "Annie"}
	rojos := []string{"Bob", "Ben", "Bill"}
	verbose := false

	p, err := pdt.NuevaMiniPartida(azules[:n>>1], rojos[:n>>1], verbose)

	// isMiniTruco := true
	// p, err := pdt.NuevaPartida(
	// 	pdt.A10,     // 10 pts
	// 	isMiniTruco, // mini
	// 	azules[:n>>1],
	// 	rojos[:n>>1],
	// 	1, // limiteEnvido
	// 	verbose,
	// )

	if err != nil {
		panic(err)
	}

	rec_play(p)

	fmt.Println("total:", terminals)
}
