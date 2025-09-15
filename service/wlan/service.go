// Package wlan provides WLAN domain services for Cisco IOS-XE Wireless Network Controller API.
package wlan

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
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
func (s Service) ListProfileConfigs(ctx context.Context) (*model.WlanCfgEntries, error) {
	return core.Get[model.WlanCfgEntries](ctx, s.Client(), routes.WLANWlanCfgEntriesPath)
}

// GetProfileConfig retrieves a specific WLAN configuration entry by profile name.
func (s Service) GetProfileConfig(ctx context.Context, profileName string) (*model.WlanCfgEntry, error) {
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.WLANWlanCfgEntryPath, profileName)
	return core.Get[model.WlanCfgEntry](ctx, s.Client(), url)
}

// ListPolicies retrieves WLAN policies.
func (s Service) ListPolicies(ctx context.Context) (*model.WlanPolicies, error) {
	return core.Get[model.WlanPolicies](ctx, s.Client(), routes.WLANWlanPoliciesPath)
}

// ListPolicyListEntries retrieves all policy list entries.
func (s Service) ListPolicyListEntries(ctx context.Context) (*model.PolicyListEntries, error) {
	return core.Get[model.PolicyListEntries](ctx, s.Client(), routes.WLANPolicyListEntriesPath)
}

// ListWirelessAAAPolicyConfigs retrieves wireless AAA policy configurations.
func (s Service) ListWirelessAAAPolicyConfigs(ctx context.Context) (*model.WirelessAaaPolicyConfigs, error) {
	return core.Get[model.WirelessAaaPolicyConfigs](ctx, s.Client(), routes.WLANWirelessAaaPolicyConfigsPath)
}

// GetConfig retrieves WLAN configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*model.WlanCfg, error) {
	return core.Get[model.WlanCfg](ctx, s.Client(), routes.WLANCfgPath)
}

// GetOperational retrieves WLAN operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.WlanGlobalOper, error) {
	return core.Get[model.WlanGlobalOper](ctx, s.Client(), routes.WLANGlobalOperPath)
}
