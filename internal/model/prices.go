// Package model models
package model

import (
	"crypto/rand"
	"math/big"
)

// 24 bits “mantissa”, otherwise known as a coefficient or significand.
const maxInt int64 = 1 << 53

// 24 bits “mantissa”, otherwise known as a coefficient or significand.
const startPrice float64 = 50

// Price info about one position
type Price struct {
	Name string
	SellingPrice,
	PurchasePrice float64
}

// GetStartPrices get start prices slice
func GetStartPrices() []*Price {
	return []*Price{
		{
			Name:          "gold",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float64(),
		},
		{
			Name:          "oil",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float64(),
		},
		{
			Name:          "tesla",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float64(),
		},
		{
			Name:          "google",
			SellingPrice:  startPrice,
			PurchasePrice: startPrice + Float64(),
		},
	}
}

// Float64 random float64 using crypto/rand
func Float64() float64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	return float64(nBig.Int64()) / float64(maxInt)
}
