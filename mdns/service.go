package mdns

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
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

// Oper returns mDNS operational data.
func (s Service) Oper(ctx context.Context) (*model.MdnsOperResponse, error) {
	var out model.MdnsOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, MdnsOperEndpoint, &out)
}

// GlobalStats returns mDNS global statistics.
func (s Service) GlobalStats(ctx context.Context) (*model.MdnsGlobalStatsResponse, error) {
	var out model.MdnsGlobalStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, MdnsGlobalStatsEndpoint, &out)
}

// WlanStats returns mDNS WLAN statistics.
func (s Service) WlanStats(ctx context.Context) (*model.MdnsWlanStatsResponse, error) {
	var out model.MdnsWlanStatsResponse
	return &out, s.c.Do(ctx, http.MethodGet, MdnsWlanStatsEndpoint, &out)
}
