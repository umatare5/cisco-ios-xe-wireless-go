package rrm

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RrmGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "RRM",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetGlobalOper",
				Method: func() (any, error) {
					return service.GetGlobalOper(ctx)
				},
			},
			{
				Name: "GetEmulOper",
				Method: func() (any, error) {
					return service.GetEmulOper(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByDeviceID",
				Method: func() (any, error) {
					return service.GetOperByDeviceID(ctx, "test-device")
				},
			},
			{
				Name: "GetOperByPhyType",
				Method: func() (any, error) {
					return service.GetOperByPhyType(ctx, "802.11n")
				},
			},
			{
				Name: "GetOperByWtpMacAndRadioSlot",
				Method: func() (any, error) {
					return service.GetOperByWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", 0)
				},
			},
			{
				Name: "GetGlobalOperBy5GWtpMacAndRadioSlot",
				Method: func() (any, error) {
					return service.GetGlobalOperBy5GWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", 1)
				},
			},
			{
				Name: "GetGlobalOperBy6GhzWtpMacAndRadioSlot",
				Method: func() (any, error) {
					return service.GetGlobalOperBy6GhzWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", 2)
				},
			},
			{
				Name: "GetGlobalOperByApMac",
				Method: func() (any, error) {
					return service.GetGlobalOperByApMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperByBandID",
				Method: func() (any, error) {
					return service.GetGlobalOperByBandID(ctx, "band-1")
				},
			},
			{
				Name: "GetGlobalOperByChannelPhyType",
				Method: func() (any, error) {
					return service.GetGlobalOperByChannelPhyType(ctx, "802.11ac")
				},
			},
			{
				Name: "GetGlobalOperByPhyType",
				Method: func() (any, error) {
					return service.GetGlobalOperByPhyType(ctx, "802.11ax")
				},
			},
			{
				Name: "GetGlobalOperByWtpMacAndRadioSlot",
				Method: func() (any, error) {
					return service.GetGlobalOperByWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", 0)
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
				Name: "GetOperByDeviceID_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByDeviceID(ctx, "test-device")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByDeviceID_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByDeviceID(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByPhyType_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByPhyType(ctx, "802.11n")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByPhyType_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByPhyType(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWtpMacAndRadioSlot_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", 0)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWtpMacAndRadioSlot_EmptyMac",
				Function: func() (any, error) {
					_, err := service.GetOperByWtpMacAndRadioSlot(ctx, "", 0)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByWtpMacAndRadioSlot_InvalidSlot",
				Function: func() (any, error) {
					_, err := service.GetOperByWtpMacAndRadioSlot(ctx, "aa:bb:cc:dd:ee:ff", -1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetGlobalOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOperByApMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetGlobalOperByApMac(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOperByApMac_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetGlobalOperByApMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetEmulOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetEmulOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
