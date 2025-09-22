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
func (s Service) GetOperational(ctx context.Context) (*MeshOper, error) {
	return core.Get[MeshOper](ctx, s.Client(), routes.MeshOperPath)
}

// GetOperationalData retrieves mesh operational data including queue stats, data rate stats, security stats, and operational data.
func (s Service) GetOperationalData(ctx context.Context) (*MeshOper, error) {
	return core.Get[MeshOper](ctx, s.Client(), routes.MeshGlobalStatsPath)
}

// ListMeshQueueStats retrieves mesh packet queue statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshQueueStats(ctx context.Context) (*MeshOperMeshQueueStats, error) {
	return core.Get[MeshOperMeshQueueStats](ctx, s.Client(), routes.MeshQueueStatsPath)
}

// ListMeshDataRateStats retrieves mesh data rate statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshDataRateStats(ctx context.Context) (*MeshOperMeshDataRateStats, error) {
	return core.Get[MeshOperMeshDataRateStats](ctx, s.Client(), routes.MeshDataRateStatsPath)
}

// ListMeshSecurityStats retrieves mesh security statistics from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshSecurityStats(ctx context.Context) (*MeshOperMeshSecurityStats, error) {
	return core.Get[MeshOperMeshSecurityStats](ctx, s.Client(), routes.MeshSecurityStatsPath)
}

// ListMeshOperationalData retrieves mesh operational data from the controller.
// Note: Based on YANG: IOS-XE 17.12.1 - may not be available on all controller versions.
func (s Service) ListMeshOperationalData(ctx context.Context) (*MeshOperMeshOperational, error) {
	return core.Get[MeshOperMeshOperational](ctx, s.Client(), routes.MeshOperationalDataPath)
}
