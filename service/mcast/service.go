package mcast

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides multicast management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Multicast service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves multicast operational data from the controller.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessMcastOper, error) {
	return core.Get[CiscoIOSXEWirelessMcastOper](ctx, s.Client(), routes.McastOperPath, opts...)
}

// GetFlexConnectMediastreamClientSummary retrieves FlexConnect mediastream client summary data from the wireless controller.
func (s Service) GetFlexConnectMediastreamClientSummary(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperFlexMediastreamClientSummary, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperFlexMediastreamClientSummary](
		ctx, s.Client(), routes.McastFlexMediastreamPath, opts...,
	)
}

// ListVLANL2MGIDs retrieves VLAN Layer 2 multicast group ID operational data from the wireless controller.
func (s Service) ListVLANL2MGIDs(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperVlanL2MgidOp, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperVlanL2MgidOp](ctx, s.Client(), routes.McastVlanL2MgidPath, opts...)
}

// GetFabricMediastreamClientSummary retrieves fabric mediastream client summary data from the wireless controller.
func (s Service) GetFabricMediastreamClientSummary(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperFabricMediaStreamClientSummary, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperFabricMediaStreamClientSummary](
		ctx, s.Client(), routes.McastFabricMediastreamPath, opts...,
	)
}

// GetMcastMgidInfo retrieves multicast MGID information from the wireless controller.
func (s Service) GetMcastMgidInfo(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperMcastMgidInfo, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperMcastMgidInfo](ctx, s.Client(), routes.McastMgidInfoPath, opts...)
}

// GetMulticastOperData retrieves multicast operational data from the wireless controller.
func (s Service) GetMulticastOperData(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperMulticastOperData, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperMulticastOperData](
		ctx,
		s.Client(),
		routes.McastMulticastOperDataPath,
		opts...,
	)
}

// ListRrcHistoryClientRecordData retrieves RRC history client record data from the controller.
func (s Service) ListRrcHistoryClientRecordData(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperRrcHistoryClientRecordData, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperRrcHistoryClientRecordData](
		ctx,
		s.Client(),
		routes.McastRrcHistoryClientRecordDataPath, opts...,
	)
}

// ListRrcSrRadioRecord retrieves RRC stream radio record data from the controller.
func (s Service) ListRrcSrRadioRecord(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperRrcSrRadioRecord, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperRrcSrRadioRecord](
		ctx,
		s.Client(),
		routes.McastRrcSrRadioRecordPath,
		opts...,
	)
}

// ListRrcStreamRecord retrieves RRC stream record data from the controller.
func (s Service) ListRrcStreamRecord(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperRrcStreamRecord, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperRrcStreamRecord](
		ctx,
		s.Client(),
		routes.McastRrcStreamRecordPath,
		opts...,
	)
}

// ListRrcStreamAdmitRecord retrieves RRC stream admit record data from the controller.
func (s Service) ListRrcStreamAdmitRecord(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperRrcStreamAdmitRecord, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperRrcStreamAdmitRecord](
		ctx,
		s.Client(),
		routes.McastRrcStreamAdmitRecordPath, opts...,
	)
}

// ListRrcStreamDenyRecord retrieves RRC stream deny record data from the controller.
func (s Service) ListRrcStreamDenyRecord(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessMcastOperRrcStreamDenyRecord, error) {
	return core.Get[CiscoIOSXEWirelessMcastOperRrcStreamDenyRecord](
		ctx,
		s.Client(),
		routes.McastRrcStreamDenyRecordPath, opts...,
	)
}
