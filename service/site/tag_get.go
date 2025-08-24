package site

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// GetSiteTag retrieves a specific site tag configuration.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the site tag to retrieve
//
// **Returns:**
//   - *SiteTagConfig: Site tag configuration
//   - error: nil on success, error otherwise
func (s *SiteTagService) GetSiteTag(ctx context.Context, tagName string) (*SiteTagConfig, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	// Use separate Get and Parse operations for better separation of concerns
	body, err := s.tagOps.Get(ctx, tagName)
	if validation.HasError(err) {
		return nil, err
	}

	if validation.IsNil(body) {
		return nil, nil
	}

	// Parse the response
	result, err := s.tagOps.Parse(body, tagName, helpers.ParseSiteTagJSONResponse)
	if validation.HasError(err) {
		return nil, err
	}

	if validation.IsNil(result) {
		return nil, nil
	}

	return s.convertToConfig(result)
}

// convertToConfig converts parsed result to SiteTagConfig using validation predicates
func (s *SiteTagService) convertToConfig(result any) (*SiteTagConfig, error) {
	jsonBytes, err := json.Marshal(result)
	if validation.HasError(err) {
		return nil, fmt.Errorf("failed to marshal site tag result: %w", err)
	}

	var config SiteTagConfig
	if err := json.Unmarshal(jsonBytes, &config); validation.HasError(err) {
		return nil, fmt.Errorf("failed to unmarshal site tag config: %w", err)
	}

	return &config, nil
}
