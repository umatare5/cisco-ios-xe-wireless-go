package rrm

import "time"

// CiscoIOSXEWirelessRRMOper represents RRM operational response data.
type CiscoIOSXEWirelessRRMOper struct {
	CiscoIOSXEWirelessRRMOperData struct {
		ApAutoRFDot11Data   []ApAutoRFDot11Data   `json:"ap-auto-rf-dot11-data,omitempty"`  // AP auto RF 802.11 data (Live: IOS-XE 17.12.6a)
		ApDot11RadarData    []ApDot11RadarData    `json:"ap-dot11-radar-data,omitempty"`    // AP radar detection data (Live: IOS-XE 17.12.6a)
		ApDot11SpectrumData []ApDot11SpectrumData `json:"ap-dot11-spectrum-data,omitempty"` // AP spectrum analysis data (Live: IOS-XE 17.12.6a)
		RRMMeasurement      []RRMMeasurement      `json:"rrm-measurement,omitempty"`        // RRM measurement data (Live: IOS-XE 17.12.6a)
		RadioSlot           []RadioSlot           `json:"radio-slot,omitempty"`             // Radio slot operational data (Live: IOS-XE 17.12.6a)
		MainData            []MainData            `json:"main-data,omitempty"`              // Main RRM data by PHY type (Live: IOS-XE 17.12.6a)
		RegDomainOper       *RegDomainOper        `json:"reg-domain-oper,omitempty"`        // Regulatory domain operational data (Live: IOS-XE 17.12.6a)
		SpectrumDeviceTable []SpectrumDeviceTable `json:"spectrum-device-table,omitempty"`  // Spectrum device detection table (Live: IOS-XE 17.12.6a)
		SpectrumAqTable     []SpectrumAqTable     `json:"spectrum-aq-table,omitempty"`      // Spectrum air quality table (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RRM operational data container
}

// CiscoIOSXEWirelessRRMOperApAutoRFDot11Data represents the AP auto RF 802.11 operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperApAutoRFDot11Data struct {
	ApAutoRFDot11Data []ApAutoRFDot11Data `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-auto-rf-dot11-data"`
}

// CiscoIOSXEWirelessRRMOperApDot11RadarData represents the AP radar detection operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperApDot11RadarData struct {
	ApDot11RadarData []ApDot11RadarData `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-radar-data"`
}

// CiscoIOSXEWirelessRRMOperApDot11SpectrumData wraps spectrum intelligence data.
type CiscoIOSXEWirelessRRMOperApDot11SpectrumData struct {
	ApDot11SpectrumData []ApDot11SpectrumData `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-spectrum-data"`
}

// CiscoIOSXEWirelessRRMOperRRMMeasurement represents the RRM measurement operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperRRMMeasurement struct {
	RRMMeasurement []RRMMeasurement `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-measurement"`
}

// CiscoIOSXEWirelessRRMOperRadioSlot represents the radio slot operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperRadioSlot struct {
	RadioSlot []RadioSlot `json:"Cisco-IOS-XE-wireless-rrm-oper:radio-slot"`
}

// CiscoIOSXEWirelessRRMOperMainData represents the main RRM operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperMainData struct {
	MainData []MainData `json:"Cisco-IOS-XE-wireless-rrm-oper:main-data"`
}

// CiscoIOSXEWirelessRRMOperRegDomainOper represents the regulatory domain operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperRegDomainOper struct {
	RegDomainOper *RegDomainOper `json:"Cisco-IOS-XE-wireless-rrm-oper:reg-domain-oper"`
}

// CiscoIOSXEWirelessRRMOperSpectrumDeviceTable represents the spectrum device table operational data (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperSpectrumDeviceTable struct {
	SpectrumDeviceTable []SpectrumDeviceTable `json:"Cisco-IOS-XE-wireless-rrm-oper:spectrum-device-table"`
}

// CiscoIOSXEWirelessRRMOperSpectrumAqTable represents the spectrum air quality table (YANG: IOS-XE 17.12.1).
type CiscoIOSXEWirelessRRMOperSpectrumAqTable struct {
	SpectrumAqTable []SpectrumAqTable `json:"Cisco-IOS-XE-wireless-rrm-oper:spectrum-aq-table"`
}

