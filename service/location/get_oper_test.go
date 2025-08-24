package location

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestLocationGetOper(t *testing.T) {
	clientSetup := client.OptionalTestClient(t)
	service := NewService(clientSetup)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Location Services Operations",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper with nil client",
				Function: func() (any, error) {
					return nilService.GetOper(ctx)
				},
				ExpectedError: true,
			},
		},
	})
}
