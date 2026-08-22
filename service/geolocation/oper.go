package geolocation

// CiscoIOSXEWirelessGeolocationOper represents the geolocation operational data.
type CiscoIOSXEWirelessGeolocationOper struct {
	CiscoIOSXEWirelessGeolocationOperData *CiscoIOSXEWirelessGeolocationOperData `json:"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"` // Geolocation operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessGeolocationOperData represents Geolocation operational data (Live: IOS-XE 17.12.6a).
type CiscoIOSXEWirelessGeolocationOperData struct {
	ApGeoLocData  []ApGeoLocData `json:"ap-geo-loc-data,omitempty"` // AP geolocation data list (YANG: IOS-XE 17.12.1)
	ApGeoLocStats *ApGeoLocStats `json:"ap-geo-loc-stats"`          // AP geolocation statistics (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessGeolocationOperApGeoLocStats represents the AP geolocation statistics.
type CiscoIOSXEWirelessGeolocationOperApGeoLocStats struct {
	ApGeoLocStats *ApGeoLocStats `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats"`
}

// ApGeoLocStats represents AP geolocation statistics.
type ApGeoLocStats struct {
	NumApGnss               int    `json:"num-ap-gnss"`               // Number of APs with GNSS module (Live: IOS-XE 17.12.6a)
	NumApManHeight          int    `json:"num-ap-man-height"`         // Number of APs with manual elevation config (Live: IOS-XE 17.12.6a)
	NumApDerived            int    `json:"num-ap-derived"`            // Number of APs with derived geolocation info (Live: IOS-XE 17.12.6a)
	LastDerivationTimestamp string `json:"last-derivation-timestamp"` // Last derivation algorithm run timestamp (Live: IOS-XE 17.12.6a)
}

// ApGeoLocData represents wireless AP geolocation data.
type ApGeoLocData struct {
	ApMAC     string           `json:"ap-mac"`              // AP MAC address (YANG: IOS-XE 17.12.1)
	Loc       *GeoLocInfo      `json:"loc,omitempty"`       // AP geolocation information (YANG: IOS-XE 17.12.1)
	Elevation *GeoLocElevation `json:"elevation,omitempty"` // AP elevation information (YANG: IOS-XE 17.12.1)
}

// GeoLocInfo represents wireless AP geolocation information.
type GeoLocInfo struct {
	Source               *string        `json:"source,omitempty"`                // AP geolocation source (YANG: IOS-XE 17.12.1)
	AreaOfUncertainty    *uint32        `json:"area-of-uncertainty,omitempty"`   // AP geolocation area of uncertainty in square meters (YANG: IOS-XE 17.12.1)
	HDOP                 *string        `json:"hdop,omitempty"`                  // AP GPS Horizontal Dilution of Precision (YANG: IOS-XE 17.12.1)
	LastRcvdTimestamp    *string        `json:"last-rcvd-timestamp,omitempty"`   // Last received GPS coordinate timestamp (YANG: IOS-XE 17.12.1)
	AnchorAP             *string        `json:"anchor-ap,omitempty"`             // MAC address of anchor AP (YANG: IOS-XE 17.12.1)
	SourceDerivedGeoloc  *string        `json:"source-derived-geoloc,omitempty"` // Source of derived geolocation (YANG: IOS-XE 17.12.1)
	DerivationTechniques *string        `json:"derivation-techniques,omitempty"` // Techniques used for deriving geolocation (YANG: IOS-XE 17.12.1)
	DerivationDistance   *uint16        `json:"derivation-distance,omitempty"`   // Distance to Anchor AP in meters (Live: IOS-XE 17.12.7a)
	Ellipse              *GeoLocEllipse `json:"ellipse,omitempty"`               // AP geolocation in ellipse format (YANG: IOS-XE 17.12.1)
	Invalid              *bool          `json:"invalid,omitempty"`               // AP invalid geolocation (YANG: IOS-XE 17.12.1)
}

// GeoLocElevation represents wireless AP elevation information.
type GeoLocElevation struct {
	LastRcvdTimestamp *string              `json:"last-rcvd-timestamp,omitempty"` // Last received height timestamp (YANG: IOS-XE 17.12.1)
	AGLData           *GeoLocElevationData `json:"agl-data,omitempty"`            // Above Ground Level elevation information (YANG: IOS-XE 17.12.1)
	MSLData           *GeoLocElevationData `json:"msl-data,omitempty"`            // Mean Sea Level elevation information (YANG: IOS-XE 17.12.1)
	Invalid           *bool                `json:"invalid,omitempty"`             // Invalid elevation information (YANG: IOS-XE 17.12.1)
}

// GeoLocElevationData represents elevation data details.
type GeoLocElevationData struct {
	Uncertainty *uint16 `json:"uncertainty,omitempty"` // Elevation uncertainty in meters (YANG: IOS-XE 17.12.1)
	Height      *int16  `json:"height,omitempty"`      // AP height in meters (YANG: IOS-XE 17.12.1)
	Source      *string `json:"source,omitempty"`      // Elevation source (YANG: IOS-XE 17.12.1)
}

// GeoLocEllipse represents wireless AP geolocation ellipse representation.
type GeoLocEllipse struct {
	Center      *GeoLocPoint `json:"center,omitempty"`      // Center of the ellipse (YANG: IOS-XE 17.12.1)
	MajorAxis   *uint16      `json:"major-axis,omitempty"`  // Major axis of the ellipse in meters (YANG: IOS-XE 17.12.1)
	MinorAxis   *uint16      `json:"minor-axis,omitempty"`  // Minor axis of the ellipse in meters (YANG: IOS-XE 17.12.1)
	Orientation *string      `json:"orientation,omitempty"` // Orientation clockwise from True North (YANG: IOS-XE 17.12.1)
}

// GeoLocPoint represents wireless AP geolocation point.
type GeoLocPoint struct {
	Longitude *string `json:"longitude,omitempty"` // Geolocation longitude in WGS 84 format (YANG: IOS-XE 17.12.1)
	Latitude  *string `json:"latitude,omitempty"`  // Geolocation latitude in WGS 84 format (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessGeolocationOperApGeoLocData wraps the ApGeoLocData structure of the geolocation operational data.
type CiscoIOSXEWirelessGeolocationOperApGeoLocData struct {
	ApGeoLocData []ApGeoLocData `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-data"`
}
