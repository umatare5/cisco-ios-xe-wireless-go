package mesh

// MeshOper represents mesh operational data container.
type MeshOper struct {
	MeshOperData MeshOperData `json:"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"`
}

// MeshOperData represents mesh operational data container.
type MeshOperData struct {
	// Direct lists from YANG mesh-oper-data container
	MeshQueueStats    []MeshQueueStats    `json:"mesh-q-stats,omitempty"`   // Mesh packet queue statistics (YANG: IOS-XE 17.12.1+)
	MeshDataRateStats []MeshDataRateStats `json:"mesh-dr-stats,omitempty"`  // Mesh data rate statistics (YANG: IOS-XE 17.12.1+)
	MeshSecurityStats []MeshSecurityStats `json:"mesh-sec-stats,omitempty"` // Mesh security statistics (YANG: IOS-XE 17.12.1+)
	MeshOperational   []MeshOperational   `json:"mesh-oper-data,omitempty"` // Mesh operational data (YANG: IOS-XE 17.12.1+)
}

// MeshQueueStats represents mesh access point packet queue statistics.
type MeshQueueStats struct {
	WTPMac     string  `json:"wtp-mac"`               // Wireless Termination Point MAC address (YANG: IOS-XE 17.12.1+)
	QueueType  string  `json:"q-type"`                // Queue type identifier (YANG: IOS-XE 17.12.1+)
	PeakLength *uint16 `json:"peak-length,omitempty"` // Peak number of packets waiting in queue (YANG: IOS-XE 17.12.1+)
	AverageLen *uint16 `json:"average-len,omitempty"` // Average number of packets waiting in queue (YANG: IOS-XE 17.12.1+)
	Overflows  *uint16 `json:"overflows,omitempty"`   // Number of queue overflows (YANG: IOS-XE 17.12.1+)
}

// MeshDataRateStats represents mesh access point data rate statistics.
type MeshDataRateStats struct {
	WTPMac        string  `json:"wtp-mac"`               // Wireless Termination Point MAC address (YANG: IOS-XE 17.12.1+)
	NeighborAPMac string  `json:"neigh-ap-mac"`          // Neighbor access point MAC address (YANG: IOS-XE 17.12.1+)
	DataRateIndex uint32  `json:"data-rate-index"`       // Data rate index value (YANG: IOS-XE 17.12.1+)
	TxSuccess     *uint32 `json:"tx-success,omitempty"`  // Successfully transmitted packets (YANG: IOS-XE 17.12.1+)
	TxAttempts    *uint32 `json:"tx-attempts,omitempty"` // Total transmission attempts (YANG: IOS-XE 17.12.1+)
}

// MeshSecurityStats represents mesh access point security statistics.
type MeshSecurityStats struct {
	WTPMac      string  `json:"wtp-mac"`                 // Wireless Termination Point MAC address (YANG: IOS-XE 17.12.1+)
	TxPktsTotal *uint32 `json:"tx-pkts-total,omitempty"` // Total transmitted packets during security negotiation (YANG: IOS-XE 17.12.1+)
	RxPktsTotal *uint32 `json:"rx-pkts-total,omitempty"` // Total received packets during security negotiation (YANG: IOS-XE 17.12.1+)
	RxPktsError *uint32 `json:"rx-pkts-error,omitempty"` // Total error packets received during security negotiation (YANG: IOS-XE 17.12.1+)
	// Additional parent/child stats containers will be added in future releases based on YANG model updates
}

// MeshOperational represents mesh access point operational data.
type MeshOperational struct {
	WTPMac                string  `json:"wtp-mac"`                            // Wireless Termination Point MAC address (YANG: IOS-XE 17.12.1+)
	BackhaulSlotID        *uint8  `json:"bhaul-slot-id,omitempty"`            // Backhaul radio slot identifier (YANG: IOS-XE 17.12.1+)
	BackhaulRateMCSs      *uint8  `json:"bhaul-rate-mcs-ss,omitempty"`        // Mesh backhaul 802.11ac MCS spatial stream (YANG: IOS-XE 17.12.1+)
	ActiveTrunkNativeVLAN *uint16 `json:"active-trunk-native-vlan,omitempty"` // Trunk native VLAN (YANG: IOS-XE 17.12.1+)
	ConfiguredRole        *string `json:"configured-role,omitempty"`          // Configured AP role (YANG: IOS-XE 17.12.1+)
	BackhaulRadioMode     *string `json:"bhaul-radio-mode,omitempty"`         // Backhaul radio mode (YANG: IOS-XE 17.12.1+)
	APMode                *string `json:"ap-mode,omitempty"`                  // AP mode (YANG: IOS-XE 17.12.1+)
	APRole                *string `json:"ap-role,omitempty"`                  // Current AP role (YANG: IOS-XE 17.12.1+)
	BackhaulRadioType     *string `json:"bhaul-radio-type,omitempty"`         // Backhaul radio type (YANG: IOS-XE 17.12.1+)
	BackhaulDataRateType  *string `json:"bhaul-data-rate-type,omitempty"`     // Backhaul data rate type (YANG: IOS-XE 17.12.1+)
}
