package ap

import "time"

// ApOper represents access point operational data response.
type ApOper struct {
	CiscoIOSXEWirelessAccessPointOperAccessPointOperData struct { // Root container of access point operational data (Live: IOS-XE 17.12.5)
		ApRadioNeighbor         []ApRadioNeighbor        `json:"ap-radio-neighbor"`          // AP radio neighbor information (Live: IOS-XE 17.12.5)
		RadioOperData           []RadioOperData          `json:"radio-oper-data"`            // Radio operational data corresponding to a radio of the 802.11 LWAPP AP (Live: IOS-XE 17.12.5)
		RadioResetStats         []RadioResetStats        `json:"radio-reset-stats"`          // Radio reset stats (Live: IOS-XE 17.12.5)
		QosClientData           []QosClientData          `json:"qos-client-data,omitempty"`  // QoS client data (YANG: IOS-XE 17.12.1)
		CapwapData              []CapwapData             `json:"capwap-data"`                // Information about the 802.11 LWAPP AP that has joined the controller (Live: IOS-XE 17.12.5)
		ApNameMacMap            []ApNameMacMap           `json:"ap-name-mac-map"`            // Mapping between AP name and radio MAC of AP (Live: IOS-XE 17.12.5)
		WtpSlotWlanStats        []WtpSlotWlanStats       `json:"wtp-slot-wlan-stats"`        // AP slot and WLAN stats (Live: IOS-XE 17.12.5)
		EthernetMacWtpMacMap    []EthernetMacWtpMacMap   `json:"ethernet-mac-wtp-mac-map"`   // Mapping between AP ethernet MAC and base radio MAC (Live: IOS-XE 17.12.5)
		RadioOperStats          []RadioOperStats         `json:"radio-oper-stats"`           // Operational statistics for a particular radio (Live: IOS-XE 17.12.5)
		EthernetIfStats         []EthernetIfStats        `json:"ethernet-if-stats"`          // Ethernet interface statistics (Live: IOS-XE 17.12.5)
		EwlcWncdStats           EwlcWncdStats            `json:"ewlc-wncd-stats"`            // AP image download and predownload statistics for EWC on AP platforms (Live: IOS-XE 17.12.5)
		ApIoxOperData           []ApIoxOperData          `json:"ap-iox-oper-data"`           // IOx application hosting operational data reported by the AP (Live: IOS-XE 17.12.5)
		QosGlobalStats          QosGlobalStats           `json:"qos-global-stats"`           // QoS Global statistics data in DB (Live: IOS-XE 17.12.5)
		OperData                []ApOperInternalData     `json:"oper-data"`                  // Operational data corresponding to an 802.11 LWAPP AP (Live: IOS-XE 17.12.5)
		RlanOper                []RlanOper               `json:"rlan-oper,omitempty"`        // LAN information of the AP (YANG: IOS-XE 17.12.1)
		EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec  `json:"ewlc-mewlc-predownload-rec"` // Embedded Wireless Controller predownload data (Live: IOS-XE 17.12.5)
		CdpCacheData            []CdpCacheData           `json:"cdp-cache-data"`             // Cached neighbor information via CDP messages on APs (Live: IOS-XE 17.12.5)
		LldpNeigh               []LldpNeigh              `json:"lldp-neigh"`                 // Cached neighbor information via LLDP messages on APs (Live: IOS-XE 17.12.5)
		TpCertInfo              TpCertInfo               `json:"tp-cert-info"`               // Trustpoint Certificate information (Live: IOS-XE 17.12.5)
		DiscData                []DiscData               `json:"disc-data"`                  // Discovery packet counters (Live: IOS-XE 17.12.5)
		CapwapPkts              []CapwapPkts             `json:"capwap-pkts"`                // CAPWAP packet counters (Live: IOS-XE 17.12.5)
		CountryOper             []CountryOper            `json:"country-oper"`               // Regulatory Domain country details (Live: IOS-XE 17.12.5)
		SuppCountryOper         []SuppCountryOper        `json:"supp-country-oper"`          // Supported Regulatory Domain country details (Live: IOS-XE 17.12.5)
		ApNhGlobalData          ApNhGlobalData           `json:"ap-nh-global-data"`          // Information about the RRM based AP clustering algorithm stats (Live: IOS-XE 17.12.5)
		ApImagePrepareLocation  []ApImagePrepareLocation `json:"ap-image-prepare-location"`  // AP image for prepare location (Live: IOS-XE 17.12.5)
		ApImageActiveLocation   []ApImageActiveLocation  `json:"ap-image-active-location"`   // AP image for active location (Live: IOS-XE 17.12.5)
		IotFirmware             []IotFirmware            `json:"iot-firmware"`               // IoT radio firmware operational data (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"` // Root container of access point operational data (Live: IOS-XE 17.12.5)
}

// ApOperApRadioNeighbor represents the access point radio neighbor response.
type ApOperApRadioNeighbor struct {
	ApRadioNeighbor []ApRadioNeighbor `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor"`
}

// ApOperRadioOperData represents the radio operational data response.
type ApOperRadioOperData struct {
	RadioOperData []RadioOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data"`
}

// ApOperRadioResetStats represents the radio reset statistics response.
type ApOperRadioResetStats struct {
	RadioResetStats []RadioResetStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-reset-stats"`
}

// ApOperQosClientData represents the QoS client data response.
type ApOperQosClientData struct {
	QosClientData []QosClientData `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data"`
}

// ApOperCapwapData represents the CAPWAP data response.
type ApOperCapwapData struct {
	CapwapData []CapwapData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
}

// ApOperApNameMacMap represents the AP name to MAC mapping response.
type ApOperApNameMacMap struct {
	ApNameMacMap []ApNameMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
}

// ApOperWtpSlotWlanStats represents the WTP slot WLAN statistics response.
type ApOperWtpSlotWlanStats struct {
	WtpSlotWlanStats []WtpSlotWlanStats `json:"Cisco-IOS-XE-wireless-access-point-oper:wtp-slot-wlan-stats"`
}

// ApOperEthernetMacWtpMacMap represents the Ethernet MAC to WTP MAC mapping response.
type ApOperEthernetMacWtpMacMap struct {
	EthernetMacWtpMacMap []EthernetMacWtpMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map"`
}

// ApOperRadioOperStats represents the radio operational statistics response.
type ApOperRadioOperStats struct {
	RadioOperStats []RadioOperStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-stats"`
}

// ApOperEthernetIfStats represents the Ethernet interface statistics response.
type ApOperEthernetIfStats struct {
	EthernetIfStats []EthernetIfStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats"`
}

// ApOperEwlcWncdStats represents the EWLC WNCD statistics response.
type ApOperEwlcWncdStats struct {
	EwlcWncdStats EwlcWncdStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-wncd-stats"`
}

// ApOperApIoxOperData represents the AP IOx operational data response.
type ApOperApIoxOperData struct {
	ApIoxOperData []ApIoxOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-iox-oper-data"`
}

// ApOperQosGlobalStats represents the QoS global statistics response.
type ApOperQosGlobalStats struct {
	QosGlobalStats QosGlobalStats `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-global-stats"`
}

// ApOperData represents the AP operational data response.
type ApOperData struct {
	OperData []ApOperInternalData `json:"Cisco-IOS-XE-wireless-access-point-oper:oper-data"`
}

// ApOperRlanOper represents the RLAN operational data response.
type ApOperRlanOper struct {
	RlanOper []RlanOper `json:"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper"`
}

// ApOperEwlcMewlcPredownloadRec represents the EWLC MEWLC predownload record response.
type ApOperEwlcMewlcPredownloadRec struct {
	EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-mewlc-predownload-rec"`
}

// ApOperCdpCacheData represents the CDP cache data response.
type ApOperCdpCacheData struct {
	CdpCacheData []CdpCacheData `json:"Cisco-IOS-XE-wireless-access-point-oper:cdp-cache-data"`
}

// ApOperLldpNeigh represents the LLDP neighbor response.
type ApOperLldpNeigh struct {
	LldpNeigh []LldpNeigh `json:"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh"`
}

// ApOperTpCertInfo represents the trustpoint certificate info response.
type ApOperTpCertInfo struct {
	TpCertInfo TpCertInfo `json:"Cisco-IOS-XE-wireless-access-point-oper:tp-cert-info"`
}

// ApOperDiscData represents the discovery data response.
type ApOperDiscData struct {
	DiscData []DiscData `json:"Cisco-IOS-XE-wireless-access-point-oper:disc-data"`
}

// ApOperCapwapPkts represents the CAPWAP packets response.
type ApOperCapwapPkts struct {
	CapwapPkts []CapwapPkts `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-pkts"`
}

// ApOperCountryOper represents the country operational data response.
type ApOperCountryOper struct {
	CountryOper []CountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:country-oper"`
}

// ApOperSuppCountryOper represents the supported country operational data response.
type ApOperSuppCountryOper struct {
	SuppCountryOper []SuppCountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:supp-country-oper"`
}

// ApOperApNhGlobalData represents the AP neighborhood global data response.
type ApOperApNhGlobalData struct {
	ApNhGlobalData ApNhGlobalData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-nh-global-data"`
}

// ApOperApImagePrepareLocation represents the AP image prepare location response.
type ApOperApImagePrepareLocation struct {
	ApImagePrepareLocation []ApImagePrepareLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-prepare-location"`
}

// ApOperApImageActiveLocation represents the AP image active location response.
type ApOperApImageActiveLocation struct {
	ApImageActiveLocation []ApImageActiveLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-active-location"`
}

// ApOperApPwrInfo represents the AP power information response.
type ApOperApPwrInfo struct {
	ApPwrInfo []ApPwrInfo `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-pwr-info"`
}

// ApOperApSensorStatus represents the AP sensor status response.
type ApOperApSensorStatus struct {
	ApSensorStatus []ApSensorStatus `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-sensor-status"`
}

// ApOperIotFirmware represents the IoT firmware response.
type ApOperIotFirmware struct {
	IotFirmware []IotFirmware `json:"Cisco-IOS-XE-wireless-access-point-oper:iot-firmware"`
}

// ApPwrInfo represents AP power information.
type ApPwrInfo struct {
	WtpMac  string    `json:"wtp-mac"`  // AP Radio MAC address (Live: IOS-XE 17.12.5)
	Status  string    `json:"status"`   // Power status (Live: IOS-XE 17.12.5)
	PpeInfo []PpeInfo `json:"ppe-info"` // Power policy entries (Live: IOS-XE 17.12.5)
}

// PpeInfo represents power policy entry information.
type PpeInfo struct {
	SeqNumber int    `json:"seq-number"` // Power policy sequence number (Live: IOS-XE 17.12.5)
	PpeResult string `json:"ppe-result"` // Power policy result (Live: IOS-XE 17.12.5)
	Ethernet  *struct {
		EthID    string `json:"eth-id"`    // Ethernet interface ID (Live: IOS-XE 17.12.5)
		EthSpeed string `json:"eth-speed"` // Ethernet speed (Live: IOS-XE 17.12.5)
	} `json:"ethernet,omitempty"` // Ethernet interface info (Live: IOS-XE 17.12.5)
	Radio *struct {
		RadioID       string  `json:"radio-id"`                 // Radio interface ID (Live: IOS-XE 17.12.5)
		SpatialStream *string `json:"spatial-stream,omitempty"` // Spatial stream config (Live: IOS-XE 17.12.5)
		State         *string `json:"state,omitempty"`          // Radio state (Live: IOS-XE 17.12.5)
	} `json:"radio,omitempty"` // Radio interface info (Live: IOS-XE 17.12.5)
	Usb *struct {
		UsbID string  `json:"usb-id"`          // USB interface ID (Live: IOS-XE 17.12.5)
		State *string `json:"state,omitempty"` // USB state (Live: IOS-XE 17.12.5)
	} `json:"usb,omitempty"` // USB interface info (Live: IOS-XE 17.12.5)
}

// ApSensorStatus represents AP sensor status information.
type ApSensorStatus struct {
	ApMac       string `json:"ap-mac"`       // AP MAC address (Live: IOS-XE 17.12.5)
	SensorType  string `json:"sensor-type"`  // Sensor type ID (Live: IOS-XE 17.12.5)
	ConfigState string `json:"config-state"` // Sensor config state (Live: IOS-XE 17.12.5)
	AdminState  string `json:"admin-state"`  // Admin state (Live: IOS-XE 17.12.5)
}

// ApRadioNeighbor represents AP radio neighbor information.
type ApRadioNeighbor struct {
	ApMac          string    `json:"ap-mac"`           // Access point MAC address (Live: IOS-XE 17.12.5)
	SlotID         int       `json:"slot-id"`          // Radio slot identifier (Live: IOS-XE 17.12.5)
	Bssid          string    `json:"bssid"`            // Basic Service Set Identifier (Live: IOS-XE 17.12.5)
	Ssid           string    `json:"ssid"`             // Service Set Identifier (Live: IOS-XE 17.12.5)
	Rssi           int       `json:"rssi"`             // Received Signal Strength Indicator (Live: IOS-XE 17.12.5)
	Channel        int       `json:"channel"`          // Operating channel number (Live: IOS-XE 17.12.5)
	PrimaryChannel int       `json:"primary-channel"`  // Primary channel number (Live: IOS-XE 17.12.5)
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Last neighbor update timestamp (Live: IOS-XE 17.12.5)
}

// RadioOperData represents radio operational data.
type RadioOperData struct {
	WtpMac       string `json:"wtp-mac"`                  // Wireless Termination Point MAC address (Live: IOS-XE 17.12.5)
	RadioSlotID  int    `json:"radio-slot-id"`            // Radio slot identifier (Live: IOS-XE 17.12.5)
	SlotID       int    `json:"slot-id,omitempty"`        // Physical slot identifier (Live: IOS-XE 17.12.5)
	RadioType    string `json:"radio-type,omitempty"`     // Radio hardware type (Live: IOS-XE 17.12.5)
	AdminState   string `json:"admin-state,omitempty"`    // Administrative state (Live: IOS-XE 17.12.5)
	OperState    string `json:"oper-state,omitempty"`     // Operational state (Live: IOS-XE 17.12.5)
	RadioMode    string `json:"radio-mode,omitempty"`     // Radio operational mode (Live: IOS-XE 17.12.5)
	RadioSubMode string `json:"radio-sub-mode,omitempty"` // Radio sub-mode details (Live: IOS-XE 17.12.5)
	RadioSubtype string `json:"radio-subtype,omitempty"`  // Radio hardware subtype (Live: IOS-XE 17.12.5)
	RadioSubband string `json:"radio-subband,omitempty"`  // Radio frequency subband (Live: IOS-XE 17.12.5)

	// Band and channel information
	CurrentBandID     int    `json:"current-band-id,omitempty"`     // Active band ID (Live: IOS-XE 17.12.5)
	CurrentActiveBand string `json:"current-active-band,omitempty"` // Active frequency band (Live: IOS-XE 17.12.5)

	// Protocol capabilities
	PhyHtCap        *PhyHtCapStruct `json:"phy-ht-cap,omitempty"`        // 802.11n HT capabilities (Live: IOS-XE 17.12.5)
	PhyHeCap        *PhyHeCapStruct `json:"phy-he-cap,omitempty"`        // 802.11ax HE capabilities (Live: IOS-XE 17.12.5)
	RadioHeCapable  bool            `json:"radio-he-capable,omitempty"`  // 802.11ax capability status (Live: IOS-XE 17.12.5)
	RadioFraCapable string          `json:"radio-fra-capable,omitempty"` // Frame aggregation capability (Live: IOS-XE 17.12.5)

	// XOR capabilities
	XorRadioMode string       `json:"xor-radio-mode,omitempty"` // XOR radio mode (Live: IOS-XE 17.12.5)
	XorPhyHtCap  *XorPhyHtCap `json:"xor-phy-ht-cap,omitempty"` // XOR HT capabilities (Live: IOS-XE 17.12.5)
	XorPhyHeCap  *XorPhyHeCap `json:"xor-phy-he-cap,omitempty"` // XOR HE capabilities (Live: IOS-XE 17.12.5)

	// Additional operational fields
	AntennaGain            int               `json:"antenna-gain,omitempty"`             // Antenna gain value in dBi (Live: IOS-XE 17.12.5)
	AntennaPid             string            `json:"antenna-pid,omitempty"`              // Antenna product identifier (Live: IOS-XE 17.12.5)
	SlotAntennaType        string            `json:"slot-antenna-type,omitempty"`        // Antenna type (internal/external) (Live: IOS-XE 17.12.5)
	RadioEnableTime        string            `json:"radio-enable-time,omitempty"`        // Last radio enable timestamp (Live: IOS-XE 17.12.5)
	HighestThroughputProto string            `json:"highest-throughput-proto,omitempty"` // Highest throughput protocol supported (Live: IOS-XE 17.12.5)
	CacActive              bool              `json:"cac-active,omitempty"`               // Channel Availability Check active status (Live: IOS-XE 17.12.5)
	MeshBackhaul           bool              `json:"mesh-backhaul,omitempty"`            // Mesh backhaul link status (Live: IOS-XE 17.12.5)
	MeshDesignatedDownlink bool              `json:"mesh-designated-downlink,omitempty"` // Mesh designated downlink status (Live: IOS-XE 17.12.5)
	MultiDomainCap         *MultiDomainCap   `json:"multi-domain-cap,omitempty"`         // Multi-domain capabilities (Live: IOS-XE 17.12.5)
	StationCfg             *StationCfg       `json:"station-cfg,omitempty"`              // Station mode configuration (Live: IOS-XE 17.12.5)
	PhyHtCfg               *PhyHtCfg         `json:"phy-ht-cfg,omitempty"`               // High Throughput configuration settings (Live: IOS-XE 17.12.5)
	ChanPwrInfo            *ChanPwrInfo      `json:"chan-pwr-info,omitempty"`            // Channel power level information (Live: IOS-XE 17.12.5)
	SnifferCfg             *SnifferCfg       `json:"sniffer-cfg,omitempty"`              // Packet sniffer configuration (Live: IOS-XE 17.12.5)
	RadioBandInfo          []RadioBandInfo   `json:"radio-band-info,omitempty"`          // Radio frequency band information (Live: IOS-XE 17.12.5)
	VapOperConfig          []VapOperConfig   `json:"vap-oper-config,omitempty"`          // VAP operational configuration (Live: IOS-XE 17.12.5)
	RegDomainCheckStatus   string            `json:"reg-domain-check-status,omitempty"`  // Regulatory compliance check status (Live: IOS-XE 17.12.5)
	Dot11nMcsRates         string            `json:"dot11n-mcs-rates,omitempty"`         // 802.11n MCS rates supported (Live: IOS-XE 17.12.5)
	DualRadioModeCfg       *DualRadioModeCfg `json:"dual-radio-mode-cfg,omitempty"`      // Dual radio mode configuration (Live: IOS-XE 17.12.5)
	BssColorCfg            *BssColorCfg      `json:"bss-color-cfg,omitempty"`            // BSS color configuration for 802.11ax (Live: IOS-XE 17.12.5)
	ObssPdCapable          bool              `json:"obss-pd-capable,omitempty"`          // OBSS Preamble Detection capability (Live: IOS-XE 17.12.5)
	NdpCap                 string            `json:"ndp-cap,omitempty"`                  // Null Data Packet capability information (Live: IOS-XE 17.12.5)
	NdpOnChannel           bool              `json:"ndp-on-channel,omitempty"`           // Null Data Packet transmission on channel status (Live: IOS-XE 17.12.5)
	BeamSelection          string            `json:"beam-selection,omitempty"`           // Antenna beam selection algorithm configuration (Live: IOS-XE 17.12.5)
	NumAntEnabled          uint8             `json:"num-ant-enabled,omitempty"`          // Number of antennas enabled (Live: IOS-XE 17.12.5)
	CurAntBitmap           string            `json:"cur-ant-bitmap,omitempty"`           // Current antenna bitmap (Live: IOS-XE 17.12.5)
	SuppAntBitmap          string            `json:"supp-ant-bitmap,omitempty"`          // Supported antenna bitmap (Live: IOS-XE 17.12.5)
	Supp160mhzAntBitmap    string            `json:"supp-160mhz-ant-bitmap,omitempty"`   // 160MHz antenna bitmap (Live: IOS-XE 17.12.5)
	MaxClientAllowed       uint16            `json:"max-client-allowed,omitempty"`       // Maximum clients allowed (Live: IOS-XE 17.12.5)
	ObssPdSrgCapable       bool              `json:"obss-pd-srg-capable,omitempty"`      // OBSS PD SRG capability (Live: IOS-XE 17.12.5)
	CoverageOverlapFactor  uint8             `json:"coverage-overlap-factor,omitempty"`  // RF coverage overlap factor (Live: IOS-XE 17.12.5)

	// 6GHz related (YANG: IOS-XE 17.12.1)
	Ap6GhzPwrMode    *string `json:"ap-6ghz-pwr-mode,omitempty"`     // 6GHz power mode (LPI/SP/VLP) (YANG: IOS-XE 17.12.1)
	Ap6GhzPwrModeCap *string `json:"ap-6ghz-pwr-mode-cap,omitempty"` // 6GHz power mode capability (YANG: IOS-XE 17.12.1)

	// AFC related
	AfcBelowTxmin    bool `json:"afc-below-txmin,omitempty"`    // AFC below minimum transmission power (YANG: IOS-XE 17.12.1)
	AfcLicenseNeeded bool `json:"afc-license-needed,omitempty"` // AFC license requirement status (YANG: IOS-XE 17.12.1)
	PushAfcRespDone  bool `json:"push-afc-resp-done,omitempty"` // AFC response push completion status (YANG: IOS-XE 17.12.1)
}

// RadioResetStats represents radio reset statistics.
type RadioResetStats struct {
	ApMac       string `json:"ap-mac"`       // Access point MAC address (Live: IOS-XE 17.12.5)
	RadioID     int    `json:"radio-id"`     // Radio interface identifier (Live: IOS-XE 17.12.5)
	Cause       string `json:"cause"`        // Reset cause description (Live: IOS-XE 17.12.5)
	DetailCause string `json:"detail-cause"` // Detailed reset cause information (Live: IOS-XE 17.12.5)
	Count       int    `json:"count"`        // Reset count since statistics clear (Live: IOS-XE 17.12.5)
}

// QosClientData represents QoS client data.
type QosClientData struct {
	ClientMac    string `json:"client-mac"` // Client MAC address (Live: IOS-XE 17.12.5)
	AaaQosParams struct {
		AaaAvgdtus   int `json:"aaa-avgdtus"`   // AAA average downstream utilization (Live: IOS-XE 17.12.5)
		AaaAvgrtdtus int `json:"aaa-avgrtdtus"` // AAA average real-time downstream utilization (Live: IOS-XE 17.12.5)
		AaaBstdtus   int `json:"aaa-bstdtus"`   // AAA burst downstream utilization (Live: IOS-XE 17.12.5)
		AaaBstrtdtus int `json:"aaa-bstrtdtus"` // AAA burst real-time downstream utilization (Live: IOS-XE 17.12.5)
		AaaAvgdtds   int `json:"aaa-avgdtds"`   // AAA average downstream data size (Live: IOS-XE 17.12.5)
		AaaAvgrtdtds int `json:"aaa-avgrtdtds"` // AAA average real-time downstream data size (Live: IOS-XE 17.12.5)
		AaaBstdtds   int `json:"aaa-bstdtds"`   // AAA burst downstream data size (Live: IOS-XE 17.12.5)
		AaaBstrtdtds int `json:"aaa-bstrtdtds"` // AAA burst real-time downstream data size (Live: IOS-XE 17.12.5)
	} `json:"aaa-qos-params"` // AAA QoS parameters (Live: IOS-XE 17.12.5)
}

// CapwapData represents CAPWAP data.
type CapwapData struct {
	WtpMac       string       `json:"wtp-mac"`       // WTP MAC address for CAPWAP session (Live: IOS-XE 17.12.5)
	IPAddr       string       `json:"ip-addr"`       // AP management IP address (Live: IOS-XE 17.12.5)
	Name         string       `json:"name"`          // AP hostname identifier (Live: IOS-XE 17.12.5)
	DeviceDetail DeviceDetail `json:"device-detail"` // Hardware device specifications (Live: IOS-XE 17.12.5)
	ApState      ApState      `json:"ap-state"`      // AP operational and admin state (Live: IOS-XE 17.12.5)

	// AP Mode Data
	ApModeData ApModeData `json:"ap-mode-data"` // AP operational mode configuration (Live: IOS-XE 17.12.5)

	// Location and Services
	ApLocation         ApLocation         `json:"ap-location"`          // Physical deployment location (Live: IOS-XE 17.12.5)
	ApServices         ApServices         `json:"ap-services"`          // Enabled AP service capabilities (Live: IOS-XE 17.12.5)
	TagInfo            TagInfo            `json:"tag-info"`             // Policy and site tag assignment (Live: IOS-XE 17.12.5)
	Tunnel             Tunnel             `json:"tunnel"`               // CAPWAP tunnel configuration (Live: IOS-XE 17.12.5)
	ExternalModuleData ExternalModuleData `json:"external-module-data"` // USB and expansion module information (Live: IOS-XE 17.12.5)
	ApTimeInfo         ApTimeInfo         `json:"ap-time-info"`         // Time synchronization data (Live: IOS-XE 17.12.5)
	ApSecurityData     ApSecurityData     `json:"ap-security-data"`     // Security configuration (Live: IOS-XE 17.12.5)
	SlidingWindow      SlidingWindow      `json:"sliding-window"`       // CAPWAP sliding window parameters (Live: IOS-XE 17.12.5)
	ApVlan             ApVlan             `json:"ap-vlan"`              // VLAN tagging configuration (Live: IOS-XE 17.12.5)
	HyperlocationData  HyperlocationData  `json:"hyperlocation-data"`   // Hyperlocation service configuration (Live: IOS-XE 17.12.5)
	RebootStats        RebootStats        `json:"reboot-stats"`         // AP reboot history and analysis (Live: IOS-XE 17.12.5)
	ProxyInfo          ProxyInfo          `json:"proxy-info"`           // HTTP proxy configuration (Live: IOS-XE 17.12.5)

	// Image Download Tracking
	ImageSizeEta           uint64 `json:"image-size-eta"`            // AP firmware image download ETA (Live: IOS-XE 17.12.5)
	ImageSizeStartTime     string `json:"image-size-start-time"`     // AP image download start timestamp (Live: IOS-XE 17.12.5)
	ImageSizePercentage    uint32 `json:"image-size-percentage"`     // AP image download progress percentage (Live: IOS-XE 17.12.5)
	WlcImageSizeEta        uint64 `json:"wlc-image-size-eta"`        // WLC firmware image download ETA (Live: IOS-XE 17.12.5)
	WlcImageSizeStartTime  string `json:"wlc-image-size-start-time"` // WLC image download start timestamp (Live: IOS-XE 17.12.5)
	WlcImageSizePercentage uint32 `json:"wlc-image-size-percentage"` // WLC image download progress percentage (Live: IOS-XE 17.12.5)

	// Local DHCP Configuration
	Ipv4Pool              Ipv4Pool          `json:"ipv4-pool"`                // Local DHCP IPv4 pool configuration (Live: IOS-XE 17.12.5)
	DisconnectDetail      DisconnectDetail  `json:"disconnect-detail"`        // AP disconnection analysis (Live: IOS-XE 17.12.5)
	StatsMonitor          StatsMonitor      `json:"stats-monitor"`            // AP statistics monitoring configuration (Live: IOS-XE 17.12.5)
	LscStatusPldSupported []interface{}     `json:"lsc-status-pld-supported"` // LSC status payload support capability (Live: IOS-XE 17.12.5)
	ApLscStatus           ApLscStatus       `json:"ap-lsc-status"`            // LSC authentication status (Live: IOS-XE 17.12.5)
	RadioStatsMonitor     RadioStatsMonitor `json:"radio-stats-monitor"`      // Radio statistics monitoring (Live: IOS-XE 17.12.5)
	ZeroWtDfs             ZeroWtDfs         `json:"zero-wt-dfs"`              // Zero Wait DFS configuration (Live: IOS-XE 17.12.5)
	GnssInfo              GnssInfo          `json:"gnss-info"`                // GNSS positioning data (Live: IOS-XE 17.12.5)

	// Basic Configuration Fields
	ApLagEnabled    bool   `json:"ap-lag-enabled"`    // LAG configuration status (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	CountryCode     string `json:"country-code"`      // Regulatory country code (Live: IOS-XE 17.12.5)
	NumRadioSlots   uint8  `json:"num-radio-slots"`   // Number of radio slots available (Live: IOS-XE 17.12.5)
	Ipv6Joined      uint8  `json:"ipv6-joined"`       // IPv6 join status (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	DartIsConnected bool   `json:"dart-is-connected"` // DART connection status (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	IsMaster        bool   `json:"is-master"`         // Master AP designation (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	CdpEnable       bool   `json:"cdp-enable"`        // CDP enablement (Live: IOS-XE 17.12.5)
	GrpcEnabled     bool   `json:"grpc-enabled"`      // gRPC streaming enablement (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	LocalDhcp       bool   `json:"local-dhcp"`        // Local DHCP server status (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)

	// Status and operational fields
	ApStationType        string `json:"ap-stationing-type,omitempty"`      // AP stationing type configuration (Live: IOS-XE 17.12.5)
	ApKeepAliveState     bool   `json:"ap-keepalive-state,omitempty"`      // CAPWAP keep-alive state (Live: IOS-XE 17.12.5)
	MaxClientsSupported  uint16 `json:"max-clients-supported,omitempty"`   // Maximum clients supported (Live: IOS-XE 17.12.5)
	MdnsGroupID          uint32 `json:"mdns-group-id,omitempty"`           // mDNS group identifier (Live: IOS-XE 17.12.5)
	MdnsRuleName         string `json:"mdns-rule-name,omitempty"`          // Applied mDNS filtering rule name (Live: IOS-XE 17.12.5)
	MdnsGroupMethod      string `json:"mdns-group-method,omitempty"`       // mDNS group assignment method (Live: IOS-XE 17.12.5)
	MerakiCapable        bool   `json:"meraki-capable,omitempty"`          // Meraki cloud capability (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	MerakiConnectStatus  string `json:"meraki-connect-status,omitempty"`   // Meraki cloud connection status (Live: IOS-XE 17.12.5)
	MerakiMonitorCapable bool   `json:"meraki-monitor-capable,omitempty"`  // Meraki monitoring capability (YANG: IOS-XE 17.18.1) (Live: IOS-XE 17.12.5)
	KernelCoredumpCount  uint16 `json:"kernel-coredump-count,omitempty"`   // Kernel coredump count (Live: IOS-XE 17.12.5)
	RegDomain            string `json:"reg-domain,omitempty"`              // Regulatory domain configuration (Live: IOS-XE 17.12.5)
	DartConStatus        string `json:"dart-con-status,omitempty"`         // DART connection status (Live: IOS-XE 17.12.5)
	ApAfcPreNotification bool   `json:"ap-afc-pre-notification,omitempty"` // AFC pre-notification capability (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	OobImgDwldMethod     string `json:"oob-img-dwld-method,omitempty"`     // Out-of-band image download method (Live: IOS-XE 17.12.5)
	WtpIP                string `json:"wtp-ip,omitempty"`                  // WTP IP address (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
}

// ApTimeInfo represents AP time related information.
type ApTimeInfo struct {
	BootTime      string `json:"boot-time"`       // Last AP reboot timestamp (Live: IOS-XE 17.12.5)
	JoinTime      string `json:"join-time"`       // AP join timestamp to controller (Live: IOS-XE 17.12.5)
	JoinTimeTaken uint32 `json:"join-time-taken"` // AP join process duration in seconds (Live: IOS-XE 17.12.5)
}

// ApSecurityData represents AP LSC (Local Significant Certificate) data.
type ApSecurityData struct {
	FipsEnabled      bool   `json:"fips-enabled"`        // FIPS compliance status (Live: IOS-XE 17.12.5)
	WlanccEnabled    bool   `json:"wlancc-enabled"`      // WLAN Common Criteria compliance (Live: IOS-XE 17.12.5)
	CertType         string `json:"cert-type"`           // AP certificate type (Live: IOS-XE 17.12.5)
	LscApAuthType    string `json:"lsc-ap-auth-type"`    // LSC authentication method (Live: IOS-XE 17.12.5)
	ApCertPolicy     string `json:"ap-cert-policy"`      // Certificate policy identifier (Live: IOS-XE 17.12.5)
	ApCertExpiryTime string `json:"ap-cert-expiry-time"` // AP certificate expiration timestamp (Live: IOS-XE 17.12.5)
	ApCertIssuerCn   string `json:"ap-cert-issuer-cn"`   // Certificate authority common name (Live: IOS-XE 17.12.5)
}

// SlidingWindow represents CAPWAP multiwindow transport information.
type SlidingWindow struct {
	MultiWindowSupport bool   `json:"multi-window-support"` // CAPWAP multiple window support (Live: IOS-XE 17.12.5)
	WindowSize         uint16 `json:"window-size"`          // CAPWAP sliding window size (Live: IOS-XE 17.12.5)
}

// ApVlan represents AP VLAN tagging details.
type ApVlan struct {
	VlanTagState string `json:"vlan-tag-state"` // AP VLAN tagging state (Live: IOS-XE 17.12.5)
	VlanTagID    uint16 `json:"vlan-tag-id"`    // 802.1Q VLAN identifier (Live: IOS-XE 17.12.5)
}

// HyperlocationData represents AP Hyperlocation details.
type HyperlocationData struct {
	HyperlocationMethod string `json:"hyperlocation-method"` // Hyperlocation positioning method (Live: IOS-XE 17.12.5)
	CmxIP               string `json:"cmx-ip,omitempty"`     // CMX server IP address (Live: IOS-XE 17.12.5)
}

// RebootStats represents AP reboot statistics.
type RebootStats struct {
	RebootReason string `json:"reboot-reason"` // Last AP reboot reason (Live: IOS-XE 17.12.5)
	RebootType   string `json:"reboot-type"`   // AP reboot type (Live: IOS-XE 17.12.5)
}

// ProxyInfo represents HTTP proxy configuration provisioned to AP.
type ProxyInfo struct {
	Hostname     string  `json:"hostname"`                // HTTP proxy server hostname or IP (Live: IOS-XE 17.12.5)
	Port         uint16  `json:"port"`                    // HTTP proxy server TCP port (Live: IOS-XE 17.12.5)
	NoProxyList  string  `json:"no-proxy-list"`           // URLs to bypass proxy (Live: IOS-XE 17.12.5)
	Username     *string `json:"username,omitempty"`      // HTTP proxy username (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	PasswordType *string `json:"password-type,omitempty"` // HTTP proxy password type (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	Password     *string `json:"password,omitempty"`      // HTTP proxy password (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
}

// Ipv4Pool represents DHCP IPv4 pool configuration.
type Ipv4Pool struct {
	Network   string `json:"network"`    // DHCP network address range (Live: IOS-XE 17.12.5)
	LeaseTime uint16 `json:"lease-time"` // DHCP lease duration in days (Live: IOS-XE 17.12.5)
	Netmask   string `json:"netmask"`    // IPv4 subnet mask (Live: IOS-XE 17.12.5)
}

// DisconnectDetail represents AP disconnect detail.
type DisconnectDetail struct {
	DisconnectReason string `json:"disconnect-reason"` // AP last disconnection reason (Live: IOS-XE 17.12.5)
}

// StatsMonitor represents AP statistics monitoring configuration.
type StatsMonitor struct {
	ActionApReload bool `json:"action-ap-reload"` // Auto AP reload on critical thresholds (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
}

// ApLscStatus represents AP LSC (Local Significant Certificate) status information.
type ApLscStatus struct {
	IsDtlsLscEnabled      bool   `json:"is-dtls-lsc-enabled"`                 // LSC enablement for CAPWAP DTLS (Live: IOS-XE 17.12.5)
	IsDot1xLscEnabled     bool   `json:"is-dot1x-lsc-enabled"`                // LSC enablement for 802.1X (Live: IOS-XE 17.12.5)
	IsDtlsLscFallback     bool   `json:"is-dtls-lsc-fallback"`                // AP fallback to default certificate (Live: IOS-XE 17.12.5)
	DtlsLscIssuerHash     string `json:"dtls-lsc-issuer-hash,omitempty"`      // CA hash for CAPWAP DTLS (Live: IOS-XE 17.12.5)
	Dot1xLscIssuerHash    string `json:"dot1x-lsc-issuer-hash,omitempty"`     // CA hash for 802.1X authentication (Live: IOS-XE 17.12.5)
	DtlsLscCertExpiryTime string `json:"dtls-lsc-cert-expiry-time,omitempty"` // DTLS LSC certificate expiration (Live: IOS-XE 17.12.5)
}

// RadioStatsMonitor represents AP radio statistics monitoring configuration.
type RadioStatsMonitor struct {
	Enable       bool          `json:"enable"`        // Radio statistics collection enable (Live: IOS-XE 17.12.5)
	SampleIntvl  uint16        `json:"sample-intvl"`  // Sampling interval in seconds (Live: IOS-XE 17.12.5)
	AlarmsEnable []interface{} `json:"alarms-enable"` // Statistics alarm enablement (YANG: IOS-XE 17.12.1) (Live: IOS-XE 17.12.5)
	RadioReset   bool          `json:"radio-reset"`   // Auto radio reset on stuck condition (Live: IOS-XE 17.12.5)
}

// ZeroWtDfs represents Zero wait DFS information of the AP.
type ZeroWtDfs struct {
	ReserveChannel ReserveChannel `json:"reserve-channel"` // DFS channel reservation data (Live: IOS-XE 17.12.5)
	Type           string         `json:"type"`            // CAC domain type classification (Live: IOS-XE 17.12.5)
	// DfsChanInclList and DfsChanExclList would be added if present in JSON
}

// ReserveChannel represents reserved CAC channel information.
type ReserveChannel struct {
	Channel      uint8  `json:"channel"`       // Reserved channel number (Live: IOS-XE 17.12.5)
	ChannelWidth string `json:"channel-width"` // Channel width for reserved CAC (Live: IOS-XE 17.12.5)
	State        string `json:"state"`         // CAC state for reserved channel (Live: IOS-XE 17.12.5)
}

// GnssInfo represents AP GNSS (Global Navigation Satellite System) information.
type GnssInfo struct {
	AntType          string  `json:"ant-type"`             // GNSS antenna type (Live: IOS-XE 17.12.5)
	AntCableLength   uint16  `json:"ant-cable-length"`     // GNSS antenna cable length in meters (Live: IOS-XE 17.12.5)
	AntennaProductID string  `json:"antenna-product-id"`   // GNSS antenna product identifier (Live: IOS-XE 17.12.5)
	AntennaSn        *string `json:"antenna-sn,omitempty"` // GNSS antenna serial number (YANG: IOS-XE 17.18.1) (Live: IOS-XE 17.12.5)
}

// ApState represents AP state information.
type ApState struct {
	ApAdminState     string `json:"ap-admin-state"`     // AP admin state (enabled/disabled/shutdown) (Live: IOS-XE 17.12.5)
	ApOperationState string `json:"ap-operation-state"` // AP operational state (Live: IOS-XE 17.12.5)
}

// ApModeData represents AP mode related information.
type ApModeData struct {
	HomeApEnabled bool         `json:"home-ap-enabled"` // Home AP feature enablement (Live: IOS-XE 17.12.5)
	ClearMode     bool         `json:"clear-mode"`      // Clear mode status (Live: IOS-XE 17.12.5)
	ApSubMode     string       `json:"ap-sub-mode"`     // AP operational sub-mode (Live: IOS-XE 17.12.5)
	WtpMode       string       `json:"wtp-mode"`        // WTP mode (local/flexconnect/monitor) (Live: IOS-XE 17.12.5)
	ApFabricData  ApFabricData `json:"ap-fabric-data"`  // SDA fabric integration attributes (Live: IOS-XE 17.12.5)
	// Ap6GhzData will be added for 17.18.1+ compatibility
}

// ApFabricData represents AP fabric related attributes.
type ApFabricData struct {
	IsFabricAp bool `json:"is-fabric-ap"` // SDA fabric-enabled AP designation (Live: IOS-XE 17.12.5)
}

// ApLocation represents AP location information.
type ApLocation struct {
	Floor             int         `json:"floor"`              // Physical floor number (Live: IOS-XE 17.12.5)
	Location          string      `json:"location"`           // AP physical placement description (Live: IOS-XE 17.12.5)
	AaaLocation       AaaLocation `json:"aaa-location"`       // AAA server location parameters (Live: IOS-XE 17.12.5)
	FloorID           int         `json:"floor-id"`           // Floor identifier for location services (Live: IOS-XE 17.12.5)
	RangingCapability int         `json:"ranging-capability"` // Location ranging capability level (Live: IOS-XE 17.12.5)
}

// AaaLocation represents AAA location information.
type AaaLocation struct {
	CivicID string `json:"civic-id"` // Civic location identifier (Live: IOS-XE 17.12.5)
	GeoID   string `json:"geo-id"`   // Geographic coordinate identifier (Live: IOS-XE 17.12.5)
	OperID  string `json:"oper-id"`  // Operator location identifier (Live: IOS-XE 17.12.5)
}

// ApServices represents AP services information.
type ApServices struct {
	MonitorModeOptType string       `json:"monitor-mode-opt-type"` // Monitor mode optimization type (Live: IOS-XE 17.12.5)
	ApDhcpServer       ApDhcpServer `json:"ap-dhcp-server"`        // Local DHCP server configuration (Live: IOS-XE 17.12.5)
	TotSnifferRadio    int          `json:"tot-sniffer-radio"`     // Total sniffer radio interfaces (Live: IOS-XE 17.12.5)
}

// ApDhcpServer represents AP DHCP server configuration.
type ApDhcpServer struct {
	IsDhcpServerEnabled bool `json:"is-dhcp-server-enabled"` // Local DHCP service enablement (Live: IOS-XE 17.12.5)
}

// XorPhyHtCap represents XOR PHY HT capabilities.
type XorPhyHtCap struct {
	Data XorPhyHtCapData `json:"data"` // XOR High Throughput capability data (Live: IOS-XE 17.12.5)
}

// XorPhyHtCapData represents XOR PHY HT capability data.
type XorPhyHtCapData struct {
	VhtCapable bool `json:"vht-capable"` // 802.11ac VHT capability support (Live: IOS-XE 17.12.5)
	HtCapable  bool `json:"ht-capable"`  // 802.11n HT capability support (Live: IOS-XE 17.12.5)
}

// XorPhyHeCap represents XOR PHY HE capabilities.
type XorPhyHeCap struct {
	Data XorPhyHeCapData `json:"data"` // XOR High Efficiency capability data (Live: IOS-XE 17.12.5)
}

// XorPhyHeCapData represents XOR PHY HE capability data.
type XorPhyHeCapData struct {
	HeEnabled              bool `json:"he-enabled"`                // 802.11ax HE protocol enablement (Live: IOS-XE 17.12.5)
	HeCapable              bool `json:"he-capable"`                // 802.11ax capability status (Live: IOS-XE 17.12.5)
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // 802.11ax single-user beamforming (Live: IOS-XE 17.12.5)
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // 802.11ax multi-user beamforming (Live: IOS-XE 17.12.5)
	HeStbcMode             int  `json:"he-stbc-mode"`              // 802.11ax STBC mode (Live: IOS-XE 17.12.5)
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // 802.11ax A-MPDU TID bitmap (Live: IOS-XE 17.12.5)
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // 802.11ax MCS and NSS capability (Live: IOS-XE 17.12.5)
}

// StationCfg represents station configuration.
type StationCfg struct {
	CfgData StationCfgData `json:"cfg-data"` // Station configuration parameters (Live: IOS-XE 17.12.5)
}

// StationCfgData represents station configuration data.
type StationCfgData struct {
	StationCfgConfigType string `json:"station-cfg-config-type"` // Station configuration type (Live: IOS-XE 17.12.5)
	MediumOccupancyLimit int    `json:"medium-occupancy-limit"`  // 802.11 medium occupancy limit (Live: IOS-XE 17.12.5)
	CfpPeriod            int    `json:"cfp-period"`              // CFP interval for PCF access control (Live: IOS-XE 17.12.5)
	CfpMaxDuration       int    `json:"cfp-max-duration"`        // Maximum CFP duration (Live: IOS-XE 17.12.5)
	Bssid                string `json:"bssid"`                   // BSS Identifier MAC address (Live: IOS-XE 17.12.5)
	BeaconPeriod         int    `json:"beacon-period"`           // 802.11 beacon transmission interval (Live: IOS-XE 17.12.5)
	CountryString        string `json:"country-string"`          // ISO country code string (Live: IOS-XE 17.12.5)
}

// MultiDomainCap represents multi-domain capability configuration.
type MultiDomainCap struct {
	CfgData MultiDomainCapData `json:"cfg-data"` // Multi-domain capability parameters (Live: IOS-XE 17.12.5)
}

// MultiDomainCapData represents multi-domain capability data.
type MultiDomainCapData struct {
	FirstChanNum    int `json:"first-chan-num"`     // First channel in regulatory domain (Live: IOS-XE 17.12.5)
	NumChannels     int `json:"num-channels"`       // Total channels in regulatory domain (Live: IOS-XE 17.12.5)
	MaxTxPowerLevel int `json:"max-tx-power-level"` // Maximum transmission power (dBm) (Live: IOS-XE 17.12.5)
}

// PhyHtCfg represents PHY HT configuration.
type PhyHtCfg struct {
	CfgData PhyHtCfgData `json:"cfg-data"` // High Throughput configuration parameters (Live: IOS-XE 17.12.5)
}

// PhyHtCfgData represents PHY HT configuration data.
type PhyHtCfgData struct {
	HtEnable               int    `json:"ht-enable"`                 // 802.11n HT protocol enablement (Live: IOS-XE 17.12.5)
	PhyHtCfgConfigType     string `json:"phy-ht-cfg-config-type"`    // Physical layer HT configuration type designation (Live: IOS-XE 17.12.5)
	CurrFreq               int    `json:"curr-freq"`                 // Current operating frequency (MHz) (Live: IOS-XE 17.12.5)
	ChanWidth              int    `json:"chan-width"`                // Channel bandwidth width in MHz (20/40/80/160) (Live: IOS-XE 17.12.5)
	ExtChan                int    `json:"ext-chan"`                  // Extension channel for 40MHz bonding (Live: IOS-XE 17.12.5)
	VhtEnable              bool   `json:"vht-enable"`                // 802.11ac Very High Throughput protocol enablement (Live: IOS-XE 17.12.5)
	LegTxBfEnabled         int    `json:"leg-tx-bf-enabled"`         // Legacy TX beamforming enablement (Live: IOS-XE 17.12.5)
	RrmChannelChangeReason string `json:"rrm-channel-change-reason"` // RRM channel change reason (Live: IOS-XE 17.12.5)
	FreqString             string `json:"freq-string"`               // Frequency designation string (Live: IOS-XE 17.12.5)
}

// PhyHtCapStruct represents PHY HT capability structure.
type PhyHtCapStruct struct {
	Data PhyHtCapStructData `json:"data"` // PHY HT capability information (Live: IOS-XE 17.12.5)
}

// PhyHtCapStructData represents PHY HT capability data.
type PhyHtCapStructData struct {
	VhtCapable bool `json:"vht-capable"` // 802.11ac capability status (Live: IOS-XE 17.12.5)
	HtCapable  bool `json:"ht-capable"`  // 802.11n capability status (Live: IOS-XE 17.12.5)
}

// PhyHeCapStruct represents PHY HE capability structure.
type PhyHeCapStruct struct {
	Data PhyHeCapStructData `json:"data"` // High Efficiency capability information (Live: IOS-XE 17.12.5)
}

// PhyHeCapStructData represents PHY HE capability data.
type PhyHeCapStructData struct {
	HeEnabled              bool `json:"he-enabled"`                // 802.11ax operational enablement status (Live: IOS-XE 17.12.5)
	HeCapable              bool `json:"he-capable"`                // 802.11ax capability status (Live: IOS-XE 17.12.5)
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // 802.11ax single-user beamforming (Live: IOS-XE 17.12.5)
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // 802.11ax multi-user beamforming (Live: IOS-XE 17.12.5)
	HeStbcMode             int  `json:"he-stbc-mode"`              // 802.11ax STBC operational mode (Live: IOS-XE 17.12.5)
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // 802.11ax A-MPDU TID bitmap (Live: IOS-XE 17.12.5)
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // 802.11ax MCS and NSS capabilities (Live: IOS-XE 17.12.5)
}

// ChanPwrInfo represents channel power information.
type ChanPwrInfo struct {
	Data ChanPwrInfoData `json:"data"` // Channel-specific power data (Live: IOS-XE 17.12.5)
}

// ChanPwrInfoData represents channel power information data.
type ChanPwrInfoData struct {
	AntennaGain    int         `json:"antenna-gain"`     // Antenna gain value in dBi (Live: IOS-XE 17.12.5)
	IntAntennaGain int         `json:"int-antenna-gain"` // Internal antenna gain in dBi (Live: IOS-XE 17.12.5)
	ExtAntennaGain int         `json:"ext-antenna-gain"` // External antenna gain in dBi (Live: IOS-XE 17.12.5)
	ChanPwrList    ChanPwrList `json:"chan-pwr-list"`    // Per-channel power level configuration (Live: IOS-XE 17.12.5)
}

// ChanPwrList represents channel power list.
type ChanPwrList struct {
	ChanPwr []ChanPwr `json:"chan-pwr"` // Channel power configurations array (Live: IOS-XE 17.12.5)
}

// ChanPwr represents individual channel power.
type ChanPwr struct {
	Chan int `json:"chan"` // Channel number for power assignment (Live: IOS-XE 17.12.5)
}

// SnifferCfg represents sniffer configuration.
type SnifferCfg struct {
	SnifferEnabled bool `json:"sniffer-enabled"` // Packet capture functionality enablement (Live: IOS-XE 17.12.5)
}

// RadioBandInfo represents radio band information.
type RadioBandInfo struct {
	BandID                 uint8          `json:"band-id"`                      // RF band identifier (2.4/5/6GHz) (Live: IOS-XE 17.12.5)
	RegDomainCode          uint16         `json:"reg-domain-code"`              // Regulatory domain code (Live: IOS-XE 17.12.5)
	RegulatoryDomain       string         `json:"regulatory-domain"`            // Regulatory domain name (Live: IOS-XE 17.12.5)
	MacOperCfg             MacOperCfg     `json:"mac-oper-cfg,omitempty"`       // MAC layer operational configuration (Live: IOS-XE 17.12.5)
	PhyTxPwrCfg            PhyTxPwrCfg    `json:"phy-tx-pwr-cfg,omitempty"`     // PHY layer TX power configuration (Live: IOS-XE 17.12.5)
	PhyTxPwrLvlCfg         PhyTxPwrLvlCfg `json:"phy-tx-pwr-lvl-cfg,omitempty"` // Multi-level TX power configuration (Live: IOS-XE 17.12.5)
	AntennaCfg             AntennaCfg     `json:"antenna-cfg,omitempty"`        // Antenna system configuration (Live: IOS-XE 17.12.5)
	Dot11acChannelWidthCap uint8          `json:"dot11ac-channel-width-cap"`    // 802.11ac max channel width capability (Live: IOS-XE 17.12.5)
	Secondary80Channel     uint16         `json:"secondary-80-channel"`         // Secondary 80MHz channel for 160MHz VHT (Live: IOS-XE 17.12.5)
	SiaParams              SiaParams      `json:"sia-params,omitempty"`         // Self-Identifying Antenna parameters (Live: IOS-XE 17.12.5)
}

// MacOperCfg represents MAC operation configuration.
type MacOperCfg struct {
	CfgData MacOperCfgData `json:"cfg-data"` // MAC layer operational data (Live: IOS-XE 17.12.5)
}

// MacOperCfgData represents MAC operation configuration data.
type MacOperCfgData struct {
	MacOperationConfigType string `json:"mac-operation-config-type"` // MAC operation configuration type (Live: IOS-XE 17.12.5)
	RtsThreshold           uint16 `json:"rts-threshold"`             // RTS threshold in bytes (Live: IOS-XE 17.12.5)
	ShortRetryLimit        uint8  `json:"short-retry-limit"`         // Max retry attempts for short frames (Live: IOS-XE 17.12.5)
	LongRetryLimit         uint8  `json:"long-retry-limit"`          // Max retry attempts for long frames (Live: IOS-XE 17.12.5)
	FragThreshold          uint16 `json:"frag-threshold"`            // Frame fragmentation threshold (Live: IOS-XE 17.12.5)
	MaxTxLifeTime          uint16 `json:"max-tx-life-time"`          // Maximum frame TX lifetime (Live: IOS-XE 17.12.5)
	MaxRxLifeTime          uint16 `json:"max-rx-life-time"`          // Max frame RX lifetime (Live: IOS-XE 17.12.5)
}

// PhyTxPwrCfg represents PHY TX power configuration.
type PhyTxPwrCfg struct {
	CfgData PhyTxPwrCfgData `json:"cfg-data"` // PHY layer TX power configuration (Live: IOS-XE 17.12.5)
}

// PhyTxPwrCfgData represents PHY TX power configuration data.
type PhyTxPwrCfgData struct {
	PhyTxPowerConfigType string `json:"phy-tx-power-config-type"` // PHY TX power configuration type (Live: IOS-XE 17.12.5)
	CurrentTxPowerLevel  uint8  `json:"current-tx-power-level"`   // Current TX power level index (Live: IOS-XE 17.12.5)
}

// PhyTxPwrLvlCfg represents PHY TX power level configuration.
type PhyTxPwrLvlCfg struct {
	CfgData PhyTxPwrLvlCfgData `json:"cfg-data"` // Multi-level TX power configuration (Live: IOS-XE 17.12.5)
}

// PhyTxPwrLvlCfgData represents PHY TX power level configuration data.
type PhyTxPwrLvlCfgData struct {
	NumSuppPowerLevels uint8 `json:"num-supp-power-levels"` // Number of supported power levels (Live: IOS-XE 17.12.5)
	TxPowerLevel1      int8  `json:"tx-power-level-1"`      // TX power level 1 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel2      int8  `json:"tx-power-level-2"`      // TX power level 2 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel3      int8  `json:"tx-power-level-3"`      // TX power level 3 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel4      int8  `json:"tx-power-level-4"`      // TX power level 4 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel5      int8  `json:"tx-power-level-5"`      // TX power level 5 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel6      int8  `json:"tx-power-level-6"`      // TX power level 6 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel7      int8  `json:"tx-power-level-7"`      // TX power level 7 (dBm) (Live: IOS-XE 17.12.5)
	TxPowerLevel8      int8  `json:"tx-power-level-8"`      // TX power level 8 (dBm) (Live: IOS-XE 17.12.5)
	CurrTxPowerInDbm   int8  `json:"curr-tx-power-in-dbm"`  // Current active TX power (dBm) (Live: IOS-XE 17.12.5)
}

// AntennaCfg represents antenna configuration.
type AntennaCfg struct {
	CfgData AntennaCfgData `json:"cfg-data"` // Antenna system configuration data (Live: IOS-XE 17.12.5)
}

// AntennaCfgData represents antenna configuration data.
type AntennaCfgData struct {
	DiversitySelection string `json:"diversity-selection"` // Antenna diversity selection algorithm (Live: IOS-XE 17.12.5)
	AntennaMode        string `json:"antenna-mode"`        // Antenna operational mode (Live: IOS-XE 17.12.5)
	NumOfAntennas      uint8  `json:"num-of-antennas"`     // Number of physical antenna elements (Live: IOS-XE 17.12.5)
}

// SiaParams represents Self Identifying Antenna parameters.
type SiaParams struct {
	IsRptncPresent bool   `json:"is-rptnc-present"` // Reverse Polarity TNC connector presence (Live: IOS-XE 17.12.5)
	IsDartPresent  bool   `json:"is-dart-present"`  // DART technology presence (Live: IOS-XE 17.12.5)
	AntennaIfType  string `json:"antenna-if-type"`  // Antenna interface type (Live: IOS-XE 17.12.5)
	AntennaGain    uint8  `json:"antenna-gain"`     // Antenna gain value (dBi) (Live: IOS-XE 17.12.5)
	Marlin4Present bool   `json:"marlin4-present"`  // Marlin4 antenna module presence (Live: IOS-XE 17.12.5)
	DmServType     string `json:"dm-serv-type"`     // Device management service type (Live: IOS-XE 17.12.5)
}

// VapOperConfig represents VAP operational configuration.
type VapOperConfig struct {
	ApVapID         uint8  `json:"ap-vap-id"`         // Virtual Access Point identifier (Live: IOS-XE 17.12.5)
	WlanID          uint8  `json:"wlan-id"`           // Wireless LAN identifier (Live: IOS-XE 17.12.5)
	BssidMac        string `json:"bssid-mac"`         // BSS Identifier MAC address (Live: IOS-XE 17.12.5)
	WtpMac          string `json:"wtp-mac"`           // Wireless Termination Point MAC address (Live: IOS-XE 17.12.5)
	WlanProfileName string `json:"wlan-profile-name"` // WLAN profile name (Live: IOS-XE 17.12.5)
	SSID            string `json:"ssid"`              // Service Set Identifier (Live: IOS-XE 17.12.5)
}

// DualRadioModeCfg represents dual radio mode configuration.
type DualRadioModeCfg struct {
	DualRadioMode    string `json:"dual-radio-mode"`    // Dual radio operational mode (Live: IOS-XE 17.12.5)
	DualRadioCapable string `json:"dual-radio-capable"` // Dual radio hardware capability (Live: IOS-XE 17.12.5)
	DualRadioModeOp  string `json:"dual-radio-mode-op"` // Dual radio operational state (Live: IOS-XE 17.12.5)
}

// BssColorCfg represents BSS color configuration.
type BssColorCfg struct {
	BssColorCapable    bool   `json:"bss-color-capable"`     // 802.11ax BSS color capability (Live: IOS-XE 17.12.5)
	BssColor           uint8  `json:"bss-color"`             // 802.11ax BSS color ID (1-63) (Live: IOS-XE 17.12.5)
	BssColorConfigType string `json:"bss-color-config-type"` // BSS color configuration type (Live: IOS-XE 17.12.5)
}

// BoardDataOpt represents board data options.
type BoardDataOpt struct {
	JoinPriority uint8 `json:"join-priority"` // Controller join priority value (Live: IOS-XE 17.12.5)
}

// DescriptorData represents descriptor data.
type DescriptorData struct {
	RadioSlotsInUse        uint8 `json:"radio-slots-in-use"`      // Number of active radio slots (Live: IOS-XE 17.12.5)
	EncryptionCapabilities bool  `json:"encryption-capabilities"` // Hardware encryption capability (Live: IOS-XE 17.12.5)
}

// ApProv represents AP provisioning information.
type ApProv struct {
	IsUniversal          bool   `json:"is-universal"`           // Universal AP provisioning status (Live: IOS-XE 17.12.5)
	UniversalPrimeStatus string `json:"universal-prime-status"` // Universal Prime licensing status (Live: IOS-XE 17.12.5)
}

// ApModels represents AP model information.
type ApModels struct {
	Model string `json:"model"` // Access point hardware model (Live: IOS-XE 17.12.5)
}

// TempInfo represents temperature information.
type TempInfo struct {
	Degree       int    `json:"degree"`        // Temperature in degrees Celsius (Live: IOS-XE 17.12.5)
	TempStatus   string `json:"temp-status"`   // Temperature operational status (Live: IOS-XE 17.12.5)
	HeaterStatus string `json:"heater-status"` // Internal heater operational status (Live: IOS-XE 17.12.5)
}

// TagInfo represents AP tag information.
type TagInfo struct {
	TagSource         string          `json:"tag-source"`          // Tag assignment source methodology (Live: IOS-XE 17.12.5)
	IsApMisconfigured bool            `json:"is-ap-misconfigured"` // AP misconfiguration detection (Live: IOS-XE 17.12.5)
	ResolvedTagInfo   ResolvedTagInfo `json:"resolved-tag-info"`   // Final resolved tag assignments (Live: IOS-XE 17.12.5)
	PolicyTagInfo     PolicyTagInfo   `json:"policy-tag-info"`     // Policy tag configuration (Live: IOS-XE 17.12.5)
	SiteTag           SiteTag         `json:"site-tag"`            // Site tag information (Live: IOS-XE 17.12.5)
	RfTag             RfTag           `json:"rf-tag"`              // RF tag configuration (Live: IOS-XE 17.12.5)
	FilterInfo        FilterInfo      `json:"filter-info"`         // Access control filter information (Live: IOS-XE 17.12.5)
	IsDtlsLscFbkAp    bool            `json:"is-dtls-lsc-fbk-ap"`  // DTLS LSC fallback AP designation (Live: IOS-XE 17.12.5)
}

// ResolvedTagInfo represents resolved tag information.
type ResolvedTagInfo struct {
	ResolvedPolicyTag string `json:"resolved-policy-tag"` // Final resolved policy tag name (Live: IOS-XE 17.12.5)
	ResolvedSiteTag   string `json:"resolved-site-tag"`   // Final resolved site tag name (Live: IOS-XE 17.12.5)
	ResolvedRfTag     string `json:"resolved-rf-tag"`     // Final resolved RF tag name (Live: IOS-XE 17.12.5)
}

// PolicyTagInfo represents policy tag information.
type PolicyTagInfo struct {
	PolicyTagName string `json:"policy-tag-name"` // Policy tag name identifier (Live: IOS-XE 17.12.5)
}

// SiteTag represents site tag information.
type SiteTag struct {
	SiteTagName string `json:"site-tag-name"` // Site tag name identifier (Live: IOS-XE 17.12.5)
	ApProfile   string `json:"ap-profile"`    // AP profile name (Live: IOS-XE 17.12.5)
	FlexProfile string `json:"flex-profile"`  // FlexConnect profile name (Live: IOS-XE 17.12.5)
}

// RfTag represents RF tag information.
type RfTag struct {
	RfTagName string `json:"rf-tag-name"` // RF tag name identifier (Live: IOS-XE 17.12.5)
}

// FilterInfo represents filter information.
type FilterInfo struct {
	FilterName string `json:"filter-name"` // Access control filter name (Live: IOS-XE 17.12.5)
}

// Tunnel represents tunnel configuration.
type Tunnel struct {
	PreferredMode string `json:"preferred-mode"` // Preferred CAPWAP tunnel mode (Live: IOS-XE 17.12.5)
	UDPLite       string `json:"udp-lite"`       // UDP-Lite protocol configuration (Live: IOS-XE 17.12.5)
}

// ExternalModuleData represents external module data.
type ExternalModuleData struct {
	XmData             XmData  `json:"xm-data"`               // External module data (Live: IOS-XE 17.12.5)
	UsbData            UsbData `json:"usb-data"`              // USB module data (Live: IOS-XE 17.12.5)
	UsbOverride        bool    `json:"usb-override"`          // USB configuration override (Live: IOS-XE 17.12.5)
	IsExtModuleEnabled bool    `json:"is-ext-module-enabled"` // External module enablement status (Live: IOS-XE 17.12.5)
}

// XmData represents XM module data.
type XmData struct {
	IsModulePresent bool `json:"is-module-present"` // External module presence detection (Live: IOS-XE 17.12.5)
	Xm              Xm   `json:"xm"`                // External module detailed information (Live: IOS-XE 17.12.5)
}

// UsbData represents USB module data.
type UsbData struct {
	IsModulePresent bool `json:"is-module-present"` // USB module presence detection (Live: IOS-XE 17.12.5)
	Xm              Xm   `json:"xm"`                // USB module detailed information (Live: IOS-XE 17.12.5)
}

// Xm represents external module information.
type Xm struct {
	NumericID          uint32 `json:"numeric-id"`           // Module unique numeric identifier (Live: IOS-XE 17.12.5)
	MaxPower           uint16 `json:"max-power"`            // Module maximum power consumption (mW) (Live: IOS-XE 17.12.5)
	SerialNumberString string `json:"serial-number-string"` // Module serial number (Live: IOS-XE 17.12.5)
	ProductIDString    string `json:"product-id-string"`    // Module product identifier (Live: IOS-XE 17.12.5)
	ModuleType         string `json:"module-type"`          // Module type classification (Live: IOS-XE 17.12.5)
	ModuleDescription  string `json:"module-description"`   // Module description (Live: IOS-XE 17.12.5)
}

// DeviceDetail represents device detail information.
type DeviceDetail struct {
	StaticInfo  StaticInfo  `json:"static-info"`  // Static hardware and firmware information (Live: IOS-XE 17.12.5)
	DynamicInfo DynamicInfo `json:"dynamic-info"` // Dynamic operational status (Live: IOS-XE 17.12.5)
	WtpVersion  WtpVersion  `json:"wtp-version"`  // WTP software version details (Live: IOS-XE 17.12.5)
}

// StaticInfo represents static information.
type StaticInfo struct {
	BoardData struct {
		WtpSerialNum string `json:"wtp-serial-num"` // AP serial number (Live: IOS-XE 17.12.5)
		WtpEnetMac   string `json:"wtp-enet-mac"`   // AP Ethernet MAC address (Live: IOS-XE 17.12.5)
		ApSysInfo    struct {
			MemType string `json:"mem-type"` // AP memory type (Live: IOS-XE 17.12.5)
			CPUType string `json:"cpu-type"` // AP CPU type (Live: IOS-XE 17.12.5)
			MemSize int    `json:"mem-size"` // AP memory size (Live: IOS-XE 17.12.5)
		} `json:"ap-sys-info"` // AP system info (Live: IOS-XE 17.12.5)
	} `json:"board-data"` // AP Board Data (Live: IOS-XE 17.12.5)
	BoardDataOpt   BoardDataOpt   `json:"board-data-opt,omitempty"`  // AP Additional Board data option (Live: IOS-XE 17.12.5)
	DescriptorData DescriptorData `json:"descriptor-data,omitempty"` // AP FW,HW information (Live: IOS-XE 17.12.5)
	ApProv         ApProv         `json:"ap-prov,omitempty"`         // AP universal provision (Live: IOS-XE 17.12.5)
	ApModels       ApModels       `json:"ap-models,omitempty"`       // AP device model (Live: IOS-XE 17.12.5)
	NumPorts       uint8          `json:"num-ports,omitempty"`       // Number of ports on AP (Live: IOS-XE 17.12.5)
	NumSlots       uint8          `json:"num-slots,omitempty"`       // Number of slots present in the access point (Live: IOS-XE 17.12.5)
	WtpModelType   uint16         `json:"wtp-model-type,omitempty"`  // AP model type (Live: IOS-XE 17.12.5)
	ApCapability   string         `json:"ap-capability,omitempty"`   // AP capabilities (Live: IOS-XE 17.12.5)
	IsMmOpt        bool           `json:"is-mm-opt,omitempty"`       // AP monitor mode optimization support (Live: IOS-XE 17.12.5)
	ApImageName    string         `json:"ap-image-name,omitempty"`   // AP Software image name (Live: IOS-XE 17.12.5)
}

// DynamicInfo represents dynamic information.
type DynamicInfo struct {
	ApCrashData struct {
		ApCrashFile           string `json:"ap-crash-file"`              // AP crash file (Live: IOS-XE 17.12.5)
		ApRadio2GCrashFile    string `json:"ap-radio-2g-crash-file"`     // AP 2 GHz radio crash file (Live: IOS-XE 17.12.5)
		ApRadio5GCrashFile    string `json:"ap-radio-5g-crash-file"`     // AP 5 GHz radio crash file (Live: IOS-XE 17.12.5)
		ApRadio6GCrashFile    string `json:"ap-radio-6g-crash-file"`     // AP 6 GHz radio crash file (Live: IOS-XE 17.12.5)
		ApRad5GSlot2CrashFile string `json:"ap-rad-5g-slot2-crash-file"` // AP 5 GHz radio slot 2 crash file (Live: IOS-XE 17.12.5)
	} `json:"ap-crash-data"` // AP crash data (Live: IOS-XE 17.12.5)
	LedStateEnabled  bool     `json:"led-state-enabled,omitempty"`  // True if LED state of AP is enabled (Live: IOS-XE 17.12.5)
	ResetButtonState bool     `json:"reset-button-state,omitempty"` // True if AP Reset button state is enabled (Live: IOS-XE 17.12.5)
	LedFlashEnabled  bool     `json:"led-flash-enabled,omitempty"`  // True if LED Flash state of AP is enabled (Live: IOS-XE 17.12.5)
	FlashSec         uint16   `json:"flash-sec,omitempty"`          // LED Flash timer duration for AP in seconds (Live: IOS-XE 17.12.5)
	TempInfo         TempInfo `json:"temp-info,omitempty"`          // AP temperature info (Live: IOS-XE 17.12.5)
	LedFlashExpiry   string   `json:"led-flash-expiry,omitempty"`   // Led Flash Expiry Date and Time (Live: IOS-XE 17.12.5)
}

// WtpVersion represents WTP version information.
type WtpVersion struct {
	BackupSwVersion struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.5)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.5)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.5)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.5)
	} `json:"backup-sw-version"` // Backup software version of the AP (Live: IOS-XE 17.12.5)
	MiniIosVersion struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.5)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.5)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.5)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.5)
	} `json:"mini-ios-version,omitempty"` // Cisco AP mini IOS version details (Live: IOS-XE 17.12.5)
	SwVer struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.5)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.5)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.5)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.5)
	} `json:"sw-ver,omitempty"` // Software version of the AP (Live: IOS-XE 17.12.5)
	BootVer struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.5)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.5)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.5)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.5)
	} `json:"boot-ver,omitempty"` // Cisco AP boot version details (Live: IOS-XE 17.12.5)
	SwVersion string `json:"sw-version,omitempty"` // Cisco AP software version details (Live: IOS-XE 17.12.5)
}

