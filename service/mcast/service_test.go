package mcast

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// TestService tests the Mcast service using the minimal test pattern for oper-only service
func TestService(t *testing.T) {
	// Create and run minimal service test suite
	suite := framework.MinimalServiceTestSuite{
		ServiceName: "Mcast",
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

// Test_McastJSONSerialization_Unit tests the JSON serialization/deserialization
