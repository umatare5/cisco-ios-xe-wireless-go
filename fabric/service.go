package fabric

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Fabric operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Fabric service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves Fabric configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.FabricCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"

	var result model.FabricCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
