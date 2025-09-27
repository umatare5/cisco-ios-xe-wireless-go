package ap

import "time"

// CiscoIOSXEWirelessAPOper represents access point operational data response.
type CiscoIOSXEWirelessAPOper struct {
	CiscoIOSXEWirelessAPOperData struct {
		ApRadioNeighbor         []ApRadioNeighbor        `json:"ap-radio-neighbor"`          // AP radio neighbor information (Live: IOS-XE 17.12.6a)
		RadioOperData           []RadioOperData          `json:"radio-oper-data"`            // Radio operational data corresponding to a radio of the 802.11 LWAPP AP (Live: IOS-XE 17.12.6a)
		RadioResetStats         []RadioResetStats        `json:"radio-reset-stats"`          // Radio reset stats (Live: IOS-XE 17.12.6a)
		QosClientData           []QosClientData          `json:"qos-client-data,omitempty"`  // QoS client data (YANG: IOS-XE 17.12.1)
		CAPWAPData              []CAPWAPData             `json:"capwap-data"`                // Information about the 802.11 LWAPP AP that has joined the controller (Live: IOS-XE 17.12.6a)
		ApNameMACMap            []ApNameMACMap           `json:"ap-name-mac-map"`            // Mapping between AP name and radio MAC of AP (Live: IOS-XE 17.12.6a)
		WtpSlotWlanStats        []WtpSlotWlanStats       `json:"wtp-slot-wlan-stats"`        // AP slot and WLAN stats (Live: IOS-XE 17.12.6a)
		EthernetMACWtpMACMap    []EthernetMACWtpMACMap   `json:"ethernet-mac-wtp-mac-map"`   // Mapping between AP ethernet MAC and base radio MAC (Live: IOS-XE 17.12.6a)
		RadioOperStats          []RadioOperStats         `json:"radio-oper-stats"`           // Operational statistics for a particular radio (Live: IOS-XE 17.12.6a)
		EthernetIfStats         []EthernetIfStats        `json:"ethernet-if-stats"`          // Ethernet interface statistics (Live: IOS-XE 17.12.6a)
		EwlcWncdStats           EwlcWncdStats            `json:"ewlc-wncd-stats"`            // AP image download and predownload statistics for EWC on AP platforms (Live: IOS-XE 17.12.6a)
		ApIoxOperData           []ApIoxOperData          `json:"ap-iox-oper-data"`           // IOx application hosting operational data reported by the AP (Live: IOS-XE 17.12.6a)
		QosGlobalStats          QosGlobalStats           `json:"qos-global-stats"`           // QoS Global statistics data in DB (Live: IOS-XE 17.12.6a)
		OperData                []ApOperInternalData     `json:"oper-data"`                  // Operational data corresponding to an 802.11 LWAPP AP (Live: IOS-XE 17.12.6a)
		RlanOper                []RlanOper               `json:"rlan-oper,omitempty"`        // LAN information of the AP (YANG: IOS-XE 17.12.1)
		EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec  `json:"ewlc-mewlc-predownload-rec"` // Embedded Wireless Controller predownload data (Live: IOS-XE 17.12.6a)
		CdpCacheData            []CdpCacheData           `json:"cdp-cache-data"`             // Cached neighbor information via CDP messages on APs (Live: IOS-XE 17.12.6a)
		LldpNeigh               []LldpNeigh              `json:"lldp-neigh"`                 // Cached neighbor information via LLDP messages on APs (Live: IOS-XE 17.12.6a)
		TpCertInfo              TpCertInfo               `json:"tp-cert-info"`               // Trustpoint Certificate information (Live: IOS-XE 17.12.6a)
		DiscData                []DiscData               `json:"disc-data"`                  // Discovery packet counters (Live: IOS-XE 17.12.6a)
		CAPWAPPkts              []CAPWAPPkts             `json:"capwap-pkts"`                // CAPWAP packet counters (Live: IOS-XE 17.12.6a)
		CountryOper             []CountryOper            `json:"country-oper"`               // Regulatory Domain country details (Live: IOS-XE 17.12.6a)
		SuppCountryOper         []SuppCountryOper        `json:"supp-country-oper"`          // Supported Regulatory Domain country details (Live: IOS-XE 17.12.6a)
		ApNhGlobalData          ApNhGlobalData           `json:"ap-nh-global-data"`          // Information about the RRM based AP clustering algorithm stats (Live: IOS-XE 17.12.6a)
		ApImagePrepareLocation  []ApImagePrepareLocation `json:"ap-image-prepare-location"`  // AP image for prepare location (Live: IOS-XE 17.12.6a)
		ApImageActiveLocation   []ApImageActiveLocation  `json:"ap-image-active-location"`   // AP image for active location (Live: IOS-XE 17.12.6a)
		IotFirmware             []IotFirmware            `json:"iot-firmware"`               // IoT radio firmware operational data (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"` // Root container of access point operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessApOperApRadioNeighbor represents the access point radio neighbor response.
type CiscoIOSXEWirelessApOperApRadioNeighbor struct {
	ApRadioNeighbor []ApRadioNeighbor `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor"`
}

// CiscoIOSXEWirelessApOperRadioOperData represents the radio operational data response.
type CiscoIOSXEWirelessApOperRadioOperData struct {
	RadioOperData []RadioOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data"`
}

// CiscoIOSXEWirelessApOperRadioResetStats represents the radio reset statistics response.
type CiscoIOSXEWirelessApOperRadioResetStats struct {
	RadioResetStats []RadioResetStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-reset-stats"`
}

// CiscoIOSXEWirelessApOperQosClientData represents the QoS client data response.
type CiscoIOSXEWirelessApOperQosClientData struct {
	QosClientData []QosClientData `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data"`
}

// CiscoIOSXEWirelessApOperCAPWAPData represents the CAPWAP data response.
type CiscoIOSXEWirelessApOperCAPWAPData struct {
	CAPWAPData []CAPWAPData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
}

// CiscoIOSXEWirelessApOperApNameMACMap represents the AP name to MAC mapping response.
type CiscoIOSXEWirelessApOperApNameMACMap struct {
	ApNameMACMap []ApNameMACMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
}

// CiscoIOSXEWirelessApOperWtpSlotWlanStats represents the WTP slot WLAN statistics response.
type CiscoIOSXEWirelessApOperWtpSlotWlanStats struct {
	WtpSlotWlanStats []WtpSlotWlanStats `json:"Cisco-IOS-XE-wireless-access-point-oper:wtp-slot-wlan-stats"`
}

// CiscoIOSXEWirelessApOperEthernetMACWtpMACMap represents the Ethernet MAC to WTP MAC mapping response.
type CiscoIOSXEWirelessApOperEthernetMACWtpMACMap struct {
	EthernetMACWtpMACMap []EthernetMACWtpMACMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map"`
}

// CiscoIOSXEWirelessApOperRadioOperStats represents the radio operational statistics response.
type CiscoIOSXEWirelessApOperRadioOperStats struct {
	RadioOperStats []RadioOperStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-stats"`
}

// CiscoIOSXEWirelessApOperEthernetIfStats represents the Ethernet interface statistics response.
type CiscoIOSXEWirelessApOperEthernetIfStats struct {
	EthernetIfStats []EthernetIfStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats"`
}

// CiscoIOSXEWirelessApOperEwlcWncdStats represents the EWLC WNCD statistics response.
type CiscoIOSXEWirelessApOperEwlcWncdStats struct {
	EwlcWncdStats EwlcWncdStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-wncd-stats"`
}

// CiscoIOSXEWirelessApOperApIoxOperData represents the AP IOx operational data response.
type CiscoIOSXEWirelessApOperApIoxOperData struct {
	ApIoxOperData []ApIoxOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-iox-oper-data"`
}

// CiscoIOSXEWirelessApOperQosGlobalStats represents the QoS global statistics response.
type CiscoIOSXEWirelessApOperQosGlobalStats struct {
	QosGlobalStats QosGlobalStats `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-global-stats"`
}

// CiscoIOSXEWirelessApOperData represents the AP operational data response.
type CiscoIOSXEWirelessApOperData struct {
	OperData []ApOperInternalData `json:"Cisco-IOS-XE-wireless-access-point-oper:oper-data"`
}

// CiscoIOSXEWirelessApOperRlanOper represents the RLAN operational data response.
type CiscoIOSXEWirelessApOperRlanOper struct {
	RlanOper []RlanOper `json:"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper"`
}

// CiscoIOSXEWirelessApOperEwlcMewlcPredownloadRec represents the EWLC MEWLC predownload record response.
type CiscoIOSXEWirelessApOperEwlcMewlcPredownloadRec struct {
	EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-mewlc-predownload-rec"`
}

// CiscoIOSXEWirelessApOperCdpCacheData represents the CDP cache data response.
type CiscoIOSXEWirelessApOperCdpCacheData struct {
	CdpCacheData []CdpCacheData `json:"Cisco-IOS-XE-wireless-access-point-oper:cdp-cache-data"`
}

// CiscoIOSXEWirelessApOperLldpNeigh represents the LLDP neighbor response.
type CiscoIOSXEWirelessApOperLldpNeigh struct {
	LldpNeigh []LldpNeigh `json:"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh"`
}

// CiscoIOSXEWirelessApOperTpCertInfo represents the trustpoint certificate info response.
type CiscoIOSXEWirelessApOperTpCertInfo struct {
	TpCertInfo TpCertInfo `json:"Cisco-IOS-XE-wireless-access-point-oper:tp-cert-info"`
}

// CiscoIOSXEWirelessApOperDiscData represents the discovery data response.
type CiscoIOSXEWirelessApOperDiscData struct {
	DiscData []DiscData `json:"Cisco-IOS-XE-wireless-access-point-oper:disc-data"`
}

// CiscoIOSXEWirelessApOperCAPWAPPkts represents the CAPWAP packets response.
type CiscoIOSXEWirelessApOperCAPWAPPkts struct {
	CAPWAPPkts []CAPWAPPkts `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-pkts"`
}

// CiscoIOSXEWirelessApOperCountryOper represents the country operational data response.
type CiscoIOSXEWirelessApOperCountryOper struct {
	CountryOper []CountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:country-oper"`
}

// CiscoIOSXEWirelessApOperSuppCountryOper represents the supported country operational data response.
type CiscoIOSXEWirelessApOperSuppCountryOper struct {
	SuppCountryOper []SuppCountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:supp-country-oper"`
}

// CiscoIOSXEWirelessApOperApNhGlobalData represents the AP neighborhood global data response.
type CiscoIOSXEWirelessApOperApNhGlobalData struct {
	ApNhGlobalData ApNhGlobalData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-nh-global-data"`
}

// CiscoIOSXEWirelessApOperApImagePrepareLocation represents the AP image prepare location response.
type CiscoIOSXEWirelessApOperApImagePrepareLocation struct {
	ApImagePrepareLocation []ApImagePrepareLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-prepare-location"`
}

// CiscoIOSXEWirelessApOperApImageActiveLocation represents the AP image active location response.
type CiscoIOSXEWirelessApOperApImageActiveLocation struct {
	ApImageActiveLocation []ApImageActiveLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-active-location"`
}

// CiscoIOSXEWirelessApOperApPwrInfo represents the AP power information response.
type CiscoIOSXEWirelessApOperApPwrInfo struct {
	ApPwrInfo []ApPwrInfo `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-pwr-info"`
}

// CiscoIOSXEWirelessApOperApSensorStatus represents the AP sensor status response.
type CiscoIOSXEWirelessApOperApSensorStatus struct {
	ApSensorStatus []ApSensorStatus `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-sensor-status"`
}

// CiscoIOSXEWirelessApOperIotFirmware represents the IoT firmware response.
type CiscoIOSXEWirelessApOperIotFirmware struct {
	IotFirmware []IotFirmware `json:"Cisco-IOS-XE-wireless-access-point-oper:iot-firmware"`
}

// ApPwrInfo represents AP power information.
type ApPwrInfo struct {
	WtpMAC  string    `json:"wtp-mac"`  // AP Radio MAC address (Live: IOS-XE 17.12.6a)
	Status  string    `json:"status"`   // Power status (Live: IOS-XE 17.12.6a)
	PpeInfo []PpeInfo `json:"ppe-info"` // Power policy entries (Live: IOS-XE 17.12.6a)
}

// PpeInfo represents power policy entry information.
type PpeInfo struct {
	SeqNumber int    `json:"seq-number"` // Power policy sequence number (Live: IOS-XE 17.12.6a)
	PpeResult string `json:"ppe-result"` // Power policy result (Live: IOS-XE 17.12.6a)
	Ethernet  *struct {
		EthID    string `json:"eth-id"`    // Ethernet interface ID (Live: IOS-XE 17.12.6a)
		EthSpeed string `json:"eth-speed"` // Ethernet speed (Live: IOS-XE 17.12.6a)
	} `json:"ethernet,omitempty"` // Ethernet interface info (Live: IOS-XE 17.12.6a)
	Radio *struct {
		RadioID       string  `json:"radio-id"`                 // Radio interface ID (Live: IOS-XE 17.12.6a)
		SpatialStream *string `json:"spatial-stream,omitempty"` // Spatial stream config (Live: IOS-XE 17.12.6a)
		State         *string `json:"state,omitempty"`          // Radio state (Live: IOS-XE 17.12.6a)
	} `json:"radio,omitempty"` // Radio interface info (Live: IOS-XE 17.12.6a)
	Usb *struct {
		UsbID string  `json:"usb-id"`          // USB interface ID (Live: IOS-XE 17.12.6a)
		State *string `json:"state,omitempty"` // USB state (Live: IOS-XE 17.12.6a)
	} `json:"usb,omitempty"` // USB interface info (Live: IOS-XE 17.12.6a)
}

// ApSensorStatus represents AP sensor status information.
type ApSensorStatus struct {
	ApMAC       string `json:"ap-mac"`       // AP MAC address (Live: IOS-XE 17.12.6a)
	SensorType  string `json:"sensor-type"`  // Sensor type ID (Live: IOS-XE 17.12.6a)
	ConfigState string `json:"config-state"` // Sensor config state (Live: IOS-XE 17.12.6a)
	AdminState  string `json:"admin-state"`  // Admin state (Live: IOS-XE 17.12.6a)
}

// ApRadioNeighbor represents AP radio neighbor information.
type ApRadioNeighbor struct {
	ApMAC          string    `json:"ap-mac"`           // Access point MAC address (Live: IOS-XE 17.12.6a)
	SlotID         int       `json:"slot-id"`          // Radio slot identifier (Live: IOS-XE 17.12.6a)
	Bssid          string    `json:"bssid"`            // Basic Service Set Identifier (Live: IOS-XE 17.12.6a)
	Ssid           string    `json:"ssid"`             // Service Set Identifier (Live: IOS-XE 17.12.6a)
	RSSI           int       `json:"rssi"`             // Received Signal Strength Indicator (Live: IOS-XE 17.12.6a)
	Channel        int       `json:"channel"`          // Operating channel number (Live: IOS-XE 17.12.6a)
	PrimaryChannel int       `json:"primary-channel"`  // Primary channel number (Live: IOS-XE 17.12.6a)
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Last neighbor update timestamp (Live: IOS-XE 17.12.6a)
}

// RadioOperData represents radio operational data.
type RadioOperData struct {
	WtpMAC       string `json:"wtp-mac"`                  // Wireless Termination Point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID  int    `json:"radio-slot-id"`            // Radio slot identifier (Live: IOS-XE 17.12.6a)
	SlotID       int    `json:"slot-id,omitempty"`        // Physical slot identifier (Live: IOS-XE 17.12.6a)
	RadioType    string `json:"radio-type,omitempty"`     // Radio hardware type (Live: IOS-XE 17.12.6a)
	AdminState   string `json:"admin-state,omitempty"`    // Administrative state (Live: IOS-XE 17.12.6a)
	OperState    string `json:"oper-state,omitempty"`     // Operational state (Live: IOS-XE 17.12.6a)
	RadioMode    string `json:"radio-mode,omitempty"`     // Radio operational mode (Live: IOS-XE 17.12.6a)
	RadioSubMode string `json:"radio-sub-mode,omitempty"` // Radio sub-mode details (Live: IOS-XE 17.12.6a)
	RadioSubtype string `json:"radio-subtype,omitempty"`  // Radio hardware subtype (Live: IOS-XE 17.12.6a)
	RadioSubband string `json:"radio-subband,omitempty"`  // Radio frequency subband (Live: IOS-XE 17.12.6a)

	// Band and channel information
	CurrentBandID     int    `json:"current-band-id,omitempty"`     // Active band ID (Live: IOS-XE 17.12.6a)
	CurrentActiveBand string `json:"current-active-band,omitempty"` // Active frequency band (Live: IOS-XE 17.12.6a)

	// Protocol capabilities
	PhyHtCap        *PhyHtCapStruct `json:"phy-ht-cap,omitempty"`        // 802.11n HT capabilities (Live: IOS-XE 17.12.6a)
	PhyHeCap        *PhyHeCapStruct `json:"phy-he-cap,omitempty"`        // 802.11ax HE capabilities (Live: IOS-XE 17.12.6a)
	RadioHeCapable  bool            `json:"radio-he-capable,omitempty"`  // 802.11ax capability status (Live: IOS-XE 17.12.6a)
	RadioFraCapable string          `json:"radio-fra-capable,omitempty"` // Frame aggregation capability (Live: IOS-XE 17.12.6a)

	// XOR capabilities
	XorRadioMode string       `json:"xor-radio-mode,omitempty"` // XOR radio mode (Live: IOS-XE 17.12.6a)
	XorPhyHtCap  *XorPhyHtCap `json:"xor-phy-ht-cap,omitempty"` // XOR HT capabilities (Live: IOS-XE 17.12.6a)
	XorPhyHeCap  *XorPhyHeCap `json:"xor-phy-he-cap,omitempty"` // XOR HE capabilities (Live: IOS-XE 17.12.6a)

	// Additional operational fields
	AntennaGain            int               `json:"antenna-gain,omitempty"`             // Antenna gain value in dBi (Live: IOS-XE 17.12.6a)
	AntennaPid             string            `json:"antenna-pid,omitempty"`              // Antenna product identifier (Live: IOS-XE 17.12.6a)
	SlotAntennaType        string            `json:"slot-antenna-type,omitempty"`        // Antenna type (internal/external) (Live: IOS-XE 17.12.6a)
	RadioEnableTime        string            `json:"radio-enable-time,omitempty"`        // Last radio enable timestamp (Live: IOS-XE 17.12.6a)
	HighestThroughputProto string            `json:"highest-throughput-proto,omitempty"` // Highest throughput protocol supported (Live: IOS-XE 17.12.6a)
	CacActive              bool              `json:"cac-active,omitempty"`               // Channel Availability Check active status (Live: IOS-XE 17.12.6a)
	MeshBackhaul           bool              `json:"mesh-backhaul,omitempty"`            // Mesh backhaul link status (Live: IOS-XE 17.12.6a)
	MeshDesignatedDownlink bool              `json:"mesh-designated-downlink,omitempty"` // Mesh designated downlink status (Live: IOS-XE 17.12.6a)
	MultiDomainCap         *MultiDomainCap   `json:"multi-domain-cap,omitempty"`         // Multi-domain capabilities (Live: IOS-XE 17.12.6a)
	StationCfg             *StationCfg       `json:"station-cfg,omitempty"`              // Station mode configuration (Live: IOS-XE 17.12.6a)
	PhyHtCfg               *PhyHtCfg         `json:"phy-ht-cfg,omitempty"`               // High Throughput configuration settings (Live: IOS-XE 17.12.6a)
	ChanPwrInfo            *ChanPwrInfo      `json:"chan-pwr-info,omitempty"`            // Channel power level information (Live: IOS-XE 17.12.6a)
	SnifferCfg             *SnifferCfg       `json:"sniffer-cfg,omitempty"`              // Packet sniffer configuration (Live: IOS-XE 17.12.6a)
	RadioBandInfo          []RadioBandInfo   `json:"radio-band-info,omitempty"`          // Radio frequency band information (Live: IOS-XE 17.12.6a)
	VapOperConfig          []VapOperConfig   `json:"vap-oper-config,omitempty"`          // VAP operational configuration (Live: IOS-XE 17.12.6a)
	RegDomainCheckStatus   string            `json:"reg-domain-check-status,omitempty"`  // Regulatory compliance check status (Live: IOS-XE 17.12.6a)
	Dot11nMcsRates         string            `json:"dot11n-mcs-rates,omitempty"`         // 802.11n MCS rates supported (Live: IOS-XE 17.12.6a)
	DualRadioModeCfg       *DualRadioModeCfg `json:"dual-radio-mode-cfg,omitempty"`      // Dual radio mode configuration (Live: IOS-XE 17.12.6a)
	BssColorCfg            *BssColorCfg      `json:"bss-color-cfg,omitempty"`            // BSS color configuration for 802.11ax (Live: IOS-XE 17.12.6a)
	ObssPdCapable          bool              `json:"obss-pd-capable,omitempty"`          // OBSS Preamble Detection capability (Live: IOS-XE 17.12.6a)
	NdpCap                 string            `json:"ndp-cap,omitempty"`                  // Null Data Packet capability information (Live: IOS-XE 17.12.6a)
	NdpOnChannel           bool              `json:"ndp-on-channel,omitempty"`           // Null Data Packet transmission on channel status (Live: IOS-XE 17.12.6a)
	BeamSelection          string            `json:"beam-selection,omitempty"`           // Antenna beam selection algorithm configuration (Live: IOS-XE 17.12.6a)
	NumAntEnabled          uint8             `json:"num-ant-enabled,omitempty"`          // Number of antennas enabled (Live: IOS-XE 17.12.6a)
	CurAntBitmap           string            `json:"cur-ant-bitmap,omitempty"`           // Current antenna bitmap (Live: IOS-XE 17.12.6a)
	SuppAntBitmap          string            `json:"supp-ant-bitmap,omitempty"`          // Supported antenna bitmap (Live: IOS-XE 17.12.6a)
	Supp160mhzAntBitmap    string            `json:"supp-160mhz-ant-bitmap,omitempty"`   // 160MHz antenna bitmap (Live: IOS-XE 17.12.6a)
	MaxClientAllowed       uint16            `json:"max-client-allowed,omitempty"`       // Maximum clients allowed (Live: IOS-XE 17.12.6a)
	ObssPdSrgCapable       bool              `json:"obss-pd-srg-capable,omitempty"`      // OBSS PD SRG capability (Live: IOS-XE 17.12.6a)
	CoverageOverlapFactor  uint8             `json:"coverage-overlap-factor,omitempty"`  // RF coverage overlap factor (Live: IOS-XE 17.12.6a)

	// 6GHz related (YANG: IOS-XE 17.12.1)
	Ap6GhzPwrMode    *string `json:"ap-6ghz-pwr-mode,omitempty"`     // 6GHz power mode (LPI/SP/VLP) (YANG: IOS-XE 17.12.1)
	Ap6GhzPwrModeCap *string `json:"ap-6ghz-pwr-mode-cap,omitempty"` // 6GHz power mode capability (YANG: IOS-XE 17.12.1)

	// AFC related
	AFCBelowTxmin    bool `json:"afc-below-txmin,omitempty"`    // AFC below minimum transmission power (YANG: IOS-XE 17.12.1)
	AFCLicenseNeeded bool `json:"afc-license-needed,omitempty"` // AFC license requirement status (YANG: IOS-XE 17.12.1)
	PushAFCRespDone  bool `json:"push-afc-resp-done,omitempty"` // AFC response push completion status (YANG: IOS-XE 17.12.1)
}

// RadioResetStats represents radio reset statistics.
type RadioResetStats struct {
	ApMAC       string `json:"ap-mac"`       // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioID     int    `json:"radio-id"`     // Radio interface identifier (Live: IOS-XE 17.12.6a)
	Cause       string `json:"cause"`        // Reset cause description (Live: IOS-XE 17.12.6a)
	DetailCause string `json:"detail-cause"` // Detailed reset cause information (Live: IOS-XE 17.12.6a)
	Count       int    `json:"count"`        // Reset count since statistics clear (Live: IOS-XE 17.12.6a)
}

// QosClientData represents QoS client data.
type QosClientData struct {
	ClientMAC    string `json:"client-mac"` // Client MAC address (Live: IOS-XE 17.12.6a)
	AaaQosParams struct {
		AaaAvgdtus   int `json:"aaa-avgdtus"`   // AAA average downstream utilization (Live: IOS-XE 17.12.6a)
		AaaAvgrtdtus int `json:"aaa-avgrtdtus"` // AAA average real-time downstream utilization (Live: IOS-XE 17.12.6a)
		AaaBstdtus   int `json:"aaa-bstdtus"`   // AAA burst downstream utilization (Live: IOS-XE 17.12.6a)
		AaaBstrtdtus int `json:"aaa-bstrtdtus"` // AAA burst real-time downstream utilization (Live: IOS-XE 17.12.6a)
		AaaAvgdtds   int `json:"aaa-avgdtds"`   // AAA average downstream data size (Live: IOS-XE 17.12.6a)
		AaaAvgrtdtds int `json:"aaa-avgrtdtds"` // AAA average real-time downstream data size (Live: IOS-XE 17.12.6a)
		AaaBstdtds   int `json:"aaa-bstdtds"`   // AAA burst downstream data size (Live: IOS-XE 17.12.6a)
		AaaBstrtdtds int `json:"aaa-bstrtdtds"` // AAA burst real-time downstream data size (Live: IOS-XE 17.12.6a)
	} `json:"aaa-qos-params"` // AAA QoS parameters (Live: IOS-XE 17.12.6a)
}

// CAPWAPData represents CAPWAP data.
type CAPWAPData struct {
	WtpMAC       string       `json:"wtp-mac"`       // WTP MAC address for CAPWAP session (Live: IOS-XE 17.12.6a)
	IPAddr       string       `json:"ip-addr"`       // AP management IP address (Live: IOS-XE 17.12.6a)
	Name         string       `json:"name"`          // AP hostname identifier (Live: IOS-XE 17.12.6a)
	DeviceDetail DeviceDetail `json:"device-detail"` // Hardware device specifications (Live: IOS-XE 17.12.6a)
	ApState      ApState      `json:"ap-state"`      // AP operational and admin state (Live: IOS-XE 17.12.6a)

	// AP Mode Data
	ApModeData ApModeData `json:"ap-mode-data"` // AP operational mode configuration (Live: IOS-XE 17.12.6a)

	// Location and Services
	ApLocation         ApLocation         `json:"ap-location"`          // Physical deployment location (Live: IOS-XE 17.12.6a)
	ApServices         ApServices         `json:"ap-services"`          // Enabled AP service capabilities (Live: IOS-XE 17.12.6a)
	TagInfo            TagInfo            `json:"tag-info"`             // Policy and site tag assignment (Live: IOS-XE 17.12.6a)
	Tunnel             Tunnel             `json:"tunnel"`               // CAPWAP tunnel configuration (Live: IOS-XE 17.12.6a)
	ExternalModuleData ExternalModuleData `json:"external-module-data"` // USB and expansion module information (Live: IOS-XE 17.12.6a)
	ApTimeInfo         ApTimeInfo         `json:"ap-time-info"`         // Time synchronization data (Live: IOS-XE 17.12.6a)
	ApSecurityData     ApSecurityData     `json:"ap-security-data"`     // Security configuration (Live: IOS-XE 17.12.6a)
	SlidingWindow      SlidingWindow      `json:"sliding-window"`       // CAPWAP sliding window parameters (Live: IOS-XE 17.12.6a)
	ApVlan             ApVlan             `json:"ap-vlan"`              // VLAN tagging configuration (Live: IOS-XE 17.12.6a)
	HyperlocationData  HyperlocationData  `json:"hyperlocation-data"`   // Hyperlocation service configuration (Live: IOS-XE 17.12.6a)
	RebootStats        RebootStats        `json:"reboot-stats"`         // AP reboot history and analysis (Live: IOS-XE 17.12.6a)
	ProxyInfo          ProxyInfo          `json:"proxy-info"`           // HTTP proxy configuration (Live: IOS-XE 17.12.6a)

	// Image Download Tracking
	ImageSizeEta           uint64 `json:"image-size-eta"`            // AP firmware image download ETA (Live: IOS-XE 17.12.6a)
	ImageSizeStartTime     string `json:"image-size-start-time"`     // AP image download start timestamp (Live: IOS-XE 17.12.6a)
	ImageSizePercentage    uint32 `json:"image-size-percentage"`     // AP image download progress percentage (Live: IOS-XE 17.12.6a)
	WlcImageSizeEta        uint64 `json:"wlc-image-size-eta"`        // WLC firmware image download ETA (Live: IOS-XE 17.12.6a)
	WlcImageSizeStartTime  string `json:"wlc-image-size-start-time"` // WLC image download start timestamp (Live: IOS-XE 17.12.6a)
	WlcImageSizePercentage uint32 `json:"wlc-image-size-percentage"` // WLC image download progress percentage (Live: IOS-XE 17.12.6a)

	// Local DHCP Configuration
	Ipv4Pool              Ipv4Pool            `json:"ipv4-pool"`                // Local DHCP IPv4 pool configuration (Live: IOS-XE 17.12.6a)
	DisconnectDetail      DisconnectDetail    `json:"disconnect-detail"`        // AP disconnection analysis (Live: IOS-XE 17.12.6a)
	StatsMonitor          StatsMonitor        `json:"stats-monitor"`            // AP statistics monitoring configuration (Live: IOS-XE 17.12.6a)
	LscStatusPldSupported *[]LscStatusPayload `json:"lsc-status-pld-supported"` // LSC status payload support capability (Live: IOS-XE 17.12.6a)
	ApLscStatus           ApLscStatus         `json:"ap-lsc-status"`            // LSC authentication status (Live: IOS-XE 17.12.6a)
	RadioStatsMonitor     RadioStatsMonitor   `json:"radio-stats-monitor"`      // Radio statistics monitoring (Live: IOS-XE 17.12.6a)
	ZeroWtDFS             ZeroWtDFS           `json:"zero-wt-dfs"`              // Zero Wait DFS configuration (Live: IOS-XE 17.12.6a)
	GnssInfo              GnssInfo            `json:"gnss-info"`                // GNSS positioning data (Live: IOS-XE 17.12.6a)

	// Basic Configuration Fields
	ApLagEnabled    bool   `json:"ap-lag-enabled"`    // LAG configuration status (Live: IOS-XE 17.12.6a)
	CountryCode     string `json:"country-code"`      // Regulatory country code (Live: IOS-XE 17.12.6a)
	NumRadioSlots   uint8  `json:"num-radio-slots"`   // Number of radio slots available (Live: IOS-XE 17.12.6a)
	Ipv6Joined      uint8  `json:"ipv6-joined"`       // IPv6 join status (Live: IOS-XE 17.12.6a)
	DartIsConnected bool   `json:"dart-is-connected"` // DART connection status (Live: IOS-XE 17.12.6a)
	IsMaster        bool   `json:"is-master"`         // Master AP designation (Live: IOS-XE 17.12.6a)
	CdpEnable       bool   `json:"cdp-enable"`        // CDP enablement (Live: IOS-XE 17.12.6a)
	GrpcEnabled     bool   `json:"grpc-enabled"`      // gRPC streaming enablement (Live: IOS-XE 17.12.6a)
	LocalDHCP       bool   `json:"local-dhcp"`        // Local DHCP server status (Live: IOS-XE 17.12.6a)

	// Status and operational fields
	ApStationType        string `json:"ap-stationing-type,omitempty"`      // AP stationing type configuration (Live: IOS-XE 17.12.6a)
	ApKeepAliveState     bool   `json:"ap-keepalive-state,omitempty"`      // CAPWAP keep-alive state (Live: IOS-XE 17.12.6a)
	MaxClientsSupported  uint16 `json:"max-clients-supported,omitempty"`   // Maximum clients supported (Live: IOS-XE 17.12.6a)
	MDNSGroupID          uint32 `json:"mdns-group-id,omitempty"`           // mDNS group identifier (Live: IOS-XE 17.12.6a)
	MDNSRuleName         string `json:"mdns-rule-name,omitempty"`          // Applied mDNS filtering rule name (Live: IOS-XE 17.12.6a)
	MDNSGroupMethod      string `json:"mdns-group-method,omitempty"`       // mDNS group assignment method (Live: IOS-XE 17.12.6a)
	MerakiCapable        bool   `json:"meraki-capable,omitempty"`          // Meraki cloud capability (YANG: IOS-XE 17.12.1)
	MerakiConnectStatus  string `json:"meraki-connect-status,omitempty"`   // Meraki cloud connection status (Live: IOS-XE 17.12.6a)
	MerakiMonitorCapable bool   `json:"meraki-monitor-capable"`            // Meraki monitoring capability (Live: IOS-XE 17.12.6a)
	KernelCoredumpCount  uint16 `json:"kernel-coredump-count,omitempty"`   // Kernel coredump count (Live: IOS-XE 17.12.6a)
	RegDomain            string `json:"reg-domain,omitempty"`              // Regulatory domain configuration (Live: IOS-XE 17.12.6a)
	DartConStatus        string `json:"dart-con-status,omitempty"`         // DART connection status (Live: IOS-XE 17.12.6a)
	ApAFCPreNotification bool   `json:"ap-afc-pre-notification,omitempty"` // AFC pre-notification capability (YANG: IOS-XE 17.12.1)
	OobImgDwldMethod     string `json:"oob-img-dwld-method,omitempty"`     // Out-of-band image download method (Live: IOS-XE 17.12.6a)
	WtpIP                string `json:"wtp-ip"`                            // WTP IP address (Live: IOS-XE 17.12.6a)
}

// ApTimeInfo represents AP time related information.
type ApTimeInfo struct {
	BootTime      string `json:"boot-time"`       // Last AP reboot timestamp (Live: IOS-XE 17.12.6a)
	JoinTime      string `json:"join-time"`       // AP join timestamp to controller (Live: IOS-XE 17.12.6a)
	JoinTimeTaken uint32 `json:"join-time-taken"` // AP join process duration in seconds (Live: IOS-XE 17.12.6a)
}

// ApSecurityData represents AP LSC (Local Significant Certificate) data.
type ApSecurityData struct {
	FipsEnabled      bool   `json:"fips-enabled"`        // FIPS compliance status (Live: IOS-XE 17.12.6a)
	WlanccEnabled    bool   `json:"wlancc-enabled"`      // WLAN Common Criteria compliance (Live: IOS-XE 17.12.6a)
	CertType         string `json:"cert-type"`           // AP certificate type (Live: IOS-XE 17.12.6a)
	LscApAuthType    string `json:"lsc-ap-auth-type"`    // LSC authentication method (Live: IOS-XE 17.12.6a)
	ApCertPolicy     string `json:"ap-cert-policy"`      // Certificate policy identifier (Live: IOS-XE 17.12.6a)
	ApCertExpiryTime string `json:"ap-cert-expiry-time"` // AP certificate expiration timestamp (Live: IOS-XE 17.12.6a)
	ApCertIssuerCn   string `json:"ap-cert-issuer-cn"`   // Certificate authority common name (Live: IOS-XE 17.12.6a)
}

// SlidingWindow represents CAPWAP multiwindow transport information.
type SlidingWindow struct {
	MultiWindowSupport bool   `json:"multi-window-support"` // CAPWAP multiple window support (Live: IOS-XE 17.12.6a)
	WindowSize         uint16 `json:"window-size"`          // CAPWAP sliding window size (Live: IOS-XE 17.12.6a)
}

// ApVlan represents AP VLAN tagging details.
type ApVlan struct {
	VlanTagState string `json:"vlan-tag-state"` // AP VLAN tagging state (Live: IOS-XE 17.12.6a)
	VlanTagID    uint16 `json:"vlan-tag-id"`    // 802.1Q VLAN identifier (Live: IOS-XE 17.12.6a)
}

// HyperlocationData represents AP Hyperlocation details.
type HyperlocationData struct {
	HyperlocationMethod string `json:"hyperlocation-method"` // Hyperlocation positioning method (Live: IOS-XE 17.12.6a)
	CmxIP               string `json:"cmx-ip,omitempty"`     // CMX server IP address (Live: IOS-XE 17.12.6a)
}

// RebootStats represents AP reboot statistics.
type RebootStats struct {
	RebootReason string `json:"reboot-reason"` // Last AP reboot reason (Live: IOS-XE 17.12.6a)
	RebootType   string `json:"reboot-type"`   // AP reboot type (Live: IOS-XE 17.12.6a)
}

// ProxyInfo represents HTTP proxy configuration provisioned to AP.
type ProxyInfo struct {
	Hostname     string  `json:"hostname"`                // HTTP proxy server hostname or IP (Live: IOS-XE 17.12.6a)
	Port         uint16  `json:"port"`                    // HTTP proxy server TCP port (Live: IOS-XE 17.12.6a)
	NoProxyList  string  `json:"no-proxy-list"`           // URLs to bypass proxy (Live: IOS-XE 17.12.6a)
	Username     *string `json:"username,omitempty"`      // HTTP proxy username (YANG: IOS-XE 17.12.1)
	PasswordType *string `json:"password-type,omitempty"` // HTTP proxy password type (YANG: IOS-XE 17.12.1)
	Password     *string `json:"password,omitempty"`      // HTTP proxy password (YANG: IOS-XE 17.12.1)
}

// Ipv4Pool represents DHCP IPv4 pool configuration.
type Ipv4Pool struct {
	Network   string `json:"network"`    // DHCP network address range (Live: IOS-XE 17.12.6a)
	LeaseTime uint16 `json:"lease-time"` // DHCP lease duration in days (Live: IOS-XE 17.12.6a)
	Netmask   string `json:"netmask"`    // IPv4 subnet mask (Live: IOS-XE 17.12.6a)
}

// DisconnectDetail represents AP disconnect detail.
type DisconnectDetail struct {
	DisconnectReason string `json:"disconnect-reason"` // AP last disconnection reason (Live: IOS-XE 17.12.6a)
}

// StatsMonitor represents AP statistics monitoring configuration.
type StatsMonitor struct {
	ActionApReload bool `json:"action-ap-reload"` // Auto AP reload on critical thresholds (Live: IOS-XE 17.12.6a)
}

// ApLscStatus represents AP LSC (Local Significant Certificate) status information.
type ApLscStatus struct {
	IsDTLSLscEnabled      bool   `json:"is-dtls-lsc-enabled"`                 // LSC enablement for CAPWAP DTLS (Live: IOS-XE 17.12.6a)
	IsDot1xLscEnabled     bool   `json:"is-dot1x-lsc-enabled"`                // LSC enablement for 802.1X (Live: IOS-XE 17.12.6a)
	IsDTLSLscFallback     bool   `json:"is-dtls-lsc-fallback"`                // AP fallback to default certificate (Live: IOS-XE 17.12.6a)
	DTLSLscIssuerHash     string `json:"dtls-lsc-issuer-hash,omitempty"`      // CA hash for CAPWAP DTLS (Live: IOS-XE 17.12.6a)
	Dot1xLscIssuerHash    string `json:"dot1x-lsc-issuer-hash,omitempty"`     // CA hash for 802.1X authentication (Live: IOS-XE 17.12.6a)
	DTLSLscCertExpiryTime string `json:"dtls-lsc-cert-expiry-time,omitempty"` // DTLS LSC certificate expiration (Live: IOS-XE 17.12.6a)
}

// RadioStatsMonitor represents AP radio statistics monitoring configuration.
type RadioStatsMonitor struct {
	Enable       bool              `json:"enable"`        // Radio statistics collection enable (Live: IOS-XE 17.12.6a)
	SampleIntvl  uint16            `json:"sample-intvl"`  // Sampling interval in seconds (Live: IOS-XE 17.12.6a)
	AlarmsEnable []AlarmEnableType `json:"alarms-enable"` // Statistics alarm enablement (Live: IOS-XE 17.12.6a)
	RadioReset   bool              `json:"radio-reset"`   // Auto radio reset on stuck condition (Live: IOS-XE 17.12.6a)
}

// ZeroWtDFS represents Zero wait DFS information of the AP.
type ZeroWtDFS struct {
	ReserveChannel ReserveChannel `json:"reserve-channel"` // DFS channel reservation data (Live: IOS-XE 17.12.6a)
	Type           string         `json:"type"`            // CAC domain type classification (Live: IOS-XE 17.12.6a)
}

// ReserveChannel represents reserved CAC channel information.
type ReserveChannel struct {
	Channel      uint8  `json:"channel"`       // Reserved channel number (Live: IOS-XE 17.12.6a)
	ChannelWidth string `json:"channel-width"` // Channel width for reserved CAC (Live: IOS-XE 17.12.6a)
	State        string `json:"state"`         // CAC state for reserved channel (Live: IOS-XE 17.12.6a)
}

// GnssInfo represents AP GNSS (Global Navigation Satellite System) information.
type GnssInfo struct {
	AntType          string  `json:"ant-type"`             // GNSS antenna type (Live: IOS-XE 17.12.6a)
	AntCableLength   uint16  `json:"ant-cable-length"`     // GNSS antenna cable length in meters (Live: IOS-XE 17.12.6a)
	AntennaProductID string  `json:"antenna-product-id"`   // GNSS antenna product identifier (Live: IOS-XE 17.12.6a)
	AntennaSn        *string `json:"antenna-sn,omitempty"` // GNSS antenna serial number (YANG: IOS-XE 17.18.1)
}

// ApState represents AP state information.
type ApState struct {
	ApAdminState     string `json:"ap-admin-state"`     // AP admin state (enabled/disabled/shutdown) (Live: IOS-XE 17.12.6a)
	ApOperationState string `json:"ap-operation-state"` // AP operational state (Live: IOS-XE 17.12.6a)
}

// ApModeData represents AP mode related information.
type ApModeData struct {
	HomeApEnabled bool         `json:"home-ap-enabled"` // Home AP feature enablement (Live: IOS-XE 17.12.6a)
	ClearMode     bool         `json:"clear-mode"`      // Clear mode status (Live: IOS-XE 17.12.6a)
	ApSubMode     string       `json:"ap-sub-mode"`     // AP operational sub-mode (Live: IOS-XE 17.12.6a)
	WtpMode       string       `json:"wtp-mode"`        // WTP mode (local/flexconnect/monitor) (Live: IOS-XE 17.12.6a)
	ApFabricData  ApFabricData `json:"ap-fabric-data"`  // SDA fabric integration attributes (Live: IOS-XE 17.12.6a)
}

// ApFabricData represents AP fabric related attributes.
type ApFabricData struct {
	IsFabricAp bool `json:"is-fabric-ap"` // SDA fabric-enabled AP designation (Live: IOS-XE 17.12.6a)
}

// ApLocation represents AP location information.
type ApLocation struct {
	Floor             int         `json:"floor"`              // Physical floor number (Live: IOS-XE 17.12.6a)
	Location          string      `json:"location"`           // AP physical placement description (Live: IOS-XE 17.12.6a)
	AaaLocation       AaaLocation `json:"aaa-location"`       // AAA server location parameters (Live: IOS-XE 17.12.6a)
	FloorID           int         `json:"floor-id"`           // Floor identifier for location services (Live: IOS-XE 17.12.6a)
	RangingCapability int         `json:"ranging-capability"` // Location ranging capability level (Live: IOS-XE 17.12.6a)
}

// AaaLocation represents AAA location information.
type AaaLocation struct {
	CivicID string `json:"civic-id"` // Civic location identifier (Live: IOS-XE 17.12.6a)
	GeoID   string `json:"geo-id"`   // Geographic coordinate identifier (Live: IOS-XE 17.12.6a)
	OperID  string `json:"oper-id"`  // Operator location identifier (Live: IOS-XE 17.12.6a)
}

// ApServices represents AP services information.
type ApServices struct {
	MonitorModeOptType string       `json:"monitor-mode-opt-type"` // Monitor mode optimization type (Live: IOS-XE 17.12.6a)
	ApDHCPServer       ApDHCPServer `json:"ap-dhcp-server"`        // Local DHCP server configuration (Live: IOS-XE 17.12.6a)
	TotSnifferRadio    int          `json:"tot-sniffer-radio"`     // Total sniffer radio interfaces (Live: IOS-XE 17.12.6a)
}

// ApDHCPServer represents AP DHCP server configuration.
type ApDHCPServer struct {
	IsDHCPServerEnabled bool `json:"is-dhcp-server-enabled"` // Local DHCP service enablement (Live: IOS-XE 17.12.6a)
}

// XorPhyHtCap represents XOR PHY HT capabilities.
type XorPhyHtCap struct {
	Data XorPhyHtCapData `json:"data"` // XOR High Throughput capability data (Live: IOS-XE 17.12.6a)
}

// XorPhyHtCapData represents XOR PHY HT capability data.
type XorPhyHtCapData struct {
	VhtCapable bool `json:"vht-capable"` // 802.11ac VHT capability support (Live: IOS-XE 17.12.6a)
	HtCapable  bool `json:"ht-capable"`  // 802.11n HT capability support (Live: IOS-XE 17.12.6a)
}

// XorPhyHeCap represents XOR PHY HE capabilities.
type XorPhyHeCap struct {
	Data XorPhyHeCapData `json:"data"` // XOR High Efficiency capability data (Live: IOS-XE 17.12.6a)
}

// XorPhyHeCapData represents XOR PHY HE capability data.
type XorPhyHeCapData struct {
	HeEnabled              bool `json:"he-enabled"`                // 802.11ax HE protocol enablement (Live: IOS-XE 17.12.6a)
	HeCapable              bool `json:"he-capable"`                // 802.11ax capability status (Live: IOS-XE 17.12.6a)
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // 802.11ax single-user beamforming (Live: IOS-XE 17.12.6a)
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // 802.11ax multi-user beamforming (Live: IOS-XE 17.12.6a)
	HeStbcMode             int  `json:"he-stbc-mode"`              // 802.11ax STBC mode (Live: IOS-XE 17.12.6a)
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // 802.11ax A-MPDU TID bitmap (Live: IOS-XE 17.12.6a)
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // 802.11ax MCS and NSS capability (Live: IOS-XE 17.12.6a)
}

// StationCfg represents station configuration.
type StationCfg struct {
	CfgData StationCfgData `json:"cfg-data"` // Station configuration parameters (Live: IOS-XE 17.12.6a)
}

// StationCfgData represents station configuration data.
type StationCfgData struct {
	StationCfgConfigType string `json:"station-cfg-config-type"` // Station configuration type (Live: IOS-XE 17.12.6a)
	MediumOccupancyLimit int    `json:"medium-occupancy-limit"`  // 802.11 medium occupancy limit (Live: IOS-XE 17.12.6a)
	CfpPeriod            int    `json:"cfp-period"`              // CFP interval for PCF access control (Live: IOS-XE 17.12.6a)
	CfpMaxDuration       int    `json:"cfp-max-duration"`        // Maximum CFP duration (Live: IOS-XE 17.12.6a)
	Bssid                string `json:"bssid"`                   // BSS Identifier MAC address (Live: IOS-XE 17.12.6a)
	BeaconPeriod         int    `json:"beacon-period"`           // 802.11 beacon transmission interval (Live: IOS-XE 17.12.6a)
	CountryString        string `json:"country-string"`          // ISO country code string (Live: IOS-XE 17.12.6a)
}

// MultiDomainCap represents multi-domain capability configuration.
type MultiDomainCap struct {
	CfgData MultiDomainCapData `json:"cfg-data"` // Multi-domain capability parameters (Live: IOS-XE 17.12.6a)
}

// MultiDomainCapData represents multi-domain capability data.
type MultiDomainCapData struct {
	FirstChanNum    int `json:"first-chan-num"`     // First channel in regulatory domain (Live: IOS-XE 17.12.6a)
	NumChannels     int `json:"num-channels"`       // Total channels in regulatory domain (Live: IOS-XE 17.12.6a)
	MaxTxPowerLevel int `json:"max-tx-power-level"` // Maximum transmission power (dBm) (Live: IOS-XE 17.12.6a)
}

// PhyHtCfg represents PHY HT configuration.
type PhyHtCfg struct {
	CfgData PhyHtCfgData `json:"cfg-data"` // High Throughput configuration parameters (Live: IOS-XE 17.12.6a)
}

// PhyHtCfgData represents PHY HT configuration data.
type PhyHtCfgData struct {
	HtEnable               int    `json:"ht-enable"`                 // 802.11n HT protocol enablement (Live: IOS-XE 17.12.6a)
	PhyHtCfgConfigType     string `json:"phy-ht-cfg-config-type"`    // Physical layer HT configuration type designation (Live: IOS-XE 17.12.6a)
	CurrFreq               int    `json:"curr-freq"`                 // Current operating frequency (MHz) (Live: IOS-XE 17.12.6a)
	ChanWidth              int    `json:"chan-width"`                // Channel bandwidth width in MHz (20/40/80/160) (Live: IOS-XE 17.12.6a)
	ExtChan                int    `json:"ext-chan"`                  // Extension channel for 40MHz bonding (Live: IOS-XE 17.12.6a)
	VhtEnable              bool   `json:"vht-enable"`                // 802.11ac Very High Throughput protocol enablement (Live: IOS-XE 17.12.6a)
	LegTxBfEnabled         int    `json:"leg-tx-bf-enabled"`         // Legacy TX beamforming enablement (Live: IOS-XE 17.12.6a)
	RRMChannelChangeReason string `json:"rrm-channel-change-reason"` // RRM channel change reason (Live: IOS-XE 17.12.6a)
	FreqString             string `json:"freq-string"`               // Frequency designation string (Live: IOS-XE 17.12.6a)
}

// PhyHtCapStruct represents PHY HT capability structure.
type PhyHtCapStruct struct {
	Data PhyHtCapStructData `json:"data"` // PHY HT capability information (Live: IOS-XE 17.12.6a)
}

// PhyHtCapStructData represents PHY HT capability data.
type PhyHtCapStructData struct {
	VhtCapable bool `json:"vht-capable"` // 802.11ac capability status (Live: IOS-XE 17.12.6a)
	HtCapable  bool `json:"ht-capable"`  // 802.11n capability status (Live: IOS-XE 17.12.6a)
}

// PhyHeCapStruct represents PHY HE capability structure.
type PhyHeCapStruct struct {
	Data PhyHeCapStructData `json:"data"` // High Efficiency capability information (Live: IOS-XE 17.12.6a)
}

// PhyHeCapStructData represents PHY HE capability data.
type PhyHeCapStructData struct {
	HeEnabled              bool `json:"he-enabled"`                // 802.11ax operational enablement status (Live: IOS-XE 17.12.6a)
	HeCapable              bool `json:"he-capable"`                // 802.11ax capability status (Live: IOS-XE 17.12.6a)
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // 802.11ax single-user beamforming (Live: IOS-XE 17.12.6a)
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // 802.11ax multi-user beamforming (Live: IOS-XE 17.12.6a)
	HeStbcMode             int  `json:"he-stbc-mode"`              // 802.11ax STBC operational mode (Live: IOS-XE 17.12.6a)
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // 802.11ax A-MPDU TID bitmap (Live: IOS-XE 17.12.6a)
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // 802.11ax MCS and NSS capabilities (Live: IOS-XE 17.12.6a)
}

// ChanPwrInfo represents channel power information.
type ChanPwrInfo struct {
	Data ChanPwrInfoData `json:"data"` // Channel-specific power data (Live: IOS-XE 17.12.6a)
}

// ChanPwrInfoData represents channel power information data.
type ChanPwrInfoData struct {
	AntennaGain    int         `json:"antenna-gain"`     // Antenna gain value in dBi (Live: IOS-XE 17.12.6a)
	IntAntennaGain int         `json:"int-antenna-gain"` // Internal antenna gain in dBi (Live: IOS-XE 17.12.6a)
	ExtAntennaGain int         `json:"ext-antenna-gain"` // External antenna gain in dBi (Live: IOS-XE 17.12.6a)
	ChanPwrList    ChanPwrList `json:"chan-pwr-list"`    // Per-channel power level configuration (Live: IOS-XE 17.12.6a)
}

// ChanPwrList represents channel power list.
type ChanPwrList struct {
	ChanPwr []ChanPwr `json:"chan-pwr"` // Channel power configurations array (Live: IOS-XE 17.12.6a)
}

// ChanPwr represents individual channel power.
type ChanPwr struct {
	Chan int `json:"chan"` // Channel number for power assignment (Live: IOS-XE 17.12.6a)
}

// SnifferCfg represents sniffer configuration.
type SnifferCfg struct {
	SnifferEnabled bool `json:"sniffer-enabled"` // Packet capture functionality enablement (Live: IOS-XE 17.12.6a)
}

// RadioBandInfo represents radio band information.
type RadioBandInfo struct {
	BandID                 uint8          `json:"band-id"`                      // RF band identifier (2.4/5/6GHz) (Live: IOS-XE 17.12.6a)
	RegDomainCode          uint16         `json:"reg-domain-code"`              // Regulatory domain code (Live: IOS-XE 17.12.6a)
	RegulatoryDomain       string         `json:"regulatory-domain"`            // Regulatory domain name (Live: IOS-XE 17.12.6a)
	MACOperCfg             MACOperCfg     `json:"mac-oper-cfg,omitempty"`       // MAC layer operational configuration (Live: IOS-XE 17.12.6a)
	PhyTxPwrCfg            PhyTxPwrCfg    `json:"phy-tx-pwr-cfg,omitempty"`     // PHY layer TX power configuration (Live: IOS-XE 17.12.6a)
	PhyTxPwrLvlCfg         PhyTxPwrLvlCfg `json:"phy-tx-pwr-lvl-cfg,omitempty"` // Multi-level TX power configuration (Live: IOS-XE 17.12.6a)
	AntennaCfg             AntennaCfg     `json:"antenna-cfg,omitempty"`        // Antenna system configuration (Live: IOS-XE 17.12.6a)
	Dot11acChannelWidthCap uint8          `json:"dot11ac-channel-width-cap"`    // 802.11ac max channel width capability (Live: IOS-XE 17.12.6a)
	Secondary80Channel     uint16         `json:"secondary-80-channel"`         // Secondary 80MHz channel for 160MHz VHT (Live: IOS-XE 17.12.6a)
	SiaParams              SiaParams      `json:"sia-params,omitempty"`         // Self-Identifying Antenna parameters (Live: IOS-XE 17.12.6a)
}

// MACOperCfg represents MAC operation configuration.
type MACOperCfg struct {
	CfgData MACOperCfgData `json:"cfg-data"` // MAC layer operational data (Live: IOS-XE 17.12.6a)
}

// MACOperCfgData represents MAC operation configuration data.
type MACOperCfgData struct {
	MACOperationConfigType string `json:"mac-operation-config-type"` // MAC operation configuration type (Live: IOS-XE 17.12.6a)
	RtsThreshold           uint16 `json:"rts-threshold"`             // RTS threshold in bytes (Live: IOS-XE 17.12.6a)
	ShortRetryLimit        uint8  `json:"short-retry-limit"`         // Max retry attempts for short frames (Live: IOS-XE 17.12.6a)
	LongRetryLimit         uint8  `json:"long-retry-limit"`          // Max retry attempts for long frames (Live: IOS-XE 17.12.6a)
	FragThreshold          uint16 `json:"frag-threshold"`            // Frame fragmentation threshold (Live: IOS-XE 17.12.6a)
	MaxTxLifeTime          uint16 `json:"max-tx-life-time"`          // Maximum frame TX lifetime (Live: IOS-XE 17.12.6a)
	MaxRxLifeTime          uint16 `json:"max-rx-life-time"`          // Max frame RX lifetime (Live: IOS-XE 17.12.6a)
}

// PhyTxPwrCfg represents PHY TX power configuration.
type PhyTxPwrCfg struct {
	CfgData PhyTxPwrCfgData `json:"cfg-data"` // PHY layer TX power configuration (Live: IOS-XE 17.12.6a)
}

// PhyTxPwrCfgData represents PHY TX power configuration data.
type PhyTxPwrCfgData struct {
	PhyTxPowerConfigType string `json:"phy-tx-power-config-type"` // PHY TX power configuration type (Live: IOS-XE 17.12.6a)
	CurrentTxPowerLevel  uint8  `json:"current-tx-power-level"`   // Current TX power level index (Live: IOS-XE 17.12.6a)
}

// PhyTxPwrLvlCfg represents PHY TX power level configuration.
type PhyTxPwrLvlCfg struct {
	CfgData PhyTxPwrLvlCfgData `json:"cfg-data"` // Multi-level TX power configuration (Live: IOS-XE 17.12.6a)
}

// PhyTxPwrLvlCfgData represents PHY TX power level configuration data.
type PhyTxPwrLvlCfgData struct {
	NumSuppPowerLevels uint8 `json:"num-supp-power-levels"` // Number of supported power levels (Live: IOS-XE 17.12.6a)
	TxPowerLevel1      int8  `json:"tx-power-level-1"`      // TX power level 1 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel2      int8  `json:"tx-power-level-2"`      // TX power level 2 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel3      int8  `json:"tx-power-level-3"`      // TX power level 3 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel4      int8  `json:"tx-power-level-4"`      // TX power level 4 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel5      int8  `json:"tx-power-level-5"`      // TX power level 5 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel6      int8  `json:"tx-power-level-6"`      // TX power level 6 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel7      int8  `json:"tx-power-level-7"`      // TX power level 7 (dBm) (Live: IOS-XE 17.12.6a)
	TxPowerLevel8      int8  `json:"tx-power-level-8"`      // TX power level 8 (dBm) (Live: IOS-XE 17.12.6a)
	CurrTxPowerInDbm   int8  `json:"curr-tx-power-in-dbm"`  // Current active TX power (dBm) (Live: IOS-XE 17.12.6a)
}

// AntennaCfg represents antenna configuration.
type AntennaCfg struct {
	CfgData AntennaCfgData `json:"cfg-data"` // Antenna system configuration data (Live: IOS-XE 17.12.6a)
}

// AntennaCfgData represents antenna configuration data.
type AntennaCfgData struct {
	DiversitySelection string `json:"diversity-selection"` // Antenna diversity selection algorithm (Live: IOS-XE 17.12.6a)
	AntennaMode        string `json:"antenna-mode"`        // Antenna operational mode (Live: IOS-XE 17.12.6a)
	NumOfAntennas      uint8  `json:"num-of-antennas"`     // Number of physical antenna elements (Live: IOS-XE 17.12.6a)
}

// SiaParams represents Self Identifying Antenna parameters.
type SiaParams struct {
	IsRptncPresent bool   `json:"is-rptnc-present"` // Reverse Polarity TNC connector presence (Live: IOS-XE 17.12.6a)
	IsDartPresent  bool   `json:"is-dart-present"`  // DART technology presence (Live: IOS-XE 17.12.6a)
	AntennaIfType  string `json:"antenna-if-type"`  // Antenna interface type (Live: IOS-XE 17.12.6a)
	AntennaGain    uint8  `json:"antenna-gain"`     // Antenna gain value (dBi) (Live: IOS-XE 17.12.6a)
	Marlin4Present bool   `json:"marlin4-present"`  // Marlin4 antenna module presence (Live: IOS-XE 17.12.6a)
	DmServType     string `json:"dm-serv-type"`     // Device management service type (Live: IOS-XE 17.12.6a)
}

// VapOperConfig represents VAP operational configuration.
type VapOperConfig struct {
	ApVapID         uint8  `json:"ap-vap-id"`         // Virtual Access Point identifier (Live: IOS-XE 17.12.6a)
	WlanID          uint8  `json:"wlan-id"`           // Wireless LAN identifier (Live: IOS-XE 17.12.6a)
	BssidMAC        string `json:"bssid-mac"`         // BSS Identifier MAC address (Live: IOS-XE 17.12.6a)
	WtpMAC          string `json:"wtp-mac"`           // Wireless Termination Point MAC address (Live: IOS-XE 17.12.6a)
	WlanProfileName string `json:"wlan-profile-name"` // WLAN profile name (Live: IOS-XE 17.12.6a)
	SSID            string `json:"ssid"`              // Service Set Identifier (Live: IOS-XE 17.12.6a)
}

// DualRadioModeCfg represents dual radio mode configuration.
type DualRadioModeCfg struct {
	DualRadioMode    string `json:"dual-radio-mode"`    // Dual radio operational mode (Live: IOS-XE 17.12.6a)
	DualRadioCapable string `json:"dual-radio-capable"` // Dual radio hardware capability (Live: IOS-XE 17.12.6a)
	DualRadioModeOp  string `json:"dual-radio-mode-op"` // Dual radio operational state (Live: IOS-XE 17.12.6a)
}

// BssColorCfg represents BSS color configuration.
type BssColorCfg struct {
	BssColorCapable    bool   `json:"bss-color-capable"`     // 802.11ax BSS color capability (Live: IOS-XE 17.12.6a)
	BssColor           uint8  `json:"bss-color"`             // 802.11ax BSS color ID (1-63) (Live: IOS-XE 17.12.6a)
	BssColorConfigType string `json:"bss-color-config-type"` // BSS color configuration type (Live: IOS-XE 17.12.6a)
}

// BoardDataOpt represents board data options.
type BoardDataOpt struct {
	JoinPriority uint8 `json:"join-priority"` // Controller join priority value (Live: IOS-XE 17.12.6a)
}

// DescriptorData represents descriptor data.
type DescriptorData struct {
	RadioSlotsInUse        uint8 `json:"radio-slots-in-use"`      // Number of active radio slots (Live: IOS-XE 17.12.6a)
	EncryptionCapabilities bool  `json:"encryption-capabilities"` // Hardware encryption capability (Live: IOS-XE 17.12.6a)
}

// ApProv represents AP provisioning information.
type ApProv struct {
	IsUniversal          bool   `json:"is-universal"`           // Universal AP provisioning status (Live: IOS-XE 17.12.6a)
	UniversalPrimeStatus string `json:"universal-prime-status"` // Universal Prime licensing status (Live: IOS-XE 17.12.6a)
}

// ApModels represents AP model information.
type ApModels struct {
	Model string `json:"model"` // Access point hardware model (Live: IOS-XE 17.12.6a)
}

// TempInfo represents temperature information.
type TempInfo struct {
	Degree       int    `json:"degree"`        // Temperature in degrees Celsius (Live: IOS-XE 17.12.6a)
	TempStatus   string `json:"temp-status"`   // Temperature operational status (Live: IOS-XE 17.12.6a)
	HeaterStatus string `json:"heater-status"` // Internal heater operational status (Live: IOS-XE 17.12.6a)
}

// TagInfo represents AP tag information.
type TagInfo struct {
	TagSource         string          `json:"tag-source"`          // Tag assignment source methodology (Live: IOS-XE 17.12.6a)
	IsApMisconfigured bool            `json:"is-ap-misconfigured"` // AP misconfiguration detection (Live: IOS-XE 17.12.6a)
	ResolvedTagInfo   ResolvedTagInfo `json:"resolved-tag-info"`   // Final resolved tag assignments (Live: IOS-XE 17.12.6a)
	PolicyTagInfo     PolicyTagInfo   `json:"policy-tag-info"`     // Policy tag configuration (Live: IOS-XE 17.12.6a)
	SiteTag           SiteTag         `json:"site-tag"`            // Site tag information (Live: IOS-XE 17.12.6a)
	RFTag             RFTag           `json:"rf-tag"`              // RF tag configuration (Live: IOS-XE 17.12.6a)
	FilterInfo        FilterInfo      `json:"filter-info"`         // Access control filter information (Live: IOS-XE 17.12.6a)
	IsDTLSLscFbkAp    bool            `json:"is-dtls-lsc-fbk-ap"`  // DTLS LSC fallback AP designation (Live: IOS-XE 17.12.6a)
}

// ResolvedTagInfo represents resolved tag information.
type ResolvedTagInfo struct {
	ResolvedPolicyTag string `json:"resolved-policy-tag"` // Final resolved policy tag name (Live: IOS-XE 17.12.6a)
	ResolvedSiteTag   string `json:"resolved-site-tag"`   // Final resolved site tag name (Live: IOS-XE 17.12.6a)
	ResolvedRFTag     string `json:"resolved-rf-tag"`     // Final resolved RF tag name (Live: IOS-XE 17.12.6a)
}

// PolicyTagInfo represents policy tag information.
type PolicyTagInfo struct {
	PolicyTagName string `json:"policy-tag-name"` // Policy tag name identifier (Live: IOS-XE 17.12.6a)
}

// SiteTag represents site tag information.
type SiteTag struct {
	SiteTagName string `json:"site-tag-name"` // Site tag name identifier (Live: IOS-XE 17.12.6a)
	ApProfile   string `json:"ap-profile"`    // AP profile name (Live: IOS-XE 17.12.6a)
	FlexProfile string `json:"flex-profile"`  // FlexConnect profile name (Live: IOS-XE 17.12.6a)
}

// RFTag represents RF tag information.
type RFTag struct {
	RFTagName string `json:"rf-tag-name"` // RF tag name identifier (Live: IOS-XE 17.12.6a)
}

// FilterInfo represents filter information.
type FilterInfo struct {
	FilterName string `json:"filter-name"` // Access control filter name (Live: IOS-XE 17.12.6a)
}

// Tunnel represents tunnel configuration.
type Tunnel struct {
	PreferredMode string `json:"preferred-mode"` // Preferred CAPWAP tunnel mode (Live: IOS-XE 17.12.6a)
	UDPLite       string `json:"udp-lite"`       // UDP-Lite protocol configuration (Live: IOS-XE 17.12.6a)
}

// ExternalModuleData represents external module data.
type ExternalModuleData struct {
	XmData             XmData  `json:"xm-data"`               // External module data (Live: IOS-XE 17.12.6a)
	UsbData            UsbData `json:"usb-data"`              // USB module data (Live: IOS-XE 17.12.6a)
	UsbOverride        bool    `json:"usb-override"`          // USB configuration override (Live: IOS-XE 17.12.6a)
	IsExtModuleEnabled bool    `json:"is-ext-module-enabled"` // External module enablement status (Live: IOS-XE 17.12.6a)
}

// XmData represents XM module data.
type XmData struct {
	IsModulePresent bool `json:"is-module-present"` // External module presence detection (Live: IOS-XE 17.12.6a)
	Xm              Xm   `json:"xm"`                // External module detailed information (Live: IOS-XE 17.12.6a)
}

// UsbData represents USB module data.
type UsbData struct {
	IsModulePresent bool `json:"is-module-present"` // USB module presence detection (Live: IOS-XE 17.12.6a)
	Xm              Xm   `json:"xm"`                // USB module detailed information (Live: IOS-XE 17.12.6a)
}

// Xm represents external module information.
type Xm struct {
	NumericID          uint32 `json:"numeric-id"`           // Module unique numeric identifier (Live: IOS-XE 17.12.6a)
	MaxPower           uint16 `json:"max-power"`            // Module maximum power consumption (mW) (Live: IOS-XE 17.12.6a)
	SerialNumberString string `json:"serial-number-string"` // Module serial number (Live: IOS-XE 17.12.6a)
	ProductIDString    string `json:"product-id-string"`    // Module product identifier (Live: IOS-XE 17.12.6a)
	ModuleType         string `json:"module-type"`          // Module type classification (Live: IOS-XE 17.12.6a)
	ModuleDescription  string `json:"module-description"`   // Module description (Live: IOS-XE 17.12.6a)
}

// DeviceDetail represents device detail information.
type DeviceDetail struct {
	StaticInfo  StaticInfo  `json:"static-info"`  // Static hardware and firmware information (Live: IOS-XE 17.12.6a)
	DynamicInfo DynamicInfo `json:"dynamic-info"` // Dynamic operational status (Live: IOS-XE 17.12.6a)
	WtpVersion  WtpVersion  `json:"wtp-version"`  // WTP software version details (Live: IOS-XE 17.12.6a)
}

// StaticInfo represents static information.
type StaticInfo struct {
	BoardData struct {
		WtpSerialNum string `json:"wtp-serial-num"` // AP serial number (Live: IOS-XE 17.12.6a)
		WtpEnetMAC   string `json:"wtp-enet-mac"`   // AP Ethernet MAC address (Live: IOS-XE 17.12.6a)
		ApSysInfo    struct {
			MemType string `json:"mem-type"` // AP memory type (Live: IOS-XE 17.12.6a)
			CPUType string `json:"cpu-type"` // AP CPU type (Live: IOS-XE 17.12.6a)
			MemSize int    `json:"mem-size"` // AP memory size (Live: IOS-XE 17.12.6a)
		} `json:"ap-sys-info"` // AP system info (Live: IOS-XE 17.12.6a)
	} `json:"board-data"` // AP Board Data (Live: IOS-XE 17.12.6a)
	BoardDataOpt   BoardDataOpt   `json:"board-data-opt,omitempty"`  // AP Additional Board data option (Live: IOS-XE 17.12.6a)
	DescriptorData DescriptorData `json:"descriptor-data,omitempty"` // AP FW,HW information (Live: IOS-XE 17.12.6a)
	ApProv         ApProv         `json:"ap-prov,omitempty"`         // AP universal provision (Live: IOS-XE 17.12.6a)
	ApModels       ApModels       `json:"ap-models,omitempty"`       // AP device model (Live: IOS-XE 17.12.6a)
	NumPorts       uint8          `json:"num-ports,omitempty"`       // Number of ports on AP (Live: IOS-XE 17.12.6a)
	NumSlots       uint8          `json:"num-slots,omitempty"`       // Number of slots present in the access point (Live: IOS-XE 17.12.6a)
	WtpModelType   uint16         `json:"wtp-model-type,omitempty"`  // AP model type (Live: IOS-XE 17.12.6a)
	ApCapability   string         `json:"ap-capability,omitempty"`   // AP capabilities (Live: IOS-XE 17.12.6a)
	IsMmOpt        bool           `json:"is-mm-opt,omitempty"`       // AP monitor mode optimization support (Live: IOS-XE 17.12.6a)
	ApImageName    string         `json:"ap-image-name,omitempty"`   // AP Software image name (Live: IOS-XE 17.12.6a)
}

// DynamicInfo represents dynamic information.
type DynamicInfo struct {
	ApCrashData struct {
		ApCrashFile           string `json:"ap-crash-file"`              // AP crash file (Live: IOS-XE 17.12.6a)
		ApRadio2GCrashFile    string `json:"ap-radio-2g-crash-file"`     // AP 2 GHz radio crash file (Live: IOS-XE 17.12.6a)
		ApRadio5GCrashFile    string `json:"ap-radio-5g-crash-file"`     // AP 5 GHz radio crash file (Live: IOS-XE 17.12.6a)
		ApRadio6GCrashFile    string `json:"ap-radio-6g-crash-file"`     // AP 6 GHz radio crash file (Live: IOS-XE 17.12.6a)
		ApRad5GSlot2CrashFile string `json:"ap-rad-5g-slot2-crash-file"` // AP 5 GHz radio slot 2 crash file (Live: IOS-XE 17.12.6a)
	} `json:"ap-crash-data"` // AP crash data (Live: IOS-XE 17.12.6a)
	LedStateEnabled  bool     `json:"led-state-enabled,omitempty"`  // True if LED state of AP is enabled (Live: IOS-XE 17.12.6a)
	ResetButtonState bool     `json:"reset-button-state,omitempty"` // True if AP Reset button state is enabled (Live: IOS-XE 17.12.6a)
	LedFlashEnabled  bool     `json:"led-flash-enabled,omitempty"`  // True if LED Flash state of AP is enabled (Live: IOS-XE 17.12.6a)
	FlashSec         uint16   `json:"flash-sec,omitempty"`          // LED Flash timer duration for AP in seconds (Live: IOS-XE 17.12.6a)
	TempInfo         TempInfo `json:"temp-info,omitempty"`          // AP temperature info (Live: IOS-XE 17.12.6a)
	LedFlashExpiry   string   `json:"led-flash-expiry,omitempty"`   // Led Flash Expiry Date and Time (Live: IOS-XE 17.12.6a)
}

// WtpVersion represents WTP version information.
type WtpVersion struct {
	BackupSwVersion struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.6a)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.6a)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.6a)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.6a)
	} `json:"backup-sw-version"` // Backup software version of the AP (Live: IOS-XE 17.12.6a)
	MiniIosVersion struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.6a)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.6a)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.6a)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.6a)
	} `json:"mini-ios-version,omitempty"` // Cisco AP mini IOS version details (Live: IOS-XE 17.12.6a)
	SwVer struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.6a)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.6a)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.6a)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.6a)
	} `json:"sw-ver,omitempty"` // Software version of the AP (Live: IOS-XE 17.12.6a)
	BootVer struct {
		Version int `json:"version"` // Version number (Live: IOS-XE 17.12.6a)
		Release int `json:"release"` // Release number (Live: IOS-XE 17.12.6a)
		Maint   int `json:"maint"`   // Maintenance number (Live: IOS-XE 17.12.6a)
		Build   int `json:"build"`   // Build number (Live: IOS-XE 17.12.6a)
	} `json:"boot-ver,omitempty"` // Cisco AP boot version details (Live: IOS-XE 17.12.6a)
	SwVersion string `json:"sw-version,omitempty"` // Cisco AP software version details (Live: IOS-XE 17.12.6a)
}

