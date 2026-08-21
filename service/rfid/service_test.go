package rfid

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfidServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
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
		"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-data-detail=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-data-detail": [
				{
					"rfid-mac-addr": "aa:bb:cc:dd:ee:ff"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data/rfid-data=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-rfid-oper:rfid-data": [
				{
					"rfid-mac-addr": "aa:bb:cc:dd:ee:ff"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data/rfid-radio-data=aa:bb:cc:dd:ee:ff,00:25:36:57:ed:cb,0": `{
			"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-radio-data": [
				{
					"rfid-mac-addr": "aa:bb:cc:dd:ee:ff",
					"ap-mac-addr": "00:25:36:57:ed:cb",
					"slot": 0
				}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

	t.Run("GetGlobalOperational", func(t *testing.T) {
		result, err := service.GetGlobalOperational(ctx)
		if err != nil {
			t.Errorf("GetGlobalOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalOperational returned nil result")
		}
	})

	// Test uncovered functions
	t.Run("GetGlobalDetailByMAC", func(t *testing.T) {
		result, err := service.GetGlobalDetailByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("GetGlobalDetailByMAC returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalDetailByMAC returned nil result")
		}
	})

	t.Run("GetRadioInfo", func(t *testing.T) {
		result, err := service.GetRadioInfo(ctx, "aa:bb:cc:dd:ee:ff", "00:25:36:57:ed:cb", 0)
		if err != nil {
			t.Errorf("GetRadioInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetRadioInfo returned nil result")
		}
	})

	t.Run("GetDetailByMAC", func(t *testing.T) {
		result, err := service.GetDetailByMAC(ctx, "aa:bb:cc:dd:ee:ff")
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
		"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data/rfid": `{
			"Cisco-IOS-XE-wireless-rfid-cfg:rfid": {
				"rfid-enable-cisco": false,
				"rfid-timeout": 120,
				"rfid-rssi-expiry": 60
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
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

// TestRfidServiceUnit_EnumSpelling_MockSuccess tests that the wire's enumeration spellings decode
// into the RFID enum types. The spellings come from the device's own schema declaration; the route
// answers 204 on every lab release, so no populated body is available to read them from.
func TestRfidServiceUnit_EnumSpelling_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data": `{
			"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data": {
				"rfid-data": [
					{
						"rfid-type": "cisco-rfid-data",
						"rfid-auto-interval": 60,
						"rfid-vendor": {
							"cisco": {"cisco-vendor-type": "rfid-type-aeroscout"}
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Fatalf("GetOperational failed: %v", err)
	}

	if result.CiscoIOSXEWirelessRFIDOperData == nil {
		t.Fatal("Expected rfid-oper-data to decode")
	}

	records := result.CiscoIOSXEWirelessRFIDOperData.RFIDData
	if len(records) != 1 {
		t.Fatalf("Expected 1 record, got %d", len(records))
	}
	if records[0].RFIDType != CiscoRFIDData {
		t.Errorf("Expected rfid-type %q, got %q", CiscoRFIDData, records[0].RFIDType)
	}
	if records[0].RFIDAutoInterval != 60 {
		t.Error("Expected a bare sibling of the enumeration to survive")
	}

	cisco := records[0].RFIDVendor.Cisco
	if cisco == nil || cisco.CiscoVendorType != RFIDTypeAeroscout {
		t.Errorf("Expected cisco-vendor-type to decode, got %+v", cisco)
	}
}
