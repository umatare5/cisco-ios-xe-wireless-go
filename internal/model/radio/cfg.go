// Package radio provides data models for radio configuration data.
package radio

// RadioCfg represents the root container for radio configuration data.
type RadioCfg struct {
	RadioCfgData RadioCfgData `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"`
}

// RadioCfgData represents radio configuration data container.
type RadioCfgData struct {
	RadioProfiles RadioProfiles `json:"radio-profiles"` // Radio profiles configuration container
}

// RadioProfiles represents radio profiles container.
type RadioProfiles struct {
	RadioProfile []RadioProfile `json:"radio-profile"` // List of radio profile configurations
}

// RadioProfile represents individual radio profile configuration.
type RadioProfile struct {
	Name                   string `json:"name"`                               // Radio profile name identifier
	Desc                   string `json:"desc,omitempty"`                     // Radio profile description
	MeshBackhaul           bool   `json:"mesh-backhaul"`                      // Mesh backhaul enable status
	BeamSteerMode          string `json:"beam-steer-mode,omitempty"`          // Beam steering mode for AP slot (YANG: IOS-XE 17.12.1+)
	NumAntEnabled          *uint8 `json:"num-ant-enabled,omitempty"`          // Number of antennas enabled for AP slot (YANG: IOS-XE 17.12.1+)
	MeshDesignatedDownlink *bool  `json:"mesh-designated-downlink,omitempty"` // Designated mesh downlink backhaul setting (YANG: IOS-XE 17.12.1+)
	DTIMPeriod             *uint8 `json:"dtim-period,omitempty"`              // DTIM interval for 6GHz radio (YANG: IOS-XE 17.12.1+)
}
