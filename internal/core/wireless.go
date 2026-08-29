package core

import (
	"fmt"
)

// RadioBand represents the radio frequency band.
type RadioBand int

// Radio configuration constants.
const (
	// RadioBand enumeration values (from iota in ap package).
	RadioBand24GHzEnum = iota
	RadioBand5GHzEnum
	RadioBand6GHzEnum

	// RadioBand24GHzValue represents the band value for 2.4 GHz radio.
	RadioBand24GHzValue = 1

	// RadioBand5GHzValue represents the band value for 5 GHz radio.
	RadioBand5GHzValue = 2

	// RadioBand6GHzValue is the band this SDK sends for RadioBand6GHz. The RPC's band names a
	// radio type, not a frequency: 3 is its dual-band selector and 4 its 6 GHz one.
	RadioBand6GHzValue = 3

	// RadioSlot24GHz is the slot this SDK sends with band 1.
	RadioSlot24GHz = 0

	// RadioSlot5GHz is the slot this SDK sends with band 2, which the RPC also admits on slot 2.
	RadioSlot5GHz = 1

	// RadioSlot6GHz is the slot this SDK sends with band 3, which the RPC also admits on slot 0.
	RadioSlot6GHz = 2

	// AdminStateEnabled represents the enabled admin state.
	AdminStateEnabled = "admin-state-enabled"

	// AdminStateDisabled represents the disabled admin state.
	AdminStateDisabled = "admin-state-disabled"

	// OperationEnable represents the enable operation status.
	OperationEnable = "enable"

	// OperationDisable represents the disable operation status.
	OperationDisable = "disable"
)

const (
	// RadioBand24GHz selects band 1 and slot 0.
	RadioBand24GHz RadioBand = RadioBand24GHzEnum
	// RadioBand5GHz selects band 2 and slot 1.
	RadioBand5GHz RadioBand = RadioBand5GHzEnum
	// RadioBand6GHz selects band 3 and slot 2.
	RadioBand6GHz RadioBand = RadioBand6GHzEnum
)

// RadioBandInfo represents radio band configuration information.
type RadioBandInfo struct {
	Band   uint32
	SlotID uint8
}

// GetRadioBandInfo converts RadioBand enum to band and slot-id values.
//
// The RPC pairs a band with a slot rather than deriving one from the other: its must clause
// admits seven (band, slot-id) pairs and only band 1 has a single slot, so the three pairs
// returned here are the three this enum can name, not the seven the RPC admits.
func GetRadioBandInfo(radioBand int) (RadioBandInfo, error) {
	switch radioBand {
	case RadioBand24GHzEnum:
		return RadioBandInfo{
			Band:   RadioBand24GHzValue,
			SlotID: RadioSlot24GHz,
		}, nil
	case RadioBand5GHzEnum:
		return RadioBandInfo{
			Band:   RadioBand5GHzValue,
			SlotID: RadioSlot5GHz,
		}, nil
	case RadioBand6GHzEnum:
		return RadioBandInfo{
			Band:   RadioBand6GHzValue,
			SlotID: RadioSlot6GHz,
		}, nil
	default:
		return RadioBandInfo{}, fmt.Errorf("radio band validation failed: %w",
			fmt.Errorf("unsupported radio band value %v, expected %d (2.4GHz) or %d (5GHz)",
				radioBand, RadioBand24GHzEnum, RadioBand5GHzEnum))
	}
}

// GetAdminStateMode returns the admin state mode string based on enabled flag.
func GetAdminStateMode(enabled bool) string {
	if enabled {
		return AdminStateEnabled
	}
	return AdminStateDisabled
}
