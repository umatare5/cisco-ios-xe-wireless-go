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
		server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
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
	// Create comprehensive mock responses for all mobility endpoints
	responses := map[string]string{
		// Root configuration data
		"Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data": `{
			"Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data": {
				"mobility-config": {
					"local-group": "test-group",
					"mac-address": "aa:bb:cc:dd:ee:ff"
				}
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data/mobility-config": `{
			"Cisco-IOS-XE-wireless-mobility-cfg:mobility-config": {
				"local-group": "test-group",
				"mac-address": "aa:bb:cc:dd:ee:ff"
			}
		}`,

		// Root operational data
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data": {
				"ap-cache": [{"ap-mac-address": "aa:bb:cc:dd:ee:ff"}],
				"ap-peer-list": [{"peer-ip": "192.168.1.100"}],
				"mm-global-data": {"tunnel-count": 5},
				"mm-if-global-msg-stats": {"total-messages": 1000},
				"mm-if-global-stats": {"total-events": 500},
				"mobility-client-data": [{"client-mac": "11:22:33:44:55:66"}],
				"mobility-client-stats": [{"client-events": 10}],
				"mobility-global-dtls-stats": {"dtls-tunnels": 3},
				"mobility-global-msg-stats": {"messages-sent": 200},
				"mobility-global-stats": {"global-events": 150},
				"wlan-client-limit": [{"wlan-id": 1, "client-limit": 50}]
			}
		}`,

		// Individual endpoint responses
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-cache": `{
			"Cisco-IOS-XE-wireless-mobility-oper:ap-cache": [
				{"ap-mac-address": "aa:bb:cc:dd:ee:ff", "mobility-role": "anchor"}
			]
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-peer-list": `{
			"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list": [
				{"peer-ip": "192.168.1.100", "ap-count": 10}
			]
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-global-data": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data": {
				"tunnel-count": 5,
				"mobility-role": "anchor"
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-if-global-msg-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-msg-stats": {
				"total-messages": 1000,
				"successful-messages": 950
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-if-global-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats": {
				"total-events": 500,
				"successful-events": 480
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-client-data": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data": [
				{"client-mac": "11:22:33:44:55:66", "mobility-status": "local"}
			]
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-client-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-stats": {
				"mm-mblty-stats": {"mm-mblty-tx-pkts": 100, "mm-mblty-rx-pkts": 200},
				"ipc-stats": [],
				"dgram-stats": []
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-dtls-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-dtls-stats": {
				"dtls-tunnels": 3,
				"active-tunnels": 2
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-msg-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-msg-stats": {
				"messages-sent": 200,
				"messages-received": 180
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-stats": `{
			"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats": {
				"global-events": 150,
				"tunnel-events": 100
			}
		}`,

		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/wlan-client-limit": `{
			"Cisco-IOS-XE-wireless-mobility-oper:wlan-client-limit": [
				{"wlan-id": 1, "client-limit": 50}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := mobility.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test configuration functions
	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Fatalf("GetConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetConfig returned nil result")
		}
	})

	t.Run("ListMobilityConfig", func(t *testing.T) {
		result, err := service.ListMobilityConfig(ctx)
		if err != nil {
			t.Fatalf("ListMobilityConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListMobilityConfig returned nil result")
		}
	})

	// Test base GetOperational function
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Fatalf("GetOperational failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetOperational returned nil result")
		}
	})

	// Test existing List/Get functions
	t.Run("ListAPCache", func(t *testing.T) {
		result, err := service.ListAPCache(ctx)
		if err != nil {
			t.Fatalf("ListAPCache failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListAPCache returned nil result")
		}
	})

	t.Run("ListAPPeers", func(t *testing.T) {
		result, err := service.ListAPPeers(ctx)
		if err != nil {
			t.Fatalf("ListAPPeers failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListAPPeers returned nil result")
		}
	})

	t.Run("GetMMGlobalInfo", func(t *testing.T) {
		result, err := service.GetMMGlobalInfo(ctx)
		if err != nil {
			t.Fatalf("GetMMGlobalInfo failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetMMGlobalInfo returned nil result")
		}
	})

	t.Run("GetMMIFGlobalStats", func(t *testing.T) {
		result, err := service.GetMMIFGlobalStats(ctx)
		if err != nil {
			t.Fatalf("GetMMIFGlobalStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetMMIFGlobalStats returned nil result")
		}
	})

	t.Run("ListClients", func(t *testing.T) {
		result, err := service.ListClients(ctx)
		if err != nil {
			t.Fatalf("ListClients failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListClients returned nil result")
		}
	})

	t.Run("GetGlobalStats", func(t *testing.T) {
		result, err := service.GetGlobalStats(ctx)
		if err != nil {
			t.Fatalf("GetGlobalStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetGlobalStats returned nil result")
		}
	})

	// Test newly implemented List* functions
	t.Run("ListMmIfGlobalMsgStats", func(t *testing.T) {
		result, err := service.ListMmIfGlobalMsgStats(ctx)
		if err != nil {
			t.Fatalf("ListMmIfGlobalMsgStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListMmIfGlobalMsgStats returned nil result")
		}
	})

	t.Run("ListClientStats", func(t *testing.T) {
		result, err := service.ListClientStats(ctx)
		if err != nil {
			t.Fatalf("ListClientStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListClientStats returned nil result")
		}
	})

	t.Run("ListGlobalDTLSStats", func(t *testing.T) {
		result, err := service.ListGlobalDTLSStats(ctx)
		if err != nil {
			t.Fatalf("ListGlobalDTLSStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListGlobalDTLSStats returned nil result")
		}
	})

	t.Run("ListGlobalMsgStats", func(t *testing.T) {
		result, err := service.ListGlobalMsgStats(ctx)
		if err != nil {
			t.Fatalf("ListGlobalMsgStats failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListGlobalMsgStats returned nil result")
		}
	})

	t.Run("ListWlanClientLimit", func(t *testing.T) {
		result, err := service.ListWlanClientLimit(ctx)
		if err != nil {
			t.Fatalf("ListWlanClientLimit failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListWlanClientLimit returned nil result")
		}
	})
}

// TestMobilityServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestMobilityServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data",
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := mobility.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test error handling for configuration functions
	t.Run("GetConfig_404Error", func(t *testing.T) {
		_, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListMobilityConfig_404Error", func(t *testing.T) {
		_, err := service.ListMobilityConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	// Test error handling for operational functions
	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListAPCache_404Error", func(t *testing.T) {
		_, err := service.ListAPCache(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListMmIfGlobalMsgStats_404Error", func(t *testing.T) {
		_, err := service.ListMmIfGlobalMsgStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})
}

// TestMobilityServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestMobilityServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMobilityConfig_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListMobilityConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

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

	t.Run("ListMmIfGlobalMsgStats_NilClient", func(t *testing.T) {
		service := mobility.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListMmIfGlobalMsgStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}

// TestMobilityServiceUnit_ListOperations_MockSuccess tests List operations using mock server.
func TestMobilityServiceUnit_ListOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC mobility data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-cache": `{
			"Cisco-IOS-XE-wireless-mobility-oper:ap-cache": [{
				"ap-mac": "aa:bb:cc:dd:ee:ff",
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

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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
