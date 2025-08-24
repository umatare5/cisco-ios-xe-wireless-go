// Package ap provides access point functionality for the Cisco IOS-XE Wireless Network Controller API.
package ap

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for AP service
func (s Service) operOps() *core.OperationalOperations[model.ApOper] {
	return core.NewOperationalOperations[model.ApOper](s.Client(), routes.APOperBasePath)
}

// GetOper retrieves the complete AP operational data
func (s Service) GetOper(ctx context.Context) (*model.ApOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperCapwapDataAll retrieves CAPWAP protocol data
func (s Service) GetOperCapwapDataAll(ctx context.Context) (*model.ApOperCapwapData, error) {
	return core.Get[model.ApOperCapwapData](ctx, s.Client(), routes.CapwapDataEndpoint)
}

// GetOperRadioStatusAll retrieves radio operational status data
func (s Service) GetOperRadioStatusAll(ctx context.Context) (*model.ApOper, error) {
	return core.Get[model.ApOper](ctx, s.Client(), routes.RadioOperDataEndpoint)
}

// GetOperCapwapDataByMac retrieves CAPWAP data for a specific WTP MAC
func (s Service) GetOperCapwapDataByMac(ctx context.Context, wtpMac string) (*model.ApOperCapwapData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CapwapDataEndpoint+"/capwap-data", wtpMac)
	return core.Get[model.ApOperCapwapData](ctx, s.Client(), url)
}

// GetOperNameMacMapByWtpName retrieves AP name-to-MAC mapping filtered by WTP name
func (s Service) GetOperNameMacMapByWtpName(
	ctx context.Context,
	wtpName string,
) (*model.ApOperApNameMacMap, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.NameMacMapEndpoint+"/ap-name-mac-map", wtpName)
	return core.Get[model.ApOperApNameMacMap](ctx, s.Client(), url)
}

// GetOperRadioStatusBySlot retrieves radio operational data by WTP MAC and slot ID
func (s Service) GetOperRadioStatusBySlot(
	ctx context.Context,
	wtpMac string,
	slotID int,
) (*model.ApOper, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryCompositeURL(routes.RadioOperDataEndpoint, wtpMac, slotID)
	return core.Get[model.ApOper](ctx, s.Client(), url)
}

// GetOperApRadioNeighborBySlotBssid retrieves AP radio neighbor information by slot and BSSID
func (s Service) GetOperApRadioNeighborBySlotBssid(
	ctx context.Context,
	apMac string,
	slotID int,
	bssid string,
) (*model.ApOperApRadioNeighbor, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildCompositeKeyURL(
		routes.APOperEndpoint, "ap-radio-neighbor", apMac, slotID, bssid)
	return core.Get[model.ApOperApRadioNeighbor](ctx, s.Client(), url)
}

// GetOperApIoxOperData retrieves AP IOx operational data by AP MAC address
func (s Service) GetOperApIoxOperData(ctx context.Context, apMac string) (*model.ApOperApIoxOperData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.APOperEndpoint, "ap-iox-oper-data", apMac)
	return core.Get[model.ApOperApIoxOperData](ctx, s.Client(), url)
}

// GetOperApImageActiveLocationOnly retrieves active image location information using fields parameter
func (s Service) GetOperApImageActiveLocationOnly(
	ctx context.Context,
) (*model.ApOperApImageActiveLocation, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APOperEndpoint, "ap-image-active-location")
	return core.Get[model.ApOperApImageActiveLocation](ctx, s.Client(), url)
}

// GetOperApImagePrepareLocationOnly retrieves only AP image prepare location data using fields parameter
func (s Service) GetOperApImagePrepareLocationOnly(
	ctx context.Context,
) (*model.ApOperApImagePrepareLocation, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APOperEndpoint, "ap-image-prepare-location")
	return core.Get[model.ApOperApImagePrepareLocation](ctx, s.Client(), url)
}

// GetOperApPwrInfoOnly retrieves only AP power information using fields parameter
func (s Service) GetOperApPwrInfoOnly(ctx context.Context) (*model.ApOperApPwrInfo, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APOperEndpoint, "ap-pwr-info")
	return core.Get[model.ApOperApPwrInfo](ctx, s.Client(), url)
}

// GetOperApSensorStatusOnly retrieves only AP sensor status using fields parameter
func (s Service) GetOperApSensorStatusOnly(ctx context.Context) (*model.ApOperApSensorStatus, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APOperEndpoint, "ap-sensor-status")
	return core.Get[model.ApOperApSensorStatus](ctx, s.Client(), url)
}

// GetOperCapwapPktsOnly retrieves only CAPWAP packets data using fields parameter
func (s Service) GetOperCapwapPktsOnly(ctx context.Context) (*model.ApOperCapwapPkts, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildFieldsURL(routes.APOperEndpoint, "capwap-pkts")
	return core.Get[model.ApOperCapwapPkts](ctx, s.Client(), url)
}
