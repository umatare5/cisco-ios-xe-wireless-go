package wnc

import (
	"context"
	"log/slog"
	"time"
)

// WirelessControllerAPI defines the comprehensive interface for the Cisco Wireless Network Controller API client.
// It combines all feature-specific interfaces to provide a unified API for controller operations.
type WirelessControllerAPI interface {
	CoreAPI
	AccessPointAPI
	GeneralAPI
	RadioResourceManagementAPI
	WirelessLANAPI
	ClientAPI
	MobilityAPI
	RogueAPI
	NetworkManagementAPI
	HyperlocationAPI
	AWIPSApi
	GeolocationAPI
	AFCApi
	mDNSAPI
	MultcastAPI
	BluetoothAPI
	LISPApi
	SiteAPI
	Dot11API
	ApfAPI
	RfAPI
	Dot15API
	CTSAPI
	LocationAPI
	RadioAPI
	MeshAPI
	FlexAPI
	FabricAPI
	RFIDAPI
}

// CoreAPI defines the core interface for basic API operations.
type CoreAPI interface {
	SendAPIRequest(ctx context.Context, endpoint string, result any) error
}

// AccessPointAPI defines the interface for Access Point operations.
type AccessPointAPI interface {
	// Access Point Operational Data
	GetApOper(ctx context.Context) (any, error)
	GetApRadioNeighbor(ctx context.Context) (any, error)
	GetApRadioOperData(ctx context.Context) (any, error)
	GetApRadioResetStats(ctx context.Context) (any, error)
	GetApQosClientData(ctx context.Context) (any, error)
	GetApCapwapData(ctx context.Context) (any, error)
	GetApNameMacMap(ctx context.Context) (any, error)
	GetApWtpSlotWlanStats(ctx context.Context) (any, error)
	GetApEthernetMacWtpMacMap(ctx context.Context) (any, error)
	GetApRadioOperStats(ctx context.Context) (any, error)
	GetApEthernetIfStats(ctx context.Context) (any, error)
	GetApEwlcWncdStats(ctx context.Context) (any, error)
	GetApIoxOperData(ctx context.Context) (any, error)
	GetApQosGlobalStats(ctx context.Context) (any, error)
	GetApOperData(ctx context.Context) (any, error)
	GetApRlanOper(ctx context.Context) (any, error)
	GetApEwlcMewlcPredownloadRec(ctx context.Context) (any, error)
	GetApCdpCacheData(ctx context.Context) (any, error)
	GetApLldpNeigh(ctx context.Context) (any, error)
	GetApTpCertInfo(ctx context.Context) (any, error)
	GetApDiscData(ctx context.Context) (any, error)
	GetApCapwapPkts(ctx context.Context) (any, error)
	GetApCountryOper(ctx context.Context) (any, error)
	GetApSuppCountryOper(ctx context.Context) (any, error)
	GetApNhGlobalData(ctx context.Context) (any, error)
	GetApImagePrepareLocation(ctx context.Context) (any, error)
	GetApImageActiveLocation(ctx context.Context) (any, error)

	// Access Point Configuration
	GetApCfg(ctx context.Context) (any, error)
	GetTagSourcePriorityConfigs(ctx context.Context) (any, error)
	GetApTagSourcePriorityConfigs(ctx context.Context) (any, error)
	GetApApTags(ctx context.Context) (any, error)
}

// GeneralAPI defines the interface for general controller operations.
type GeneralAPI interface {
	// General Operational Data
	GetGeneralOper(ctx context.Context) (any, error)
	GetGeneralOperMgmtIntfData(ctx context.Context) (any, error)

	// General Configuration
	GetGeneralCfg(ctx context.Context) (any, error)
	GetGeneralMewlcConfig(ctx context.Context) (any, error)
	GetGeneralCacConfig(ctx context.Context) (any, error)
	GetGeneralMfp(ctx context.Context) (any, error)
	GetGeneralFipsCfg(ctx context.Context) (any, error)
	GetGeneralWsaApClientEvent(ctx context.Context) (any, error)
	GetGeneralSimL3InterfaceCacheData(ctx context.Context) (any, error)
	GetGeneralWlcManagementData(ctx context.Context) (any, error)
	GetGeneralLaginfo(ctx context.Context) (any, error)
	GetGeneralMulticastConfig(ctx context.Context) (any, error)
	GetGeneralFeatureUsageCfg(ctx context.Context) (any, error)
	GetGeneralThresholdWarnCfg(ctx context.Context) (any, error)
	GetGeneralApLocRangingCfg(ctx context.Context) (any, error)
	GetGeneralGeolocationCfg(ctx context.Context) (any, error)
}

