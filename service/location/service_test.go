package location_test

import (
	"errors"
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
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(responses)
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
		"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/operator-locations": `{
			"Cisco-IOS-XE-wireless-location-cfg:operator-locations": []
		}`,
		"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/nmsp-config": `{
			"Cisco-IOS-XE-wireless-location-cfg:nmsp-config": {}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := location.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("ListProfileConfigs", func(t *testing.T) {
		result, err := service.ListProfileConfigs(ctx)
		if err != nil {
			t.Errorf("ListProfileConfigs returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListProfileConfigs returned nil result")
		}
	})

	t.Run("ListServerConfigs", func(t *testing.T) {
		result, err := service.ListServerConfigs(ctx)
		if err != nil {
			t.Errorf("ListServerConfigs returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListServerConfigs returned nil result")
		}
	})
}

// TestLocationServiceUnit_GetOperations_ErrorHandling tests error scenarios for implemented operations.
func TestLocationServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := location.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("ListProfileConfigs_404Error", func(t *testing.T) {
		result, err := service.ListProfileConfigs(ctx)
		if err == nil {
			t.Error("Expected error for ListProfileConfigs, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListServerConfigs_404Error", func(t *testing.T) {
		result, err := service.ListServerConfigs(ctx)
		if err == nil {
			t.Error("Expected error for ListServerConfigs, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestLocationServiceUnit_NotImplementedOperations_ResourceNotFound tests operations that return ErrResourceNotFound.
func TestLocationServiceUnit_NotImplementedOperations_ResourceNotFound(t *testing.T) {
	t.Parallel()

	// Create mock server and test client
	responses := map[string]string{
		"test-endpoint": `{"status": "success"}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := location.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NotImplemented", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for GetConfig, got nil")
		}
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected ErrResourceNotFound, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for not implemented operation")
		}
	})

	t.Run("GetSettingsConfig_NotImplemented", func(t *testing.T) {
		result, err := service.GetSettingsConfig(ctx)
		if err == nil {
			t.Error("Expected error for GetSettingsConfig, got nil")
		}
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected ErrResourceNotFound, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for not implemented operation")
		}
	})

	t.Run("GetOperational_NotImplemented", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
		}
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected ErrResourceNotFound, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for not implemented operation")
		}
	})

	t.Run("GetStats_NotImplemented", func(t *testing.T) {
		result, err := service.GetStats(ctx)
		if err == nil {
			t.Error("Expected error for GetStats, got nil")
		}
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected ErrResourceNotFound, got: %v", err)
		}
		if result != nil {
			t.Error("Expected nil result for not implemented operation")
		}
	})
}

// TestLocationServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestLocationServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("ListProfileConfigs_NilClient", func(t *testing.T) {
		service := location.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListProfileConfigs(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListServerConfigs_NilClient", func(t *testing.T) {
		service := location.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListServerConfigs(ctx)
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
