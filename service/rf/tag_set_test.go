package rf

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

func TestRFTagService_SetRFTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "RFTagService",
		Operation:   "SetRFTag",
		NilClientTest: func() error {
			service := NewRFTagService(nil)
			config := RFTagConfig{TagName: "test-rf-tag"}
			return service.SetRFTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &RFTagService{}
			config := RFTagConfig{} // Empty tag name
			return service.SetRFTag(context.Background(), config)
		},
	})
}

func TestRFTagService_SetRFTag_StatusCode(t *testing.T) {
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
			tagName := "test-rf-tag"
			endpoint := fmt.Sprintf("%s=%s", routes.RfTagsEndpoint, tagName)
			server := mock.NewRESTCONFErrorServer([]string{endpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewRFTagService(testClient)
			config := RFTagConfig{
				TagName:             tagName,
				Description:         "Test RF tag updated via status code test",
				Dot11ARfProfileName: "updated-11a-profile",
				Dot11BRfProfileName: "updated-11b-profile",
			}

			ctx := context.Background()
			err := service.SetRFTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestRFTagService_SetRFTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewRFTagService(testClient)

	// Use timestamp to ensure unique tag name
	tagName := fmt.Sprintf("test-rf-set-%d", time.Now().Unix())

	// Create initial tag
	initialConfig := RFTagConfig{
		TagName:             tagName,
		Description:         "Test RF tag for set operation", // 64 chars limit
		Dot11ARfProfileName: "default-11a-profile",
		Dot11BRfProfileName: "default-11b-profile",
	}

	ctx := context.Background()

	// First create a tag to update
	err := service.CreateRFTag(ctx, initialConfig)
	if err != nil {
		t.Fatalf("Failed to create RF tag for set test: %v", err)
	}

	// Cleanup - attempt to delete the test tag
	defer func() {
		if deleteErr := service.DeleteRFTag(ctx, tagName); deleteErr != nil {
			t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
		}
	}()

	// Update the tag with new configuration
	updatedConfig := RFTagConfig{
		TagName:             tagName,
		Description:         "Updated RF tag for testing", // 64 chars limit
		Dot11ARfProfileName: "updated-11a-profile",
		Dot11BRfProfileName: "updated-11b-profile",
	}

	// Test setting/updating
	err = service.SetRFTag(ctx, updatedConfig)
	if err != nil {
		t.Errorf("Failed to set/update RF tag: %v", err)
	}
}