// ApNameMacMap represents AP name to MAC address mapping.
type ApNameMacMap struct {
	WtpName string `json:"wtp-name"` // WTP administrative name (Live: IOS-XE 17.12.5)
	WtpMac  string `json:"wtp-mac"`  // WTP radio interface MAC address (Live: IOS-XE 17.12.5)
	EthMac  string `json:"eth-mac"`  // Ethernet interface MAC address (Live: IOS-XE 17.12.5)
}

// WtpSlotWlanStats represents WTP slot WLAN statistics.
type WtpSlotWlanStats struct {
	WtpMac      string `json:"wtp-mac"`      // WTP MAC address for radio interface (Live: IOS-XE 17.12.5)
	SlotID      int    `json:"slot-id"`      // Radio slot identifier (Live: IOS-XE 17.12.5)
	WlanID      int    `json:"wlan-id"`      // WLAN identifier (Live: IOS-XE 17.12.5)
	BssidMac    string `json:"bssid-mac"`    // BSS Identifier MAC address (Live: IOS-XE 17.12.5)
	Ssid        string `json:"ssid"`         // Service Set Identifier name (Live: IOS-XE 17.12.5)
	BytesRx     string `json:"bytes-rx"`     // Total bytes received on WLAN interface (Live: IOS-XE 17.12.5)
	BytesTx     string `json:"bytes-tx"`     // Total bytes transmitted on WLAN interface (Live: IOS-XE 17.12.5)
	PktsRx      string `json:"pkts-rx"`      // Total packets received on WLAN interface (Live: IOS-XE 17.12.5)
	PktsTx      string `json:"pkts-tx"`      // Total packets transmitted on WLAN interface (Live: IOS-XE 17.12.5)
	DataRetries string `json:"data-retries"` // Data frame retransmission count (Live: IOS-XE 17.12.5)
}

