package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rfid"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// configOps provides high-level configuration operations for rfid service
func (s Service) configOps() *core.ConfigOperations[model.RfidCfg] {
	return core.NewConfigOperations[model.RfidCfg](s.Client(), routes.RFIDCfgBasePath)
}

// GetCfg retrieves RFID configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RfidCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgRfidConfig retrieves the RFID configuration settings.
func (s Service) GetCfgRfidConfig(ctx context.Context) (*model.RfidCfgRfidConfig, error) {
	subOps := core.NewConfigOperations[model.RfidCfgRfidConfig](s.Client(), routes.RFIDCfgBasePath)
	return subOps.GetSubRes(ctx, "rfid-config")
}
