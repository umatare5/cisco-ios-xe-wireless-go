package flex

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_FlexGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "FlexConnect Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfgByPolicyName",
				Method: func() (any, error) {
					return service.GetCfgByPolicyName(ctx, "test-flex-policy")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgByPolicyName with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByPolicyName(ctx, "test-flex-policy")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByPolicyName with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByPolicyName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByPolicyName with whitespace parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByPolicyName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
