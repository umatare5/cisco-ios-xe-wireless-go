package ap

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Service provides access point operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new AP service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the complete AP configuration data.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessAPCfg, error) {
	return core.Get[CiscoIOSXEWirelessAPCfg](ctx, s.Client(), routes.APCfgPath, opts...)
}

// ListTagConfigs retrieves access point tag configurations.
func (s Service) ListTagConfigs(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessApCfgApTags, error) {
	return core.Get[CiscoIOSXEWirelessApCfgApTags](ctx, s.Client(), routes.APTagsPath, opts...)
}

// GetTagConfigByMAC retrieves AP tag configuration filtered by AP MAC address.
func (s Service) GetTagConfigByMAC(
	ctx context.Context,
	mac string,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApCfgApTag, error) {
	if err := validation.ValidateMACAddress(mac); err != nil {
		return nil, fmt.Errorf(ErrInvalidAPMacFormat, mac)
	}
	normalizedMAC, err := validation.NormalizeMACAddress(mac)
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidAPMacFormat, mac)
	}

	// Build correct RESTCONF path: /ap-cfg-data/ap-tags/ap-tag=MAC
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APTagQueryPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApCfgApTag](ctx, s.Client(), url, opts...)
}

// ListTagSourcePriorityConfigs retrieves tag source priority configurations.
func (s Service) ListTagSourcePriorityConfigs(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApCfgTagSourcePriorityConfigs, error) {
	return core.Get[CiscoIOSXEWirelessApCfgTagSourcePriorityConfigs](
		ctx,
		s.Client(),
		routes.APTagSourcePriorityConfigsPath,
		opts...,
	)
}

// GetGlobalOperational retrieves the complete AP global operational data.
func (s Service) GetGlobalOperational(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessAPGlobalOper, error) {
	return core.Get[CiscoIOSXEWirelessAPGlobalOper](ctx, s.Client(), routes.APGlobalOperPath, opts...)
}

// GetEWLCAPStats retrieves EWLC AP statistics.
func (s Service) GetEWLCAPStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperEwlcApStats, error) {
	return core.Get[CiscoIOSXEWirelessApGlobalOperEwlcApStats](ctx, s.Client(), routes.APEwlcApStatsPath, opts...)
}

// ListAPHistoryByEthernetMAC retrieves AP history data filtered by ethernet MAC address.
func (s Service) ListAPHistoryByEthernetMAC(
	ctx context.Context,
	ethernetMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperApHistory, error) {
	normalizedMAC, err := service.RequireMACAddress(ethernetMAC)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APHistoryQueryPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApGlobalOperApHistory](ctx, s.Client(), url, opts...)
}

// GetAPJoinStatsByWTPMAC retrieves AP join statistics filtered by WTP MAC address.
func (s Service) GetAPJoinStatsByWTPMAC(
	ctx context.Context, mac string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperApJoinStats, error) {
	if mac == "" {
		return nil, core.ErrResourceNotFound
	}
	if strings.TrimSpace(mac) == "" {
		return nil, core.ErrResourceNotFound
	}

	normalizedMAC, err := validation.NormalizeMACAddress(mac)
	if err != nil {
		return nil, fmt.Errorf("invalid MAC address %s: %w", mac, err)
	}

	// Build URL with RESTCONF path for ap-join-stats filtered by wtp-mac
	url := s.Client().RESTCONFBuilder().BuildQueryURL(
		routes.APJoinStatsPath,
		normalizedMAC,
	)
	return core.Get[CiscoIOSXEWirelessApGlobalOperApJoinStats](ctx, s.Client(), url, opts...)
}

// GetWLANClientStatsByWLANID retrieves WLAN client statistics filtered by WLAN ID.
func (s Service) GetWLANClientStatsByWLANID(
	ctx context.Context,
	wlanID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperWlanClientStats, error) {
	if wlanID <= 0 {
		return nil, core.ErrResourceNotFound
	}
	url := s.Client().RESTCONFBuilder().BuildQueryURL(
		routes.APWlanClientStatsQueryPath,
		strconv.Itoa(wlanID),
	)
	return core.Get[CiscoIOSXEWirelessApGlobalOperWlanClientStats](ctx, s.Client(), url, opts...)
}

// ListAPHistory retrieves only AP history data using fields parameter.
func (s Service) ListAPHistory(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperApHistory, error) {
	return core.Get[CiscoIOSXEWirelessApGlobalOperApHistory](ctx, s.Client(), routes.APHistoryPath, opts...)
}

// ListAPJoinStats retrieves only AP join statistics using fields parameter.
func (s Service) ListAPJoinStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperApJoinStats, error) {
	return core.Get[CiscoIOSXEWirelessApGlobalOperApJoinStats](ctx, s.Client(), routes.APJoinStatsPath, opts...)
}

