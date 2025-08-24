// Package apf provides Application Policy Framework configuration operations for the Cisco IOS-XE Wireless Network Controller API.
package apf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/apf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// configOps provides high-level configuration operations for apf service
func (s Service) configOps() *core.ConfigOperations[model.ApfCfg] {
	return core.NewConfigOperations[model.ApfCfg](s.Client(), routes.APFCfgBasePath)
}

// GetCfg retrieves the complete APF configuration data
func (s Service) GetCfg(ctx context.Context) (*model.ApfCfg, error) {
	return s.configOps().GetAll(ctx)
}
