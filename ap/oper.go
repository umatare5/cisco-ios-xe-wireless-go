// Package ap provides access point operational data functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"fmt"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

// Access Point operational data API endpoints
const (
	// ApOperBasePath is the base path for access point operational data endpoints
	ApOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"
	// ApOperEndpoint retrieves complete access point operational data
	ApOperEndpoint = ApOperBasePath
	// ApRadioNeighborEndpoint retrieves access point radio neighbor information
	ApRadioNeighborEndpoint = ApOperBasePath + "/ap-radio-neighbor"
	// RadioOperDataEndpoint retrieves radio operational data for access points
	RadioOperDataEndpoint = ApOperBasePath + "/radio-oper-data"
	// RadioResetStatsEndpoint retrieves radio reset statistics
	RadioResetStatsEndpoint = ApOperBasePath + "/radio-reset-stats"
	// QosClientDataEndpoint retrieves QoS client data information
	QosClientDataEndpoint = ApOperBasePath + "/qos-client-data"
	// CapwapDataEndpoint retrieves CAPWAP data for access points
	CapwapDataEndpoint = ApOperBasePath + "/capwap-data"
	// ApNameMacMapEndpoint retrieves AP name to MAC address mapping
	ApNameMacMapEndpoint = ApOperBasePath + "/ap-name-mac-map"
	// WtpSlotWlanStatsEndpoint retrieves WTP slot WLAN statistics
	WtpSlotWlanStatsEndpoint = ApOperBasePath + "/wtp-slot-wlan-stats"
	// EthernetMacWtpMacMapEndpoint retrieves Ethernet MAC to WTP MAC mapping
	EthernetMacWtpMacMapEndpoint = ApOperBasePath + "/ethernet-mac-wtp-mac-map"
	// RadioOperStatsEndpoint retrieves radio operational statistics
	RadioOperStatsEndpoint = ApOperBasePath + "/radio-oper-stats"
	// EthernetIfStatsEndpoint retrieves Ethernet interface statistics
	EthernetIfStatsEndpoint = ApOperBasePath + "/ethernet-if-stats"
	// EwlcWncdStatsEndpoint retrieves EWLC WNCD statistics
	EwlcWncdStatsEndpoint = ApOperBasePath + "/ewlc-wncd-stats"
	// ApIoxOperDataEndpoint retrieves AP IOx operational data
	ApIoxOperDataEndpoint = ApOperBasePath + "/ap-iox-oper-data"
	// QosGlobalStatsEndpoint retrieves global QoS statistics
	QosGlobalStatsEndpoint = ApOperBasePath + "/qos-global-stats"
	// OperDataEndpoint retrieves general operational data
	OperDataEndpoint = ApOperBasePath + "/oper-data"
	// RlanOperEndpoint retrieves RLAN operational data
	RlanOperEndpoint = ApOperBasePath + "/rlan-oper"
	// EwlcMewlcPredownloadRecEndpoint retrieves EWLC MEWLC predownload records
	EwlcMewlcPredownloadRecEndpoint = ApOperBasePath + "/ewlc-mewlc-predownload-rec"
	// CdpCacheDataEndpoint retrieves CDP cache data
	CdpCacheDataEndpoint = ApOperBasePath + "/cdp-cache-data"
	// LldpNeighEndpoint retrieves LLDP neighbor information
	LldpNeighEndpoint = ApOperBasePath + "/lldp-neigh"
	// TpCertInfoEndpoint retrieves trustpoint certificate information
	TpCertInfoEndpoint = ApOperBasePath + "/tp-cert-info"
	// DiscDataEndpoint retrieves discovery data
	DiscDataEndpoint = ApOperBasePath + "/disc-data"
	// CapwapPktsEndpoint retrieves CAPWAP packet statistics
	CapwapPktsEndpoint = ApOperBasePath + "/capwap-pkts"
	// CountryOperEndpoint retrieves country operational data
	CountryOperEndpoint = ApOperBasePath + "/country-oper"
	// SuppCountryOperEndpoint retrieves supported country operational data
	SuppCountryOperEndpoint = ApOperBasePath + "/supp-country-oper"
	// ApNhGlobalDataEndpoint retrieves AP NH global data
	ApNhGlobalDataEndpoint = ApOperBasePath + "/ap-nh-global-data"
	// ApImagePrepareLocationEndpoint retrieves AP image prepare location data
	ApImagePrepareLocationEndpoint = ApOperBasePath + "/ap-image-prepare-location"
	// ApImageActiveLocationEndpoint retrieves AP image active location data
	ApImageActiveLocationEndpoint = ApOperBasePath + "/ap-image-active-location"
)

// ApOperResponse represents the complete access point operational data response.
// This structure contains comprehensive operational information for all access points
// in the wireless network, including radio data, statistics, neighbor information,
// CAPWAP details, QoS metrics, and various operational counters.
type ApOperResponse struct {
	CiscoIOSXEWirelessAccessPointOperAccessPointOperData struct {
		ApRadioNeighbor         []ApRadioNeighbor        `json:"ap-radio-neighbor"`
		RadioOperData           []RadioOperData          `json:"radio-oper-data"`
		RadioResetStats         []RadioResetStats        `json:"radio-reset-stats"`
		QosClientData           []QosClientData          `json:"qos-client-data"`
		CapwapData              []CapwapData             `json:"capwap-data"`
		ApNameMacMap            []ApNameMacMap           `json:"ap-name-mac-map"`
		WtpSlotWlanStats        []WtpSlotWlanStats       `json:"wtp-slot-wlan-stats"`
		EthernetMacWtpMacMap    []EthernetMacWtpMacMap   `json:"ethernet-mac-wtp-mac-map"`
		RadioOperStats          []RadioOperStats         `json:"radio-oper-stats"`
		EthernetIfStats         []EthernetIfStats        `json:"ethernet-if-stats"`
		EwlcWncdStats           EwlcWncdStats            `json:"ewlc-wncd-stats"`
		ApIoxOperData           []ApIoxOperData          `json:"ap-iox-oper-data"`
		QosGlobalStats          QosGlobalStats           `json:"qos-global-stats"`
		OperData                []ApOperData             `json:"oper-data"`
		RlanOper                []RlanOper               `json:"rlan-oper"`
		EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec  `json:"ewlc-mewlc-predownload-rec"`
		CdpCacheData            []CdpCacheData           `json:"cdp-cache-data"`
		LldpNeigh               []LldpNeigh              `json:"lldp-neigh"`
		TpCertInfo              TpCertInfo               `json:"tp-cert-info"`
		DiscData                []DiscData               `json:"disc-data"`
		CapwapPkts              []CapwapPkts             `json:"capwap-pkts"`
		CountryOper             []CountryOper            `json:"country-oper"`
		SuppCountryOper         []SuppCountryOper        `json:"supp-country-oper"`
		ApNhGlobalData          ApNhGlobalData           `json:"ap-nh-global-data"`
		ApImagePrepareLocation  []ApImagePrepareLocation `json:"ap-image-prepare-location"`
		ApImageActiveLocation   []ApImageActiveLocation  `json:"ap-image-active-location"`
	} `json:"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"`
}

// ApOperApRadioNeighborResponse represents the access point radio neighbor response
type ApOperApRadioNeighborResponse struct {
	ApRadioNeighbor []ApRadioNeighbor `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor"`
}

// ApOperRadioOperDataResponse represents the radio operational data response
type ApOperRadioOperDataResponse struct {
	RadioOperData []RadioOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data"`
}

// ApOperRadioResetStatsResponse represents the radio reset statistics response
type ApOperRadioResetStatsResponse struct {
	RadioResetStats []RadioResetStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-reset-stats"`
}

// ApOperQosClientDataResponse represents the QoS client data response
type ApOperQosClientDataResponse struct {
	QosClientData []QosClientData `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data"`
}

// ApOperCapwapDataResponse represents the CAPWAP data response
type ApOperCapwapDataResponse struct {
	CapwapData []CapwapData `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
}

// ApOperApNameMacMapResponse represents the AP name to MAC mapping response
type ApOperApNameMacMapResponse struct {
	ApNameMacMap []ApNameMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
}

// ApOperWtpSlotWlanStatsResponse represents the WTP slot WLAN statistics response
type ApOperWtpSlotWlanStatsResponse struct {
	WtpSlotWlanStats []WtpSlotWlanStats `json:"Cisco-IOS-XE-wireless-access-point-oper:wtp-slot-wlan-stats"`
}

// ApOperEthernetMacWtpMacMapResponse represents the Ethernet MAC to WTP MAC mapping response
type ApOperEthernetMacWtpMacMapResponse struct {
	EthernetMacWtpMacMap []EthernetMacWtpMacMap `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map"`
}

// ApOperRadioOperStatsResponse represents the radio operational statistics response
type ApOperRadioOperStatsResponse struct {
	RadioOperStats []RadioOperStats `json:"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-stats"`
}

// ApOperEthernetIfStatsResponse represents the Ethernet interface statistics response
type ApOperEthernetIfStatsResponse struct {
	EthernetIfStats []EthernetIfStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats"`
}

// ApOperEwlcWncdStatsResponse represents the EWLC WNCD statistics response
type ApOperEwlcWncdStatsResponse struct {
	EwlcWncdStats EwlcWncdStats `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-wncd-stats"`
}

