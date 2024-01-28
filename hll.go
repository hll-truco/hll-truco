package hll

import (
	"crypto/sha256"
	"math"
)

// The number of buckets used in HyperLogLog is typically a power of 2, and the
// bucket size is usually between 5 and 8 bits 1.

type HyperLogLog struct {
	M uint
}

func NewHyperLogLog() *HyperLogLog {
	return &HyperLogLog{
		M: 0,
	}
}

func countConsecutiveZeroBits(slice []byte) uint {
	var count uint = 0
	for _, b := range slice {
		for i := 0; i < 8; i++ {
			if b&(1<<i) != 0 {
				return count
			}
			count++
		}
	}
	return count
}

func (h *HyperLogLog) Add(x []byte) {
	hash := sha256.New()
	hash.Write(x)
	res := hash.Sum(nil)
	rho := countConsecutiveZeroBits(res) + 1
	if rho > h.M {
		h.M = rho
	}
}

func alpha(m int) float64 {
	if m == 16 {
		return 0.673
	} else if m == 32 {
		return 0.697
	} else if m == 64 {
		return 0.709
	} else {
		return 0.7213 / (1 + 1.079/float64(m))
	}
}

func countZeroes(buckets []uint) int {
	count := 0
	for _, b := range buckets {
		if b == 0 {
			count++
		}
	}
	return count
}

func sumInverse(buckets []uint) float64 {
	s := 0.0
	for _, b := range buckets {
		s += math.Pow(2.0, -float64(b))
	}
	return s
}

// todo lo que sea 32 es del hash...

func Calc(buckets []uint) int {
	m := len(buckets)
	var z float64 = alpha(m) * float64(m*m) / sumInverse(buckets)
	if z <= 2.5*float64(m) {
		v := countZeroes(buckets)
		if v != 0 {
			return int(math.Round(float64(m) * math.Log(float64(m)/float64(v))))
		} else {
			return int(math.Round(z))
		}
	} else if z <= float64(1<<32)/30.0 {
		return int(math.Round(z))
	} else {
		return int(math.Round(-1 * float64(1<<32) * math.Log(1-z/float64(1<<32))))
	}
}
