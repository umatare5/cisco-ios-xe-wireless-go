package rf

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := NewService(testClient.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected valid client, got nil")
	}
}

func TestRfServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC data from tmp/live_rf_cfg_data.json
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
				}
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
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
}

func TestRfServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