// RadioResourceManagementAPI defines the interface for RRM operations.
type RadioResourceManagementAPI interface {
	// RRM Operational Data
	GetRrmOper(ctx context.Context) (any, error)
	GetApAutoRfDot11Data(ctx context.Context) (any, error)
	GetApDot11RadarData(ctx context.Context) (any, error)
	GetApDot11SpectrumData(ctx context.Context) (any, error)
	GetRrmMeasurement(ctx context.Context) (any, error)
	GetRadioSlot(ctx context.Context) (any, error)
	GetMainData(ctx context.Context) (any, error)
	GetSpectrumDeviceTable(ctx context.Context) (any, error)
	GetSpectrumAqTable(ctx context.Context) (any, error)
	GetRegDomainOper(ctx context.Context) (any, error)

	// RRM Configuration
	GetRrmCfg(ctx context.Context) (any, error)
	GetRrmRrms(ctx context.Context) (any, error)
	GetRrmMgrCfgEntries(ctx context.Context) (any, error)
}

// WirelessLANAPI defines the interface for WLAN configuration operations.
type WirelessLANAPI interface {
	GetWlanCfg(ctx context.Context) (any, error)
	GetWlanCfgEntries(ctx context.Context) (any, error)
	GetWlanPolicies(ctx context.Context) (any, error)
	GetPolicyListEntries(ctx context.Context) (any, error)
	GetWirelessAaaPolicyConfigs(ctx context.Context) (any, error)
}

// ClientAPI defines the interface for client operations.
type ClientAPI interface {
	GetClientOper(ctx context.Context) (any, error)
	GetClientOperCommonOperData(ctx context.Context) (any, error)
	GetClientOperDot11OperData(ctx context.Context) (any, error)
	GetClientOperMobilityOperData(ctx context.Context) (any, error)
	GetClientOperMmIfClientStats(ctx context.Context) (any, error)
	GetClientOperMmIfClientHistory(ctx context.Context) (any, error)
	GetClientOperTrafficStats(ctx context.Context) (any, error)
	GetClientOperPolicyData(ctx context.Context) (any, error)
	GetClientOperSisfDbMac(ctx context.Context) (any, error)
	GetClientOperDcInfo(ctx context.Context) (any, error)
}

// MobilityAPI defines the interface for mobility management operations.
type MobilityAPI interface {
	GetMobilityOper(ctx context.Context) (any, error)
	GetMobilityMmIfGlobalStats(ctx context.Context) (any, error)
	GetMobilityMmIfGlobalMsgStats(ctx context.Context) (any, error)
	GetMobilityGlobalStats(ctx context.Context) (any, error)
	GetMobilityMmGlobalData(ctx context.Context) (any, error)
	GetMobilityGlobalMsgStats(ctx context.Context) (any, error)
	GetMobilityClientData(ctx context.Context) (any, error)
	GetMobilityApCache(ctx context.Context) (any, error)
	GetMobilityApPeerList(ctx context.Context) (any, error)
	GetMobilityClientStats(ctx context.Context) (any, error)
	GetMobilityWlanClientLimit(ctx context.Context) (any, error)
	GetMobilityGlobalDTLSStats(ctx context.Context) (any, error)
}

// RogueAPI defines the interface for rogue detection operations.
type RogueAPI interface {
	GetRogueOper(ctx context.Context) (any, error)
	GetRogueStats(ctx context.Context) (any, error)
	GetRogueData(ctx context.Context) (any, error)
	GetRogueClientData(ctx context.Context) (any, error)
	GetRldpStats(ctx context.Context) (any, error)
}

// NetworkManagementAPI defines the interface for network management operations.
type NetworkManagementAPI interface {
	GetNmspOper(ctx context.Context) (any, error)
	GetNmspClientRegistration(ctx context.Context) (any, error)
	GetNmspCmxConnection(ctx context.Context) (any, error)
	GetNmspCmxCloudInfo(ctx context.Context) (any, error)
}

// HyperlocationAPI defines the interface for hyperlocation operations.
type HyperlocationAPI interface {
	GetHyperlocationOper(ctx context.Context) (any, error)
	GetHyperlocationProfiles(ctx context.Context) (any, error)
}

// AWIPSApi defines the interface for AWIPS operations.
type AWIPSApi interface {
	GetAwipsOper(ctx context.Context) (any, error)
	GetAwipsPerApInfo(ctx context.Context) (any, error)
	GetAwipsDwldStatus(ctx context.Context) (any, error)
	GetAwipsApDwldStatus(ctx context.Context) (any, error)
}

// GeolocationAPI defines the interface for geolocation operations.
type GeolocationAPI interface {
	GetGeolocationOper(ctx context.Context) (any, error)
	GetGeolocationOperApGeoLocStats(ctx context.Context) (any, error)
}

// AFCApi defines the interface for AFC operations.
type AFCApi interface {
	GetAfcOper(ctx context.Context) (any, error)
	GetAfcEwlcAfcApResp(ctx context.Context) (any, error)
}

// mDNSAPI defines the interface for mDNS operations.
type mDNSAPI interface {
	GetMdnsOper(ctx context.Context) (any, error)
	GetMdnsGlobalStats(ctx context.Context) (any, error)
	GetMdnsWlanStats(ctx context.Context) (any, error)
}

