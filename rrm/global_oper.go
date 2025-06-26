// Package rrm provides Radio Resource Management global operational data functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"

	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// RrmGlobalOperBasePath defines the base path for RRM global operational data endpoints.
	RrmGlobalOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"
	// RrmGlobalOperEndpoint defines the endpoint for RRM global operational data.
	RrmGlobalOperEndpoint = RrmGlobalOperBasePath
	// RrmOneShotCountersEndpoint defines the endpoint for RRM one-shot counters.
	RrmOneShotCountersEndpoint = RrmGlobalOperBasePath + "/rrm-one-shot-counters"
	// RrmChannelParamsEndpoint defines the endpoint for RRM channel parameters.
	RrmChannelParamsEndpoint = RrmGlobalOperBasePath + "/rrm-channel-params"
	// SpectrumAqWorstTableEndpoint defines the endpoint for spectrum air quality worst table.
	SpectrumAqWorstTableEndpoint = RrmGlobalOperBasePath + "/spectrum-aq-worst-table"
	// RadioOperData24GEndpoint defines the endpoint for 2.4GHz radio operational data.
	RadioOperData24GEndpoint = RrmGlobalOperBasePath + "/radio-oper-data-24g"
	// RadioOperData5GEndpoint defines the endpoint for 5GHz radio operational data.
	RadioOperData5GEndpoint = RrmGlobalOperBasePath + "/radio-oper-data-5g"
	// RadioOperData6GEndpoint defines the endpoint for 6GHz radio operational data.
	RadioOperData6GEndpoint = RrmGlobalOperBasePath + "/radio-oper-data-6g"
	// SpectrumBandConfigDataEndpoint defines the endpoint for spectrum band configuration data.
	SpectrumBandConfigDataEndpoint = RrmGlobalOperBasePath + "/spectrum-band-config-data"
	// RadioOperDataDualbandEndpoint defines the endpoint for dual-band radio operational data.
	RadioOperDataDualbandEndpoint = RrmGlobalOperBasePath + "/radio-oper-data-dualband"
	// RrmClientDataEndpoint defines the endpoint for RRM client data.
	RrmClientDataEndpoint = RrmGlobalOperBasePath + "/rrm-client-data"
	// RrmFraStatsEndpoint defines the endpoint for RRM FRA statistics.
	RrmFraStatsEndpoint = RrmGlobalOperBasePath + "/rrm-fra-stats"
	// RrmCoverageEndpoint defines the endpoint for RRM coverage data.
	RrmCoverageEndpoint = RrmGlobalOperBasePath + "/rrm-coverage"
)

// RrmGlobalOperResponse represents the response structure for RRM global operational data.
type RrmGlobalOperResponse struct {
	CiscoIOSXEWirelessRrmGlobalOperData struct {
		RrmOneShotCounters     []RrmOneShotCounters     `json:"rrm-one-shot-counters"`
		RrmChannelParams       []RrmChannelParams       `json:"rrm-channel-params"`
		SpectrumAqWorstTable   []SpectrumAqWorstTable   `json:"spectrum-aq-worst-table"`
		RadioOperData24G       []RadioOperData24G       `json:"radio-oper-data-24g"`
		RadioOperData5G        []RadioOperData5G        `json:"radio-oper-data-5g"`
		RadioOperData6G        []RadioOperData6G        `json:"radio-oper-data-6g"`
		SpectrumBandConfigData []SpectrumBandConfigData `json:"spectrum-band-config-data"`
		RadioOperDataDualband  []RadioOperDataDualband  `json:"radio-oper-data-dualband"`
		RrmClientData          []RrmClientData          `json:"rrm-client-data"`
		RrmFraStats            RrmGlobalOperRrmFraStats `json:"rrm-fra-stats"`
		RrmCoverage            []RrmCoverage            `json:"rrm-coverage"`
	} `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"`
}

// RrmOneShotCountersResponse represents the response structure for RRM one-shot counters.
type RrmOneShotCountersResponse struct {
	RrmOneShotCounters []RrmOneShotCounters `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-one-shot-counters"`
}

// RrmChannelParamsResponse represents the response structure for RRM channel parameters.
type RrmChannelParamsResponse struct {
	RrmChannelParams []RrmChannelParams `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-channel-params"`
}

// SpectrumAqWorstTableResponse represents the response structure for spectrum air quality worst table.
type SpectrumAqWorstTableResponse struct {
	SpectrumAqWorstTable []SpectrumAqWorstTable `json:"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-aq-worst-table"`
}

type RadioOperData24GResponse struct {
	RadioOperData24G []RadioOperData24G `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-24g"`
}

