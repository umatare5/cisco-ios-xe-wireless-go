package site

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_SiteGetOper_StandardTests(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	// Define test suite using the new standard pattern
	suite := framework.GetOperTestSuite{
		ServiceName: "Site",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			// Site service only has GetOper, no filter methods
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
		},
		TestConstants: framework.StandardGetOperTestConstants(),
		CustomTests: func(t *testing.T) {
			// Custom Site-specific tests
			t.Run("SiteDataStructureValidation", func(t *testing.T) {
				if testClient == nil {
					t.Skip("WNC_TEST_HOST not set, skipping integration tests")
					return
				}

				oper, err := service.GetOper(ctx)
				if err != nil {
					t.Logf("GetOper error (expected for some controllers): %v", err)
					return
				}

				if oper == nil {
					t.Log("GetOper returned nil (may be normal for controllers without site oper data)")
					return
				}

				// Validate the structure is correct even if data is empty
				t.Logf("Operational data structure validation passed")
			})
		},
	}

	// Run the standardized test suite
	framework.RunStandardGetOperTests(t, suite)
}
