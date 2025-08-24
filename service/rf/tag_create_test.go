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

func TestRFTagService_CreateRFTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "RFTagService",
		Operation:   "CreateRFTag",
		NilClientTest: func() error {
			service := NewRFTagService(nil)
			config := RFTagConfig{TagName: "test-rf-tag"}
			return service.CreateRFTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &RFTagService{}
			config := RFTagConfig{} // Empty tag name
			return service.CreateRFTag(context.Background(), config)
		},
	})
}

func TestRFTagService_CreateRFTag_StatusCode(t *testing.T) {
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
			server := mock.NewRESTCONFErrorServer([]string{routes.RfTagsEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewRFTagService(testClient)
			config := RFTagConfig{
				TagName:             "test-rf-tag",
				Description:         "Test RF tag created via status code test",
				Dot11ARfProfileName: "default-11a-profile",
				Dot11BRfProfileName: "default-11b-profile",
			}

			ctx := context.Background()
			err := service.CreateRFTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestRFTagService_CreateRFTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewRFTagService(testClient)

	// Use timestamp to ensure unique tag name
	tagName := fmt.Sprintf("test-rf-create-%d", time.Now().Unix())

	config := RFTagConfig{
		TagName: tagName,
		Description: fmt.Sprintf(
			"Integration test RF tag created at %s",
			time.Now().Format("2006-01-02 15:04:05"),
		),
		Dot11ARfProfileName: "default-11a-profile",
		Dot11BRfProfileName: "default-11b-profile",
	}

	ctx := context.Background()

	// Test creation
	err := service.CreateRFTag(ctx, config)
	if err != nil {
		t.Errorf("Failed to create RF tag: %v", err)
		return
	}

	// Cleanup - attempt to delete the test tag
	if deleteErr := service.DeleteRFTag(ctx, tagName); deleteErr != nil {
		t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
	}
}
