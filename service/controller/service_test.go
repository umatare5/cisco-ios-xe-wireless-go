package controller

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_ControllerService_StandardTests(t *testing.T) {
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Controller Service",
		ParameterTests: []framework.ErrorFunctionValidationCase{
			{
				Name: "Reload with nil client",
				Function: func() (any, error) {
					err := nilService.Reload(ctx, "Scheduled maintenance reload", false)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "ReloadWithReason with nil client",
				Function: func() (any, error) {
					err := nilService.ReloadWithReason(ctx, "System upgrade reload")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "Reload with empty reason",
				Function: func() (any, error) {
					err := nilService.Reload(ctx, "", false)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "ReloadWithReason with empty reason",
				Function: func() (any, error) {
					err := nilService.ReloadWithReason(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
		},
		CustomTests: func(t *testing.T) {
			t.Run("AdministrativeOperations", func(t *testing.T) {
				t.Log("Administrative operations (Reload, ReloadWithReason) tested in ParameterTests")
			})
		},
	})
}
