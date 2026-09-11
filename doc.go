// ANNOTATION BLOCK
// File: doc.go — package overview and the canonical usage example.
// Why it exists: this is the first thing `go doc` (and a new caller) sees. It
// explains what the package is, the isolation guarantee, the two surfaces
// (minimal ingest on *Client, full API via sub-clients), auth, and errors, then
// shows a runnable end-to-end example. No code lives here beyond the Example.

// Package auditclient is an isolated, typed Go client for the ab0t Audit Service
// (the compliance/audit-log product served at https://audit.service.ab0t.com).
//
// # Isolation
//
// This package is its own Go module (github.com/ab0t-com/ab0t-audit-go) and depends
// ONLY on the standard library. Any ab0t service can adopt it without pulling in
// audit-service internals.
//
// # Two surfaces
//
// The minimal INGEST surface lives directly on *Client for emitters that only
// need to write events, and is fail-open friendly:
//
//	c := auditclient.New("", auditclient.WithAPIKey(key))
//	_ = c.Ingest(ctx, auditclient.Event{
//		Action:    "stripe__create_invoice",
//		EventType: "tools/call",
//		Service:   "mcp-gateway",
//		Outcome:   "success",
//	})
//
// The FULL Audit Service surface — all 249 operations across 226 paths — is
// grouped into typed sub-clients hung off the same *Client, one per API
// designation:
//
//	c.Logs            // ingest + query audit events, event stats
//	c.Search          // full-text search, facets, suggestions, tool inventory
//	c.Compliance      // controls, dashboards, reports, DSAR/GDPR, SOC2, integrity
//	c.ComplianceAutomation // framework packs, judge policies, ML, risk cases, callbacks
//	c.Export          // export jobs and downloads
//	c.LegalHold       // place / release / list legal holds
//	c.AuditorShare    // evidence packages and auditor share links
//	c.Analytics       // volume, flow, geo, security, compliance analytics
//	c.Admin           // retention, purge, ad-hoc query, org summaries, status
//	c.Settings        // redaction / classification / notification preferences
//	c.Notifications   // notification feed + mark-read
//	c.SavedSearches   // saved searches
//	c.ExplorerQueries // data-explorer saved queries + history
//	c.ReportSchedules // scheduled compliance reports
//	c.RuntimeEvents   // runtime event ingest + stream
//	c.Health          // health / liveness / readiness
//	c.Capabilities    // capability discovery
//	c.ComplianceCopilot, c.ComplianceExtras, c.ComplianceTemplates,
//	c.ComplianceUI, c.UI // copilot, auditor-requests, templates, HTML pages
//
// Every method is context-first, takes typed path/query/body parameters and
// returns a typed response value (or an *APIError). Query parameters for a
// method live in a dedicated <Receiver><Method>Params struct; pass nil for
// "no filters". Path parameters are positional arguments. Request bodies are
// typed structs from types.go.
//
//	dash, err := c.Compliance.GetComplianceDashboard(ctx, &auditclient.ComplianceGetComplianceDashboardParams{
//		OrgID:     auditclient.String("org_123"),
//		Framework: auditclient.String("soc2"),
//	})
//
// # Authentication
//
// WithAPIKey sets the service key; it is sent as both `Authorization: Bearer
// <key>` and `X-API-Key: <key>` on every request, covering both accepted forms.
//
// # Errors
//
// Any non-2xx response is returned as an *APIError carrying the status, method,
// path and parsed code/message. Branch with the helpers rather than on the raw
// status:
//
//	if _, err := c.Logs.GetAuditEvent(ctx, id, nil); err != nil {
//		switch {
//		case auditclient.IsNotFound(err):    // 404
//		case auditclient.IsRateLimited(err): // 429
//		case auditclient.IsRetryable(err):   // 429 or 5xx
//		}
//	}
//
// # Pagination
//
// Offset-paginated endpoints return a PaginatedResponse_* value (Data, Total,
// HasMore, NextCursor). See CollectAll and the pagination helpers.
package auditclient

// String, Int, Bool and Float64 return pointers to their arguments. They are
// convenience constructors for the optional (pointer) fields on the *Params
// structs and on request bodies.
func String(s string) *string    { return &s }
func Int(i int) *int             { return &i }
func Bool(b bool) *bool          { return &b }
func Float64(f float64) *float64 { return &f }
