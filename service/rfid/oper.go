package rfid

import "time"

// RFIDCiscoVendorType represents the RFID Cisco vendor type enumeration.
type RFIDCiscoVendorType int

const (
	RFIDTypeG2        RFIDCiscoVendorType = 0 // G2 vendor
	RFIDTypeAeroscout RFIDCiscoVendorType = 1 // Aero scout vendor
	RFIDTypeUnknown   RFIDCiscoVendorType = 2 // Unknown vendor
)

// RFIDDataType represents the RFID data type enumeration.
type RFIDDataType int

const (
	BluesoftRFIDData RFIDDataType = 0 // Bluesoft RFID tag
	CiscoRFIDData    RFIDDataType = 1 // Cisco RFID tag
)

// CiscoIOSXEWirelessRFIDOper represents RFID operational data.
type CiscoIOSXEWirelessRFIDOper struct {
	CiscoIOSXEWirelessRFIDOperData struct {
		RFIDData []RFIDData `json:"rfid-data"` // RFID packet parameters list (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rfid-oper:rfid-oper-data"` // Wireless RFID operational data (YANG: IOS-XE 17.12.1)
}

// RFIDData represents RFID packet parameters.
type RFIDData struct {
	RFIDMACAddr         string             `json:"rfid-mac-addr"`                 // RFID device MAC address (YANG: IOS-XE 17.12.1)
	RFIDType            RFIDDataType       `json:"rfid-type"`                     // RFID tag type (YANG: IOS-XE 17.12.1)
	RFIDAutoInterval    uint16             `json:"rfid-auto-interval"`            // RFID packet transmission interval (YANG: IOS-XE 17.12.1)
	RFIDBytesRx         uint32             `json:"rfid-bytes-rx"`                 // Received RFID packet bytes count (YANG: IOS-XE 17.12.1)
	RFIDPacketsRx       uint32             `json:"rfid-packets-rx"`               // Received RFID packets count (YANG: IOS-XE 17.12.1)
	RFIDLastHeardSecond time.Time          `json:"rfid-last-heard-second"`        // Last RFID packet timestamp (YANG: IOS-XE 17.12.1)
	RFIDVendor          RFIDVendorSpecData `json:"rfid-vendor"`                   // RFID vendor-specific data (YANG: IOS-XE 17.12.1)
	ApHighRSSI          RFIDApHighestRSSI  `json:"ap-high-rssi"`                  // Highest RSSI AP parameters (YANG: IOS-XE 17.12.1)
	ApData              []RFIDApEntryList  `json:"ap-data"`                       // AP information list for RFID packet (YANG: IOS-XE 17.12.1)
	TxPower             int8               `json:"tx-power"`                      // AP transmit power for RFID packet (YANG: IOS-XE 17.12.1)
	TimerHandle         uint64             `json:"timer-handle"`                  // RFID record timer handle (YANG: IOS-XE 17.12.1)
	ApListLowestRSSI    *RFIDApLowestRSSI  `json:"ap-list-lowest-rssi,omitempty"` // Lowest RSSI AP parameters (YANG: IOS-XE 17.12.1)
}

// RFIDVendorSpecData represents vendor-specific RFID packet data.
type RFIDVendorSpecData struct {
	Bluesoft *RFIDBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1)
	Cisco    *RFIDCiscoData    `json:"cisco,omitempty"`    // Cisco RFID tag information (YANG: IOS-XE 17.12.1)
}

// RFIDBluesoftData represents Bluesoft RFID tag data.
type RFIDBluesoftData struct {
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1)
	TagType    uint8 `json:"tag-type"`     // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1)
}

