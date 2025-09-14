package rfid

import "time"

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
	RfidOperData struct {
		RfidData []RfidData `json:"rfid-data"` // RFID packet parameters list (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data"` // Wireless RFID operational data (YANG: IOS-XE 17.12.1)
}

// RfidData represents RFID packet parameters.
type RfidData struct {
	RfidMacAddr         string             `json:"rfid-mac-addr"`                 // RFID device MAC address (YANG: IOS-XE 17.12.1)
	RfidType            RfidDataType       `json:"rfid-type"`                     // RFID tag type (YANG: IOS-XE 17.12.1)
	RfidAutoInterval    uint16             `json:"rfid-auto-interval"`            // RFID packet transmission interval (YANG: IOS-XE 17.12.1)
	RfidBytesRx         uint32             `json:"rfid-bytes-rx"`                 // Received RFID packet bytes count (YANG: IOS-XE 17.12.1)
	RfidPacketsRx       uint32             `json:"rfid-packets-rx"`               // Received RFID packets count (YANG: IOS-XE 17.12.1)
	RfidLastHeardSecond time.Time          `json:"rfid-last-heard-second"`        // Last RFID packet timestamp (YANG: IOS-XE 17.12.1)
	RfidVendor          RfidVendorSpecData `json:"rfid-vendor"`                   // RFID vendor-specific data (YANG: IOS-XE 17.12.1)
	ApHighRssi          RfidApHighestRssi  `json:"ap-high-rssi"`                  // Highest RSSI AP parameters (YANG: IOS-XE 17.12.1)
	ApData              []RfidApEntryList  `json:"ap-data"`                       // AP information list for RFID packet (YANG: IOS-XE 17.12.1)
	TxPower             int8               `json:"tx-power"`                      // AP transmit power for RFID packet (YANG: IOS-XE 17.12.1)
	TimerHandle         uint64             `json:"timer-handle"`                  // RFID record timer handle (YANG: IOS-XE 17.12.1)
	ApListLowestRssi    *RfidApLowestRssi  `json:"ap-list-lowest-rssi,omitempty"` // Lowest RSSI AP parameters (YANG: IOS-XE 17.12.1)
}

// RfidVendorSpecData represents vendor-specific RFID packet data.
type RfidVendorSpecData struct {
	Bluesoft *RfidBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1)
	Cisco    *RfidCiscoData    `json:"cisco,omitempty"`    // Cisco RFID tag information (YANG: IOS-XE 17.12.1)
}

// RfidBluesoftData represents Bluesoft RFID tag data.
type RfidBluesoftData struct {
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1)
	TagType    uint8 `json:"tag-type"`     // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1)
}

// RfidCiscoData represents Cisco RFID tag data.
type RfidCiscoData struct {
	RfidCiscoHdr    RfidCiscoContHdr    `json:"rfid-cisco-hdr"`    // Cisco tag header information (YANG: IOS-XE 17.12.1)
	SeqControl      RfidApfSeqControl   `json:"seq-control"`       // RFID tag sequence control (YANG: IOS-XE 17.12.1)
	PayloadLen      uint16              `json:"payload-len"`       // RFID packet payload length (YANG: IOS-XE 17.12.1)
	CcxPayload      RfidCcxPayloadList  `json:"ccx-payload"`       // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1)
	CiscoVendorType RfidCiscoVendorType `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1)
}

// RfidCiscoContHdr represents Cisco RFID tag context header attributes.
type RfidCiscoContHdr struct {
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1)
	TxPower    int8  `json:"tx-power"`    // RFID transmit power (YANG: IOS-XE 17.12.1)
	Channel    uint8 `json:"channel"`     // RFID channel identifier (YANG: IOS-XE 17.12.1)
	RegClass   uint8 `json:"reg-class"`   // RFID regulatory class (YANG: IOS-XE 17.12.1)
	BurstLen   uint8 `json:"burst-len"`   // RFID burst length (YANG: IOS-XE 17.12.1)
}

// RfidApfSeqControl represents RFID sequence and fragmentation control.
type RfidApfSeqControl struct {
	SeqNum  uint16 `json:"seq-num"`  // RFID tag packet sequence number (YANG: IOS-XE 17.12.1)
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1)
}

// RfidCcxPayloadList represents CCX payload data container for RFID.
type RfidCcxPayloadList struct {
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1)
}

// RfidAvgRssi represents RFID RSSI measurement information.
type RfidAvgRssi struct {
	Num      int16  `json:"num"`       // RSSI numerator value (YANG: IOS-XE 17.12.1)
	Denom    uint16 `json:"denom"`     // RSSI denominator value (YANG: IOS-XE 17.12.1)
	Value    int8   `json:"value"`     // RSSI measurement value (YANG: IOS-XE 17.12.1)
	LastSent int8   `json:"last-sent"` // Last transmitted RSSI value (YANG: IOS-XE 17.12.1)
}

// RfidAvgSnr represents RFID SNR measurement information.
type RfidAvgSnr struct {
	Num   int16  `json:"num"`   // SNR numerator value (YANG: IOS-XE 17.12.1)
	Denom uint16 `json:"denom"` // SNR denominator value (YANG: IOS-XE 17.12.1)
	Value int8   `json:"value"` // SNR measurement value (YANG: IOS-XE 17.12.1)
}

// RfidApEntryList represents AP data entry for RFID detection (max 16 entries).
type RfidApEntryList struct {
	Rssi      RfidAvgRssi `json:"rssi"`        // RFID RSSI measurement parameters (YANG: IOS-XE 17.12.1)
	Slot      uint8       `json:"slot"`        // Access point slot number (YANG: IOS-XE 17.12.1)
	Band      uint8       `json:"band"`        // AP radio band identifier (YANG: IOS-XE 17.12.1)
	Snr       RfidAvgSnr  `json:"snr"`         // RFID SNR measurement value (YANG: IOS-XE 17.12.1)
	TimeStamp time.Time   `json:"time-stamp"`  // RFID packet reception timestamp (YANG: IOS-XE 17.12.1)
	Expired   bool        `json:"expired"`     // Entry expiration status flag (YANG: IOS-XE 17.12.1)
	ApMacAddr string      `json:"ap-mac-addr"` // Access point MAC address (YANG: IOS-XE 17.12.1)
}

// RfidApHighestRssi represents access point with highest RFID RSSI measurement.
type RfidApHighestRssi struct {
	ApMacAddr string      `json:"ap-mac-addr"` // Highest RSSI AP MAC address (YANG: IOS-XE 17.12.1)
	Rssi      RfidAvgRssi `json:"rssi"`        // Highest RSSI measurement value (YANG: IOS-XE 17.12.1)
}

// RfidApLowestRssi represents access point with lowest RFID RSSI measurement.
type RfidApLowestRssi struct {
	ApMacAddr string      `json:"ap-mac-addr"` // Lowest RSSI AP MAC address (YANG: IOS-XE 17.12.1)
	Rssi      RfidAvgRssi `json:"rssi"`        // Lowest RSSI measurement value (YANG: IOS-XE 17.12.1)
}
