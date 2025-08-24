package flex

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// TestService tests the Flex service using the minimal test pattern for cfg-only service
func TestService(t *testing.T) {
	// Create and run minimal service test suite
	suite := framework.MinimalServiceTestSuite{
		ServiceName: "Flex",
		NewServiceFunc: func(client any) any {
			if client == nil {
				return NewService(nil)
			}
			return NewService(client.(*core.Client))
		},
		GetClientFunc: func(service any) any {
			if service == nil {
				return nil
			}
			s := service.(Service)
			return s.Client()
		},
	}

	framework.RunMinimalServiceTests(t, suite)
}

// Test_FlexJSONSerialization_Unit tests the JSON serialization/deserialization
