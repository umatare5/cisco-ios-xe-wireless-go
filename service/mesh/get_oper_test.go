package mesh

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_MeshGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Mesh",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperNodes",
				Method: func() (any, error) {
					return service.GetOperNodes(ctx)
				},
			},
			{
				Name: "GetOperStats",
				Method: func() (any, error) {
					return service.GetOperStats(ctx)
				},
			},
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfgByProfileName",
				Method: func() (any, error) {
					return service.GetCfgByProfileName(ctx, "default-mesh-profile")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByProfileName_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByProfileName(ctx, "default-mesh-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByProfileName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetCfgByProfileName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
