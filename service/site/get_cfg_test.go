package site

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_SiteGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Site Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetAPCfgProfiles",
				Method: func() (any, error) {
					return service.GetAPCfgProfiles(ctx)
				},
			},
			{
				Name: "GetSiteTagConfigs",
				Method: func() (any, error) {
					return service.GetSiteTagConfigs(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetAPCfgProfileByName",
				Method: func() (any, error) {
					return service.GetAPCfgProfileByName(ctx, "test-ap-profile")
				},
			},
			{
				Name: "GetSiteTagConfigByName",
				Method: func() (any, error) {
					return service.GetSiteTagConfigByName(ctx, "test-site-tag")
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
				Name: "GetAPCfgProfiles_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetAPCfgProfiles(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetSiteTagConfigs_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetSiteTagConfigs(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetAPCfgProfileByName_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetAPCfgProfileByName(ctx, "test-ap-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetSiteTagConfigByName_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetSiteTagConfigByName(ctx, "test-site-tag")
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetAPCfgProfileByName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetAPCfgProfileByName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetAPCfgProfileByName_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetAPCfgProfileByName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetSiteTagConfigByName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetSiteTagConfigByName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetSiteTagConfigByName_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetSiteTagConfigByName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
