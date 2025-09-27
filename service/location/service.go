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
func (s Service) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessLocationCfg, error) {
	return core.Get[CiscoIOSXEWirelessLocationCfg](ctx, s.Client(), routes.LocationCfgPath)
}

// ListOperatorLocations retrieves location profile configuration data from the wireless controller.
func (s Service) ListOperatorLocations(ctx context.Context) (*OperatorLocations, error) {
	return core.Get[OperatorLocations](ctx, s.Client(), routes.LocationOperatorLocationsPath)
}

// ListNMSPConfig retrieves location NMSP configuration data from the wireless controller.
func (s Service) ListNMSPConfig(ctx context.Context) (*CiscoIOSXEWirelessLocationCfgNMSPConfig, error) {
	return core.Get[CiscoIOSXEWirelessLocationCfgNMSPConfig](ctx, s.Client(), routes.LocationNMSPConfigPath)
}

// GetLocation retrieves location settings configuration data from the wireless controller.
func (s Service) GetLocation(ctx context.Context) (*CiscoIOSXEWirelessLocationSettings, error) {
	return core.Get[CiscoIOSXEWirelessLocationSettings](ctx, s.Client(), routes.LocationPath)
}

// GetOperational retrieves all location operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessLocationOper, error) {
	return core.Get[CiscoIOSXEWirelessLocationOper](ctx, s.Client(), routes.LocationOperPath)
}

// LocationRSSIMeasurements retrieves location statistics operational data from the wireless controller.
func (s Service) LocationRSSIMeasurements(
	ctx context.Context,
) (*CiscoIOSXEWirelessLocationLocationRSSIMeasurements, error) {
	return core.Get[CiscoIOSXEWirelessLocationLocationRSSIMeasurements](
		ctx,
		s.Client(),
		routes.LocationRSSIMeasurementsPath,
	)
}
