package main_test

import (
	"fmt"
	"testing"

	"github.com/hll-truco/hll-truco"
	"golang.org/x/crypto/sha3"
)

func TestByteSlice(t *testing.T) {
	bs := []byte{0xFF, 0x45, 0x3A}
	fmt.Println(bs)
}

func TestBytes(t *testing.T) {
	{
		x := uint64(0)
		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 0 && rho == 61; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := hll.Cast8BytesAsUInt64([]byte{0b11111110, 0, 0, 0, 0, 0, 0, 0})
		// como lo intrepreto como little endian se "da vuelta"
		// 00000000 00000000 00000000 00000000 00000000 00000000 00000000 11111110
		// ---- estos los iterepta como i
		//     ---- -------- -------- -------- -------- -------- -------- -------- estos como los ceros? 4 ceros + 7*(8 ceros) = 61
		//
		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 0 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := hll.Cast8BytesAsUInt64([]byte{0b11111110, 0, 0, 0, 0, 0, 0, 0x30})
		// como lo intrepreto como little endian se "da vuelta"
		// 00110000 00000000 00000000 00000000 00000000 00000000 00000000 11111110
		// ---- estos los iterepta como i = 3
		//     ---- -------- -------- -------- -------- -------- -------- -------- estos como los ceros? 4 ceros + 6*(8 ceros) = 52

		// retorna cantidad de 0's + 1s

		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 3 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		x := hll.Cast8BytesAsUInt64([]byte{0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b0000000, 0b11110000})
		// como lo intrepreto como little endian se "da vuelta"
		// 11110000 00000000 00000000 00000000 00000000 00000000 00000000 00000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- -------- -------- -------- -------- -------- 4 ceros + 7*(8 ceros) = 60

		// retorna cantidad de 0's + 1s

		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 60+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		//                               |
		x := hll.Cast8BytesAsUInt64([]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b0000000, 0b11110000})
		//                               ^ este lo cambio a 1
		// como lo intrepreto como little endian se "da vuelta"
		// 11110000 00000000 00000000 00000000 00000000 00000000 00000000 10000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- -------- -------- -------- -------- -------- 4 ceros + 6*(8 ceros) = 52

		// retorna cantidad de 0's + 1s

		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 52+1; !ok {
			t.Error("got something diff than expected")
		}
	}

	{
		//                               a                                                   b                               ****
		//                               |                                               ====|       ========    ========    ||||====
		x := hll.Cast8BytesAsUInt64([]byte{0b10000000, 0b00000000, 0b00000000, 0b00000000, 0b00001000, 0b00000000, 0b00000000, 0b11110000})
		//                               ^----------------estos lo cambio a 1----------------^
		// como lo intrepreto como little endian se "da vuelta"

		// ****                           b                               a
		// ||||==== ======== ======== ====|                               |
		// 11110000 00000000 00000000 00001000 00000000 00000000 00000000 10000000
		// ---- estos los iterepta como i = 15
		//     ---- -------- -------- ---- 4 + 2*8 + 4 = 24

		// retorna cantidad de 0's + 1s

		i, rho := hll.GetPosVal(x, 4)
		t.Logf("for x:%d -> i:%d rho:%d\n", x, i, rho)
		if ok := i == 15 && rho == 24+1; !ok {
			t.Error("got something diff than expected")
		}
	}
}

func TestFullRho(t *testing.T) {
	// test: para un uint64 que se le aplica `getPosVal` (i.e., el primer ui64)
	//       su "Val"/Rho es todo nulo sii Rho == (64-p)+1 (ya que siempre se le suma 1)

	hash := []byte{
		// primer uint64 ~ 8 bytes (su rho es 61, por lo que es todo nulo)
		//                                                                                    ****
		//                                                                                    ||||
		0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b11110000, // 4+7*8=60 (+1)
		// segundo uint64 ~ 8 bytes (su `bits.LeadingZeros64` es 64 por lo que es todo nulo, se suma a rho, y se pasa al siguiente)
		0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, // + 64
		0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000001, 0b00000000, // + (8 + 7) = 15 BREAK
		//                                                     le agrego un 1 acá -------^
		0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000,
	}

	// number of bits to interpret as the bucket index
	p := uint8(4)
	i, totalRho := hll.GetPosValDynamic(hash, p)
	t.Logf("i:%d, totalRho:%d\n", i, totalRho)

	if ok := i == 15 && totalRho == 60+1+64+15; !ok {
		t.Error("got something diff than expected")
	}
}

func TestSha3(t *testing.T) {

	data := []byte("your data here")

	// 1 = (1 byte / 8 bits) = (8b/1B) = 1
	// so...
	// if we want a 600 bits hash,
	// 600 bits =  600b / (8b/1B) = 600/8 B = 75B

	// if we want a 256 bits hash, 256/8 = 32B

	hash256bits := make([]byte, 32)
	hash640bits := make([]byte, 80)
	hash1024bits := make([]byte, 128)

	sha3.ShakeSum256(hash256bits, data)
	sha3.ShakeSum256(hash640bits, data)
	sha3.ShakeSum256(hash1024bits, data)

	t.Logf("hash256bits: %x\n", hash256bits)
	t.Logf("hash640bits: %x\n", hash640bits)
	t.Logf("hash1024bits: %x\n", hash1024bits)
}