// ApNameMACMap represents AP name to MAC address mapping.
type ApNameMACMap struct {
	WtpName string `json:"wtp-name"` // WTP administrative name (Live: IOS-XE 17.12.6a)
	WtpMAC  string `json:"wtp-mac"`  // WTP radio interface MAC address (Live: IOS-XE 17.12.6a)
	EthMAC  string `json:"eth-mac"`  // Ethernet interface MAC address (Live: IOS-XE 17.12.6a)
}

// WtpSlotWlanStats represents WTP slot WLAN statistics.
type WtpSlotWlanStats struct {
	WtpMAC      string `json:"wtp-mac"`      // WTP MAC address for radio interface (Live: IOS-XE 17.12.6a)
	SlotID      int    `json:"slot-id"`      // Radio slot identifier (Live: IOS-XE 17.12.6a)
	WlanID      int    `json:"wlan-id"`      // WLAN identifier (Live: IOS-XE 17.12.6a)
	BssidMAC    string `json:"bssid-mac"`    // BSS Identifier MAC address (Live: IOS-XE 17.12.6a)
	Ssid        string `json:"ssid"`         // Service Set Identifier name (Live: IOS-XE 17.12.6a)
	BytesRx     string `json:"bytes-rx"`     // Total bytes received on WLAN interface (Live: IOS-XE 17.12.6a)
	BytesTx     string `json:"bytes-tx"`     // Total bytes transmitted on WLAN interface (Live: IOS-XE 17.12.6a)
	PktsRx      string `json:"pkts-rx"`      // Total packets received on WLAN interface (Live: IOS-XE 17.12.6a)
	PktsTx      string `json:"pkts-tx"`      // Total packets transmitted on WLAN interface (Live: IOS-XE 17.12.6a)
	DataRetries string `json:"data-retries"` // Data frame retransmission count (Live: IOS-XE 17.12.6a)
}

