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
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessGeneralOper, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOper](ctx, s.Client(), routes.GeneralOperPath, opts...)
}

// GetManagementInterfaceState retrieves management interface operational data.
func (s Service) GetManagementInterfaceState(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralOperMgmtIntfData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath, opts...)
}

// GetConfig retrieves complete general configuration data from the controller.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*GeneralCfg, error) {
	return core.Get[GeneralCfg](ctx, s.Client(), routes.GeneralCfgPath, opts...)
}

// Configuration data retrieval methods

// GetAPLocationRangingConfig retrieves AP location ranging configuration data.
func (s Service) GetAPLocationRangingConfig(ctx context.Context, opts ...core.GetOption) (*ApLocRangingCfg, error) {
	return core.Get[ApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath, opts...)
}

// GetCACConfig retrieves CAC configuration data.
func (s Service) GetCACConfig(ctx context.Context, opts ...core.GetOption) (*CacConfig, error) {
	return core.Get[CacConfig](ctx, s.Client(), routes.GeneralCacConfigPath, opts...)
}

// GetFeatureUsageConfig retrieves feature usage configuration data.
func (s Service) GetFeatureUsageConfig(ctx context.Context, opts ...core.GetOption) (*FeatureUsageCfg, error) {
	return core.Get[FeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath, opts...)
}

// GetFIPSConfig retrieves FIPS configuration data.
func (s Service) GetFIPSConfig(ctx context.Context, opts ...core.GetOption) (*FipsCfg, error) {
	return core.Get[FipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath, opts...)
}

// GetGeolocationConfig retrieves geolocation configuration data.
func (s Service) GetGeolocationConfig(ctx context.Context, opts ...core.GetOption) (*GeolocationCfg, error) {
	return core.Get[GeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath, opts...)
}

// GetLAGInfo retrieves LAG (Link Aggregation) information.
func (s Service) GetLAGInfo(ctx context.Context, opts ...core.GetOption) (*Laginfo, error) {
	return core.Get[Laginfo](ctx, s.Client(), routes.GeneralLaginfoPath, opts...)
}

// GetMEWLCConfig retrieves MEWLC configuration data.
func (s Service) GetMEWLCConfig(ctx context.Context, opts ...core.GetOption) (*MewlcConfig, error) {
	return core.Get[MewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath, opts...)
}

// GetMFPConfig retrieves MFP (Management Frame Protection) configuration data.
func (s Service) GetMFPConfig(ctx context.Context, opts ...core.GetOption) (*Mfp, error) {
	return core.Get[Mfp](ctx, s.Client(), routes.GeneralMfpPath, opts...)
}

// GetMulticastConfig retrieves multicast configuration data.
func (s Service) GetMulticastConfig(ctx context.Context, opts ...core.GetOption) (*MulticastConfig, error) {
	return core.Get[MulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath, opts...)
}

// ListSIML3InterfaceCache returns SIM L3 interface cache data.
func (s Service) ListSIML3InterfaceCache(
	ctx context.Context,
	opts ...core.GetOption,
) (*SimL3InterfaceCacheData, error) {
	return core.Get[SimL3InterfaceCacheData](ctx, s.Client(), routes.GeneralSimL3InterfaceCacheDataPath, opts...)
}

// GetThresholdWarningConfig retrieves threshold warning configuration data.
func (s Service) GetThresholdWarningConfig(ctx context.Context, opts ...core.GetOption) (*ThresholdWarnCfg, error) {
	return core.Get[ThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath, opts...)
}

// GetWLCManagementInfo retrieves WLC management data.
func (s Service) GetWLCManagementInfo(ctx context.Context, opts ...core.GetOption) (*WlcManagementData, error) {
	return core.Get[WlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath, opts...)
}

