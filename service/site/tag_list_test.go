package site

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/mock"
)

func TestSiteTagService_ListSiteTags_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "SiteTagService",
		Operation:   "ListSiteTags",
		NilClientTest: func() error {
			service := NewSiteTagService(nil)
			_, err := service.ListSiteTags(context.Background())
			return err
		},
		EmptyParamTest: func() error {
			service := &SiteTagService{}
			_, err := service.ListSiteTags(context.Background())
			return err
		},
	})
}

func TestSiteTagService_ListSiteTags_StatusCode(t *testing.T) {
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
			server := mock.NewRESTCONFErrorServer([]string{routes.SiteTagConfigsEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewSiteTagService(testClient)

			ctx := context.Background()
			_, err := service.ListSiteTags(ctx)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestSiteTagService_ListSiteTags_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewSiteTagService(testClient)
	ctx := context.Background()

	// Test listing site tags
	tags, err := service.ListSiteTags(ctx)
	if err != nil {
		t.Errorf("Failed to list site tags: %v", err)
	}

	// Basic validation
	if tags == nil {
		t.Error("Expected non-nil response for site tags list")
		return
	}

	t.Logf("Successfully listed site tags (count: %d)", len(tags))
}
