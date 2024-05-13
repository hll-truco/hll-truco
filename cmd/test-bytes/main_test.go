package main_test

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"testing"
)

func cast8BytesAsUInt64(bs []byte) uint64 {
	return binary.LittleEndian.Uint64(bs[:8])
}

func bextr(v uint64, start, length uint8) uint64 {
	return (v >> start) & ((1 << length) - 1)
}

func getPosVal(x uint64, p uint8) (uint64, uint8) {
	i := bextr(x, 64-p, p) // {x63,...,x64-p} // [63..60][]
	w := x<<p | 1<<(p-1)   // {x63-p,...,x0}
	rho := uint8(bits.LeadingZeros64(w)) + 1
	return i, rho
}

func TestByteSlice(t *testing.T) {
	bs := []byte{0xFF, 0x45, 0x3A}
	fmt.Println(bs)
}

func TestBytes(t *testing.T) {
	{
		x := uint64(0)
		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 0 && rho == 61; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := cast8BytesAsUInt64([]byte{0b11111110, 0, 0, 0, 0, 0, 0, 0})
		// como lo intrepreto como little endian se "da vuelta"
		// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 11111110
		// ---- estos los iterepta como i
		//     ---- -------- -------- -------- -------- -------- -------- -------- estos como los ceros? 4 ceros + 7*(8 ceros) = 61
		//
		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 0 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := cast8BytesAsUInt64([]byte{0b11111110, 0, 0, 0, 0, 0, 0, 0x30})
		// como lo intrepreto como little endian se "da vuelta"
		// 00110000 00000000 00000000 00000000 00000000 00000000 00000000 11111110
		// ---- estos los iterepta como i = 3
		//     ---- -------- -------- -------- -------- -------- -------- -------- estos como los ceros? 4 ceros + 6*(8 ceros) = 52

		// retorna cantidad de 0's + 1s

		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 3 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := cast8BytesAsUInt64([]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b0000000, 0b11110000})
		// como lo intrepreto como little endian se "da vuelta"
		// 11110000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- -------- -------- -------- -------- -------- 4 ceros + 7*(8 ceros) = 60

		// retorna cantidad de 0's + 1s

		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 60+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		//                               |
		x := cast8BytesAsUInt64([]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b0000000, 0b11110000})
		//                               ^ este lo cambio a 1
		// como lo intrepreto como little endian se "da vuelta"
		// 11110000 00000000 00000000 00000000 00000000 00000000 00000000 10000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- -------- -------- -------- -------- -------- 4 ceros + 6*(8 ceros) = 52

		// retorna cantidad de 0's + 1s

		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		//                               a                                                   b                               ****
		//                               |                                               ====|       ========    ========    ||||====
		x := cast8BytesAsUInt64([]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00001000, 0b00000000, 0b00000000, 0b11110000})
		//                               ^----------------estos lo cambio a 1----------------^
		// como lo intrepreto como little endian se "da vuelta"

		// ****                           b                               a
		// ||||==== ======== ======== ====|                               |
		// 11110000 00000000 00000000 00001000 00000000 00000000 00000000 10000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- ---- 4 + 2*8 + 4 = 24

		// retorna cantidad de 0's + 1s

		i, rho := getPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 24+1; !ok {
			t.Error("got something diff than expected")
		}
	}
}
