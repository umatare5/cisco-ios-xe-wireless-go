package model

import "time"

// UrwbnetOper represents the complete URWB network operational data from YANG 17.18.1+.
type UrwbnetOper struct {
	UrwbnetOperData *UrwbnetOperData `json:"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data"`
}

// UrwbnetOperUrwbnetStats represents the URWB network stats container from YANG 17.18.1+.
type UrwbnetOperUrwbnetStats struct {
	UrwbnetStats []UrwbnetStats `json:"Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-stats"`
}

// UrwbnetOperData represents the URWB network operational data container from YANG 17.18.1+.
type UrwbnetOperData struct {
	UrwbnetStats []UrwbnetStats `json:"urwbnet-stats"`
	UrwbnetNodeG []UrwbnetNodeG `json:"urwbnet-node-g,omitempty"`
}

// UrwbnetStats represents URWB network operational data per coordinator from YANG 17.18.1+.
type UrwbnetStats struct {
	Mac                string              `json:"mac"`                    // Coordinator MAC address
	Gateway            *UrwbnetMeshID      `json:"gw,omitempty"`           // Gateway identifier
	NodesCount         int                 `json:"nodes-cnt,omitempty"`    // URWB network node count
	UrwbnetNode        []UrwbnetNode       `json:"urwbnet-node,omitempty"` // URWB network nodes
	CoordRouteCount    int                 `json:"coord-rt-cnt,omitempty"` // Coordinator route count
	UrwbnetCoordRoutes []UrwbnetCoordRoute `json:"urwbnet-coord-routes,omitempty"`
}

// UrwbnetMeshID represents mesh identifier information.
type UrwbnetMeshID struct {
	Address   string `json:"addr"`                 // Network address
	MeshName  string `json:"mesh-name,omitempty"`  // Mesh network name
	NetworkID string `json:"network-id,omitempty"` // Network identifier
}

// UrwbnetNode represents URWB network node information.
type UrwbnetNode struct {
	// From st-urwbnet-mesh-id (inherited by key "addr")
	Address string `json:"addr"` // Node address (key field)

	// From st-urwbnet-node grouping
	ID           *UrwbnetMeshID    `json:"id,omitempty"`             // URWB network node id
	Name         string            `json:"name,omitempty"`           // URWB network node name
	DeviceModel  string            `json:"device-model,omitempty"`   // URWB network node device model
	Role         string            `json:"role,omitempty"`           // uar enum - URWB network node role
	IPAddr       string            `json:"ip-addr,omitempty"`        // URWB network node IP address
	NetMask      string            `json:"net-mask,omitempty"`       // URWB network node netmask
	NodeIf       []UrwbnetNodeIntf `json:"node-if,omitempty"`        // URWB network node interface list (max 16)
	ConnDevCount uint8             `json:"conn-dev-count,omitempty"` // Connected device count
}

// UrwbnetNodeIntf represents URWB network node interface statistics.
type UrwbnetNodeIntf struct {
	IntfID      *UrwbnetMeshID `json:"intf-id,omitempty"`      // URWB network node interface id
	IntfEnabled bool           `json:"intf-enabled,omitempty"` // Interface enable status
	Role        string         `json:"role,omitempty"`         // urr enum - radio role
	Freq        int32          `json:"freq,omitempty"`         // MHz - radio frequency
	Bw          string         `json:"bw,omitempty"`           // urwbnet-bw enum - bandwidth
	TxPower     int32          `json:"tx-power,omitempty"`     // dBm - tx power
	IsTdma      bool           `json:"is-tdma,omitempty"`      // time division multiple access
	FreqScan    bool           `json:"freq-scan,omitempty"`    // frequency scan status
}

// UrwbnetNodeMetrics represents node performance metrics.
type UrwbnetNodeMetrics struct {
	BytesSent       uint64    `json:"bytes-sent,omitempty"`
	BytesReceived   uint64    `json:"bytes-received,omitempty"`
	PacketsSent     uint64    `json:"packets-sent,omitempty"`
	PacketsReceived uint64    `json:"packets-received,omitempty"`
	ErrorCount      uint64    `json:"error-count,omitempty"`
	RetryCount      uint64    `json:"retry-count,omitempty"`
	Uptime          time.Time `json:"uptime,omitempty"`
}

// UrwbnetCoordRoute represents coordinator routing information.
type UrwbnetCoordRoute struct {
	Destination string    `json:"destination"`            // Destination address
	NextHop     string    `json:"next-hop,omitempty"`     // Next hop address
	Metric      int       `json:"metric,omitempty"`       // Route metric
	Interface   string    `json:"interface,omitempty"`    // Output interface
	RouteType   string    `json:"route-type,omitempty"`   // Route type (direct, mesh, etc.)
	LastUpdated time.Time `json:"last-updated,omitempty"` // Route last update time
}

// UrwbnetNodeG represents URWB network node group information.
type UrwbnetNodeG struct {
	GroupID        string        `json:"group-id"`                  // Group identifier
	GroupName      string        `json:"group-name,omitempty"`      // Group name
	CoordinatorMac string        `json:"coordinator-mac,omitempty"` // Coordinator MAC
	Members        []UrwbnetNode `json:"members,omitempty"`         // Group members
	Status         string        `json:"status,omitempty"`          // Group status
	Formation      time.Time     `json:"formation,omitempty"`       // Group formation time
}
