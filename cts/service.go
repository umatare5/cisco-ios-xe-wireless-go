package cts

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to CTS (Cisco TrustSec) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new CTS service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves CTS configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.CtsCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-cts-cfg:cts-cfg-data"

	var result model.CtsCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