// ListWLANClientStats retrieves only WLAN client statistics using fields parameter.
func (s Service) ListWLANClientStats(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApGlobalOperWlanClientStats, error) {
	return core.Get[CiscoIOSXEWirelessApGlobalOperWlanClientStats](
		ctx,
		s.Client(),
		routes.APWlanClientStatsPath,
		opts...,
	)
}

// GetOperational retrieves the complete AP operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessAPOper, error) {
	return core.Get[CiscoIOSXEWirelessAPOper](ctx, s.Client(), routes.APOperPath, opts...)
}

// ListApOperData retrieves AP operational data.
func (s Service) ListApOperData(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessApOperData, error) {
	return core.Get[CiscoIOSXEWirelessApOperData](ctx, s.Client(), routes.APOperDataPath, opts...)
}

// ListCAPWAPData retrieves CAPWAP protocol data.
func (s Service) ListCAPWAPData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCAPWAPData, error) {
	return core.Get[CiscoIOSXEWirelessApOperCAPWAPData](ctx, s.Client(), routes.APCapwapDataPath, opts...)
}

// GetCAPWAPDataByWTPMAC retrieves CAPWAP data for a specific WTP MAC.
func (s Service) GetCAPWAPDataByWTPMAC(
	ctx context.Context,
	wtpMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCAPWAPData, error) {
	normalizedMAC, err := service.RequireMACAddress(wtpMAC)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APCapwapDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperCAPWAPData](ctx, s.Client(), url, opts...)
}

// ListNameMACMaps retrieves AP name-to-MAC mapping data.
func (s Service) ListNameMACMaps(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApNameMACMap, error) {
	return core.Get[CiscoIOSXEWirelessApOperApNameMACMap](ctx, s.Client(), routes.APApNameMACMapPath, opts...)
}

// GetNameMACMapByWTPName retrieves AP name-to-MAC mapping filtered by WTP name.
func (s Service) GetNameMACMapByWTPName(
	ctx context.Context,
	wtpName string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApNameMACMap, error) {
	if err := service.RequireAPName(wtpName); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APApNameMACMapPath, wtpName)
	return core.Get[CiscoIOSXEWirelessApOperApNameMACMap](ctx, s.Client(), url, opts...)
}

// ListRadioData retrieves radio operational data.
func (s Service) ListRadioData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioOperData, error) {
	return core.Get[CiscoIOSXEWirelessApOperRadioOperData](ctx, s.Client(), routes.APRadioOperDataPath, opts...)
}

// GetRadioStatusByWTPMACAndSlot retrieves radio operational data by WTP MAC and slot ID.
func (s Service) GetRadioStatusByWTPMACAndSlot(
	ctx context.Context, wtpMAC string, slotID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioOperData, error) {
	normalizedMAC, err := service.RequireMACAddress(wtpMAC)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(routes.APRadioOperDataPath, normalizedMAC, slotID)
	return core.Get[CiscoIOSXEWirelessApOperRadioOperData](ctx, s.Client(), url, opts...)
}

// ListRadioNeighbors retrieves all AP radio neighbor information.
func (s Service) ListRadioNeighbors(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApRadioNeighbor, error) {
	return core.Get[CiscoIOSXEWirelessApOperApRadioNeighbor](ctx, s.Client(), routes.APRadioNeighborPath, opts...)
}

// GetRadioNeighborByAPMACSlotAndBSSID retrieves AP radio neighbor information for a specific AP MAC, slot ID and BSSID.
// This follows the YANG model key structure: "ap-mac slot-id bssid".
func (s Service) GetRadioNeighborByAPMACSlotAndBSSID(
	ctx context.Context, apMAC string, slotID int, bssid string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApRadioNeighbor, error) {
	normalizedAPMAC, err := service.RequireMACAddress(apMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid AP MAC address: %w", err)
	}
	normalizedBSSID, err := service.RequireMACAddress(bssid)
	if err != nil {
		return nil, fmt.Errorf("invalid BSSID: %w", err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APRadioNeighborPath,
		normalizedAPMAC,
		slotID,
		normalizedBSSID,
	)
	return core.Get[CiscoIOSXEWirelessApOperApRadioNeighbor](ctx, s.Client(), url, opts...)
}

// ListActiveImageLocations retrieves active image location information using fields parameter.
func (s Service) ListActiveImageLocations(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApImageActiveLocation, error) {
	return core.Get[CiscoIOSXEWirelessApOperApImageActiveLocation](
		ctx,
		s.Client(),
		routes.APImageActiveLocationPath,
		opts...,
	)
}

// ListPreparedImageLocations retrieves only AP image prepare location data using fields parameter.
func (s Service) ListPreparedImageLocations(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApImagePrepareLocation, error) {
	return core.Get[CiscoIOSXEWirelessApOperApImagePrepareLocation](
		ctx,
		s.Client(),
		routes.APImagePrepareLocationPath,
		opts...,
	)
}

// ListPowerInfo retrieves only AP power information using fields parameter.
func (s Service) ListPowerInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApPwrInfo, error) {
	return core.Get[CiscoIOSXEWirelessApOperApPwrInfo](ctx, s.Client(), routes.APPwrInfoPath, opts...)
}

