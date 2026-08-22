package mesh

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides wireless mesh networking operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Mesh service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves all mesh configuration data.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*MeshCfg, error) {
	return core.Get[MeshCfg](ctx, s.Client(), routes.MeshCfgPath, opts...)
}

// GetOperational retrieves all mesh operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessMeshOper, error) {
	return core.Get[CiscoIOSXEWirelessMeshOper](ctx, s.Client(), routes.MeshOperPath, opts...)
}

// ListMeshQueueStats retrieves mesh packet queue statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshQueueStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshQueueStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshQueueStats](ctx, s.Client(), routes.MeshQueueStatsPath, opts...)
}

// ListMeshDataRateStats retrieves mesh data rate statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshDataRateStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshDataRateStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshDataRateStats](ctx, s.Client(), routes.MeshDataRateStatsPath, opts...)
}

// ListMeshSecurityStats retrieves mesh security statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshSecurityStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshSecurityStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshSecurityStats](ctx, s.Client(), routes.MeshSecurityStatsPath, opts...)
}

// ListMeshOperationalData retrieves mesh operational data from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshOperationalData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshOperational, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshOperational](ctx, s.Client(), routes.MeshOperationalDataPath, opts...)
}

// GetGlobalStats retrieves mesh global statistics from the controller.
func (s Service) GetGlobalStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshGlobalStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshGlobalStats](ctx, s.Client(), routes.MeshGlobalStatsPath, opts...)
}

// ListApCacInfo retrieves mesh AP CAC information from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListApCacInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshApCacInfo, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApCacInfo](ctx, s.Client(), routes.MeshApCacInfoPath, opts...)
}

// ListApPathInfo retrieves mesh AP path information from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListApPathInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshApPathInfo, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApPathInfo](ctx, s.Client(), routes.MeshApPathInfoPath, opts...)
}

// ListApTreeData retrieves mesh AP tree data from the controller.
func (s Service) ListApTreeData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMeshOperMeshApTreeData, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApTreeData](ctx, s.Client(), routes.MeshApTreeDataPath, opts...)
}
