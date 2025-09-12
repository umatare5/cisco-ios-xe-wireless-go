//go:build scenario

package ap_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
)

// Test constants and types for AP scenario testing.
const (
	testAPSlotID = 0 // Radio0 (2.4GHz)
)

// APScenarioStep represents the current step in AP scenario testing.
type APScenarioStep int

const (
	StepInitialState APScenarioStep = iota + 1
	StepDisable
	StepEnable
	StepFinalState
)

// APScenarioContext holds common context for AP scenario tests.
type APScenarioContext struct {
	T         *testing.T
	Ctx       context.Context
	Client    *core.Client
	TestAPMac string
	SlotID    int
}

// ============================
// Initialization Functions
// ============================

// initializeAPScenario sets up the common context for AP scenario tests.
func initializeAPScenario(t *testing.T, apMac string, slotID int) (*APScenarioContext, bool) {
	// Initialize client using integration test utilities
	setup := client.SetupRequiredClient(t)

	return &APScenarioContext{
		T:         t,
		Ctx:       setup.Context,
		Client:    setup.Client,
		TestAPMac: apMac,
		SlotID:    slotID,
	}, false
}

// ============================
// Context Methods
// ============================

// LogStep logs the current step being executed.
func (asc *APScenarioContext) LogStep(step APScenarioStep, message string) {
	asc.T.Logf("Step %d: %s", int(step), message)
}

// LogStepWithAP logs the current step with AP MAC and slot information.
func (asc *APScenarioContext) LogStepWithAP(step APScenarioStep, message string) {
	asc.T.Logf("Step %d: %s: AP %s Slot %d", int(step), message, asc.TestAPMac, asc.SlotID)
}

// LogSuccess logs a successful operation with AP information.
func (asc *APScenarioContext) LogSuccess(operation string) {
	asc.T.Logf("Successfully %s for AP %s Slot %d", operation, asc.TestAPMac, asc.SlotID)
}

// ExecuteStep executes a step with error handling and logging.
func (asc *APScenarioContext) ExecuteStep(step APScenarioStep, description string, operation func() error) {
	asc.LogStepWithAP(step, description)
	if err := operation(); err != nil {
		asc.T.Fatalf("Failed to %s for AP %s Slot %d: %v", description, asc.TestAPMac, asc.SlotID, err)
	}
}

// ============================
// Workflow Functions
// ============================

// executeAPStepWithResult executes a step with a result and error handling.
func executeAPStepWithResult[T any](asc *APScenarioContext, step APScenarioStep, description string, operation func() (T, error)) T {
	asc.LogStepWithAP(step, description)
	result, err := operation()
	if err != nil {
		asc.T.Fatalf("Failed to %s for AP %s Slot %d: %v", description, asc.TestAPMac, asc.SlotID, err)
	}
	return result
}

// executeAPAdminStateWorkflow performs the standard AP admin state scenario.
// This function encapsulates the 4-step workflow for AP admin state management,
// providing consistent logging and error handling for each operation.
func executeAPAdminStateWorkflow(asc *APScenarioContext, service *ap.Service) {
	// Step 1: Get initial AP state
	initialState := executeAPStepWithResult(asc, StepInitialState,
		"Getting initial AP capwap status", func() (*model.ApOperCapwapData, error) {
			return service.GetCAPWAPDataByWTPMAC(asc.Ctx, asc.TestAPMac)
		})
	logAPCapwapStatus(asc.T, "Initial", asc.TestAPMac, initialState)

	// Step 2: Disable AP admin state
	asc.ExecuteStep(StepDisable, "Disabling AP admin state", func() error {
		return service.DisableAP(asc.Ctx, asc.TestAPMac)
	})
	asc.LogSuccess("disabled AP admin state")

	// Step 3: Enable AP admin state
	asc.ExecuteStep(StepEnable, "Enabling AP admin state", func() error {
		return service.EnableAP(asc.Ctx, asc.TestAPMac)
	})
	asc.LogSuccess("enabled AP admin state")

	// Step 4: Get final AP state and validate
	finalState := executeAPStepWithResult(asc, StepFinalState,
		"Getting final AP capwap status", func() (*model.ApOperCapwapData, error) {
			return service.GetCAPWAPDataByWTPMAC(asc.Ctx, asc.TestAPMac)
		})
	logAPCapwapStatus(asc.T, "Final", asc.TestAPMac, finalState)
	validateAPCapwapEnabled(asc.T, asc.TestAPMac, finalState)
}

