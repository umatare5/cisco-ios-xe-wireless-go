//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/awips"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestAWIPSServiceIntegration_GetOperationalOperations_Success validates AWIPS service
// operational data retrieval against live WNC controller.
func TestAWIPSServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "AWIPS",
			ServiceConstructor: func(client any) any {
				return awips.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(awips.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
