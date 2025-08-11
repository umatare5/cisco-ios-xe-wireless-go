// Package ap provides access point test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestApService tests the AP service using the standardized testing framework
func TestApService(t *testing.T) {
	// Obtain an integration client opportunistically (nil if env not set).
	client := tests.OptionalTestClient(t)

	service := NewService(client)

	// ========================================
	// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
	// ========================================

	// Configure test methods
	testMethods := []tests.TestMethod{
		{
			Name: "GetCfg",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetCfg(ctx)
			},
		},
		{
			Name: "GetTagSourcePriorityConfigs",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetTagSourcePriorityConfigs(ctx)
			},
		},
		{
			Name: "GetApTags",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetApTags(ctx)
			},
		},
		{
			Name: "GetOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetOper(ctx)
			},
		},
		{
			Name: "GetRadioNeighbor",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetRadioNeighbor(ctx)
			},
		},
		{
			Name: "GetNameMacMap",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetNameMacMap(ctx)
			},
		},
		{
			Name: "GetCapwapData",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetCapwapData(ctx)
			},
		},
		{
			Name: "GetGlobalOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetGlobalOper(ctx)
			},
		},
		{
			Name: "GetHistory",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetHistory(ctx)
			},
		},
		{
			Name: "GetEwlcApStats",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetEwlcApStats(ctx)
			},
		},
	}

	// Configure JSON test cases using standard helper
	jsonTestCases := tests.StandardJSONTestCases("ap")

	// Configure and run tests
	config := tests.ServiceTestConfig{
		ServiceName:    "AP",
		TestMethods:    testMethods,
		JSONTestCases:  jsonTestCases,
		SkipShortTests: true,
	}

	tests.RunServiceTests(t, config)

	// Additional unit test: wrapper methods with mocked RESTCONF servers
	// Align naming with standard category for visibility
	t.Run("Method_Tests/WrapperMethods_SuccessAndErrorBranches", func(t *testing.T) {
		runAPWrapperMethodsTests(t)
	})
}

// TestApServiceSpecific contains AP-specific tests that don't fit the standard pattern
func TestApServiceSpecific(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("CfgResponseType", func(t *testing.T) {
		result, err := service.GetCfg(ctx)
		if err != nil {
			t.Logf("Cfg returned error (expected in test env): %v", err)
			return
		}

		tests.AssertNonNilResult(t, result, "Cfg")
		tests.LogMethodResult(t, "Cfg", result, err)
	})

	t.Run("OperResponseType", func(t *testing.T) {
		result, err := service.GetOper(ctx)
		if err != nil {
			t.Logf("Oper returned error (expected in test env): %v", err)
			return
		}

		tests.AssertNonNilResult(t, result, "Oper")
		tests.LogMethodResult(t, "Oper", result, err)
	})

	t.Run("ServiceCreationPattern", func(t *testing.T) {
		// Test the standard service creation pattern
		service1 := NewService(client)
		service2 := NewService(nil)

		if service1.c == nil && client != nil {
			t.Error("Service with valid client should have non-nil internal client")
		}

		if service2.c != nil {
			t.Error("Service with nil client should have nil internal client")
		}
	})

	t.Run("StructTypeValidation", func(t *testing.T) {
		// Test different response types with ValidateStructType
		if client != nil {
			// Test Cfg response structure
			cfgResult, cfgErr := service.GetCfg(ctx)
			if cfgErr == nil && cfgResult != nil {
				tests.ValidateStructType(t, cfgResult)
			}

			// Test Oper response structure
			operResult, operErr := service.GetOper(ctx)
			if operErr == nil && operResult != nil {
				tests.ValidateStructType(t, operResult)
			}
		}
	})
}

// Additional explicit sections for fail-fast validations and integration parity with other services
func TestApService_FailFastAndIntegration(t *testing.T) {
	client := tests.OptionalTestClient(t)
	service := NewService(client)

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================
	t.Run("Critical_Validations", func(t *testing.T) {
		// Nil client should error
		t.Run("NilClient", func(t *testing.T) {
			s := NewService(nil)
			ctx := context.Background()
			if _, err := s.GetOper(ctx); err == nil {
				t.Fatal("expected error with nil client, got none")
			}
		})

		// Nil context should error
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 intentionally nil for testing
			if _, err := service.GetOper(nilCtx); err == nil {
				t.Fatal("expected error with nil context, got none")
			}
		})

		// Canceled context should error
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			if _, err := service.GetOper(ctx); err == nil {
				t.Fatal("expected error with canceled context, got none")
			}
		})
	})

	// ========================================
	// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
	// ========================================
	t.Run("Integration_Test", func(t *testing.T) {
		if testing.Short() {
			t.Skip("skipping integration test in short mode")
		}
		if client == nil {
			t.Skip("no controller env; skipping integration")
		}
		ctx := tests.TestContext(t)
		if _, err := service.GetOper(ctx); err != nil {
			t.Logf("integration GetOper returned error: %v", err)
		}
	})
}

// Note: TLS client/server helpers are provided by internal/tests.

// runAPWrapperMethodsTests exercises wrapper methods that require post-processing of core.Get result.
func runAPWrapperMethodsTests(t *testing.T) {
	t.Helper()
	ctx := tests.TestContext(t)

	// Success server: returns minimal valid JSON for both endpoints
	successSrv := tests.NewRESTCONFSuccessServer(map[string]string{
		APNameMacMapEndpoint: `{"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map":[]}`,
		CapwapDataEndpoint:   `{"Cisco-IOS-XE-wireless-access-point-oper:capwap-data":[]}`,
	})
	defer successSrv.Close()

	// Error server: both endpoints return 500
	errorSrv := tests.NewRESTCONFErrorServer(
		[]string{APNameMacMapEndpoint, CapwapDataEndpoint},
		http.StatusInternalServerError,
	)
	defer errorSrv.Close()

	t.Run("Success", func(t *testing.T) {
		client := tests.NewTLSClientForServer(t, successSrv)
		svc := NewService(client)

		// GetNameMacMap success branch (covers return of &resp.Data)
		nm, err := svc.GetNameMacMap(ctx)
		if err != nil {
			t.Fatalf("GetNameMacMap success expected, got error: %v", err)
		}
		if nm == nil {
			t.Fatalf("GetNameMacMap returned nil slice pointer on success")
		}

		// GetCapwapData success branch
		cd, err := svc.GetCapwapData(ctx)
		if err != nil {
			t.Fatalf("GetCapwapData success expected, got error: %v", err)
		}
		if cd == nil {
			t.Fatalf("GetCapwapData returned nil slice pointer on success")
		}
	})

	t.Run("Error", func(t *testing.T) {
		client := tests.NewTLSClientForServer(t, errorSrv)
		svc := NewService(client)

		// GetNameMacMap error branch (covers err != nil path)
		nm, err := svc.GetNameMacMap(ctx)
		if err == nil || nm != nil {
			t.Fatalf("GetNameMacMap expected error branch, got result=%v err=%v", nm, err)
		}

		// GetCapwapData error branch
		cd, err := svc.GetCapwapData(ctx)
		if err == nil || cd != nil {
			t.Fatalf("GetCapwapData expected error branch, got result=%v err=%v", cd, err)
		}
	})
}
