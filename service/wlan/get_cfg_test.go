package wlan

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_WlanGetCfg_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "WLAN Configuration",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetCfgEntries",
				Method: func() (any, error) {
					return service.GetCfgEntries(ctx)
				},
			},
			{
				Name: "GetCfgPolicies",
				Method: func() (any, error) {
					return service.GetCfgPolicies(ctx)
				},
			},
			{
				Name: "GetCfgPolicyListEntries",
				Method: func() (any, error) {
					return service.GetCfgPolicyListEntries(ctx)
				},
			},
			{
				Name: "GetCfgWirelessAaaPolicyConfigs",
				Method: func() (any, error) {
					return service.GetCfgWirelessAaaPolicyConfigs(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetCfgByProfileName",
				Method: func() (any, error) {
					return service.GetCfgByProfileName(ctx, "test-wlan-profile")
				},
			},
			{
				Name: "GetCfgPoliciesByPolicyProfileName",
				Method: func() (any, error) {
					return service.GetCfgPoliciesByPolicyProfileName(ctx, "test-policy-profile")
				},
			},
			{
				Name: "GetCfgByID",
				Method: func() (any, error) {
					return service.GetCfgByID(ctx, 1)
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
				Name: "GetCfgEntries_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgEntries(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgPolicies_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgPolicies(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgPolicyListEntries_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgPolicyListEntries(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgWirelessAaaPolicyConfigs_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgWirelessAaaPolicyConfigs(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByProfileName_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByProfileName(ctx, "test-wlan-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgPoliciesByPolicyProfileName_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgPoliciesByPolicyProfileName(ctx, "test-policy-profile")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByID_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetCfgByID(ctx, 1)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetCfgByProfileName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetCfgByProfileName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByProfileName_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetCfgByProfileName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgPoliciesByPolicyProfileName_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetCfgPoliciesByPolicyProfileName(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgPoliciesByPolicyProfileName_WhitespaceParam",
				Function: func() (any, error) {
					_, err := service.GetCfgPoliciesByPolicyProfileName(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfgByID_NegativeParam",
				Function: func() (any, error) {
					_, err := service.GetCfgByID(ctx, -1)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
