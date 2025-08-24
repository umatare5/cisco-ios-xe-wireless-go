package dot11

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestService(t *testing.T) {
	// Create and run minimal service test suite for configuration-only service
	suite := framework.MinimalServiceTestSuite{
		ServiceName: "Dot11",
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
