// Package model provides data models for fabric configuration data.
package model

// FabricCfg  represents the Fabric configuration data.
type FabricCfg struct {
	FabricCfgData FabricCfgData `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"`
}

// FabricCfgFabric  represents the fabric configuration.
type FabricCfgFabric struct {
	Fabric *FabricConfig `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric,omitempty"`
}

// FabricCfgFabricProfiles  represents the fabric profiles.
type FabricCfgFabricProfiles struct {
	FabricProfiles *FabricProfiles `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-profiles,omitempty"`
}

// FabricCfgFabricControlplaneNames  represents the fabric control plane names.
type FabricCfgFabricControlplaneNames struct {
	FabricControlplaneNames *FabricControlplanes `json:"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names,omitempty"`
}

type FabricCfgData struct {
	Fabric                  *FabricConfig        `json:"fabric,omitempty"`
	FabricProfiles          *FabricProfiles      `json:"fabric-profiles,omitempty"`
	FabricControlplaneNames *FabricControlplanes `json:"fabric-controlplane-names,omitempty"`
}

type FabricConfig struct {
	FabricEnabled         bool                   `json:"fabric-enabled,omitempty"`
	FabricNameVNIDEntries *FabricNameVNIDEntries `json:"fabric-name-vnid-entries,omitempty"`
}

type FabricProfiles struct {
	FabricProfile []FabricProfile `json:"fabric-profile,omitempty"`
}

type FabricProfile struct {
	FabricProfileName string `json:"fabric-profile-name"`
	// Other fabric profile fields would be added here based on the YANG model
}

type FabricControlplanes struct {
	FabricControlplaneName []FabricControlplaneName `json:"fabric-controlplane-name,omitempty"`
}

type FabricControlplaneName struct {
	ControlPlaneName            string                       `json:"control-plane-name"`
	Description                 string                       `json:"description,omitempty"`
	FabricControlPlaneIPConfigs *FabricControlPlaneIPConfigs `json:"fabric-control-plane-ip-configs,omitempty"`
}

type FabricControlPlaneIPConfigs struct {
	FabricControlPlaneIPConfig []FabricControlPlaneIPConfig `json:"fabric-control-plane-ip-config,omitempty"`
}

type FabricControlPlaneIPConfig struct {
	ControlPlaneIP string `json:"control-plane-ip"`
	PSKKey         string `json:"psk-key,omitempty"`
	PSKKeyType     string `json:"psk-key-type,omitempty"`
}

type FabricNameVNIDEntries struct {
	FabricNameVNIDEntry []FabricNameVNIDEntry `json:"fabric-name-vnid-entry,omitempty"`
}

type FabricNameVNIDEntry struct {
	Name   string `json:"name"`
	L2VNID int    `json:"l2-vnid"`
}
