package urwb_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/urwb"
)

// TestUrwbServiceUnit_Constructor_Success tests service constructor functionality.
func TestUrwbServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := urwb.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := urwb.NewService(nil)

		// Test that service can be created with nil client
		if service.Client() != nil {
			t.Error("Expected service to have nil client")
		}
	})
}

// TestUrwbServiceUnit_GetConfig_MockSuccess tests GetConfig method with mock success.
func TestUrwbServiceUnit_GetConfig_MockSuccess(t *testing.T) {
	// Mock server response that returns empty data (simulating successful response with no content)
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": `{}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := urwb.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetConfig(ctx)
	// Verify no error and result is not nil
	if err != nil {
		t.Errorf("GetConfig returned unexpected error: %v", err)
	}

	if result == nil {
		t.Error("GetConfig returned nil result")
	}
}

// TestUrwbServiceUnit_GetConfig_ErrorHandling tests GetConfig error handling.
func TestUrwbServiceUnit_GetConfig_ErrorHandling(t *testing.T) {
	// Mock server that returns an error for URWB config requests
	mockServer := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data", testutil.ResponseConfig{
			StatusCode: 400,
			Body:       "uri keypath not found",
		}),
	)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := urwb.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetConfig(ctx)

	// Verify error is returned and result is nil
	if err == nil {
		t.Error("GetConfig should have returned an error")
	}

	if result != nil {
		t.Error("GetConfig should have returned nil result on error")
	}
}

// TestUrwbServiceUnit_GetURWBNetOperational_MockSuccess tests GetURWBNetOperational method with mock success.
func TestUrwbServiceUnit_GetURWBNetOperational_MockSuccess(t *testing.T) {
	// Mock server response that returns empty operational data
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": `{}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := urwb.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetURWBNetOperational(ctx)
	// Verify no error and result is not nil
	if err != nil {
		t.Errorf("GetURWBNetOperational returned unexpected error: %v", err)
	}

	if result == nil {
		t.Error("GetURWBNetOperational returned nil result")
	}
}
