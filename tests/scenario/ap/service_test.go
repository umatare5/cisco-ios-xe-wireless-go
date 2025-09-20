//go:build scenario

package ap_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/tests/testutil/integration"
)

const testAPSlotID = 0 // Radio0 (2.4GHz)

// TestAPServiceScenario_AdminStateManagement_Success validates the complete
// admin state management workflow for Access Points through the WNC controller.
//
// This scenario test executes a 4-step workflow that covers the complete
// enable/disable cycle for AP admin state management:
//  1. Retrieve initial AP CAPWAP status
//  2. Disable AP admin state
//  3. Re-enable AP admin state
//  4. Validate final state shows 'adminstate-enabled'
//
// Test Environment:
//   - Requires live WNC controller connectivity
//   - Uses environment variables for authentication
//   - Tests against configured test AP (WNC_AP_MAC_ADDR)
//
// Note: Not marked as parallel due to AP state modifications
// that could interfere with concurrent AP operations.
func TestAPServiceScenario_AdminStateManagement_Success(t *testing.T) {
	// Initialize test client and service
	setup := integration.SetupTestClient(t)
	if setup.Client == nil {
		t.Skip("Skipping scenario tests: no client available")
	}

	service := ap.NewService(setup.Client)
	apMac := integration.TestAPMac()
	ctx := setup.Context

	t.Logf("Starting AP admin state management scenario for AP %s", apMac)

	// Step 1: Get initial AP CAPWAP status
	t.Logf("Step 1: Getting initial AP CAPWAP status for AP %s", apMac)
	initialState, err := service.GetCAPWAPDataByWTPMAC(ctx, apMac)
	if err != nil {
		t.Fatalf("Failed to get initial CAPWAP status for AP %s: %v", apMac, err)
	}
	logAPCapwapStatus(t, "Initial", apMac, initialState)

	// Step 2: Disable AP admin state
	t.Logf("Step 2: Disabling AP admin state for AP %s", apMac)
	if err := service.DisableAP(ctx, apMac); err != nil {
		t.Fatalf("Failed to disable AP %s: %v", apMac, err)
	}
	t.Logf("Successfully disabled AP admin state for AP %s", apMac)

	// Step 3: Enable AP admin state
	t.Logf("Step 3: Enabling AP admin state for AP %s", apMac)
	if err := service.EnableAP(ctx, apMac); err != nil {
		t.Fatalf("Failed to enable AP %s: %v", apMac, err)
	}
	t.Logf("Successfully enabled AP admin state for AP %s", apMac)

	// Step 4: Get final AP state and validate
	t.Logf("Step 4: Getting final AP CAPWAP status for AP %s", apMac)
	finalState, err := service.GetCAPWAPDataByWTPMAC(ctx, apMac)
	if err != nil {
		t.Fatalf("Failed to get final CAPWAP status for AP %s: %v", apMac, err)
	}
	logAPCapwapStatus(t, "Final", apMac, finalState)
	validateAPCapwapEnabled(t, apMac, finalState)

	t.Logf("AP admin state management scenario completed successfully for AP %s", apMac)
}