// ApOperApIoxOperDataResponse represents the AP IOx operational data response
type ApOperApIoxOperDataResponse struct {
	ApIoxOperData []ApIoxOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-iox-oper-data"`
}

// ApOperQosGlobalStatsResponse represents the global QoS statistics response
type ApOperQosGlobalStatsResponse struct {
	QosGlobalStats QosGlobalStats `json:"Cisco-IOS-XE-wireless-access-point-oper:qos-global-stats"`
}

// ApOperOperDataResponse represents the general operational data response
type ApOperOperDataResponse struct {
	OperData []ApOperData `json:"Cisco-IOS-XE-wireless-access-point-oper:oper-data"`
}

// ApOperRlanOperResponse represents the RLAN operational data response
type ApOperRlanOperResponse struct {
	RlanOper []RlanOper `json:"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper"`
}

// ApOperEwlcMewlcPredownloadRecResponse represents the EWLC MEWLC predownload record response
type ApOperEwlcMewlcPredownloadRecResponse struct {
	EwlcMewlcPredownloadRec EwlcMewlcPredownloadRec `json:"Cisco-IOS-XE-wireless-access-point-oper:ewlc-mewlc-predownload-rec"`
}

// ApOperCdpCacheDataResponse represents the CDP cache data response
type ApOperCdpCacheDataResponse struct {
	CdpCacheData []CdpCacheData `json:"Cisco-IOS-XE-wireless-access-point-oper:cdp-cache-data"`
}

// ApOperLldpNeighResponse represents the LLDP neighbor response
type ApOperLldpNeighResponse struct {
	LldpNeigh []LldpNeigh `json:"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh"`
}

// ApOperTpCertInfoResponse represents the trustpoint certificate information response
type ApOperTpCertInfoResponse struct {
	TpCertInfo TpCertInfo `json:"Cisco-IOS-XE-wireless-access-point-oper:tp-cert-info"`
}

// ApOperDiscDataResponse represents the discovery data response
type ApOperDiscDataResponse struct {
	DiscData []DiscData `json:"Cisco-IOS-XE-wireless-access-point-oper:disc-data"`
}

// ApOperCapwapPktsResponse represents the CAPWAP packet statistics response
type ApOperCapwapPktsResponse struct {
	CapwapPkts []CapwapPkts `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-pkts"`
}

// ApOperCountryOperResponse represents the country operational data response
type ApOperCountryOperResponse struct {
	CountryOper []CountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:country-oper"`
}

// ApOperSuppCountryOperResponse represents the supported country operational data response
type ApOperSuppCountryOperResponse struct {
	SuppCountryOper []SuppCountryOper `json:"Cisco-IOS-XE-wireless-access-point-oper:supp-country-oper"`
}

// ApOperApNhGlobalDataResponse represents the AP NH global data response
type ApOperApNhGlobalDataResponse struct {
	ApNhGlobalData ApNhGlobalData `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-nh-global-data"`
}

// ApOperApImagePrepareLocationResponse represents the AP image prepare location response
type ApOperApImagePrepareLocationResponse struct {
	ApImagePrepareLocation []ApImagePrepareLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-prepare-location"`
}

// ApOperApImageActiveLocationResponse represents the AP image active location response
type ApOperApImageActiveLocationResponse struct {
	ApImageActiveLocation []ApImageActiveLocation `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-image-active-location"`
}

// ApRadioNeighbor represents access point radio neighbor information
type ApRadioNeighbor struct {
	ApMac          string    `json:"ap-mac"`           // MAC address of the neighboring access point
	SlotID         int       `json:"slot-id"`          // Radio slot identifier
	Bssid          string    `json:"bssid"`            // Basic Service Set Identifier
	Ssid           string    `json:"ssid"`             // Service Set Identifier
	Rssi           int       `json:"rssi"`             // Received Signal Strength Indicator in dBm
	Channel        int       `json:"channel"`          // Radio channel number
	PrimaryChannel int       `json:"primary-channel"`  // Primary channel number
	LastUpdateRcvd time.Time `json:"last-update-rcvd"` // Timestamp of last update received
}