// GetWSAAPClientEventConfig retrieves WSA AP client event configuration data.
func (s Service) GetWSAAPClientEventConfig(ctx context.Context, opts ...core.GetOption) (*WsaApClientEvent, error) {
	return core.Get[WsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath, opts...)
}

// ListCfgMewlcConfig retrieves MEWLC configuration data wrapper.
func (s Service) ListCfgMewlcConfig(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgMewlcConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath, opts...)
}

// ListCfgCacConfig retrieves CAC configuration data wrapper.
func (s Service) ListCfgCacConfig(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgCacConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgCacConfig](ctx, s.Client(), routes.GeneralCacConfigPath, opts...)
}

// ListCfgMfp retrieves MFP configuration data wrapper.
func (s Service) ListCfgMfp(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessGeneralCfgMfp, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMfp](ctx, s.Client(), routes.GeneralMfpPath, opts...)
}

// ListCfgFipsCfg retrieves FIPS configuration data wrapper.
func (s Service) ListCfgFipsCfg(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgFipsCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgFipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath, opts...)
}

// ListCfgWsaApClientEvent retrieves WSA AP client event configuration data wrapper.
func (s Service) ListCfgWsaApClientEvent(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgWsaApClientEvent, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgWsaApClientEvent](
		ctx,
		s.Client(),
		routes.GeneralWsaApClientEventPath,
		opts...)
}

// ListCfgSimL3InterfaceCacheData retrieves SIM L3 interface cache data wrapper.
func (s Service) ListCfgSimL3InterfaceCacheData(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgSimL3InterfaceCacheData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgSimL3InterfaceCacheData](
		ctx,
		s.Client(),
		routes.GeneralSimL3InterfaceCacheDataPath, opts...,
	)
}

// ListCfgWlcManagementData retrieves WLC management data wrapper.
func (s Service) ListCfgWlcManagementData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgWlcManagementData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgWlcManagementData](
		ctx,
		s.Client(),
		routes.GeneralWlcManagementDataPath,
		opts...)
}

// ListCfgLaginfo retrieves LAG information wrapper.
func (s Service) ListCfgLaginfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgLaginfo, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgLaginfo](ctx, s.Client(), routes.GeneralLaginfoPath, opts...)
}

// ListCfgMulticastConfig retrieves multicast configuration data wrapper.
func (s Service) ListCfgMulticastConfig(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgMulticastConfig, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgMulticastConfig](
		ctx,
		s.Client(),
		routes.GeneralMulticastConfigPath,
		opts...)
}

// ListCfgFeatureUsageCfg retrieves feature usage configuration data wrapper.
func (s Service) ListCfgFeatureUsageCfg(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgFeatureUsageCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgFeatureUsageCfg](
		ctx,
		s.Client(),
		routes.GeneralFeatureUsageCfgPath,
		opts...)
}

// ListCfgThresholdWarnCfg retrieves threshold warning configuration data wrapper.
func (s Service) ListCfgThresholdWarnCfg(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgThresholdWarnCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgThresholdWarnCfg](
		ctx,
		s.Client(),
		routes.GeneralThresholdWarnCfgPath,
		opts...)
}

// ListCfgApLocRangingCfg retrieves AP location ranging configuration data wrapper.
func (s Service) ListCfgApLocRangingCfg(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgApLocRangingCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgApLocRangingCfg](
		ctx,
		s.Client(),
		routes.GeneralApLocRangingCfgPath,
		opts...)
}

// ListCfgGeolocationCfg retrieves geolocation configuration data wrapper.
func (s Service) ListCfgGeolocationCfg(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralCfgGeolocationCfg, error) {
	return core.Get[CiscoIOSXEWirelessGeneralCfgGeolocationCfg](
		ctx,
		s.Client(),
		routes.GeneralGeolocationCfgPath,
		opts...)
}

