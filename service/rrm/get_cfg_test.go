package rrm

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RrmGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Radio Resource Management (RRM) Configuration",
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
				Name: "GetCfgByBand",
				Method: func() (any, error) {
					return service.GetCfgByBand(ctx, "2.4")
				},
			},
			{
				Name: "GetCfgByMgrBand",
				Method: func() (any, error) {
					return service.GetCfgByMgrBand(ctx, "5")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgByBand with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByBand(ctx, "2.4")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByMgrBand with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByMgrBand(ctx, "5")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByBand with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByBand(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByMgrBand with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByMgrBand(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
