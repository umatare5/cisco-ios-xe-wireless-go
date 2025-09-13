//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/hyperlocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestHyperlocationServiceIntegration_GetOperationalOperations_Success validates Hyperlocation service
// operational data retrieval against live WNC controller.
func TestHyperlocationServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Hyperlocation",
			ServiceConstructor: func(client any) any {
				return hyperlocation.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(hyperlocation.Service).GetOperational(ctx)
				},
			},
			{
				Name: "ListProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(hyperlocation.Service).ListProfiles(ctx)
				},
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
