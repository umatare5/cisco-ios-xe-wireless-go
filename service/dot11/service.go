package dot11

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides IEEE 802.11 wireless configuration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new 802.11 service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves dot11 configuration data from the controller.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessDot11Cfg, error) {
	return core.Get[CiscoIOSXEWirelessDot11Cfg](ctx, s.Client(), routes.Dot11CfgPath, opts...)
}

// ListCfgConfiguredCountries retrieves dot11 configured countries wrapper data.
func (s Service) ListCfgConfiguredCountries(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessDot11CfgConfiguredCountries, error) {
	return core.Get[CiscoIOSXEWirelessDot11CfgConfiguredCountries](
		ctx,
		s.Client(),
		routes.Dot11ConfiguredCountriesPath,
		opts...,
	)
}

// ListCfgDot11Entries retrieves dot11 entries wrapper data.
func (s Service) ListCfgDot11Entries(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessDot11CfgDot11Entries, error) {
	return core.Get[CiscoIOSXEWirelessDot11CfgDot11Entries](ctx, s.Client(), routes.Dot11EntriesPath, opts...)
}

// ListCfgDot11acMcsEntries retrieves dot11ac MCS entries wrapper data.
func (s Service) ListCfgDot11acMcsEntries(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessDot11CfgDot11acMcsEntries, error) {
	return core.Get[CiscoIOSXEWirelessDot11CfgDot11acMcsEntries](ctx, s.Client(), routes.Dot11AcMcsEntriesPath, opts...)
}
