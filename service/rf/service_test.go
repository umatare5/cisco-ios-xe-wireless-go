package rf

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := NewService(testClient.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected valid client, got nil")
	}
}

func TestRfServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC data from RF configuration
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
							"tag-name": "default-rf-tag",
							"description": "Preconfigured default RF tag"
						}
					]
				},
				"rf-profiles": {
					"rf-profile": [
						{
							"profile-name": "labo-rf-5gh-inside",
							"rf-band": "dot11-5ghz-band",
							"description": "RF profile for 5GHz indoor"
						}
					]
				},
				"multi-bssid-profiles": {
					"multi-bssid-profile": []
				},
				"atf-policies": {
					"atf-policy": []
				},
				"rf-profile-default-entries": {
					"rf-profile-default-entry": []
				}
			}
		}`,
		// Individual wrapper endpoints
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
						"tag-name": "default-rf-tag",
						"description": "Preconfigured default RF tag"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profiles": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-profiles": {
				"rf-profile": [
					{
						"profile-name": "labo-rf-5gh-inside",
						"rf-band": "dot11-5ghz-band",
						"description": "RF profile for 5GHz indoor"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/multi-bssid-profiles": `{
			"Cisco-IOS-XE-wireless-rf-cfg:multi-bssid-profiles": {
				"multi-bssid-profile": []
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/atf-policies": `{
			"Cisco-IOS-XE-wireless-rf-cfg:atf-policies": {
				"atf-policy": []
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profile-default-entries": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-profile-default-entries": {
				"rf-profile-default-entry": []
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
		}
	})

	t.Run("RFTag", func(t *testing.T) {
		rfTagService := service.RFTag()
		if rfTagService == nil {
			t.Error("RFTag service returned nil")
		}
	})
	t.Run("ListRFTags", func(t *testing.T) {
		result, err := service.ListRFTags(ctx)
		if err != nil {
			t.Errorf("ListRFTags returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFTags returned nil result")
		}
	})

	t.Run("ListRFProfiles", func(t *testing.T) {
		result, err := service.ListRFProfiles(ctx)
		if err != nil {
			t.Errorf("ListRFProfiles returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFProfiles returned nil result")
		}
	})

	t.Run("ListMultiBssidProfiles", func(t *testing.T) {
		result, err := service.ListMultiBssidProfiles(ctx)
		if err != nil {
			t.Errorf("ListMultiBssidProfiles returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMultiBssidProfiles returned nil result")
		}
	})

	t.Run("ListAtfPolicies", func(t *testing.T) {
		result, err := service.ListAtfPolicies(ctx)
		if err != nil {
			t.Errorf("ListAtfPolicies returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListAtfPolicies returned nil result")
		}
	})

	t.Run("ListRFProfileDefaultEntries", func(t *testing.T) {
		result, err := service.ListRFProfileDefaultEntries(ctx)
		if err != nil {
			t.Errorf("ListRFProfileDefaultEntries returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFProfileDefaultEntries returned nil result")
		}
	})
}

func TestRfServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFTags_NilClient", func(t *testing.T) {
		result, err := service.ListRFTags(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFProfiles_NilClient", func(t *testing.T) {
		result, err := service.ListRFProfiles(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMultiBssidProfiles_NilClient", func(t *testing.T) {
		result, err := service.ListMultiBssidProfiles(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListAtfPolicies_NilClient", func(t *testing.T) {
		result, err := service.ListAtfPolicies(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFProfileDefaultEntries_NilClient", func(t *testing.T) {
		result, err := service.ListRFProfileDefaultEntries(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
