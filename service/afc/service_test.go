package afc_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/afc"
)

// TestAfcServiceUnit_Constructor_Success tests service constructor functionality.
func TestAfcServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := afc.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := afc.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestAfcServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestAfcServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with AFC endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": `{
			"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": {
				"ewlc-afc-info": {
					"afc-enable": true,
					"afc-server-url": "https://afc.example.com"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp": [{
				"ap-mac": "aa:bb:cc:dd:ee:ff",
				"response-status": "success"
			}]
		}`,
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req": `{
			"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-req": [{
				"ap-mac": "aa:bb:cc:dd:ee:ff",
				"request-status": "pending"
			}]
		}`,
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": `{
			"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": {
				"afc-cloud-enable": true
			}
		}`,
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats": `{
			"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-stats": {
				"requests-sent": 100,
				"responses-received": 95
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := afc.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test all AFC Get operations
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetOperational, got nil")
		}
	})

	t.Run("ListAPResponses", func(t *testing.T) {
		result, err := service.ListAPResponses(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPResponses, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPResponses, got nil")
		}
	})

	t.Run("ListAPRequests", func(t *testing.T) {
		result, err := service.ListAPRequests(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPRequests, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPRequests, got nil")
		}
	})

	t.Run("GetCloudInfo", func(t *testing.T) {
		result, err := service.GetCloudInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetCloudInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCloudInfo, got nil")
		}
	})

	t.Run("GetCloudStats", func(t *testing.T) {
		result, err := service.GetCloudStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetCloudStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCloudStats, got nil")
		}
	})
}

// TestAfcServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestAfcServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for AFC endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
		"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-req",
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data",
		"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := afc.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
		}
	})

	t.Run("ListAPResponses", func(t *testing.T) {
		_, err := service.ListAPResponses(ctx)
		if err == nil {
			t.Error("Expected error for ListAPResponses, got nil")
		}
	})

	t.Run("ListAPRequests", func(t *testing.T) {
		_, err := service.ListAPRequests(ctx)
		if err == nil {
			t.Error("Expected error for ListAPRequests, got nil")
		}
	})

	t.Run("GetCloudInfo", func(t *testing.T) {
		_, err := service.GetCloudInfo(ctx)
		if err == nil {
			t.Error("Expected error for GetCloudInfo, got nil")
		}
	})

	t.Run("GetCloudStats", func(t *testing.T) {
		_, err := service.GetCloudStats(ctx)
		if err == nil {
			t.Error("Expected error for GetCloudStats, got nil")
		}
	})
}
