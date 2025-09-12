//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/lisp"
)

// TestLISPServiceIntegration_GetOperationalOperations_Success validates LISP service
// operational data retrieval against live WNC controller.
func TestLISPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "LISP",
			ServiceConstructor: func(client any) any {
				return lisp.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // LISP may not be configured
			},
			{
				Name: "GetCapabilities",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).GetCapabilities(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
