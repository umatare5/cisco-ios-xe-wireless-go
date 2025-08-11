// Package wlan provides WLAN test functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestWLANService tests the WLAN service using the standardized testing framework
func TestWLANService(t *testing.T) {
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
			Name: "GetCfgEntries",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetCfgEntries(ctx)
			},
		},
		{
			Name: "GetPolicies",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetPolicies(ctx)
			},
		},
		{
			Name: "GetPolicyListEntries",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetPolicyListEntries(ctx)
			},
		},
		{
			Name: "GetWirelessAaaPolicyConfigs",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetWirelessAaaPolicyConfigs(ctx)
			},
		},
		{
			Name: "GetGlobalOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetGlobalOper(ctx)
			},
		},
	}

	// Configure JSON test cases using standard helper
	jsonTestCases := tests.StandardJSONTestCases("wlan")

	// Configure and run tests
	config := tests.ServiceTestConfig{
		ServiceName:    "WLAN",
		TestMethods:    testMethods,
		JSONTestCases:  jsonTestCases,
		SkipShortTests: true,
	}

	tests.RunServiceTests(t, config)
}

// TestWLANServiceSpecific contains WLAN-specific tests that don't fit the standard pattern
func TestWLANServiceSpecific(t *testing.T) {
	client := tests.OptionalTestClient(t)
	if client == nil { // if no integration env, skip these specifics gracefully
		return
	}
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

	t.Run("GlobalOperResponseType", func(t *testing.T) {
		result, err := service.GetGlobalOper(ctx)
		if err != nil {
			t.Logf("GlobalOper returned error (expected in test env): %v", err)
			return
		}

		tests.AssertNonNilResult(t, result, "GlobalOper")
		tests.LogMethodResult(t, "GlobalOper", result, err)
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

			// Test GlobalOper response structure
			operResult, operErr := service.GetGlobalOper(ctx)
			if operErr == nil && operResult != nil {
				tests.ValidateStructType(t, operResult)
			}
		}
	})
}

// Dedicated fail-fast and integration coverage to match repository test regulation
func TestWLANService_FailFastAndIntegration(t *testing.T) {
	client := tests.OptionalTestClient(t)
	service := NewService(client)

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================
	t.Run("Critical_Validations", func(t *testing.T) {
		t.Run("NilClient", func(t *testing.T) {
			s := NewService(nil)
			ctx := context.Background()
			if _, err := s.GetCfg(ctx); err == nil {
				t.Fatal("expected error with nil client, got none")
			}
		})

		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 intentionally nil for testing
			if _, err := service.GetCfg(nilCtx); err == nil {
				t.Fatal("expected error with nil context, got none")
			}
		})

		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			if _, err := service.GetCfg(ctx); err == nil {
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
		if _, err := service.GetCfg(ctx); err != nil {
			t.Logf("integration GetCfg returned error: %v", err)
		}
	})
}
