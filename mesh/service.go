package mesh

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// MeshOperBasePath defines the base path for mesh operational endpoints
	MeshOperBasePath = "Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"
	// MeshOperEndpoint defines the endpoint for mesh operational data
	MeshOperEndpoint = MeshOperBasePath

	// MeshCfgBasePath defines the base path for mesh configuration endpoints
	MeshCfgBasePath = "Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"
	// MeshCfgEndpoint defines the endpoint for mesh configuration data
	MeshCfgEndpoint = MeshCfgBasePath
)

// Service provides Mesh operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Oper returns Mesh operational data.
func (s Service) Oper(ctx context.Context) (*model.MeshOperResponse, error) {
	var result model.MeshOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, MeshOperEndpoint, &result)
}

// Cfg returns Mesh configuration data.
func (s Service) Cfg(ctx context.Context) (*model.MeshCfgResponse, error) {
	var result model.MeshCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, MeshCfgEndpoint, &result)
}
