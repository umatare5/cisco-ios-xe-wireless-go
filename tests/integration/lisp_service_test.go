//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/lisp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestLISPServiceIntegration_GetOperationalOperations_Success validates LISP service
// operational data retrieval against live WNC controller.
func TestLISPServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "LISP",
			ServiceConstructor: func(client any) any {
				return lisp.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // Not Verified on IOS-XE 17.12.5
			},
			{
				Name: "GetMemoryStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).GetMemoryStats(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "GetCapabilities",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).GetCapabilities(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
			{
				Name: "ListAPCapabilities",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(lisp.Service).ListAPCapabilities(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}
