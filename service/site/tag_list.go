package site

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// ListSiteTags retrieves all site tag configurations.
//
// **Returns:**
//   - []SiteTagConfig: List of all site tags
//   - error: nil on success, error otherwise
func (s *SiteTagService) ListSiteTags(ctx context.Context) ([]SiteTagConfig, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	resp, err := core.Get[model.SiteCfgSiteTagConfigs](ctx, s.Client(), routes.SiteTagConfigsEndpoint)
	if validation.HasError(err) {
		return nil, fmt.Errorf("failed to list site tags: %w", err)
	}

	// Use validation predicate for nil check
	if validation.IsNil(resp) {
		return []SiteTagConfig{}, nil
	}

	return s.convertSiteTagConfigsToServiceModel(resp.SiteTagConfigs.SiteTagConfig), nil
}

// convertSiteTagConfigsToServiceModel converts model SiteTagConfig to service SiteTagConfig
func (s *SiteTagService) convertSiteTagConfigsToServiceModel(configs []model.SiteTagConfig) []SiteTagConfig {
	result := make([]SiteTagConfig, len(configs))
	for i, config := range configs {
		// Convert bool to *bool for IsLocalSite field
		isLocalSite := config.IsLocalSite
		result[i] = SiteTagConfig{
			SiteTagName:   config.SiteTagName,
			Description:   config.Description,
			FlexProfile:   config.FlexProfile,
			APJoinProfile: config.ApJoinProfile,
			IsLocalSite:   &isLocalSite,
		}
	}
	return result
}
