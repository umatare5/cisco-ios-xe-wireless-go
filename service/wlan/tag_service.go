package wlan

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// PolicyTagService provides Policy Tag management operations.
type PolicyTagService struct {
	service.BaseService
}

// NewPolicyTagService creates a new PolicyTagService instance.
func NewPolicyTagService(c *core.Client) *PolicyTagService {
	return &PolicyTagService{
		BaseService: service.NewBaseService(c),
	}
}

// GetPolicyTag retrieves a specific policy tag configuration.
func (s *PolicyTagService) GetPolicyTag(ctx context.Context, tagName string) (*model.PolicyListEntry, error) {
	if err := s.validateTagName(tagName); err != nil {
		return nil, err
	}

	// Get all policy tags and find the specific one
	allTags, err := s.ListPolicyTags(ctx)
	if err != nil {
		return nil, err
	}

	// Find the tag with the matching name
	for _, tag := range allTags {
		if tag.TagName == tagName {
			return &tag, nil
		}
	}

	// Tag not found
	return nil, nil
}

// ListPolicyTags retrieves all policy tag configurations.
func (s *PolicyTagService) ListPolicyTags(ctx context.Context) ([]model.PolicyListEntry, error) {
	result, err := core.Get[model.WlanCfgPolicyListEntries](ctx, s.Client(), routes.WLANPolicyListEntriesPath)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return []model.PolicyListEntry{}, nil
	}

	if result.PolicyListEntries == nil {
		return []model.PolicyListEntry{}, nil
	}

	if result.PolicyListEntries.PolicyListEntry == nil {
		return []model.PolicyListEntry{}, nil
	}

	return result.PolicyListEntries.PolicyListEntry, nil
}

// CreatePolicyTag creates a new policy tag configuration.
func (s *PolicyTagService) CreatePolicyTag(ctx context.Context, config *model.PolicyListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if config.TagName == "" {
		return errors.New("policy tag name cannot be empty")
	}

	if err := s.validateTagName(config.TagName); err != nil {
		return err
	}

	// Use the model directly without conversion
	payload := s.buildPayload(*config)
	return core.PostVoid(ctx, s.Client(), routes.WLANPolicyListEntriesPath, payload)
}

// SetPolicyTag configures a policy tag on the wireless controller using PUT operation.
func (s *PolicyTagService) SetPolicyTag(ctx context.Context, config *model.PolicyListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if config.TagName == "" {
		return errors.New("policy tag name cannot be empty")
	}

	if err := s.validateTagName(config.TagName); err != nil {
		return err
	}

	// Use the model directly without conversion
	payload := s.buildPayload(*config)
	return core.PatchVoid(ctx, s.Client(), s.buildTagURL(config.TagName), payload)
}

// SetPolicyProfile sets the policy profile for a specific WLAN in a policy tag.
func (s *PolicyTagService) SetPolicyProfile(
	ctx context.Context,
	tagName, wlanProfileName, policyProfileName string,
) error {
	// Get existing policy tag
	config, err := s.GetPolicyTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("policy tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", tagName, err))
	}

	if config == nil {
		return fmt.Errorf("policy tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", tagName))
	}

	// Initialize WLANPolicies if nil
	if config.WLANPolicies == nil {
		config.WLANPolicies = &model.WLANPolicies{
			WLANPolicy: []model.WLANPolicyMap{},
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
		config.WLANPolicies.WLANPolicy = append(config.WLANPolicies.WLANPolicy, model.WLANPolicyMap{
			WLANProfileName:   wlanProfileName,
			PolicyProfileName: policyProfileName,
		})
	}

	return s.SetPolicyTag(ctx, config)
}

// SetDescription sets the description for a policy tag.
func (s *PolicyTagService) SetDescription(ctx context.Context, tagName, description string) error {
	config, err := s.GetPolicyTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("policy tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", tagName, err))
	}

	if config == nil {
		return fmt.Errorf("policy tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", tagName))
	}

	// Update description
	config.Description = description
	return s.SetPolicyTag(ctx, config)
}

// DeletePolicyTag deletes a policy tag configuration.
func (s *PolicyTagService) DeletePolicyTag(ctx context.Context, tagName string) error {
	if err := s.validateTagName(tagName); err != nil {
		return err
	}
	return core.Delete(ctx, s.Client(), s.buildTagURL(tagName))
}

// validateTagName validates policy tag name.
func (s *PolicyTagService) validateTagName(tagName string) error {
	if tagName == "" {
		return errors.New("policy tag name cannot be empty")
	}
	if len(tagName) > 32 {
		return fmt.Errorf("policy tag validation failed: %w",
			fmt.Errorf("tag name length validation failed: name too long (max 32 characters): '%s'", tagName))
	}
	return nil
}

// buildTagURL builds URL for specific tag operations using RESTCONF builder.
func (s *PolicyTagService) buildTagURL(tagName string) string {
	return s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.WLANPolicyListEntriesPath,
		"policy-list-entry",
		tagName,
	)
}

// buildPayload builds a payload for tag operations directly.
func (s *PolicyTagService) buildPayload(config model.PolicyListEntry) map[string]interface{} {
	return map[string]interface{}{
		"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry": config,
	}
}
