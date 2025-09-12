package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/location"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides location services operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Location service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

func (s Service) GetConfig(ctx context.Context) (*model.LocationCfg, error) {
	return nil, core.ErrResourceNotFound
}

// ListProfileConfigs retrieves location profile configuration data from the wireless controller.
func (s Service) ListProfileConfigs(ctx context.Context) (*model.LocationCfg, error) {
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationOperatorLocationsPath)
}

// ListServerConfigs retrieves location server configuration data from the wireless controller.
func (s Service) ListServerConfigs(ctx context.Context) (*model.LocationCfg, error) {
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationNmspConfigPath)
}

// GetSettingsConfig retrieves location settings configuration data from the wireless controller.
func (s Service) GetSettingsConfig(ctx context.Context) (*model.LocationCfg, error) {
	return nil, core.ErrResourceNotFound
}

// GetOperational retrieves all location operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*model.LocationCfg, error) {
	// Return error for now since this endpoint doesn't exist in this WNC version
	return nil, core.ErrResourceNotFound
}

// GetStats retrieves location statistics operational data from the wireless controller.
func (s Service) GetStats(ctx context.Context) (*model.LocationCfg, error) {
	// Return error for now since this endpoint doesn't exist in this WNC version
	return nil, core.ErrResourceNotFound
}
