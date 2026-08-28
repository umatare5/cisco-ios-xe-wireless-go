package ap_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
)

// The bodies below are the same record twice: once with every published leaf on the wire, once
// with those leaves omitted and their containers and siblings left in place. The siblings are the
// control — they prove the container was reached, so a nil leaf is the leaf's own absence rather
// than a body the reader never decoded.
const (
	apJoinStatsEndpoint = "Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats"
	capwapDataEndpoint  = "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"
	radioOperEndpoint   = "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data"

	joinStatsAllLeavesSent = `{
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"ap-join-info": {
					"ap-name": "AP-CONTROL",
					"is-joined": true,
					"num-join-req-recvd": 11,
					"num-config-req-recvd": 12,
					"num-succ-join-resp-sent": 13,
					"num-unsucc-join-req-procn": 0,
					"num-succ-conf-resp-sent": 15,
					"num-unsucc-conf-req-procn": 16
				},
				"ap-discovery-info": {
					"last-disc-failure-type": "DISC-CONTROL",
					"num-discovery-req-recvd": 21,
					"num-succ-disc-resp-sent": 22,
					"num-err-disc-req": 0
				},
				"dtls-sess-info": {
					"mac-addr": "00:00:00:00:00:01",
					"data-dtls-setup-req": 31,
					"data-dtls-success": 32,
					"data-dtls-failure": 0,
					"ctrl-dtls-setup-req": 34,
					"ctrl-dtls-success": 35,
					"ctrl-dtls-failure": 36,
					"ctrl-dtls-decrypt-err": 37,
					"ctrl-dtls-anti-replay-err": 38,
					"data-dtls-decrypt-err": 39,
					"data-dtls-anti-replay-err": 40
				}
			}
		]
	}`

	joinStatsPublishedLeavesOmitted = `{
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"ap-join-info": {"ap-name": "AP-CONTROL"},
				"ap-discovery-info": {"last-disc-failure-type": "DISC-CONTROL"},
				"dtls-sess-info": {"mac-addr": "00:00:00:00:00:01"}
			}
		]
	}`

	capwapDataAllLeavesSent = `{
		"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"tag-info": {
					"tag-source": "TAG-CONTROL",
					"is-ap-misconfigured": true,
					"ap-misconfig": "country-misconfig"
				}
			},
			{
				"wtp-mac": "00:00:00:00:00:02",
				"tag-info": {
					"tag-source": "TAG-CONTROL",
					"is-ap-misconfigured": false,
					"ap-misconfig": "apmgr-no-misconfig"
				}
			}
		]
	}`

	capwapDataPublishedLeavesOmitted = `{
		"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"tag-info": {"tag-source": "TAG-CONTROL"}
			}
		]
	}`

	radioOperAllLeavesSent = `{
		"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"radio-slot-id": 0,
				"slot-id": 0,
				"radio-type": "radio-80211abgn",
				"radio-band-info": [
					{
						"band-id": 0,
						"phy-tx-pwr-lvl-cfg": {
							"cfg-data": {
								"num-supp-power-levels": 8,
								"tx-power-level-1": 23,
								"curr-tx-power-in-dbm": 0
							}
						}
					}
				]
			}
		]
	}`

	radioOperPublishedLeavesOmitted = `{
		"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [
			{
				"wtp-mac": "00:00:00:00:00:01",
				"radio-slot-id": 0,
				"radio-type": "radio-remote-lan",
				"radio-band-info": [
					{
						"band-id": 0,
						"phy-tx-pwr-lvl-cfg": {"cfg-data": {"num-supp-power-levels": 8}}
					}
				]
			}
		]
	}`
)

// countedLeaf is one published counter: the name it carries on the wire, the field it decodes
// into, and the value the all-leaves-sent body puts there.
type countedLeaf struct {
	name string
	got  *int
	want int
}

