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

// Service provides rogue access point detection operational data access.
type Service struct {
	c *wnc.Client
}

// NewService creates a new rogue service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Oper retrieves rogue operational data.
func (s Service) Oper(ctx context.Context) (*model.RogueOperResponse, error) {
	var out model.RogueOperResponse
	err := s.c.Do(ctx, http.MethodGet, RogueOperEndpoint, &out)
	return &out, err
}

// Stats retrieves rogue statistics.
func (s Service) Stats(ctx context.Context) (*model.RogueStatsResponse, error) {
	var out model.RogueStatsResponse
	err := s.c.Do(ctx, http.MethodGet, RogueStatsEndpoint, &out)
	return &out, err
}

// Data retrieves rogue data.
func (s Service) Data(ctx context.Context) (*model.RogueDataResponse, error) {
	var out model.RogueDataResponse
	err := s.c.Do(ctx, http.MethodGet, RogueDataEndpoint, &out)
	return &out, err
}

// ClientData retrieves rogue client data.
func (s Service) ClientData(ctx context.Context) (*model.RogueClientDataResponse, error) {
	var out model.RogueClientDataResponse
	err := s.c.Do(ctx, http.MethodGet, RogueClientDataEndpoint, &out)
	return &out, err
}

// RldpStats retrieves RLDP statistics.
func (s Service) RldpStats(ctx context.Context) (*model.RldpStatsResponse, error) {
	var out model.RldpStatsResponse
	err := s.c.Do(ctx, http.MethodGet, RldpStatsEndpoint, &out)
	return &out, err
}
