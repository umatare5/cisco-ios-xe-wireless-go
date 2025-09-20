package general

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides general system information operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new General service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves general operational data.
func (s Service) GetOperational(ctx context.Context) (*GeneralOper, error) {
	return core.Get[GeneralOper](ctx, s.Client(), routes.GeneralOperPath)
}

// GetManagementInterfaceState retrieves management interface operational data.
func (s Service) GetManagementInterfaceState(ctx context.Context) (*GeneralOperMgmtIntfData, error) {
	return core.Get[GeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath)
}

// GetConfig retrieves complete general configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*GeneralCfg, error) {
	return core.Get[GeneralCfg](ctx, s.Client(), routes.GeneralCfgPath)
}

// Configuration data retrieval methods

// GetAPLocationRangingConfig retrieves AP location ranging configuration data.
func (s Service) GetAPLocationRangingConfig(ctx context.Context) (*ApLocRangingCfg, error) {
	return core.Get[ApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath)
}

// GetCACConfig retrieves CAC configuration data.
func (s Service) GetCACConfig(ctx context.Context) (*CacConfig, error) {
	return core.Get[CacConfig](ctx, s.Client(), routes.GeneralCacConfigPath)
}

// GetFeatureUsageConfig retrieves feature usage configuration data.
func (s Service) GetFeatureUsageConfig(ctx context.Context) (*FeatureUsageCfg, error) {
	return core.Get[FeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath)
}

// GetFIPSConfig retrieves FIPS configuration data.
func (s Service) GetFIPSConfig(ctx context.Context) (*FipsCfg, error) {
	return core.Get[FipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath)
}

// GetGeolocationConfig retrieves geolocation configuration data.
func (s Service) GetGeolocationConfig(ctx context.Context) (*GeolocationCfg, error) {
	return core.Get[GeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath)
}

// GetLAGInfo retrieves LAG (Link Aggregation) information.
func (s Service) GetLAGInfo(ctx context.Context) (*Laginfo, error) {
	return core.Get[Laginfo](ctx, s.Client(), routes.GeneralLaginfoPath)
}

// GetMEWLCConfig retrieves MEWLC configuration data.
func (s Service) GetMEWLCConfig(ctx context.Context) (*MewlcConfig, error) {
	return core.Get[MewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath)
}

// GetMFPConfig retrieves MFP (Management Frame Protection) configuration data.
func (s Service) GetMFPConfig(ctx context.Context) (*Mfp, error) {
	return core.Get[Mfp](ctx, s.Client(), routes.GeneralMfpPath)
}

// GetMulticastConfig retrieves multicast configuration data.
func (s Service) GetMulticastConfig(ctx context.Context) (*MulticastConfig, error) {
	return core.Get[MulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath)
}

// ListSIML3InterfaceCache returns SIM L3 interface cache data.
func (s Service) ListSIML3InterfaceCache(ctx context.Context) (*SimL3InterfaceCacheData, error) {
	return core.Get[SimL3InterfaceCacheData](ctx, s.Client(), routes.GeneralSimL3InterfaceCacheDataPath)
}

// GetThresholdWarningConfig retrieves threshold warning configuration data.
func (s Service) GetThresholdWarningConfig(ctx context.Context) (*ThresholdWarnCfg, error) {
	return core.Get[ThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath)
}

// GetWLCManagementInfo retrieves WLC management data.
func (s Service) GetWLCManagementInfo(ctx context.Context) (*WlcManagementData, error) {
	return core.Get[WlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath)
}

// GetWSAAPClientEventConfig retrieves WSA AP client event configuration data.
func (s Service) GetWSAAPClientEventConfig(ctx context.Context) (*WsaApClientEvent, error) {
	return core.Get[WsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath)
}
