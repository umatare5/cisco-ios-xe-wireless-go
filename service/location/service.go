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

// GetConfig retrieves complete location configuration data from the wireless controller.
func (s Service) GetConfig(ctx context.Context) (*model.LocationCfg, error) {
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationCfgPath)
}

// ListOperatorLocations retrieves location profile configuration data from the wireless controller.
func (s Service) ListOperatorLocations(ctx context.Context) (*model.OperatorLocations, error) {
	return core.Get[model.OperatorLocations](ctx, s.Client(), routes.LocationOperatorLocationsPath)
}

// ListNmspConfig retrieves location NMSP configuration data from the wireless controller.
func (s Service) ListNmspConfig(ctx context.Context) (*model.LocationCfgNmspConfig, error) {
	return core.Get[model.LocationCfgNmspConfig](ctx, s.Client(), routes.LocationNmspConfigPath)
}

// GetLocation retrieves location settings configuration data from the wireless controller.
func (s Service) GetLocation(ctx context.Context) (*model.LocationSettings, error) {
	return core.Get[model.LocationSettings](ctx, s.Client(), routes.LocationPath)
}

// GetOperational retrieves all location operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*model.LocationOper, error) {
	return core.Get[model.LocationOper](ctx, s.Client(), routes.LocationOperPath)
}

// LocationRssiMeasurements retrieves location statistics operational data from the wireless controller.
func (s Service) LocationRssiMeasurements(ctx context.Context) (*model.LocationLocationRssiMeasurements, error) {
	return core.Get[model.LocationLocationRssiMeasurements](ctx, s.Client(), routes.LocationRssiMeasurementsPath)
}
