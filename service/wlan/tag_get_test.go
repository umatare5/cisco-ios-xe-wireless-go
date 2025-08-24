package wlan

import (
	"context"
	"fmt"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestPolicyTagService_GetPolicyTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "PolicyTagService",
		Operation:   "GetPolicyTag",
		NilClientTest: func() error {
			service := NewPolicyTagService(nil)
			_, err := service.GetPolicyTag(context.Background(), "test-policy-tag")
			return err
		},
		EmptyParamTest: func() error {
			service := &PolicyTagService{}
			_, err := service.GetPolicyTag(context.Background(), "") // Empty tag name
			return err
		},
	})
}

func TestPolicyTagService_GetPolicyTag_StatusCode(t *testing.T) {
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
			// Create mock server with tag-specific endpoint
			tagName := "test-policy-tag"
			endpoint := fmt.Sprintf("%s/policy-list-entry=%s", routes.PolicyListEntriesEndpoint, tagName)
			server := mock.NewRESTCONFErrorServer([]string{endpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewPolicyTagService(testClient)

			ctx := context.Background()
			_, err := service.GetPolicyTag(ctx, tagName)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestPolicyTagService_GetPolicyTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)
	ctx := context.Background()

	// Test getting existing policy tags
	// Note: This test will use existing policy tags from the system
	testCases := []string{
		"labo-wlan-flex",     // Known existing tag from live system
		"default-policy-tag", // Known existing tag from live system
	}

	for _, tagName := range testCases {
		t.Run(fmt.Sprintf("GetExisting_%s", tagName), func(t *testing.T) {
			tag, err := service.GetPolicyTag(ctx, tagName)
			if err != nil {
				t.Errorf("Failed to get policy tag '%s': %v", tagName, err)
				return
			}

			// Basic validation
			if tag == nil {
				t.Error("Expected non-nil response for policy tag get")
				return
			}

			if tag.TagName != tagName {
				t.Errorf("Expected tag name %s, got %s", tagName, tag.TagName)
			}

			t.Logf("Successfully retrieved policy tag: %s", tagName)
			if tag.Description != "" {
				t.Logf("  Description: %s", tag.Description)
			}
			if tag.WLANPolicies != nil && len(tag.WLANPolicies.WLANPolicy) > 0 {
				t.Logf("  WLAN Policies count: %d", len(tag.WLANPolicies.WLANPolicy))
			}
		})
	}

	// Test getting non-existent tag
	t.Run("GetNonExistent", func(t *testing.T) {
		nonExistentTag := "non-existent-policy-tag-test"
		tag, err := service.GetPolicyTag(ctx, nonExistentTag)
		switch {
		case err != nil:
			// This is expected behavior for non-existent tags
			t.Logf("Expected error for non-existent tag: %v", err)
		case tag == nil:
			t.Logf("Non-existent tag returned nil (acceptable)")
		default:
			t.Errorf("Unexpected result for non-existent tag: %+v", tag)
		}
	})
}
