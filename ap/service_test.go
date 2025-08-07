// Package ap provides access point test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"os"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestApService tests the AP service using the standardized testing framework
func TestApService(t *testing.T) {
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
			Name: "TagSourcePriorityConfigs",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.TagSourcePriorityConfigs(ctx)
			},
		},
		{
			Name: "ApTags",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.ApTags(ctx)
			},
		},
		{
			Name: "Oper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.Oper(ctx)
			},
		},
		{
			Name: "RadioNeighbor",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.RadioNeighbor(ctx)
			},
		},
		{
			Name: "NameMacMap",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.NameMacMap(ctx)
			},
		},
		{
			Name: "CapwapData",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.CapwapData(ctx)
			},
		},
		{
			Name: "GlobalOper",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.GlobalOper(ctx)
			},
		},
		{
			Name: "History",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.History(ctx)
			},
		},
		{
			Name: "EwlcApStats",
			Method: func() (interface{}, error) {
				ctx := tests.TestContext(t)
				return service.EwlcApStats(ctx)
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
		result, err := service.Cfg(ctx)
		if err != nil {
			t.Logf("Cfg returned error (expected in test env): %v", err)
			return
		}

		tests.AssertNonNilResult(t, result, "Cfg")
		tests.LogMethodResult(t, "Cfg", result, err)
	})

	t.Run("OperResponseType", func(t *testing.T) {
		result, err := service.Oper(ctx)
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
			cfgResult, cfgErr := service.Cfg(ctx)
			if cfgErr == nil && cfgResult != nil {
				tests.ValidateStructType(t, cfgResult)
			}

			// Test Oper response structure
			operResult, operErr := service.Oper(ctx)
			if operErr == nil && operResult != nil {
				tests.ValidateStructType(t, operResult)
			}
		}
	})
}
