package afc

// AfcOper represents AFC operational data container.
type AfcOper struct {
	CiscoIOSXEWirelessAfcOperAfcOperData struct {
		EwlcAfcApResp []EwlcAfcApResp `json:"ewlc-afc-ap-resp"`          // AFC response list (Live: IOS-XE 17.12.5)
		EwlcAfcApReq  []EwlcAfcApReq  `json:"ewlc-afc-ap-req,omitempty"` // AFC request list (YANG: IOS-XE 17.18.1)
	} `json:"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data"` // AFC operational data (Live: IOS-XE 17.12.5)
}

// AfcOperEwlcAfcApResp represents AFC AP response data container.
type AfcOperEwlcAfcApResp struct {
	EwlcAfcApResp []EwlcAfcApResp `json:"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp"`
}

// EwlcAfcApReq represents AFC request information from access point.
type EwlcAfcApReq struct {
	ApMac           string          `json:"ap-mac"`                      // Access point MAC address (YANG: IOS-XE 17.18.1)
	ReqData         *AfcRequestData `json:"req-data,omitempty"`          // AFC request data structure (YANG: IOS-XE 17.18.1)
	RequestStatus   string          `json:"request-status,omitempty"`    // AFC request status (YANG: IOS-XE 17.18.1)
	RequestStatusTS string          `json:"request-status-ts,omitempty"` // Request status timestamp (YANG: IOS-XE 17.18.1)
	ReqIDSent       *uint64         `json:"req-id-sent,omitempty"`       // Sent request ID (YANG: IOS-XE 17.18.1)
}

// AfcRequestData represents AFC request data structure.
type AfcRequestData struct {
	Device          *AfcDeviceDescriptor `json:"device,omitempty"`            // AFC device descriptor (YANG: IOS-XE 17.18.1)
	Location        *AfcLocation         `json:"location,omitempty"`          // AFC device location information (YANG: IOS-XE 17.18.1)
	Band20          *AfcBandRequest      `json:"band20,omitempty"`            // 20MHz band request (YANG: IOS-XE 17.18.1)
	Band40          *AfcBandRequest      `json:"band40,omitempty"`            // 40MHz band request (YANG: IOS-XE 17.18.1)
	Band80          *AfcBandRequest      `json:"band80,omitempty"`            // 80MHz band request (YANG: IOS-XE 17.18.1)
	Band160         *AfcBandRequest      `json:"band160,omitempty"`           // 160MHz band request (YANG: IOS-XE 17.18.1)
	Band80Plus      *AfcBandRequest      `json:"band80plus,omitempty"`        // 80+ MHz band request (YANG: IOS-XE 17.18.1)
	MinDesiredPower *float64             `json:"min-desired-power,omitempty"` // Minimum desired power level (YANG: IOS-XE 17.18.1)

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1)
	Band320 *AfcBandRequest `json:"band320,omitempty"` // 320MHz band request (YANG: IOS-XE 17.18.1)
}

// AfcDeviceDescriptor represents AFC device descriptor information.
type AfcDeviceDescriptor struct {
	SerialNumber string      `json:"serial-number,omitempty"` // Device serial number (YANG: IOS-XE 17.18.1)
	CertID       []AfcCertID `json:"cert-id,omitempty"`       // Certification ID list (YANG: IOS-XE 17.18.1)
}

// AfcCertID represents AFC certification identifier.
type AfcCertID struct {
	ID        string `json:"id,omitempty"`         // Certification ID (YANG: IOS-XE 17.18.1)
	RulesetID string `json:"ruleset-id,omitempty"` // Ruleset identifier (YANG: IOS-XE 17.18.1)
}

// AfcLocation represents AFC device location information.
type AfcLocation struct {
	LocType           string            `json:"loc-type,omitempty"`            // Location type (YANG: IOS-XE 17.18.1)
	Ellipse           *AfcEllipse       `json:"ellipse,omitempty"`             // Ellipse location data (YANG: IOS-XE 17.18.1)
	LinearPol         *AfcLinearPolygon `json:"linear-pol,omitempty"`          // Linear polygon location data (YANG: IOS-XE 17.18.1)
	Elevation         *AfcElevation     `json:"elevation,omitempty"`           // Device elevation data (YANG: IOS-XE 17.18.1)
	Deployment        string            `json:"deployment,omitempty"`          // Deployment type (YANG: IOS-XE 17.18.1)
	AreaOfUncertainty *uint32           `json:"area-of-uncertainty,omitempty"` // Area of uncertainty in meters (YANG: IOS-XE 17.18.1)
}

// AfcEllipse represents device ellipse location coordinates.
type AfcEllipse struct {
	Center      *AfcPoint `json:"center,omitempty"`      // Ellipse center point (YANG: IOS-XE 17.18.1)
	MajorAxis   *uint16   `json:"major-axis,omitempty"`  // Major axis length (YANG: IOS-XE 17.18.1)
	MinorAxis   *uint16   `json:"minor-axis,omitempty"`  // Minor axis length (YANG: IOS-XE 17.18.1)
	Orientation *float64  `json:"orientation,omitempty"` // Ellipse orientation angle (YANG: IOS-XE 17.18.1)
}

