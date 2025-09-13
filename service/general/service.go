package general

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/general"
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
func (s Service) GetOperational(ctx context.Context) (*model.GeneralOper, error) {
	return core.Get[model.GeneralOper](ctx, s.Client(), routes.GeneralOperPath)
}

// GetManagementInterfaceState retrieves management interface operational data.
func (s Service) GetManagementInterfaceState(ctx context.Context) (*model.GeneralOperMgmtIntfData, error) {
	return core.Get[model.GeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath)
}

// GetConfig retrieves complete general configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*model.GeneralCfg, error) {
	return core.Get[model.GeneralCfg](ctx, s.Client(), routes.GeneralCfgPath)
}

// Configuration data retrieval methods

// GetAPLocationRangingConfig retrieves AP location ranging configuration data.
func (s Service) GetAPLocationRangingConfig(ctx context.Context) (*model.ApLocRangingCfg, error) {
	return core.Get[model.ApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath)
}

// GetCACConfig retrieves CAC configuration data.
func (s Service) GetCACConfig(ctx context.Context) (*model.CacConfig, error) {
	return core.Get[model.CacConfig](ctx, s.Client(), routes.GeneralCacConfigPath)
}

// GetFeatureUsageConfig retrieves feature usage configuration data.
func (s Service) GetFeatureUsageConfig(ctx context.Context) (*model.FeatureUsageCfg, error) {
	return core.Get[model.FeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath)
}

// GetFIPSConfig retrieves FIPS configuration data.
func (s Service) GetFIPSConfig(ctx context.Context) (*model.FipsCfg, error) {
	return core.Get[model.FipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath)
}

// GetGeolocationConfig retrieves geolocation configuration data.
func (s Service) GetGeolocationConfig(ctx context.Context) (*model.GeolocationCfg, error) {
	return core.Get[model.GeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath)
}

// GetLAGInfo retrieves LAG (Link Aggregation) information.
func (s Service) GetLAGInfo(ctx context.Context) (*model.Laginfo, error) {
	return core.Get[model.Laginfo](ctx, s.Client(), routes.GeneralLaginfoPath)
}

// GetMEWLCConfig retrieves MEWLC configuration data.
func (s Service) GetMEWLCConfig(ctx context.Context) (*model.MewlcConfig, error) {
	return core.Get[model.MewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath)
}

// GetMFPConfig retrieves MFP (Management Frame Protection) configuration data.
func (s Service) GetMFPConfig(ctx context.Context) (*model.Mfp, error) {
	return core.Get[model.Mfp](ctx, s.Client(), routes.GeneralMfpPath)
}

// GetMulticastConfig retrieves multicast configuration data.
func (s Service) GetMulticastConfig(ctx context.Context) (*model.MulticastConfig, error) {
	return core.Get[model.MulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath)
}

// ListSIML3InterfaceCache returns SIM L3 interface cache data.
func (s Service) ListSIML3InterfaceCache(ctx context.Context) (*model.SimL3InterfaceCacheData, error) {
	return core.Get[model.SimL3InterfaceCacheData](ctx, s.Client(), routes.GeneralSimL3InterfaceCacheDataPath)
}

// GetThresholdWarningConfig retrieves threshold warning configuration data.
func (s Service) GetThresholdWarningConfig(ctx context.Context) (*model.ThresholdWarnCfg, error) {
	return core.Get[model.ThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath)
}

// GetWLCManagementInfo retrieves WLC management data.
func (s Service) GetWLCManagementInfo(ctx context.Context) (*model.WlcManagementData, error) {
	return core.Get[model.WlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath)
}

// GetWSAAPClientEventConfig retrieves WSA AP client event configuration data.
func (s Service) GetWSAAPClientEventConfig(ctx context.Context) (*model.WsaApClientEvent, error) {
	return core.Get[model.WsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath)
}
