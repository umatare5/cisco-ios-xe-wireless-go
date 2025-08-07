package model

// LispOperResponse represents the response structure for LISP operational data.
type LispOperResponse struct {
	LispOperData LispOperData `json:"Cisco-IOS-XE-wireless-lisp-oper:lisp-oper-data"`
}

// LispOperData contains LISP operational data
type LispOperData struct {
	LispInstances []LispInstance `json:"lisp-instances"`
	LispStats     LispStats      `json:"lisp-stats"`
}

// LispInstance represents LISP instance information
type LispInstance struct {
	InstanceID int         `json:"instance-id"`
	VrfName    string      `json:"vrf-name"`
	Status     string      `json:"status"`
	EidTables  []EidTable  `json:"eid-tables"`
	MapServers []MapServer `json:"map-servers"`
}

// EidTable represents EID table information
type EidTable struct {
	EidPrefix     string   `json:"eid-prefix"`
	LocalizedSets []string `json:"localized-sets"`
	RemoteSets    []string `json:"remote-sets"`
}

// MapServer represents Map Server information
type MapServer struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	Status     string `json:"status"`
	Registered bool   `json:"registered"`
}

// LispStats represents LISP statistics
type LispStats struct {
	TotalMappings  int `json:"total-mappings"`
	LocalMappings  int `json:"local-mappings"`
	RemoteMappings int `json:"remote-mappings"`
	MapRequests    int `json:"map-requests"`
	MapReplies     int `json:"map-replies"`
}