// ListSensorStatus retrieves only AP sensor status using fields parameter.
func (s Service) ListSensorStatus(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApSensorStatus, error) {
	return core.Get[CiscoIOSXEWirelessApOperApSensorStatus](ctx, s.Client(), routes.APSensorStatusPath, opts...)
}

// ListCAPWAPPackets retrieves only CAPWAP packets data using fields parameter.
func (s Service) ListCAPWAPPackets(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCAPWAPPkts, error) {
	return core.Get[CiscoIOSXEWirelessApOperCAPWAPPkts](ctx, s.Client(), routes.APCapwapPktsPath, opts...)
}

// ListIotFirmware retrieves IoT firmware information for all access points.
func (s Service) ListIotFirmware(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperIotFirmware, error) {
	return core.Get[CiscoIOSXEWirelessApOperIotFirmware](ctx, s.Client(), routes.APIotFirmwarePath, opts...)
}

// ListRadioResetStats retrieves radio reset statistics for all access points.
func (s Service) ListRadioResetStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioResetStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperRadioResetStats](ctx, s.Client(), routes.APRadioResetStatsPath, opts...)
}

// GetRadioResetStatsByAPMACAndRadioID retrieves radio reset statistics for a specific AP MAC and radio ID.
func (s Service) GetRadioResetStatsByAPMACAndRadioID(
	ctx context.Context, apMAC string, radioID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioResetStats, error) {
	if apMAC == "" || strings.TrimSpace(apMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(apMAC); err != nil {
		return nil, fmt.Errorf("invalid AP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid AP MAC address %s: %w", apMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APRadioResetStatsPath,
		normalizedMAC,
		radioID,
	)
	return core.Get[CiscoIOSXEWirelessApOperRadioResetStats](ctx, s.Client(), url, opts...)
}

// ListQosClientData retrieves QoS client data for all access points.
func (s Service) ListQosClientData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperQosClientData, error) {
	return core.Get[CiscoIOSXEWirelessApOperQosClientData](ctx, s.Client(), routes.APQosClientDataPath, opts...)
}

// GetQosClientDataByClientMAC retrieves QoS client data for a specific client MAC address.
func (s Service) GetQosClientDataByClientMAC(
	ctx context.Context, clientMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperQosClientData, error) {
	if clientMAC == "" || strings.TrimSpace(clientMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(clientMAC); err != nil {
		return nil, fmt.Errorf("invalid client MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(clientMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid client MAC address %s: %w", clientMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APQosClientDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperQosClientData](ctx, s.Client(), url, opts...)
}

// ListWtpSlotWlanStats retrieves WTP slot WLAN statistics for all access points.
func (s Service) ListWtpSlotWlanStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperWtpSlotWlanStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperWtpSlotWlanStats](ctx, s.Client(), routes.APWtpSlotWlanStatsPath, opts...)
}

// GetWtpSlotWlanStatsByWTPMACSlotAndWLANID retrieves WTP slot WLAN statistics for a specific WTP MAC, slot ID, and WLAN ID.
func (s Service) GetWtpSlotWlanStatsByWTPMACSlotAndWLANID(
	ctx context.Context,
	wtpMAC string,
	slotID int,
	wlanID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperWtpSlotWlanStats, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APWtpSlotWlanStatsPath,
		normalizedMAC,
		slotID,
		wlanID,
	)
	return core.Get[CiscoIOSXEWirelessApOperWtpSlotWlanStats](ctx, s.Client(), url, opts...)
}

// ListEthernetMACWtpMACMaps retrieves Ethernet MAC to WTP MAC mapping for all access points.
func (s Service) ListEthernetMACWtpMACMaps(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEthernetMACWtpMACMap, error) {
	return core.Get[CiscoIOSXEWirelessApOperEthernetMACWtpMACMap](
		ctx,
		s.Client(),
		routes.APEthernetMACWtpMACMapPath,
		opts...,
	)
}

