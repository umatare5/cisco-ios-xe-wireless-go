package mobility

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mobility"
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

// GetOperational retrieves mobility operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.MobilityOper, error) {
	return core.Get[model.MobilityOper](ctx, s.Client(), routes.MobilityOperPath)
}

// ListAPCache retrieves AP cache data.
func (s Service) ListAPCache(ctx context.Context) (*model.MobilityOperApCache, error) {
	return core.Get[model.MobilityOperApCache](ctx, s.Client(), routes.MobilityApCachePath)
}

// ListAPPeers retrieves AP peer list data.
func (s Service) ListAPPeers(ctx context.Context) (*model.MobilityOperApPeerList, error) {
	return core.Get[model.MobilityOperApPeerList](ctx, s.Client(), routes.MobilityApPeerListPath)
}

// GetMMGlobalInfo retrieves MM global information.
func (s Service) GetMMGlobalInfo(ctx context.Context) (*model.MobilityOperMmGlobalData, error) {
	return core.Get[model.MobilityOperMmGlobalData](ctx, s.Client(), routes.MobilityMmGlobalDataPath)
}

// GetMMIFGlobalStats retrieves MM interface global statistics.
func (s Service) GetMMIFGlobalStats(ctx context.Context) (*model.MobilityOperMmIfGlobalStats, error) {
	return core.Get[model.MobilityOperMmIfGlobalStats](
		ctx, s.Client(), routes.MobilityMmIfGlobalStatsPath)
}

// ListClients retrieves mobility client data.
func (s Service) ListClients(ctx context.Context) (*model.MobilityOperMobilityClientData, error) {
	return core.Get[model.MobilityOperMobilityClientData](
		ctx, s.Client(), routes.MobilityClientDataPath)
}

// GetGlobalStats retrieves mobility global statistics.
func (s Service) GetGlobalStats(
	ctx context.Context,
) (*model.MobilityOperMobilityGlobalStats, error) {
	return core.Get[model.MobilityOperMobilityGlobalStats](
		ctx, s.Client(), routes.MobilityGlobalStatsPath,
	)
}
