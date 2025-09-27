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
func (s Service) GetConfig(ctx context.Context) (*MeshCfg, error) {
	return core.Get[MeshCfg](ctx, s.Client(), routes.MeshCfgPath)
}

// GetOperational retrieves all mesh operational data.
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessMeshOper, error) {
	return core.Get[CiscoIOSXEWirelessMeshOper](ctx, s.Client(), routes.MeshOperPath)
}

// GetOperationalData retrieves mesh operational data including queue stats, data rate stats, security stats, and operational data.
func (s Service) GetOperationalData(ctx context.Context) (*CiscoIOSXEWirelessMeshOper, error) {
	return core.Get[CiscoIOSXEWirelessMeshOper](ctx, s.Client(), routes.MeshGlobalStatsPath)
}

// ListMeshQueueStats retrieves mesh packet queue statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshQueueStats(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshQueueStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshQueueStats](ctx, s.Client(), routes.MeshQueueStatsPath)
}

// ListMeshDataRateStats retrieves mesh data rate statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshDataRateStats(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshDataRateStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshDataRateStats](ctx, s.Client(), routes.MeshDataRateStatsPath)
}

// ListMeshSecurityStats retrieves mesh security statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshSecurityStats(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshSecurityStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshSecurityStats](ctx, s.Client(), routes.MeshSecurityStatsPath)
}

// ListMeshOperationalData retrieves mesh operational data from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshOperationalData(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshOperational, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshOperational](ctx, s.Client(), routes.MeshOperationalDataPath)
}

// GetGlobalStats retrieves mesh global statistics from the controller.
func (s Service) GetGlobalStats(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshGlobalStats, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshGlobalStats](ctx, s.Client(), routes.MeshGlobalStatsPath)
}

// ListApCacInfo retrieves mesh AP CAC information from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListApCacInfo(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshApCacInfo, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApCacInfo](ctx, s.Client(), routes.MeshApCacInfoPath)
}

// ListApPathInfo retrieves mesh AP path information from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListApPathInfo(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshApPathInfo, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApPathInfo](ctx, s.Client(), routes.MeshApPathInfoPath)
}

// ListApTreeData retrieves mesh AP tree data from the controller.
func (s Service) ListApTreeData(ctx context.Context) (*CiscoIOSXEWirelessMeshOperMeshApTreeData, error) {
	return core.Get[CiscoIOSXEWirelessMeshOperMeshApTreeData](ctx, s.Client(), routes.MeshApTreeDataPath)
}
