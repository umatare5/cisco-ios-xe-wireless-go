// Package model provides data models for LISP operational data.
package model

// LispOper  represents the structure for LISP operational data.
type LispOper struct {
	LispOperData LispOperData `json:"Cisco-IOS-XE-wireless-lisp-oper:lisp-oper-data"`
}

// LispOperLispInstances  represents the structure for LISP instances.
type LispOperLispInstances struct {
	LispInstances []LispInstance `json:"Cisco-IOS-XE-wireless-lisp-oper:lisp-instances"`
}

// LispOperLispStats  represents the structure for LISP statistics.
type LispOperLispStats struct {
	LispStats LispStats `json:"Cisco-IOS-XE-wireless-lisp-oper:lisp-stats"`
}

type LispOperData struct {
	LispInstances []LispInstance `json:"lisp-instances"`
	LispStats     LispStats      `json:"lisp-stats"`
}

type LispInstance struct {
	InstanceID int         `json:"instance-id"`
	VrfName    string      `json:"vrf-name"`
	Status     string      `json:"status"`
	EidTables  []EidTable  `json:"eid-tables"`
	MapServers []MapServer `json:"map-servers"`
}

type EidTable struct {
	EidPrefix     string   `json:"eid-prefix"`
	LocalizedSets []string `json:"localized-sets"`
	RemoteSets    []string `json:"remote-sets"`
}

type MapServer struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	Status     string `json:"status"`
	Registered bool   `json:"registered"`
}

type LispStats struct {
	TotalMappings  int `json:"total-mappings"`
	LocalMappings  int `json:"local-mappings"`
	RemoteMappings int `json:"remote-mappings"`
	MapRequests    int `json:"map-requests"`
	MapReplies     int `json:"map-replies"`
}
