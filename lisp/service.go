package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to LISP (Locator/Identifier Separation Protocol) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new LISP service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves LISP operational data.
func (s *Service) Oper(ctx context.Context) (*model.LispOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-lisp-oper:lisp-oper-data"

	var result model.LispOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
