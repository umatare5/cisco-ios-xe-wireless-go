package radio

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/radio"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for radio service
func (s Service) configOps() *core.ConfigOperations[model.RadioCfg] {
	return core.NewConfigOperations[model.RadioCfg](s.Client(), routes.RadioCfgBasePath)
}

// GetCfg retrieves all radio configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RadioCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgProfiles retrieves radio profiles configuration data.
func (s Service) GetCfgProfiles(ctx context.Context) (*model.RadioCfgRadioProfiles, error) {
	subOps := core.NewConfigOperations[model.RadioCfgRadioProfiles](s.Client(), routes.RadioCfgBasePath)
	return subOps.GetSubRes(ctx, "radio-profiles")
}

// GetCfgByName retrieves radio configuration data filtered by radio profile name.
func (s Service) GetCfgByName(ctx context.Context, name string) (*model.RadioCfg, error) {
	if err := validation.ValidateProfileName(name, "radio"); err != nil {
		return nil, err
	}
	subOps := core.NewConfigOperations[model.RadioCfg](s.Client(), routes.RadioCfgProfilesEndpoint)
	return subOps.GetByID(ctx, "radio-profile", name)
}
