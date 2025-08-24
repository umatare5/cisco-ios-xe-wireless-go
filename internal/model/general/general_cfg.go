// Package model provides type definitions for Cisco IOS-XE wireless controller operations.
package model

// GeneralCfg  represents the general configuration response
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

// GeneralCfgMewlcConfig  represents the corresponding data structure.
type GeneralCfgMewlcConfig struct {
	MewlcConfig MewlcConfig `json:"Cisco-IOS-XE-wireless-general-cfg:mewlc-config"`
}

// GeneralCfgCacConfig  represents the corresponding data structure.
type GeneralCfgCacConfig struct {
	CacConfig CacConfig `json:"Cisco-IOS-XE-wireless-general-cfg:cac-config"`
}

// GeneralCfgMfp  represents the corresponding data structure.
type GeneralCfgMfp struct {
	Mfp Mfp `json:"Cisco-IOS-XE-wireless-general-cfg:mfp"`
}

// GeneralCfgFipsCfg  represents the corresponding data structure.
type GeneralCfgFipsCfg struct {
	FipsCfg FipsCfg `json:"Cisco-IOS-XE-wireless-general-cfg:fips-cfg"`
}

// GeneralCfgWsaApClientEvent  represents the corresponding data structure.
type GeneralCfgWsaApClientEvent struct {
	WsaApClientEvent WsaApClientEvent `json:"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event"`
}

// GeneralCfgSimL3InterfaceCacheData  represents the corresponding data structure.
type GeneralCfgSimL3InterfaceCacheData struct {
	SimL3InterfaceCacheData *SimL3InterfaceCacheData `json:"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data,omitempty"`
}

// GeneralCfgWlcManagementData  represents the corresponding data structure.
type GeneralCfgWlcManagementData struct {
	WlcManagementData *WlcManagementData `json:"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data,omitempty"`
}

// GeneralCfgLaginfo  represents the corresponding data structure.
type GeneralCfgLaginfo struct {
	Laginfo Laginfo `json:"Cisco-IOS-XE-wireless-general-cfg:laginfo"`
}

// GeneralCfgMulticastConfig  represents the corresponding data structure.
type GeneralCfgMulticastConfig struct {
	MulticastConfig *MulticastConfig `json:"Cisco-IOS-XE-wireless-general-cfg:multicast-config,omitempty"`
}

// GeneralCfgFeatureUsageCfg  represents the corresponding data structure.
type GeneralCfgFeatureUsageCfg struct {
	FeatureUsageCfg FeatureUsageCfg `json:"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg"`
}

// GeneralCfgThresholdWarnCfg  represents the corresponding data structure.
type GeneralCfgThresholdWarnCfg struct {
	ThresholdWarnCfg ThresholdWarnCfg `json:"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg"`
}

// GeneralCfgApLocRangingCfg  represents the corresponding data structure.
type GeneralCfgApLocRangingCfg struct {
	ApLocRangingCfg ApLocRangingCfg `json:"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg"`
}

// GeneralCfgGeolocationCfg  represents the corresponding data structure.
type GeneralCfgGeolocationCfg struct {
	GeolocationCfg GeolocationCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg"`
}

// MewlcConfig  represents the MEWLC configuration response
type MewlcConfig struct {
	MewlcConfigData MewlcConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:mewlc-config"`
}

// CacConfig  represents the CAC configuration response
type CacConfig struct {
	CacConfigData CacConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:cac-config"`
}

// Mfp  represents the Management Frame Protection configuration response
type Mfp struct {
	MfpData MfpData `json:"Cisco-IOS-XE-wireless-general-cfg:mfp"`
}

type FipsCfg struct {
	FipsCfgData FipsCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:fips-cfg"`
}

type WsaApClientEvent struct {
	WsaApClientEventData WsaApClientEventData `json:"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event"`
}

type SimL3InterfaceCacheData struct {
	SimL3InterfaceCacheDataInfo SimL3InterfaceCacheDataInfo `json:"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data"`
}

type WlcManagementData struct {
	WlcManagementDataInfo WlcManagementDataInfo `json:"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data"`
}

type Laginfo struct {
	LaginfoData LaginfoData `json:"Cisco-IOS-XE-wireless-general-cfg:laginfo"`
}

type MulticastConfig struct {
	MulticastConfigData MulticastConfigData `json:"Cisco-IOS-XE-wireless-general-cfg:multicast-config"`
}

type FeatureUsageCfg struct {
	FeatureUsageCfgData FeatureUsageCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg"`
}

type ThresholdWarnCfg struct {
	ThresholdWarnCfgData ThresholdWarnCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg"`
}

type ApLocRangingCfg struct {
	ApLocRangingCfgData ApLocRangingCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg"`
}

type GeolocationCfg struct {
	GeolocationCfgData GeolocationCfgData `json:"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg"`
}

// MewlcConfigData represents MEWLC configuration data structure.
type MewlcConfigData struct{}

type CacConfigData struct{}

type MfpData struct{}

type FipsCfgData struct{}

type WsaApClientEventData struct{}

type SimL3InterfaceCacheDataInfo struct {
	InterfaceName string `json:"interface-name"`
}

type WlcManagementDataInfo struct {
	PkiTrustpointName string `json:"pki-trustpoint-name"`
}

type LaginfoData struct{}

type MulticastConfigData struct {
	IsMdnsEnabled bool `json:"is-mdns-enabled"`
}

type FeatureUsageCfgData struct{}

type ThresholdWarnCfgData struct{}

type ApLocRangingCfgData struct{}

type GeolocationCfgData struct{}
