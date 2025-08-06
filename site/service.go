package site

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to Site operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new Site service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves Site operational data.
func (s *Service) Oper(ctx context.Context) (*model.SiteOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-site-oper:site-oper-data"

	var result model.SiteOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
