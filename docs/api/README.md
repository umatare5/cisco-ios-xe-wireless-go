# 📚 API Overview

Unified client exposes typed wireless domain services (read‑only GET).

## 🧱 Layers

1. Client (transport, TLS)
2. Services (thin wrappers)
3. Models (YANG‑aligned)

## 🔧 Client

```go
client, _ := wnc.NewClient(host, token, wnc.WithTimeout(30*time.Second))
ctx := context.Background()
_, _ = client.General().Oper(ctx)
```

Options: timeout, insecure (dev), logger.

## 🔁 Generic Helper

`core.Get[T]` centralizes GET + JSON decode.

## 📦 Groups

Core services: `services_core.md`
Extended: `services_extended.md`

## 📑 Models

Pointers mark optional leaves. Always nil‑check.

## 🛡 Auth

Base64 `user:pass` token. TLS verify ON by default.

## 🕒 Context

Use per‑call deadline contexts for cancellation.

## � Additional (Collapsed)

<details><summary>Notes</summary>

Reuse a single client (connection reuse). Avoid disabling TLS outside dev. Check nil pointers for optional YANG leaves. Keep service methods thin.

</details>

## �🔗 Related

Testing: `../testing/` · Security: `../security/` · Scripts: `../scripts/`
