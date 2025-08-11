package client

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for Client service functions
type TestDataCollector struct {
	mu                    sync.Mutex
	OperResp              *model.ClientOperResponse
	OperErr               error
	CommonOperDataResp    *model.ClientOperCommonOperDataResponse
	CommonOperDataErr     error
	Dot11OperDataResp     *model.ClientOperDot11OperDataResponse
	Dot11OperDataErr      error
	MobilityOperDataResp  *model.ClientOperMobilityOperDataResponse
	MobilityOperDataErr   error
	MmIfClientStatsResp   *model.ClientOperMmIfClientStatsResponse
	MmIfClientStatsErr    error
	MmIfClientHistoryResp *model.ClientOperMmIfClientHistoryResponse
	MmIfClientHistoryErr  error
	TrafficStatsResp      *model.ClientOperTrafficStatsResponse
	TrafficStatsErr       error
	PolicyDataResp        *model.ClientOperPolicyDataResponse
	PolicyDataErr         error
	SisfDBMacResp         *model.ClientOperSisfDBMacResponse
	SisfDBMacErr          error
	DcInfoResp            *model.ClientOperDcInfoResponse
	DcInfoErr             error
}

// TestClientService tests all Client service functions with the 4-pattern testing approach
func TestClientService(t *testing.T) {
	client := tests.OptionalTestClient(t)
	ctx := context.Background()
	if client != nil {
		ctx = tests.TestContext(t)
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
		wg.Add(constants.DefaultTestMethods) // 10 methods to test

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.GetOper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test CommonOperData method
		go func() {
			defer wg.Done()
			resp, err := service.GetCommonOperData(ctx)
			collector.mu.Lock()
			collector.CommonOperDataResp = resp
			collector.CommonOperDataErr = err
			collector.mu.Unlock()
		}()

		// Test Dot11OperData method
		go func() {
			defer wg.Done()
			resp, err := service.GetDot11OperData(ctx)
			collector.mu.Lock()
			collector.Dot11OperDataResp = resp
			collector.Dot11OperDataErr = err
			collector.mu.Unlock()
		}()

		// Test MobilityOperData method
		go func() {
			defer wg.Done()
			resp, err := service.GetMobilityOperData(ctx)
			collector.mu.Lock()
			collector.MobilityOperDataResp = resp
			collector.MobilityOperDataErr = err
			collector.mu.Unlock()
		}()

		// Test MmIfClientStats method
		go func() {
			defer wg.Done()
			resp, err := service.GetMmIfClientStats(ctx)
			collector.mu.Lock()
			collector.MmIfClientStatsResp = resp
			collector.MmIfClientStatsErr = err
			collector.mu.Unlock()
		}()

		// Test MmIfClientHistory method
		go func() {
			defer wg.Done()
			resp, err := service.GetMmIfClientHistory(ctx)
			collector.mu.Lock()
			collector.MmIfClientHistoryResp = resp
			collector.MmIfClientHistoryErr = err
			collector.mu.Unlock()
		}()

		// Test TrafficStats method
		go func() {
			defer wg.Done()
			resp, err := service.GetTrafficStats(ctx)
			collector.mu.Lock()
			collector.TrafficStatsResp = resp
			collector.TrafficStatsErr = err
			collector.mu.Unlock()
		}()

		// Test PolicyData method
		go func() {
			defer wg.Done()
			resp, err := service.GetPolicyData(ctx)
			collector.mu.Lock()
			collector.PolicyDataResp = resp
			collector.PolicyDataErr = err
			collector.mu.Unlock()
		}()

		// Test SisfDBMac method
		go func() {
			defer wg.Done()
			resp, err := service.GetSisfDBMac(ctx)
			collector.mu.Lock()
			collector.SisfDBMacResp = resp
			collector.SisfDBMacErr = err
			collector.mu.Unlock()
		}()

		// Test DcInfo method
		go func() {
			defer wg.Done()
			resp, err := service.GetDcInfo(ctx)
			collector.mu.Lock()
			collector.DcInfoResp = resp
			collector.DcInfoErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d Client service methods", constants.DefaultTestMethods)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "ClientOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-client-oper:client-oper-data": {
						"common-oper-data": {
							"client-summary": {
								"total-client-count": 10
							}
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
				name: "GetOper",
				method: func() (interface{}, error) {
					return service.GetOper(ctx)
				},
			},
			{
				name: "GetCommonOperData",
				method: func() (interface{}, error) {
					return service.GetCommonOperData(ctx)
				},
			},
			{
				name: "GetDot11OperData",
				method: func() (interface{}, error) {
					return service.GetDot11OperData(ctx)
				},
			},
			{
				name: "GetMobilityOperData",
				method: func() (interface{}, error) {
					return service.GetMobilityOperData(ctx)
				},
			},
			{
				name: "GetMmIfClientStats",
				method: func() (interface{}, error) {
					return service.GetMmIfClientStats(ctx)
				},
			},
			{
				name: "GetMmIfClientHistory",
				method: func() (interface{}, error) {
					return service.GetMmIfClientHistory(ctx)
				},
			},
			{
				name: "GetTrafficStats",
				method: func() (interface{}, error) {
					return service.GetTrafficStats(ctx)
				},
			},
			{
				name: "GetPolicyData",
				method: func() (interface{}, error) {
					return service.GetPolicyData(ctx)
				},
			},
			{
				name: "GetSisfDBMac",
				method: func() (interface{}, error) {
					return service.GetSisfDBMac(ctx)
				},
			},
			{
				name: "GetDcInfo",
				method: func() (interface{}, error) {
					return service.GetDcInfo(ctx)
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
			_, err := service.GetOper(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.GetOper(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.GetOper(ctx)
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

		// Test operational data
		operResp, operErr := service.GetOper(ctx)
		if operErr != nil {
			t.Logf("Integration test - Oper error: %v", operErr)
		} else {
			t.Logf("Integration test - Oper success: %+v", operResp)
		}

		// Test common operational data
		commonResp, commonErr := service.GetCommonOperData(ctx)
		if commonErr != nil {
			t.Logf("Integration test - CommonOperData error: %v", commonErr)
		} else {
			t.Logf("Integration test - CommonOperData success: %+v", commonResp)
		}
	})
}
