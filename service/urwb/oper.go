package urwb

import "time"

// CiscoIOSXEWirelessURWBOper represents the complete URWB network operational data from YANG 17.18.1.
type CiscoIOSXEWirelessURWBOper struct {
	CiscoIOSXEWirelessURWBOperData *struct {
		UrwbnetStats []URWBnetStats `json:"urwbnet-stats"`
		UrwbnetNodeG []URWBnetNodeG `json:"urwbnet-node-g,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data"`
}

// CiscoIOSXEWirelessURWBnetOperUrwbnetStats represents the URWB network stats container from YANG 17.18.1.
type CiscoIOSXEWirelessURWBnetOperUrwbnetStats struct {
	UrwbnetStats []URWBnetStats `json:"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-stats"`
}

// URWBnetStats represents URWB network operational data per coordinator from YANG 17.18.1.
type URWBnetStats struct {
	MAC                string              `json:"mac"`                    // Coordinator MAC address
	Gateway            *URWBnetMeshID      `json:"gw,omitempty"`           // Gateway identifier
	NodesCount         int                 `json:"nodes-cnt,omitempty"`    // URWB network node count
	UrwbnetNode        []URWBnetNode       `json:"urwbnet-node,omitempty"` // URWB network nodes
	CoordRouteCount    int                 `json:"coord-rt-cnt,omitempty"` // Coordinator route count
	UrwbnetCoordRoutes []URWBnetCoordRoute `json:"urwbnet-coord-routes,omitempty"`
}

// URWBnetMeshID represents mesh identifier information from YANG 17.18.1.
type URWBnetMeshID struct {
	Address   string `json:"addr"`                 // Network address
	MeshName  string `json:"mesh-name,omitempty"`  // Mesh network name
	NetworkID string `json:"network-id,omitempty"` // Network identifier
}

// URWBnetNode represents URWB network node information from YANG 17.18.1.
type URWBnetNode struct {
	// From st-urwbnet-mesh-id (inherited by key "addr")
	Address string `json:"addr"` // Node address (key field)

	// From st-urwbnet-node grouping
	ID           *URWBnetMeshID    `json:"id,omitempty"`             // URWB network node id
	Name         string            `json:"name,omitempty"`           // URWB network node name
	DeviceModel  string            `json:"device-model,omitempty"`   // URWB network node device model
	Role         string            `json:"role,omitempty"`           // uar enum - URWB network node role
	IPAddr       string            `json:"ip-addr,omitempty"`        // URWB network node IP address
	NetMask      string            `json:"net-mask,omitempty"`       // URWB network node netmask
	NodeIf       []URWBnetNodeIntf `json:"node-if,omitempty"`        // URWB network node interface list (max 16)
	ConnDevCount uint8             `json:"conn-dev-count,omitempty"` // Connected device count
}

// URWBnetNodeIntf represents URWB network node interface statistics from YANG 17.18.1.
type URWBnetNodeIntf struct {
	IntfID      *URWBnetMeshID `json:"intf-id,omitempty"`      // URWB network node interface id
	IntfEnabled bool           `json:"intf-enabled,omitempty"` // Interface enable status
	Role        string         `json:"role,omitempty"`         // urr enum - radio role
	Freq        int32          `json:"freq,omitempty"`         // MHz - radio frequency
	Bw          string         `json:"bw,omitempty"`           // urwbnet-bw enum - bandwidth
	TxPower     int32          `json:"tx-power,omitempty"`     // dBm - tx power
	IsTdma      bool           `json:"is-tdma,omitempty"`      // time division multiple access
	FreqScan    bool           `json:"freq-scan,omitempty"`    // frequency scan status
}

// URWBnetNodeMetrics represents node performance metrics from YANG 17.18.1.
type URWBnetNodeMetrics struct {
	BytesSent       uint64    `json:"bytes-sent,omitempty"`
	BytesReceived   uint64    `json:"bytes-received,omitempty"`
	PacketsSent     uint64    `json:"packets-sent,omitempty"`
	PacketsReceived uint64    `json:"packets-received,omitempty"`
	ErrorCount      uint64    `json:"error-count,omitempty"`
	RetryCount      uint64    `json:"retry-count,omitempty"`
	Uptime          time.Time `json:"uptime,omitempty"`
}

// URWBnetCoordRoute represents coordinator routing information from YANG 17.18.1.
type URWBnetCoordRoute struct {
	Destination string    `json:"destination"`            // Destination address
	NextHop     string    `json:"next-hop,omitempty"`     // Next hop address
	Metric      int       `json:"metric,omitempty"`       // Route metric
	Interface   string    `json:"interface,omitempty"`    // Output interface
	RouteType   string    `json:"route-type,omitempty"`   // Route type (direct, mesh, etc.)
	LastUpdated time.Time `json:"last-updated,omitempty"` // Route last update time
}

// URWBnetNodeG represents URWB network node group information from YANG 17.18.1.
type URWBnetNodeG struct {
	GroupID        string        `json:"group-id"`                  // Group identifier
	GroupName      string        `json:"group-name,omitempty"`      // Group name
	CoordinatorMAC string        `json:"coordinator-mac,omitempty"` // Coordinator MAC
	Members        []URWBnetNode `json:"members,omitempty"`         // Group members
	Status         string        `json:"status,omitempty"`          // Group status
	Formation      time.Time     `json:"formation,omitempty"`       // Group formation time
}
