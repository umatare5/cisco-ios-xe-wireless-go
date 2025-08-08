package afc

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// AFCOperBasePath defines the base path for AFC operational data endpoints
	AFCOperBasePath = constants.YANGModelPrefix + "afc-oper:afc-oper-data"
	// AFCOperEndpoint retrieves overall AFC operational data
	AFCOperEndpoint = AFCOperBasePath
	// AFCApRespEndpoint retrieves per-AP AFC response data
	AFCApRespEndpoint = AFCOperBasePath + "/ewlc-afc-ap-resp"

	// AFCCloudOperBasePath defines the base path for AFC cloud operational data endpoints
	AFCCloudOperBasePath = constants.YANGModelPrefix + "afc-cloud-oper:afc-cloud-oper-data"
	// AFCCloudOperEndpoint retrieves AFC cloud operational data
	AFCCloudOperEndpoint = AFCCloudOperBasePath
	// AFCCloudStatsEndpoint retrieves AFC cloud statistics
	AFCCloudStatsEndpoint = AFCCloudOperBasePath + "/afc-cloud-stats"
)

// Service provides AFC operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// GetOper returns overall AFC operational data.
func (s Service) GetOper(ctx context.Context) (*model.AfcOperResponse, error) {
	return core.Get[model.AfcOperResponse](ctx, s.c, AFCOperEndpoint)
}

// GetAPResp returns per-AP AFC response data.
func (s Service) GetAPResp(ctx context.Context) (*model.AfcOperEwlcAfcApRespResponse, error) {
	return core.Get[model.AfcOperEwlcAfcApRespResponse](ctx, s.c, AFCApRespEndpoint)
}

// GetCloudOper returns AFC cloud operational data.
func (s Service) GetCloudOper(ctx context.Context) (*model.AfcCloudOperResponse, error) {
	return core.Get[model.AfcCloudOperResponse](ctx, s.c, AFCCloudOperEndpoint)
}

// GetCloudStats returns AFC cloud statistics.
func (s Service) GetCloudStats(ctx context.Context) (*model.AfcCloudOperAfcCloudStatsResponse, error) {
	return core.Get[model.AfcCloudOperAfcCloudStatsResponse](ctx, s.c, AFCCloudStatsEndpoint)
}