// EthernetMACWtpMACMap represents Ethernet MAC to WTP MAC mapping.
type EthernetMACWtpMACMap struct {
	EthernetMAC string `json:"ethernet-mac"` // Ethernet interface MAC address (Live: IOS-XE 17.12.6a)
	WtpMAC      string `json:"wtp-mac"`      // WTP MAC address for radio interface (Live: IOS-XE 17.12.6a)
}

// ApIoxOperData represents AP IOx operational data.
type ApIoxOperData struct {
	ApMAC        string `json:"ap-mac"`        // Access point MAC address (Live: IOS-XE 17.12.6a)
	ApphostState string `json:"apphost-state"` // Application hosting service state (Live: IOS-XE 17.12.6a)
	CafToken     string `json:"caf-token"`     // CAF authentication token (Live: IOS-XE 17.12.6a)
	CafPort      int    `json:"caf-port"`      // CAF service communication port (Live: IOS-XE 17.12.6a)
}

// QosGlobalStats represents QoS global statistics.
type QosGlobalStats struct {
	QosClientVoiceStats struct {
		TotalNumOfTspecRcvd       int `json:"total-num-of-tspec-rcvd"`        // Total Number of TSPEC requests received (Live: IOS-XE 17.12.6a)
		NewTspecFromAssocReq      int `json:"new-tspec-from-assoc-req"`       // Number of New TSPEC received from Assoc Request (Live: IOS-XE 17.12.6a)
		TspecRenewalFromAssocReq  int `json:"tspec-renewal-from-assoc-req"`   // TSPEC renewal from Assoc Request (Live: IOS-XE 17.12.6a)
		NewTspecAsAddTS           int `json:"new-tspec-as-add-ts"`            // Number of new Add TS requests received (Live: IOS-XE 17.12.6a)
		TspecRenewalFromAddTS     int `json:"tspec-renewal-from-add-ts"`      // Number of Add TS renewal requests received (Live: IOS-XE 17.12.6a)
		NumOfActiveTspecCalls     int `json:"num-of-active-tspec-calls"`      // Total Number of active TSPEC calls (Live: IOS-XE 17.12.6a)
		NumOfActiveSIPCalls       int `json:"num-of-active-sip-calls"`        // Total Number of active SIP calls (Live: IOS-XE 17.12.6a)
		NumOfCallsAccepted        int `json:"num-of-calls-accepted"`          // Total Number of calls accepted (Live: IOS-XE 17.12.6a)
		NumOfCallsRejectedInsufBw int `json:"num-of-calls-rejected-insuf-bw"` // Number of calls rejected due to Insufficient BW (Live: IOS-XE 17.12.6a)
		NumOfCallsRejectedPhyRate int `json:"num-of-calls-rejected-phy-rate"` // Number of calls rejected due to PHY rate (Live: IOS-XE 17.12.6a)
		NumOfCallsRejectedQos     int `json:"num-of-calls-rejected-qos"`      // Number of calls rejected due to QoS policy (Live: IOS-XE 17.12.6a)
		NumOfCallsRejInvalidTspec int `json:"num-of-calls-rej-invalid-tspec"` // Number of calls rejected due to Invalid TSPEC (Live: IOS-XE 17.12.6a)
		NumOfRoamCallsAccepted    int `json:"num-of-roam-calls-accepted"`     // Total Number of roam calls accepted (Live: IOS-XE 17.12.6a)
		NumOfRoamCallsRejected    int `json:"num-of-roam-calls-rejected"`     // Total Number of roam calls rejected (Live: IOS-XE 17.12.6a)
		TspecProcessFailedGetRec  int `json:"tspec-process-failed-get-rec"`   // Number of DB failures while processing TSPEC (Live: IOS-XE 17.12.6a)
		TotalNumOfCallReport      int `json:"total-num-of-call-report"`       // Total number of call-report received (Live: IOS-XE 17.12.6a)
		TotalSIPFailureTrapSend   int `json:"total-sip-failure-trap-send"`    // Total number of SIP failure trap send (Live: IOS-XE 17.12.6a)
		TotalSIPInviteOnCaller    int `json:"total-sip-invite-on-caller"`     // Total number of SIP Invite received on Caller (Live: IOS-XE 17.12.6a)
		TotalSIPInviteOnCallee    int `json:"total-sip-invite-on-callee"`     // Total number of SIP Invite received on Callee (Live: IOS-XE 17.12.6a)
	} `json:"qos-client-voice-stats"`
}

