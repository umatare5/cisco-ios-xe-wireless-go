package site

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

func TestSiteTagService_GetSiteTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "SiteTagService",
		Operation:   "GetSiteTag",
		NilClientTest: func() error {
			service := NewSiteTagService(nil)
			_, err := service.GetSiteTag(context.Background(), "test-site-tag")
			return err
		},
		EmptyParamTest: func() error {
			service := &SiteTagService{}
			_, err := service.GetSiteTag(context.Background(), "") // Empty tag name
			return err
		},
	})
}

func TestSiteTagService_GetSiteTag_StatusCode(t *testing.T) {
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
			tagName := "test-site-tag"
			endpoint := fmt.Sprintf("%s/site-tag-config=%s", routes.SiteTagConfigsEndpoint, tagName)
			server := mock.NewRESTCONFErrorServer([]string{endpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewSiteTagService(testClient)

			ctx := context.Background()
			_, err := service.GetSiteTag(ctx, tagName)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestSiteTagService_GetSiteTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewSiteTagService(testClient)
	ctx := context.Background()

	// Create a tag first, then try to get it
	tagName := fmt.Sprintf("get-site-test-%d", time.Now().Unix()%100000000)
	config := SiteTagConfig{
		SiteTagName:   tagName,
		Description:   "Test site tag for get operation",
		APJoinProfile: "default-ap-join-profile",
		IsLocalSite:   &[]bool{true}[0],
	}

	// Create the tag
	err := service.CreateSiteTag(ctx, config)
	if err != nil {
		t.Errorf("Failed to create site tag for get test: %v", err)
		return
	}

	t.Logf("Created test tag: %s", tagName)

	// Now try to get it
	tag, err := service.GetSiteTag(ctx, tagName)
	switch {
	case err != nil:
		t.Errorf("Failed to get site tag '%s': %v", tagName, err)
	case tag == nil:
		t.Error("Expected non-nil response for site tag get")
	default:
		// Additional validation
		if tag.SiteTagName != tagName {
			t.Errorf("Expected tag name '%s', got '%s'", tagName, tag.SiteTagName)
		}
		t.Logf("Successfully retrieved site tag: %s", tagName)
	}

	// Clean up: delete the test tag
	if deleteErr := service.DeleteSiteTag(ctx, tagName); deleteErr != nil {
		t.Logf("Warning: Failed to clean up test tag '%s': %v", tagName, deleteErr)
	} else {
		t.Logf("Successfully cleaned up test tag: %s", tagName)
	}
}