// EthernetMacWtpMacMap represents Ethernet MAC to WTP MAC mapping.
type EthernetMacWtpMacMap struct {
	EthernetMac string `json:"ethernet-mac"` // Ethernet interface MAC address (Live: IOS-XE 17.12.5)
	WtpMac      string `json:"wtp-mac"`      // WTP MAC address for radio interface (Live: IOS-XE 17.12.5)
}

// ApIoxOperData represents AP IOx operational data.
type ApIoxOperData struct {
	ApMac        string `json:"ap-mac"`        // Access point MAC address (Live: IOS-XE 17.12.5)
	ApphostState string `json:"apphost-state"` // Application hosting service state (Live: IOS-XE 17.12.5)
	CafToken     string `json:"caf-token"`     // CAF authentication token (Live: IOS-XE 17.12.5)
	CafPort      int    `json:"caf-port"`      // CAF service communication port (Live: IOS-XE 17.12.5)
}

// QosGlobalStats represents QoS global statistics.
type QosGlobalStats struct {
	QosClientVoiceStats struct {
		TotalNumOfTspecRcvd       int `json:"total-num-of-tspec-rcvd"`        // Total Number of TSPEC requests received (Live: IOS-XE 17.12.5)
		NewTspecFromAssocReq      int `json:"new-tspec-from-assoc-req"`       // Number of New TSPEC received from Assoc Request (Live: IOS-XE 17.12.5)
		TspecRenewalFromAssocReq  int `json:"tspec-renewal-from-assoc-req"`   // TSPEC renewal from Assoc Request (Live: IOS-XE 17.12.5)
		NewTspecAsAddTS           int `json:"new-tspec-as-add-ts"`            // Number of new Add TS requests received (Live: IOS-XE 17.12.5)
		TspecRenewalFromAddTS     int `json:"tspec-renewal-from-add-ts"`      // Number of Add TS renewal requests received (Live: IOS-XE 17.12.5)
		NumOfActiveTspecCalls     int `json:"num-of-active-tspec-calls"`      // Total Number of active TSPEC calls (Live: IOS-XE 17.12.5)
		NumOfActiveSIPCalls       int `json:"num-of-active-sip-calls"`        // Total Number of active SIP calls (Live: IOS-XE 17.12.5)
		NumOfCallsAccepted        int `json:"num-of-calls-accepted"`          // Total Number of calls accepted (Live: IOS-XE 17.12.5)
		NumOfCallsRejectedInsufBw int `json:"num-of-calls-rejected-insuf-bw"` // Number of calls rejected due to Insufficient BW (Live: IOS-XE 17.12.5)
		NumOfCallsRejectedPhyRate int `json:"num-of-calls-rejected-phy-rate"` // Number of calls rejected due to PHY rate (Live: IOS-XE 17.12.5)
		NumOfCallsRejectedQos     int `json:"num-of-calls-rejected-qos"`      // Number of calls rejected due to QoS policy (Live: IOS-XE 17.12.5)
		NumOfCallsRejInvalidTspec int `json:"num-of-calls-rej-invalid-tspec"` // Number of calls rejected due to Invalid TSPEC (Live: IOS-XE 17.12.5)
		NumOfRoamCallsAccepted    int `json:"num-of-roam-calls-accepted"`     // Total Number of roam calls accepted (Live: IOS-XE 17.12.5)
		NumOfRoamCallsRejected    int `json:"num-of-roam-calls-rejected"`     // Total Number of roam calls rejected (Live: IOS-XE 17.12.5)
		TspecProcessFailedGetRec  int `json:"tspec-process-failed-get-rec"`   // Number of DB failures while processing TSPEC (Live: IOS-XE 17.12.5)
		TotalNumOfCallReport      int `json:"total-num-of-call-report"`       // Total number of call-report received (Live: IOS-XE 17.12.5)
		TotalSIPFailureTrapSend   int `json:"total-sip-failure-trap-send"`    // Total number of SIP failure trap send (Live: IOS-XE 17.12.5)
		TotalSIPInviteOnCaller    int `json:"total-sip-invite-on-caller"`     // Total number of SIP Invite received on Caller (Live: IOS-XE 17.12.5)
		TotalSIPInviteOnCallee    int `json:"total-sip-invite-on-callee"`     // Total number of SIP Invite received on Callee (Live: IOS-XE 17.12.5)
	} `json:"qos-client-voice-stats"`
}

