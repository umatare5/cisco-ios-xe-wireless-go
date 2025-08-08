# 📚 API Reference

Fast map. Full detail under `docs/api/`.

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

## � Additional (Collapsed)

<details><summary>Design notes</summary>

`core.Get[T]` centralizes HTTP + JSON decode. Models mirror YANG; optional leaves are pointers (nil‑check). Reuse a single client instance.

</details>

## 🔗 Related

`docs/testing/` · `docs/security/` · `docs/scripts/`
