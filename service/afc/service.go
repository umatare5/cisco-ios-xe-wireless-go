package afc

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/afc"
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

// GetOperational retrieves the complete AFC operational data.
func (s Service) GetOperational(ctx context.Context) (*model.AFCOper, error) {
	return core.Get[model.AFCOper](ctx, s.Client(), routes.AFCOperPath)
}

// ListAPResponses retrieves AFC AP response data.
func (s Service) ListAPResponses(ctx context.Context) (*model.AFCOperEwlcAFCApResp, error) {
	return core.Get[model.AFCOperEwlcAFCApResp](ctx, s.Client(), routes.AFCEwlcAFCApRespPath)
}

// GetCloudInfo retrieves AFC cloud operational data.
func (s Service) GetCloudInfo(ctx context.Context) (*model.AFCCloudOper, error) {
	return core.Get[model.AFCCloudOper](ctx, s.Client(), routes.AFCCloudOperPath)
}

// GetCloudStats retrieves AFC cloud statistics.
func (s Service) GetCloudStats(ctx context.Context) (*model.AFCCloudOperAFCCloudStats, error) {
	return core.Get[model.AFCCloudOperAFCCloudStats](ctx, s.Client(), routes.AFCAfcCloudStatsPath)
}
