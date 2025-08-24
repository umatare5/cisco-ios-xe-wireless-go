package rfid

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RfidGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "RFID Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetCfgRfidConfig",
				Method: func() (any, error) {
					return service.GetCfgRfidConfig(ctx)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgRfidConfig with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgRfidConfig(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