// ApAutoRFDot11Data represents AP auto RF 802.11 data.
type ApAutoRFDot11Data struct {
	WtpMAC            string             `json:"wtp-mac"`                       // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID       int                `json:"radio-slot-id"`                 // Radio slot identifier (Live: IOS-XE 17.12.6a)
	NeighborRadioInfo *NeighborRadioInfo `json:"neighbor-radio-info,omitempty"` // Neighbor radio information (Live: IOS-XE 17.12.6a)
}

// NeighborRadioInfo represents neighbor radio information.
type NeighborRadioInfo struct {
	NeighborRadioList []NeighborRadioListItem `json:"neighbor-radio-list,omitempty"` // List of neighboring radios (Live: IOS-XE 17.12.6a)
}

// NeighborRadioListItem represents a single neighbor radio entry.
type NeighborRadioListItem struct {
	NeighborRadioInfo NeighborRadioDetail `json:"neighbor-radio-info"` // Detailed neighbor radio information (Live: IOS-XE 17.12.6a)
}

// NeighborRadioDetail represents detailed neighbor radio information.
type NeighborRadioDetail struct {
	NeighborRadioMAC    string `json:"neighbor-radio-mac"`     // Neighbor radio MAC address (Live: IOS-XE 17.12.6a)
	NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"` // Neighbor radio slot identifier (Live: IOS-XE 17.12.6a)
	RSSI                int    `json:"rssi"`                   // Received Signal Strength Indicator in dBm (Live: IOS-XE 17.12.6a)
	SNR                 int    `json:"snr"`                    // Signal-to-Noise Ratio in dB (Live: IOS-XE 17.12.6a)
	Channel             int    `json:"channel"`                // Operating channel number (Live: IOS-XE 17.12.6a)
	Power               int    `json:"power"`                  // Transmit power level in dBm (Live: IOS-XE 17.12.6a)
	GroupLeaderIP       string `json:"group-leader-ip"`        // RRM group leader IP address (Live: IOS-XE 17.12.6a)
	ChanWidth           string `json:"chan-width"`             // Channel width setting (Live: IOS-XE 17.12.6a)
	SensorCovered       bool   `json:"sensor-covered"`         // Sensor coverage status (Live: IOS-XE 17.12.6a)
}

// ApDot11RadarData represents AP radar data.
type ApDot11RadarData struct {
	WtpMAC           string    `json:"wtp-mac"`             // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID      int       `json:"radio-slot-id"`       // Radio slot identifier (Live: IOS-XE 17.12.6a)
	LastRadarOnRadio time.Time `json:"last-radar-on-radio"` // Timestamp of last radar detection (Live: IOS-XE 17.12.6a)
}

// ApDot11SpectrumData represents AP spectrum data.
type ApDot11SpectrumData struct {
	WtpMAC      string          `json:"wtp-mac"`          // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID int             `json:"radio-slot-id"`    // Radio slot identifier (Live: IOS-XE 17.12.6a)
	Config      *SpectrumConfig `json:"config,omitempty"` // Spectrum analysis configuration (Live: IOS-XE 17.12.6a)
}

// SpectrumConfig represents spectrum configuration.
type SpectrumConfig struct {
	SpectrumIntelligenceEnable bool   `json:"spectrum-intelligence-enable"` // Enable spectrum intelligence feature (Live: IOS-XE 17.12.6a)
	SpectrumWtpCaSiCapable     string `json:"spectrum-wtp-ca-si-capable"`   // WTP spectrum intelligence capability (Live: IOS-XE 17.12.6a)
	SpectrumOperationState     string `json:"spectrum-operation-state"`     // Current spectrum operation state (Live: IOS-XE 17.12.6a)
	SpectrumAdminState         bool   `json:"spectrum-admin-state"`         // Administrative state of spectrum analysis (Live: IOS-XE 17.12.6a)
	SpectrumCapable            bool   `json:"spectrum-capable"`             // Spectrum analysis capability (Live: IOS-XE 17.12.6a)
	RapidUpdateEnable          bool   `json:"rapid-update-enable"`          // Enable rapid update mode (Live: IOS-XE 17.12.6a)
	SensordOperationalStatus   int    `json:"sensord-operational-status"`   // Sensor daemon operational status (Live: IOS-XE 17.12.6a)
	ScanRadioType              string `json:"scan-radio-type"`              // Radio type for spectrum scanning (Live: IOS-XE 17.12.6a)
}

