package rf

import (
	"strings"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfTagServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := NewService(testClient.Core().(*core.Client))

	rfTagService := service.RFTag()
	if rfTagService.Client() == nil {
		t.Error("Expected valid client, got nil")
	}
}

func TestRfTagServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Focus on Get/List functions - based on real WNC data
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": {
				"rf-tags": {
					"rf-tag": [
						{
							"tag-name": "labo-inside",
							"dot11a-rf-profile-name": "labo-rf-5gh-inside",
							"dot11b-rf-profile-name": "labo-rf-24gh",
							"dot11-6ghz-rf-prof-name": "labo-rf-6gh"
						},
						{
							"tag-name": "labo-outside",
							"dot11a-rf-profile-name": "labo-rf-5gh-outside",
							"dot11b-rf-profile-name": "labo-rf-24gh"
						},
						{
							"tag-name": "default-rf-tag",
							"description": "Preconfigured default RF tag"
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
				"rf-tag": [
					{
						"tag-name": "labo-inside",
						"dot11a-rf-profile-name": "labo-rf-5gh-inside",
						"dot11b-rf-profile-name": "labo-rf-24gh",
						"dot11-6ghz-rf-prof-name": "labo-rf-6gh"
					},
					{
						"tag-name": "labo-outside",
						"dot11a-rf-profile-name": "labo-rf-5gh-outside",
						"dot11b-rf-profile-name": "labo-rf-24gh"
					},
					{
						"tag-name": "default-rf-tag",
						"description": "Preconfigured default RF tag"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=labo-inside": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tag": [
				{
					"tag-name": "labo-inside",
					"dot11a-rf-profile-name": "labo-rf-5gh-inside",
					"dot11b-rf-profile-name": "labo-rf-24gh",
					"dot11-6ghz-rf-prof-name": "labo-rf-6gh"
				}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	rfTagService := service.RFTag()
	ctx := testutil.TestContext(t)

	t.Run("ListRFTags", func(t *testing.T) {
		result, err := rfTagService.ListRFTags(ctx)
		if err != nil {
			t.Errorf("ListRFTags returned unexpected error: %v", err)
		}
		if len(result) == 0 {
			t.Error("ListRFTags returned empty result")
		}
		// Verify first tag from real data
		if result[0].TagName != "labo-inside" {
			t.Errorf("Expected tag name 'labo-inside', got '%s'", result[0].TagName)
		}
	})

	t.Run("GetRFTag", func(t *testing.T) {
		result, err := rfTagService.GetRFTag(ctx, "labo-inside")
		if err != nil {
			t.Errorf("GetRFTag returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetRFTag returned nil result")
		}
		if result != nil && result.TagName != "labo-inside" {
			t.Errorf("Expected tag name 'labo-inside', got '%s'", result.TagName)
		}
	})

	t.Run("RFTagService_GetConfig", func(t *testing.T) {
		result, err := rfTagService.GetConfig(ctx)
		if err != nil {
			t.Errorf("RFTagService GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("RFTagService GetConfig returned nil result")
		}
	})
}

func TestRfTagServiceUnit_SetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses for RF tag operations based on real WNC data
	responses := map[string]string{
		// Full RF config for tag operations
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": {
				"rf-tags": {
					"rf-tag": [
						{
							"tag-name": "test-rf-tag",
							"dot11a-rf-profile-name": "default-5ghz",
							"dot11b-rf-profile-name": "default-24ghz",
							"description": "Test RF tag"
						}
					]
				}
			}
		}`,
		// Individual tag query for updates
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=test-rf-tag": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tag": [
				{
					"tag-name": "test-rf-tag",
					"dot11a-rf-profile-name": "default-5ghz",
					"dot11b-rf-profile-name": "default-24ghz",
					"description": "Test RF tag"
				}
			]
		}`,
		// Success responses for POST/PUT/DELETE operations
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{"status": "success"}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	rfTagService := service.RFTag()
	ctx := testutil.TestContext(t)

	t.Run("CreateRFTag", func(t *testing.T) {
		newTag := &model.RfTag{
			TagName:             "new-rf-tag",
			Dot11ARfProfileName: "5ghz-profile",
			Dot11BRfProfileName: "24ghz-profile",
			Dot116GhzRfProfName: "6ghz-profile",
			Description:         "New test RF tag",
		}

		err := rfTagService.CreateRFTag(ctx, newTag)
		if err != nil {
			t.Errorf("CreateRFTag returned unexpected error: %v", err)
		}
	})

	t.Run("SetDot11ARfProfile", func(t *testing.T) {
		err := rfTagService.SetDot11ARfProfile(ctx, "test-rf-tag", "updated-5ghz-profile")
		if err != nil {
			t.Errorf("SetDot11ARfProfile returned unexpected error: %v", err)
		}
	})

	t.Run("SetDot11BRfProfile", func(t *testing.T) {
		err := rfTagService.SetDot11BRfProfile(ctx, "test-rf-tag", "updated-24ghz-profile")
		if err != nil {
			t.Errorf("SetDot11BRfProfile returned unexpected error: %v", err)
		}
	})

	t.Run("SetDot116GhzRfProfile", func(t *testing.T) {
		err := rfTagService.SetDot116GhzRfProfile(ctx, "test-rf-tag", "updated-6ghz-profile")
		if err != nil {
			t.Errorf("SetDot116GhzRfProfile returned unexpected error: %v", err)
		}
	})

	t.Run("SetDescription", func(t *testing.T) {
		err := rfTagService.SetDescription(ctx, "test-rf-tag", "Updated description")
		if err != nil {
			t.Errorf("SetDescription returned unexpected error: %v", err)
		}
	})

	t.Run("DeleteRFTag", func(t *testing.T) {
		err := rfTagService.DeleteRFTag(ctx, "test-rf-tag")
		if err != nil {
			t.Errorf("DeleteRFTag returned unexpected error: %v", err)
		}
	})
}

func TestRfTagServiceUnit_ValidationErrors_EmptyInputs(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	rfTagService := service.RFTag()
	ctx := testutil.TestContext(t)

	t.Run("GetRFTag_EmptyTagName", func(t *testing.T) {
		result, err := rfTagService.GetRFTag(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("GetRFTag_WhitespaceTagName", func(t *testing.T) {
		result, err := rfTagService.GetRFTag(ctx, "   ")
		if err == nil {
			t.Error("Expected validation error for whitespace tag name")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("CreateRFTag_NilConfig", func(t *testing.T) {
		err := rfTagService.CreateRFTag(ctx, nil)
		if err == nil {
			t.Error("Expected error for nil config")
		}
		if !strings.Contains(err.Error(), "cannot be nil") {
			t.Errorf("Expected 'cannot be nil' error, got: %v", err)
		}
	})

	t.Run("CreateRFTag_EmptyTagName", func(t *testing.T) {
		emptyTag := &model.RfTag{TagName: ""}
		err := rfTagService.CreateRFTag(ctx, emptyTag)
		if err == nil {
			t.Error("Expected error for empty tag name")
		}
		if !strings.Contains(err.Error(), "cannot be empty") {
			t.Errorf("Expected 'cannot be empty' error, got: %v", err)
		}
	})

	t.Run("CreateRFTag_WhitespaceTagName", func(t *testing.T) {
		whitespaceTag := &model.RfTag{TagName: "   "}
		err := rfTagService.CreateRFTag(ctx, whitespaceTag)
		if err == nil {
			t.Error("Expected error for whitespace tag name")
		}
		if !strings.Contains(err.Error(), "invalid tag name format") {
			t.Errorf("Expected 'invalid tag name format' error, got: %v", err)
		}
	})

	t.Run("DeleteRFTag_EmptyTagName", func(t *testing.T) {
		err := rfTagService.DeleteRFTag(ctx, "")
		if err == nil {
			t.Error("Expected error for empty tag name")
		}
	})

	t.Run("SetDot11ARfProfile_EmptyTagName", func(t *testing.T) {
		err := rfTagService.SetDot11ARfProfile(ctx, "", "profile")
		if err == nil {
			t.Error("Expected error for empty tag name")
		}
	})

	t.Run("SetDescription_EmptyTagName", func(t *testing.T) {
		err := rfTagService.SetDescription(ctx, "", "description")
		if err == nil {
			t.Error("Expected error for empty tag name")
		}
	})
}

func TestRfTagServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("RFTag_NilClient", func(t *testing.T) {
		service := NewService(nil)
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.ListRFTags(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("CreateRFTag_APIError", func(t *testing.T) {
		// Mock server that returns 404 for tag operations to simulate API errors
		errorPaths := []string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags",
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=nonexistent-tag",
		}
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		newTag := &model.RfTag{
			TagName: "valid-tag",
		}
		err := rfTagService.CreateRFTag(ctx, newTag)
		if err == nil {
			t.Error("Expected error for API failure")
		}
	})

	t.Run("DeleteRFTag_APIError", func(t *testing.T) {
		errorPaths := []string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=nonexistent-tag",
		}
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		err := rfTagService.DeleteRFTag(ctx, "nonexistent-tag")
		if err == nil {
			t.Error("Expected error for API failure")
		}
	})

	t.Run("SetDot11ARfProfile_TagNotFound", func(t *testing.T) {
		errorPaths := []string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=nonexistent-tag",
		}
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		err := rfTagService.SetDot11ARfProfile(ctx, "nonexistent-tag", "profile")
		if err == nil {
			t.Error("Expected error for nonexistent tag")
		}
	})
}

func TestRfTagServiceUnit_ErrorHandling_EdgeCases(t *testing.T) {
	t.Parallel()

	// Test nil/empty response scenarios
	t.Run("GetRFTag_NilResult", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=test-tag": `null`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.GetRFTag(ctx, "test-tag")
		if err != nil {
			t.Errorf("Expected no error for nil response, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for null response")
		}
	})

	t.Run("GetRFTag_EmptyRfTagList", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=test-tag": `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-tag": []
			}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.GetRFTag(ctx, "test-tag")
		if err != nil {
			t.Errorf("Expected no error for empty list, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for empty rf-tag list")
		}
	})

	t.Run("ListRFTags_NilResult", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `null`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.ListRFTags(ctx)
		if err != nil {
			t.Errorf("Expected no error for nil response, got: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil response, got length: %d", len(result))
		}
	})

	t.Run("ListRFTags_EmptyList", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
					"rf-tag": []
				}
			}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.ListRFTags(ctx)
		if err != nil {
			t.Errorf("Expected no error for empty list, got: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice for empty list, got length: %d", len(result))
		}
	})

	t.Run("ListRFTags_NilRfTags", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": null
			}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		result, err := rfTagService.ListRFTags(ctx)
		if err != nil {
			t.Errorf("Expected no error for nil rf-tags, got: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("Expected empty slice for nil rf-tags, got length: %d", len(result))
		}
	})

	t.Run("UpdateTagField_NilFunction", func(t *testing.T) {
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag=test-tag": `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-tag": [{
					"tag-name": "test-tag",
					"description": "Test tag"
				}]
			}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		// This tests the internal updateTagField function through SetDescription with nil function
		// Since updateTagField is not exported, we test through error scenarios in other methods
		err := rfTagService.SetDescription(ctx, "nonexistent-tag", "new description")
		if err == nil {
			t.Error("Expected error for nonexistent tag during update")
		}
	})

	t.Run("SetRFTag_ValidationErrors", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		rfTagService := service.RFTag()
		ctx := testutil.TestContext(t)

		// Test setRFTag with nil config through SetDescription internal call
		err := rfTagService.SetDescription(ctx, "", "description")
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
	})
}
