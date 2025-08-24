package radio

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestService_MinimalServiceTestSuite(t *testing.T) {
	framework.RunMinimalServiceTests(t, framework.MinimalServiceTestSuite{
		ServiceName: "radio",
		NewServiceFunc: func(client any) any {
			if client == nil {
				return NewService(nil)
			}
			return NewService(client.(*core.Client))
		},
		GetClientFunc: func(service any) any {
			radioService, ok := service.(Service)
			if !ok {
				return nil
			}
			return radioService.Client()
		},
	})
}
