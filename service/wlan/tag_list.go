package wlan

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// ListPolicyTags retrieves all policy tag configurations.
//
// **Returns:**
//   - []PolicyTagConfig: List of all policy tags
//   - error: nil on success, error otherwise
func (s *PolicyTagService) ListPolicyTags(ctx context.Context) ([]PolicyTagConfig, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	resp, err := core.Get[model.WlanCfgPolicyListEntries](ctx, s.Client(), routes.PolicyListEntriesEndpoint)
	if validation.HasError(err) {
		return nil, fmt.Errorf("failed to list policy tags: %w", err)
	}

	// Use validation predicate for nil check
	if validation.IsNil(resp) ||
		validation.IsNil(resp.PolicyListEntries) ||
		validation.IsNil(resp.PolicyListEntries.PolicyListEntry) {
		return []PolicyTagConfig{}, nil
	}

	return s.convertToPolicyTagConfigs(resp.PolicyListEntries.PolicyListEntry), nil
}

// convertToPolicyTagConfigs converts model PolicyListEntry to service PolicyTagConfig
func (s *PolicyTagService) convertToPolicyTagConfigs(entries []model.PolicyListEntry) []PolicyTagConfig {
	result := make([]PolicyTagConfig, len(entries))
	for i, entry := range entries {
		// Debug: Check entry structure
		var wlanPolicies *WLANPolicies
		if entry.WLANPolicies != nil {
			wlanPolicies = convertWLANPolicies(entry.WLANPolicies)
		}

		result[i] = PolicyTagConfig{
			TagName:      entry.TagName,
			Description:  entry.Description,
			WLANPolicies: wlanPolicies,
		}
	}
	return result
}

// convertWLANPolicies converts model WLANPolicies to service WLANPolicies
func convertWLANPolicies(modelPolicies *model.WLANPolicies) *WLANPolicies {
	if validation.IsNil(modelPolicies) {
		return nil
	}

	if len(modelPolicies.WLANPolicy) == 0 {
		return nil
	}

	servicePolicies := &WLANPolicies{
		WLANPolicy: make([]WLANPolicyMap, len(modelPolicies.WLANPolicy)),
	}

	for i, policy := range modelPolicies.WLANPolicy {
		servicePolicies.WLANPolicy[i] = WLANPolicyMap{
			WLANProfileName:   policy.WLANProfileName,
			PolicyProfileName: policy.PolicyProfileName,
		}
	}

	return servicePolicies
}
