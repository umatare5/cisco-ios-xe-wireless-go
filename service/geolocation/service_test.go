package geolocation_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/geolocation"
)

// TestGeolocationServiceUnit_Constructor_Success tests service constructor functionality.
func TestGeolocationServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(responses)
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := geolocation.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := geolocation.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestGeolocationServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestGeolocationServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Geolocation endpoints based on live WNC data
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data": {
				"ap-geo-loc-stats": {
					"num-ap-gnss": 0,
					"num-ap-man-height": 0,
					"num-ap-derived": 0,
					"last-derivation-timestamp": "2025-09-10T17:04:29.717868+00:00"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data/ap-geo-loc-stats": `{
			"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats": {
				"num-ap-gnss": 0,
				"num-ap-man-height": 0,
				"num-ap-derived": 0,
				"last-derivation-timestamp": "2025-09-10T17:04:29.717868+00:00"
			}
		}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock GetOperational, got: %v", err)
	}
	if result == nil {
		t.Error("Expected result for mock GetOperational, got nil")
	}

	// Test ListAPGeolocationStats operation
	statsResult, err := service.ListAPGeolocationStats(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock ListAPGeolocationStats, got: %v", err)
	}
	if statsResult == nil {
		t.Error("Expected result for mock ListAPGeolocationStats, got nil")
	}
}

// TestGeolocationServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestGeolocationServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Geolocation endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data",
	}
	mockServer := testutil.NewMockErrorServer(errorPaths, 404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := geolocation.NewService(testClient.Core().(*core.Client))
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
