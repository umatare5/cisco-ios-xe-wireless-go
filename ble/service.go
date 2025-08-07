package ble

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// BleOperBasePath defines the base path for BLE operational data endpoints
	BleOperBasePath = constants.YANGModelPrefix + "ble-oper:ble-oper-data"
	// BleOperEndpoint defines the endpoint for BLE operational data
	BleOperEndpoint = BleOperBasePath
)

// Service provides BLE operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// Oper returns BLE operational data.
func (s Service) Oper(ctx context.Context) (*model.BleOperResponse, error) {
	var result model.BleOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, BleOperEndpoint, &result)
}
