package ble_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ble"
)

// TestBleServiceUnit_Constructor_Success tests service constructor functionality.
func TestBleServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := ble.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := ble.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestBleServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestBleServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with BLE endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data": `{
			"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data": {
				"ble-ltx-ap-antenna": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:f0",
						"ble-slot-id": 0,
						"ble-antenna-id": 1,
						"is-ble-antenna-present": true
					}
				],
				"ble-ltx-ap": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:f0",
						"admin": {
							"enable": true
						}
					}
				]
			}
		}`,
		// Individual wrapper endpoints for BLE LTX operations
		"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap": `{
			"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:f0",
					"admin": {
						"enable": true
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ble-ltx-ap-antenna": `{
			"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:f0",
					"ble-slot-id": 0,
					"ble-antenna-id": 1,
					"is-ble-antenna-present": true
				}
			]
		}`,
		// BLE Management endpoints
		"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data": `{
			"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data": {
				"ble-mgmt-ap": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:f0",
						"is-new": false,
						"cmx-id": 1,
						"oper-state": true
					}
				],
				"ble-mgmt-cmx": [
					{
						"cmx-id": 1,
						"oper-state": true,
						"admin-state": true
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data/ble-mgmt-ap": `{
			"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-ap": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:f0",
					"is-new": false,
					"cmx-id": 1,
					"oper-state": true
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-oper-data/ble-mgmt-cmx": `{
			"Cisco-IOS-XE-wireless-ble-mgmt-oper:ble-mgmt-cmx": [
				{
					"cmx-id": 1,
					"oper-state": true,
					"admin-state": true
				}
			]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := ble.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock GetOperational, got nil")
		}
	})

	t.Run("ListBLELtxAp", func(t *testing.T) {
		result, err := service.ListBLELtxAp(ctx)
		if err != nil {
			t.Errorf("ListBLELtxAp returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListBLELtxAp returned nil result")
		}
	})

	t.Run("ListBLELtxApAntenna", func(t *testing.T) {
		result, err := service.ListBLELtxApAntenna(ctx)
		if err != nil {
			t.Errorf("ListBLELtxApAntenna returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListBLELtxApAntenna returned nil result")
		}
	})

	t.Run("GetMgmtOperational", func(t *testing.T) {
		result, err := service.GetMgmtOperational(ctx)
		if err != nil {
			t.Errorf("GetMgmtOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMgmtOperational returned nil result")
		}
	})

	t.Run("ListBLEMgmtAp", func(t *testing.T) {
		result, err := service.ListBLEMgmtAp(ctx)
		if err != nil {
			t.Errorf("ListBLEMgmtAp returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListBLEMgmtAp returned nil result")
		}
	})

	t.Run("ListBLEMgmtCmx", func(t *testing.T) {
		result, err := service.ListBLEMgmtCmx(ctx)
		if err != nil {
			t.Errorf("ListBLEMgmtCmx returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListBLEMgmtCmx returned nil result")
		}
	})
}

// TestBleServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestBleServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for BLE endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ble.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListBLELtxAp_404Error", func(t *testing.T) {
		_, err := service.ListBLELtxAp(ctx)
		if err == nil {
			t.Error("Expected error for ListBLELtxAp, got nil")
		}
	})

	t.Run("ListBLELtxApAntenna_404Error", func(t *testing.T) {
		_, err := service.ListBLELtxApAntenna(ctx)
		if err == nil {
			t.Error("Expected error for ListBLELtxApAntenna, got nil")
		}
	})

	t.Run("GetMgmtOperational_404Error", func(t *testing.T) {
		_, err := service.GetMgmtOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetMgmtOperational, got nil")
		}
	})

	t.Run("ListBLEMgmtAp_404Error", func(t *testing.T) {
		_, err := service.ListBLEMgmtAp(ctx)
		if err == nil {
			t.Error("Expected error for ListBLEMgmtAp, got nil")
		}
	})

	t.Run("ListBLEMgmtCmx_404Error", func(t *testing.T) {
		_, err := service.ListBLEMgmtCmx(ctx)
		if err == nil {
			t.Error("Expected error for ListBLEMgmtCmx, got nil")
		}
	})
}

// TestBleServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestBleServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := ble.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListBLELtxAp_NilClient", func(t *testing.T) {
		result, err := service.ListBLELtxAp(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListBLELtxApAntenna_NilClient", func(t *testing.T) {
		result, err := service.ListBLELtxApAntenna(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetMgmtOperational_NilClient", func(t *testing.T) {
		result, err := service.GetMgmtOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListBLEMgmtAp_NilClient", func(t *testing.T) {
		result, err := service.ListBLEMgmtAp(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListBLEMgmtCmx_NilClient", func(t *testing.T) {
		result, err := service.ListBLEMgmtCmx(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
