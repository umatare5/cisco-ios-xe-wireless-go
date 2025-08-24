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

func TestPolicyTagService_DeletePolicyTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "PolicyTagService",
		Operation:   "DeletePolicyTag",
		NilClientTest: func() error {
			service := NewPolicyTagService(nil)
			return service.DeletePolicyTag(context.Background(), "test-policy-tag")
		},
		EmptyParamTest: func() error {
			service := &PolicyTagService{}
			return service.DeletePolicyTag(context.Background(), "") // Empty tag name
		},
	})
}

func TestPolicyTagService_DeletePolicyTag_StatusCode(t *testing.T) {
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
			err := service.DeletePolicyTag(ctx, tagName)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestPolicyTagService_DeletePolicyTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewPolicyTagService(testClient)

	// Use timestamp to ensure unique tag name with required format for this test
	tagName := fmt.Sprintf("test-policy-tag-%s", time.Now().Format("20060102-150405"))

	config := &PolicyTagConfig{
		TagName:     tagName,
		Description: "Test policy tag for deletion",
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

	// First create a tag to delete
	err := service.CreatePolicyTag(ctx, config)
	if err != nil {
		t.Fatalf("Failed to create policy tag for deletion test: %v", err)
	}

	t.Logf("Successfully created policy tag for deletion: %s", tagName)

	// Test deletion
	err = service.DeletePolicyTag(ctx, tagName)
	if err != nil {
		t.Errorf("Failed to delete policy tag: %v", err)
	} else {
		t.Logf("Successfully deleted policy tag: %s", tagName)
	}
}
