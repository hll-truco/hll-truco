package hll

import (
	"encoding/binary"
	"math/bits"
)

func Cast8BytesAsUInt64(bs []byte) uint64 {
	return binary.LittleEndian.Uint64(bs[:8])
}

func Bextr(v uint64, start, length uint8) uint64 {
	return (v >> start) & ((1 << length) - 1)
}

func GetPosVal(x uint64, p uint8) (uint64, uint8) {
	i := Bextr(x, 64-p, p) // {x63,...,x64-p} // [63..60][]
	w := x<<p | 1<<(p-1)   // {x63-p,...,x0}
	rho := uint8(bits.LeadingZeros64(w)) + 1
	return i, rho
}

func GetPosValDynamic(hash []byte, p uint8) (uint64, uint) {

	var (
		i        = uint64(0)
		totalRho = uint(0)
	)

	// uint64 are 8 bytes
	// so i break the hash in chunks of 8 bytes
	for offset := 0; offset < len(hash)/8; offset++ {

		ix := 8 * offset // 8 porque 1 uint64 = 8 bytes
		x := Cast8BytesAsUInt64(hash[ix : ix+8])

		if isFirstInt := offset == 0; isFirstInt {

			pos, rho := GetPosVal(x, p)

			i = pos
			totalRho += uint(rho)

			if shouldBreak := rho < (64-p)+1; shouldBreak {
				break
			}

		} else {
			rho := uint(bits.LeadingZeros64(x))
			totalRho += rho
			if shouldBreak := rho < 64; shouldBreak {
				break
			}
		}

	}

	return i, totalRho
}
