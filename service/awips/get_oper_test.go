package awips

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// Test constants
const testAPMac = "aa:bb:cc:dd:ee:ff"

func Test_AWIPSGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	// Define test suite using the new standard pattern
	suite := framework.GetOperTestSuite{
		ServiceName: "AWIPS",
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
					return service.GetOperByApMac(ctx, testAPMac)
				},
			},
			{
				Name: "GetOperByApMacDownloadStatus",
				Method: func() (any, error) {
					return service.GetOperByApMacDownloadStatus(ctx, testAPMac)
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			// Nil client tests
			{
				Name: "GetOper_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMac_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByApMac(ctx, testAPMac)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacDownloadStatus_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByApMacDownloadStatus(ctx, testAPMac)
					return nil, err
				},
				ExpectedError: true,
			},
			// Parameter validation tests
			{
				Name: "GetOperByApMac_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMac(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMac_WhitespaceMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMac(ctx, "   ")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByApMacDownloadStatus_EmptyMAC",
				Function: func() (any, error) {
					_, err := service.GetOperByApMacDownloadStatus(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	}

	// Run the standardized test suite
	framework.RunStandardGetOperTests(t, suite)
}
