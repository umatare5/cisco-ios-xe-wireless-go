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
func (s Service) GetConfig(ctx context.Context) (*ApCfg, error) {
	return core.Get[ApCfg](ctx, s.Client(), routes.APCfgPath)
}

// ListTagConfigs retrieves access point tag configurations.
func (s Service) ListTagConfigs(ctx context.Context) (*ApCfgApTag, error) {
	return core.Get[ApCfgApTag](ctx, s.Client(), routes.APTagsPath)
}

// GetTagConfigByMAC retrieves AP tag configuration filtered by AP MAC address.
func (s Service) GetTagConfigByMAC(ctx context.Context, mac string) (*ApCfgApTag, error) {
	if err := validation.ValidateMACAddress(mac); err != nil {
		return nil, fmt.Errorf(ErrInvalidAPMacFormat, mac)
	}
	normalizedMAC, err := validation.NormalizeMACAddress(mac)
	if err != nil {
		return nil, fmt.Errorf(ErrInvalidAPMacFormat, mac)
	}

	// Build correct RESTCONF path: /ap-cfg-data/ap-tags/ap-tag=MAC
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APTagQueryPath, normalizedMAC)
	return core.Get[ApCfgApTag](ctx, s.Client(), url)
}

// ListTagSourcePriorityConfigs retrieves tag source priority configurations.
func (s Service) ListTagSourcePriorityConfigs(ctx context.Context) (*TagSourcePriorityConfigs, error) {
	return core.Get[TagSourcePriorityConfigs](ctx, s.Client(), routes.APTagSourcePriorityConfigsPath)
}

// GetTagSourcePriorityConfigByPriority retrieves tag source priority configuration filtered by priority.
func (s Service) GetTagSourcePriorityConfigByPriority(
	ctx context.Context,
	priority int,
) (*ApCfgTagSourcePriorityConfigs, error) {
	url := s.Client().RESTCONFBuilder().BuildQueryURL(
		routes.APTagSourcePriorityConfigQueryPath,
		strconv.Itoa(priority),
	)
	return core.Get[ApCfgTagSourcePriorityConfigs](ctx, s.Client(), url)
}

// GetGlobalInfo retrieves the complete AP global operational data.
func (s Service) GetGlobalInfo(ctx context.Context) (*ApGlobalOper, error) {
	return core.Get[ApGlobalOper](ctx, s.Client(), routes.APGlobalOperPath)
}

// GetEWLCAPStats retrieves EWLC AP statistics.
func (s Service) GetEWLCAPStats(ctx context.Context) (*ApGlobalOperEwlcApStats, error) {
	return core.Get[ApGlobalOperEwlcApStats](ctx, s.Client(), routes.APEwlcApStatsPath)
}

// ListAPHistoryByEthernetMAC retrieves AP history data filtered by ethernet MAC address.
func (s Service) ListAPHistoryByEthernetMAC(
	ctx context.Context,
	ethernetMAC string,
) (*ApGlobalOperApHistory, error) {
	if ethernetMAC == "" {
		return nil, core.ErrResourceNotFound
	}
	if strings.TrimSpace(ethernetMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APHistoryQueryPath, ethernetMAC)
	return core.Get[ApGlobalOperApHistory](ctx, s.Client(), url)
}

// GetAPJoinStatsByWTPMAC retrieves AP join statistics filtered by WTP MAC address.
func (s Service) GetAPJoinStatsByWTPMAC(
	ctx context.Context, mac string,
) (*ApGlobalOperApJoinStats, error) {
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
	return core.Get[ApGlobalOperApJoinStats](ctx, s.Client(), url)
}

// GetWLANClientStatsByWLANID retrieves WLAN client statistics filtered by WLAN ID.
func (s Service) GetWLANClientStatsByWLANID(
	ctx context.Context,
	wlanID int,
) (*ApGlobalOperWlanClientStats, error) {
	if wlanID <= 0 {
		return nil, core.ErrResourceNotFound
	}
	url := s.Client().RESTCONFBuilder().BuildQueryURL(
		routes.APWlanClientStatsQueryPath,
		strconv.Itoa(wlanID),
	)
	return core.Get[ApGlobalOperWlanClientStats](ctx, s.Client(), url)
}

// ListAPHistory retrieves only AP history data using fields parameter.
func (s Service) ListAPHistory(ctx context.Context) (*ApGlobalOperApHistory, error) {
	return core.Get[ApGlobalOperApHistory](ctx, s.Client(), routes.APHistoryPath)
}

