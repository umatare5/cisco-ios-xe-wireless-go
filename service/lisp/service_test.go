package lisp_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/lisp"
)

// TestLispServiceUnit_Constructor_Success tests service constructor functionality.
func TestLispServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := lisp.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := lisp.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestLispServiceUnit_GetOperations_MockSuccess tests Get operations using mock server with real WNC data structure.
func TestLispServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC LISP data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data": `{
			"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data": {
				"lisp-agent-memory-stats": {
					"malloc-psk-buf": "0",
					"free-psk-buf": "0",
					"malloc-map-reg-msg": "0",
					"free-map-reg-msg": "0",
					"malloc-map-req-msg": "0",
					"free-map-req-msg": "0",
					"malloc-lisp-ha-node": "0",
					"free-lisp-ha-node": "0",
					"malloc-map-server-ctxt": "0",
					"free-map-server-ctxt": "0"
				},
				"lisp-wlc-capabilities": {
					"fabric-capable": true
				},
				"lisp-ap-capabilities": [
					{
						"ap-type": 35,
						"fabric-capable": true
					},
					{
						"ap-type": 36,
						"fabric-capable": true
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-agent-memory-stats": `{
			"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-memory-stats": {
				"malloc-psk-buf": "0",
				"free-psk-buf": "0",
				"malloc-map-reg-msg": "0",
				"free-map-reg-msg": "0",
				"malloc-map-req-msg": "0",
				"free-map-req-msg": "0",
				"malloc-lisp-ha-node": "0",
				"free-lisp-ha-node": "0",
				"malloc-map-server-ctxt": "0",
				"free-map-server-ctxt": "0"
			}
		}`,
		"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-wlc-capabilities": `{
			"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-wlc-capabilities": {
				"fabric-capable": true
			}
		}`,
		"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-agent-oper-data/lisp-ap-capabilities": `{
			"Cisco-IOS-XE-wireless-lisp-agent-oper:lisp-ap-capabilities": [
				{
					"ap-type": 35,
					"fabric-capable": true
				},
				{
					"ap-type": 36,
					"fabric-capable": true
				}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := lisp.NewService(testClient.Core().(*core.Client))
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

	t.Run("GetMemoryStats", func(t *testing.T) {
		result, err := service.GetMemoryStats(ctx)
		if err != nil {
			t.Errorf("GetMemoryStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMemoryStats returned nil result")
		}
	})

	t.Run("GetCapabilities", func(t *testing.T) {
		result, err := service.GetCapabilities(ctx)
		if err != nil {
			t.Errorf("GetCapabilities returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetCapabilities returned nil result")
		}
	})

	t.Run("ListAPCapabilities", func(t *testing.T) {
		result, err := service.ListAPCapabilities(ctx)
		if err != nil {
			t.Errorf("ListAPCapabilities returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListAPCapabilities returned nil result")
		}
	})
}

// TestLispServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestLispServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := lisp.NewService(testClient.Core().(*core.Client))
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

	t.Run("GetMemoryStats_404Error", func(t *testing.T) {
		result, err := service.GetMemoryStats(ctx)
		if err == nil {
			t.Error("Expected error for GetMemoryStats, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetCapabilities_404Error", func(t *testing.T) {
		result, err := service.GetCapabilities(ctx)
		if err == nil {
			t.Error("Expected error for GetCapabilities, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListAPCapabilities_404Error", func(t *testing.T) {
		result, err := service.ListAPCapabilities(ctx)
		if err == nil {
			t.Error("Expected error for ListAPCapabilities, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestLispServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestLispServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		service := lisp.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetMemoryStats_NilClient", func(t *testing.T) {
		service := lisp.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetMemoryStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetCapabilities_NilClient", func(t *testing.T) {
		service := lisp.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetCapabilities(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
