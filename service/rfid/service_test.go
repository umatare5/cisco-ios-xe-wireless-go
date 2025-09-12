package rfid

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfidServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(map[string]string{})
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := NewService(testClient.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected valid client, got nil")
	}
}

func TestRfidServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on YANG model structure for RFID
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data": `{
			"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data": {
				"rfid": {
					"enable": false
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data": `{
			"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data": {
				"rfid-summary": {
					"total-tags": 0,
					"active-tags": 0
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data": `{
			"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data": {
				"rfid-global-info": {
					"enabled": false,
					"total-aps": 0
				}
			}
		}`,
		// Add mock responses for MAC-based queries
		"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-data-detail=28:ac:9e:bb:3c:80": `{
			"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-emltd-data": {
				"mac-address": "28:ac:9e:bb:3c:80",
				"location": "Building A"
			}
		}`,
		"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data/rfid-data=28:ac:9e:bb:3c:80": `{
			"Cisco-IOS-XE-wireless-rfid-oper:rfid-data": {
				"mac-address": "28:ac:9e:bb:3c:80",
				"status": "active"
			}
		}`,
		"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-radio-data=28:ac:9e:bb:3c:80,00:25:36:57:ed:cb,0": `{
			"Cisco-IOS-XE-wireless-rfid-oper:rfid-radio-data": {
				"mac-address": "28:ac:9e:bb:3c:80",
				"ap-mac-address": "00:25:36:57:ed:cb",
				"slot": 0
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
		}
	})

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
		}
	})

	t.Run("GetGlobalInfo", func(t *testing.T) {
		result, err := service.GetGlobalInfo(ctx)
		if err != nil {
			t.Errorf("GetGlobalInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalInfo returned nil result")
		}
	})

	// Test uncovered functions
	t.Run("GetGlobalDetailByMAC", func(t *testing.T) {
		result, err := service.GetGlobalDetailByMAC(ctx, "28:ac:9e:bb:3c:80")
		if err != nil {
			t.Errorf("GetGlobalDetailByMAC returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalDetailByMAC returned nil result")
		}
	})

	t.Run("GetRadioInfo", func(t *testing.T) {
		result, err := service.GetRadioInfo(ctx, "28:ac:9e:bb:3c:80", "00:25:36:57:ed:cb", 0)
		if err != nil {
			t.Errorf("GetRadioInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetRadioInfo returned nil result")
		}
	})

	t.Run("GetDetailByMAC", func(t *testing.T) {
		result, err := service.GetDetailByMAC(ctx, "28:ac:9e:bb:3c:80")
		if err != nil {
			t.Errorf("GetDetailByMAC returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetDetailByMAC returned nil result")
		}
	})
}

func TestRfidServiceUnit_GetConfigSettings_MockSuccess(t *testing.T) {
	t.Parallel()

	responses := map[string]string{
		"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data/rfid-config": `{
			"Cisco-IOS-XE-wireless-rfid-cfg:rfid-config": {
				"enable": false,
				"timeout": 120,
				"interval": 60
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfigSettings", func(t *testing.T) {
		result, err := service.GetConfigSettings(ctx)
		if err != nil {
			t.Errorf("GetConfigSettings returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfigSettings returned nil result")
		}
	})
}

func TestRfidServiceUnit_ValidationErrors_InvalidInputs(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(map[string]string{})
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetGlobalDetailByMAC_InvalidMAC", func(t *testing.T) {
		result, err := service.GetGlobalDetailByMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected validation error for invalid MAC address")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("GetDetailByMAC_EmptyMAC", func(t *testing.T) {
		result, err := service.GetDetailByMAC(ctx, "")
		if err == nil {
			t.Error("Expected validation error for empty MAC address")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("GetRadioInfo_InvalidMAC", func(t *testing.T) {
		result, err := service.GetRadioInfo(ctx, "invalid", "11:22:33:44:55:66", 0)
		if err == nil {
			t.Error("Expected validation error for invalid MAC address")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})

	t.Run("GetRadioInfo_InvalidAPMAC", func(t *testing.T) {
		result, err := service.GetRadioInfo(ctx, "aa:bb:cc:dd:ee:ff", "invalid", 0)
		if err == nil {
			t.Error("Expected validation error for invalid AP MAC address")
		}
		if result != nil {
			t.Error("Expected nil result for invalid input")
		}
	})
}

func TestRfidServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		service := NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
