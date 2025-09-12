// Package site provides data models for site configuration data.
package site

// SiteCfg represents site configuration data container.
type SiteCfg struct {
	SiteCfgData SiteCfgData `json:"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"`
}

// SiteCfgSiteListEntries represents site list entries container wrapper.
type SiteCfgSiteListEntries struct {
	SiteListEntries SiteListEntries `json:"Cisco-IOS-XE-wireless-site-cfg:site-list-entries"`
}

// SiteCfgApCfgProfiles represents AP configuration profiles container.
type SiteCfgApCfgProfiles struct {
	ApCfgProfiles ApCfgProfiles `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles"`
}

// SiteCfgSiteListEntrys represents site tag configurations container.
type SiteCfgSiteListEntrys struct {
	SiteListEntrys SiteListEntrys `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs"`
}

// SiteCfgApCfgProfile represents individual AP configuration profile wrapper.
type SiteCfgApCfgProfile struct {
	ApCfgProfile []ApCfgProfile `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profile"`
}

// SiteCfgSiteListEntry represents individual site tag configuration wrapper.
type SiteCfgSiteListEntry struct {
	SiteListEntry []SiteListEntry `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"`
}

// SiteListEntrysPayload represents request structure for site-tag-configs endpoint.
type SiteListEntrysPayload struct {
	SiteListEntry SiteListEntry `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"`
}

// SiteCfgData represents site configuration data container.
type SiteCfgData struct {
	ApCfgProfiles  ApCfgProfiles  `json:"ap-cfg-profiles"`
	SiteListEntrys SiteListEntrys `json:"site-tag-configs"`
}

// SiteListEntrys represents site tag configurations container.
type SiteListEntrys struct {
	SiteListEntry []SiteListEntry `json:"site-tag-config"`
}

// ApCfgProfiles represents access point configuration profiles container.
type ApCfgProfiles struct {
	ApCfgProfile []ApCfgProfile `json:"ap-cfg-profile"`
}

// ApCfgProfile represents individual access point configuration profile.
type ApCfgProfile struct {
	ProfileName        string            `json:"profile-name"`                  // AP configuration profile name
	Description        *string           `json:"description,omitempty"`         // Profile description
	BleBeaconInterval  *int              `json:"ble-beacon-interval,omitempty"` // BLE beacon interval (YANG: IOS-XE 17.12.1+)
	BleBeaconAdvpwr    *int              `json:"ble-beacon-advpwr,omitempty"`   // BLE beacon advertising power (YANG: IOS-XE 17.12.1+)
	DataEncryptionFlag bool              `json:"data-encryption-flag"`          // Data encryption enable flag
	StatsTimer         StatsTimer        `json:"stats-timer"`                   // Statistics timer configuration
	LedState           *LedState         `json:"led-state,omitempty"`           // LED state configuration (YANG: IOS-XE 17.12.1+)
	LinkLatency        *LinkLatency      `json:"link-latency,omitempty"`        // Link latency configuration (YANG: IOS-XE 17.12.1+)
	JumboMtu           JumboMtu          `json:"jumbo-mtu"`                     // Jumbo MTU configuration
	ApMode             *ApMode           `json:"ap-mode,omitempty"`             // AP mode configuration (YANG: IOS-XE 17.12.1+)
	Poe                *Poe              `json:"poe,omitempty"`                 // Power over Ethernet configuration (YANG: IOS-XE 17.12.1+)
	DeviceMgmt         DeviceMgmt        `json:"device-mgmt"`                   // Device management configuration
	UserMgmt           UserMgmt          `json:"user-mgmt"`                     // User management configuration
	Tunnel             Tunnel            `json:"tunnel"`                        // Tunnel configuration
	CapwapTimer        CapwapTimer       `json:"capwap-timer"`                  // CAPWAP timer configuration
	Syslog             Syslog            `json:"syslog"`                        // Syslog configuration
	Hyperlocation      Hyperlocation     `json:"hyperlocation"`                 // Hyperlocation configuration
	RogueDetection     RogueDetection    `json:"rogue-detection"`               // Rogue detection configuration
	TftpDownGrade      TftpDownGrade     `json:"tftp-down-grade"`               // TFTP downgrade configuration
	ReportingInterval  ReportingInterval `json:"reporting-interval"`            // Reporting interval configuration
	GasRateLimit       GasRateLimit      `json:"gas-rate-limit"`                // GAS rate limit configuration
	NtpServerInfo      NtpServerInfo     `json:"ntp-server-info"`               // NTP server configuration
	PublicIPDiscovery  bool              `json:"public-ip-discovery"`           // Public IP discovery enable flag
	Oeap               Oeap              `json:"oeap"`                          // OEAP configuration
	ApTzConfig         ApTzConfig        `json:"ap-tz-config"`                  // AP timezone configuration
	RadioStatsMonitor  RadioStatsMonitor `json:"radio-stats-monitor"`           // Radio statistics monitoring configuration
	ApProfPpCfg        ApProfPpCfg       `json:"ap-prof-pp-cfg"`                // AP profile power configuration
}