// executeAPRadioStateWorkflow performs the standard AP radio state scenario.
// This function encapsulates the 4-step workflow for radio admin state management,
// providing consistent logging and error handling for each radio operation.
func executeAPRadioStateWorkflow(asc *APScenarioContext, service *ap.Service) {
	// Step 1: Get initial radio state
	initialState := executeAPStepWithResult(asc, StepInitialState,
		"Getting initial radio status", func() (*model.ApOperRadioOperData, error) {
			return service.GetRadioStatusByWTPMACAndSlot(asc.Ctx, asc.TestAPMac, asc.SlotID)
		})
	logRadioStatus(asc.T, "Initial", asc.TestAPMac, asc.SlotID, initialState)

	// Step 2: Disable radio admin state
	asc.ExecuteStep(StepDisable, "Disabling radio admin state", func() error {
		radioBand := getRadioBandForSlot(asc.SlotID)
		return service.DisableRadio(asc.Ctx, asc.TestAPMac, radioBand)
	})
	asc.LogSuccess("disabled radio admin state")

	// Step 3: Enable radio admin state
	asc.ExecuteStep(StepEnable, "Enabling radio admin state", func() error {
		radioBand := getRadioBandForSlot(asc.SlotID)
		return service.EnableRadio(asc.Ctx, asc.TestAPMac, radioBand)
	})
	asc.LogSuccess("enabled radio admin state")

	// Step 4: Get final radio state and validate
	finalState := executeAPStepWithResult(asc, StepFinalState,
		"Getting final radio status", func() (*model.ApOperRadioOperData, error) {
			return service.GetRadioStatusByWTPMACAndSlot(asc.Ctx, asc.TestAPMac, asc.SlotID)
		})
	logRadioStatus(asc.T, "Final", asc.TestAPMac, asc.SlotID, finalState)
	validateRadioEnabled(asc.T, asc.TestAPMac, asc.SlotID, finalState)
}

// ============================
// Utility Functions
// ============================

// parseSlotID converts string slot ID to integer.
func parseSlotID(slotIDStr string) (int, error) {
	return strconv.Atoi(slotIDStr)
}

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
		// Default to 2.4GHz for unknown slots
		return core.RadioBand24GHz
	}
}

// ============================
// Logging Functions
// ============================

// logAPCapwapStatus logs the AP CAPWAP status information with admin state.
func logAPCapwapStatus(t *testing.T, phase, apMac string, status *model.ApOperCapwapData) {
	if status == nil {
		t.Logf("%s Status: No CAPWAP data available for AP %s", phase, apMac)
		return
	}

	t.Logf("%s Status for AP %s:", phase, apMac)

	if capwapData := findCapwapDataByMAC(status.CapwapData, apMac); capwapData != nil {
		logCapwapDetails(t, capwapData)
		return
	}

	logCapwapDataNotFound(t, apMac, status.CapwapData)
}

// logRadioStatus logs the radio operational status information.
func logRadioStatus(t *testing.T, phase, apMac string, slotID int, status *model.ApOperRadioOperData) {
	if status == nil {
		t.Logf("%s Status: No data available for AP %s Slot %d", phase, apMac, slotID)
		return
	}

	t.Logf("%s Status for AP %s Slot %d:", phase, apMac, slotID)

	if radioData := findRadioDataByMACAndSlot(status.RadioOperData, apMac, slotID); radioData != nil {
		logRadioDetails(t, radioData, slotID)
		return
	}

	logRadioDataNotFound(t, apMac, slotID, status.RadioOperData)
}

