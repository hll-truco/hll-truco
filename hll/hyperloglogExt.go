// package hll implements the HyperLogLog and HyperLogLog++ cardinality
// estimation algorithms.
// These algorithms are used for accurately estimating the cardinality of a
// multiset using constant memory. HyperLogLog++ has multiple improvements over
// HyperLogLog, with a much lower error rate for smaller cardinalities.
//
// HyperLogLog is described here:
// http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf
//
// HyperLogLog++ is described here:
// http://research.google.com/pubs/pub40671.html
package hll

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
)

type HyperLogLogExt struct {
	reg []uint // regs
	m   uint32 // 2^precision
	p   uint8  // precision
}

// New returns a new initialized HyperLogLogExt.
func NewExt(precision uint8) (*HyperLogLogExt, error) {
	fmt.Println("init hll fork")
	if precision > 16 || precision < 4 {
		return nil, errors.New("precision must be between 4 and 16")
	}

	h := &HyperLogLogExt{}
	h.p = precision
	h.m = 1 << precision
	h.reg = make([]uint, h.m)
	return h, nil
}

// Clear sets HyperLogLogExt h back to its initial state.
func (h *HyperLogLogExt) Clear() {
	h.reg = make([]uint, h.m)
}

// Add adds a new item to HyperLogLogExt h.
func (h *HyperLogLogExt) Add(hash []byte) bool {
	i, zeroBits := GetPosValDynamic(hash, h.p)

	if zeroBits > h.reg[i] {
		h.reg[i] = zeroBits
		return true
	}

	return false
}

// Merge takes another HyperLogLogExt and combines it with HyperLogLogExt h.
// returns true if there's bump in any of its regs after the merge
func (h *HyperLogLogExt) Merge(other *HyperLogLogExt) (bool, error) {
	if h.p != other.p {
		return false, errors.New("precisions must be equal")
	}

	bump := false
	for i, v := range other.reg {
		if v > h.reg[i] {
			h.reg[i] = v
			bump = true
		}
	}
	return bump, nil
}

// Count returns the cardinality estimate.
func (h *HyperLogLogExt) Count() uint64 {
	est := calculateEstimateExt(h.reg)
	if est <= float64(h.m)*2.5 {
		if v := countZerosExt(h.reg); v != 0 {
			return uint64(linearCounting(h.m, v))
		}
		return uint64(est)
	} else if est < two32/30 {
		return uint64(est)
	}
	return uint64(-two32 * math.Log(1-est/two32))
}

var (
	twoPointFive = big.NewFloat(2.5)
	two32Big     = big.NewFloat(0).SetUint64(1 << 32)
	two32Div30   = big.NewFloat(0).Quo(two32Big, big.NewFloat(30))
	negative     = big.NewFloat(-1)
)

// Count returns the cardinality estimate.
func (h *HyperLogLogExt) CountBig() *big.Float {
	est := calculateEstimateExtBig(h.reg)

	m := new(big.Float).SetInt64(int64(h.m))
	tmp := new(big.Float).Mul(m, twoPointFive)

	// [<,=,>] ~ [-1,0,1]
	// so, since we need `<=` the cases we are intereted in are [-1,0]
	// so, basically, we ask if `cmp is <= 0`
	if estIsLesserOrEqualThanTmp := est.Cmp(tmp) <= 0; estIsLesserOrEqualThanTmp {
		// `countZerosExt` just count the number of zeros in a slice
		if v := countZerosExt(h.reg); v != 0 {
			return linearCountingBig(h.m, v)
		}
		return est
	} else if isLesser := est.Cmp(two32Div30) == -1; isLesser {
		return est
	}

	return new(big.Float).Mul(
		new(big.Float).Mul(negative, two32Big),
		bigfloat.Log(new(big.Float).Quo(new(big.Float).Sub(big.NewFloat(1), est), two32Big)),
	)
}

// Encode HyperLogLogExt into a gob
func (h *HyperLogLogExt) GobEncode() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(h.reg); err != nil {
		return nil, err
	}
	if err := enc.Encode(h.m); err != nil {
		return nil, err
	}
	if err := enc.Encode(h.p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode gob into a HyperLogLogExt structure
func (h *HyperLogLogExt) GobDecode(b []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	if err := dec.Decode(&h.reg); err != nil {
		return err
	}
	if err := dec.Decode(&h.m); err != nil {
		return err
	}
	if err := dec.Decode(&h.p); err != nil {
		return err
	}
	return nil
}

func (h *HyperLogLogExt) String() string {
	return fmt.Sprintf("%v", h.reg)
}
