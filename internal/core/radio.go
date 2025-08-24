// Package core provides radio-related helper functions and constants.
package core

import (
	"fmt"
)

// RadioBand represents the radio frequency band
type RadioBand int

// Radio configuration constants
const (
	// RadioBand enumeration values (from iota in ap package)
	RadioBand24GHzEnum = 0 // RadioBand24GHz = iota (0)
	RadioBand5GHzEnum  = 1 // RadioBand5GHz = iota (1)

	// RadioBand24GHzValue represents the band value for 2.4 GHz radio
	RadioBand24GHzValue = 1

	// RadioBand5GHzValue represents the band value for 5 GHz radio
	RadioBand5GHzValue = 2

	// RadioSlot24GHz represents the slot ID for 2.4 GHz radio
	RadioSlot24GHz = 0

	// RadioSlot5GHz represents the slot ID for 5 GHz radio
	RadioSlot5GHz = 1

	// AdminStateEnabled represents the enabled admin state
	AdminStateEnabled = "admin-state-enabled"

	// AdminStateDisabled represents the disabled admin state
	AdminStateDisabled = "admin-state-disabled"

	// OperationEnable represents the enable operation status
	OperationEnable = "enable"

	// OperationDisable represents the disable operation status
	OperationDisable = "disable"
)

const (
	// RadioBand24GHz represents 2.4GHz band (slot 0)
	RadioBand24GHz RadioBand = RadioBand24GHzEnum
	// RadioBand5GHz represents 5GHz band (slot 1)
	RadioBand5GHz RadioBand = RadioBand5GHzEnum
)

// RadioBandInfo represents radio band configuration information
type RadioBandInfo struct {
	Band   uint32
	SlotID uint8
}

// GetRadioBandInfo converts RadioBand enum to band and slot-id values
func GetRadioBandInfo(radioBand int) (RadioBandInfo, error) {
	switch radioBand {
	case RadioBand24GHzEnum: // RadioBand24GHz = 0 (iota)
		return RadioBandInfo{
			Band:   RadioBand24GHzValue,
			SlotID: RadioSlot24GHz,
		}, nil
	case RadioBand5GHzEnum: // RadioBand5GHz = 1 (iota)
		return RadioBandInfo{
			Band:   RadioBand5GHzValue,
			SlotID: RadioSlot5GHz,
		}, nil
	default:
		return RadioBandInfo{}, fmt.Errorf("unsupported radio band: %v", radioBand)
	}
}

// GetAdminStateMode returns the admin state mode string based on enabled flag
func GetAdminStateMode(enabled bool) string {
	if enabled {
		return AdminStateEnabled
	}
	return AdminStateDisabled
}
