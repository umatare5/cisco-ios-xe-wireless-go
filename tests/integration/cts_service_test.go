//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/cts"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestCTSServiceIntegration_GetOperationalOperations_Success validates CTS service
// configuration retrieval against live WNC controller.
func TestCTSServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "CTS",
			ServiceConstructor: func(client any) any {
				return cts.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(cts.Service).GetConfig(ctx)
				},
				LogResult: true,
			},

			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(cts.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "ListFlexModeApSxpConnectionStatus",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(cts.Service).ListFlexModeApSxpConnectionStatus(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
