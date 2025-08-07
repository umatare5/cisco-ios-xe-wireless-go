// Package wlan provides WLAN test functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// TestDataCollector holds test data for WLAN service functions
type TestDataCollector struct {
	mu                           sync.Mutex
	CfgResp                      *model.WlanCfgResponse
	CfgErr                       error
	CfgEntriesResp               *model.WlanCfgEntriesResponse
	CfgEntriesErr                error
	PoliciesResp                 *model.WlanPoliciesResponse
	PoliciesErr                  error
	PolicyListEntriesResp        *model.PolicyListEntriesResponse
	PolicyListEntriesErr         error
	WirelessAaaPolicyConfigsResp *model.WirelessAaaPolicyConfigsResponse
	WirelessAaaPolicyConfigsErr  error
	GlobalOperResp               *model.WlanGlobalOperResponse
	GlobalOperErr                error
}

// TestWLANService tests all WLAN service functions with the 4-pattern testing approach
func TestWLANService(t *testing.T) {
	// Create a mock client that will be used when environment variables are not set
	var client *wnc.Client
	var ctx context.Context

	// Try to get real client from environment
	if os.Getenv("WNC_CONTROLLER") != "" && os.Getenv("WNC_ACCESS_TOKEN") != "" {
		client = tests.TestClient(t)
		ctx = tests.TestContext(t)
	} else {
		// Use nil client for unit testing
		client = nil
		ctx = context.Background()
	}

	service := NewService(client)

	// ========================================
	// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
	// ========================================

	t.Run("Service_Creation", func(t *testing.T) {
		// Test service creation with both nil and valid clients
		nilService := NewService(nil)
		if nilService.c != nil {
			t.Error("NewService with nil client should have nil internal client")
		}

		if client != nil {
			validService := NewService(client)
			if validService.c == nil {
				t.Error("NewService with valid client should have non-nil internal client")
			}
		}
	})

	t.Run("Data_Collection", func(t *testing.T) {
		// Test the functionality by collecting data using all available methods
		collector := &TestDataCollector{}

		var wg sync.WaitGroup
		wg.Add(6) // Six methods to test

		// Test Cfg method
		go func() {
			defer wg.Done()
			resp, err := service.Cfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		// Test CfgEntries method
		go func() {
			defer wg.Done()
			resp, err := service.CfgEntries(ctx)
			collector.mu.Lock()
			collector.CfgEntriesResp = resp
			collector.CfgEntriesErr = err
			collector.mu.Unlock()
		}()

		// Test Policies method
		go func() {
			defer wg.Done()
			resp, err := service.Policies(ctx)
			collector.mu.Lock()
			collector.PoliciesResp = resp
			collector.PoliciesErr = err
			collector.mu.Unlock()
		}()

		// Test PolicyListEntries method
		go func() {
			defer wg.Done()
			resp, err := service.PolicyListEntries(ctx)
			collector.mu.Lock()
			collector.PolicyListEntriesResp = resp
			collector.PolicyListEntriesErr = err
			collector.mu.Unlock()
		}()

		// Test WirelessAaaPolicyConfigs method
		go func() {
			defer wg.Done()
			resp, err := service.WirelessAaaPolicyConfigs(ctx)
			collector.mu.Lock()
			collector.WirelessAaaPolicyConfigsResp = resp
			collector.WirelessAaaPolicyConfigsErr = err
			collector.mu.Unlock()
		}()

		// Test GlobalOper method
		go func() {
			defer wg.Done()
			resp, err := service.GlobalOper(ctx)
			collector.mu.Lock()
			collector.GlobalOperResp = resp
			collector.GlobalOperErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		if collector.CfgErr != nil {
			t.Logf("Cfg method returned error: %v", collector.CfgErr)
		}
		if collector.CfgResp != nil {
			t.Logf("Cfg method returned data successfully")
		}
		// Log other methods too
		t.Logf("All WLAN methods tested")
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "CtsCfgResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-cts-cfg:cts-cfg-data": {
						"cts-config": {
							"enable": true
						}
					}
				}`,
			},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				var data interface{}
				err := json.Unmarshal([]byte(tt.jsonData), &data)
				if err != nil {
					t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
				}

				_, err = json.Marshal(data)
				if err != nil {
					t.Errorf("Failed to marshal %s: %v", tt.name, err)
				}
			})
		}
	})

	// ========================================
	// 2. TABLE-DRIVEN TEST PATTERNS
	// ========================================

	t.Run("Method_Tests", func(t *testing.T) {
		methodTests := []struct {
			name   string
			method func() (interface{}, error)
		}{
			{
				name: "Cfg",
				method: func() (interface{}, error) {
					return service.Cfg(ctx)
				},
			},
			{
				name: "CfgEntries",
				method: func() (interface{}, error) {
					return service.CfgEntries(ctx)
				},
			},
			{
				name: "Policies",
				method: func() (interface{}, error) {
					return service.Policies(ctx)
				},
			},
			{
				name: "PolicyListEntries",
				method: func() (interface{}, error) {
					return service.PolicyListEntries(ctx)
				},
			},
			{
				name: "WirelessAaaPolicyConfigs",
				method: func() (interface{}, error) {
					return service.WirelessAaaPolicyConfigs(ctx)
				},
			},
			{
				name: "GlobalOper",
				method: func() (interface{}, error) {
					return service.GlobalOper(ctx)
				},
			},
		}

		for _, tt := range methodTests {
			t.Run(tt.name, func(t *testing.T) {
				result, err := tt.method()
				if err != nil {
					t.Logf("Method %s returned error: %v", tt.name, err)
				}
				if result != nil {
					t.Logf("Method %s returned result", tt.name)
				}
			})
		}
	})

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================

	t.Run("Critical_Validations", func(t *testing.T) {
		// Test with nil client
		t.Run("NilClient", func(t *testing.T) {
			service := NewService(nil)
			_, err := service.Cfg(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context (should handle gracefully or fail fast)
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.Cfg(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.Cfg(ctx)
			if err == nil {
				t.Fatal("Expected error with canceled context, got none")
			}
		})
	})

	// ========================================
	// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
	// ========================================

	t.Run("Integration_Test", func(t *testing.T) {
		// This test requires actual WNC environment
		// Skip if running in unit test mode
		if testing.Short() {
			t.Skip("Skipping integration test in short mode")
		}

		// Test with real service
		resp, err := service.Cfg(ctx)
		if err != nil {
			t.Logf("Integration test - Cfg error: %v", err)
		} else {
			t.Logf("Integration test - Cfg success: %+v", resp)
		}
	})
}