// ApOperInternalData represents internal AP operational data.
type ApOperInternalData struct {
	WtpMac                 string         `json:"wtp-mac"`                             // MAC Address of the AP Radio (Live: IOS-XE 17.12.5)
	RadioID                int            `json:"radio-id"`                            // AP radio identifier (Live: IOS-XE 17.12.5)
	ApAntennaBandMode      string         `json:"ap-antenna-band-mode"`                // AP antenna band mode configuration (Live: IOS-XE 17.12.5)
	LinkEncryptionEnabled  bool           `json:"link-encryption-enabled"`             // Controller-AP link encryption status (Live: IOS-XE 17.12.5)
	ApRemoteDebugMode      bool           `json:"ap-remote-debug-mode"`                // Remote debugging status for the AP (Live: IOS-XE 17.12.5)
	ApRole                 string         `json:"ap-role"`                             // AP role in PMK push (Live: IOS-XE 17.12.5)
	ApIndoorMode           bool           `json:"ap-indoor-mode"`                      // Identifier for indoor AP mode (Live: IOS-XE 17.12.5)
	MaxClientsAllowed      int            `json:"max-clients-allowed"`                 // Maximum clients allowed on an AP (Live: IOS-XE 17.12.5)
	IsLocalNet             bool           `json:"is-local-net"`                        // Identifier for local access in OEAP AP (Live: IOS-XE 17.12.5)
	Ipv4TcpMss             TCPMssConfig   `json:"ipv4-tcp-mss"`                        // Configured IPv4 TCP MSS value for client (Live: IOS-XE 17.12.5)
	Ipv6TcpMss             TCPMssConfig   `json:"ipv6-tcp-mss"`                        // Configured IPv6 TCP MSS value for client (Live: IOS-XE 17.12.5)
	RangingMode            string         `json:"ranging-mode"`                        // Ranging mode - normal or accurate (Live: IOS-XE 17.12.5)
	PowerProfile           string         `json:"power-profile"`                       // Power profile applied to the AP (Live: IOS-XE 17.12.5)
	PwrProfType            string         `json:"pwr-prof-type"`                       // Power profile type (Live: IOS-XE 17.12.5)
	PwrCalProfile          string         `json:"pwr-cal-profile"`                     // Calendar profile associated to power profile (Live: IOS-XE 17.12.5)
	PersistentSsid         PersistentSsid `json:"persistent-ssid"`                     // Persistent SSID broadcast operation information (Live: IOS-XE 17.12.5)
	ProvSsid               bool           `json:"prov-ssid"`                           // Office Extended AP Provisional SSID status (Live: IOS-XE 17.12.5)
	PrimingProfile         string         `json:"priming-profile"`                     // Applied AP priming profile name (Live: IOS-XE 17.12.5)
	PrimingProfileSrc      string         `json:"priming-profile-src"`                 // AP priming profile configuration source (Live: IOS-XE 17.12.5)
	PrimingFilter          string         `json:"priming-filter"`                      // AP priming filter name (Live: IOS-XE 17.12.5)
	PmkBsSenderAddr        string         `json:"pmk-bs-sender-addr"`                  // PMK bulk sync sender AP MAC address (Live: IOS-XE 17.12.5)
	PmkBsReceiverAddr      string         `json:"pmk-bs-receiver-addr"`                // PMK bulk sync receiver AP MAC address (Live: IOS-XE 17.12.5)
	Accounting             interface{}    `json:"accounting,omitempty"`                // Accounting info to be sent to radius server (Live: IOS-XE 17.12.5)
	ApDnaData              interface{}    `json:"ap-dna-data,omitempty"`               // Cisco-DNA related data (Live: IOS-XE 17.12.5)
	ApGasRateLimitCfg      interface{}    `json:"ap-gas-rate-limit-cfg,omitempty"`     // Cisco AP Generic Advertisement Services (GAS) rate configuration (Live: IOS-XE 17.12.5)
	ApIPData               interface{}    `json:"ap-ip-data,omitempty"`                // AP IP address configuration (Live: IOS-XE 17.12.5)
	ApLoginCredentials     interface{}    `json:"ap-login-credentials,omitempty"`      // Login credentials configured on an AP (Live: IOS-XE 17.12.5)
	ApMgmt                 interface{}    `json:"ap-mgmt,omitempty"`                   // AP management data (Live: IOS-XE 17.12.5)
	ApNtpServerInfoCfg     interface{}    `json:"ap-ntp-server-info-cfg,omitempty"`    // NTP server information to be used by AP (Live: IOS-XE 17.12.5)
	ApNtpSyncStatus        interface{}    `json:"ap-ntp-sync-status,omitempty"`        // AP NTP synchronization status (Live: IOS-XE 17.12.5)
	ApPmkPropagationStatus interface{}    `json:"ap-pmk-propagation-status,omitempty"` // AP PMK push propagation status (Live: IOS-XE 17.12.5)
	ApPow                  interface{}    `json:"ap-pow,omitempty"`                    // AP power related data (Live: IOS-XE 17.12.5)
	ApPrimeInfo            interface{}    `json:"ap-prime-info,omitempty"`             // Controller configuration for the AP (Live: IOS-XE 17.12.5)
	ApPrimingOverride      interface{}    `json:"ap-priming-override,omitempty"`       // AP priming override flag status (Live: IOS-XE 17.12.5)
	ApSysStats             interface{}    `json:"ap-sys-stats,omitempty"`              // AP system statistics (Live: IOS-XE 17.12.5)
	ApTzConfig             interface{}    `json:"ap-tz-config,omitempty"`              // AP timezone configuration (Live: IOS-XE 17.12.5)
	ApUdpliteInfo          interface{}    `json:"ap-udplite-info,omitempty"`           // UDP-Lite operational information (Live: IOS-XE 17.12.5)
	AuxClientInterfaceData interface{}    `json:"aux-client-interface-data,omitempty"` // Auxiliary Client Interface data (Live: IOS-XE 17.12.5)
	InfrastructureMfp      interface{}    `json:"infrastructure-mfp,omitempty"`        // Cisco AP Management Frame Protection (Live: IOS-XE 17.12.5)
	KernelCoredump         interface{}    `json:"kernel-coredump,omitempty"`           // Kernel coredump configuration (Live: IOS-XE 17.12.5)
	LinkAudit              interface{}    `json:"link-audit,omitempty"`                // Link audit data (Live: IOS-XE 17.12.5)
	OeapAudit              interface{}    `json:"oeap-audit,omitempty"`                // On-demand Office Extended AP link test data (Live: IOS-XE 17.12.5)
	Retransmit             interface{}    `json:"retransmit,omitempty"`                // AP retransmission parameters (Live: IOS-XE 17.12.5)
	Syslog                 interface{}    `json:"syslog,omitempty"`                    // Cisco AP System Logging (Live: IOS-XE 17.12.5)
	Timer                  interface{}    `json:"timer,omitempty"`                     // Cisco access point timer data (Live: IOS-XE 17.12.5)
}

