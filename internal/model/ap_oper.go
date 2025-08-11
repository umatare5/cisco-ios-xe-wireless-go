// Package model contains generated response structures for the Cisco WNC API.
package model

import "time"

// ApOperResponse represents the complete access point operational data response.
type ApOperResponse struct {
	CiscoIOSXEWirelessAccessPointOperAccessPointOperData struct {
		ApRadioNeighbor         []ApRadioNeighbor        `json:"ap-radio-neighbor"`
		RadioOperData           []RadioOperData          `json:"radio-oper-data"`
		RadioResetStats         []RadioResetStats        `json:"radio-reset-stats"`
		QosClientData           []QosClientData          `json:"qos-client-data"`
		CapwapData              []CapwapData             `json:"capwap-data"`
		ApNameMacMap            []ApNameMacMap           `json:"ap-name-mac-map"`
		WtpSlotWlanStats        []WtpSlotWlanStats       `json:"wtp-slot-wlan-stats"`
		EthernetMacWtpMacMap    []EthernetMacWtpMacMap   `json:"ethernet-mac-wtp-mac-map"`
		RadioOperStats          []RadioOperStats         `json:"radio-oper-stats"`
		EthernetIfStats         []EthernetIfStats        `json:"ethernet-if-stats"`
		EwlcWncdStats           EwlcWncdStats            `json:"ewlc-wncd-stats"`
		ApIoxOperData           []ApIoxOperData          `json:"ap-iox-oper-data"`
		QosGlobalStats          QosGlobalStats           `json:"qos-global-stats"`
		OperData                []ApOperData             `json:"oper-data"`
		RlanOper                []RlanOper               `json:"rlan-oper"`
		EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec  `json:"ewlc-mewlc-predownload-rec"`
		CdpCacheData            []CdpCacheData           `json:"cdp-cache-data"`
		LldpNeigh               []LldpNeigh              `json:"lldp-neigh"`
		TpCertInfo              TpCertInfo               `json:"tp-cert-info"`
		DiscData                []DiscData               `json:"disc-data"`
		CapwapPkts              []CapwapPkts             `json:"capwap-pkts"`
		CountryOper             []CountryOper            `json:"country-oper"`
		SuppCountryOper         []SuppCountryOper        `json:"supp-country-oper"`
		ApNhGlobalData          ApNhGlobalData           `json:"ap-nh-global-data"`
		ApImagePrepareLocation  []ApImagePrepareLocation `json:"ap-image-prepare-location"`
		ApImageActiveLocation   []ApImageActiveLocation  `json:"ap-image-active-location"`
	} `json:"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"`
}

// ApOperApRadioNeighborResponse represents the access point radio neighbor response
type ApOperApRadioNeighborResponse struct {
	ApRadioNeighbor []ApRadioNeighbor `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor"`
}

// ApRadioNeighbor represents access point radio neighbor information
type ApRadioNeighbor struct {
	ApMac          string    `json:"ap-mac"`
	SlotID         int       `json:"slot-id"`
	Bssid          string    `json:"bssid"`
	Ssid           string    `json:"ssid"`
	Rssi           int       `json:"rssi"`
	Channel        int       `json:"channel"`
	PrimaryChannel int       `json:"primary-channel"`
	LastUpdateRcvd time.Time `json:"last-update-rcvd"`
}

// Additional response wrapper types for individual endpoints would be defined here...
// (Truncated for brevity - these would include all the response types from ap/oper.go)

// RadioOperData represents radio operational data for an access point
type RadioOperData struct {
	WtpMac      string `json:"wtp-mac"`
	RadioSlotID int    `json:"radio-slot-id"`
	SlotID      int    `json:"slot-id,omitempty"`
	RadioType   string `json:"radio-type"`
	AdminState  string `json:"admin-state,omitempty"`
	OperState   string `json:"oper-state,omitempty"`
	RadioMode   string `json:"radio-mode,omitempty"`
	// Additional fields would be included here...
}

// RadioResetStats represents radio reset statistics data
type RadioResetStats struct {
	ApMac       string `json:"ap-mac"`
	RadioID     int    `json:"radio-id"`
	Cause       string `json:"cause"`
	DetailCause string `json:"detail-cause"`
	Count       int    `json:"count"`
}

type QosClientData struct {
	ClientMac    string `json:"client-mac"`
	AaaQosParams struct {
		AaaAvgdtus   int `json:"aaa-avgdtus"`
		AaaAvgrtdtus int `json:"aaa-avgrtdtus"`
		AaaBstdtus   int `json:"aaa-bstdtus"`
		AaaBstrtdtus int `json:"aaa-bstrtdtus"`
		AaaAvgdtds   int `json:"aaa-avgdtds"`
		AaaAvgrtdtds int `json:"aaa-avgrtdtds"`
		AaaBstdtds   int `json:"aaa-bstdtds"`
		AaaBstrtdtds int `json:"aaa-bstrtdtds"`
	} `json:"aaa-qos-params"`
}

