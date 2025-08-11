package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"
)

// TestStartEnvMissing covers missing env branches.
func TestStartEnvMissing(t *testing.T) {
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Unsetenv("WNC_CONTROLLER")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := start(); code != 1 {
		t.Fatalf("expected exit 1 for missing controller, got %d", code)
	}
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	if code := start(); code != 1 {
		t.Fatalf("expected exit 1 for missing token, got %d", code)
	}
}

// TestRunSuccess covers success path using stubbed fetch function.
func TestRunSuccess(t *testing.T) {
	// Spin up a minimal HTTPS server that serves valid AP oper JSON
	handler := http.NewServeMux()
	handler.HandleFunc("/restconf/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data" {
			w.Header().Set("Content-Type", "application/yang-data+json")
			fmt.Fprint(w, `{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data":{"oper-data":[{},{}]}}`)
			return
		}
		http.NotFound(w, r)
	})
	ts := httptest.NewTLSServer(handler)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	controller := u.Host

	count, err := run(controller, "abcd")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2, got %d", count)
	}
}

// TestRunFailure ensures AP oper failure propagates.
func TestRunFailure(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/restconf/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data" {
			http.Error(w, "boom", http.StatusInternalServerError)
			return
		}
		http.NotFound(w, r)
	})
	ts := httptest.NewTLSServer(handler)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	controller := u.Host

	_, err := run(controller, "abcd")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

// TestRunClientCreationError covers the error path when WithTimeout returns error (timeout <=0).
func TestRunClientCreationError(t *testing.T) {
	// Invalid controller should fail client creation fast.
	_, err := run("", "abcd")
	if err == nil {
		t.Fatalf("expected client creation error")
	}
}

// TestStartSuccess covers fully successful execution.
func TestStartSuccess(t *testing.T) {
	// Use a local server to back start() path as well.
	handler := http.NewServeMux()
	handler.HandleFunc("/restconf/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data" {
			w.Header().Set("Content-Type", "application/yang-data+json")
			fmt.Fprint(w, `{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data":{"oper-data":[{}]}}`)
			return
		}
		http.NotFound(w, r)
	})
	ts := httptest.NewTLSServer(handler)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	ctrl := u.Host

	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", ctrl)
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := start(); code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

// TestStartRunFailure covers failure after run() returns error.
func TestStartRunFailure(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/restconf/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data" {
			http.Error(w, "boom", http.StatusInternalServerError)
			return
		}
		http.NotFound(w, r)
	})
	ts := httptest.NewTLSServer(handler)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	ctrl := u.Host

	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", ctrl)
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := start(); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

// TestMainFunction covers the main() function via exitFunc injection.
func TestMainFunction(t *testing.T) {
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	codeCh := make(chan int, 1)
	exitFunc = func(code int) { codeCh <- code }

	// Back main() with a local server
	handler := http.NewServeMux()
	handler.HandleFunc("/restconf/data/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data" {
			w.Header().Set("Content-Type", "application/yang-data+json")
			fmt.Fprint(w, `{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data":{"oper-data":[{},{}]}}`)
			return
		}
		http.NotFound(w, r)
	})
	ts := httptest.NewTLSServer(handler)
	defer ts.Close()
	u, _ := url.Parse(ts.URL)
	ctrl := u.Host

	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", ctrl)
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	main()
	select {
	case code := <-codeCh:
		if code != 0 {
			t.Fatalf("expected exit code 0, got %d", code)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("main did not call exitFunc")
	}
}
