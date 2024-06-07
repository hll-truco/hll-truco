package testbigint_test

import (
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	// Create a new big.Int with value 2
	base := big.NewInt(2)

	// Create a new big.Int with value 512
	exponent := big.NewInt(512)

	// Allocate a new big.Int to store the result
	result := new(big.Int)

	// Calculate 2^512 and store the result in result
	result.Exp(base, exponent, nil)

	t.Log(result)
}
