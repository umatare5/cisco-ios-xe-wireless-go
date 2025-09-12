package rfid

import "time"

// Enum types for RFID

// RfidCiscoVendorType represents the RFID Cisco vendor type enumeration.
type RfidCiscoVendorType int

const (
	RfidTypeG2        RfidCiscoVendorType = 0 // G2 vendor
	RfidTypeAeroscout RfidCiscoVendorType = 1 // Aero scout vendor
	RfidTypeUnknown   RfidCiscoVendorType = 2 // Unknown vendor
)

// RfidDataType represents the RFID data type enumeration.
type RfidDataType int

const (
	BluesoftRfidData RfidDataType = 0 // Bluesoft RFID tag
	CiscoRfidData    RfidDataType = 1 // Cisco RFID tag
)

// RfidOper represents RFID operational data.
type RfidOper struct {
	RfidOperData RfidOperData `json:"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data"`
}

// RfidOperData represents RFID operational data container.
type RfidOperData struct {
	RfidData []RfidData `json:"rfid-data"`
}

// RfidData represents RFID packet parameters.
type RfidData struct {
	// RfidMacAddr represents mac address of the RFID device.
	RfidMacAddr string `json:"rfid-mac-addr"` // RFID device MAC address (YANG: IOS-XE 17.12.1+)

	// RfidType represents type of RFID tag.
	RfidType RfidDataType `json:"rfid-type"` // RFID tag type (YANG: IOS-XE 17.12.1+)

	// RfidAutoInterval represents RFID packet interval.
	RfidAutoInterval uint16 `json:"rfid-auto-interval"` // RFID packet transmission interval (YANG: IOS-XE 17.12.1+)

	// RfidBytesRx represents number of bytes received in RFID packet.
	RfidBytesRx uint32 `json:"rfid-bytes-rx"` // Received RFID packet bytes count (YANG: IOS-XE 17.12.1+)

	// RfidPacketsRx represents number of RFID packets received for same RFID.
	RfidPacketsRx uint32 `json:"rfid-packets-rx"` // Received RFID packets count (YANG: IOS-XE 17.12.1+)

	// RfidLastHeardSecond represents time stamp of last RFID packet received.
	RfidLastHeardSecond time.Time `json:"rfid-last-heard-second"` // Last RFID packet timestamp (YANG: IOS-XE 17.12.1+)

	// RfidVendor represents vendor type of RFID received like cisco tag or bluesoft.
	RfidVendor RfidVendorSpecData `json:"rfid-vendor"` // RFID vendor-specific data (YANG: IOS-XE 17.12.1+)

	// ApHighRssi represents parameters of highest RSSI AP.
	ApHighRssi RfidApHighestRssi `json:"ap-high-rssi"` // Highest RSSI AP parameters (YANG: IOS-XE 17.12.1+)

	// ApData represents AP information of RFID packet received.
	ApData []RfidApEntryList `json:"ap-data"` // AP information list for RFID packet (YANG: IOS-XE 17.12.1+)

	// TxPower represents tx power of AP which send RFID packet.
	TxPower int8 `json:"tx-power"` // AP transmit power for RFID packet (YANG: IOS-XE 17.12.1+)

	// TimerHandle represents timer handle of the record.
	TimerHandle uint64 `json:"timer-handle"` // RFID record timer handle (YANG: IOS-XE 17.12.1+)

	// ApListLowestRssi represents parameters of the lowest RSSI AP.
	ApListLowestRssi *RfidApLowestRssi `json:"ap-list-lowest-rssi,omitempty"` // Lowest RSSI AP parameters (YANG: IOS-XE 17.12.1+)
}

// RfidVendorSpecData represents vendor-specific RFID packet data.
type RfidVendorSpecData struct {
	// Bluesoft represents update of bluesoft RFID information.
	Bluesoft *RfidBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1+)

	// Cisco represents cisco RFID tag information.
	Cisco *RfidCiscoData `json:"cisco,omitempty"` // Cisco RFID tag information (YANG: IOS-XE 17.12.1+)
}

// RfidBluesoftData represents Bluesoft RFID tag data.
type RfidBluesoftData struct {
	// LastSeqNum represents last sequence number of RFID tag.
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1+)

	// TagType represents type of bluesoft RFID tag.
	TagType uint8 `json:"tag-type"` // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1+)
}