// RadioOperData represents radio operational data for an access point
type RadioOperData struct {
	WtpMac            string `json:"wtp-mac"`                       // Wireless Termination Point MAC address
	RadioSlotID       int    `json:"radio-slot-id"`                 // Radio slot identifier
	SlotID            int    `json:"slot-id,omitempty"`             // Slot identifier (optional)
	RadioType         string `json:"radio-type"`                    // Type of radio (e.g., 802.11ac, 802.11ax)
	AdminState        string `json:"admin-state,omitempty"`         // Administrative state of the radio
	OperState         string `json:"oper-state,omitempty"`          // Operational state of the radio
	RadioMode         string `json:"radio-mode,omitempty"`          // Current radio mode
	RadioSubMode      string `json:"radio-sub-mode,omitempty"`      // Radio sub-mode configuration
	RadioSubtype      string `json:"radio-subtype,omitempty"`       // Radio subtype information
	CurrentBandID     int    `json:"current-band-id,omitempty"`     // Current operating band identifier
	CurrentActiveBand string `json:"current-active-band,omitempty"` // Current active frequency band
	XorRadioMode      string `json:"xor-radio-mode,omitempty"`      // XOR radio mode setting
	StationCfg        struct {
		CfgData struct {
			StationCfgConfigType string `json:"station-cfg-config-type"`
			MediumOccupancyLimit int    `json:"medium-occupancy-limit"`
			CfpPeriod            int    `json:"cfp-period"`
			CfpMaxDuration       int    `json:"cfp-max-duration"`
			Bssid                string `json:"bssid"`
			BeaconPeriod         int    `json:"beacon-period"`
			CountryString        string `json:"country-string"`
		} `json:"cfg-data"`
	} `json:"station-cfg,omitempty"`
	MultiDomainCap struct {
		MultiDomainCapCfgData struct {
			FirstChanNum    int `json:"first-chan-num"`
			NumChannels     int `json:"num-channels"`
			MaxTxPowerLevel int `json:"max-tx-power-level"`
		} `json:"cfg-data"`
	} `json:"multi-domain-cap,omitempty"`
	PhyHtCfg struct {
		PhyHtCfgCfgData struct {
			HtEnable               int    `json:"ht-enable"`
			PhyHtCfgConfigType     string `json:"phy-ht-cfg-config-type"`
			CurrFreq               int    `json:"curr-freq"`
			ChanWidth              int    `json:"chan-width"`
			ExtChan                int    `json:"ext-chan"`
			VhtEnable              bool   `json:"vht-enable"`
			LegTxBfEnabled         int    `json:"leg-tx-bf-enabled"`
			RrmChannelChangeReason string `json:"rrm-channel-change-reason"`
			FreqString             string `json:"freq-string"`
		} `json:"cfg-data"`
	} `json:"phy-ht-cfg,omitempty"`
	PhyHtCap struct {
		Data struct {
			VhtCapable bool `json:"vht-capable"`
			HtCapable  bool `json:"ht-capable"`
		} `json:"data"`
	} `json:"phy-ht-cap,omitempty"`
	XorPhyHtCap struct {
		XorPhyHtCapData struct {
			VhtCapable bool `json:"vht-capable"`
			HtCapable  bool `json:"ht-capable"`
		} `json:"data"`
	} `json:"xor-phy-ht-cap,omitempty"`
	PhyHeCap struct {
		PhyHeCapData struct {
			HeEnabled              bool `json:"he-enabled"`
			HeCapable              bool `json:"he-capable"`
			HeSingleUserBeamformer int  `json:"he-single-user-beamformer"`
			HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`
			HeStbcMode             int  `json:"he-stbc-mode"`
			HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`
			HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`
		} `json:"data"`
	} `json:"phy-he-cap,omitempty"`
	RadioHeCapable bool `json:"radio-he-capable,omitempty"`
	ChanPwrInfo    struct {
		ChanPwrInfoData struct {
			AntennaGain    int `json:"antenna-gain"`
			IntAntennaGain int `json:"int-antenna-gain"`
			ExtAntennaGain int `json:"ext-antenna-gain"`
			ChanPwrList    struct {
				ChanPwr []struct {
					Chan int `json:"chan"`
				} `json:"chan-pwr"`
			} `json:"chan-pwr-list"`
		} `json:"data"`
	} `json:"chan-pwr-info,omitempty"`
	SnifferCfg struct {
		SnifferEnabled bool `json:"sniffer-enabled"`
	} `json:"sniffer-cfg,omitempty"`
	AntennaPid    string `json:"antenna-pid,omitempty"`
	RadioBandInfo []struct {
		BandID           int    `json:"band-id"`
		RegDomainCode    int    `json:"reg-domain-code"`
		RegulatoryDomain string `json:"regulatory-domain"`
		MacOperCfg       struct {
			MacOperCfgCfgData struct {
				MacOperationConfigType string `json:"mac-operation-config-type"`
				RtsThreshold           int    `json:"rts-threshold"`
				ShortRetryLimit        int    `json:"short-retry-limit"`
				LongRetryLimit         int    `json:"long-retry-limit"`
				FragThreshold          int    `json:"frag-threshold"`
				MaxTxLifeTime          int    `json:"max-tx-life-time"`
				MaxRxLifeTime          int    `json:"max-rx-life-time"`
			} `json:"cfg-data"`
		} `json:"mac-oper-cfg"`
		PhyTxPwrCfg struct {
			PhyTxPwrCfgCfgData struct {
				PhyTxPowerConfigType string `json:"phy-tx-power-config-type"`
				CurrentTxPowerLevel  int    `json:"current-tx-power-level"`
			} `json:"cfg-data"`
		} `json:"phy-tx-pwr-cfg"`
		PhyTxPwrLvlCfg struct {
			PhyTxPwrLvlCfgCfgData struct {
				NumSuppPowerLevels int `json:"num-supp-power-levels"`
				TxPowerLevel1      int `json:"tx-power-level-1"`
				TxPowerLevel2      int `json:"tx-power-level-2"`
				TxPowerLevel3      int `json:"tx-power-level-3"`
				TxPowerLevel4      int `json:"tx-power-level-4"`
				TxPowerLevel5      int `json:"tx-power-level-5"`
				TxPowerLevel6      int `json:"tx-power-level-6"`
				TxPowerLevel7      int `json:"tx-power-level-7"`
				TxPowerLevel8      int `json:"tx-power-level-8"`
				CurrTxPowerInDbm   int `json:"curr-tx-power-in-dbm"`
			} `json:"cfg-data"`
		} `json:"phy-tx-pwr-lvl-cfg"`
		AntennaCfg struct {
			AntennaCfgCfgData struct {
				DiversitySelection string `json:"diversity-selection"`
				AntennaMode        string `json:"antenna-mode"`
				NumOfAntennas      int    `json:"num-of-antennas"`
			} `json:"cfg-data"`
		} `json:"antenna-cfg"`
		Dot11AcChannelWidthCap int `json:"dot11ac-channel-width-cap"`
		Secondary80Channel     int `json:"secondary-80-channel"`
		SiaParams              struct {
			IsRptncPresent bool   `json:"is-rptnc-present"`
			IsDartPresent  bool   `json:"is-dart-present"`
			AntennaIfType  string `json:"antenna-if-type"`
			AntennaGain    int    `json:"antenna-gain"`
			Marlin4Present bool   `json:"marlin4-present"`
			DmServType     string `json:"dm-serv-type"`
		} `json:"sia-params"`
	} `json:"radio-band-info,omitempty"`
	VapOperConfig []struct {
		ApVapID         int    `json:"ap-vap-id"`
		WlanID          int    `json:"wlan-id"`
		BssidMac        string `json:"bssid-mac"`
		WtpMac          string `json:"wtp-mac"`
		WlanProfileName string `json:"wlan-profile-name"`
		Ssid            string `json:"ssid"`
	} `json:"vap-oper-config,omitempty"`
	RegDomainCheckStatus string    `json:"reg-domain-check-status,omitempty"`
	AntennaGain          int       `json:"antenna-gain,omitempty"`
	SlotAntennaType      string    `json:"slot-antenna-type,omitempty"`
	RadioEnableTime      time.Time `json:"radio-enable-time,omitempty"`
	Dot11NMcsRates       string    `json:"dot11n-mcs-rates,omitempty"`
	DualRadioModeCfg     struct {
		DualRadioMode    string `json:"dual-radio-mode"`
		DualRadioCapable string `json:"dual-radio-capable"`
		DualRadioModeOp  string `json:"dual-radio-mode-op"`
	} `json:"dual-radio-mode-cfg,omitempty"`
	RadioFraCapable string `json:"radio-fra-capable,omitempty"`
	BssColorCfg     struct {
		BssColorCapable    bool   `json:"bss-color-capable"`
		BssColor           int    `json:"bss-color"`
		BssColorConfigType string `json:"bss-color-config-type"`
	} `json:"bss-color-cfg,omitempty"`
	HighestThroughputProto string `json:"highest-throughput-proto,omitempty"`
	CacActive              bool   `json:"cac-active,omitempty"`
	ObssPdCapable          bool   `json:"obss-pd-capable,omitempty"`
	NdpCap                 string `json:"ndp-cap,omitempty"`
	NdpOnChannel           bool   `json:"ndp-on-channel,omitempty"`
	BeamSelection          string `json:"beam-selection,omitempty"`
	NumAntEnabled          int    `json:"num-ant-enabled,omitempty"`
	CurAntBitmap           string `json:"cur-ant-bitmap,omitempty"`
	SuppAntBitmap          string `json:"supp-ant-bitmap,omitempty"`
	Supp160MhzAntBitmap    string `json:"supp-160mhz-ant-bitmap,omitempty"`
	MeshBackhaul           bool   `json:"mesh-backhaul,omitempty"`
	MeshDesignatedDownlink bool   `json:"mesh-designated-downlink,omitempty"`
	MaxClientAllowed       int    `json:"max-client-allowed,omitempty"`
	ObssPdSrgCapable       bool   `json:"obss-pd-srg-capable,omitempty"`
	XorPhyHeCap            struct {
		XorPhyHeCapData struct {
			HeEnabled              bool `json:"he-enabled"`
			HeCapable              bool `json:"he-capable"`
			HeSingleUserBeamformer int  `json:"he-single-user-beamformer"`
			HeMultiUserBeamformer  int  `json:"he-multi-user-beamformer"`
			HeStbcMode             int  `json:"he-stbc-mode"`
			HeAmpduTidBitmap       int  `json:"he-ampdu-tid-bitmap"`
			HeCapTxRxMcsNss        int  `json:"he-cap-tx-rx-mcs-nss"`
		} `json:"data"`
	} `json:"xor-phy-he-cap,omitempty"`
	CoverageOverlapFactor int    `json:"coverage-overlap-factor,omitempty"`
	Ap6GhzPwrMode         string `json:"ap-6ghz-pwr-mode,omitempty"`
	Ap6GhzPwrModeCap      string `json:"ap-6ghz-pwr-mode-cap,omitempty"`
	AfcLicenseNeeded      bool   `json:"afc-license-needed,omitempty"`
	PushAfcRespDone       bool   `json:"push-afc-resp-done,omitempty"`
	AfcBelowTxmin         bool   `json:"afc-below-txmin,omitempty"`
	RadioSubband          string `json:"radio-subband,omitempty"`
}

// RadioResetStats represents radio reset statistics for troubleshooting
type RadioResetStats struct {
	ApMac       string `json:"ap-mac"`       // Access point MAC address
	RadioID     int    `json:"radio-id"`     // Radio identifier
	Cause       string `json:"cause"`        // Cause of the radio reset
	DetailCause string `json:"detail-cause"` // Detailed cause description
	Count       int    `json:"count"`        // Number of resets for this cause
}

// QosClientData represents Quality of Service data for a client
type QosClientData struct {
	ClientMac    string `json:"client-mac"` // Client MAC address
	AaaQosParams struct {
		AaaAvgdtus   int `json:"aaa-avgdtus"`   // AAA average downstream units
		AaaAvgrtdtus int `json:"aaa-avgrtdtus"` // AAA average round-trip downstream units
		AaaBstdtus   int `json:"aaa-bstdtus"`   // AAA burst downstream units
		AaaBstrtdtus int `json:"aaa-bstrtdtus"` // AAA burst round-trip downstream units
		AaaAvgdtds   int `json:"aaa-avgdtds"`   // AAA average downstream data
		AaaAvgrtdtds int `json:"aaa-avgrtdtds"` // AAA average round-trip downstream data
		AaaBstdtds   int `json:"aaa-bstdtds"`   // AAA burst downstream data
		AaaBstrtdtds int `json:"aaa-bstrtdtds"` // AAA burst round-trip downstream data
	} `json:"aaa-qos-params"` // AAA Quality of Service parameters
}

type StaticInfo struct {
	BoardData struct {
		WtpSerialNum string `json:"wtp-serial-num"`
		WtpEnetMac   string `json:"wtp-enet-mac"`
		ApSysInfo    struct {
			MemType string `json:"mem-type"`
			CPUType string `json:"cpu-type"`
			MemSize int    `json:"mem-size"`
		} `json:"ap-sys-info"`
	} `json:"board-data"`
	BoardDataOpt struct {
		JoinPriority int `json:"join-priority"`
	} `json:"board-data-opt"`
	DescriptorData struct {
		RadioSlotsInUse        int  `json:"radio-slots-in-use"`
		EncryptionCapabilities bool `json:"encryption-capabilities"`
	} `json:"descriptor-data"`
	ApProv struct {
		IsUniversal          bool   `json:"is-universal"`
		UniversalPrimeStatus string `json:"universal-prime-status"`
	} `json:"ap-prov"`
	ApModels struct {
		Model string `json:"model"`
	} `json:"ap-models"`
	NumPorts     int    `json:"num-ports"`
	NumSlots     int    `json:"num-slots"`
	WtpModelType int    `json:"wtp-model-type"`
	ApCapability string `json:"ap-capability"`
	IsMmOpt      bool   `json:"is-mm-opt"`
	ApImageName  string `json:"ap-image-name"`
}

type DynamicInfo struct {
	ApCrashData struct {
		ApCrashFile           string `json:"ap-crash-file"`
		ApRadio2GCrashFile    string `json:"ap-radio-2g-crash-file"`
		ApRadio5GCrashFile    string `json:"ap-radio-5g-crash-file"`
		ApRadio6GCrashFile    string `json:"ap-radio-6g-crash-file"`
		ApRad5GSlot2CrashFile string `json:"ap-rad-5g-slot2-crash-file"`
	} `json:"ap-crash-data"`
	LedStateEnabled  bool `json:"led-state-enabled"`
	ResetButtonState bool `json:"reset-button-state"`
	LedFlashEnabled  bool `json:"led-flash-enabled"`
	FlashSec         int  `json:"flash-sec"`
	TempInfo         struct {
		Degree       int    `json:"degree"`
		TempStatus   string `json:"temp-status"`
		HeaterStatus string `json:"heater-status"`
	} `json:"temp-info"`
	LedFlashExpiry time.Time `json:"led-flash-expiry"`
}

type WtpVersion struct {
	BackupSwVersion struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"backup-sw-version"`
	MiniIosVersion struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"mini-ios-version"`
	SwVer struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"sw-ver"`
	BootVer struct {
		Version int `json:"version"`
		Release int `json:"release"`
		Maint   int `json:"maint"`
		Build   int `json:"build"`
	} `json:"boot-ver"`
	SwVersion string `json:"sw-version"`
}

type DeviceDetail struct {
	StaticInfo  StaticInfo  `json:"static-info"`
	DynamicInfo DynamicInfo `json:"dynamic-info"`
	WtpVersion  WtpVersion  `json:"wtp-version"`
}

type AaaLocation struct {
	CivicID string `json:"civic-id"`
	GeoID   string `json:"geo-id"`
	OperID  string `json:"oper-id"`
}

type ApLocation struct {
	Floor             int         `json:"floor"`
	Location          string      `json:"location"`
	AaaLocation       AaaLocation `json:"aaa-location"`
	FloorID           int         `json:"floor-id"`
	RangingCapability int         `json:"ranging-capability"`
}

type ApDhcpServer struct {
	IsDhcpServerEnabled bool `json:"is-dhcp-server-enabled"`
}

type ApServices struct {
	MonitorModeOptType string       `json:"monitor-mode-opt-type"`
	ApDhcpServer       ApDhcpServer `json:"ap-dhcp-server"`
	TotSnifferRadio    int          `json:"tot-sniffer-radio"`
}

type TagInfo struct {
	TagSource         string `json:"tag-source"`
	IsApMisconfigured bool   `json:"is-ap-misconfigured"`
	ResolvedTagInfo   struct {
		ResolvedPolicyTag string `json:"resolved-policy-tag"`
		ResolvedSiteTag   string `json:"resolved-site-tag"`
		ResolvedRfTag     string `json:"resolved-rf-tag"`
	} `json:"resolved-tag-info"`
	PolicyTagInfo struct {
		PolicyTagName string `json:"policy-tag-name"`
	} `json:"policy-tag-info"`
	SiteTag struct {
		SiteTagName string `json:"site-tag-name"`
		ApProfile   string `json:"ap-profile"`
		FlexProfile string `json:"flex-profile"`
	} `json:"site-tag"`
	RfTag struct {
		RfTagName string `json:"rf-tag-name"`
	} `json:"rf-tag"`
	FilterInfo struct {
		FilterName string `json:"filter-name"`
	} `json:"filter-info"`
	IsDtlsLscFbkAp bool `json:"is-dtls-lsc-fbk-ap"`
}

type CapwapData struct {
	WtpMac       string       `json:"wtp-mac"`
	IPAddr       string       `json:"ip-addr"`
	Name         string       `json:"name"`
	DeviceDetail DeviceDetail `json:"device-detail"`
	ApLagEnabled bool         `json:"ap-lag-enabled"`
	ApLocation   ApLocation   `json:"ap-location"`
	ApServices   ApServices   `json:"ap-services"`
	TagInfo      TagInfo      `json:"tag-info"`
	Tunnel       struct {
		PreferredMode string `json:"preferred-mode"`
		UDPLite       string `json:"udp-lite"`
	} `json:"tunnel"`
	ExternalModuleData struct {
		XmData struct {
			IsModulePresent bool `json:"is-module-present"`
			Xm              struct {
				NumericID          int    `json:"numeric-id"`
				MaxPower           int    `json:"max-power"`
				SerialNumberString string `json:"serial-number-string"`
				ProductIDString    string `json:"product-id-string"`
				ModuleType         string `json:"module-type"`
				ModuleDescription  string `json:"module-description"`
			} `json:"xm"`
		} `json:"xm-data"`
		UsbData struct {
			IsModulePresent bool `json:"is-module-present"`
			Xm              struct {
				NumericID          int    `json:"numeric-id"`
				MaxPower           int    `json:"max-power"`
				SerialNumberString string `json:"serial-number-string"`
				ProductIDString    string `json:"product-id-string"`
				ModuleType         string `json:"module-type"`
				ModuleDescription  string `json:"module-description"`
			} `json:"xm"`
		} `json:"usb-data"`
		UsbOverride        bool `json:"usb-override"`
		IsExtModuleEnabled bool `json:"is-ext-module-enabled"`
	} `json:"external-module-data"`
	Ipv6Joined int `json:"ipv6-joined"`
	ApState    struct {
		ApAdminState     string `json:"ap-admin-state"`
		ApOperationState string `json:"ap-operation-state"`
	} `json:"ap-state"`
	ApModeData struct {
		HomeApEnabled bool   `json:"home-ap-enabled"`
		ClearMode     bool   `json:"clear-mode"`
		ApSubMode     string `json:"ap-sub-mode"`
		WtpMode       string `json:"wtp-mode"`
		ApFabricData  struct {
			IsFabricAp bool `json:"is-fabric-ap"`
		} `json:"ap-fabric-data"`
	} `json:"ap-mode-data"`
	ApTimeInfo struct {
		BootTime      time.Time `json:"boot-time"`
		JoinTime      time.Time `json:"join-time"`
		JoinTimeTaken int       `json:"join-time-taken"`
	} `json:"ap-time-info"`
	CountryCode    string `json:"country-code"`
	ApSecurityData struct {
		FipsEnabled      bool      `json:"fips-enabled"`
		WlanccEnabled    bool      `json:"wlancc-enabled"`
		CertType         string    `json:"cert-type"`
		LscApAuthType    string    `json:"lsc-ap-auth-type"`
		ApCertPolicy     string    `json:"ap-cert-policy"`
		ApCertExpiryTime time.Time `json:"ap-cert-expiry-time"`
		ApCertIssuerCn   string    `json:"ap-cert-issuer-cn"`
	} `json:"ap-security-data"`
	NumRadioSlots   int  `json:"num-radio-slots"`
	DartIsConnected bool `json:"dart-is-connected"`
	IsMaster        bool `json:"is-master"`
	SlidingWindow   struct {
		MultiWindowSupport bool `json:"multi-window-support"`
		WindowSize         int  `json:"window-size"`
	} `json:"sliding-window"`
	ApVlan struct {
		VlanTagState string `json:"vlan-tag-state"`
		VlanTagID    int    `json:"vlan-tag-id"`
	} `json:"ap-vlan"`
	HyperlocationData struct {
		HyperlocationMethod string `json:"hyperlocation-method"`
	} `json:"hyperlocation-data"`
	CdpEnable        bool   `json:"cdp-enable"`
	ApStationingType string `json:"ap-stationing-type"`
	RebootStats      struct {
		RebootReason string `json:"reboot-reason"`
		RebootType   string `json:"reboot-type"`
	} `json:"reboot-stats"`
	ProxyInfo struct {
		Hostname    string `json:"hostname"`
		Port        int    `json:"port"`
		NoProxyList string `json:"no-proxy-list"`
	} `json:"proxy-info"`
	GrpcEnabled         bool      `json:"grpc-enabled"`
	ImageSizeEta        int       `json:"image-size-eta"`
	ImageSizeStartTime  time.Time `json:"image-size-start-time"`
	ImageSizePercentage int       `json:"image-size-percentage"`
	MdnsGroupID         int       `json:"mdns-group-id"`
	MdnsRuleName        string    `json:"mdns-rule-name"`
	ApKeepaliveState    bool      `json:"ap-keepalive-state"`
	LocalDhcp           bool      `json:"local-dhcp"`
	Ipv4Pool            struct {
		Network   string `json:"network"`
		LeaseTime int    `json:"lease-time"`
		Netmask   string `json:"netmask"`
	} `json:"ipv4-pool"`
	WlcImageSizeEta        int       `json:"wlc-image-size-eta"`
	WlcImageSizeStartTime  time.Time `json:"wlc-image-size-start-time"`
	WlcImageSizePercentage int       `json:"wlc-image-size-percentage"`
	DisconnectDetail       struct {
		DisconnectReason string `json:"disconnect-reason"`
	} `json:"disconnect-detail"`
	WtpIP        string `json:"wtp-ip"`
	StatsMonitor struct {
		ActionApReload bool `json:"action-ap-reload"`
	} `json:"stats-monitor"`
	LscStatusPldSupported []any `json:"lsc-status-pld-supported"`
	ApLscStatus           struct {
		IsDtlsLscEnabled  bool `json:"is-dtls-lsc-enabled"`
		IsDot1XLscEnabled bool `json:"is-dot1x-lsc-enabled"`
		IsDtlsLscFallback bool `json:"is-dtls-lsc-fallback"`
	} `json:"ap-lsc-status"`
	RadioStatsMonitor struct {
		Enable       bool  `json:"enable"`
		SampleIntvl  int   `json:"sample-intvl"`
		AlarmsEnable []any `json:"alarms-enable"`
		RadioReset   bool  `json:"radio-reset"`
	} `json:"radio-stats-monitor"`
	OobImgDwldMethod string `json:"oob-img-dwld-method"`
	ZeroWtDfs        struct {
		ReserveChannel struct {
			Channel      int    `json:"channel"`
			ChannelWidth string `json:"channel-width"`
			State        string `json:"state"`
		} `json:"reserve-channel"`
		Type string `json:"type"`
	} `json:"zero-wt-dfs"`
	MaxClientsSupported int    `json:"max-clients-supported"`
	MerakiCapable       bool   `json:"meraki-capable"`
	MdnsGroupMethod     string `json:"mdns-group-method"`
	GnssInfo            struct {
		AntType          string `json:"ant-type"`
		AntCableLength   int    `json:"ant-cable-length"`
		AntennaProductID string `json:"antenna-product-id"`
	} `json:"gnss-info"`
	DartConStatus        string `json:"dart-con-status"`
	ApAfcPreNotification bool   `json:"ap-afc-pre-notification"`
	KernelCoredumpCount  int    `json:"kernel-coredump-count"`
	RegDomain            string `json:"reg-domain"`
	MerakiMonitorCapable bool   `json:"meraki-monitor-capable"`
	MerakiConnectStatus  string `json:"meraki-connect-status"`
}

type ApNameMacMap struct {
	WtpName string `json:"wtp-name"`
	WtpMac  string `json:"wtp-mac"`
	EthMac  string `json:"eth-mac"`
}

type WtpSlotWlanStats struct {
	WtpMac      string `json:"wtp-mac"`
	SlotID      int    `json:"slot-id"`
	WlanID      int    `json:"wlan-id"`
	BssidMac    string `json:"bssid-mac"`
	Ssid        string `json:"ssid"`
	BytesRx     string `json:"bytes-rx"`
	BytesTx     string `json:"bytes-tx"`
	PktsRx      string `json:"pkts-rx"`
	PktsTx      string `json:"pkts-tx"`
	DataRetries string `json:"data-retries"`
}

type EthernetMacWtpMacMap struct {
	EthernetMac string `json:"ethernet-mac"`
	WtpMac      string `json:"wtp-mac"`
}

type RadioOperStats struct {
	ApMac                 string `json:"ap-mac"`
	SlotID                int    `json:"slot-id"`
	AidUserList           int    `json:"aid-user-list"`
	TxFragmentCount       int    `json:"tx-fragment-count"`
	MulticastTxFrameCnt   int    `json:"multicast-tx-frame-cnt"`
	FailedCount           int    `json:"failed-count"`
	RetryCount            int    `json:"retry-count"`
	MultipleRetryCount    int    `json:"multiple-retry-count"`
	FrameDuplicateCount   int    `json:"frame-duplicate-count"`
	RtsSuccessCount       int    `json:"rts-success-count"`
	RtsFailureCount       int    `json:"rts-failure-count"`
	AckFailureCount       int    `json:"ack-failure-count"`
	RxFragmentCount       int    `json:"rx-fragment-count"`
	MulticastRxFrameCnt   int    `json:"multicast-rx-frame-cnt"`
	FcsErrorCount         int    `json:"fcs-error-count"`
	TxFrameCount          int    `json:"tx-frame-count"`
	WepUndecryptableCount int    `json:"wep-undecryptable-count"`
	RxErrorFrameCount     int    `json:"rx-error-frame-count"`
	MacMicErrFrameCount   int    `json:"mac-mic-err-frame-count"`
	RxMgmtFrameCount      int    `json:"rx-mgmt-frame-count"`
	RxCtrlFrameCount      int    `json:"rx-ctrl-frame-count"`
	RxDataFrameCount      int    `json:"rx-data-frame-count"`
	TxMgmtFrameCount      int    `json:"tx-mgmt-frame-count"`
	TxCtrlFrameCount      int    `json:"tx-ctrl-frame-count"`
	TxDataFrameCount      int    `json:"tx-data-frame-count"`
	RxDataPktCount        int    `json:"rx-data-pkt-count"`
	TxDataPktCount        int    `json:"tx-data-pkt-count"`
	NoiseFloor            int    `json:"noise-floor"`
	ApRadioStats          struct {
		StuckTs            time.Time `json:"stuck-ts"`
		LastTs             time.Time `json:"last-ts"`
		NumRadioStuckReset int       `json:"num-radio-stuck-reset"`
	} `json:"ap-radio-stats"`
	MacDecryErrFrameCount int `json:"mac-decry-err-frame-count"`
}

type EthernetIfStats struct {
	WtpMac           string `json:"wtp-mac"`
	IfIndex          int    `json:"if-index"`
	IfName           string `json:"if-name"`
	RxPkts           int    `json:"rx-pkts"`
	TxPkts           int    `json:"tx-pkts"`
	OperStatus       string `json:"oper-status"`
	RxUcastPkts      int    `json:"rx-ucast-pkts"`
	RxNonUcastPkts   int    `json:"rx-non-ucast-pkts"`
	TxUcastPkts      int    `json:"tx-ucast-pkts"`
	TxNonUcastPkts   int    `json:"tx-non-ucast-pkts"`
	Duplex           int    `json:"duplex"`
	LinkSpeed        int    `json:"link-speed"`
	RxTotalBytes     int64  `json:"rx-total-bytes"`
	TxTotalBytes     int64  `json:"tx-total-bytes"`
	InputCrc         int    `json:"input-crc"`
	InputAborts      int    `json:"input-aborts"`
	InputErrors      int    `json:"input-errors"`
	InputFrames      int    `json:"input-frames"`
	InputOverrun     int    `json:"input-overrun"`
	InputDrops       int    `json:"input-drops"`
	InputResource    int    `json:"input-resource"`
	UnknownProtocol  int    `json:"unknown-protocol"`
	Runts            int    `json:"runts"`
	Giants           int    `json:"giants"`
	Throttle         int    `json:"throttle"`
	Resets           int    `json:"resets"`
	OutputCollision  int    `json:"output-collision"`
	OutputNoBuffer   int    `json:"output-no-buffer"`
	OutputResource   int    `json:"output-resource"`
	OutputUnderrun   int    `json:"output-underrun"`
	OutputErrors     int    `json:"output-errors"`
	OutputTotalDrops int    `json:"output-total-drops"`
}

type EwlcWncdStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`
		NumInProgress           int  `json:"num-in-progress"`
		NumComplete             int  `json:"num-complete"`
		NumUnsupported          int  `json:"num-unsupported"`
		NumFailed               int  `json:"num-failed"`
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"`
		NumTotal                int  `json:"num-total"`
	} `json:"predownload-stats"`
	DownloadsInProgress int `json:"downloads-in-progress"`
	DownloadsComplete   int `json:"downloads-complete"`
	WlcPredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`
		NumInProgress           int  `json:"num-in-progress"`
		NumComplete             int  `json:"num-complete"`
		NumUnsupported          int  `json:"num-unsupported"`
		NumFailed               int  `json:"num-failed"`
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"`
		NumTotal                int  `json:"num-total"`
	} `json:"wlc-predownload-stats"`
}

