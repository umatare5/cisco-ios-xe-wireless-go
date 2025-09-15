package site

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
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
func (s Service) GetConfig(ctx context.Context) (*model.SiteCfg, error) {
	return core.Get[model.SiteCfg](ctx, s.Client(), routes.SiteCfgPath)
}

// ListAPProfileConfigs retrieves all AP configuration profiles from the site configuration.
func (s Service) ListAPProfileConfigs(ctx context.Context) (*model.SiteCfgApCfgProfiles, error) {
	return core.Get[model.SiteCfgApCfgProfiles](ctx, s.Client(), routes.APProfilesPath)
}

// ListSiteTagConfigs retrieves all site tag configurations from the site configuration.
func (s Service) ListSiteTagConfigs(ctx context.Context) (*model.SiteCfgSiteTagConfigs, error) {
	return core.Get[model.SiteCfgSiteTagConfigs](ctx, s.Client(), routes.SiteTagConfigsPath)
}
