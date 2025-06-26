// Package rrm provides Radio Resource Management operational data functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"

	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// RrmOperBasePath defines the base path for RRM operational data endpoints.
	RrmOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"
	// RrmOperEndpoint defines the endpoint for RRM operational data.
	RrmOperEndpoint = RrmOperBasePath
	// RrmOperApAutoRfDot11DataEndpoint defines the endpoint for AP auto RF 802.11 data.
	RrmOperApAutoRfDot11DataEndpoint = RrmOperBasePath + "/ap-auto-rf-dot11-data"
	// RrmOperApDot11RadarDataEndpoint defines the endpoint for AP 802.11 radar data.
	RrmOperApDot11RadarDataEndpoint = RrmOperBasePath + "/ap-dot11-radar-data"
	// RrmOperApDot11SpectrumDataEndpoint defines the endpoint for AP 802.11 spectrum data.
	RrmOperApDot11SpectrumDataEndpoint = RrmOperBasePath + "/ap-dot11-spectrum-data"
	// RrmOperRrmMeasurementEndpoint defines the endpoint for RRM measurement data.
	RrmOperRrmMeasurementEndpoint = RrmOperBasePath + "/rrm-measurement"
	// RrmOperRadioSlotEndpoint defines the endpoint for radio slot data.
	RrmOperRadioSlotEndpoint = RrmOperBasePath + "/radio-slot"
	// RrmOperMainDataEndpoint defines the endpoint for RRM main data.
	RrmOperMainDataEndpoint = RrmOperBasePath + "/main-data"
	// RrmOperSpectrumDeviceTableEndpoint defines the endpoint for spectrum device table data.
	RrmOperSpectrumDeviceTableEndpoint = RrmOperBasePath + "/spectrum-device-table"
	// RrmOperSpectrumAqTableEndpoint defines the endpoint for spectrum air quality table data.
	RrmOperSpectrumAqTableEndpoint = RrmOperBasePath + "/spectrum-aq-table"
	// RrmOperRegDomainOperEndpoint defines the endpoint for regulatory domain operational data.
	RrmOperRegDomainOperEndpoint = RrmOperBasePath + "/reg-domain-oper"
)

// RrmOperResponse represents the response structure for RRM operational data.
type RrmOperResponse struct {
	CiscoIOSXEWirelessRrmOperRrmOperData struct {
		ApAutoRfDot11Data   []ApAutoRfDot11Data   `json:"ap-auto-rf-dot11-data"`
		ApDot11RadarData    []ApDot11RadarData    `json:"ap-dot11-radar-data"`
		ApDot11SpectrumData []ApDot11SpectrumData `json:"ap-dot11-spectrum-data"`
		RrmMeasurement      []RrmMeasurement      `json:"rrm-measurement"`
		RadioSlot           []RadioSlot           `json:"radio-slot"`
		MainData            []MainData            `json:"main-data"`
		SpectrumDeviceTable []SpectrumDeviceTable `json:"spectrum-device-table"`
		SpectrumAqTable     []SpectrumAqTable     `json:"spectrum-aq-table"`
		RegDomainOper       RegDomainOper         `json:"reg-domain-oper"`
	} `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"`
}

// ApAutoRfDot11DataResponse represents the response structure for AP auto RF 802.11 data.
type ApAutoRfDot11DataResponse struct {
	ApAutoRfDot11Data []ApAutoRfDot11Data `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-auto-rf-dot11-data"`
}

// ApDot11RadarDataResponse represents the response structure for AP 802.11 radar data.
type ApDot11RadarDataResponse struct {
	ApDot11RadarData []ApDot11RadarData `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-radar-data"`
}

// ApDot11SpectrumDataResponse represents the response structure for AP 802.11 spectrum data.
type ApDot11SpectrumDataResponse struct {
	ApDot11SpectrumData []ApDot11SpectrumData `json:"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-spectrum-data"`
}