// GetEthernetMACWtpMACMapByEthernetMAC retrieves Ethernet MAC to WTP MAC mapping for a specific Ethernet MAC address.
func (s Service) GetEthernetMACWtpMACMapByEthernetMAC(
	ctx context.Context, ethernetMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEthernetMACWtpMACMap, error) {
	if ethernetMAC == "" || strings.TrimSpace(ethernetMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(ethernetMAC); err != nil {
		return nil, fmt.Errorf("invalid Ethernet MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(ethernetMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid Ethernet MAC address %s: %w", ethernetMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APEthernetMACWtpMACMapPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperEthernetMACWtpMACMap](ctx, s.Client(), url, opts...)
}

// ListRadioOperStats retrieves radio operational statistics for all access points.
func (s Service) ListRadioOperStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioOperStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperRadioOperStats](ctx, s.Client(), routes.APRadioOperStatsPath, opts...)
}

// GetRadioOperStatsByWTPMACAndSlot retrieves radio operational statistics for a specific WTP MAC and slot ID.
func (s Service) GetRadioOperStatsByWTPMACAndSlot(
	ctx context.Context,
	wtpMAC string,
	slotID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRadioOperStats, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APRadioOperStatsPath,
		normalizedMAC,
		slotID,
	)
	return core.Get[CiscoIOSXEWirelessApOperRadioOperStats](ctx, s.Client(), url, opts...)
}

// ListEthernetIfStats retrieves Ethernet interface statistics for all access points.
func (s Service) ListEthernetIfStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEthernetIfStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperEthernetIfStats](ctx, s.Client(), routes.APEthernetIfStatsPath, opts...)
}

// GetEthernetIfStatsByWTPMACAndInterfaceID retrieves Ethernet interface statistics for a specific WTP MAC and interface ID.
func (s Service) GetEthernetIfStatsByWTPMACAndInterfaceID(
	ctx context.Context,
	wtpMAC string,
	interfaceID string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEthernetIfStats, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}
	if interfaceID == "" || strings.TrimSpace(interfaceID) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APEthernetIfStatsPath,
		normalizedMAC,
		interfaceID,
	)
	return core.Get[CiscoIOSXEWirelessApOperEthernetIfStats](ctx, s.Client(), url, opts...)
}

// ListEwlcWncdStats retrieves EWLC WNCD statistics information.
func (s Service) ListEwlcWncdStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEwlcWncdStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperEwlcWncdStats](ctx, s.Client(), routes.APEwlcWncdStatsPath, opts...)
}

// ListApIoxOperData retrieves AP IOx operational data for all access points.
func (s Service) ListApIoxOperData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApIoxOperData, error) {
	return core.Get[CiscoIOSXEWirelessApOperApIoxOperData](ctx, s.Client(), routes.APApIoxOperDataPath, opts...)
}

// GetApIoxOperDataByWTPMAC retrieves AP IOx operational data for a specific WTP MAC address.
func (s Service) GetApIoxOperDataByWTPMAC(
	ctx context.Context, wtpMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApIoxOperData, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APApIoxOperDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperApIoxOperData](ctx, s.Client(), url, opts...)
}

// ListQosGlobalStats retrieves QoS global statistics information.
func (s Service) ListQosGlobalStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperQosGlobalStats, error) {
	return core.Get[CiscoIOSXEWirelessApOperQosGlobalStats](ctx, s.Client(), routes.APQosGlobalStatsPath, opts...)
}

// ListRlanOper retrieves RLAN operational data for all access points.
func (s Service) ListRlanOper(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessApOperRlanOper, error) {
	return core.Get[CiscoIOSXEWirelessApOperRlanOper](ctx, s.Client(), routes.APRlanOperPath, opts...)
}

// GetRlanOperByWTPMACAndPortID retrieves RLAN operational data for a specific WTP MAC and port ID.
func (s Service) GetRlanOperByWTPMACAndPortID(
	ctx context.Context, wtpMAC string, portID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperRlanOper, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APRlanOperPath,
		normalizedMAC,
		portID,
	)
	return core.Get[CiscoIOSXEWirelessApOperRlanOper](ctx, s.Client(), url, opts...)
}

// ListEwlcMewlcPredownloadRec retrieves EWLC MEWLC predownload record information.
// Note: Available on 17.12.6a, but unavailable on 17.15.4b.
func (s Service) ListEwlcMewlcPredownloadRec(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperEwlcMewlcPredownloadRec, error) {
	return core.Get[CiscoIOSXEWirelessApOperEwlcMewlcPredownloadRec](
		ctx,
		s.Client(),
		routes.APEwlcMewlcPredownloadRecPath, opts...,
	)
}

