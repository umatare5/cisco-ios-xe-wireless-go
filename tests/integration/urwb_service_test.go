//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/urwb"
)

// TestURWBServiceIntegration_GetConfigOperations_Success validates URWB service
// configuration retrieval against live WNC controller (IOS-XE 17.18.1+).
// Note: This service requires IOS-XE 17.18.1+ and uses MockErrorServer for unsupported versions.
func TestURWBServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "URWB",
			ServiceConstructor: func(client any) any {
				return urwb.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "ListProfiles",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).ListProfiles(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetProfile",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).GetProfile(ctx, "test-profile")
				},
				ExpectNotFound: true, // Profile may not exist
			},
		},
		ValidationTests: []client.ValidationTestMethod{
			{
				Name: "GetProfile_EmptyName",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(urwb.Service).GetProfile(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "empty", "profile"},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}

// TestURWBServiceIntegration_GetOperationalOperations_Success validates URWB service
// operational data retrieval against live WNC controller (IOS-XE 17.18.1+).
func TestURWBServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "URWB",
			ServiceConstructor: func(client any) any {
				return urwb.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).GetOperational(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "ListStats",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).ListStats(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "ListNodeGroups",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(urwb.Service).ListNodeGroups(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
		},
		FilterMethods:   []client.IntegrationTestMethod{},
		ValidationTests: []client.ValidationTestMethod{},
	}

	client.RunIntegrationTestSuite(t, suite)
}
