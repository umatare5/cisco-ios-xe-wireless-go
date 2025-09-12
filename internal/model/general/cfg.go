// Package general provides type definitions for Cisco IOS-XE wireless controller operations.
package general

// GeneralCfg represents the general configuration response.
type GeneralCfg struct {
	CiscoIOSXEWirelessGeneralCfgGeneralCfgData struct {
		MewlcConfig             MewlcConfig              `json:"mewlc-config"`
		CacConfig               CacConfig                `json:"cac-config"`
		Mfp                     Mfp                      `json:"mfp"`
		FipsCfg                 FipsCfg                  `json:"fips-cfg"`
		WsaApClientEvent        WsaApClientEvent         `json:"wsa-ap-client-event"`
		SimL3InterfaceCacheData *SimL3InterfaceCacheData `json:"sim-l3-interface-cache-data,omitempty"`
		WlcManagementData       *WlcManagementData       `json:"wlc-management-data,omitempty"`
		Laginfo                 Laginfo                  `json:"laginfo"`
		MulticastConfig         *MulticastConfig         `json:"multicast-config,omitempty"`
		FeatureUsageCfg         FeatureUsageCfg          `json:"feature-usage-cfg"`
		ThresholdWarnCfg        ThresholdWarnCfg         `json:"threshold-warn-cfg"`
		ApLocRangingCfg         ApLocRangingCfg          `json:"ap-loc-ranging-cfg"`
		GeolocationCfg          GeolocationCfg           `json:"geolocation-cfg"`
	} `json:"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data"`
}

// GeneralCfgMewlcConfig represents the corresponding data structure.
type GeneralCfgMewlcConfig struct {
	MewlcConfig MewlcConfig `json:"Cisco-IOS-XE-wireless-general-cfg:mewlc-config"`
}

// GeneralCfgCacConfig represents the corresponding data structure.
type GeneralCfgCacConfig struct {
	CacConfig CacConfig `json:"Cisco-IOS-XE-wireless-general-cfg:cac-config"`
}

// GeneralCfgMfp represents the corresponding data structure.
type GeneralCfgMfp struct {
	Mfp Mfp `json:"Cisco-IOS-XE-wireless-general-cfg:mfp"`
}

// GeneralCfgFipsCfg represents the corresponding data structure.
type GeneralCfgFipsCfg struct {
	FipsCfg FipsCfg `json:"Cisco-IOS-XE-wireless-general-cfg:fips-cfg"`
}

// GeneralCfgWsaApClientEvent represents the corresponding data structure.
type GeneralCfgWsaApClientEvent struct {
	WsaApClientEvent WsaApClientEvent `json:"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event"`
}

// GeneralCfgSimL3InterfaceCacheData represents the corresponding data structure.
type GeneralCfgSimL3InterfaceCacheData struct {
	SimL3InterfaceCacheData *SimL3InterfaceCacheData `json:"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data,omitempty"`
}

// GeneralCfgWlcManagementData represents the corresponding data structure.
type GeneralCfgWlcManagementData struct {
	WlcManagementData *WlcManagementData `json:"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data,omitempty"`
}

// GeneralCfgLaginfo represents the corresponding data structure.
type GeneralCfgLaginfo struct {
	Laginfo Laginfo `json:"Cisco-IOS-XE-wireless-general-cfg:laginfo"`
}

// GeneralCfgMulticastConfig represents the corresponding data structure.
type GeneralCfgMulticastConfig struct {
	MulticastConfig *MulticastConfig `json:"Cisco-IOS-XE-wireless-general-cfg:multicast-config,omitempty"`
}

// GeneralCfgFeatureUsageCfg represents the corresponding data structure.
type GeneralCfgFeatureUsageCfg struct {
	FeatureUsageCfg FeatureUsageCfg `json:"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg"`
}

// GeneralCfgThresholdWarnCfg represents the corresponding data structure.
type GeneralCfgThresholdWarnCfg struct {
	ThresholdWarnCfg ThresholdWarnCfg `json:"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg"`
}

// GeneralCfgApLocRangingCfg represents the corresponding data structure.
type GeneralCfgApLocRangingCfg struct {
	ApLocRangingCfg ApLocRangingCfg `json:"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg"`
}