// ListAPJoinStats retrieves only AP join statistics using fields parameter.
func (s Service) ListAPJoinStats(ctx context.Context) (*ApGlobalOperApJoinStats, error) {
	return core.Get[ApGlobalOperApJoinStats](ctx, s.Client(), routes.APJoinStatsPath)
}

// ListWLANClientStats retrieves only WLAN client statistics using fields parameter.
func (s Service) ListWLANClientStats(
	ctx context.Context,
) (*ApGlobalOperWlanClientStats, error) {
	return core.Get[ApGlobalOperWlanClientStats](ctx, s.Client(), routes.APWlanClientStatsPath)
}

// GetOperational retrieves the complete AP operational data.
func (s Service) GetOperational(ctx context.Context) (*ApOper, error) {
	return core.Get[ApOper](ctx, s.Client(), routes.APOperPath)
}

// ListApOperData retrieves AP operational data.
func (s Service) ListApOperData(ctx context.Context) (*ApOperData, error) {
	return core.Get[ApOperData](ctx, s.Client(), routes.APOperDataPath)
}

// ListCAPWAPData retrieves CAPWAP protocol data.
func (s Service) ListCAPWAPData(ctx context.Context) (*ApOperCAPWAPData, error) {
	return core.Get[ApOperCAPWAPData](ctx, s.Client(), routes.APCapwapDataPath)
}

// GetCAPWAPDataByWTPMAC retrieves CAPWAP data for a specific WTP MAC.
func (s Service) GetCAPWAPDataByWTPMAC(ctx context.Context, wtpMAC string) (*ApOperCAPWAPData, error) {
	if wtpMAC == "" || strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APCapwapDataPath, wtpMAC)
	return core.Get[ApOperCAPWAPData](ctx, s.Client(), url)
}

// ListNameMACMaps retrieves AP name-to-MAC mapping data.
func (s Service) ListNameMACMaps(ctx context.Context) (*ApOperApNameMACMap, error) {
	return core.Get[ApOperApNameMACMap](ctx, s.Client(), routes.APApNameMACMapPath)
}

// GetNameMACMapByWTPName retrieves AP name-to-MAC mapping filtered by WTP name.
func (s Service) GetNameMACMapByWTPName(ctx context.Context, wtpName string) (*ApOperApNameMACMap, error) {
	if wtpName == "" {
		return nil, core.ErrResourceNotFound
	}
	if strings.TrimSpace(wtpName) == "" {
		return nil, core.ErrResourceNotFound
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APApNameMACMapPath, wtpName)
	return core.Get[ApOperApNameMACMap](ctx, s.Client(), url)
}

// ListRadioData retrieves radio operational data.
func (s Service) ListRadioData(ctx context.Context) (*ApOperRadioOperData, error) {
	return core.Get[ApOperRadioOperData](ctx, s.Client(), routes.APRadioOperDataPath)
}

// GetRadioStatusByWTPMACAndSlot retrieves radio operational data by WTP MAC and slot ID.
func (s Service) GetRadioStatusByWTPMACAndSlot(
	ctx context.Context, wtpMAC string, slotID int,
) (*ApOperRadioOperData, error) {
	if wtpMAC == "" {
		return nil, core.ErrResourceNotFound
	}
	if strings.TrimSpace(wtpMAC) == "" {
		return nil, core.ErrResourceNotFound
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(routes.APRadioOperDataPath, wtpMAC, slotID)
	return core.Get[ApOperRadioOperData](ctx, s.Client(), url)
}

// ListRadioNeighbors retrieves all AP radio neighbor information.
func (s Service) ListRadioNeighbors(ctx context.Context) (*ApOperApRadioNeighbor, error) {
	return core.Get[ApOperApRadioNeighbor](ctx, s.Client(), routes.APRadioNeighborPath)
}

// GetRadioNeighborByAPMACSlotAndBSSID retrieves AP radio neighbor information for a specific AP MAC, slot ID and BSSID.
// This follows the YANG model key structure: "ap-mac slot-id bssid".
func (s Service) GetRadioNeighborByAPMACSlotAndBSSID(
	ctx context.Context, apMAC string, slotID int, bssid string,
) (*ApOperApRadioNeighbor, error) {
	if apMAC == "" || strings.TrimSpace(apMAC) == "" {
		return nil, errors.New("AP MAC address cannot be empty")
	}
	if bssid == "" || strings.TrimSpace(bssid) == "" {
		return nil, errors.New("BSSID cannot be empty")
	}

	if err := validation.ValidateMACAddress(apMAC); err != nil {
		return nil, fmt.Errorf("invalid AP MAC address: %w", err)
	}
	if err := validation.ValidateMACAddress(bssid); err != nil {
		return nil, fmt.Errorf("invalid BSSID: %w", err)
	}

	normalizedAPMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid AP MAC address %s: %w", apMAC, err)
	}
	normalizedBSSID, err := validation.NormalizeMACAddress(bssid)
	if err != nil {
		return nil, fmt.Errorf("invalid BSSID %s: %w", bssid, err)
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.APRadioNeighborPath,
		normalizedAPMAC,
		slotID,
		normalizedBSSID,
	)
	return core.Get[ApOperApRadioNeighbor](ctx, s.Client(), url)
}

