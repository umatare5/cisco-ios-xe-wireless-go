package core

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetGenericSuccess covers the happy path of the generic Get helper.
func TestGetGenericSuccess(t *testing.T) {
	// minimal test server returning JSON
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// Expect path /restconf/data/test due to builder
		if !strings.HasSuffix(r.URL.Path, "/test") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"value":123}`))
	}))
	defer ts.Close()

	host := strings.TrimPrefix(strings.TrimPrefix(ts.URL, "https://"), "http://")
	c, err := New(host, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	type resp struct {
		Value int `json:"value"`
	}
	out, err := Get[resp](context.Background(), c, "/test")
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
	if out == nil || out.Value != 123 {
		t.Fatalf("unexpected response: %+v", out)
	}
}

// TestGetGenericError covers underlying Do returning an error (404 JSON parse ok still counts as success path vs error HTTP).
func TestGetGenericError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error":"x"}`))
	}))
	defer ts.Close()

	host := strings.TrimPrefix(strings.TrimPrefix(ts.URL, "https://"), "http://")
	c, err := New(host, "token", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	type resp struct {
		Value int `json:"value"`
	}
	out, err := Get[resp](context.Background(), c, "/err")
	if err == nil { // expect HTTP error
		t.Fatalf("expected error, got nil (%+v)", out)
	}
}