// GeneralCfgGeolocationCfg represents the corresponding data structure.
type GeneralCfgGeolocationCfg struct {
	GeolocationCfg GeolocationCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg"`
}

// MewlcConfig represents the MEWLC configuration response.
type MewlcConfig struct {
	MewlcConfigData MewlcConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:mewlc-config"`
}

// CacConfig represents the CAC configuration response.
type CacConfig struct {
	CacConfigData CacConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:cac-config"`
}

// Mfp represents the Management Frame Protection configuration response.
type Mfp struct {
	MfpData MfpData `json:"Cisco-IOS-XE-wireless-general-cfg:mfp"`
}

// FipsCfg represents FIPS configuration.
type FipsCfg struct {
	FipsCfgData FipsCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:fips-cfg"`
}

// WsaApClientEvent represents WSA AP client event configuration.
type WsaApClientEvent struct {
	WsaApClientEventData WsaApClientEventData `json:"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event"`
}

// SimL3InterfaceCacheData represents SIM L3 interface cache data.
type SimL3InterfaceCacheData struct {
	SimL3InterfaceCacheDataInfo SimL3InterfaceCacheDataInfo `json:"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data"`
}

// WlcManagementData represents WLC management data.
type WlcManagementData struct {
	WlcManagementDataInfo WlcManagementDataInfo `json:"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data"`
}

// Laginfo represents LAG information.
type Laginfo struct {
	LaginfoData LaginfoData `json:"Cisco-IOS-XE-wireless-general-cfg:laginfo"`
}

// MulticastConfig represents multicast configuration.
type MulticastConfig struct {
	MulticastConfigData MulticastConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:multicast-config"`
}

// FeatureUsageCfg represents feature usage configuration.
type FeatureUsageCfg struct {
	FeatureUsageCfgData FeatureUsageCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg"`
}

// ThresholdWarnCfg represents threshold warning configuration.
type ThresholdWarnCfg struct {
	ThresholdWarnCfgData ThresholdWarnCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg"`
}

// ApLocRangingCfg represents AP location ranging configuration.
type ApLocRangingCfg struct {
	ApLocRangingCfgData ApLocRangingCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg"`
}

// GeolocationCfg represents geolocation configuration.
type GeolocationCfg struct {
	GeolocationCfgData GeolocationCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg"`
}

// MewlcConfigData represents MEWLC configuration data structure.
type MewlcConfigData struct {
	MewlcPlatform       *bool   `json:"mewlc-platform,omitempty"`        // MEWLC platform enable flag (YANG: IOS-XE 17.12.1+)
	MewlcVrrpVrid       *int    `json:"mewlc-vrrp-vrid,omitempty"`       // MEWLC VRRP VRID value (YANG: IOS-XE 17.12.1+)
	PreferredMasterName *string `json:"preferred-master-name,omitempty"` // Preferred master controller name (YANG: IOS-XE 17.12.1+)
	PImgDwnld           *bool   `json:"p-img-dwnld,omitempty"`           // Image download permission flag (YANG: IOS-XE 17.12.1+)
}

// CacConfigData represents CAC configuration data structure.
type CacConfigData struct {
	IplearnqHighThreshold *int `json:"iplearnq-high-threshold,omitempty"` // IP learning queue high threshold (YANG: IOS-XE 17.12.1+)
	AaaHighThreshold      *int `json:"aaa-high-threshold,omitempty"`      // AAA high threshold (YANG: IOS-XE 17.12.1+)
	IpcHighThreshold      *int `json:"ipc-high-threshold,omitempty"`      // IPC high threshold (YANG: IOS-XE 17.12.1+)
	IpcLowThreshold       *int `json:"ipc-low-threshold,omitempty"`       // IPC low threshold (YANG: IOS-XE 17.12.1+)
}

// MfpData represents Management Frame Protection data structure.
type MfpData struct {
	GlobalMfpState           *bool `json:"global-mfp-state,omitempty"`           // Global MFP enable state (YANG: IOS-XE 17.12.1+)
	ApImpersonationDetection *bool `json:"ap-impersonation-detection,omitempty"` // AP impersonation detection enable flag (YANG: IOS-XE 17.12.1+)
	MfpKeyRefreshInterval    *int  `json:"mfp-key-refresh-interval,omitempty"`   // MFP key refresh interval in seconds (YANG: IOS-XE 17.12.1+)
}

