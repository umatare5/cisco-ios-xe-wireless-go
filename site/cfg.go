// Package site provides site configuration functionality for the Cisco Wireless Network Controller API.
package site

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// SiteCfgBasePath defines the base path for site configuration endpoints
	SiteCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"
	// ApCfgProfilesEndpoint retrieves AP configuration profiles
	ApCfgProfilesEndpoint = SiteCfgBasePath + "/ap-cfg-profiles"
	// SiteTagConfigsEndpoint retrieves site tag configurations
	SiteTagConfigsEndpoint = SiteCfgBasePath + "/site-tag-configs"
)

// SiteCfgResponse represents the complete site configuration response
type SiteCfgResponse struct {
	CiscoIOSXEWirelessSiteCfgData struct {
		ApCfgProfiles  []ApCfgProfile  `json:"ap-cfg-profiles"`
		SiteTagConfigs []SiteTagConfig `json:"site-tag-configs"`
	} `json:"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data"`
}

// SiteApCfgProfilesResponse represents the AP configuration profiles response
type SiteApCfgProfilesResponse struct {
	ApCfgProfiles []ApCfgProfile `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles"`
}

// SiteTagConfigsResponse represents the site tag configurations response
type SiteTagConfigsResponse struct {
	SiteTagConfigs []SiteTagConfig `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs"`
}

// ApCfgProfile represents an access point configuration profile with various settings
type ApCfgProfile struct {
	ProfileName string `json:"profile-name"`          // Profile name identifier
	Description string `json:"description,omitempty"` // Optional profile description
	StatsTimer  *struct {
		StatsTimer int `json:"stats-timer"` // Statistics collection timer interval
	} `json:"stats-timer,omitempty"`
	UserMgmt *struct {
		Username     string `json:"username"`      // Management username
		Password     string `json:"password"`      // Management password
		PasswordType string `json:"password-type"` // Password type (clear/encrypted)
		Secret       string `json:"secret"`        // Enable secret
		SecretType   string `json:"secret-type"`   // Secret type (clear/encrypted)
	} `json:"user-mgmt,omitempty"`
	Tunnel *struct {
		PreferredMode string `json:"preferred-mode"` // Preferred tunnel mode
	} `json:"tunnel,omitempty"`
	CapwapTimer *struct {
		FastHeartBeatTimeout int `json:"fast-heart-beat-timeout"` // CAPWAP fast heartbeat timeout
	} `json:"capwap-timer,omitempty"`
	Dot1xEapTypeInfo *struct {
		Dot1xEapType string `json:"dot1x-eap-type"` // 802.1X EAP type
	} `json:"dot1x-eap-type-info,omitempty"`
	Syslog *struct {
		FacilityValue string `json:"facility-value,omitempty"`
		LogLevel      string `json:"log-level,omitempty"`
		Host          string `json:"host,omitempty"`
	} `json:"syslog,omitempty"`
	Hyperlocation *struct {
		HyperlocationEnable bool `json:"hyperlocation-enable"`
	} `json:"hyperlocation,omitempty"`
	RogueDetection *struct {
		ApRogueDetectionMinRssi int `json:"ap-rogue-detection-min-rssi"`
	} `json:"rogue-detection,omitempty"`
	ReportingInterval *struct {
		Radio24Ghz int `json:"radio-24ghz"`
		Radio5Ghz  int `json:"radio-5ghz"`
	} `json:"reporting-interval,omitempty"`
	PublicIPDiscovery *bool `json:"public-ip-discovery,omitempty"`
	Oeap              *struct {
		OeapDataEncr bool `json:"oeap-data-encr"`
		ProvSsid     bool `json:"prov-ssid"`
	} `json:"oeap,omitempty"`
	RadioStatsMonitor *struct {
		Enable       bool  `json:"enable"`
		AlarmsEnable []any `json:"alarms-enable"`
		RadioReset   bool  `json:"radio-reset"`
	} `json:"radio-stats-monitor,omitempty"`
	DeviceMgmt *struct {
		SSH bool `json:"ssh"`
	} `json:"device-mgmt,omitempty"`
}

type SiteTagConfig struct {
	SiteTagName   string `json:"site-tag-name"`
	Description   string `json:"description,omitempty"`
	FlexProfile   string `json:"flex-profile,omitempty"`
	ApJoinProfile string `json:"ap-join-profile,omitempty"`
	IsLocalSite   bool   `json:"is-local-site"`
}

func GetSiteCfg(client *wnc.Client, ctx context.Context) (*SiteCfgResponse, error) {
	var data SiteCfgResponse
	if err := client.SendAPIRequest(ctx, SiteCfgBasePath, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetSiteApCfgProfiles(client *wnc.Client, ctx context.Context) (*SiteApCfgProfilesResponse, error) {
	var data SiteApCfgProfilesResponse
	if err := client.SendAPIRequest(ctx, ApCfgProfilesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetSiteTagConfigs(client *wnc.Client, ctx context.Context) (*SiteTagConfigsResponse, error) {
	var data SiteTagConfigsResponse
	if err := client.SendAPIRequest(ctx, SiteTagConfigsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