// RlanOper represents RLAN operational data.
type RlanOper struct {
	WtpMac         string `json:"wtp-mac"`          // Radio MAC address of the AP (Live: IOS-XE 17.12.5)
	RlanPortID     int    `json:"rlan-port-id"`     // RLAN port identifier (Live: IOS-XE 17.12.5)
	RlanOperState  bool   `json:"rlan-oper-state"`  // Status of the LAN port (Live: IOS-XE 17.12.5)
	RlanPortStatus bool   `json:"rlan-port-status"` // Remote LAN status of the LAN port (Live: IOS-XE 17.12.5)
	RlanVlanValid  bool   `json:"rlan-vlan-valid"`  // LAN port valid or not (Live: IOS-XE 17.12.5)
	RlanVlanID     int    `json:"rlan-vlan-id"`     // VLAN id of the LAN port (Live: IOS-XE 17.12.5)
	RlanPoeState   string `json:"rlan-poe-state"`   // PoE state of the LAN port (Live: IOS-XE 17.12.5)
	PowerLevelID   int    `json:"power-level-id"`   // Power level of the LAN port (Live: IOS-XE 17.12.5)
}

// EwlcMewlcPredownloadRec represents EWLC MEWLC predownload record.
type EwlcMewlcPredownloadRec struct {
	PredState                    string `json:"pred-state"`                     // Embedded Wireless Controller predownload state (Live: IOS-XE 17.12.5)
	MeCapableApCount             int    `json:"me-capable-ap-count"`            // Total EWC capable AP count (Live: IOS-XE 17.12.5)
	ControllerPredownloadVersion string `json:"controller-predownload-version"` // Embedded Wireless Controller predownload version (Live: IOS-XE 17.12.5)
}

