// Package flex provides configuration test functionality for the Cisco Wireless Network Controller API.
package flex

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestFlexService tests the Flex service using standardized test patterns
func TestFlexService(t *testing.T) {
	client := tests.TestClient(t)
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
	}

	// Configure JSON test cases using standard helper
	jsonTestCases := tests.StandardJSONTestCases("flex")

	// Configure and run tests
	config := tests.ServiceTestConfig{
		ServiceName:    "Flex",
		TestMethods:    testMethods,
		JSONTestCases:  jsonTestCases,
		SkipShortTests: true,
	}

	tests.RunServiceTests(t, config)
}

// Explicit fail-fast and integration coverage aligned with repository standard
func TestFlexService_FailFastAndIntegration(t *testing.T) {
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
