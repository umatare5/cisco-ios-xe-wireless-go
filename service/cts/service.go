package cts

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides CTS (Cisco TrustSec) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new CTS service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves CTS configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*CTSCfg, error) {
	return core.Get[CTSCfg](ctx, s.Client(), routes.CTSCfgPath)
}
