package controller_test

import (
	"testing"
	"time"

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

// bootTimeEndpoint is the endpoint GetBootTime reads, without the RESTCONF data prefix the mock
// server strips.
const bootTimeEndpoint = "Cisco-IOS-XE-device-hardware-oper:device-hardware-data/" +
	"device-hardware/device-system-data/boot-time"

// wireBootTime is the shape a controller sends: RFC 3339 to the second, with UTC written as an
// explicit "+00:00" offset rather than "Z", and with no fractional part.
const wireBootTime = "2026-01-02T03:04:05+00:00"

// TestControllerServiceUnit_GetBootTime_MockSuccess tests that a boot instant decodes and that an
// answer holding nothing decodes to nil rather than to the year 1.
func TestControllerServiceUnit_GetBootTime_MockSuccess(t *testing.T) {
	t.Run("InstantIsDecoded", func(t *testing.T) {
		responses := map[string]string{
			bootTimeEndpoint: `{"Cisco-IOS-XE-device-hardware-oper:boot-time": "` + wireBootTime + `"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := controller.NewService(testClient.Core().(*core.Client))

		result, err := service.GetBootTime(testutil.TestContext(t))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}
		if result == nil || result.BootTime == nil {
			t.Fatal("Expected a boot instant, got nil")
		}

		want, err := time.Parse(time.RFC3339, wireBootTime)
		if err != nil {
			t.Fatalf("Test constant is not RFC 3339: %v", err)
		}
		if !result.BootTime.Equal(want) {
			t.Errorf("Boot instant mismatch: expected %s, got %s", want, result.BootTime)
		}
	})

	// A read answered with no body is a successful read of a node holding nothing, so the field has
	// to be able to say so. Were it a value, this would be the year 1 behind a nil error.
	t.Run("EmptyAnswerIsNotAnInstant", func(t *testing.T) {
		responses := map[string]string{
			bootTimeEndpoint: "",
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := controller.NewService(testClient.Core().(*core.Client))

		result, err := service.GetBootTime(testutil.TestContext(t))
		if err != nil {
			t.Fatalf("Expected no error for an empty answer, got: %v", err)
		}
		if result == nil {
			t.Fatal("Expected a result for an empty answer, got nil")
		}
		if result.BootTime != nil {
			t.Errorf("Expected nil boot instant for an empty answer, got %s", result.BootTime)
		}
	})
}

// TestControllerServiceUnit_GetBootTime_ErrorHandling tests that a body keyed for another node is
// an error rather than a zero instant.
func TestControllerServiceUnit_GetBootTime_ErrorHandling(t *testing.T) {
	responses := map[string]string{
		bootTimeEndpoint: `{"Cisco-IOS-XE-device-hardware-oper:up-time": "` + wireBootTime + `"}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := controller.NewService(testClient.Core().(*core.Client))

	if _, err := service.GetBootTime(testutil.TestContext(t)); err == nil {
		t.Error("Expected an error for a body keyed for another node, got nil")
	}
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

// TestControllerServiceUnit_SaveConfig_MockAnswers tests the three answers the RPC can give.
func TestControllerServiceUnit_SaveConfig_MockAnswers(t *testing.T) {
	// The fixture is the body the controller sends, pretty-printed at two spaces: 78 bytes with its
	// newline, where a compact one would be 63.
	t.Run("ResultIsDecoded", func(t *testing.T) {
		responses := map[string]string{
			"cisco-ia:save-config": `{
  "cisco-ia:output": {
    "result": "Save running-config successful"
  }
}
`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		service := controller.NewService(testutil.NewTestClient(mockServer).Core().(*core.Client))

		out, err := service.SaveConfig(testutil.TestContext(t))
		if err != nil {
			t.Fatalf("Expected no error for mock save, got: %v", err)
		}
		if out == nil || out.Output == nil {
			t.Fatal("Expected the output container to decode, got nil")
		}
		if out.Output.Result != "Save running-config successful" {
			t.Errorf("Result = %q, want the controller's account", out.Output.Result)
		}
	})

	// An answer with no body is a success carrying no account, not an error: decode returns the
	// zero envelope, so Output stays nil and the caller tests the leaf rather than the error.
	t.Run("BodilessAnswerLeavesOutputNil", func(t *testing.T) {
		responses := map[string]string{"cisco-ia:save-config": ""}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		service := controller.NewService(testutil.NewTestClient(mockServer).Core().(*core.Client))

		out, err := service.SaveConfig(testutil.TestContext(t))
		if err != nil {
			t.Fatalf("Expected no error for a bodiless answer, got: %v", err)
		}
		if out == nil {
			t.Fatal("Expected a non-nil envelope for a bodiless answer, got nil")
		}
		if out.Output != nil {
			t.Errorf("Output = %+v, want nil for a bodiless answer", out.Output)
		}
	})

	t.Run("ControllerRefusalIsAnError", func(t *testing.T) {
		errorPaths := []string{"cisco-ia:save-config"}
		mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 500))
		defer mockServer.Close()

		service := controller.NewService(testutil.NewTestClient(mockServer).Core().(*core.Client))

		if _, err := service.SaveConfig(testutil.TestContext(t)); err == nil {
			t.Error("Expected error for 500 response, got nil")
		}
	})
}
