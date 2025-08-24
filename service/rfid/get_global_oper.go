package rfid

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rfid"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// GetGlobalOper retrieves RFID global operational data.
func (s Service) GetGlobalOper(ctx context.Context) (*model.RfidGlobalOper, error) {
	return core.Get[model.RfidGlobalOper](ctx, s.Client(), routes.RFIDGlobalOperEndpoint)
}

// GetGlobalOperByMac retrieves RFID global operational data filtered by RFID MAC address.
func (s Service) GetGlobalOperByMac(ctx context.Context, macAddr string) (*model.RfidGlobalOper, error) {
	if err := validation.ValidateRFIDMAC(macAddr); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/rfid-data-detail=%s", routes.RFIDGlobalOperBasePath, macAddr)
	return core.Get[model.RfidGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByRadioKey retrieves RFID global operational data filtered by radio key combination.
func (s Service) GetGlobalOperByRadioKey(
	ctx context.Context,
	macAddr, apMacAddr string,
	slot int,
) (*model.RfidGlobalOper, error) {
	if err := validation.ValidateBothMACs(macAddr, apMacAddr); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/rfid-radio-data=%s,%s,%d", routes.RFIDGlobalOperBasePath, macAddr, apMacAddr, slot)
	return core.Get[model.RfidGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperRfidDataDetail retrieves specific RFID data detail by MAC address.
func (s Service) GetGlobalOperRfidDataDetail(
	ctx context.Context,
	macAddr string,
) (*model.RfidGlobalOperRfidDataDetail, error) {
	if err := validation.ValidateRFIDMAC(macAddr); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/rfid-data-detail=%s", routes.RFIDGlobalOperBasePath, macAddr)
	return core.Get[model.RfidGlobalOperRfidDataDetail](ctx, s.Client(), url)
}

// GetGlobalOperRfidRadioData retrieves RFID radio data by radio key combination.
func (s Service) GetGlobalOperRfidRadioData(
	ctx context.Context,
	macAddr, apMacAddr string,
	slot int,
) (*model.RfidGlobalOperRfidRadioData, error) {
	if err := validation.ValidateBothMACs(macAddr, apMacAddr); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/rfid-radio-data=%s,%s,%d", routes.RFIDGlobalOperBasePath, macAddr, apMacAddr, slot)
	return core.Get[model.RfidGlobalOperRfidRadioData](ctx, s.Client(), url)
}
