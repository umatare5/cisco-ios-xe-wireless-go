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

func TestPolicyTagService_CreatePolicyTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "PolicyTagService",
		Operation:   "CreatePolicyTag",
		NilClientTest: func() error {
			service := NewPolicyTagService(nil)
			config := &PolicyTagConfig{TagName: "test-policy-tag"}
			return service.CreatePolicyTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &PolicyTagService{}
			config := &PolicyTagConfig{} // Empty tag name
			return service.CreatePolicyTag(context.Background(), config)
		},
	})
}

func TestPolicyTagService_CreatePolicyTag_StatusCode(t *testing.T) {
	// Test data
	testCases := []struct {
		name        string
		statusCode  int
		expectError bool
	}{
		{"Success_201", 201, false},
		{"Success_204", 204, false},
		{"BadRequest_400", 400, true},
		{"Unauthorized_401", 401, true},
		{"Forbidden_403", 403, true},
		{"NotFound_404", 404, true},
		{"Conflict_409", 409, true},
		{"InternalError_500", 500, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock server
			server := mock.NewRESTCONFErrorServer([]string{routes.PolicyListEntriesEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewPolicyTagService(testClient)
			config := &PolicyTagConfig{
				TagName:     "test-policy-tag",
				Description: "Test policy tag created via status code test",
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
			err := service.CreatePolicyTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestPolicyTagService_CreatePolicyTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)

	// Use timestamp to ensure unique tag name with required format for this test
	tagName := fmt.Sprintf("test-policy-tag-%s", time.Now().Format("20060102-150405"))

	config := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Integration test policy tag",
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

	// Test creation
	err := service.CreatePolicyTag(ctx, config)
	if err != nil {
		t.Errorf("Failed to create policy tag: %v", err)
		return
	}

	t.Logf("Successfully created policy tag: %s", tagName)

	// Cleanup - attempt to delete the test tag
	if deleteErr := service.DeletePolicyTag(ctx, tagName); deleteErr != nil {
		t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
	}
}
