// Package model models
package model

import (
	"crypto/rand"
	"math/big"
)

// 24 bits “mantissa”, otherwise known as a coefficient or significand.
const maxInt int64 = 1 << 24

// 24 bits “mantissa”, otherwise known as a coefficient or significand.
const startPrice float32 = 50

// Price info about one position
type Price struct {
	Name string
	SellingPrice,
	PurchasePrice float32
}

// GetStartPrices get start prices slice
func GetStartPrices() []*Price {
	return []*Price{
		{
			Name:          "gold",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float32(),
		},
		{
			Name:          "oil",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float32(),
		},
		{
			Name:          "tesla",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float32(),
		},
		{
			Name:          "google",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float32(),
		},
	}
}

// Float32 random float32 using crypto/rand
func Float32() float32 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	return float32(nBig.Int64()) / float32(maxInt)
}
