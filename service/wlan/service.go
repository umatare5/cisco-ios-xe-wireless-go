// Package wlan provides WLAN domain services for Cisco IOS-XE Wireless Network Controller API.
package wlan

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides WLAN operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new WLAN service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// PolicyTag returns a PolicyTagService for policy tag operations.
func (s Service) PolicyTag() *PolicyTagService {
	return NewPolicyTagService(s.Client())
}

// ListProfileConfigs retrieves WLAN configuration entries.
func (s Service) ListProfileConfigs(ctx context.Context) (*WlanCfgEntries, error) {
	return core.Get[WlanCfgEntries](ctx, s.Client(), routes.WLANWlanCfgEntriesPath)
}

// GetProfileConfig retrieves a specific WLAN configuration entry by profile name.
func (s Service) GetProfileConfig(ctx context.Context, profileName string) (*WlanCfgEntry, error) {
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.WLANWlanCfgEntryPath, profileName)
	return core.Get[WlanCfgEntry](ctx, s.Client(), url)
}

// ListPolicies retrieves WLAN policies.
func (s Service) ListPolicies(ctx context.Context) (*WlanPolicies, error) {
	return core.Get[WlanPolicies](ctx, s.Client(), routes.WLANWlanPoliciesPath)
}

// ListPolicyListEntries retrieves all policy list entries.
func (s Service) ListPolicyListEntries(ctx context.Context) (*PolicyListEntries, error) {
	return core.Get[PolicyListEntries](ctx, s.Client(), routes.WLANPolicyListEntriesPath)
}

// ListWirelessAAAPolicyConfigs retrieves wireless AAA policy configurations.
func (s Service) ListWirelessAAAPolicyConfigs(ctx context.Context) (*WirelessAaaPolicyConfigs, error) {
	return core.Get[WirelessAaaPolicyConfigs](ctx, s.Client(), routes.WLANWirelessAaaPolicyConfigsPath)
}

// GetConfig retrieves WLAN configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*WlanCfg, error) {
	return core.Get[WlanCfg](ctx, s.Client(), routes.WLANCfgPath)
}

// GetOperational retrieves WLAN operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*WlanGlobalOper, error) {
	return core.Get[WlanGlobalOper](ctx, s.Client(), routes.WLANGlobalOperPath)
}
