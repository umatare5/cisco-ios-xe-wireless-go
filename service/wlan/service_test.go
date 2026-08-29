package wlan

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	assert "github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestWlanServiceUnit_Constructor_Success(t *testing.T) {
	service := NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client service")
	}
}

func TestWlanServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Mock server with basic WLAN response structure
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {
				"global-params": {
					"country-code": "US"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{
						"profile-name": "test-wlan",
						"ssid": "TEST_SSID",
						"admin-status": true
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries/wlan-cfg-entry=test-wlan": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entry": [
				{
					"profile-name": "test-wlan",
					"ssid": "TEST_SSID",
					"admin-status": true
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {
				"wlan-policy": [
					{
						"policy-name": "test-policy",
						"description": "Test policy"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
				"policy-list-entry": [
					{
						"tag-name": "test-policy-tag",
						"description": "Test policy tag"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wireless-aaa-policy-configs": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs": {
				"wireless-aaa-policy-config": [
					{
						"policy-name": "test-aaa-policy",
						"description": "Test AAA policy"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": `{
			"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": {
				"global-stats": {
					"total-wlans": 2,
					"active-wlans": 1
				}
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetConfig
	config, err := service.GetConfig(ctx)
	if err != nil {
		t.Errorf("GetConfig failed: %v", err)
		return
	}

	if config == nil {
		t.Error("GetConfig returned nil result")
		return
	}

	// Test GetOperational
	operational, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
		return
	}

	if operational == nil {
		t.Error("GetOperational returned nil result")
		return
	}

	// Test ListWlanCfgEntries
	cfgEntries, err := service.ListWlanCfgEntries(ctx)
	if err != nil {
		t.Errorf("ListWlanCfgEntries failed: %v", err)
		return
	}

	if cfgEntries == nil {
		t.Error("ListWlanCfgEntries returned nil result")
		return
	}

	// Test ListWlanPolicies
	wlanPolicies, err := service.ListWlanPolicies(ctx)
	if err != nil {
		t.Errorf("ListWlanPolicies failed: %v", err)
		return
	}

	if wlanPolicies == nil {
		t.Error("ListWlanPolicies returned nil result")
		return
	}

	// Test ListCfgPolicyListEntries
	cfgPolicyEntries, err := service.ListCfgPolicyListEntries(ctx)
	if err != nil {
		t.Errorf("ListCfgPolicyListEntries failed: %v", err)
		return
	}

	if cfgPolicyEntries == nil {
		t.Error("ListCfgPolicyListEntries returned nil result")
		return
	}

	// Test ListCfgWirelessAaaPolicyConfigs
	cfgAaaConfigs, err := service.ListCfgWirelessAaaPolicyConfigs(ctx)
	if err != nil {
		t.Errorf("ListCfgWirelessAaaPolicyConfigs failed: %v", err)
		return
	}

	if cfgAaaConfigs == nil {
		t.Error("ListCfgWirelessAaaPolicyConfigs returned nil result")
		return
	}

	// Test ListDot11beProfiles (skip if not supported by mock server)
	dot11beProfiles, err := service.ListDot11beProfiles(ctx)
	if err != nil {
		// Wi-Fi 7 endpoints may not be supported by all mock servers
		t.Logf("ListDot11beProfiles failed (expected for older mock servers): %v", err)
	} else if dot11beProfiles == nil {
		t.Error("ListDot11beProfiles returned nil result")
		return
	}

	// Test ListWlanInfo (skip if not supported by mock server)
	wlanInfo, err := service.ListWlanInfo(ctx)
	if err != nil {
		// WlanInfo endpoint may not be supported by all mock servers
		t.Logf("ListWlanInfo failed (expected for older mock servers): %v", err)
	} else if wlanInfo == nil {
		t.Error("ListWlanInfo returned nil result")
		return
	}

	t.Logf("All get operations returned valid WLAN data")
}

// TestWlanServiceUnit_SecretRedaction_Success pins the four methods that keep a pre-shared key
// out of a log and the one verb that still carries it. Each direction is a separate assertion
// because each is held by a different method: dropping Secret.LogValue leaves the JSON handler
// leaking the field, and dropping WlanCfgEntry.LogValue leaves it leaking the whole entry.
func TestWlanServiceUnit_SecretRedaction_Success(t *testing.T) {
	const key = "not-a-real-key"

	entry := WlanCfgEntry{WlanID: 7, ProfileName: "profile-redaction", PSK: Secret(key)}

	if got := entry.PSK.String(); got != redacted {
		t.Errorf("Secret.String() = %q, want %q", got, redacted)
	}
	if got := entry.PSK.Reveal(); got != key {
		t.Errorf("Secret.Reveal() = %q, want the key back", got)
	}
	if got := fmt.Sprintf("%+v", entry); strings.Contains(got, key) {
		t.Errorf("%%+v of the entry carried the key: %s", got)
	}

	// The entry is held by WlanCfgEntry.LogValue, which renders the identifying leaves and no
	// configuration at all, so the assertion is absence of the key rather than presence of the
	// redaction. Without that method the handler renders the entry through json.Marshal.
	if got := logJSON(t, "entry", entry); strings.Contains(got, key) {
		t.Errorf("slog.JSONHandler carried the key for the entry: %s", got)
	}

	// The field on its own is held by Secret.LogValue, which does render the redaction.
	got := logJSON(t, "psk", entry.PSK)
	if strings.Contains(got, key) {
		t.Errorf("slog.JSONHandler carried the key for the field: %s", got)
	}
	if !strings.Contains(got, redacted) {
		t.Errorf("slog.JSONHandler rendered no redaction for the field: %s", got)
	}

	// The JSON path is held by Secret.MarshalJSON, which covers the containers above the entry
	// as well: WlanCfgEntries and the response wrapper have no LogValue of their own.
	body, err := json.Marshal(entry)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	if strings.Contains(string(body), key) {
		t.Errorf("json.Marshal carried the key: %s", body)
	}
	if !strings.Contains(string(body), redacted) {
		t.Errorf("json.Marshal rendered no redaction: %s", body)
	}

	// The empty case is asserted on the bare value because encoding/json omits an empty omitempty
	// field without calling the marshaler, so the struct cannot show what the method returns.
	empty, err := json.Marshal(Secret(""))
	if err != nil {
		t.Fatalf("json.Marshal of an empty Secret failed: %v", err)
	}
	if string(empty) != `""` {
		t.Errorf("json.Marshal of an empty Secret = %s, want an empty string", empty)
	}

	// The one residual, asserted in the direction it must keep: %#v prints the underlying string
	// of any named type.
	if got := fmt.Sprintf("%#v", entry); !strings.Contains(got, key) {
		t.Errorf("%%#v redacted, so the doc comment naming it a residual is now wrong")
	}
}

// logJSON returns the one record a slog.JSONHandler writes for a single attribute.
func logJSON(t *testing.T, key string, value any) string {
	t.Helper()

	var buf bytes.Buffer
	slog.New(slog.NewJSONHandler(&buf, nil)).Info("redaction", key, value)

	return buf.String()
}

func TestWlanServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := NewService(nil)
	ctx := testutil.TestContext(t)

	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetConfig")
	}

	_, err = service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetOperational")
	}

	_, err = service.ListWlanCfgEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListWlanCfgEntries")
	}

	_, err = service.ListWlanPolicies(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListWlanPolicies")
	}

	_, err = service.ListCfgPolicyListEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListCfgPolicyListEntries")
	}

	_, err = service.ListCfgWirelessAaaPolicyConfigs(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListCfgWirelessAaaPolicyConfigs")
	}

	// Note: ListDot11beProfiles and ListWlanInfo are not tested with nil client as they may not be supported by all mock servers
}

// TestWlanServiceUnit_OmittedSecurityLeaf_MockSuccess tests that an omitted security leaf stays
// nil while an explicitly configured false decodes to a non-nil false.
func TestWlanServiceUnit_OmittedSecurityLeaf_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{"wlan-id": 1, "profile-name": "profile-default", "apf-vap-id-data": {"ssid": "ssid-default"}},
					{
						"wlan-id": 2,
						"profile-name": "profile-explicit",
						"wpa2-enabled": false,
						"wpa2-aes": false,
						"security-wpa": false,
						"okc": false,
						"wlan-11k-neigh-list": false,
						"apf-vap-802-11v-data": {"dot11v-dms": false},
						"apf-vap-id-data": {"ssid": "ssid-explicit", "wlan-status": false, "broadcast-ssid": false}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.ListWlanCfgEntries(ctx)
	if err != nil {
		t.Fatalf("ListWlanCfgEntries failed: %v", err)
	}
	if result.WlanCfgEntries == nil || len(result.WlanCfgEntries.WlanCfgEntry) != 2 {
		t.Fatalf("Expected 2 entries, got %+v", result.WlanCfgEntries)
	}

	entries := result.WlanCfgEntries.WlanCfgEntry
	if entries[0].WPA2Enabled != nil {
		t.Error("Expected an omitted wpa2-enabled to stay nil: the default in force is not false")
	}
	if entries[0].Wlan11kNeighList != nil {
		t.Error("Expected an omitted wlan-11k-neigh-list to stay nil")
	}
	if entries[0].APFVap80211vData != nil {
		t.Error("Expected an omitted apf-vap-802-11v-data container to stay nil")
	}
	// Every leaf this test retypes is asserted in both directions. An omission-only assertion
	// leaves the JSON tag unreachable, so renaming it breaks nothing a nil-check can observe.
	if entries[1].WPA2Enabled == nil || *entries[1].WPA2Enabled {
		t.Error("Expected an explicit false to decode to a non-nil false")
	}
	if entries[1].Wlan11kNeighList == nil || *entries[1].Wlan11kNeighList {
		t.Error("Expected an explicit false wlan-11k-neigh-list to decode to a non-nil false")
	}
	if entries[1].APFVap80211vData == nil || entries[1].APFVap80211vData.Dot11vDms == nil ||
		*entries[1].APFVap80211vData.Dot11vDms {
		t.Error("Expected an explicit false dot11v-dms to decode to a non-nil false")
	}
	if entries[1].APFVapIDData == nil || entries[1].APFVapIDData.WlanStatus == nil {
		t.Fatal("Expected wlan-status to decode")
	}
	if *entries[1].APFVapIDData.WlanStatus {
		t.Error("Expected the explicit false for wlan-status")
	}

	// The three default-true leaves this type gains on the entry are read through one map per
	// direction, so a failure names its leaf rather than its line.
	for leaf, got := range map[string]*bool{
		"security-wpa": entries[0].SecurityWPA,
		"wpa2-aes":     entries[0].WPA2AES,
		"okc":          entries[0].OKC,
	} {
		if got != nil {
			t.Errorf("Expected an omitted %s to stay nil: the default in force is not false", leaf)
		}
	}

	for leaf, got := range map[string]*bool{
		"security-wpa": entries[1].SecurityWPA,
		"wpa2-aes":     entries[1].WPA2AES,
		"okc":          entries[1].OKC,
	} {
		if got == nil || *got {
			t.Errorf("Expected an explicit false %s to decode to a non-nil false", leaf)
		}
	}

	// broadcast-ssid's absence is read off a container that did arrive, with ssid as the control: a
	// nil container would leave the leaf itself unasserted.
	if entries[0].APFVapIDData == nil || entries[0].APFVapIDData.SSID != "ssid-default" {
		t.Fatalf("Expected the first record to carry apf-vap-id-data with ssid, got %+v",
			entries[0].APFVapIDData)
	}
	if entries[0].APFVapIDData.BroadcastSSID != nil {
		t.Error("Expected an omitted broadcast-ssid to stay nil")
	}
	if entries[1].APFVapIDData.BroadcastSSID == nil || *entries[1].APFVapIDData.BroadcastSSID {
		t.Error("Expected an explicit false broadcast-ssid to decode to a non-nil false")
	}
}

// TestWlanServiceUnit_SecurityLeafTags_MockSuccess reaches every security leaf this type declares
// as a value, which no nil-check can reach: a value field decodes an absent leaf and a misspelled
// tag to the same zero, so the tag is only exercised by a body that sets the leaf to its non-zero.
func TestWlanServiceUnit_SecurityLeafTags_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{
						"wlan-id": 3,
						"profile-name": "profile-every-leaf",
						"auth-key-mgmt-psk-sha256": true,
						"auth-key-mgmt-cckm": true,
						"auth-key-mgmt-sae": true,
						"auth-key-mgmt-ft-psk": true,
						"auth-key-mgmt-ft-dot1x": true,
						"auth-key-mgmt-ft-sae": true,
						"auth-key-mgmt-suite-b": true,
						"auth-key-mgmt-suite-b-192": true,
						"akm-owe": true,
						"akm-sae-ext-key": true,
						"akm-ft-sae-ext-key": true,
						"wpa1-enabled": true,
						"wep-enabled": true,
						"osen": true,
						"dot11-auth-type": "apf-vap-80211-auth-open",
						"apf-vap-id-data": {
							"ssid": "ssid-every-leaf",
							"p2p-block-action": "p2p-blocking-action-drop"
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.ListWlanCfgEntries(ctx)
	if err != nil {
		t.Fatalf("ListWlanCfgEntries failed: %v", err)
	}
	if result.WlanCfgEntries == nil || len(result.WlanCfgEntries.WlanCfgEntry) != 1 {
		t.Fatalf("Expected 1 entry, got %+v", result.WlanCfgEntries)
	}

	entry := result.WlanCfgEntries.WlanCfgEntry[0]

	for leaf, got := range map[string]bool{
		"auth-key-mgmt-psk-sha256":  entry.AuthKeyMgmtPskSha256,
		"auth-key-mgmt-cckm":        entry.AuthKeyMgmtCckm,
		"auth-key-mgmt-sae":         entry.AuthKeyMgmtSae,
		"auth-key-mgmt-ft-psk":      entry.AuthKeyMgmtFtPsk,
		"auth-key-mgmt-ft-dot1x":    entry.AuthKeyMgmtFtDot1x,
		"auth-key-mgmt-ft-sae":      entry.AuthKeyMgmtFtSae,
		"auth-key-mgmt-suite-b":     entry.AuthKeyMgmtSuiteB,
		"auth-key-mgmt-suite-b-192": entry.AuthKeyMgmtSuiteB192,
		"akm-owe":                   entry.AkmOwe,
		"akm-sae-ext-key":           entry.AkmSaeExtKey,
		"akm-ft-sae-ext-key":        entry.AkmFtSaeExtKey,
		"wpa1-enabled":              entry.WPA1Enabled,
		"wep-enabled":               entry.WEPEnabled,
		"osen":                      entry.OSEN,
	} {
		if !got {
			t.Errorf("%s decoded to false, so the field does not carry that tag", leaf)
		}
	}

	if entry.Dot11AuthType != "apf-vap-80211-auth-open" {
		t.Errorf("dot11-auth-type decoded to %q, so the field does not carry that tag",
			entry.Dot11AuthType)
	}
	if entry.APFVapIDData == nil || entry.APFVapIDData.P2PBlockAction != "p2p-blocking-action-drop" {
		t.Errorf("p2p-block-action did not decode, so the field does not carry that tag: %+v",
			entry.APFVapIDData)
	}
}

// TestWlanServiceUnit_PartialTimeoutContainer_MockSuccess tests that a container arriving with only
// some of its leaves leaves the omitted ones nil, and that an explicit zero stays distinguishable
// from an omission. Both record shapes are the measured ones: a plain read sends wlan-timeout with
// session-timeout alone and wlan-switching-policy with two of its four central-* leaves, while a
// with-defaults read sends every leaf and idle-threshold arrives as zero.
func TestWlanServiceUnit_PartialTimeoutContainer_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {
				"wlan-policy": [
					{
						"policy-profile-name": "policy-omitted-leaves",
						"wlan-timeout": {"session-timeout": 1800},
						"wlan-switching-policy": {"central-switching": false, "central-dhcp": false}
					},
					{
						"policy-profile-name": "policy-every-leaf",
						"wlan-timeout": {"session-timeout": 1800, "idle-timeout": 300, "idle-threshold": 0},
						"wlan-switching-policy": {
							"central-switching": false,
							"central-authentication": false,
							"central-dhcp": false,
							"central-assoc-enable": false
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.ListWlanPolicies(ctx)
	if err != nil {
		t.Fatalf("ListWlanPolicies failed: %v", err)
	}
	if result.WlanPolicies == nil || len(result.WlanPolicies.WlanPolicy) != 2 {
		t.Fatalf("Expected 2 policies, got %+v", result.WlanPolicies)
	}

	partial := result.WlanPolicies.WlanPolicy[0]
	if partial.WlanTimeout == nil || partial.WlanTimeout.SessionTimeout == nil {
		t.Fatal("Expected session-timeout to decode")
	}
	if partial.WlanTimeout.IdleTimeout != nil || partial.WlanTimeout.IdleThreshold != nil {
		t.Error("Expected the idle leaves omitted from a present container to stay nil")
	}
	if partial.WlanSwitchingPolicy == nil || partial.WlanSwitchingPolicy.CentralSwitching == nil {
		t.Fatal("Expected central-switching to decode")
	}
	if partial.WlanSwitchingPolicy.CentralAuthentication != nil ||
		partial.WlanSwitchingPolicy.CentralAssocEnable != nil {
		t.Error("Expected an omitted central-* leaf to stay nil: the default in force is true")
	}

	// central-dhcp arrives on both records, so it is asserted in the present direction. An
	// omission-only assertion leaves the tag unreachable: renaming it changes nothing a nil-check
	// can see.
	if partial.WlanSwitchingPolicy.CentralDHCP == nil || *partial.WlanSwitchingPolicy.CentralDHCP {
		t.Error("Expected an explicit false central-dhcp to decode to a non-nil false")
	}

	complete := result.WlanPolicies.WlanPolicy[1]
	if complete.WlanTimeout.IdleThreshold == nil || *complete.WlanTimeout.IdleThreshold != 0 {
		t.Error("Expected an explicit zero to decode to a non-nil zero")
	}
	if complete.WlanTimeout.IdleTimeout == nil || *complete.WlanTimeout.IdleTimeout != 300 {
		t.Error("Expected the idle-timeout of a complete container to decode to its value")
	}
	if complete.WlanSwitchingPolicy.CentralAuthentication == nil ||
		complete.WlanSwitchingPolicy.CentralAssocEnable == nil {
		t.Error("Expected every central-* leaf of a complete container to decode")
	}
	if complete.WlanSwitchingPolicy.CentralDHCP == nil || *complete.WlanSwitchingPolicy.CentralDHCP {
		t.Error("Expected the central-dhcp of a complete container to decode as a non-nil false")
	}
}

// queryRecorder stores the raw query string of the request the server received last.
type queryRecorder struct {
	mu  sync.Mutex
	raw string
}

func (q *queryRecorder) set(raw string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.raw = raw
}

func (q *queryRecorder) get() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.raw
}

// newRecordingService starts a server that records the query of each request and answers
// with body, then returns a service bound to it. A local recorder needs no per-route handler
// registration, so one body serves a table over several routes; body must be the envelope of
// the node under read, because core.Get holds the sole top-level key against the endpoint.
func newRecordingService(t *testing.T, body string) (Service, *queryRecorder) {
	t.Helper()

	recorder := &queryRecorder{}
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder.set(r.URL.RawQuery)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	parsed, err := url.Parse(server.URL)
	assert.AssertNoError(t, err, "parse test server URL")

	client, err := core.New(parsed.Host, "test-token", core.WithInsecureSkipVerify(true))
	assert.AssertNoError(t, err, "create core client")

	return NewService(client), recorder
}

// assertQueryBothDirections calls read twice: once without an option, where the wire must
// carry no query, and once with WithDefaults, where the wire must carry the parameter.
func assertQueryBothDirections(t *testing.T, read func(ctx context.Context, opts ...core.GetOption) error,
	recorder *queryRecorder,
) {
	t.Helper()
	ctx := context.Background()

	assert.AssertNoError(t, read(ctx), "read without option")
	assert.AssertStringEquals(t, recorder.get(), "", "RawQuery without option")

	assert.AssertNoError(t, read(ctx, core.WithDefaults(core.DefaultsReportAll)), "read with option")
	assert.AssertStringEquals(t, recorder.get(), "with-defaults=report-all", "RawQuery with option")
}

func TestWlanServiceUnit_GetOptions_ConfigRouteWireQuery(t *testing.T) {
	tests := []struct {
		name string
		body string
		read func(s Service) func(ctx context.Context, opts ...core.GetOption) error
	}{
		{
			name: "GetConfig",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.GetConfig(ctx, opts...)
					return err
				}
			},
		},
		{
			name: "ListWlanCfgEntries",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.ListWlanCfgEntries(ctx, opts...)
					return err
				}
			},
		},
		{
			name: "ListWlanPolicies",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.ListWlanPolicies(ctx, opts...)
					return err
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, recorder := newRecordingService(t, tt.body)
			assertQueryBothDirections(t, tt.read(service), recorder)
		})
	}
}

func TestWlanServiceUnit_GetOptions_OperRouteWireQuery(t *testing.T) {
	service, recorder := newRecordingService(t,
		`{"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": {}}`)

	assertQueryBothDirections(t, func(ctx context.Context, opts ...core.GetOption) error {
		_, err := service.GetOperational(ctx, opts...)
		return err
	}, recorder)
}
