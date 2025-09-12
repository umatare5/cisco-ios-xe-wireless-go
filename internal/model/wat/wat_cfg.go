package model

// WatCfg represents the complete WAT configuration from YANG 17.18.1+.
type WatCfg struct {
	WatCfgData *WatCfgData `json:"Cisco-IOS-XE-wireless-wat-cfg:wat-cfg-data"`
}

// WatCfgWatConfig represents the WAT configuration container from YANG 17.18.1+.
type WatCfgWatConfig struct {
	WatConfig *WatConfig `json:"Cisco-IOS-XE-wireless-wat-cfg:wat-config,omitempty"`
}

// WatCfgData represents the WAT configurational data container from YANG 17.18.1+.
type WatCfgData struct {
	WatConfig *WatConfig `json:"wat-config,omitempty"`
}

// WatConfig represents Wireless Active Testing (WAT) ThousandEyes configurations from YANG 17.18.1+.
type WatConfig struct {
	WatEnabled         bool   `json:"wat-enable,omitempty"`           // Enable WAT at global level
	TeConnectionString string `json:"te-conn-str,omitempty"`          // ThousandEyes connection string
	TeDownloadURL      string `json:"te-download-url,omitempty"`      // ThousandEyes download URL
	TeAgentVersion     string `json:"te-agent-version,omitempty"`     // ThousandEyes agent version
	TeCloudEndpoint    string `json:"te-cloud-endpoint,omitempty"`    // ThousandEyes cloud endpoint
	TePollInterval     int    `json:"te-poll-interval,omitempty"`     // Polling interval (seconds)
	TeTimeout          int    `json:"te-timeout,omitempty"`           // Connection timeout (seconds)
	TeRetryAttempts    int    `json:"te-retry-attempts,omitempty"`    // Number of retry attempts
	TeLogLevel         string `json:"te-log-level,omitempty"`         // Log level (debug, info, warn, error)
	TeDataCollection   bool   `json:"te-data-collection,omitempty"`   // Enable data collection
	TeAnalyticsEnabled bool   `json:"te-analytics-enabled,omitempty"` // Enable analytics
}

// WatTestProfile represents a WAT test profile configuration from YANG 17.18.1+.
type WatTestProfile struct {
	ProfileName      string             `json:"profile-name"`                // Test profile name
	Description      string             `json:"description,omitempty"`       // Profile description
	Enabled          bool               `json:"enabled,omitempty"`           // Profile enabled state
	TestType         string             `json:"test-type,omitempty"`         // Type of test (http, ping, traceroute, etc.)
	TargetURL        string             `json:"target-url,omitempty"`        // Target URL for testing
	TestInterval     int                `json:"test-interval,omitempty"`     // Test execution interval (seconds)
	TestTimeout      int                `json:"test-timeout,omitempty"`      // Test timeout (seconds)
	SuccessThreshold int                `json:"success-threshold,omitempty"` // Success threshold percentage
	FailureThreshold int                `json:"failure-threshold,omitempty"` // Failure threshold percentage
	AlertingEnabled  bool               `json:"alerting-enabled,omitempty"`  // Enable alerting
	TestParameters   *WatTestParameters `json:"test-parameters,omitempty"`   // Test-specific parameters
}

// WatTestParameters represents test-specific parameters from YANG 17.18.1+.
type WatTestParameters struct {
	HTTPMethod           string            `json:"http-method,omitempty"`            // HTTP method (GET, POST, etc.)
	HTTPHeaders          map[string]string `json:"http-headers,omitempty"`           // HTTP headers
	HTTPBody             string            `json:"http-body,omitempty"`              // HTTP request body
	ExpectedStatusCode   int               `json:"expected-status-code,omitempty"`   // Expected HTTP status code
	ExpectedResponseTime int               `json:"expected-response-time,omitempty"` // Expected response time (ms)
	PingPacketSize       int               `json:"ping-packet-size,omitempty"`       // Ping packet size (bytes)
	PingCount            int               `json:"ping-count,omitempty"`             // Number of ping packets
	TracerouteMaxHops    int               `json:"traceroute-max-hops,omitempty"`    // Maximum hops for traceroute
	DNSServer            string            `json:"dns-server,omitempty"`             // DNS server for resolution
	FollowRedirects      bool              `json:"follow-redirects,omitempty"`       // Follow HTTP redirects
}

// WatSchedule represents WAT test scheduling configuration from YANG 17.18.1+.
type WatSchedule struct {
	ScheduleName       string   `json:"schedule-name"`                  // Schedule name
	Description        string   `json:"description,omitempty"`          // Schedule description
	Enabled            bool     `json:"enabled,omitempty"`              // Schedule enabled state
	TestProfiles       []string `json:"test-profiles,omitempty"`        // Associated test profiles
	CronExpression     string   `json:"cron-expression,omitempty"`      // Cron expression for scheduling
	TimeZone           string   `json:"time-zone,omitempty"`            // Time zone
	StartDate          string   `json:"start-date,omitempty"`           // Schedule start date
	EndDate            string   `json:"end-date,omitempty"`             // Schedule end date
	MaxConcurrentTests int      `json:"max-concurrent-tests,omitempty"` // Maximum concurrent tests
	RetryFailedTests   bool     `json:"retry-failed-tests,omitempty"`   // Retry failed tests
}

// WatReport represents WAT test reporting configuration from YANG 17.18.1+.
type WatReport struct {
	ReportName         string   `json:"report-name"`                   // Report name
	Description        string   `json:"description,omitempty"`         // Report description
	Enabled            bool     `json:"enabled,omitempty"`             // Report enabled state
	ReportType         string   `json:"report-type,omitempty"`         // Report type (summary, detailed, etc.)
	TestProfiles       []string `json:"test-profiles,omitempty"`       // Test profiles to include
	ReportFormat       string   `json:"report-format,omitempty"`       // Report format (json, xml, csv, etc.)
	GenerationInterval int      `json:"generation-interval,omitempty"` // Report generation interval (hours)
	RetentionPeriod    int      `json:"retention-period,omitempty"`    // Report retention period (days)
	EmailNotification  bool     `json:"email-notification,omitempty"`  // Enable email notification
	EmailRecipients    []string `json:"email-recipients,omitempty"`    // Email recipients
}
