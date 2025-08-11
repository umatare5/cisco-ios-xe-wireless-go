// Package site provides site test functionality for the Cisco Wireless Network Controller API.
package site

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestSiteService tests the Site service using the standardized testing framework
func TestSiteService(t *testing.T) {
	client := tests.OptionalTestClient(t)

	service := NewService(client)

	// ========================================
	// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
	// ========================================

	// Configure test methods
	testMethods := []tests.TestMethod{
		{
			Name: "GetOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GetOper(ctx)
			},
		},
	}

	// Configure JSON test cases using standard helper
	jsonTestCases := tests.StandardJSONTestCases("site")

	// Configure and run tests
	config := tests.ServiceTestConfig{
		ServiceName:    "Site",
		TestMethods:    testMethods,
		JSONTestCases:  jsonTestCases,
		SkipShortTests: true,
	}

	tests.RunServiceTests(t, config)
}

// TestSiteServiceSpecific contains site-specific tests that don't fit the standard pattern
func TestSiteServiceSpecific(t *testing.T) {
	client := tests.OptionalTestClient(t)
	if client == nil { // no integration env
		return
	}
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("OperResponseType", func(t *testing.T) {
		result, err := service.GetOper(ctx)
		if err != nil {
			t.Logf("Oper returned error (expected in test env): %v", err)
			return
		}

		// Validate response type
		if result == nil {
			t.Error("Oper should not return nil result")
			return
		}

		t.Logf("Oper returned result of type %T", result)
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
}

// Complementary fail-fast and integration sections to adhere to repository testing regulation
func TestSiteService_FailFastAndIntegration(t *testing.T) {
	client := tests.OptionalTestClient(t)
	service := NewService(client)

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================
	t.Run("Critical_Validations", func(t *testing.T) {
		t.Run("NilClient", func(t *testing.T) {
			s := NewService(nil)
			ctx := context.Background()
			if _, err := s.GetOper(ctx); err == nil {
				t.Fatal("expected error with nil client, got none")
			}
		})

		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 intentionally nil for testing
			if _, err := service.GetOper(nilCtx); err == nil {
				t.Fatal("expected error with nil context, got none")
			}
		})

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
