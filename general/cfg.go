// Package general provides general configuration functionality for the Cisco Wireless Network Controller API.
package general

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// GeneralCfgBasePath defines the base path for general configuration endpoints
	GeneralCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data"
	// GeneralCfgEndpoint retrieves complete general configuration data
	GeneralCfgEndpoint = GeneralCfgBasePath
	// MewlcConfigEndpoint retrieves MEWLC configuration
	MewlcConfigEndpoint = GeneralCfgBasePath + "/mewlc-config"
	// CacConfigEndpoint retrieves CAC configuration
	CacConfigEndpoint = GeneralCfgBasePath + "/cac-config"
	// MfpEndpoint retrieves Management Frame Protection configuration
	MfpEndpoint = GeneralCfgBasePath + "/mfp"
	// FipsCfgEndpoint retrieves FIPS configuration
	FipsCfgEndpoint = GeneralCfgBasePath + "/fips-cfg"
	// WsaApClientEventEndpoint retrieves WSA AP client event configuration
	WsaApClientEventEndpoint = GeneralCfgBasePath + "/wsa-ap-client-event"
	// SimL3InterfaceCacheDataEndpoint retrieves SIM L3 interface cache data
	SimL3InterfaceCacheDataEndpoint = GeneralCfgBasePath + "/sim-l3-interface-cache-data"
	// WlcManagementDataEndpoint retrieves WLC management data
	WlcManagementDataEndpoint = GeneralCfgBasePath + "/wlc-management-data"
	// LaginfoEndpoint retrieves LAG information
	LaginfoEndpoint = GeneralCfgBasePath + "/laginfo"
	// MulticastConfigEndpoint retrieves multicast configuration
	MulticastConfigEndpoint = GeneralCfgBasePath + "/multicast-config"
	// FeatureUsageCfgEndpoint retrieves feature usage configuration
	FeatureUsageCfgEndpoint = GeneralCfgBasePath + "/feature-usage-cfg"
	// ThresholdWarnCfgEndpoint retrieves threshold warning configuration
	ThresholdWarnCfgEndpoint = GeneralCfgBasePath + "/threshold-warn-cfg"
	// ApLocRangingCfgEndpoint retrieves AP location ranging configuration
	ApLocRangingCfgEndpoint = GeneralCfgBasePath + "/ap-loc-ranging-cfg"
	// GeolocationCfgEndpoint retrieves geolocation configuration
	GeolocationCfgEndpoint = GeneralCfgBasePath + "/geolocation-cfg"
)

// GeneralCfgResponse represents the complete general configuration response
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

// GetGeneralCfg retrieves general configuration with context support
func GetGeneralCfg(client *wnc.Client, ctx context.Context) (*GeneralCfgResponse, error) {
	var data GeneralCfgResponse
	if err := client.SendAPIRequest(ctx, GeneralCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralMewlcConfig(client *wnc.Client, ctx context.Context) (*MewlcConfigResponse, error) {
	var data MewlcConfigResponse
	if err := client.SendAPIRequest(ctx, MewlcConfigEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralCacConfig(client *wnc.Client, ctx context.Context) (*CacConfigResponse, error) {
	var data CacConfigResponse
	if err := client.SendAPIRequest(ctx, CacConfigEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralMfp(client *wnc.Client, ctx context.Context) (*MfpResponse, error) {
	var data MfpResponse
	if err := client.SendAPIRequest(ctx, MfpEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralFipsCfg(client *wnc.Client, ctx context.Context) (*FipsCfgResponse, error) {
	var data FipsCfgResponse
	if err := client.SendAPIRequest(ctx, FipsCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralWsaApClientEvent(client *wnc.Client, ctx context.Context) (*WsaApClientEventResponse, error) {
	var data WsaApClientEventResponse
	if err := client.SendAPIRequest(ctx, WsaApClientEventEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralSimL3InterfaceCacheData(client *wnc.Client, ctx context.Context) (*SimL3InterfaceCacheDataResponse, error) {
	var data SimL3InterfaceCacheDataResponse
	if err := client.SendAPIRequest(ctx, SimL3InterfaceCacheDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralWlcManagementData(client *wnc.Client, ctx context.Context) (*WlcManagementDataResponse, error) {
	var data WlcManagementDataResponse
	if err := client.SendAPIRequest(ctx, WlcManagementDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralLaginfo(client *wnc.Client, ctx context.Context) (*LaginfoResponse, error) {
	var data LaginfoResponse
	if err := client.SendAPIRequest(ctx, LaginfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralMulticastConfig(client *wnc.Client, ctx context.Context) (*MulticastConfigResponse, error) {
	var data MulticastConfigResponse
	if err := client.SendAPIRequest(ctx, MulticastConfigEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralFeatureUsageCfg(client *wnc.Client, ctx context.Context) (*FeatureUsageCfgResponse, error) {
	var data FeatureUsageCfgResponse
	if err := client.SendAPIRequest(ctx, FeatureUsageCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralThresholdWarnCfg(client *wnc.Client, ctx context.Context) (*ThresholdWarnCfgResponse, error) {
	var data ThresholdWarnCfgResponse
	if err := client.SendAPIRequest(ctx, ThresholdWarnCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralApLocRangingCfg(client *wnc.Client, ctx context.Context) (*ApLocRangingCfgResponse, error) {
	var data ApLocRangingCfgResponse
	if err := client.SendAPIRequest(ctx, ApLocRangingCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetGeneralGeolocationCfg(client *wnc.Client, ctx context.Context) (*GeolocationCfgResponse, error) {
	var data GeolocationCfgResponse
	if err := client.SendAPIRequest(ctx, GeolocationCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
