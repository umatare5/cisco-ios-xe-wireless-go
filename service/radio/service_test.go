package radio_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/radio"
)

// TestRadioServiceUnit_Constructor_Success tests service constructor with different client scenarios.
func TestRadioServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		server := testutil.NewMockServer(map[string]string{})
		defer server.Close()
		testClient := testutil.NewTestClient(server)
		service := radio.NewService(testClient.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected valid client, got nil")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := radio.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected nil client, got non-nil")
		}
	})
}

// TestRadioServiceUnit_GetConfigOperations_MockSuccess tests Get configuration operations using mock server.
func TestRadioServiceUnit_GetConfigOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC radio data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data": `{
			"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data": {
				"radio-profiles": {
					"radio-profile": [
						{
							"name": "default-radio-profile",
							"desc": "Preconfigured default radio profile",
							"mesh-backhaul": false
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data/radio-profiles": `{
			"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles": {
				"radio-profile": [
					{
						"name": "default-radio-profile",
						"desc": "Preconfigured default radio profile",
						"mesh-backhaul": false
					}
				]
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := radio.NewService(testClient.Core().(*core.Client))
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

	t.Run("ListProfileConfigs", func(t *testing.T) {
		result, err := service.ListProfileConfigs(ctx)
		if err != nil {
			t.Errorf("ListProfileConfigs returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListProfileConfigs returned nil result")
		}
	})
}

// TestRadioServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestRadioServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := radio.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_404Error", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for GetConfig, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListProfileConfigs_404Error", func(t *testing.T) {
		result, err := service.ListProfileConfigs(ctx)
		if err == nil {
			t.Error("Expected error for ListProfileConfigs, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestRadioServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestRadioServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := radio.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListProfileConfigs_NilClient", func(t *testing.T) {
		service := radio.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListProfileConfigs(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
