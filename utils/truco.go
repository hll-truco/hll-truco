package utils

import (
	"os"
	"strconv"

	"github.com/truquito/gotruco/enco"
	"github.com/truquito/gotruco/pdt"
)

// retorna true si empieza una nueva ronda
func RondaIsDone(pkts []enco.Envelope) bool {
	for _, pkt := range pkts {
		if pkt.Message.Cod() == enco.TNuevaRonda {
			return true
		}
	}
	return false
}

func ParseCartas(

	muestraID int,
	misCartasIDs,
	opCartasIDs []int,

) (

	muestra pdt.Carta,
	manojos []pdt.Manojo,

) {
	muestra = pdt.NuevaCarta(pdt.CartaID(muestraID))

	var (
		miManojo = &pdt.Manojo{}
		opManojo = &pdt.Manojo{}
	)

	// mi manojo
	for i, cartaID := range misCartasIDs {
		c := pdt.NuevaCarta(pdt.CartaID(cartaID))
		miManojo.Cartas[i] = &c
	}
	// manojo ok
	for i, cartaID := range opCartasIDs {
		c := pdt.NuevaCarta(pdt.CartaID(cartaID))
		opManojo.Cartas[i] = &c
	}

	return muestra, []pdt.Manojo{*miManojo, *opManojo}
}

func SetCartasRonda(

	p *pdt.Partida,
	muestraID int,
	misCartasIDs,
	opCartasIDs []int,

) {

	muestra, manojos := ParseCartas(muestraID, misCartasIDs, opCartasIDs)
	// seteo las cartas
	p.Ronda.SetMuestra(muestra)
	p.Ronda.SetManojos(manojos)

}

func NuevaPartida(
	pts pdt.Puntuacion,
	azules, rojos []string,
	limEnvite int,
	verbose bool,
) *pdt.Partida {
	// gotruco
	// p, err := pdt.NuevaPartida(pts, azules, rojos, limEnvite, verbose)

	// minitruco
	deckSize, _ := strconv.Atoi(os.Getenv("DECK"))
	if deckSize == 0 {
		deckSize = 14
	}
	p, err := pdt.NuevaPartida(
		pts,
		true,           // isMini
		Deck(deckSize), // decksize
		azules,
		rojos,
		limEnvite,
		verbose)

	if err != nil {
		panic(err)
	}

	return p
}
