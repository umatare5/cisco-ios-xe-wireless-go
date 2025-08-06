package radio

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Radio operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Radio service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves Radio configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.RadioCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"

	var result model.RadioCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