// CdpCacheData represents CDP cache data.
type CdpCacheData struct {
	MacAddr                string       `json:"mac-addr"`                   // MAC address (Live: IOS-XE 17.12.5)
	CdpCacheDeviceID       string       `json:"cdp-cache-device-id"`        // CDP device identifier (Live: IOS-XE 17.12.5)
	ApName                 string       `json:"ap-name"`                    // AP Name (Live: IOS-XE 17.12.5)
	LastUpdatedTime        time.Time    `json:"last-updated-time"`          // Last updated time (Live: IOS-XE 17.12.5)
	Version                int          `json:"version"`                    // Cisco Discovery Protocol version (Live: IOS-XE 17.12.5)
	WtpMacAddr             string       `json:"wtp-mac-addr"`               // WTP MAC address (Live: IOS-XE 17.12.5)
	DeviceIndex            int          `json:"device-index"`               // Device index (Live: IOS-XE 17.12.5)
	IPAddress              CdpIPAddress `json:"ip-address"`                 // Device network addresses from CDP message (Live: IOS-XE 17.12.5)
	CdpAddrCount           int          `json:"cdp-addr-count"`             // Neighbor IP count (Live: IOS-XE 17.12.5)
	CdpCacheApAddress      string       `json:"cdp-cache-ap-address"`       // CDP cache address type for the AP (Live: IOS-XE 17.12.5)
	CdpCacheDevicePort     string       `json:"cdp-cache-device-port"`      // Device outgoing port (Live: IOS-XE 17.12.5)
	CdpCacheDuplex         string       `json:"cdp-cache-duplex"`           // CDP cache duplex type (Live: IOS-XE 17.12.5)
	CdpCacheIfIndex        int          `json:"cdp-cache-if-index"`         // CDP cache interface index (Live: IOS-XE 17.12.5)
	CdpCacheInterfaceSpeed int          `json:"cdp-cache-interface-speed"`  // CDP cache interface speed (Live: IOS-XE 17.12.5)
	CdpCacheIPAddressValue string       `json:"cdp-cache-ip-address-value"` // Entry address(es) (Live: IOS-XE 17.12.5)
	CdpCacheLocalPort      string       `json:"cdp-cache-local-port"`       // Device interface port (Live: IOS-XE 17.12.5)
	CdpCachePlatform       string       `json:"cdp-cache-platform"`         // CDP cache platform (Live: IOS-XE 17.12.5)
	CdpCacheVersion        string       `json:"cdp-cache-version"`          // CDP cache version (Live: IOS-XE 17.12.5)
	CdpCapabilitiesString  string       `json:"cdp-capabilities-string"`    // CDP cache capabilities (Live: IOS-XE 17.12.5)
}