// ListCdpCacheData retrieves CDP cache data for all access points.
func (s Service) ListCdpCacheData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCdpCacheData, error) {
	return core.Get[CiscoIOSXEWirelessApOperCdpCacheData](ctx, s.Client(), routes.APCdpCacheDataPath, opts...)
}

// GetCdpCacheDataByWTPMAC retrieves CDP cache data for a specific WTP MAC address.
func (s Service) GetCdpCacheDataByWTPMAC(
	ctx context.Context, wtpMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCdpCacheData, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APCdpCacheDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperCdpCacheData](ctx, s.Client(), url, opts...)
}

// ListLldpNeigh retrieves LLDP neighbor information for all access points.
func (s Service) ListLldpNeigh(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperLldpNeigh, error) {
	return core.Get[CiscoIOSXEWirelessApOperLldpNeigh](ctx, s.Client(), routes.APLldpNeighPath, opts...)
}

// GetLldpNeighByWTPMAC retrieves LLDP neighbor information for a specific WTP MAC address.
func (s Service) GetLldpNeighByWTPMAC(
	ctx context.Context, wtpMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperLldpNeigh, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APLldpNeighPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperLldpNeigh](ctx, s.Client(), url, opts...)
}

// ListTpCertInfo retrieves trustpoint certificate info information.
func (s Service) ListTpCertInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperTpCertInfo, error) {
	return core.Get[CiscoIOSXEWirelessApOperTpCertInfo](ctx, s.Client(), routes.APTpCertInfoPath, opts...)
}

// ListDiscData retrieves discovery data for all access points.
func (s Service) ListDiscData(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessApOperDiscData, error) {
	return core.Get[CiscoIOSXEWirelessApOperDiscData](ctx, s.Client(), routes.APDiscDataPath, opts...)
}

// GetDiscDataByWTPMAC retrieves discovery data for a specific WTP MAC address.
func (s Service) GetDiscDataByWTPMAC(
	ctx context.Context, wtpMAC string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperDiscData, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APDiscDataPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessApOperDiscData](ctx, s.Client(), url, opts...)
}

// ListCountryOper retrieves country operational data for all access points.
func (s Service) ListCountryOper(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCountryOper, error) {
	return core.Get[CiscoIOSXEWirelessApOperCountryOper](ctx, s.Client(), routes.APCountryOperPath, opts...)
}

// GetCountryOperByWTPMACAndRadioID retrieves country operational data for a specific WTP MAC and radio ID.
func (s Service) GetCountryOperByWTPMACAndRadioID(
	ctx context.Context, wtpMAC string, radioID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperCountryOper, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APCountryOperPath,
		normalizedMAC,
		radioID,
	)
	return core.Get[CiscoIOSXEWirelessApOperCountryOper](ctx, s.Client(), url, opts...)
}

// ListSuppCountryOper retrieves supported country operational data for all access points.
func (s Service) ListSuppCountryOper(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperSuppCountryOper, error) {
	return core.Get[CiscoIOSXEWirelessApOperSuppCountryOper](ctx, s.Client(), routes.APSuppCountryOperPath, opts...)
}

// GetSuppCountryOperByWTPMACAndRadioID retrieves supported country operational data for a specific WTP MAC and radio ID.
func (s Service) GetSuppCountryOperByWTPMACAndRadioID(
	ctx context.Context, wtpMAC string, radioID int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperSuppCountryOper, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	if err := validation.ValidateMACAddress(wtpMAC); err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address: %w", err)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(wtpMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid WTP MAC address %s: %w", wtpMAC, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APSuppCountryOperPath,
		normalizedMAC,
		radioID,
	)
	return core.Get[CiscoIOSXEWirelessApOperSuppCountryOper](ctx, s.Client(), url, opts...)
}

// ListApNhGlobalData retrieves AP neighborhood global data.
func (s Service) ListApNhGlobalData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessApOperApNhGlobalData, error) {
	return core.Get[CiscoIOSXEWirelessApOperApNhGlobalData](ctx, s.Client(), routes.APApNhGlobalDataPath, opts...)
}

// EnableAPByMAC enables the administrative state of the access point with this address.
//
// The By suffix names the arm of the RPC's mandatory choice this fills rather than a list key,
// which is why the tag writes on this service carry no suffix: ap-mac is their list key and there
// is nothing to choose.
func (s Service) EnableAPByMAC(ctx context.Context, apMAC string) error {
	return s.setAPAdminStateByMAC(ctx, apMAC, true)
}

// DisableAPByMAC disables the administrative state of the access point with this address.
func (s Service) DisableAPByMAC(ctx context.Context, apMAC string) error {
	return s.setAPAdminStateByMAC(ctx, apMAC, false)
}

