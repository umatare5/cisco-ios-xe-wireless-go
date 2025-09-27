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
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessGeneralOper, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOper](ctx, s.Client(), routes.GeneralOperPath)
}

// GetManagementInterfaceState retrieves management interface operational data.
func (s Service) GetManagementInterfaceState(ctx context.Context) (*CiscoIOSXEWirelessGeneralOperMgmtIntfData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath)
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

// ListCfgMewlcConfig retrieves MEWLC configuration data wrapper.
func (s Service) ListCfgMewlcConfig(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgMewlcConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath)
}

// ListCfgCacConfig retrieves CAC configuration data wrapper.
func (s Service) ListCfgCacConfig(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgCacConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgCacConfig](ctx, s.Client(), routes.GeneralCacConfigPath)
}

// ListCfgMfp retrieves MFP configuration data wrapper.
func (s Service) ListCfgMfp(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgMfp, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMfp](ctx, s.Client(), routes.GeneralMfpPath)
}

// ListCfgFipsCfg retrieves FIPS configuration data wrapper.
func (s Service) ListCfgFipsCfg(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgFipsCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgFipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath)
}

// ListCfgWsaApClientEvent retrieves WSA AP client event configuration data wrapper.
func (s Service) ListCfgWsaApClientEvent(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgWsaApClientEvent, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgWsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath)
}

// ListCfgSimL3InterfaceCacheData retrieves SIM L3 interface cache data wrapper.
func (s Service) ListCfgSimL3InterfaceCacheData(
	ctx context.Context,
) (*CiscoIOSXEWirelessGeneralCfgSimL3InterfaceCacheData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgSimL3InterfaceCacheData](
		ctx,
		s.Client(),
		routes.GeneralSimL3InterfaceCacheDataPath,
	)
}

// ListCfgWlcManagementData retrieves WLC management data wrapper.
func (s Service) ListCfgWlcManagementData(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgWlcManagementData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgWlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath)
}

// ListCfgLaginfo retrieves LAG information wrapper.
func (s Service) ListCfgLaginfo(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgLaginfo, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgLaginfo](ctx, s.Client(), routes.GeneralLaginfoPath)
}

// ListCfgMulticastConfig retrieves multicast configuration data wrapper.
func (s Service) ListCfgMulticastConfig(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgMulticastConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath)
}

// ListCfgFeatureUsageCfg retrieves feature usage configuration data wrapper.
func (s Service) ListCfgFeatureUsageCfg(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgFeatureUsageCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgFeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath)
}

// ListCfgThresholdWarnCfg retrieves threshold warning configuration data wrapper.
func (s Service) ListCfgThresholdWarnCfg(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgThresholdWarnCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath)
}

// ListCfgApLocRangingCfg retrieves AP location ranging configuration data wrapper.
func (s Service) ListCfgApLocRangingCfg(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgApLocRangingCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath)
}

// ListCfgGeolocationCfg retrieves geolocation configuration data wrapper.
func (s Service) ListCfgGeolocationCfg(ctx context.Context) (*CiscoIOSXEWirelessGeneralCfgGeolocationCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgGeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath)
}

// ListOperMgmtIntfData retrieves management interface operational data wrapper.
func (s Service) ListOperMgmtIntfData(ctx context.Context) (*CiscoIOSXEWirelessGeneralOperMgmtIntfData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath)
}

// ListMewlcConfig retrieves MEWLC configuration data.
func (s Service) ListMewlcConfig(ctx context.Context) (*MewlcConfig, error) {
	return core.Get[MewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath)
}

// ListCacConfig retrieves CAC configuration data.
func (s Service) ListCacConfig(ctx context.Context) (*CacConfig, error) {
	return core.Get[CacConfig](ctx, s.Client(), routes.GeneralCacConfigPath)
}

// ListMfp retrieves MFP configuration data.
func (s Service) ListMfp(ctx context.Context) (*Mfp, error) {
	return core.Get[Mfp](ctx, s.Client(), routes.GeneralMfpPath)
}

// ListFipsCfg retrieves FIPS configuration data.
func (s Service) ListFipsCfg(ctx context.Context) (*FipsCfg, error) {
	return core.Get[FipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath)
}

// ListWsaApClientEvent retrieves WSA AP client event configuration data.
func (s Service) ListWsaApClientEvent(ctx context.Context) (*WsaApClientEvent, error) {
	return core.Get[WsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath)
}

// ListWlcManagementData retrieves WLC management data.
func (s Service) ListWlcManagementData(ctx context.Context) (*WlcManagementData, error) {
	return core.Get[WlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath)
}

// ListLaginfo retrieves LAG information.
func (s Service) ListLaginfo(ctx context.Context) (*Laginfo, error) {
	return core.Get[Laginfo](ctx, s.Client(), routes.GeneralLaginfoPath)
}

// ListMulticastConfig retrieves multicast configuration data.
func (s Service) ListMulticastConfig(ctx context.Context) (*MulticastConfig, error) {
	return core.Get[MulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath)
}

// ListFeatureUsageCfg retrieves feature usage configuration data.
func (s Service) ListFeatureUsageCfg(ctx context.Context) (*FeatureUsageCfg, error) {
	return core.Get[FeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath)
}

// ListThresholdWarnCfg retrieves threshold warning configuration data.
func (s Service) ListThresholdWarnCfg(ctx context.Context) (*ThresholdWarnCfg, error) {
	return core.Get[ThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath)
}

// ListApLocRangingCfg retrieves AP location ranging configuration data.
func (s Service) ListApLocRangingCfg(ctx context.Context) (*ApLocRangingCfg, error) {
	return core.Get[ApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath)
}

// ListGeolocationCfg retrieves geolocation configuration data.
func (s Service) ListGeolocationCfg(ctx context.Context) (*GeolocationCfg, error) {
	return core.Get[GeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath)
}
