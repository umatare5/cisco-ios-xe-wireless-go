// Package location provides data models for location configuration data.
package location

type LocationCfgData struct {
	// Live WNC 17.12.5: contains nmsp-config (empty object {})
	NmspConfig *NmspConfig `json:"nmsp-config,omitempty"`
	// YANG 17.12.1: location container (not present in live data)
	LocationConfig *LocationConfig `json:"location,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// YANG 17.12.1: operator-locations container (not present in live data)
	OperatorLocations *OperatorLocations `json:"operator-locations,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// LocationCfg represents the structure for Location configuration data from WNC 17.12.5.
type LocationCfg struct {
	LocationCfgData LocationCfgData `json:"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"`
}

// OperatorLocations represents AAA Location Operators container.
type OperatorLocations struct {
	// List of operator location configurations
	OperatorLocation []OperatorLocation `json:"operator-location,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// LocationServerConfigResponse represents the response structure for nmsp-config endpoint.
type LocationServerConfigResponse struct {
	NmspConfig *NmspConfig `json:"Cisco-IOS-XE-wireless-location-cfg:nmsp-config,omitempty"`
}

// LocationSettings represents the response structure for location settings endpoint (HTTP 204).
type LocationSettings struct {
	LocationConfig *LocationConfig `json:"Cisco-IOS-XE-wireless-location-cfg:location,omitempty"`
}

// LocationCfgLocationConfig represents the corresponding data structure for location config from WNC 17.12.5.
type LocationCfgLocationConfig struct {
	LocationConfig *LocationConfig `json:"Cisco-IOS-XE-wireless-location-cfg:location,omitempty"`
}

// LocationCfgNmspConfig represents the corresponding data structure for NMSP config from WNC 17.12.5.
type LocationCfgNmspConfig struct {
	NmspConfig *NmspConfig `json:"Cisco-IOS-XE-wireless-location-cfg:nmsp-config,omitempty"`
}

type LocationConfig struct {
	// Location algorithm enum type
	LocationAlgorithm *string `json:"location-algorithm,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Notification threshold for clients (0-10)
	LocationNotifyClientThreshold *uint32 `json:"location-notify-client-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI threshold for clients
	LocationRssiClientThreshold *string `json:"location-rssi-client-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI threshold for calibrating clients
	LocationRssiCalClientThreshold *string `json:"location-rssi-cal-client-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI threshold for rogue APs
	LocationRssiRogueApThreshold *string `json:"location-rssi-rogue-ap-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Expiry timeout for RSSI values (5-3600 seconds)
	LocationExpiryClientThreshold *uint32 `json:"location-expiry-client-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Expiry timeout for calibrating clients (1-3600 seconds)
	LocationExpiryCalClientThreshold *uint32 `json:"location-expiry-cal-client-threshold,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// NmspConfig represents NMSP notification parameters.
type NmspConfig struct {
	// Enable NMSP server
	Enable *bool `json:"enable,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Enable strong ciphers for NMSP server
	IsSecure *bool `json:"is-secure,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// RSSI measurement notify interval (1-180 seconds)
	RssiMeasurementNotifyInterval *uint32 `json:"rssi-measurement-notify-interval,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Client notification interval (1-180 seconds)
	ClientNotifyInterval *uint32 `json:"client-notify-interval,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Rogue client notification interval (1-180 seconds)
	RogueClientNotifyInterval *uint32 `json:"rogue-client-notify-interval,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Rogue AP notification interval (1-180 seconds)
	RogueApNotifyInterval *uint32 `json:"rogue-ap-notify-interval,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Spectrum notification interval (1-180 seconds)
	SpectrumNotifyInterval *uint32 `json:"spectrum-notify-interval,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// CMX cloud parameters
	CloudParams *NmspCloudParams `json:"cloud-params,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// NmspCloudParams represents CMX cloud service parameters.
type NmspCloudParams struct {
	// Enable NMSP CMX cloud services
	Enable *bool `json:"enable,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Server URL for CMX cloud services
	ServerURL *string `json:"server-url,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Hostname for proxy
	ProxyHostname *string `json:"proxy-hostname,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Port for proxy (1-65535)
	ProxyPort *uint16 `json:"proxy-port,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// Authentication token for CMX cloud services
	AuthToken *string `json:"auth-token,omitempty"` // (YANG: IOS-XE 17.12.1+)
}

// OperatorLocation represents AAA Location Operator parameters.
type OperatorLocation struct {
	// AAA Location Operator ID (key, 0-215 chars)
	LocOperID string `json:"loc-oper-id"` // (YANG: IOS-XE 17.12.1+)
	// AAA Location Operator Name (0-250 chars)
	LocOperName *string `json:"loc-oper-name,omitempty"` // (YANG: IOS-XE 17.12.1+)
	// AAA Location Namespace ID enum
	LocNsID *string `json:"loc-ns-id,omitempty"` // (YANG: IOS-XE 17.12.1+)
}
