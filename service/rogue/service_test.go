package rogue

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestService_MinimalServiceTestSuite(t *testing.T) {
	framework.RunMinimalServiceTests(t, framework.MinimalServiceTestSuite{
		ServiceName: "rogue",
		NewServiceFunc: func(client any) any {
			if client == nil {
				return NewService(nil)
			}
			return NewService(client.(*core.Client))
		},
		GetClientFunc: func(service any) any {
			rogueService, ok := service.(Service)
			if !ok {
				return nil
			}
			return rogueService.Client()
		},
	})
}
