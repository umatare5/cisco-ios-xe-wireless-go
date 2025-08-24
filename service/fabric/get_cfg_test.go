package fabric

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_FabricGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Fabric Configuration",
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
				Name: "GetCfgByFabricProfileName",
				Method: func() (any, error) {
					return service.GetCfgByFabricProfileName(ctx, "test-profile")
				},
			},
			{
				Name: "GetCfgByControlPlaneName",
				Method: func() (any, error) {
					return service.GetCfgByControlPlaneName(ctx, "test-controlplane")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgByFabricProfileName with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByFabricProfileName(ctx, "test-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByControlPlaneName with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByControlPlaneName(ctx, "test-controlplane")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByFabricProfileName with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByFabricProfileName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByControlPlaneName with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByControlPlaneName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
