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

// TestMeshServiceUnit_GetOperations_MockSuccess tests Get configuration and operational operations using a mock server.
func TestMeshServiceUnit_GetOperations_MockSuccess(t *testing.T) {
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
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data": `{
			"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data": {
				"mesh-q-stats": [
					{
						"wtp-mac": "00:11:22:33:44:55",
						"q-type": "data",
						"peak-length": 100,
						"average-len": 50,
						"overflows": 2
					}
				],
				"mesh-dr-stats": [
					{
						"wtp-mac": "00:11:22:33:44:55",
						"neigh-ap-mac": "00:66:77:88:99:aa",
						"data-rate-index": 1,
						"tx-success": 1000,
						"tx-attempts": 1050
					}
				],
				"mesh-sec-stats": [
					{
						"wtp-mac": "00:11:22:33:44:55",
						"tx-pkts-total": 5000,
						"rx-pkts-total": 4800,
						"rx-pkts-error": 5
					}
				],
				"mesh-oper-data": [
					{
						"wtp-mac": "00:11:22:33:44:55",
						"bhaul-slot-id": 0,
						"configured-role": "MAP",
						"ap-mode": "bridge"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-q-stats": `{
			"Cisco-IOS-XE-wireless-mesh-oper:mesh-q-stats": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"q-type": "data",
					"peak-length": 100,
					"average-len": 50,
					"overflows": 2
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-dr-stats": `{
			"Cisco-IOS-XE-wireless-mesh-oper:mesh-dr-stats": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"neigh-ap-mac": "00:66:77:88:99:aa",
					"data-rate-index": 1,
					"tx-success": 1000,
					"tx-attempts": 1050
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-sec-stats": `{
			"Cisco-IOS-XE-wireless-mesh-oper:mesh-sec-stats": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"tx-pkts-total": 5000,
					"rx-pkts-total": 4800,
					"rx-pkts-error": 5
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-oper-data": `{
			"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"bhaul-slot-id": 0,
					"configured-role": "MAP",
					"ap-mode": "bridge"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-global-stats": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-stats": {
				"num-of-bridge-aps": 5,
				"num-of-maps": 10,
				"num-of-raps": 3,
				"num-of-flex-bridge-aps": 2,
				"num-of-flex-bridge-maps": 8,
				"num-of-flex-bridge-raps": 2
			}
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-cac-info": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-cac-info": [
				{
					"wtp-name": "TEST-AP01"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-path-info": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-path-info": [
				{
					"wtp-name": "TEST-AP02"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-tree-data": `{
			"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-tree-data": [
				{
					"sector-number": 1,
					"wtp-mac": "00:11:22:33:44:55",
					"mesh-ap-count": 5,
					"rap-count": 2,
					"map-count": 3,
					"mesh-ap-list": [
						{
							"ap-name": "TEST-AP03",
							"ap-role": "MAP",
							"bridge-group-name": "default",
							"pref-parent-ap-name": "TEST-RAP01"
						}
					]
				}
			]
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

	t.Run("ListMeshQueueStats", func(t *testing.T) {
		result, err := service.ListMeshQueueStats(ctx)
		if err != nil {
			t.Errorf("ListMeshQueueStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMeshQueueStats returned nil result")
		}
	})

	t.Run("ListMeshDataRateStats", func(t *testing.T) {
		result, err := service.ListMeshDataRateStats(ctx)
		if err != nil {
			t.Errorf("ListMeshDataRateStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMeshDataRateStats returned nil result")
		}
	})

	t.Run("ListMeshSecurityStats", func(t *testing.T) {
		result, err := service.ListMeshSecurityStats(ctx)
		if err != nil {
			t.Errorf("ListMeshSecurityStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMeshSecurityStats returned nil result")
		}
	})

	t.Run("ListMeshOperationalData", func(t *testing.T) {
		result, err := service.ListMeshOperationalData(ctx)
		if err != nil {
			t.Errorf("ListMeshOperationalData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMeshOperationalData returned nil result")
		}
	})

	t.Run("GetGlobalStats", func(t *testing.T) {
		result, err := service.GetGlobalStats(ctx)
		if err != nil {
			t.Errorf("GetGlobalStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalStats returned nil result")
		}
	})

	t.Run("ListApCacInfo", func(t *testing.T) {
		result, err := service.ListApCacInfo(ctx)
		if err != nil {
			t.Errorf("ListApCacInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApCacInfo returned nil result")
		}
	})

	t.Run("ListApPathInfo", func(t *testing.T) {
		result, err := service.ListApPathInfo(ctx)
		if err != nil {
			t.Errorf("ListApPathInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApPathInfo returned nil result")
		}
	})

	t.Run("ListApTreeData", func(t *testing.T) {
		result, err := service.ListApTreeData(ctx)
		if err != nil {
			t.Errorf("ListApTreeData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApTreeData returned nil result")
		}
	})
}

// TestMeshServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestMeshServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service with error responses
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data",
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data",
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-global-stats",
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-q-stats",
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-dr-stats",
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-sec-stats",
		"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-oper-data",
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-cac-info",
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-path-info",
		"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-tree-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
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

	t.Run("ListMeshQueueStats_404Error", func(t *testing.T) {
		result, err := service.ListMeshQueueStats(ctx)
		if err == nil {
			t.Error("Expected error for ListMeshQueueStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListMeshDataRateStats_404Error", func(t *testing.T) {
		result, err := service.ListMeshDataRateStats(ctx)
		if err == nil {
			t.Error("Expected error for ListMeshDataRateStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListMeshSecurityStats_404Error", func(t *testing.T) {
		result, err := service.ListMeshSecurityStats(ctx)
		if err == nil {
			t.Error("Expected error for ListMeshSecurityStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListMeshOperationalData_404Error", func(t *testing.T) {
		result, err := service.ListMeshOperationalData(ctx)
		if err == nil {
			t.Error("Expected error for ListMeshOperationalData, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetGlobalStats_404Error", func(t *testing.T) {
		result, err := service.GetGlobalStats(ctx)
		if err == nil {
			t.Error("Expected error for GetGlobalStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListApCacInfo_404Error", func(t *testing.T) {
		result, err := service.ListApCacInfo(ctx)
		if err == nil {
			t.Error("Expected error for ListApCacInfo, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListApPathInfo_404Error", func(t *testing.T) {
		result, err := service.ListApPathInfo(ctx)
		if err == nil {
			t.Error("Expected error for ListApPathInfo, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListApTreeData_404Error", func(t *testing.T) {
		result, err := service.ListApTreeData(ctx)
		if err == nil {
			t.Error("Expected error for ListApTreeData, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestMeshServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestMeshServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := mesh.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperationalData_NilClient", func(t *testing.T) {
		result, err := service.GetOperationalData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMeshQueueStats_NilClient", func(t *testing.T) {
		result, err := service.ListMeshQueueStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMeshDataRateStats_NilClient", func(t *testing.T) {
		result, err := service.ListMeshDataRateStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMeshSecurityStats_NilClient", func(t *testing.T) {
		result, err := service.ListMeshSecurityStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMeshOperationalData_NilClient", func(t *testing.T) {
		result, err := service.ListMeshOperationalData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetGlobalStats_NilClient", func(t *testing.T) {
		result, err := service.GetGlobalStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListApCacInfo_NilClient", func(t *testing.T) {
		result, err := service.ListApCacInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListApPathInfo_NilClient", func(t *testing.T) {
		result, err := service.ListApPathInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListApTreeData_NilClient", func(t *testing.T) {
		result, err := service.ListApTreeData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
