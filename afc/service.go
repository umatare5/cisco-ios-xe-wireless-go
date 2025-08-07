package afc

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
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
	return &out, s.c.Do(ctx, http.MethodGet,
		"/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		&out)
}

// APResp returns per-AP AFC response data.
func (s Service) APResp(ctx context.Context) (*model.AfcOperEwlcAfcApRespResponse, error) {
	var out model.AfcOperEwlcAfcApRespResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
		&out)
}

// CloudOper returns AFC cloud operational data.
func (s Service) CloudOper(ctx context.Context) (*model.AfcCloudOperResponse, error) {
	var out model.AfcCloudOperResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data",
		&out)
}

// CloudStats returns AFC cloud statistics.
func (s Service) CloudStats(ctx context.Context) (*model.AfcCloudOperAfcCloudStatsResponse, error) {
	var out model.AfcCloudOperAfcCloudStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats",
		&out)
}
