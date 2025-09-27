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

// GetConfig retrieves the complete WLAN configuration.
func (s Service) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessWlanCfg, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfg](ctx, s.Client(), routes.WLANCfgPath)
}

// GetOperational retrieves WLAN operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*WlanGlobalOper, error) {
	return core.Get[WlanGlobalOper](ctx, s.Client(), routes.WLANGlobalOperPath)
}

// ListConfigEntries retrieves WLAN configuration entries.
func (s Service) ListConfigEntries(ctx context.Context) (*WlanCfgEntries, error) {
	return core.Get[WlanCfgEntries](ctx, s.Client(), routes.WLANWlanCfgEntriesPath)
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

// ListWlanCfgEntries retrieves WLAN configuration entries using WlanCfgWlanCfgEntries wrapper.
func (s Service) ListWlanCfgEntries(ctx context.Context) (*CiscoIOSXEWirelessWlanCfgWlanCfgEntries, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfgWlanCfgEntries](ctx, s.Client(), routes.WLANWlanCfgEntriesPath)
}

// ListWlanPolicies retrieves WLAN policies using WlanCfgWlanPolicies wrapper.
func (s Service) ListWlanPolicies(ctx context.Context) (*CiscoIOSXEWirelessWlanCfgWlanPolicies, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfgWlanPolicies](ctx, s.Client(), routes.WLANWlanPoliciesPath)
}

// ListCfgPolicyListEntries retrieves policy list entries using WlanCfgPolicyListEntries wrapper.
func (s Service) ListCfgPolicyListEntries(ctx context.Context) (*CiscoIOSXEWirelessWlanCfgPolicyListEntries, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfgPolicyListEntries](ctx, s.Client(), routes.WLANPolicyListEntriesPath)
}

// ListCfgWirelessAaaPolicyConfigs retrieves wireless AAA policy configurations using WlanCfgWirelessAaaPolicyConfigs wrapper.
func (s Service) ListCfgWirelessAaaPolicyConfigs(
	ctx context.Context,
) (*CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs](
		ctx,
		s.Client(),
		routes.WLANWirelessAaaPolicyConfigsPath,
	)
}

// ListDot11beProfiles retrieves Wi-Fi 7 / 802.11be profiles using WlanCfgDot11beProfiles wrapper.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListDot11beProfiles(ctx context.Context) (*CiscoIOSXEWirelessWlanCfgDot11beProfiles, error) {
	return core.Get[CiscoIOSXEWirelessWlanCfgDot11beProfiles](ctx, s.Client(), routes.WLANDot11beProfilesPath)
}

// ListWlanInfo retrieves WLAN information using WlanGlobalOperWlanInfo wrapper.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListWlanInfo(ctx context.Context) (*CiscoIOSXEWirelessWlanGlobalOperWlanInfo, error) {
	return core.Get[CiscoIOSXEWirelessWlanGlobalOperWlanInfo](ctx, s.Client(), routes.WLANWlanInfoPath)
}