// ListActiveImageLocations retrieves active image location information using fields parameter.
func (s Service) ListActiveImageLocations(
	ctx context.Context,
) (*ApOperApImageActiveLocation, error) {
	return core.Get[ApOperApImageActiveLocation](ctx, s.Client(), routes.APImageActiveLocationPath)
}

// ListPreparedImageLocations retrieves only AP image prepare location data using fields parameter.
func (s Service) ListPreparedImageLocations(
	ctx context.Context,
) (*ApOperApImagePrepareLocation, error) {
	return core.Get[ApOperApImagePrepareLocation](ctx, s.Client(), routes.APImagePrepareLocationPath)
}

// ListPowerInfo retrieves only AP power information using fields parameter.
func (s Service) ListPowerInfo(ctx context.Context) (*ApOperApPwrInfo, error) {
	return core.Get[ApOperApPwrInfo](ctx, s.Client(), routes.APPwrInfoPath)
}

// ListSensorStatus retrieves only AP sensor status using fields parameter.
func (s Service) ListSensorStatus(ctx context.Context) (*ApOperApSensorStatus, error) {
	return core.Get[ApOperApSensorStatus](ctx, s.Client(), routes.APSensorStatusPath)
}

// ListCAPWAPPackets retrieves only CAPWAP packets data using fields parameter.
func (s Service) ListCAPWAPPackets(ctx context.Context) (*ApOperCAPWAPPkts, error) {
	return core.Get[ApOperCAPWAPPkts](ctx, s.Client(), routes.APCapwapPktsPath)
}

// ListIotFirmware retrieves IoT firmware information for all access points.
func (s Service) ListIotFirmware(ctx context.Context) (*ApOperIotFirmware, error) {
	return core.Get[ApOperIotFirmware](ctx, s.Client(), routes.APIotFirmwarePath)
}

// ListRadioResetStats retrieves radio reset statistics for all access points.
func (s Service) ListRadioResetStats(ctx context.Context) (*ApOperRadioResetStats, error) {
	return core.Get[ApOperRadioResetStats](ctx, s.Client(), routes.APRadioResetStatsPath)
}

// GetRadioResetStatsByAPMACAndRadioID retrieves radio reset statistics for a specific AP MAC and radio ID.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetRadioResetStatsByAPMACAndRadioID(
	ctx context.Context, apMAC string, radioID int,
) (*ApOperRadioResetStats, error) {
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
	return core.Get[ApOperRadioResetStats](ctx, s.Client(), url)
}

// ListQosClientData retrieves QoS client data for all access points.
func (s Service) ListQosClientData(ctx context.Context) (*ApOperQosClientData, error) {
	return core.Get[ApOperQosClientData](ctx, s.Client(), routes.APQosClientDataPath)
}

// GetQosClientDataByClientMAC retrieves QoS client data for a specific client MAC address.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetQosClientDataByClientMAC(
	ctx context.Context, clientMAC string,
) (*ApOperQosClientData, error) {
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
	return core.Get[ApOperQosClientData](ctx, s.Client(), url)
}

// ListWtpSlotWlanStats retrieves WTP slot WLAN statistics for all access points.
func (s Service) ListWtpSlotWlanStats(ctx context.Context) (*ApOperWtpSlotWlanStats, error) {
	return core.Get[ApOperWtpSlotWlanStats](ctx, s.Client(), routes.APWtpSlotWlanStatsPath)
}

// GetWtpSlotWlanStatsByWTPMACSlotAndWLANID retrieves WTP slot WLAN statistics for a specific WTP MAC, slot ID, and WLAN ID.
func (s Service) GetWtpSlotWlanStatsByWTPMACSlotAndWLANID(
	ctx context.Context,
	wtpMAC string,
	slotID int,
	wlanID int,
) (*ApOperWtpSlotWlanStats, error) {
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
	return core.Get[ApOperWtpSlotWlanStats](ctx, s.Client(), url)
}

