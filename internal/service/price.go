// Package service serv
package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/OVantsevich/PriceProvider/internal/model"
)

// MQ stream interface for user service
//
//go:generate mockery --name=MQ --case=underscore --output=./mocks
type MQ interface {
	PublishPrices(ctx context.Context, prices []*model.Price) error
}

// Prices price service
type Prices struct {
	messageQueue MQ
	mu           sync.RWMutex
	prices       []*model.Price
	pricesMAP    map[string]*model.Price
	maxChange    float64
}

// NewPrices constructor
func NewPrices(mq MQ, mch float64) *Prices {
	pr := &Prices{messageQueue: mq, prices: model.GetStartPrices(), maxChange: mch}
	pr.pricesMAP = make(map[string]*model.Price, len(pr.prices))
	for _, p := range pr.prices {
		pr.pricesMAP[p.Name] = p
	}
	return pr
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
//
//nolint:gomnd
func (p *Prices) RandPrices() {
	p.mu.Lock()
	for _, pr := range p.prices {
		chg := -p.maxChange + (2*p.maxChange)*model.Float64()
		pr.SellingPrice += chg
		pr.PurchasePrice += chg
	}
	p.mu.Unlock()
}

// GetCurrentPrices get current prices
func (p *Prices) GetCurrentPrices(ctx context.Context, names []string) (map[string]*model.Price, error) {
	p.mu.RLock()

	result := make(map[string]*model.Price, len(names))

	for _, n := range names {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceld")
		default:
			price, ok := p.pricesMAP[n]
			if !ok {
				return nil, fmt.Errorf("no such price available")
			}
			result[n] = &(*price)
		}
	}
	p.mu.RUnlock()

	return result, nil
}
