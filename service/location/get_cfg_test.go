package location

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestLocationGetCfg(t *testing.T) {
	clientSetup := client.OptionalTestClient(t)
	service := NewService(clientSetup)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Location Services Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfg with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