// ApOperInternalData represents internal AP operational data.
type ApOperInternalData struct {
	WtpMAC                 string                  `json:"wtp-mac"`                             // MAC Address of the AP Radio (Live: IOS-XE 17.12.6a)
	RadioID                int                     `json:"radio-id"`                            // AP radio identifier (Live: IOS-XE 17.12.6a)
	ApAntennaBandMode      string                  `json:"ap-antenna-band-mode"`                // AP antenna band mode configuration (Live: IOS-XE 17.12.6a)
	LinkEncryptionEnabled  bool                    `json:"link-encryption-enabled"`             // Controller-AP link encryption status (Live: IOS-XE 17.12.6a)
	ApRemoteDebugMode      bool                    `json:"ap-remote-debug-mode"`                // Remote debugging status for the AP (Live: IOS-XE 17.12.6a)
	ApRole                 string                  `json:"ap-role"`                             // AP role in PMK push (Live: IOS-XE 17.12.6a)
	ApIndoorMode           bool                    `json:"ap-indoor-mode"`                      // Identifier for indoor AP mode (Live: IOS-XE 17.12.6a)
	MaxClientsAllowed      int                     `json:"max-clients-allowed"`                 // Maximum clients allowed on an AP (Live: IOS-XE 17.12.6a)
	IsLocalNet             bool                    `json:"is-local-net"`                        // Identifier for local access in OEAP AP (Live: IOS-XE 17.12.6a)
	Ipv4TcpMss             TCPMssConfig            `json:"ipv4-tcp-mss"`                        // Configured IPv4 TCP MSS value for client (Live: IOS-XE 17.12.6a)
	Ipv6TcpMss             TCPMssConfig            `json:"ipv6-tcp-mss"`                        // Configured IPv6 TCP MSS value for client (Live: IOS-XE 17.12.6a)
	RangingMode            string                  `json:"ranging-mode"`                        // Ranging mode - normal or accurate (Live: IOS-XE 17.12.6a)
	PowerProfile           string                  `json:"power-profile"`                       // Power profile applied to the AP (Live: IOS-XE 17.12.6a)
	PwrProfType            string                  `json:"pwr-prof-type"`                       // Power profile type (Live: IOS-XE 17.12.6a)
	PwrCalProfile          string                  `json:"pwr-cal-profile"`                     // Calendar profile associated to power profile (Live: IOS-XE 17.12.6a)
	PersistentSsid         PersistentSsid          `json:"persistent-ssid"`                     // Persistent SSID broadcast operation information (Live: IOS-XE 17.12.6a)
	ProvSsid               bool                    `json:"prov-ssid"`                           // Office Extended AP Provisional SSID status (Live: IOS-XE 17.12.6a)
	PrimingProfile         string                  `json:"priming-profile"`                     // Applied AP priming profile name (Live: IOS-XE 17.12.6a)
	PrimingProfileSrc      string                  `json:"priming-profile-src"`                 // AP priming profile configuration source (Live: IOS-XE 17.12.6a)
	PrimingFilter          string                  `json:"priming-filter"`                      // AP priming filter name (Live: IOS-XE 17.12.6a)
	PmkBsSenderAddr        string                  `json:"pmk-bs-sender-addr"`                  // PMK bulk sync sender AP MAC address (Live: IOS-XE 17.12.6a)
	PmkBsReceiverAddr      string                  `json:"pmk-bs-receiver-addr"`                // PMK bulk sync receiver AP MAC address (Live: IOS-XE 17.12.6a)
	Accounting             *AccountingInfo         `json:"accounting,omitempty"`                // Accounting info to be sent to radius server (Live: IOS-XE 17.12.6a)
	ApDnaData              *ApDnaData              `json:"ap-dna-data,omitempty"`               // Cisco-DNA related data (Live: IOS-XE 17.12.6a)
	ApGasRateLimitCfg      *ApGasRateLimitConfig   `json:"ap-gas-rate-limit-cfg,omitempty"`     // Cisco AP Generic Advertisement Services (GAS) rate configuration (Live: IOS-XE 17.12.6a)
	ApIPData               *ApIPData               `json:"ap-ip-data,omitempty"`                // AP IP address configuration (Live: IOS-XE 17.12.6a)
	ApLoginCredentials     *ApLoginCredentials     `json:"ap-login-credentials,omitempty"`      // Login credentials configured on an AP (Live: IOS-XE 17.12.6a)
	ApMgmt                 *ApManagement           `json:"ap-mgmt,omitempty"`                   // AP management data (Live: IOS-XE 17.12.6a)
	ApNtpServerInfoCfg     *ApNtpServerInfo        `json:"ap-ntp-server-info-cfg,omitempty"`    // NTP server information to be used by AP (Live: IOS-XE 17.12.6a)
	ApNtpSyncStatus        *ApNtpSyncStatus        `json:"ap-ntp-sync-status,omitempty"`        // AP NTP synchronization status (Live: IOS-XE 17.12.6a)
	ApPmkPropagationStatus *bool                   `json:"ap-pmk-propagation-status,omitempty"` // AP PMK push propagation status (Live: IOS-XE 17.12.6a)
	ApPow                  *ApPowData              `json:"ap-pow,omitempty"`                    // AP power related data (Live: IOS-XE 17.12.6a)
	ApPrimeInfo            *ApPrimeInfo            `json:"ap-prime-info,omitempty"`             // Controller configuration for the AP (Live: IOS-XE 17.12.6a)
	ApPrimingOverride      *bool                   `json:"ap-priming-override,omitempty"`       // AP priming override flag status (Live: IOS-XE 17.12.6a)
	ApSysStats             *ApSystemStats          `json:"ap-sys-stats,omitempty"`              // AP system statistics (Live: IOS-XE 17.12.6a)
	ApTzConfig             *ApTimezoneConfig       `json:"ap-tz-config,omitempty"`              // AP timezone configuration (Live: IOS-XE 17.12.6a)
	ApUdpliteInfo          *string                 `json:"ap-udplite-info,omitempty"`           // UDP-Lite operational information (Live: IOS-XE 17.12.6a)
	AuxClientInterfaceData *AuxClientInterfaceData `json:"aux-client-interface-data,omitempty"` // Auxiliary Client Interface data (Live: IOS-XE 17.12.6a)
	InfrastructureMfp      *InfrastructureMfp      `json:"infrastructure-mfp,omitempty"`        // Cisco AP Management Frame Protection (Live: IOS-XE 17.12.6a)
	KernelCoredump         *KernelCoredumpConfig   `json:"kernel-coredump,omitempty"`           // Kernel coredump configuration (Live: IOS-XE 17.12.6a)
	LinkAudit              *LinkAuditData          `json:"link-audit,omitempty"`                // Link audit data (Live: IOS-XE 17.12.6a)
	OeapAudit              *OeapAuditData          `json:"oeap-audit,omitempty"`                // On-demand Office Extended AP link test data (Live: IOS-XE 17.12.6a)
	Retransmit             *RetransmitConfig       `json:"retransmit,omitempty"`                // AP retransmission parameters (Live: IOS-XE 17.12.6a)
	Syslog                 *SyslogConfig           `json:"syslog,omitempty"`                    // Cisco AP System Logging (Live: IOS-XE 17.12.6a)
	Timer                  *ApTimerConfig          `json:"timer,omitempty"`                     // Cisco access point timer data (Live: IOS-XE 17.12.6a)
}

