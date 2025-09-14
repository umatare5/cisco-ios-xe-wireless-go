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
				"ble-ltx-summary": {
					"total-ble-beacons": 10,
					"enabled-interfaces": 2
				}
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := ble.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock GetOperational, got: %v", err)
	}
	if result == nil {
		t.Error("Expected result for mock GetOperational, got nil")
	}
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
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}
}
