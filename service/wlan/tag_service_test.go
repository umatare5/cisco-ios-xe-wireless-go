package wlan

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestWlanPolicyTagServiceUnit_Constructor_Success(t *testing.T) {
	service := NewService(nil)
	if service.PolicyTag() == nil {
		t.Error("Expected PolicyTag service, got nil")
	}
}

func TestWlanPolicyTagServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Mock server with policy tag responses
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
				"policy-list-entry": [
					{
						"tag-name": "test-policy-tag",
						"description": "Test policy tag",
						"wlan-policies": {
							"wlan-policy": [
								{
									"wlan-profile-name": "test-wlan",
									"policy-profile-name": "test-profile"
								}
							]
						}
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=test-policy-tag": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry": [
				{
					"tag-name": "test-policy-tag",
					"description": "Test policy tag",
					"wlan-policies": {
						"wlan-policy": [
							{
								"wlan-profile-name": "test-wlan",
								"policy-profile-name": "test-profile"
							}
						]
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=new-tag": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry": [
				{
					"tag-name": "new-tag",
					"description": "New policy tag"
				}
			]
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := NewService(client.Core().(*core.Client))
	policyTag := service.PolicyTag()
	ctx := testutil.TestContext(t)

	// Test ListPolicyTags
	tags, err := policyTag.ListPolicyTags(ctx)
	if err != nil {
		t.Errorf("ListPolicyTags failed: %v", err)
		return
	}

	if len(tags) == 0 {
		t.Error("ListPolicyTags returned empty result")
		return
	}

	// Test GetPolicyTag
	tag, err := policyTag.GetPolicyTag(ctx, "test-policy-tag")
	if err != nil {
		t.Errorf("GetPolicyTag failed: %v", err)
		return
	}

	if tag == nil {
		t.Error("GetPolicyTag returned nil result")
		return
	}

	if tag.TagName != "test-policy-tag" {
		t.Errorf("Expected tag name 'test-policy-tag', got %s", tag.TagName)
	}

	// Test GetPolicyTag for non-existent tag
	nonExistentTag, err := policyTag.GetPolicyTag(ctx, "non-existent")
	if err != nil {
		t.Errorf("GetPolicyTag failed for non-existent tag: %v", err)
		return
	}

	if nonExistentTag != nil {
		t.Error("Expected nil for non-existent tag, got result")
	}

	t.Logf("Policy tag operations returned valid WLAN policy tag data")
}

func TestWlanPolicyTagServiceUnit_SetOperations_MockSuccess(t *testing.T) {
	// Mock server with policy tag CRUD responses
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
				"policy-list-entry": [
					{
						"tag-name": "existing-tag",
						"description": "Existing policy tag",
						"wlan-policies": {
							"wlan-policy": [
								{
									"wlan-profile-name": "test-wlan",
									"policy-profile-name": "test-profile"
								}
							]
						}
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=existing-tag": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry": [
				{
					"tag-name": "existing-tag",
					"description": "Existing policy tag",
					"wlan-policies": {
						"wlan-policy": [
							{
								"wlan-profile-name": "test-wlan",
								"policy-profile-name": "test-profile"
							}
						]
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=updated-tag": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry": [
				{
					"tag-name": "updated-tag",
					"description": "Updated description",
					"wlan-policies": {
						"wlan-policy": [
							{
								"wlan-profile-name": "new-wlan",
								"policy-profile-name": "new-profile"
							}
						]
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=test-tag": `{}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := NewService(client.Core().(*core.Client))
	policyTag := service.PolicyTag()
	ctx := testutil.TestContext(t)

	// Test CreatePolicyTag with valid config using proper model
	config := &PolicyListEntry{
		TagName:     "new-tag",
		Description: "New policy tag",
	}

	err := policyTag.CreatePolicyTag(ctx, config)
	if err != nil {
		t.Errorf("CreatePolicyTag failed: %v", err)
		return
	}

	// Test SetPolicyTag with valid config using proper model
	updateConfig := &PolicyListEntry{
		TagName:     "updated-tag",
		Description: "Updated description",
	}

	err = policyTag.SetPolicyTag(ctx, updateConfig)
	if err != nil {
		t.Errorf("SetPolicyTag failed: %v", err)
		return
	}

	// Test SetPolicyProfile with existing tag
	err = policyTag.SetPolicyProfile(ctx, "existing-tag", "new-wlan", "new-profile")
	if err != nil {
		t.Errorf("SetPolicyProfile failed: %v", err)
		return
	}

	// Test SetDescription with existing tag
	err = policyTag.SetDescription(ctx, "existing-tag", "Updated description")
	if err != nil {
		t.Errorf("SetDescription failed: %v", err)
		return
	}

	// Test DeletePolicyTag
	err = policyTag.DeletePolicyTag(ctx, "test-tag")
	if err != nil {
		t.Errorf("DeletePolicyTag failed: %v", err)
		return
	}

	t.Logf("Policy tag CRUD operations completed successfully")
}

func TestWlanPolicyTagServiceUnit_ValidationErrors_EmptyTagName(t *testing.T) {
	service := NewService(nil)
	policyTag := service.PolicyTag()
	ctx := testutil.TestContext(t)

	// Test CreatePolicyTag with empty tag name config using proper model
	configEmptyName := &PolicyListEntry{
		TagName:     "",
		Description: "Test description",
	}

	err := policyTag.CreatePolicyTag(ctx, configEmptyName)
	if err == nil {
		t.Error("Expected error with empty tag name for CreatePolicyTag, got nil")
	}

	// Test SetPolicyTag with empty tag name config using proper model
	err = policyTag.SetPolicyTag(ctx, configEmptyName)
	if err == nil {
		t.Error("Expected error with empty tag name for SetPolicyTag, got nil")
	}

	// Test CreatePolicyTag with long tag name config using proper model
	configLongName := &PolicyListEntry{
		TagName:     "this-is-a-very-long-tag-name-that-exceeds-thirty-two-characters",
		Description: "Test description",
	}

	err = policyTag.CreatePolicyTag(ctx, configLongName)
	if err == nil {
		t.Error("Expected error with long tag name for CreatePolicyTag, got nil")
	}

	// Test SetPolicyTag with long tag name config using proper model
	err = policyTag.SetPolicyTag(ctx, configLongName)
	if err == nil {
		t.Error("Expected error with long tag name for SetPolicyTag, got nil")
	}
}

func TestWlanPolicyTagServiceUnit_ErrorHandling_NonExistentTag(t *testing.T) {
	// Mock server with error responses for SetPolicyProfile and SetDescription operations
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
				"policy-list-entry": []
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := NewService(client.Core().(*core.Client))
	policyTag := service.PolicyTag()
	ctx := testutil.TestContext(t)

	// Test SetPolicyProfile with non-existent tag
	err := policyTag.SetPolicyProfile(ctx, "non-existent-tag", "wlan", "profile")
	if err == nil {
		t.Error("Expected error with non-existent tag for SetPolicyProfile, got nil")
	}

	// Test SetDescription with non-existent tag
	err = policyTag.SetDescription(ctx, "non-existent-tag", "description")
	if err == nil {
		t.Error("Expected error with non-existent tag for SetDescription, got nil")
	}
}

func TestWlanPolicyTagServiceUnit_GetOperations_EdgeCases(t *testing.T) {
	// Mock server with various edge case responses
	tests := []struct {
		name         string
		mockResponse string
		expectEmpty  bool
	}{
		{
			name:         "NilResult",
			mockResponse: `null`,
			expectEmpty:  true,
		},
		{
			name: "EmptyPolicyListEntries",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": null
			}`,
			expectEmpty: true,
		},
		{
			name: "EmptyPolicyListEntry",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
					"policy-list-entry": null
				}
			}`,
			expectEmpty: true,
		},
		{
			name: "EmptyArray",
			mockResponse: `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
					"policy-list-entry": []
				}
			}`,
			expectEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
				"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": tt.mockResponse,
			}))
			defer mockServer.Close()

			client := testutil.NewTestClient(mockServer)
			service := NewService(client.Core().(*core.Client))
			policyTag := service.PolicyTag()
			ctx := testutil.TestContext(t)

			tags, err := policyTag.ListPolicyTags(ctx)
			if err != nil {
				t.Errorf("ListPolicyTags failed: %v", err)
				return
			}

			if tt.expectEmpty && len(tags) != 0 {
				t.Errorf("Expected empty tags list for %s, got %d tags", tt.name, len(tags))
			}
		})
	}
}

func TestWlanPolicyTagServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := NewService(nil)
	policyTag := service.PolicyTag()
	ctx := testutil.TestContext(t)

	_, err := policyTag.GetPolicyTag(ctx, "test")
	if err == nil {
		t.Error("Expected error with nil client for GetPolicyTag")
	}

	_, err = policyTag.ListPolicyTags(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListPolicyTags")
	}

	err = policyTag.CreatePolicyTag(ctx, nil)
	if err == nil {
		t.Error("Expected error with nil client for CreatePolicyTag")
	}

	err = policyTag.SetPolicyTag(ctx, nil)
	if err == nil {
		t.Error("Expected error with nil client for SetPolicyTag")
	}

	err = policyTag.SetPolicyProfile(ctx, "test", "wlan", "profile")
	if err == nil {
		t.Error("Expected error with nil client for SetPolicyProfile")
	}

	err = policyTag.SetDescription(ctx, "test", "desc")
	if err == nil {
		t.Error("Expected error with nil client for SetDescription")
	}

	err = policyTag.DeletePolicyTag(ctx, "test")
	if err == nil {
		t.Error("Expected error with nil client for DeletePolicyTag")
	}
}

func TestWlanPolicyTagServiceUnit_ValidationErrors_AdvancedScenarios(t *testing.T) {
	t.Parallel()

	t.Run("ValidationErrors_TagName", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test validateTagName with whitespace-only string
		_, err := policyTagService.GetPolicyTag(ctx, "   ")
		if err == nil {
			t.Error("Expected validation error for whitespace-only tag name")
		}

		err = policyTagService.DeletePolicyTag(ctx, "   ")
		if err == nil {
			t.Error("Expected validation error for whitespace-only tag name in DeletePolicyTag")
		}
	})

	t.Run("GetPolicyTag_NonExistent", func(t *testing.T) {
		// Mock server that returns empty result for non-existent tag
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
					"policy-list-entry": []
				}
			}`,
		}

		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test GetPolicyTag with non-existent tag should return nil
		result, err := policyTagService.GetPolicyTag(ctx, "non-existent")
		if err != nil {
			t.Errorf("GetPolicyTag should not error for non-existent tag: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for non-existent policy tag")
		}
	})

	t.Run("ListPolicyTags_EmptyResponse", func(t *testing.T) {
		// Mock server that returns empty array
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
					"policy-list-entry": []
				}
			}`,
		}

		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test ListPolicyTags with empty response
		results, err := policyTagService.ListPolicyTags(ctx)
		if err != nil {
			t.Errorf("ListPolicyTags should not error with empty response: %v", err)
		}
		if len(results) != 0 {
			t.Errorf("Expected empty results, got %d items", len(results))
		}
	})

	t.Run("ListPolicyTags_NilResponse", func(t *testing.T) {
		// Mock server that returns null policy-list-entries
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": null
			}`,
		}

		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test ListPolicyTags with null response
		results, err := policyTagService.ListPolicyTags(ctx)
		if err != nil {
			t.Errorf("ListPolicyTags should not error with null response: %v", err)
		}
		if len(results) != 0 {
			t.Errorf("Expected empty results for null response, got %d items", len(results))
		}
	})

	t.Run("SetPolicyProfile_GetError", func(t *testing.T) {
		// Mock server that returns error for GET operations (ListPolicyTags call)
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries",
		}, 500))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test SetPolicyProfile with server error during tag retrieval
		err := policyTagService.SetPolicyProfile(ctx, "test-tag", "wlan", "new-profile")
		if err == nil {
			t.Error("Expected error for SetPolicyProfile with server error")
		}
	})

	t.Run("SetPolicyProfile_TagNotFound", func(t *testing.T) {
		// Mock server that returns empty result for non-existent tag
		mockResponses := map[string]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
				"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
					"policy-list-entry": []
				}
			}`,
		}

		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(mockResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test SetPolicyProfile with non-existent tag
		err := policyTagService.SetPolicyProfile(ctx, "non-existent", "wlan", "profile")
		if err == nil {
			t.Error("Expected error for SetPolicyProfile with non-existent tag")
		}
	})

	t.Run("DeletePolicyTag_ValidationError", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test DeletePolicyTag with empty tag name
		err := policyTagService.DeletePolicyTag(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty tag name")
		}
	})

	t.Run("DeletePolicyTag_ServerError", func(t *testing.T) {
		// Mock server that returns error for DELETE operations
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry=test-tag",
		}, 500))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		policyTagService := service.PolicyTag()
		ctx := testutil.TestContext(t)

		// Test DeletePolicyTag with server error
		err := policyTagService.DeletePolicyTag(ctx, "test-tag")
		if err == nil {
			t.Error("Expected error for DeletePolicyTag with server error")
		}
	})
}
