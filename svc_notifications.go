// ANNOTATION BLOCK
// File: svc_notifications.go — the `Notifications` designation of the Audit Service.
// Why it exists: typed methods for every Notifications endpoint, reached via
// c.Notifications.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// NotificationsGetNotificationsParams holds the query parameters for Notifications.GetNotifications.
type NotificationsGetNotificationsParams struct {
	OrgID               *string `json:"org_id,omitempty"`                // Organization ID (defaults to caller's org)
	LookbackDays        *int    `json:"lookback_days,omitempty"`         // How far back to look for report failures / security alerts
	DeadlineHorizonDays *int    `json:"deadline_horizon_days,omitempty"` // Surface DSARs whose deadline falls within this window
	Limit               *int    `json:"limit,omitempty"`                 // Max items to return across all sources
}

func (p *NotificationsGetNotificationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.LookbackDays != nil {
		v.Set("lookback_days", strconv.Itoa(*p.LookbackDays))
	}
	if p.DeadlineHorizonDays != nil {
		v.Set("deadline_horizon_days", strconv.Itoa(*p.DeadlineHorizonDays))
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// NotificationsGetNotifications2Params holds the query parameters for Notifications.GetNotifications2.
type NotificationsGetNotifications2Params struct {
	OrgID               *string `json:"org_id,omitempty"`                // Organization ID (defaults to caller's org)
	LookbackDays        *int    `json:"lookback_days,omitempty"`         // How far back to look for report failures / security alerts
	DeadlineHorizonDays *int    `json:"deadline_horizon_days,omitempty"` // Surface DSARs whose deadline falls within this window
	Limit               *int    `json:"limit,omitempty"`                 // Max items to return across all sources
}

func (p *NotificationsGetNotifications2Params) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.LookbackDays != nil {
		v.Set("lookback_days", strconv.Itoa(*p.LookbackDays))
	}
	if p.DeadlineHorizonDays != nil {
		v.Set("deadline_horizon_days", strconv.Itoa(*p.DeadlineHorizonDays))
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// NotificationsMarkNotificationsReadParams holds the query parameters for Notifications.MarkNotificationsRead.
type NotificationsMarkNotificationsReadParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID (defaults to caller's org)
}

func (p *NotificationsMarkNotificationsReadParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// GetNotifications calls GET /notifications/.
// Get Notifications
//
// Aggregate the audit service's own actionable signals into the header bell
// feed. **When to use**: the global header Notifications bell on every app
// page. Renders a real in-app panel of items that need the user's
// attention; the unread dot reflects `unread_count` (hidden when zero).
// **Returns**: NotificationFeed — `unread_count` plus a `items[]` list of
// {id, kind, title, detail, severity, created_at, deep_link},
// severity-then-recency ordered (critical first), capped at `limit`.
// **Sources** (all org-scoped, all from this service's own data): -
// `report_failed` failed/unsupported compliance report jobs (last
// `lookback_days`) - `auditor_request` external-auditor requests still in
// the `new` state - `dsar_deadline` open DSARs within
// `deadline_horizon_days` of (or past) their deadline - `security_alert`
// recent auth-failure / admin / security-violation events **Auth**:
// AuditUIUser (audit.ui.read). Org-scoped to the caller. **Resilience**: a
// failure in any single source is logged and skipped, so the feed never
// fails wholesale; an empty feed renders the panel's empty-state.
func (s *NotificationsClient) GetNotifications(ctx context.Context, params *NotificationsGetNotificationsParams) (*NotificationFeed, error) {
	b, err := s.c.do(ctx, "GET", "/notifications/", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out NotificationFeed
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/notifications/", Err: err}
	}
	return &out, nil
}

// GetNotifications2 calls GET /notifications.
// Get Notifications
//
// Aggregate the audit service's own actionable signals into the header bell
// feed. **When to use**: the global header Notifications bell on every app
// page. Renders a real in-app panel of items that need the user's
// attention; the unread dot reflects `unread_count` (hidden when zero).
// **Returns**: NotificationFeed — `unread_count` plus a `items[]` list of
// {id, kind, title, detail, severity, created_at, deep_link},
// severity-then-recency ordered (critical first), capped at `limit`.
// **Sources** (all org-scoped, all from this service's own data): -
// `report_failed` failed/unsupported compliance report jobs (last
// `lookback_days`) - `auditor_request` external-auditor requests still in
// the `new` state - `dsar_deadline` open DSARs within
// `deadline_horizon_days` of (or past) their deadline - `security_alert`
// recent auth-failure / admin / security-violation events **Auth**:
// AuditUIUser (audit.ui.read). Org-scoped to the caller. **Resilience**: a
// failure in any single source is logged and skipped, so the feed never
// fails wholesale; an empty feed renders the panel's empty-state.
func (s *NotificationsClient) GetNotifications2(ctx context.Context, params *NotificationsGetNotifications2Params) (*NotificationFeed, error) {
	b, err := s.c.do(ctx, "GET", "/notifications", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out NotificationFeed
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/notifications", Err: err}
	}
	return &out, nil
}

// MarkNotificationsRead calls POST /notifications/read.
// Mark Notifications Read
//
// Acknowledge the notification feed — record that this user has seen
// everything up to now, clearing the header unread dot. **When to use**:
// called by the header bell when the user opens the panel (and by an
// explicit "Mark all read" action). After this returns, the next `GET
// /notifications` reports `unread_count = 0` until a notification with a
// `created_at` newer than this moment appears. **Effect**: upserts a
// per-user, org-scoped `last_seen_at = now` marker in ClickHouse
// (`notification_read_markers`, ReplacingMergeTree). Idempotent — calling
// it repeatedly just advances the marker to the latest `now`. **Auth**:
// AuditUIUser (audit.ui.read). Org-scoped to the caller.
func (s *NotificationsClient) MarkNotificationsRead(ctx context.Context, params *NotificationsMarkNotificationsReadParams) (*NotificationReadResponse, error) {
	b, err := s.c.do(ctx, "POST", "/notifications/read", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out NotificationReadResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/notifications/read", Err: err}
	}
	return &out, nil
}
