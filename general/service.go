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

// Oper returns general operational data.
func (s Service) Oper(ctx context.Context) (*model.GeneralOperResponse, error) {
	return core.Get[model.GeneralOperResponse](ctx, s.c, GeneralOperEndpoint)
}

// MgmtIntfData returns management interface operational data.
func (s Service) MgmtIntfData(ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	return core.Get[model.GeneralOperMgmtIntfDataResponse](ctx, s.c, MgmtIntfDataEndpoint)
}

// Configuration Data Methods

// Cfg returns general configuration data.
func (s Service) Cfg(ctx context.Context) (*model.GeneralCfgResponse, error) {
	return core.Get[model.GeneralCfgResponse](ctx, s.c, GeneralCfgEndpoint)
}

// MewlcConfig returns MEWLC configuration data.
func (s Service) MewlcConfig(ctx context.Context) (*model.MewlcConfigResponse, error) {
	return core.Get[model.MewlcConfigResponse](ctx, s.c, MewlcConfigEndpoint)
}

// CacConfig returns CAC configuration data.
func (s Service) CacConfig(ctx context.Context) (*model.CacConfigResponse, error) {
	return core.Get[model.CacConfigResponse](ctx, s.c, CacConfigEndpoint)
}

// Mfp returns MFP (Management Frame Protection) configuration data.
func (s Service) Mfp(ctx context.Context) (*model.MfpResponse, error) {
	return core.Get[model.MfpResponse](ctx, s.c, MfpEndpoint)
}

// FipsCfg returns FIPS configuration data.
func (s Service) FipsCfg(ctx context.Context) (*model.FipsCfgResponse, error) {
	return core.Get[model.FipsCfgResponse](ctx, s.c, FipsCfgEndpoint)
}

// WsaApClientEvent returns WSA AP client event configuration data.
func (s Service) WsaApClientEvent(ctx context.Context) (*model.WsaApClientEventResponse, error) {
	return core.Get[model.WsaApClientEventResponse](ctx, s.c, WsaApClientEventEndpoint)
}

// SimL3InterfaceCacheData returns SIM L3 interface cache data.
func (s Service) SimL3InterfaceCacheData(ctx context.Context) (*model.SimL3InterfaceCacheDataResponse, error) {
	return core.Get[model.SimL3InterfaceCacheDataResponse](ctx, s.c, SimL3InterfaceCacheDataEndpoint)
}

// WlcManagementData returns WLC management data.
func (s Service) WlcManagementData(ctx context.Context) (*model.WlcManagementDataResponse, error) {
	return core.Get[model.WlcManagementDataResponse](ctx, s.c, WlcManagementDataEndpoint)
}

// Laginfo returns LAG (Link Aggregation) information.
func (s Service) Laginfo(ctx context.Context) (*model.LaginfoResponse, error) {
	return core.Get[model.LaginfoResponse](ctx, s.c, LaginfoEndpoint)
}

// MulticastConfig returns multicast configuration data.
func (s Service) MulticastConfig(ctx context.Context) (*model.MulticastConfigResponse, error) {
	return core.Get[model.MulticastConfigResponse](ctx, s.c, MulticastConfigEndpoint)
}

// FeatureUsageCfg returns feature usage configuration data.
func (s Service) FeatureUsageCfg(ctx context.Context) (*model.FeatureUsageCfgResponse, error) {
	return core.Get[model.FeatureUsageCfgResponse](ctx, s.c, FeatureUsageCfgEndpoint)
}

// ThresholdWarnCfg returns threshold warning configuration data.
func (s Service) ThresholdWarnCfg(ctx context.Context) (*model.ThresholdWarnCfgResponse, error) {
	return core.Get[model.ThresholdWarnCfgResponse](ctx, s.c, ThresholdWarnCfgEndpoint)
}

// ApLocRangingCfg returns AP location ranging configuration data.
func (s Service) ApLocRangingCfg(ctx context.Context) (*model.ApLocRangingCfgResponse, error) {
	return core.Get[model.ApLocRangingCfgResponse](ctx, s.c, ApLocRangingCfgEndpoint)
}

// GeolocationCfg returns geolocation configuration data.
func (s Service) GeolocationCfg(ctx context.Context) (*model.GeolocationCfgResponse, error) {
	return core.Get[model.GeolocationCfgResponse](ctx, s.c, GeolocationCfgEndpoint)
}
