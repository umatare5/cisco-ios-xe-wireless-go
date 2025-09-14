package rfid

import "time"

// RfidGlobalOper represents RFID global operational data.
type RfidGlobalOper struct {
	RfidGlobalOperData struct {
		RfidTotalCount *RfidCountData  `json:"rfid-total-count,omitempty"` // Total unique RFID entries count (YANG: IOS-XE 17.12.1)
		RfidDataDetail []RfidEmltdData `json:"rfid-data-detail"`           // Detailed RFID data entries (YANG: IOS-XE 17.12.1)
		RfidRadioData  []RfidRadioData `json:"rfid-radio-data"`            // Known RFID tags operational data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data"` // RFID global operational data (YANG: IOS-XE 17.12.1)
}

// RfidCountData represents RFID count information.
type RfidCountData struct {
	TotalRfidCount uint32 `json:"total-rfid-count"` // Total unique RFID devices count (YANG: IOS-XE 17.12.1)
}

// RfidEmltdData represents RFID emulated parameters.
type RfidEmltdData struct {
	RfidMacAddr         string                   `json:"rfid-mac-addr"`          // RFID device MAC address (YANG: IOS-XE 17.12.1)
	RfidType            RfidDataType             `json:"rfid-type"`              // RFID tag type identifier (YANG: IOS-XE 17.12.1)
	RfidAutoInterval    uint16                   `json:"rfid-auto-interval"`     // RFID packet transmission interval (YANG: IOS-XE 17.12.1)
	RfidBytesRx         uint32                   `json:"rfid-bytes-rx"`          // Received RFID packet bytes count (YANG: IOS-XE 17.12.1)
	RfidPacketsRx       uint32                   `json:"rfid-packets-rx"`        // Received RFID packets count (YANG: IOS-XE 17.12.1)
	RfidLastHeardSecond time.Time                `json:"rfid-last-heard-second"` // Last RFID packet timestamp (YANG: IOS-XE 17.12.1)
	RfidVendor          RfidGlobalVendorSpecData `json:"rfid-vendor"`            // RFID vendor-specific data (YANG: IOS-XE 17.12.1)
	ApData              []RfidEmltdApData        `json:"ap-data"`                // AP information list for RFID packet (YANG: IOS-XE 17.12.1)
}

// RfidGlobalVendorSpecData represents global vendor-specific RFID packet data.
type RfidGlobalVendorSpecData struct {
	Bluesoft *RfidGlobalBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1)
	Cisco    *RfidGlobalCiscoData    `json:"cisco,omitempty"`    // Cisco RFID tag information (YANG: IOS-XE 17.12.1)
}

// RfidGlobalBluesoftData represents global Bluesoft RFID tag data.
type RfidGlobalBluesoftData struct {
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1)
	TagType    uint8 `json:"tag-type"`     // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1)
}

// RfidGlobalCiscoData represents global Cisco RFID tag data.
type RfidGlobalCiscoData struct {
	RfidCiscoHdr    RfidGlobalCiscoContHdr   `json:"rfid-cisco-hdr"`    // Cisco tag header information (YANG: IOS-XE 17.12.1)
	SeqControl      RfidGlobalApfSeqControl  `json:"seq-control"`       // RFID tag sequence control (YANG: IOS-XE 17.12.1)
	PayloadLen      uint16                   `json:"payload-len"`       // RFID packet payload length (YANG: IOS-XE 17.12.1)
	CcxPayload      RfidGlobalCcxPayloadList `json:"ccx-payload"`       // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1)
	CiscoVendorType RfidCiscoVendorType      `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1)
}

// RfidGlobalCiscoContHdr represents global Cisco RFID tag context header attributes.
type RfidGlobalCiscoContHdr struct {
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1)
	TxPower    int8  `json:"tx-power"`    // RFID transmit power (YANG: IOS-XE 17.12.1)
	Channel    uint8 `json:"channel"`     // RFID channel identifier (YANG: IOS-XE 17.12.1)
	RegClass   uint8 `json:"reg-class"`   // RFID regulatory class (YANG: IOS-XE 17.12.1)
	BurstLen   uint8 `json:"burst-len"`   // RFID burst length (YANG: IOS-XE 17.12.1)
}

// RfidGlobalApfSeqControl represents global RFID sequence and fragmentation control.
type RfidGlobalApfSeqControl struct {
	SeqNum  uint16 `json:"seq-num"`  // RFID tag packet sequence number (YANG: IOS-XE 17.12.1)
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1)
}

// RfidGlobalCcxPayloadList represents global CCX payload data container for RFID.
type RfidGlobalCcxPayloadList struct {
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1)
}

// RfidEmltdApData represents emulated AP data entry for RFID detection (max 16 entries).
type RfidEmltdApData struct {
	RssiValue      int8      `json:"rssi-value"`       // RFID RSSI measurement value (YANG: IOS-XE 17.12.1)
	Channel        uint16    `json:"channel"`          // RFID channel identifier (YANG: IOS-XE 17.12.1)
	SlotID         uint8     `json:"slot-id"`          // Access point slot identifier (YANG: IOS-XE 17.12.1)
	ApName         string    `json:"ap-name"`          // Access point name (YANG: IOS-XE 17.12.1)
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Last update received timestamp (YANG: IOS-XE 17.12.1)
	WtpMode        string    `json:"wtp-mode"`         // Wireless termination point mode (YANG: IOS-XE 17.12.1)
}

// RfidRadioData represents RFID radio detection entry data.
type RfidRadioData struct {
	RfidMacAddr   string    `json:"rfid-mac-addr"`   // RFID device MAC address (YANG: IOS-XE 17.12.1)
	ApMacAddr     string    `json:"ap-mac-addr"`     // Access point MAC address (YANG: IOS-XE 17.12.1)
	Slot          uint8     `json:"slot"`            // Access point slot number (YANG: IOS-XE 17.12.1)
	ApName        string    `json:"ap-name"`         // Access point name (YANG: IOS-XE 17.12.1)
	RssiLastHeard time.Time `json:"rssi-last-heard"` // RFID last heard timestamp (YANG: IOS-XE 17.12.1)
	RssiValue     int8      `json:"rssi-value"`      // RFID RSSI value in dBm (YANG: IOS-XE 17.12.1)
	SnrValue      int8      `json:"snr-value"`       // RFID SNR value in dBm (YANG: IOS-XE 17.12.1)
	RadioIntType  string    `json:"radio-int-type"`  // Detecting AP radio interface type (YANG: IOS-XE 17.12.1)
}
