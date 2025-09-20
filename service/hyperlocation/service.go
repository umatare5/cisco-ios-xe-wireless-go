package hyperlocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides Hyperlocation operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Hyperlocation service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves hyperlocation operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*HyperlocationOper, error) {
	return core.Get[HyperlocationOper](ctx, s.Client(), routes.HyperlocationOperPath)
}

// ListProfiles retrieves hyperlocation profiles.
func (s Service) ListProfiles(ctx context.Context) (*HyperlocationProfiles, error) {
	return core.Get[HyperlocationProfiles](ctx, s.Client(), routes.HyperlocationProfilesPath)
}
