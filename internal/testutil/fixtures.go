package testutil

// Canonical fixture identities.
//
// Every JSON fixture in this repository draws its MAC addresses from a single
// locally administered block, aa:bb:cc:dd:ee:00/40. The last octet encodes the
// role of the device, so a fixture can be read without cross-referencing the
// response it came from:
//
//	:00        wireless LAN controller management interface
//	:01 - :0f  access point radio base MAC   (TEST-APnn -> :nn)
//	:11 - :1f  access point Ethernet MAC     (TEST-APnn -> :1n)
//	:a1 - :af  associated client stations
//	:b1 - :bf  BSSIDs advertised by fixture access points
//	:c1 - :cf  RRM grouping cluster identifiers
//	:f1 - :ff  rogue and otherwise unmanaged devices
//
// The block satisfies both bits the 802 address format defines: the first octet
// 0xaa has the I/G bit clear (unicast) and the U/L bit set (locally
// administered), so no fixture can collide with a real vendor assignment.
const (
	// TestWLCMgmtMAC is the management interface MAC of the fixture controller.
	TestWLCMgmtMAC = "aa:bb:cc:dd:ee:00"

	// TestAP01RadioMAC is the radio base MAC of TEST-AP01.
	TestAP01RadioMAC = "aa:bb:cc:dd:ee:01"

	// TestAP02RadioMAC is the radio base MAC of TEST-AP02.
	TestAP02RadioMAC = "aa:bb:cc:dd:ee:02"

	// TestAP03RadioMAC is the radio base MAC of TEST-AP03.
	TestAP03RadioMAC = "aa:bb:cc:dd:ee:03"

	// TestAP01EthernetMAC is the wired MAC of TEST-AP01.
	TestAP01EthernetMAC = "aa:bb:cc:dd:ee:11"

	// TestAP02EthernetMAC is the wired MAC of TEST-AP02.
	TestAP02EthernetMAC = "aa:bb:cc:dd:ee:12"

	// TestClient01MAC is the station MAC of the first associated client.
	TestClient01MAC = "aa:bb:cc:dd:ee:a1"

	// TestClient02MAC is the station MAC of the second associated client.
	TestClient02MAC = "aa:bb:cc:dd:ee:a2"

	// TestBSSID is the BSSID advertised by a fixture access point radio.
	TestBSSID = "aa:bb:cc:dd:ee:b1"

	// TestClusterID is the RRM grouping cluster identifier.
	TestClusterID = "aa:bb:cc:dd:ee:c1"

	// TestRogueAPMAC is the MAC of a rogue access point.
	TestRogueAPMAC = "aa:bb:cc:dd:ee:f1"

	// TestRogueClientMAC is the MAC of a client associated to a rogue access point.
	TestRogueClientMAC = "aa:bb:cc:dd:ee:f2"

	// TestSpectrumDeviceID identifies a non-Wi-Fi interferer seen by CleanAir.
	TestSpectrumDeviceID = "aa:bb:cc:dd:ee:f3"
)

// Canonical fixture host identities.
const (
	// TestAP01Name and its siblings name the access points owned by the fixtures.
	TestAP01Name = "TEST-AP01"
	TestAP02Name = "TEST-AP02"
	TestAP03Name = "TEST-AP03"

	// TestWLCName is the hostname reported by the fixture controller.
	TestWLCName = "wnc1"

	// TestWLCFQDN is the address every fixture and example dials.
	TestWLCFQDN = "wnc1.example.internal"

	// TestWLCMgmtIP is the management address of the fixture controller.
	TestWLCMgmtIP = "192.168.1.10"

	// TestAP01IP and TestAP02IP are the addresses of the fixture access points.
	TestAP01IP = "192.168.1.11"
	TestAP02IP = "192.168.1.12"
)
