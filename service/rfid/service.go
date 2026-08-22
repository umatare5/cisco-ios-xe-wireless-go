package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides RFID (Radio Frequency Identification) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new RFID service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves RFID configuration data.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRFIDCfg, error) {
	return core.Get[CiscoIOSXEWirelessRFIDCfg](ctx, s.Client(), routes.RFIDCfgPath, opts...)
}

// GetConfigSettings retrieves the RFID configuration settings.
func (s Service) GetConfigSettings(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessRFIDCfgRFID, error) {
	return core.Get[CiscoIOSXEWirelessRFIDCfgRFID](ctx, s.Client(), routes.RFIDCfgRFIDConfigPath, opts...)
}

// GetGlobalOperational retrieves RFID global information.
func (s Service) GetGlobalOperational(ctx context.Context, opts ...core.GetOption) (*RFIDGlobalOper, error) {
	return core.Get[RFIDGlobalOper](ctx, s.Client(), routes.RFIDGlobalOperPath, opts...)
}

// GetGlobalDetailByMAC retrieves specific RFID data detail by MAC address.
func (s Service) GetGlobalDetailByMAC(
	ctx context.Context,
	macAddr string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRFIDGlobalOperRFIDDataDetail, error) {
	normalizedMAC, err := service.RequireMACAddress(macAddr)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataDetailQueryPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessRFIDGlobalOperRFIDDataDetail](ctx, s.Client(), url, opts...)
}

// GetRadioInfo retrieves RFID radio information by radio key combination.
func (s Service) GetRadioInfo(
	ctx context.Context,
	macAddr, apMACAddr string,
	slot int, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRFIDGlobalOperRFIDRadioData, error) {
	normalizedMAC, err := service.RequireMACAddress(macAddr)
	if err != nil {
		return nil, err
	}
	normalizedAPMAC, err := service.RequireMACAddress(apMACAddr)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.RFIDRadioDataPath,
		normalizedMAC,
		normalizedAPMAC,
		slot,
	)
	return core.Get[CiscoIOSXEWirelessRFIDGlobalOperRFIDRadioData](ctx, s.Client(), url, opts...)
}

// GetOperational retrieves RFID operational data.
func (s Service) GetOperational(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRFIDOper, error) {
	return core.Get[CiscoIOSXEWirelessRFIDOper](ctx, s.Client(), routes.RFIDOperPath, opts...)
}

// GetDetailByMAC retrieves specific RFID data based on MAC address.
func (s Service) GetDetailByMAC(
	ctx context.Context,
	macAddr string, opts ...core.GetOption,
) (*CiscoIOSXEWirelessRFIDOperRFIDData, error) {
	normalizedMAC, err := service.RequireMACAddress(macAddr)
	if err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataQueryPath, normalizedMAC)
	return core.Get[CiscoIOSXEWirelessRFIDOperRFIDData](ctx, s.Client(), url, opts...)
}
