//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wlan"
)

// TestWLANServiceIntegration_GetConfigOperations_Success validates WLAN service
// configuration retrieval against live WNC controller.
func TestWLANServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "WLAN",
			ServiceConstructor: func(client any) any {
				service := wlan.NewService(client.(*core.Client))
				return &service
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListProfileConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).ListProfileConfigs(ctx)
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
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}

// TestWLANServiceIntegration_GetOperationalOperations_Success validates WLAN service
// operational data retrieval against live WNC controller.
func TestWLANServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "WLAN",
			ServiceConstructor: func(client any) any {
				service := wlan.NewService(client.(*core.Client))
				return &service
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(*wlan.Service).GetOperational(ctx)
				},
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
