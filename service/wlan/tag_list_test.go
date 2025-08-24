package wlan

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestPolicyTagService_ListPolicyTags_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "PolicyTagService",
		Operation:   "ListPolicyTags",
		NilClientTest: func() error {
			service := NewPolicyTagService(nil)
			_, err := service.ListPolicyTags(context.Background())
			return err
		},
		EmptyParamTest: func() error {
			service := &PolicyTagService{}
			_, err := service.ListPolicyTags(context.Background())
			return err
		},
	})
}

func TestPolicyTagService_ListPolicyTags_StatusCode(t *testing.T) {
	// Test data
	testCases := []struct {
		name        string
		statusCode  int
		expectError bool
	}{
		{"Success_200", 200, false},
		{"Unauthorized_401", 401, true},
		{"Forbidden_403", 403, true},
		{"NotFound_404", 404, true},
		{"InternalError_500", 500, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock server
			server := mock.NewRESTCONFErrorServer([]string{routes.PolicyListEntriesEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewPolicyTagService(testClient)

			ctx := context.Background()
			_, err := service.ListPolicyTags(ctx)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestPolicyTagService_ListPolicyTags_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)
	ctx := context.Background()

	// Test listing policy tags
	tags, err := service.ListPolicyTags(ctx)
	if err != nil {
		t.Errorf("Failed to list policy tags: %v", err)
		return
	}

	// Basic validation
	if tags == nil {
		t.Error("Expected non-nil response for policy tags list")
		return
	}

	tagCount := len(tags)
	t.Logf("Successfully listed policy tags (count: %d)", tagCount)

	// Validate that we have at least some known tags
	if tagCount == 0 {
		t.Error("Expected at least some policy tags in the system")
		return
	}

	// Check for known tags from the live system
	expectedTags := map[string]bool{
		"labo-wlan-flex":     false,
		"test-create-error":  false,
		"default-policy-tag": false,
	}

	for _, tag := range tags {
		if tag.TagName == "" {
			t.Error("Found policy tag with empty name")
		}

		// Mark expected tags as found
		if _, exists := expectedTags[tag.TagName]; exists {
			expectedTags[tag.TagName] = true
		}

		// Log details for verification
		t.Logf("  Tag: %s", tag.TagName)
		if tag.Description != "" {
			t.Logf("    Description: %s", tag.Description)
		}
		if tag.WLANPolicies != nil && len(tag.WLANPolicies.WLANPolicy) > 0 {
			t.Logf("    WLAN Policies: %d", len(tag.WLANPolicies.WLANPolicy))
			for _, policy := range tag.WLANPolicies.WLANPolicy {
				t.Logf("      WLAN: %s -> Policy: %s", policy.WLANProfileName, policy.PolicyProfileName)
			}
		}
	}

	// Check if we found expected tags
	for tagName, found := range expectedTags {
		if !found {
			t.Logf("Expected tag '%s' not found in list", tagName)
		}
	}
}