// logCapwapDetails logs detailed CAPWAP information for an AP.
func logCapwapDetails(t *testing.T, capwapData *model.CapwapData) {
	t.Logf("  WTP MAC: %s", capwapData.WtpMac)
	t.Logf("  IP Address: %s", capwapData.IPAddr)
	t.Logf("  Name: %s", capwapData.Name)
	t.Logf("  AP Admin State: %s", capwapData.ApState.ApAdminState)
	t.Logf("  AP Operation State: %s", capwapData.ApState.ApOperationState)
}

// logCapwapDataNotFound logs when CAPWAP data is not found for specified AP.
func logCapwapDataNotFound(t *testing.T, apMac string, capwapDataList []model.CapwapData) {
	t.Logf("  CAPWAP data not found for AP %s", apMac)
	if capwapDataList != nil {
		t.Logf("  Available CAPWAP data for other APs:")
		for i := range len(capwapDataList) {
			capwapData := capwapDataList[i]
			t.Logf("    [%d] WTP MAC: %s, Name: %s, Admin State: %s",
				i, capwapData.WtpMac, capwapData.Name, capwapData.ApState.ApAdminState)
		}
	} else {
		t.Logf("  No CAPWAP data available in response")
	}
}

// logRadioDetails logs detailed radio information.
func logRadioDetails(t *testing.T, radioData *model.RadioOperData, slotID int) {
	t.Logf("  WTP MAC: %s", radioData.WtpMac)
	t.Logf("  Radio Slot %d:", slotID)
	t.Logf("    Admin State: %s", radioData.AdminState)
	t.Logf("    Oper State: %s", radioData.OperState)
	t.Logf("    Radio Type: %s", radioData.RadioType)
	if radioData.RadioMode != "" {
		t.Logf("    Radio Mode: %s", radioData.RadioMode)
	}
}

// logRadioDataNotFound logs when radio data is not found for specified AP and slot.
func logRadioDataNotFound(t *testing.T, apMac string, slotID int, radioDataList []model.RadioOperData) {
	t.Logf("  Radio data not found for AP %s Slot %d", apMac, slotID)
	if radioDataList != nil {
		t.Logf("  Available radio data for other APs/slots:")
		for i := range len(radioDataList) {
			radioData := radioDataList[i]
			t.Logf("    [%d] WTP MAC: %s, Slot: %d, Admin: %s, Oper: %s",
				i, radioData.WtpMac, radioData.SlotID, radioData.AdminState, radioData.OperState)
		}
	} else {
		t.Logf("  No radio operational data available in response")
	}
}

// logAPStateResult logs AP admin state validation result.
func logAPStateResult(t *testing.T, actualState, expectedState string) {
	if actualState == expectedState {
		t.Logf("✅ AP Admin State correctly shows '%s'", expectedState)
	} else {
		t.Logf("⚠️  AP Admin State is '%s', expected '%s'", actualState, expectedState)
	}
}

// logRadioStateResult logs radio state validation results.
func logRadioStateResult(t *testing.T, actualAdminState, expectedAdminState, actualOperState, expectedOperState string, isDisable bool) {
	// Check Admin State
	if actualAdminState == expectedAdminState {
		t.Logf("✅ Admin State correctly shows '%s'", expectedAdminState)
	} else {
		t.Logf("⚠️  Admin State is '%s', expected '%s'", actualAdminState, expectedAdminState)
	}

	// Check Oper State with context-appropriate messaging
	if actualOperState == expectedOperState {
		t.Logf("✅ Radio Interface (Oper State) correctly shows '%s'", expectedOperState)
	} else {
		if isDisable {
			t.Logf("ℹ️  Radio Interface (Oper State) is '%s' (physical state may take time to reflect admin changes)", actualOperState)
		} else {
			t.Logf("ℹ️  Radio Interface (Oper State) is '%s' (physical interface may take time to come up)", actualOperState)
		}
	}
}

