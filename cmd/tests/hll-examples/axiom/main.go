package main

import (
	"fmt"

	"github.com/axiomhq/hyperloglog"
)

func estimateError(got, exp uint64) float64 {
	var delta uint64
	if got > exp {
		delta = got - exp
	} else {
		delta = exp - got
	}
	return float64(delta) / float64(exp)
}

func main() {
	axiom := hyperloglog.New16()
	// for 1e8 ~   100.000.000 got 100_807_828
	// for 1e9 ~ 1.000.000.000 got ?
	exact := uint64(1e9)
	step := uint64(1e8)

	for i := uint64(1); i <= exact; i++ {
		str := fmt.Sprintf("stream-%d", i)
		axiom.Insert([]byte(str))

		if i%step == 0 || i == 1e9 {
			// step *= 5
			res := axiom.Estimate()
			ratio := 100 * estimateError(res, exact)
			progress := 100 * float64(i) / float64(exact)
			fmt.Printf("Exact: %d, iter: %d (%.1f), estimate: %d (%.4f%% off)\n",
				exact,
				i,
				progress,
				res,
				ratio)
		}
	}

	res := axiom.Estimate()
	ratio := 100 * estimateError(res, exact)
	fmt.Printf("AxiomHQ HLL total size:\t %d (%.4f%% off)\n", res, ratio)

	data2, err := axiom.MarshalBinary()
	if err != nil {
		panic(err)
	}

	fmt.Printf("AxiomHQ HLL total size:\t %d\n", len(data2))
}
