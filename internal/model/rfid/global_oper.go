package rfid

import "time"

// RFIDGlobalOper represents RFID global operational data.
type RFIDGlobalOper struct {
	RFIDGlobalOperData struct {
		RFIDTotalCount *RFIDCountData  `json:"rfid-total-count,omitempty"` // Total unique RFID entries count (YANG: IOS-XE 17.12.1)
		RFIDDataDetail []RFIDEmltdData `json:"rfid-data-detail"`           // Detailed RFID data entries (YANG: IOS-XE 17.12.1)
		RFIDRadioData  []RFIDRadioData `json:"rfid-radio-data"`            // Known RFID tags operational data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data"` // RFID global operational data (YANG: IOS-XE 17.12.1)
}

// RFIDCountData represents RFID count information.
type RFIDCountData struct {
	TotalRFIDCount uint32 `json:"total-rfid-count"` // Total unique RFID devices count (YANG: IOS-XE 17.12.1)
}

// RFIDEmltdData represents RFID emulated parameters.
type RFIDEmltdData struct {
	RFIDMACAddr         string                   `json:"rfid-mac-addr"`          // RFID device MAC address (YANG: IOS-XE 17.12.1)
	RFIDType            RFIDDataType             `json:"rfid-type"`              // RFID tag type identifier (YANG: IOS-XE 17.12.1)
	RFIDAutoInterval    uint16                   `json:"rfid-auto-interval"`     // RFID packet transmission interval (YANG: IOS-XE 17.12.1)
	RFIDBytesRx         uint32                   `json:"rfid-bytes-rx"`          // Received RFID packet bytes count (YANG: IOS-XE 17.12.1)
	RFIDPacketsRx       uint32                   `json:"rfid-packets-rx"`        // Received RFID packets count (YANG: IOS-XE 17.12.1)
	RFIDLastHeardSecond time.Time                `json:"rfid-last-heard-second"` // Last RFID packet timestamp (YANG: IOS-XE 17.12.1)
	RFIDVendor          RFIDGlobalVendorSpecData `json:"rfid-vendor"`            // RFID vendor-specific data (YANG: IOS-XE 17.12.1)
	ApData              []RFIDEmltdApData        `json:"ap-data"`                // AP information list for RFID packet (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalVendorSpecData represents global vendor-specific RFID packet data.
type RFIDGlobalVendorSpecData struct {
	Bluesoft *RFIDGlobalBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1)
	Cisco    *RFIDGlobalCiscoData    `json:"cisco,omitempty"`    // Cisco RFID tag information (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalBluesoftData represents global Bluesoft RFID tag data.
type RFIDGlobalBluesoftData struct {
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1)
	TagType    uint8 `json:"tag-type"`     // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalCiscoData represents global Cisco RFID tag data.
type RFIDGlobalCiscoData struct {
	RFIDCiscoHdr    RFIDGlobalCiscoContHdr   `json:"rfid-cisco-hdr"`    // Cisco tag header information (YANG: IOS-XE 17.12.1)
	SeqControl      RFIDGlobalAPFSeqControl  `json:"seq-control"`       // RFID tag sequence control (YANG: IOS-XE 17.12.1)
	PayloadLen      uint16                   `json:"payload-len"`       // RFID packet payload length (YANG: IOS-XE 17.12.1)
	CcxPayload      RFIDGlobalCcxPayloadList `json:"ccx-payload"`       // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1)
	CiscoVendorType RFIDCiscoVendorType      `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalCiscoContHdr represents global Cisco RFID tag context header attributes.
type RFIDGlobalCiscoContHdr struct {
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1)
	TxPower    int8  `json:"tx-power"`    // RFID transmit power (YANG: IOS-XE 17.12.1)
	Channel    uint8 `json:"channel"`     // RFID channel identifier (YANG: IOS-XE 17.12.1)
	RegClass   uint8 `json:"reg-class"`   // RFID regulatory class (YANG: IOS-XE 17.12.1)
	BurstLen   uint8 `json:"burst-len"`   // RFID burst length (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalAPFSeqControl represents global RFID sequence and fragmentation control.
type RFIDGlobalAPFSeqControl struct {
	SeqNum  uint16 `json:"seq-num"`  // RFID tag packet sequence number (YANG: IOS-XE 17.12.1)
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1)
}

// RFIDGlobalCcxPayloadList represents global CCX payload data container for RFID.
type RFIDGlobalCcxPayloadList struct {
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1)
}

// RFIDEmltdApData represents emulated AP data entry for RFID detection (max 16 entries).
type RFIDEmltdApData struct {
	RSSIValue      int8      `json:"rssi-value"`       // RFID RSSI measurement value (YANG: IOS-XE 17.12.1)
	Channel        uint16    `json:"channel"`          // RFID channel identifier (YANG: IOS-XE 17.12.1)
	SlotID         uint8     `json:"slot-id"`          // Access point slot identifier (YANG: IOS-XE 17.12.1)
	ApName         string    `json:"ap-name"`          // Access point name (YANG: IOS-XE 17.12.1)
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Last update received timestamp (YANG: IOS-XE 17.12.1)
	WtpMode        string    `json:"wtp-mode"`         // Wireless termination point mode (YANG: IOS-XE 17.12.1)
}

// RFIDRadioData represents RFID radio detection entry data.
type RFIDRadioData struct {
	RFIDMACAddr   string    `json:"rfid-mac-addr"`   // RFID device MAC address (YANG: IOS-XE 17.12.1)
	ApMACAddr     string    `json:"ap-mac-addr"`     // Access point MAC address (YANG: IOS-XE 17.12.1)
	Slot          uint8     `json:"slot"`            // Access point slot number (YANG: IOS-XE 17.12.1)
	ApName        string    `json:"ap-name"`         // Access point name (YANG: IOS-XE 17.12.1)
	RSSILastHeard time.Time `json:"rssi-last-heard"` // RFID last heard timestamp (YANG: IOS-XE 17.12.1)
	RSSIValue     int8      `json:"rssi-value"`      // RFID RSSI value in dBm (YANG: IOS-XE 17.12.1)
	SNRValue      int8      `json:"snr-value"`       // RFID SNR value in dBm (YANG: IOS-XE 17.12.1)
	RadioIntType  string    `json:"radio-int-type"`  // Detecting AP radio interface type (YANG: IOS-XE 17.12.1)
}