// RlanOper represents RLAN operational data.
type RlanOper struct {
	WtpMAC         string `json:"wtp-mac"`          // Radio MAC address of the AP (Live: IOS-XE 17.12.6a)
	RlanPortID     int    `json:"rlan-port-id"`     // RLAN port identifier (Live: IOS-XE 17.12.6a)
	RlanOperState  bool   `json:"rlan-oper-state"`  // Status of the LAN port (Live: IOS-XE 17.12.6a)
	RlanPortStatus bool   `json:"rlan-port-status"` // Remote LAN status of the LAN port (Live: IOS-XE 17.12.6a)
	RlanVlanValid  bool   `json:"rlan-vlan-valid"`  // LAN port valid or not (Live: IOS-XE 17.12.6a)
	RlanVlanID     int    `json:"rlan-vlan-id"`     // VLAN id of the LAN port (Live: IOS-XE 17.12.6a)
	RlanPoeState   string `json:"rlan-poe-state"`   // PoE state of the LAN port (Live: IOS-XE 17.12.6a)
	PowerLevelID   int    `json:"power-level-id"`   // Power level of the LAN port (Live: IOS-XE 17.12.6a)
}

// EwlcMewlcPredownloadRec represents EWLC MEWLC predownload record.
type EwlcMewlcPredownloadRec struct {
	PredState                    string `json:"pred-state"`                     // Embedded Wireless Controller predownload state (Live: IOS-XE 17.12.6a)
	MeCapableApCount             int    `json:"me-capable-ap-count"`            // Total EWC capable AP count (Live: IOS-XE 17.12.6a)
	ControllerPredownloadVersion string `json:"controller-predownload-version"` // Embedded Wireless Controller predownload version (Live: IOS-XE 17.12.6a)
}