// ListOperMgmtIntfData retrieves management interface operational data wrapper.
func (s Service) ListOperMgmtIntfData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessGeneralOperMgmtIntfData, error) {
	return core.Get[CiscoIOSXEWirelessGeneralOperMgmtIntfData](ctx, s.Client(), routes.GeneralMgmtIntfDataPath, opts...)
}

// ListMewlcConfig retrieves MEWLC configuration data.
func (s Service) ListMewlcConfig(ctx context.Context, opts ...core.GetOption) (*MewlcConfig, error) {
	return core.Get[MewlcConfig](ctx, s.Client(), routes.GeneralMewlcConfigPath, opts...)
}

// ListCacConfig retrieves CAC configuration data.
func (s Service) ListCacConfig(ctx context.Context, opts ...core.GetOption) (*CacConfig, error) {
	return core.Get[CacConfig](ctx, s.Client(), routes.GeneralCacConfigPath, opts...)
}

// ListMfp retrieves MFP configuration data.
func (s Service) ListMfp(ctx context.Context, opts ...core.GetOption) (*Mfp, error) {
	return core.Get[Mfp](ctx, s.Client(), routes.GeneralMfpPath, opts...)
}

// ListFipsCfg retrieves FIPS configuration data.
func (s Service) ListFipsCfg(ctx context.Context, opts ...core.GetOption) (*FipsCfg, error) {
	return core.Get[FipsCfg](ctx, s.Client(), routes.GeneralFipsCfgPath, opts...)
}

// ListWsaApClientEvent retrieves WSA AP client event configuration data.
func (s Service) ListWsaApClientEvent(ctx context.Context, opts ...core.GetOption) (*WsaApClientEvent, error) {
	return core.Get[WsaApClientEvent](ctx, s.Client(), routes.GeneralWsaApClientEventPath, opts...)
}

// ListWlcManagementData retrieves WLC management data.
func (s Service) ListWlcManagementData(ctx context.Context, opts ...core.GetOption) (*WlcManagementData, error) {
	return core.Get[WlcManagementData](ctx, s.Client(), routes.GeneralWlcManagementDataPath, opts...)
}

// ListLaginfo retrieves LAG information.
func (s Service) ListLaginfo(ctx context.Context, opts ...core.GetOption) (*Laginfo, error) {
	return core.Get[Laginfo](ctx, s.Client(), routes.GeneralLaginfoPath, opts...)
}

// ListMulticastConfig retrieves multicast configuration data.
func (s Service) ListMulticastConfig(ctx context.Context, opts ...core.GetOption) (*MulticastConfig, error) {
	return core.Get[MulticastConfig](ctx, s.Client(), routes.GeneralMulticastConfigPath, opts...)
}

// ListFeatureUsageCfg retrieves feature usage configuration data.
func (s Service) ListFeatureUsageCfg(ctx context.Context, opts ...core.GetOption) (*FeatureUsageCfg, error) {
	return core.Get[FeatureUsageCfg](ctx, s.Client(), routes.GeneralFeatureUsageCfgPath, opts...)
}

// ListThresholdWarnCfg retrieves threshold warning configuration data.
func (s Service) ListThresholdWarnCfg(ctx context.Context, opts ...core.GetOption) (*ThresholdWarnCfg, error) {
	return core.Get[ThresholdWarnCfg](ctx, s.Client(), routes.GeneralThresholdWarnCfgPath, opts...)
}

// ListApLocRangingCfg retrieves AP location ranging configuration data.
func (s Service) ListApLocRangingCfg(ctx context.Context, opts ...core.GetOption) (*ApLocRangingCfg, error) {
	return core.Get[ApLocRangingCfg](ctx, s.Client(), routes.GeneralApLocRangingCfgPath, opts...)
}

// ListGeolocationCfg retrieves geolocation configuration data.
func (s Service) ListGeolocationCfg(ctx context.Context, opts ...core.GetOption) (*GeolocationCfg, error) {
	return core.Get[GeolocationCfg](ctx, s.Client(), routes.GeneralGeolocationCfgPath, opts...)
}
