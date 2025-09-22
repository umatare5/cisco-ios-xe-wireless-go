package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides RRM (Radio Resource Management) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new RRM service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves complete RRM configuration data.
func (s Service) GetConfig(ctx context.Context) (*RRMCfg, error) {
	return core.Get[RRMCfg](ctx, s.Client(), routes.RRMCfgPath)
}

// GetOperational retrieves RRM operational data.
func (s Service) GetOperational(ctx context.Context) (*RRMOper, error) {
	return core.Get[RRMOper](ctx, s.Client(), routes.RRMOperPath)
}

// GetGlobalOperational retrieves RRM global operational information.
func (s Service) GetGlobalOperational(ctx context.Context) (*RRMGlobalOper, error) {
	return core.Get[RRMGlobalOper](ctx, s.Client(), routes.RRMGlobalOperPath)
}

// GetEmulationOperational retrieves RRM emulation operational information.
func (s Service) GetEmulationOperational(ctx context.Context) (*RRMEmulOper, error) {
	return core.Get[RRMEmulOper](ctx, s.Client(), routes.RRMEmulOperPath)
}

// ListRrms retrieves RRM configuration data.
func (s Service) ListRrms(ctx context.Context) (*RRMCfgRrms, error) {
	return core.Get[RRMCfgRrms](ctx, s.Client(), routes.RRMCfgRrmsPath)
}

// ListRRMMgrCfgEntries retrieves RRM manager configuration entries.
func (s Service) ListRRMMgrCfgEntries(ctx context.Context) (*RRMCfgRRMMgrCfgEntries, error) {
	return core.Get[RRMCfgRRMMgrCfgEntries](ctx, s.Client(), routes.RRMCfgRRMMgrCfgEntriesPath)
}

// ListApAutoRFDot11Data retrieves AP auto RF 802.11 data.
func (s Service) ListApAutoRFDot11Data(ctx context.Context) (*RRMOperApAutoRFDot11Data, error) {
	return core.Get[RRMOperApAutoRFDot11Data](ctx, s.Client(), routes.RRMOperApAutoRFDot11DataPath)
}

// ListApDot11RadarData retrieves AP radar detection data.
func (s Service) ListApDot11RadarData(ctx context.Context) (*RRMOperApDot11RadarData, error) {
	return core.Get[RRMOperApDot11RadarData](ctx, s.Client(), routes.RRMOperApDot11RadarDataPath)
}

// ListApDot11SpectrumData retrieves AP spectrum analysis data.
func (s Service) ListApDot11SpectrumData(ctx context.Context) (*RRMOperApDot11SpectrumData, error) {
	return core.Get[RRMOperApDot11SpectrumData](ctx, s.Client(), routes.RRMOperApDot11SpectrumDataPath)
}

// ListRRMMeasurement retrieves RRM measurement data.
func (s Service) ListRRMMeasurement(ctx context.Context) (*RRMOperRRMMeasurement, error) {
	return core.Get[RRMOperRRMMeasurement](ctx, s.Client(), routes.RRMOperRRMMeasurementPath)
}

// ListRadioSlot retrieves radio slot operational data.
func (s Service) ListRadioSlot(ctx context.Context) (*RRMOperRadioSlot, error) {
	return core.Get[RRMOperRadioSlot](ctx, s.Client(), routes.RRMOperRadioSlotPath)
}

// ListMainData retrieves main RRM data by PHY type.
func (s Service) ListMainData(ctx context.Context) (*RRMOperMainData, error) {
	return core.Get[RRMOperMainData](ctx, s.Client(), routes.RRMOperMainDataPath)
}

// ListRegDomainOper retrieves regulatory domain operational data.
func (s Service) ListRegDomainOper(ctx context.Context) (*RRMOperRegDomainOper, error) {
	return core.Get[RRMOperRegDomainOper](ctx, s.Client(), routes.RRMOperRegDomainOperPath)
}

// ListSpectrumDeviceTable retrieves spectrum device detection table.
func (s Service) ListSpectrumDeviceTable(ctx context.Context) (*RRMOperSpectrumDeviceTable, error) {
	return core.Get[RRMOperSpectrumDeviceTable](ctx, s.Client(), routes.RRMOperSpectrumDeviceTablePath)
}

// ListSpectrumAqTable retrieves spectrum air quality table.
func (s Service) ListSpectrumAqTable(ctx context.Context) (*RRMOperSpectrumAqTable, error) {
	return core.Get[RRMOperSpectrumAqTable](ctx, s.Client(), routes.RRMOperSpectrumAqTablePath)
}

