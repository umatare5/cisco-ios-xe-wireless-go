package lisp

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_LispGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	// Define test suite using the new standard pattern
	suite := framework.GetOperTestSuite{
		ServiceName: "LISP",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperMemoryStats",
				Method: func() (any, error) {
					return service.GetOperMemoryStats(ctx)
				},
			},
			{
				Name: "GetOperCapabilities",
				Method: func() (any, error) {
					return service.GetOperCapabilities(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			// LISP service only has basic methods, no filter methods
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
				Name: "GetOperMemoryStats_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperMemoryStats(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperCapabilities_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperCapabilities(ctx)
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