// ListEthernetMACWtpMACMaps retrieves Ethernet MAC to WTP MAC mapping for all access points.
func (s Service) ListEthernetMACWtpMACMaps(ctx context.Context) (*ApOperEthernetMACWtpMACMap, error) {
	return core.Get[ApOperEthernetMACWtpMACMap](ctx, s.Client(), routes.APEthernetMACWtpMACMapPath)
}

// GetEthernetMACWtpMACMapByEthernetMAC retrieves Ethernet MAC to WTP MAC mapping for a specific Ethernet MAC address.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetEthernetMACWtpMACMapByEthernetMAC(
	ctx context.Context, ethernetMAC string,
) (*ApOperEthernetMACWtpMACMap, error) {
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
	return core.Get[ApOperEthernetMACWtpMACMap](ctx, s.Client(), url)
}

// ListRadioOperStats retrieves radio operational statistics for all access points.
func (s Service) ListRadioOperStats(ctx context.Context) (*ApOperRadioOperStats, error) {
	return core.Get[ApOperRadioOperStats](ctx, s.Client(), routes.APRadioOperStatsPath)
}

// GetRadioOperStatsByWTPMACAndSlot retrieves radio operational statistics for a specific WTP MAC and slot ID.
func (s Service) GetRadioOperStatsByWTPMACAndSlot(
	ctx context.Context,
	wtpMAC string,
	slotID int,
) (*ApOperRadioOperStats, error) {
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
	return core.Get[ApOperRadioOperStats](ctx, s.Client(), url)
}

// ListEthernetIfStats retrieves Ethernet interface statistics for all access points.
func (s Service) ListEthernetIfStats(ctx context.Context) (*ApOperEthernetIfStats, error) {
	return core.Get[ApOperEthernetIfStats](ctx, s.Client(), routes.APEthernetIfStatsPath)
}

// GetEthernetIfStatsByWTPMACAndInterfaceID retrieves Ethernet interface statistics for a specific WTP MAC and interface ID.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetEthernetIfStatsByWTPMACAndInterfaceID(
	ctx context.Context,
	wtpMAC string,
	interfaceID string,
) (*ApOperEthernetIfStats, error) {
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
	return core.Get[ApOperEthernetIfStats](ctx, s.Client(), url)
}

// ListEwlcWncdStats retrieves EWLC WNCD statistics information.
func (s Service) ListEwlcWncdStats(ctx context.Context) (*ApOperEwlcWncdStats, error) {
	return core.Get[ApOperEwlcWncdStats](ctx, s.Client(), routes.APEwlcWncdStatsPath)
}

// ListApIoxOperData retrieves AP IOx operational data for all access points.
func (s Service) ListApIoxOperData(ctx context.Context) (*ApOperApIoxOperData, error) {
	return core.Get[ApOperApIoxOperData](ctx, s.Client(), routes.APApIoxOperDataPath)
}

// GetApIoxOperDataByWTPMAC retrieves AP IOx operational data for a specific WTP MAC address.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetApIoxOperDataByWTPMAC(
	ctx context.Context, wtpMAC string,
) (*ApOperApIoxOperData, error) {
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
	return core.Get[ApOperApIoxOperData](ctx, s.Client(), url)
}

// ListQosGlobalStats retrieves QoS global statistics information.
func (s Service) ListQosGlobalStats(ctx context.Context) (*ApOperQosGlobalStats, error) {
	return core.Get[ApOperQosGlobalStats](ctx, s.Client(), routes.APQosGlobalStatsPath)
}

// ListRlanOper retrieves RLAN operational data for all access points.
func (s Service) ListRlanOper(ctx context.Context) (*ApOperRlanOper, error) {
	return core.Get[ApOperRlanOper](ctx, s.Client(), routes.APRlanOperPath)
}

// GetRlanOperByWTPMACAndPortID retrieves RLAN operational data for a specific WTP MAC and port ID.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetRlanOperByWTPMACAndPortID(
	ctx context.Context, wtpMAC string, portID int,
) (*ApOperRlanOper, error) {
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
	return core.Get[ApOperRlanOper](ctx, s.Client(), url)
}