// countedLeaves pairs every published counter of a join-stats record with its expected value. The
// sent case and the omitted case read the same list, so neither can drift away from the other, and
// three of the values are 0 on the wire: a counter that reads zero has to stay distinguishable
// from a counter the controller never sent.
func countedLeaves(record ap.ApJoinStats) []countedLeaf {
	join, disc, dtls := record.ApJoinInfo, record.ApDiscoveryInfo, record.DTLSSessInfo

	return []countedLeaf{
		{"num-join-req-recvd", join.NumJoinReqRecvd, 11},
		{"num-config-req-recvd", join.NumConfigReqRecvd, 12},
		{"num-succ-join-resp-sent", join.NumSuccJoinRespSent, 13},
		{"num-unsucc-join-req-procn", join.NumUnsuccJoinReqProcn, 0},
		{"num-succ-conf-resp-sent", join.NumSuccConfRespSent, 15},
		{"num-unsucc-conf-req-procn", join.NumUnsuccConfReqProcn, 16},
		{"num-discovery-req-recvd", disc.NumDiscoveryReqRecvd, 21},
		{"num-succ-disc-resp-sent", disc.NumSuccDiscRespSent, 22},
		{"num-err-disc-req", disc.NumErrDiscReq, 0},
		{"data-dtls-setup-req", dtls.DataDTLSSetupReq, 31},
		{"data-dtls-success", dtls.DataDTLSSuccess, 32},
		{"data-dtls-failure", dtls.DataDTLSFailure, 0},
		{"ctrl-dtls-setup-req", dtls.CtrlDTLSSetupReq, 34},
		{"ctrl-dtls-success", dtls.CtrlDTLSSuccess, 35},
		{"ctrl-dtls-failure", dtls.CtrlDTLSFailure, 36},
		{"ctrl-dtls-decrypt-err", dtls.CtrlDTLSDecryptErr, 37},
		{"ctrl-dtls-anti-replay-err", dtls.CtrlDTLSAntiReplayErr, 38},
		{"data-dtls-decrypt-err", dtls.DataDTLSDecryptErr, 39},
		{"data-dtls-anti-replay-err", dtls.DataDTLSAntiReplayErr, 40}, //nolint:mnd // wire values
	}
}

// TestApServiceUnit_OmittedJoinLeaf_MockSuccess tests that an omitted join, discovery or DTLS leaf
// stays nil instead of decoding as a zero counter or an AP that never joined, while a leaf the
// controller does send — zero included — still decodes.
func TestApServiceUnit_OmittedJoinLeaf_MockSuccess(t *testing.T) {
	t.Run("SentLeavesDecode", func(t *testing.T) {
		record := soleJoinStatsRecord(t, joinStatsAllLeavesSent)

		if record.ApJoinInfo.ApName != "AP-CONTROL" {
			t.Fatalf("ap-name = %q, so the container was not decoded", record.ApJoinInfo.ApName)
		}

		if record.ApJoinInfo.IsJoined == nil || !*record.ApJoinInfo.IsJoined {
			t.Error("is-joined: want a non-nil true")
		}

		for _, leaf := range countedLeaves(record) {
			if leaf.got == nil {
				t.Errorf("%s: nil, want %d", leaf.name, leaf.want)
				continue
			}
			if *leaf.got != leaf.want {
				t.Errorf("%s = %d, want %d", leaf.name, *leaf.got, leaf.want)
			}
		}
	})

	t.Run("OmittedLeavesStayNil", func(t *testing.T) {
		record := soleJoinStatsRecord(t, joinStatsPublishedLeavesOmitted)

		// The siblings are the control: without them a nil leaf would only prove the reader
		// never saw the container.
		if record.ApJoinInfo.ApName != "AP-CONTROL" ||
			record.ApDiscoveryInfo.LastDiscFailureType != "DISC-CONTROL" ||
			record.DTLSSessInfo.MACAddr == "" {
			t.Fatalf("a sibling leaf is missing, so the containers were not decoded: %+v", record)
		}

		if record.ApJoinInfo.IsJoined != nil {
			t.Error("is-joined: want nil rather than an AP reported as never joined")
		}

		for _, leaf := range countedLeaves(record) {
			if leaf.got != nil {
				t.Errorf("%s = %d, want nil: an absent counter must not read as a value", leaf.name, *leaf.got)
			}
		}
	})
}

