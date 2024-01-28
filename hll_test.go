package hll

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestRho(t *testing.T) {
	// from Flajolet paper:
	// b1000 ~ rho(b1000) = 4
	// and we know that
	// b1000 ~ d8 ~ 0x8
	i := []byte{0x8}
	rho := countConsecutiveZeroBits(i) + 1
	var exp uint = 4
	if ok := rho == exp; !ok {
		t.Error()
	}
}

func meanAndStdev(nums []uint) (float64, float64) {
	sum := 0.0
	for _, num := range nums {
		sum += float64(num)
	}
	mean := sum / float64(len(nums))

	variance := 0.0
	for _, num := range nums {
		variance += math.Pow(float64(num)-mean, 2)
	}
	stdev := math.Sqrt(variance / float64(len(nums)-1))

	return mean, stdev
}

func TestRhoAvg(t *testing.T) {
	N := 1_000_000
	xs := make([]uint, 0, N)
	for i := 0; i < N; i++ {
		h := sha256.New()
		h.Write([]byte(strconv.Itoa(i)))
		res := h.Sum(nil)
		res_s := hex.EncodeToString(res)
		rho := countConsecutiveZeroBits([]byte(res_s)) + 1
		xs = append(xs, rho)
	}

	m, s := meanAndStdev(xs)
	if ok := m < 1.9 && m < 2.1; !ok {
		t.Error()
	}

	t.Logf("m: %.5f %.5f", m, s)
}

func TestEstimateCardinality10M(t *testing.T) {
	for bits_buckets := 1; bits_buckets <= 14; bits_buckets++ {
		m := int(math.Pow(2, float64(bits_buckets)))
		start := time.Now()
		T := 10000000 - 10
		delta := int(math.Round(float64(T) / float64(m)))

		buckets := make([]uint, m)

		for i := 0; i < m; i++ {
			_from, _to := delta*i, delta*(i+1)
			hll := NewHyperLogLog()
			for j := _from; j <= _to; j++ {
				hll.Add([]byte(strconv.Itoa(j)))
			}
			buckets[i] = hll.M
		}

		c := Calc(buckets)
		tDelta := time.Since(start)
		n := len(buckets)
		if n > 10 {
			n = 10
		}
		fmt.Printf("Estimated cardinality for b=%d is %d (%v) %v\n",
			bits_buckets,
			c,
			tDelta,
			buckets[:n])
	}
}
