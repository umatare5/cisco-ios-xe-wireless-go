package general

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/general"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for general service
func (s Service) operOps() *core.OperationalOperations[model.GeneralOper] {
	return core.NewOperationalOperations[model.GeneralOper](s.Client(), routes.GeneralOperBasePath)
}

// GetOper retrieves general operational data.
func (s Service) GetOper(ctx context.Context) (*model.GeneralOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperMgmtIntfData retrieves management interface operational data.
func (s Service) GetOperMgmtIntfData(ctx context.Context) (*model.GeneralOperMgmtIntfData, error) {
	return core.Get[model.GeneralOperMgmtIntfData](ctx, s.Client(), routes.MgmtIntfDataEndpoint)
}

// Configuration data retrieval methods

// GetOperApLocRangingCfg retrieves AP location ranging configuration data.
func (s Service) GetOperApLocRangingCfg(ctx context.Context) (*model.ApLocRangingCfg, error) {
	return core.Get[model.ApLocRangingCfg](ctx, s.Client(), routes.ApLocRangingCfgEndpoint)
}

// GetOperCacConfig retrieves CAC configuration data.
func (s Service) GetOperCacConfig(ctx context.Context) (*model.CacConfig, error) {
	return core.Get[model.CacConfig](ctx, s.Client(), routes.CacConfigEndpoint)
}

// GetOperFeatureUsageCfg retrieves feature usage configuration data.
func (s Service) GetOperFeatureUsageCfg(ctx context.Context) (*model.FeatureUsageCfg, error) {
	return core.Get[model.FeatureUsageCfg](ctx, s.Client(), routes.FeatureUsageCfgEndpoint)
}

// GetOperFipsCfg retrieves FIPS configuration data.
func (s Service) GetOperFipsCfg(ctx context.Context) (*model.FipsCfg, error) {
	return core.Get[model.FipsCfg](ctx, s.Client(), routes.FipsCfgEndpoint)
}

// GetOperGeolocationCfg retrieves geolocation configuration data.
func (s Service) GetOperGeolocationCfg(ctx context.Context) (*model.GeolocationCfg, error) {
	return core.Get[model.GeolocationCfg](ctx, s.Client(), routes.GeolocationCfgEndpoint)
}

// GetOperLaginfo retrieves LAG (Link Aggregation) information.
func (s Service) GetOperLaginfo(ctx context.Context) (*model.Laginfo, error) {
	return core.Get[model.Laginfo](ctx, s.Client(), routes.LaginfoEndpoint)
}

// GetOperMewlcConfig retrieves MEWLC configuration data.
func (s Service) GetOperMewlcConfig(ctx context.Context) (*model.MewlcConfig, error) {
	return core.Get[model.MewlcConfig](ctx, s.Client(), routes.MewlcConfigEndpoint)
}

// GetOperMfp retrieves MFP (Management Frame Protection) configuration data.
func (s Service) GetOperMfp(ctx context.Context) (*model.Mfp, error) {
	return core.Get[model.Mfp](ctx, s.Client(), routes.MfpEndpoint)
}

// GetOperMulticastConfig retrieves multicast configuration data.
func (s Service) GetOperMulticastConfig(ctx context.Context) (*model.MulticastConfig, error) {
	return core.Get[model.MulticastConfig](ctx, s.Client(), routes.MulticastConfigEndpoint)
}

// GetOperSimL3InterfaceCacheData returns SIM L3 interface cache data.
func (s Service) GetOperSimL3InterfaceCacheData(ctx context.Context) (*model.SimL3InterfaceCacheData, error) {
	return core.Get[model.SimL3InterfaceCacheData](ctx, s.Client(), routes.SimL3InterfaceCacheDataEndpoint)
}

// GetOperThresholdWarnCfg retrieves threshold warning configuration data.
func (s Service) GetOperThresholdWarnCfg(ctx context.Context) (*model.ThresholdWarnCfg, error) {
	return core.Get[model.ThresholdWarnCfg](ctx, s.Client(), routes.ThresholdWarnCfgEndpoint)
}

// GetOperWlcManagementData retrieves WLC management data.
func (s Service) GetOperWlcManagementData(ctx context.Context) (*model.WlcManagementData, error) {
	return core.Get[model.WlcManagementData](ctx, s.Client(), routes.WlcManagementDataEndpoint)
}

// GetOperWsaApClientEvent retrieves WSA AP client event configuration data.
func (s Service) GetOperWsaApClientEvent(ctx context.Context) (*model.WsaApClientEvent, error) {
	return core.Get[model.WsaApClientEvent](ctx, s.Client(), routes.WsaApClientEventEndpoint)
}
