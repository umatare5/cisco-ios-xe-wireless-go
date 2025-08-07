package rogue

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// RogueOperBasePath defines the base path for rogue operational data endpoints.
	RogueOperBasePath = "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data"
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
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Oper returns rogue operational data.
func (s Service) Oper(ctx context.Context) (*model.RogueOperResponse, error) {
	var out model.RogueOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, RogueOperEndpoint, &out)
}

// Stats returns rogue statistics.
func (s Service) Stats(ctx context.Context) (*model.RogueStatsResponse, error) {
	var out model.RogueStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, RogueStatsEndpoint, &out)
}

// Data returns rogue data.
func (s Service) Data(ctx context.Context) (*model.RogueDataResponse, error) {
	var out model.RogueDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, RogueDataEndpoint, &out)
}

// ClientData returns rogue client data.
func (s Service) ClientData(ctx context.Context) (*model.RogueClientDataResponse, error) {
	var out model.RogueClientDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, RogueClientDataEndpoint, &out)
}

// RldpStats returns RLDP statistics.
func (s Service) RldpStats(ctx context.Context) (*model.RldpStatsResponse, error) {
	var out model.RldpStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, RldpStatsEndpoint, &out)
}
