package cts

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_CtsGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Cisco TrustSec (CTS) Configuration",
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
				Name: "GetCfgBySxpProfileName",
				Method: func() (any, error) {
					return service.GetCfgBySxpProfileName(ctx, "test-sxp-profile")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgBySxpProfileName with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgBySxpProfileName(ctx, "test-sxp-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgBySxpProfileName with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgBySxpProfileName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgBySxpProfileName with whitespace parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgBySxpProfileName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
