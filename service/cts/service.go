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

// GetConfig retrieves CTS SXP configuration.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessCTSCfg, error) {
	return core.Get[CiscoIOSXEWirelessCTSCfg](ctx, s.Client(), routes.CTSCfgPath, opts...)
}

// GetOperational retrieves CTS operational data from the controller.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessCTSOper, error) {
	return core.Get[CiscoIOSXEWirelessCTSOper](ctx, s.Client(), routes.CTSOperPath, opts...)
}

// ListFlexModeApSxpConnectionStatus retrieves FlexConnect AP SXP connection status.
func (s Service) ListFlexModeApSxpConnectionStatus(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessCTSOperFlexModeApSxpConnectionStatus, error) {
	return core.Get[CiscoIOSXEWirelessCTSOperFlexModeApSxpConnectionStatus](
		ctx,
		s.Client(),
		routes.CTSFlexModeApSxpConnectionStatusPath, opts...,
	)
}
