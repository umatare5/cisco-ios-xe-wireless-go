package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
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
func (s Service) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessRFIDCfg, error) {
	return core.Get[CiscoIOSXEWirelessRFIDCfg](ctx, s.Client(), routes.RFIDCfgPath)
}

// GetConfigSettings retrieves the RFID configuration settings.
func (s Service) GetConfigSettings(ctx context.Context) (*RFIDConfig, error) {
	return core.Get[RFIDConfig](ctx, s.Client(), routes.RFIDCfgRFIDConfigPath)
}

// GetGlobalOperational retrieves RFID global information.
func (s Service) GetGlobalOperational(ctx context.Context) (*RFIDGlobalOper, error) {
	return core.Get[RFIDGlobalOper](ctx, s.Client(), routes.RFIDGlobalOperPath)
}

// GetGlobalDetailByMAC retrieves specific RFID data detail by MAC address.
func (s Service) GetGlobalDetailByMAC(
	ctx context.Context,
	macAddr string,
) (*RFIDEmltdData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataDetailQueryPath, macAddr)
	return core.Get[RFIDEmltdData](ctx, s.Client(), url)
}

// GetRadioInfo retrieves RFID radio information by radio key combination.
func (s Service) GetRadioInfo(
	ctx context.Context,
	macAddr, apMACAddr string,
	slot int,
) (*RFIDRadioData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}
	if err := validation.ValidateMACAddress(apMACAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.RFIDRadioDataPath,
		macAddr,
		apMACAddr,
		slot,
	)
	return core.Get[RFIDRadioData](ctx, s.Client(), url)
}

// GetOperational retrieves RFID operational data.
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessRFIDOper, error) {
	return core.Get[CiscoIOSXEWirelessRFIDOper](ctx, s.Client(), routes.RFIDOperPath)
}

// GetDetailByMAC retrieves specific RFID data based on MAC address.
func (s Service) GetDetailByMAC(ctx context.Context, macAddr string) (*RFIDData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataQueryPath, macAddr)
	return core.Get[RFIDData](ctx, s.Client(), url)
}
