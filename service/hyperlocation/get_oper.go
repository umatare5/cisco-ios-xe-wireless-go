package hyperlocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/hyperlocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for hyperlocation service
func (s Service) operOps() *core.OperationalOperations[model.HyperlocationOper] {
	return core.NewOperationalOperations[model.HyperlocationOper](s.Client(), routes.HyperlocationOperBasePath)
}

// GetOper retrieves hyperlocation operational data.
func (s Service) GetOper(ctx context.Context) (*model.HyperlocationOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByName retrieves hyperlocation operational data filtered by profile name.
func (s Service) GetOperByName(ctx context.Context, name string) (*model.HyperlocationOper, error) {
	if err := validation.ValidateNonEmptyString(name, "profile name"); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "ewlc-hyperlocation-profile", name)
}

// GetOperProfiles retrieves hyperlocation profiles.
func (s Service) GetOperProfiles(ctx context.Context) (*model.HyperlocationProfiles, error) {
	return core.Get[model.HyperlocationProfiles](ctx, s.Client(), routes.HyperlocationProfilesEndpoint)
}
