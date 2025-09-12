package geolocation

// GeolocationOper represents the geolocation operational data.
type GeolocationOper struct {
	CiscoIOSXEWirelessGeolocationOperGeolocationOperData struct {
		// AP geolocation data list (YANG: IOS-XE 17.12.1+)
		ApGeoLocData []ApGeoLocData `json:"ap-geo-loc-data,omitempty"`
		// AP geolocation statistics
		ApGeoLocStats *ApGeoLocStats `json:"ap-geo-loc-stats,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-geolocation-oper:geolocation-oper-data"`
}

// GeolocationOperApGeoLocStats represents the AP geolocation statistics.
type GeolocationOperApGeoLocStats struct {
	ApGeoLocStats ApGeoLocStats `json:"Cisco-IOS-XE-wireless-geolocation-oper:ap-geo-loc-stats"`
}

// ApGeoLocStats represents AP geolocation statistics.
type ApGeoLocStats struct {
	// Number of APs with GNSS module
	NumApGnss int `json:"num-ap-gnss"`
	// Number of APs with manual elevation configuration
	NumApManHeight int `json:"num-ap-man-height"`
	// Number of APs with derived geolocation information
	NumApDerived int `json:"num-ap-derived"`
	// Last run timestamp of the geolocation derivation algorithm
	LastDerivationTimestamp string `json:"last-derivation-timestamp"`
}

// ApGeoLocData represents wireless AP geolocation data.
type ApGeoLocData struct {
	// AP MAC address (YANG: IOS-XE 17.12.1+)
	ApMac string `json:"ap-mac"`
	// AP geolocation information (YANG: IOS-XE 17.12.1+)
	Loc *GeoLocInfo `json:"loc,omitempty"`
	// AP elevation information (YANG: IOS-XE 17.12.1+)
	Elevation *GeoLocElevation `json:"elevation,omitempty"`
}

// GeoLocInfo represents wireless AP geolocation information.
type GeoLocInfo struct {
	// AP geolocation source (YANG: IOS-XE 17.12.1+)
	Source *string `json:"source,omitempty"`
	// AP geolocation area of uncertainty in square meters (YANG: IOS-XE 17.12.1+)
	AreaOfUncertainty *int `json:"area-of-uncertainty,omitempty"`
	// AP geolocation GPS Horizontal Dilution of Precision (YANG: IOS-XE 17.12.1+)
	HDOP *float64 `json:"hdop,omitempty"`
	// Last received GPS coordinate timestamp (YANG: IOS-XE 17.12.1+)
	LastRcvdTimestamp *string `json:"last-rcvd-timestamp,omitempty"`
	// MAC address of anchor AP (YANG: IOS-XE 17.12.1+)
	AnchorAP *string `json:"anchor-ap,omitempty"`
	// Source of derived geolocation (YANG: IOS-XE 17.12.1+)
	SourceDerivedGeoloc *string `json:"source-derived-geoloc,omitempty"`
	// Techniques used for deriving geolocation, LLDP support added (YANG: IOS-XE 17.18.1+)
	DerivationTechniques *string `json:"derivation-techniques,omitempty"`
	// Distance to Anchor AP in meters (YANG: IOS-XE 17.12.1+)
	DerivationDistance *int `json:"derivation-distance,omitempty"`
	// AP geolocation in ellipse format (YANG: IOS-XE 17.12.1+)
	Ellipse *GeoLocEllipse `json:"ellipse,omitempty"`
	// AP invalid geolocation (YANG: IOS-XE 17.12.1+)
	Invalid *bool `json:"invalid,omitempty"`
}

// GeoLocElevation represents wireless AP elevation information.
type GeoLocElevation struct {
	// Last received height timestamp (YANG: IOS-XE 17.12.1+)
	LastRcvdTimestamp *string `json:"last-rcvd-timestamp,omitempty"`
	// Above Ground Level elevation information (YANG: IOS-XE 17.12.1+)
	AGLData *GeoLocElevationData `json:"agl-data,omitempty"`
	// Mean Sea Level elevation information (YANG: IOS-XE 17.12.1+)
	MSLData *GeoLocElevationData `json:"msl-data,omitempty"`
	// Invalid elevation information (YANG: IOS-XE 17.12.1+)
	Invalid *bool `json:"invalid,omitempty"`
}

// GeoLocElevationData represents elevation data details.
type GeoLocElevationData struct {
	// Elevation uncertainty in meters (YANG: IOS-XE 17.12.1+)
	Uncertainty *int `json:"uncertainty,omitempty"`
	// AP height in meters (YANG: IOS-XE 17.12.1+)
	Height *int `json:"height,omitempty"`
	// Elevation source (YANG: IOS-XE 17.12.1+)
	Source *string `json:"source,omitempty"`
}

// GeoLocEllipse represents wireless AP geolocation ellipse representation.
type GeoLocEllipse struct {
	// Center of the ellipse (YANG: IOS-XE 17.12.1+)
	Center *GeoLocPoint `json:"center,omitempty"`
	// Major axis of the ellipse in meters (YANG: IOS-XE 17.12.1+)
	MajorAxis *int `json:"major-axis,omitempty"`
	// Minor axis of the ellipse in meters (YANG: IOS-XE 17.12.1+)
	MinorAxis *int `json:"minor-axis,omitempty"`
	// Orientation of the ellipse measured clockwise from True North in degrees (YANG: IOS-XE 17.12.1+)
	Orientation *float64 `json:"orientation,omitempty"`
}

// GeoLocPoint represents wireless AP geolocation point.
type GeoLocPoint struct {
	// Geolocation longitude in WGS 84 format in degrees (YANG: IOS-XE 17.12.1+)
	Longitude *float64 `json:"longitude,omitempty"`
	// Geolocation latitude in WGS 84 format in degrees (YANG: IOS-XE 17.12.1+)
	Latitude *float64 `json:"latitude,omitempty"`
}