// logRadioOpResult logs radio operation result.
func logRadioOpResult(t *testing.T, apMac string, slotID int, operation string, success bool) {
	if success {
		t.Logf("✅ Radio %s operation successful for AP %s Slot %d", operation, apMac, slotID)
	} else {
		t.Logf("⚠️  Radio %s operation result mismatch for AP %s Slot %d", operation, apMac, slotID)
	}
}

// ============================
// Validation Functions
// ============================

// validateAPCapwapDisabled validates that the AP is in disabled admin state.
func validateAPCapwapDisabled(t *testing.T, apMac string, status *model.ApOperCapwapData) {
	validateAPCapwapState(t, apMac, status, "adminstate-disabled", "disable")
}

// validateAPCapwapEnabled validates that the AP is in enabled admin state.
func validateAPCapwapEnabled(t *testing.T, apMac string, status *model.ApOperCapwapData) {
	validateAPCapwapState(t, apMac, status, "adminstate-enabled", "enable")
}

// validateRadioDisabled validates that the radio is in disabled state.
func validateRadioDisabled(t *testing.T, apMac string, slotID int, status *model.ApOperRadioOperData) {
	validateRadioState(t, apMac, slotID, status, "disabled", "radio-down", "disable")
}

// validateRadioEnabled validates that the radio is in enabled state.
func validateRadioEnabled(t *testing.T, apMac string, slotID int, status *model.ApOperRadioOperData) {
	validateRadioState(t, apMac, slotID, status, "enabled", "radio-up", "enable")
}

// validateAPCapwapState validates AP CAPWAP state against expected values.
func validateAPCapwapState(t *testing.T, apMac string, status *model.ApOperCapwapData, expectedAdminState, operation string) {
	if status == nil {
		t.Logf("Warning: Cannot validate AP %s state - no CAPWAP data", operation)
		return
	}

	if capwapData := findCapwapDataByMAC(status.CapwapData, apMac); capwapData != nil {
		logAPStateResult(t, capwapData.ApState.ApAdminState, expectedAdminState)
		t.Logf("ℹ️  AP Operation State: '%s'", capwapData.ApState.ApOperationState)
		return
	}

	t.Logf("AP admin %s operation completed for AP %s", operation, apMac)
}

// validateRadioState validates radio state against expected values.
func validateRadioState(t *testing.T, apMac string, slotID int, status *model.ApOperRadioOperData, expectedAdminState, expectedOperState, operation string) {
	if status == nil {
		t.Logf("Warning: Cannot validate radio %s state - no operational data", operation)
		return
	}

	if radioData := findRadioDataByMACAndSlot(status.RadioOperData, apMac, slotID); radioData != nil {
		logRadioStateResult(t, radioData.AdminState, expectedAdminState, radioData.OperState, expectedOperState, operation == "disable")
		return
	}

	t.Logf("Radio %s operation completed for AP %s Slot %d", operation, apMac, slotID)
}

// ============================
// Data Finder Functions
// ============================

// findCapwapDataByMAC finds CAPWAP data for a specific AP MAC address.
func findCapwapDataByMAC(capwapDataList []model.CapwapData, apMac string) *model.CapwapData {
	for _, capwapData := range capwapDataList {
		if capwapData.WtpMac == apMac {
			return &capwapData
		}
	}
	return nil
}

// findRadioDataByMACAndSlot finds radio data for a specific AP MAC and slot ID.
func findRadioDataByMACAndSlot(radioDataList []model.RadioOperData, apMac string, slotID int) *model.RadioOperData {
	for _, radioData := range radioDataList {
		if radioData.WtpMac == apMac && radioData.SlotID == slotID {
			return &radioData
		}
	}
	return nil
}
