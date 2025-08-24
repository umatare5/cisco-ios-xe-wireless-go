package general

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_GeneralGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "General",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperMgmtIntfData",
				Method: func() (any, error) {
					return service.GetOperMgmtIntfData(ctx)
				},
			},
			{
				Name: "GetOperLaginfo",
				Method: func() (any, error) {
					return service.GetOperLaginfo(ctx)
				},
			},
			{
				Name: "GetOperMfp",
				Method: func() (any, error) {
					return service.GetOperMfp(ctx)
				},
			},
			{
				Name: "GetOperMewlcConfig",
				Method: func() (any, error) {
					return service.GetOperMewlcConfig(ctx)
				},
			},
			{
				Name: "GetOperMulticastConfig",
				Method: func() (any, error) {
					return service.GetOperMulticastConfig(ctx)
				},
			},
			{
				Name: "GetOperSimL3InterfaceCacheData",
				Method: func() (any, error) {
					return service.GetOperSimL3InterfaceCacheData(ctx)
				},
			},
			{
				Name: "GetOperWlcManagementData",
				Method: func() (any, error) {
					return service.GetOperWlcManagementData(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperApLocRangingCfg",
				Method: func() (any, error) {
					return service.GetOperApLocRangingCfg(ctx)
				},
			},
			{
				Name: "GetOperCacConfig",
				Method: func() (any, error) {
					return service.GetOperCacConfig(ctx)
				},
			},
			{
				Name: "GetOperFeatureUsageCfg",
				Method: func() (any, error) {
					return service.GetOperFeatureUsageCfg(ctx)
				},
			},
			{
				Name: "GetOperFipsCfg",
				Method: func() (any, error) {
					return service.GetOperFipsCfg(ctx)
				},
			},
			{
				Name: "GetOperGeolocationCfg",
				Method: func() (any, error) {
					return service.GetOperGeolocationCfg(ctx)
				},
			},
			{
				Name: "GetOperThresholdWarnCfg",
				Method: func() (any, error) {
					return service.GetOperThresholdWarnCfg(ctx)
				},
			},
			{
				Name: "GetOperWsaApClientEvent",
				Method: func() (any, error) {
					return service.GetOperWsaApClientEvent(ctx)
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
				Name: "GetOperMgmtIntfData_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMgmtIntfData(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperLaginfo_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperLaginfo(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperMfp_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMfp(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
