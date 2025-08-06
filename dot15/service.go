package dot15

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to 802.15 standard operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new 802.15 service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves 802.15 configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.Dot15CfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"

	var result model.Dot15CfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
