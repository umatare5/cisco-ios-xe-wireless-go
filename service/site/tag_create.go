package site

import (
	"context"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
)

// CreateSiteTag creates a new site tag configuration.
// This function creates a site tag that can be associated with Access Points.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Site tag configuration to create
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) CreateSiteTag(ctx context.Context, config SiteTagConfig) error {
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

	return s.tagOps.Create(ctx, internalConfig, config.SiteTagName)
}
