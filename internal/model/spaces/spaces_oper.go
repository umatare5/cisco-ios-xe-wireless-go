package model

import "time"

// CiscoSpacesOper represents the complete Cisco Spaces operational data.
type CiscoSpacesOper struct {
	CiscoSpacesOperData *CiscoSpacesOperData `json:"Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data"`
}

// CiscoSpacesOperSpacesConnectionDetail represents the Cisco Spaces connection detail container.
type CiscoSpacesOperSpacesConnectionDetail struct {
	SpacesConnectionDetail *SpacesConnectionDetail `json:"Cisco-IOS-XE-wireless-cisco-spaces-oper:spaces-connection-detail,omitempty"`
}

// CiscoSpacesOperData represents the Cisco Spaces operational data container.
type CiscoSpacesOperData struct {
	SpacesConnectionDetail *SpacesConnectionDetail `json:"spaces-connection-detail,omitempty"`
}

// SpacesConnectionDetail represents Cisco wireless controller to spaces connection details.
type SpacesConnectionDetail struct {
	SpacesHealthURL     string                 `json:"spaces-health-url,omitempty"` // Cisco spaces health check URL
	ConnectionState     string                 `json:"con-state,omitempty"`         // Onboarding web-socket connection state
	LastConnectionError string                 `json:"last-conn-err,omitempty"`     // Last connection error
	ConnectionEstTime   time.Time              `json:"conn-estb-time,omitempty"`    // Connection establishment time
	Stats               *SpacesConnectionStats `json:"stats,omitempty"`             // Connection statistics
	Tenant              *SpacesTenant          `json:"tenant,omitempty"`            // Tenant details
}

// SpacesConnectionStats represents Cisco Spaces onboarding web-socket connection statistics.
type SpacesConnectionStats struct {
	TotalConnectionAttempts   uint64    `json:"total-con-attempts,omitempty"`   // Total connection attempts
	ConnectionAttemptsSuccess uint64    `json:"con-attempts-success,omitempty"` // Successful connection attempts
	ConnectionAttemptsFailure uint64    `json:"con-attempts-failure,omitempty"` // Failed connection attempts
	TotalDisconnections       uint64    `json:"total-discon,omitempty"`         // Total disconnections
	TotalMessagesReceived     uint64    `json:"total-msg-rcvd,omitempty"`       // Total messages received
	TotalMessagesSent         uint64    `json:"total-msg-sent,omitempty"`       // Total messages sent
	TotalDataReceived         uint64    `json:"total-data-rcvd,omitempty"`      // Total data received (bytes)
	TotalDataSent             uint64    `json:"total-data-sent,omitempty"`      // Total data sent (bytes)
	LastHeartbeatTime         time.Time `json:"last-heartbeat-time,omitempty"`  // Last heartbeat time
	AverageResponseTime       int       `json:"avg-response-time,omitempty"`    // Average response time (ms)
}

// SpacesTenant represents Cisco Spaces onboarding web-socket tenant details.
type SpacesTenant struct {
	TenantID            string    `json:"tenant-id,omitempty"`            // Tenant identifier
	TenantName          string    `json:"tenant-name,omitempty"`          // Tenant name
	OrganizationID      string    `json:"organization-id,omitempty"`      // Organization identifier
	OrganizationName    string    `json:"organization-name,omitempty"`    // Organization name
	RegistrationStatus  string    `json:"registration-status,omitempty"`  // Registration status
	LastSyncTime        time.Time `json:"last-sync-time,omitempty"`       // Last synchronization time
	SyncInterval        int       `json:"sync-interval,omitempty"`        // Sync interval (seconds)
	ConfigVersion       string    `json:"config-version,omitempty"`       // Configuration version
	CapabilitiesEnabled []string  `json:"capabilities-enabled,omitempty"` // Enabled capabilities
}

// SpacesCapabilities represents the capabilities available for Cisco Spaces integration.
type SpacesCapabilities struct {
	LocationAnalytics    bool `json:"location-analytics,omitempty"`    // Location analytics capability
	PresenceAnalytics    bool `json:"presence-analytics,omitempty"`    // Presence analytics capability
	BehaviorAnalytics    bool `json:"behavior-analytics,omitempty"`    // Behavior analytics capability
	AssetTracking        bool `json:"asset-tracking,omitempty"`        // Asset tracking capability
	GuestEngagement      bool `json:"guest-engagement,omitempty"`      // Guest engagement capability
	EnvironmentalMetrics bool `json:"environmental-metrics,omitempty"` // Environmental metrics capability
}

// SpacesLocationData represents location-based analytics data.
type SpacesLocationData struct {
	BuildingID     string             `json:"building-id,omitempty"`
	FloorID        string             `json:"floor-id,omitempty"`
	ZoneID         string             `json:"zone-id,omitempty"`
	DeviceCount    int                `json:"device-count,omitempty"`
	UniqueVisitors int                `json:"unique-visitors,omitempty"`
	DwellTime      int                `json:"dwell-time,omitempty"` // Average dwell time (minutes)
	Coordinates    *SpacesCoordinates `json:"coordinates,omitempty"`
	LastUpdated    time.Time          `json:"last-updated,omitempty"`
}

// SpacesCoordinates represents coordinate information for location services.
type SpacesCoordinates struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Accuracy  float64 `json:"accuracy,omitempty"` // Accuracy in meters
	Altitude  float64 `json:"altitude,omitempty"` // Altitude in meters
	Floor     int     `json:"floor,omitempty"`    // Floor number
}
