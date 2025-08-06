package apf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to APF (Application Policy Framework) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new APF service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves APF configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.ApfCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"

	var result model.ApfCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
