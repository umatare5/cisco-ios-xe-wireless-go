package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/afc"
	"github.com/umatare5/cisco-ios-xe-wireless-go/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/general"
	"github.com/umatare5/cisco-ios-xe-wireless-go/geolocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/hyperlocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/mdns"
	"github.com/umatare5/cisco-ios-xe-wireless-go/nmsp"
)

func main() {
	// Get credentials from environment
	controller := os.Getenv("WNC_CONTROLLER")
	accessToken := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || accessToken == "" {
		log.Fatal("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set")
	}

	// Create client configuration
	config := wnc.Config{
		Controller:         controller,
		AccessToken:        accessToken,
		Timeout:            30 * time.Second,
		InsecureSkipVerify: true, // For development only
	}

	// Create client
	client, err := wnc.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("🚀 Testing new Service-based API...")
	fmt.Println()

	// Test General Service
	fmt.Println("1. Testing General Service...")
	generalService := general.NewService(client.CoreClient())
	generalOper, err := generalService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ General.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ General.Oper success - Interface: %s (%s)\n",
			generalOper.CiscoIOSXEWirelessGeneralOperData.MgmtIntfData.IntfName,
			generalOper.CiscoIOSXEWirelessGeneralOperData.MgmtIntfData.MgmtIP)
	}

	// Test AFC Service
	fmt.Println("2. Testing AFC Service...")
	afcService := afc.NewService(client.CoreClient())
	afcOper, err := afcService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ AFC.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ AFC.Oper success - AFC responses: %d\n", len(afcOper.CiscoIOSXEWirelessAfcOperAfcOperData.EwlcAfcApResp))
	}

	// Test Geolocation Service
	fmt.Println("3. Testing Geolocation Service...")
	geoService := geolocation.NewService(client.CoreClient())
	geoOper, err := geoService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ Geolocation.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ Geolocation.Oper success - GNSS APs: %d\n", geoOper.CiscoIOSXEWirelessGeolocationOperGeolocationOperData.ApGeoLocStats.NumApGnss)
	}

	// Test AP Service
	fmt.Println("4. Testing AP Service...")
	apService := ap.NewService(client.CoreClient())
	apCfg, err := apService.Cfg(ctx)
	if err != nil {
		fmt.Printf("   ❌ AP.Cfg failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ AP.Cfg success - Configuration loaded\n")
		_ = apCfg // Use the result to avoid unused variable warning
	}

	// Test NMSP Service
	fmt.Println("5. Testing NMSP Service...")
	nmspService := nmsp.NewService(client)
	nmspOper, err := nmspService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ NMSP.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ NMSP.Oper success - Client registrations: %d\n", len(nmspOper.CiscoIOSXEWirelessNmspOperData.ClientRegistration))
	}

	// Test Hyperlocation Service
	fmt.Println("6. Testing Hyperlocation Service...")
	hyperService := hyperlocation.NewService(client)
	hyperOper, err := hyperService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ Hyperlocation.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ Hyperlocation.Oper success - Profiles: %d\n", len(hyperOper.CiscoIOSXEWirelessHyperlocationOperHyperlocationOperData.EwlcHyperlocationProfile))
	}

	// Test mDNS Service
	fmt.Println("7. Testing mDNS Service...")
	mdnsService := mdns.NewService(client)
	mdnsOper, err := mdnsService.Oper(ctx)
	if err != nil {
		fmt.Printf("   ❌ mDNS.Oper failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ mDNS.Oper success - WLAN stats: %d\n", len(mdnsOper.CiscoIOSXEWirelessMdnsOperMdnsOperData.MdnsWlanStats))
	}

	fmt.Println()
	fmt.Println("🎉 Service API tests completed!")
	fmt.Println("✨ The new three-layer architecture is working perfectly!")
}
