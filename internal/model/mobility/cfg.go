package mobility

// MobilityCfg represents the root container for mobility configuration data.
type MobilityCfg struct {
	MobilityCfgData MobilityCfgData `json:"Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data"`
}

// MobilityCfgData represents mobility configuration data container.
type MobilityCfgData struct {
	MobilityConfig MobilityConfig `json:"mobility-config"`
}

// MobilityConfig represents local mobility configuration.
type MobilityConfig struct {
	LocalGroup                string          `json:"local-group"`                             // Local mobility group name
	LocalMcastAddrEnabled     *bool           `json:"local-mcast-addr-enabled,omitempty"`      // Enable IPv4 multicast support for local mobility group (YANG: IOS-XE 17.12.1+)
	LocalMulticastAddress     *string         `json:"local-multicast-address,omitempty"`       // Local mobility Multicast IPv4 address (YANG: IOS-XE 17.12.1+)
	LocalIPv6McastAddrEnabled *bool           `json:"local-ipv6-mcast-addr-enabled,omitempty"` // Enable IPv6 multicast support for local mobility group (YANG: IOS-XE 17.12.1+)
	LocalIPv6MulticastAddress *string         `json:"local-ipv6-multicast-address,omitempty"`  // Local mobility Multicast IPv6 address (YANG: IOS-XE 17.12.1+)
	MobilityKeepaliveInterval *uint16         `json:"mobility-keepalive-interval,omitempty"`   // Mobility Keep Alive interval (YANG: IOS-XE 17.12.1+)
	MobilityKeepaliveCount    *uint16         `json:"mobility-keepalive-count,omitempty"`      // Mobility Keep Alive count (YANG: IOS-XE 17.12.1+)
	MobilityDSCP              *uint8          `json:"mobility-dscp,omitempty"`                 // mobility dcsp value (YANG: IOS-XE 17.12.1+)
	MACAddress                string          `json:"mac-address"`                             // Mobility local Mac Address
	MobilityPeers             *MobilityPeers  `json:"mobility-peers,omitempty"`                // Mobility Peers configuration (YANG: IOS-XE 17.12.1+)
	MobilityGroups            *MobilityGroups `json:"mobility-groups,omitempty"`               // Mobility groups configuration (YANG: IOS-XE 17.12.1+)
	MmDTLSHighCipher          *bool           `json:"mm-dtls-high-cipher,omitempty"`           // Enable/Disable DTLS high cipher for mobility (YANG: IOS-XE 17.12.1+)
}

// MobilityPeers represents mobility peers configuration container.
type MobilityPeers struct {
	MobilityPeer []MobilityPeer `json:"mobility-peer,omitempty"` // List of Mobility Peers (YANG: IOS-XE 17.12.1+)
}

// MobilityPeer represents mobility peer configuration.
type MobilityPeer struct {
	MACAddr   string  `json:"mac-addr"`         // Mobility peer MAC address (YANG: IOS-XE 17.12.1+)
	IPAddress string  `json:"ip-address"`       // Peer IP Address (YANG: IOS-XE 17.12.1+)
	GroupName string  `json:"group-name"`       // Peer group name (YANG: IOS-XE 17.12.1+)
	NatIP     *string `json:"nat-ip,omitempty"` // NAT IP address for peer (YANG: IOS-XE 17.12.1+)
}

// MobilityGroups represents mobility groups configuration container.
type MobilityGroups struct {
	MobilityGroup []MobilityGroup `json:"mobility-group,omitempty"` // List of mobility groups (YANG: IOS-XE 17.12.1+)
}

// MobilityGroup represents remote mobility group configuration.
type MobilityGroup struct {
	GroupName          string  `json:"group-name"`                     // Group name for remote mobility group (YANG: IOS-XE 17.12.1+)
	McastAddrEnabled   *bool   `json:"mcast-addr-enabled,omitempty"`   // Enable IPv4 multicast support for mobility group (YANG: IOS-XE 17.12.1+)
	MulticastAddress   *string `json:"multicast-address,omitempty"`    // Mobility multicast ipv4 address for remote group (YANG: IOS-XE 17.12.1+)
	McastAddrv6Enabled *bool   `json:"mcast-addrv6-enabled,omitempty"` // Enable IPv6 multicast support for mobility group (YANG: IOS-XE 17.12.1+)
	MulticastAddressv6 *string `json:"multicast-addressv6,omitempty"`  // Mobility multicast ipv6 address for remote group (YANG: IOS-XE 17.12.1+)
}
