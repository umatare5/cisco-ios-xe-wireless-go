package mdns

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_MdnsGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "mDNS",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperGlobalStats",
				Method: func() (any, error) {
					return service.GetOperGlobalStats(ctx)
				},
			},
			{
				Name: "GetOperWlanStats",
				Method: func() (any, error) {
					return service.GetOperWlanStats(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByWlanID",
				Method: func() (any, error) {
					return service.GetOperByWlanID(ctx, "1")
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
				Name: "GetOperGlobalStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperGlobalStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperWlanStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperWlanStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWlanID_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByWlanID(ctx, "1")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWlanID_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByWlanID(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWlanID_InvalidWlanID",
				Function: func() (any, error) {
					_, err := service.GetOperByWlanID(ctx, "invalid")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
