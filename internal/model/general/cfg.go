// Package general provides type definitions for Cisco IOS-XE wireless controller operations.
package general

// GeneralCfg represents the general configuration response.
type GeneralCfg struct {
	CiscoIOSXEWirelessGeneralCfgGeneralCfgData struct {
		MewlcConfig             MewlcConfig              `json:"mewlc-config"`                          // Embedded Wireless Controller configuration (Live: IOS-XE 17.12.5)
		CacConfig               CacConfig                `json:"cac-config"`                            // CAC resources values configuration (Live: IOS-XE 17.12.5)
		Mfp                     Mfp                      `json:"mfp"`                                   // Management Frame Protection configuration (Live: IOS-XE 17.12.5)
		FipsCfg                 FipsCfg                  `json:"fips-cfg"`                              // DTLS for AP join configurations (Live: IOS-XE 17.12.5)
		WsaApClientEvent        WsaApClientEvent         `json:"wsa-ap-client-event"`                   // Client event config parameters for AP (Live: IOS-XE 17.12.5)
		SimL3InterfaceCacheData *SimL3InterfaceCacheData `json:"sim-l3-interface-cache-data,omitempty"` // Wireless management interface data (Live: IOS-XE 17.12.5)
		WlcManagementData       *WlcManagementData       `json:"wlc-management-data,omitempty"`         // WLC management certificate and authorization config (Live: IOS-XE 17.12.5)
		Laginfo                 Laginfo                  `json:"laginfo"`                               // AP LAG information (Live: IOS-XE 17.12.5)
		MulticastConfig         *MulticastConfig         `json:"multicast-config,omitempty"`            // Broadcast/Multicast configuration (Live: IOS-XE 17.12.5)
		FeatureUsageCfg         FeatureUsageCfg          `json:"feature-usage-cfg"`                     // Wireless feature usage monitoring configuration (Live: IOS-XE 17.12.5)
		ThresholdWarnCfg        ThresholdWarnCfg         `json:"threshold-warn-cfg"`                    // Threshold warnings configuration (Live: IOS-XE 17.12.5)
		ApLocRangingCfg         ApLocRangingCfg          `json:"ap-loc-ranging-cfg"`                    // Location calendar profile configuration (Live: IOS-XE 17.12.5)
		GeolocationCfg          GeolocationCfg           `json:"geolocation-cfg"`                       // Wireless geolocation configuration (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data"` // General configuration data (Live: IOS-XE 17.12.5)
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
	MewlcPlatform       *bool   `json:"mewlc-platform,omitempty"`        // Embedded Wireless Controller Platform (Live: IOS-XE 17.12.5)
	MewlcVrrpVrid       *int    `json:"mewlc-vrrp-vrid,omitempty"`       // EWC VRRP Virtual Router Identifier (Live: IOS-XE 17.12.5)
	PreferredMasterName *string `json:"preferred-master-name,omitempty"` // Preferred Master AP Name (Live: IOS-XE 17.12.5)
	PImgDwnld           *bool   `json:"p-img-dwnld,omitempty"`           // Download AP image in parallel on Image Master APs (Live: IOS-XE 17.12.5)
}

// CacConfigData represents CAC configuration data structure.
type CacConfigData struct {
	IplearnqHighThreshold *int `json:"iplearnq-high-threshold,omitempty"` // High threshold value for IplearnQ (Live: IOS-XE 17.12.5)
	AaaHighThreshold      *int `json:"aaa-high-threshold,omitempty"`      // High threshold for AAA outstanding request queue (Live: IOS-XE 17.12.5)
	IpcHighThreshold      *int `json:"ipc-high-threshold,omitempty"`      // High threshold value for IPC resources (Live: IOS-XE 17.12.5)
	IpcLowThreshold       *int `json:"ipc-low-threshold,omitempty"`       // Low threshold value for IPC resources (Live: IOS-XE 17.12.5)
}

// MfpData represents Management Frame Protection data structure.
type MfpData struct {
	GlobalMfpState           *bool `json:"global-mfp-state,omitempty"`           // Global MFP enabled status (Live: IOS-XE 17.12.5)
	ApImpersonationDetection *bool `json:"ap-impersonation-detection,omitempty"` // AP impersonation detection status (Live: IOS-XE 17.12.5)
	MfpKeyRefreshInterval    *int  `json:"mfp-key-refresh-interval,omitempty"`   // Key refresh interval in hours (Live: IOS-XE 17.12.5)
}

// FipsCfgData represents FIPS configuration data structure.
type FipsCfgData struct {
	DTLSVersion     *string `json:"dtls-version,omitempty"`     // DTLS version used for AP join (Live: IOS-XE 17.12.5)
	DTLSCiphersuite *string `json:"dtls-ciphersuite,omitempty"` // Configure DTLS ciphersuite (Live: IOS-XE 17.12.5)
}

// WsaApClientEventData represents WSA AP client event data structure.
type WsaApClientEventData struct {
	Frequency *int `json:"frequency,omitempty"` // Frequency in sec at which events from AP sent to WLC (Live: IOS-XE 17.12.5)
}

// SimL3InterfaceCacheDataInfo represents SIM L3 interface cache data information.
type SimL3InterfaceCacheDataInfo struct {
	InterfaceName      string  `json:"interface-name"`                 // Wireless management interface name (Live: IOS-XE 17.12.5)
	NatIPAddress       *string `json:"nat-ip-address,omitempty"`       // NAT IP address (Live: IOS-XE 17.12.5)
	NatEnable          *bool   `json:"nat-enable,omitempty"`           // NAT IP address is enabled or disabled (Live: IOS-XE 17.12.5)
	PrivateIPDiscovery *bool   `json:"private-ip-discovery,omitempty"` // Discovery response from private IP enabled/disabled (Live: IOS-XE 17.12.5)
}

// WlcManagementDataInfo represents WLC management data information.
type WlcManagementDataInfo struct {
	PkiTrustpointName string  `json:"pki-trustpoint-name"`           // Wireless management trustpoint name (Live: IOS-XE 17.12.5)
	SscAuthToken      *string `json:"ssc-auth-token,omitempty"`      // SSC authorization token (Live: IOS-XE 17.12.5)
	SscAuthTokenType  *string `json:"ssc-auth-token-type,omitempty"` // SSC authorization token encryption type (Live: IOS-XE 17.12.5)
}

// LaginfoData represents LAG information data structure.
type LaginfoData struct {
	Enabled *bool `json:"enabled,omitempty"` // Global lag status (Live: IOS-XE 17.12.5)
}

// MulticastConfigData represents multicast configuration data structure.
type MulticastConfigData struct {
	IsMDNSEnabled                  bool    `json:"is-mdns-enabled"`                              // Flag to enable or disable mdns (Live: IOS-XE 17.12.5)
	MulticastOverMulticastIPv4Addr *string `json:"multicast-over-multicast-ipv4-addr,omitempty"` // IPv4 multicast group address for CAPWAP used by APs (Live: IOS-XE 17.12.5)
	MulticastOverMulticastIPv6Addr *string `json:"multicast-over-multicast-ipv6-addr,omitempty"` // IPv6 multicast group address for CAPWAP used by APs (Live: IOS-XE 17.12.5)
	IsMcastEnabled                 *bool   `json:"is-mcast-enabled,omitempty"`                   // Multicast enable/disable (Live: IOS-XE 17.12.5)
	IsNonIPMulticastEnabled        *bool   `json:"is-non-ip-multicast-enabled,omitempty"`        // Non-ip multicast enable/disable (Live: IOS-XE 17.12.5)
}

// FeatureUsageCfgData represents feature usage configuration data structure.
type FeatureUsageCfgData struct {
	Enable *bool `json:"enable,omitempty"` // Enable wireless feature usage monitoring (Live: IOS-XE 17.12.5)
}

// ThresholdWarnCfgData represents threshold warning configuration data structure.
type ThresholdWarnCfgData struct {
	ThresholdWarning *bool `json:"threshold-warning,omitempty"` // Enable or disable threshold warning functionality (Live: IOS-XE 17.12.5)
	ClientsThreshold *int  `json:"clients-threshold,omitempty"` // Configure clients threshold (Live: IOS-XE 17.12.5)
	WarningPeriod    *int  `json:"warning-period,omitempty"`    // Configure warning check periodicity (Live: IOS-XE 17.12.5)
}

// ApLocRangingCfgData represents AP location ranging configuration data structure.
type ApLocRangingCfgData struct {
	DeriveGeolocation *bool `json:"derive-geolocation,omitempty"` // Enable/disable geolocation derivation (Live: IOS-XE 17.12.5)
}

// GeolocationCfgData represents geolocation configuration data structure.
type GeolocationCfgData struct {
	EnableDerivation *bool `json:"enable-derivation,omitempty"` // Enable wireless geolocation derivation (Live: IOS-XE 17.12.5)
}
