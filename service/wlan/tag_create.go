package wlan

import (
	"context"
	"errors"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
)

// CreatePolicyTag creates a new policy tag configuration.
// This function creates a policy tag that can be associated with Access Points.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Policy tag configuration to create
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *PolicyTagService) CreatePolicyTag(ctx context.Context, config *PolicyTagConfig) error {
	// Validate input parameters first (independent of client)
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if config.TagName == "" {
		return errors.New("policy tag name cannot be empty")
	}

	// Validate client after input validation
	if err := s.ValidateClient(); err != nil {
		return err
	}

	// Convert service model to internal model
	internalConfig := model.PolicyListEntry{
		TagName:      config.TagName,
		Description:  config.Description,
		WLANPolicies: convertToModelWLANPolicies(config.WLANPolicies),
	}

	return s.tagOps.Create(ctx, internalConfig, config.TagName)
}
