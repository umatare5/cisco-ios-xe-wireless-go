package rrm

import "time"

// RRMOper represents RRM operational response data.
type RRMOper struct {
	CiscoIOSXEWirelessRRMOperRRMOperData RRMOperData `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"` // RRM operational data container
}

// RRMOperData represents the main RRM operational data container.
type RRMOperData struct {
	ApAutoRFDot11Data   []ApAutoRFDot11Data   `json:"ap-auto-rf-dot11-data,omitempty"`  // AP auto RF 802.11 data
	ApDot11RadarData    []ApDot11RadarData    `json:"ap-dot11-radar-data,omitempty"`    // AP radar detection data
	ApDot11SpectrumData []ApDot11SpectrumData `json:"ap-dot11-spectrum-data,omitempty"` // AP spectrum analysis data
	RRMMeasurement      []RRMMeasurement      `json:"rrm-measurement,omitempty"`        // RRM measurement data
	RadioSlot           []RadioSlot           `json:"radio-slot,omitempty"`             // Radio slot operational data
	MainData            []MainData            `json:"main-data,omitempty"`              // Main RRM data by PHY type
	RegDomainOper       *RegDomainOper        `json:"reg-domain-oper,omitempty"`        // Regulatory domain operational data
}

// ApAutoRFDot11Data represents AP auto RF 802.11 data.
type ApAutoRFDot11Data struct {
	WtpMAC            string             `json:"wtp-mac"`                       // Access point MAC address
	RadioSlotID       int                `json:"radio-slot-id"`                 // Radio slot identifier
	NeighborRadioInfo *NeighborRadioInfo `json:"neighbor-radio-info,omitempty"` // Neighbor radio information
}

// NeighborRadioInfo represents neighbor radio information.
type NeighborRadioInfo struct {
	NeighborRadioList []NeighborRadioListItem `json:"neighbor-radio-list,omitempty"` // List of neighboring radios
}

// NeighborRadioListItem represents a single neighbor radio entry.
type NeighborRadioListItem struct {
	NeighborRadioInfo NeighborRadioDetail `json:"neighbor-radio-info"` // Detailed neighbor radio information
}

// NeighborRadioDetail represents detailed neighbor radio information.
type NeighborRadioDetail struct {
	NeighborRadioMAC    string `json:"neighbor-radio-mac"`     // Neighbor radio MAC address
	NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"` // Neighbor radio slot identifier
	RSSI                int    `json:"rssi"`                   // Received Signal Strength Indicator in dBm
	SNR                 int    `json:"snr"`                    // Signal-to-Noise Ratio in dB
	Channel             int    `json:"channel"`                // Operating channel number
	Power               int    `json:"power"`                  // Transmit power level in dBm
	GroupLeaderIP       string `json:"group-leader-ip"`        // RRM group leader IP address
	ChanWidth           string `json:"chan-width"`             // Channel width setting
	SensorCovered       bool   `json:"sensor-covered"`         // Sensor coverage status
}

// ApDot11RadarData represents AP radar data.
type ApDot11RadarData struct {
	WtpMAC           string    `json:"wtp-mac"`             // Access point MAC address
	RadioSlotID      int       `json:"radio-slot-id"`       // Radio slot identifier
	LastRadarOnRadio time.Time `json:"last-radar-on-radio"` // Timestamp of last radar detection
}

// ApDot11SpectrumData represents AP spectrum data.
type ApDot11SpectrumData struct {
	WtpMAC      string          `json:"wtp-mac"`          // Access point MAC address
	RadioSlotID int             `json:"radio-slot-id"`    // Radio slot identifier
	Config      *SpectrumConfig `json:"config,omitempty"` // Spectrum analysis configuration
}

// SpectrumConfig represents spectrum configuration.
type SpectrumConfig struct {
	SpectrumIntelligenceEnable bool   `json:"spectrum-intelligence-enable"` // Enable spectrum intelligence feature
	SpectrumWtpCaSiCapable     string `json:"spectrum-wtp-ca-si-capable"`   // WTP spectrum intelligence capability
	SpectrumOperationState     string `json:"spectrum-operation-state"`     // Current spectrum operation state
	SpectrumAdminState         bool   `json:"spectrum-admin-state"`         // Administrative state of spectrum analysis
	SpectrumCapable            bool   `json:"spectrum-capable"`             // Spectrum analysis capability
	RapidUpdateEnable          bool   `json:"rapid-update-enable"`          // Enable rapid update mode
	SensordOperationalStatus   int    `json:"sensord-operational-status"`   // Sensor daemon operational status
	ScanRadioType              string `json:"scan-radio-type"`              // Radio type for spectrum scanning
}