// TestApServiceUnit_OmittedMisconfiguredLeaf_MockSuccess tests that an omitted misconfiguration
// flag stays nil instead of decoding as a correctly configured AP, while an explicit false still
// decodes.
func TestApServiceUnit_OmittedMisconfiguredLeaf_MockSuccess(t *testing.T) {
	t.Run("SentLeafDecodes", func(t *testing.T) {
		records := capwapDataRecords(t, capwapDataAllLeavesSent)
		if len(records) != 2 {
			t.Fatalf("decoded %d records, want 2", len(records))
		}

		if records[0].TagInfo.IsApMisconfigured == nil || !*records[0].TagInfo.IsApMisconfigured {
			t.Error("is-ap-misconfigured: want a non-nil true")
		}

		if records[1].TagInfo.IsApMisconfigured == nil || *records[1].TagInfo.IsApMisconfigured {
			t.Error("is-ap-misconfigured: want an explicit false to decode to a non-nil false")
		}

		if records[0].TagInfo.ApMisconfig == nil || *records[0].TagInfo.ApMisconfig != ap.ApMisconfigCountry {
			t.Errorf("ap-misconfig: want a non-nil %q", ap.ApMisconfigCountry)
		}

		// The no-misconfiguration member is the one the controller sends on a healthy AP, so it
		// has to decode to a non-nil value rather than reading like an absent leaf.
		if records[1].TagInfo.ApMisconfig == nil || *records[1].TagInfo.ApMisconfig != ap.ApMisconfigNone {
			t.Errorf("ap-misconfig: want a non-nil %q", ap.ApMisconfigNone)
		}
	})

	t.Run("OmittedLeafStaysNil", func(t *testing.T) {
		records := capwapDataRecords(t, capwapDataPublishedLeavesOmitted)
		if len(records) != 1 {
			t.Fatalf("decoded %d records, want 1", len(records))
		}

		if records[0].TagInfo.TagSource != "TAG-CONTROL" {
			t.Fatalf("tag-source = %q, so tag-info was not decoded", records[0].TagInfo.TagSource)
		}

		if records[0].TagInfo.IsApMisconfigured != nil {
			t.Error("is-ap-misconfigured: want nil rather than an AP reported as correctly configured")
		}

		if records[0].TagInfo.ApMisconfig != nil {
			t.Error("ap-misconfig: want nil, which is the controller not serving the leaf")
		}
	})
}

// TestApServiceUnit_OmittedTxPowerLeaf_MockSuccess tests that an omitted TX power leaf stays nil
// instead of decoding as 0 dBm, which is a power a radio can genuinely be at.
func TestApServiceUnit_OmittedTxPowerLeaf_MockSuccess(t *testing.T) {
	t.Run("SentLeavesDecode", func(t *testing.T) {
		cfg := soleRadioBandCfgData(t, radioOperAllLeavesSent)

		if cfg.TxPowerLevel1 == nil || *cfg.TxPowerLevel1 != 23 {
			t.Errorf("tx-power-level-1 = %v, want a non-nil 23", cfg.TxPowerLevel1)
		}

		if cfg.CurrTxPowerInDbm == nil || *cfg.CurrTxPowerInDbm != 0 {
			t.Errorf("curr-tx-power-in-dbm = %v, want a non-nil 0", cfg.CurrTxPowerInDbm)
		}
	})

	t.Run("OmittedLeavesStayNil", func(t *testing.T) {
		cfg := soleRadioBandCfgData(t, radioOperPublishedLeavesOmitted)

		if cfg.NumSuppPowerLevels != 8 {
			t.Fatalf("num-supp-power-levels = %d, so cfg-data was not decoded", cfg.NumSuppPowerLevels)
		}

		if cfg.TxPowerLevel1 != nil {
			t.Errorf("tx-power-level-1 = %d, want nil rather than a fabricated dBm", *cfg.TxPowerLevel1)
		}

		if cfg.CurrTxPowerInDbm != nil {
			t.Errorf("curr-tx-power-in-dbm = %d, want nil rather than a fabricated dBm", *cfg.CurrTxPowerInDbm)
		}
	})
}

