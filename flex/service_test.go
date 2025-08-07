// Package flex provides configuration test functionality for the Cisco Wireless Network Controller API.
package flex

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestFLEXService tests the FLEX service using standardized test patterns
func TestFLEXService(t *testing.T) {
	client := tests.TestClient(t)
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
	}

	// Configure JSON test cases using standard helper
	jsonTestCases := tests.StandardJSONTestCases("flex")

	// Configure and run tests
	config := tests.ServiceTestConfig{
		ServiceName:    "FLEX",
		TestMethods:    testMethods,
		JSONTestCases:  jsonTestCases,
		SkipShortTests: true,
	}

	tests.RunServiceTests(t, config)
}
