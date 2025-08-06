package awips

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to AWIPS (Advanced Weather Interactive Processing System) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new AWIPS service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves AWIPS operational data.
func (s *Service) Oper(ctx context.Context) (*model.AwipsOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"

	var result model.AwipsOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
