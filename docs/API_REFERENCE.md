# 📚 API Reference

Primary detail lives in `docs/api/`. This page is a fast map.

| Topic | Path |
|-------|------|
| Overview | `docs/api/README.md` |
| Core services | `docs/api/services_core.md` |
| Extended services | `docs/api/services_extended.md` |

## 🔧 Minimal Use

```go
client, _ := wnc.NewClient(host, token)
ctx := context.Background()
resp, _ := client.General().Oper(ctx)
_ = resp
```

## 🔍 Notes (Collapsed)

<details><summary>Design details</summary>

All simple GETs route through `core.Get[T]` for HTTP + JSON decode. Models mirror YANG; optional leaves are pointers. Reuse a single client instance for connection reuse.

</details>

## 🔗 Related

`docs/testing/` · `docs/security/` · `docs/scripts/`
