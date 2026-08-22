package mobility

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides client mobility management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Mobility service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the mobility configuration.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessMobilityCfg, error) {
	return core.Get[CiscoIOSXEWirelessMobilityCfg](ctx, s.Client(), routes.MobilityCfgPath, opts...)
}

// ListMobilityConfig retrieves mobility configuration data using wrapper structure.
func (s Service) ListMobilityConfig(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityCfgMobilityConfig, error) {
	return core.Get[CiscoIOSXEWirelessMobilityCfgMobilityConfig](ctx, s.Client(), routes.MobilityConfigPath, opts...)
}

// GetOperational retrieves mobility operational data from the controller.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessMobilityOper, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOper](ctx, s.Client(), routes.MobilityOperPath, opts...)
}

// ListAPCache retrieves AP cache data.
func (s Service) ListAPCache(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperApCache, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperApCache](ctx, s.Client(), routes.MobilityApCachePath, opts...)
}

// ListAPPeers retrieves AP peer list data.
func (s Service) ListAPPeers(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperApPeerList, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperApPeerList](ctx, s.Client(), routes.MobilityApPeerListPath, opts...)
}

// GetMMGlobalInfo retrieves MM global information.
func (s Service) GetMMGlobalInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMmGlobalData, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMmGlobalData](
		ctx,
		s.Client(),
		routes.MobilityMmGlobalDataPath,
		opts...,
	)
}

// GetMMIFGlobalStats retrieves MM interface global statistics.
func (s Service) GetMMIFGlobalStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMmIfGlobalStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMmIfGlobalStats](
		ctx, s.Client(), routes.MobilityMmIfGlobalStatsPath, opts...,
	)
}

// ListClients retrieves mobility client data.
func (s Service) ListClients(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMobilityClientData, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMobilityClientData](
		ctx, s.Client(), routes.MobilityClientDataPath, opts...,
	)
}

// GetGlobalStats retrieves mobility global statistics.
func (s Service) GetGlobalStats(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMobilityGlobalStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMobilityGlobalStats](
		ctx, s.Client(), routes.MobilityGlobalStatsPath, opts...,
	)
}

// ListMmIfGlobalMsgStats retrieves MM interface global message statistics.
func (s Service) ListMmIfGlobalMsgStats(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMmIfGlobalMsgStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMmIfGlobalMsgStats](
		ctx,
		s.Client(),
		routes.MobilityMmIfGlobalMsgStatsPath, opts...,
	)
}

// ListClientStats retrieves mobility client statistics.
func (s Service) ListClientStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMobilityClientStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMobilityClientStats](
		ctx,
		s.Client(),
		routes.MobilityClientStatsPath,
		opts...,
	)
}

// ListGlobalDTLSStats retrieves mobility global DTLS statistics.
func (s Service) ListGlobalDTLSStats(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMobilityGlobalDTLSStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMobilityGlobalDTLSStats](
		ctx,
		s.Client(),
		routes.MobilityGlobalDTLSStatsPath, opts...,
	)
}

// ListGlobalMsgStats retrieves mobility global message statistics.
func (s Service) ListGlobalMsgStats(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperMobilityGlobalMsgStats, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperMobilityGlobalMsgStats](
		ctx,
		s.Client(),
		routes.MobilityGlobalMsgStatsPath, opts...,
	)
}

// ListWlanClientLimit retrieves WLAN client limit data.
func (s Service) ListWlanClientLimit(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMobilityOperWlanClientLimit, error) {
	return core.Get[CiscoIOSXEWirelessMobilityOperWlanClientLimit](
		ctx,
		s.Client(),
		routes.MobilityWlanClientLimitPath,
		opts...,
	)
}
