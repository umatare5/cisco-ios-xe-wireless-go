package ap

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_ApGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "AP Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetCfgApTagsOnly",
				Method: func() (any, error) {
					return service.GetCfgApTagsOnly(ctx)
				},
			},
			{
				Name: "GetCfgTagSourcePriorityConfigsOnly",
				Method: func() (any, error) {
					return service.GetCfgTagSourcePriorityConfigsOnly(ctx)
				},
			},
			{
				Name: "GetApTagsCfg",
				Method: func() (any, error) {
					return service.GetApTagsCfg(ctx)
				},
			},
			{
				Name: "GetTagSourcePriorityCfg",
				Method: func() (any, error) {
					return service.GetTagSourcePriorityCfg(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfgApTagByMac",
				Method: func() (any, error) {
					return service.GetCfgApTagByMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetCfgTagSourcePriorityByPriority",
				Method: func() (any, error) {
					return service.GetCfgTagSourcePriorityByPriority(ctx, 1)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			// Nil client tests
			{
				Name: "GetCfg_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgApTagsOnly_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgApTagsOnly(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgTagSourcePriorityConfigsOnly_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgTagSourcePriorityConfigsOnly(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetApTagsCfg_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetApTagsCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetTagSourcePriorityCfg_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetTagSourcePriorityCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgApTagByMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgApTagByMac(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgTagSourcePriorityByPriority_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgTagSourcePriorityByPriority(ctx, 1)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetCfgApTagByMac_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetCfgApTagByMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgApTagByMac_InvalidMAC",
				Function: func() (any, error) {
					_, err := service.GetCfgApTagByMac(ctx, "invalid-mac")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgApTagByMac_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetCfgApTagByMac(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgTagSourcePriorityByPriority_NegativeParam",
				Function: func() (any, error) {
					_, err := service.GetCfgTagSourcePriorityByPriority(ctx, -1)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
