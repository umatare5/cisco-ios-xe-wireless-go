package ap

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// APCfgBasePath defines the base path for access point configuration endpoints
	APCfgBasePath = constants.YANGModelPrefix + "ap-cfg:ap-cfg-data"
	// APCfgEndpoint retrieves complete access point configuration data
	APCfgEndpoint = APCfgBasePath
	// TagSourcePriorityConfigsEndpoint retrieves tag source priority configurations
	TagSourcePriorityConfigsEndpoint = APCfgBasePath + "/tag-source-priority-configs"
	// APTagsEndpoint retrieves access point tag configurations
	APTagsEndpoint = APCfgBasePath + "/ap-tags"

	// APOperBasePath defines the base path for access point operational endpoints
	APOperBasePath = constants.YANGModelPrefix + "access-point-oper:access-point-oper-data"
	// APOperEndpoint retrieves complete access point operational data
	APOperEndpoint = APOperBasePath
	// APRadioNeighborEndpoint retrieves access point radio neighbor information
	APRadioNeighborEndpoint = APOperBasePath + "/ap-radio-neighbor"
	// RadioOperDataEndpoint retrieves radio operational data for access points
	RadioOperDataEndpoint = APOperBasePath + "/radio-oper-data"
	// QosClientDataEndpoint retrieves QoS client data information
	QosClientDataEndpoint = APOperBasePath + "/qos-client-data"
	// CapwapDataEndpoint retrieves CAPWAP data for access points
	CapwapDataEndpoint = APOperBasePath + "/capwap-data"
	// APNameMacMapEndpoint retrieves AP name to MAC address mapping
	APNameMacMapEndpoint = APOperBasePath + "/ap-name-mac-map"

	// APGlobalOperBasePath defines the base path for global access point operational endpoints
	APGlobalOperBasePath = constants.YANGModelPrefix + "ap-global-oper:ap-global-oper-data"
	// APGlobalOperEndpoint retrieves complete AP global operational data
	APGlobalOperEndpoint = APGlobalOperBasePath
	// APHistoryEndpoint retrieves AP history data
	APHistoryEndpoint = APGlobalOperBasePath + "/ap-history"
	// EwlcAPStatsEndpoint retrieves EWLC AP statistics
	EwlcAPStatsEndpoint = APGlobalOperBasePath + "/ewlc-ap-stats"
)

// Service provides access point operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Configuration Methods

// GetCfg returns complete access point configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.ApCfgResponse, error) {
	return core.Get[model.ApCfgResponse](ctx, s.c, APCfgEndpoint)
}

// GetTagSourcePriorityConfigs returns tag source priority configurations.
func (s Service) GetTagSourcePriorityConfigs(ctx context.Context) (*model.TagSourcePriorityConfigs, error) {
	return core.Get[model.TagSourcePriorityConfigs](ctx, s.c, TagSourcePriorityConfigsEndpoint)
}

// GetApTags returns access point tag configurations.
func (s Service) GetApTags(ctx context.Context) (*model.ApCfgApTagsResponse, error) {
	return core.Get[model.ApCfgApTagsResponse](ctx, s.c, APTagsEndpoint)
}

// Operational Methods

// GetOper returns complete access point operational data.
func (s Service) GetOper(ctx context.Context) (*model.ApOperResponse, error) {
	return core.Get[model.ApOperResponse](ctx, s.c, APOperEndpoint)
}

// GetRadioNeighbor returns access point radio neighbor information.
func (s Service) GetRadioNeighbor(ctx context.Context) (*model.ApOperApRadioNeighborResponse, error) {
	return core.Get[model.ApOperApRadioNeighborResponse](ctx, s.c, APRadioNeighborEndpoint)
}

// GetNameMacMap returns the mapping between AP names and MAC addresses.
func (s Service) GetNameMacMap(ctx context.Context) (*[]model.ApNameMacMap, error) {
	// local response wrapper to align with core.Get usage pattern
	type apNameMacMapResponse struct {
		Data []model.ApNameMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
	}
	resp, err := core.Get[apNameMacMapResponse](ctx, s.c, APNameMacMapEndpoint)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// GetCapwapData returns CAPWAP protocol data.
func (s Service) GetCapwapData(ctx context.Context) (*[]model.CapwapData, error) {
	// local response wrapper to align with core.Get usage pattern
	type capwapDataResponse struct {
		Data []model.CapwapData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}
	resp, err := core.Get[capwapDataResponse](ctx, s.c, CapwapDataEndpoint)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// Global Operational Methods

// GetGlobalOper returns complete AP global operational data.
func (s Service) GetGlobalOper(ctx context.Context) (*model.ApGlobalOperResponse, error) {
	return core.Get[model.ApGlobalOperResponse](ctx, s.c, APGlobalOperEndpoint)
}

// GetHistory returns AP history data.
func (s Service) GetHistory(ctx context.Context) (*model.ApGlobalOperApHistoryResponse, error) {
	return core.Get[model.ApGlobalOperApHistoryResponse](ctx, s.c, APHistoryEndpoint)
}

// GetEwlcApStats returns EWLC AP statistics.
func (s Service) GetEwlcApStats(ctx context.Context) (*model.ApGlobalOperEwlcApStatsResponse, error) {
	return core.Get[model.ApGlobalOperEwlcApStatsResponse](ctx, s.c, EwlcAPStatsEndpoint)
}
