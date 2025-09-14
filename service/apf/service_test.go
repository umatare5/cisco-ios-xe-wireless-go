package apf_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/apf"
)

// TestUapfServiceUnit_Constructor_Success tests service constructor functionality.
func TestUapfServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := apf.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := apf.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestUapfServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestUapfServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Uapf endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data": `{
			"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data": {
				"config": {
					"enable": true
				}
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := apf.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetConfig operation
	result, err := service.GetConfig(ctx)
	if err != nil {
		t.Errorf("Expected no error for GetConfig, got: %v", err)
	}
	if result == nil {
		t.Error("Expected result for GetConfig, got nil")
	}
}

// TestUapfServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestUapfServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Uapf endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := apf.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetConfig properly handles 404 errors
	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}
}
