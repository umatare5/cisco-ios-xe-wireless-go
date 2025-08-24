package dot11

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_Dot11GetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "IEEE 802.11 Configuration",
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
				Name: "GetCfgByCountryCode",
				Method: func() (any, error) {
					return service.GetCfgByCountryCode(ctx, "US")
				},
			},
			{
				Name: "GetCfgByBand",
				Method: func() (any, error) {
					return service.GetCfgByBand(ctx, "2.4GHz")
				},
			},
			{
				Name: "GetCfgBySpatialStreamAndIndex",
				Method: func() (any, error) {
					return service.GetCfgBySpatialStreamAndIndex(ctx, 2, "mcs-index-7")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetCfgByCountryCode with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByCountryCode(ctx, "US")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByBand with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByBand(ctx, "2.4GHz")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgBySpatialStreamAndIndex with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfgBySpatialStreamAndIndex(ctx, 2, "mcs-index-7")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByCountryCode with empty parameter",
				Function: func() (any, error) {
					_, err := service.GetCfgByCountryCode(ctx, "")
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
				Name: "GetCfgBySpatialStreamAndIndex with invalid stream",
				Function: func() (any, error) {
					_, err := service.GetCfgBySpatialStreamAndIndex(ctx, -1, "mcs-index-7")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgBySpatialStreamAndIndex with empty index",
				Function: func() (any, error) {
					_, err := service.GetCfgBySpatialStreamAndIndex(ctx, 2, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
	})
}