// LldpNeigh represents LLDP neighbor information.
type LldpNeigh struct {
	WtpMac          string `json:"wtp-mac"`          // Radio MAC address of the AP (Live: IOS-XE 17.12.5)
	NeighMac        string `json:"neigh-mac"`        // MAC address of the LLDP neighbor device (Live: IOS-XE 17.12.5)
	PortID          string `json:"port-id"`          // LLDP neighbor port name or ID (Live: IOS-XE 17.12.5)
	LocalPort       string `json:"local-port"`       // AP interface sending/receiving LLDP PDUs (Live: IOS-XE 17.12.5)
	SystemName      string `json:"system-name"`      // LLDP neighbor name (Live: IOS-XE 17.12.5)
	PortDescription string `json:"port-description"` // LLDP neighbor port description (Live: IOS-XE 17.12.5)
	Capabilities    string `json:"capabilities"`     // LLDP device capabilities (Live: IOS-XE 17.12.5)
	MgmtAddr        string `json:"mgmt-addr"`        // Management IPv4 address of LLDP neighbor (Live: IOS-XE 17.12.5)
}

// TpCertInfo represents trustpoint certificate information.
type TpCertInfo struct {
	Trustpoint Trustpoint `json:"trustpoint"` // Trustpoint certificate information (Live: IOS-XE 17.12.5)
}

// Trustpoint represents trustpoint information.
type Trustpoint struct {
	TrustpointName     string  `json:"trustpoint-name"`      // Trustpoint name (Live: IOS-XE 17.12.5)
	IsCertAvailable    bool    `json:"is-cert-available"`    // Is certificate available (Live: IOS-XE 17.12.5)
	IsPrivkeyAvailable bool    `json:"is-privkey-available"` // Is private key available (Live: IOS-XE 17.12.5)
	CertHash           string  `json:"cert-hash"`            // Certificate hash (Live: IOS-XE 17.12.5)
	CertType           string  `json:"cert-type"`            // Certificate type (Live: IOS-XE 17.12.5)
	FipsSuitability    string  `json:"fips-suitability"`     // FIPS Suitability (Live: IOS-XE 17.12.5)
	State              *string `json:"state,omitempty"`      // Trustpoint state (Live: IOS-XE 17.12.5)
}

// DiscData represents discovery data.
type DiscData struct {
	WtpMac           string `json:"wtp-mac"`            // Wireless termination point MAC address (Live: IOS-XE 17.12.5)
	DiscoveryPkts    string `json:"discovery-pkts"`     // Discovery packet count (Live: IOS-XE 17.12.5)
	DiscoveryErrPkts string `json:"discovery-err-pkts"` // Discovery error packet count (Live: IOS-XE 17.12.5)
}

// CapwapPkts represents CAPWAP packet statistics.
type CapwapPkts struct {
	WtpMac            string `json:"wtp-mac"`              // Wireless termination point MAC address (Live: IOS-XE 17.12.5)
	CntrlPkts         string `json:"cntrl-pkts"`           // Control packet count (Live: IOS-XE 17.12.5)
	DataKeepAlivePkts string `json:"data-keep-alive-pkts"` // Data keep-alive packet count (Live: IOS-XE 17.12.5)
	CapwapErrorPkts   string `json:"capwap-error-pkts"`    // CAPWAP error packet count (Live: IOS-XE 17.12.5)
	ArpPkts           string `json:"arp-pkts"`             // ARP packet count (Live: IOS-XE 17.12.5)
	DhcpPkts          string `json:"dhcp-pkts"`            // DHCP packet count (Live: IOS-XE 17.12.5)
	Dot1xCtrlPkts     string `json:"dot1x-ctrl-pkts"`      // 802.1X control packet count (Live: IOS-XE 17.12.5)
	Dot1xEapPkts      string `json:"dot1x-eap-pkts"`       // 802.1X EAP packet count (Live: IOS-XE 17.12.5)
	Dot1xKeyTypePkts  string `json:"dot1x-key-type-pkts"`  // 802.1X key type packet count (Live: IOS-XE 17.12.5)
	Dot1xMgmtPkts     string `json:"dot1x-mgmt-pkts"`      // 802.1X management packet count (Live: IOS-XE 17.12.5)
	IappPkts          string `json:"iapp-pkts"`            // IAPP packet count (Live: IOS-XE 17.12.5)
	IPPkts            string `json:"ip-pkts"`              // IP packet count (Live: IOS-XE 17.12.5)
	Ipv6Pkts          string `json:"ipv6-pkts"`            // IPv6 packet count (Live: IOS-XE 17.12.5)
	RfidPkts          string `json:"rfid-pkts"`            // RFID packet count (Live: IOS-XE 17.12.5)
	RrmPkts           string `json:"rrm-pkts"`             // Radio resource management packet count (Live: IOS-XE 17.12.5)
}

// CountryOper represents country operational data.
type CountryOper struct {
	CountryCode         string      `json:"country-code"`                 // Country code for regulatory compliance (Live: IOS-XE 17.12.5)
	CountryString       string      `json:"country-string"`               // Country string representation (Live: IOS-XE 17.12.5)
	RegDomainStr80211Bg string      `json:"reg-domain-str-80211bg"`       // Regulatory domain string for 802.11bg (Live: IOS-XE 17.12.5)
	RegDomainStr80211A  string      `json:"reg-domain-str-80211a"`        // Regulatory domain string for 802.11a (Live: IOS-XE 17.12.5)
	CountrySupported    bool        `json:"country-supported"`            // Country support status in regulatory database (Live: IOS-XE 17.12.5)
	Channels11a         interface{} `json:"channels-11a,omitempty"`       // Available channels for 802.11a operation (Live: IOS-XE 17.12.5)
	Channels11bg        interface{} `json:"channels-11bg,omitempty"`      // Available channels for 802.11bg operation (Live: IOS-XE 17.12.5)
	ChannelsString11a   string      `json:"channels-string-11a"`          // Channel string representation for 802.11a (Live: IOS-XE 17.12.5)
	ChannelsString11bg  string      `json:"channels-string-11bg"`         // Channel string representation for 802.11bg (Live: IOS-XE 17.12.5)
	DcaChannels11a      interface{} `json:"dca-channels-11a,omitempty"`   // DCA channels for 802.11a band (Live: IOS-XE 17.12.5)
	DcaChannels11bg     interface{} `json:"dca-channels-11bg,omitempty"`  // DCA channels for 802.11bg band (Live: IOS-XE 17.12.5)
	RadarChannels11a    interface{} `json:"radar-channels-11a,omitempty"` // Radar-affected channels for 802.11a (Live: IOS-XE 17.12.5)
	RegDom6ghz          interface{} `json:"reg-dom-6ghz,omitempty"`       // Regulatory domain information for 6GHz (Live: IOS-XE 17.12.5)
	ChanInfo6ghz        interface{} `json:"chan-info-6ghz,omitempty"`     // Channel information for 6GHz band (Live: IOS-XE 17.12.5)
}

// SuppCountryOper represents supported country operational data.
type SuppCountryOper struct {
	CountryCode      string      `json:"country-code"`                  // Supported country code for regulatory compliance (Live: IOS-XE 17.12.5)
	CountryString    string      `json:"country-string"`                // Supported country string representation (Live: IOS-XE 17.12.5)
	CountryCodeIso   string      `json:"country-code-iso"`              // ISO standard country code (Live: IOS-XE 17.12.5)
	ChanList24ghz    interface{} `json:"chan-list-24ghz,omitempty"`     // Channel list for 2.4GHz band (Live: IOS-XE 17.12.5)
	ChanList5ghz     interface{} `json:"chan-list-5ghz,omitempty"`      // Channel list for 5GHz band (Live: IOS-XE 17.12.5)
	ChanList6ghz     interface{} `json:"chan-list-6ghz,omitempty"`      // Channel list for 6GHz band (Live: IOS-XE 17.12.5)
	ChanListDca24ghz interface{} `json:"chan-list-dca-24ghz,omitempty"` // DCA channel list for 2.4GHz band (Live: IOS-XE 17.12.5)
	ChanListDca5ghz  interface{} `json:"chan-list-dca-5ghz,omitempty"`  // DCA channel list for 5GHz band (Live: IOS-XE 17.12.5)
	ChanListDca6ghz  interface{} `json:"chan-list-dca-6ghz,omitempty"`  // DCA channel list for 6GHz band (Live: IOS-XE 17.12.5)
	ChanListPsc6ghz  interface{} `json:"chan-list-psc-6ghz,omitempty"`  // PSC channel list for 6GHz band (Live: IOS-XE 17.12.5)
	RegDom24ghz      interface{} `json:"reg-dom-24ghz,omitempty"`       // Regulatory domain for 2.4GHz band (Live: IOS-XE 17.12.5)
	RegDom5ghz       interface{} `json:"reg-dom-5ghz,omitempty"`        // Regulatory domain for 5GHz band (Live: IOS-XE 17.12.5)
	RegDom6ghz       interface{} `json:"reg-dom-6ghz,omitempty"`        // Regulatory domain for 6GHz band (Live: IOS-XE 17.12.5)
}

// ApNhGlobalData represents AP neighborhood global data.
type ApNhGlobalData struct {
	AlgorithmRunning   bool `json:"algorithm-running"`     // Status of neighborhood algorithm execution (Live: IOS-XE 17.12.5)
	AlgorithmItrCount  int  `json:"algorithm-itr-count"`   // Total AP neighborhood algorithm iteration count (Live: IOS-XE 17.12.5)
	IdealCapacityPerRg int  `json:"ideal-capacity-per-rg"` // Ideal capacity of APs per resource group (Live: IOS-XE 17.12.5)
	NumOfNeighborhood  int  `json:"num-of-neighborhood"`   // Total number of calculated neighborhood areas (Live: IOS-XE 17.12.5)
}

