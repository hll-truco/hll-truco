package utils

import (
	"github.com/truquito/truco/enco"
	"github.com/truquito/truco/pdt"
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
