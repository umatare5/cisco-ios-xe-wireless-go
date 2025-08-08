package rogue

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RogueOperBasePath defines the base path for rogue operational data endpoints.
	RogueOperBasePath = constants.YANGModelPrefix + "rogue-oper:rogue-oper-data"
	// RogueOperEndpoint defines the endpoint for rogue operational data.
	RogueOperEndpoint = RogueOperBasePath
	// RogueStatsEndpoint defines the endpoint for rogue statistics.
	RogueStatsEndpoint = RogueOperBasePath + "/rogue-stats"
	// RogueDataEndpoint defines the endpoint for rogue data.
	RogueDataEndpoint = RogueOperBasePath + "/rogue-data"
	// RogueClientDataEndpoint defines the endpoint for rogue client data.
	RogueClientDataEndpoint = RogueOperBasePath + "/rogue-client-data"
	// RldpStatsEndpoint defines the endpoint for RLDP (Rogue Location Discovery Protocol) statistics.
	RldpStatsEndpoint = RogueOperBasePath + "/rldp-stats"
)

// Service provides Rogue operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// GetOper returns rogue operational data.
func (s Service) GetOper(ctx context.Context) (*model.RogueOperResponse, error) {
	return core.Get[model.RogueOperResponse](ctx, s.c, RogueOperEndpoint)
}

// GetStats returns rogue statistics.
func (s Service) GetStats(ctx context.Context) (*model.RogueStatsResponse, error) {
	return core.Get[model.RogueStatsResponse](ctx, s.c, RogueStatsEndpoint)
}

// GetData returns rogue data.
func (s Service) GetData(ctx context.Context) (*model.RogueDataResponse, error) {
	return core.Get[model.RogueDataResponse](ctx, s.c, RogueDataEndpoint)
}

// GetClientData returns rogue client data.
func (s Service) GetClientData(ctx context.Context) (*model.RogueClientDataResponse, error) {
	return core.Get[model.RogueClientDataResponse](ctx, s.c, RogueClientDataEndpoint)
}

// GetRldpStats returns RLDP statistics.
func (s Service) GetRldpStats(ctx context.Context) (*model.RldpStatsResponse, error) {
	return core.Get[model.RldpStatsResponse](ctx, s.c, RldpStatsEndpoint)
}