type ApIoxOperData struct {
	ApMac        string `json:"ap-mac"`
	ApphostState string `json:"apphost-state"`
	CafToken     string `json:"caf-token"`
	CafPort      int    `json:"caf-port"`
}

type QosGlobalStats struct {
	QosClientVoiceStats struct {
		TotalNumOfTspecRcvd       int `json:"total-num-of-tspec-rcvd"`
		NewTspecFromAssocReq      int `json:"new-tspec-from-assoc-req"`
		TspecRenewalFromAssocReq  int `json:"tspec-renewal-from-assoc-req"`
		NewTspecAsAddTs           int `json:"new-tspec-as-add-ts"`
		TspecRenewalFromAddTs     int `json:"tspec-renewal-from-add-ts"`
		TspecProcessFailedGetRec  int `json:"tspec-process-failed-get-rec"`
		TotalSipInviteOnCaller    int `json:"total-sip-invite-on-caller"`
		TotalSipInviteOnCallee    int `json:"total-sip-invite-on-callee"`
		TotalNumOfCallReport      int `json:"total-num-of-call-report"`
		TotalSipFailureTrapSend   int `json:"total-sip-failure-trap-send"`
		NumOfCallsAccepted        int `json:"num-of-calls-accepted"`
		NumOfCallsRejectedInsufBw int `json:"num-of-calls-rejected-insuf-bw"`
		NumOfCallsRejectedQos     int `json:"num-of-calls-rejected-qos"`
		NumOfCallsRejectedPhyRate int `json:"num-of-calls-rejected-phy-rate"`
		NumOfCallsRejInvalidTspec int `json:"num-of-calls-rej-invalid-tspec"`
		NumOfRoamCallsAccepted    int `json:"num-of-roam-calls-accepted"`
		NumOfRoamCallsRejected    int `json:"num-of-roam-calls-rejected"`
		NumOfActiveSipCalls       int `json:"num-of-active-sip-calls"`
		NumOfActiveTspecCalls     int `json:"num-of-active-tspec-calls"`
	} `json:"qos-client-voice-stats"`
}

