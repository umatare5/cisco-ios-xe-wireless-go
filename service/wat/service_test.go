package wat_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wat"
)

// TestWatServiceUnit_Constructor_Success tests service constructor functionality.
func TestWatServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := wat.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := wat.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestWatServiceUnit_GetOperations_ErrorExpected tests Get operations expecting errors for IOS-XE 17.12.x.
func TestWatServiceUnit_GetOperations_ErrorExpected(t *testing.T) {
	// WAT is experimental feature for IOS-XE 17.18.1+, expect "uri keypath not found" errors on IOS-XE 17.12.x
	// Based on live WNC response: {"ietf-restconf:errors":{"error":[{"error-type":"application","error-tag":"invalid-value","error-message":"uri keypath not found"}]}}
	mockServer := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data", testutil.ResponseConfig{
			StatusCode: 400,
			Body:       `{"ietf-restconf:errors":{"error":[{"error-type":"application","error-tag":"invalid-value","error-message":"uri keypath not found"}]}}`,
		}),
	)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := wat.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetConfig - expect error for unsupported feature on IOS-XE 17.12.x
	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error for unsupported WAT feature on IOS-XE 17.12.x, got nil")
	}
}

// TestWatServiceUnit_GetOperations_MockSuccess tests Get operations using mock server with expected IOS-XE 17.18.1+ data.
func TestWatServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with WAT endpoints using expected IOS-XE 17.18.1+ data structure
	// Note: This is hypothetical data structure based on YANG model since WAT is not available on current IOS-XE 17.12.x
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data": `{
			"Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data": {
				"wat-config": {
					"wat-enable": true,
					"te-conn-str": "test-connection-string",
					"te-download-url": "https://downloads.thousandeyes.com",
					"te-agent-version": "1.0.0",
					"te-cloud-endpoint": "api.thousandeyes.com",
					"te-poll-interval": 300,
					"te-timeout": 30,
					"te-retry-attempts": 3,
					"te-log-level": "info",
					"te-data-collection": true,
					"te-analytics-enabled": true
				}
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := wat.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetConfig operation with expected IOS-XE 17.18.1+ response structure
	result, err := service.GetConfig(ctx)
	if err != nil {
		t.Errorf("GetConfig failed: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result from GetConfig")
	}
}

// TestWatServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestWatServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for WAT endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := wat.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetConfig properly handles 404 errors
	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}
