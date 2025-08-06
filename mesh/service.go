package mesh

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Mesh operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Mesh service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves Mesh operational data.
func (s *Service) Oper(ctx context.Context) (*model.MeshOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"

	var result model.MeshOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}

// Cfg retrieves Mesh configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.MeshCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"

	var result model.MeshCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
