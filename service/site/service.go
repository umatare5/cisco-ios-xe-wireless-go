package site

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides site configuration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Site service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// SiteTag returns a site tag service instance for site tag management operations.
func (s Service) SiteTag() *SiteTagService {
	return NewSiteTagService(s.Client())
}

// GetConfig retrieves site configuration data including AP configuration profiles and site tag configurations.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessSiteCfg, error) {
	return core.Get[CiscoIOSXEWirelessSiteCfg](ctx, s.Client(), routes.SiteCfgPath, opts...)
}

// ListAPProfileConfigs retrieves all AP configuration profiles from the site configuration.
func (s Service) ListAPProfileConfigs(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessSiteCfgApCfgProfiles, error) {
	return core.Get[CiscoIOSXEWirelessSiteCfgApCfgProfiles](ctx, s.Client(), routes.APProfilesPath, opts...)
}

// ListSiteTagConfigs retrieves all site tag configurations from the site configuration.
func (s Service) ListSiteTagConfigs(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessSiteCfgSiteTagConfigs, error) {
	return core.Get[CiscoIOSXEWirelessSiteCfgSiteTagConfigs](ctx, s.Client(), routes.SiteTagConfigsPath, opts...)
}
