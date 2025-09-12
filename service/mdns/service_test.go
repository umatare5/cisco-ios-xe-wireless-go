package mdns_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/mdns"
)

// TestMdnsServiceUnit_Constructor_Success tests service constructor functionality.
func TestMdnsServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create test server and client
		server := testutil.NewMockServer(map[string]string{})
		defer server.Close()

		testClient := testutil.NewTestClient(server)
		service := mdns.NewService(testClient.Core().(*core.Client))

		// Verify service creation
		if service.Client() == nil {
			t.Error("Expected valid client, got nil")
		}
		if service.Client() != testClient.Core().(*core.Client) {
			t.Error("Expected client to match provided client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := mdns.NewService(nil)

		// Verify service creation (should handle nil gracefully)
		if service.Client() != nil {
			t.Error("Expected nil client, got non-nil")
		}
	})
}

// TestMdnsServiceUnit_GetOperations_MockSuccess tests Get operations using mock server with real WNC data structure.
func TestMdnsServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC MDNS data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data": `{
			"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data": {
				"mdns-global-stats": {
					"stats-global": {
						"pak-sent": "0",
						"pak-sent-v4": "0",
						"pak-sent-advt-v4": "0",
						"pak-sent-query-v4": "0",
						"pak-received": "0",
						"pak-received-advt": "0",
						"pak-received-query": "0",
						"pak-dropped": "0",
						"ptr-query": "0",
						"srv-query": "0",
						"a-query": "0",
						"aaaa-query": "0",
						"txt-query": "0",
						"any-query": "0",
						"other-query": "0"
					},
					"last-clear-time": "2025-09-06T04:13:50+00:00"
				},
				"mdns-wlan-stats": [
					{
						"wlan-id": 0,
						"stats-wlan": {
							"pak-sent": "0",
							"pak-received": "0",
							"pak-dropped": "0"
						},
						"last-clear-time": "2025-09-06T04:13:50+00:00"
					},
					{
						"wlan-id": 1,
						"stats-wlan": {
							"pak-sent": "0",
							"pak-received": "0",
							"pak-dropped": "0"
						},
						"last-clear-time": "2025-09-06T04:14:28+00:00"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data/mdns-global-stats": `{
			"Cisco-IOS-XE-wireless-mdns-oper:mdns-global-stats": {
				"stats-global": {
					"pak-sent": "0",
					"pak-sent-v4": "0",
					"pak-sent-advt-v4": "0",
					"pak-sent-query-v4": "0",
					"pak-sent-v6": "0",
					"pak-sent-advt-v6": "0",
					"pak-sent-query-v6": "0",
					"pak-sent-mcast": "0",
					"pak-sent-mcast-v4": "0",
					"pak-sent-mcast-v6": "0",
					"pak-received": "0",
					"pak-received-advt": "0",
					"pak-received-query": "0",
					"pak-received-v4": "0",
					"pak-received-advt-v4": "0",
					"pak-received-query-v4": "0",
					"pak-received-v6": "0",
					"pak-received-advt-v6": "0",
					"pak-received-query-v6": "0",
					"pak-dropped": "0",
					"ptr-query": "0",
					"srv-query": "0",
					"a-query": "0",
					"aaaa-query": "0",
					"txt-query": "0",
					"any-query": "0",
					"other-query": "0"
				},
				"last-clear-time": "2025-09-06T04:13:50+00:00"
			}
		}`,
		"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data/mdns-wlan-stats": `{
			"Cisco-IOS-XE-wireless-mdns-oper:mdns-wlan-stats": [
				{
					"wlan-id": 0,
					"stats-wlan": {
						"pak-sent": "0",
						"pak-sent-v4": "0",
						"pak-sent-advt-v4": "0",
						"pak-sent-query-v4": "0",
						"pak-sent-v6": "0",
						"pak-sent-advt-v6": "0",
						"pak-sent-query-v6": "0",
						"pak-sent-mcast": "0",
						"pak-sent-mcast-v4": "0",
						"pak-sent-mcast-v6": "0",
						"pak-received": "0",
						"pak-received-advt": "0",
						"pak-received-query": "0",
						"pak-received-v4": "0",
						"pak-received-advt-v4": "0",
						"pak-received-query-v4": "0",
						"pak-received-v6": "0",
						"pak-received-advt-v6": "0",
						"pak-received-query-v6": "0",
						"pak-dropped": "0",
						"ptr-query": "0",
						"srv-query": "0",
						"a-query": "0",
						"aaaa-query": "0",
						"txt-query": "0",
						"any-query": "0",
						"other-query": "0"
					},
					"last-clear-time": "2025-09-06T04:13:50+00:00"
				},
				{
					"wlan-id": 1,
					"stats-wlan": {
						"pak-sent": "0",
						"pak-sent-v4": "0",
						"pak-sent-advt-v4": "0",
						"pak-sent-query-v4": "0",
						"pak-sent-v6": "0",
						"pak-sent-advt-v6": "0",
						"pak-sent-query-v6": "0",
						"pak-sent-mcast": "0",
						"pak-sent-mcast-v4": "0",
						"pak-sent-mcast-v6": "0",
						"pak-received": "0",
						"pak-received-advt": "0",
						"pak-received-query": "0",
						"pak-received-v4": "0",
						"pak-received-advt-v4": "0",
						"pak-received-query-v4": "0",
						"pak-received-v6": "0",
						"pak-received-advt-v6": "0",
						"pak-received-query-v6": "0",
						"pak-dropped": "0",
						"ptr-query": "0",
						"srv-query": "0",
						"a-query": "0",
						"aaaa-query": "0",
						"txt-query": "0",
						"any-query": "0",
						"other-query": "0"
					},
					"last-clear-time": "2025-09-06T04:14:28+00:00"
				}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := mdns.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
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

	t.Run("ListWLANStats", func(t *testing.T) {
		result, err := service.ListWLANStats(ctx)
		if err != nil {
			t.Errorf("ListWLANStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListWLANStats returned nil result")
		}
	})
}

// TestMdnsServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestMdnsServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := mdns.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational_404Error", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
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

	t.Run("ListWLANStats_404Error", func(t *testing.T) {
		result, err := service.ListWLANStats(ctx)
		if err == nil {
			t.Error("Expected error for ListWLANStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestMdnsServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestMdnsServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		service := mdns.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetGlobalStats_NilClient", func(t *testing.T) {
		service := mdns.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetGlobalStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListWLANStats_NilClient", func(t *testing.T) {
		service := mdns.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListWLANStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
