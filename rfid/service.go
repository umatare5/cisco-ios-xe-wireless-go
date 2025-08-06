package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to RFID operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new RFID service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Cfg retrieves RFID configuration data.
func (s *Service) Cfg(ctx context.Context) (*model.RfidCfgResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"

	var result model.RfidCfgResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
