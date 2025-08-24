package site

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// GetCfg retrieves site configuration data including AP configuration profiles and site tag configurations.
//
// This method returns all site configuration information including AP configuration
// profiles, site tag configurations, packet capture profiles, and other site-related
// settings. The response contains the full site configuration hierarchy as defined
// in the Cisco-IOS-XE-wireless-site-cfg YANG model.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.SiteCfg: Complete site configuration data structure
//   - error: Any error encountered during the operation
//
// Example:
//
//	cfg, err := siteService.GetCfg(ctx)
//	if err != nil {
//		log.Printf("Failed to get site configuration: %v", err)
//		return err
//	}
//
//	// Access AP configuration profiles
//	for _, profile := range cfg.SiteCfgData.APCfgProfiles.APCfgProfile {
//		fmt.Printf("Profile: %s\n", profile.ProfileName)
//	}
func (s Service) GetCfg(ctx context.Context) (*model.SiteCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.SiteCfg](ctx, s.Client(), routes.SiteCfgBasePath)
}

// GetAPCfgProfiles retrieves all AP configuration profiles from the site configuration.
//
// This method provides access to the complete set of AP configuration profiles
// defined within the site configuration. Each profile contains settings for AP
// behavior, security, management, and operational parameters.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.SiteCfgApCfgProfiles: All AP configuration profiles data
//   - error: Any error encountered during the operation
//
// Example:
//
//	profiles, err := siteService.GetAPCfgProfiles(ctx)
//	if err != nil {
//		log.Printf("Failed to get AP config profiles: %v", err)
//		return err
//	}
//
//	for _, profile := range profiles.ApCfgProfiles.ApCfgProfile {
//		fmt.Printf("AP Profile: %s, SSH: %v\n", profile.ProfileName, profile.DeviceMgmt.SSH)
//	}
func (s Service) GetAPCfgProfiles(ctx context.Context) (*model.SiteCfgApCfgProfiles, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.SiteCfgApCfgProfiles](ctx, s.Client(), routes.APCfgProfilesEndpoint)
}

// GetSiteTagConfigs retrieves all site tag configurations from the site configuration.
//
// This method returns complete site tag configuration data including site tag names,
// associated flex profiles, AP join profiles, and local site settings. Site tags
// are used to group and manage access points based on physical location or
// administrative boundaries.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.SiteCfgSiteTagConfigs: All site tag configurations data
//   - error: Any error encountered during the operation
//
// Example:
//
//	tags, err := siteService.GetSiteTagConfigs(ctx)
//	if err != nil {
//		log.Printf("Failed to get site tag configs: %v", err)
//		return err
//	}
//
//	for _, tag := range tags.SiteTagConfigs.SiteTagConfig {
//		fmt.Printf("Site Tag: %s, Local: %v\n", tag.SiteTagName, tag.IsLocalSite)
//	}
func (s Service) GetSiteTagConfigs(ctx context.Context) (*model.SiteCfgSiteTagConfigs, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.SiteCfgSiteTagConfigs](ctx, s.Client(), routes.SiteTagConfigsEndpoint)
}

// GetAPCfgProfileByName retrieves a specific AP configuration profile by name.
//
// This method returns detailed configuration information for a single AP profile
// identified by its profile name. The response includes all settings and parameters
// defined for that specific profile.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - profileName: Name of the AP configuration profile to retrieve
//
// Returns:
//   - *model.SiteCfgApCfgProfile: Specific AP configuration profile data
//   - error: Any error encountered during the operation, including profile not found
//
// Example:
//
//	profile, err := siteService.GetAPCfgProfileByName(ctx, "labo-common")
//	if err != nil {
//		log.Printf("Failed to get AP profile: %v", err)
//		return err
//	}
//
//	fmt.Printf("Profile: %s, SSH enabled: %v\n", profile.ApCfgProfile[0].ProfileName, profile.ApCfgProfile[0].DeviceMgmt.SSH)
func (s Service) GetAPCfgProfileByName(
	ctx context.Context,
	profileName string,
) (*model.SiteCfgApCfgProfile, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/ap-cfg-profile=%s", routes.APCfgProfilesEndpoint, profileName)
	return core.Get[model.SiteCfgApCfgProfile](ctx, s.Client(), url)
}

// GetSiteTagConfigByName retrieves a specific site tag configuration by name.
//
// This method returns detailed configuration information for a single site tag
// identified by its tag name. The response includes all settings and associations
// defined for that specific site tag.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - siteTagName: Name of the site tag configuration to retrieve
//
// Returns:
//   - *model.SiteCfgSiteTagConfig: Specific site tag configuration data
//   - error: Any error encountered during the operation, including tag not found
//
// Example:
//
//	tag, err := siteService.GetSiteTagConfigByName(ctx, "labo-site-flex")
//	if err != nil {
//		log.Printf("Failed to get site tag: %v", err)
//		return err
//	}
//
//	fmt.Printf("Tag: %s, Flex Profile: %s\n", tag.SiteTagConfig[0].SiteTagName, tag.SiteTagConfig[0].FlexProfile)
func (s Service) GetSiteTagConfigByName(
	ctx context.Context,
	siteTagName string,
) (*model.SiteCfgSiteTagConfig, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/site-tag-config=%s", routes.SiteTagConfigsEndpoint, siteTagName)
	return core.Get[model.SiteCfgSiteTagConfig](ctx, s.Client(), url)
}
