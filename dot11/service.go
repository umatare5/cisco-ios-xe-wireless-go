package dot11

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to 802.11 wireless standard operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new 802.11 service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves 802.11 configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.Dot11CfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data"

	var result model.Dot11CfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
