// Package model provides data models for site configuration data.
package model

// SiteCfg  represents the Site configuration data.
type SiteCfg struct {
	SiteCfgData SiteCfgData `json:"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"`
}

// SiteCfgApCfgProfiles  represents the AP configuration profiles.
type SiteCfgApCfgProfiles struct {
	ApCfgProfiles ApCfgProfiles `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles"`
}

// SiteCfgSiteTagConfigs  represents the site tag configurations.
type SiteCfgSiteTagConfigs struct {
	SiteTagConfigs SiteTagConfigs `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs"`
}

// SiteCfgApCfgProfile  represents a single AP configuration profile.
type SiteCfgApCfgProfile struct {
	ApCfgProfile []ApCfgProfile `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profile"`
}

// SiteCfgSiteTagConfig  represents a single site tag configuration.
type SiteCfgSiteTagConfig struct {
	SiteTagConfig []SiteTagConfig `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"`
}

type SiteCfgData struct {
	ApCfgProfiles  ApCfgProfiles  `json:"ap-cfg-profiles"`
	SiteTagConfigs SiteTagConfigs `json:"site-tag-configs"`
}

type ApCfgProfiles struct {
	ApCfgProfile []ApCfgProfile `json:"ap-cfg-profile"`
}

type ApCfgProfile struct {
	ProfileName        string            `json:"profile-name"`
	DataEncryptionFlag bool              `json:"data-encryption-flag"`
	StatsTimer         StatsTimer        `json:"stats-timer"`
	JumboMtu           JumboMtu          `json:"jumbo-mtu"`
	DeviceMgmt         DeviceMgmt        `json:"device-mgmt"`
	UserMgmt           UserMgmt          `json:"user-mgmt"`
	Tunnel             Tunnel            `json:"tunnel"`
	CapwapTimer        CapwapTimer       `json:"capwap-timer"`
	Syslog             Syslog            `json:"syslog"`
	Hyperlocation      Hyperlocation     `json:"hyperlocation"`
	RogueDetection     RogueDetection    `json:"rogue-detection"`
	TftpDownGrade      TftpDownGrade     `json:"tftp-down-grade"`
	ReportingInterval  ReportingInterval `json:"reporting-interval"`
	GasRateLimit       GasRateLimit      `json:"gas-rate-limit"`
	NtpServerInfo      NtpServerInfo     `json:"ntp-server-info"`
	PublicIPDiscovery  bool              `json:"public-ip-discovery"`
	Oeap               Oeap              `json:"oeap"`
	ApTzConfig         ApTzConfig        `json:"ap-tz-config"`
	RadioStatsMonitor  RadioStatsMonitor `json:"radio-stats-monitor"`
	ApProfPpCfg        ApProfPpCfg       `json:"ap-prof-pp-cfg"`
}

// SiteCfgSiteListEntries  represents the site list entries.
type SiteCfgSiteListEntries struct {
	SiteListEntries SiteListEntries `json:"Cisco-IOS-XE-wireless-site-cfg:site-list-entries"`
}

type SiteListEntries struct {
	SiteListEntry []SiteListEntry `json:"site-list-entry"`
}

type SiteListEntry struct {
	SiteTagName   string `json:"site-tag-name"`
	Description   string `json:"description,omitempty"`
	FlexProfile   string `json:"flex-profile,omitempty"`
	ApJoinProfile string `json:"ap-join-profile,omitempty"`
	IsLocalSite   bool   `json:"is-local-site,omitempty"`
}

type SiteTagConfigs struct {
	SiteTagConfig []SiteTagConfig `json:"site-tag-config"`
}

type SiteTagConfig struct {
	SiteTagName   string `json:"site-tag-name"`
	Description   string `json:"description,omitempty"`
	FlexProfile   string `json:"flex-profile,omitempty"`
	ApJoinProfile string `json:"ap-join-profile,omitempty"`
	IsLocalSite   bool   `json:"is-local-site,omitempty"`
}

type StatsTimer struct {
	StatsTimer int `json:"stats-timer"`
}

type JumboMtu struct {
	JumboMtu bool `json:"jumbo-mtu"`
}

type DeviceMgmt struct {
	SSH bool `json:"ssh"`
}

type UserMgmt struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordType string `json:"password-type"`
	Secret       string `json:"secret"`
	SecretType   string `json:"secret-type"`
}

type Tunnel struct {
	PreferredMode string `json:"preferred-mode"`
}

type CapwapTimer struct {
	FastHeartBeatTimeout int `json:"fast-heart-beat-timeout"`
}

type Syslog struct {
	FacilityValue string `json:"facility-value"`
	LogLevel      string `json:"log-level"`
	Host          string `json:"host"`
	TLSMode       bool   `json:"tls-mode"`
}

type Hyperlocation struct {
	HyperlocationEnable bool `json:"hyperlocation-enable"`
}

type RogueDetection struct {
	RogueDetectionMonitorModeReportInterval int `json:"rogue-detection-monitor-mode-report-interval"`
	ApRogueDetectionMinRssi                 int `json:"ap-rogue-detection-min-rssi"`
	ApRogueDetectionTransientInterval       int `json:"ap-rogue-detection-transient-interval"`
}

type TftpDownGrade struct {
	TftpDowngradeIPAddress string `json:"tftp-downgrade-ip-address"`
}

type ReportingInterval struct {
	Radio24GHz int `json:"radio-24ghz"`
	Radio5GHz  int `json:"radio-5ghz"`
}

type GasRateLimit struct {
	NumReqPerInterval int `json:"num-req-per-interval"`
	IntervalMsec      int `json:"interval-msec"`
}

type NtpServerInfo struct {
	NtpAddress string `json:"ntp-address"`
}

type Oeap struct {
	OeapDataEncr bool `json:"oeap-data-encr"`
	IsLocalNet   bool `json:"is-local-net"`
	ProvSsid     bool `json:"prov-ssid"`
}

type ApTzConfig struct {
	TzEnabled bool   `json:"tz-enabled"`
	Mode      string `json:"mode"`
}

type RadioStatsMonitor struct {
	Enable       bool        `json:"enable"`
	AlarmsEnable interface{} `json:"alarms-enable"` // interface{} needed for Cisco API that returns bool/string/number types
	RadioReset   bool        `json:"radio-reset"`
}

type ApProfPpCfg struct {
	PowerProfileName string `json:"power-profile-name"`
}
