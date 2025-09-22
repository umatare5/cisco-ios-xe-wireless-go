//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rrm"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestRRMServiceIntegration_GetConfigOperations_Success validates RRM service
// configuration retrieval against live WNC controller.
func TestRRMServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RRM",
			ServiceConstructor: func(client any) any {
				return rrm.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetConfig(ctx)
				},
			},
			{
				Name: "ListRrms",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRrms(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRRMMgrCfgEntries",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRRMMgrCfgEntries(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}

// TestRRMServiceIntegration_GetOperationalOperations_Success validates RRM service
// operational data retrieval against live WNC controller.
func TestRRMServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RRM",
			ServiceConstructor: func(client any) any {
				return rrm.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetGlobalOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetEmulationOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).GetEmulationOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRRMOneShotCounters",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRRMOneShotCounters(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRadioOperData5g",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRadioOperData5g(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRRMChannelParams",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRRMChannelParams(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListRadioOperData24g",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rrm.Service).ListRadioOperData24g(ctx)
				},
				LogResult: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
