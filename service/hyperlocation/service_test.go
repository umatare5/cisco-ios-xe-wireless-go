package hyperlocation_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/hyperlocation"
)

// TestHyperlocationServiceUnit_Constructor_Success tests service constructor functionality.
func TestHyperlocationServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := hyperlocation.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := hyperlocation.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestHyperlocationServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestHyperlocationServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Hyperlocation endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data": `{
			"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data": {
				"ewlc-hyperlocation-profile": [
					{
						"name": "labo-common",
						"hyperlocation-data": {
							"hyperlocation-enable": true,
							"pak-rssi-threshold-detection": -100,
							"pak-rssi-threshold-trigger": 10,
							"pak-rssi-threshold-reset": 8
						},
						"ntp-server": "0.0.0.0",
						"status": false,
						"reason-down": "hyperlocation-reason-ntp"
					},
					{
						"name": "default-ap-profile",
						"hyperlocation-data": {
							"hyperlocation-enable": false,
							"pak-rssi-threshold-detection": -100,
							"pak-rssi-threshold-trigger": 10,
							"pak-rssi-threshold-reset": 8
						},
						"ntp-server": "0.0.0.0",
						"status": false,
						"reason-down": "hyperlocation-reason-disabled"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data/ewlc-hyperlocation-profile": `{
			"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile": [
				{
					"name": "labo-common",
					"hyperlocation-data": {
						"hyperlocation-enable": true,
						"pak-rssi-threshold-detection": -100,
						"pak-rssi-threshold-trigger": 10,
						"pak-rssi-threshold-reset": 8
					},
					"ntp-server": "0.0.0.0",
					"status": false,
					"reason-down": "hyperlocation-reason-ntp"
				},
				{
					"name": "default-ap-profile",
					"hyperlocation-data": {
						"hyperlocation-enable": false,
						"pak-rssi-threshold-detection": -100,
						"pak-rssi-threshold-trigger": 10,
						"pak-rssi-threshold-reset": 8
					},
					"ntp-server": "0.0.0.0",
					"status": false,
					"reason-down": "hyperlocation-reason-disabled"
				}
			]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := hyperlocation.NewService(testClient.Core().(*core.Client))
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

	// Test ListProfiles operation
	t.Run("ListProfiles", func(t *testing.T) {
		result, err := service.ListProfiles(ctx)
		if err != nil {
			t.Errorf("Expected no error for mock ListProfiles, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for mock ListProfiles, got nil")
		}
	})
}

// TestHyperlocationServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestHyperlocationServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Hyperlocation endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data",
		"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data/ewlc-hyperlocation-profile",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := hyperlocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}

		// Verify error contains expected information
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	// Test that ListProfiles properly handles 404 errors
	t.Run("ListProfiles_404Error", func(t *testing.T) {
		_, err := service.ListProfiles(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}

		// Verify error contains expected information
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}