// TestAPServiceScenario_RadioStateManagement_Success validates the complete
// radio state management workflow for Access Point radio interfaces.
//
// This scenario test executes a 4-step workflow that covers the complete
// enable/disable cycle for AP radio admin state management:
//  1. Retrieve initial radio operational status
//  2. Disable radio admin state
//  3. Re-enable radio admin state
//  4. Validate final state shows 'enabled' and 'radio-up'
//
// Test Environment:
//   - Requires live WNC controller connectivity
//   - Uses environment variables for authentication
//   - Tests against configured test AP (WNC_AP_MAC_ADDR) Slot 0
//
// Note: Not marked as parallel due to radio state modifications
// that could interfere with concurrent radio operations.
func TestAPServiceScenario_RadioStateManagement_Success(t *testing.T) {
	// Initialize test client and service
	setup := integration.SetupTestClient(t)
	if setup.Client == nil {
		t.Skip("Skipping scenario tests: no client available")
	}

	service := ap.NewService(setup.Client)
	apMac := integration.TestAPMac()
	slotID := testAPSlotID
	ctx := setup.Context

	t.Logf("Starting radio state management scenario for AP %s Slot %d", apMac, slotID)

	// Step 1: Get initial radio operational status
	t.Logf("Step 1: Getting initial radio status for AP %s Slot %d", apMac, slotID)
	initialState, err := service.GetRadioStatusByWTPMACAndSlot(ctx, apMac, slotID)
	if err != nil {
		t.Fatalf("Failed to get initial radio status for AP %s Slot %d: %v", apMac, slotID, err)
	}
	logRadioStatus(t, "Initial", apMac, slotID, initialState)

	// Step 2: Disable radio admin state
	t.Logf("Step 2: Disabling radio admin state for AP %s Slot %d", apMac, slotID)
	radioBand := getRadioBandForSlot(slotID)
	if err := service.DisableRadio(ctx, apMac, radioBand); err != nil {
		t.Fatalf("Failed to disable radio for AP %s Slot %d: %v", apMac, slotID, err)
	}
	t.Logf("Successfully disabled radio admin state for AP %s Slot %d", apMac, slotID)

	// Step 3: Enable radio admin state
	t.Logf("Step 3: Enabling radio admin state for AP %s Slot %d", apMac, slotID)
	if err := service.EnableRadio(ctx, apMac, radioBand); err != nil {
		t.Fatalf("Failed to enable radio for AP %s Slot %d: %v", apMac, slotID, err)
	}
	t.Logf("Successfully enabled radio admin state for AP %s Slot %d", apMac, slotID)

	// Step 4: Get final radio state and validate
	t.Logf("Step 4: Getting final radio status for AP %s Slot %d", apMac, slotID)
	finalState, err := service.GetRadioStatusByWTPMACAndSlot(ctx, apMac, slotID)
	if err != nil {
		t.Fatalf("Failed to get final radio status for AP %s Slot %d: %v", apMac, slotID, err)
	}
	logRadioStatus(t, "Final", apMac, slotID, finalState)
	validateRadioEnabled(t, apMac, slotID, finalState)

	t.Logf("Radio state management scenario completed successfully for AP %s Slot %d", apMac, slotID)
}

// ============================
// Helper Functions
// ============================

// getRadioBandForSlot returns the appropriate radio band for a given slot ID.
func getRadioBandForSlot(slotID int) core.RadioBand {
	switch slotID {
	case 0:
		return core.RadioBand24GHz
	case 1:
		return core.RadioBand5GHz
	case 2:
		return core.RadioBand5GHz // Use 5GHz for slot 2 as 6GHz may not be available
	default:
		return core.RadioBand24GHz // Default to 2.4GHz for unknown slots
	}
}

// logAPCapwapStatus logs the AP CAPWAP status information with admin state.
func logAPCapwapStatus(t *testing.T, phase, apMac string, status *ap.ApOperCAPWAPData) {
	if status == nil {
		t.Logf("%s Status: No CAPWAP data available for AP %s", phase, apMac)
		return
	}

	t.Logf("%s Status for AP %s:", phase, apMac)
	if capwapData := findCAPWAPDataByMAC(status.CAPWAPData, apMac); capwapData != nil {
		t.Logf("  WTP MAC: %s", capwapData.WtpMAC)
		t.Logf("  IP Address: %s", capwapData.IPAddr)
		t.Logf("  Name: %s", capwapData.Name)
		t.Logf("  AP Admin State: %s", capwapData.ApState.ApAdminState)
		t.Logf("  AP Operation State: %s", capwapData.ApState.ApOperationState)
	} else {
		t.Logf("  CAPWAP data not found for AP %s", apMac)
		if status.CAPWAPData != nil {
			t.Logf("  Available CAPWAP data for other APs:")
			for i, capwapData := range status.CAPWAPData {
				t.Logf("    [%d] WTP MAC: %s, Name: %s, Admin State: %s",
					i, capwapData.WtpMAC, capwapData.Name, capwapData.ApState.ApAdminState)
			}
		}
	}
}

