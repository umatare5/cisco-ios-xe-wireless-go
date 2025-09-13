package urwb_test

import (
	"context"
	"strings"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/urwb"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/urwb"
)

func TestUrwbServiceUnit_Constructor_Success(t *testing.T) {
	service := urwb.NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client service")
	}

	// Test with valid client
	mockServer := testutil.NewMockServer(map[string]string{
		"test": `{"data": {}}`,
	})
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service = urwb.NewService(client.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected service to have client, got nil")
	}
}

func TestUrwbServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	responses := map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": `{
			"urwb-profiles": {
				"urwb-profile": [
					{
						"profile-name": "test-profile",
						"descp": "Test profile",
						"enabled": true,
						"passphr": "testpassword123",
						"mpo": {
							"max-links": 1
						}
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": `{
			"urwbnet-stats": [],
			"urwbnet-node-g": []
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := urwb.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
		}
		if result != nil && len(result.UrwbProfile) == 0 {
			t.Error("GetConfig returned empty profiles")
		}
	})

	t.Run("ListProfiles", func(t *testing.T) {
		result, err := service.ListProfiles(ctx)
		if err != nil {
			t.Errorf("ListProfiles returned unexpected error: %v", err)
		}
		if len(result) == 0 {
			t.Error("ListProfiles returned empty result")
			return
		}
		if result[0].ProfileName != "test-profile" {
			t.Errorf("Expected profile name 'test-profile', got '%s'", result[0].ProfileName)
		}
	})

	t.Run("GetProfile", func(t *testing.T) {
		result, err := service.GetProfile(ctx, "test-profile")
		if err != nil {
			t.Errorf("GetProfile returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetProfile returned nil result")
		}
		if result != nil && result.ProfileName != "test-profile" {
			t.Errorf("Expected profile name 'test-profile', got '%s'", result.ProfileName)
		}
	})

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Log("GetOperational returned nil result (expected when not configured)")
		} else {
			t.Log("GetOperational returned data")
		}
	})

	t.Run("ListStats", func(t *testing.T) {
		result, err := service.ListStats(ctx)
		if err != nil {
			t.Errorf("ListStats returned unexpected error: %v", err)
		}
		if len(result) == 0 {
			t.Log("ListStats returned empty result (expected when not configured)")
			return
		}
		if result[0].Mac == "" {
			t.Error("Expected non-empty MAC address")
		}
	})

	t.Run("ListNodeGroups", func(t *testing.T) {
		result, err := service.ListNodeGroups(ctx)
		if err != nil {
			t.Errorf("ListNodeGroups returned unexpected error: %v", err)
		}
		if len(result) == 0 {
			t.Log("ListNodeGroups returned empty result (expected when not configured)")
		} else if result[0].GroupID == "" {
			t.Error("Expected non-empty group ID")
		}
	})
}

func TestUrwbServiceUnit_ValidationErrors(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(map[string]string{})
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := urwb.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetProfile_EmptyName", func(t *testing.T) {
		result, err := service.GetProfile(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty profile name")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("UpsertProfile_EmptyName", func(t *testing.T) {
		err := service.UpsertProfile(ctx, &model.UrwbProfile{ProfileName: ""})
		if err == nil {
			t.Error("Expected validation error for empty profile name")
		}
	})

	t.Run("DeleteProfile_EmptyName", func(t *testing.T) {
		err := service.DeleteProfile(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty profile name")
		}
	})
}

func TestUrwbServiceUnit_NotFoundHandling(t *testing.T) {
	t.Parallel()

	responses := map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
				"urwb-profiles": {
					"urwb-profile": [
						{
							"profile-name": "existing-profile",
							"coordinator-id": 1
						}
					]
				}
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := urwb.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetProfile_NotFound", func(t *testing.T) {
		result, err := service.GetProfile(ctx, "nonexistent-profile")
		if err == nil {
			t.Error("Expected error for non-existent profile")
		}
		if result != nil {
			t.Error("Expected nil result for non-existent profile")
		}
	})
}

func TestUrwbServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	service := urwb.NewService(nil)
	ctx := testutil.TestContext(t)

	// Test with nil client
	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error with nil client, got nil")
	}
}

// TestUrwbServiceUnit_GetConfig_RealDataSuccess tests GetConfig with real IOS-XE 17.18.1 data.
func TestUrwbServiceUnit_GetConfig_RealDataSuccess(t *testing.T) {
	// Real data from IOS-XE 17.18.1 Live WNC
	mockResponse := `{
		"urwb-profiles": {
			"urwb-profile": [
				{
					"profile-name": "custom-profile",
					"descp": "A custom profile",
					"enabled": true,
					"passphr": "QOWLDUVhLCEgYdhea_G[\\U_RPfILAXM\\MPAJAAB",
					"mpo": {
						"max-links": 1
					}
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetConfig(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if len(result.UrwbProfile) == 0 {
		t.Fatal("Expected at least one profile, got empty")
	}

	profile := result.UrwbProfile[0]
	if profile.ProfileName != "custom-profile" {
		t.Errorf("Expected profile name 'custom-profile', got '%s'", profile.ProfileName)
	}

	if !profile.Enabled {
		t.Error("Expected profile to be enabled")
	}

	if profile.Description != "A custom profile" {
		t.Errorf("Expected description 'A custom profile', got '%s'", profile.Description)
	}
}

// TestUrwbServiceUnit_GetProfile_RealDataSpecific tests GetProfile with specific profile data.
func TestUrwbServiceUnit_GetProfile_RealDataSpecific(t *testing.T) {
	// Real URWB profiles data from IOS-XE 17.18.1
	mockResponse := `{
		"urwb-profiles": {
			"urwb-profile": [
				{
					"profile-name": "custom-profile",
					"descp": "A custom profile",
					"enabled": true,
					"passphr": "QOWLDUVhLCEgYdhea_G[\\U_RPfILAXM\\MPAJAAB",
					"mpo": {
						"max-links": 1
					}
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetProfile(ctx, "custom-profile")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.ProfileName != "custom-profile" {
		t.Errorf("Expected profile name 'custom-profile', got '%s'", result.ProfileName)
	}

	if result.Description != "A custom profile" {
		t.Errorf("Expected description 'A custom profile', got '%s'", result.Description)
	}

	if !result.Enabled {
		t.Error("Expected profile to be enabled")
	}

	if result.Mpo == nil {
		t.Fatal("Expected MPO configuration, got nil")
	}

	if result.Mpo.MaxLinks != 1 {
		t.Errorf("Expected max links 1, got %d", result.Mpo.MaxLinks)
	}
}

// TestUrwbServiceUnit_UpsertProfile_RealDataCreateNew tests UpsertProfile for creating new profile.
func TestUrwbServiceUnit_UpsertProfile_RealDataCreateNew(t *testing.T) {
	// Empty initial response
	getResponse := `{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
			"urwb-profiles": {
				"urwb-profile": []
			}
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": getResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	newProfile := &model.UrwbProfile{
		ProfileName: "new-profile",
		Description: "Test profile for new creation",
		Enabled:     true,
	}

	err := service.UpsertProfile(ctx, newProfile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// TestUrwbServiceUnit_DeleteProfile_RealDataSuccess tests DeleteProfile with existing profile.
func TestUrwbServiceUnit_DeleteProfile_RealDataSuccess(t *testing.T) {
	getResponse := `{
		"urwb-profiles": {
			"urwb-profile": [
				{
					"profile-name": "custom-profile",
					"descp": "A custom profile",
					"enabled": true,
					"passphr": "QOWLDUVhLCEgYdhea_G[\\U_RPfILAXM\\MPAJAAB",
					"mpo": {
						"max-links": 1
					}
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": getResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	err := service.DeleteProfile(ctx, "custom-profile")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// TestUrwbServiceUnit_GetOperational_RealDataSuccess tests GetOperational with operational data.
func TestUrwbServiceUnit_GetOperational_RealDataSuccess(t *testing.T) {
	// Real data from IOS-XE 17.18.1 Live WNC
	// Empty response indicates URWB operational data is not configured
	mockResponse := `{}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Empty response is expected when URWB operational data is not configured
	if result == nil {
		t.Log("Note: Empty URWB operational data (expected when not configured)")
		return
	}

	// If result is not nil, validate the structure
	if len(result.UrwbnetStats) == 0 {
		t.Log("Note: No URWB stats available (expected when not configured)")
	}

	if len(result.UrwbnetNodeG) == 0 {
		t.Log("Note: No URWB node groups available (expected when not configured)")
	}
}

// TestUrwbServiceUnit_ErrorHandling_RealDataHTTPErrors tests various HTTP error scenarios.
func TestUrwbServiceUnit_ErrorHandling_RealDataHTTPErrors(t *testing.T) {
	tests := []struct {
		name      string
		operation func(service urwb.Service) error
	}{
		{
			name: "GetConfig_InternalServerError",
			operation: func(service urwb.Service) error {
				_, err := service.GetConfig(context.Background())
				return err
			},
		},
		{
			name: "ListProfiles_NotFound",
			operation: func(service urwb.Service) error {
				_, err := service.ListProfiles(context.Background())
				return err
			},
		},
		{
			name: "GetOperational_ServiceUnavailable",
			operation: func(service urwb.Service) error {
				_, err := service.GetOperational(context.Background())
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorPaths := []string{
				"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data",
				"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data",
			}

			mockServer := testutil.NewMockErrorServer(errorPaths, 500)
			defer mockServer.Close()

			testClient := mockServer.NewTestClient(t)
			client := testClient.Core().(*core.Client)
			service := urwb.NewService(client)

			err := tt.operation(service)
			if err == nil {
				t.Fatal("Expected error, got nil")
			}
		})
	}
}

// TestUrwbServiceUnit_SetConfig_Success tests SetConfig with profile data.
func TestUrwbServiceUnit_SetConfig_Success(t *testing.T) {
	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": "",
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	profiles := &model.UrwbProfiles{
		UrwbProfile: []model.UrwbProfile{
			{
				ProfileName: "test-profile",
				Description: "Test profile",
				Enabled:     true,
			},
		},
	}

	err := service.SetConfig(ctx, profiles)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// TestUrwbServiceUnit_SetConfig_HTTPErrors tests SetConfig error handling.
func TestUrwbServiceUnit_SetConfig_HTTPErrors(t *testing.T) {
	errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
	mockServer := testutil.NewMockErrorServer(errorPaths, 500)
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	profiles := &model.UrwbProfiles{UrwbProfile: []model.UrwbProfile{}}
	err := service.SetConfig(ctx, profiles)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// TestUrwbServiceUnit_UpsertProfile_UpdateExisting tests UpsertProfile for updating existing profile.
func TestUrwbServiceUnit_UpsertProfile_UpdateExisting(t *testing.T) {
	getResponse := `{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
			"urwb-profiles": {
				"urwb-profile": [
					{
						"profile-name": "existing-profile",
						"coordinator-id": 1,
						"node-id": 2
					}
				]
			}
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": getResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	updatedProfile := &model.UrwbProfile{
		ProfileName: "existing-profile",
		Description: "Updated description",
		Enabled:     true,
	}

	err := service.UpsertProfile(ctx, updatedProfile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// TestUrwbServiceUnit_UpsertProfile_ErrorHandling tests UpsertProfile error scenarios.
func TestUrwbServiceUnit_UpsertProfile_ErrorHandling(t *testing.T) {
	t.Run("GetProfilesError", func(t *testing.T) {
		errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
		mockServer := testutil.NewMockErrorServer(errorPaths, 500)
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		profile := &model.UrwbProfile{ProfileName: "test-profile"}
		err := service.UpsertProfile(ctx, profile)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("SetConfigError", func(t *testing.T) {
		// Use errorPaths approach for simpler error simulation
		errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
		errorServer := testutil.NewMockErrorServer(errorPaths, 500)
		defer errorServer.Close()

		testClient := errorServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		profile := &model.UrwbProfile{ProfileName: "test-profile"}
		err := service.UpsertProfile(ctx, profile)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

// TestUrwbServiceUnit_DeleteProfile_ErrorHandling tests DeleteProfile error scenarios.
func TestUrwbServiceUnit_DeleteProfile_ErrorHandling(t *testing.T) {
	t.Run("GetProfilesError", func(t *testing.T) {
		errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
		mockServer := testutil.NewMockErrorServer(errorPaths, 500)
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		err := service.DeleteProfile(ctx, "test-profile")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("ProfileNotFound", func(t *testing.T) {
		getResponse := `{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {"urwb-profiles": {"urwb-profile": []}}}`
		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": getResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		err := service.DeleteProfile(ctx, "nonexistent-profile")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("SetConfigError", func(t *testing.T) {
		// Use errorPaths approach for simpler error simulation
		errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
		errorServer := testutil.NewMockErrorServer(errorPaths, 500)
		defer errorServer.Close()

		testClient := errorServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		err := service.DeleteProfile(ctx, "profile-to-delete")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

// TestUrwbServiceUnit_ListStats_Success tests ListStats functionality.
func TestUrwbServiceUnit_ListStats_Success(t *testing.T) {
	mockResponse := `{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": {
			"urwbnet-stats": [
				{
					"mac": "aa:bb:cc:dd:ee:ff",
					"nodes-cnt": 15,
					"coordinator-id": 1
				},
				{
					"mac": "28:ac:9e:bb:3c:81",
					"nodes-cnt": 8,
					"coordinator-id": 2
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.ListStats(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Expected 2 stats, got %d", len(result))
	}

	if result[0].Mac != "aa:bb:cc:dd:ee:ff" {
		t.Errorf("Expected MAC 'aa:bb:cc:dd:ee:ff', got '%s'", result[0].Mac)
	}

	if result[0].NodesCount != 15 {
		t.Errorf("Expected nodes count 15, got %d", result[0].NodesCount)
	}
}

// TestUrwbServiceUnit_ListStats_HTTPErrors tests ListStats error handling.
func TestUrwbServiceUnit_ListStats_HTTPErrors(t *testing.T) {
	errorPaths := []string{"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data"}
	mockServer := testutil.NewMockErrorServer(errorPaths, 500)
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	_, err := service.ListStats(ctx)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// TestUrwbServiceUnit_ListNodeGroups_Success tests ListNodeGroups functionality.
func TestUrwbServiceUnit_ListNodeGroups_Success(t *testing.T) {
	mockResponse := `{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": {
			"urwbnet-node-g": [
				{
					"group-id": "group-001",
					"group-name": "primary-group",
					"status": "active"
				},
				{
					"group-id": "group-002",
					"group-name": "secondary-group",
					"status": "standby"
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.ListNodeGroups(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Expected 2 node groups, got %d", len(result))
	}

	if result[0].GroupID != "group-001" {
		t.Errorf("Expected group ID 'group-001', got '%s'", result[0].GroupID)
	}

	if result[0].GroupName != "primary-group" {
		t.Errorf("Expected group name 'primary-group', got '%s'", result[0].GroupName)
	}
}

// TestUrwbServiceUnit_ListNodeGroups_HTTPErrors tests ListNodeGroups error handling.
func TestUrwbServiceUnit_ListNodeGroups_HTTPErrors(t *testing.T) {
	errorPaths := []string{"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data"}
	mockServer := testutil.NewMockErrorServer(errorPaths, 500)
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	_, err := service.ListNodeGroups(ctx)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// TestUrwbServiceUnit_EmptyResponse_RealDataHandling tests handling of empty or minimal responses.
func TestUrwbServiceUnit_EmptyResponse_RealDataHandling(t *testing.T) {
	tests := []struct {
		name      string
		responses map[string]string
		operation func(service urwb.Service) (any, error)
		validate  func(t *testing.T, result any, err error)
	}{
		{
			name: "ListProfiles_EmptyResponse",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": `{
					"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {}
				}`,
			},
			operation: func(service urwb.Service) (any, error) {
				return service.ListProfiles(context.Background())
			},
			validate: func(t *testing.T, result any, err error) {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				profiles := result.([]model.UrwbProfile)
				if len(profiles) > 0 {
					t.Errorf("Expected empty profiles list, got %d profiles", len(profiles))
				}
			},
		},
		{
			name: "GetOperational_EmptyStats",
			responses: map[string]string{
				"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": `{
					"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": {
						"urwbnet-oper-data": {
							"urwbnet-stats": [],
							"urwbnet-node-g": []
						}
					}
				}`,
			},
			operation: func(service urwb.Service) (any, error) {
				return service.GetOperational(context.Background())
			},
			validate: func(t *testing.T, result any, err error) {
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				operData := result.(*model.UrwbnetOperData)
				if operData == nil {
					t.Fatal("Expected result, got nil")
				}
				if len(operData.UrwbnetStats) != 0 {
					t.Errorf("Expected empty stats, got %d items", len(operData.UrwbnetStats))
				}
				if len(operData.UrwbnetNodeG) != 0 {
					t.Errorf("Expected empty node groups, got %d items", len(operData.UrwbnetNodeG))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := testutil.NewMockServer(tt.responses)
			defer mockServer.Close()

			testClient := mockServer.NewTestClient(t)
			client := testClient.Core().(*core.Client)
			service := urwb.NewService(client)

			result, err := tt.operation(service)
			tt.validate(t, result, err)
		})
	}
}

// TestUrwbServiceUnit_GetConfig_NilHandling tests GetConfig with nil responses.
func TestUrwbServiceUnit_GetConfig_NilHandling(t *testing.T) {
	mockResponse := `{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetConfig(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
}

// TestUrwbServiceUnit_GetOperational_NilHandling tests GetOperational with nil responses.
func TestUrwbServiceUnit_GetOperational_NilHandling(t *testing.T) {
	mockResponse := `{}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected nil result for empty response, got non-nil")
	}
}

// TestUrwbServiceUnit_GetProfile_AdditionalErrorHandling tests GetProfile edge cases.
func TestUrwbServiceUnit_GetProfile_AdditionalErrorHandling(t *testing.T) {
	t.Run("EmptyProfileList", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
				"urwb-profiles": {
					"urwb-profile": []
				}
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		_, err := service.GetProfile(ctx, "non-existent")
		if err == nil {
			t.Fatal("Expected error for non-existent profile, got nil")
		}
	})

	t.Run("NilProfilesContainer", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		_, err := service.GetProfile(ctx, "any-profile")
		if err == nil {
			t.Fatal("Expected error for nil profiles, got nil")
		}
	})
}

// TestUrwbServiceUnit_UpsertProfile_AdditionalErrorHandling tests UpsertProfile edge cases.
func TestUrwbServiceUnit_UpsertProfile_AdditionalErrorHandling(t *testing.T) {
	t.Run("NilProfilesContainer", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		profile := &model.UrwbProfile{
			ProfileName: "new-profile",
			Description: "Test profile",
		}

		err := service.UpsertProfile(ctx, profile)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("EmptyProfileList", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
				"urwb-profiles": {
					"urwb-profile": []
				}
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		profile := &model.UrwbProfile{
			ProfileName: "new-profile",
			Description: "Test profile",
		}

		err := service.UpsertProfile(ctx, profile)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})
}

// TestUrwbServiceUnit_DeleteProfile_AdditionalErrorHandling tests DeleteProfile edge cases.
func TestUrwbServiceUnit_DeleteProfile_AdditionalErrorHandling(t *testing.T) {
	t.Run("NilProfilesContainer", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		err := service.DeleteProfile(ctx, "any-profile")
		if err == nil {
			t.Fatal("Expected error for nil profiles, got nil")
		}
	})

	t.Run("EmptyProfileList", func(t *testing.T) {
		mockResponse := `{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
				"urwb-profiles": {
					"urwb-profile": []
				}
			}
		}`

		mockServer := testutil.NewMockServer(map[string]string{
			"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
		})
		defer mockServer.Close()

		testClient := mockServer.NewTestClient(t)
		client := testClient.Core().(*core.Client)
		service := urwb.NewService(client)
		ctx := context.Background()

		err := service.DeleteProfile(ctx, "non-existent")
		if err == nil {
			t.Fatal("Expected error for non-existent profile, got nil")
		}
	})
}

// TestUrwbServiceUnit_ListProfiles_NilHandling tests ListProfiles with nil responses.
func TestUrwbServiceUnit_ListProfiles_NilHandling(t *testing.T) {
	mockResponse := `{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": {
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.ListProfiles(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
}

// TestUrwbServiceUnit_ListStats_NilHandling tests ListStats with nil responses.
func TestUrwbServiceUnit_ListStats_NilHandling(t *testing.T) {
	mockResponse := `{}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.ListStats(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected nil result for empty response, got non-nil")
	}
}

// TestUrwbServiceUnit_ListNodeGroups_NilHandling tests ListNodeGroups with nil responses.
func TestUrwbServiceUnit_ListNodeGroups_NilHandling(t *testing.T) {
	mockResponse := `{}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.ListNodeGroups(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result != nil {
		t.Error("Expected nil result for empty response, got non-nil")
	}
}

// TestUrwbServiceUnit_DeleteProfile_HTTPErrorHandling tests DeleteProfile HTTP errors.
func TestUrwbServiceUnit_DeleteProfile_HTTPErrorHandling(t *testing.T) {
	errorPaths := []string{"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"}
	mockServer := testutil.NewMockErrorServer(errorPaths, 500)
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	err := service.DeleteProfile(ctx, "test-profile")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// TestUrwbServiceUnit_GetProfile_ProfileNotFound tests GetProfile when profile doesn't exist.
func TestUrwbServiceUnit_GetProfile_ProfileNotFound(t *testing.T) {
	// Empty profiles data
	mockResponse := `{
		"urwb-profiles": {
			"urwb-profile": []
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	result, err := service.GetProfile(ctx, "nonexistent-profile")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if result != nil {
		t.Fatal("Expected nil result, got profile")
	}

	// Verify it's a NotFoundError
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Expected 'not found' error, got: %v", err)
	}
}

// TestUrwbServiceUnit_UpsertProfile_AddNewProfile tests UpsertProfile adding a new profile.
func TestUrwbServiceUnit_UpsertProfile_AddNewProfile(t *testing.T) {
	// Start with empty profiles
	emptyResponse := `{
		"urwb-profiles": {
			"urwb-profile": []
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": emptyResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	// Create a new profile
	newProfile := &model.UrwbProfile{
		ProfileName: "new-profile",
		Description: "New test profile",
		Enabled:     true,
	}

	err := service.UpsertProfile(ctx, newProfile)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// TestUrwbServiceUnit_DeleteProfile_ProfileNotFound tests DeleteProfile when profile doesn't exist.
func TestUrwbServiceUnit_DeleteProfile_ProfileNotFound(t *testing.T) {
	// Empty profiles data
	mockResponse := `{
		"urwb-profiles": {
			"urwb-profile": []
		}
	}`

	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data": mockResponse,
	})
	defer mockServer.Close()

	testClient := mockServer.NewTestClient(t)
	client := testClient.Core().(*core.Client)
	service := urwb.NewService(client)
	ctx := context.Background()

	err := service.DeleteProfile(ctx, "nonexistent-profile")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Verify it's a NotFoundError
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("Expected 'not found' error, got: %v", err)
	}
}