// MultcastAPI defines the interface for multicast operations.
type MultcastAPI interface {
	GetMcastOper(ctx context.Context) (any, error)
	GetMcastFlexMediastreamClientSummary(ctx context.Context) (any, error)
	GetMcastVlanL2MgidOp(ctx context.Context) (any, error)
}

// BluetoothAPI defines the interface for Bluetooth Low Energy operations.
type BluetoothAPI interface {
	GetBleLtxOper(ctx context.Context) (any, error)
	GetBleLtxApAntenna(ctx context.Context) (any, error)
	GetBleLtxAp(ctx context.Context) (any, error)
}

// LISPApi defines the interface for LISP operations.
type LISPApi interface {
	GetLispAgentOper(ctx context.Context) (any, error)
	GetLispAgentMemoryStats(ctx context.Context) (any, error)
	GetLispWlcCapabilities(ctx context.Context) (any, error)
	GetLispApCapabilities(ctx context.Context) (any, error)
}

// SiteAPI defines the interface for site configuration operations.
type SiteAPI interface {
	GetSiteCfg(ctx context.Context) (any, error)
	GetSiteApCfgProfiles(ctx context.Context) (any, error)
	GetSiteTagConfigs(ctx context.Context) (any, error)
}

// Dot11API defines the interface for 802.11 configuration operations.
type Dot11API interface {
	GetDot11Cfg(ctx context.Context) (any, error)
	GetDot11ConfiguredCountries(ctx context.Context) (any, error)
	GetDot11acMcsEntries(ctx context.Context) (any, error)
	GetDot11Entries(ctx context.Context) (any, error)
}

// ApfAPI defines the interface for APF configuration operations.
type ApfAPI interface {
	GetApfCfg(ctx context.Context) (any, error)
	GetApf(ctx context.Context) (any, error)
}

// RfAPI defines the interface for RF configuration operations.
type RfAPI interface {
	GetRfCfg(ctx context.Context) (any, error)
	GetRfMultiBssidProfiles(ctx context.Context) (any, error)
	GetRfAtfPolicies(ctx context.Context) (any, error)
	GetRfTags(ctx context.Context) (any, error)
	GetRfProfiles(ctx context.Context) (any, error)
	GetRfProfileDefaultEntries(ctx context.Context) (any, error)
}

// Dot15API defines the interface for 802.15 configuration operations.
type Dot15API interface {
	GetDot15Cfg(ctx context.Context) (any, error)
	GetDot15GlobalConfig(ctx context.Context) (any, error)
}

// CTSAPI defines the interface for CTS configuration operations.
type CTSAPI interface {
	GetCtsSxpCfg(ctx context.Context) (any, error)
	GetCtsSxpConfiguration(ctx context.Context) (any, error)
}

// LocationAPI defines the interface for location configuration operations.
type LocationAPI interface {
	GetLocationCfg(ctx context.Context) (any, error)
	GetLocationNmspConfig(ctx context.Context) (any, error)
}

// RadioAPI defines the interface for radio configuration operations.
type RadioAPI interface {
	GetRadioCfg(ctx context.Context) (any, error)
	GetRadioProfiles(ctx context.Context) (any, error)
}

// MeshAPI defines the interface for mesh configuration operations.
type MeshAPI interface {
	GetMeshCfg(ctx context.Context) (any, error)
	GetMesh(ctx context.Context) (any, error)
	GetMeshProfiles(ctx context.Context) (any, error)
}

// FlexAPI defines the interface for Flex configuration operations.
type FlexAPI interface {
	GetFlexCfg(ctx context.Context) (any, error)
	GetFlexCfgData(ctx context.Context) (any, error)
}

// FabricAPI defines the interface for Fabric configuration operations.
type FabricAPI interface {
	GetFabricCfg(ctx context.Context) (any, error)
	GetFabricControlplaneNames(ctx context.Context) (any, error)
	GetFabric(ctx context.Context) (any, error)
}

// RFIDAPI defines the interface for RFID configuration operations.
type RFIDAPI interface {
	GetRfidCfg(ctx context.Context) (any, error)
}

// WNCClient defines the core interface for the Cisco Wireless Network Controller API client.
// Individual feature packages (ap, client, rrm, wlan, etc.) extend this client with specific functionality.
// This interface is kept for backward compatibility and basic operations.
type WNCClient interface {
	CoreAPI
}

// Client represents a WNC API client with configuration and logging capabilities.
// This is the main client structure used to interact with the Cisco Wireless Network Controller.
// It implements the comprehensive WirelessControllerAPI interface providing unified access to all features.
type Client struct {
	controller         string        // WNC hostname or IP address
	accessToken        string        // Authentication token for API access
	timeout            time.Duration // Request timeout duration
	insecureSkipVerify bool          // Whether to skip TLS certificate verification
	logger             *slog.Logger  // Logger instance for debugging and monitoring
}
