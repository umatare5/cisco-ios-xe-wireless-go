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
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessAFCOper, error) {
	return core.Get[CiscoIOSXEWirelessAFCOper](ctx, s.Client(), routes.AFCOperPath)
}

// ListAPResponses retrieves AFC AP response data.
func (s Service) ListAPResponses(ctx context.Context) (*CiscoIOSXEWirelessAFCOperEwlcAFCApResp, error) {
	return core.Get[CiscoIOSXEWirelessAFCOperEwlcAFCApResp](ctx, s.Client(), routes.AFCEwlcAFCApRespPath)
}

// ListAPRequests retrieves AFC AP request data.
func (s Service) ListAPRequests(ctx context.Context) (*CiscoIOSXEWirelessAFCOperEwlcAFCApReq, error) {
	return core.Get[CiscoIOSXEWirelessAFCOperEwlcAFCApReq](ctx, s.Client(), routes.AFCEwlcAFCApReqPath)
}

// GetCloudInfo retrieves AFC cloud operational data.
func (s Service) GetCloudInfo(ctx context.Context) (*CiscoIOSXEWirelessAFCCloudOper, error) {
	return core.Get[CiscoIOSXEWirelessAFCCloudOper](ctx, s.Client(), routes.AFCCloudOperPath)
}

// GetCloudStats retrieves AFC cloud statistics.
func (s Service) GetCloudStats(ctx context.Context) (*CiscoIOSXEWirelessAFCCloudOperAFCCloudStats, error) {
	return core.Get[CiscoIOSXEWirelessAFCCloudOperAFCCloudStats](ctx, s.Client(), routes.AFCAfcCloudStatsPath)
}
