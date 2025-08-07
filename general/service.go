package general

import (
	"context"
	"net/http"

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
	var out model.GeneralOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, GeneralOperEndpoint, &out)
}

// MgmtIntfData returns management interface operational data.
func (s Service) MgmtIntfData(ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	var out model.GeneralOperMgmtIntfDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, MgmtIntfDataEndpoint, &out)
}

// Configuration Data Methods

// Cfg returns general configuration data.
func (s Service) Cfg(ctx context.Context) (*model.GeneralCfgResponse, error) {
	var out model.GeneralCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, GeneralCfgEndpoint, &out)
}

// MewlcConfig returns MEWLC configuration data.
func (s Service) MewlcConfig(ctx context.Context) (*model.MewlcConfigResponse, error) {
	var out model.MewlcConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet, MewlcConfigEndpoint, &out)
}

// CacConfig returns CAC configuration data.
func (s Service) CacConfig(ctx context.Context) (*model.CacConfigResponse, error) {
	var out model.CacConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet, CacConfigEndpoint, &out)
}

// Mfp returns MFP (Management Frame Protection) configuration data.
func (s Service) Mfp(ctx context.Context) (*model.MfpResponse, error) {
	var out model.MfpResponse
	return &out, s.c.Do(ctx, http.MethodGet, MfpEndpoint, &out)
}

// FipsCfg returns FIPS configuration data.
func (s Service) FipsCfg(ctx context.Context) (*model.FipsCfgResponse, error) {
	var out model.FipsCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, FipsCfgEndpoint, &out)
}

// WsaApClientEvent returns WSA AP client event configuration data.
func (s Service) WsaApClientEvent(ctx context.Context) (*model.WsaApClientEventResponse, error) {
	var out model.WsaApClientEventResponse
	return &out, s.c.Do(ctx, http.MethodGet, WsaApClientEventEndpoint, &out)
}

// SimL3InterfaceCacheData returns SIM L3 interface cache data.
func (s Service) SimL3InterfaceCacheData(ctx context.Context) (*model.SimL3InterfaceCacheDataResponse, error) {
	var out model.SimL3InterfaceCacheDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, SimL3InterfaceCacheDataEndpoint, &out)
}

// WlcManagementData returns WLC management data.
func (s Service) WlcManagementData(ctx context.Context) (*model.WlcManagementDataResponse, error) {
	var out model.WlcManagementDataResponse
	return &out, s.c.Do(ctx, http.MethodGet, WlcManagementDataEndpoint, &out)
}

// Laginfo returns LAG (Link Aggregation) information.
func (s Service) Laginfo(ctx context.Context) (*model.LaginfoResponse, error) {
	var out model.LaginfoResponse
	return &out, s.c.Do(ctx, http.MethodGet, LaginfoEndpoint, &out)
}

// MulticastConfig returns multicast configuration data.
func (s Service) MulticastConfig(ctx context.Context) (*model.MulticastConfigResponse, error) {
	var out model.MulticastConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet, MulticastConfigEndpoint, &out)
}

// FeatureUsageCfg returns feature usage configuration data.
func (s Service) FeatureUsageCfg(ctx context.Context) (*model.FeatureUsageCfgResponse, error) {
	var out model.FeatureUsageCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, FeatureUsageCfgEndpoint, &out)
}

// ThresholdWarnCfg returns threshold warning configuration data.
func (s Service) ThresholdWarnCfg(ctx context.Context) (*model.ThresholdWarnCfgResponse, error) {
	var out model.ThresholdWarnCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, ThresholdWarnCfgEndpoint, &out)
}

// ApLocRangingCfg returns AP location ranging configuration data.
func (s Service) ApLocRangingCfg(ctx context.Context) (*model.ApLocRangingCfgResponse, error) {
	var out model.ApLocRangingCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, ApLocRangingCfgEndpoint, &out)
}

// GeolocationCfg returns geolocation configuration data.
func (s Service) GeolocationCfg(ctx context.Context) (*model.GeolocationCfgResponse, error) {
	var out model.GeolocationCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet, GeolocationCfgEndpoint, &out)
}
