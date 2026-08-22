package rogue

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRogueOper, error) {
	return core.Get[CiscoIOSXEWirelessRogueOper](ctx, s.Client(), routes.RogueOperPath, opts...)
}

// ListRogues retrieves rogue client data.
func (s Service) ListRogues(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRogueData, error) {
	return core.Get[CiscoIOSXEWirelessRogueData](ctx, s.Client(), routes.RogueDataPath, opts...)
}

// GetRogueByMAC retrieves rogue data filtered by rogue address.
func (s Service) GetRogueByMAC(
	ctx context.Context,
	mac string,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRogueData, error) {
	normalizedMAC, err := service.RequireMACAddress(mac)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RogueDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessRogueData](ctx, s.Client(), url, opts...)
}

// ListRogueClients retrieves rogue client data.
func (s Service) ListRogueClients(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRogueClientData, error) {
	return core.Get[CiscoIOSXEWirelessRogueClientData](ctx, s.Client(), routes.RogueClientDataPath, opts...)
}

// GetRogueClientByMAC retrieves rogue data filtered by rogue address.
func (s Service) GetRogueClientByMAC(
	ctx context.Context,
	mac string,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRogueClientData, error) {
	normalizedMAC, err := service.RequireMACAddress(mac)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RogueClientDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessRogueClientData](ctx, s.Client(), url, opts...)
}

// GetStats retrieves rogue statistics.
func (s Service) GetStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRogueOperRogueStats, error) {
	return core.Get[CiscoIOSXEWirelessRogueOperRogueStats](ctx, s.Client(), routes.RogueStatsPath, opts...)
}

// GetRLDPStats retrieves RLDP statistics.
//
// The controller declares this node as a presence container with status deprecated in every
// release measured (17.12, 17.15, 17.18 and 26.1), and answers it with data on all four.
// A presence container exists only once the feature instantiates it, so a 404 here reports a
// chassis where RLDP was never enabled, not a route the release lacks.
func (s Service) GetRLDPStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRogueOperRLDPStats, error) {
	return core.Get[CiscoIOSXEWirelessRogueOperRLDPStats](ctx, s.Client(), routes.RogueRLDPStatsPath, opts...)
}
