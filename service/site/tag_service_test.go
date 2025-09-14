package site

import (
	"strings"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestSiteTagServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	siteTagService := service.SiteTag()
	if siteTagService == nil {
		t.Error("SiteTag service should not be nil")
	}
}

func TestSiteTagServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	mockResponses := map[string]string{
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=test-site": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": [{
				"site-tag-name": "test-site",
				"description": "Test Site",
				"ap-join-profile": "default-ap-policy-profile",
				"local-site": false
			}]
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
				"site-tag-config": [{
					"site-tag-name": "test-site",
					"description": "Test Site",
					"ap-join-profile": "default-ap-policy-profile",
					"local-site": false
				}]
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("GetSiteTag", func(t *testing.T) {
		result, err := siteTagService.GetSiteTag(ctx, "test-site")
		if err != nil {
			t.Errorf("GetSiteTag returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetSiteTag returned nil result")
		}
	})

	t.Run("ListSiteTags", func(t *testing.T) {
		result, err := siteTagService.ListSiteTags(ctx)
		if err != nil {
			t.Errorf("ListSiteTags returned unexpected error: %v", err)
		}
		if len(result) == 0 {
			t.Error("ListSiteTags returned empty result")
		}
	})
}

func TestSiteTagServiceUnit_SetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	mockResponses := map[string]string{
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=test-site": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": [{
				"site-tag-name": "test-site",
				"description": "Test Site for Operations",
				"ap-join-profile": "labo-common",
				"flex-profile": "labo-flex",
				"is-local-site": false
			}]
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
				"site-tag-config": [{
					"site-tag-name": "test-site",
					"description": "Test Site for Operations",
					"ap-join-profile": "labo-common",
					"flex-profile": "labo-flex",
					"is-local-site": false
				}]
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("CreateSiteTag", func(t *testing.T) {
		newSiteTag := &model.SiteListEntry{
			SiteTagName:   "new-site-tag",
			Description:   strPtr("New Site Tag"),
			ApJoinProfile: strPtr("test-ap-policy-profile"),
			IsLocalSite:   boolPtr(false),
		}

		err := siteTagService.CreateSiteTag(ctx, newSiteTag)
		if err != nil {
			t.Errorf("CreateSiteTag returned unexpected error: %v", err)
		}
	})

	t.Run("SetAPJoinProfile", func(t *testing.T) {
		err := siteTagService.SetAPJoinProfile(ctx, "test-site", "new-ap-profile")
		if err != nil {
			t.Errorf("SetAPJoinProfile returned unexpected error: %v", err)
		}
	})

	t.Run("SetFlexProfile", func(t *testing.T) {
		err := siteTagService.SetFlexProfile(ctx, "test-site", "new-flex-profile")
		if err != nil {
			t.Errorf("SetFlexProfile returned unexpected error: %v", err)
		}
	})

	t.Run("SetLocalSite", func(t *testing.T) {
		err := siteTagService.SetLocalSite(ctx, "test-site", true)
		if err != nil {
			t.Errorf("SetLocalSite returned unexpected error: %v", err)
		}
	})

	t.Run("SetDescription", func(t *testing.T) {
		err := siteTagService.SetDescription(ctx, "test-site", "Updated description")
		if err != nil {
			t.Errorf("SetDescription returned unexpected error: %v", err)
		}
	})

	t.Run("DeleteSiteTag", func(t *testing.T) {
		err := siteTagService.DeleteSiteTag(ctx, "test-site")
		if err != nil {
			t.Errorf("DeleteSiteTag returned unexpected error: %v", err)
		}
	})
}

