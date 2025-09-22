//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot11"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestDot11ServiceIntegration_GetConfigOperations_Success validates 802.11 service
// configuration retrieval against live WNC controller.
func TestDot11ServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "DOT11",
			ServiceConstructor: func(client any) any {
				return dot11.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgConfiguredCountries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListCfgConfiguredCountries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgDot11Entries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListCfgDot11Entries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListCfgDot11acMcsEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListCfgDot11acMcsEntries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListConfiguredCountries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListConfiguredCountries(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListDot11Entries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListDot11Entries(ctx)
				},
				LogResult: true,
			},

			{
				Name: "ListCfgFilters",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot11.Service).ListCfgFilters(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
