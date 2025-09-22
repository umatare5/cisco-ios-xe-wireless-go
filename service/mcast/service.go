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
func (s Service) GetOperational(ctx context.Context) (*McastOper, error) {
	return core.Get[McastOper](ctx, s.Client(), routes.McastOperPath)
}

// GetFlexConnectMediastreamClientSummary retrieves FlexConnect mediastream client summary data from the wireless controller.
func (s Service) GetFlexConnectMediastreamClientSummary(
	ctx context.Context,
) (*McastOperFlexMediastreamClientSummary, error) {
	return core.Get[McastOperFlexMediastreamClientSummary](
		ctx, s.Client(), routes.McastFlexMediastreamPath,
	)
}

// ListVLANL2MGIDs retrieves VLAN Layer 2 multicast group ID operational data from the wireless controller.
func (s Service) ListVLANL2MGIDs(ctx context.Context) (*McastOperVlanL2MgidOp, error) {
	return core.Get[McastOperVlanL2MgidOp](ctx, s.Client(), routes.McastVlanL2MgidPath)
}

// GetFabricMediastreamClientSummary retrieves fabric mediastream client summary data from the wireless controller.
func (s Service) GetFabricMediastreamClientSummary(
	ctx context.Context,
) (*McastOperFabricMediaStreamClientSummary, error) {
	return core.Get[McastOperFabricMediaStreamClientSummary](
		ctx, s.Client(), routes.McastFabricMediastreamPath,
	)
}

// GetMcastMgidInfo retrieves multicast MGID information from the wireless controller.
func (s Service) GetMcastMgidInfo(ctx context.Context) (*McastOperMcastMgidInfo, error) {
	return core.Get[McastOperMcastMgidInfo](ctx, s.Client(), routes.McastMgidInfoPath)
}

// GetMulticastOperData retrieves multicast operational data from the wireless controller.
func (s Service) GetMulticastOperData(ctx context.Context) (*McastOperMulticastOperData, error) {
	return core.Get[McastOperMulticastOperData](ctx, s.Client(), routes.McastMulticastOperDataPath)
}

// ListRrcHistoryClientRecordData retrieves RRC history client record data from the controller.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListRrcHistoryClientRecordData(ctx context.Context) (*McastOperRrcHistoryClientRecordData, error) {
	return core.Get[McastOperRrcHistoryClientRecordData](ctx, s.Client(), routes.McastRrcHistoryClientRecordDataPath)
}

// ListRrcSrRadioRecord retrieves RRC stream radio record data from the controller.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListRrcSrRadioRecord(ctx context.Context) (*McastOperRrcSrRadioRecord, error) {
	return core.Get[McastOperRrcSrRadioRecord](ctx, s.Client(), routes.McastRrcSrRadioRecordPath)
}

// ListRrcStreamRecord retrieves RRC stream record data from the controller.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListRrcStreamRecord(ctx context.Context) (*McastOperRrcStreamRecord, error) {
	return core.Get[McastOperRrcStreamRecord](ctx, s.Client(), routes.McastRrcStreamRecordPath)
}

// ListRrcStreamAdmitRecord retrieves RRC stream admit record data from the controller.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListRrcStreamAdmitRecord(ctx context.Context) (*McastOperRrcStreamAdmitRecord, error) {
	return core.Get[McastOperRrcStreamAdmitRecord](ctx, s.Client(), routes.McastRrcStreamAdmitRecordPath)
}

// ListRrcStreamDenyRecord retrieves RRC stream deny record data from the controller.
// Note: Not Verified on IOS-XE 17.12.5 - may return 404 errors on some controller versions.
func (s Service) ListRrcStreamDenyRecord(ctx context.Context) (*McastOperRrcStreamDenyRecord, error) {
	return core.Get[McastOperRrcStreamDenyRecord](ctx, s.Client(), routes.McastRrcStreamDenyRecordPath)
}
