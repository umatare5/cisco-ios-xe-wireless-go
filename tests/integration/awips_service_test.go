//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/awips"
)

// TestAWIPSServiceIntegration_GetOperationalOperations_Success validates AWIPS service
// operational data retrieval against live WNC controller.
func TestAWIPSServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "AWIPS",
			ServiceConstructor: func(client any) any {
				return awips.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(awips.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
