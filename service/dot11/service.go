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
func (s Service) GetConfig(ctx context.Context) (*Dot11Cfg, error) {
	return core.Get[Dot11Cfg](ctx, s.Client(), routes.Dot11CfgPath)
}

// ListCfgFilters retrieves dot11 configuration filters.
func (s Service) ListCfgFilters(ctx context.Context) (*Dot11CfgFilter, error) {
	return core.Get[Dot11CfgFilter](ctx, s.Client(), routes.Dot11CfgPath)
}

// ListCfgConfiguredCountries retrieves dot11 configured countries wrapper data.
func (s Service) ListCfgConfiguredCountries(ctx context.Context) (*Dot11CfgConfiguredCountries, error) {
	return core.Get[Dot11CfgConfiguredCountries](ctx, s.Client(), routes.Dot11ConfiguredCountriesPath)
}

// ListCfgDot11Entries retrieves dot11 entries wrapper data.
func (s Service) ListCfgDot11Entries(ctx context.Context) (*Dot11CfgDot11Entries, error) {
	return core.Get[Dot11CfgDot11Entries](ctx, s.Client(), routes.Dot11EntriesPath)
}

// ListCfgDot11acMcsEntries retrieves dot11ac MCS entries wrapper data.
func (s Service) ListCfgDot11acMcsEntries(ctx context.Context) (*Dot11CfgDot11acMcsEntries, error) {
	return core.Get[Dot11CfgDot11acMcsEntries](ctx, s.Client(), routes.Dot11AcMcsEntriesPath)
}

// ListConfiguredCountries retrieves configured countries data.
func (s Service) ListConfiguredCountries(ctx context.Context) (*Dot11ConfiguredCountries, error) {
	return core.Get[Dot11ConfiguredCountries](ctx, s.Client(), routes.Dot11ConfiguredCountriesPath)
}

// ListDot11Entries retrieves 802.11 entries data.
func (s Service) ListDot11Entries(ctx context.Context) (*Dot11Entries, error) {
	return core.Get[Dot11Entries](ctx, s.Client(), routes.Dot11EntriesPath)
}
