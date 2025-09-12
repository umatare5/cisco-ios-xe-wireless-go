package site

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestSiteServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client")
	}

	// Test with valid client
	mockServer := testutil.NewMockServer(map[string]string{})
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service = NewService(client.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected service to have client, got nil")
	}
}

func TestSiteServiceUnit_GetOperations_Success(t *testing.T) {
	t.Parallel()

	mockResponses := map[string]string{
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data": {
				"ap-cfg-profiles": {"ap-cfg-profile": []},
				"site-tag-configs": {
					"site-tag-config": [{
						"site-tag-name": "default-site-tag",
						"description": "Default Site Tag for WNC",
						"ap-join-profile": "default-ap-policy-profile",
						"local-site": false
					}]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/ap-cfg-profiles": `{
			"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles": {"ap-cfg-profile": []}
		}`,
		"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs": `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
				"site-tag-config": [{
					"site-tag-name": "default-site-tag",
					"description": "Default Site Tag for WNC",
					"ap-join-profile": "default-ap-policy-profile",
					"local-site": false
				}]
			}
		}`,
	}

	mockServer := testutil.NewMockServer(mockResponses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		_, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
	})

	t.Run("ListAPProfileConfigs", func(t *testing.T) {
		_, err := service.ListAPProfileConfigs(ctx)
		if err != nil {
			t.Errorf("ListAPProfileConfigs returned unexpected error: %v", err)
		}
	})

	t.Run("ListSiteTagConfigs", func(t *testing.T) {
		_, err := service.ListSiteTagConfigs(ctx)
		if err != nil {
			t.Errorf("ListSiteTagConfigs returned unexpected error: %v", err)
		}
	})

	t.Run("SiteTag_Service", func(t *testing.T) {
		siteTagService := service.SiteTag()
		if siteTagService == nil {
			t.Error("SiteTag service should not be nil")
		}
	})
}

func TestSiteServiceUnit_ValidationErrors_EmptyInputs(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(map[string]string{})
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	// Basic service validation tests only - site tag tests moved to tag_service_test.go
	t.Run("Service_Constructor", func(t *testing.T) {
		if service.Client() == nil {
			t.Error("Expected service to have client, got nil")
		}
	})
}

func TestSiteServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		_, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error with nil client for GetConfig")
		}
	})

	t.Run("ListAPProfileConfigs_NilClient", func(t *testing.T) {
		_, err := service.ListAPProfileConfigs(ctx)
		if err == nil {
			t.Error("Expected error with nil client for ListAPProfileConfigs")
		}
	})

	t.Run("ListSiteTagConfigs_NilClient", func(t *testing.T) {
		_, err := service.ListSiteTagConfigs(ctx)
		if err == nil {
			t.Error("Expected error with nil client for ListSiteTagConfigs")
		}
	})
}
