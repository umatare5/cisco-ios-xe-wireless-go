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

func TestRFTagService_DeleteRFTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "RFTagService",
		Operation:   "DeleteRFTag",
		NilClientTest: func() error {
			service := NewRFTagService(nil)
			return service.DeleteRFTag(context.Background(), "test-rf-tag")
		},
		EmptyParamTest: func() error {
			service := &RFTagService{}
			return service.DeleteRFTag(context.Background(), "") // Empty tag name
		},
	})
}

func TestRFTagService_DeleteRFTag_StatusCode(t *testing.T) {
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
			tagName := "test-rf-tag"
			endpoint := fmt.Sprintf("%s=%s", routes.RfTagsEndpoint, tagName)
			server := mock.NewRESTCONFErrorServer([]string{endpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewRFTagService(testClient)

			ctx := context.Background()
			err := service.DeleteRFTag(ctx, tagName)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestRFTagService_DeleteRFTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewRFTagService(testClient)

	// Use timestamp to ensure unique tag name
	tagName := fmt.Sprintf("test-rf-delete-%d", time.Now().Unix())

	config := RFTagConfig{
		TagName:             tagName,
		Description:         "Test RF tag for deletion", // 64 chars limit
		Dot11ARfProfileName: "default-11a-profile",
		Dot11BRfProfileName: "default-11b-profile",
	}

	ctx := context.Background()

	// First create a tag to delete
	err := service.CreateRFTag(ctx, config)
	if err != nil {
		t.Fatalf("Failed to create RF tag for deletion test: %v", err)
	}

	// Test deletion
	err = service.DeleteRFTag(ctx, tagName)
	if err != nil {
		t.Errorf("Failed to delete RF tag: %v", err)
	}
}
