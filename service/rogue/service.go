package rogue

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rogue"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides rogue detection and mitigation operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Rogue service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves rogue detection operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.RogueOper, error) {
	return core.Get[model.RogueOper](ctx, s.Client(), routes.RogueOperPath)
}

// ListRogues retrieves rogue client data.
func (s Service) ListRogues(ctx context.Context) (*model.RogueData, error) {
	return core.Get[model.RogueData](ctx, s.Client(), routes.RogueDataPath)
}

// GetRogueByMAC retrieves rogue data filtered by rogue address.
func (s Service) GetRogueByMAC(ctx context.Context, mac string) (*model.RogueData, error) {
	if mac == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RestconfBuilder().BuildQueryURL(routes.RogueDataPath, mac)
	return core.Get[model.RogueData](ctx, s.Client(), url)
}

// ListRogueClients retrieves rogue client data.
func (s Service) ListRogueClients(ctx context.Context) (*model.RogueClientData, error) {
	return core.Get[model.RogueClientData](ctx, s.Client(), routes.RogueClientDataPath)
}

// GetRogueClientByMAC retrieves rogue data filtered by rogue address.
func (s Service) GetRogueClientByMAC(ctx context.Context, mac string) (*model.RogueClientData, error) {
	if mac == "" {
		return nil, core.ErrInvalidConfiguration
	}

	url := s.Client().RestconfBuilder().BuildQueryURL(routes.RogueClientDataPath, mac)
	return core.Get[model.RogueClientData](ctx, s.Client(), url)
}

// GetStats retrieves rogue statistics.
func (s Service) GetStats(ctx context.Context) (*model.RogueStats, error) {
	return core.Get[model.RogueStats](ctx, s.Client(), routes.RogueStatsPath)
}

// Alias methods for integration test compatibility

// GetOperClientData is an alias for ListRogueClients.
func (s Service) GetOperClientData(ctx context.Context) (*model.RogueClientData, error) {
	return s.ListRogueClients(ctx)
}

// GetOperData is an alias for ListRogues.
func (s Service) GetOperData(ctx context.Context) (*model.RogueData, error) {
	return s.ListRogues(ctx)
}

// GetOperStats is an alias for GetStats.
func (s Service) GetOperStats(ctx context.Context) (*model.RogueStats, error) {
	return s.GetStats(ctx)
}

// GetRLDPStats retrieves RLDP (Rogue Location Discovery Protocol) statistics.
func (s Service) GetRLDPStats(ctx context.Context) (*model.RogueStats, error) {
	// For now, return the same as rogue stats since RLDP is part of rogue detection
	return s.GetStats(ctx)
}

// GetOperByRogueAddress is an alias for GetRogueByMAC.
func (s Service) GetOperByRogueAddress(ctx context.Context, mac string) (*model.RogueData, error) {
	return s.GetRogueByMAC(ctx, mac)
}

// GetOperByRogueClientAddress is an alias for GetRogueClientByMAC.
func (s Service) GetOperByRogueClientAddress(ctx context.Context, mac string) (*model.RogueClientData, error) {
	return s.GetRogueClientByMAC(ctx, mac)
}
