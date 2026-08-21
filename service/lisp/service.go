package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides LISP (Locator/ID Separation Protocol) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new LISP service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves LISP operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessLISPOper, error) {
	return core.Get[CiscoIOSXEWirelessLISPOper](ctx, s.Client(), routes.LISPOperPath, opts...)
}

// GetMemoryStats retrieves LISP agent memory statistics using wrapper struct.
func (s Service) GetMemoryStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessLISPOperLISPAgentMemoryStats, error) {
	return core.Get[CiscoIOSXEWirelessLISPOperLISPAgentMemoryStats](
		ctx,
		s.Client(),
		routes.LISPMemoryStatsPath,
		opts...,
	)
}

// GetCapabilities retrieves LISP WLC capabilities using wrapper struct.
func (s Service) GetCapabilities(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessLISPOperLISPWLCCapabilities, error) {
	return core.Get[CiscoIOSXEWirelessLISPOperLISPWLCCapabilities](
		ctx,
		s.Client(),
		routes.LISPCapabilitiesPath,
		opts...,
	)
}

// ListAPCapabilities retrieves LISP AP capabilities list using wrapper struct.
func (s Service) ListAPCapabilities(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessLISPOperLISPAPCapabilities, error) {
	return core.Get[CiscoIOSXEWirelessLISPOperLISPAPCapabilities](
		ctx,
		s.Client(),
		routes.LISPAPCapabilitiesPath,
		opts...,
	)
}
