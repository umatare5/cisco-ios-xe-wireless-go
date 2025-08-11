package tests

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRESTCONFServersHelpers(t *testing.T) {
	okSrv := NewRESTCONFSuccessServer(map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data": `{"ap-oper": "ok"}`,
	})
	defer okSrv.Close()
	badSrv := NewRESTCONFErrorServer(
		[]string{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"},
		http.StatusInternalServerError,
	)
	defer badSrv.Close()

	// Success client
	cOK := okSrv.Client()
	cOK.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402 - test env
	resp, err := cOK.Get(okSrv.URL + "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	t.Cleanup(func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	})
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	_ = json.NewDecoder

	// Error client
	cErr := badSrv.Client()
	cErr.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402 - test env
	resp2, err := cErr.Get(badSrv.URL + "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	t.Cleanup(func() {
		if resp2 != nil && resp2.Body != nil {
			_ = resp2.Body.Close()
		}
	})
	if resp2.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp2.StatusCode)
	}
}

// TestRESTCONFServersHelpers_NotFoundBranches covers the non-matching path (404) branches
// for both success and error servers, and verifies Content-Type on success path.
func TestRESTCONFServersHelpers_NotFoundBranches(t *testing.T) {
	okSrv := NewRESTCONFSuccessServer(map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": `{"Cisco-IOS-XE-wireless-access-point-oper:capwap-data":[]}`,
	})
	defer okSrv.Close()

	badSrv := NewRESTCONFErrorServer(
		[]string{"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"},
		http.StatusBadGateway,
	)
	defer badSrv.Close()

	// Success server: hit defined endpoint (200) and undefined endpoint (404)
	c1 := okSrv.Client()
	c1.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402 - test env
	respOK, err := c1.Get(okSrv.URL + "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:capwap-data")
	if err != nil {
		t.Fatalf("ok request failed: %v", err)
	}
	if ct := respOK.Header.Get("Content-Type"); ct == "" {
		t.Errorf("expected Content-Type to be set, got empty")
	}
	_ = respOK.Body.Close()

	respNF, err := c1.Get(okSrv.URL + "/restconf/data/unknown")
	if err != nil {
		t.Fatalf("nf request failed: %v", err)
	}
	if respNF.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", respNF.StatusCode)
	}
	_ = respNF.Body.Close()

	// Error server: defined endpoint -> provided status, undefined -> 404
	c2 := badSrv.Client()
	c2.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402 - test env
	respErr, err := c2.Get(badSrv.URL + "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map")
	if err != nil {
		t.Fatalf("err request failed: %v", err)
	}
	if respErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("expected 502, got %d", respErr.StatusCode)
	}
	_ = respErr.Body.Close()

	respNF2, err := c2.Get(badSrv.URL + "/restconf/data/unknown")
	if err != nil {
		t.Fatalf("nf2 request failed: %v", err)
	}
	if respNF2.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", respNF2.StatusCode)
	}
	_ = respNF2.Body.Close()
}
