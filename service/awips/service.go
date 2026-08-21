package awips

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides AWIPS (Automated Wireless IPS) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new AWIPS service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves the complete AWIPS operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessAWIPSOper, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOper](ctx, s.Client(), routes.AWIPSOperPath, opts...)
}

// ListAWIPSPerApInfo retrieves AWIPS per AP information.
func (s Service) ListAWIPSPerApInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSPerApInfo, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSPerApInfo](ctx, s.Client(), routes.AWIPSPerApInfoPath, opts...)
}

// ListAWIPSDwldStatus retrieves AWIPS download status.
// Note: Available on 17.12.6a, but unavailable on 17.15.4b.
func (s Service) ListAWIPSDwldStatus(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatus, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatus](ctx, s.Client(), routes.AWIPSDwldStatusPath, opts...)
}

// ListAWIPSApDwldStatus retrieves AWIPS per AP download status.
func (s Service) ListAWIPSApDwldStatus(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSApDwldStatus, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSApDwldStatus](
		ctx,
		s.Client(),
		routes.AWIPSApDownloadStatusPath,
		opts...,
	)
}

// ListAWIPSPerSignStats retrieves AWIPS per signature statistics.
func (s Service) ListAWIPSPerSignStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSPerSignStats, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSPerSignStats](
		ctx,
		s.Client(),
		routes.AWIPSPerSignStatsPath,
		opts...,
	)
}

// ListAWIPSGlobStats retrieves AWIPS global statistics.
func (s Service) ListAWIPSGlobStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSGlobStats, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSGlobStats](ctx, s.Client(), routes.AWIPSGlobStatsPath, opts...)
}

// ListAWIPSDwldStatusWncd retrieves AWIPS download status for WNCD.
func (s Service) ListAWIPSDwldStatusWncd(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatusWncd, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatusWncd](
		ctx,
		s.Client(),
		routes.AWIPSDwldStatusWncdPath,
		opts...,
	)
}
