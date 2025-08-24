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

func TestSiteTagService_CreateSiteTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "SiteTagService",
		Operation:   "CreateSiteTag",
		NilClientTest: func() error {
			service := NewSiteTagService(nil)
			config := SiteTagConfig{SiteTagName: "test-site-tag"}
			return service.CreateSiteTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &SiteTagService{}
			config := SiteTagConfig{} // Empty tag name
			return service.CreateSiteTag(context.Background(), config)
		},
	})
}

func TestSiteTagService_CreateSiteTag_StatusCode(t *testing.T) {
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
			server := mock.NewRESTCONFErrorServer([]string{routes.SiteTagConfigsEndpoint}, tc.statusCode)
			defer server.Close()

			testClient := mock.NewTLSClientForServer(t, server)
			service := NewSiteTagService(testClient)
			config := SiteTagConfig{
				SiteTagName:   "test-site-tag",
				Description:   "Test Site tag created via status code test",
				APJoinProfile: "default-ap-join-profile",
				IsLocalSite:   &[]bool{true}[0],
			}

			ctx := context.Background()
			err := service.CreateSiteTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestSiteTagService_CreateSiteTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewSiteTagService(testClient)

	// Use timestamp to ensure unique tag name (limit to 32 chars)
	tagName := fmt.Sprintf("crt-site-nf-%d", time.Now().Unix()%100000000)

	config := SiteTagConfig{
		SiteTagName: tagName,
		Description: fmt.Sprintf(
			"Integration test Site tag created at %s",
			time.Now().Format("2006-01-02 15:04:05"),
		),
		APJoinProfile: "default-ap-join-profile",
		IsLocalSite:   &[]bool{true}[0], // Local site, no flex-profile needed
	}

	ctx := context.Background()

	// Test creation
	err := service.CreateSiteTag(ctx, config)
	if err != nil {
		t.Errorf("Failed to create Site tag: %v", err)
		return
	}

	t.Logf("Successfully created site tag: %s", tagName)

	// Cleanup - attempt to delete the test tag
	if deleteErr := service.DeleteSiteTag(ctx, tagName); deleteErr != nil {
		t.Logf("Cleanup failed - could not delete test tag %s: %v", tagName, deleteErr)
	}
}
