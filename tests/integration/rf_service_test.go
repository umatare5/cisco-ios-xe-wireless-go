//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestRFServiceIntegration_GetConfigOperations_Success validates RF service
// configuration retrieval against live WNC controller.
func TestRFServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RF",
			ServiceConstructor: func(client any) any {
				return rf.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}
