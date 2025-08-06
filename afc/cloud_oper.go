// Package afc provides Automated Frequency Coordination cloud operational data functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"
	"fmt"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

// Constants for backward compatibility with tests
const (
	// AfcCloudOperBasePath defines the base path for AFC cloud operational data endpoints.
	AfcCloudOperBasePath = "Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data"
	// AfcCloudOperEndpoint defines the endpoint for AFC cloud operational data.
	AfcCloudOperEndpoint = AfcCloudOperBasePath
	// AfcCloudStatsEndpoint defines the endpoint for AFC cloud statistics.
	AfcCloudStatsEndpoint = AfcCloudOperBasePath + "/afc-cloud-stats"
)

// Deprecated: use service-based API instead.
func GetAfcCloudOper(client *wnc.Client, ctx context.Context) (*model.AfcCloudOperResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.CloudOper(ctx)
}

// Deprecated: use service-based API instead.
func GetAfcCloudStats(client *wnc.Client, ctx context.Context) (*model.AfcCloudOperAfcCloudStatsResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
	}
	service := NewService(client.CoreClient())
	return service.CloudStats(ctx)
}
