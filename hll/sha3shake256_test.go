package hll_test

import (
	"testing"

	"github.com/hll-truco/hll-truco/hll"
	"golang.org/x/crypto/sha3"
)

func TestSha3(t *testing.T) {

	// 1 = (1 byte / 8 bits) = (8b/1B) = 1
	// so...
	// if we want a 600 bits hash,
	// 600 bits =  600b / (8b/1B) = 600/8 B = 75B

	// if we want a 256 bits hash, 256/8 = 32B

	hash256bits := make([]byte, 32)
	hash640bits := make([]byte, 80)
	hash1024bits := make([]byte, 128)
	fn := hll.NewSha3Hash(128)

	data := []byte("your data here")

	sha3.ShakeSum256(hash256bits, data)
	sha3.ShakeSum256(hash640bits, data)
	sha3.ShakeSum256(hash1024bits, data)

	fn.Write(data)
	hashResult := fn.Sum(nil)
	fn.Reset()

	// b0cc2f07c96d2f3cb071bed8796f80116de4d4fde513562401529b689334d54e
	t.Logf("hash256bits: %x\n", hash256bits)
	t.Logf("hash640bits: %x\n", hash640bits)
	t.Logf("hash1024bits: %x\n", hash1024bits)
	t.Logf("fn1024bits: %x\n", hashResult)

	data2 := []byte("your data here2")

	sha3.ShakeSum256(hash256bits, data2)
	sha3.ShakeSum256(hash640bits, data2)
	sha3.ShakeSum256(hash1024bits, data2)

	fn.Write(data2)
	hashResult = fn.Sum(nil)
	fn.Reset()

	// 47bf56dbe47761393a27013ba60029cca9a2ce5c94e7bee47bc0219db8a485df
	t.Logf("hash256bits: %x\n", hash256bits)
	t.Logf("hash640bits: %x\n", hash640bits)
	t.Logf("hash1024bits: %x\n", hash1024bits)
	t.Logf("fn1024bits: %x\n", hashResult)

	data3 := []byte("your data here")

	sha3.ShakeSum256(hash256bits, data3)
	sha3.ShakeSum256(hash640bits, data3)
	sha3.ShakeSum256(hash1024bits, data3)

	fn.Write(data3)
	hashResult = fn.Sum(nil)
	fn.Reset()

	// b0cc2f07c96d2f3cb071bed8796f80116de4d4fde513562401529b689334d54e
	t.Logf("hash256bits: %x\n", hash256bits)
	t.Logf("hash640bits: %x\n", hash640bits)
	t.Logf("hash1024bits: %x\n", hash1024bits)
	t.Logf("fn1024bits: %x\n", hashResult)
}

func TestSha3Imlp(t *testing.T) {
	hash256bits := make([]byte, 32)
	data := []byte("your data here2")
	sha3.ShakeSum256(hash256bits, data)
	// 47bf56dbe47761393a27013ba60029cca9a2ce5c94e7bee47bc0219db8a485df
	t.Logf("hash256bits: %x\n", hash256bits)

	fn := hll.NewSha3Hash(128)
	fn.Write([]byte("your data here2"))
	hashResult := fn.Sum(nil)
	t.Logf("hash1024bits: %x\n", hashResult)
	if ok := string(hashResult[:32]) == string(hash256bits); !ok {
		t.Error("hashes do not match")
	}
}