// RRMMeasurement represents RRM measurement data.
type RRMMeasurement struct {
	WtpMAC      string   `json:"wtp-mac"`           // Access point MAC address
	RadioSlotID int      `json:"radio-slot-id"`     // Radio slot identifier
	Foreign     *Foreign `json:"foreign,omitempty"` // Foreign interference measurements
	Noise       *Noise   `json:"noise,omitempty"`   // Noise measurements
	Load        *Load    `json:"load,omitempty"`    // Load measurements
}

// Foreign represents foreign data measurements.
type Foreign struct {
	Foreign ForeignData `json:"foreign"` // Foreign interference data
}

// ForeignData represents foreign interference data.
type ForeignData struct {
	ForeignData []ForeignDataItem `json:"foreign-data,omitempty"` // List of foreign interference measurements
}

// ForeignDataItem represents a single foreign data measurement.
type ForeignDataItem struct {
	Chan                int `json:"chan"`                   // Channel number
	Power               int `json:"power"`                  // Signal power in dBm
	Rogue20Count        int `json:"rogue-20-count"`         // Count of 20MHz rogue signals
	Rogue40PrimaryCount int `json:"rogue-40-primary-count"` // Count of 40MHz primary rogue signals
	Rogue80PrimaryCount int `json:"rogue-80-primary-count"` // Count of 80MHz primary rogue signals
	ChanUtil            int `json:"chan-util"`              // Channel utilization percentage
}

// Noise represents noise measurements.
type Noise struct {
	Noise NoiseData `json:"noise"` // Noise measurement data
}

// NoiseData represents noise data collection.
type NoiseData struct {
	NoiseData []NoiseDataItem `json:"noise-data,omitempty"` // List of noise measurements by channel
}

// NoiseDataItem represents a single noise measurement.
type NoiseDataItem struct {
	Chan  int `json:"chan"`  // Channel number
	Noise int `json:"noise"` // Noise level in dBm
}

// Load represents load measurements.
type Load struct {
	RxUtilPercentage          int `json:"rx-util-percentage"`           // Receive utilization percentage
	TxUtilPercentage          int `json:"tx-util-percentage"`           // Transmit utilization percentage
	CcaUtilPercentage         int `json:"cca-util-percentage"`          // Clear Channel Assessment utilization percentage
	Stations                  int `json:"stations"`                     // Number of associated stations
	RxNoiseChannelUtilization int `json:"rx-noise-channel-utilization"` // Receive noise channel utilization
	NonWifiInter              int `json:"non-wifi-inter"`               // Non-WiFi interference level
}

// RadioSlot represents radio slot data.
type RadioSlot struct {
	WtpMAC      string     `json:"wtp-mac"`              // Access point MAC address
	RadioSlotID int        `json:"radio-slot-id"`        // Radio slot identifier
	RadioData   *RadioData `json:"radio-data,omitempty"` // Detailed radio operational data
}

// RadioData represents detailed radio operational data.
type RadioData struct {
	BestTxPwrLevel            int       `json:"best-tx-pwr-level"`           // Best transmit power level
	BestRtsThresh             int       `json:"best-rts-thresh"`             // Best RTS threshold
	BestFragThresh            int       `json:"best-frag-thresh"`            // Best fragmentation threshold
	LoadProfPassed            bool      `json:"load-prof-passed"`            // Load profile test result
	CoverageProfilePassed     bool      `json:"coverage-profile-passed"`     // Coverage profile test result
	InterferenceProfilePassed bool      `json:"interference-profile-passed"` // Interference profile test result
	NoiseProfilePassed        bool      `json:"noise-profile-passed"`        // Noise profile test result
	DCAStats                  *DCAStats `json:"dca-stats,omitempty"`         // Dynamic Channel Assignment statistics
	CoverageOverlapFactor     string    `json:"coverage-overlap-factor"`     // Coverage overlap factor
	SensorCoverageFactor      string    `json:"sensor-coverage-factor"`      // Sensor coverage factor
}

// DCAStats represents Dynamic Channel Assignment statistics.
type DCAStats struct {
	BestChan          int `json:"best-chan"`           // Best channel selection
	CurrentChanEnergy int `json:"current-chan-energy"` // Current channel energy level
	LastChanEnergy    int `json:"last-chan-energy"`    // Last channel energy level
	ChanChanges       int `json:"chan-changes"`        // Number of channel changes
}

