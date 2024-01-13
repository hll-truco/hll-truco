package main

import (
	"fmt"

	"github.com/filevich/combinatronics"
	"github.com/hll-truco/experiments/utils"
)

func main() {
	// mazo de cartas primas
	miniIDs := []int{10, 20, 11, 21, 14, 24, 15, 25, 16, 26, 17, 27, 18, 28}

	var totalManojos uint64 = 0

	// todas las muestras posibles
	for _, muestraID := range miniIDs {
		resto := utils.CopyWithoutThese(miniIDs, muestraID)
		// todos mis manojos posibles
		for _, miManojoIDs := range combinatronics.Combs(resto, 3) {
			resto2 := utils.CopyWithoutThese(resto, miManojoIDs...)
			// todos sus manojos posibles
			totalManojos += uint64(len(combinatronics.Combs(resto2, 3)))
		}
	}

	fmt.Println("totalManojos", totalManojos)
}
