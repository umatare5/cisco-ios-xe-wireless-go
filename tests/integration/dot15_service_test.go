//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot15"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestDot15ServiceIntegration_GetConfigOperations_Success validates 802.15.4 service
// configuration retrieval against live WNC controller.
func TestDot15ServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "DOT15",
			ServiceConstructor: func(client any) any {
				return dot15.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot15.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
			{
				Name: "ListDot15GlobalConfigs",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(dot15.Service).ListDot15GlobalConfigs(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
		},
	}

	integration.RunTestSuite(t, suite)
}