// CdpCacheData represents CDP cache data.
type CdpCacheData struct {
	MACAddr                string       `json:"mac-addr"`                   // MAC address (Live: IOS-XE 17.12.6a)
	CdpCacheDeviceID       string       `json:"cdp-cache-device-id"`        // CDP device identifier (Live: IOS-XE 17.12.6a)
	ApName                 string       `json:"ap-name"`                    // AP Name (Live: IOS-XE 17.12.6a)
	LastUpdatedTime        time.Time    `json:"last-updated-time"`          // Last updated time (Live: IOS-XE 17.12.6a)
	Version                int          `json:"version"`                    // Cisco Discovery Protocol version (Live: IOS-XE 17.12.6a)
	WtpMACAddr             string       `json:"wtp-mac-addr"`               // WTP MAC address (Live: IOS-XE 17.12.6a)
	DeviceIndex            int          `json:"device-index"`               // Device index (Live: IOS-XE 17.12.6a)
	IPAddress              CdpIPAddress `json:"ip-address"`                 // Device network addresses from CDP message (Live: IOS-XE 17.12.6a)
	CdpAddrCount           int          `json:"cdp-addr-count"`             // Neighbor IP count (Live: IOS-XE 17.12.6a)
	CdpCacheApAddress      string       `json:"cdp-cache-ap-address"`       // CDP cache address type for the AP (Live: IOS-XE 17.12.6a)
	CdpCacheDevicePort     string       `json:"cdp-cache-device-port"`      // Device outgoing port (Live: IOS-XE 17.12.6a)
	CdpCacheDuplex         string       `json:"cdp-cache-duplex"`           // CDP cache duplex type (Live: IOS-XE 17.12.6a)
	CdpCacheIfIndex        int          `json:"cdp-cache-if-index"`         // CDP cache interface index (Live: IOS-XE 17.12.6a)
	CdpCacheInterfaceSpeed int          `json:"cdp-cache-interface-speed"`  // CDP cache interface speed (Live: IOS-XE 17.12.6a)
	CdpCacheIPAddressValue string       `json:"cdp-cache-ip-address-value"` // Entry address(es) (Live: IOS-XE 17.12.6a)
	CdpCacheLocalPort      string       `json:"cdp-cache-local-port"`       // Device interface port (Live: IOS-XE 17.12.6a)
	CdpCachePlatform       string       `json:"cdp-cache-platform"`         // CDP cache platform (Live: IOS-XE 17.12.6a)
	CdpCacheVersion        string       `json:"cdp-cache-version"`          // CDP cache version (Live: IOS-XE 17.12.6a)
	CdpCapabilitiesString  string       `json:"cdp-capabilities-string"`    // CDP cache capabilities (Live: IOS-XE 17.12.6a)
}

// LldpNeigh represents LLDP neighbor information.
type LldpNeigh struct {
	WtpMAC          string `json:"wtp-mac"`          // Radio MAC address of the AP (Live: IOS-XE 17.12.6a)
	NeighMAC        string `json:"neigh-mac"`        // MAC address of the LLDP neighbor device (Live: IOS-XE 17.12.6a)
	PortID          string `json:"port-id"`          // LLDP neighbor port name or ID (Live: IOS-XE 17.12.6a)
	LocalPort       string `json:"local-port"`       // AP interface sending/receiving LLDP PDUs (Live: IOS-XE 17.12.6a)
	SystemName      string `json:"system-name"`      // LLDP neighbor name (Live: IOS-XE 17.12.6a)
	PortDescription string `json:"port-description"` // LLDP neighbor port description (Live: IOS-XE 17.12.6a)
	Capabilities    string `json:"capabilities"`     // LLDP device capabilities (Live: IOS-XE 17.12.6a)
	MgmtAddr        string `json:"mgmt-addr"`        // Management IPv4 address of LLDP neighbor (Live: IOS-XE 17.12.6a)
}

// TpCertInfo represents trustpoint certificate information.
type TpCertInfo struct {
	Trustpoint Trustpoint `json:"trustpoint"` // Trustpoint certificate information (Live: IOS-XE 17.12.6a)
}

// Trustpoint represents trustpoint information.
type Trustpoint struct {
	TrustpointName     string  `json:"trustpoint-name"`      // Trustpoint name (Live: IOS-XE 17.12.6a)
	IsCertAvailable    bool    `json:"is-cert-available"`    // Is certificate available (Live: IOS-XE 17.12.6a)
	IsPrivkeyAvailable bool    `json:"is-privkey-available"` // Is private key available (Live: IOS-XE 17.12.6a)
	CertHash           string  `json:"cert-hash"`            // Certificate hash (Live: IOS-XE 17.12.6a)
	CertType           string  `json:"cert-type"`            // Certificate type (Live: IOS-XE 17.12.6a)
	FipsSuitability    string  `json:"fips-suitability"`     // FIPS Suitability (Live: IOS-XE 17.12.6a)
	State              *string `json:"state,omitempty"`      // Trustpoint state (Live: IOS-XE 17.12.6a)
}

// DiscData represents discovery data.
type DiscData struct {
	WtpMAC           string `json:"wtp-mac"`            // Wireless termination point MAC address (Live: IOS-XE 17.12.6a)
	DiscoveryPkts    string `json:"discovery-pkts"`     // Discovery packet count (Live: IOS-XE 17.12.6a)
	DiscoveryErrPkts string `json:"discovery-err-pkts"` // Discovery error packet count (Live: IOS-XE 17.12.6a)
}

// CAPWAPPkts represents CAPWAP packet statistics.
type CAPWAPPkts struct {
	WtpMAC            string `json:"wtp-mac"`              // Wireless termination point MAC address (Live: IOS-XE 17.12.6a)
	CntrlPkts         string `json:"cntrl-pkts"`           // Control packet count (Live: IOS-XE 17.12.6a)
	DataKeepAlivePkts string `json:"data-keep-alive-pkts"` // Data keep-alive packet count (Live: IOS-XE 17.12.6a)
	CAPWAPErrorPkts   string `json:"capwap-error-pkts"`    // CAPWAP error packet count (Live: IOS-XE 17.12.6a)
	ArpPkts           string `json:"arp-pkts"`             // ARP packet count (Live: IOS-XE 17.12.6a)
	DHCPPkts          string `json:"dhcp-pkts"`            // DHCP packet count (Live: IOS-XE 17.12.6a)
	Dot1xCtrlPkts     string `json:"dot1x-ctrl-pkts"`      // 802.1X control packet count (Live: IOS-XE 17.12.6a)
	Dot1xEapPkts      string `json:"dot1x-eap-pkts"`       // 802.1X EAP packet count (Live: IOS-XE 17.12.6a)
	Dot1xKeyTypePkts  string `json:"dot1x-key-type-pkts"`  // 802.1X key type packet count (Live: IOS-XE 17.12.6a)
	Dot1xMgmtPkts     string `json:"dot1x-mgmt-pkts"`      // 802.1X management packet count (Live: IOS-XE 17.12.6a)
	IappPkts          string `json:"iapp-pkts"`            // IAPP packet count (Live: IOS-XE 17.12.6a)
	IPPkts            string `json:"ip-pkts"`              // IP packet count (Live: IOS-XE 17.12.6a)
	Ipv6Pkts          string `json:"ipv6-pkts"`            // IPv6 packet count (Live: IOS-XE 17.12.6a)
	RFIDPkts          string `json:"rfid-pkts"`            // RFID packet count (Live: IOS-XE 17.12.6a)
	RRMPkts           string `json:"rrm-pkts"`             // Radio resource management packet count (Live: IOS-XE 17.12.6a)
}

// CountryOper represents country operational data.
type CountryOper struct {
	CountryCode         string  `json:"country-code"`                 // Country code for regulatory compliance (Live: IOS-XE 17.12.6a)
	CountryString       string  `json:"country-string"`               // Country string representation (Live: IOS-XE 17.12.6a)
	RegDomainStr80211Bg string  `json:"reg-domain-str-80211bg"`       // Regulatory domain string for 802.11bg (Live: IOS-XE 17.12.6a)
	RegDomainStr80211A  string  `json:"reg-domain-str-80211a"`        // Regulatory domain string for 802.11a (Live: IOS-XE 17.12.6a)
	CountrySupported    bool    `json:"country-supported"`            // Country support status in regulatory database (Live: IOS-XE 17.12.6a)
	Channels11a         *string `json:"channels-11a,omitempty"`       // Available channels for 802.11a operation (Live: IOS-XE 17.12.6a)
	Channels11bg        *string `json:"channels-11bg,omitempty"`      // Available channels for 802.11bg operation (Live: IOS-XE 17.12.6a)
	ChannelsString11a   string  `json:"channels-string-11a"`          // Channel string representation for 802.11a (Live: IOS-XE 17.12.6a)
	ChannelsString11bg  string  `json:"channels-string-11bg"`         // Channel string representation for 802.11bg (Live: IOS-XE 17.12.6a)
	DCAChannels11a      *string `json:"dca-channels-11a,omitempty"`   // DCA channels for 802.11a band (Live: IOS-XE 17.12.6a)
	DCAChannels11bg     *string `json:"dca-channels-11bg,omitempty"`  // DCA channels for 802.11bg band (Live: IOS-XE 17.12.6a)
	RadarChannels11a    *string `json:"radar-channels-11a,omitempty"` // Radar-affected channels for 802.11a (Live: IOS-XE 17.12.6a)
	RegDom6ghz          *string `json:"reg-dom-6ghz,omitempty"`       // Regulatory domain information for 6GHz (Live: IOS-XE 17.12.6a)
	ChanInfo6ghz        *string `json:"chan-info-6ghz,omitempty"`     // Channel information for 6GHz band (Live: IOS-XE 17.12.6a)
}

