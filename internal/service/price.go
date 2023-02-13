// Package service serv
package service

import (
	"Price-Provider/internal/model"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

// MQ stream interface for user service
//
//go:generate mockery --name=MQ --case=underscore --output=./mocks
type MQ interface {
	PublishPrices(ctx context.Context, prices model.Prices) error
}

// 24 bits “mantissa”, otherwise known as a coefficient or significand.
const maxInt int64 = 1 << 24

// Prices price service
type Prices struct {
	messageQueue MQ
	prices       model.Prices
	maxPrice     float32
}

// NewPrices constructor
func NewPrices(mq MQ, mp float32) *Prices {
	return &Prices{messageQueue: mq, prices: make(model.Prices), maxPrice: mp}
}

// PublishPrices publishing prices into MQ
func (p *Prices) PublishPrices(ctx context.Context) error {
	err := p.messageQueue.PublishPrices(ctx, p.prices)
	if err != nil {
		return fmt.Errorf("prices - PublishPrices - PublishPrices: %w", err)
	}

	return nil
}

// RandPrices random price generation
func (p *Prices) RandPrices() {
	for _, key := range p.prices.GetPricesKeys() {
		p.prices[key] = p.maxPrice * Float32()
	}
}

// Float32 random float32 using crypto/rand
func Float32() float32 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(maxInt))
	return float32(nBig.Int64()) / float32(maxInt)
}
