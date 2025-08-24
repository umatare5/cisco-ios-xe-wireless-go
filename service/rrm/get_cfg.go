package rrm

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rrm"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for rrm service
func (s Service) configOps() *core.ConfigOperations[model.RrmCfg] {
	return core.NewConfigOperations[model.RrmCfg](s.Client(), routes.RRMCfgBasePath)
}

// GetCfg retrieves complete RRM configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RrmCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgByBand retrieves RRM configuration data filtered by band.
func (s Service) GetCfgByBand(ctx context.Context, band string) (*model.RrmCfg, error) {
	if err := validation.ValidateNonEmptyString(band, "band"); err != nil {
		return nil, err
	}
	subOps := core.NewConfigOperations[model.RrmCfg](s.Client(), routes.RRMByBandEndpoint)
	return subOps.GetByID(ctx, "band", band)
}

// GetCfgByMgrBand retrieves RRM manager configuration data filtered by band.
func (s Service) GetCfgByMgrBand(ctx context.Context, band string) (*model.RrmCfg, error) {
	if err := validation.ValidateNonEmptyString(band, "band"); err != nil {
		return nil, err
	}
	subOps := core.NewConfigOperations[model.RrmCfg](s.Client(), routes.RRMMgrByBandEndpoint)
	return subOps.GetByID(ctx, "band", band)
}