// SuppCountryOper represents supported country operational data.
type SuppCountryOper struct {
	CountryCode      string          `json:"country-code"`                  // Supported country code for regulatory compliance (Live: IOS-XE 17.12.6a)
	CountryString    string          `json:"country-string"`                // Supported country string representation (Live: IOS-XE 17.12.6a)
	CountryCodeIso   string          `json:"country-code-iso"`              // ISO standard country code (Live: IOS-XE 17.12.6a)
	ChanList24ghz    *ChannelList    `json:"chan-list-24ghz,omitempty"`     // Channel list for 2.4GHz band (Live: IOS-XE 17.12.6a)
	ChanList5ghz     *ChannelList    `json:"chan-list-5ghz,omitempty"`      // Channel list for 5GHz band (Live: IOS-XE 17.12.6a)
	ChanList6ghz     *ChannelList    `json:"chan-list-6ghz,omitempty"`      // Channel list for 6GHz band (Live: IOS-XE 17.12.6a)
	ChanListDCA24ghz *ChannelList    `json:"chan-list-dca-24ghz,omitempty"` // DCA channel list for 2.4GHz band (Live: IOS-XE 17.12.6a)
	ChanListDCA5ghz  *ChannelList    `json:"chan-list-dca-5ghz,omitempty"`  // DCA channel list for 5GHz band (Live: IOS-XE 17.12.6a)
	ChanListDCA6ghz  *ChannelList    `json:"chan-list-dca-6ghz,omitempty"`  // DCA channel list for 6GHz band (Live: IOS-XE 17.12.6a)
	ChanListPsc6ghz  *ChannelList    `json:"chan-list-psc-6ghz,omitempty"`  // PSC channel list for 6GHz band (Live: IOS-XE 17.12.6a)
	RegDom24ghz      *RegDomainCodes `json:"reg-dom-24ghz,omitempty"`       // Regulatory domain for 2.4GHz band (Live: IOS-XE 17.12.6a)
	RegDom5ghz       *RegDomainCodes `json:"reg-dom-5ghz,omitempty"`        // Regulatory domain for 5GHz band (Live: IOS-XE 17.12.6a)
	RegDom6ghz       *RegDomainCodes `json:"reg-dom-6ghz,omitempty"`        // Regulatory domain for 6GHz band (Live: IOS-XE 17.12.6a)
}

// ApNhGlobalData represents AP neighborhood global data.
type ApNhGlobalData struct {
	AlgorithmRunning   bool `json:"algorithm-running"`     // Status of neighborhood algorithm execution (Live: IOS-XE 17.12.6a)
	AlgorithmItrCount  int  `json:"algorithm-itr-count"`   // Total AP neighborhood algorithm iteration count (Live: IOS-XE 17.12.6a)
	IdealCapacityPerRg int  `json:"ideal-capacity-per-rg"` // Ideal capacity of APs per resource group (Live: IOS-XE 17.12.6a)
	NumOfNeighborhood  int  `json:"num-of-neighborhood"`   // Total number of calculated neighborhood areas (Live: IOS-XE 17.12.6a)
}

// ApImagePrepareLocation represents AP image prepare location.
type ApImagePrepareLocation struct {
	Index     int         `json:"index"`      // AP image index identifier for prepare location (Live: IOS-XE 17.12.6a)
	ImageFile string      `json:"image-file"` // AP image file name for prepare location (Live: IOS-XE 17.12.6a)
	ImageData []ImageData `json:"image-data"` // AP image info for prepare location (Live: IOS-XE 17.12.6a)
}

// ImageData represents image data information.
type ImageData struct {
	ImageName     string   `json:"image-name"`     // AP image name identifier (Live: IOS-XE 17.12.6a)
	ImageLocation string   `json:"image-location"` // AP image storage location path (Live: IOS-XE 17.12.6a)
	ImageVersion  string   `json:"image-version"`  // AP image version identifier (Live: IOS-XE 17.12.6a)
	IsNew         bool     `json:"is-new"`         // New image flag for install operation (Live: IOS-XE 17.12.6a)
	FileSize      string   `json:"file-size"`      // AP image file size (Live: IOS-XE 17.12.6a)
	ApModelList   []string `json:"ap-model-list"`  // List of supported AP models for this image (Live: IOS-XE 17.12.6a)
}

// ApImageActiveLocation represents AP image active location.
type ApImageActiveLocation struct {
	Index                          int    `json:"index"`      // AP image index identifier for active location (Live: IOS-XE 17.12.6a)
	ImageFile                      string `json:"image-file"` // AP image file name for active location (Live: IOS-XE 17.12.6a)
	ApImageActiveLocationImageData []struct {
		ImageName                                 string   `json:"image-name"`     // AP image name identifier (Live: IOS-XE 17.12.6a)
		ImageLocation                             string   `json:"image-location"` // AP image storage location path (Live: IOS-XE 17.12.6a)
		ImageVersion                              string   `json:"image-version"`  // AP image version identifier (Live: IOS-XE 17.12.6a)
		IsNew                                     bool     `json:"is-new"`         // New image flag for install operation (Live: IOS-XE 17.12.6a)
		FileSize                                  string   `json:"file-size"`      // AP image file size (Live: IOS-XE 17.12.6a)
		ApImageActiveLocationImageDataApModelList []string `json:"ap-model-list"`  // List of supported AP models for this image (Live: IOS-XE 17.12.6a)
	} `json:"image-data"` // AP image info for active location (Live: IOS-XE 17.12.6a)
}

// TCPMssConfig represents TCP MSS adjustment configuration.
type TCPMssConfig struct {
	TCPAdjustMssState bool `json:"tcp-adjust-mss-state"` // TCP MSS clamping state for CAPWAP (Live: IOS-XE 17.12.6a)
	TCPAdjustMssSize  int  `json:"tcp-adjust-mss-size"`  // TCP MSS clamp size in bytes (Live: IOS-XE 17.12.6a)
}

// PersistentSsid represents persistent SSID configuration.
type PersistentSsid struct {
	IsPersistentSsidEnabled bool `json:"is-persistent-ssid-enabled"` // SSID persistence across reboots/failover (Live: IOS-XE 17.12.6a)
}

// CdpIPAddress represents CDP IP address information.
type CdpIPAddress struct {
	IPAddressValue []string `json:"ip-address-value"` // CDP discovered neighbor IP addresses (Live: IOS-XE 17.12.6a)
}

// RadioOperStats represents radio operational statistics.
type RadioOperStats struct {
	ApMAC                 string        `json:"ap-mac"`                    // Access point MAC address (Live: IOS-XE 17.12.6a)
	SlotID                int           `json:"slot-id"`                   // Radio slot identifier (Live: IOS-XE 17.12.6a)
	AidUserList           *int          `json:"aid-user-list,omitempty"`   // Association ID user list for this radio (Live: IOS-XE 17.12.6a)
	TxFragmentCount       int           `json:"tx-fragment-count"`         // Number of transmitted frame fragments (Live: IOS-XE 17.12.6a)
	MultipleRetryCount    int           `json:"multiple-retry-count"`      // Multi-retry frame count (Live: IOS-XE 17.12.6a)
	MulticastTxFrameCnt   int           `json:"multicast-tx-frame-cnt"`    // Number of multicast frames transmitted (Live: IOS-XE 17.12.6a)
	FailedCount           int           `json:"failed-count"`              // Number of failed transmission attempts (Live: IOS-XE 17.12.6a)
	RetryCount            int           `json:"retry-count"`               // Number of frame retransmission attempts (Live: IOS-XE 17.12.6a)
	FrameDuplicateCount   int           `json:"frame-duplicate-count"`     // Number of duplicate frames received (Live: IOS-XE 17.12.6a)
	AckFailureCount       int           `json:"ack-failure-count"`         // Number of acknowledgment failures (Live: IOS-XE 17.12.6a)
	FcsErrorCount         int           `json:"fcs-error-count"`           // Number of frames with frame check sequence errors (Live: IOS-XE 17.12.6a)
	MACDecryErrFrameCount int           `json:"mac-decry-err-frame-count"` // Number of frames with MAC decryption errors (Live: IOS-XE 17.12.6a)
	MACMicErrFrameCount   int           `json:"mac-mic-err-frame-count"`   // MAC MIC error frame count (Live: IOS-XE 17.12.6a)
	MulticastRxFrameCnt   int           `json:"multicast-rx-frame-cnt"`    // Number of multicast frames received (Live: IOS-XE 17.12.6a)
	NoiseFloor            int           `json:"noise-floor"`               // Current noise floor level in dBm (Live: IOS-XE 17.12.6a)
	RtsFailureCount       int           `json:"rts-failure-count"`         // Number of Request to Send (RTS) failures (Live: IOS-XE 17.12.6a)
	RtsSuccessCount       int           `json:"rts-success-count"`         // Number of successful Request to Send (RTS) transmissions (Live: IOS-XE 17.12.6a)
	RxCtrlFrameCount      int           `json:"rx-ctrl-frame-count"`       // Number of control frames received (Live: IOS-XE 17.12.6a)
	RxDataFrameCount      int           `json:"rx-data-frame-count"`       // Number of data frames received (Live: IOS-XE 17.12.6a)
	RxDataPktCount        int           `json:"rx-data-pkt-count"`         // Number of data packets received (Live: IOS-XE 17.12.6a)
	RxErrorFrameCount     int           `json:"rx-error-frame-count"`      // Number of frames received with errors (Live: IOS-XE 17.12.6a)
	RxFragmentCount       int           `json:"rx-fragment-count"`         // Number of frame fragments received (Live: IOS-XE 17.12.6a)
	RxMgmtFrameCount      int           `json:"rx-mgmt-frame-count"`       // Number of management frames received (Live: IOS-XE 17.12.6a)
	TxCtrlFrameCount      int           `json:"tx-ctrl-frame-count"`       // Number of control frames transmitted (Live: IOS-XE 17.12.6a)
	TxDataFrameCount      int           `json:"tx-data-frame-count"`       // Number of data frames transmitted (Live: IOS-XE 17.12.6a)
	TxDataPktCount        int           `json:"tx-data-pkt-count"`         // Number of data packets transmitted (Live: IOS-XE 17.12.6a)
	TxFrameCount          int           `json:"tx-frame-count"`            // Total number of frames transmitted (Live: IOS-XE 17.12.6a)
	TxMgmtFrameCount      int           `json:"tx-mgmt-frame-count"`       // Number of management frames transmitted (Live: IOS-XE 17.12.6a)
	WepUndecryptableCount int           `json:"wep-undecryptable-count"`   // Number of WEP frames that could not be decrypted (Live: IOS-XE 17.12.6a)
	ApRadioStats          *ApRadioStats `json:"ap-radio-stats,omitempty"`  // Additional access point radio statistics (Live: IOS-XE 17.12.6a)
}

// EthernetIfStats represents Ethernet interface statistics.
type EthernetIfStats struct {
	WtpMAC           string `json:"wtp-mac"`            // Wireless termination point MAC address (Live: IOS-XE 17.12.6a)
	IfIndex          int    `json:"if-index"`           // Interface index identifier (Live: IOS-XE 17.12.6a)
	IfName           string `json:"if-name"`            // Interface name identifier (Live: IOS-XE 17.12.6a)
	RxPkts           int    `json:"rx-pkts"`            // Total packets received on interface (Live: IOS-XE 17.12.6a)
	TxPkts           int    `json:"tx-pkts"`            // Total packets transmitted on interface (Live: IOS-XE 17.12.6a)
	OperStatus       string `json:"oper-status"`        // Current operational status of interface (Live: IOS-XE 17.12.6a)
	RxUcastPkts      int    `json:"rx-ucast-pkts"`      // Unicast packets received (Live: IOS-XE 17.12.6a)
	RxNonUcastPkts   int    `json:"rx-non-ucast-pkts"`  // Non-unicast packets received (broadcast/multicast) (Live: IOS-XE 17.12.6a)
	TxUcastPkts      int    `json:"tx-ucast-pkts"`      // Unicast packets transmitted (Live: IOS-XE 17.12.6a)
	TxNonUcastPkts   int    `json:"tx-non-ucast-pkts"`  // Non-unicast packets transmitted (broadcast/multicast) (Live: IOS-XE 17.12.6a)
	Duplex           int    `json:"duplex"`             // Duplex mode of interface (full/half duplex) (Live: IOS-XE 17.12.6a)
	LinkSpeed        int    `json:"link-speed"`         // Current link speed in bits per second (Live: IOS-XE 17.12.6a)
	RxTotalBytes     int    `json:"rx-total-bytes"`     // Total bytes received on interface (Live: IOS-XE 17.12.6a)
	TxTotalBytes     int    `json:"tx-total-bytes"`     // Total bytes transmitted on interface (Live: IOS-XE 17.12.6a)
	InputCrc         int    `json:"input-crc"`          // Input cyclic redundancy check errors (Live: IOS-XE 17.12.6a)
	InputAborts      int    `json:"input-aborts"`       // Input packets aborted during reception (Live: IOS-XE 17.12.6a)
	InputErrors      int    `json:"input-errors"`       // Total input errors on interface (Live: IOS-XE 17.12.6a)
	InputFrames      int    `json:"input-frames"`       // Input framing errors (Live: IOS-XE 17.12.6a)
	InputOverrun     int    `json:"input-overrun"`      // Input overrun errors (Live: IOS-XE 17.12.6a)
	InputDrops       int    `json:"input-drops"`        // Input packets dropped by interface (Live: IOS-XE 17.12.6a)
	InputResource    int    `json:"input-resource"`     // Input packets dropped due to resource limitations (Live: IOS-XE 17.12.6a)
	UnknownProtocol  int    `json:"unknown-protocol"`   // Packets with unknown or unsupported protocol (Live: IOS-XE 17.12.6a)
	Runts            int    `json:"runts"`              // Packets smaller than minimum frame size (Live: IOS-XE 17.12.6a)
	Giants           int    `json:"giants"`             // Packets larger than maximum frame size (Live: IOS-XE 17.12.6a)
	Throttle         int    `json:"throttle"`           // Times interface was throttled (Live: IOS-XE 17.12.6a)
	Resets           int    `json:"resets"`             // Number of interface resets performed (Live: IOS-XE 17.12.6a)
	OutputCollision  int    `json:"output-collision"`   // Output collision detection events (Live: IOS-XE 17.12.6a)
	OutputNoBuffer   int    `json:"output-no-buffer"`   // Output packets dropped due to no buffer space (Live: IOS-XE 17.12.6a)
	OutputResource   int    `json:"output-resource"`    // Output packets dropped due to resource limits (Live: IOS-XE 17.12.6a)
	OutputUnderrun   int    `json:"output-underrun"`    // Output underrun errors (Live: IOS-XE 17.12.6a)
	OutputErrors     int    `json:"output-errors"`      // Total output errors on interface (Live: IOS-XE 17.12.6a)
	OutputTotalDrops int    `json:"output-total-drops"` // Total output packets dropped (Live: IOS-XE 17.12.6a)
}

// EwlcWncdStats represents EWLC WNCD statistics.
type EwlcWncdStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // Number of predownload sessions initiated (Live: IOS-XE 17.12.6a)
		NumInProgress           int  `json:"num-in-progress"`            // Predownload sessions in progress (Live: IOS-XE 17.12.6a)
		NumComplete             int  `json:"num-complete"`               // Predownload sessions completed (Live: IOS-XE 17.12.6a)
		NumUnsupported          int  `json:"num-unsupported"`            // Number of unsupported predownload requests (Live: IOS-XE 17.12.6a)
		NumFailed               int  `json:"num-failed"`                 // Number of predownload sessions that failed (Live: IOS-XE 17.12.6a)
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // Predownload operation active status (Live: IOS-XE 17.12.6a)
		NumTotal                int  `json:"num-total"`                  // Total number of predownload sessions attempted (Live: IOS-XE 17.12.6a)
	} `json:"predownload-stats"` // EWC predownload statistics (Live: IOS-XE 17.12.6a)
	DownloadsComplete   int                 `json:"downloads-complete"`              // Total number of completed downloads (Live: IOS-XE 17.12.6a)
	DownloadsInProgress int                 `json:"downloads-in-progress"`           // Number of downloads currently in progress (Live: IOS-XE 17.12.6a)
	WlcPredownloadStats *WlcPredownloadStat `json:"wlc-predownload-stats,omitempty"` // Wireless LAN Controller predownload statistics (Live: IOS-XE 17.12.6a)
}

