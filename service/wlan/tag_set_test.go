package wlan

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestPolicyTagService_SetPolicyTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "PolicyTagService",
		Operation:   "SetPolicyTag",
		NilClientTest: func() error {
			service := NewPolicyTagService(nil)
			config := &PolicyTagConfig{TagName: "test-policy-tag"}
			return service.SetPolicyTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &PolicyTagService{}
			config := &PolicyTagConfig{} // Empty tag name
			return service.SetPolicyTag(context.Background(), config)
		},
	})
}

func TestPolicyTagService_SetPolicyTag_StatusCode(t *testing.T) {
	// Test data
	testCases := []struct {
		name        string
		statusCode  int
		expectError bool
	}{
		{"Success_204", 204, false},
		{"Success_200", 200, false},
		{"BadRequest_400", 400, true},
		{"Unauthorized_401", 401, true},
		{"Forbidden_403", 403, true},
		{"NotFound_404", 404, true},
		{"Conflict_409", 409, true},
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
			config := &PolicyTagConfig{
				TagName:     tagName,
				Description: "Test policy tag updated via status code test",
				WLANPolicies: &WLANPolicies{
					WLANPolicy: []WLANPolicyMap{
						{
							WLANProfileName:   "test-wlan",
							PolicyProfileName: "labo-wlan-profile",
						},
					},
				},
			}

			ctx := context.Background()
			err := service.SetPolicyTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestPolicyTagService_SetPolicyTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)

	// Use timestamp to ensure unique tag name with required format for this test
	tagName := fmt.Sprintf("test-policy-tag-%s", time.Now().Format("20060102-150405"))

	// Create initial tag
	initialConfig := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Initial policy tag for set operation",
		WLANPolicies: &WLANPolicies{
			WLANPolicy: []WLANPolicyMap{
				{
					WLANProfileName:   "test-wlan",
					PolicyProfileName: "labo-wlan-profile",
				},
			},
		},
	}

	ctx := context.Background()

	// First create a tag to update
	err := service.CreatePolicyTag(ctx, initialConfig)
	if err != nil {
		t.Fatalf("Failed to create policy tag for set test: %v", err)
	}

	t.Logf("Successfully created policy tag for set operation: %s", tagName)

	// Cleanup - attempt to delete the test tag
	defer func() {
		if deleteErr := service.DeletePolicyTag(ctx, tagName); deleteErr != nil {
			t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
		}
	}()

	// Update the tag with new configuration
	updatedConfig := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Updated policy tag for testing",
		WLANPolicies: &WLANPolicies{
			WLANPolicy: []WLANPolicyMap{
				{
					WLANProfileName:   "test-wlan",
					PolicyProfileName: "labo-wlan-profile",
				},
			},
		},
	}

	// Test setting/updating
	err = service.SetPolicyTag(ctx, updatedConfig)
	if err != nil {
		t.Errorf("Failed to set/update policy tag: %v", err)
	} else {
		t.Logf("Successfully updated policy tag: %s", tagName)
	}
}

func TestPolicyTagService_SetPolicyProfile_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)

	// Use timestamp to ensure unique tag name with required format for this test
	tagName := fmt.Sprintf("test-policy-tag-%s", time.Now().Format("20060102-150405"))

	// Create initial tag without policies
	initialConfig := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Policy tag for SetPolicyProfile test",
	}

	ctx := context.Background()

	// First create a tag
	err := service.CreatePolicyTag(ctx, initialConfig)
	if err != nil {
		t.Fatalf("Failed to create policy tag for SetPolicyProfile test: %v", err)
	}

	t.Logf("Successfully created policy tag for SetPolicyProfile test: %s", tagName)

	// Cleanup - attempt to delete the test tag
	defer func() {
		if deleteErr := service.DeletePolicyTag(ctx, tagName); deleteErr != nil {
			t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
		}
	}()

	// Test setting policy profile
	err = service.SetPolicyProfile(ctx, tagName, "test-wlan", "labo-wlan-profile")
	if err != nil {
		t.Errorf("Failed to set policy profile: %v", err)
	} else {
		t.Logf("Successfully set policy profile for tag: %s", tagName)
	}

	// Verify the policy was set by getting the tag
	updatedTag, err := service.GetPolicyTag(ctx, tagName)
	if err != nil {
		t.Errorf("Failed to get updated policy tag: %v", err)
	} else if updatedTag != nil {
		if updatedTag.WLANPolicies != nil && len(updatedTag.WLANPolicies.WLANPolicy) > 0 {
			found := false
			for _, policy := range updatedTag.WLANPolicies.WLANPolicy {
				if policy.WLANProfileName == "test-wlan" && policy.PolicyProfileName == "labo-wlan-profile" {
					found = true
					break
				}
			}
			if !found {
				t.Error("Policy profile was not properly set")
			} else {
				t.Log("Policy profile was successfully verified")
			}
		} else {
			t.Error("Expected policy to be set, but WLANPolicies is empty")
		}
	}
}

func TestPolicyTagService_SetDescription_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)

	// Use timestamp to ensure unique tag name with required format for this test
	tagName := fmt.Sprintf("test-policy-tag-%s", time.Now().Format("20060102-150405"))

	// Create initial tag
	initialConfig := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Initial description",
	}

	ctx := context.Background()

	// First create a tag
	err := service.CreatePolicyTag(ctx, initialConfig)
	if err != nil {
		t.Fatalf("Failed to create policy tag for SetDescription test: %v", err)
	}

	t.Logf("Successfully created policy tag for SetDescription test: %s", tagName)

	// Cleanup - attempt to delete the test tag
	defer func() {
		if deleteErr := service.DeletePolicyTag(ctx, tagName); deleteErr != nil {
			t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
		}
	}()

	// Test setting description
	newDescription := "Updated description"
	err = service.SetDescription(ctx, tagName, newDescription)
	if err != nil {
		t.Errorf("Failed to set description: %v", err)
	} else {
		t.Logf("Successfully set description for tag: %s", tagName)
	}

	// Verify the description was set by getting the tag
	updatedTag, err := service.GetPolicyTag(ctx, tagName)
	if err != nil {
		t.Errorf("Failed to get updated policy tag: %v", err)
	} else if updatedTag != nil {
		if updatedTag.Description != newDescription {
			t.Errorf("Expected description '%s', got '%s'", newDescription, updatedTag.Description)
		} else {
			t.Log("Description was successfully verified")
		}
	}
}
