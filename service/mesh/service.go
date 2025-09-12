package mesh

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mesh"
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
func (s Service) GetConfig(ctx context.Context) (*model.MeshCfg, error) {
	return core.Get[model.MeshCfg](ctx, s.Client(), routes.MeshCfgPath)
}

// GetOperational retrieves all mesh operational data.
func (s Service) GetOperational(ctx context.Context) (*model.MeshOper, error) {
	return core.Get[model.MeshOper](ctx, s.Client(), routes.MeshOperPath)
}

// GetOperationalData retrieves mesh operational data including queue stats, data rate stats, security stats, and operational data.
func (s Service) GetOperationalData(ctx context.Context) (*model.MeshOper, error) {
	return core.Get[model.MeshOper](ctx, s.Client(), routes.MeshGlobalStatsPath)
}
