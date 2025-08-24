package mesh

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mesh"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for mesh service
func (s Service) operOps() *core.OperationalOperations[model.MeshOper] {
	return core.NewOperationalOperations[model.MeshOper](s.Client(), routes.MeshOperBasePath)
}

// GetOper retrieves all mesh operational data.
func (s Service) GetOper(ctx context.Context) (*model.MeshOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperNodes retrieves mesh nodes operational data.
func (s Service) GetOperNodes(ctx context.Context) (*model.MeshOperMeshNodes, error) {
	return core.Get[model.MeshOperMeshNodes](ctx, s.Client(), routes.MeshOperNodesEndpoint)
}

// GetOperStats retrieves mesh statistics operational data.
func (s Service) GetOperStats(ctx context.Context) (*model.MeshOperMeshStats, error) {
	return core.Get[model.MeshOperMeshStats](ctx, s.Client(), routes.MeshOperStatsEndpoint)
}

// GetCfg retrieves all mesh configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.MeshCfg, error) {
	return core.Get[model.MeshCfg](ctx, s.Client(), routes.MeshCfgEndpoint)
}

// GetCfgByProfileName retrieves mesh configuration data filtered by profile name.
func (s Service) GetCfgByProfileName(ctx context.Context, profileName string) (*model.MeshCfg, error) {
	if err := validation.ValidateNonEmptyString(profileName, "profile name"); err != nil {
		return nil, err
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.MeshCfgEndpoint+"/mesh-profiles", "mesh-profile", profileName)
	return core.Get[model.MeshCfg](ctx, s.Client(), url)
}
