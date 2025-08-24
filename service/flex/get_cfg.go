package flex

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/flex"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for flex service
func (s Service) configOps() *core.ConfigOperations[model.FlexCfg] {
	return core.NewConfigOperations[model.FlexCfg](s.Client(), routes.FlexCfgBasePath)
}

// GetCfg retrieves FlexConnect configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.FlexCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgByPolicyName retrieves FlexConnect configuration data filtered by policy name.
func (s Service) GetCfgByPolicyName(ctx context.Context, policyName string) (*model.FlexCfg, error) {
	if err := validation.ValidateNonEmptyString(policyName, "policy name"); err != nil {
		return nil, err
	}
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.FlexCfgPolicyEntriesEndpoint, "flex-policy-entry", policyName)
	return core.Get[model.FlexCfg](ctx, s.Client(), url)
}