// ApImagePrepareLocation represents AP image prepare location.
type ApImagePrepareLocation struct {
	Index     int         `json:"index"`      // AP image index identifier for prepare location (Live: IOS-XE 17.12.5)
	ImageFile string      `json:"image-file"` // AP image file name for prepare location (Live: IOS-XE 17.12.5)
	ImageData []ImageData `json:"image-data"` // AP image info for prepare location (Live: IOS-XE 17.12.5)
}

// ImageData represents image data information.
type ImageData struct {
	ImageName     string   `json:"image-name"`     // AP image name identifier (Live: IOS-XE 17.12.5)
	ImageLocation string   `json:"image-location"` // AP image storage location path (Live: IOS-XE 17.12.5)
	ImageVersion  string   `json:"image-version"`  // AP image version identifier (Live: IOS-XE 17.12.5)
	IsNew         bool     `json:"is-new"`         // New image flag for install operation (Live: IOS-XE 17.12.5)
	FileSize      string   `json:"file-size"`      // AP image file size (Live: IOS-XE 17.12.5)
	ApModelList   []string `json:"ap-model-list"`  // List of supported AP models for this image (Live: IOS-XE 17.12.5)
}

// ApImageActiveLocation represents AP image active location.
type ApImageActiveLocation struct {
	Index                          int    `json:"index"`      // AP image index identifier for active location (Live: IOS-XE 17.12.5)
	ImageFile                      string `json:"image-file"` // AP image file name for active location (Live: IOS-XE 17.12.5)
	ApImageActiveLocationImageData []struct {
		ImageName                                 string   `json:"image-name"`     // AP image name identifier (Live: IOS-XE 17.12.5)
		ImageLocation                             string   `json:"image-location"` // AP image storage location path (Live: IOS-XE 17.12.5)
		ImageVersion                              string   `json:"image-version"`  // AP image version identifier (Live: IOS-XE 17.12.5)
		IsNew                                     bool     `json:"is-new"`         // New image flag for install operation (Live: IOS-XE 17.12.5)
		FileSize                                  string   `json:"file-size"`      // AP image file size (Live: IOS-XE 17.12.5)
		ApImageActiveLocationImageDataApModelList []string `json:"ap-model-list"`  // List of supported AP models for this image (Live: IOS-XE 17.12.5)
	} `json:"image-data"` // AP image info for active location (Live: IOS-XE 17.12.5)
}

// TCPMssConfig represents TCP MSS adjustment configuration.
type TCPMssConfig struct {
	TCPAdjustMssState bool `json:"tcp-adjust-mss-state"` // TCP MSS clamping state for CAPWAP (Live: IOS-XE 17.12.5)
	TCPAdjustMssSize  int  `json:"tcp-adjust-mss-size"`  // TCP MSS clamp size in bytes (Live: IOS-XE 17.12.5)
}

// PersistentSsid represents persistent SSID configuration.
type PersistentSsid struct {
	IsPersistentSsidEnabled bool `json:"is-persistent-ssid-enabled"` // SSID persistence across reboots/failover (Live: IOS-XE 17.12.5)
}

// CdpIPAddress represents CDP IP address information.
type CdpIPAddress struct {
	IPAddressValue []string `json:"ip-address-value"` // CDP discovered neighbor IP addresses (Live: IOS-XE 17.12.5)
}

// RadioOperStats represents radio operational statistics.
type RadioOperStats struct {
	ApMac                 string      `json:"ap-mac"`                    // Access point MAC address (Live: IOS-XE 17.12.5)
	SlotID                int         `json:"slot-id"`                   // Radio slot identifier (Live: IOS-XE 17.12.5)
	AidUserList           interface{} `json:"aid-user-list"`             // Association ID user list for this radio (Live: IOS-XE 17.12.5)
	TxFragmentCount       int         `json:"tx-fragment-count"`         // Number of transmitted frame fragments (Live: IOS-XE 17.12.5)
	MultipleRetryCount    int         `json:"multiple-retry-count"`      // Multi-retry frame count (Live: IOS-XE 17.12.5)
	MulticastTxFrameCnt   int         `json:"multicast-tx-frame-cnt"`    // Number of multicast frames transmitted (Live: IOS-XE 17.12.5)
	FailedCount           int         `json:"failed-count"`              // Number of failed transmission attempts (Live: IOS-XE 17.12.5)
	RetryCount            int         `json:"retry-count"`               // Number of frame retransmission attempts (Live: IOS-XE 17.12.5)
	FrameDuplicateCount   int         `json:"frame-duplicate-count"`     // Number of duplicate frames received (Live: IOS-XE 17.12.5)
	AckFailureCount       int         `json:"ack-failure-count"`         // Number of acknowledgment failures (Live: IOS-XE 17.12.5)
	FcsErrorCount         int         `json:"fcs-error-count"`           // Number of frames with frame check sequence errors (Live: IOS-XE 17.12.5)
	MacDecryErrFrameCount int         `json:"mac-decry-err-frame-count"` // Number of frames with MAC decryption errors (Live: IOS-XE 17.12.5)
	MacMicErrFrameCount   int         `json:"mac-mic-err-frame-count"`   // MAC MIC error frame count (Live: IOS-XE 17.12.5)
	MulticastRxFrameCnt   int         `json:"multicast-rx-frame-cnt"`    // Number of multicast frames received (Live: IOS-XE 17.12.5)
	NoiseFloor            int         `json:"noise-floor"`               // Current noise floor level in dBm (Live: IOS-XE 17.12.5)
	RtsFailureCount       int         `json:"rts-failure-count"`         // Number of Request to Send (RTS) failures (Live: IOS-XE 17.12.5)
	RtsSuccessCount       int         `json:"rts-success-count"`         // Number of successful Request to Send (RTS) transmissions (Live: IOS-XE 17.12.5)
	RxCtrlFrameCount      int         `json:"rx-ctrl-frame-count"`       // Number of control frames received (Live: IOS-XE 17.12.5)
	RxDataFrameCount      int         `json:"rx-data-frame-count"`       // Number of data frames received (Live: IOS-XE 17.12.5)
	RxDataPktCount        int         `json:"rx-data-pkt-count"`         // Number of data packets received (Live: IOS-XE 17.12.5)
	RxErrorFrameCount     int         `json:"rx-error-frame-count"`      // Number of frames received with errors (Live: IOS-XE 17.12.5)
	RxFragmentCount       int         `json:"rx-fragment-count"`         // Number of frame fragments received (Live: IOS-XE 17.12.5)
	RxMgmtFrameCount      int         `json:"rx-mgmt-frame-count"`       // Number of management frames received (Live: IOS-XE 17.12.5)
	TxCtrlFrameCount      int         `json:"tx-ctrl-frame-count"`       // Number of control frames transmitted (Live: IOS-XE 17.12.5)
	TxDataFrameCount      int         `json:"tx-data-frame-count"`       // Number of data frames transmitted (Live: IOS-XE 17.12.5)
	TxDataPktCount        int         `json:"tx-data-pkt-count"`         // Number of data packets transmitted (Live: IOS-XE 17.12.5)
	TxFrameCount          int         `json:"tx-frame-count"`            // Total number of frames transmitted (Live: IOS-XE 17.12.5)
	TxMgmtFrameCount      int         `json:"tx-mgmt-frame-count"`       // Number of management frames transmitted (Live: IOS-XE 17.12.5)
	WepUndecryptableCount int         `json:"wep-undecryptable-count"`   // Number of WEP frames that could not be decrypted (Live: IOS-XE 17.12.5)
	ApRadioStats          interface{} `json:"ap-radio-stats,omitempty"`  // Additional access point radio statistics (Live: IOS-XE 17.12.5)
}

// EthernetIfStats represents Ethernet interface statistics.
type EthernetIfStats struct {
	WtpMac           string `json:"wtp-mac"`            // Wireless termination point MAC address (Live: IOS-XE 17.12.5)
	IfIndex          int    `json:"if-index"`           // Interface index identifier (Live: IOS-XE 17.12.5)
	IfName           string `json:"if-name"`            // Interface name identifier (Live: IOS-XE 17.12.5)
	RxPkts           int    `json:"rx-pkts"`            // Total packets received on interface (Live: IOS-XE 17.12.5)
	TxPkts           int    `json:"tx-pkts"`            // Total packets transmitted on interface (Live: IOS-XE 17.12.5)
	OperStatus       string `json:"oper-status"`        // Current operational status of interface (Live: IOS-XE 17.12.5)
	RxUcastPkts      int    `json:"rx-ucast-pkts"`      // Unicast packets received (Live: IOS-XE 17.12.5)
	RxNonUcastPkts   int    `json:"rx-non-ucast-pkts"`  // Non-unicast packets received (broadcast/multicast) (Live: IOS-XE 17.12.5)
	TxUcastPkts      int    `json:"tx-ucast-pkts"`      // Unicast packets transmitted (Live: IOS-XE 17.12.5)
	TxNonUcastPkts   int    `json:"tx-non-ucast-pkts"`  // Non-unicast packets transmitted (broadcast/multicast) (Live: IOS-XE 17.12.5)
	Duplex           int    `json:"duplex"`             // Duplex mode of interface (full/half duplex) (Live: IOS-XE 17.12.5)
	LinkSpeed        int    `json:"link-speed"`         // Current link speed in bits per second (Live: IOS-XE 17.12.5)
	RxTotalBytes     int    `json:"rx-total-bytes"`     // Total bytes received on interface (Live: IOS-XE 17.12.5)
	TxTotalBytes     int    `json:"tx-total-bytes"`     // Total bytes transmitted on interface (Live: IOS-XE 17.12.5)
	InputCrc         int    `json:"input-crc"`          // Input cyclic redundancy check errors (Live: IOS-XE 17.12.5)
	InputAborts      int    `json:"input-aborts"`       // Input packets aborted during reception (Live: IOS-XE 17.12.5)
	InputErrors      int    `json:"input-errors"`       // Total input errors on interface (Live: IOS-XE 17.12.5)
	InputFrames      int    `json:"input-frames"`       // Input framing errors (Live: IOS-XE 17.12.5)
	InputOverrun     int    `json:"input-overrun"`      // Input overrun errors (Live: IOS-XE 17.12.5)
	InputDrops       int    `json:"input-drops"`        // Input packets dropped by interface (Live: IOS-XE 17.12.5)
	InputResource    int    `json:"input-resource"`     // Input packets dropped due to resource limitations (Live: IOS-XE 17.12.5)
	UnknownProtocol  int    `json:"unknown-protocol"`   // Packets with unknown or unsupported protocol (Live: IOS-XE 17.12.5)
	Runts            int    `json:"runts"`              // Packets smaller than minimum frame size (Live: IOS-XE 17.12.5)
	Giants           int    `json:"giants"`             // Packets larger than maximum frame size (Live: IOS-XE 17.12.5)
	Throttle         int    `json:"throttle"`           // Times interface was throttled (Live: IOS-XE 17.12.5)
	Resets           int    `json:"resets"`             // Number of interface resets performed (Live: IOS-XE 17.12.5)
	OutputCollision  int    `json:"output-collision"`   // Output collision detection events (Live: IOS-XE 17.12.5)
	OutputNoBuffer   int    `json:"output-no-buffer"`   // Output packets dropped due to no buffer space (Live: IOS-XE 17.12.5)
	OutputResource   int    `json:"output-resource"`    // Output packets dropped due to resource limits (Live: IOS-XE 17.12.5)
	OutputUnderrun   int    `json:"output-underrun"`    // Output underrun errors (Live: IOS-XE 17.12.5)
	OutputErrors     int    `json:"output-errors"`      // Total output errors on interface (Live: IOS-XE 17.12.5)
	OutputTotalDrops int    `json:"output-total-drops"` // Total output packets dropped (Live: IOS-XE 17.12.5)
}

// EwlcWncdStats represents EWLC WNCD statistics.
type EwlcWncdStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // Number of predownload sessions initiated (Live: IOS-XE 17.12.5)
		NumInProgress           int  `json:"num-in-progress"`            // Predownload sessions in progress (Live: IOS-XE 17.12.5)
		NumComplete             int  `json:"num-complete"`               // Predownload sessions completed (Live: IOS-XE 17.12.5)
		NumUnsupported          int  `json:"num-unsupported"`            // Number of unsupported predownload requests (Live: IOS-XE 17.12.5)
		NumFailed               int  `json:"num-failed"`                 // Number of predownload sessions that failed (Live: IOS-XE 17.12.5)
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // Predownload operation active status (Live: IOS-XE 17.12.5)
		NumTotal                int  `json:"num-total"`                  // Total number of predownload sessions attempted (Live: IOS-XE 17.12.5)
	} `json:"predownload-stats"` // EWC predownload statistics (Live: IOS-XE 17.12.5)
	DownloadsComplete   int         `json:"downloads-complete"`              // Total number of completed downloads (Live: IOS-XE 17.12.5)
	DownloadsInProgress int         `json:"downloads-in-progress"`           // Number of downloads currently in progress (Live: IOS-XE 17.12.5)
	WlcPredownloadStats interface{} `json:"wlc-predownload-stats,omitempty"` // Wireless LAN Controller predownload statistics (Live: IOS-XE 17.12.5)
}

// IotFirmware represents IoT firmware information for access points.
type IotFirmware struct {
	ApMac      string    `json:"ap-mac"`      // Access point MAC address (Live: IOS-XE 17.12.5)
	IfName     string    `json:"if-name"`     // Interface name for IoT radio (Live: IOS-XE 17.12.5)
	IsDefault  EmptyType `json:"is-default"`  // Default firmware status (Live: IOS-XE 17.12.5)
	Version    string    `json:"version"`     // Firmware version string (Live: IOS-XE 17.12.5)
	VendorName string    `json:"vendor-name"` // Firmware vendor name (Live: IOS-XE 17.12.5)
	Type       string    `json:"type"`        // Firmware type identifier (Live: IOS-XE 17.12.5)
	Desc       string    `json:"desc"`        // Firmware description (Live: IOS-XE 17.12.5)
}

// EmptyType represents YANG empty type fields appearing as null arrays in RESTCONF JSON.
type EmptyType []interface{} // appears as [null]