// soleJoinStatsRecord serves body from the join-stats endpoint and returns the one record it holds.
func soleJoinStatsRecord(t *testing.T, body string) ap.ApJoinStats {
	t.Helper()

	service, ctx := absenceService(t, apJoinStatsEndpoint, body)

	result, err := service.ListAPJoinStats(ctx)
	if err != nil {
		t.Fatalf("ListAPJoinStats failed: %v", err)
	}

	if len(result.ApJoinStats) != 1 {
		t.Fatalf("decoded %d records, want 1", len(result.ApJoinStats))
	}

	return result.ApJoinStats[0]
}

// capwapDataRecords serves body from the capwap-data endpoint and returns the records it holds.
func capwapDataRecords(t *testing.T, body string) []ap.CAPWAPData {
	t.Helper()

	service, ctx := absenceService(t, capwapDataEndpoint, body)

	result, err := service.ListCAPWAPData(ctx)
	if err != nil {
		t.Fatalf("ListCAPWAPData failed: %v", err)
	}

	return result.CAPWAPData
}

// soleRadioBandCfgData serves body from the radio-oper-data endpoint and returns the TX power
// configuration of the one radio band it holds.
func soleRadioBandCfgData(t *testing.T, body string) ap.PhyTxPwrLvlCfgData {
	t.Helper()

	service, ctx := absenceService(t, radioOperEndpoint, body)

	result, err := service.ListRadioData(ctx)
	if err != nil {
		t.Fatalf("ListRadioData failed: %v", err)
	}

	if len(result.RadioOperData) != 1 || len(result.RadioOperData[0].RadioBandInfo) != 1 {
		t.Fatalf("want one radio with one band, got %+v", result.RadioOperData)
	}

	return result.RadioOperData[0].RadioBandInfo[0].PhyTxPwrLvlCfg.CfgData
}

// TestApServiceUnit_OmittedSlotLeaf_MockSuccess tests that the withheld physical slot stays nil
// while the list key beside it still decodes, so a join cannot silently land on radio 0.
func TestApServiceUnit_OmittedSlotLeaf_MockSuccess(t *testing.T) {
	t.Run("SentLeafDecodes", func(t *testing.T) {
		radio := soleRadioRecord(t, radioOperAllLeavesSent)

		if radio.SlotID == nil || *radio.SlotID != 0 {
			t.Errorf("slot-id = %v, want a non-nil 0", radio.SlotID)
		}
	})

	t.Run("OmittedLeafStaysNil", func(t *testing.T) {
		radio := soleRadioRecord(t, radioOperPublishedLeavesOmitted)

		if radio.RadioType != ap.RadioTypeRemoteLAN {
			t.Fatalf("radio-type = %q, so the record was not decoded", radio.RadioType)
		}

		if radio.RadioSlotID != 0 {
			t.Errorf("radio-slot-id = %d, want the list key to keep decoding", radio.RadioSlotID)
		}

		if radio.SlotID != nil {
			t.Errorf("slot-id = %d, want nil rather than radio 0", *radio.SlotID)
		}
	})
}

// soleRadioRecord serves body from the radio-oper-data endpoint and returns the one radio it holds.
func soleRadioRecord(t *testing.T, body string) ap.RadioOperData {
	t.Helper()

	service, ctx := absenceService(t, radioOperEndpoint, body)

	result, err := service.ListRadioData(ctx)
	if err != nil {
		t.Fatalf("ListRadioData failed: %v", err)
	}

	if len(result.RadioOperData) != 1 {
		t.Fatalf("decoded %d records, want 1", len(result.RadioOperData))
	}

	return result.RadioOperData[0]
}

// absenceService stands up a server that answers endpoint alone, so a body that fails to match
// arrives as a 404 rather than as an empty record no assertion could tell from an absent leaf.
func absenceService(t *testing.T, endpoint, body string) (ap.Service, context.Context) {
	t.Helper()

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{endpoint: body}))
	t.Cleanup(mockServer.Close)

	testClient := testutil.NewTestClient(mockServer)

	return ap.NewService(testClient.Core().(*core.Client)), testutil.TestContext(t)
}