func TestSiteTagServiceUnit_ValidationErrors_EmptyInputs(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("GetSiteTag_EmptyTagName", func(t *testing.T) {
		_, err := siteTagService.GetSiteTag(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
	})

	t.Run("CreateSiteTag_NilConfig", func(t *testing.T) {
		err := siteTagService.CreateSiteTag(ctx, nil)
		if err == nil {
			t.Error("Expected validation error for nil config")
		}
	})

	t.Run("CreateSiteTag_EmptyTagName", func(t *testing.T) {
		config := &model.SiteListEntry{SiteTagName: ""}
		err := siteTagService.CreateSiteTag(ctx, config)
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
	})

	t.Run("DeleteSiteTag_EmptyTagName", func(t *testing.T) {
		err := siteTagService.DeleteSiteTag(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
	})
}

func TestSiteTagServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("GetSiteTag_NilClient", func(t *testing.T) {
		_, err := siteTagService.GetSiteTag(ctx, "test-site")
		if err == nil {
			t.Error("Expected error with nil client for GetSiteTag")
		}
	})

	t.Run("ListSiteTags_NilClient", func(t *testing.T) {
		_, err := siteTagService.ListSiteTags(ctx)
		if err == nil {
			t.Error("Expected error with nil client for ListSiteTags")
		}
	})

	t.Run("CreateSiteTag_NilClient", func(t *testing.T) {
		config := &model.SiteListEntry{
			SiteTagName: "test-site",
		}
		err := siteTagService.CreateSiteTag(ctx, config)
		if err == nil {
			t.Error("Expected error with nil client for CreateSiteTag")
		}
	})
}

func TestSiteTagServiceUnit_GetOperations_RealWNCData(t *testing.T) {
	t.Parallel()

	// Based on real WNC 17.12.5 data structure
	mockResponses := map[string]string{
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
				"site-tag-config": [
					{
						"site-tag-name": "labo-site-flex",
						"flex-profile": "labo-flex",
						"ap-join-profile": "labo-common",
						"is-local-site": false
					},
					{
						"site-tag-name": "default-site-tag",
						"description": "Test description update",
						"ap-join-profile": "test-ap-profile",
						"is-local-site": false
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=labo-site-flex": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": [{
				"site-tag-name": "labo-site-flex",
				"flex-profile": "labo-flex",
				"ap-join-profile": "labo-common",
				"is-local-site": false
			}]
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=default-site-tag": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": [{
				"site-tag-name": "default-site-tag",
				"description": "Test description update",
				"ap-join-profile": "test-ap-profile",
				"is-local-site": false
			}]
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=non-existent": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": []
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("GetSiteTag_RealDataStructure", func(t *testing.T) {
		// Test with labo-site-flex (has flex-profile)
		result, err := siteTagService.GetSiteTag(ctx, "labo-site-flex")
		if err != nil {
			t.Errorf("GetSiteTag failed for labo-site-flex: %v", err)
		}
		if result == nil {
			t.Error("Expected site tag result, got nil")
		}
		if result != nil && result.SiteTagName != "labo-site-flex" {
			t.Errorf("Expected site tag name 'labo-site-flex', got %s", result.SiteTagName)
		}

		// Test with default-site-tag (has description)
		result2, err := siteTagService.GetSiteTag(ctx, "default-site-tag")
		if err != nil {
			t.Errorf("GetSiteTag failed for default-site-tag: %v", err)
		}
		if result2 == nil {
			t.Error("Expected site tag result, got nil")
		}
	})

	t.Run("GetSiteTag_NonExistent", func(t *testing.T) {
		// Test non-existent site tag returns nil without error
		result, err := siteTagService.GetSiteTag(ctx, "non-existent")
		if err != nil {
			t.Errorf("GetSiteTag failed for non-existent tag: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for non-existent site tag")
		}
	})

	t.Run("ListSiteTags_RealDataStructure", func(t *testing.T) {
		results, err := siteTagService.ListSiteTags(ctx)
		if err != nil {
			t.Errorf("ListSiteTags failed: %v", err)
		}
		if len(results) != 2 {
			t.Errorf("Expected 2 site tags, got %d", len(results))
		}

		// Verify real data structure fields
		hasLaboSite := false
		hasDefaultSite := false
		for _, tag := range results {
			if tag.SiteTagName == "labo-site-flex" {
				hasLaboSite = true
				if tag.FlexProfile == nil || *tag.FlexProfile != "labo-flex" {
					t.Error("Expected labo-site-flex to have flex-profile 'labo-flex'")
				}
			}
			if tag.SiteTagName == "default-site-tag" {
				hasDefaultSite = true
				if tag.Description == nil || *tag.Description != "Test description update" {
					t.Error("Expected default-site-tag to have proper description")
				}
			}
		}
		if !hasLaboSite {
			t.Error("Expected to find labo-site-flex in results")
		}
		if !hasDefaultSite {
			t.Error("Expected to find default-site-tag in results")
		}
	})
}

func TestSiteTagServiceUnit_GetOperations_EdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		mockResponse string
		expectEmpty  bool
	}{
		{
			name: "NilSiteTagConfigs",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": null
			}`,
			expectEmpty: true,
		},
		{
			name: "EmptySiteTagConfig",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
					"site-tag-config": null
				}
			}`,
			expectEmpty: true,
		},
		{
			name: "EmptyArray",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
					"site-tag-config": []
				}
			}`,
			expectEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
				"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": tt.mockResponse,
			}))
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := NewService(testClient.Core().(*core.Client))
			siteTagService := service.SiteTag()
			ctx := testutil.TestContext(t)

			results, err := siteTagService.ListSiteTags(ctx)
			if err != nil {
				t.Errorf("ListSiteTags failed for %s: %v", tt.name, err)
			}

			if tt.expectEmpty && len(results) != 0 {
				t.Errorf("Expected empty results for %s, got %d", tt.name, len(results))
			}
		})
	}
}

func TestSiteTagServiceUnit_SetOperations_AdvancedScenarios(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC data patterns
	mockResponses := map[string]string{
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=test-site": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": [{
				"site-tag-name": "test-site",
				"description": "Test Site for Operations",
				"ap-join-profile": "labo-common",
				"flex-profile": "labo-flex",
				"is-local-site": false
			}]
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
				"site-tag-config": [{
					"site-tag-name": "test-site",
					"description": "Test Site for Operations",
					"ap-join-profile": "labo-common",
					"flex-profile": "labo-flex",
					"is-local-site": false
				}]
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	siteTagService := service.SiteTag()
	ctx := testutil.TestContext(t)

	t.Run("SetAPJoinProfile_ExistingTag", func(t *testing.T) {
		err := siteTagService.SetAPJoinProfile(ctx, "test-site", "new-ap-profile")
		if err != nil {
			t.Errorf("SetAPJoinProfile failed: %v", err)
		}
	})

	t.Run("SetFlexProfile_ExistingTag", func(t *testing.T) {
		err := siteTagService.SetFlexProfile(ctx, "test-site", "new-flex-profile")
		if err != nil {
			t.Errorf("SetFlexProfile failed: %v", err)
		}
	})

	t.Run("SetLocalSite_ExistingTag", func(t *testing.T) {
		err := siteTagService.SetLocalSite(ctx, "test-site", true)
		if err != nil {
			t.Errorf("SetLocalSite failed: %v", err)
		}
	})

	t.Run("SetDescription_ExistingTag", func(t *testing.T) {
		err := siteTagService.SetDescription(ctx, "test-site", "Updated description from test")
		if err != nil {
			t.Errorf("SetDescription failed: %v", err)
		}
	})
}

func TestSiteTagServiceUnit_ErrorHandling_ComprehensiveScenarios(t *testing.T) {
	t.Parallel()

	t.Run("ValidationErrors_TagName", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		siteTagService := service.SiteTag()
		ctx := testutil.TestContext(t)

		// Test validateTagName with whitespace-only string
		_, err := siteTagService.GetSiteTag(ctx, "   ")
		if err == nil {
			t.Error("Expected validation error for whitespace-only tag name")
		}

		err = siteTagService.DeleteSiteTag(ctx, "   ")
		if err == nil {
			t.Error("Expected validation error for whitespace-only tag name in DeleteSiteTag")
		}
	})

	t.Run("SetOperations_TagNotFound", func(t *testing.T) {
		// Mock server that returns empty result for non-existent tag
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=non-existent": `{
				"Cisco-IOS-XE-wireless-site-cfg:site-tag-config": []
			}`,
		}

		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		siteTagService := service.SiteTag()
		ctx := testutil.TestContext(t)

		// Test SetAPJoinProfile with non-existent tag
		err := siteTagService.SetAPJoinProfile(ctx, "non-existent", "new-profile")
		if err == nil {
			t.Error("Expected error for SetAPJoinProfile with non-existent tag")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Expected 'not found' error, got: %v", err)
		}

		// Test SetFlexProfile with non-existent tag
		err = siteTagService.SetFlexProfile(ctx, "non-existent", "new-flex")
		if err == nil {
			t.Error("Expected error for SetFlexProfile with non-existent tag")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Expected 'not found' error, got: %v", err)
		}

		// Test SetLocalSite with non-existent tag
		err = siteTagService.SetLocalSite(ctx, "non-existent", true)
		if err == nil {
			t.Error("Expected error for SetLocalSite with non-existent tag")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Expected 'not found' error, got: %v", err)
		}

		// Test SetDescription with non-existent tag
		err = siteTagService.SetDescription(ctx, "non-existent", "new description")
		if err == nil {
			t.Error("Expected error for SetDescription with non-existent tag")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Expected 'not found' error, got: %v", err)
		}
	})

	t.Run("SetOperations_GetTagError", func(t *testing.T) {
		// Mock server that returns error for GET operations using error server
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
			"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config=test-site",
		}, 500))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		siteTagService := service.SiteTag()
		ctx := testutil.TestContext(t)

		// Test SetAPJoinProfile with server error during tag retrieval
		err := siteTagService.SetAPJoinProfile(ctx, "test-site", "new-profile")
		if err == nil {
			t.Error("Expected error for SetAPJoinProfile with server error")
		}
		if !strings.Contains(err.Error(), "tag retrieval failed") {
			t.Errorf("Expected 'tag retrieval failed' error, got: %v", err)
		}

		// Test SetFlexProfile with server error during tag retrieval
		err = siteTagService.SetFlexProfile(ctx, "test-site", "new-flex")
		if err == nil {
			t.Error("Expected error for SetFlexProfile with server error")
		}
		if !strings.Contains(err.Error(), "tag retrieval failed") {
			t.Errorf("Expected 'tag retrieval failed' error, got: %v", err)
		}

		// Test SetLocalSite with server error during tag retrieval
		err = siteTagService.SetLocalSite(ctx, "test-site", true)
		if err == nil {
			t.Error("Expected error for SetLocalSite with server error")
		}
		if !strings.Contains(err.Error(), "tag retrieval failed") {
			t.Errorf("Expected 'tag retrieval failed' error, got: %v", err)
		}

		// Test SetDescription with server error during tag retrieval
		err = siteTagService.SetDescription(ctx, "test-site", "new description")
		if err == nil {
			t.Error("Expected error for SetDescription with server error")
		}
		if !strings.Contains(err.Error(), "tag retrieval failed") {
			t.Errorf("Expected 'tag retrieval failed' error, got: %v", err)
		}
	})
}

// Helper functions for pointer creation.
func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
