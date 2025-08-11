package mdns

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// MdnsOperBasePath defines the base path for mDNS operational data endpoints.
	MdnsOperBasePath = constants.YANGModelPrefix + "mdns-oper:mdns-oper-data"
	// MdnsOperEndpoint defines the endpoint for mDNS operational data.
	MdnsOperEndpoint = MdnsOperBasePath
	// MdnsGlobalStatsEndpoint defines the endpoint for mDNS global statistics.
	MdnsGlobalStatsEndpoint = MdnsOperBasePath + "/mdns-global-stats"
	// MdnsWlanStatsEndpoint defines the endpoint for mDNS WLAN statistics.
	MdnsWlanStatsEndpoint = MdnsOperBasePath + "/mdns-wlan-stats"
)

// Service provides MDNS operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// GetOper returns mDNS operational data.
func (s Service) GetOper(ctx context.Context) (*model.MdnsOperResponse, error) {
	return core.Get[model.MdnsOperResponse](ctx, s.c, MdnsOperEndpoint)
}

// GetGlobalStats returns mDNS global statistics.
func (s Service) GetGlobalStats(ctx context.Context) (*model.MdnsGlobalStatsResponse, error) {
	return core.Get[model.MdnsGlobalStatsResponse](ctx, s.c, MdnsGlobalStatsEndpoint)
}

// GetWlanStats returns mDNS WLAN statistics.
func (s Service) GetWlanStats(ctx context.Context) (*model.MdnsWlanStatsResponse, error) {
	return core.Get[model.MdnsWlanStatsResponse](ctx, s.c, MdnsWlanStatsEndpoint)
}