// SiteListEntries represents site list entries container.
type SiteListEntries struct {
	SiteListEntry []SiteListEntry `json:"site-list-entry"`
}

// SiteListEntry represents individual site list entry.
type SiteListEntry struct {
	SiteTagName              string  `json:"site-tag-name"`                         // Site tag name identifier
	Description              *string `json:"description,omitempty"`                 // Site description
	FlexProfile              *string `json:"flex-profile,omitempty"`                // FlexConnect profile name
	ApJoinProfile            *string `json:"ap-join-profile,omitempty"`             // AP join profile name
	IsLocalSite              *bool   `json:"is-local-site,omitempty"`               // Local site flag
	FabricControlPlaneName   *string `json:"fabric-control-plane-name,omitempty"`   // Fabric control plane name (YANG: IOS-XE 17.12.1+)
	ImageDownloadProfileName *string `json:"image-download-profile-name,omitempty"` // Image download profile name (YANG: IOS-XE 17.12.1+)
	ArpCaching               *bool   `json:"arp-caching,omitempty"`                 // ARP caching enable flag (YANG: IOS-XE 17.12.1+)
	DHCPBcast                *bool   `json:"dhcp-bcast,omitempty"`                  // DHCP broadcast enable flag (YANG: IOS-XE 17.12.1+)
	FabricMcastIPv4Addr      *string `json:"fabric-mcast-ipv4-addr,omitempty"`      // Fabric multicast IPv4 address (YANG: IOS-XE 17.12.1+)
	Load                     *int    `json:"load,omitempty"`                        // Site load value (YANG: IOS-XE 17.12.1+)
}

// StatsTimer represents statistics timer configuration.
type StatsTimer struct {
	StatsTimer int `json:"stats-timer"` // Statistics timer interval in seconds
}

// JumboMtu represents jumbo MTU configuration.
type JumboMtu struct {
	JumboMtu bool `json:"jumbo-mtu"` // Jumbo MTU enable flag
}

// DeviceMgmt represents device management configuration.
type DeviceMgmt struct {
	SSH bool `json:"ssh"` // SSH access enable flag
}

// UserMgmt represents user management configuration.
type UserMgmt struct {
	Username     string `json:"username"`      // Admin username
	Password     string `json:"password"`      // Admin password
	PasswordType string `json:"password-type"` // Password encryption type
	Secret       string `json:"secret"`        // Enable secret
	SecretType   string `json:"secret-type"`   // Secret encryption type
}

// Tunnel represents tunnel configuration.
type Tunnel struct {
	PreferredMode string `json:"preferred-mode"` // Tunnel preferred mode
}

// CapwapTimer represents CAPWAP timer configuration.
type CapwapTimer struct {
	FastHeartBeatTimeout int `json:"fast-heart-beat-timeout"` // Fast heartbeat timeout in seconds
}

// Syslog represents syslog configuration.
type Syslog struct {
	FacilityValue string `json:"facility-value"` // Syslog facility value
	LogLevel      string `json:"log-level"`      // Syslog level
	Host          string `json:"host"`           // Syslog server host
	TLSMode       bool   `json:"tls-mode"`       // TLS mode enable flag
}

