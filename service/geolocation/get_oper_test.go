package geolocation

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_GeolocationGetOper_StandardTests(t *testing.T) {
	clientSetup := client.OptionalTestClient(t)
	service := NewService(clientSetup)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Geolocation",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			// Geolocation service doesn't have filter methods currently
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					return nilService.GetOper(ctx)
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
