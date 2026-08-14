package afc

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides AFC (Automated Frequency Coordination) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new AFC service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves AFC operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessAFCOper, error) {
	return core.Get[CiscoIOSXEWirelessAFCOper](ctx, s.Client(), routes.AFCOperPath, opts...)
}

// ListAPResponses retrieves AFC AP response data.
func (s Service) ListAPResponses(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAFCOperEwlcAFCApResp, error) {
	return core.Get[CiscoIOSXEWirelessAFCOperEwlcAFCApResp](ctx, s.Client(), routes.AFCEwlcAFCApRespPath, opts...)
}

// ListAPRequests retrieves AFC AP request data.
func (s Service) ListAPRequests(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAFCOperEwlcAFCApReq, error) {
	return core.Get[CiscoIOSXEWirelessAFCOperEwlcAFCApReq](ctx, s.Client(), routes.AFCEwlcAFCApReqPath, opts...)
}

// GetCloudInfo retrieves AFC cloud operational data.
func (s Service) GetCloudInfo(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessAFCCloudOper, error) {
	return core.Get[CiscoIOSXEWirelessAFCCloudOper](ctx, s.Client(), routes.AFCCloudOperPath, opts...)
}

// GetCloudStats retrieves AFC cloud statistics.
func (s Service) GetCloudStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAFCCloudOperAFCCloudStats, error) {
	return core.Get[CiscoIOSXEWirelessAFCCloudOperAFCCloudStats](ctx, s.Client(), routes.AFCAfcCloudStatsPath, opts...)
}
