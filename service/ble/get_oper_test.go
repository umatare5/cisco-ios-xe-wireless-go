package ble

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_BLEGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "BLE",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByApMac",
				Method: func() (any, error) {
					return service.GetOperByApMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetOperByApMacSlotAntenna",
				Method: func() (any, error) {
					return service.GetOperByApMacSlotAntenna(ctx, "aa:bb:cc:dd:ee:ff", 0, 1)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMac_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByApMac(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacSlotAntenna_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacSlotAntenna(ctx, "", 0, 1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacSlotAntenna_NegativeSlot",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacSlotAntenna(ctx, "aa:bb:cc:dd:ee:ff", -1, 1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacSlotAntenna_NegativeAntenna",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacSlotAntenna(ctx, "aa:bb:cc:dd:ee:ff", 0, -1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacSlotAntenna_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByApMacSlotAntenna(ctx, "aa:bb:cc:dd:ee:ff", 0, 1)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
