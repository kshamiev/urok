package testb

import (
	"crypto/rand"
	"math/big"
)

func GenInt(x int64) int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}
