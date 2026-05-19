package helper

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// Simple algebraic key sharding mechanism for multi-enclave demonstration
func SplitKey(secretHex string,curveOrder *big.Int) ([]byte, []byte, []byte) {
	secretBytes, _ := hex.DecodeString(secretHex)
	secretInt := new(big.Int).SetBytes(secretBytes)
	a1, _ := rand.Int(rand.Reader, curveOrder)

	calcShare := func(x int64) []byte {
		xInt := big.NewInt(x)
		tmp := new(big.Int).Mul(a1, xInt)
		tmp.Add(tmp, secretInt)
		tmp.Mod(tmp, curveOrder)
		return tmp.Bytes()
	}
	return calcShare(1), calcShare(2), calcShare(3)
}