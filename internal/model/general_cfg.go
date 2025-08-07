// Package model contains generated response structures for the Cisco WNC API.
// This package is part of the three-layer architecture providing Generated Type separation.
package model

// General Configuration Response Types

type GeneralCfgResponse struct {
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

// MewlcConfigResponse represents the MEWLC configuration response
type MewlcConfigResponse struct {
	MewlcConfig MewlcConfig `json:"Cisco-IOS-XE-wireless-general-cfg:mewlc-config"`
}

// CacConfigResponse represents the CAC configuration response
type CacConfigResponse struct {
	CacConfig CacConfig `json:"Cisco-IOS-XE-wireless-general-cfg:cac-config"`
}

// MfpResponse represents the Management Frame Protection configuration response
type MfpResponse struct {
	Mfp Mfp `json:"Cisco-IOS-XE-wireless-general-cfg:mfp"`
}

type FipsCfgResponse struct {
	FipsCfg FipsCfg `json:"Cisco-IOS-XE-wireless-general-cfg:fips-cfg"`
}

type WsaApClientEventResponse struct {
	WsaApClientEvent WsaApClientEvent `json:"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event"`
}

type SimL3InterfaceCacheDataResponse struct {
	SimL3InterfaceCacheData SimL3InterfaceCacheData `json:"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data"`
}

type WlcManagementDataResponse struct {
	WlcManagementData WlcManagementData `json:"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data"`
}

type LaginfoResponse struct {
	Laginfo Laginfo `json:"Cisco-IOS-XE-wireless-general-cfg:laginfo"`
}

type MulticastConfigResponse struct {
	MulticastConfig MulticastConfig `json:"Cisco-IOS-XE-wireless-general-cfg:multicast-config"`
}

type FeatureUsageCfgResponse struct {
	FeatureUsageCfg FeatureUsageCfg `json:"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg"`
}

type ThresholdWarnCfgResponse struct {
	ThresholdWarnCfg ThresholdWarnCfg `json:"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg"`
}

type ApLocRangingCfgResponse struct {
	ApLocRangingCfg ApLocRangingCfg `json:"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg"`
}

type GeolocationCfgResponse struct {
	GeolocationCfg GeolocationCfg `json:"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg"`
}

// General Configuration Supporting Types

type MewlcConfig struct{}

type CacConfig struct{}

type Mfp struct{}

type FipsCfg struct{}

type WsaApClientEvent struct{}

type SimL3InterfaceCacheData struct {
	InterfaceName string `json:"interface-name"`
}

type WlcManagementData struct {
	PkiTrustpointName string `json:"pki-trustpoint-name"`
}

type Laginfo struct{}

type MulticastConfig struct {
	IsMdnsEnabled bool `json:"is-mdns-enabled"`
}

type FeatureUsageCfg struct{}

type ThresholdWarnCfg struct{}

type ApLocRangingCfg struct{}

type GeolocationCfg struct{}