// ListEwlcMewlcPredownloadRec retrieves EWLC MEWLC predownload record information.
func (s Service) ListEwlcMewlcPredownloadRec(ctx context.Context) (*ApOperEwlcMewlcPredownloadRec, error) {
	return core.Get[ApOperEwlcMewlcPredownloadRec](ctx, s.Client(), routes.APEwlcMewlcPredownloadRecPath)
}

// ListCdpCacheData retrieves CDP cache data for all access points.
func (s Service) ListCdpCacheData(ctx context.Context) (*ApOperCdpCacheData, error) {
	return core.Get[ApOperCdpCacheData](ctx, s.Client(), routes.APCdpCacheDataPath)
}

// GetCdpCacheDataByWTPMAC retrieves CDP cache data for a specific WTP MAC address.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetCdpCacheDataByWTPMAC(
	ctx context.Context, wtpMAC string,
) (*ApOperCdpCacheData, error) {
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
	return core.Get[ApOperCdpCacheData](ctx, s.Client(), url)
}

// ListLldpNeigh retrieves LLDP neighbor information for all access points.
func (s Service) ListLldpNeigh(ctx context.Context) (*ApOperLldpNeigh, error) {
	return core.Get[ApOperLldpNeigh](ctx, s.Client(), routes.APLldpNeighPath)
}

// GetLldpNeighByWTPMAC retrieves LLDP neighbor information for a specific WTP MAC address.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetLldpNeighByWTPMAC(
	ctx context.Context, wtpMAC string,
) (*ApOperLldpNeigh, error) {
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
	return core.Get[ApOperLldpNeigh](ctx, s.Client(), url)
}

// ListTpCertInfo retrieves trustpoint certificate info information.
func (s Service) ListTpCertInfo(ctx context.Context) (*ApOperTpCertInfo, error) {
	return core.Get[ApOperTpCertInfo](ctx, s.Client(), routes.APTpCertInfoPath)
}

// ListDiscData retrieves discovery data for all access points.
func (s Service) ListDiscData(ctx context.Context) (*ApOperDiscData, error) {
	return core.Get[ApOperDiscData](ctx, s.Client(), routes.APDiscDataPath)
}

// GetDiscDataByWTPMAC retrieves discovery data for a specific WTP MAC address.
func (s Service) GetDiscDataByWTPMAC(
	ctx context.Context, wtpMAC string,
) (*ApOperDiscData, error) {
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
	return core.Get[ApOperDiscData](ctx, s.Client(), url)
}

// ListCountryOper retrieves country operational data for all access points.
func (s Service) ListCountryOper(ctx context.Context) (*ApOperCountryOper, error) {
	return core.Get[ApOperCountryOper](ctx, s.Client(), routes.APCountryOperPath)
}

// GetCountryOperByWTPMACAndRadioID retrieves country operational data for a specific WTP MAC and radio ID.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetCountryOperByWTPMACAndRadioID(
	ctx context.Context, wtpMAC string, radioID int,
) (*ApOperCountryOper, error) {
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
	return core.Get[ApOperCountryOper](ctx, s.Client(), url)
}

// ListSuppCountryOper retrieves supported country operational data for all access points.
func (s Service) ListSuppCountryOper(ctx context.Context) (*ApOperSuppCountryOper, error) {
	return core.Get[ApOperSuppCountryOper](ctx, s.Client(), routes.APSuppCountryOperPath)
}

// GetSuppCountryOperByWTPMACAndRadioID retrieves supported country operational data for a specific WTP MAC and radio ID.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) GetSuppCountryOperByWTPMACAndRadioID(
	ctx context.Context, wtpMAC string, radioID int,
) (*ApOperSuppCountryOper, error) {
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
	return core.Get[ApOperSuppCountryOper](ctx, s.Client(), url)
}

// ListApNhGlobalData retrieves AP neighborhood global data.
func (s Service) ListApNhGlobalData(ctx context.Context) (*ApOperApNhGlobalData, error) {
	return core.Get[ApOperApNhGlobalData](ctx, s.Client(), routes.APApNhGlobalDataPath)
}

// DisableAP disables the administrative state of an access point.
func (s Service) DisableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-disabled")
}

// EnableAP enables the administrative state of an access point.
func (s Service) EnableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-enabled")
}

// EnableRadio enables a radio on an Access Point using MAC address.
func (s Service) EnableRadio(ctx context.Context, apMAC string, radioBand core.RadioBand) error {
	if err := validation.ValidateMACAddress(apMAC); err != nil {
		return err
	}
	return s.updateRadioState(ctx, apMAC, &radioBand, true)
}