// Hyperlocation represents hyperlocation configuration.
type Hyperlocation struct {
	HyperlocationEnable bool `json:"hyperlocation-enable"` // Hyperlocation enable flag
}

// RogueDetection represents rogue detection configuration.
type RogueDetection struct {
	RogueDetectionMonitorModeReportInterval int `json:"rogue-detection-monitor-mode-report-interval"` // Rogue detection report interval in seconds
	ApRogueDetectionMinRssi                 int `json:"ap-rogue-detection-min-rssi"`                  // Minimum RSSI threshold for rogue detection
	ApRogueDetectionTransientInterval       int `json:"ap-rogue-detection-transient-interval"`        // Transient interval for rogue detection in seconds
}

// TftpDownGrade represents TFTP downgrade configuration.
type TftpDownGrade struct {
	TftpDowngradeIPAddress string `json:"tftp-downgrade-ip-address"` // TFTP server IP address for downgrade
}

// ReportingInterval represents reporting interval configuration.
type ReportingInterval struct {
	Radio24GHz int `json:"radio-24ghz"` // 2.4GHz radio reporting interval in seconds
	Radio5GHz  int `json:"radio-5ghz"`  // 5GHz radio reporting interval in seconds
}

// GasRateLimit represents GAS rate limit configuration.
type GasRateLimit struct {
	NumReqPerInterval int `json:"num-req-per-interval"` // Number of requests per interval
	IntervalMsec      int `json:"interval-msec"`        // Interval in milliseconds
}

// NtpServerInfo represents NTP server configuration.
type NtpServerInfo struct {
	NtpAddress string `json:"ntp-address"` // NTP server IP address
}

// Oeap represents OEAP configuration.
type Oeap struct {
	OeapDataEncr bool `json:"oeap-data-encr"` // OEAP data encryption enable flag
	IsLocalNet   bool `json:"is-local-net"`   // Local network flag
	ProvSsid     bool `json:"prov-ssid"`      // Provisioning SSID enable flag
}

// ApTzConfig represents access point timezone configuration.
type ApTzConfig struct {
	TzEnabled bool   `json:"tz-enabled"` // Timezone enable flag
	Mode      string `json:"mode"`       // Timezone mode
}

// RadioStatsMonitor represents radio statistics monitoring configuration.
type RadioStatsMonitor struct {
	Enable       bool    `json:"enable"`        // Radio statistics monitoring enable flag
	AlarmsEnable []*bool `json:"alarms-enable"` // Alarms enable flags array
	RadioReset   bool    `json:"radio-reset"`   // Radio reset enable flag
}

// ApProfPpCfg represents access point profile power configuration.
type ApProfPpCfg struct {
	PowerProfileName string `json:"power-profile-name"` // Power profile name
}

// LedState represents AP LED configuration based on YANG wireless-ap-types:st-led-state.
type LedState struct {
	LedState bool `json:"led-state"` // LED state enable flag (YANG: IOS-XE 17.12.1+)
}

// LinkLatency represents AP link auditing configuration based on YANG wireless-ap-types:st-link-latency.
type LinkLatency struct {
	LinkLatencyFlag string `json:"link-latency-flag"` // Link latency flag configuration (YANG: IOS-XE 17.12.1+)
}

// ApMode represents AP mode configuration based on YANG wireless-ap-types:st-mode.
type ApMode struct {
	ApSubMode   string `json:"ap-sub-mode"`            // AP sub-mode type (YANG: IOS-XE 17.12.1+)
	FastChannel *int   `json:"fast-channel,omitempty"` // Fast channel configuration (YANG: IOS-XE 17.12.1+)
}

// Poe represents Power over Ethernet configuration based on YANG wireless-ap-types:st-poe-cfg.
type Poe struct {
	PreStandard8023afSwitchFlag bool   `json:"pre-standard8023af-switch-flag"` // Pre-standard 802.3af switch flag (YANG: IOS-XE 17.12.1+)
	PowerInjectorState          bool   `json:"power-injector-state"`           // Power injector state flag (YANG: IOS-XE 17.12.1+)
	PowerInjectorSelection      string `json:"power-injector-selection"`       // Power injector selection type (YANG: IOS-XE 17.12.1+)
}