// RRMMeasurement represents RRM measurement data.
type RRMMeasurement struct {
	WtpMAC      string   `json:"wtp-mac"`           // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID int      `json:"radio-slot-id"`     // Radio slot identifier (Live: IOS-XE 17.12.6a)
	Foreign     *Foreign `json:"foreign,omitempty"` // Foreign interference measurements (Live: IOS-XE 17.12.6a)
	Noise       *Noise   `json:"noise,omitempty"`   // Noise measurements (Live: IOS-XE 17.12.6a)
	Load        *Load    `json:"load,omitempty"`    // Load measurements (Live: IOS-XE 17.12.6a)
}

// Foreign represents foreign data measurements.
type Foreign struct {
	Foreign ForeignData `json:"foreign"` // Foreign interference data (Live: IOS-XE 17.12.6a)
}

// ForeignData represents foreign interference data.
type ForeignData struct {
	ForeignData []ForeignDataItem `json:"foreign-data,omitempty"` // List of foreign interference measurements (Live: IOS-XE 17.12.6a)
}

// ForeignDataItem represents a single foreign data measurement.
type ForeignDataItem struct {
	Chan                int `json:"chan"`                   // Channel number (Live: IOS-XE 17.12.6a)
	Power               int `json:"power"`                  // Signal power in dBm (Live: IOS-XE 17.12.6a)
	Rogue20Count        int `json:"rogue-20-count"`         // Count of 20MHz rogue signals (Live: IOS-XE 17.12.6a)
	Rogue40PrimaryCount int `json:"rogue-40-primary-count"` // Count of 40MHz primary rogue signals (Live: IOS-XE 17.12.6a)
	Rogue80PrimaryCount int `json:"rogue-80-primary-count"` // Count of 80MHz primary rogue signals (Live: IOS-XE 17.12.6a)
	ChanUtil            int `json:"chan-util"`              // Channel utilization percentage (Live: IOS-XE 17.12.6a)
}

// Noise represents noise measurements.
type Noise struct {
	Noise NoiseData `json:"noise"` // Noise measurement data (Live: IOS-XE 17.12.6a)
}

// NoiseData represents noise data collection.
type NoiseData struct {
	NoiseData []NoiseDataItem `json:"noise-data,omitempty"` // List of noise measurements by channel (Live: IOS-XE 17.12.6a)
}

// NoiseDataItem represents a single noise measurement.
type NoiseDataItem struct {
	Chan  int `json:"chan"`  // Channel number (Live: IOS-XE 17.12.6a)
	Noise int `json:"noise"` // Noise level in dBm (Live: IOS-XE 17.12.6a)
}

// Load represents load measurements.
type Load struct {
	RxUtilPercentage          int `json:"rx-util-percentage"`           // Receive utilization percentage (Live: IOS-XE 17.12.6a)
	TxUtilPercentage          int `json:"tx-util-percentage"`           // Transmit utilization percentage (Live: IOS-XE 17.12.6a)
	CcaUtilPercentage         int `json:"cca-util-percentage"`          // Clear Channel Assessment utilization percentage (Live: IOS-XE 17.12.6a)
	Stations                  int `json:"stations"`                     // Number of associated stations (Live: IOS-XE 17.12.6a)
	RxNoiseChannelUtilization int `json:"rx-noise-channel-utilization"` // Receive noise channel utilization (Live: IOS-XE 17.12.6a)
	NonWifiInter              int `json:"non-wifi-inter"`               // Non-WiFi interference level (Live: IOS-XE 17.12.6a)
}

