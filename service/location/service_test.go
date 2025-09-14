package location_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/location"
)

// TestLocationServiceUnit_Constructor_Success tests service constructor functionality.
func TestLocationServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		testResponses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(testResponses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := location.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := location.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestLocationServiceUnit_GetConfigOperations_MockSuccess tests Get configuration operations using mock server.
func TestLocationServiceUnit_GetConfigOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC location data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data": `{
			"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data": {
				"nmsp-config": {}
			}
		}`,
		"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/operator-locations": `{
			"Cisco-IOS-XE-wireless-location-cfg:operator-locations": []
		}`,
		"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/nmsp-config": `{
			"Cisco-IOS-XE-wireless-location-cfg:nmsp-config": {}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := location.NewService(testClient.Core().(*core.Client))
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

	t.Run("ListOperatorLocations", func(t *testing.T) {
		result, err := service.ListOperatorLocations(ctx)
		if err != nil {
			t.Errorf("ListOperatorLocations returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListOperatorLocations returned nil result")
		}
	})

	t.Run("ListServerConfigs", func(t *testing.T) {
		result, err := service.ListNmspConfig(ctx)
		if err != nil {
			t.Errorf("ListServerConfigs returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListServerConfigs returned nil result")
		}
	})
}

// TestLocationServiceUnit_GetConfigOperations_MockNoContentSuccess tests operations that return HTTP 204 (No Content) responses.
func TestLocationServiceUnit_GetConfigOperations_MockNoContentSuccess(t *testing.T) {
	t.Parallel()

	mockServerNoContent := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/location",
			testutil.ResponseConfig{
				StatusCode: 204,
				Body:       "",
			}),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-location-oper:location-oper-data",
			testutil.ResponseConfig{
				StatusCode: 204,
				Body:       "",
			}),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-location-oper:location-oper-data/location-rssi-measurements",
			testutil.ResponseConfig{
				StatusCode: 204,
				Body:       "",
			}),
	)
	defer mockServerNoContent.Close()

	testClient := testutil.NewTestClient(mockServerNoContent)
	serviceNoContent := location.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetLocation_NoContent", func(t *testing.T) {
		result, err := serviceNoContent.GetLocation(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetLocation, got: %v", err)
		}
		if result == nil {
			t.Error("Expected non-nil result for GetLocation")
		} else if result.LocationConfig != nil {
			t.Error("Expected nil LocationConfig for HTTP 204 response")
		}
	})

	t.Run("GetOperational_NoContent", func(t *testing.T) {
		result, err := serviceNoContent.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected non-nil result for GetOperational")
		} else if result.CiscoIOSXEWirelessLocationOperData != nil {
			t.Error("Expected nil CiscoIOSXEWirelessLocationOperData for HTTP 204 response")
		}
	})

	t.Run("LocationRssiMeasurements_NoContent", func(t *testing.T) {
		result, err := serviceNoContent.LocationRssiMeasurements(ctx)
		if err != nil {
			t.Errorf("Expected no error for LocationRssiMeasurements, got: %v", err)
		}
		if result == nil {
			t.Error("Expected non-nil result for LocationRssiMeasurements")
		} else if len(result.LocationRssiMeasurements) > 0 {
			t.Error("Expected empty LocationRssiMeasurements for HTTP 204 response")
		}
	})
}

// TestLocationServiceUnit_GetOperations_ErrorHandling tests error scenarios for implemented operations.
func TestLocationServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	errorServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer errorServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(errorServer)
	service := location.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_404Error", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("GetConfig should return error for 404")
		}
		if result != nil {
			t.Error("GetConfig should return nil result on error")
		}
	})

	// Note: ListProfileConfigs and ListServerConfigs with empty responses
	// are handled by HTTP 204 (No Content) in live environment,
	// which core.Get handles gracefully by returning empty structs.
}

// TestLocationServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestLocationServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("ListOperatorLocations_NilClient", func(t *testing.T) {
		service := location.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListOperatorLocations(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListNmspConfig_NilClient", func(t *testing.T) {
		service := location.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListNmspConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := location.NewService(nil)
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
