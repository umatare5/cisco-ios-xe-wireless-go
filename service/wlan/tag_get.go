package wlan

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// GetPolicyTag retrieves a specific policy tag configuration.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the policy tag to retrieve
//
// **Returns:**
//   - *PolicyTagConfig: Policy tag configuration
//   - error: nil on success, error otherwise
func (s *PolicyTagService) GetPolicyTag(ctx context.Context, tagName string) (*PolicyTagConfig, error) {
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
	result, err := s.tagOps.Parse(body, tagName, helpers.ParsePolicyTagJSONResponse)
	if validation.HasError(err) {
		return nil, err
	}

	if validation.IsNil(result) {
		return nil, nil
	}

	return s.convertToConfig(result)
}

// convertToConfig converts parsed result to PolicyTagConfig using validation predicates
func (s *PolicyTagService) convertToConfig(result any) (*PolicyTagConfig, error) {
	jsonBytes, err := json.Marshal(result)
	if validation.HasError(err) {
		return nil, fmt.Errorf("failed to marshal policy tag result: %w", err)
	}

	var config PolicyTagConfig
	if err := json.Unmarshal(jsonBytes, &config); validation.HasError(err) {
		return nil, fmt.Errorf("failed to unmarshal policy tag config: %w", err)
	}

	return &config, nil
}
