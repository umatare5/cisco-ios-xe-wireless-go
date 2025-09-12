// Package rfid provides data models for RFID global operational data.
package rfid

import "time"

// RfidGlobalOper represents RFID global operational data.
type RfidGlobalOper struct {
	RfidGlobalOperData RfidGlobalOperData `json:"Cisco-IOS-XE-wireless-rfid-global-oper:rfid-global-oper-data"`
}

// RfidGlobalOperData represents RFID global operational data container.
type RfidGlobalOperData struct {
	// RfidTotalCount represents total number of unique RFID entries.
	RfidTotalCount *RfidCountData `json:"rfid-total-count,omitempty"` // Total unique RFID entries count (YANG: IOS-XE 17.12.1+)

	// RfidDataDetail represents RFID data detail.
	RfidDataDetail []RfidEmltdData `json:"rfid-data-detail"` // Detailed RFID data entries (YANG: IOS-XE 17.12.1+)

	// RfidRadioData represents operational data of known RFID tags.
	RfidRadioData []RfidRadioData `json:"rfid-radio-data"` // Known RFID tags operational data (YANG: IOS-XE 17.12.1+)
}

// RfidCountData represents RFID count information.
type RfidCountData struct {
	// TotalRfidCount is the total number of unique RFID found in the RFID radio entry data.
	TotalRfidCount uint32 `json:"total-rfid-count"` // Total unique RFID devices count (YANG: IOS-XE 17.12.1+)
}

