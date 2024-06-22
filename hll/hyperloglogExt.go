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
	"log/slog"
	"math"
	"math/big"

	"github.com/ALTree/bigfloat"
)

type HyperLogLogExt struct {
	reg []uint // regs
	m   uint32 // 2^precision
	p   uint8  // precision
	max uint   // max number of leading zeros so far
}

// New returns a new initialized HyperLogLogExt.
func NewExt(precision uint8) (*HyperLogLogExt, error) {
	fmt.Println("init hll fork")
	if precision > 16 || precision < 4 {
		return nil, errors.New("precision must be between 4 and 16")
	}

	h := &HyperLogLogExt{}
	h.max = uint(0)
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

	if zeroBits > h.max {
		slog.Debug("NEW_RECORD", "m", zeroBits) // comment this for no output
		h.max = zeroBits
	}

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

	// deprecated
	Two32Big   = big.NewFloat(0).SetUint64(1 << 32)
	Two32Div30 = big.NewFloat(0).Quo(Two32Big, big.NewFloat(30))

	// new
	// 2^1024
	TwoBig         = big.NewInt(2)
	ExpBig         = big.NewInt(1024)
	Two1024Big     = new(big.Float).SetInt(new(big.Int).Exp(TwoBig, ExpBig, nil))
	Two1024Div1022 = big.NewFloat(0).Quo(Two1024Big, big.NewFloat(1022))

	negative = big.NewFloat(-1)
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
	} else if isLesser := est.Cmp(Two1024Div1022) == -1; isLesser {
		return est
	}

	return new(big.Float).Mul(
		new(big.Float).Mul(negative, Two1024Big),
		bigfloat.Log(
			new(big.Float).Sub(
				big.NewFloat(1),
				new(big.Float).Quo(est, Two1024Big),
			),
		),
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

// dynm hll

// stategies
func MaxDynm(h *HyperLogLogExt) float64 {
	_m := uint(0)
	for _, v := range h.reg {
		if v > _m {
			_m = v
		}
	}
	return float64(_m)
}

// func Mean(h *HyperLogLogExt) float64 {
// 	s := uint(0)
// 	for _, r := range h.reg {
// 		s += r
// 	}
// 	avg := float64(s) / float64(len(h.reg))
// 	return avg
// }

func Fixed(h *HyperLogLogExt, f float64) float64 {
	return f
}

func MaxPlusDelta(h *HyperLogLogExt, d float64) float64 {
	return MaxDynm(h) + d
}

func MinSafePlusDelta(h *HyperLogLogExt, d float64) float64 {
	est := calculateEstimateExt(h.reg)
	exp := MaxDynm(h)

	for base := math.Exp2(exp); est > base; exp++ {
		base = math.Exp2(exp)
	}
	return exp + d
}

func MinSafePlusDeltaBig(h *HyperLogLogExt, d float64) float64 {
	est := calculateEstimateExtBig(h.reg)
	exp := uint(math.Ceil(MaxDynm(h)))
	base := calc2Pow(exp)

	for {
		if isEnough := est.Cmp(base) <= 0; isEnough {
			break
		}
		exp++
		base = calc2Pow(exp)
	}

	return float64(exp) + d
}

func Dynm(h *HyperLogLogExt) float64 {

	// 1. vanilla hll
	// return Fixed(h, 32.0)

	// 2. max + sqrt(p)
	// delta := math.Sqrt(float64(h.p))
	// return MaxPlusDelta(h, delta)

	// 3. max + 10
	// return MaxPlusDelta(h, 10.0)

	// 4. min safe
	return MinSafePlusDelta(h, 5.0) // non-big
	// return MinSafePlusDeltaBig(h, 0.0) // big

}

func (h *HyperLogLogExt) CountDynm() uint64 {
	est := calculateEstimateExt(h.reg)
	exp := Dynm(h)
	base := math.Exp2(exp)

	slog.Debug(
		"VALUES_NORMAL",
		"m", h.max,
		"max", MaxDynm(h),
		"exp", exp,
		"est", est)

	e := uint64(-base * math.Log(1-est/base))
	return e
}

var _oneBig = big.NewInt(1)

func calc2Pow(e uint) *big.Float {
	result := new(big.Int).Lsh(_oneBig, e)
	return new(big.Float).SetInt(result)
}

func (h *HyperLogLogExt) CountBigDynm() *big.Float {
	est := calculateEstimateExtBig(h.reg)
	exp := Dynm(h)
	base := calc2Pow(uint(exp))

	slog.Info(
		"VALUES_BIG",
		"m", h.max,
		"exp", exp,
		"est", est,
		"base", base)

	return new(big.Float).Mul(
		new(big.Float).Mul(negative, base),
		bigfloat.Log(
			new(big.Float).Sub(
				big.NewFloat(1),
				new(big.Float).Quo(est, base),
			),
		),
	)
}
