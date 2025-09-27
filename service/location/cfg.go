package location

// CiscoIOSXEWirelessLocationCfg represents the structure for Location configuration data.
type CiscoIOSXEWirelessLocationCfg struct {
	CiscoIOSXEWirelessLocationCfgData struct {
		NMSPConfig        NMSPConfig         `json:"nmsp-config"`                  // NMSP notification parameters (Live: IOS-XE 17.12.6a)
		LocationConfig    *LocationConfig    `json:"location,omitempty"`           // Location config parameters (YANG: IOS-XE 17.12.1)
		OperatorLocations *OperatorLocations `json:"operator-locations,omitempty"` // AAA Location Operators (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"` // Location configuration data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessLocationSettings represents the response structure for location settings endpoint (HTTP 204).
type CiscoIOSXEWirelessLocationSettings struct {
	LocationConfig *LocationConfig `json:"Cisco-IOS-XE-wireless-location-cfg:location,omitempty"`
}

// CiscoIOSXEWirelessLocationCfgNMSPConfig represents the corresponding data structure for NMSP config.
type CiscoIOSXEWirelessLocationCfgNMSPConfig struct {
	NMSPConfig NMSPConfig `json:"Cisco-IOS-XE-wireless-location-cfg:nmsp-config"`
}

// OperatorLocations represents AAA Location Operators container.
type OperatorLocations struct {
	OperatorLocation []OperatorLocation `json:"operator-location,omitempty"`
}

type LocationConfig struct {
	LocationAlgorithm                *string `json:"location-algorithm,omitempty"`                   // Algorithm to average RSSI and SNR values (YANG: IOS-XE 17.12.1)
	LocationNotifyClientThreshold    *uint32 `json:"location-notify-client-threshold,omitempty"`     // NMSP notification threshold for clients (YANG: IOS-XE 17.12.1)
	LocationRSSIClientThreshold      *string `json:"location-rssi-client-threshold,omitempty"`       // NMSP notification threshold for clients (YANG: IOS-XE 17.12.1)
	LocationRSSICalClientThreshold   *string `json:"location-rssi-cal-client-threshold,omitempty"`   // Half life for calibrating clients (YANG: IOS-XE 17.12.1)
	LocationRSSIRogueApThreshold     *string `json:"location-rssi-rogue-ap-threshold,omitempty"`     // Half life for Rogue APs (YANG: IOS-XE 17.12.1)
	LocationExpiryClientThreshold    *uint32 `json:"location-expiry-client-threshold,omitempty"`     // Timeout for RSSI values (YANG: IOS-XE 17.12.1)
	LocationExpiryCalClientThreshold *uint32 `json:"location-expiry-cal-client-threshold,omitempty"` // Timeout for calibrating clients (YANG: IOS-XE 17.12.1)
}

// NMSPConfig represents NMSP notification parameters.
type NMSPConfig struct {
	Enable                        *bool            `json:"enable,omitempty"`                           // Enable NMSP server (YANG: IOS-XE 17.12.1)
	IsSecure                      *bool            `json:"is-secure,omitempty"`                        // Enable strong ciphers for NMSP server (YANG: IOS-XE 17.12.1)
	RSSIMeasurementNotifyInterval *uint32          `json:"rssi-measurement-notify-interval,omitempty"` // RSSI measurement notify interval (YANG: IOS-XE 17.12.1)
	ClientNotifyInterval          *uint32          `json:"client-notify-interval,omitempty"`           // Measurement interval for clients in seconds (YANG: IOS-XE 17.12.1)
	RogueClientNotifyInterval     *uint32          `json:"rogue-client-notify-interval,omitempty"`     // Measurement interval for rogue clients (YANG: IOS-XE 17.12.1)
	RogueApNotifyInterval         *uint32          `json:"rogue-ap-notify-interval,omitempty"`         // Measurement interval for rogue APs (YANG: IOS-XE 17.12.1)
	SpectrumNotifyInterval        *uint32          `json:"spectrum-notify-interval,omitempty"`         // Measurement interval for spectrum interferers (YANG: IOS-XE 17.12.1)
	CloudParams                   *NMSPCloudParams `json:"cloud-params,omitempty"`                     // Parameters for CMX cloud (YANG: IOS-XE 17.12.1)
}

// NMSPCloudParams represents CMX cloud service parameters.
type NMSPCloudParams struct {
	Enable        *bool   `json:"enable,omitempty"`         // Enable NMSP CMX cloud services (YANG: IOS-XE 17.12.1)
	ServerURL     *string `json:"server-url,omitempty"`     // Server URL for CMX cloud services (YANG: IOS-XE 17.12.1)
	ProxyHostname *string `json:"proxy-hostname,omitempty"` // Hostname for the proxy (YANG: IOS-XE 17.12.1)
	ProxyPort     *uint16 `json:"proxy-port,omitempty"`     // Port to use for the proxy (YANG: IOS-XE 17.12.1)
	AuthToken     *string `json:"auth-token,omitempty"`     // Authentication token for the CMX cloud services (YANG: IOS-XE 17.12.1)
}

// OperatorLocation represents AAA Location Operator parameters.
type OperatorLocation struct {
	LocOperID   string  `json:"loc-oper-id"`             // AAA Location Operator ID (YANG: IOS-XE 17.12.1)
	LocOperName *string `json:"loc-oper-name,omitempty"` // AAA Location Operator Name (YANG: IOS-XE 17.12.1)
	LocNsID     *string `json:"loc-ns-id,omitempty"`     // AAA Location Namespace ID (YANG: IOS-XE 17.12.1)
}