type CapwapData struct {
	WtpMac       string       `json:"wtp-mac"`
	IPAddr       string       `json:"ip-addr"`
	Name         string       `json:"name"`
	DeviceDetail DeviceDetail `json:"device-detail"`
	// Additional fields would be included here...
}

type DeviceDetail struct {
	StaticInfo  StaticInfo  `json:"static-info"`
	DynamicInfo DynamicInfo `json:"dynamic-info"`
	WtpVersion  WtpVersion  `json:"wtp-version"`
}

type StaticInfo struct {
	BoardData struct {
		WtpSerialNum string `json:"wtp-serial-num"`
		WtpEnetMac   string `json:"wtp-enet-mac"`
		ApSysInfo    struct {
			MemType string `json:"mem-type"`
			CPUType string `json:"cpu-type"`
			MemSize int    `json:"mem-size"`
		} `json:"ap-sys-info"`
	} `json:"board-data"`
	// Additional fields...
}

type DynamicInfo struct {
	ApCrashData struct {
		ApCrashFile           string `json:"ap-crash-file"`
		ApRadio2GCrashFile    string `json:"ap-radio-2g-crash-file"`
		ApRadio5GCrashFile    string `json:"ap-radio-5g-crash-file"`
		ApRadio6GCrashFile    string `json:"ap-radio-6g-crash-file"`
		ApRad5GSlot2CrashFile string `json:"ap-rad-5g-slot2-crash-file"`
	} `json:"ap-crash-data"`
	// Additional fields...
}

type WtpVersion struct {
	BackupSwVersion struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"backup-sw-version"`
	// Additional fields...
}

// ApNameMacMap represents AP name to MAC address mapping data
type ApNameMacMap struct {
	WtpName string `json:"wtp-name"`
	WtpMac  string `json:"wtp-mac"`
	EthMac  string `json:"eth-mac"`
}

type WtpSlotWlanStats struct {
	WtpMac      string `json:"wtp-mac"`
	SlotID      int    `json:"slot-id"`
	WlanID      int    `json:"wlan-id"`
	BssidMac    string `json:"bssid-mac"`
	Ssid        string `json:"ssid"`
	BytesRx     string `json:"bytes-rx"`
	BytesTx     string `json:"bytes-tx"`
	PktsRx      string `json:"pkts-rx"`
	PktsTx      string `json:"pkts-tx"`
	DataRetries string `json:"data-retries"`
}

type EthernetMacWtpMacMap struct {
	EthernetMac string `json:"ethernet-mac"`
	WtpMac      string `json:"wtp-mac"`
}

type RadioOperStats struct {
	ApMac               string `json:"ap-mac"`
	SlotID              int    `json:"slot-id"`
	TxFragmentCount     int    `json:"tx-fragment-count"`
	MulticastTxFrameCnt int    `json:"multicast-tx-frame-cnt"`
	FailedCount         int    `json:"failed-count"`
	RetryCount          int    `json:"retry-count"`
	// Additional fields...
}

type EthernetIfStats struct {
	WtpMac     string `json:"wtp-mac"`
	IfIndex    int    `json:"if-index"`
	IfName     string `json:"if-name"`
	RxPkts     int    `json:"rx-pkts"`
	TxPkts     int    `json:"tx-pkts"`
	OperStatus string `json:"oper-status"`
	// Additional fields...
}

type EwlcWncdStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`
		NumInProgress           int  `json:"num-in-progress"`
		NumComplete             int  `json:"num-complete"`
		NumUnsupported          int  `json:"num-unsupported"`
		NumFailed               int  `json:"num-failed"`
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"`
		NumTotal                int  `json:"num-total"`
	} `json:"predownload-stats"`
	// Additional fields...
}

type ApIoxOperData struct {
	ApMac        string `json:"ap-mac"`
	ApphostState string `json:"apphost-state"`
	CafToken     string `json:"caf-token"`
	CafPort      int    `json:"caf-port"`
}

type QosGlobalStats struct {
	QosClientVoiceStats struct {
		TotalNumOfTspecRcvd      int `json:"total-num-of-tspec-rcvd"`
		NewTspecFromAssocReq     int `json:"new-tspec-from-assoc-req"`
		TspecRenewalFromAssocReq int `json:"tspec-renewal-from-assoc-req"`
		// Additional fields...
	} `json:"qos-client-voice-stats"`
}

