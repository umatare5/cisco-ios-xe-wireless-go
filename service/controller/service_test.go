package controller_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/controller"
)

// TestControllerServiceUnit_Constructor_Success tests service constructor functionality.
func TestControllerServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := controller.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := controller.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestControllerServiceUnit_ReloadOperations_MockSuccess tests Reload operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestControllerServiceUnit_ReloadOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Controller RPC endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-rpc:reload": `{
			"Cisco-IOS-XE-rpc:output": {
				"result": "success",
				"message": "Controller reload initiated successfully"
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := controller.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ReloadWithReason operation
	err := service.ReloadWithReason(ctx, "Test reload for mock testing")
	if err != nil {
		t.Errorf("Expected no error for mock reload, got: %v", err)
	}

	// Test Reload operation with force flag
	err = service.Reload(ctx, "Test reload with force", true)
	if err != nil {
		t.Errorf("Expected no error for mock reload with force, got: %v", err)
	}

	// Test Reload operation without force flag
	err = service.Reload(ctx, "Test reload without force", false)
	if err != nil {
		t.Errorf("Expected no error for mock reload without force, got: %v", err)
	}
}

// TestControllerServiceUnit_ReloadOperations_ErrorHandling tests error scenarios using mock server.
func TestControllerServiceUnit_ReloadOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 500 for Controller RPC endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-rpc:reload",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 500))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := controller.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that Reload properly handles 500 errors
	err := service.ReloadWithReason(ctx, "Test error handling")
	if err == nil {
		t.Error("Expected error for 500 response, got nil")
	}
}

// TestControllerServiceUnit_ReloadOperations_ValidationErrors tests input validation.
func TestControllerServiceUnit_ReloadOperations_ValidationErrors(t *testing.T) {
	// Create mock RESTCONF server
	responses := map[string]string{
		"Cisco-IOS-XE-rpc:reload": `{
			"Cisco-IOS-XE-rpc:output": {
				"result": "success",
				"message": "Controller reload initiated successfully"
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := controller.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test empty reason validation
	err := service.ReloadWithReason(ctx, "")
	if err == nil {
		t.Error("Expected error for empty reason, got nil")
	}

	// Test empty reason validation for Reload with force
	err = service.Reload(ctx, "", true)
	if err == nil {
		t.Error("Expected error for empty reason in Reload, got nil")
	}

	// Test whitespace-only reason validation
	err = service.Reload(ctx, "   ", false)
	if err == nil {
		t.Error("Expected error for whitespace-only reason in Reload, got nil")
	}

	// Test whitespace-only reason validation
	err = service.ReloadWithReason(ctx, "   ")
	if err == nil {
		t.Error("Expected error for whitespace-only reason, got nil")
	}
}
