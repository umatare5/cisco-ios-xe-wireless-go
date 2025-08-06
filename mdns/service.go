// Package mdns provides multicast DNS service for the Cisco Wireless Network Controller API.
// This package implements the Domain Service layer of the three-layer architecture.
package mdns

import (
	"context"
	"net/http"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

// Service provides mDNS (multicast DNS) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new mDNS service instance.
func NewService(c *wnc.Client) *Service {
	return &Service{c: c}
}

// Oper retrieves mDNS operational data.
func (s *Service) Oper(ctx context.Context) (*model.MdnsOperResponse, error) {
	var out model.MdnsOperResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, MdnsOperEndpoint, &out)
	return &out, err
}

// GlobalStats retrieves mDNS global statistics.
func (s *Service) GlobalStats(ctx context.Context) (*model.MdnsGlobalStatsResponse, error) {
	var out model.MdnsGlobalStatsResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, MdnsGlobalStatsEndpoint, &out)
	return &out, err
}

// WlanStats retrieves mDNS WLAN statistics.
func (s *Service) WlanStats(ctx context.Context) (*model.MdnsWlanStatsResponse, error) {
	var out model.MdnsWlanStatsResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, MdnsWlanStatsEndpoint, &out)
	return &out, err
}
