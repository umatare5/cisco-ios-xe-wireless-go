package fabric

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides SD-Access Fabric operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Fabric service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves fabric configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessFabricCfg, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfg](ctx, s.Client(), routes.FabricCfgPath)
}

// ListCfgFabric retrieves fabric configuration wrapper data.
func (s Service) ListCfgFabric(ctx context.Context) (*CiscoIOSXEWirelessFabricCfgFabric, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabric](ctx, s.Client(), routes.FabricPath)
}

// ListCfgFabricProfiles retrieves fabric profiles wrapper data.
func (s Service) ListCfgFabricProfiles(ctx context.Context) (*CiscoIOSXEWirelessFabricCfgFabricProfiles, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricProfiles](ctx, s.Client(), routes.FabricProfilesPath)
}

// ListCfgFabricControlplaneNames retrieves fabric controlplane names wrapper data.
func (s Service) ListCfgFabricControlplaneNames(
	ctx context.Context,
) (*CiscoIOSXEWirelessFabricCfgFabricControlplaneNames, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricControlplaneNames](
		ctx,
		s.Client(),
		routes.FabricControlplaneNamesPath,
	)
}

// ListFabricConfig retrieves fabric global configuration data.
func (s Service) ListFabricConfig(ctx context.Context) (*FabricConfig, error) {
	return core.Get[FabricConfig](ctx, s.Client(), routes.FabricPath)
}

// ListFabricProfiles retrieves fabric profiles data.
func (s Service) ListFabricProfiles(ctx context.Context) (*FabricProfiles, error) {
	return core.Get[FabricProfiles](ctx, s.Client(), routes.FabricProfilesPath)
}

// ListFabricProfile retrieves individual fabric profile entries.
func (s Service) ListFabricProfile(ctx context.Context) (*FabricProfile, error) {
	return core.Get[FabricProfile](ctx, s.Client(), routes.FabricProfilesPath)
}

// ListFabricControlplanes retrieves fabric controlplanes data.
func (s Service) ListFabricControlplanes(ctx context.Context) (*FabricControlplanes, error) {
	return core.Get[FabricControlplanes](ctx, s.Client(), routes.FabricControlplaneNamesPath)
}

// ListFabricControlplaneName retrieves individual fabric controlplane name entries.
func (s Service) ListFabricControlplaneName(ctx context.Context) (*FabricControlplaneName, error) {
	return core.Get[FabricControlplaneName](ctx, s.Client(), routes.FabricControlplaneNamesPath)
}
