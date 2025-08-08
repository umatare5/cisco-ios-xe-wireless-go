# 🧩 Core Services

Shortest path to common operational state.

| Service | Method | YANG Focus           | Notes                  |
| ------- | ------ | -------------------- | ---------------------- |
| General | `Oper` | Platform/env summary | Health signal          |
| Client  | `Info` | Session meta         | Library build/version  |
| Radio   | `Oper` | RF slots & status    | Channels, admin, power |
| RRM     | `Oper` | RRM state            | DCA/TPC summaries      |
| AP      | `Oper` | AP inventory         | Join status, type      |
| WLAN    | `Oper` | WLAN list            | SSIDs, status          |
| Rogue   | `Oper` | Rogue detections     | Counts only            |
| Site    | `Oper` | Hierarchy tree       | Floors / tags          |

## 🔄 Pattern

```go
radios, err := client.Radio().Oper(ctx)
```

Returned slices are zero when no data (never nil).

## 🚫 Mutations

Write (POST/PUT/PATCH/DELETE) intentionally omitted.

## 🔽 Additional (Collapsed)

<details><summary>Error & design notes</summary>

Errors wrapped with context. If future sentinel errors are added prefer `errors.Is`. Keep methods thin; heavy logic belongs outside service wrappers.

</details>
