// Package mdns provides multicast DNS operational data functionality for the Cisco Wireless Network Controller API.
package mdns

import (
	"context"
	"fmt"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// MdnsOperBasePath defines the base path for mDNS operational data endpoints.
	MdnsOperBasePath = "Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"
	// MdnsOperEndpoint defines the endpoint for mDNS operational data.
	MdnsOperEndpoint = MdnsOperBasePath
	// MdnsGlobalStatsEndpoint defines the endpoint for mDNS global statistics.
	MdnsGlobalStatsEndpoint = MdnsOperBasePath + "/mdns-global-stats"
	// MdnsWlanStatsEndpoint defines the endpoint for mDNS WLAN statistics.
	MdnsWlanStatsEndpoint = MdnsOperBasePath + "/mdns-wlan-stats"
)

// Type aliases for backward compatibility - will be removed in v2.0.0
type (
	// Deprecated: Use model.MdnsOperResponse instead. Will be removed in v2.0.0.
	MdnsOperResponse = model.MdnsOperResponse
	// Deprecated: Use model.MdnsGlobalStatsResponse instead. Will be removed in v2.0.0.
	MdnsGlobalStatsResponse = model.MdnsGlobalStatsResponse
	// Deprecated: Use model.MdnsWlanStatsResponse instead. Will be removed in v2.0.0.
	MdnsWlanStatsResponse = model.MdnsWlanStatsResponse
	// Deprecated: Use model.MdnsGlobalStats instead. Will be removed in v2.0.0.
	MdnsGlobalStats = model.MdnsGlobalStats
	// Deprecated: Use model.MdnsWlanStat instead. Will be removed in v2.0.0.
	MdnsWlanStat = model.MdnsWlanStat
	// Deprecated: Use model.MdnsStats instead. Will be removed in v2.0.0.
	MdnsStats = model.MdnsStats
)

// Deprecated: Use mdns.NewService(client).Oper(ctx) instead. Will be removed in v2.0.0.
func GetMdnsOper(client *wnc.Client, ctx context.Context) (*model.MdnsOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	return NewService(client).Oper(ctx)
}

// Deprecated: Use mdns.NewService(client).GlobalStats(ctx) instead. Will be removed in v2.0.0.
func GetMdnsGlobalStats(client *wnc.Client, ctx context.Context) (*model.MdnsGlobalStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	return NewService(client).GlobalStats(ctx)
}

// Deprecated: Use mdns.NewService(client).WlanStats(ctx) instead. Will be removed in v2.0.0.
func GetMdnsWlanStats(client *wnc.Client, ctx context.Context) (*model.MdnsWlanStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	return NewService(client).WlanStats(ctx)
}
