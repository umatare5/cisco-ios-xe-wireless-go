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
func (s Service) GetOperational(ctx context.Context) (*AWIPSOper, error) {
	return core.Get[AWIPSOper](ctx, s.Client(), routes.AWIPSOperPath)
}

// ListAWIPSPerApInfo retrieves AWIPS per AP information.
func (s Service) ListAWIPSPerApInfo(ctx context.Context) (*AWIPSOperAWIPSPerApInfo, error) {
	return core.Get[AWIPSOperAWIPSPerApInfo](ctx, s.Client(), routes.AWIPSPerApInfoPath)
}

// ListAWIPSDwldStatus retrieves AWIPS download status.
func (s Service) ListAWIPSDwldStatus(ctx context.Context) (*AWIPSOperAWIPSDwldStatus, error) {
	return core.Get[AWIPSOperAWIPSDwldStatus](ctx, s.Client(), routes.AWIPSDwldStatusPath)
}

// ListAWIPSApDwldStatus retrieves AWIPS per AP download status.
func (s Service) ListAWIPSApDwldStatus(ctx context.Context) (*AWIPSOperAWIPSApDwldStatus, error) {
	return core.Get[AWIPSOperAWIPSApDwldStatus](ctx, s.Client(), routes.AWIPSApDownloadStatusPath)
}

// ListAWIPSPerSignStats retrieves AWIPS per signature statistics.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSPerSignStats(ctx context.Context) (*AWIPSOperAWIPSPerSignStats, error) {
	return core.Get[AWIPSOperAWIPSPerSignStats](ctx, s.Client(), routes.AWIPSPerSignStatsPath)
}

// ListAWIPSGlobStats retrieves AWIPS global statistics.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSGlobStats(ctx context.Context) (*AWIPSOperAWIPSGlobStats, error) {
	return core.Get[AWIPSOperAWIPSGlobStats](ctx, s.Client(), routes.AWIPSGlobStatsPath)
}

// ListAWIPSDwldStatusWncd retrieves AWIPS download status for WNCD.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListAWIPSDwldStatusWncd(ctx context.Context) (*AWIPSOperAWIPSDwldStatusWncd, error) {
	return core.Get[AWIPSOperAWIPSDwldStatusWncd](ctx, s.Client(), routes.AWIPSDwldStatusWncdPath)
}
