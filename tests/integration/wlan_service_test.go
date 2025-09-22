//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestWLANServiceIntegration_GetConfigOperations_Success validates WLAN service
// configuration retrieval against live WNC controller.
func TestWLANServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "WLAN",
			ServiceConstructor: func(client any) any {
				service := wlan.NewService(client.(*core.Client))
				return &service
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListConfigEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListConfigEntries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListPolicies",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListPolicies(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListPolicyListEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListPolicyListEntries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListWirelessAAAPolicyConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListWirelessAAAPolicyConfigs(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListWlanCfgEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListWlanCfgEntries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListWlanPolicies",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListWlanPolicies(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgPolicyListEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListCfgPolicyListEntries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgWirelessAaaPolicyConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListCfgWirelessAaaPolicyConfigs(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListDot11beProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListDot11beProfiles(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}

// TestWLANServiceIntegration_GetOperationalOperations_Success validates WLAN service
// operational data retrieval against live WNC controller.
func TestWLANServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "WLAN",
			ServiceConstructor: func(client any) any {
				service := wlan.NewService(client.(*core.Client))
				return &service
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListWlanInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListWlanInfo(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
