package radio

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RadioGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Radio Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetCfgProfiles",
				Method: func() (any, error) {
					return service.GetCfgProfiles(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfgByName",
				Method: func() (any, error) {
					return service.GetCfgByName(ctx, framework.StandardGetOperTestConstants().TestAPMac)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgByName with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByName(ctx, "test-radio-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByName with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByName with whitespace parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
