# 📚 API Reference (Stub)

Full content moved to modular docs under `docs/api/`.

| Area | File |
|------|------|
| Overview | `docs/api/README.md` |
| Core services | `docs/api/services_core.md` |
| Extended services | `docs/api/services_extended.md` |

## Sample

```go
client, _ := wnc.NewClient(host, token)
ctx := context.Background()
resp, _ := client.General().Oper(ctx)
_ = resp
```

## Notes

- Simple GET → internal `core.Get[T]`
- Strongly typed YANG models
- Reuse a single client instance

## See Also

`docs/testing/` · `docs/security/` · `docs/scripts/`
