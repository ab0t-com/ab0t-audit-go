# auditclient

An isolated, typed, **stdlib-only** Go client for the **ab0t Audit Service** — the
compliance / audit-log product served at `https://audit.service.ab0t.com`.

It covers the **entire** service contract: **all 249 operations** (226 paths) and
**all 215 schemas** from the OpenAPI spec. The minimal ingest surface used by
event emitters is preserved on the root `*Client`; the full product surface is
grouped into typed sub-clients.

## Install

```go
import "github.com/ab0t-com/ab0t-audit-go"
```

This is a self-contained module (`github.com/ab0t-com/ab0t-audit-go`) with **no
third-party dependencies**. To vendor it locally, add a replace directive:

```
require github.com/ab0t-com/ab0t-audit-go v0.0.0
replace github.com/ab0t-com/ab0t-audit-go => ../shared/auditclient
```

## Shape

```go
c := auditclient.New("", auditclient.WithAPIKey(os.Getenv("AUDIT_API_KEY")))
```

`New(baseURL, ...Option)` — empty `baseURL` uses `DefaultBaseURL`. Options:
`WithAPIKey(key)` (sent as both `Authorization: Bearer` and `X-API-Key`) and
`WithHTTPClient(h)`.

Every method is **context-first**, takes typed path/query/body params, and
returns a typed value or an `*APIError`. Query params live in a per-method
`<Receiver><Method>Params` struct — pass `nil` for no filters. Use the
`String/Int/Bool/Float64` helpers for the optional pointer fields.

Sub-clients (one per designation): `Logs`, `Search`, `Compliance`,
`ComplianceAutomation`, `Export`, `LegalHold`, `AuditorShare`, `Analytics`,
`Admin`, `Settings`, `Notifications`, `SavedSearches`, `ExplorerQueries`,
`ReportSchedules`, `RuntimeEvents`, `Health`, `Capabilities`, `ComplianceCopilot`,
`ComplianceExtras`, `ComplianceTemplates`, `ComplianceUI`, `UI`.

## Examples

### 1. Ingest an event (emitter surface)

```go
// Minimal, fail-open friendly — on the root client:
err := c.Ingest(ctx, auditclient.Event{
    Action: "stripe__create_invoice", EventType: "tools/call",
    Service: "mcp-gateway", Outcome: "success",
})

// Or the full typed ingest, with a typed 202 response:
resp, err := c.Logs.IngestAuditEvent(ctx, auditclient.AuditEventCreate{
    Action: "user.login", EventType: "auth", Service: "auth-service",
})
_ = resp.EventID
```

### 2. Query & search audit events (logs + search)

```go
page, err := c.Logs.GetAuditEvents(ctx, &auditclient.LogsGetAuditEventsParams{
    OrgID:  auditclient.String("org_123"),
    Outcome: auditclient.String("failure"),
    Limit:  auditclient.Int(100),
})

res, err := c.Search.SearchAuditEvents(ctx, auditclient.SearchQuery{
    Query:    "login",
    Services: []string{"auth-service"},
    Limit:    auditclient.Int(50),
})
_ = res.Total
```

### 3. Compliance dashboards, exports & legal holds

```go
dash, err := c.Compliance.GetComplianceDashboard(ctx, &auditclient.ComplianceGetComplianceDashboardParams{
    OrgID: auditclient.String("org_123"), Framework: auditclient.String("soc2"),
})
_ = dash.ComplianceScore

job, err := c.Export.CreateExportJob(ctx, auditclient.ExportRequestInput{ /* ... */ })

_, err = c.LegalHold.PlaceLegalHold(ctx, auditclient.LegalHoldCreate{ /* ... */ }, nil)
```

### 4. Errors & pagination

```go
if _, err := c.Logs.GetAuditEvent(ctx, id, nil); err != nil {
    switch {
    case auditclient.IsNotFound(err):    // 404
    case auditclient.IsRateLimited(err): // 429
    case auditclient.IsRetryable(err):   // 429 or 5xx
    }
    if ae, ok := auditclient.AsAPIError(err); ok {
        log.Printf("audit %d: %s", ae.StatusCode, ae.Message)
    }
}

// Walk an offset-paginated endpoint to completion:
all, err := auditclient.CollectAll(ctx, func(ctx context.Context, offset int) (auditclient.Page[auditclient.AuditEvent], error) {
    r, err := c.Logs.GetAuditEvents(ctx, &auditclient.LogsGetAuditEventsParams{Limit: auditclient.Int(500), Offset: auditclient.Int(offset)})
    if err != nil { return auditclient.Page[auditclient.AuditEvent]{}, err }
    return auditclient.Page[auditclient.AuditEvent]{Data: r.Data, Total: r.Total, HasMore: r.HasMore}, nil
})
```

## Layout

| File | Contents |
|------|----------|
| `auditclient.go` | Root `*Client`, `New`, options, sub-client fields, legacy `Ingest`/`IngestBatch`/`Event` |
| `services.go` | Sub-client types, `initServices`, the shared `do()` round-trip |
| `types.go` | All 215 schemas (structs, enums, `PaginatedResponse_*` / `ListEnvelope_*`) |
| `svc_<designation>.go` | Typed methods + `<Method>Params` for each of the 22 designations |
| `misc.go` | Root / `/metrics` / `/llm.txt` (untagged) methods |
| `errors.go` | `APIError`, `EncodeError`, `DecodeError`, `Is*` helpers |
| `pagination.go` | `CollectAll` / `CollectCursor` generic drivers |
| `doc.go` | Package overview + pointer helpers (`String`/`Int`/`Bool`/`Float64`) |
| `client_test.go` | `httptest` wiring tests across designations |

## Notes

- Endpoints documented as returning `any` in the spec return `RawMessage`
  (`= json.RawMessage`) — unmarshal into a concrete type yourself.
- HTML `ui/*` pages return `string`.
- Date-time fields are typed as `string` (RFC 3339) to remain robust against
  nullability; parse with `time.Parse` as needed.
- Coverage is **complete** — no designation is partial.