// RfidCiscoData represents Cisco RFID tag data.
type RfidCiscoData struct {
	// RfidCiscoHdr represents cisco tag header.
	RfidCiscoHdr RfidCiscoContHdr `json:"rfid-cisco-hdr"` // Cisco tag header information (YANG: IOS-XE 17.12.1+)

	// SeqControl represents sequence control number of RFID tag.
	SeqControl RfidApfSeqControl `json:"seq-control"` // RFID tag sequence control (YANG: IOS-XE 17.12.1+)

	// PayloadLen represents payload length of RFID packet.
	PayloadLen uint16 `json:"payload-len"` // RFID packet payload length (YANG: IOS-XE 17.12.1+)

	// CcxPayload represents CCX payload length of RFID packet.
	CcxPayload RfidCcxPayloadList `json:"ccx-payload"` // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1+)

	// CiscoVendorType represents name of RFID vendor.
	CiscoVendorType RfidCiscoVendorType `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1+)
}

// RfidCiscoContHdr represents Cisco RFID tag context header attributes.
type RfidCiscoContHdr struct {
	// CcxVersion represents CCX version of RFID.
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1+)

	// TxPower represents tx power of RFID.
	TxPower int8 `json:"tx-power"` // RFID transmit power (YANG: IOS-XE 17.12.1+)

	// Channel represents channel id of RFID.
	Channel uint8 `json:"channel"` // RFID channel identifier (YANG: IOS-XE 17.12.1+)

	// RegClass represents reg_class of RFID.
	RegClass uint8 `json:"reg-class"` // RFID regulatory class (YANG: IOS-XE 17.12.1+)

	// BurstLen represents burst length of RFID.
	BurstLen uint8 `json:"burst-len"` // RFID burst length (YANG: IOS-XE 17.12.1+)
}

// RfidApfSeqControl represents RFID sequence and fragmentation control.
type RfidApfSeqControl struct {
	// SeqNum represents sequence number of RFID tag packet.
	SeqNum uint16 `json:"seq-num"` // RFID tag packet sequence number (YANG: IOS-XE 17.12.1+)

	// FragNum represents fragmentation number of RFID tag packet.
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1+)
}

// RfidCcxPayloadList represents CCX payload data container for RFID.
type RfidCcxPayloadList struct {
	// Data contains one byte of CCX payload data (max 500 elements).
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1+)
}

// RfidAvgRssi represents RFID RSSI measurement information.
type RfidAvgRssi struct {
	// Num represents RSSI num.
	Num int16 `json:"num"` // RSSI numerator value (YANG: IOS-XE 17.12.1+)

	// Denom represents RSSI denom.
	Denom uint16 `json:"denom"` // RSSI denominator value (YANG: IOS-XE 17.12.1+)

	// Value represents value of RSSI.
	Value int8 `json:"value"` // RSSI measurement value (YANG: IOS-XE 17.12.1+)

	// LastSent represents last sent RSSI.
	LastSent int8 `json:"last-sent"` // Last transmitted RSSI value (YANG: IOS-XE 17.12.1+)
}

// RfidAvgSnr represents RFID SNR measurement information.
type RfidAvgSnr struct {
	// Num of SNR.
	Num int16 `json:"num"` // SNR numerator value (YANG: IOS-XE 17.12.1+)

	// Denom of SNR.
	Denom uint16 `json:"denom"` // SNR denominator value (YANG: IOS-XE 17.12.1+)

	// Value of SNR.
	Value int8 `json:"value"` // SNR measurement value (YANG: IOS-XE 17.12.1+)
}

// RfidApEntryList represents AP data entry for RFID detection (max 16 entries).
type RfidApEntryList struct {
	// Rssi represents RSSI parameters of RFID.
	Rssi RfidAvgRssi `json:"rssi"` // RFID RSSI measurement parameters (YANG: IOS-XE 17.12.1+)

	// Slot represents AP slot.
	Slot uint8 `json:"slot"` // Access point slot number (YANG: IOS-XE 17.12.1+)

	// Band represents radio band of AP.
	Band uint8 `json:"band"` // AP radio band identifier (YANG: IOS-XE 17.12.1+)

	// Snr represents value of SNR.
	Snr RfidAvgSnr `json:"snr"` // RFID SNR measurement value (YANG: IOS-XE 17.12.1+)

	// TimeStamp represents time stamp of RFID packet received by the AP.
	TimeStamp time.Time `json:"time-stamp"` // RFID packet reception timestamp (YANG: IOS-XE 17.12.1+)

	// Expired represents stale node in the list, true if time stamp time expired.
	Expired bool `json:"expired"` // Entry expiration status flag (YANG: IOS-XE 17.12.1+)

	// ApMacAddr represents AP mac address of AP.
	ApMacAddr string `json:"ap-mac-addr"` // Access point MAC address (YANG: IOS-XE 17.12.1+)
}

// RfidApHighestRssi represents access point with highest RFID RSSI measurement.
type RfidApHighestRssi struct {
	// ApMacAddr represents AP mac address of highest RSSI AP.
	ApMacAddr string `json:"ap-mac-addr"` // Highest RSSI AP MAC address (YANG: IOS-XE 17.12.1+)

	// Rssi represents RSSI value of highest RSSI AP.
	Rssi RfidAvgRssi `json:"rssi"` // Highest RSSI measurement value (YANG: IOS-XE 17.12.1+)
}

// RfidApLowestRssi represents access point with lowest RFID RSSI measurement.
type RfidApLowestRssi struct {
	// ApMacAddr represents MAC address of the lowest RSSI AP.
	ApMacAddr string `json:"ap-mac-addr"` // Lowest RSSI AP MAC address (YANG: IOS-XE 17.12.1+)

	// Rssi represents lowest reported RSSI among all APs.
	Rssi RfidAvgRssi `json:"rssi"` // Lowest RSSI measurement value (YANG: IOS-XE 17.12.1+)
}