type RadioOperData5GResponse struct {
	RadioOperData5G []RadioOperData5G `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-5g"`
}

type SpectrumBandConfigDataResponse struct {
	SpectrumBandConfigData []SpectrumBandConfigData `json:"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-band-config-data"`
}

type RadioOperDataDualbandResponse struct {
	RadioOperDataDualband []RadioOperDataDualband `json:"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-dualband"`
}

type RrmClientDataResponse struct {
	RrmClientData []RrmClientData `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-client-data"`
}

type RrmFraStatsResponse struct {
	RrmFraStats RrmGlobalOperRrmFraStats `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-fra-stats"`
}

type RrmCoverageResponse struct {
	RrmCoverage []RrmCoverage `json:"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-coverage"`
}

type RrmOneShotCounters struct {
	PhyType      string `json:"phy-type"`
	PowerCounter int    `json:"power-counter"`
}

type RrmChannelParams struct {
	PhyType        string `json:"phy-type"`
	MinDwell       int    `json:"min-dwell"`
	AvgDwell       int    `json:"avg-dwell"`
	MaxDwell       int    `json:"max-dwell"`
	MinRssi        int    `json:"min-rssi"`
	MaxRssi        int    `json:"max-rssi"`
	AvgRssi        int    `json:"avg-rssi"`
	ChannelCounter int    `json:"channel-counter"`
}

type SpectrumAqWorstTable struct {
	BandID               int    `json:"band-id"`
	DetectingApName      string `json:"detecting-ap-name"`
	ChannelNum           int    `json:"channel-num"`
	MinAqi               int    `json:"min-aqi"`
	Aqi                  int    `json:"aqi"`
	TotalIntfDeviceCount int    `json:"total-intf-device-count"`
	WtpCaSiCapable       string `json:"wtp-ca-si-capable"`
	ScanRadioType        string `json:"scan-radio-type"`
}

type RadioOperData24G struct {
	WtpMac          string    `json:"wtp-mac"`
	RadioSlotID     int       `json:"radio-slot-id"`
	ApMac           string    `json:"ap-mac"`
	SlotID          int       `json:"slot-id"`
	Name            string    `json:"name"`
	SpectrumCapable []any     `json:"spectrum-capable"`
	NumSlots        int       `json:"num-slots"`
	MeshRadioRole   string    `json:"mesh-radio-role"`
	ApUpTime        time.Time `json:"ap-up-time"`
	CapwapUpTime    time.Time `json:"capwap-up-time"`
}

type RadioOperData5G struct {
	WtpMac          string    `json:"wtp-mac"`
	RadioSlotID     int       `json:"radio-slot-id"`
	ApMac           string    `json:"ap-mac"`
	SlotID          int       `json:"slot-id"`
	Name            string    `json:"name"`
	SpectrumCapable []any     `json:"spectrum-capable"`
	NumSlots        int       `json:"num-slots"`
	MeshRadioRole   string    `json:"mesh-radio-role"`
	ApUpTime        time.Time `json:"ap-up-time"`
	CapwapUpTime    time.Time `json:"capwap-up-time"`
}

type RadioOperData6G struct {
	WtpMac          string    `json:"wtp-mac"`
	RadioSlotID     int       `json:"radio-slot-id"`
	ApMac           string    `json:"ap-mac"`
	SlotID          int       `json:"slot-id"`
	Name            string    `json:"name"`
	SpectrumCapable []any     `json:"spectrum-capable"`
	NumSlots        int       `json:"num-slots"`
	MeshRadioRole   string    `json:"mesh-radio-role"`
	ApUpTime        time.Time `json:"ap-up-time"`
	CapwapUpTime    time.Time `json:"capwap-up-time"`
}

type SpectrumBandConfigData struct {
	ApMac              string `json:"ap-mac"`
	SpectrumBandConfig []struct {
		BandID             string `json:"band-id"`
		SpectrumAdminState bool   `json:"spectrum-admin-state"`
	} `json:"spectrum-band-config"`
}

type RadioOperDataDualband struct {
	WtpMac          string    `json:"wtp-mac"`
	RadioSlotID     int       `json:"radio-slot-id"`
	ApMac           string    `json:"ap-mac"`
	SlotID          int       `json:"slot-id"`
	Name            string    `json:"name"`
	SpectrumCapable []any     `json:"spectrum-capable"`
	NumSlots        int       `json:"num-slots"`
	MeshRadioRole   string    `json:"mesh-radio-role"`
	ApUpTime        time.Time `json:"ap-up-time"`
	CapwapUpTime    time.Time `json:"capwap-up-time"`
}

