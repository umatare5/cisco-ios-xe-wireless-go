package mobility_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mobility"
)

// TestMobilityServiceUnit_Constructor_Success tests service constructor with different client scenarios.
func TestMobilityServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create test server and client
		server := testutil.NewMockServer(map[string]string{})
		defer server.Close()

		testClient := testutil.NewTestClient(server)
		service := mobility.NewService(testClient.Core().(*core.Client))

		// Verify service creation
		if service.Client() == nil {
			t.Error("Expected valid client, got nil")
		}
		if service.Client() != testClient.Core().(*core.Client) {
			t.Error("Expected client to match provided client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := mobility.NewService(nil)

		// Verify service creation (should handle nil gracefully)
		if service.Client() != nil {
			t.Error("Expected nil client, got non-nil")
		}
	})
}

// TestMobilityServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
func TestMobilityServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock server and service
	mockServer := testutil.NewMockServer(map[string]string{})
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := mobility.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	result, err := service.GetOperational(ctx)

	// Verify results based on mobility service implementation - expects HTTP 404 from mock server
	if err == nil {
		t.Error("Expected error for GetOperational, got nil")
	}
	if result != nil {
		t.Error("Expected nil result for GetOperational, got non-nil result")
	}
}

// TestMobilityServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestMobilityServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create test server and service
	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := mobility.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test error scenarios
	result, err := service.GetOperational(ctx)

	// Verify error handling
	if err == nil {
		t.Error("Expected error for GetOperational, got nil")
	}
	if result != nil {
		t.Error("Expected nil result on error, got non-nil result")
	}
}

// TestMobilityServiceUnit_ListOperations_MockSuccess tests List operations using mock server.
func TestMobilityServiceUnit_ListOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC mobility data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-cache": `{
			"Cisco-IOS-XE-wireless-mobility-oper:ap-cache": [{
				"ap-mac": "28:ac:9e:bb:3c:80",
				"ap-name": "TEST-AP01"
			}]
		}`,
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-peer-list": `{
			"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list": [{
				"peer-ip": "192.168.255.1",
				"peer-status": "up"
			}]
		}`,
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-global-data": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data": {
				"mm-mac-addr": "00:1e:49:96:4c:ff"
			}
		}`,
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-if-global-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats": {
				"mblty-stats": {
					"event-data-allocs": 1434,
					"event-data-frees": 1434,
					"intra-wncd-roam-count": 85
				},
				"mblty-domain-info": {
					"mobility-domain-id": 54441
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-client-data": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data": [{
				"client-mac": "aa:bb:cc:dd:ee:ff",
				"mobility-state": "local"
			}]
		}`,
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats": {
				"total-handoffs": 0,
				"successful-handoffs": 0,
				"failed-handoffs": 0
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := mobility.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("ListAPCache", func(t *testing.T) {
		result, err := service.ListAPCache(ctx)
		if err != nil {
			t.Errorf("ListAPCache returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListAPCache returned nil result")
		}
	})

	t.Run("ListAPPeers", func(t *testing.T) {
		result, err := service.ListAPPeers(ctx)
		if err != nil {
			t.Errorf("ListAPPeers returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListAPPeers returned nil result")
		}
	})

	t.Run("GetMMGlobalInfo", func(t *testing.T) {
		result, err := service.GetMMGlobalInfo(ctx)
		if err != nil {
			t.Errorf("GetMMGlobalInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMMGlobalInfo returned nil result")
		}
	})

	t.Run("GetMMIFGlobalStats", func(t *testing.T) {
		result, err := service.GetMMIFGlobalStats(ctx)
		if err != nil {
			t.Errorf("GetMMIFGlobalStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMMIFGlobalStats returned nil result")
		}
	})

	t.Run("ListClients", func(t *testing.T) {
		result, err := service.ListClients(ctx)
		if err != nil {
			t.Errorf("ListClients returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListClients returned nil result")
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
}

// TestMobilityServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestMobilityServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("ListAPCache_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListAPCache(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetMMGlobalInfo_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetMMGlobalInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetGlobalStats_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetGlobalStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