// RrmMeasurementResponse represents the response structure for RRM measurement data.
type RrmMeasurementResponse struct {
	RrmMeasurement []RrmMeasurement `json:"Cisco-IOS-XE-wireless-rrm-oper:rrm-measurement"`
}

type RadioSlotResponse struct {
	RadioSlot []RadioSlot `json:"Cisco-IOS-XE-wireless-rrm-oper:radio-slot"`
}

type MainDataResponse struct {
	MainData []MainData `json:"Cisco-IOS-XE-wireless-rrm-oper:main-data"`
}

type SpectrumDeviceTableResponse struct {
	SpectrumDeviceTable []SpectrumDeviceTable `json:"Cisco-IOS-XE-wireless-rrm-oper:spectrum-device-table"`
}

type SpectrumAqTableResponse struct {
	SpectrumAqTable []SpectrumAqTable `json:"Cisco-IOS-XE-wireless-rrm-oper:spectrum-aq-table"`
}

type RegDomainOperResponse struct {
	RegDomainOper RegDomainOper `json:"Cisco-IOS-XE-wireless-rrm-oper:reg-domain-oper"`
}

type ApAutoRfDot11Data struct {
	WtpMac            string `json:"wtp-mac"`
	RadioSlotID       int    `json:"radio-slot-id"`
	NeighborRadioInfo struct {
		NeighborRadioList []struct {
			NeighborRadioInfo struct {
				NeighborRadioMac    string `json:"neighbor-radio-mac"`
				NeighborRadioSlotID int    `json:"neighbor-radio-slot-id"`
				Rssi                int    `json:"rssi"`
				Snr                 int    `json:"snr"`
				Channel             int    `json:"channel"`
				Power               int    `json:"power"`
				GroupLeaderIP       string `json:"group-leader-ip"`
				ChanWidth           string `json:"chan-width"`
				SensorCovered       bool   `json:"sensor-covered"`
			} `json:"neighbor-radio-info"`
		} `json:"neighbor-radio-list"`
	} `json:"neighbor-radio-info,omitempty"`
}

type ApDot11RadarData struct {
	WtpMac           string    `json:"wtp-mac"`
	RadioSlotID      int       `json:"radio-slot-id"`
	LastRadarOnRadio time.Time `json:"last-radar-on-radio"`
}

type ApDot11SpectrumData struct {
	WtpMac      string `json:"wtp-mac"`
	RadioSlotID int    `json:"radio-slot-id"`
	Config      struct {
		SpectrumIntelligenceEnable bool   `json:"spectrum-intelligence-enable"`
		SpectrumWtpCaSiCapable     string `json:"spectrum-wtp-ca-si-capable"`
		SpectrumOperationState     string `json:"spectrum-operation-state"`
		SpectrumAdminState         bool   `json:"spectrum-admin-state"`
		SpectrumCapable            bool   `json:"spectrum-capable"`
		RapidUpdateEnable          bool   `json:"rapid-update-enable"`
		SensordOperationalStatus   int    `json:"sensord-operational-status"`
		ScanRadioType              string `json:"scan-radio-type"`
	} `json:"config"`
}

type RrmMeasurement struct {
	WtpMac      string `json:"wtp-mac"`
	RadioSlotID int    `json:"radio-slot-id"`
	Foreign     struct {
		ForeignForeign struct {
			ForeignData []struct {
				Chan                int `json:"chan"`
				Power               int `json:"power"`
				Rogue20Count        int `json:"rogue-20-count"`
				Rogue40PrimaryCount int `json:"rogue-40-primary-count"`
				Rogue80PrimaryCount int `json:"rogue-80-primary-count"`
				ChanUtil            int `json:"chan-util"`
			} `json:"foreign-data"`
		} `json:"foreign"`
	} `json:"foreign"`
	Noise struct {
		NoiseNoise struct {
			NoiseData []struct {
				Chan  int `json:"chan"`
				Noise int `json:"noise"`
			} `json:"noise-data"`
		} `json:"noise"`
	} `json:"noise"`
	Load struct {
		RxUtilPercentage          int `json:"rx-util-percentage"`
		TxUtilPercentage          int `json:"tx-util-percentage"`
		CcaUtilPercentage         int `json:"cca-util-percentage"`
		Stations                  int `json:"stations"`
		RxNoiseChannelUtilization int `json:"rx-noise-channel-utilization"`
		NonWifiInter              int `json:"non-wifi-inter"`
	} `json:"load"`
}

