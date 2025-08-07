// Package wlan provides WLAN test functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"os"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestWLANService tests the WLAN service using the standardized testing framework
func TestWLANService(t *testing.T) {
	// Get test client and create service
	var client *core.Client

	// Try to get real client from environment
	if os.Getenv("WNC_CONTROLLER") != "" && os.Getenv("WNC_ACCESS_TOKEN") != "" {
		client = tests.TestClient(t)
	}

	service := NewService(client)

	// Configure test methods
	testMethods := []tests.TestMethod{
		{
			Name: "Cfg",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.Cfg(ctx)
			},
		},
		{
			Name: "CfgEntries",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.CfgEntries(ctx)
			},
		},
		{
			Name: "Policies",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.Policies(ctx)
			},
		},
		{
			Name: "PolicyListEntries",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.PolicyListEntries(ctx)
			},
		},
		{
			Name: "WirelessAaaPolicyConfigs",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.WirelessAaaPolicyConfigs(ctx)
			},
		},
		{
			Name: "GlobalOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GlobalOper(ctx)
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
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("CfgResponseType", func(t *testing.T) {
		result, err := service.Cfg(ctx)
		if err != nil {
			t.Logf("Cfg returned error (expected in test env): %v", err)
			return
		}

		tests.AssertNonNilResult(t, result, "Cfg")
		tests.LogMethodResult(t, "Cfg", result, err)
	})

	t.Run("GlobalOperResponseType", func(t *testing.T) {
		result, err := service.GlobalOper(ctx)
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
			cfgResult, cfgErr := service.Cfg(ctx)
			if cfgErr == nil && cfgResult != nil {
				tests.ValidateStructType(t, cfgResult)
			}

			// Test GlobalOper response structure
			operResult, operErr := service.GlobalOper(ctx)
			if operErr == nil && operResult != nil {
				tests.ValidateStructType(t, operResult)
			}
		}
	})
}