// IotFirmware represents IoT firmware information for access points.
type IotFirmware struct {
	ApMAC      string    `json:"ap-mac"`      // Access point MAC address (Live: IOS-XE 17.12.6a)
	IfName     string    `json:"if-name"`     // Interface name for IoT radio (Live: IOS-XE 17.12.6a)
	IsDefault  EmptyType `json:"is-default"`  // Default firmware status (Live: IOS-XE 17.12.6a)
	Version    string    `json:"version"`     // Firmware version string (Live: IOS-XE 17.12.6a)
	VendorName string    `json:"vendor-name"` // Firmware vendor name (Live: IOS-XE 17.12.6a)
	Type       string    `json:"type"`        // Firmware type identifier (Live: IOS-XE 17.12.6a)
	Desc       string    `json:"desc"`        // Firmware description (Live: IOS-XE 17.12.6a)
}

// EmptyType represents YANG empty type fields appearing as null arrays in RESTCONF JSON.
type EmptyType []string

// AccountingInfo represents accounting information to be sent to RADIUS server.
type AccountingInfo struct {
	MethodList string `json:"method-list"` // Accounting method list (Live: IOS-XE 17.12.6a)
}

// ApDnaData represents Cisco-DNA related data.
type ApDnaData struct {
	GrpcStatus        string `json:"grpc-status"`         // gRPC status for DNA connection (Live: IOS-XE 17.12.6a)
	PacketsTxAttempts string `json:"packets-tx-attempts"` // Number of transmission attempts (Live: IOS-XE 17.12.6a)
	PacketsTxFailures string `json:"packets-tx-failures"` // Number of transmission failures (Live: IOS-XE 17.12.6a)
	PacketsRx         string `json:"packets-rx"`          // Number of packets received (Live: IOS-XE 17.12.6a)
	PacketsRxFailures string `json:"packets-rx-failures"` // Number of receive failures (Live: IOS-XE 17.12.6a)
}

// ApGasRateLimitConfig represents Generic Advertisement Service (GAS) rate limiting configuration.
type ApGasRateLimitConfig struct {
	IsGasRateLimitEnabled bool `json:"is-gas-rate-limit-enabled"` // GAS rate limiting status (Live: IOS-XE 17.12.6a)
	NumReqPerInterval     int  `json:"num-req-per-interval"`      // Number of requests per interval (Live: IOS-XE 17.12.6a)
	IntervalMsec          int  `json:"interval-msec"`             // Rate limiting interval in milliseconds (Live: IOS-XE 17.12.6a)
}

// ApIPData represents AP IP address configuration.
type ApIPData struct {
	ApPrefix         int    `json:"ap-prefix"`           // AP IP prefix length (Live: IOS-XE 17.12.6a)
	Mtu              int    `json:"mtu"`                 // Maximum transmission unit (Live: IOS-XE 17.12.6a)
	IsStaticApIPAddr bool   `json:"is-static-ap-ipaddr"` // Static IP address configuration status (Live: IOS-XE 17.12.6a)
	DomainName       string `json:"domain-name"`         // Domain name configuration (Live: IOS-XE 17.12.6a)
	ApIPAddr         string `json:"ap-ip-addr"`          // AP IP address (Live: IOS-XE 17.12.6a)
	ApIPv6Addr       string `json:"ap-ipv6-addr"`        // AP IPv6 address (Live: IOS-XE 17.12.6a)
	ApIPNetmask      string `json:"ap-ip-netmask"`       // AP IP netmask (Live: IOS-XE 17.12.6a)
	ApIPGateway      string `json:"ap-ip-gateway"`       // AP IP gateway (Live: IOS-XE 17.12.6a)
	ApIPv6Gateway    string `json:"ap-ipv6-gateway"`     // AP IPv6 gateway (Live: IOS-XE 17.12.6a)
	ApNameServerType string `json:"ap-name-server-type"` // Name server type (Live: IOS-XE 17.12.6a)
	ApIPv6Method     string `json:"ap-ipv6-method"`      // IPv6 configuration method (Live: IOS-XE 17.12.6a)
	StaticIP         string `json:"static-ip"`           // Static IP address (Live: IOS-XE 17.12.6a)
	StaticGwIP       string `json:"static-gw-ip"`        // Static gateway IP address (Live: IOS-XE 17.12.6a)
	StaticNetmask    string `json:"static-netmask"`      // Static netmask (Live: IOS-XE 17.12.6a)
	StaticPrefix     int    `json:"static-prefix"`       // Static prefix length (Live: IOS-XE 17.12.6a)
}

// ApLoginCredentials represents login credentials configured on an AP.
type ApLoginCredentials struct {
	Dot1xEapType  string `json:"dot1x-eap-type"` // 802.1X EAP type (Live: IOS-XE 17.12.6a)
	UserName      string `json:"user-name"`      // Username for AP login (Live: IOS-XE 17.12.6a)
	Dot1xUsername string `json:"dot1x-username"` // 802.1X username (Live: IOS-XE 17.12.6a)
}

// ApManagement represents AP management configuration.
type ApManagement struct {
	IsTelnetEnabled  bool `json:"is-telnet-enabled"`  // Telnet access status (Live: IOS-XE 17.12.6a)
	IsSSHEnabled     bool `json:"is-ssh-enabled"`     // SSH access status (Live: IOS-XE 17.12.6a)
	IsConsoleEnabled bool `json:"is-console-enabled"` // Console access status (Live: IOS-XE 17.12.6a)
}

// ApNtpServerInfo represents NTP server information for AP.
type ApNtpServerInfo struct {
	NtpServerAddress string `json:"ntp-server-address"` // NTP server IP address (Live: IOS-XE 17.12.6a)
	TrustKey         string `json:"trust-key"`          // NTP trust key (Live: IOS-XE 17.12.6a)
	KeyID            int    `json:"key-id"`             // NTP key identifier (Live: IOS-XE 17.12.6a)
	KeyType          string `json:"key-type"`           // NTP key type (Live: IOS-XE 17.12.6a)
	KeyFormat        string `json:"key-format"`         // NTP key format (Live: IOS-XE 17.12.6a)
	TrustKeyType     string `json:"trust-key-type"`     // NTP trust key type (Live: IOS-XE 17.12.6a)
}

// ApNtpSyncStatus represents AP NTP synchronization status.
type ApNtpSyncStatus struct {
	Enabled          bool   `json:"enabled"`             // NTP synchronization enabled status (Live: IOS-XE 17.12.6a)
	Stratum          int    `json:"stratum"`             // NTP stratum level (Live: IOS-XE 17.12.6a)
	Status           string `json:"status"`              // NTP synchronization status (Live: IOS-XE 17.12.6a)
	SecSinceLastSync int    `json:"sec-since-last-sync"` // Seconds since last synchronization (Live: IOS-XE 17.12.6a)
	SyncOffset       int    `json:"sync-offset"`         // Synchronization offset (Live: IOS-XE 17.12.6a)
	RxTS             string `json:"rx-ts"`               // Last receive timestamp (Live: IOS-XE 17.12.6a)
}

// ApPowData represents AP power related data.
type ApPowData struct {
	PowerInjectorSel     string `json:"power-injector-sel"`     // Power injector selection (Live: IOS-XE 17.12.6a)
	PowerInjectorMacaddr string `json:"power-injector-macaddr"` // Power injector MAC address (Live: IOS-XE 17.12.6a)
	PreStdSwitchEnabled  bool   `json:"pre-std-switch-enabled"` // Pre-standard switch enabled status (Live: IOS-XE 17.12.6a)
	PowerInjectorEnabled bool   `json:"power-injector-enabled"` // Power injector enabled status (Live: IOS-XE 17.12.6a)
	PowerType            string `json:"power-type"`             // Power source type (Live: IOS-XE 17.12.6a)
	PowerMode            string `json:"power-mode"`             // Power mode configuration (Live: IOS-XE 17.12.6a)
}

// ApPrimeInfo represents controller configuration for the AP.
type ApPrimeInfo struct {
	PrimaryControllerName     string `json:"primary-controller-name"`      // Primary controller name (Live: IOS-XE 17.12.6a)
	SecondaryControllerName   string `json:"secondary-controller-name"`    // Secondary controller name (Live: IOS-XE 17.12.6a)
	PrimaryControllerIPAddr   string `json:"primary-controller-ip-addr"`   // Primary controller IP address (Live: IOS-XE 17.12.6a)
	SecondaryControllerIPAddr string `json:"secondary-controller-ip-addr"` // Secondary controller IP address (Live: IOS-XE 17.12.6a)
	TertiaryControllerName    string `json:"tertiary-controller-name"`     // Tertiary controller name (Live: IOS-XE 17.12.6a)
	TertiaryControllerIPAddr  string `json:"tertiary-controller-ip-addr"`  // Tertiary controller IP address (Live: IOS-XE 17.12.6a)
	ApFallbackIP              string `json:"ap-fallback-ip"`               // AP fallback IP address (Live: IOS-XE 17.12.6a)
	FallbackEnabled           bool   `json:"fallback-enabled"`             // Fallback enabled status (Live: IOS-XE 17.12.6a)
}

// ApSystemStats represents AP system statistics.
type ApSystemStats struct {
	CPUUsage       int              `json:"cpu-usage"`        // Current CPU usage percentage (Live: IOS-XE 17.12.6a)
	MemoryUsage    int              `json:"memory-usage"`     // Current memory usage percentage (Live: IOS-XE 17.12.6a)
	AvgCPUUsage    int              `json:"avg-cpu-usage"`    // Average CPU usage percentage (Live: IOS-XE 17.12.6a)
	AvgMemoryUsage int              `json:"avg-memory-usage"` // Average memory usage percentage (Live: IOS-XE 17.12.6a)
	WindowSize     int              `json:"window-size"`      // Statistics window size (Live: IOS-XE 17.12.6a)
	LastTS         string           `json:"last-ts"`          // Last timestamp (Live: IOS-XE 17.12.6a)
	Memory         ApAlarmThreshold `json:"memory"`           // Memory alarm configuration (Live: IOS-XE 17.12.6a)
	CPU            ApAlarmThreshold `json:"cpu"`              // CPU alarm configuration (Live: IOS-XE 17.12.6a)
}

// ApAlarmThreshold represents alarm threshold configuration.
type ApAlarmThreshold struct {
	AlarmStatus   string `json:"alarm-status"`    // Alarm status (Live: IOS-XE 17.12.6a)
	RaiseTicks    string `json:"raise-ticks"`     // Alarm raise timestamp (Live: IOS-XE 17.12.6a)
	ClearTicks    string `json:"clear-ticks"`     // Alarm clear timestamp (Live: IOS-XE 17.12.6a)
	LastSendTicks string `json:"last-send-ticks"` // Last send timestamp (Live: IOS-XE 17.12.6a)
	Type          string `json:"type"`            // Alarm type (Live: IOS-XE 17.12.6a)
}

// ApTimezoneConfig represents AP timezone configuration.
type ApTimezoneConfig struct {
	TzEnabled  bool   `json:"tz-enabled"`  // Timezone configuration enabled status (Live: IOS-XE 17.12.6a)
	OffsetHour int    `json:"offset-hour"` // Timezone offset hours (Live: IOS-XE 17.12.6a)
	OffsetMin  int    `json:"offset-min"`  // Timezone offset minutes (Live: IOS-XE 17.12.6a)
	Mode       string `json:"mode"`        // Timezone mode (Live: IOS-XE 17.12.6a)
}

// AuxClientInterfaceData represents auxiliary client interface data.
type AuxClientInterfaceData struct {
	IsEnabled bool `json:"is-enabled"` // Auxiliary client interface enabled status (Live: IOS-XE 17.12.6a)
}

// InfrastructureMfp represents AP Management Frame Protection configuration.
type InfrastructureMfp struct {
	MfpValidation bool `json:"mfp-validation"` // MFP validation enabled status (Live: IOS-XE 17.12.6a)
	MfpProtection bool `json:"mfp-protection"` // MFP protection enabled status (Live: IOS-XE 17.12.6a)
}

// KernelCoredumpConfig represents kernel coredump configuration.
type KernelCoredumpConfig struct {
	KernelCoredumpLimit int `json:"kernel-coredump-limit"` // Kernel coredump limit (Live: IOS-XE 17.12.6a)
}

// LinkAuditData represents link audit configuration and data.
type LinkAuditData struct {
	LinkauditFlag      string `json:"linkaudit-flag"`       // Link audit flag status (Live: IOS-XE 17.12.6a)
	LinkauditDelayTime int    `json:"linkaudit-delay-time"` // Link audit delay time (Live: IOS-XE 17.12.6a)
	LinkauditMaxTime   int    `json:"linkaudit-max-time"`   // Link audit maximum time (Live: IOS-XE 17.12.6a)
	LinkauditMinTime   int    `json:"linkaudit-min-time"`   // Link audit minimum time (Live: IOS-XE 17.12.6a)
	LinkauditRcvTime   int    `json:"linkaudit-rcv-time"`   // Link audit receive time (Live: IOS-XE 17.12.6a)
}

// OeapAuditData represents Office Extended AP audit data.
type OeapAuditData struct {
	LastRun    string `json:"last-run"`    // Last audit run timestamp (Live: IOS-XE 17.12.6a)
	State      string `json:"state"`       // Audit state (Live: IOS-XE 17.12.6a)
	DTLSUpload string `json:"dtls-upload"` // DTLS upload information (Live: IOS-XE 17.12.6a)
	Latency    int    `json:"latency"`     // Network latency (Live: IOS-XE 17.12.6a)
	Jitter     int    `json:"jitter"`      // Network jitter (Live: IOS-XE 17.12.6a)
}

// RetransmitConfig represents AP retransmission parameters.
type RetransmitConfig struct {
	Count    int `json:"count"`    // Retransmission count (Live: IOS-XE 17.12.6a)
	Interval int `json:"interval"` // Retransmission interval in seconds (Live: IOS-XE 17.12.6a)
}

// SyslogConfig represents AP system logging configuration.
type SyslogConfig struct {
	LogHostIPAddr    string `json:"log-host-ipaddr"`    // Syslog host IP address (Live: IOS-XE 17.12.6a)
	LogTrapLevel     string `json:"log-trap-level"`     // Syslog trap level (Live: IOS-XE 17.12.6a)
	LogTLSMode       bool   `json:"log-tls-mode"`       // TLS mode enabled status (Live: IOS-XE 17.12.6a)
	LogFacilityLevel string `json:"log-facility-level"` // Syslog facility level (Live: IOS-XE 17.12.6a)
}

// ApTimerConfig represents AP timer configuration.
type ApTimerConfig struct {
	StatsTimer ApStatsTimer `json:"stats-timer"` // Statistics timer configuration (Live: IOS-XE 17.12.6a)
}

// ApStatsTimer represents AP statistics timer.
type ApStatsTimer struct {
	StatTmr int `json:"stat-tmr"` // Statistics timer value (Live: IOS-XE 17.12.6a)
}

// ChannelList represents a list of wireless channels.
type ChannelList struct {
	Channel []int `json:"channel"` // List of supported channels (Live: IOS-XE 17.12.6a)
}

// RegDomainCodes represents regulatory domain codes.
type RegDomainCodes struct {
	RegDomainCode []string `json:"reg-domain-code"` // List of regulatory domain codes (Live: IOS-XE 17.12.6a)
}

// ApRadioStats represents additional AP radio statistics.
type ApRadioStats struct {
	StuckTS            string `json:"stuck-ts"`              // Last stuck timestamp (Live: IOS-XE 17.12.6a)
	LastTS             string `json:"last-ts"`               // Last update timestamp (Live: IOS-XE 17.12.6a)
	NumRadioStuckReset int    `json:"num-radio-stuck-reset"` // Number of radio stuck resets (Live: IOS-XE 17.12.6a)
}

// LscStatusPayload represents LSC status payload support types.
type LscStatusPayload string

// AlarmEnableType represents alarm enablement settings for radio statistics.
type AlarmEnableType bool

// WlcPredownloadStat represents Wireless LAN Controller predownload statistics.
type WlcPredownloadStat struct {
	CompletedCount    *int    `json:"completed-count,omitempty"`     // Number of completed predownload operations
	InProgressCount   *int    `json:"in-progress-count,omitempty"`   // Number of predownload operations in progress
	ErrorCount        *int    `json:"error-count,omitempty"`         // Number of failed predownload operations
	LastOperationTime *string `json:"last-operation-time,omitempty"` // Timestamp of last predownload operation
}
