package main

import (
	"log"

	"github.com/filevich/combinatronics"
	"github.com/hll-truco/experiments/utils"
	"github.com/truquito/truco/pdt"
)

func setCartas(m *pdt.Manojo, cartasIDs []int) {
	for i, id := range cartasIDs {
		c := pdt.NuevaCarta(pdt.CartaID(id))
		m.Cartas[i] = &c
	}
}

func main() {
	// mazo de cartas primas
	// miniIDs := []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}
	miniIDs := []int{20, 0, 26, 36, 12, 16, 5}

	m := &pdt.Manojo{}

	// todas las muestras posibles
	for _, muestraID := range miniIDs {
		muestra := pdt.NuevaCarta(pdt.CartaID(muestraID))
		resto := utils.CopyWithoutThese(miniIDs, muestraID)
		// todos mis manojos posibles
		for _, miManojoIDs := range combinatronics.Combs(resto, 3) {

			setCartas(m, miManojoIDs)
			if tieneFlor, _ := m.TieneFlor(muestra); tieneFlor {
				log.Printf("bajo la muestra `%s` el manojo %v con ids %v tiene flor\n",
					muestra,
					m.Cartas,
					miManojoIDs)
				return
			}

		}
	}
}