// AfcLinearPolygon represents device linear polygon location.
type AfcLinearPolygon struct {
	Points []AfcPoint `json:"points,omitempty"` // Polygon boundary points (YANG: IOS-XE 17.18.1)
}

// AfcPoint represents AFC geographic point coordinates.
type AfcPoint struct {
	Longitude *float64 `json:"longitude,omitempty"` // Longitude coordinate (YANG: IOS-XE 17.18.1)
	Latitude  *float64 `json:"latitude,omitempty"`  // Latitude coordinate (YANG: IOS-XE 17.18.1)
}

// AfcElevation represents device elevation information.
type AfcElevation struct {
	Height      *int16  `json:"height,omitempty"`      // Height in meters (YANG: IOS-XE 17.18.1)
	HeightType  string  `json:"height-type,omitempty"` // Height measurement type (YANG: IOS-XE 17.18.1)
	Uncertainty *uint16 `json:"uncertainty,omitempty"` // Height uncertainty in meters (YANG: IOS-XE 17.18.1)
}

// AfcBandRequest represents AFC frequency band request parameters.
type AfcBandRequest struct {
	GlobalOperClass *uint16  `json:"global-oper-class,omitempty"` // Global operating class (YANG: IOS-XE 17.18.1)
	ChannelCFI      []uint16 `json:"channel-cfi,omitempty"`       // Channel center frequency indices (YANG: IOS-XE 17.18.1)
	Enabled         *bool    `json:"enabled,omitempty"`           // Band request enabled status (YANG: IOS-XE 17.18.1)
}

// AfcChannelResponse represents AFC channel response information.
type AfcChannelResponse struct {
	AvailChannelCFI *uint16  `json:"avail-channel-cfi,omitempty"` // Available channel center frequency index (YANG: IOS-XE 17.18.1)
	MaxEIRP         *float64 `json:"max-eirp,omitempty"`          // Maximum effective isotropic radiated power (YANG: IOS-XE 17.18.1)
}

// EwlcAfcApResp represents AFC response from access point.
type EwlcAfcApResp struct {
	ApMac    string `json:"ap-mac"` // Access point MAC address (Live: IOS-XE 17.12.5)
	RespData struct {
		RequestID string `json:"request-id"` // AFC request identifier (Live: IOS-XE 17.12.5)
		RulesetID string `json:"ruleset-id"` // AFC ruleset identifier (YANG: IOS-XE 17.18.1)
		RespCode  struct {
			Code             int    `json:"code"`              // AFC response code (Live: IOS-XE 17.12.5)
			Description      string `json:"description"`       // Response code description (Live: IOS-XE 17.12.5)
			SupplementalInfo string `json:"supplemental-info"` // Additional response information (Live: IOS-XE 17.12.5)
		} `json:"resp-code"` // AFC response code details (Live: IOS-XE 17.12.5)
		Band20 struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 20MHz global operating class (Live: IOS-XE 17.12.5)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 20MHz channels (Live: IOS-XE 17.12.5)
		} `json:"band20"` // 20MHz band response (Live: IOS-XE 17.12.5)
		Band40 struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 40MHz global operating class (Live: IOS-XE 17.12.5)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 40MHz channels (Live: IOS-XE 17.12.5)
		} `json:"band40"` // 40MHz band response (Live: IOS-XE 17.12.5)
		Band80 struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 80MHz global operating class (Live: IOS-XE 17.12.5)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 80MHz channels (Live: IOS-XE 17.12.5)
		} `json:"band80"` // 80MHz band response (Live: IOS-XE 17.12.5)
		Band160 struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 160MHz global operating class (Live: IOS-XE 17.12.5)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 160MHz channels (Live: IOS-XE 17.12.5)
		} `json:"band160"` // 160MHz band response (Live: IOS-XE 17.12.5)
		Band80Plus struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 80+ MHz global operating class (Live: IOS-XE 17.12.5)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 80+ MHz channels (Live: IOS-XE 17.12.5)
		} `json:"band80plus"` // 80+ MHz band response (Live: IOS-XE 17.12.5)
		Band320 *struct {
			GlobalOperClass int                  `json:"global-oper-class"`  // 320MHz global operating class (YANG: IOS-XE 17.18.1)
			Channels        []AfcChannelResponse `json:"channels,omitempty"` // Available 320MHz channels (YANG: IOS-XE 17.18.1)
		} `json:"band320,omitempty"` // 320MHz band response (YANG: IOS-XE 17.18.1)
		ExpireTime        string `json:"expire-time"`         // Response expiration time (Live: IOS-XE 17.12.5)
		RespRcvdTimestamp string `json:"resp-rcvd-timestamp"` // Response received timestamp (Live: IOS-XE 17.12.5)
	} `json:"resp-data"` // AFC response data (Live: IOS-XE 17.12.5)
	Slot int `json:"slot"` // Access point slot number (Live: IOS-XE 17.12.5)
}
