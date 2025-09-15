//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rfid"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestRFIDServiceIntegration_GetConfigOperations_Success validates RFID service
// configuration retrieval against live WNC controller.
func TestRFIDServiceIntegration_GetConfigOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RFID",
			ServiceConstructor: func(client any) any {
				return rfid.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetConfig",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetConfig(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetConfigSettings",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetConfigSettings(ctx)
				},
				ExpectNotFound: true, // RFID config endpoint may not be available
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}

// TestRFIDServiceIntegration_GetOperationalOperations_Success validates RFID service
// operational data retrieval against live WNC controller.
func TestRFIDServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "RFID",
			ServiceConstructor: func(client any) any {
				return rfid.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetGlobalInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetGlobalInfo(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods: []integration.TestMethod{
			{
				Name: "GetRadioInfo",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetRadioInfo(ctx, integration.TestAPMac(), integration.TestAPMac(), 0)
				},
				ExpectNotFound: true,
			},
			{
				Name: "GetDetailByMAC",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(rfid.Service).GetDetailByMAC(ctx, integration.TestAPMac())
				},
				ExpectNotFound: true,
			},
		},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
