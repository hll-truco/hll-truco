package sqrts_test

import (
	"math"
	"testing"
)

func TestSqrtRounds(t *testing.T) {
	minPrec := uint(4)
	maxPrec := uint(16)
	for p := minPrec; p <= maxPrec; p++ {
		base := uint(math.Sqrt(float64(p)))
		t.Logf("for p:%d got:%d\n", p, base)
	}
}

// sqrts_test.go:13: for p:4 got:2
// sqrts_test.go:13: for p:5 got:2
// sqrts_test.go:13: for p:6 got:2
// sqrts_test.go:13: for p:7 got:2
// sqrts_test.go:13: for p:8 got:2
// sqrts_test.go:13: for p:9 got:3
// sqrts_test.go:13: for p:10 got:3
// sqrts_test.go:13: for p:11 got:3
// sqrts_test.go:13: for p:12 got:3
// sqrts_test.go:13: for p:13 got:3
// sqrts_test.go:13: for p:14 got:3
// sqrts_test.go:13: for p:15 got:3
// sqrts_test.go:13: for p:16 got:4
