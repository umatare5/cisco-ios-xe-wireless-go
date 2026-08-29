package wlan

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
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
func (s *PolicyTagService) GetPolicyTag(
	ctx context.Context,
	tagName string,
	opts ...core.GetOption,
) (*PolicyListEntry, error) {
	if err := validation.ValidateTagName(tagName); err != nil {
		return nil, err
	}

	// Read the keyed URL rather than filtering the whole container client-side. Filtering would
	// make a GetOption that prunes tag-name, or one that cuts the list depth, report an existing
	// tag as absent with a nil error.
	result, err := core.Get[CiscoIOSXEWirelessWlanCfgPolicyListEntry](
		ctx, s.Client(), s.buildTagURL(tagName), opts...,
	)
	if err != nil {
		return nil, err
	}
	if result == nil || len(result.PolicyListEntry) == 0 {
		return nil, nil
	}

	return &result.PolicyListEntry[0], nil
}

// ListPolicyTags retrieves all policy tag configurations.
func (s *PolicyTagService) ListPolicyTags(
	ctx context.Context,
	opts ...core.GetOption,
) ([]PolicyListEntry, error) {
	wlanService := NewService(s.Client())

	result, err := wlanService.ListCfgPolicyListEntries(ctx, opts...)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return []PolicyListEntry{}, nil
	}
	if result.PolicyListEntries == nil {
		return []PolicyListEntry{}, nil
	}
	if result.PolicyListEntries.PolicyListEntry == nil {
		return []PolicyListEntry{}, nil
	}

	return result.PolicyListEntries.PolicyListEntry, nil
}

// CreatePolicyTag creates a new policy tag configuration.
func (s *PolicyTagService) CreatePolicyTag(ctx context.Context, config *PolicyListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if err := validation.ValidateTagName(config.TagName); err != nil {
		return err
	}

	payload := s.buildPayload(*config)
	return core.PostVoid(ctx, s.Client(), routes.WLANPolicyListEntriesPath, payload)
}

// SetPolicyTag applies a policy tag with a merge PATCH, so a leaf the payload omits keeps the
// value the controller already holds instead of being cleared.
func (s *PolicyTagService) SetPolicyTag(ctx context.Context, config *PolicyListEntry) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	if err := validation.ValidateTagName(config.TagName); err != nil {
		return err
	}

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

	// Take the address rather than the value: a merge PATCH omits an empty string, so assigning
	// one through a value field cannot clear a description that is already set.
	config.Description = &description
	return s.SetPolicyTag(ctx, config)
}

// DeletePolicyTag deletes a policy tag configuration.
func (s *PolicyTagService) DeletePolicyTag(ctx context.Context, tagName string) error {
	if err := validation.ValidateTagName(tagName); err != nil {
		return err
	}
	return core.Delete(ctx, s.Client(), s.buildTagURL(tagName))
}

// buildTagURL builds URL for specific tag operations using RESTCONF builder.
func (s *PolicyTagService) buildTagURL(tagName string) string {
	return s.Client().RESTCONFBuilder().BuildQueryURL(
		routes.WLANPolicyListEntryQueryPath,
		tagName,
	)
}

// buildPayload builds a payload for tag operations using the request.
func (s *PolicyTagService) buildPayload(config PolicyListEntry) CiscoIOSXEWirelessWlanPolicyListEntriesPayload {
	return CiscoIOSXEWirelessWlanPolicyListEntriesPayload{
		PolicyListEntry: config,
	}
}
