package hll

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

// Sha3Hash represents the SHA-3 hasher.
type Sha3Hash struct {
	sha3 sha3.ShakeHash
	size int
}

// NewSha3Hash creates a new Sha3Hash with the specified output size.
func NewSha3Hash(size int) hash.Hash {
	h := &Sha3Hash{
		sha3: sha3.NewShake256(),
		size: size,
	}
	return h
}

// Write adds more data to the running hash.
func (h *Sha3Hash) Write(p []byte) (int, error) {
	return h.sha3.Write(p)
}

// Sum appends the current hash to b and returns the resulting slice.
func (h *Sha3Hash) Sum(b []byte) []byte {
	hash := make([]byte, h.size)
	h.sha3.Read(hash)
	return append(b, hash...)
}

// Reset resets the Hash to its initial state.
func (h *Sha3Hash) Reset() {
	h.sha3.Reset()
	h.sha3 = sha3.NewShake256()
}

// Size returns the number of bytes Sum will return.
func (h *Sha3Hash) Size() int {
	return h.size
}

// BlockSize returns the hash's underlying block size.
func (h *Sha3Hash) BlockSize() int {
	return h.sha3.BlockSize()
}
