package fabric

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/fabric"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for fabric service
func (s Service) configOps() *core.ConfigOperations[model.FabricCfg] {
	return core.NewConfigOperations[model.FabricCfg](s.Client(), routes.FabricCfgBasePath)
}

// GetCfg retrieves Fabric configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.FabricCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgByFabricProfileName retrieves Fabric configuration data filtered by fabric profile name.
func (s Service) GetCfgByFabricProfileName(
	ctx context.Context, fabricProfileName string,
) (*model.FabricCfg, error) {
	if err := validation.ValidateNonEmptyString(fabricProfileName, "fabric profile name"); err != nil {
		return nil, err
	}
	subOps := core.NewConfigOperations[model.FabricCfg](s.Client(), routes.FabricCfgProfilesEndpoint)
	return subOps.GetByID(ctx, "fabric-profile", fabricProfileName)
}

// GetCfgByControlPlaneName retrieves Fabric configuration data filtered by control plane name.
func (s Service) GetCfgByControlPlaneName(
	ctx context.Context, controlPlaneName string,
) (*model.FabricCfg, error) {
	if err := validation.ValidateNonEmptyString(controlPlaneName, "control plane name"); err != nil {
		return nil, err
	}
	subOps := core.NewConfigOperations[model.FabricCfg](s.Client(), routes.FabricCfgControlplaneNamesEndpoint)
	return subOps.GetByID(ctx, "fabric-controlplane-name", controlPlaneName)
}
