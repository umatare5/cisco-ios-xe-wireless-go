//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/apf"
)

// TestAPFServiceIntegration_GetConfigOperations_Success validates APF service
// configuration retrieval operations against live WNC controller.
//
// This test verifies that basic configuration operations return valid data
// structures and can communicate with the WNC API endpoint successfully.
func TestAPFServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	// Define the test suite configuration
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "APF Config",
			ServiceConstructor: func(client any) any {
				return apf.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(apf.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
		},
		// No filter methods for APF configuration - it's a simple get configuration
		FilterMethods: []client.IntegrationTestMethod{},
		// No specific validation tests for APF - it has simple configuration
		ValidationTests: []client.ValidationTestMethod{},
	}

	// Run the unified test suite
	client.RunIntegrationTestSuite(t, suite)
}
