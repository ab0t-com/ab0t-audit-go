// ANNOTATION BLOCK
// File: svc_logs.go — the `Logs` designation of the Audit Service.
// Why it exists: typed methods for every Logs endpoint, reached via
// c.Logs.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// LogsGetAuditEventParams holds the query parameters for Logs.GetAuditEvent.
type LogsGetAuditEventParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *LogsGetAuditEventParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// LogsGetAuditEventsParams holds the query parameters for Logs.GetAuditEvents.
type LogsGetAuditEventsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID to query
	StartTime *string `json:"start_time,omitempty"` // Filter events after this time
	EndTime   *string `json:"end_time,omitempty"`   // Filter events before this time
	EventType *string `json:"event_type,omitempty"` // Filter by event type (e.g. 'user.login')
	UserID    *string `json:"user_id,omitempty"`    // Filter by specific user
	Action    *string `json:"action,omitempty"`     // Filter by action (e.g. 'create', 'delete')
	Outcome   *string `json:"outcome,omitempty"`    // Filter by outcome ('success' or 'failure')
	Limit     *int    `json:"limit,omitempty"`      // Max events to return (1-1000)
	Offset    *int    `json:"offset,omitempty"`     // Skip first N events (offset pagination)
	Cursor    *string `json:"cursor,omitempty"`     // Cursor for keyset pagination (from next_cursor in previous response). When provided, of...
}

func (p *LogsGetAuditEventsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.StartTime != nil {
		v.Set("start_time", *p.StartTime)
	}
	if p.EndTime != nil {
		v.Set("end_time", *p.EndTime)
	}
	if p.EventType != nil {
		v.Set("event_type", *p.EventType)
	}
	if p.UserID != nil {
		v.Set("user_id", *p.UserID)
	}
	if p.Action != nil {
		v.Set("action", *p.Action)
	}
	if p.Outcome != nil {
		v.Set("outcome", *p.Outcome)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", strconv.Itoa(*p.Offset))
	}
	if p.Cursor != nil {
		v.Set("cursor", *p.Cursor)
	}
	return v
}

// LogsGetAuditEventsCountParams holds the query parameters for Logs.GetAuditEventsCount.
type LogsGetAuditEventsCountParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // Count events after this time
	EndTime   *string `json:"end_time,omitempty"`   // Count events before this time
	EventType *string `json:"event_type,omitempty"` // Count only this event type
}

func (p *LogsGetAuditEventsCountParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.StartTime != nil {
		v.Set("start_time", *p.StartTime)
	}
	if p.EndTime != nil {
		v.Set("end_time", *p.EndTime)
	}
	if p.EventType != nil {
		v.Set("event_type", *p.EventType)
	}
	return v
}

// LogsIngestStreamEndpointParams holds the query parameters for Logs.IngestStreamEndpoint.
type LogsIngestStreamEndpointParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID for all events in stream
}

func (p *LogsIngestStreamEndpointParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// GetAuditEvent calls GET /logs/events/{event_id}.
// Get Audit Event
//
// Retrieve complete details of a single audit event by its ID. **When to
// use**: After receiving an event_id from POST /logs/ingest, or to get full
// details from search results. **Returns**: Full event record including
// request/response data and metadata. **Next step**: Use the event data for
// investigation, compliance reports, or incident response.
func (s *LogsClient) GetAuditEvent(ctx context.Context, eventID string, params *LogsGetAuditEventParams) (*AuditEvent, error) {
	path := fmt.Sprintf("/logs/events/%s", url.PathEscape(eventID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out AuditEvent
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetAuditEvents calls GET /logs/events.
// Get Audit Events
//
// Search and retrieve audit events with filters. Use this to find specific
// events or review activity. **When to use**: Investigating user activity,
// reviewing recent changes, troubleshooting issues. **Returns**: Array of
// matching events, newest first. Supports offset and cursor pagination.
// **Cursor pagination**: Pass `cursor` from `next_cursor` in the previous
// response for efficient paging through large result sets. **Next step**:
// To get full details of one event, use GET /logs/events/{event_id}.
// **Tip**: Combine filters (e.g. user_id + time range) to narrow results.
func (s *LogsClient) GetAuditEvents(ctx context.Context, params *LogsGetAuditEventsParams) (*PaginatedResponseAuditEvent, error) {
	b, err := s.c.do(ctx, "GET", "/logs/events", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseAuditEvent
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/logs/events", Err: err}
	}
	return &out, nil
}

// GetAuditEventsCount calls GET /logs/stats/events.
// Get Audit Events Count
//
// Get total count of audit events matching your filters without fetching
// full records. **When to use**: Check how many events exist before
// fetching them, or for dashboard metrics. **Returns**: Total number of
// matching events. **Next step**: If count is large, use GET /logs/events
// with pagination to retrieve actual events. **Tip**: Use same filters here
// as you'll use in GET /logs/events to preview result size.
func (s *LogsClient) GetAuditEventsCount(ctx context.Context, params *LogsGetAuditEventsCountParams) (*EventStatsResponse, error) {
	b, err := s.c.do(ctx, "GET", "/logs/stats/events", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out EventStatsResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/logs/stats/events", Err: err}
	}
	return &out, nil
}

// IngestAuditEvent calls POST /logs/ingest.
// Ingest Audit Event
//
// Record a single audit event. Use this to log user actions, system events,
// or security incidents. **When to use**: After any action that needs audit
// trail (user login, data access, config change). **Returns**: Confirmation
// with event_id. Event is processed asynchronously. **Next step**: To find
// this event later, use POST /logs/search with the event_id or filters.
func (s *LogsClient) IngestAuditEvent(ctx context.Context, body AuditEventCreate) (*IngestLogResponse, error) {
	b, err := s.c.do(ctx, "POST", "/logs/ingest", nil, body)
	if err != nil {
		return nil, err
	}
	var out IngestLogResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/logs/ingest", Err: err}
	}
	return &out, nil
}

// IngestAuditEventsBatch calls POST /logs/ingest/batch.
// Ingest Audit Events Batch
//
// Record multiple audit events in one request. Use this for bulk event
// ingestion. **When to use**: Logging multiple user actions, batch imports,
// or high-volume activity. **Returns**: Immediately with event IDs while
// events are processed asynchronously. **Next step**: Events are
// automatically stored. No further action needed.
func (s *LogsClient) IngestAuditEventsBatch(ctx context.Context, body []AuditEventCreate) (*IngestLogBatchResponse, error) {
	b, err := s.c.do(ctx, "POST", "/logs/ingest/batch", nil, body)
	if err != nil {
		return nil, err
	}
	var out IngestLogBatchResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/logs/ingest/batch", Err: err}
	}
	return &out, nil
}

// IngestStreamEndpoint calls POST /logs/ingest/stream.
// Ingest Stream Endpoint
//
// Stream large volumes of events in one HTTP request. Accepts
// newline-delimited JSON (JSONL). **When to use**: Importing historical
// data, migrating from another system, or bulk uploading logs. **Format**:
// Each line must be a complete JSON object. Use Content-Type:
// application/jsonl. **Returns**: Summary with count, processing rate, and
// any errors encountered. **Next step**: Events are queued for processing.
// No further action needed. **Performance**: Can handle thousands of events
// per second.
func (s *LogsClient) IngestStreamEndpoint(ctx context.Context, params *LogsIngestStreamEndpointParams) (*IngestStreamResponse, error) {
	b, err := s.c.do(ctx, "POST", "/logs/ingest/stream", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out IngestStreamResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/logs/ingest/stream", Err: err}
	}
	return &out, nil
}
