package wlan

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
)

// SetPolicyTag configures a policy tag on the wireless controller using PUT operation.
// This method is intended for AP usage and will overwrite existing configuration.
//
// Parameters:
//   - ctx: Context for request lifecycle management
//   - config: Policy tag configuration including TagName and settings
//
// Returns:
//   - error: nil on success, detailed error on failure
//
// YANG Path: /Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry
func (s *PolicyTagService) SetPolicyTag(ctx context.Context, config *PolicyTagConfig) error {
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

	return s.tagOps.Update(ctx, internalConfig, config.TagName)
}

// SetPolicyProfile sets the policy profile for a specific WLAN in a policy tag.
//
// Parameters:
//   - ctx: Context for the request
//   - tagName: Name of the policy tag
//   - wlanProfileName: WLAN profile name
//   - policyProfileName: Policy profile name to set
//
// Returns:
//   - error: nil on success, error otherwise
func (s *PolicyTagService) SetPolicyProfile(
	ctx context.Context,
	tagName, wlanProfileName, policyProfileName string,
) error {
	// Get existing policy tag
	config, err := s.GetPolicyTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get policy tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("policy tag %s not found", tagName)
	}

	// Initialize WLANPolicies if nil
	if config.WLANPolicies == nil {
		config.WLANPolicies = &WLANPolicies{
			WLANPolicy: []WLANPolicyMap{},
		}
	}

	// Find existing WLAN policy or add new one
	found := false
	for i := range config.WLANPolicies.WLANPolicy {
		if config.WLANPolicies.WLANPolicy[i].WLANProfileName == wlanProfileName {
			config.WLANPolicies.WLANPolicy[i].PolicyProfileName = policyProfileName
			found = true
			break
		}
	}

	if !found {
		config.WLANPolicies.WLANPolicy = append(config.WLANPolicies.WLANPolicy, WLANPolicyMap{
			WLANProfileName:   wlanProfileName,
			PolicyProfileName: policyProfileName,
		})
	}

	return s.SetPolicyTag(ctx, config)
}

// SetDescription sets the description for a policy tag.
//
// Parameters:
//   - ctx: Context for the request
//   - tagName: Name of the policy tag
//   - description: Description to set
//
// Returns:
//   - error: nil on success, error otherwise
func (s *PolicyTagService) SetDescription(ctx context.Context, tagName, description string) error {
	// Get existing policy tag
	config, err := s.GetPolicyTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get policy tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("policy tag %s not found", tagName)
	}

	// Update description
	config.Description = description
	return s.SetPolicyTag(ctx, config)
}

// ConfigurePolicyTag provides a complete configuration interface for policy tags.
// This method either creates a new policy tag or updates an existing one.
//
// Parameters:
//   - ctx: Context for the request
//   - config: Complete policy tag configuration
//
// Returns:
//   - error: nil on success, error otherwise
func (s *PolicyTagService) ConfigurePolicyTag(ctx context.Context, config *PolicyTagConfig) error {
	// Try to set first, if fails with 404, create new
	err := s.SetPolicyTag(ctx, config)
	if err != nil {
		// Check if error is 404 (not found)
		if core.IsNotFoundError(err) {
			return s.CreatePolicyTag(ctx, config)
		}
		return err
	}

	return nil
}

// convertToModelWLANPolicies converts service WLANPolicies to model WLANPolicies
func convertToModelWLANPolicies(servicePolicies *WLANPolicies) *model.WLANPolicies {
	if servicePolicies == nil {
		return nil
	}

	modelPolicies := &model.WLANPolicies{
		WLANPolicy: make([]model.WLANPolicyMap, len(servicePolicies.WLANPolicy)),
	}

	for i, policy := range servicePolicies.WLANPolicy {
		modelPolicies.WLANPolicy[i] = model.WLANPolicyMap{
			WLANProfileName:   policy.WLANProfileName,
			PolicyProfileName: policy.PolicyProfileName,
		}
	}

	return modelPolicies
}