// EnableAPByName enables the administrative state of the access point with this name.
func (s Service) EnableAPByName(ctx context.Context, apName string) error {
	return s.setAPAdminStateByName(ctx, apName, true)
}

// DisableAPByName disables the administrative state of the access point with this name.
func (s Service) DisableAPByName(ctx context.Context, apName string) error {
	return s.setAPAdminStateByName(ctx, apName, false)
}

// EnableRadioByMAC enables one radio on the access point with this address.
//
// slotID is the radio-slot-id the controller reports for that radio and radioType is the
// radio-type leaf of the same record. The RPC's band follows the radio the slot holds rather than
// the band it is serving, so both come from the record and neither is derived from the other;
// whether the controller accepts the pair is the controller's to answer.
func (s Service) EnableRadioByMAC(ctx context.Context, apMAC string, slotID int, radioType RadioType) error {
	return s.setRadioAdminStateByMAC(ctx, apMAC, slotID, radioType, true)
}

// DisableRadioByMAC disables one radio on the access point with this address.
func (s Service) DisableRadioByMAC(ctx context.Context, apMAC string, slotID int, radioType RadioType) error {
	return s.setRadioAdminStateByMAC(ctx, apMAC, slotID, radioType, false)
}

// EnableRadioByName enables one radio on the access point with this name.
func (s Service) EnableRadioByName(ctx context.Context, apName string, slotID int, radioType RadioType) error {
	return s.setRadioAdminStateByName(ctx, apName, slotID, radioType, true)
}

// DisableRadioByName disables one radio on the access point with this name.
func (s Service) DisableRadioByName(ctx context.Context, apName string, slotID int, radioType RadioType) error {
	return s.setRadioAdminStateByName(ctx, apName, slotID, radioType, false)
}

// AssignSiteTag assigns a site tag to an Access Point using MAC address.
func (s Service) AssignSiteTag(ctx context.Context, apMAC, siteTag string) error {
	if !validation.IsValidTagAssignment(siteTag, "site") {
		return ierrors.RequiredParameterError("site tag")
	}
	tags := ApTag{SiteTag: siteTag}
	return s.assignTags(ctx, apMAC, tags)
}

// AssignPolicyTag assigns a policy tag to an Access Point using MAC address.
func (s Service) AssignPolicyTag(ctx context.Context, apMAC, policyTag string) error {
	if !validation.IsValidTagAssignment(policyTag, "policy") {
		return ierrors.RequiredParameterError("policy tag")
	}
	tags := ApTag{PolicyTag: policyTag}
	return s.assignTags(ctx, apMAC, tags)
}

// AssignRFTag assigns an RF tag to an Access Point using MAC address.
func (s Service) AssignRFTag(ctx context.Context, apMAC, rfTag string) error {
	if !validation.IsValidTagAssignment(rfTag, "rf") {
		return ierrors.RequiredParameterError("RF tag")
	}
	tags := ApTag{RFTag: rfTag}
	return s.assignTags(ctx, apMAC, tags)
}

// ReloadByMAC restarts the access point with this address, interrupting service on it.
//
// The RPC accepts a name or an address and this resolves the name, because the address arm has
// never been observed to complete. The resolve is the keyed capwap-data read and carries no
// fields expression: the record is one row either way, and pruning to name alone would drop the
// wtp-mac this address was matched on.
func (s Service) ReloadByMAC(ctx context.Context, apMAC string) error {
	normalizedMAC, err := service.RequireMACAddress(apMAC)
	if err != nil {
		return err
	}

	resp, err := s.GetCAPWAPDataByWTPMAC(ctx, normalizedMAC)
	if core.IsNotFoundError(err) {
		return fmt.Errorf(ErrAPNotFoundByMAC, apMAC)
	}
	if err != nil {
		return fmt.Errorf(ErrFailedGetCAPWAPData, err)
	}
	if resp == nil {
		return errors.New(ErrCAPWAPDataUnavailable)
	}
	if len(resp.CAPWAPData) == 0 {
		return fmt.Errorf(ErrAPNotFoundByMAC, apMAC)
	}

	return s.ReloadByName(ctx, resp.CAPWAPData[0].Name)
}

// ReloadByName restarts the access point with this name, interrupting service on it.
func (s Service) ReloadByName(ctx context.Context, apName string) error {
	if err := service.RequireAPName(apName); err != nil {
		return err
	}

	payload := APReloadRPCPayload{Input: APReloadRPCInput{APName: apName}}

	return core.PostRPCVoid(ctx, s.Client(), routes.APApResetRPC, payload)
}