// ListRRMOneShotCounters retrieves RRM one-shot counters.
func (s Service) ListRRMOneShotCounters(ctx context.Context) (*RRMGlobalOperRRMOneShotCounters, error) {
	return core.Get[RRMGlobalOperRRMOneShotCounters](ctx, s.Client(), routes.RRMGlobalOperRRMOneShotCountersPath)
}

// ListRRMChannelParams retrieves RRM channel parameters.
func (s Service) ListRRMChannelParams(ctx context.Context) (*RRMGlobalOperRRMChannelParams, error) {
	return core.Get[RRMGlobalOperRRMChannelParams](ctx, s.Client(), routes.RRMGlobalOperRRMChannelParamsPath)
}

// ListRadioOperData24g retrieves 2.4GHz radio operational data.
func (s Service) ListRadioOperData24g(ctx context.Context) (*RRMGlobalOperRadioOperData24g, error) {
	return core.Get[RRMGlobalOperRadioOperData24g](ctx, s.Client(), routes.RRMGlobalOperRadioOperData24gPath)
}

// ListRadioOperData5g retrieves 5GHz radio operational data.
func (s Service) ListRadioOperData5g(ctx context.Context) (*RRMGlobalOperRadioOperData5g, error) {
	return core.Get[RRMGlobalOperRadioOperData5g](ctx, s.Client(), routes.RRMGlobalOperRadioOperData5gPath)
}

// ListRadioOperData6ghz retrieves 6GHz radio operational data.
func (s Service) ListRadioOperData6ghz(ctx context.Context) (*RRMGlobalOperRadioOperData6ghz, error) {
	return core.Get[RRMGlobalOperRadioOperData6ghz](ctx, s.Client(), routes.RRMGlobalOperRadioOperData6ghzPath)
}

// ListRadioOperDataDualband retrieves dual-band radio operational data.
func (s Service) ListRadioOperDataDualband(ctx context.Context) (*RRMGlobalOperRadioOperDataDualband, error) {
	return core.Get[RRMGlobalOperRadioOperDataDualband](ctx, s.Client(), routes.RRMGlobalOperRadioOperDataDualbandPath)
}

// ListSpectrumBandConfigData retrieves spectrum band configuration data.
func (s Service) ListSpectrumBandConfigData(
	ctx context.Context,
) (*RRMGlobalOperSpectrumBandConfigData, error) {
	return core.Get[RRMGlobalOperSpectrumBandConfigData](
		ctx,
		s.Client(),
		routes.RRMGlobalOperSpectrumBandConfigDataPath,
	)
}

// ListRRMClientData retrieves RRM client data.
func (s Service) ListRRMClientData(ctx context.Context) (*RRMGlobalOperRRMClientData, error) {
	return core.Get[RRMGlobalOperRRMClientData](ctx, s.Client(), routes.RRMGlobalOperRRMClientDataPath)
}

// ListRRMFraStats retrieves RRM flexible radio assignment statistics from global operational data.
func (s Service) ListRRMFraStats(ctx context.Context) (*RRMGlobalOperRRMFraStats, error) {
	return core.Get[RRMGlobalOperRRMFraStats](ctx, s.Client(), routes.RRMGlobalOperRRMFraStatsPath)
}

// ListRRMCoverage retrieves RRM coverage information.
func (s Service) ListRRMCoverage(ctx context.Context) (*RRMGlobalOperRRMCoverage, error) {
	return core.Get[RRMGlobalOperRRMCoverage](ctx, s.Client(), routes.RRMGlobalOperRRMCoveragePath)
}

// ListSpectrumAqWorstTable retrieves spectrum air quality worst table.
func (s Service) ListSpectrumAqWorstTable(ctx context.Context) (*RRMGlobalOperSpectrumAqWorstTable, error) {
	return core.Get[RRMGlobalOperSpectrumAqWorstTable](ctx, s.Client(), routes.RRMGlobalOperSpectrumAqWorstTablePath)
}

// ListRRMFraStatsFromEmul retrieves RRM flexible radio assignment statistics from emulation operational data.
func (s Service) ListRRMFraStatsFromEmul(ctx context.Context) (*RRMEmulOperRRMFraStats, error) {
	return core.Get[RRMEmulOperRRMFraStats](ctx, s.Client(), routes.RRMEmulOperRRMFraStatsPath)
}