type RadioSlot struct {
	WtpMac      string `json:"wtp-mac"`
	RadioSlotID int    `json:"radio-slot-id"`
	RadioData   struct {
		BestTxPwrLevel            int  `json:"best-tx-pwr-level"`
		BestRtsThresh             int  `json:"best-rts-thresh"`
		BestFragThresh            int  `json:"best-frag-thresh"`
		LoadProfPassed            bool `json:"load-prof-passed"`
		CoverageProfilePassed     bool `json:"coverage-profile-passed"`
		InterferenceProfilePassed bool `json:"interference-profile-passed"`
		NoiseProfilePassed        bool `json:"noise-profile-passed"`
		DcaStats                  struct {
			BestChan          int `json:"best-chan"`
			CurrentChanEnergy int `json:"current-chan-energy"`
			LastChanEnergy    int `json:"last-chan-energy"`
			ChanChanges       int `json:"chan-changes"`
		} `json:"dca-stats"`
		CoverageOverlapFactor string `json:"coverage-overlap-factor"`
		SensorCoverageFactor  string `json:"sensor-coverage-factor"`
	} `json:"radio-data"`
}

type MainData struct {
	PhyType string `json:"phy-type"`
	Grp     struct {
		CurrentState string    `json:"current-state"`
		LastRun      time.Time `json:"last-run"`
		Dca          struct {
			DcaLastRun time.Time `json:"dca-last-run"`
		} `json:"dca"`
		Txpower struct {
			DpcLastRun time.Time `json:"dpc-last-run"`
			RunTime    int       `json:"run-time"`
		} `json:"txpower"`
		CurrentGroupingMode   string `json:"current-grouping-mode"`
		JoinProtocolVer       int    `json:"join-protocol-ver"`
		CurrentGroupingRole   string `json:"current-grouping-role"`
		CntrlrName            string `json:"cntrlr-name"`
		CntrlrIPAddr          string `json:"cntrlr-ip-addr"`
		CntrlrSecondaryIPAddr string `json:"cntrlr-secondary-ip-addr"`
		IsStaticMember        string `json:"is-static-member"`
		DpcConfig             struct {
			Rf struct {
				Mode              string `json:"mode"`
				UpdateCounter     int    `json:"update-counter"`
				UpdateIntervalSec int    `json:"update-interval-sec"`
				Contribution      int    `json:"contribution"`
			} `json:"rf"`
			DpcMinTxPowerLimit      int `json:"dpc-min-tx-power-limit"`
			DpcMaxTxPowerLimit      int `json:"dpc-max-tx-power-limit"`
			TxPowerControlThreshold int `json:"tx-power-control-threshold"`
		} `json:"dpc-config"`
		FraSensorCoverage int `json:"fra-sensor-coverage"`
		ProtocolVer       int `json:"protocol-ver"`
		HdrVer            int `json:"hdr-ver"`
	} `json:"grp"`
	RfName          string `json:"rf-name"`
	RrmMgrGrpMember []struct {
		MemberIP       string `json:"member-ip"`
		MaxRadioCnt    int    `json:"max-radio-cnt"`
		CurrRadioCnt   int    `json:"curr-radio-cnt"`
		Name           string `json:"name"`
		DtlsConnStatus string `json:"dtls-conn-status"`
	} `json:"rrm-mgr-grp-member"`
	OperData struct {
		DcaThreshVal       int `json:"dca-thresh-val"`
		DefaultDcaChannels struct {
			Channel []int `json:"channel"`
		} `json:"default-dca-channels"`
		DefaultNonDcaChannels struct {
			DefaultNonDcaChannelsChannel []int `json:"channel"`
		} `json:"default-non-dca-channels"`
		FraOperState bool `json:"fra-oper-state"`
	} `json:"oper-data,omitempty"`
}

