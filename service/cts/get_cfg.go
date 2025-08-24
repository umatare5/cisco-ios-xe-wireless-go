package cts

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/cts"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for cts service
func (s Service) configOps() *core.ConfigOperations[model.CtsCfg] {
	return core.NewConfigOperations[model.CtsCfg](s.Client(), routes.CTSCfgBasePath)
}

// GetCfg retrieves CTS SXP configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.CtsCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgBySxpProfileName retrieves CTS configuration data filtered by SXP profile name.
func (s Service) GetCfgBySxpProfileName(
	ctx context.Context, sxpProfileName string,
) (*model.CtsCfgFilter, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateNonEmptyString(sxpProfileName, "SXP profile name"); err != nil {
		return nil, ierrors.ServiceOperationError(ierrors.ActionValidate, EntityTypeCTS, "SXP profile name", err)
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.CTSCfgBasePath, "cts-sxp-profile", sxpProfileName)
	return core.Get[model.CtsCfgFilter](ctx, s.Client(), url)
}
