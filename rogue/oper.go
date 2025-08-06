// Package rogue provides rogue access point detection operational data functionality for the Cisco Wireless Network Controller API.
package rogue

import (
	"context"
	"fmt"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
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

// Type aliases for backward compatibility.
type (
	// Deprecated: Use model.RogueOperResponse instead.
	RogueOperResponse = model.RogueOperResponse
	// Deprecated: Use model.RogueStatsResponse instead.
	RogueStatsResponse = model.RogueStatsResponse
	// Deprecated: Use model.RogueDataResponse instead.
	RogueDataResponse = model.RogueDataResponse
	// Deprecated: Use model.RogueClientDataResponse instead.
	RogueClientDataResponse = model.RogueClientDataResponse
	// Deprecated: Use model.RldpStatsResponse instead.
	RldpStatsResponse = model.RldpStatsResponse
	// Deprecated: Use model.RogueStats instead.
	RogueStats = model.RogueStats
	// Deprecated: Use model.RogueData instead.
	RogueData = model.RogueData
	// Deprecated: Use model.RogueClientData instead.
	RogueClientData = model.RogueClientData
	// Deprecated: Use model.RldpStats instead.
	RldpStats = model.RldpStats
)

// Deprecated: Use client.Rogue().Oper(ctx) instead.
// GetRogueOper retrieves rogue operational data.
func GetRogueOper(client *wnc.Client, ctx context.Context) (*RogueOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.Oper(ctx)
}

// Deprecated: Use client.Rogue().Stats(ctx) instead.
// GetRogueStats retrieves rogue statistics.
func GetRogueStats(client *wnc.Client, ctx context.Context) (*RogueStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.Stats(ctx)
}

// Deprecated: Use client.Rogue().Data(ctx) instead.
// GetRogueData retrieves rogue data.
func GetRogueData(client *wnc.Client, ctx context.Context) (*RogueDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.Data(ctx)
}

// Deprecated: Use client.Rogue().ClientData(ctx) instead.
// GetRogueClientData retrieves rogue client data.
func GetRogueClientData(client *wnc.Client, ctx context.Context) (*RogueClientDataResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.ClientData(ctx)
}

// Deprecated: Use client.Rogue().RldpStats(ctx) instead.
// GetRldpStats retrieves RLDP statistics.
func GetRldpStats(client *wnc.Client, ctx context.Context) (*RldpStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.RldpStats(ctx)
}
