package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rfid"
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
func (s Service) GetConfig(ctx context.Context) (*model.RfidCfg, error) {
	return core.Get[model.RfidCfg](ctx, s.Client(), routes.RFIDCfgPath)
}

// GetConfigSettings retrieves the RFID configuration settings.
func (s Service) GetConfigSettings(ctx context.Context) (*model.RfidConfig, error) {
	url := routes.RFIDCfgPath + "/rfid-config"
	return core.Get[model.RfidConfig](ctx, s.Client(), url)
}

// GetGlobalInfo retrieves RFID global information.
func (s Service) GetGlobalInfo(ctx context.Context) (*model.RfidGlobalOper, error) {
	return core.Get[model.RfidGlobalOper](ctx, s.Client(), routes.RFIDGlobalOperPath)
}

// GetGlobalDetailByMAC retrieves specific RFID data detail by MAC address.
func (s Service) GetGlobalDetailByMAC(
	ctx context.Context,
	macAddr string,
) (*model.RfidEmltdData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataDetailQueryPath, macAddr)
	return core.Get[model.RfidEmltdData](ctx, s.Client(), url)
}

// GetRadioInfo retrieves RFID radio information by radio key combination.
func (s Service) GetRadioInfo(
	ctx context.Context,
	macAddr, apMacAddr string,
	slot int,
) (*model.RfidRadioData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}
	if err := validation.ValidateMACAddress(apMacAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryCompositeURL(
		routes.RFIDRadioDataPath,
		macAddr,
		apMacAddr,
		slot,
	)
	return core.Get[model.RfidRadioData](ctx, s.Client(), url)
}

// GetOperational retrieves RFID operational data.
func (s Service) GetOperational(ctx context.Context) (*model.RfidOper, error) {
	return core.Get[model.RfidOper](ctx, s.Client(), routes.RFIDOperPath)
}

// GetDetailByMAC retrieves specific RFID data based on MAC address.
func (s Service) GetDetailByMAC(ctx context.Context, macAddr string) (*model.RfidData, error) {
	if err := validation.ValidateMACAddress(macAddr); err != nil {
		return nil, err
	}

	url := s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFIDDataQueryPath, macAddr)
	return core.Get[model.RfidData](ctx, s.Client(), url)
}
