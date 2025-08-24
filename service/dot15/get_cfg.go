package dot15

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/dot15"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// configOps provides high-level configuration operations for dot15 service
func (s Service) configOps() *core.ConfigOperations[model.Dot15Cfg] {
	return core.NewConfigOperations[model.Dot15Cfg](s.Client(), routes.Dot15CfgBasePath)
}

// GetCfg retrieves 802.15 configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.Dot15Cfg, error) {
	return s.configOps().GetAll(ctx)
}