// MainData represents main RRM data by PHY type.
type MainData struct {
	PhyType         string            `json:"phy-type"`                     // PHY type identifier
	Grp             *GroupData        `json:"grp,omitempty"`                // RRM group information
	OperData        *OperationalData  `json:"oper-data,omitempty"`          // Operational data
	RFName          string            `json:"rf-name"`                      // RF profile name
	RRMMgrGrpMember []RRMMgrGrpMember `json:"rrm-mgr-grp-member,omitempty"` // RRM manager group members
}

// GroupData represents RRM group information.
type GroupData struct {
	CurrentState          string       `json:"current-state"`            // Current RRM group state
	LastRun               time.Time    `json:"last-run"`                 // Last RRM algorithm run timestamp
	DCA                   *DCAInfo     `json:"dca,omitempty"`            // Dynamic Channel Assignment information
	Txpower               *TxPowerInfo `json:"txpower,omitempty"`        // Transmit power information
	CurrentGroupingMode   string       `json:"current-grouping-mode"`    // Current grouping mode
	JoinProtocolVer       int          `json:"join-protocol-ver"`        // Join protocol version
	CurrentGroupingRole   string       `json:"current-grouping-role"`    // Current grouping role
	CntrlrName            string       `json:"cntrlr-name"`              // Controller name
	CntrlrIPAddr          string       `json:"cntrlr-ip-addr"`           // Controller IP address
	CntrlrSecondaryIPAddr string       `json:"cntrlr-secondary-ip-addr"` // Controller secondary IP address
	IsStaticMember        string       `json:"is-static-member"`         // Static member status
	DpcConfig             *DpcConfig   `json:"dpc-config,omitempty"`     // Dynamic Power Control configuration
	FraSensorCoverage     int          `json:"fra-sensor-coverage"`      // FRA sensor coverage
	ProtocolVer           int          `json:"protocol-ver"`             // Protocol version
	HdrVer                int          `json:"hdr-ver"`                  // Header version
}

// DCAInfo represents DCA information.
type DCAInfo struct {
	DCALastRun time.Time `json:"dca-last-run"` // Last DCA algorithm run timestamp
}

// TxPowerInfo represents transmit power information.
type TxPowerInfo struct {
	DpcLastRun time.Time `json:"dpc-last-run"` // Last DPC algorithm run timestamp
	RunTime    int       `json:"run-time"`     // Algorithm run time in seconds
}

// DpcConfig represents Dynamic Power Control configuration.
type DpcConfig struct {
	RF                      *RFConfig `json:"rf,omitempty"`               // RF configuration
	DpcMinTxPowerLimit      int       `json:"dpc-min-tx-power-limit"`     // Minimum transmit power limit in dBm
	DpcMaxTxPowerLimit      int       `json:"dpc-max-tx-power-limit"`     // Maximum transmit power limit in dBm
	TxPowerControlThreshold int       `json:"tx-power-control-threshold"` // Transmit power control threshold
}

// RFConfig represents RF configuration.
type RFConfig struct {
	Mode              string `json:"mode"`                // RF mode setting
	UpdateCounter     int    `json:"update-counter"`      // Update counter
	UpdateIntervalSec int    `json:"update-interval-sec"` // Update interval in seconds
	Contribution      int    `json:"contribution"`        // Contribution value
}

// OperationalData represents operational data.
type OperationalData struct {
	DCAThreshVal          int          `json:"dca-thresh-val"`                     // DCA threshold value
	DefaultDCAChannels    *ChannelList `json:"default-dca-channels,omitempty"`     // Default DCA channels
	DefaultNonDCAChannels *ChannelList `json:"default-non-dca-channels,omitempty"` // Default non-DCA channels
	FraOperState          bool         `json:"fra-oper-state"`                     // FRA operational state
}

// ChannelList represents a list of channels.
type ChannelList struct {
	Channel []int `json:"channel,omitempty"` // List of channel numbers
}

// RRMMgrGrpMember represents RRM manager group member.
type RRMMgrGrpMember struct {
	MemberIP       string `json:"member-ip"`        // Member IP address
	MaxRadioCnt    int    `json:"max-radio-cnt"`    // Maximum radio count
	CurrRadioCnt   int    `json:"curr-radio-cnt"`   // Current radio count
	Name           string `json:"name"`             // Member name
	DTLSConnStatus string `json:"dtls-conn-status"` // DTLS connection status
}

// RegDomainOper represents regulatory domain operational data.
type RegDomainOper struct {
	CountryList string `json:"country-list"` // Supported country list
}
