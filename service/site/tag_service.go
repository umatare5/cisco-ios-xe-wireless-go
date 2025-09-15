package site

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// SiteTagService provides site tag management operations.
type SiteTagService struct {
	service.BaseService
}

// NewSiteTagService creates a new site tag service.
func NewSiteTagService(client *core.Client) *SiteTagService {
	return &SiteTagService{
		BaseService: service.NewBaseService(client),
	}
}

// GetSiteTag retrieves a specific site tag configuration.
func (s *SiteTagService) GetSiteTag(ctx context.Context, tagName string) (*model.SiteListEntry, error) {
	if err := s.validateTagName(tagName); err != nil {
		return nil, err
	}

	result, err := core.Get[model.SiteCfgSiteTagConfig](ctx, s.Client(), s.buildTagURL(tagName))
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	if len(result.SiteListEntry) == 0 {
		return nil, nil
	}

	return &result.SiteListEntry[0], nil
}

// ListSiteTags retrieves all site tag configurations.
func (s *SiteTagService) ListSiteTags(ctx context.Context) ([]model.SiteListEntry, error) {
	type siteTagsResponse struct {
		SiteTagConfigs model.SiteTagConfigs `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs"`
	}

	result, err := core.Get[siteTagsResponse](ctx, s.Client(), routes.SiteTagConfigsPath)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return []model.SiteListEntry{}, nil
	}
	if len(result.SiteTagConfigs.SiteTagConfig) == 0 {
		return []model.SiteListEntry{}, nil
	}

	return result.SiteTagConfigs.SiteTagConfig, nil
}

// CreateSiteTag creates a new site tag configuration.
func (s *SiteTagService) CreateSiteTag(ctx context.Context, config *model.SiteListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if config.SiteTagName == "" {
		return errors.New("site tag name cannot be empty")
	}
	if err := s.validateTagName(config.SiteTagName); err != nil {
		return err
	}

	// Convert service payload model to internal model
	payload := s.buildPayload(*config)
	return core.PostVoid(ctx, s.Client(), routes.SiteTagConfigsPath, payload)
}

// SetAPJoinProfile sets the AP join profile for a site tag.
func (s *SiteTagService) SetAPJoinProfile(ctx context.Context, siteTagName, apJoinProfile string) error {
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", siteTagName, err))
	}
	if config == nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", siteTagName))
	}

	// Update AP join profile
	config.ApJoinProfile = &apJoinProfile
	return s.setSiteTag(ctx, config)
}

// SetFlexProfile sets the flex profile for a site tag.
func (s *SiteTagService) SetFlexProfile(ctx context.Context, siteTagName, flexProfile string) error {
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", siteTagName, err))
	}

	if config == nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", siteTagName))
	}

	// Update flex profile
	config.FlexProfile = &flexProfile
	// Explicitly set to false for flex-profile compatibility
	isLocalSite := false
	config.IsLocalSite = &isLocalSite
	return s.setSiteTag(ctx, config)
}

// SetLocalSite sets the local site mode for a site tag.
func (s *SiteTagService) SetLocalSite(ctx context.Context, siteTagName string, enabled bool) error {
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", siteTagName, err))
	}

	if config == nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", siteTagName))
	}

	// Update local site
	config.IsLocalSite = &enabled
	return s.setSiteTag(ctx, config)
}

// SetDescription sets the description for a site tag.
func (s *SiteTagService) SetDescription(ctx context.Context, siteTagName, description string) error {
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", siteTagName, err))
	}
	if config == nil {
		return fmt.Errorf("site tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", siteTagName))
	}

	// Update description
	config.Description = &description
	return s.setSiteTag(ctx, config)
}

// DeleteSiteTag deletes a site tag configuration.
func (s *SiteTagService) DeleteSiteTag(ctx context.Context, siteTagName string) error {
	if err := s.validateTagName(siteTagName); err != nil {
		return err
	}
	return core.Delete(ctx, s.Client(), s.buildTagURL(siteTagName))
}

// setSiteTag sets/updates an existing site tag configuration.
func (s *SiteTagService) setSiteTag(ctx context.Context, config *model.SiteListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if config.SiteTagName == "" {
		return errors.New("site tag name cannot be empty")
	}
	if err := s.validateTagName(config.SiteTagName); err != nil {
		return err
	}

	// Convert service payload model to internal model
	payload := s.buildPayload(*config)
	return core.PatchVoid(ctx, s.Client(), s.buildTagURL(config.SiteTagName), payload)
}

// validateTagName validates site tag name.
func (s *SiteTagService) validateTagName(tagName string) error {
	if tagName == "" {
		return errors.New("site tag name cannot be empty")
	}
	if strings.TrimSpace(tagName) == "" {
		return fmt.Errorf("site tag validation failed: %w",
			fmt.Errorf("invalid tag name format: '%s'", tagName))
	}
	return nil
}

// buildTagURL builds URL for specific tag operations using RESTCONF builder.
func (s *SiteTagService) buildTagURL(tagName string) string {
	return s.Client().RESTCONFBuilder().BuildQueryURL(routes.SiteTagConfigQueryPath, tagName)
}

// buildPayload builds a payload for tag operations using the request model.
func (s *SiteTagService) buildPayload(config model.SiteListEntry) model.SiteTagConfigsPayload {
	return model.SiteTagConfigsPayload{
		SiteListEntry: config,
	}
}
