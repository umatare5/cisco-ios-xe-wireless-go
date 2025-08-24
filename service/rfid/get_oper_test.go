package rfid

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RfidGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "RFID",
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
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByMac",
				Method: func() (any, error) {
					return service.GetOperByMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetOperRfidData",
				Method: func() (any, error) {
					return service.GetOperRfidData(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperByMac",
				Method: func() (any, error) {
					return service.GetGlobalOperByMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperByRadioKey",
				Method: func() (any, error) {
					return service.GetGlobalOperByRadioKey(ctx, "aa:bb:cc:dd:ee:ff", "bb:cc:dd:ee:ff:aa", 0)
				},
			},
			{
				Name: "GetGlobalOperRfidDataDetail",
				Method: func() (any, error) {
					return service.GetGlobalOperRfidDataDetail(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperRfidRadioData",
				Method: func() (any, error) {
					return service.GetGlobalOperRfidRadioData(ctx, "aa:bb:cc:dd:ee:ff", "bb:cc:dd:ee:ff:aa", 0)
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
				Name: "GetOperByMac_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByMac_InvalidMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByMac(ctx, "invalid-mac")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperRfidData_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperRfidData(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperRfidData_InvalidMAC",
				Function: func() (any, error) {
					_, err := service.GetOperRfidData(ctx, "invalid-mac")
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
				Name: "GetGlobalOperByMac_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetGlobalOperByMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOperByMac_InvalidMAC",
				Function: func() (any, error) {
					_, err := service.GetGlobalOperByMac(ctx, "invalid-mac")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOperByRadioKey_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetGlobalOperByRadioKey(ctx, "", "bb:cc:dd:ee:ff:aa", 0)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetGlobalOperByRadioKey_EmptyAPMAC",
				Function: func() (any, error) {
					_, err := service.GetGlobalOperByRadioKey(ctx, "aa:bb:cc:dd:ee:ff", "", 0)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
