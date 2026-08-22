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
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRRMCfg, error) {
	return core.Get[CiscoIOSXEWirelessRRMCfg](ctx, s.Client(), routes.RRMCfgPath, opts...)
}

// GetOperational retrieves RRM operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRRMOper, error) {
	return core.Get[CiscoIOSXEWirelessRRMOper](ctx, s.Client(), routes.RRMOperPath, opts...)
}

// GetGlobalOperational retrieves RRM global operational information.
func (s Service) GetGlobalOperational(ctx context.Context, opts ...core.GetOption) (*RRMGlobalOper, error) {
	return core.Get[RRMGlobalOper](ctx, s.Client(), routes.RRMGlobalOperPath, opts...)
}

// GetEmulationOperational retrieves RRM emulation operational information.
func (s Service) GetEmulationOperational(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMEmulOper, error) {
	return core.Get[CiscoIOSXEWirelessRRMEmulOper](ctx, s.Client(), routes.RRMEmulOperPath, opts...)
}

// ListRrms retrieves RRM configuration data.
func (s Service) ListRrms(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRRMCfgRrms, error) {
	return core.Get[CiscoIOSXEWirelessRRMCfgRrms](ctx, s.Client(), routes.RRMCfgRrmsPath, opts...)
}

// ListRRMMgrCfgEntries retrieves RRM manager configuration entries.
func (s Service) ListRRMMgrCfgEntries(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMCfgRRMMgrCfgEntries, error) {
	return core.Get[CiscoIOSXEWirelessRRMCfgRRMMgrCfgEntries](
		ctx,
		s.Client(),
		routes.RRMCfgRRMMgrCfgEntriesPath,
		opts...,
	)
}

// ListApAutoRFDot11Data retrieves AP auto RF 802.11 data.
func (s Service) ListApAutoRFDot11Data(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperApAutoRFDot11Data, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperApAutoRFDot11Data](
		ctx,
		s.Client(),
		routes.RRMOperApAutoRFDot11DataPath,
		opts...,
	)
}

// ListApDot11RadarData retrieves AP radar detection data.
func (s Service) ListApDot11RadarData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperApDot11RadarData, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperApDot11RadarData](
		ctx,
		s.Client(),
		routes.RRMOperApDot11RadarDataPath,
		opts...,
	)
}

// ListApDot11SpectrumData retrieves AP spectrum analysis data.
func (s Service) ListApDot11SpectrumData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperApDot11SpectrumData, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperApDot11SpectrumData](
		ctx,
		s.Client(),
		routes.RRMOperApDot11SpectrumDataPath, opts...,
	)
}

// ListRRMMeasurement retrieves RRM measurement data.
func (s Service) ListRRMMeasurement(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperRRMMeasurement, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperRRMMeasurement](ctx, s.Client(), routes.RRMOperRRMMeasurementPath, opts...)
}

// ListRadioSlot retrieves radio slot operational data.
func (s Service) ListRadioSlot(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperRadioSlot, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperRadioSlot](ctx, s.Client(), routes.RRMOperRadioSlotPath, opts...)
}

// ListMainData retrieves main RRM data by PHY type.
func (s Service) ListMainData(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRRMOperMainData, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperMainData](ctx, s.Client(), routes.RRMOperMainDataPath, opts...)
}

// ListRegDomainOper retrieves regulatory domain operational data.
func (s Service) ListRegDomainOper(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperRegDomainOper, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperRegDomainOper](ctx, s.Client(), routes.RRMOperRegDomainOperPath, opts...)
}

// ListSpectrumDeviceTable retrieves spectrum device detection table.
func (s Service) ListSpectrumDeviceTable(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperSpectrumDeviceTable, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperSpectrumDeviceTable](
		ctx,
		s.Client(),
		routes.RRMOperSpectrumDeviceTablePath, opts...,
	)
}

// ListSpectrumAqTable retrieves spectrum air quality table.
func (s Service) ListSpectrumAqTable(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMOperSpectrumAqTable, error) {
	return core.Get[CiscoIOSXEWirelessRRMOperSpectrumAqTable](
		ctx,
		s.Client(),
		routes.RRMOperSpectrumAqTablePath,
		opts...,
	)
}

// ListRRMOneShotCounters retrieves RRM one-shot counters.
func (s Service) ListRRMOneShotCounters(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRRMOneShotCounters, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRRMOneShotCounters](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRRMOneShotCountersPath, opts...,
	)
}

// ListRRMChannelParams retrieves RRM channel parameters.
func (s Service) ListRRMChannelParams(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRRMChannelParams, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRRMChannelParams](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRRMChannelParamsPath, opts...,
	)
}

// ListRadioOperData24g retrieves 2.4GHz radio operational data.
func (s Service) ListRadioOperData24g(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRadioOperData24g, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRadioOperData24g](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRadioOperData24gPath, opts...,
	)
}

// ListRadioOperData5g retrieves 5GHz radio operational data.
func (s Service) ListRadioOperData5g(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRadioOperData5g, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRadioOperData5g](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRadioOperData5gPath, opts...,
	)
}

// ListRadioOperData6ghz retrieves 6GHz radio operational data.
func (s Service) ListRadioOperData6ghz(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRadioOperData6ghz, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRadioOperData6ghz](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRadioOperData6ghzPath, opts...,
	)
}

// ListRadioOperDataDualband retrieves dual-band radio operational data.
func (s Service) ListRadioOperDataDualband(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRadioOperDataDualband, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRadioOperDataDualband](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRadioOperDataDualbandPath, opts...,
	)
}

// ListSpectrumBandConfigData retrieves spectrum band configuration data.
func (s Service) ListSpectrumBandConfigData(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperSpectrumBandConfigData, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperSpectrumBandConfigData](
		ctx,
		s.Client(),
		routes.RRMGlobalOperSpectrumBandConfigDataPath, opts...,
	)
}

// ListRRMClientData retrieves RRM client data.
func (s Service) ListRRMClientData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRRMClientData, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRRMClientData](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRRMClientDataPath, opts...,
	)
}

// ListRRMFraStats retrieves RRM flexible radio assignment statistics from global operational data.
func (s Service) ListRRMFraStats(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRRMFraStats, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRRMFraStats](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRRMFraStatsPath,
		opts...,
	)
}

// ListRRMCoverage retrieves RRM coverage information.
func (s Service) ListRRMCoverage(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperRRMCoverage, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperRRMCoverage](
		ctx,
		s.Client(),
		routes.RRMGlobalOperRRMCoveragePath,
		opts...,
	)
}

// ListSpectrumAqWorstTable retrieves spectrum air quality worst table.
func (s Service) ListSpectrumAqWorstTable(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMGlobalOperSpectrumAqWorstTable, error) {
	return core.Get[CiscoIOSXEWirelessRRMGlobalOperSpectrumAqWorstTable](
		ctx,
		s.Client(),
		routes.RRMGlobalOperSpectrumAqWorstTablePath, opts...,
	)
}

// ListRRMFraStatsFromEmul retrieves RRM flexible radio assignment statistics from emulation operational data.
func (s Service) ListRRMFraStatsFromEmul(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRRMEmulOperRRMFraStats, error) {
	return core.Get[CiscoIOSXEWirelessRRMEmulOperRRMFraStats](
		ctx,
		s.Client(),
		routes.RRMEmulOperRRMFraStatsPath,
		opts...,
	)
}
