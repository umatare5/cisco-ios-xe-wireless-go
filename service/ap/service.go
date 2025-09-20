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

// ListRadioStatus retrieves radio operational status data.
func (s Service) ListRadioStatus(ctx context.Context) (*RadioOperData, error) {
	return core.Get[RadioOperData](ctx, s.Client(), routes.APRadioOperDataPath)
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

// EnableAP enables the administrative state of an access point.
func (s Service) EnableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-enabled")
}

// DisableAP disables the administrative state of an access point.
func (s Service) DisableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-disabled")
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
