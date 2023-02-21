// Package service serv
package service

import (
	"Price-Provider/internal/model"
	"context"
	"fmt"
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
	prices       []*model.Price
	maxChange    float32
}

// NewPrices constructor
func NewPrices(mq MQ, mch float32) *Prices {
	return &Prices{messageQueue: mq, prices: model.GetStartPrices(), maxChange: mch}
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
	for _, pr := range p.prices {
		chg := -p.maxChange + (2*p.maxChange)*model.Float32()
		pr.SellingPrice += chg
		pr.PurchasePrice += chg
	}
}
