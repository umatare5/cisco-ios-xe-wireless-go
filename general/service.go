package general

import (
	"context"
	"net/http"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides General domain operations for wireless controller configuration and operational data.
type Service struct {
	c *wnc.Client
}

// NewService creates a new General service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Operational Data Methods

// Oper returns general operational data.
func (s Service) Oper(ctx context.Context) (*model.GeneralOperResponse, error) {
	var out model.GeneralOperResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data", &out)
}

// MgmtIntfData returns management interface operational data.
func (s Service) MgmtIntfData(ctx context.Context) (*model.GeneralOperMgmtIntfDataResponse, error) {
	var out model.GeneralOperMgmtIntfDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data/mgmt-intf-data", &out)
}

// Configuration Data Methods

// Cfg returns general configuration data.
func (s Service) Cfg(ctx context.Context) (*model.GeneralCfgResponse, error) {
	var out model.GeneralCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data", &out)
}

// MewlcConfig returns MEWLC configuration data.
func (s Service) MewlcConfig(ctx context.Context) (*model.MewlcConfigResponse, error) {
	var out model.MewlcConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config", &out)
}

// CacConfig returns CAC configuration data.
func (s Service) CacConfig(ctx context.Context) (*model.CacConfigResponse, error) {
	var out model.CacConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config", &out)
}

// Mfp returns MFP (Management Frame Protection) configuration data.
func (s Service) Mfp(ctx context.Context) (*model.MfpResponse, error) {
	var out model.MfpResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp", &out)
}

// FipsCfg returns FIPS configuration data.
func (s Service) FipsCfg(ctx context.Context) (*model.FipsCfgResponse, error) {
	var out model.FipsCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg", &out)
}

// WsaApClientEvent returns WSA AP client event configuration data.
func (s Service) WsaApClientEvent(ctx context.Context) (*model.WsaApClientEventResponse, error) {
	var out model.WsaApClientEventResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event", &out)
}

// SimL3InterfaceCacheData returns SIM L3 interface cache data.
func (s Service) SimL3InterfaceCacheData(ctx context.Context) (*model.SimL3InterfaceCacheDataResponse, error) {
	var out model.SimL3InterfaceCacheDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data", &out)
}

// WlcManagementData returns WLC management data.
func (s Service) WlcManagementData(ctx context.Context) (*model.WlcManagementDataResponse, error) {
	var out model.WlcManagementDataResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data", &out)
}

// Laginfo returns LAG (Link Aggregation) information.
func (s Service) Laginfo(ctx context.Context) (*model.LaginfoResponse, error) {
	var out model.LaginfoResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo", &out)
}

// MulticastConfig returns multicast configuration data.
func (s Service) MulticastConfig(ctx context.Context) (*model.MulticastConfigResponse, error) {
	var out model.MulticastConfigResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config", &out)
}

// FeatureUsageCfg returns feature usage configuration data.
func (s Service) FeatureUsageCfg(ctx context.Context) (*model.FeatureUsageCfgResponse, error) {
	var out model.FeatureUsageCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg", &out)
}

// ThresholdWarnCfg returns threshold warning configuration data.
func (s Service) ThresholdWarnCfg(ctx context.Context) (*model.ThresholdWarnCfgResponse, error) {
	var out model.ThresholdWarnCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg", &out)
}

// ApLocRangingCfg returns AP location ranging configuration data.
func (s Service) ApLocRangingCfg(ctx context.Context) (*model.ApLocRangingCfgResponse, error) {
	var out model.ApLocRangingCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg", &out)
}

// GeolocationCfg returns geolocation configuration data.
func (s Service) GeolocationCfg(ctx context.Context) (*model.GeolocationCfgResponse, error) {
	var out model.GeolocationCfgResponse
	return &out, s.c.Do(ctx, http.MethodGet,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg", &out)
}
