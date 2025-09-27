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
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOper, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOper](ctx, s.Client(), routes.AWIPSOperPath)
}

// ListAWIPSPerApInfo retrieves AWIPS per AP information.
func (s Service) ListAWIPSPerApInfo(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSPerApInfo, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSPerApInfo](ctx, s.Client(), routes.AWIPSPerApInfoPath)
}

// ListAWIPSDwldStatus retrieves AWIPS download status.
// Note: Available on 17.12.5, Unavailable on 17.15.4b.
func (s Service) ListAWIPSDwldStatus(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatus, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatus](ctx, s.Client(), routes.AWIPSDwldStatusPath)
}

// ListAWIPSApDwldStatus retrieves AWIPS per AP download status.
func (s Service) ListAWIPSApDwldStatus(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSApDwldStatus, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSApDwldStatus](ctx, s.Client(), routes.AWIPSApDownloadStatusPath)
}

// ListAWIPSPerSignStats retrieves AWIPS per signature statistics.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSPerSignStats(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSPerSignStats, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSPerSignStats](ctx, s.Client(), routes.AWIPSPerSignStatsPath)
}

// ListAWIPSGlobStats retrieves AWIPS global statistics.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSGlobStats(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSGlobStats, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSGlobStats](ctx, s.Client(), routes.AWIPSGlobStatsPath)
}

// ListAWIPSDwldStatusWncd retrieves AWIPS download status for WNCD.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSDwldStatusWncd(ctx context.Context) (*CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatusWncd, error) {
	return core.Get[CiscoIOSXEWirelessAWIPSOperAWIPSDwldStatusWncd](ctx, s.Client(), routes.AWIPSDwldStatusWncdPath)
}
