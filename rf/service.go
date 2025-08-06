package rf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to RF (Radio Frequency) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new RF service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves RF configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.RfCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"

	var result model.RfCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