// ResetCAPWAPByMAC tears down and re-establishes the CAPWAP session of the access point with this
// address, which does not reboot it.
//
// The address arm is what the RPC declares and what this sends; no completed write through it is
// on record, only the controller naming mac-addr in its own refusal of an input carrying neither
// arm. ResetCAPWAPByName is the arm that has been measured.
func (s Service) ResetCAPWAPByMAC(ctx context.Context, apMAC string) error {
	normalizedMAC, err := service.RequireMACAddress(apMAC)
	if err != nil {
		return err
	}

	return s.resetCAPWAP(ctx, APCAPWAPResetRPCInput{MACAddr: normalizedMAC})
}

// ResetCAPWAPByName tears down and re-establishes the CAPWAP session of the access point with this
// name, which does not reboot it.
//
// Measured on 17.12.8 and again on 17.18: the access point leaves within five seconds and rejoins
// within ten, and boot-time moves by about a second because the controller re-derives it from the
// access point's uptime at join.
func (s Service) ResetCAPWAPByName(ctx context.Context, apName string) error {
	if err := service.RequireAPName(apName); err != nil {
		return err
	}

	return s.resetCAPWAP(ctx, APCAPWAPResetRPCInput{APName: apName})
}

// resetCAPWAP posts one CAPWAP-reset RPC. The arm is the caller's: the input's choice is
// mandatory, and an input carrying neither is refused with 400.
func (s Service) resetCAPWAP(ctx context.Context, input APCAPWAPResetRPCInput) error {
	payload := APCAPWAPResetRPCPayload{Input: input}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.APSetRadCAPWAPResetRPC, payload); err != nil {
		return ierrors.ServiceOperationError("reset", "AP", "CAPWAP session", err)
	}

	return nil
}

// setAPAdminStateByMAC fills the input's mac-addr arm.
func (s Service) setAPAdminStateByMAC(ctx context.Context, apMAC string, enabled bool) error {
	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return fmt.Errorf(ErrInvalidAPMacFormat, apMAC)
	}

	return s.updateAPState(ctx, APConfigRPCInput{
		Mode:    core.GetAdminStateMode(enabled),
		MACAddr: normalizedMAC,
	})
}

// setAPAdminStateByName fills the input's ap-name arm.
func (s Service) setAPAdminStateByName(ctx context.Context, apName string, enabled bool) error {
	if err := service.RequireAPName(apName); err != nil {
		return err
	}

	return s.updateAPState(ctx, APConfigRPCInput{
		Mode:   core.GetAdminStateMode(enabled),
		APName: apName,
	})
}

// updateAPState posts one admin-state RPC. The arm is the caller's: the input's choice is
// mandatory, and an input carrying both names two access points.
func (s Service) updateAPState(ctx context.Context, input APConfigRPCInput) error {
	payload := APConfigRPCPayload{Input: input}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.APSetApAdminStateRPC, payload); err != nil {
		return ierrors.ServiceOperationError("update", "AP", "admin state", err)
	}

	return nil
}

// The band numbers set-ap-slot-admin-state declares, identical on 17.12, 17.15 and 17.18.
const (
	radioBand24GHz = 1
	radioBand5GHz  = 2
	radioBandDual  = 3
	radioBand6GHz  = 4
)

// radioBandNumber returns the band number set-ap-slot-admin-state takes for a radio type, and
// false when this SDK has no number for it.
//
// The RPC's band leaf names a kind of radio and not a frequency — 1 is 2.4 GHz, 2 is 5 GHz, 3 is
// dual band and 4 is 6 GHz — and its must clause admits seven (band, slot-id) pairs. Which of
// them a given access point accepts is the controller's to arbitrate, so nothing here refuses a
// slot: a dual-band radio in slot 0 takes band 3 where the served band would derive band 1.
//
// radio-invalid, radio-uwb and radio-remote-lan have no number because the leaf's domain is 1..4
// and none of the four names them. radio-80211-xor-24-6ghz has none either: both the dual-band 3
// and the 6 GHz 4 fit a 2.4/6 GHz XOR radio and no controller has been asked which.
func radioBandNumber(radioType RadioType) (int, bool) {
	switch radioType {
	case RadioType80211BG:
		return radioBand24GHz, true
	case RadioType80211A:
		return radioBand5GHz, true
	case RadioType80211ABGN, RadioTypeXOR5And6GHz:
		return radioBandDual, true
	case RadioType6GHz:
		return radioBand6GHz, true
	default:
		return 0, false
	}
}

// setRadioAdminStateByMAC fills the input's mac-addr arm.
func (s Service) setRadioAdminStateByMAC(
	ctx context.Context,
	apMAC string, slotID int, radioType RadioType, enabled bool,
) error {
	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return fmt.Errorf("invalid AP MAC address %s: %w", apMAC, err)
	}

	input, err := buildSlotAdminInput(slotID, radioType, enabled)
	if err != nil {
		return err
	}
	input.MACAddr = normalizedMAC

	return s.updateRadioState(ctx, input)
}

