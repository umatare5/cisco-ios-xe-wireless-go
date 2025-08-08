# 🧪 Unit Test Details

## Goals

- Validate constructors
- Exercise service wrappers
- Ensure JSON tags
- Edge errors (nil client, nil ctx)

## Helpers

`internal/tests` provides:

| Helper | Purpose |
|--------|---------|
| `TestClient(t)` | Synthetic client |
| `TestContext(t)` | Deadline ctx |
| `SaveTestDataToFile` | Persist samples |

## Pattern

```go
func TestRadioOper(t *testing.T) {
    c := tests.TestClient(t)
    ctx := context.Background()
    got, err := c.Radio().Oper(ctx)
    if err != nil { t.Fatalf("oper: %v", err) }
    if got == nil { t.Fatal("nil resp") }
}
```

## Nil Context

```go
var nilCtx context.Context
_, err := c.Radio().Oper(nilCtx)
if err == nil { t.Fatal("expected err") }
```

## Table Tests

Prefer consistent naming:

```go
cases := []struct{ name string; fn func(context.Context) error }{
  {"oper", func(ctx context.Context) error { _, e := s.Oper(ctx); return e }},
}
for _, tc := range cases {
  t.Run(tc.name, func(t *testing.T){ if err := tc.fn(ctx); err!=nil { t.Fatal(err) } })
}
```

## Fail Fast

Use `t.Fatalf` for setup errors, `t.Errorf` for per‑assert issues.

## Marshalling

Round‑trip where useful:

```go
b, _ := json.Marshal(got)
if !bytes.Contains(b, []byte("radio")) { t.Log("sanity") }
```

## Skips

Integration‑only logic MUST NOT silently skip; unit tests never depend on env.