type ApOperData struct {
	WtpMac                string `json:"wtp-mac"`
	RadioID               int    `json:"radio-id"`
	ApAntennaBandMode     string `json:"ap-antenna-band-mode"`
	LinkEncryptionEnabled bool   `json:"link-encryption-enabled"`
	// Additional fields...
}

type RlanOper struct {
	WtpMac         string `json:"wtp-mac"`
	RlanPortID     int    `json:"rlan-port-id"`
	RlanOperState  bool   `json:"rlan-oper-state"`
	RlanPortStatus bool   `json:"rlan-port-status"`
	RlanVlanValid  bool   `json:"rlan-vlan-valid"`
	RlanVlanID     int    `json:"rlan-vlan-id"`
	RlanPoeState   string `json:"rlan-poe-state"`
	PowerLevelID   int    `json:"power-level-id"`
}

type EwlcMewlcPredownloadRec struct {
	PredState                    string `json:"pred-state"`
	MeCapableApCount             int    `json:"me-capable-ap-count"`
	ControllerPredownloadVersion string `json:"controller-predownload-version"`
}

type CdpCacheData struct {
	MacAddr          string    `json:"mac-addr"`
	CdpCacheDeviceID string    `json:"cdp-cache-device-id"`
	ApName           string    `json:"ap-name"`
	LastUpdatedTime  time.Time `json:"last-updated-time"`
	Version          int       `json:"version"`
	// Additional fields...
}

type LldpNeigh struct {
	WtpMac          string `json:"wtp-mac"`
	NeighMac        string `json:"neigh-mac"`
	PortID          string `json:"port-id"`
	LocalPort       string `json:"local-port"`
	SystemName      string `json:"system-name"`
	PortDescription string `json:"port-description"`
	Capabilities    string `json:"capabilities"`
	MgmtAddr        string `json:"mgmt-addr"`
}

type TpCertInfo struct {
	Trustpoint Trustpoint `json:"trustpoint"`
}

type Trustpoint struct {
	TrustpointName     string `json:"trustpoint-name"`
	IsCertAvailable    bool   `json:"is-cert-available"`
	IsPrivkeyAvailable bool   `json:"is-privkey-available"`
	CertHash           string `json:"cert-hash"`
	CertType           string `json:"cert-type"`
	FipsSuitability    string `json:"fips-suitability"`
}

type DiscData struct {
	WtpMac           string `json:"wtp-mac"`
	DiscoveryPkts    string `json:"discovery-pkts"`
	DiscoveryErrPkts string `json:"discovery-err-pkts"`
}

type CapwapPkts struct {
	WtpMac            string `json:"wtp-mac"`
	CntrlPkts         string `json:"cntrl-pkts"`
	DataKeepAlivePkts string `json:"data-keep-alive-pkts"`
	CapwapErrorPkts   string `json:"capwap-error-pkts"`
	// Additional fields...
}

type CountryOper struct {
	CountryCode         string `json:"country-code"`
	CountryString       string `json:"country-string"`
	RegDomainStr80211Bg string `json:"reg-domain-str-80211bg"`
	RegDomainStr80211A  string `json:"reg-domain-str-80211a"`
	CountrySupported    bool   `json:"country-supported"`
	// Additional fields...
}

type SuppCountryOper struct {
	CountryCode    string `json:"country-code"`
	CountryString  string `json:"country-string"`
	CountryCodeIso string `json:"country-code-iso"`
	// Additional fields...
}

type ApNhGlobalData struct {
	AlgorithmRunning   bool `json:"algorithm-running"`
	AlgorithmItrCount  int  `json:"algorithm-itr-count"`
	IdealCapacityPerRg int  `json:"ideal-capacity-per-rg"`
	NumOfNeighborhood  int  `json:"num-of-neighborhood"`
}

type ApImagePrepareLocation struct {
	Index     int         `json:"index"`
	ImageFile string      `json:"image-file"`
	ImageData []ImageData `json:"image-data"`
}

type ImageData struct {
	ImageName     string   `json:"image-name"`
	ImageLocation string   `json:"image-location"`
	ImageVersion  string   `json:"image-version"`
	IsNew         bool     `json:"is-new"`
	FileSize      string   `json:"file-size"`
	ApModelList   []string `json:"ap-model-list"`
}

type ApImageActiveLocation struct {
	Index                          int    `json:"index"`
	ImageFile                      string `json:"image-file"`
	ApImageActiveLocationImageData []struct {
		ImageName                                 string   `json:"image-name"`
		ImageLocation                             string   `json:"image-location"`
		ImageVersion                              string   `json:"image-version"`
		IsNew                                     bool     `json:"is-new"`
		FileSize                                  string   `json:"file-size"`
		ApImageActiveLocationImageDataApModelList []string `json:"ap-model-list"`
	} `json:"image-data"`
}
