package geolocation

// GeolocationOper represents the geolocation operational data.
type GeolocationOper struct {
	CiscoIOSXEWirelessGeolocationOperGeolocationOperData struct {
		ApGeoLocData  []ApGeoLocData `json:"ap-geo-loc-data,omitempty"` // AP geolocation data list (YANG: IOS-XE 17.12.1)
		ApGeoLocStats *ApGeoLocStats `json:"ap-geo-loc-stats"`          // AP geolocation statistics (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"` // Geolocation operational data (Live: IOS-XE 17.12.5)
}

// GeolocationOperApGeoLocStats represents the AP geolocation statistics.
type GeolocationOperApGeoLocStats struct {
	ApGeoLocStats ApGeoLocStats `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats"`
}

// ApGeoLocStats represents AP geolocation statistics.
type ApGeoLocStats struct {
	NumApGnss               int    `json:"num-ap-gnss"`               // Number of APs with GNSS module (Live: IOS-XE 17.12.5)
	NumApManHeight          int    `json:"num-ap-man-height"`         // Number of APs with manual elevation config (Live: IOS-XE 17.12.5)
	NumApDerived            int    `json:"num-ap-derived"`            // Number of APs with derived geolocation info (Live: IOS-XE 17.12.5)
	LastDerivationTimestamp string `json:"last-derivation-timestamp"` // Last derivation algorithm run timestamp (Live: IOS-XE 17.12.5)
}

// ApGeoLocData represents wireless AP geolocation data.
type ApGeoLocData struct {
	ApMac     string           `json:"ap-mac"`              // AP MAC address (YANG: IOS-XE 17.12.1)
	Loc       *GeoLocInfo      `json:"loc,omitempty"`       // AP geolocation information (YANG: IOS-XE 17.12.1)
	Elevation *GeoLocElevation `json:"elevation,omitempty"` // AP elevation information (YANG: IOS-XE 17.12.1)
}

// GeoLocInfo represents wireless AP geolocation information.
type GeoLocInfo struct {
	Source               *string        `json:"source,omitempty"`                // AP geolocation source (YANG: IOS-XE 17.12.1)
	AreaOfUncertainty    *int           `json:"area-of-uncertainty,omitempty"`   // AP geolocation area of uncertainty (YANG: IOS-XE 17.12.1)
	HDOP                 *float64       `json:"hdop,omitempty"`                  // AP GPS Horizontal Dilution of Precision (YANG: IOS-XE 17.12.1)
	LastRcvdTimestamp    *string        `json:"last-rcvd-timestamp,omitempty"`   // Last received GPS coordinate timestamp (YANG: IOS-XE 17.12.1)
	AnchorAP             *string        `json:"anchor-ap,omitempty"`             // MAC address of anchor AP (YANG: IOS-XE 17.12.1)
	SourceDerivedGeoloc  *string        `json:"source-derived-geoloc,omitempty"` // Source of derived geolocation (YANG: IOS-XE 17.12.1)
	DerivationTechniques *string        `json:"derivation-techniques,omitempty"` // Techniques used for deriving geolocation (YANG: IOS-XE 17.18.1)
	DerivationDistance   *int           `json:"derivation-distance,omitempty"`   // Distance to Anchor AP in meters (YANG: IOS-XE 17.12.1)
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
	Uncertainty *int    `json:"uncertainty,omitempty"` // Elevation uncertainty in meters (YANG: IOS-XE 17.12.1)
	Height      *int    `json:"height,omitempty"`      // AP height in meters (YANG: IOS-XE 17.12.1)
	Source      *string `json:"source,omitempty"`      // Elevation source (YANG: IOS-XE 17.12.1)
}

// GeoLocEllipse represents wireless AP geolocation ellipse representation.
type GeoLocEllipse struct {
	Center      *GeoLocPoint `json:"center,omitempty"`      // Center of the ellipse (YANG: IOS-XE 17.12.1)
	MajorAxis   *int         `json:"major-axis,omitempty"`  // Major axis of the ellipse in meters (YANG: IOS-XE 17.12.1)
	MinorAxis   *int         `json:"minor-axis,omitempty"`  // Minor axis of the ellipse in meters (YANG: IOS-XE 17.12.1)
	Orientation *float64     `json:"orientation,omitempty"` // Orientation clockwise from True North (YANG: IOS-XE 17.12.1)
}

// GeoLocPoint represents wireless AP geolocation point.
type GeoLocPoint struct {
	Longitude *float64 `json:"longitude,omitempty"` // Geolocation longitude in WGS 84 format (YANG: IOS-XE 17.12.1)
	Latitude  *float64 `json:"latitude,omitempty"`  // Geolocation latitude in WGS 84 format (YANG: IOS-XE 17.12.1)
}
