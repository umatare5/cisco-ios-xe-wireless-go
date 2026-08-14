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
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessFabricCfg, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfg](ctx, s.Client(), routes.FabricCfgPath, opts...)
}

// ListCfgFabric retrieves fabric configuration wrapper data.
func (s Service) ListCfgFabric(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabric, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabric](ctx, s.Client(), routes.FabricPath, opts...)
}

// ListCfgFabricProfiles retrieves fabric profiles wrapper data.
func (s Service) ListCfgFabricProfiles(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabricProfiles, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricProfiles](ctx, s.Client(), routes.FabricProfilesPath, opts...)
}

// ListCfgFabricControlplaneNames retrieves fabric controlplane names wrapper data.
func (s Service) ListCfgFabricControlplaneNames(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabricControlplaneNames, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricControlplaneNames](
		ctx,
		s.Client(),
		routes.FabricControlplaneNamesPath, opts...,
	)
}

// ListFabricConfig retrieves fabric global configuration data.
func (s Service) ListFabricConfig(ctx context.Context, opts ...core.GetOption) (*FabricConfig, error) {
	return core.Get[FabricConfig](ctx, s.Client(), routes.FabricPath, opts...)
}

// ListFabricProfiles retrieves fabric profiles data.
func (s Service) ListFabricProfiles(ctx context.Context, opts ...core.GetOption) (*FabricProfiles, error) {
	return core.Get[FabricProfiles](ctx, s.Client(), routes.FabricProfilesPath, opts...)
}

// ListFabricProfile retrieves individual fabric profile entries.
func (s Service) ListFabricProfile(ctx context.Context, opts ...core.GetOption) (*FabricProfile, error) {
	return core.Get[FabricProfile](ctx, s.Client(), routes.FabricProfilesPath, opts...)
}

// ListFabricControlplanes retrieves fabric controlplanes data.
func (s Service) ListFabricControlplanes(ctx context.Context, opts ...core.GetOption) (*FabricControlplanes, error) {
	return core.Get[FabricControlplanes](ctx, s.Client(), routes.FabricControlplaneNamesPath, opts...)
}

// ListFabricControlplaneName retrieves individual fabric controlplane name entries.
func (s Service) ListFabricControlplaneName(
	ctx context.Context,
	opts ...core.GetOption,
) (*FabricControlplaneName, error) {
	return core.Get[FabricControlplaneName](ctx, s.Client(), routes.FabricControlplaneNamesPath, opts...)
}
