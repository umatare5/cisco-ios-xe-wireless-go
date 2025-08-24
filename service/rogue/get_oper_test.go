package rogue

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

func Test_RogueGetOper_StandardTests(t *testing.T) {
	clientSetup := client.OptionalTestClient(t)
	service := NewService(clientSetup)
	nilService := NewService(nil)
	ctx := client.TestContext(t)

	framework.RunStandardGetOperTests(t, framework.GetOperTestSuite{
		ServiceName: "Rogue",
		BasicMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOper",
				Method: func() (any, error) {
					return service.GetOper(ctx)
				},
			},
			{
				Name: "GetOperClientData",
				Method: func() (any, error) {
					return service.GetOperClientData(ctx)
				},
			},
			{
				Name: "GetOperData",
				Method: func() (any, error) {
					return service.GetOperData(ctx)
				},
			},
			{
				Name: "GetOperRldpStats",
				Method: func() (any, error) {
					return service.GetOperRldpStats(ctx)
				},
			},
			{
				Name: "GetOperStats",
				Method: func() (any, error) {
					return service.GetOperStats(ctx)
				},
			},
		},
		FilterMethods: []framework.GetOperTestMethod{
			{
				Name: "GetOperByRogueAddress",
				Method: func() (any, error) {
					return service.GetOperByRogueAddress(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetOperByRogueClientAddress",
				Method: func() (any, error) {
					return service.GetOperByRogueClientAddress(ctx, "aa:bb:cc:dd:ee:ff")
				},
			},
			{
				Name: "GetOperByClassType",
				Method: func() (any, error) {
					return service.GetOperByClassType(ctx, "malicious")
				},
			},
			{
				Name: "GetOperByContainmentLevel",
				Method: func() (any, error) {
					return service.GetOperByContainmentLevel(ctx, 1)
				},
			},
			{
				Name: "GetOperWithFields",
				Method: func() (any, error) {
					return service.GetOperWithFields(ctx, []string{"rogue-data", "stats"})
				},
			},
			{
				Name: "GetOperStatsWithFields",
				Method: func() (any, error) {
					return service.GetOperStatsWithFields(ctx, []string{"total-rogues"})
				},
			},
			{
				Name: "GetOperDataWithFields",
				Method: func() (any, error) {
					return service.GetOperDataWithFields(ctx, []string{"mac-address"})
				},
			},
			{
				Name: "GetOperClientDataWithFields",
				Method: func() (any, error) {
					return service.GetOperClientDataWithFields(ctx, []string{"client-mac"})
				},
			},
			{
				Name: "GetOperRldpStatsWithFields",
				Method: func() (any, error) {
					return service.GetOperRldpStatsWithFields(ctx, []string{"rldp-count"})
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
				Name: "GetOperByRogueAddress_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByRogueAddress(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByRogueAddress_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByRogueAddress(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByRogueClientAddress_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperByRogueClientAddress(ctx, "aa:bb:cc:dd:ee:ff")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByRogueClientAddress_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByRogueClientAddress(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByClassType_EmptyParam",
				Function: func() (any, error) {
					_, err := service.GetOperByClassType(ctx, "")
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperByContainmentLevel_InvalidLevel",
				Function: func() (any, error) {
					_, err := service.GetOperByContainmentLevel(ctx, -1)
					return nil, err
				},
				ExpectedError: true,
			},
			{
				Name: "GetOperWithFields_NilClient",
				Function: func() (any, error) {
					_, err := nilService.GetOperWithFields(ctx, []string{"rogue-data"})
					return nil, err
				},
				ExpectedError: true,
			},
		},
		TestConstants: framework.StandardGetOperTestConstants(),
	})
}
