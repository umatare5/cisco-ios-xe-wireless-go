package nmsp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/nmsp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for nmsp service
func (s Service) operOps() *core.OperationalOperations[model.NmspOper] {
	return core.NewOperationalOperations[model.NmspOper](s.Client(), routes.NMSPOperBasePath)
}

// GetOper retrieves all NMSP operational data.
func (s Service) GetOper(ctx context.Context) (*model.NmspOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperClientRegistration retrieves NMSP client registration data.
func (s Service) GetOperClientRegistration(ctx context.Context) (*model.NmspClientRegistration, error) {
	return core.Get[model.NmspClientRegistration](ctx, s.Client(), routes.EndpointGetOperClientRegistration)
}

// GetOperCmxConnection retrieves NMSP CMX connection data.
func (s Service) GetOperCmxConnection(ctx context.Context) (*model.NmspCmxConnection, error) {
	return core.Get[model.NmspCmxConnection](ctx, s.Client(), routes.EndpointGetOperCmxConnection)
}

// GetOperCmxCloudInfo retrieves NMSP CMX cloud information.
func (s Service) GetOperCmxCloudInfo(ctx context.Context) (*model.NmspCmxCloudInfo, error) {
	return core.Get[model.NmspCmxCloudInfo](ctx, s.Client(), routes.EndpointGetOperCmxCloudInfo)
}

// GetOperByClientID retrieves NMSP operational data filtered by client registration ID.
func (s Service) GetOperByClientID(ctx context.Context, clientID string) (*model.NmspOper, error) {
	if err := validation.ValidateNonEmptyString(clientID, "client ID"); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "client-registration", clientID)
}
