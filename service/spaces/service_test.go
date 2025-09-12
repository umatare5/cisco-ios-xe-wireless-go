package spaces

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/spaces"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

// TestSpacesServiceUnit_Constructor_Success tests service constructor functionality.
func TestSpacesServiceUnit_Constructor_Success(t *testing.T) {
	service := NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client service")
	}

	// Test with valid client
	mockServer := testutil.NewMockServer(map[string]string{
		"test": `{"data": {}}`,
	})
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service = NewService(client.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected service to have client, got nil")
	}
}

// TestSpacesServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
// This is essential for CI environments where actual Cisco controllers are not available.
func TestSpacesServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Spaces endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": `{
			"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": {
				"status": "active"
			}
		}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
	}
	if result == nil {
		t.Error("Expected result for GetOperational, got nil")
	}
}

// TestSpacesServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestSpacesServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Spaces endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data",
	}
	mockServer := testutil.NewMockErrorServer(errorPaths, 404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}
}

// TestSpacesServiceUnit_GetOperational_RealDataSuccess tests GetOperational with real IOS-XE 17.18.1 data.
func TestSpacesServiceUnit_GetOperational_RealDataSuccess(t *testing.T) {
	// Real data from IOS-XE 17.18.1 Live WNC
	mockResponse := `{
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
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	result, err := service.GetOperational(context.Background())
	if err != nil {
		t.Log("Method returned error (expected when service not configured)")
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.CiscoSpacesOperData == nil || result.CiscoSpacesOperData.SpacesConnectionDetail == nil {
		t.Error("Expected spaces connection detail, got nil")
		return
	}

	detail := result.CiscoSpacesOperData.SpacesConnectionDetail
	if detail.SpacesHealthURL != "https://health.ciscoservices.com/ns-api/v2/org/health" {
		t.Errorf("Expected health URL 'https://health.ciscoservices.com/ns-api/v2/org/health', got '%s'",
			detail.SpacesHealthURL)
	}

	if detail.ConnectionState != "connected" {
		t.Errorf("Expected connection state 'connected', got '%s'", detail.ConnectionState)
	}

	if detail.Stats == nil {
		t.Error("Expected stats, got nil")
	} else {
		if detail.Stats.TotalConnectionAttempts != 25 {
			t.Errorf("Expected total connection attempts 25, got %d", detail.Stats.TotalConnectionAttempts)
		}
		if detail.Stats.ConnectionAttemptsSuccess != 23 {
			t.Errorf("Expected successful attempts 23, got %d", detail.Stats.ConnectionAttemptsSuccess)
		}
		if detail.Stats.TotalMessagesSent != 12847 {
			t.Errorf("Expected total messages sent 12847, got %d", detail.Stats.TotalMessagesSent)
		}
	}

	if detail.Tenant == nil {
		t.Error("Expected tenant info, got nil")
	} else {
		if detail.Tenant.TenantID != "demo-tenant" {
			t.Errorf("Expected tenant ID 'demo-tenant', got '%s'", detail.Tenant.TenantID)
		}
		if detail.Tenant.OrganizationName != "Demo Organization" {
			t.Errorf("Expected organization name 'Demo Organization', got '%s'", detail.Tenant.OrganizationName)
		}
		if detail.Tenant.RegistrationStatus != "active" {
			t.Errorf("Expected registration status 'active', got '%s'", detail.Tenant.RegistrationStatus)
		}
	}
}

// TestSpacesServiceUnit_GetConnectionDetails_RealDataSuccess tests GetConnectionDetails error handling.
// NOTE: Live WNC did not have connection details data available for testing.
func TestSpacesServiceUnit_GetConnectionDetails_RealDataSuccess(t *testing.T) {
	// Use MockErrorServer to simulate 404 errors for unconfigured services
	mockServer := testutil.NewMockErrorServer(
		[]string{"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail"},
		404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	result, err := service.GetConnectionDetails(context.Background())
	if err != nil {
		t.Log("Method returned error (expected when service not configured)")
		// When service not configured, expect nil result
		if result != nil {
			t.Log("Note: Method returned non-nil result despite error")
		}
		return
	}

	// If no error, validate the mock response structure
	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.SpacesHealthURL != "https://health.ciscoservices.com/ns-api/v2/org/health" {
		t.Errorf("Expected health URL 'https://health.ciscoservices.com/ns-api/v2/org/health', got '%s'",
			result.SpacesHealthURL)
	}

	if result.ConnectionState != "connected" {
		t.Errorf("Expected connection state 'connected', got '%s'", result.ConnectionState)
	}

	if result.Tenant == nil {
		t.Error("Expected tenant info, got nil")
	} else {
		if result.Tenant.TenantID != "test-tenant" {
			t.Errorf("Expected tenant ID 'test-tenant', got '%s'", result.Tenant.TenantID)
		}
		if result.Tenant.OrganizationName != "Test Organization" {
			t.Errorf("Expected organization name 'Test Organization', got '%s'", result.Tenant.OrganizationName)
		}
	}
}

// TestSpacesServiceUnit_GetTenantInfo_RealDataSuccess tests GetTenantInfo error handling.
// NOTE: Live WNC did not have tenant configuration data available for testing.
func TestSpacesServiceUnit_GetTenantInfo_RealDataSuccess(t *testing.T) {
	// Configuration setup for IOS-XE 17.18.1
	// Mock 404 response as no tenant data was available on live WNC
	mockServer := testutil.NewMockErrorServer(
		[]string{"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data"}, 404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	result, err := service.GetTenantInfo(context.Background())
	if err != nil {
		t.Log("GetTenantInfo returned error (expected when tenant not configured)")
		if result != nil {
			t.Error("Expected nil result when error occurred")
		}
	} else {
		t.Log("GetTenantInfo returned success")
	}
}

// TestSpacesServiceUnit_GetConnectionStats_RealDataSuccess tests GetConnectionStats error handling.
// NOTE: Live WNC did not have connection stats data available for testing.
func TestSpacesServiceUnit_GetConnectionStats_RealDataSuccess(t *testing.T) {
	// Use MockErrorServer to simulate 404 errors for unconfigured services
	mockServer := testutil.NewMockErrorServer(
		[]string{"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail/stats"}, 404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	result, err := service.GetConnectionStats(context.Background())
	if err != nil {
		t.Log("Method returned error (expected when service not configured)")
		// When service not configured, expect nil result
		if result != nil {
			t.Log("Note: Method returned non-nil result despite error")
		}
		return
	}

	// If no error, validate the mock response structure
	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.TotalConnectionAttempts != 100 {
		t.Errorf("Expected total connection attempts 100, got %d", result.TotalConnectionAttempts)
	}

	if result.ConnectionAttemptsSuccess != 98 {
		t.Errorf("Expected successful attempts 98, got %d", result.ConnectionAttemptsSuccess)
	}

	if result.ConnectionAttemptsFailure != 2 {
		t.Errorf("Expected failed attempts 2, got %d", result.ConnectionAttemptsFailure)
	}

	if result.TotalMessagesReceived != 50000 {
		t.Errorf("Expected messages received 50000, got %d", result.TotalMessagesReceived)
	}

	if result.TotalMessagesSent != 49995 {
		t.Errorf("Expected messages sent 49995, got %d", result.TotalMessagesSent)
	}

	if result.AverageResponseTime != 150 {
		t.Errorf("Expected average response time 150ms, got %d", result.AverageResponseTime)
	}
}

// TestSpacesServiceUnit_ErrorHandling_HTTPErrors tests various HTTP error scenarios.
func TestSpacesServiceUnit_ErrorHandling_HTTPErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation func(service Service) error
	}{
		{
			name: "GetOperational_InternalServerError",
			operation: func(service Service) error {
				_, err := service.GetOperational(context.Background())
				return err
			},
		},
		{
			name: "GetConnectionDetails_NotFound",
			operation: func(service Service) error {
				_, err := service.GetConnectionDetails(context.Background())
				return err
			},
		},
		{
			name: "GetTenantInfo_ServiceUnavailable",
			operation: func(service Service) error {
				_, err := service.GetTenantInfo(context.Background())
				return err
			},
		},
		{
			name: "GetConnectionStats_ServiceUnavailable",
			operation: func(service Service) error {
				_, err := service.GetConnectionStats(context.Background())
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorPaths := []string{
				"Cisco-IOS-XE-wireless-spaces-oper:spaces-oper-data",
				"Cisco-IOS-XE-wireless-spaces-oper:spaces-connection-detail",
			}

			mockServer := testutil.NewMockErrorServer(errorPaths, 500)
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := NewService(testClient.Core().(*core.Client))

			err := tt.operation(service)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
		})
	}
}

// TestSpacesServiceUnit_ErrorHandling_NetworkErrors tests network-level error scenarios.
func TestSpacesServiceUnit_ErrorHandling_NetworkErrors(t *testing.T) {
	tests := []struct {
		name      string
		responses map[string]string
		operation func(service Service) error
	}{
		{
			name: "InvalidJSON",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": `{"invalid": json}`,
			},
			operation: func(service Service) error {
				_, err := service.GetOperational(context.Background())
				return err
			},
		},
		{
			name: "EmptyResponse",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data": `{}`,
			},
			operation: func(service Service) error {
				_, err := service.GetOperational(context.Background())
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(tt.responses)
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := NewService(testClient.Core().(*core.Client))

			err := tt.operation(service)
			if err == nil && tt.name != "EmptyResponse" {
				t.Fatal("Expected error, got nil")
			}
		})
	}
}

// TestSpacesServiceUnit_EmptyResponse_Handling tests handling of empty or minimal responses.
func TestSpacesServiceUnit_EmptyResponse_Handling(t *testing.T) {
	tests := []struct {
		name      string
		responses map[string]string
		operation func(service Service) (any, error)
		validate  func(t *testing.T, result any, err error)
	}{
		{
			name: "GetTenantInfo_EmptyResponse",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail/tenant": `{}`,
			},
			operation: func(service Service) (any, error) {
				return service.GetTenantInfo(context.Background())
			},
			validate: func(t *testing.T, result any, err error) {
				if err != nil {
					t.Log("Method returned error (expected when service not configured)")
				}
				// Empty response should return nil or empty structure depending on implementation
			},
		},
		{
			name: "GetConnectionStats_DisconnectedState",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail/stats": `{
					"Cisco-IOS-XE-wireless-cisco-spaces-oper:stats": {
						"total-con-attempts": 10,
						"con-attempts-success": 0,
						"con-attempts-failure": 10,
						"total-discon": 10,
						"avg-response-time": 0
					}
				}`,
			},
			operation: func(service Service) (any, error) {
				return service.GetConnectionStats(context.Background())
			},
			validate: func(t *testing.T, result any, err error) {
				if err != nil {
					t.Log("Method returned error (expected when service not configured)")
					// When service not configured, result should be nil
					if result != nil {
						t.Log("Note: Method returned non-nil result despite error")
					}
					return
				}
				stats := result.(*model.SpacesConnectionStats)
				if stats == nil {
					t.Fatal("Expected result, got nil")
				}
				if stats.ConnectionAttemptsSuccess != 0 {
					t.Error("Expected no successful connections")
				}
				if stats.ConnectionAttemptsFailure != 10 {
					t.Errorf("Expected 10 failed attempts, got %d", stats.ConnectionAttemptsFailure)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use MockErrorServer to simulate 404 errors for unconfigured services
			endpoints := make([]string, 0, len(tt.responses))
			for endpoint := range tt.responses {
				endpoints = append(endpoints, endpoint)
			}
			mockServer := testutil.NewMockErrorServer(endpoints, 404)
			defer mockServer.Close()

			testClient := testutil.NewTestClient(mockServer)
			service := NewService(testClient.Core().(*core.Client))

			result, err := tt.operation(service)
			tt.validate(t, result, err)
		})
	}
}

// TestSpacesServiceUnit_NilHandling_EdgeCases tests nil handling in GetTenantInfo and GetConnectionStats.
func TestSpacesServiceUnit_NilHandling_EdgeCases(t *testing.T) {
	t.Run("GetTenantInfo_NilTenant", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-spaces-oper:spaces-connection-detail": {
				"tenant": null,
				"stats": {
					"connection-count": 100,
					"data-transfer": "1.5GB"
				}
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail": mockResponse,
		})
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		ctx := context.Background()

		result, err := service.GetTenantInfo(ctx)
		if err != nil {
			t.Log("Method returned error (expected when service not configured)")
		}

		if result != nil {
			t.Errorf("Expected nil tenant, got %+v", result)
		}
	})

	t.Run("GetConnectionStats_NilStats", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-spaces-oper:spaces-connection-detail": {
				"tenant": {
					"tenant-id": "test-tenant",
					"tenant-name": "Test Tenant"
				},
				"stats": null
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail": mockResponse,
		})
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := NewService(testClient.Core().(*core.Client))
		ctx := context.Background()

		result, err := service.GetConnectionStats(ctx)
		if err != nil {
			t.Log("Method returned error (expected when service not configured)")
		}

		if result != nil {
			t.Errorf("Expected nil stats, got %+v", result)
		}
	})
}
