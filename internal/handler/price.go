// Package handler handle
package handler

import (
	"context"

	pr "github.com/OVantsevich/PriceProvider/proto"

	"github.com/OVantsevich/PriceProvider/internal/model"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PriceService service interface for prices handler
//
//go:generate mockery --name=PriceService --case=underscore --output=./mocks
type PriceService interface {
	GetCurrentPrices(ctx context.Context, names []string) (map[string]*model.Price, error)
}

// Prices handler
type Prices struct {
	pr.UnimplementedPriceProviderServer
	service PriceService
}

// NewPrice constructor
func NewPrice(s PriceService) *Prices {
	return &Prices{service: s}
}

// GetCurrentPrices get current prices from price provider
func (h *Prices) GetCurrentPrices(ctx context.Context, request *pr.GetPricesRequest) (*pr.GetPricesResponse, error) {
	result, err := h.service.GetCurrentPrices(ctx, request.Names)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request.Names": request.Names,
		}).Errorf("prices - GetCurrentPrices - GetCurrentPrices: %v", err)
		return nil, status.Error(codes.Unknown, err.Error())
	}
	response := &pr.GetPricesResponse{
		Prices: make(map[string]*pr.Price),
	}
	for _, r := range result {
		response.Prices[r.Name] = &pr.Price{
			Name:          r.Name,
			SellingPrice:  r.SellingPrice,
			PurchasePrice: r.PurchasePrice,
		}
	}
	return response, err
}
