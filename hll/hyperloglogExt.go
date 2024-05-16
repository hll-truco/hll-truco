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
func (h *HyperLogLogExt) Add(hash []byte) {
	i, zeroBits := GetPosValDynamic(hash, h.p)

	if zeroBits > h.reg[i] {
		h.reg[i] = zeroBits
	}
}

// Merge takes another HyperLogLogExt and combines it with HyperLogLogExt h.
func (h *HyperLogLogExt) Merge(other *HyperLogLogExt) error {
	if h.p != other.p {
		return errors.New("precisions must be equal")
	}

	for i, v := range other.reg {
		if v > h.reg[i] {
			h.reg[i] = v
		}
	}
	return nil
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
