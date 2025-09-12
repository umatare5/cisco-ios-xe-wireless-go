//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/wat"
)

// TestWATServiceIntegration_GetConfigOperations_Success validates WAT (Wireless Assurance Testing) service
// configuration retrieval against live WNC controller (IOS-XE 17.18.1+).
// Note: This service requires IOS-XE 17.18.1+ and uses MockErrorServer for unsupported versions.
func TestWATServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel() // Safe for parallel execution as read-only operations
	suite := client.IntegrationTestSuite{
		Config: client.TestSuiteConfig{
			ServiceName: "WAT",
			ServiceConstructor: func(client any) any {
				return wat.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []client.IntegrationTestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(wat.Service).GetConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
			{
				Name: "GetThousandeyesConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(wat.Service).GetThousandeyesConfig(ctx)
				},
				LogResult:      true,
				ExpectNotFound: true, // IOS-XE 17.18.1+ feature
			},
		},
		FilterMethods: []client.IntegrationTestMethod{
			{
				Name: "GetTestProfile",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(wat.Service).GetTestProfile(ctx, "test-profile")
				},
				ExpectNotFound: true, // Profile may not exist
			},
			{
				Name: "GetSchedule",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(wat.Service).GetSchedule(ctx, "test-schedule")
				},
				ExpectNotFound: true, // Schedule may not exist
			},
			{
				Name: "GetReportTemplate",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(wat.Service).GetReportTemplate(ctx, "test-report")
				},
				ExpectNotFound: true, // Report may not exist
			},
		},
		ValidationTests: []client.ValidationTestMethod{
			{
				Name: "GetTestProfile_EmptyName",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(wat.Service).GetTestProfile(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "empty", "profile"},
			},
			{
				Name: "GetSchedule_EmptyName",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(wat.Service).GetSchedule(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "empty", "schedule"},
			},
			{
				Name: "GetReportTemplate_EmptyName",
				Method: func(ctx context.Context, service any) error {
					_, err := service.(wat.Service).GetReportTemplate(ctx, "")
					return err
				},
				ExpectedError: true,
				ErrorKeywords: []string{"invalid", "empty", "report"},
			},
		},
	}

	client.RunIntegrationTestSuite(t, suite)
}