type ApOperData struct {
	WtpMac                string `json:"wtp-mac"`
	RadioID               int    `json:"radio-id"`
	ApAntennaBandMode     string `json:"ap-antenna-band-mode"`
	LinkEncryptionEnabled bool   `json:"link-encryption-enabled"`
	ApRemoteDebugMode     bool   `json:"ap-remote-debug-mode"`
	ApIPData              struct {
		ApPrefix         int    `json:"ap-prefix"`
		Mtu              int    `json:"mtu"`
		IsStaticApIpaddr bool   `json:"is-static-ap-ipaddr"`
		DomainName       string `json:"domain-name"`
		ApIPAddr         string `json:"ap-ip-addr"`
		ApIpv6Addr       string `json:"ap-ipv6-addr"`
		ApIPNetmask      string `json:"ap-ip-netmask"`
		ApIPGateway      string `json:"ap-ip-gateway"`
		ApIpv6Gateway    string `json:"ap-ipv6-gateway"`
		ApNameServerType string `json:"ap-name-server-type"`
		ApIpv6Method     string `json:"ap-ipv6-method"`
		StaticIP         string `json:"static-ip"`
		StaticGwIP       string `json:"static-gw-ip"`
		StaticNetmask    string `json:"static-netmask"`
		StaticPrefix     int    `json:"static-prefix"`
	} `json:"ap-ip-data"`
	ApPrimeInfo struct {
		PrimaryControllerName     string `json:"primary-controller-name"`
		SecondaryControllerName   string `json:"secondary-controller-name"`
		PrimaryControllerIPAddr   string `json:"primary-controller-ip-addr"`
		SecondaryControllerIPAddr string `json:"secondary-controller-ip-addr"`
		TertiaryControllerName    string `json:"tertiary-controller-name"`
		TertiaryControllerIPAddr  string `json:"tertiary-controller-ip-addr"`
		ApFallbackIP              string `json:"ap-fallback-ip"`
		FallbackEnabled           bool   `json:"fallback-enabled"`
	} `json:"ap-prime-info"`
	ApMgmt struct {
		IsTelnetEnabled  bool `json:"is-telnet-enabled"`
		IsSSHEnabled     bool `json:"is-ssh-enabled"`
		IsConsoleEnabled bool `json:"is-console-enabled"`
	} `json:"ap-mgmt"`
	ApLoginCredentials struct {
		Dot1XEapType  string `json:"dot1x-eap-type"`
		UserName      string `json:"user-name"`
		Dot1XUsername string `json:"dot1x-username"`
	} `json:"ap-login-credentials"`
	ApPow struct {
		PowerInjectorSel     string `json:"power-injector-sel"`
		PowerInjectorMacaddr string `json:"power-injector-macaddr"`
		PreStdSwitchEnabled  bool   `json:"pre-std-switch-enabled"`
		PowerInjectorEnabled bool   `json:"power-injector-enabled"`
		PowerType            string `json:"power-type"`
		PowerMode            string `json:"power-mode"`
	} `json:"ap-pow"`
	ApSysStats struct {
		CPUUsage       int       `json:"cpu-usage"`
		MemoryUsage    int       `json:"memory-usage"`
		AvgCPUUsage    int       `json:"avg-cpu-usage"`
		AvgMemoryUsage int       `json:"avg-memory-usage"`
		WindowSize     int       `json:"window-size"`
		LastTs         time.Time `json:"last-ts"`
		Memory         struct {
			AlarmStatus   string    `json:"alarm-status"`
			RaiseTicks    time.Time `json:"raise-ticks"`
			ClearTicks    time.Time `json:"clear-ticks"`
			LastSendTicks time.Time `json:"last-send-ticks"`
			Type          string    `json:"type"`
		} `json:"memory"`
		CPU struct {
			AlarmStatus   string    `json:"alarm-status"`
			RaiseTicks    time.Time `json:"raise-ticks"`
			ClearTicks    time.Time `json:"clear-ticks"`
			LastSendTicks time.Time `json:"last-send-ticks"`
			Type          string    `json:"type"`
		} `json:"cpu"`
		UplinkIf struct {
			Intf []struct {
				IfName           string `json:"if-name"`
				IfUpstreamRate   int    `json:"if-upstream-rate"`
				IfDownstreamRate int    `json:"if-downstream-rate"`
			} `json:"intf"`
		} `json:"uplink-if"`
	} `json:"ap-sys-stats"`
	Ipv4TCPMss struct {
		TCPAdjustMssState bool `json:"tcp-adjust-mss-state"`
		TCPAdjustMssSize  int  `json:"tcp-adjust-mss-size"`
	} `json:"ipv4-tcp-mss"`
	Ipv6TCPMss struct {
		TCPAdjustMssState bool `json:"tcp-adjust-mss-state"`
		TCPAdjustMssSize  int  `json:"tcp-adjust-mss-size"`
	} `json:"ipv6-tcp-mss"`
	LinkAudit struct {
		LinkauditFlag      string `json:"linkaudit-flag"`
		LinkauditDelayTime int    `json:"linkaudit-delay-time"`
		LinkauditMaxTime   int    `json:"linkaudit-max-time"`
		LinkauditMinTime   int    `json:"linkaudit-min-time"`
		LinkauditRcvTime   int    `json:"linkaudit-rcv-time"`
	} `json:"link-audit"`
	Timer struct {
		StatsTimer struct {
			StatTmr int `json:"stat-tmr"`
		} `json:"stats-timer"`
	} `json:"timer"`
	Retransmit struct {
		Count    int `json:"count"`
		Interval int `json:"interval"`
	} `json:"retransmit"`
	Syslog struct {
		LogHostIpaddr    string `json:"log-host-ipaddr"`
		LogTrapLevel     string `json:"log-trap-level"`
		LogTLSMode       bool   `json:"log-tls-mode"`
		LogFacilityLevel string `json:"log-facility-level"`
	} `json:"syslog"`
	InfrastructureMfp struct {
		MfpValidation bool `json:"mfp-validation"`
		MfpProtection bool `json:"mfp-protection"`
	} `json:"infrastructure-mfp"`
	PersistentSsid struct {
		IsPersistentSsidEnabled bool `json:"is-persistent-ssid-enabled"`
	} `json:"persistent-ssid"`
	ApGasRateLimitCfg struct {
		IsGasRateLimitEnabled bool `json:"is-gas-rate-limit-enabled"`
		NumReqPerInterval     int  `json:"num-req-per-interval"`
		IntervalMsec          int  `json:"interval-msec"`
	} `json:"ap-gas-rate-limit-cfg"`
	ApNtpServerInfoCfg struct {
		NtpServerAddress string `json:"ntp-server-address"`
		TrustKey         string `json:"trust-key"`
		KeyID            int    `json:"key-id"`
		KeyType          string `json:"key-type"`
		KeyFormat        string `json:"key-format"`
		TrustKeyType     string `json:"trust-key-type"`
	} `json:"ap-ntp-server-info-cfg"`
	ApUdpliteInfo string `json:"ap-udplite-info"`
	Accounting    struct {
		MethodList string `json:"method-list"`
	} `json:"accounting"`
	AuxClientInterfaceData struct {
		IsEnabled bool `json:"is-enabled"`
	} `json:"aux-client-interface-data"`
	ApDnaData struct {
		GrpcStatus        string `json:"grpc-status"`
		PacketsTxAttempts string `json:"packets-tx-attempts"`
		PacketsTxFailures string `json:"packets-tx-failures"`
		PacketsRx         string `json:"packets-rx"`
		PacketsRxFailures string `json:"packets-rx-failures"`
	} `json:"ap-dna-data"`
	ApIndoorMode bool `json:"ap-indoor-mode"`
	IsLocalNet   bool `json:"is-local-net"`
	OeapAudit    struct {
		LastRun    time.Time `json:"last-run"`
		State      string    `json:"state"`
		DtlsUpload string    `json:"dtls-upload"`
		Latency    int       `json:"latency"`
		Jitter     int       `json:"jitter"`
	} `json:"oeap-audit"`
	ProvSsid        bool `json:"prov-ssid"`
	ApNtpSyncStatus struct {
		Enabled          bool      `json:"enabled"`
		Stratum          int       `json:"stratum"`
		Status           string    `json:"status"`
		SecSinceLastSync int       `json:"sec-since-last-sync"`
		SyncOffset       int       `json:"sync-offset"`
		RxTs             time.Time `json:"rx-ts"`
	} `json:"ap-ntp-sync-status"`
	ApTzConfig struct {
		TzEnabled  bool   `json:"tz-enabled"`
		OffsetHour int    `json:"offset-hour"`
		OffsetMin  int    `json:"offset-min"`
		Mode       string `json:"mode"`
	} `json:"ap-tz-config"`
	ApPmkPropagationStatus bool   `json:"ap-pmk-propagation-status"`
	ApRole                 string `json:"ap-role"`
	PmkBsReceiverAddr      string `json:"pmk-bs-receiver-addr"`
	PmkBsSenderAddr        string `json:"pmk-bs-sender-addr"`
	PowerProfile           string `json:"power-profile"`
	PwrCalProfile          string `json:"pwr-cal-profile"`
	MaxClientsAllowed      int    `json:"max-clients-allowed"`
	PrimingProfileSrc      string `json:"priming-profile-src"`
	PrimingProfile         string `json:"priming-profile"`
	PrimingFilter          string `json:"priming-filter"`
	ApPrimingOverride      bool   `json:"ap-priming-override"`
	KernelCoredump         struct {
		KernelCoredumpLimit int `json:"kernel-coredump-limit"`
	} `json:"kernel-coredump"`
	RangingMode string `json:"ranging-mode"`
	PwrProfType string `json:"pwr-prof-type"`
}

