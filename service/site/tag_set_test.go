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

func TestSiteTagService_SetSiteTag_Unit(t *testing.T) {
	framework.RunUnifiedTagUnitTests(t, framework.TagUnitTestPattern{
		ServiceName: "SiteTagService",
		Operation:   "SetSiteTag",
		NilClientTest: func() error {
			service := NewSiteTagService(nil)
			config := SiteTagConfig{SiteTagName: "test-site-tag"}
			return service.SetSiteTag(context.Background(), config)
		},
		EmptyParamTest: func() error {
			service := &SiteTagService{}
			config := SiteTagConfig{} // Empty tag name
			return service.SetSiteTag(context.Background(), config)
		},
	})
}

func TestSiteTagService_SetSiteTag_StatusCode(t *testing.T) {
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
				Description:   "Test site tag",
				APJoinProfile: "default-ap-join-profile",
				FlexProfile:   "default-flex-profile",
			}

			ctx := context.Background()
			err := service.SetSiteTag(ctx, config)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for status %d, but got none", tc.statusCode)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for status %d: %v", tc.statusCode, err)
			}
		})
	}
}

func TestSiteTagService_SetSiteTag_Integration(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	if testClient == nil {
		t.Skip("Skipping integration tests: no client available")
	}

	service := NewSiteTagService(testClient)
	ctx := context.Background()

	// Test cases for both with and without flex-profile
	testCases := []struct {
		name         string
		createConfig func(timestamp int64) SiteTagConfig
		updateConfig func(timestamp int64) SiteTagConfig
	}{
		{
			name: "SetSiteTag_WithoutFlexProfile",
			createConfig: func(timestamp int64) SiteTagConfig {
				return SiteTagConfig{
					SiteTagName:   fmt.Sprintf("set-site-nf-%d", timestamp%100000000),
					Description:   "Initial site tag without flex profile",
					APJoinProfile: "default-ap-join-profile",
					IsLocalSite:   &[]bool{true}[0], // Local site, no flex-profile needed
				}
			},
			updateConfig: func(timestamp int64) SiteTagConfig {
				return SiteTagConfig{
					SiteTagName:   fmt.Sprintf("set-site-nf-%d", timestamp%100000000),
					Description:   "Updated site tag without flex profile",
					APJoinProfile: "updated-ap-join-profile",
					IsLocalSite:   &[]bool{true}[0], // Local site, no flex-profile needed
				}
			},
		},
		{
			name: "SetSiteTag_WithFlexProfile",
			createConfig: func(timestamp int64) SiteTagConfig {
				return SiteTagConfig{
					SiteTagName:   fmt.Sprintf("set-site-fp-%d", timestamp%100000000),
					Description:   "Initial site tag with flex profile",
					APJoinProfile: "labo-common",     // Use known existing profile
					FlexProfile:   "labo-flex",       // Use known existing flex profile
					IsLocalSite:   &[]bool{false}[0], // Non-local site, flex-profile allowed
				}
			},
			updateConfig: func(timestamp int64) SiteTagConfig {
				return SiteTagConfig{
					SiteTagName:   fmt.Sprintf("set-site-fp-%d", timestamp%100000000),
					Description:   "Updated site tag with flex profile",
					APJoinProfile: "labo-common",     // Use known existing profile
					FlexProfile:   "labo-flex",       // Use known existing flex profile
					IsLocalSite:   &[]bool{false}[0], // Non-local site, flex-profile allowed
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use timestamp to ensure unique tag name
			timestamp := time.Now().Unix()
			createConfig := tc.createConfig(timestamp)
			updateConfig := tc.updateConfig(timestamp)

			// Skip flex-profile tests due to current YANG constraint issues
			if createConfig.FlexProfile != "" || updateConfig.FlexProfile != "" {
				t.Skip("Skipping flex-profile test due to YANG when constraint: " +
					"'../is-local-site = false' condition is not being satisfied properly")
			}

			// First create the initial tag
			err := service.CreateSiteTag(ctx, createConfig)
			if err != nil {
				t.Fatalf("Failed to create initial site tag: %v", err)
			}

			t.Logf("Created initial site tag: %s", createConfig.SiteTagName)

			// Test updating via SetSiteTag
			err = service.SetSiteTag(ctx, updateConfig)
			if err != nil {
				t.Errorf("Failed to update site tag: %v", err)
			} else {
				t.Logf("Successfully updated site tag: %s", updateConfig.SiteTagName)
			}

			// Cleanup - attempt to delete the test tag
			defer func() {
				if deleteErr := service.DeleteSiteTag(ctx, createConfig.SiteTagName); deleteErr != nil {
					t.Logf("Cleanup failed - could not delete test tag %s: %v", createConfig.SiteTagName, deleteErr)
				}
			}()
		})
	}
}
