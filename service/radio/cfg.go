package radio

// CiscoIOSXEWirelessRadioCfg represents the root container for radio configuration data.
type CiscoIOSXEWirelessRadioCfg struct {
	CiscoIOSXEWirelessRadioCfgData struct {
		RadioProfiles RadioProfiles `json:"radio-profiles"` // Radio profiles configuration container (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"` // Radio configuration data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessRadioCfgRadioProfiles represents the radio profiles configuration wrapper structure.
type CiscoIOSXEWirelessRadioCfgRadioProfiles struct {
	RadioProfiles RadioProfiles `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles"`
}

// RadioProfiles represents radio profiles container.
type RadioProfiles struct {
	RadioProfile []RadioProfile `json:"radio-profile"` // List of radio profile configurations (Live: IOS-XE 17.12.6a)
}

// RadioProfile represents individual radio profile configuration.
type RadioProfile struct {
	Name                   string `json:"name"`                               // Name of the radio profile (Live: IOS-XE 17.12.6a)
	Desc                   string `json:"desc,omitempty"`                     // Description for the radio profile (Live: IOS-XE 17.12.6a)
	MeshBackhaul           bool   `json:"mesh-backhaul"`                      // Enable mesh backhaul on this radio (Live: IOS-XE 17.12.6a)
	BeamSteerMode          string `json:"beam-steer-mode,omitempty"`          // Beam steering mode for the AP slot (YANG: IOS-XE 17.12.1)
	NumAntEnabled          *uint8 `json:"num-ant-enabled,omitempty"`          // Number of antennas to be enabled for AP slot (YANG: IOS-XE 17.12.1)
	MeshDesignatedDownlink *bool  `json:"mesh-designated-downlink,omitempty"` // Use radio as designated mesh downlink backhaul (YANG: IOS-XE 17.12.1)
	DTIMPeriod             *uint8 `json:"dtim-period,omitempty"`              // DTIM interval for 6GHz radio (YANG: IOS-XE 17.12.1)
}
