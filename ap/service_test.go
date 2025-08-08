// Package ap provides access point test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestApService tests the AP service using the standardized testing framework
func TestApService(t *testing.T) {
	// Obtain an integration client opportunistically (nil if env not set).
	client := tests.OptionalTestClient(t)

	service := NewService(client)

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
