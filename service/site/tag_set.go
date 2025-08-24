package site

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
)

// SetSiteTag sets/updates an existing site tag configuration.
// This function modifies an existing site tag configuration.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Updated site tag configuration
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) SetSiteTag(ctx context.Context, config SiteTagConfig) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	// Convert service model to internal model
	internalConfig := model.SiteTagConfig{
		SiteTagName:   config.SiteTagName,
		Description:   config.Description,
		FlexProfile:   config.FlexProfile,
		ApJoinProfile: config.APJoinProfile,
	}

	// Convert *bool to bool for IsLocalSite
	if config.IsLocalSite != nil {
		internalConfig.IsLocalSite = *config.IsLocalSite
	}

	return s.tagOps.Update(ctx, internalConfig, config.SiteTagName)
}

// SetAPJoinProfile sets the AP join profile for a site tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - siteTagName: Name of the site tag
//   - apJoinProfile: AP join profile name to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) SetAPJoinProfile(ctx context.Context, siteTagName, apJoinProfile string) error {
	// Get existing site tag
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("failed to get site tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("site tag %s not found", siteTagName)
	}

	// Update AP join profile
	config.APJoinProfile = apJoinProfile
	return s.SetSiteTag(ctx, *config)
}

// SetFlexProfile sets the flex profile for a site tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - siteTagName: Name of the site tag
//   - flexProfile: Flex profile name to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) SetFlexProfile(ctx context.Context, siteTagName, flexProfile string) error {
	// Get existing site tag
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("failed to get site tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("site tag %s not found", siteTagName)
	}

	// Update flex profile
	config.FlexProfile = flexProfile
	return s.SetSiteTag(ctx, *config)
}

// SetLocalSite sets the local site mode for a site tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - siteTagName: Name of the site tag
//   - enabled: Whether to enable local site mode
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) SetLocalSite(ctx context.Context, siteTagName string, enabled bool) error {
	// Get existing site tag
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("failed to get site tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("site tag %s not found", siteTagName)
	}

	// Update local site setting
	config.IsLocalSite = &enabled
	return s.SetSiteTag(ctx, *config)
}

// SetDescription sets the description for a site tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - siteTagName: Name of the site tag
//   - description: Description to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) SetDescription(ctx context.Context, siteTagName, description string) error {
	// Get existing site tag
	config, err := s.GetSiteTag(ctx, siteTagName)
	if err != nil {
		return fmt.Errorf("failed to get site tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("site tag %s not found", siteTagName)
	}

	// Update description
	config.Description = description
	return s.SetSiteTag(ctx, *config)
}

// ConfigureSiteTag provides a complete configuration interface for site tags.
// This function either creates a new site tag or updates an existing one.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Complete site tag configuration
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) ConfigureSiteTag(ctx context.Context, config SiteTagConfig) error {
	// Try to set first, if fails with 404, create new
	err := s.SetSiteTag(ctx, config)
	if err != nil {
		// Check if error is 404 (not found)
		if core.IsNotFoundError(err) {
			return s.CreateSiteTag(ctx, config)
		}
		return err
	}

	return nil
}
