package awips_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/awips"
)

// TestAwipsServiceUnit_Constructor_Success tests service constructor functionality.
func TestAwipsServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := awips.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := awips.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestAwipsServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestAwipsServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with AWIPS endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data": {
				"awips-per-ap-info": [
					{
						"ap-mac": "00:11:22:33:44:55",
						"awips-status": "enabled",
						"alarm-count": "0",
						"forensic-capture-status": "disabled"
					}
				],
				"awips-dwld-status": {
					"last-success-timestamp": "2023-01-01T00:00:00Z",
					"last-failed-timestamp": "",
					"num-of-failure-attempts": 0,
					"last-failure-reason": 0,
					"wlc-version": "17.12.5",
					"max-file-ver": 1,
					"latest-file-version": 1
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-ap-info": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info": [
				{
					"ap-mac": "00:11:22:33:44:55",
					"awips-status": "enabled",
					"alarm-count": "0",
					"forensic-capture-status": "disabled"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status": {
				"last-success-timestamp": "2023-01-01T00:00:00Z",
				"last-failed-timestamp": "",
				"num-of-failure-attempts": 0,
				"last-failure-reason": 0,
				"wlc-version": "17.12.5",
				"max-file-ver": 1,
				"latest-file-version": 1
			}
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-ap-dwld-status": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-ap-dwld-status": [
				{
					"ap-mac": "00:11:22:33:44:55",
					"status": "success"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-sign-stats": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-per-sign-stats": [
				{
					"signature-id": 1,
					"match-count": 5
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-glob-stats": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-glob-stats": {
				"total-detections": 10,
				"active-threats": 2
			}
		}`,
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status-wncd": `{
			"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status-wncd": {
				"internal-status": "active"
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := awips.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetOperational, got nil")
		}
	})

	// Test ListAWIPSPerApInfo operation
	t.Run("ListAWIPSPerApInfo", func(t *testing.T) {
		result, err := service.ListAWIPSPerApInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSPerApInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSPerApInfo, got nil")
		}
	})

	// Test ListAWIPSDwldStatus operation
	t.Run("ListAWIPSDwldStatus", func(t *testing.T) {
		result, err := service.ListAWIPSDwldStatus(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSDwldStatus, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSDwldStatus, got nil")
		}
	})

	// Test ListAWIPSApDwldStatus operation
	t.Run("ListAWIPSApDwldStatus", func(t *testing.T) {
		result, err := service.ListAWIPSApDwldStatus(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSApDwldStatus, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSApDwldStatus, got nil")
		}
	})

	// Test ListAWIPSPerSignStats operation
	t.Run("ListAWIPSPerSignStats", func(t *testing.T) {
		result, err := service.ListAWIPSPerSignStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSPerSignStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSPerSignStats, got nil")
		}
	})

	// Test ListAWIPSGlobStats operation
	t.Run("ListAWIPSGlobStats", func(t *testing.T) {
		result, err := service.ListAWIPSGlobStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSGlobStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSGlobStats, got nil")
		}
	})

	// Test ListAWIPSDwldStatusWncd operation
	t.Run("ListAWIPSDwldStatusWncd", func(t *testing.T) {
		result, err := service.ListAWIPSDwldStatusWncd(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAWIPSDwldStatusWncd, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAWIPSDwldStatusWncd, got nil")
		}
	})
}

// TestAwipsServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestAwipsServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for AWIPS endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-ap-info",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-ap-dwld-status",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-sign-stats",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-glob-stats",
		"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status-wncd",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := awips.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test representative functions from each category
	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListAWIPSPerApInfo_404Error", func(t *testing.T) {
		_, err := service.ListAWIPSPerApInfo(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListAWIPSGlobStats_404Error", func(t *testing.T) {
		_, err := service.ListAWIPSGlobStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}

// TestAwipsServiceUnit_ErrorHandling_NilClient tests operations with nil client.
func TestAwipsServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := awips.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})

	t.Run("ListAWIPSPerApInfo_NilClient", func(t *testing.T) {
		result, err := service.ListAWIPSPerApInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})

	t.Run("ListAWIPSGlobStats_NilClient", func(t *testing.T) {
		result, err := service.ListAWIPSGlobStats(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})
}
