package ap

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// ApCfgBasePath defines the base path for access point configuration endpoints
	ApCfgBasePath = constants.YANGModelPrefix + "ap-cfg:ap-cfg-data"
	// ApCfgEndpoint retrieves complete access point configuration data
	ApCfgEndpoint = ApCfgBasePath
	// TagSourcePriorityConfigsEndpoint retrieves tag source priority configurations
	TagSourcePriorityConfigsEndpoint = ApCfgBasePath + "/tag-source-priority-configs"
	// ApTagsEndpoint retrieves access point tag configurations
	ApTagsEndpoint = ApCfgBasePath + "/ap-tags"

	// ApOperBasePath defines the base path for access point operational endpoints
	ApOperBasePath = constants.YANGModelPrefix + "access-point-oper:access-point-oper-data"
	// ApOperEndpoint retrieves complete access point operational data
	ApOperEndpoint = ApOperBasePath
	// ApRadioNeighborEndpoint retrieves access point radio neighbor information
	ApRadioNeighborEndpoint = ApOperBasePath + "/ap-radio-neighbor"
	// RadioOperDataEndpoint retrieves radio operational data for access points
	RadioOperDataEndpoint = ApOperBasePath + "/radio-oper-data"
	// QosClientDataEndpoint retrieves QoS client data information
	QosClientDataEndpoint = ApOperBasePath + "/qos-client-data"
	// CapwapDataEndpoint retrieves CAPWAP data for access points
	CapwapDataEndpoint = ApOperBasePath + "/capwap-data"
	// ApNameMacMapEndpoint retrieves AP name to MAC address mapping
	ApNameMacMapEndpoint = ApOperBasePath + "/ap-name-mac-map"

	// ApGlobalOperBasePath defines the base path for global access point operational endpoints
	ApGlobalOperBasePath = constants.YANGModelPrefix + "ap-global-oper:ap-global-oper-data"
	// ApGlobalOperEndpoint retrieves complete AP global operational data
	ApGlobalOperEndpoint = ApGlobalOperBasePath
	// ApHistoryEndpoint retrieves AP history data
	ApHistoryEndpoint = ApGlobalOperBasePath + "/ap-history"
	// EwlcApStatsEndpoint retrieves EWLC AP statistics
	EwlcApStatsEndpoint = ApGlobalOperBasePath + "/ewlc-ap-stats"
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

// Cfg returns complete access point configuration data.
func (s Service) Cfg(ctx context.Context) (*model.ApCfgResponse, error) {
	return core.Get[model.ApCfgResponse](ctx, s.c, ApCfgEndpoint)
}

// TagSourcePriorityConfigs returns tag source priority configurations.
func (s Service) TagSourcePriorityConfigs(ctx context.Context) (*model.TagSourcePriorityConfigs, error) {
	return core.Get[model.TagSourcePriorityConfigs](ctx, s.c, TagSourcePriorityConfigsEndpoint)
}

// ApTags returns access point tag configurations.
func (s Service) ApTags(ctx context.Context) (*model.ApCfgApTagsResponse, error) {
	return core.Get[model.ApCfgApTagsResponse](ctx, s.c, ApTagsEndpoint)
}

// Operational Methods

// Oper returns complete access point operational data.
func (s Service) Oper(ctx context.Context) (*model.ApOperResponse, error) {
	return core.Get[model.ApOperResponse](ctx, s.c, ApOperEndpoint)
}

// RadioNeighbor returns access point radio neighbor information.
func (s Service) RadioNeighbor(ctx context.Context) (*model.ApOperApRadioNeighborResponse, error) {
	return core.Get[model.ApOperApRadioNeighborResponse](ctx, s.c, ApRadioNeighborEndpoint)
}

// NameMacMap returns the mapping between AP names and MAC addresses.
func (s Service) NameMacMap(ctx context.Context) (*[]model.ApNameMacMap, error) {
	var resp struct {
		Data []model.ApNameMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
	}
	err := s.c.Do(ctx, http.MethodGet, ApNameMacMapEndpoint, &resp)
	return &resp.Data, err
}

// CapwapData returns CAPWAP protocol data.
func (s Service) CapwapData(ctx context.Context) (*[]model.CapwapData, error) {
	var resp struct {
		Data []model.CapwapData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}
	err := s.c.Do(ctx, http.MethodGet, CapwapDataEndpoint, &resp)
	return &resp.Data, err
}

// Global Operational Methods

// GlobalOper returns complete AP global operational data.
func (s Service) GlobalOper(ctx context.Context) (*model.ApGlobalOperResponse, error) {
	return core.Get[model.ApGlobalOperResponse](ctx, s.c, ApGlobalOperEndpoint)
}

// History returns AP history data.
func (s Service) History(ctx context.Context) (*model.ApGlobalOperApHistoryResponse, error) {
	return core.Get[model.ApGlobalOperApHistoryResponse](ctx, s.c, ApHistoryEndpoint)
}

// EwlcApStats returns EWLC AP statistics.
func (s Service) EwlcApStats(ctx context.Context) (*model.ApGlobalOperEwlcApStatsResponse, error) {
	return core.Get[model.ApGlobalOperEwlcApStatsResponse](ctx, s.c, EwlcApStatsEndpoint)
}