type RlanOper struct {
	WtpMac         string `json:"wtp-mac"`
	RlanPortID     int    `json:"rlan-port-id"`
	RlanOperState  bool   `json:"rlan-oper-state"`
	RlanPortStatus bool   `json:"rlan-port-status"`
	RlanVlanValid  bool   `json:"rlan-vlan-valid"`
	RlanVlanID     int    `json:"rlan-vlan-id"`
	RlanPoeState   string `json:"rlan-poe-state"`
	PowerLevelID   int    `json:"power-level-id"`
}

type EwlcMewlcPredownloadRec struct {
	PredState                    string `json:"pred-state"`
	MeCapableApCount             int    `json:"me-capable-ap-count"`
	ControllerPredownloadVersion string `json:"controller-predownload-version"`
}

type CdpCacheData struct {
	MacAddr                string    `json:"mac-addr"`
	CdpCacheDeviceID       string    `json:"cdp-cache-device-id"`
	ApName                 string    `json:"ap-name"`
	LastUpdatedTime        time.Time `json:"last-updated-time"`
	Version                int       `json:"version"`
	CdpCacheIfIndex        int       `json:"cdp-cache-if-index"`
	CdpCacheVersion        string    `json:"cdp-cache-version"`
	CdpCacheDevicePort     string    `json:"cdp-cache-device-port"`
	CdpCacheLocalPort      string    `json:"cdp-cache-local-port"`
	CdpCachePlatform       string    `json:"cdp-cache-platform"`
	CdpCapabilitiesString  string    `json:"cdp-capabilities-string"`
	CdpCacheApAddress      string    `json:"cdp-cache-ap-address"`
	CdpAddrCount           int       `json:"cdp-addr-count"`
	CdpCacheIPAddressValue string    `json:"cdp-cache-ip-address-value"`
	IPAddress              struct {
		IPAddressValue []string `json:"ip-address-value"`
	} `json:"ip-address"`
	CdpCacheDuplex         string `json:"cdp-cache-duplex"`
	CdpCacheInterfaceSpeed int    `json:"cdp-cache-interface-speed"`
	WtpMacAddr             string `json:"wtp-mac-addr"`
	DeviceIndex            int    `json:"device-index"`
}

