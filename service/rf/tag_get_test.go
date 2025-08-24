package rf

import (
	"context"
	"fmt"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestRFTagService_GetRFTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "RFTagService",
		Operation:   "GetRFTag",
		NilClientTest: func() error {
			service := NewRFTagService(nil)
			_, err := service.GetRFTag(context.Background(), "test-rf-tag")
			return err
		},
		EmptyParamTest: func() error {
			service := &RFTagService{}
			_, err := service.GetRFTag(context.Background(), "") // Empty tag name
			return err
		},
	})
}

func TestRFTagService_GetRFTag_StatusCode(t *testing.T) {
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
			tagName := "test-rf-tag"
			endpoint := fmt.Sprintf("%s=%s", routes.RfTagsEndpoint, tagName)
			server := mock.NewRESTCONFErrorServer([]string{endpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewRFTagService(testClient)

			ctx := context.Background()
			_, err := service.GetRFTag(ctx, tagName)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestRFTagService_GetRFTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewRFTagService(testClient)
	ctx := context.Background()

	// Test getting a specific RF tag
	// Note: This assumes at least one RF tag exists, or we could create one first
	tagName := "labo-inside" // Using an existing tag name from the system

	tag, err := service.GetRFTag(ctx, tagName)
	if err != nil {
		t.Errorf("Failed to get RF tag '%s': %v", tagName, err)
		return
	}

	// Basic validation
	if tag == nil {
		t.Error("Expected non-nil response for RF tag get")
		return
	}

	t.Logf("Successfully retrieved RF tag: %s", tagName)
}
