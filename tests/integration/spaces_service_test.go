//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/spaces"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestSpacesServiceIntegration_GetOperationalOperations_Success validates Cisco Spaces service
// operational data retrieval against live WNC controller (IOS-XE 17.18.1+).
// Note: This service requires IOS-XE 17.18.1+ and uses MockErrorServer for unsupported versions.
func TestSpacesServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "Spaces",
			ServiceConstructor: func(client any) any {
				return spaces.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(spaces.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "GetConnectionDetails",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(spaces.Service).GetConnectionDetails(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "GetTenantInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(spaces.Service).GetTenantInfo(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "GetConnectionStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(spaces.Service).GetConnectionStats(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