type LldpNeigh struct {
	WtpMac          string `json:"wtp-mac"`
	NeighMac        string `json:"neigh-mac"`
	PortID          string `json:"port-id"`
	LocalPort       string `json:"local-port"`
	SystemName      string `json:"system-name"`
	PortDescription string `json:"port-description"`
	Capabilities    string `json:"capabilities"`
	MgmtAddr        string `json:"mgmt-addr"`
}

type Trustpoint struct {
	TrustpointName     string `json:"trustpoint-name"`
	IsCertAvailable    bool   `json:"is-cert-available"`
	IsPrivkeyAvailable bool   `json:"is-privkey-available"`
	CertHash           string `json:"cert-hash"`
	CertType           string `json:"cert-type"`
	FipsSuitability    string `json:"fips-suitability"`
}

type TpCertInfo struct {
	Trustpoint Trustpoint `json:"trustpoint"`
}

type DiscData struct {
	WtpMac           string `json:"wtp-mac"`
	DiscoveryPkts    string `json:"discovery-pkts"`
	DiscoveryErrPkts string `json:"discovery-err-pkts"`
}

type CapwapPkts struct {
	WtpMac            string `json:"wtp-mac"`
	CntrlPkts         string `json:"cntrl-pkts"`
	DataKeepAlivePkts string `json:"data-keep-alive-pkts"`
	CapwapErrorPkts   string `json:"capwap-error-pkts"`
	RfidPkts          string `json:"rfid-pkts"`
	Dot1XEapPkts      string `json:"dot1x-eap-pkts"`
	Dot1XMgmtPkts     string `json:"dot1x-mgmt-pkts"`
	Dot1XKeyTypePkts  string `json:"dot1x-key-type-pkts"`
	ArpPkts           string `json:"arp-pkts"`
	IPPkts            string `json:"ip-pkts"`
	IappPkts          string `json:"iapp-pkts"`
	DhcpPkts          string `json:"dhcp-pkts"`
	RrmPkts           string `json:"rrm-pkts"`
	Ipv6Pkts          string `json:"ipv6-pkts"`
	Dot1XCtrlPkts     string `json:"dot1x-ctrl-pkts"`
}

type CountryOper struct {
	CountryCode         string `json:"country-code"`
	CountryString       string `json:"country-string"`
	RegDomainStr80211Bg string `json:"reg-domain-str-80211bg"`
	RegDomainStr80211A  string `json:"reg-domain-str-80211a"`
	CountrySupported    bool   `json:"country-supported"`
	ChannelsString11Bg  string `json:"channels-string-11bg"`
	ChannelsString11A   string `json:"channels-string-11a"`
	Channels11Bg        string `json:"channels-11bg"`
	Channels11A         string `json:"channels-11a"`
	DcaChannels11Bg     string `json:"dca-channels-11bg"`
	DcaChannels11A      string `json:"dca-channels-11a"`
	RadarChannels11A    string `json:"radar-channels-11a"`
	RegDom6Ghz          string `json:"reg-dom-6ghz"`
	ChanInfo6Ghz        string `json:"chan-info-6ghz"`
}

type SuppCountryOper struct {
	CountryCode    string `json:"country-code"`
	CountryString  string `json:"country-string"`
	CountryCodeIso string `json:"country-code-iso"`
	RegDom24Ghz    struct {
		RegDomainCode []string `json:"reg-domain-code"`
	} `json:"reg-dom-24ghz"`
	RegDom5Ghz struct {
		RegDom5GhzRegDomainCode []string `json:"reg-domain-code"`
	} `json:"reg-dom-5ghz"`
	RegDom6Ghz struct {
		RegDom6GhzRegDomainCode []string `json:"reg-domain-code"`
	} `json:"reg-dom-6ghz"`
	ChanList24Ghz struct {
		Channel []int `json:"channel"`
	} `json:"chan-list-24ghz"`
	ChanList5Ghz struct {
		ChanList5GhzChannel []int `json:"channel"`
	} `json:"chan-list-5ghz,omitempty"`
	ChanList6Ghz struct {
		ChanList6GhzChannel []int `json:"channel"`
	} `json:"chan-list-6ghz,omitempty"`
	ChanListDca24Ghz struct {
		ChanListDca24GhzChannel []int `json:"channel"`
	} `json:"chan-list-dca-24ghz"`
	ChanListDca5Ghz struct {
		ChanListDca5GhzChannel []int `json:"channel"`
	} `json:"chan-list-dca-5ghz,omitempty"`
	ChanListDca6Ghz struct {
		ChanListDca6GhzChannel []int `json:"channel"`
	} `json:"chan-list-dca-6ghz,omitempty"`
	ChanListPsc6Ghz struct {
		ChanListPsc6GhzChannel []int `json:"channel"`
	} `json:"chan-list-psc-6ghz,omitempty"`
}

type ApNhGlobalData struct {
	AlgorithmRunning   bool `json:"algorithm-running"`
	AlgorithmItrCount  int  `json:"algorithm-itr-count"`
	IdealCapacityPerRg int  `json:"ideal-capacity-per-rg"`
	NumOfNeighborhood  int  `json:"num-of-neighborhood"`
}

type ImageData struct {
	ImageName     string   `json:"image-name"`
	ImageLocation string   `json:"image-location"`
	ImageVersion  string   `json:"image-version"`
	IsNew         bool     `json:"is-new"`
	FileSize      string   `json:"file-size"`
	ApModelList   []string `json:"ap-model-list"`
}

type ApImagePrepareLocation struct {
	Index     int         `json:"index"`
	ImageFile string      `json:"image-file"`
	ImageData []ImageData `json:"image-data"`
}

type ApImageActiveLocation struct {
	Index                          int    `json:"index"`
	ImageFile                      string `json:"image-file"`
	ApImageActiveLocationImageData []struct {
		ImageName                                 string   `json:"image-name"`
		ImageLocation                             string   `json:"image-location"`
		ImageVersion                              string   `json:"image-version"`
		IsNew                                     bool     `json:"is-new"`
		FileSize                                  string   `json:"file-size"`
		ApImageActiveLocationImageDataApModelList []string `json:"ap-model-list"`
	} `json:"image-data"`
}

