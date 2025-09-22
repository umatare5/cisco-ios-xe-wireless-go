//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/general"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestGeneralServiceIntegration_GetOperationalOperations_Success validates General service
// operational data retrieval against live WNC controller.
func TestGeneralServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "General",
			ServiceConstructor: func(client any) any {
				return general.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetManagementInterfaceState",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetManagementInterfaceState(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetAPLocationRangingConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetAPLocationRangingConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetCACConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetCACConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetFeatureUsageConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetFeatureUsageConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetFIPSConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetFIPSConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGeolocationConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetGeolocationConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetLAGInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetLAGInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMEWLCConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetMEWLCConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMFPConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetMFPConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMulticastConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetMulticastConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListSIML3InterfaceCache",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListSIML3InterfaceCache(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetThresholdWarningConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetThresholdWarningConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetWLCManagementInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetWLCManagementInfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetWSAAPClientEventConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetWSAAPClientEventConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgMewlcConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListCfgMewlcConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgMfp",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListCfgMfp(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgLaginfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListCfgLaginfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgMulticastConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListCfgMulticastConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListOperMgmtIntfData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListOperMgmtIntfData(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMewlcConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListMewlcConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMfp",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListMfp(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListLaginfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListLaginfo(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMulticastConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).ListMulticastConfig(ctx)
				},
				LogResult: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}

// TestGeneralServiceIntegration_GetConfigOperations_Success validates General service
// configuration data retrieval against live WNC controller.
func TestGeneralServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "General Config",
			ServiceConstructor: func(client any) any {
				return general.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(general.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
