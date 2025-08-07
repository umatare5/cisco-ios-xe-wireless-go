package afc

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// AfcOperBasePath defines the base path for AFC operational data endpoints
	AfcOperBasePath = constants.YANGModelPrefix + "afc-oper:afc-oper-data"
	// AfcOperEndpoint retrieves overall AFC operational data
	AfcOperEndpoint = AfcOperBasePath
	// AfcApRespEndpoint retrieves per-AP AFC response data
	AfcApRespEndpoint = AfcOperBasePath + "/ewlc-afc-ap-resp"

	// AfcCloudOperBasePath defines the base path for AFC cloud operational data endpoints
	AfcCloudOperBasePath = constants.YANGModelPrefix + "afc-cloud-oper:afc-cloud-oper-data"
	// AfcCloudOperEndpoint retrieves AFC cloud operational data
	AfcCloudOperEndpoint = AfcCloudOperBasePath
	// AfcCloudStatsEndpoint retrieves AFC cloud statistics
	AfcCloudStatsEndpoint = AfcCloudOperBasePath + "/afc-cloud-stats"
)

// Service provides AFC operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Oper returns overall AFC operational data.
func (s Service) Oper(ctx context.Context) (*model.AfcOperResponse, error) {
	var out model.AfcOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, AfcOperEndpoint, &out)
}

// APResp returns per-AP AFC response data.
func (s Service) APResp(ctx context.Context) (*model.AfcOperEwlcAfcApRespResponse, error) {
	var out model.AfcOperEwlcAfcApRespResponse
	return &out, s.c.Do(ctx, http.MethodGet, AfcApRespEndpoint, &out)
}

// CloudOper returns AFC cloud operational data.
func (s Service) CloudOper(ctx context.Context) (*model.AfcCloudOperResponse, error) {
	var out model.AfcCloudOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, AfcCloudOperEndpoint, &out)
}

// CloudStats returns AFC cloud statistics.
func (s Service) CloudStats(ctx context.Context) (*model.AfcCloudOperAfcCloudStatsResponse, error) {
	var out model.AfcCloudOperAfcCloudStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, AfcCloudStatsEndpoint, &out)
}