// RFIDCiscoData represents Cisco RFID tag data.
type RFIDCiscoData struct {
	RFIDCiscoHdr    RFIDCiscoContHdr    `json:"rfid-cisco-hdr"`    // Cisco tag header information (YANG: IOS-XE 17.12.1)
	SeqControl      RFIDApfSeqControl   `json:"seq-control"`       // RFID tag sequence control (YANG: IOS-XE 17.12.1)
	PayloadLen      uint16              `json:"payload-len"`       // RFID packet payload length (YANG: IOS-XE 17.12.1)
	CcxPayload      RFIDCcxPayloadList  `json:"ccx-payload"`       // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1)
	CiscoVendorType RFIDCiscoVendorType `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1)
}

// RFIDCiscoContHdr represents Cisco RFID tag context header attributes.
type RFIDCiscoContHdr struct {
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1)
	TxPower    int8  `json:"tx-power"`    // RFID transmit power (YANG: IOS-XE 17.12.1)
	Channel    uint8 `json:"channel"`     // RFID channel identifier (YANG: IOS-XE 17.12.1)
	RegClass   uint8 `json:"reg-class"`   // RFID regulatory class (YANG: IOS-XE 17.12.1)
	BurstLen   uint8 `json:"burst-len"`   // RFID burst length (YANG: IOS-XE 17.12.1)
}

// RFIDApfSeqControl represents RFID sequence and fragmentation control.
type RFIDApfSeqControl struct {
	SeqNum  uint16 `json:"seq-num"`  // RFID tag packet sequence number (YANG: IOS-XE 17.12.1)
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1)
}

// RFIDCcxPayloadList represents CCX payload data container for RFID.
type RFIDCcxPayloadList struct {
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1)
}

// RFIDAvgRSSI represents RFID RSSI measurement information.
type RFIDAvgRSSI struct {
	Num      int16  `json:"num"`       // RSSI numerator value (YANG: IOS-XE 17.12.1)
	Denom    uint16 `json:"denom"`     // RSSI denominator value (YANG: IOS-XE 17.12.1)
	Value    int8   `json:"value"`     // RSSI measurement value (YANG: IOS-XE 17.12.1)
	LastSent int8   `json:"last-sent"` // Last transmitted RSSI value (YANG: IOS-XE 17.12.1)
}

// RFIDAvgSNR represents RFID SNR measurement information.
type RFIDAvgSNR struct {
	Num   int16  `json:"num"`   // SNR numerator value (YANG: IOS-XE 17.12.1)
	Denom uint16 `json:"denom"` // SNR denominator value (YANG: IOS-XE 17.12.1)
	Value int8   `json:"value"` // SNR measurement value (YANG: IOS-XE 17.12.1)
}

// RFIDApEntryList represents AP data entry for RFID detection (max 16 entries).
type RFIDApEntryList struct {
	RSSI      RFIDAvgRSSI `json:"rssi"`        // RFID RSSI measurement parameters (YANG: IOS-XE 17.12.1)
	Slot      uint8       `json:"slot"`        // Access point slot number (YANG: IOS-XE 17.12.1)
	Band      uint8       `json:"band"`        // AP radio band identifier (YANG: IOS-XE 17.12.1)
	SNR       RFIDAvgSNR  `json:"snr"`         // RFID SNR measurement value (YANG: IOS-XE 17.12.1)
	TimeStamp time.Time   `json:"time-stamp"`  // RFID packet reception timestamp (YANG: IOS-XE 17.12.1)
	Expired   bool        `json:"expired"`     // Entry expiration status flag (YANG: IOS-XE 17.12.1)
	ApMACAddr string      `json:"ap-mac-addr"` // Access point MAC address (YANG: IOS-XE 17.12.1)
}

// RFIDApHighestRSSI represents access point with highest RFID RSSI measurement.
type RFIDApHighestRSSI struct {
	ApMACAddr string      `json:"ap-mac-addr"` // Highest RSSI AP MAC address (YANG: IOS-XE 17.12.1)
	RSSI      RFIDAvgRSSI `json:"rssi"`        // Highest RSSI measurement value (YANG: IOS-XE 17.12.1)
}

// RFIDApLowestRSSI represents access point with lowest RFID RSSI measurement.
type RFIDApLowestRSSI struct {
	ApMACAddr string      `json:"ap-mac-addr"` // Lowest RSSI AP MAC address (YANG: IOS-XE 17.12.1)
	RSSI      RFIDAvgRSSI `json:"rssi"`        // Lowest RSSI measurement value (YANG: IOS-XE 17.12.1)
}
