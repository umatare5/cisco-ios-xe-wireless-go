// Package fabric provides data models for fabric configuration data.
package fabric

// FabricCfg represents fabric configuration data container.
type FabricCfg struct {
	FabricCfgData FabricCfgData `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"`
}

// FabricCfgFabric represents fabric configuration container.
type FabricCfgFabric struct {
	Fabric *FabricConfig `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric,omitempty"`
}

// FabricCfgFabricProfiles represents fabric profiles container.
type FabricCfgFabricProfiles struct {
	FabricProfiles *FabricProfiles `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-profiles,omitempty"`
}

// FabricCfgFabricControlplaneNames represents fabric control plane names container.
type FabricCfgFabricControlplaneNames struct {
	FabricControlplaneNames *FabricControlplanes `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names,omitempty"`
}

// FabricCfgData represents fabric configuration data container.
type FabricCfgData struct {
	Fabric                  *FabricConfig        `json:"fabric,omitempty"`
	FabricProfiles          *FabricProfiles      `json:"fabric-profiles,omitempty"`
	FabricControlplaneNames *FabricControlplanes `json:"fabric-controlplane-names,omitempty"`
}

// FabricConfig represents fabric global configuration.
type FabricConfig struct {
	FabricEnabled         bool                   `json:"fabric-enabled,omitempty"`           // Fabric feature enable status (YANG: IOS-XE 17.12.1+)
	FabricNameVNIDEntries *FabricNameVNIDEntries `json:"fabric-name-vnid-entries,omitempty"` // Fabric name to VNID mappings (YANG: IOS-XE 17.12.1+)
}

// FabricProfiles represents fabric profiles container.
type FabricProfiles struct {
	FabricProfile []FabricProfile `json:"fabric-profile,omitempty"`
}

// FabricProfile represents individual fabric profile configuration.
type FabricProfile struct {
	FabricProfileName string  `json:"fabric-profile-name"`      // Fabric profile name identifier (YANG: IOS-XE 17.12.1+)
	Description       *string `json:"description,omitempty"`    // Fabric profile description (YANG: IOS-XE 17.12.1+)
	EncapType         *string `json:"encap-type,omitempty"`     // Encapsulation type (YANG: IOS-XE 17.12.1+)
	SGTTag            *int    `json:"sgt-tag,omitempty"`        // Security Group Tag (YANG: IOS-XE 17.12.1+)
	ClientL2VNID      *int    `json:"client-l2-vnid,omitempty"` // Client Layer 2 VNID (YANG: IOS-XE 17.12.1+)
}

// FabricControlplanes represents fabric control planes container.
type FabricControlplanes struct {
	FabricControlplaneName []FabricControlplaneName `json:"fabric-controlplane-name,omitempty"`
}

// FabricControlplaneName represents individual fabric control plane configuration.
type FabricControlplaneName struct {
	ControlPlaneName            string                       `json:"control-plane-name"`                        // Fabric control plane name
	Description                 string                       `json:"description,omitempty"`                     // Control plane description
	FabricControlPlaneIPConfigs *FabricControlPlaneIPConfigs `json:"fabric-control-plane-ip-configs,omitempty"` // Control plane IP configurations (YANG: IOS-XE 17.12.1+)
}

// FabricControlPlaneIPConfigs represents fabric control plane IP configurations container.
type FabricControlPlaneIPConfigs struct {
	FabricControlPlaneIPConfig []FabricControlPlaneIPConfig `json:"fabric-control-plane-ip-config,omitempty"`
}

// FabricControlPlaneIPConfig represents individual fabric control plane IP configuration.
type FabricControlPlaneIPConfig struct {
	ControlPlaneIP string `json:"control-plane-ip"`       // Control plane IP address (YANG: IOS-XE 17.12.1+)
	PSKKey         string `json:"psk-key,omitempty"`      // Pre-shared key (YANG: IOS-XE 17.12.1+)
	PSKKeyType     string `json:"psk-key-type,omitempty"` // Pre-shared key type (YANG: IOS-XE 17.12.1+)
}

// FabricNameVNIDEntries represents fabric name VNID entries container.
type FabricNameVNIDEntries struct {
	FabricNameVNIDEntry []FabricNameVNIDEntry `json:"fabric-name-vnid-entry,omitempty"`
}

// FabricNameVNIDEntry represents individual fabric name VNID mapping.
type FabricNameVNIDEntry struct {
	Name             string  `json:"name"`                         // Fabric name identifier (YANG: IOS-XE 17.12.1+)
	L2VNID           int     `json:"l2-vnid"`                      // Layer 2 VNID (YANG: IOS-XE 17.12.1+)
	L3VNID           *int    `json:"l3-vnid,omitempty"`            // Layer 3 VNID (YANG: IOS-XE 17.12.1+)
	Netmask          *string `json:"netmask,omitempty"`            // Network mask of AP subnet (YANG: IOS-XE 17.12.1+)
	NetworkIP        *string `json:"network-ip,omitempty"`         // IP address of AP subnet (YANG: IOS-XE 17.12.1+)
	ControlPlaneName *string `json:"control-plane-name,omitempty"` // Fabric control plane name (YANG: IOS-XE 17.12.1+)
}
