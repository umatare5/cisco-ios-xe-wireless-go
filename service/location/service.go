package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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
func (s Service) GetConfig(ctx context.Context) (*LocationCfg, error) {
	return core.Get[LocationCfg](ctx, s.Client(), routes.LocationCfgPath)
}

// ListOperatorLocations retrieves location profile configuration data from the wireless controller.
func (s Service) ListOperatorLocations(ctx context.Context) (*OperatorLocations, error) {
	return core.Get[OperatorLocations](ctx, s.Client(), routes.LocationOperatorLocationsPath)
}

// ListNMSPConfig retrieves location NMSP configuration data from the wireless controller.
func (s Service) ListNMSPConfig(ctx context.Context) (*LocationCfgNMSPConfig, error) {
	return core.Get[LocationCfgNMSPConfig](ctx, s.Client(), routes.LocationNMSPConfigPath)
}

// GetLocation retrieves location settings configuration data from the wireless controller.
func (s Service) GetLocation(ctx context.Context) (*LocationSettings, error) {
	return core.Get[LocationSettings](ctx, s.Client(), routes.LocationPath)
}

// GetOperational retrieves all location operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*LocationOper, error) {
	return core.Get[LocationOper](ctx, s.Client(), routes.LocationOperPath)
}

// LocationRSSIMeasurements retrieves location statistics operational data from the wireless controller.
func (s Service) LocationRSSIMeasurements(ctx context.Context) (*LocationLocationRSSIMeasurements, error) {
	return core.Get[LocationLocationRSSIMeasurements](ctx, s.Client(), routes.LocationRSSIMeasurementsPath)
}