type SpectrumDeviceTable struct {
	DeviceID        string    `json:"device-id"`
	ClusterID       string    `json:"cluster-id"`
	LastUpdatedTime time.Time `json:"last-updated-time"`
	IdrData         struct {
		DetectingApMac      string `json:"detecting-ap-mac"`
		AffectedChannelList string `json:"affected-channel-list"`
		IsPersistent        bool   `json:"is-persistent"`
		ClassTypeEnum       string `json:"class-type-enum"`
	} `json:"idr-data"`
}

type SpectrumAqTable struct {
	WtpMac          string `json:"wtp-mac"`
	Band            string `json:"band"`
	ReportingApName string `json:"reporting-ap-name"`
	PerRadioAqData  struct {
		ChannelCount     int `json:"channel-count"`
		PerChannelAqList []struct {
			ChannelNum           int       `json:"channel-num"`
			MinAqi               int       `json:"min-aqi"`
			Aqi                  int       `json:"aqi"`
			TotalIntfDeviceCount int       `json:"total-intf-device-count"`
			SpectrumTimestamp    time.Time `json:"spectrum-timestamp"`
		} `json:"per-channel-aq-list"`
	} `json:"per-radio-aq-data"`
	WtpCaSiCapable string `json:"wtp-ca-si-capable"`
	ScanRadioType  string `json:"scan-radio-type"`
}

type RegDomainOper struct {
	CountryList string `json:"country-list"`
}

// GetRrmOper retrieves RRM operational data with context.
func GetRrmOper(client *wnc.Client, ctx context.Context) (*RrmOperResponse, error) {
	var result RrmOperResponse
	err := client.SendAPIRequest(ctx, RrmOperEndpoint, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetApAutoRfDot11Data(client *wnc.Client, ctx context.Context) (*ApAutoRfDot11DataResponse, error) {
	var data ApAutoRfDot11DataResponse
	err := client.SendAPIRequest(ctx, RrmOperApAutoRfDot11DataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetApDot11RadarData(client *wnc.Client, ctx context.Context) (*ApDot11RadarDataResponse, error) {
	var data ApDot11RadarDataResponse
	err := client.SendAPIRequest(ctx, RrmOperApDot11RadarDataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetApDot11SpectrumData(client *wnc.Client, ctx context.Context) (*ApDot11SpectrumDataResponse, error) {
	var data ApDot11SpectrumDataResponse
	err := client.SendAPIRequest(ctx, RrmOperApDot11SpectrumDataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRrmMeasurement(client *wnc.Client, ctx context.Context) (*RrmMeasurementResponse, error) {
	var data RrmMeasurementResponse
	err := client.SendAPIRequest(ctx, RrmOperRrmMeasurementEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRadioSlot(client *wnc.Client, ctx context.Context) (*RadioSlotResponse, error) {
	var data RadioSlotResponse
	err := client.SendAPIRequest(ctx, RrmOperRadioSlotEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetMainData(client *wnc.Client, ctx context.Context) (*MainDataResponse, error) {
	var data MainDataResponse
	err := client.SendAPIRequest(ctx, RrmOperMainDataEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetSpectrumDeviceTable(client *wnc.Client, ctx context.Context) (*SpectrumDeviceTableResponse, error) {
	var data SpectrumDeviceTableResponse
	err := client.SendAPIRequest(ctx, RrmOperSpectrumDeviceTableEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetSpectrumAqTable(client *wnc.Client, ctx context.Context) (*SpectrumAqTableResponse, error) {
	var data SpectrumAqTableResponse
	err := client.SendAPIRequest(ctx, RrmOperSpectrumAqTableEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRegDomainOper(client *wnc.Client, ctx context.Context) (*RegDomainOperResponse, error) {
	var data RegDomainOperResponse
	err := client.SendAPIRequest(ctx, RrmOperRegDomainOperEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
