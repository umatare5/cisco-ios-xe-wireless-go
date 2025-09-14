package mesh_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mesh"
)

// TestMeshServiceUnit_Constructor_Success tests service constructor functionality.
func TestMeshServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := mesh.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := mesh.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestMeshServiceUnit_GetConfigOperations_MockSuccess tests Get configuration operations using mock server.
func TestMeshServiceUnit_GetConfigOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC mesh data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data": `{
			"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data": {
				"mesh": {},
				"mesh-profiles": {
					"mesh-profile": [
						{
							"profile-name": "default-mesh-profile",
							"description": "Preconfigured default mesh profile"
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data": {
				"mesh-stats": {
					"total-packets": 0,
					"total-bytes": 0
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-global-stats": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-stats": {
				"queue-stats": {
					"tx-packets": 0,
					"rx-packets": 0
				}
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := mesh.NewService(testClient.Core().(*core.Client))
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

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
		}
	})

	t.Run("GetOperationalData", func(t *testing.T) {
		result, err := service.GetOperationalData(ctx)
		if err != nil {
			t.Errorf("GetOperationalData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperationalData returned nil result")
		}
	})
}

// TestMeshServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestMeshServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close() // Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := mesh.NewService(testClient.Core().(*core.Client))
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

	t.Run("GetOperational_404Error", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetOperationalData_404Error", func(t *testing.T) {
		result, err := service.GetOperationalData(ctx)
		if err == nil {
			t.Error("Expected error for GetOperationalData, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestMeshServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestMeshServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := mesh.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		service := mesh.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperationalData_NilClient", func(t *testing.T) {
		service := mesh.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperationalData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
