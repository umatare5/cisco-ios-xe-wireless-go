package rf

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestRFTagService_ListRFTags_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "RFTagService",
		Operation:   "ListRFTags",
		NilClientTest: func() error {
			service := NewRFTagService(nil)
			_, err := service.ListRFTags(context.Background())
			return err
		},
		EmptyParamTest: func() error {
			service := &RFTagService{}
			_, err := service.ListRFTags(context.Background())
			return err
		},
	})
}

func TestRFTagService_ListRFTags_StatusCode(t *testing.T) {
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
			server := mock.NewRESTCONFErrorServer([]string{routes.RfTagsEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewRFTagService(testClient)

			ctx := context.Background()
			_, err := service.ListRFTags(ctx)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestRFTagService_ListRFTags_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewRFTagService(testClient)
	ctx := context.Background()

	// Test listing RF tags
	tags, err := service.ListRFTags(ctx)
	if err != nil {
		t.Errorf("Failed to list RF tags: %v", err)
	}

	// Basic validation
	if tags == nil {
		t.Error("Expected non-nil response for RF tags list")
		return
	}

	var tagCount int
	if tags.RfTags != nil {
		tagCount = len(tags.RfTags.RfTag)
	} else {
		tagCount = len(tags.RfTag)
	}

	t.Logf("Successfully listed RF tags (count: %d)", tagCount)
}
