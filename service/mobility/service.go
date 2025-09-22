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

// GetConfig retrieves complete mobility configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*MobilityCfg, error) {
	return core.Get[MobilityCfg](ctx, s.Client(), routes.MobilityCfgPath)
}

// ListMobilityConfig retrieves mobility configuration data using wrapper structure.
func (s Service) ListMobilityConfig(ctx context.Context) (*MobilityCfgMobilityConfig, error) {
	return core.Get[MobilityCfgMobilityConfig](ctx, s.Client(), routes.MobilityConfigPath)
}

// GetOperational retrieves mobility operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*MobilityOper, error) {
	return core.Get[MobilityOper](ctx, s.Client(), routes.MobilityOperPath)
}

// ListAPCache retrieves AP cache data.
func (s Service) ListAPCache(ctx context.Context) (*MobilityOperApCache, error) {
	return core.Get[MobilityOperApCache](ctx, s.Client(), routes.MobilityApCachePath)
}

// ListAPPeers retrieves AP peer list data.
func (s Service) ListAPPeers(ctx context.Context) (*MobilityOperApPeerList, error) {
	return core.Get[MobilityOperApPeerList](ctx, s.Client(), routes.MobilityApPeerListPath)
}

// GetMMGlobalInfo retrieves MM global information.
func (s Service) GetMMGlobalInfo(ctx context.Context) (*MobilityOperMmGlobalData, error) {
	return core.Get[MobilityOperMmGlobalData](ctx, s.Client(), routes.MobilityMmGlobalDataPath)
}

// GetMMIFGlobalStats retrieves MM interface global statistics.
func (s Service) GetMMIFGlobalStats(ctx context.Context) (*MobilityOperMmIfGlobalStats, error) {
	return core.Get[MobilityOperMmIfGlobalStats](
		ctx, s.Client(), routes.MobilityMmIfGlobalStatsPath)
}

// ListClients retrieves mobility client data.
func (s Service) ListClients(ctx context.Context) (*MobilityOperMobilityClientData, error) {
	return core.Get[MobilityOperMobilityClientData](
		ctx, s.Client(), routes.MobilityClientDataPath)
}

// GetGlobalStats retrieves mobility global statistics.
func (s Service) GetGlobalStats(
	ctx context.Context,
) (*MobilityOperMobilityGlobalStats, error) {
	return core.Get[MobilityOperMobilityGlobalStats](
		ctx, s.Client(), routes.MobilityGlobalStatsPath,
	)
}

// ListMmIfGlobalMsgStats retrieves MM interface global message statistics.
func (s Service) ListMmIfGlobalMsgStats(ctx context.Context) (*MobilityOperMmIfGlobalMsgStats, error) {
	return core.Get[MobilityOperMmIfGlobalMsgStats](ctx, s.Client(), routes.MobilityMmIfGlobalMsgStatsPath)
}

// ListClientStats retrieves mobility client statistics.
func (s Service) ListClientStats(ctx context.Context) (*MobilityOperMobilityClientStats, error) {
	return core.Get[MobilityOperMobilityClientStats](ctx, s.Client(), routes.MobilityClientStatsPath)
}

// ListGlobalDTLSStats retrieves mobility global DTLS statistics.
func (s Service) ListGlobalDTLSStats(ctx context.Context) (*MobilityOperMobilityGlobalDTLSStats, error) {
	return core.Get[MobilityOperMobilityGlobalDTLSStats](ctx, s.Client(), routes.MobilityGlobalDTLSStatsPath)
}

// ListGlobalMsgStats retrieves mobility global message statistics.
func (s Service) ListGlobalMsgStats(ctx context.Context) (*MobilityOperMobilityGlobalMsgStats, error) {
	return core.Get[MobilityOperMobilityGlobalMsgStats](ctx, s.Client(), routes.MobilityGlobalMsgStatsPath)
}

// ListWlanClientLimit retrieves WLAN client limit data.
func (s Service) ListWlanClientLimit(ctx context.Context) (*MobilityOperWlanClientLimit, error) {
	return core.Get[MobilityOperWlanClientLimit](ctx, s.Client(), routes.MobilityWlanClientLimitPath)
}