// GetApOper retrieves complete access point operational data from the wireless controller.
// Returns all AP operational information including radio neighbors, operational data,
// reset statistics, QoS data, CAPWAP data, and various other AP-related metrics.
func GetApOper(client *wnc.Client, ctx context.Context) (*ApOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	var result ApOperResponse
	err := client.SendAPIRequest(ctx, ApOperEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApRadioNeighbor retrieves access point radio neighbor information.
// Returns data about neighboring access points detected by each radio.
func GetApRadioNeighbor(client *wnc.Client, ctx context.Context) (*ApOperApRadioNeighborResponse, error) {
	var result ApOperApRadioNeighborResponse
	err := client.SendAPIRequest(ctx, ApRadioNeighborEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApRadioOperData retrieves radio operational data for access points.
// Returns detailed operational information for each radio including configuration,
// capabilities, and current operational state.
func GetApRadioOperData(client *wnc.Client, ctx context.Context) (*ApOperRadioOperDataResponse, error) {
	var result ApOperRadioOperDataResponse
	err := client.SendAPIRequest(ctx, RadioOperDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApRadioResetStats retrieves radio reset statistics for troubleshooting.
// Returns information about radio resets including causes and occurrence counts.
func GetApRadioResetStats(client *wnc.Client, ctx context.Context) (*ApOperRadioResetStatsResponse, error) {
	var result ApOperRadioResetStatsResponse
	err := client.SendAPIRequest(ctx, RadioResetStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApQosClientData retrieves Quality of Service data for clients connected to access points.
// Returns QoS parameters and statistics for client traffic management.
func GetApQosClientData(client *wnc.Client, ctx context.Context) (*ApOperQosClientDataResponse, error) {
	var result ApOperQosClientDataResponse
	err := client.SendAPIRequest(ctx, QosClientDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApCapwapData retrieves CAPWAP (Control and Provisioning of Wireless Access Points) protocol data.
// Returns CAPWAP tunnel and communication information between APs and the controller.
func GetApCapwapData(client *wnc.Client, ctx context.Context) (*ApOperCapwapDataResponse, error) {
	var result ApOperCapwapDataResponse
	err := client.SendAPIRequest(ctx, CapwapDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApNameMacMap retrieves the mapping between access point names and MAC addresses.
// Returns the correlation between AP logical names and their physical MAC addresses.
func GetApNameMacMap(client *wnc.Client, ctx context.Context) (*ApOperApNameMacMapResponse, error) {
	var result ApOperApNameMacMapResponse
	err := client.SendAPIRequest(ctx, ApNameMacMapEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApWtpSlotWlanStats retrieves WTP (Wireless Termination Point) slot WLAN statistics.
// Returns per-slot WLAN statistics for wireless termination points.
func GetApWtpSlotWlanStats(client *wnc.Client, ctx context.Context) (*ApOperWtpSlotWlanStatsResponse, error) {
	var result ApOperWtpSlotWlanStatsResponse
	err := client.SendAPIRequest(ctx, WtpSlotWlanStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApEthernetMacWtpMacMap retrieves the mapping between Ethernet MAC addresses and WTP MAC addresses.
// Returns the correlation between wired and wireless interface MAC addresses for access points.
func GetApEthernetMacWtpMacMap(client *wnc.Client, ctx context.Context) (*ApOperEthernetMacWtpMacMapResponse, error) {
	var result ApOperEthernetMacWtpMacMapResponse
	err := client.SendAPIRequest(ctx, EthernetMacWtpMacMapEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApRadioOperStats retrieves radio operational statistics for access points.
// Returns detailed operational metrics and performance statistics for AP radios.
func GetApRadioOperStats(client *wnc.Client, ctx context.Context) (*ApOperRadioOperStatsResponse, error) {
	var result ApOperRadioOperStatsResponse
	err := client.SendAPIRequest(ctx, RadioOperStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApEthernetIfStats retrieves Ethernet interface statistics for access points.
// Returns network interface statistics for AP wired connections.
func GetApEthernetIfStats(client *wnc.Client, ctx context.Context) (*ApOperEthernetIfStatsResponse, error) {
	var result ApOperEthernetIfStatsResponse
	err := client.SendAPIRequest(ctx, EthernetIfStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApEwlcWncdStats retrieves EWLC (Embedded Wireless LAN Controller) WNCD statistics.
// Returns statistics for the wireless network control daemon on embedded controllers.
func GetApEwlcWncdStats(client *wnc.Client, ctx context.Context) (*ApOperEwlcWncdStatsResponse, error) {
	var result ApOperEwlcWncdStatsResponse
	err := client.SendAPIRequest(ctx, EwlcWncdStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApIoxOperData retrieves IOx (Cisco's application framework) operational data for access points.
// Returns information about applications and containers running on IOx-enabled access points.
func GetApIoxOperData(client *wnc.Client, ctx context.Context) (*ApOperApIoxOperDataResponse, error) {
	var result ApOperApIoxOperDataResponse
	err := client.SendAPIRequest(ctx, ApIoxOperDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApQosGlobalStats retrieves global Quality of Service statistics for access points.
// Returns system-wide QoS metrics and performance statistics.
func GetApQosGlobalStats(client *wnc.Client, ctx context.Context) (*ApOperQosGlobalStatsResponse, error) {
	var result ApOperQosGlobalStatsResponse
	err := client.SendAPIRequest(ctx, QosGlobalStatsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApOperData retrieves general operational data for access points.
// Returns comprehensive operational information about AP status and configuration.
func GetApOperData(client *wnc.Client, ctx context.Context) (*ApOperOperDataResponse, error) {
	var result ApOperOperDataResponse
	err := client.SendAPIRequest(ctx, OperDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApRlanOper retrieves RLAN (Remote LAN) operational data for access points.
// Returns information about remote LAN connections and bridging.
func GetApRlanOper(client *wnc.Client, ctx context.Context) (*ApOperRlanOperResponse, error) {
	var result ApOperRlanOperResponse
	err := client.SendAPIRequest(ctx, RlanOperEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApEwlcMewlcPredownloadRec retrieves EWLC MEWLC predownload records.
// Returns information about image predownload operations for mobility controllers.
func GetApEwlcMewlcPredownloadRec(client *wnc.Client, ctx context.Context) (*ApOperEwlcMewlcPredownloadRecResponse, error) {
	var result ApOperEwlcMewlcPredownloadRecResponse
	err := client.SendAPIRequest(ctx, EwlcMewlcPredownloadRecEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApCdpCacheData retrieves CDP (Cisco Discovery Protocol) cache data for access points.
// Returns information about neighboring Cisco devices discovered via CDP.
func GetApCdpCacheData(client *wnc.Client, ctx context.Context) (*ApOperCdpCacheDataResponse, error) {
	var result ApOperCdpCacheDataResponse
	err := client.SendAPIRequest(ctx, CdpCacheDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApLldpNeigh retrieves LLDP (Link Layer Discovery Protocol) neighbor information.
// Returns data about neighboring devices discovered through LLDP.
func GetApLldpNeigh(client *wnc.Client, ctx context.Context) (*ApOperLldpNeighResponse, error) {
	var result ApOperLldpNeighResponse
	err := client.SendAPIRequest(ctx, LldpNeighEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApTpCertInfo retrieves trustpoint certificate information for access points.
// Returns PKI certificate data and trust relationships.
func GetApTpCertInfo(client *wnc.Client, ctx context.Context) (*ApOperTpCertInfoResponse, error) {
	var result ApOperTpCertInfoResponse
	err := client.SendAPIRequest(ctx, TpCertInfoEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApDiscData retrieves discovery data for access points.
// Returns information about AP discovery process and neighbor detection.
func GetApDiscData(client *wnc.Client, ctx context.Context) (*ApOperDiscDataResponse, error) {
	var result ApOperDiscDataResponse
	err := client.SendAPIRequest(ctx, DiscDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApCapwapPkts retrieves CAPWAP packet statistics for access points.
// Returns detailed packet-level statistics for CAPWAP protocol communications.
func GetApCapwapPkts(client *wnc.Client, ctx context.Context) (*ApOperCapwapPktsResponse, error) {
	var result ApOperCapwapPktsResponse
	err := client.SendAPIRequest(ctx, CapwapPktsEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApCountryOper retrieves AP country operational data with context.
func GetApCountryOper(client *wnc.Client, ctx context.Context) (*ApOperCountryOperResponse, error) {
	var result ApOperCountryOperResponse
	err := client.SendAPIRequest(ctx, CountryOperEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApSuppCountryOper retrieves AP supported country operational data with context.
func GetApSuppCountryOper(client *wnc.Client, ctx context.Context) (*ApOperSuppCountryOperResponse, error) {
	var result ApOperSuppCountryOperResponse
	err := client.SendAPIRequest(ctx, SuppCountryOperEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApNhGlobalData retrieves access point next hop global data.
// Returns global routing and next hop information for access points.
func GetApNhGlobalData(client *wnc.Client, ctx context.Context) (*ApOperApNhGlobalDataResponse, error) {
	var result ApOperApNhGlobalDataResponse
	err := client.SendAPIRequest(ctx, ApNhGlobalDataEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApImagePrepareLocation retrieves access point image prepare location information.
// Returns data about AP image preparation and staging locations.
func GetApImagePrepareLocation(client *wnc.Client, ctx context.Context) (*ApOperApImagePrepareLocationResponse, error) {
	var result ApOperApImagePrepareLocationResponse
	err := client.SendAPIRequest(ctx, ApImagePrepareLocationEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetApImageActiveLocation retrieves access point image active location information.
// Returns data about currently active AP image locations and versions.
func GetApImageActiveLocation(client *wnc.Client, ctx context.Context) (*ApOperApImageActiveLocationResponse, error) {
	var result ApOperApImageActiveLocationResponse
	err := client.SendAPIRequest(ctx, ApImageActiveLocationEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
