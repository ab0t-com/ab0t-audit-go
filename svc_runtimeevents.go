// ANNOTATION BLOCK
// File: svc_runtimeevents.go — the `RuntimeEvents` designation of the Audit Service.
// Why it exists: typed methods for every RuntimeEvents endpoint, reached via
// c.RuntimeEvents.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"net/url"
)

// RuntimeEventsStreamAuditEventsParams holds the query parameters for RuntimeEvents.StreamAuditEvents.
type RuntimeEventsStreamAuditEventsParams struct {
	OrgID           *string `json:"org_id,omitempty"`            // Organization ID to scope the stream to (defaults to the caller's org).
	EventTypePrefix *string `json:"event_type_prefix,omitempty"` // Only push events whose event_type starts with this prefix (e.g. 'ai.tool').
}

func (p *RuntimeEventsStreamAuditEventsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.EventTypePrefix != nil {
		v.Set("event_type_prefix", *p.EventTypePrefix)
	}
	return v
}

// IngestRuntimeEvent calls POST /events/ingest.
// Ingest Runtime Event
//
// Fast fire-and-forget event ingestion for runtime audit events. Accepts
// any JSON with at least an event_type field. NO Pydantic validation -
// worker does all processing. **Performance:** < 2ms response time **Used
// by:** Auth service (events_v2), frontends, app logs, third-party clients
// **Table:** runtime_audit_events (separate from AI agent events)
// **Example:** ```json { "event_type": "auth.login", "org_id": "org_123",
// "user_id": "user_456", "any_other_field": "stored in event_data blob" }
// ``` **Returns:** 202 Accepted with event_id
func (s *RuntimeEventsClient) IngestRuntimeEvent(ctx context.Context) (*IngestEventResponse, error) {
	b, err := s.c.do(ctx, "POST", "/events/ingest", nil, nil)
	if err != nil {
		return nil, err
	}
	var out IngestEventResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/events/ingest", Err: err}
	}
	return &out, nil
}

// IngestRuntimeEventsBatch calls POST /events/batch.
// Ingest Runtime Events Batch
//
// Fast batch ingestion for runtime audit events. NO per-event validation -
// maximum throughput. **Performance:** Handles 1000s of events in single
// request **Used by:** High-volume event sources **Example:** ```json [
// {"event_type": "auth.login", "org_id": "org_123"}, {"event_type":
// "auth.logout", "org_id": "org_123"} ] ``` **Returns:** 202 Accepted with
// list of event_ids
func (s *RuntimeEventsClient) IngestRuntimeEventsBatch(ctx context.Context) (*IngestBatchResponse, error) {
	b, err := s.c.do(ctx, "POST", "/events/batch", nil, nil)
	if err != nil {
		return nil, err
	}
	var out IngestBatchResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/events/batch", Err: err}
	}
	return &out, nil
}

// StreamAuditEvents calls GET /events/stream.
// Stream Audit Events
//
// Live Server-Sent-Events stream of audit events as they land. **When to
// use:** Powering a real-time "Live" tail of the audit log without repeated
// full searches. Open it from the browser with `EventSource`. **Auth:**
// Requires `audit.search.read`; results are org-scoped to the caller
// (cross-tenant principals may pass `org_id` to scope to a specific org).
// **Frames:** - `event: ready` — handshake; sent once on connect. - `event:
// event` — one new audit event (same shape as /search/ rows). - `event:
// heartbeat` — periodic keep-alive when the stream is quiet. **Behavior:**
// Only events newer than connection time are emitted (the feed tails
// forward, it does not replay history). The connection stays open until the
// client disconnects.
func (s *RuntimeEventsClient) StreamAuditEvents(ctx context.Context, params *RuntimeEventsStreamAuditEventsParams) (RawMessage, error) {
	b, err := s.c.do(ctx, "GET", "/events/stream", params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}
