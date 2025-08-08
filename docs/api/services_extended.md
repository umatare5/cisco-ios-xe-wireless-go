# 🧮 Extended Services

Less frequent operational domains.

| Service | Method | Domain |
|---------|--------|--------|
| AFC | `Oper` | CleanAir / RF analytics |
| APF | `Oper` | AP filter groups |
| APF (Flex) | `Oper` | (If model present) |
| APF? | - | Placeholder removal if absent |
| AWIPS | `Oper` | IPS wireless |
| BLE | `Oper` | BLE beacons |
| CTS | `Oper` | TrustSec |
| DOT11 | `Oper` | 802.11 stats |
| DOT15 | `Oper` | 802.15 data |
| FABRIC | `Oper` | SD‑Access fabric |
| FLEX | `Oper` | Flex profiles |
| GEOLOCATION | `Oper` | Location data |
| HYPERLOCATION | `Oper` | Fine‑grain location |
| LISP | `Oper` | LISP state |
| LOCATION | `Oper` | General location |
| MCAST | `Oper` | Multicast |
| MDNS | `Oper` | mDNS services |
| MESH | `Oper` | Mesh links |
| MOBILITY | `Oper` | Mobility groups |
| NMSP | `Oper` | NMSP feeds |
| RF | `Oper` | RF summary |
| RFID | `Oper` | RFID tags |
| ROGUE | `Oper` | Expanded rogue data |
| RRM | `Oper` | (Also core) advanced detail |
| SITE | `Oper` | (Also core) extended tree |

## 🔍 Usage

```go
mesh, err := client.Mesh().Oper(ctx)
```

## 🧪 Testing

Prefer focused unit tests using canned JSON under `*/test_data/`.