// RadioSlot represents radio slot data.
type RadioSlot struct {
	WtpMAC      string     `json:"wtp-mac"`              // Access point MAC address (Live: IOS-XE 17.12.6a)
	RadioSlotID int        `json:"radio-slot-id"`        // Radio slot identifier (Live: IOS-XE 17.12.6a)
	RadioData   *RadioData `json:"radio-data,omitempty"` // Detailed radio operational data (Live: IOS-XE 17.12.6a)
}

// RadioData represents detailed radio operational data.
type RadioData struct {
	BestTxPwrLevel            int       `json:"best-tx-pwr-level"`           // Best transmit power level (Live: IOS-XE 17.12.6a)
	BestRtsThresh             int       `json:"best-rts-thresh"`             // Best RTS threshold (Live: IOS-XE 17.12.6a)
	BestFragThresh            int       `json:"best-frag-thresh"`            // Best fragmentation threshold (Live: IOS-XE 17.12.6a)
	LoadProfPassed            bool      `json:"load-prof-passed"`            // Load profile test result (Live: IOS-XE 17.12.6a)
	CoverageProfilePassed     bool      `json:"coverage-profile-passed"`     // Coverage profile test result (Live: IOS-XE 17.12.6a)
	InterferenceProfilePassed bool      `json:"interference-profile-passed"` // Interference profile test result (Live: IOS-XE 17.12.6a)
	NoiseProfilePassed        bool      `json:"noise-profile-passed"`        // Noise profile test result (Live: IOS-XE 17.12.6a)
	DCAStats                  *DCAStats `json:"dca-stats,omitempty"`         // Dynamic Channel Assignment statistics (Live: IOS-XE 17.12.6a)
	CoverageOverlapFactor     string    `json:"coverage-overlap-factor"`     // Coverage overlap factor (Live: IOS-XE 17.12.6a)
	SensorCoverageFactor      string    `json:"sensor-coverage-factor"`      // Sensor coverage factor (Live: IOS-XE 17.12.6a)
}

// DCAStats represents Dynamic Channel Assignment statistics.
type DCAStats struct {
	BestChan          int `json:"best-chan"`           // Best channel selection (Live: IOS-XE 17.12.6a)
	CurrentChanEnergy int `json:"current-chan-energy"` // Current channel energy level (Live: IOS-XE 17.12.6a)
	LastChanEnergy    int `json:"last-chan-energy"`    // Last channel energy level (Live: IOS-XE 17.12.6a)
	ChanChanges       int `json:"chan-changes"`        // Number of channel changes (Live: IOS-XE 17.12.6a)
}

// MainData represents main RRM data by PHY type.
type MainData struct {
	PhyType         string            `json:"phy-type"`                     // PHY type identifier (Live: IOS-XE 17.12.6a)
	Grp             *GroupData        `json:"grp,omitempty"`                // RRM group information (Live: IOS-XE 17.12.6a)
	OperData        *OperationalData  `json:"oper-data,omitempty"`          // Operational data (Live: IOS-XE 17.12.6a)
	RFName          string            `json:"rf-name"`                      // RF profile name (Live: IOS-XE 17.12.6a)
	RRMMgrGrpMember []RRMMgrGrpMember `json:"rrm-mgr-grp-member,omitempty"` // RRM manager group members (Live: IOS-XE 17.12.6a)
}