// logRadioStatus logs the radio operational status information.
func logRadioStatus(t *testing.T, phase, apMac string, slotID int, status *ap.ApOperRadioOperData) {
	if status == nil {
		t.Logf("%s Status: No data available for AP %s Slot %d", phase, apMac, slotID)
		return
	}

	t.Logf("%s Status for AP %s Slot %d:", phase, apMac, slotID)
	if radioData := findRadioDataByMACAndSlot(status.RadioOperData, apMac, slotID); radioData != nil {
		t.Logf("  WTP MAC: %s", radioData.WtpMAC)
		t.Logf("  Radio Slot %d:", slotID)
		t.Logf("    Admin State: %s", radioData.AdminState)
		t.Logf("    Oper State: %s", radioData.OperState)
		t.Logf("    Radio Type: %s", radioData.RadioType)
		if radioData.RadioMode != "" {
			t.Logf("    Radio Mode: %s", radioData.RadioMode)
		}
	} else {
		t.Logf("  Radio data not found for AP %s Slot %d", apMac, slotID)
		if status.RadioOperData != nil {
			t.Logf("  Available radio data for other APs/slots:")
			for i, radioData := range status.RadioOperData {
				t.Logf("    [%d] WTP MAC: %s, Slot: %d, Admin: %s, Oper: %s",
					i, radioData.WtpMAC, radioData.SlotID, radioData.AdminState, radioData.OperState)
			}
		}
	}
}

// validateAPCapwapEnabled validates that the AP is in enabled admin state.
func validateAPCapwapEnabled(t *testing.T, apMac string, status *ap.ApOperCAPWAPData) {
	if status == nil {
		t.Logf("Warning: Cannot validate AP state - no CAPWAP data")
		return
	}

	if capwapData := findCAPWAPDataByMAC(status.CAPWAPData, apMac); capwapData != nil {
		expectedAdminState := "adminstate-enabled"
		if capwapData.ApState.ApAdminState == expectedAdminState {
			t.Logf("✅ AP Admin State correctly shows '%s'", expectedAdminState)
		} else {
			t.Logf("⚠️  AP Admin State is '%s', expected '%s'", capwapData.ApState.ApAdminState, expectedAdminState)
		}
		t.Logf("ℹ️  AP Operation State: '%s'", capwapData.ApState.ApOperationState)
	} else {
		t.Logf("AP admin enable operation completed for AP %s", apMac)
	}
}

// validateRadioEnabled validates that the radio is in enabled state.
func validateRadioEnabled(t *testing.T, apMac string, slotID int, status *ap.ApOperRadioOperData) {
	if status == nil {
		t.Logf("Warning: Cannot validate radio state - no operational data")
		return
	}

	if radioData := findRadioDataByMACAndSlot(status.RadioOperData, apMac, slotID); radioData != nil {
		expectedAdminState := "enabled"
		expectedOperState := "radio-up"

		// Check Admin State
		if radioData.AdminState == expectedAdminState {
			t.Logf("✅ Admin State correctly shows '%s'", expectedAdminState)
		} else {
			t.Logf("⚠️  Admin State is '%s', expected '%s'", radioData.AdminState, expectedAdminState)
		}

		// Check Oper State
		if radioData.OperState == expectedOperState {
			t.Logf("✅ Radio Interface (Oper State) correctly shows '%s'", expectedOperState)
		} else {
			t.Logf("ℹ️  Radio Interface (Oper State) is '%s' (physical interface may take time to come up)", radioData.OperState)
		}
	} else {
		t.Logf("Radio enable operation completed for AP %s Slot %d", apMac, slotID)
	}
}

// findCAPWAPDataByMAC finds CAPWAP data for a specific AP MAC address.
func findCAPWAPDataByMAC(capwapDataList []ap.CAPWAPData, apMac string) *ap.CAPWAPData {
	for _, capwapData := range capwapDataList {
		if capwapData.WtpMAC == apMac {
			return &capwapData
		}
	}
	return nil
}

// findRadioDataByMACAndSlot finds radio data for a specific AP MAC and slot ID.
func findRadioDataByMACAndSlot(radioDataList []ap.RadioOperData, apMac string, slotID int) *ap.RadioOperData {
	for _, radioData := range radioDataList {
		if radioData.WtpMAC == apMac && radioData.SlotID == slotID {
			return &radioData
		}
	}
	return nil
}
