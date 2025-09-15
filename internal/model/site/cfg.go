package site

// SiteCfg represents site configuration data container (Live: IOS-XE 17.12.5).
type SiteCfg struct {
	SiteCfgData struct {
		ApCfgProfiles  ApCfgProfiles  `json:"ap-cfg-profiles"`  // AP config profiles container (Live: IOS-XE 17.12.5)
		SiteTagConfigs SiteTagConfigs `json:"site-tag-configs"` // Site tag configs container (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"`
}

// SiteCfgApCfgProfiles represents AP config profiles container (Live: IOS-XE 17.12.5).
type SiteCfgApCfgProfiles struct {
	ApCfgProfiles ApCfgProfiles `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles"` // AP config profiles list (Live: IOS-XE 17.12.5)
}

// SiteCfgSiteTagConfigs represents site tag configs container (Live: IOS-XE 17.12.5).
type SiteCfgSiteTagConfigs struct {
	SiteTagConfigs SiteTagConfigs `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs"` // Site tag configs list (Live: IOS-XE 17.12.5)
}

// SiteCfgSiteTagConfig represents site tag config wrapper (Live: IOS-XE 17.12.5).
type SiteCfgSiteTagConfig struct {
	SiteListEntry []SiteListEntry `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"` // Site tag config entries (Live: IOS-XE 17.12.5)
}

// ApCfgProfiles represents AP config profiles container (Live: IOS-XE 17.12.5).
type ApCfgProfiles struct {
	ApCfgProfile []ApCfgProfile `json:"ap-cfg-profile"` // List of AP config profiles (Live: IOS-XE 17.12.5)
}

// SiteTagConfigs represents site tag configs container (Live: IOS-XE 17.12.5).
type SiteTagConfigs struct {
	SiteTagConfig []SiteListEntry `json:"site-tag-config"` // List of site tag configs (Live: IOS-XE 17.12.5)
}

// SiteListEntries represents site list entries container (Live: IOS-XE 17.12.5).
type SiteListEntries struct {
	SiteListEntry []SiteListEntry `json:"site-list-entry"` // List of site entries (Live: IOS-XE 17.12.5)
}

// ApCfgProfile represents AP config profile (Live: IOS-XE 17.12.5).
type ApCfgProfile struct {
	ProfileName        string            `json:"profile-name"`                  // AP config profile name (Live: IOS-XE 17.12.5)
	Description        *string           `json:"description,omitempty"`         // Profile description (Live: IOS-XE 17.12.5)
	BLEBeaconInterval  *int              `json:"ble-beacon-interval,omitempty"` // BLE beacon interval (YANG: IOS-XE 17.12.1)
	BLEBeaconAdvpwr    *int              `json:"ble-beacon-advpwr,omitempty"`   // BLE beacon advertising power (YANG: IOS-XE 17.12.1)
	DataEncryptionFlag bool              `json:"data-encryption-flag"`          // Data encryption status of AP (Live: IOS-XE 17.12.5)
	StatsTimer         StatsTimer        `json:"stats-timer"`                   // Stats timer for the AP (Live: IOS-XE 17.12.5)
	LedState           *LedState         `json:"led-state,omitempty"`           // LED state configuration (YANG: IOS-XE 17.12.1)
	LinkLatency        *LinkLatency      `json:"link-latency,omitempty"`        // Link auditing options (YANG: IOS-XE 17.12.1)
	JumboMtu           JumboMtu          `json:"jumbo-mtu"`                     // Jumbo MTU enable flag (Live: IOS-XE 17.12.5)
	ApMode             *ApMode           `json:"ap-mode,omitempty"`             // Mode of operation of the AP (YANG: IOS-XE 17.12.1)
	Poe                *Poe              `json:"poe,omitempty"`                 // Power over ethernet config (YANG: IOS-XE 17.12.1)
	DeviceMgmt         DeviceMgmt        `json:"device-mgmt"`                   // Device management related config (Live: IOS-XE 17.12.5)
	UserMgmt           UserMgmt          `json:"user-mgmt"`                     // User management related config (Live: IOS-XE 17.12.5)
	Tunnel             Tunnel            `json:"tunnel"`                        // Global capwap prefer-mode (Live: IOS-XE 17.12.5)
	CAPWAPTimer        CAPWAPTimer       `json:"capwap-timer"`                  // CAPWAP timer related config (Live: IOS-XE 17.12.5)
	Syslog             Syslog            `json:"syslog"`                        // AP sys log related config (Live: IOS-XE 17.12.5)
	Hyperlocation      Hyperlocation     `json:"hyperlocation"`                 // Hyperlocation config for the AP (Live: IOS-XE 17.12.5)
	RogueDetection     RogueDetection    `json:"rogue-detection"`               // Rogue detection related params (Live: IOS-XE 17.12.5)
	TftpDownGrade      TftpDownGrade     `json:"tftp-down-grade"`               // Tftp downgrade related config (Live: IOS-XE 17.12.5)
	ReportingInterval  ReportingInterval `json:"reporting-interval"`            // Interval at which AP should send stats (Live: IOS-XE 17.12.5)
	GasRateLimit       GasRateLimit      `json:"gas-rate-limit"`                // GAS rate limiting for Hotspot 2.0 (Live: IOS-XE 17.12.5)
	NtpServerInfo      NtpServerInfo     `json:"ntp-server-info"`               // NTP server info to be used by AP (Live: IOS-XE 17.12.5)
	PublicIPDiscovery  bool              `json:"public-ip-discovery"`           // Discovery Response from public IP enabled (Live: IOS-XE 17.12.5)
	Oeap               Oeap              `json:"oeap"`                          // Office Extended AP config (Live: IOS-XE 17.12.5)
	ApTzConfig         ApTzConfig        `json:"ap-tz-config"`                  // AP timezone config (Live: IOS-XE 17.12.5)
	RadioStatsMonitor  RadioStatsMonitor `json:"radio-stats-monitor"`           // AP radio statistics monitoring config (Live: IOS-XE 17.12.5)
	ApProfPpCfg        ApProfPpCfg       `json:"ap-prof-pp-cfg"`                // Power profile config per AP profile (Live: IOS-XE 17.12.5)
}

// SiteListEntry represents site list entry (Live: IOS-XE 17.12.5).
type SiteListEntry struct {
	SiteTagName              string  `json:"site-tag-name"`                         // Site tag name identifier (Live: IOS-XE 17.12.5)
	Description              *string `json:"description,omitempty"`                 // Description of the Site Tag (Live: IOS-XE 17.12.5)
	FlexProfile              *string `json:"flex-profile,omitempty"`                // Flex profile part of Site tag (Live: IOS-XE 17.12.5)
	ApJoinProfile            *string `json:"ap-join-profile,omitempty"`             // AP Join profile part of Site tag (Live: IOS-XE 17.12.5)
	IsLocalSite              *bool   `json:"is-local-site,omitempty"`               // Parameter to enable local site (Live: IOS-XE 17.12.5)
	FabricControlPlaneName   *string `json:"fabric-control-plane-name,omitempty"`   // Fabric Control Plane Name (YANG: IOS-XE 17.12.1)
	ImageDownloadProfileName *string `json:"image-download-profile-name,omitempty"` // Image Download Profile Name (YANG: IOS-XE 17.12.1)
	ArpCaching               *bool   `json:"arp-caching,omitempty"`                 // Enable or disable AP ARP caching (YANG: IOS-XE 17.12.1)
	DHCPBcast                *bool   `json:"dhcp-bcast,omitempty"`                  // Enable or disable fabric AP DHCP broadcast (YANG: IOS-XE 17.12.1)
	FabricMcastIPv4Addr      *string `json:"fabric-mcast-ipv4-addr,omitempty"`      // Fabric multicast group IPv4 address (YANG: IOS-XE 17.12.1)
	Load                     *int    `json:"load,omitempty"`                        // Estimate of relative load by site (YANG: IOS-XE 17.12.1)
}

// StatsTimer represents statistics timer config (Live: IOS-XE 17.12.5).
type StatsTimer struct {
	StatsTimer int `json:"stats-timer"` // Stats timer for the AP (Live: IOS-XE 17.12.5)
}

// JumboMtu represents jumbo MTU config (Live: IOS-XE 17.12.5).
type JumboMtu struct {
	JumboMtu bool `json:"jumbo-mtu"` // Jumbo MTU enable flag (Live: IOS-XE 17.12.5)
}

// DeviceMgmt represents device management config (Live: IOS-XE 17.12.5).
type DeviceMgmt struct {
	SSH bool `json:"ssh"` // SSH access enable flag (Live: IOS-XE 17.12.5)
}

// UserMgmt represents user management config (Live: IOS-XE 17.12.5).
type UserMgmt struct {
	Username     string `json:"username"`      // Admin username (Live: IOS-XE 17.12.5)
	Password     string `json:"password"`      // Admin password (Live: IOS-XE 17.12.5)
	PasswordType string `json:"password-type"` // Password encryption type (Live: IOS-XE 17.12.5)
	Secret       string `json:"secret"`        // Enable secret (Live: IOS-XE 17.12.5)
	SecretType   string `json:"secret-type"`   // Secret encryption type (Live: IOS-XE 17.12.5)
}

// Tunnel represents tunnel config (Live: IOS-XE 17.12.5).
type Tunnel struct {
	PreferredMode string `json:"preferred-mode"` // Tunnel preferred mode (Live: IOS-XE 17.12.5)
}

// CAPWAPTimer represents CAPWAP timer config (Live: IOS-XE 17.12.5).
type CAPWAPTimer struct {
	FastHeartBeatTimeout int `json:"fast-heart-beat-timeout"` // Fast heartbeat timeout in seconds (Live: IOS-XE 17.12.5)
}

// Syslog represents syslog config (Live: IOS-XE 17.12.5).
type Syslog struct {
	FacilityValue string `json:"facility-value"` // Syslog facility value (Live: IOS-XE 17.12.5)
	LogLevel      string `json:"log-level"`      // Syslog level (Live: IOS-XE 17.12.5)
	Host          string `json:"host"`           // Syslog server host (Live: IOS-XE 17.12.5)
	TLSMode       bool   `json:"tls-mode"`       // TLS mode enable flag (Live: IOS-XE 17.12.5)
}

// Hyperlocation represents hyperlocation config (Live: IOS-XE 17.12.5).
type Hyperlocation struct {
	HyperlocationEnable bool `json:"hyperlocation-enable"` // Hyperlocation enable flag (Live: IOS-XE 17.12.5)
}

// RogueDetection represents rogue detection config (Live: IOS-XE 17.12.5).
type RogueDetection struct {
	RogueDetectionMonitorModeReportInterval int `json:"rogue-detection-monitor-mode-report-interval"` // Rogue detection report interval seconds (Live: IOS-XE 17.12.5)
	ApRogueDetectionMinRSSI                 int `json:"ap-rogue-detection-min-rssi"`                  // Minimum RSSI threshold for rogue detect (Live: IOS-XE 17.12.5)
	ApRogueDetectionTransientInterval       int `json:"ap-rogue-detection-transient-interval"`        // Transient interval for rogue detect seconds (Live: IOS-XE 17.12.5)
}

// TftpDownGrade represents TFTP downgrade config (Live: IOS-XE 17.12.5).
type TftpDownGrade struct {
	TftpDowngradeIPAddress string `json:"tftp-downgrade-ip-address"` // TFTP server IP address for downgrade (Live: IOS-XE 17.12.5)
}

// ReportingInterval represents reporting interval config (Live: IOS-XE 17.12.5).
type ReportingInterval struct {
	Radio24GHz int `json:"radio-24ghz"` // 2.4GHz radio reporting interval seconds (Live: IOS-XE 17.12.5)
	Radio5GHz  int `json:"radio-5ghz"`  // 5GHz radio reporting interval seconds (Live: IOS-XE 17.12.5)
}

// GasRateLimit represents GAS rate limit config (Live: IOS-XE 17.12.5).
type GasRateLimit struct {
	NumReqPerInterval int `json:"num-req-per-interval"` // Number of requests per interval (Live: IOS-XE 17.12.5)
	IntervalMsec      int `json:"interval-msec"`        // Interval in milliseconds (Live: IOS-XE 17.12.5)
}

// NtpServerInfo represents NTP server config (Live: IOS-XE 17.12.5).
type NtpServerInfo struct {
	NtpAddress string `json:"ntp-address"` // NTP server IP address (Live: IOS-XE 17.12.5)
}

// Oeap represents OEAP config (Live: IOS-XE 17.12.5).
type Oeap struct {
	OeapDataEncr bool `json:"oeap-data-encr"` // OEAP data encryption enable flag (Live: IOS-XE 17.12.5)
	IsLocalNet   bool `json:"is-local-net"`   // Local network flag (Live: IOS-XE 17.12.5)
	ProvSsid     bool `json:"prov-ssid"`      // Provisioning SSID enable flag (Live: IOS-XE 17.12.5)
}

// ApTzConfig represents AP timezone config (Live: IOS-XE 17.12.5).
type ApTzConfig struct {
	TzEnabled bool   `json:"tz-enabled"` // Timezone enable flag (Live: IOS-XE 17.12.5)
	Mode      string `json:"mode"`       // Timezone mode (Live: IOS-XE 17.12.5)
}

// RadioStatsMonitor represents radio stats monitoring config (Live: IOS-XE 17.12.5).
type RadioStatsMonitor struct {
	Enable       bool    `json:"enable"`        // Radio statistics monitoring enable flag (Live: IOS-XE 17.12.5)
	AlarmsEnable []*bool `json:"alarms-enable"` // Alarms enable flags array (Live: IOS-XE 17.12.5)
	RadioReset   bool    `json:"radio-reset"`   // Radio reset enable flag (Live: IOS-XE 17.12.5)
}

// ApProfPpCfg represents AP profile power config (Live: IOS-XE 17.12.5).
type ApProfPpCfg struct {
	PowerProfileName string `json:"power-profile-name"` // Power profile name (Live: IOS-XE 17.12.5)
}

// LedState represents AP LED config based on st-led-state (YANG: IOS-XE 17.12.1).
type LedState struct {
	LedState bool `json:"led-state"` // LED state enable flag (YANG: IOS-XE 17.12.1)
}

// LinkLatency represents AP link auditing config based on st-link-latency (YANG: IOS-XE 17.12.1).
type LinkLatency struct {
	LinkLatencyFlag string `json:"link-latency-flag"` // Link latency flag config (YANG: IOS-XE 17.12.1)
}

// ApMode represents AP mode config based on st-mode (YANG: IOS-XE 17.12.1).
type ApMode struct {
	ApSubMode   string `json:"ap-sub-mode"`            // AP sub-mode type (YANG: IOS-XE 17.12.1)
	FastChannel *int   `json:"fast-channel,omitempty"` // Fast channel config (YANG: IOS-XE 17.12.1)
}

// Poe represents PoE config based on st-poe-cfg (YANG: IOS-XE 17.12.1).
type Poe struct {
	PreStandard8023afSwitchFlag bool   `json:"pre-standard8023af-switch-flag"` // Pre-standard 802.3af switch flag (YANG: IOS-XE 17.12.1)
	PowerInjectorState          bool   `json:"power-injector-state"`           // Power injector state flag (YANG: IOS-XE 17.12.1)
	PowerInjectorSelection      string `json:"power-injector-selection"`       // Power injector selection type (YANG: IOS-XE 17.12.1)
}
