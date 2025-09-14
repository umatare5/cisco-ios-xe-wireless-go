package spaces_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/spaces"
)

// TestSpacesServiceUnit_Constructor_Success tests service constructor functionality.
func TestSpacesServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := spaces.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := spaces.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestSpacesServiceUnit_GetOperations_MockSuccess tests Get operations using mock server with real IOS-XE 17.18.1+ data.
func TestSpacesServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Spaces endpoints using real IOS-XE 17.18.1+ data structure
	// Note: This is real data from IOS-XE 17.18.1 Live WNC, preserved from original test implementation
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": `{
			"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": {
				"spaces-connection-detail": {
					"spaces-health-url": "https://health.ciscoservices.com/ns-api/v2/org/health",
					"con-state": "connected",
					"conn-estb-time": "2025-01-11T10:30:45Z",
					"stats": {
						"total-con-attempts": 25,
						"con-attempts-success": 23,
						"con-attempts-failure": 2,
						"total-msg-sent": 12847,
						"total-msg-rcvd": 12844,
						"last-heartbeat-time": "2025-01-11T10:30:45Z"
					},
					"tenant": {
						"tenant-id": "demo-tenant",
						"organization-name": "Demo Organization",
						"registration-status": "active"
					}
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail": `{
			"Cisco-IOS-XE-wireless-cisco-spaces-oper:spaces-connection-detail": {
				"spaces-health-url": "https://health.ciscoservices.com/ns-api/v2/org/health",
				"con-state": "connected",
				"conn-estb-time": "2025-01-11T10:30:45Z",
				"stats": {
					"total-con-attempts": 25,
					"con-attempts-success": 23,
					"con-attempts-failure": 2,
					"total-msg-sent": 12847,
					"total-msg-rcvd": 12844,
					"last-heartbeat-time": "2025-01-11T10:30:45Z"
				},
				"tenant": {
					"tenant-id": "demo-tenant",
					"organization-name": "Demo Organization",
					"registration-status": "active"
				}
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := spaces.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation with real IOS-XE 17.18.1+ response structure
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result from GetOperational")
	}

	// Test GetConnectionDetails operation with real IOS-XE 17.18.1+ response structure
	connectionDetail, err := service.GetConnectionDetails(ctx)
	if err != nil {
		t.Errorf("GetConnectionDetails failed: %v", err)
	}
	if connectionDetail == nil {
		t.Error("Expected non-nil result from GetConnectionDetails")
	}
}

// TestSpacesServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestSpacesServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Spaces endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data",
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := spaces.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Test that GetConnectionDetails properly handles 404 errors
	_, err = service.GetConnectionDetails(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}