// FipsCfgData represents FIPS configuration data structure.
type FipsCfgData struct {
	DtlsVersion     *string `json:"dtls-version,omitempty"`     // DTLS version specification (YANG: IOS-XE 17.12.1+)
	DtlsCiphersuite *string `json:"dtls-ciphersuite,omitempty"` // DTLS cipher suite specification (YANG: IOS-XE 17.12.1+)
}

// WsaApClientEventData represents WSA AP client event data structure.
type WsaApClientEventData struct {
	Frequency *int `json:"frequency,omitempty"` // Event frequency in seconds (YANG: IOS-XE 17.12.1+)
}

// SimL3InterfaceCacheDataInfo represents SIM L3 interface cache data information.
type SimL3InterfaceCacheDataInfo struct {
	InterfaceName      string  `json:"interface-name"`                 // SIM L3 interface name
	NatIPAddress       *string `json:"nat-ip-address,omitempty"`       // NAT IP address (YANG: IOS-XE 17.12.1+)
	NatEnable          *bool   `json:"nat-enable,omitempty"`           // NAT enable status (YANG: IOS-XE 17.12.1+)
	PrivateIPDiscovery *bool   `json:"private-ip-discovery,omitempty"` // Private IP discovery enable flag (YANG: IOS-XE 17.12.1+)
}

// WlcManagementDataInfo represents WLC management data information.
type WlcManagementDataInfo struct {
	PkiTrustpointName string  `json:"pki-trustpoint-name"`           // PKI trustpoint name
	SscAuthToken      *string `json:"ssc-auth-token,omitempty"`      // SSC authentication token (YANG: IOS-XE 17.12.1+)
	SscAuthTokenType  *string `json:"ssc-auth-token-type,omitempty"` // SSC authentication token type (YANG: IOS-XE 17.12.1+)
}

// LaginfoData represents LAG information data structure.
type LaginfoData struct {
	// LAG configuration fields would be defined here based on YANG model (YANG: IOS-XE 17.12.1+)
}

// MulticastConfigData represents multicast configuration data structure.
type MulticastConfigData struct {
	IsMdnsEnabled                  bool    `json:"is-mdns-enabled"`                              // mDNS enable status
	MulticastOverMulticastIPv4Addr *string `json:"multicast-over-multicast-ipv4-addr,omitempty"` // Multicast over multicast IPv4 address (YANG: IOS-XE 17.12.1+)
	MulticastOverMulticastIPv6Addr *string `json:"multicast-over-multicast-ipv6-addr,omitempty"` // Multicast over multicast IPv6 address (YANG: IOS-XE 17.12.1+)
	IsMcastEnabled                 *bool   `json:"is-mcast-enabled,omitempty"`                   // Multicast enable status (YANG: IOS-XE 17.12.1+)
	IsNonIPMulticastEnabled        *bool   `json:"is-non-ip-multicast-enabled,omitempty"`        // Non-IP multicast enable status (YANG: IOS-XE 17.12.1+)
}

// FeatureUsageCfgData represents feature usage configuration data structure.
type FeatureUsageCfgData struct {
	Enable *bool `json:"enable,omitempty"` // Feature usage enable flag (YANG: IOS-XE 17.12.1+)
}

// ThresholdWarnCfgData represents threshold warning configuration data structure.
type ThresholdWarnCfgData struct {
	ThresholdWarning *bool `json:"threshold-warning,omitempty"` // Threshold warning enable flag (YANG: IOS-XE 17.12.1+)
	ClientsThreshold *int  `json:"clients-threshold,omitempty"` // Client count threshold value (YANG: IOS-XE 17.12.1+)
	WarningPeriod    *int  `json:"warning-period,omitempty"`    // Warning period in seconds (YANG: IOS-XE 17.12.1+)
}

// ApLocRangingCfgData represents AP location ranging configuration data structure.
type ApLocRangingCfgData struct {
	// AP location ranging configuration fields would be defined here based on YANG model (YANG: IOS-XE 17.12.1+)
}

// GeolocationCfgData represents geolocation configuration data structure.
type GeolocationCfgData struct {
	// Geolocation configuration fields would be defined here based on YANG model (YANG: IOS-XE 17.12.1+)
}
