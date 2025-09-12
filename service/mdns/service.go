package mdns

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mdns"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides mDNS (Multicast DNS) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new MDNS service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves all mDNS operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*model.MdnsOper, error) {
	return core.Get[model.MdnsOper](ctx, s.Client(), routes.MDNSOperPath)
}

// GetGlobalStats retrieves mDNS global statistics from the wireless controller.
func (s Service) GetGlobalStats(ctx context.Context) (*model.MdnsGlobalStats, error) {
	return core.Get[model.MdnsGlobalStats](ctx, s.Client(), routes.MDNSGlobalStatsPath)
}

// ListWLANStats retrieves mDNS WLAN statistics from the wireless controller.
func (s Service) ListWLANStats(ctx context.Context) (*model.MdnsWlanStats, error) {
	return core.Get[model.MdnsWlanStats](ctx, s.Client(), routes.MDNSWlanStatsPath)
}