// RfidEmltdData represents RFID emulated parameters.
type RfidEmltdData struct {
	// RfidMacAddr represents RFID device mac address.
	RfidMacAddr string `json:"rfid-mac-addr"` // RFID device MAC address (YANG: IOS-XE 17.12.1+)

	// RfidType represents type of RFID tag.
	RfidType RfidDataType `json:"rfid-type"` // RFID tag type identifier (YANG: IOS-XE 17.12.1+)

	// RfidAutoInterval represents RFID packet interval.
	RfidAutoInterval uint16 `json:"rfid-auto-interval"` // RFID packet transmission interval (YANG: IOS-XE 17.12.1+)

	// RfidBytesRx represents number of bytes received in RFID packet.
	RfidBytesRx uint32 `json:"rfid-bytes-rx"` // Received RFID packet bytes count (YANG: IOS-XE 17.12.1+)

	// RfidPacketsRx represents number of RFID packets received.
	RfidPacketsRx uint32 `json:"rfid-packets-rx"` // Received RFID packets count (YANG: IOS-XE 17.12.1+)

	// RfidLastHeardSecond represents time stamp of last RFID packet received.
	RfidLastHeardSecond time.Time `json:"rfid-last-heard-second"` // Last RFID packet timestamp (YANG: IOS-XE 17.12.1+)

	// RfidVendor represents vendor type of RFID received.
	RfidVendor RfidGlobalVendorSpecData `json:"rfid-vendor"` // RFID vendor-specific data (YANG: IOS-XE 17.12.1+)

	// ApData represents AP information of RFID packet received.
	ApData []RfidEmltdApData `json:"ap-data"` // AP information list for RFID packet (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalVendorSpecData represents global vendor-specific RFID packet data.
type RfidGlobalVendorSpecData struct {
	// Bluesoft represents update of bluesoft RFID information.
	Bluesoft *RfidGlobalBluesoftData `json:"bluesoft,omitempty"` // Bluesoft RFID tag information (YANG: IOS-XE 17.12.1+)

	// Cisco represents Cisco RFID tag information.
	Cisco *RfidGlobalCiscoData `json:"cisco,omitempty"` // Cisco RFID tag information (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalBluesoftData represents global Bluesoft RFID tag data.
type RfidGlobalBluesoftData struct {
	// LastSeqNum represents last sequence number of RFID tag.
	LastSeqNum uint8 `json:"last-seq-num"` // Last sequence number of RFID tag (YANG: IOS-XE 17.12.1+)

	// TagType represents type of bluesoft RFID tag.
	TagType uint8 `json:"tag-type"` // Bluesoft RFID tag type (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalCiscoData represents global Cisco RFID tag data.
type RfidGlobalCiscoData struct {
	// RfidCiscoHdr represents Cisco tag header.
	RfidCiscoHdr RfidGlobalCiscoContHdr `json:"rfid-cisco-hdr"` // Cisco tag header information (YANG: IOS-XE 17.12.1+)

	// SeqControl represents sequence control number of RFID tag.
	SeqControl RfidGlobalApfSeqControl `json:"seq-control"` // RFID tag sequence control (YANG: IOS-XE 17.12.1+)

	// PayloadLen represents payload length of RFID packet.
	PayloadLen uint16 `json:"payload-len"` // RFID packet payload length (YANG: IOS-XE 17.12.1+)

	// CcxPayload represents CCX payload length of RFID packet.
	CcxPayload RfidGlobalCcxPayloadList `json:"ccx-payload"` // CCX payload data for RFID packet (YANG: IOS-XE 17.12.1+)

	// CiscoVendorType represents name of RFID vendor.
	CiscoVendorType RfidCiscoVendorType `json:"cisco-vendor-type"` // Cisco RFID vendor type (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalCiscoContHdr represents global Cisco RFID tag context header attributes.
type RfidGlobalCiscoContHdr struct {
	// CcxVersion represents CCX version of RFID.
	CcxVersion uint8 `json:"ccx-version"` // CCX version of RFID (YANG: IOS-XE 17.12.1+)

	// TxPower represents Tx power of RFID.
	TxPower int8 `json:"tx-power"` // RFID transmit power (YANG: IOS-XE 17.12.1+)

	// Channel represents Channel id of RFID.
	Channel uint8 `json:"channel"` // RFID channel identifier (YANG: IOS-XE 17.12.1+)

	// RegClass represents Regulatory class of RFID.
	RegClass uint8 `json:"reg-class"` // RFID regulatory class (YANG: IOS-XE 17.12.1+)

	// BurstLen represents Burst length of RFID.
	BurstLen uint8 `json:"burst-len"` // RFID burst length (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalApfSeqControl represents global RFID sequence and fragmentation control.
type RfidGlobalApfSeqControl struct {
	// SeqNum represents sequence number of RFID tag packet.
	SeqNum uint16 `json:"seq-num"` // RFID tag packet sequence number (YANG: IOS-XE 17.12.1+)

	// FragNum represents fragmentation number of RFID tag packet.
	FragNum uint16 `json:"frag-num"` // RFID tag packet fragmentation number (YANG: IOS-XE 17.12.1+)
}

// RfidGlobalCcxPayloadList represents global CCX payload data container for RFID.
type RfidGlobalCcxPayloadList struct {
	// Data contains one byte of CCX payload data (max 500 elements).
	Data []uint8 `json:"data"` // CCX payload data bytes (YANG: IOS-XE 17.12.1+)
}

// RfidEmltdApData represents emulated AP data entry for RFID detection (max 16 entries).
type RfidEmltdApData struct {
	// RssiValue represents RSSI parameters of RFID.
	RssiValue int8 `json:"rssi-value"` // RFID RSSI measurement value (YANG: IOS-XE 17.12.1+)

	// Channel represents RFID channel.
	Channel uint16 `json:"channel"` // RFID channel identifier (YANG: IOS-XE 17.12.1+)

	// SlotID represents AP slot id.
	SlotID uint8 `json:"slot-id"` // Access point slot identifier (YANG: IOS-XE 17.12.1+)

	// ApName represents value of AP name.
	ApName string `json:"ap-name"` // Access point name (YANG: IOS-XE 17.12.1+)

	// LastUpdateRcvd represents last update received.
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Last update received timestamp (YANG: IOS-XE 17.12.1+)

	// WtpMode represents AP mode (Note: this uses wireless types which should be imported).
	WtpMode string `json:"wtp-mode"` // Wireless termination point mode (YANG: IOS-XE 17.12.1+)
}

// RfidRadioData represents RFID radio detection entry data.
type RfidRadioData struct {
	// RfidMacAddr represents RFID MAC address (key).
	RfidMacAddr string `json:"rfid-mac-addr"` // RFID device MAC address (YANG: IOS-XE 17.12.1+)

	// ApMacAddr represents AP MAC address (key).
	ApMacAddr string `json:"ap-mac-addr"` // Access point MAC address (YANG: IOS-XE 17.12.1+)

	// Slot represents Slot ID (key, range 0..3).
	Slot uint8 `json:"slot"` // Access point slot number (YANG: IOS-XE 17.12.1+)

	// ApName represents AP Name.
	ApName string `json:"ap-name"` // Access point name (YANG: IOS-XE 17.12.1+)

	// RssiLastHeard represents time at which RFID was last heard by controller.
	RssiLastHeard time.Time `json:"rssi-last-heard"` // RFID last heard timestamp (YANG: IOS-XE 17.12.1+)

	// RssiValue represents RSSI value of the RFID tag when last heard by controller (dBm).
	RssiValue int8 `json:"rssi-value"` // RFID RSSI value in dBm (YANG: IOS-XE 17.12.1+)

	// SnrValue represents SNR signal to noise ratio value of RFID tag (dBm).
	SnrValue int8 `json:"snr-value"` // RFID SNR value in dBm (YANG: IOS-XE 17.12.1+)

	// RadioIntType represents radio type of the detecting AP that detected RFID.
	RadioIntType string `json:"radio-int-type"` // Detecting AP radio interface type (YANG: IOS-XE 17.12.1+)
}
