//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestSiteServiceIntegration_GetConfigOperations_Success validates Site service
// configuration retrieval against live WNC controller.
func TestSiteServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Site",
			ServiceConstructor: func(client any) any {
				return site.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(site.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAPProfileConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(site.Service).ListAPProfileConfigs(ctx)
				},
			},
			{
				Name: "ListSiteTagConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(site.Service).ListSiteTagConfigs(ctx)
				},
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
