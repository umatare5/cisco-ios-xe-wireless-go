package ap

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func TestService(t *testing.T) {
	testClient := client.OptionalTestClient(t)
	service := NewService(testClient)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "AP Service",
		BasicMethods: []framework.GetOperTestMethod{
			// Basic operational methods
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperCapwapDataAll",
				Method: func() (any, error) {
					return service.GetOperCapwapDataAll(ctx)
				},
			},
			{
				Name: "GetOperRadioStatusAll",
				Method: func() (any, error) {
					return service.GetOperRadioStatusAll(ctx)
				},
			},
			// Global operational methods
			{
				Name: "GetGlobalOper",
				Method: func() (any, error) {
					return service.GetGlobalOper(ctx)
				},
			},
			{
				Name: "GetGlobalOperEwlcApStats",
				Method: func() (any, error) {
					return service.GetGlobalOperEwlcApStats(ctx)
				},
			},
			// Configuration methods
			{
				Name: "GetCfg",
				Method: func() (any, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				Name: "GetApTagsCfg",
				Method: func() (any, error) {
					return service.GetApTagsCfg(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetGlobalOperApHistoryByEthernetMAC",
				Method: func() (any, error) {
					return service.GetGlobalOperApHistoryByEthernetMAC(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperApJoinStats",
				Method: func() (any, error) {
					return service.GetGlobalOperApJoinStats(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetGlobalOperWlanClientStatsByWlanID",
				Method: func() (any, error) {
					return service.GetGlobalOperWlanClientStatsByWlanID(ctx, 1)
				},
			},
			{
				Name: "GetGlobalOperApLocationStatsByLocation",
				Method: func() (any, error) {
					return service.GetGlobalOperApLocationStatsByLocation(ctx, "Building-A")
				},
			},
			{
				Name: "GetCfgApTagByMac",
				Method: func() (any, error) {
					return service.GetCfgApTagByMac(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
		},
		ParameterTests: []framework.ErrorFunctionValidationCase{
			// Nil client tests
			{
				Name: "GetOper with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetOper(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetCfg with nil client",
				Function: func() (any, error) {
					_, err := nilService.GetCfg(ctx)
					return nil, err
				},
				ExpectedError: true,
			},
		},
		CustomTests: func(t *testing.T) {
			// Administrative operations
			t.Run("AdministrativeOperations", func(t *testing.T) {
				testMAC := "aa:bb:cc:dd:ee:ff"

				t.Run("EnableAP", func(t *testing.T) {
					err := nilService.EnableAP(ctx, testMAC)
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})

				t.Run("DisableAP", func(t *testing.T) {
					err := nilService.DisableAP(ctx, testMAC)
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})
			})

			// Tag assignment operations
			t.Run("TagAssignmentOperations", func(t *testing.T) {
				testMAC := "aa:bb:cc:dd:ee:ff"

				t.Run("AssignSiteTag", func(t *testing.T) {
					err := nilService.AssignSiteTag(ctx, testMAC, "Site-A")
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})

				t.Run("AssignPolicyTag", func(t *testing.T) {
					err := nilService.AssignPolicyTag(ctx, testMAC, "Policy-A")
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})

				t.Run("AssignRFTag", func(t *testing.T) {
					err := nilService.AssignRFTag(ctx, testMAC, "RF-A")
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})
			})

			// Maintenance operations
			t.Run("MaintenanceOperations", func(t *testing.T) {
				testMAC := "aa:bb:cc:dd:ee:ff"

				t.Run("ReloadAP", func(t *testing.T) {
					err := nilService.Reload(ctx, testMAC)
					if err == nil {
						t.Error("Expected error with nil client, got nil")
					}
				})
			})
		},
	})
}
