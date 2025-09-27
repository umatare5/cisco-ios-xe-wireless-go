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
			{
				Name: "ListRFTags",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).ListRFTags(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRFProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).ListRFProfiles(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListMultiBssidProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).ListMultiBssidProfiles(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListAtfPolicies",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).ListAtfPolicies(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRFProfileDefaultEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).ListRFProfileDefaultEntries(ctx)
				},
				LogResult: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}

// TestRFServiceIntegration_GetOperationalOperations_Success validates RF service
// operational data retrieval against live WNC controller.
func TestRFServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RF Operational",
			ServiceConstructor: func(client any) any {
				return rf.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetAutoRFDot11Data",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).GetAutoRFDot11Data(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetRadarDetectionData",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rf.Service).GetRadarDetectionData(ctx)
				},
				LogResult: true,
			},
		},
	}

	integration.RunTestSuite(t, suite)
}