// GroupData represents RRM group information.
type GroupData struct {
	CurrentState          string       `json:"current-state"`            // Current RRM group state (Live: IOS-XE 17.12.6a)
	LastRun               time.Time    `json:"last-run"`                 // Last RRM algorithm run timestamp (Live: IOS-XE 17.12.6a)
	DCA                   *DCAInfo     `json:"dca,omitempty"`            // Dynamic Channel Assignment information (Live: IOS-XE 17.12.6a)
	Txpower               *TxPowerInfo `json:"txpower,omitempty"`        // Transmit power information (Live: IOS-XE 17.12.6a)
	CurrentGroupingMode   string       `json:"current-grouping-mode"`    // Current grouping mode (Live: IOS-XE 17.12.6a)
	JoinProtocolVer       int          `json:"join-protocol-ver"`        // Join protocol version (Live: IOS-XE 17.12.6a)
	CurrentGroupingRole   string       `json:"current-grouping-role"`    // Current grouping role (Live: IOS-XE 17.12.6a)
	CntrlrName            string       `json:"cntrlr-name"`              // Controller name (Live: IOS-XE 17.12.6a)
	CntrlrIPAddr          string       `json:"cntrlr-ip-addr"`           // Controller IP address (Live: IOS-XE 17.12.6a)
	CntrlrSecondaryIPAddr string       `json:"cntrlr-secondary-ip-addr"` // Controller secondary IP address (Live: IOS-XE 17.12.6a)
	IsStaticMember        string       `json:"is-static-member"`         // Static member status (Live: IOS-XE 17.12.6a)
	DpcConfig             *DpcConfig   `json:"dpc-config,omitempty"`     // Dynamic Power Control configuration (Live: IOS-XE 17.12.6a)
	FraSensorCoverage     int          `json:"fra-sensor-coverage"`      // FRA sensor coverage (Live: IOS-XE 17.12.6a)
	ProtocolVer           int          `json:"protocol-ver"`             // Protocol version (Live: IOS-XE 17.12.6a)
	HdrVer                int          `json:"hdr-ver"`                  // Header version (Live: IOS-XE 17.12.6a)
}

// DCAInfo represents DCA information.
type DCAInfo struct {
	DCALastRun time.Time `json:"dca-last-run"` // Last DCA algorithm run timestamp (Live: IOS-XE 17.12.6a)
}

// TxPowerInfo represents transmit power information.
type TxPowerInfo struct {
	DpcLastRun time.Time `json:"dpc-last-run"` // Last DPC algorithm run timestamp (Live: IOS-XE 17.12.6a)
	RunTime    int       `json:"run-time"`     // Algorithm run time in seconds (Live: IOS-XE 17.12.6a)
}

// DpcConfig represents Dynamic Power Control configuration.
type DpcConfig struct {
	RF                      *RFConfig `json:"rf,omitempty"`               // RF configuration (Live: IOS-XE 17.12.6a)
	DpcMinTxPowerLimit      int       `json:"dpc-min-tx-power-limit"`     // Minimum transmit power limit in dBm (Live: IOS-XE 17.12.6a)
	DpcMaxTxPowerLimit      int       `json:"dpc-max-tx-power-limit"`     // Maximum transmit power limit in dBm (Live: IOS-XE 17.12.6a)
	TxPowerControlThreshold int       `json:"tx-power-control-threshold"` // Transmit power control threshold (Live: IOS-XE 17.12.6a)
}

// RFConfig represents RF configuration.
type RFConfig struct {
	Mode              string `json:"mode"`                // RF mode setting (Live: IOS-XE 17.12.6a)
	UpdateCounter     int    `json:"update-counter"`      // Update counter (Live: IOS-XE 17.12.6a)
	UpdateIntervalSec int    `json:"update-interval-sec"` // Update interval in seconds (Live: IOS-XE 17.12.6a)
	Contribution      int    `json:"contribution"`        // Contribution value (Live: IOS-XE 17.12.6a)
}

// OperationalData represents operational data.
type OperationalData struct {
	DCAThreshVal          int          `json:"dca-thresh-val"`                     // DCA threshold value (Live: IOS-XE 17.12.6a)
	DefaultDCAChannels    *ChannelList `json:"default-dca-channels,omitempty"`     // Default DCA channels (Live: IOS-XE 17.12.6a)
	DefaultNonDCAChannels *ChannelList `json:"default-non-dca-channels,omitempty"` // Default non-DCA channels (Live: IOS-XE 17.12.6a)
	FraOperState          bool         `json:"fra-oper-state"`                     // FRA operational state (Live: IOS-XE 17.12.6a)
}

// ChannelList represents a list of channels.
type ChannelList struct {
	Channel []int `json:"channel,omitempty"` // List of channel numbers (Live: IOS-XE 17.12.6a)
}

