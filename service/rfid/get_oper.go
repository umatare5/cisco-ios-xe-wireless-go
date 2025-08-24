package rfid

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rfid"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for rfid service
func (s Service) operOps() *core.OperationalOperations[model.RfidOper] {
	return core.NewOperationalOperations[model.RfidOper](s.Client(), routes.RFIDOperBasePath)
}

// GetOper retrieves RFID operational data.
func (s Service) GetOper(ctx context.Context) (*model.RfidOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByMac retrieves RFID operational data filtered by RFID MAC address.
func (s Service) GetOperByMac(ctx context.Context, macAddr string) (*model.RfidOper, error) {
	if err := validation.ValidateRFIDMAC(macAddr); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "rfid-data", macAddr)
}

// GetOperRfidData retrieves specific RFID data based on MAC address.
func (s Service) GetOperRfidData(ctx context.Context, macAddr string) (*model.RfidOperRfidData, error) {
	if err := validation.ValidateRFIDMAC(macAddr); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/rfid-data=%s", routes.RFIDOperBasePath, macAddr)
	return core.Get[model.RfidOperRfidData](ctx, s.Client(), url)
}
