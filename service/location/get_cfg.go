package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/location"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for location service
func (s Service) configOps() *core.ConfigOperations[model.LocationCfg] {
	return core.NewConfigOperations[model.LocationCfg](s.Client(), routes.LocationCfgBasePath)
}

// GetCfg retrieves all location configuration data from the wireless controller.
// This function returns comprehensive location configuration information including
// profiles, servers, and settings.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Complete location configuration data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	data, err := service.GetCfg(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location configuration: %+v\n", data)
func (s Service) GetCfg(ctx context.Context) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return s.configOps().GetAll(ctx)
}

// GetCfgProfiles retrieves location profile configuration data from the wireless controller.
// This function returns information about configured location profiles.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Location profiles configuration data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	profiles, err := service.GetCfgProfiles(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location profiles: %+v\n", profiles)
func (s Service) GetCfgProfiles(ctx context.Context) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationCfgProfilesEndpoint)
}

// GetCfgServers retrieves location server configuration data from the wireless controller.
// This function returns information about configured location servers.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Location servers configuration data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	servers, err := service.GetCfgServers(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location servers: %+v\n", servers)
func (s Service) GetCfgServers(ctx context.Context) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationCfgServersEndpoint)
}

// GetCfgSettings retrieves location settings configuration data from the wireless controller.
// This function returns information about location service settings.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Location settings configuration data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	settings, err := service.GetCfgSettings(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location settings: %+v\n", settings)
func (s Service) GetCfgSettings(ctx context.Context) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationCfgSettingsEndpoint)
}

// GetCfgByProfileName retrieves location configuration data for a specific profile name.
// This function provides targeted access to location configuration for a particular profile.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//   - profileName: Specific location profile name to retrieve configuration for
//
// Returns:
//   - *model.LocationCfg: Location configuration data for the specified profile
//   - error: Error if the operation fails or profileName is invalid
//
// Example:
//
//	service := location.NewService(client)
//	data, err := service.GetCfgByProfileName(ctx, "profile-1")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location profile configuration: %+v\n", data)
func (s Service) GetCfgByProfileName(ctx context.Context, profileName string) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	if err := validation.ValidateNonEmptyString(profileName, "profile name"); err != nil {
		return nil, err
	}
	endpoint := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.LocationCfgProfilesEndpoint, "profile", profileName)
	return core.Get[model.LocationCfg](ctx, s.Client(), endpoint)
}

// GetCfgByServerName retrieves location configuration data for a specific server name.
// This function provides targeted access to location configuration for a particular server.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//   - serverName: Specific location server name to retrieve configuration for
//
// Returns:
//   - *model.LocationCfg: Location configuration data for the specified server
//   - error: Error if the operation fails or serverName is invalid
//
// Example:
//
//	service := location.NewService(client)
//	data, err := service.GetCfgByServerName(ctx, "server-1")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location server configuration: %+v\n", data)
func (s Service) GetCfgByServerName(ctx context.Context, serverName string) (*model.LocationCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	if err := validation.ValidateNonEmptyString(serverName, "server name"); err != nil {
		return nil, err
	}
	endpoint := s.Client().RestconfBuilder().BuildPathQueryURL(routes.LocationCfgServersEndpoint, "server", serverName)
	return core.Get[model.LocationCfg](ctx, s.Client(), endpoint)
}