// setRadioAdminStateByName fills the input's ap-name arm.
func (s Service) setRadioAdminStateByName(
	ctx context.Context,
	apName string, slotID int, radioType RadioType, enabled bool,
) error {
	if err := service.RequireAPName(apName); err != nil {
		return err
	}

	input, err := buildSlotAdminInput(slotID, radioType, enabled)
	if err != nil {
		return err
	}
	input.APName = apName

	return s.updateRadioState(ctx, input)
}

// buildSlotAdminInput fills every leaf but the arm, which the caller sets.
func buildSlotAdminInput(slotID int, radioType RadioType, enabled bool) (APSlotConfigRPCInput, error) {
	band, ok := radioBandNumber(radioType)
	if !ok {
		return APSlotConfigRPCInput{}, ierrors.ValidationError("radio type", string(radioType))
	}

	return APSlotConfigRPCInput{
		Mode:   core.GetAdminStateMode(enabled),
		SlotID: slotID,
		Band:   strconv.Itoa(band),
	}, nil
}

// updateRadioState posts one slot-admin RPC. The arm is the caller's: the input's choice is
// mandatory, and an input carrying both names two access points.
func (s Service) updateRadioState(ctx context.Context, input APSlotConfigRPCInput) error {
	payload := APSlotConfigRPCPayload{Input: input}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.APSetApSlotAdminStateRPC, payload); err != nil {
		return ierrors.ServiceOperationError("set", "AP radio", "state", err)
	}

	return nil
}

// assignTags assigns multiple tags to an Access Point (internal implementation).
func (s Service) assignTags(ctx context.Context, apMAC string, tags ApTag) error {
	if !validation.IsValidMACAddr(apMAC) {
		return ierrors.ValidationError("AP MAC address", apMAC)
	}
	if !validation.HasValidTags(tags.SiteTag, tags.PolicyTag, tags.RFTag) {
		return ierrors.RequiredParameterError("at least one tag")
	}

	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return fmt.Errorf("invalid AP MAC address %s: %w", apMAC, err)
	}
	tagData, err := s.resolveAPTagData(ctx, normalizedMAC, tags)
	if err != nil {
		return err
	}
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APTagQueryPath, normalizedMAC)

	// Execute operation with direct error propagation
	if err := core.PutVoid(ctx, s.Client(), url, APTagPayload{ApTag: tagData}); err != nil {
		return ierrors.ServiceOperationError("assign", "AP", "tags", err)
	}
	return nil
}

// resolveAPTagData returns the entry to write. The request replaces the whole entry, so a tag
// the caller did not name is carried over from the AP's current one; the defaults apply only
// where the AP has no entry to carry over.
func (s Service) resolveAPTagData(
	ctx context.Context,
	normalizedMAC string,
	tags ApTag,
) (APCfgApTagData, error) {
	current, err := s.GetTagConfigByMAC(ctx, normalizedMAC)
	if err != nil && !core.IsNotFoundError(err) {
		return APCfgApTagData{}, ierrors.ServiceOperationError("assign", "AP", "tags", err)
	}
	if err != nil || current == nil || len(current.ApTag) == 0 {
		return buildAPCfgApTagData(normalizedMAC, tags), nil
	}

	// A tag holding its default is omitted from the read, so the carried-over value still goes
	// through the defaults: the controller rejects a payload naming a tag with an empty string.
	existing := current.ApTag[0]
	return buildAPCfgApTagData(normalizedMAC, ApTag{
		SiteTag:        validation.SelectNonEmptyValue(tags.SiteTag, existing.SiteTag),
		PolicyTag:      validation.SelectNonEmptyValue(tags.PolicyTag, existing.PolicyTag),
		RFTag:          validation.SelectNonEmptyValue(tags.RFTag, existing.RFTag),
		PrimingProfile: existing.PrimingProfile,
	}), nil
}

// buildAPCfgApTagData constructs the payload for tag assignment requests.
func buildAPCfgApTagData(normalizedMAC string, tags ApTag) APCfgApTagData {
	return APCfgApTagData{
		APMac:          normalizedMAC,
		SiteTag:        validation.SelectNonEmptyValue(tags.SiteTag, validation.DefaultSiteTag),
		PolicyTag:      validation.SelectNonEmptyValue(tags.PolicyTag, validation.DefaultPolicyTag),
		RFTag:          validation.SelectNonEmptyValue(tags.RFTag, validation.DefaultRFTag),
		PrimingProfile: tags.PrimingProfile,
	}
}
