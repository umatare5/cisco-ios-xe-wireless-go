package fabric

// CiscoIOSXEWirelessFabricCfg represents fabric configuration data container.
type CiscoIOSXEWirelessFabricCfg struct {
	CiscoIOSXEWirelessFabricCfgData struct {
		Fabric                  *FabricConfig        `json:"fabric"`                    // Fabric specific configuration (Live: IOS-XE 17.12.5)
		FabricProfiles          *FabricProfiles      `json:"fabric-profiles,omitempty"` // Fabric profiles configuration (YANG: IOS-XE 17.12.1)
		FabricControlplaneNames *FabricControlplanes `json:"fabric-controlplane-names"` // Fabric control plane name configuration (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"` // Fabric configuration data (Live: IOS-XE 17.12.5)
}

// CiscoIOSXEWirelessFabricCfgFabric represents fabric configuration container.
type CiscoIOSXEWirelessFabricCfgFabric struct {
	Fabric *FabricConfig `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric,omitempty"`
}

// CiscoIOSXEWirelessFabricCfgFabricProfiles represents fabric profiles container.
type CiscoIOSXEWirelessFabricCfgFabricProfiles struct {
	FabricProfiles *FabricProfiles `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-profiles,omitempty"`
}

// CiscoIOSXEWirelessFabricCfgFabricControlplaneNames represents fabric control plane names container.
type CiscoIOSXEWirelessFabricCfgFabricControlplaneNames struct {
	FabricControlplaneNames *FabricControlplanes `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names,omitempty"`
}

// FabricConfig represents fabric global configuration.
type FabricConfig struct {
	FabricEnabled         bool                   `json:"fabric-enabled,omitempty"`           // Global fabric enable status (YANG: IOS-XE 17.12.1)
	FabricNameVNIDEntries *FabricNameVNIDEntries `json:"fabric-name-vnid-entries,omitempty"` // Global VNID name configuration (YANG: IOS-XE 17.12.1)
}

// FabricProfiles represents fabric profiles container.
type FabricProfiles struct {
	FabricProfile []FabricProfile `json:"fabric-profile,omitempty"` // List of fabric profiles (YANG: IOS-XE 17.12.1)
}

// FabricProfile represents individual fabric profile configuration.
type FabricProfile struct {
	FabricProfileName string  `json:"fabric-profile-name"`      // Fabric profile name (YANG: IOS-XE 17.12.1)
	Description       *string `json:"description,omitempty"`    // Fabric profile description (YANG: IOS-XE 17.12.1)
	EncapType         *string `json:"encap-type,omitempty"`     // Encapsulation type used in profile (YANG: IOS-XE 17.12.1)
	SGTTag            *int    `json:"sgt-tag,omitempty"`        // Group tag associated with fabric profile (YANG: IOS-XE 17.12.1)
	ClientL2VNID      *int    `json:"client-l2-vnid,omitempty"` // Client virtual network ID (YANG: IOS-XE 17.12.1)
}

// FabricControlplanes represents fabric control planes container.
type FabricControlplanes struct {
	FabricControlplaneName []FabricControlplaneName `json:"fabric-controlplane-name,omitempty"` // List of fabric control plane configurations (Live: IOS-XE 17.12.5)
}

// FabricControlplaneName represents individual fabric control plane configuration.
type FabricControlplaneName struct {
	ControlPlaneName            string                       `json:"control-plane-name"`                        // Fabric control plane name (Live: IOS-XE 17.12.5)
	Description                 string                       `json:"description,omitempty"`                     // Fabric control plane description (Live: IOS-XE 17.12.5)
	FabricControlPlaneIPConfigs *FabricControlPlaneIPConfigs `json:"fabric-control-plane-ip-configs,omitempty"` // Fabric control plane configuration (YANG: IOS-XE 17.12.1)
}

// FabricControlPlaneIPConfigs represents fabric control plane IP configurations container.
type FabricControlPlaneIPConfigs struct {
	FabricControlPlaneIPConfig []FabricControlPlaneIPConfig `json:"fabric-control-plane-ip-config,omitempty"` // List of fabric control plane configurations (YANG: IOS-XE 17.12.1)
}

// FabricControlPlaneIPConfig represents individual fabric control plane IP configuration.
type FabricControlPlaneIPConfig struct {
	ControlPlaneIP string `json:"control-plane-ip"`       // IP address of the control plane (YANG: IOS-XE 17.12.1)
	PSKKey         string `json:"psk-key,omitempty"`      // PSK associated with control plane entity (YANG: IOS-XE 17.12.1)
	PSKKeyType     string `json:"psk-key-type,omitempty"` // PSK type for the control plane (YANG: IOS-XE 17.12.1)
}

// FabricNameVNIDEntries represents fabric name VNID entries container.
type FabricNameVNIDEntries struct {
	FabricNameVNIDEntry []FabricNameVNIDEntry `json:"fabric-name-vnid-entry,omitempty"` // Global VNID name configuration (YANG: IOS-XE 17.12.1)
}

// FabricNameVNIDEntry represents individual fabric name VNID mapping.
type FabricNameVNIDEntry struct {
	Name             string  `json:"name"`                         // Fabric name for AP join and VNID override (YANG: IOS-XE 17.12.1)
	L2VNID           int     `json:"l2-vnid"`                      // VNID for client subnet (YANG: IOS-XE 17.12.1)
	L3VNID           *int    `json:"l3-vnid,omitempty"`            // VNID for AP subnet (YANG: IOS-XE 17.12.1)
	Netmask          *string `json:"netmask,omitempty"`            // Network mask of AP subnet (YANG: IOS-XE 17.12.1)
	NetworkIP        *string `json:"network-ip,omitempty"`         // IP address of AP subnet (YANG: IOS-XE 17.12.1)
	ControlPlaneName *string `json:"control-plane-name,omitempty"` // Fabric control plane name (YANG: IOS-XE 17.12.1)
}
