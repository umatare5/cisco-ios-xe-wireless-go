package general

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// GeneralOperBasePath defines the base path for general operational data endpoints
	GeneralOperBasePath = constants.YANGModelPrefix + "general-oper:general-oper-data"
	// GeneralOperEndpoint retrieves general operational data
	GeneralOperEndpoint = GeneralOperBasePath
	// MgmtIntfDataEndpoint retrieves management interface operational data
	MgmtIntfDataEndpoint = GeneralOperBasePath + "/mgmt-intf-data"

	// GeneralCfgBasePath defines the base path for general configuration data endpoints
	GeneralCfgBasePath = constants.YANGModelPrefix + "general-cfg:general-cfg-data"
	// GeneralCfgEndpoint retrieves general configuration data
	GeneralCfgEndpoint = GeneralCfgBasePath
	// MewlcConfigEndpoint retrieves MEWLC configuration data
	MewlcConfigEndpoint = GeneralCfgBasePath + "/mewlc-config"
	// CacConfigEndpoint retrieves CAC configuration data
	CacConfigEndpoint = GeneralCfgBasePath + "/cac-config"
	// MfpEndpoint retrieves MFP (Management Frame Protection) configuration data
	MfpEndpoint = GeneralCfgBasePath + "/mfp"
	// FipsCfgEndpoint retrieves FIPS configuration data
	FipsCfgEndpoint = GeneralCfgBasePath + "/fips-cfg"
	// WsaApClientEventEndpoint retrieves WSA AP client event configuration data
	WsaApClientEventEndpoint = GeneralCfgBasePath + "/wsa-ap-client-event"
	// SimL3InterfaceCacheDataEndpoint retrieves SIM L3 interface cache data
	SimL3InterfaceCacheDataEndpoint = GeneralCfgBasePath + "/sim-l3-interface-cache-data"
	// WlcManagementDataEndpoint retrieves WLC management data
	WlcManagementDataEndpoint = GeneralCfgBasePath + "/wlc-management-data"
	// LaginfoEndpoint retrieves LAG (Link Aggregation) information
	LaginfoEndpoint = GeneralCfgBasePath + "/laginfo"
	// MulticastConfigEndpoint retrieves multicast configuration data
	MulticastConfigEndpoint = GeneralCfgBasePath + "/multicast-config"
	// FeatureUsageCfgEndpoint retrieves feature usage configuration data
	FeatureUsageCfgEndpoint = GeneralCfgBasePath + "/feature-usage-cfg"
	// ThresholdWarnCfgEndpoint retrieves threshold warning configuration data
	ThresholdWarnCfgEndpoint = GeneralCfgBasePath + "/threshold-warn-cfg"
	// ApLocRangingCfgEndpoint retrieves AP location ranging configuration data
	ApLocRangingCfgEndpoint = GeneralCfgBasePath + "/ap-loc-ranging-cfg"
	// GeolocationCfgEndpoint retrieves geolocation configuration data
	GeolocationCfgEndpoint = GeneralCfgBasePath + "/geolocation-cfg"
)

// Service provides General operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Operational Data Methods

// GetOper returns general operational data.
func (s Service) GetOper(ctx context.Context) (*model.GeneralOperResponse, error) {
	return core.Get[model.GeneralOperResponse](ctx, s.c, GeneralOperEndpoint)
}

// GetMgmtIntfData returns management interface operational data.
func (s Service) GetMgmtIntfData(ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	return core.Get[model.GeneralOperMgmtIntfDataResponse](ctx, s.c, MgmtIntfDataEndpoint)
}

// Configuration Data Methods

// GetCfg returns general configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.GeneralCfgResponse, error) {
	return core.Get[model.GeneralCfgResponse](ctx, s.c, GeneralCfgEndpoint)
}

// GetMewlcConfig returns MEWLC configuration data.
func (s Service) GetMewlcConfig(ctx context.Context) (*model.MewlcConfigResponse, error) {
	return core.Get[model.MewlcConfigResponse](ctx, s.c, MewlcConfigEndpoint)
}

// GetCacConfig returns CAC configuration data.
func (s Service) GetCacConfig(ctx context.Context) (*model.CacConfigResponse, error) {
	return core.Get[model.CacConfigResponse](ctx, s.c, CacConfigEndpoint)
}

// GetMfp returns MFP (Management Frame Protection) configuration data.
func (s Service) GetMfp(ctx context.Context) (*model.MfpResponse, error) {
	return core.Get[model.MfpResponse](ctx, s.c, MfpEndpoint)
}

// GetFipsCfg returns FIPS configuration data.
func (s Service) GetFipsCfg(ctx context.Context) (*model.FipsCfgResponse, error) {
	return core.Get[model.FipsCfgResponse](ctx, s.c, FipsCfgEndpoint)
}

// GetWsaApClientEvent returns WSA AP client event configuration data.
func (s Service) GetWsaApClientEvent(ctx context.Context) (*model.WsaApClientEventResponse, error) {
	return core.Get[model.WsaApClientEventResponse](ctx, s.c, WsaApClientEventEndpoint)
}

// GetSimL3InterfaceCacheData returns SIM L3 interface cache data.
func (s Service) GetSimL3InterfaceCacheData(ctx context.Context) (*model.SimL3InterfaceCacheDataResponse, error) {
	return core.Get[model.SimL3InterfaceCacheDataResponse](ctx, s.c, SimL3InterfaceCacheDataEndpoint)
}

// GetWlcManagementData returns WLC management data.
func (s Service) GetWlcManagementData(ctx context.Context) (*model.WlcManagementDataResponse, error) {
	return core.Get[model.WlcManagementDataResponse](ctx, s.c, WlcManagementDataEndpoint)
}

// GetLaginfo returns LAG (Link Aggregation) information.
func (s Service) GetLaginfo(ctx context.Context) (*model.LaginfoResponse, error) {
	return core.Get[model.LaginfoResponse](ctx, s.c, LaginfoEndpoint)
}

// GetMulticastConfig returns multicast configuration data.
func (s Service) GetMulticastConfig(ctx context.Context) (*model.MulticastConfigResponse, error) {
	return core.Get[model.MulticastConfigResponse](ctx, s.c, MulticastConfigEndpoint)
}

// GetFeatureUsageCfg returns feature usage configuration data.
func (s Service) GetFeatureUsageCfg(ctx context.Context) (*model.FeatureUsageCfgResponse, error) {
	return core.Get[model.FeatureUsageCfgResponse](ctx, s.c, FeatureUsageCfgEndpoint)
}

// GetThresholdWarnCfg returns threshold warning configuration data.
func (s Service) GetThresholdWarnCfg(ctx context.Context) (*model.ThresholdWarnCfgResponse, error) {
	return core.Get[model.ThresholdWarnCfgResponse](ctx, s.c, ThresholdWarnCfgEndpoint)
}

// GetApLocRangingCfg returns AP location ranging configuration data.
func (s Service) GetApLocRangingCfg(ctx context.Context) (*model.ApLocRangingCfgResponse, error) {
	return core.Get[model.ApLocRangingCfgResponse](ctx, s.c, ApLocRangingCfgEndpoint)
}

// GetGeolocationCfg returns geolocation configuration data.
func (s Service) GetGeolocationCfg(ctx context.Context) (*model.GeolocationCfgResponse, error) {
	return core.Get[model.GeolocationCfgResponse](ctx, s.c, GeolocationCfgEndpoint)
}
