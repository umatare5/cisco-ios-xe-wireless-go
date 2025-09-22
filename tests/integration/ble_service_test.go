//go:build integration

package integration_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ble"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

// TestBLEServiceIntegration_GetOperationalOperations_Success validates BLE service
// operational data retrieval against live WNC controller.
func TestBLEServiceIntegration_GetOperationalOperations_Success(t *testing.T) {
	t.Parallel()

	// Define the test suite configuration
	suite := integration.TestSuite{
		Config: integration.TestSuiteConfig{
			ServiceName: "BLE",
			ServiceConstructor: func(client any) any {
				return ble.NewService(client.(*core.Client))
			},
			UseTimeout: true,
		},
		BasicMethods: []integration.TestMethod{
			{
				Name: "GetOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).GetOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListBLELtxAp",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).ListBLELtxAp(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListBLELtxApAntenna",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).ListBLELtxApAntenna(ctx)
				},
				LogResult: true,
			},
			{
				Name: "GetMgmtOperational",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).GetMgmtOperational(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListBLEMgmtAp",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).ListBLEMgmtAp(ctx)
				},
				LogResult: true,
			},
			{
				Name: "ListBLEMgmtCmx",
				Method: func(ctx context.Context, service any) (any, error) {
					return service.(ble.Service).ListBLEMgmtCmx(ctx)
				},
				LogResult: true,
			},
		},
		FilterMethods:   []integration.TestMethod{},
		ValidationTests: []integration.ValidationTestMethod{},
	}

	integration.RunTestSuite(t, suite)
}
