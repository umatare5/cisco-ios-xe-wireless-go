package ble

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// Service provides access to BLE (Bluetooth Low Energy) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new BLE service instance.
func NewService(client *wnc.Client) *Service {
	return &Service{c: client}
}

// Oper retrieves BLE operational data.
func (s *Service) Oper(ctx context.Context) (*model.BleOperResponse, error) {
	const endpoint = "Cisco-IOS-XE-wireless-ble-oper:ble-oper-data"

	var result model.BleOperResponse
	err := s.c.Do(ctx, "GET", endpoint, &result)
	return &result, err
}