// RRMMgrGrpMember represents RRM manager group member.
type RRMMgrGrpMember struct {
	MemberIP       string `json:"member-ip"`        // Member IP address (Live: IOS-XE 17.12.6a)
	MaxRadioCnt    int    `json:"max-radio-cnt"`    // Maximum radio count (Live: IOS-XE 17.12.6a)
	CurrRadioCnt   int    `json:"curr-radio-cnt"`   // Current radio count (Live: IOS-XE 17.12.6a)
	Name           string `json:"name"`             // Member name (Live: IOS-XE 17.12.6a)
	DTLSConnStatus string `json:"dtls-conn-status"` // DTLS connection status (Live: IOS-XE 17.12.6a)
}

// RegDomainOper represents regulatory domain operational data.
type RegDomainOper struct {
	CountryList string `json:"country-list"` // Supported country list (Live: IOS-XE 17.12.6a)
}

// SpectrumDeviceTable represents spectrum device detection table entry.
type SpectrumDeviceTable struct {
	DeviceID        string    `json:"device-id"`          // Device identifier (Live: IOS-XE 17.12.6a)
	ClusterID       string    `json:"cluster-id"`         // Cluster identifier (Live: IOS-XE 17.12.6a)
	LastUpdatedTime time.Time `json:"last-updated-time"`  // Last update timestamp (Live: IOS-XE 17.12.6a)
	IDRData         *IDRData  `json:"idr-data,omitempty"` // Interference device recognition data (Live: IOS-XE 17.12.6a)
}

// IDRData represents interference device recognition data.
type IDRData struct {
	DetectingApMAC      string `json:"detecting-ap-mac"`      // Detecting AP MAC address (Live: IOS-XE 17.12.6a)
	AffectedChannelList string `json:"affected-channel-list"` // Affected channel list (Live: IOS-XE 17.12.6a)
	IsPersistent        bool   `json:"is-persistent"`         // Persistent interference indicator (Live: IOS-XE 17.12.6a)
	ClassTypeEnum       string `json:"class-type-enum"`       // Device class type enumeration (Live: IOS-XE 17.12.6a)
}

// SpectrumAqTable represents spectrum air quality table entry.
type SpectrumAqTable struct {
	WtpMAC          string          `json:"wtp-mac"`                     // Access point MAC address (Live: IOS-XE 17.12.6a)
	Band            string          `json:"band"`                        // Radio band identifier (Live: IOS-XE 17.12.6a)
	ReportingApName string          `json:"reporting-ap-name"`           // Reporting AP name (Live: IOS-XE 17.12.6a)
	PerRadioAqData  *PerRadioAqData `json:"per-radio-aq-data,omitempty"` // Per radio air quality data (Live: IOS-XE 17.12.6a)
	WtpCaSiCapable  string          `json:"wtp-ca-si-capable"`           // WTP CleanAir/SI capability (Live: IOS-XE 17.12.6a)
	ScanRadioType   string          `json:"scan-radio-type"`             // Scan radio type (Live: IOS-XE 17.12.6a)
}

// PerRadioAqData represents per radio air quality data.
type PerRadioAqData struct {
	ChannelCount     int                `json:"channel-count"`                 // Number of channels (Live: IOS-XE 17.12.6a)
	PerChannelAqList []PerChannelAqList `json:"per-channel-aq-list,omitempty"` // Per channel air quality list (Live: IOS-XE 17.12.6a)
}

// PerChannelAqList represents per channel air quality list entry.
type PerChannelAqList struct {
	ChannelNum           int       `json:"channel-num"`             // Channel number (Live: IOS-XE 17.12.6a)
	MinAqi               int       `json:"min-aqi"`                 // Minimum air quality index (Live: IOS-XE 17.12.6a)
	Aqi                  int       `json:"aqi"`                     // Air quality index (Live: IOS-XE 17.12.6a)
	TotalIntfDeviceCount int       `json:"total-intf-device-count"` // Total interference device count (Live: IOS-XE 17.12.6a)
	SpectrumTimestamp    time.Time `json:"spectrum-timestamp"`      // Spectrum analysis timestamp (Live: IOS-XE 17.12.6a)
}
