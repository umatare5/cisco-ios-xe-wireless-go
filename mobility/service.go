package mobility

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Mobility operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Mobility service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves Mobility operational data.
func (s *Service) Oper(ctx context.Context) (*model.MobilityOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"

	var result model.MobilityOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
