package ap

import "time"

// ApOper represents access point operational data response.
type ApOper struct {
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
		OperData                []ApOperInternalData     `json:"oper-data"`
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
		IotFirmware             interface{}              `json:"iot-firmware,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"`
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

// ApPwrInfo represents AP power information.
type ApPwrInfo struct {
	WtpMac  string    `json:"wtp-mac"`  // MAC Address of the AP Radio (Live: IOS-XE 17.12.5)
	Status  string    `json:"status"`   // Power status information (Live: IOS-XE 17.12.5)
	PpeInfo []PpeInfo `json:"ppe-info"` // Power policy entry information (Live: IOS-XE 17.12.5)
}

// PpeInfo represents power policy entry information.
type PpeInfo struct {
	SeqNumber int    `json:"seq-number"` // Sequence number for power policy entry (Live: IOS-XE 17.12.5)
	PpeResult string `json:"ppe-result"` // Power policy entry result (Live: IOS-XE 17.12.5)
	Ethernet  *struct {
		EthID    string `json:"eth-id"`    // Ethernet interface ID (Live: IOS-XE 17.12.5)
		EthSpeed string `json:"eth-speed"` // Ethernet speed (Live: IOS-XE 17.12.5)
	} `json:"ethernet,omitempty"` // Ethernet interface information (Live: IOS-XE 17.12.5)
	Radio *struct {
		RadioID       string  `json:"radio-id"`                 // Radio interface ID (Live: IOS-XE 17.12.5)
		SpatialStream *string `json:"spatial-stream,omitempty"` // Spatial stream configuration (Live: IOS-XE 17.12.5)
		State         *string `json:"state,omitempty"`          // Radio state (Live: IOS-XE 17.12.5)
	} `json:"radio,omitempty"` // Radio interface information (Live: IOS-XE 17.12.5)
	Usb *struct {
		UsbID string  `json:"usb-id"`          // USB interface ID (Live: IOS-XE 17.12.5)
		State *string `json:"state,omitempty"` // USB state (Live: IOS-XE 17.12.5)
	} `json:"usb,omitempty"` // USB interface information (Live: IOS-XE 17.12.5)
}

// ApSensorStatus represents AP sensor status information.
type ApSensorStatus struct {
	ApMac       string `json:"ap-mac"`       // Access point MAC address (Live: IOS-XE 17.12.5)
	SensorType  string `json:"sensor-type"`  // Sensor type identifier (Live: IOS-XE 17.12.5)
	ConfigState string `json:"config-state"` // Sensor configuration state (Live: IOS-XE 17.12.5)
	AdminState  string `json:"admin-state"`  // Administrative state (Live: IOS-XE 17.12.5)
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
	WtpMac       string `json:"wtp-mac"`                  // Wireless Termination Point MAC address
	RadioSlotID  int    `json:"radio-slot-id"`            // Radio slot identifier
	SlotID       int    `json:"slot-id,omitempty"`        // Radio slot identifier (Live: IOS-XE 17.12.5)
	RadioType    string `json:"radio-type,omitempty"`     // Radio type information (Live: IOS-XE 17.12.5)
	AdminState   string `json:"admin-state,omitempty"`    // Administrative state (Live: IOS-XE 17.12.5)
	OperState    string `json:"oper-state,omitempty"`     // Operational state (Live: IOS-XE 17.12.5)
	RadioMode    string `json:"radio-mode,omitempty"`     // Radio mode configuration (Live: IOS-XE 17.12.5)
	RadioSubMode string `json:"radio-sub-mode,omitempty"` // Radio sub-mode configuration (Live: IOS-XE 17.12.5)
	RadioSubtype string `json:"radio-subtype,omitempty"`  // Radio subtype information (Live: IOS-XE 17.12.5)
	RadioSubband string `json:"radio-subband,omitempty"`  // Radio subband information (Live: IOS-XE 17.12.5)

	// Band and channel information
	CurrentBandID     int    `json:"current-band-id,omitempty"`     // Current radio band identifier
	CurrentActiveBand string `json:"current-active-band,omitempty"` // Currently active band designation

	// Protocol capabilities
	PhyHtCap        *PhyHtCapStruct `json:"phy-ht-cap,omitempty"`        // Physical layer HT capabilities
	PhyHeCap        *PhyHeCapStruct `json:"phy-he-cap,omitempty"`        // Physical layer HE capabilities
	RadioHeCapable  bool            `json:"radio-he-capable,omitempty"`  // Radio HE capability status
	RadioFraCapable string          `json:"radio-fra-capable,omitempty"` // Radio frame aggregation capability

	// XOR capabilities
	XorRadioMode string       `json:"xor-radio-mode,omitempty"` // XOR radio mode configuration
	XorPhyHtCap  *XorPhyHtCap `json:"xor-phy-ht-cap,omitempty"` // XOR physical HT capabilities
	XorPhyHeCap  *XorPhyHeCap `json:"xor-phy-he-cap,omitempty"` // XOR physical HE capabilities

	// Additional operational fields
	AntennaGain            int               `json:"antenna-gain,omitempty"`             // Antenna gain value in dBi
	AntennaPid             string            `json:"antenna-pid,omitempty"`              // Antenna product identifier
	SlotAntennaType        string            `json:"slot-antenna-type,omitempty"`        // Slot antenna type designation
	RadioEnableTime        string            `json:"radio-enable-time,omitempty"`        // Radio enable timestamp
	HighestThroughputProto string            `json:"highest-throughput-proto,omitempty"` // Highest throughput protocol supported
	CacActive              bool              `json:"cac-active,omitempty"`               // Channel availability check active status
	MeshBackhaul           bool              `json:"mesh-backhaul,omitempty"`            // Mesh backhaul configuration status
	MeshDesignatedDownlink bool              `json:"mesh-designated-downlink,omitempty"` // Mesh designated downlink status
	MultiDomainCap         *MultiDomainCap   `json:"multi-domain-cap,omitempty"`         // Multi-domain capabilities
	StationCfg             *StationCfg       `json:"station-cfg,omitempty"`              // Station configuration parameters
	PhyHtCfg               *PhyHtCfg         `json:"phy-ht-cfg,omitempty"`               // Physical HT configuration
	ChanPwrInfo            *ChanPwrInfo      `json:"chan-pwr-info,omitempty"`            // Channel power information
	SnifferCfg             *SnifferCfg       `json:"sniffer-cfg,omitempty"`              // Sniffer configuration parameters
	RadioBandInfo          []RadioBandInfo   `json:"radio-band-info,omitempty"`          // Radio band information array
	VapOperConfig          []VapOperConfig   `json:"vap-oper-config,omitempty"`          // VAP operational configuration
	RegDomainCheckStatus   string            `json:"reg-domain-check-status,omitempty"`  // Regulatory domain check status
	Dot11nMcsRates         string            `json:"dot11n-mcs-rates,omitempty"`         // 802.11n MCS rates supported
	DualRadioModeCfg       *DualRadioModeCfg `json:"dual-radio-mode-cfg,omitempty"`      // Dual radio mode configuration
	BssColorCfg            *BssColorCfg      `json:"bss-color-cfg,omitempty"`            // BSS color configuration
	ObssPdCapable          bool              `json:"obss-pd-capable,omitempty"`          // OBSS preamble detection capable
	NdpCap                 string            `json:"ndp-cap,omitempty"`                  // NDP capability information
	NdpOnChannel           bool              `json:"ndp-on-channel,omitempty"`           // NDP on channel status
	BeamSelection          string            `json:"beam-selection,omitempty"`           // Beam selection algorithm
	NumAntEnabled          uint8             `json:"num-ant-enabled,omitempty"`          // Number of antennas enabled
	CurAntBitmap           string            `json:"cur-ant-bitmap,omitempty"`           // Current antenna bitmap
	SuppAntBitmap          string            `json:"supp-ant-bitmap,omitempty"`          // Supported antenna bitmap
	Supp160mhzAntBitmap    string            `json:"supp-160mhz-ant-bitmap,omitempty"`   // Supported 160MHz antenna bitmap
	MaxClientAllowed       uint16            `json:"max-client-allowed,omitempty"`       // Maximum clients allowed
	ObssPdSrgCapable       bool              `json:"obss-pd-srg-capable,omitempty"`      // OBSS PD SRG capability
	CoverageOverlapFactor  uint8             `json:"coverage-overlap-factor,omitempty"`  // Coverage overlap factor

	// 6GHz related (YANG: IOS-XE 17.12.1+)
	Ap6GhzPwrMode    *string `json:"ap-6ghz-pwr-mode,omitempty"`     // 6GHz power mode configuration (YANG: IOS-XE 17.12.1+)
	Ap6GhzPwrModeCap *string `json:"ap-6ghz-pwr-mode-cap,omitempty"` // 6GHz power mode capability (YANG: IOS-XE 17.12.1+)

	// AFC related
	AfcBelowTxmin    bool `json:"afc-below-txmin,omitempty"`    // AFC below minimum transmission power (YANG: IOS-XE 17.12.1+)
	AfcLicenseNeeded bool `json:"afc-license-needed,omitempty"` // AFC license requirement status (YANG: IOS-XE 17.12.1+)
	PushAfcRespDone  bool `json:"push-afc-resp-done,omitempty"` // AFC response push completion status (YANG: IOS-XE 17.12.1+)
}

// RadioResetStats represents radio reset statistics.
type RadioResetStats struct {
	ApMac       string `json:"ap-mac"`       // Access point MAC address
	RadioID     int    `json:"radio-id"`     // Radio interface identifier
	Cause       string `json:"cause"`        // Reset cause description
	DetailCause string `json:"detail-cause"` // Detailed reset cause information
	Count       int    `json:"count"`        // Number of reset occurrences
}

// QosClientData represents QoS client data.
type QosClientData struct {
	ClientMac    string `json:"client-mac"` // Client MAC address
	AaaQosParams struct {
		AaaAvgdtus   int `json:"aaa-avgdtus"`   // AAA average downstream utilization (seconds)
		AaaAvgrtdtus int `json:"aaa-avgrtdtus"` // AAA average real-time downstream utilization (seconds)
		AaaBstdtus   int `json:"aaa-bstdtus"`   // AAA burst downstream utilization (seconds)
		AaaBstrtdtus int `json:"aaa-bstrtdtus"` // AAA burst real-time downstream utilization (seconds)
		AaaAvgdtds   int `json:"aaa-avgdtds"`   // AAA average downstream data size
		AaaAvgrtdtds int `json:"aaa-avgrtdtds"` // AAA average real-time downstream data size
		AaaBstdtds   int `json:"aaa-bstdtds"`   // AAA burst downstream data size
		AaaBstrtdtds int `json:"aaa-bstrtdtds"` // AAA burst real-time downstream data size
	} `json:"aaa-qos-params"` // AAA QoS parameters
}

// CapwapData represents CAPWAP data.
type CapwapData struct {
	WtpMac       string       `json:"wtp-mac"`       // Wireless termination point MAC address
	IPAddr       string       `json:"ip-addr"`       // AP management IP address
	Name         string       `json:"name"`          // Access point hostname
	DeviceDetail DeviceDetail `json:"device-detail"` // Hardware device details
	ApState      ApState      `json:"ap-state"`      // Access point state information

	// AP Mode Data
	ApModeData ApModeData `json:"ap-mode-data"` // Access point mode configuration

	// Location and Services
	ApLocation         ApLocation         `json:"ap-location"`          // Physical location information
	ApServices         ApServices         `json:"ap-services"`          // Enabled AP services
	TagInfo            TagInfo            `json:"tag-info"`             // AP tag assignment information
	Tunnel             Tunnel             `json:"tunnel"`               // Tunnel configuration data
	ExternalModuleData ExternalModuleData `json:"external-module-data"` // External module information
	ApTimeInfo         ApTimeInfo         `json:"ap-time-info"`         // Time synchronization information
	ApSecurityData     ApSecurityData     `json:"ap-security-data"`     // Security configuration data
	SlidingWindow      SlidingWindow      `json:"sliding-window"`       // Sliding window configuration
	ApVlan             ApVlan             `json:"ap-vlan"`              // VLAN configuration
	HyperlocationData  HyperlocationData  `json:"hyperlocation-data"`   // Hyperlocation service data
	RebootStats        RebootStats        `json:"reboot-stats"`         // AP reboot statistics
	ProxyInfo          ProxyInfo          `json:"proxy-info"`           // HTTP proxy configuration

	// Image Download Tracking
	ImageSizeEta           uint64 `json:"image-size-eta"`            // ETA for image download (timeticks)
	ImageSizeStartTime     string `json:"image-size-start-time"`     // Image download start time
	ImageSizePercentage    uint32 `json:"image-size-percentage"`     // Percentage of image download completed
	WlcImageSizeEta        uint64 `json:"wlc-image-size-eta"`        // ETA for controller image download (timeticks)
	WlcImageSizeStartTime  string `json:"wlc-image-size-start-time"` // Controller image download start time
	WlcImageSizePercentage uint32 `json:"wlc-image-size-percentage"` // Percentage of controller image download completed

	// Local DHCP Configuration
	Ipv4Pool              Ipv4Pool          `json:"ipv4-pool"`                // DHCP IPv4 pool configuration
	DisconnectDetail      DisconnectDetail  `json:"disconnect-detail"`        // AP disconnect detail
	StatsMonitor          StatsMonitor      `json:"stats-monitor"`            // AP statistics monitoring configuration
	LscStatusPldSupported []interface{}     `json:"lsc-status-pld-supported"` // AP platform support for LSC-status payload (empty type)
	ApLscStatus           ApLscStatus       `json:"ap-lsc-status"`            // AP LSC status information
	RadioStatsMonitor     RadioStatsMonitor `json:"radio-stats-monitor"`      // AP radio statistics monitoring configuration
	ZeroWtDfs             ZeroWtDfs         `json:"zero-wt-dfs"`              // Zero wait DFS information
	GnssInfo              GnssInfo          `json:"gnss-info"`                // AP GNSS information

	// Basic Configuration Fields
	ApLagEnabled    bool   `json:"ap-lag-enabled"`    // AP LAG configuration status (YANG: IOS-XE 17.12.1+)
	CountryCode     string `json:"country-code"`      // Country code configuration
	NumRadioSlots   uint8  `json:"num-radio-slots"`   // Number of radio slots available
	Ipv6Joined      uint8  `json:"ipv6-joined"`       // IPv6 join status (YANG: IOS-XE 17.12.1+)
	DartIsConnected bool   `json:"dart-is-connected"` // DART connection status (YANG: IOS-XE 17.12.1+)
	IsMaster        bool   `json:"is-master"`         // Master AP designation status (YANG: IOS-XE 17.12.1+)
	CdpEnable       bool   `json:"cdp-enable"`        // CDP protocol enable status
	GrpcEnabled     bool   `json:"grpc-enabled"`      // gRPC service enable status (YANG: IOS-XE 17.12.1+)
	LocalDhcp       bool   `json:"local-dhcp"`        // Local DHCP service status (YANG: IOS-XE 17.12.1+)

	// Status and operational fields
	ApStationType        string `json:"ap-stationing-type,omitempty"`      // AP stationing type configuration
	ApKeepAliveState     bool   `json:"ap-keepalive-state,omitempty"`      // AP keepalive state status
	MaxClientsSupported  uint16 `json:"max-clients-supported,omitempty"`   // Maximum supported client count
	MdnsGroupID          uint32 `json:"mdns-group-id,omitempty"`           // mDNS group identifier
	MdnsRuleName         string `json:"mdns-rule-name,omitempty"`          // mDNS rule name configuration
	MdnsGroupMethod      string `json:"mdns-group-method,omitempty"`       // mDNS group method specification
	MerakiCapable        bool   `json:"meraki-capable,omitempty"`          // Meraki integration capability (YANG: IOS-XE 17.12.1+)
	MerakiConnectStatus  string `json:"meraki-connect-status,omitempty"`   // Meraki connection status
	MerakiMonitorCapable bool   `json:"meraki-monitor-capable,omitempty"`  // Meraki monitoring capability (YANG: IOS-XE 17.18.1+)
	KernelCoredumpCount  uint16 `json:"kernel-coredump-count,omitempty"`   // Kernel core dump count
	RegDomain            string `json:"reg-domain,omitempty"`              // Regulatory domain configuration
	DartConStatus        string `json:"dart-con-status,omitempty"`         // DART connection status
	ApAfcPreNotification bool   `json:"ap-afc-pre-notification,omitempty"` // AP AFC pre-notification status (YANG: IOS-XE 17.12.1+)
	OobImgDwldMethod     string `json:"oob-img-dwld-method,omitempty"`     // Out-of-band image download method
	WtpIP                string `json:"wtp-ip,omitempty"`                  // IP address of the AP (YANG: IOS-XE 17.12.1+)
}

// ApTimeInfo represents AP time related information.
type ApTimeInfo struct {
	BootTime      string `json:"boot-time"`       // Last AP reboot Date and Time
	JoinTime      string `json:"join-time"`       // Date and Time at which AP joined
	JoinTimeTaken uint32 `json:"join-time-taken"` // Time taken by AP to join in seconds
}

// ApSecurityData represents AP LSC (Local Significant Certificate) data.
type ApSecurityData struct {
	FipsEnabled      bool   `json:"fips-enabled"`        // Cisco AP FIPS enabled status
	WlanccEnabled    bool   `json:"wlancc-enabled"`      // Cisco AP CC enabled status
	CertType         string `json:"cert-type"`           // AP Certificate Type
	LscApAuthType    string `json:"lsc-ap-auth-type"`    // AP LSC authentication state
	ApCertPolicy     string `json:"ap-cert-policy"`      // Certificate policy used during AP join
	ApCertExpiryTime string `json:"ap-cert-expiry-time"` // AP certificate expiry time
	ApCertIssuerCn   string `json:"ap-cert-issuer-cn"`   // AP certificate issuer common name
}

// SlidingWindow represents CAPWAP multiwindow transport information.
type SlidingWindow struct {
	MultiWindowSupport bool   `json:"multi-window-support"` // True if CAPWAP multiwindow is enabled on AP
	WindowSize         uint16 `json:"window-size"`          // Window size for CAPWAP multiwindow transport
}

// ApVlan represents AP VLAN tagging details.
type ApVlan struct {
	VlanTagState string `json:"vlan-tag-state"` // AP VLAN tagging state
	VlanTagID    uint16 `json:"vlan-tag-id"`    // VLAN ID for the AP
}

// HyperlocationData represents AP Hyperlocation details.
type HyperlocationData struct {
	HyperlocationMethod string `json:"hyperlocation-method"` // AP hyperlocation method
	CmxIP               string `json:"cmx-ip,omitempty"`     // Connected Mobile Experiences (CMX) IP address
}

// RebootStats represents AP reboot statistics.
type RebootStats struct {
	RebootReason string `json:"reboot-reason"` // Reason for last AP reboot
	RebootType   string `json:"reboot-type"`   // AP specified last reboot type
}

// ProxyInfo represents HTTP proxy configuration provisioned to AP.
type ProxyInfo struct {
	Hostname     string  `json:"hostname"`                // HTTP proxy hostname
	Port         uint16  `json:"port"`                    // HTTP proxy port
	NoProxyList  string  `json:"no-proxy-list"`           // List of URLs to be excluded from proxying
	Username     *string `json:"username,omitempty"`      // AP proxy username (YANG: IOS-XE 17.12.1+)
	PasswordType *string `json:"password-type,omitempty"` // Password type for AP proxy (YANG: IOS-XE 17.12.1+)
	Password     *string `json:"password,omitempty"`      // Password for AP proxy (YANG: IOS-XE 17.12.1+)
}

// Ipv4Pool represents DHCP IPv4 pool configuration.
type Ipv4Pool struct {
	Network   string `json:"network"`    // DHCP pool start address
	LeaseTime uint16 `json:"lease-time"` // Total lease time in days
	Netmask   string `json:"netmask"`    // Subnet mask address
}

// DisconnectDetail represents AP disconnect detail.
type DisconnectDetail struct {
	DisconnectReason string `json:"disconnect-reason"` // AP specified last disconnect reason
}

// StatsMonitor represents AP statistics monitoring configuration.
type StatsMonitor struct {
	ActionApReload bool `json:"action-ap-reload"` // AP reload action on high CPU or high memory usage (YANG: IOS-XE 17.12.1+)
}

// ApLscStatus represents AP LSC (Local Significant Certificate) status information.
type ApLscStatus struct {
	IsDtlsLscEnabled      bool   `json:"is-dtls-lsc-enabled"`                 // LSC enable status for CAPWAP DTLS handshake
	IsDot1xLscEnabled     bool   `json:"is-dot1x-lsc-enabled"`                // LSC enable status for dot1x port authentication
	IsDtlsLscFallback     bool   `json:"is-dtls-lsc-fallback"`                // AP fallback state to default certificate instead of LSC
	DtlsLscIssuerHash     string `json:"dtls-lsc-issuer-hash,omitempty"`      // Issuer certificate hash for CAPWAP DTLS
	Dot1xLscIssuerHash    string `json:"dot1x-lsc-issuer-hash,omitempty"`     // Issuer certificate hash for dot1x authentication
	DtlsLscCertExpiryTime string `json:"dtls-lsc-cert-expiry-time,omitempty"` // Certificate expiry time for CAPWAP DTLS
}

// RadioStatsMonitor represents AP radio statistics monitoring configuration.
type RadioStatsMonitor struct {
	Enable       bool          `json:"enable"`        // AP radio stats collection enabled
	SampleIntvl  uint16        `json:"sample-intvl"`  // Sampling interval of radio statistics (seconds)
	AlarmsEnable []interface{} `json:"alarms-enable"` // AP radio statistics alarms enabled (YANG: IOS-XE 17.12.1+)
	RadioReset   bool          `json:"radio-reset"`   // Enable AP radio reset on radio stuck
}

// ZeroWtDfs represents Zero wait DFS information of the AP.
type ZeroWtDfs struct {
	ReserveChannel ReserveChannel `json:"reserve-channel"` // Reserved channel data
	Type           string         `json:"type"`            // CAC domain type
	// DfsChanInclList and DfsChanExclList would be added if present in JSON
}

// ReserveChannel represents reserved CAC channel information.
type ReserveChannel struct {
	Channel      uint8  `json:"channel"`       // Channel reserved for CAC
	ChannelWidth string `json:"channel-width"` // Channel width reserved for CAC
	State        string `json:"state"`         // CAC status of the reserved channel
}

// GnssInfo represents AP GNSS (Global Navigation Satellite System) information.
type GnssInfo struct {
	AntType          string  `json:"ant-type"`             // AP GNSS antenna type
	AntCableLength   uint16  `json:"ant-cable-length"`     // AP GNSS external antenna cable length (meters)
	AntennaProductID string  `json:"antenna-product-id"`   // AP GNSS antenna product id
	AntennaSn        *string `json:"antenna-sn,omitempty"` // AP GNSS antenna serial number (YANG: IOS-XE 17.18.1+)
}

// ApState represents AP state information.
type ApState struct {
	ApAdminState     string `json:"ap-admin-state"`     // Administrative state of the access point
	ApOperationState string `json:"ap-operation-state"` // Operational state of the access point
}

// ApModeData represents AP mode related information.
type ApModeData struct {
	HomeApEnabled bool         `json:"home-ap-enabled"` // Home AP feature enabled status
	ClearMode     bool         `json:"clear-mode"`      // Clear mode configuration status
	ApSubMode     string       `json:"ap-sub-mode"`     // Access point sub-mode
	WtpMode       string       `json:"wtp-mode"`        // Wireless termination point mode
	ApFabricData  ApFabricData `json:"ap-fabric-data"`  // AP fabric related attributes
	// Ap6GhzData will be added for 17.18.1+ compatibility
}

// ApFabricData represents AP fabric related attributes.
type ApFabricData struct {
	IsFabricAp bool `json:"is-fabric-ap"` // Fabric AP designation status
}

// ApLocation represents AP location information.
type ApLocation struct {
	Floor             int         `json:"floor"`              // Floor number
	Location          string      `json:"location"`           // Location description
	AaaLocation       AaaLocation `json:"aaa-location"`       // AAA location information
	FloorID           int         `json:"floor-id"`           // Floor identifier
	RangingCapability int         `json:"ranging-capability"` // Location ranging capability
}

// AaaLocation represents AAA location information.
type AaaLocation struct {
	CivicID string `json:"civic-id"` // Civic location identifier
	GeoID   string `json:"geo-id"`   // Geographic location identifier
	OperID  string `json:"oper-id"`  // Operator location identifier
}

// ApServices represents AP services information.
type ApServices struct {
	MonitorModeOptType string       `json:"monitor-mode-opt-type"` // Monitor mode optimization type
	ApDhcpServer       ApDhcpServer `json:"ap-dhcp-server"`        // AP DHCP server configuration
	TotSnifferRadio    int          `json:"tot-sniffer-radio"`     // Total number of sniffer radios
}

// ApDhcpServer represents AP DHCP server configuration.
type ApDhcpServer struct {
	IsDhcpServerEnabled bool `json:"is-dhcp-server-enabled"` // DHCP server enabled status
}

// XorPhyHtCap represents XOR PHY HT capabilities.
type XorPhyHtCap struct {
	Data XorPhyHtCapData `json:"data"`
}

// XorPhyHtCapData represents XOR PHY HT capability data.
type XorPhyHtCapData struct {
	VhtCapable bool `json:"vht-capable"` // VHT capability status
	HtCapable  bool `json:"ht-capable"`  // HT capability status
}

// XorPhyHeCap represents XOR PHY HE capabilities.
type XorPhyHeCap struct {
	Data XorPhyHeCapData `json:"data"`
}

// XorPhyHeCapData represents XOR PHY HE capability data.
type XorPhyHeCapData struct {
	HeEnabled              bool `json:"he-enabled"`                // HE protocol enable status
	HeCapable              bool `json:"he-capable"`                // HE capability status
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // Single user beamformer capability
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // Multi user beamformer capability
	HeStbcMode             int  `json:"he-stbc-mode"`              // HE STBC mode configuration
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // HE AMPDU TID bitmap
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // HE TX/RX MCS NSS capability
}

// StationCfg represents station configuration.
type StationCfg struct {
	CfgData StationCfgData `json:"cfg-data"`
}

// StationCfgData represents station configuration data.
type StationCfgData struct {
	StationCfgConfigType string `json:"station-cfg-config-type"` // Station configuration type
	MediumOccupancyLimit int    `json:"medium-occupancy-limit"`  // Medium occupancy limit
	CfpPeriod            int    `json:"cfp-period"`              // Contention Free Period
	CfpMaxDuration       int    `json:"cfp-max-duration"`        // CFP maximum duration
	Bssid                string `json:"bssid"`                   // Basic Service Set Identifier
	BeaconPeriod         int    `json:"beacon-period"`           // Beacon period interval
	CountryString        string `json:"country-string"`          // Country string identifier
}

// MultiDomainCap represents multi-domain capability configuration.
type MultiDomainCap struct {
	CfgData MultiDomainCapData `json:"cfg-data"`
}

// MultiDomainCapData represents multi-domain capability data.
type MultiDomainCapData struct {
	FirstChanNum    int `json:"first-chan-num"`     // First channel number
	NumChannels     int `json:"num-channels"`       // Number of channels
	MaxTxPowerLevel int `json:"max-tx-power-level"` // Maximum transmission power level
}

// PhyHtCfg represents PHY HT configuration.
type PhyHtCfg struct {
	CfgData PhyHtCfgData `json:"cfg-data"`
}

// PhyHtCfgData represents PHY HT configuration data.
type PhyHtCfgData struct {
	HtEnable               int    `json:"ht-enable"`                 // HT protocol enable status
	PhyHtCfgConfigType     string `json:"phy-ht-cfg-config-type"`    // PHY HT configuration type
	CurrFreq               int    `json:"curr-freq"`                 // Current frequency
	ChanWidth              int    `json:"chan-width"`                // Channel width
	ExtChan                int    `json:"ext-chan"`                  // Extension channel
	VhtEnable              bool   `json:"vht-enable"`                // VHT protocol enable status
	LegTxBfEnabled         int    `json:"leg-tx-bf-enabled"`         // Legacy TX beamforming enable status
	RrmChannelChangeReason string `json:"rrm-channel-change-reason"` // RRM channel change reason
	FreqString             string `json:"freq-string"`               // Frequency string
}

// PhyHtCapStruct represents PHY HT capability structure.
type PhyHtCapStruct struct {
	Data PhyHtCapStructData `json:"data"`
}

// PhyHtCapStructData represents PHY HT capability data.
type PhyHtCapStructData struct {
	VhtCapable bool `json:"vht-capable"` // VHT capability status
	HtCapable  bool `json:"ht-capable"`  // HT capability status
}

// PhyHeCapStruct represents PHY HE capability structure.
type PhyHeCapStruct struct {
	Data PhyHeCapStructData `json:"data"`
}

// PhyHeCapStructData represents PHY HE capability data.
type PhyHeCapStructData struct {
	HeEnabled              bool `json:"he-enabled"`                // HE protocol enable status
	HeCapable              bool `json:"he-capable"`                // HE capability status
	HeSingleUserBeamformer int  `json:"he-single-user-beamformer"` // Single user beamformer capability
	HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`  // Multi user beamformer capability
	HeStbcMode             int  `json:"he-stbc-mode"`              // HE STBC mode configuration
	HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`       // HE AMPDU TID bitmap
	HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`      // HE TX/RX MCS NSS capability
}

// ChanPwrInfo represents channel power information.
type ChanPwrInfo struct {
	Data ChanPwrInfoData `json:"data"`
}

// ChanPwrInfoData represents channel power information data.
type ChanPwrInfoData struct {
	AntennaGain    int         `json:"antenna-gain"`     // Antenna gain value
	IntAntennaGain int         `json:"int-antenna-gain"` // Internal antenna gain
	ExtAntennaGain int         `json:"ext-antenna-gain"` // External antenna gain
	ChanPwrList    ChanPwrList `json:"chan-pwr-list"`    // Channel power list
}

// ChanPwrList represents channel power list.
type ChanPwrList struct {
	ChanPwr []ChanPwr `json:"chan-pwr"`
}

// ChanPwr represents individual channel power.
type ChanPwr struct {
	Chan int `json:"chan"` // Channel number
}

// SnifferCfg represents sniffer configuration.
type SnifferCfg struct {
	SnifferEnabled bool `json:"sniffer-enabled"` // Sniffer enable status
}

// RadioBandInfo represents radio band information.
type RadioBandInfo struct {
	BandID                 uint8          `json:"band-id"`                      // Radio band identifier
	RegDomainCode          uint16         `json:"reg-domain-code"`              // Regulatory domain code
	RegulatoryDomain       string         `json:"regulatory-domain"`            // Regulatory domain name
	MacOperCfg             MacOperCfg     `json:"mac-oper-cfg,omitempty"`       // MAC operation configuration
	PhyTxPwrCfg            PhyTxPwrCfg    `json:"phy-tx-pwr-cfg,omitempty"`     // PHY TX power configuration
	PhyTxPwrLvlCfg         PhyTxPwrLvlCfg `json:"phy-tx-pwr-lvl-cfg,omitempty"` // PHY TX power level configuration
	AntennaCfg             AntennaCfg     `json:"antenna-cfg,omitempty"`        // Antenna configuration
	Dot11acChannelWidthCap uint8          `json:"dot11ac-channel-width-cap"`    // 802.11ac channel width capability
	Secondary80Channel     uint16         `json:"secondary-80-channel"`         // Secondary 80MHz channel
	SiaParams              SiaParams      `json:"sia-params,omitempty"`         // Self Identifying Antenna parameters
}

// MacOperCfg represents MAC operation configuration.
type MacOperCfg struct {
	CfgData MacOperCfgData `json:"cfg-data"`
}

// MacOperCfgData represents MAC operation configuration data.
type MacOperCfgData struct {
	MacOperationConfigType string `json:"mac-operation-config-type"` // MAC operation configuration type
	RtsThreshold           uint16 `json:"rts-threshold"`             // RTS threshold value
	ShortRetryLimit        uint8  `json:"short-retry-limit"`         // Short retry limit
	LongRetryLimit         uint8  `json:"long-retry-limit"`          // Long retry limit
	FragThreshold          uint16 `json:"frag-threshold"`            // Fragmentation threshold
	MaxTxLifeTime          uint16 `json:"max-tx-life-time"`          // Maximum transmission lifetime
	MaxRxLifeTime          uint16 `json:"max-rx-life-time"`          // Maximum reception lifetime
}

// PhyTxPwrCfg represents PHY TX power configuration.
type PhyTxPwrCfg struct {
	CfgData PhyTxPwrCfgData `json:"cfg-data"`
}

// PhyTxPwrCfgData represents PHY TX power configuration data.
type PhyTxPwrCfgData struct {
	PhyTxPowerConfigType string `json:"phy-tx-power-config-type"` // PHY TX power configuration type
	CurrentTxPowerLevel  uint8  `json:"current-tx-power-level"`   // Current transmission power level
}

// PhyTxPwrLvlCfg represents PHY TX power level configuration.
type PhyTxPwrLvlCfg struct {
	CfgData PhyTxPwrLvlCfgData `json:"cfg-data"`
}

// PhyTxPwrLvlCfgData represents PHY TX power level configuration data.
type PhyTxPwrLvlCfgData struct {
	NumSuppPowerLevels uint8 `json:"num-supp-power-levels"` // Number of supported power levels
	TxPowerLevel1      int8  `json:"tx-power-level-1"`      // Transmission power level 1
	TxPowerLevel2      int8  `json:"tx-power-level-2"`      // Transmission power level 2
	TxPowerLevel3      int8  `json:"tx-power-level-3"`      // Transmission power level 3
	TxPowerLevel4      int8  `json:"tx-power-level-4"`      // Transmission power level 4
	TxPowerLevel5      int8  `json:"tx-power-level-5"`      // Transmission power level 5
	TxPowerLevel6      int8  `json:"tx-power-level-6"`      // Transmission power level 6
	TxPowerLevel7      int8  `json:"tx-power-level-7"`      // Transmission power level 7
	TxPowerLevel8      int8  `json:"tx-power-level-8"`      // Transmission power level 8
	CurrTxPowerInDbm   int8  `json:"curr-tx-power-in-dbm"`  // Current transmission power in dBm
}

// AntennaCfg represents antenna configuration.
type AntennaCfg struct {
	CfgData AntennaCfgData `json:"cfg-data"`
}

// AntennaCfgData represents antenna configuration data.
type AntennaCfgData struct {
	DiversitySelection string `json:"diversity-selection"` // Antenna diversity selection
	AntennaMode        string `json:"antenna-mode"`        // Antenna mode configuration
	NumOfAntennas      uint8  `json:"num-of-antennas"`     // Number of antennas
}

// SiaParams represents Self Identifying Antenna parameters.
type SiaParams struct {
	IsRptncPresent bool   `json:"is-rptnc-present"` // RPTNC presence status
	IsDartPresent  bool   `json:"is-dart-present"`  // DART presence status
	AntennaIfType  string `json:"antenna-if-type"`  // Antenna interface type
	AntennaGain    uint8  `json:"antenna-gain"`     // Antenna gain value
	Marlin4Present bool   `json:"marlin4-present"`  // Marlin4 presence status
	DmServType     string `json:"dm-serv-type"`     // Device management service type
}

// VapOperConfig represents VAP operational configuration.
type VapOperConfig struct {
	ApVapID         uint8  `json:"ap-vap-id"`         // AP VAP identifier
	WlanID          uint8  `json:"wlan-id"`           // WLAN identifier
	BssidMac        string `json:"bssid-mac"`         // BSSID MAC address
	WtpMac          string `json:"wtp-mac"`           // MAC address of AP radio interface
	WlanProfileName string `json:"wlan-profile-name"` // WLAN profile name
	SSID            string `json:"ssid"`              // Service Set Identifier
}

// DualRadioModeCfg represents dual radio mode configuration.
type DualRadioModeCfg struct {
	DualRadioMode    string `json:"dual-radio-mode"`    // Dual radio mode configuration
	DualRadioCapable string `json:"dual-radio-capable"` // Dual radio capability status
	DualRadioModeOp  string `json:"dual-radio-mode-op"` // Dual radio mode operation
}

// BssColorCfg represents BSS color configuration.
type BssColorCfg struct {
	BssColorCapable    bool   `json:"bss-color-capable"`     // BSS color capability status
	BssColor           uint8  `json:"bss-color"`             // BSS color value
	BssColorConfigType string `json:"bss-color-config-type"` // BSS color configuration type
}

// BoardDataOpt represents board data options.
type BoardDataOpt struct {
	JoinPriority uint8 `json:"join-priority"` // Join priority value
}

// DescriptorData represents descriptor data.
type DescriptorData struct {
	RadioSlotsInUse        uint8 `json:"radio-slots-in-use"`      // Number of radio slots in use
	EncryptionCapabilities bool  `json:"encryption-capabilities"` // Encryption capabilities status
}

// ApProv represents AP provisioning information.
type ApProv struct {
	IsUniversal          bool   `json:"is-universal"`           // Universal AP status
	UniversalPrimeStatus string `json:"universal-prime-status"` // Universal prime status
}

// ApModels represents AP model information.
type ApModels struct {
	Model string `json:"model"` // Access point model name
}

// TempInfo represents temperature information.
type TempInfo struct {
	Degree       int    `json:"degree"`        // Temperature in degrees
	TempStatus   string `json:"temp-status"`   // Temperature status indicator
	HeaterStatus string `json:"heater-status"` // Heater operational status
}

// TagInfo represents AP tag information.
type TagInfo struct {
	TagSource         string          `json:"tag-source"`          // Tag assignment source
	IsApMisconfigured bool            `json:"is-ap-misconfigured"` // AP misconfiguration status
	ResolvedTagInfo   ResolvedTagInfo `json:"resolved-tag-info"`   // Resolved tag information
	PolicyTagInfo     PolicyTagInfo   `json:"policy-tag-info"`     // Policy tag information
	SiteTag           SiteTag         `json:"site-tag"`            // Site tag information
	RfTag             RfTag           `json:"rf-tag"`              // RF tag information
	FilterInfo        FilterInfo      `json:"filter-info"`         // Filter information
	IsDtlsLscFbkAp    bool            `json:"is-dtls-lsc-fbk-ap"`  // DTLS LSC fallback AP status
}

// ResolvedTagInfo represents resolved tag information.
type ResolvedTagInfo struct {
	ResolvedPolicyTag string `json:"resolved-policy-tag"` // Resolved policy tag name
	ResolvedSiteTag   string `json:"resolved-site-tag"`   // Resolved site tag name
	ResolvedRfTag     string `json:"resolved-rf-tag"`     // Resolved RF tag name
}

// PolicyTagInfo represents policy tag information.
type PolicyTagInfo struct {
	PolicyTagName string `json:"policy-tag-name"` // Policy tag name
}

// SiteTag represents site tag information.
type SiteTag struct {
	SiteTagName string `json:"site-tag-name"` // Site tag name
	ApProfile   string `json:"ap-profile"`    // AP profile name
	FlexProfile string `json:"flex-profile"`  // FlexConnect profile name
}

// RfTag represents RF tag information.
type RfTag struct {
	RfTagName string `json:"rf-tag-name"` // RF tag name
}

// FilterInfo represents filter information.
type FilterInfo struct {
	FilterName string `json:"filter-name"` // Filter name
}

// Tunnel represents tunnel configuration.
type Tunnel struct {
	PreferredMode string `json:"preferred-mode"` // Preferred tunnel mode
	UDPLite       string `json:"udp-lite"`       // UDP-Lite configuration
}

// ExternalModuleData represents external module data.
type ExternalModuleData struct {
	XmData             XmData  `json:"xm-data"`               // XM module data
	UsbData            UsbData `json:"usb-data"`              // USB module data
	UsbOverride        bool    `json:"usb-override"`          // USB override configuration
	IsExtModuleEnabled bool    `json:"is-ext-module-enabled"` // External module enabled status
}

// XmData represents XM module data.
type XmData struct {
	IsModulePresent bool `json:"is-module-present"` // XM module presence status
	Xm              Xm   `json:"xm"`                // XM module information
}

// UsbData represents USB module data.
type UsbData struct {
	IsModulePresent bool `json:"is-module-present"` // USB module presence status
	Xm              Xm   `json:"xm"`                // USB module information
}

// Xm represents external module information.
type Xm struct {
	NumericID          uint32 `json:"numeric-id"`           // Module numeric identifier
	MaxPower           uint16 `json:"max-power"`            // Module maximum power consumption
	SerialNumberString string `json:"serial-number-string"` // Module serial number
	ProductIDString    string `json:"product-id-string"`    // Module product ID
	ModuleType         string `json:"module-type"`          // Module type description
	ModuleDescription  string `json:"module-description"`   // Module description
}

// DeviceDetail represents device detail information.
type DeviceDetail struct {
	StaticInfo  StaticInfo  `json:"static-info"`  // Static device information
	DynamicInfo DynamicInfo `json:"dynamic-info"` // Dynamic device information
	WtpVersion  WtpVersion  `json:"wtp-version"`  // WTP version information
}

// StaticInfo represents static information.
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
	BoardDataOpt   BoardDataOpt   `json:"board-data-opt,omitempty"`
	DescriptorData DescriptorData `json:"descriptor-data,omitempty"`
	ApProv         ApProv         `json:"ap-prov,omitempty"`
	ApModels       ApModels       `json:"ap-models,omitempty"`
	NumPorts       uint8          `json:"num-ports,omitempty"`
	NumSlots       uint8          `json:"num-slots,omitempty"`
	WtpModelType   uint16         `json:"wtp-model-type,omitempty"`
	ApCapability   string         `json:"ap-capability,omitempty"`
	IsMmOpt        bool           `json:"is-mm-opt,omitempty"`
	ApImageName    string         `json:"ap-image-name,omitempty"`
}

// DynamicInfo represents dynamic information.
type DynamicInfo struct {
	ApCrashData struct {
		ApCrashFile           string `json:"ap-crash-file"`
		ApRadio2GCrashFile    string `json:"ap-radio-2g-crash-file"`
		ApRadio5GCrashFile    string `json:"ap-radio-5g-crash-file"`
		ApRadio6GCrashFile    string `json:"ap-radio-6g-crash-file"`
		ApRad5GSlot2CrashFile string `json:"ap-rad-5g-slot2-crash-file"`
	} `json:"ap-crash-data"`
	LedStateEnabled  bool     `json:"led-state-enabled,omitempty"`
	ResetButtonState bool     `json:"reset-button-state,omitempty"`
	LedFlashEnabled  bool     `json:"led-flash-enabled,omitempty"`
	FlashSec         uint16   `json:"flash-sec,omitempty"`
	TempInfo         TempInfo `json:"temp-info,omitempty"`
	LedFlashExpiry   string   `json:"led-flash-expiry,omitempty"`
}

// WtpVersion represents WTP version information.
type WtpVersion struct {
	BackupSwVersion struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"backup-sw-version"`
	MiniIosVersion struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"mini-ios-version,omitempty"`
	SwVer struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"sw-ver,omitempty"`
	BootVer struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"boot-ver,omitempty"`
	SwVersion string `json:"sw-version,omitempty"`
}

// ApNameMacMap represents AP name to MAC address mapping.
type ApNameMacMap struct {
	WtpName string `json:"wtp-name"` // Wireless termination point name
	WtpMac  string `json:"wtp-mac"`  // WTP MAC address
	EthMac  string `json:"eth-mac"`  // Ethernet MAC address
}

// WtpSlotWlanStats represents WTP slot WLAN statistics.
type WtpSlotWlanStats struct {
	WtpMac      string `json:"wtp-mac"`      // WTP MAC address
	SlotID      int    `json:"slot-id"`      // Radio slot identifier
	WlanID      int    `json:"wlan-id"`      // WLAN identifier
	BssidMac    string `json:"bssid-mac"`    // BSSID MAC address
	Ssid        string `json:"ssid"`         // Service Set Identifier
	BytesRx     string `json:"bytes-rx"`     // Bytes received
	BytesTx     string `json:"bytes-tx"`     // Bytes transmitted
	PktsRx      string `json:"pkts-rx"`      // Packets received
	PktsTx      string `json:"pkts-tx"`      // Packets transmitted
	DataRetries string `json:"data-retries"` // Data retry count
}

// EthernetMacWtpMacMap represents Ethernet MAC to WTP MAC mapping.
type EthernetMacWtpMacMap struct {
	EthernetMac string `json:"ethernet-mac"` // Ethernet MAC address
	WtpMac      string `json:"wtp-mac"`      // WTP MAC address
}

// RadioOperStats represents radio operational statistics.
type RadioOperStats struct {
	ApMac                 string      `json:"ap-mac"`
	SlotID                int         `json:"slot-id"`
	AidUserList           interface{} `json:"aid-user-list"` // Number of users associated with this radio
	TxFragmentCount       int         `json:"tx-fragment-count"`
	MultipleRetryCount    int         `json:"multiple-retry-count"` // Multiple retry count
	MulticastTxFrameCnt   int         `json:"multicast-tx-frame-cnt"`
	FailedCount           int         `json:"failed-count"`
	RetryCount            int         `json:"retry-count"`
	FrameDuplicateCount   int         `json:"frame-duplicate-count"`
	AckFailureCount       int         `json:"ack-failure-count"`
	FcsErrorCount         int         `json:"fcs-error-count"`
	MacDecryErrFrameCount int         `json:"mac-decry-err-frame-count"`
	MacMicErrFrameCount   int         `json:"mac-mic-err-frame-count"`
	MulticastRxFrameCnt   int         `json:"multicast-rx-frame-cnt"`
	NoiseFloor            int         `json:"noise-floor"`
	RtsFailureCount       int         `json:"rts-failure-count"`
	RtsSuccessCount       int         `json:"rts-success-count"`
	RxCtrlFrameCount      int         `json:"rx-ctrl-frame-count"`
	RxDataFrameCount      int         `json:"rx-data-frame-count"`
	RxDataPktCount        int         `json:"rx-data-pkt-count"`
	RxErrorFrameCount     int         `json:"rx-error-frame-count"`
	RxFragmentCount       int         `json:"rx-fragment-count"`
	RxMgmtFrameCount      int         `json:"rx-mgmt-frame-count"`
	TxCtrlFrameCount      int         `json:"tx-ctrl-frame-count"`
	TxDataFrameCount      int         `json:"tx-data-frame-count"`
	TxDataPktCount        int         `json:"tx-data-pkt-count"`
	TxFrameCount          int         `json:"tx-frame-count"`
	TxMgmtFrameCount      int         `json:"tx-mgmt-frame-count"`
	WepUndecryptableCount int         `json:"wep-undecryptable-count"`
	ApRadioStats          interface{} `json:"ap-radio-stats,omitempty"`
}

// EthernetIfStats represents Ethernet interface statistics.
type EthernetIfStats struct {
	WtpMac           string `json:"wtp-mac"`
	IfIndex          int    `json:"if-index"`
	IfName           string `json:"if-name"`
	RxPkts           int    `json:"rx-pkts"`
	TxPkts           int    `json:"tx-pkts"`
	OperStatus       string `json:"oper-status"`
	RxUcastPkts      int    `json:"rx-ucast-pkts"`
	RxNonUcastPkts   int    `json:"rx-non-ucast-pkts"`
	TxUcastPkts      int    `json:"tx-ucast-pkts"`
	TxNonUcastPkts   int    `json:"tx-non-ucast-pkts"`
	Duplex           int    `json:"duplex"`
	LinkSpeed        int    `json:"link-speed"`
	RxTotalBytes     int    `json:"rx-total-bytes"`
	TxTotalBytes     int    `json:"tx-total-bytes"`
	InputCrc         int    `json:"input-crc"`
	InputAborts      int    `json:"input-aborts"`
	InputErrors      int    `json:"input-errors"`
	InputFrames      int    `json:"input-frames"`
	InputOverrun     int    `json:"input-overrun"`
	InputDrops       int    `json:"input-drops"`
	InputResource    int    `json:"input-resource"`
	UnknownProtocol  int    `json:"unknown-protocol"`
	Runts            int    `json:"runts"`
	Giants           int    `json:"giants"`
	Throttle         int    `json:"throttle"`
	Resets           int    `json:"resets"`
	OutputCollision  int    `json:"output-collision"`
	OutputNoBuffer   int    `json:"output-no-buffer"`
	OutputResource   int    `json:"output-resource"`
	OutputUnderrun   int    `json:"output-underrun"`
	OutputErrors     int    `json:"output-errors"`
	OutputTotalDrops int    `json:"output-total-drops"`
}

// EwlcWncdStats represents EWLC WNCD statistics.
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
	DownloadsComplete   int         `json:"downloads-complete"`
	DownloadsInProgress int         `json:"downloads-in-progress"`
	WlcPredownloadStats interface{} `json:"wlc-predownload-stats,omitempty"`
}

// ApIoxOperData represents AP IOx operational data.
type ApIoxOperData struct {
	ApMac        string `json:"ap-mac"`        // Access point MAC address
	ApphostState string `json:"apphost-state"` // Application host state
	CafToken     string `json:"caf-token"`     // CAF authentication token
	CafPort      int    `json:"caf-port"`      // CAF service port
}

// QosGlobalStats represents QoS global statistics.
type QosGlobalStats struct {
	QosClientVoiceStats struct {
		TotalNumOfTspecRcvd       int `json:"total-num-of-tspec-rcvd"`
		NewTspecFromAssocReq      int `json:"new-tspec-from-assoc-req"`
		TspecRenewalFromAssocReq  int `json:"tspec-renewal-from-assoc-req"`
		NewTspecAsAddTS           int `json:"new-tspec-as-add-ts"`
		TspecRenewalFromAddTS     int `json:"tspec-renewal-from-add-ts"`
		NumOfActiveTspecCalls     int `json:"num-of-active-tspec-calls"`
		NumOfActiveSIPCalls       int `json:"num-of-active-sip-calls"`
		NumOfCallsAccepted        int `json:"num-of-calls-accepted"`
		NumOfCallsRejectedInsufBw int `json:"num-of-calls-rejected-insuf-bw"`
		NumOfCallsRejectedPhyRate int `json:"num-of-calls-rejected-phy-rate"`
		NumOfCallsRejectedQos     int `json:"num-of-calls-rejected-qos"`
		NumOfCallsRejInvalidTspec int `json:"num-of-calls-rej-invalid-tspec"`
		NumOfRoamCallsAccepted    int `json:"num-of-roam-calls-accepted"`
		NumOfRoamCallsRejected    int `json:"num-of-roam-calls-rejected"`
		TspecProcessFailedGetRec  int `json:"tspec-process-failed-get-rec"`
		TotalNumOfCallReport      int `json:"total-num-of-call-report"`
		TotalSIPFailureTrapSend   int `json:"total-sip-failure-trap-send"`
		TotalSIPInviteOnCaller    int `json:"total-sip-invite-on-caller"`
		TotalSIPInviteOnCallee    int `json:"total-sip-invite-on-callee"`
	} `json:"qos-client-voice-stats"`
}

// ApOperInternalData represents internal AP operational data.
type ApOperInternalData struct {
	WtpMac                 string         `json:"wtp-mac"`
	RadioID                int            `json:"radio-id"`
	ApAntennaBandMode      string         `json:"ap-antenna-band-mode"`
	LinkEncryptionEnabled  bool           `json:"link-encryption-enabled"`
	ApRemoteDebugMode      bool           `json:"ap-remote-debug-mode"`
	ApRole                 string         `json:"ap-role"`
	ApIndoorMode           bool           `json:"ap-indoor-mode"`
	MaxClientsAllowed      int            `json:"max-clients-allowed"`
	IsLocalNet             bool           `json:"is-local-net"`
	Ipv4TcpMss             TCPMssConfig   `json:"ipv4-tcp-mss"`
	Ipv6TcpMss             TCPMssConfig   `json:"ipv6-tcp-mss"`
	RangingMode            string         `json:"ranging-mode"`
	PowerProfile           string         `json:"power-profile"`
	PwrProfType            string         `json:"pwr-prof-type"`
	PwrCalProfile          string         `json:"pwr-cal-profile"`
	PersistentSsid         PersistentSsid `json:"persistent-ssid"`
	ProvSsid               bool           `json:"prov-ssid"`
	PrimingProfile         string         `json:"priming-profile"`
	PrimingProfileSrc      string         `json:"priming-profile-src"`
	PrimingFilter          string         `json:"priming-filter"`
	PmkBsSenderAddr        string         `json:"pmk-bs-sender-addr"`
	PmkBsReceiverAddr      string         `json:"pmk-bs-receiver-addr"`
	Accounting             interface{}    `json:"accounting,omitempty"`
	ApDnaData              interface{}    `json:"ap-dna-data,omitempty"`
	ApGasRateLimitCfg      interface{}    `json:"ap-gas-rate-limit-cfg,omitempty"`
	ApIPData               interface{}    `json:"ap-ip-data,omitempty"`
	ApLoginCredentials     interface{}    `json:"ap-login-credentials,omitempty"`
	ApMgmt                 interface{}    `json:"ap-mgmt,omitempty"`
	ApNtpServerInfoCfg     interface{}    `json:"ap-ntp-server-info-cfg,omitempty"`
	ApNtpSyncStatus        interface{}    `json:"ap-ntp-sync-status,omitempty"`
	ApPmkPropagationStatus interface{}    `json:"ap-pmk-propagation-status,omitempty"`
	ApPow                  interface{}    `json:"ap-pow,omitempty"`
	ApPrimeInfo            interface{}    `json:"ap-prime-info,omitempty"`
	ApPrimingOverride      interface{}    `json:"ap-priming-override,omitempty"`
	ApSysStats             interface{}    `json:"ap-sys-stats,omitempty"`
	ApTzConfig             interface{}    `json:"ap-tz-config,omitempty"`
	ApUdpliteInfo          interface{}    `json:"ap-udplite-info,omitempty"`
	AuxClientInterfaceData interface{}    `json:"aux-client-interface-data,omitempty"`
	InfrastructureMfp      interface{}    `json:"infrastructure-mfp,omitempty"`
	KernelCoredump         interface{}    `json:"kernel-coredump,omitempty"`
	LinkAudit              interface{}    `json:"link-audit,omitempty"`
	OeapAudit              interface{}    `json:"oeap-audit,omitempty"`
	Retransmit             interface{}    `json:"retransmit,omitempty"`
	Syslog                 interface{}    `json:"syslog,omitempty"`
	Timer                  interface{}    `json:"timer,omitempty"`
}

// RlanOper represents RLAN operational data.
type RlanOper struct {
	WtpMac         string `json:"wtp-mac"`          // WTP MAC address
	RlanPortID     int    `json:"rlan-port-id"`     // RLAN port identifier
	RlanOperState  bool   `json:"rlan-oper-state"`  // RLAN operational state
	RlanPortStatus bool   `json:"rlan-port-status"` // RLAN port status
	RlanVlanValid  bool   `json:"rlan-vlan-valid"`  // RLAN VLAN validity
	RlanVlanID     int    `json:"rlan-vlan-id"`     // RLAN VLAN identifier
	RlanPoeState   string `json:"rlan-poe-state"`   // RLAN PoE state
	PowerLevelID   int    `json:"power-level-id"`   // Power level identifier
}

// EwlcMewlcPredownloadRec represents EWLC MEWLC predownload record.
type EwlcMewlcPredownloadRec struct {
	PredState                    string `json:"pred-state"`                     // Predownload state
	MeCapableApCount             int    `json:"me-capable-ap-count"`            // ME capable AP count
	ControllerPredownloadVersion string `json:"controller-predownload-version"` // Controller predownload version
}

// CdpCacheData represents CDP cache data.
type CdpCacheData struct {
	MacAddr                string       `json:"mac-addr"`                   // MAC address
	CdpCacheDeviceID       string       `json:"cdp-cache-device-id"`        // CDP device identifier
	ApName                 string       `json:"ap-name"`                    // Access point name
	LastUpdatedTime        time.Time    `json:"last-updated-time"`          // Last update timestamp
	Version                int          `json:"version"`                    // Version number
	WtpMacAddr             string       `json:"wtp-mac-addr"`               // WTP MAC address
	DeviceIndex            int          `json:"device-index"`               // Device index
	IPAddress              CdpIPAddress `json:"ip-address"`                 // IP address information
	CdpAddrCount           int          `json:"cdp-addr-count"`             // CDP address count
	CdpCacheApAddress      string       `json:"cdp-cache-ap-address"`       // CDP cache AP address
	CdpCacheDevicePort     string       `json:"cdp-cache-device-port"`      // CDP cache device port
	CdpCacheDuplex         string       `json:"cdp-cache-duplex"`           // CDP cache duplex mode
	CdpCacheIfIndex        int          `json:"cdp-cache-if-index"`         // CDP cache interface index
	CdpCacheInterfaceSpeed int          `json:"cdp-cache-interface-speed"`  // CDP cache interface speed
	CdpCacheIPAddressValue string       `json:"cdp-cache-ip-address-value"` // CDP cache IP address value
	CdpCacheLocalPort      string       `json:"cdp-cache-local-port"`       // CDP cache local port
	CdpCachePlatform       string       `json:"cdp-cache-platform"`         // CDP cache platform
	CdpCacheVersion        string       `json:"cdp-cache-version"`          // CDP cache version
	CdpCapabilitiesString  string       `json:"cdp-capabilities-string"`    // CDP capabilities string
}

// LldpNeigh represents LLDP neighbor information.
type LldpNeigh struct {
	WtpMac          string `json:"wtp-mac"`          // WTP MAC address
	NeighMac        string `json:"neigh-mac"`        // Neighbor MAC address
	PortID          string `json:"port-id"`          // Port identifier
	LocalPort       string `json:"local-port"`       // Local port
	SystemName      string `json:"system-name"`      // System name
	PortDescription string `json:"port-description"` // Port description
	Capabilities    string `json:"capabilities"`     // Device capabilities
	MgmtAddr        string `json:"mgmt-addr"`        // Management address
}

// TpCertInfo represents trustpoint certificate information.
type TpCertInfo struct {
	Trustpoint Trustpoint `json:"trustpoint"`
}

// Trustpoint represents trustpoint information.
type Trustpoint struct {
	TrustpointName     string  `json:"trustpoint-name"`      // Trustpoint identifier (Live: IOS-XE 17.12.5)
	IsCertAvailable    bool    `json:"is-cert-available"`    // Certificate availability status (Live: IOS-XE 17.12.5)
	IsPrivkeyAvailable bool    `json:"is-privkey-available"` // Private key availability status (Live: IOS-XE 17.12.5)
	CertHash           string  `json:"cert-hash"`            // Certificate hash (Live: IOS-XE 17.12.5)
	CertType           string  `json:"cert-type"`            // Certificate type (Live: IOS-XE 17.12.5)
	FipsSuitability    string  `json:"fips-suitability"`     // FIPS suitability status (Live: IOS-XE 17.12.5)
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
	CountryCode         string      `json:"country-code"`                 // Country code
	CountryString       string      `json:"country-string"`               // Country string
	RegDomainStr80211Bg string      `json:"reg-domain-str-80211bg"`       // Regulatory domain string for 802.11bg
	RegDomainStr80211A  string      `json:"reg-domain-str-80211a"`        // Regulatory domain string for 802.11a
	CountrySupported    bool        `json:"country-supported"`            // Country support status
	Channels11a         interface{} `json:"channels-11a,omitempty"`       // Available channels for 802.11a
	Channels11bg        interface{} `json:"channels-11bg,omitempty"`      // Available channels for 802.11bg
	ChannelsString11a   string      `json:"channels-string-11a"`          // Channel string for 802.11a
	ChannelsString11bg  string      `json:"channels-string-11bg"`         // Channel string for 802.11bg
	DcaChannels11a      interface{} `json:"dca-channels-11a,omitempty"`   // DCA channels for 802.11a
	DcaChannels11bg     interface{} `json:"dca-channels-11bg,omitempty"`  // DCA channels for 802.11bg
	RadarChannels11a    interface{} `json:"radar-channels-11a,omitempty"` // Radar channels for 802.11a
	RegDom6ghz          interface{} `json:"reg-dom-6ghz,omitempty"`       // Regulatory domain for 6GHz
	ChanInfo6ghz        interface{} `json:"chan-info-6ghz,omitempty"`     // Channel information for 6GHz
}

// SuppCountryOper represents supported country operational data.
type SuppCountryOper struct {
	CountryCode      string      `json:"country-code"`                  // Country code
	CountryString    string      `json:"country-string"`                // Country string
	CountryCodeIso   string      `json:"country-code-iso"`              // ISO country code
	ChanList24ghz    interface{} `json:"chan-list-24ghz,omitempty"`     // Channel list for 2.4GHz
	ChanList5ghz     interface{} `json:"chan-list-5ghz,omitempty"`      // Channel list for 5GHz
	ChanList6ghz     interface{} `json:"chan-list-6ghz,omitempty"`      // Channel list for 6GHz
	ChanListDca24ghz interface{} `json:"chan-list-dca-24ghz,omitempty"` // DCA channel list for 2.4GHz
	ChanListDca5ghz  interface{} `json:"chan-list-dca-5ghz,omitempty"`  // DCA channel list for 5GHz
	ChanListDca6ghz  interface{} `json:"chan-list-dca-6ghz,omitempty"`  // DCA channel list for 6GHz
	ChanListPsc6ghz  interface{} `json:"chan-list-psc-6ghz,omitempty"`  // PSC channel list for 6GHz
	RegDom24ghz      interface{} `json:"reg-dom-24ghz,omitempty"`       // Regulatory domain for 2.4GHz
	RegDom5ghz       interface{} `json:"reg-dom-5ghz,omitempty"`        // Regulatory domain for 5GHz
	RegDom6ghz       interface{} `json:"reg-dom-6ghz,omitempty"`        // Regulatory domain for 6GHz
}

// ApNhGlobalData represents AP next hop global data.
type ApNhGlobalData struct {
	AlgorithmRunning   bool `json:"algorithm-running"`     // Algorithm running status
	AlgorithmItrCount  int  `json:"algorithm-itr-count"`   // Algorithm iteration count
	IdealCapacityPerRg int  `json:"ideal-capacity-per-rg"` // Ideal capacity per RG
	NumOfNeighborhood  int  `json:"num-of-neighborhood"`   // Number of neighborhoods
}

// ApImagePrepareLocation represents AP image prepare location.
type ApImagePrepareLocation struct {
	Index     int         `json:"index"`      // Index
	ImageFile string      `json:"image-file"` // Image file name
	ImageData []ImageData `json:"image-data"` // Image data information
}

// ImageData represents image data information.
type ImageData struct {
	ImageName     string   `json:"image-name"`     // Image name
	ImageLocation string   `json:"image-location"` // Image location
	ImageVersion  string   `json:"image-version"`  // Image version
	IsNew         bool     `json:"is-new"`         // New image flag
	FileSize      string   `json:"file-size"`      // File size
	ApModelList   []string `json:"ap-model-list"`  // AP model list
}

// ApImageActiveLocation represents AP image active location.
type ApImageActiveLocation struct {
	Index                          int    `json:"index"`      // Index
	ImageFile                      string `json:"image-file"` // Image file name
	ApImageActiveLocationImageData []struct {
		ImageName                                 string   `json:"image-name"`     // Image name
		ImageLocation                             string   `json:"image-location"` // Image location
		ImageVersion                              string   `json:"image-version"`  // Image version
		IsNew                                     bool     `json:"is-new"`         // New image flag
		FileSize                                  string   `json:"file-size"`      // File size
		ApImageActiveLocationImageDataApModelList []string `json:"ap-model-list"`  // AP model list
	} `json:"image-data"` // Image data information
}

// TCPMssConfig represents TCP MSS adjustment configuration.
type TCPMssConfig struct {
	TCPAdjustMssState bool `json:"tcp-adjust-mss-state"` // TCP MSS adjustment state
	TCPAdjustMssSize  int  `json:"tcp-adjust-mss-size"`  // TCP MSS adjustment size
}

// PersistentSsid represents persistent SSID configuration.
type PersistentSsid struct {
	IsPersistentSsidEnabled bool `json:"is-persistent-ssid-enabled"` // Persistent SSID enabled status
}

// CdpIPAddress represents CDP IP address information.
type CdpIPAddress struct {
	IPAddressValue []string `json:"ip-address-value"` // CDP IP address values
}