// DisableRadio disables a radio on an Access Point using MAC address.
func (s Service) DisableRadio(ctx context.Context, apMAC string, radioBand core.RadioBand) error {
	if err := validation.ValidateMACAddress(apMAC); err != nil {
		return err
	}
	return s.updateRadioState(ctx, apMAC, &radioBand, false)
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

// Reload restarts an Access Point by MAC address causing temporary service interruption.
func (s Service) Reload(ctx context.Context, apMAC string) error {
	if !validation.IsValidMACAddr(apMAC) {
		return fmt.Errorf(ErrInvalidAPMacFormat, apMAC)
	}

	resp, err := s.ListCAPWAPData(ctx)
	if err != nil {
		return fmt.Errorf(ErrFailedGetCAPWAPData, err)
	}
	if resp == nil {
		return errors.New(ErrCAPWAPDataUnavailable)
	}

	apName, found := findAPByMAC(resp, apMAC)
	if !found {
		return fmt.Errorf(ErrAPNotFoundByMAC, apMAC)
	}

	return s.reload(ctx, apName)
}

// updateAPState handles AP admin state changes with mac and mode parameters.
func (s Service) updateAPState(ctx context.Context, mac, mode string) error {
	if err := validation.ValidateMACAddress(mac); err != nil {
		return fmt.Errorf("invalid AP MAC address: %s", mac)
	}

	normalizedMAC, err := validation.NormalizeMACAddress(mac)
	if err != nil {
		return fmt.Errorf("invalid AP MAC address: %s", mac)
	}

	payload := APConfigRPCPayload{
		Input: APConfigRPCInput{
			Mode:    mode,
			MACAddr: normalizedMAC,
		},
	}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.APSetApAdminStateRPC, payload); err != nil {
		return ierrors.ServiceOperationError("update", "AP", "admin state", err)
	}

	return nil
}

// updateRadioState handles radio-level state changes.
func (s Service) updateRadioState(ctx context.Context, apMAC string, radioBand *core.RadioBand, enabled bool) error {
	if radioBand == nil {
		return ierrors.RequiredParameterError("radio band")
	}

	normalizedMAC, err := validation.NormalizeMACAddress(apMAC)
	if err != nil {
		return fmt.Errorf("invalid AP MAC address %s: %w", apMAC, err)
	}

	radioBandInfo, err := core.GetRadioBandInfo(int(*radioBand))
	if err != nil {
		return err
	}

	payload := APSlotConfigRPCPayload{
		Input: APSlotConfigRPCInput{
			Mode:    core.GetAdminStateMode(enabled),
			SlotID:  int(radioBandInfo.SlotID),
			Band:    strconv.Itoa(int(radioBandInfo.Band)),
			MACAddr: normalizedMAC,
		},
	}

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
	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.APTagPath, normalizedMAC)
	tagData := buildAPCfgApTagData(normalizedMAC, tags)

	// Execute operation with direct error propagation
	if err := core.PutVoid(ctx, s.Client(), url, APTagPayload{ApTag: tagData}); err != nil {
		return ierrors.ServiceOperationError("assign", "AP", "tags", err)
	}
	return nil
}

// reload is the internal helper function for AP reload operations.
func (s Service) reload(ctx context.Context, apName string) error {
	requestBody := APReloadRPCPayload{
		Input: APReloadRPCInput{
			APName: apName,
		},
	}
	return core.PostRPCVoid(ctx, s.Client(), routes.APApResetRPC, requestBody)
}

// findAPByMAC searches for an AP with the given MAC address in CAPWAP data.
func findAPByMAC(capwapData *ApOperCAPWAPData, apMAC string) (string, bool) {
	if capwapData == nil {
		return "", false
	}
	for _, data := range capwapData.CAPWAPData {
		if data.WtpMAC == apMAC {
			return data.Name, true
		}
	}
	return "", false
}

// buildAPCfgApTagData constructs the payload for tag assignment requests.
func buildAPCfgApTagData(normalizedMAC string, tags ApTag) APCfgApTagData {
	return APCfgApTagData{
		APMac:     normalizedMAC,
		SiteTag:   validation.SelectNonEmptyValue(tags.SiteTag, validation.DefaultSiteTag),
		PolicyTag: validation.SelectNonEmptyValue(tags.PolicyTag, validation.DefaultPolicyTag),
		RFTag:     validation.SelectNonEmptyValue(tags.RFTag, validation.DefaultRFTag),
	}
}