type RrmClientData struct {
	PhyType         string    `json:"phy-type"`
	LastChdRun      time.Time `json:"last-chd-run"`
	Disassociations int       `json:"disassociations"`
	Rejections      int       `json:"rejections"`
}

type RrmGlobalOperRrmFraStats struct {
	DualBandMonitorTo24Ghz int `json:"dual-band-monitor-to-24ghz"`
	DualBandMonitorTo5Ghz  int `json:"dual-band-monitor-to-5ghz"`
	DualBand24GhzTo5Ghz    int `json:"dual-band-24ghz-to-5ghz"`
	DualBand24GhzToMonitor int `json:"dual-band-24ghz-to-monitor"`
	DualBand5GhzTo24Ghz    int `json:"dual-band-5ghz-to-24ghz"`
	DualBand5GhzToMonitor  int `json:"dual-band-5ghz-to-monitor"`
	SecRadioMonitorTo5Ghz  int `json:"sec-radio-monitor-to-5ghz"`
	SecRadio5GhzToMonitor  int `json:"sec-radio-5ghz-to-monitor"`
	DualBand6GhzTo5Ghz     int `json:"dual-band-6ghz-to-5ghz"`
	DualBand5GhzTo6Ghz     int `json:"dual-band-5ghz-to-6ghz"`
}

type RrmCoverage struct {
	WtpMac            string `json:"wtp-mac"`
	RadioSlotID       int    `json:"radio-slot-id"`
	FailedClientCount int    `json:"failed-client-count"`
	SnrInfo           []struct {
		SNR        int `json:"snr"`
		NumClients int `json:"num-clients"`
	} `json:"snr-info"`
	RssiInfo []struct {
		RSSI       int `json:"rssi"`
		NumClients int `json:"num-clients"`
	} `json:"rssi-info"`
}

func GetRrmGlobalOper(client *wnc.Client, ctx context.Context) (*RrmGlobalOperResponse, error) {
	var data RrmGlobalOperResponse
	err := client.SendAPIRequest(ctx, RrmGlobalOperEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalOneShotCounters(client *wnc.Client, ctx context.Context) (*RrmOneShotCountersResponse, error) {
	var data RrmOneShotCountersResponse
	err := client.SendAPIRequest(ctx, RrmOneShotCountersEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalChannelParams(client *wnc.Client, ctx context.Context) (*RrmChannelParamsResponse, error) {
	var data RrmChannelParamsResponse
	err := client.SendAPIRequest(ctx, RrmChannelParamsEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalSpectrumAqWorstTable(client *wnc.Client, ctx context.Context) (*SpectrumAqWorstTableResponse, error) {
	var data SpectrumAqWorstTableResponse
	err := client.SendAPIRequest(ctx, SpectrumAqWorstTableEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalRadioOperData24G(client *wnc.Client, ctx context.Context) (*RadioOperData24GResponse, error) {
	var data RadioOperData24GResponse
	err := client.SendAPIRequest(ctx, RadioOperData24GEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalRadioOperData5G(client *wnc.Client, ctx context.Context) (*RadioOperData5GResponse, error) {
	var data RadioOperData5GResponse
	err := client.SendAPIRequest(ctx, RadioOperData5GEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalRadioOperData6G(client *wnc.Client, ctx context.Context) (*RadioOperData5GResponse, error) {
	var data RadioOperData5GResponse
	err := client.SendAPIRequest(ctx, RadioOperData6GEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalSpectrumBandConfigData(client *wnc.Client, ctx context.Context) (*SpectrumBandConfigDataResponse, error) {
	var data SpectrumBandConfigDataResponse
	err := client.SendAPIRequest(ctx, SpectrumBandConfigDataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalRadioOperDataDualband(client *wnc.Client, ctx context.Context) (*RadioOperDataDualbandResponse, error) {
	var data RadioOperDataDualbandResponse
	err := client.SendAPIRequest(ctx, RadioOperDataDualbandEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalClientData(client *wnc.Client, ctx context.Context) (*RrmClientDataResponse, error) {
	var data RrmClientDataResponse
	err := client.SendAPIRequest(ctx, RrmClientDataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalFraStats(client *wnc.Client, ctx context.Context) (*RrmFraStatsResponse, error) {
	var data RrmFraStatsResponse
	err := client.SendAPIRequest(ctx, RrmFraStatsEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmGlobalCoverage(client *wnc.Client, ctx context.Context) (*RrmCoverageResponse, error) {
	var data RrmCoverageResponse
	err := client.SendAPIRequest(ctx, RrmCoverageEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
