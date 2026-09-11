// ANNOTATION BLOCK
// File: svc_ui.go — the `UI` designation of the Audit Service.
// Why it exists: typed methods for every UI endpoint, reached via
// c.UI.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"fmt"
	"net/url"
)

// UIAppComplianceDemoPageParams holds the query parameters for UI.AppComplianceDemoPage.
type UIAppComplianceDemoPageParams struct {
	Scenario *string `json:"scenario,omitempty"` // Demo scenario identifier
}

func (p *UIAppComplianceDemoPageParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Scenario != nil {
		v.Set("scenario", *p.Scenario)
	}
	return v
}

// UIAppGenericDemoPageParams holds the query parameters for UI.AppGenericDemoPage.
type UIAppGenericDemoPageParams struct {
	Scenario *string `json:"scenario,omitempty"` // Demo scenario identifier
}

func (p *UIAppGenericDemoPageParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Scenario != nil {
		v.Set("scenario", *p.Scenario)
	}
	return v
}

// UIComplianceDashboardDemoPageParams holds the query parameters for UI.ComplianceDashboardDemoPage.
type UIComplianceDashboardDemoPageParams struct {
	OrgID    *string `json:"org_id,omitempty"`   // Organization ID
	Scenario *string `json:"scenario,omitempty"` // Demo scenario identifier
}

func (p *UIComplianceDashboardDemoPageParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Scenario != nil {
		v.Set("scenario", *p.Scenario)
	}
	return v
}

// UIComplianceOverviewPageParams holds the query parameters for UI.ComplianceOverviewPage.
type UIComplianceOverviewPageParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *UIComplianceOverviewPageParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// UILoginRedirectParams holds the query parameters for UI.LoginRedirect.
type UILoginRedirectParams struct {
	Next *string `json:"next,omitempty"` // Path to return to after login
}

func (p *UILoginRedirectParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Next != nil {
		v.Set("next", *p.Next)
	}
	return v
}

// AIAgentsPage calls GET /ui/ai-agents.
// Ai Agents Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). 301-redirects to
// /ui/v2/ai-agents; v1 template (ai-agents.html) moved to
// app/templates/.archive/.
func (s *UIClient) AIAgentsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/ai-agents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AIAgentsV2Page calls GET /ui/v2/ai-agents.
// Ai Agents V2 Page
//
// AI Agents V2 - Intent-based auditing focus. Focus: Why not just what,
// reasoning chains, human oversight logging.
func (s *UIClient) AIAgentsV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/ai-agents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AIPlaygroundPage calls GET /ui/ai-agents/playground.
// Ai Playground Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). 301-redirects to
// /ui/v2/ai-agents/playground; v1 template (ai-playground.html) moved to
// app/templates/.archive/.
func (s *UIClient) AIPlaygroundPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/ai-agents/playground", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AIPlaygroundV2Page calls GET /ui/v2/ai-agents/playground.
// Ai Playground V2 Page
//
// AI Playground V2 - Interactive AI agent audit demo. Features: Intent
// capture, MCP tool call logging, reasoning chain visualization.
func (s *UIClient) AIPlaygroundV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/ai-agents/playground", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAPIKeysPage calls GET /ui/app/api-keys.
// App Api Keys Page
//
// API Keys - Manage API keys for audit event ingestion. Features: API key
// listing with masked values, key creation/rotation/revocation, permission
// scopes, usage statistics, environment labels (prod/staging/dev).
func (s *UIClient) AppAPIKeysPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/api-keys", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAccountPage calls GET /ui/app/account.
// App Account Page
//
// My account (c4a9 §1a) — identity, session, credential handoff.
// Identity/session render client-side from the Auth Mesh SDK session
// (auth.getUser() + the bearer's JWT claims). Credentials (password / MFA)
// are OWNED by the auth service — the page hands off there rather than
// rebuilding credential management.
func (s *UIClient) AppAccountPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/account", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAdminPage calls GET /ui/app/admin.
// App Admin Page
//
// T16 — Admin dashboard surfacing event volume, security metrics, DB
// metrics, callback delivery log, and false-positive feedback.
func (s *UIClient) AppAdminPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/admin", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAgentDetailPage calls GET /ui/app/agents/{agent_id}.
// App Agent Detail Page
//
// Per-actor (agent / user) audit profile. Features: summary card strip
// (events, last seen, distinct tools, failures), top-tools bar chart,
// top-compliance-controls bar chart, recent activity stream using the
// shared event-row component.
func (s *UIClient) AppAgentDetailPage(ctx context.Context, agentID string) (string, error) {
	path := fmt.Sprintf("/ui/app/agents/%s", url.PathEscape(agentID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAgentsPage calls GET /ui/app/agents.
// App Agents Page
//
// Agent Monitor - Monitor AI agents and behavioral baselines. Features:
// Agent cards, behavioral baseline visualization, anomaly detection.
func (s *UIClient) AppAgentsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/agents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAlertsPage calls GET /ui/app/alerts.
// App Alerts Page
//
// Alerts - Real-time alerting for AI agent audit events. Features: Active
// alerts with severity levels, alert rules configuration, notification
// channels (Email, Slack, PagerDuty), alert history with status tracking.
func (s *UIClient) AppAlertsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/alerts", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppApprovalsPage calls GET /ui/app/approvals.
// App Approvals Page
//
// Human Oversight & Approvals - Manage workflow approvals and oversight
// modes. Features: Pending approval queue, multi-step plan review,
// HITL/HOTL/HOVL modes, approval history, risk-based filtering.
func (s *UIClient) AppApprovalsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/approvals", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAuditLogsPage calls GET /ui/app/audit-logs.
// App Audit Logs Page
//
// Audit Logs - Search and explore all audit events. Features: Filters,
// search, expandable event details, pagination.
func (s *UIClient) AppAuditLogsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/audit-logs", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAuditPeriodPage calls GET /ui/app/compliance/audit-period.
// App Audit Period Page
//
// Audit period configuration — sets the org-active SOC2 / annual review
// window. Server-backed: reads/writes GET/PUT/DELETE /settings/audit-period
// (org-scoped, table org_audit_periods), so the period is shared org-wide
// rather than per-browser.
func (s *UIClient) AppAuditPeriodPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/audit-period", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppAuditorInboxPage calls GET /ui/app/compliance/auditor-inbox.
// App Auditor Inbox Page
//
// External-auditor question tracker. Backend live (CS-014): GET/POST/PATCH
// /compliance/auditor-requests and POST
// /compliance/auditor-requests/{id}/satisfy are all implemented and persist
// to the compliance_auditor_requests table
// (app/api/compliance_extras.py:199,234,292,369).
func (s *UIClient) AppAuditorInboxPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/auditor-inbox", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppBehavioralBaselinesPage calls GET /ui/app/behavioral-baselines.
// App Behavioral Baselines Page
//
// Behavioral Baselines - Monitor agent behavioral patterns and anomalies.
// Features: Agent baseline visualization, anomaly detection alerts,
// deviation investigation, baseline configuration per agent.
func (s *UIClient) AppBehavioralBaselinesPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/behavioral-baselines", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppCallbacksPage calls GET /ui/app/callbacks.
// App Callbacks Page
//
// Event callbacks — manage the webhook subscriptions that push compliance
// events out. Surfaces the already-live callback subscription lifecycle
// endpoints (CS-053): GET /compliance-automation/callbacks/subscriptions
// list subscriptions POST /compliance-automation/callbacks/subscriptions
// create a subscription POST
// /compliance-automation/callbacks/subscriptions/{id}/pause pause a noisy
// one POST /compliance-automation/callbacks/subscriptions/{id}/test send a
// test delivery GET /compliance-automation/callbacks/deliveries delivery
// log Previously the admin page showed only the delivery LOG — the result
// of subscriptions that could only be created / paused / tested via
// hand-crafted API calls. This page lets an admin register an endpoint, see
// existing subscriptions, fire a test delivery, pause a chatty one, and
// jump to the delivery log filtered to a single subscription.
func (s *UIClient) AppCallbacksPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/callbacks", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComparePage calls GET /ui/app/compare.
// App Compare Page
//
// T21 — Side-by-side comparison of two time windows.
func (s *UIClient) AppComparePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compare", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceCalendarPage calls GET /ui/app/compliance/calendar.
// App Compliance Calendar Page
//
// Compliance calendar — month grid aggregating DSAR deadlines, report
// cadences, and audit-period boundaries.
func (s *UIClient) AppComplianceCalendarPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/calendar", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceControlDetailPage calls GET /ui/app/compliance/controls/{control_id}.
// App Compliance Control Detail Page
//
// Single compliance control detail. Features: status / framework /
// evidence-count / last-evidence summary cards, bundled control definition
// with link to authoritative reference, action shortcuts (view all events,
// build evidence package, export CSV, customise mapping), evidence stream
// rendered as expandable event-row pairs filtered to events tagged with
// this control. Uses :path converter so codes like `SOC2.CC6.1` and
// `GDPR.Art32.1.b` survive the URL routing.
func (s *UIClient) AppComplianceControlDetailPage(ctx context.Context, controlID string) (string, error) {
	path := fmt.Sprintf("/ui/app/compliance/controls/%s", url.PathEscape(controlID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceControlMappingsPage calls GET /ui/app/compliance/control-mappings.
// App Compliance Control Mappings Page
//
// Control mapping editor — override the hardcoded event→control map per
// org. Backend live (CS-007): GET/POST/DELETE /compliance/control-mappings
// (app/api/compliance.py:3339 list, :3388 create, :3462 delete; writes need
// audit.control_mappings.manage). UI shows the hardcoded defaults from
// app/core/control_mapper.py merged with per-org custom rules and lets ops
// add or remove them.
func (s *UIClient) AppComplianceControlMappingsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/control-mappings", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceCopilotPage calls GET /ui/app/compliance/copilot.
// App Compliance Copilot Page
//
// T30 — Compliance Copilot: NL ↔ ClickHouse chat.
func (s *UIClient) AppComplianceCopilotPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/copilot", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceDSARDetailPage calls GET /ui/app/compliance/dsar/{request_id}.
// App Compliance Dsar Detail Page
//
// Per-DSAR processing view. Features: subject identity, scope, processing
// notes, SLA countdown, status-workflow stepper (pending → in_review →
// in_progress → completed), full status-history timeline, advance-status
// button, add-note button, "Deliver data package" → POST /export/ with
// subject scope, "View related events" deep-link into audit-logs filtered
// by subject.
func (s *UIClient) AppComplianceDSARDetailPage(ctx context.Context, requestID string) (string, error) {
	path := fmt.Sprintf("/ui/app/compliance/dsar/%s", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceDSARPage calls GET /ui/app/compliance/dsar.
// App Compliance Dsar Page
//
// GDPR Data Subject Request queue — Article 15 (access), 17 (erasure), 20
// (portability), 16 (rectification). Features: status / type /
// deadline-window filters, 30-day SLA countdown chip per row, overdue /
// due-soon summary banner, click-through to per-request processing view.
func (s *UIClient) AppComplianceDSARPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/dsar", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceDemoPage calls GET /ui/demo/app/compliance.
// App Compliance Demo Page
//
// Demo entry for app compliance page. Sets a backend demo-scenario cookie
// so existing API endpoints can return realistic mock fixtures without
// embedding demo data in HTML.
func (s *UIClient) AppComplianceDemoPage(ctx context.Context, params *UIAppComplianceDemoPageParams) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/demo/app/compliance", params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceEvidenceBuilderPage calls GET /ui/app/compliance/evidence/new.
// App Compliance Evidence Builder Page
//
// Evidence package builder. 4-step form: package metadata, controls,
// time-range, sealing. Submits to POST /compliance/evidence-packages (live:
// create_evidence_package, app/api/auditor_share.py:186).
func (s *UIClient) AppComplianceEvidenceBuilderPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/evidence/new", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceEvidencePage calls GET /ui/app/compliance/evidence.
// App Compliance Evidence Page
//
// Evidence vault — sealed compliance evidence packages ready for auditor
// sharing. Backend live (CS-025): lists via GET
// /compliance/evidence-packages (list_evidence_packages,
// app/api/auditor_share.py:88). UI shows a real "no packages yet → build
// one" empty-state for 0 rows, and actionable auth/transient errors
// otherwise.
func (s *UIClient) AppComplianceEvidencePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/evidence", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceFrameworksPage calls GET /ui/app/compliance/frameworks.
// App Compliance Frameworks Page
//
// Framework manager — activate/deactivate framework packs per org. Backed
// by live endpoints: - GET /compliance-automation/framework-packs (list) -
// POST /compliance-automation/framework-packs/{framework}/activate - POST
// /compliance-automation/framework-packs/{framework}/deactivate Handlers in
// app/modules/compliance_automation/api_handlers.py
// (activate_framework_pack / deactivate_framework_pack); both require the
// audit.frameworks.manage permission. The page surfaces actionable errors
// (403 → missing permission, transient → retry).
func (s *UIClient) AppComplianceFrameworksPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/frameworks", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppCompliancePage calls GET /ui/app/compliance.
// App Compliance Page
//
// Compliance Controls - Monitor regulatory compliance across frameworks.
// Features: Framework score cards (SOC2, GDPR, EU AI Act, ISO 42001,
// HIPAA), control categories with pass/fail status, open issues, assessment
// scheduling.
func (s *UIClient) AppCompliancePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceReadinessPage calls GET /ui/app/compliance/readiness.
// App Compliance Readiness Page
//
// Framework readiness matrix — per-control pass/fail grid for the selected
// framework. Backed by GET /compliance/controls?framework=X. No new
// backend; this is the "what controls are passing/failing right now for my
// org" page that compliance managers ask for in standups. Drill into a row
// → /ui/app/compliance/controls/{id}.
func (s *UIClient) AppComplianceReadinessPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/readiness", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceSOC2TrailPage calls GET /ui/app/compliance/soc2-trail.
// App Compliance Soc2 Trail Page
//
// SOC2 Audit Trail viewer — tamper-evident hash chain over compliance
// events. Features: date-range + control-objective filters, summary cards
// (total events, unique actors, event-type coverage, hash-chain integrity
// status), chronological timeline with per-event hash-chain verification
// badge, auditor-ready CSV export.
func (s *UIClient) AppComplianceSOC2TrailPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/soc2-trail", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppComplianceTemplatesPage calls GET /ui/app/compliance/templates.
// App Compliance Templates Page
//
// T31 — Compliance template library.
func (s *UIClient) AppComplianceTemplatesPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/templates", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppDashboardPage calls GET /ui/app/dashboard.
// App Dashboard Page
//
// Application Dashboard - Main overview of audit platform. Shows: Event
// stats, agent activity, recent events, compliance status, alerts.
func (s *UIClient) AppDashboardPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/dashboard", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppDataExplorerPage calls GET /ui/app/data-explorer.
// App Data Explorer Page
//
// Data Explorer - Natural language interface to query audit data. Features:
// Conversational AI interface, saved queries, query history, results as
// tables/charts, export capabilities. Powered by multi-step agentic
// workflow for translating natural language to database queries.
func (s *UIClient) AppDataExplorerPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/data-explorer", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppDemoPagePayload calls GET /ui/demo/app-data/{page_slug}.
// App Demo Page Payload
//
// Backend demo payload for app pages. Allows static pages to hydrate
// scenario-specific content without embedding mock fixtures in HTML
// templates.
func (s *UIClient) AppDemoPagePayload(ctx context.Context, pageSlug string) (RawMessage, error) {
	path := fmt.Sprintf("/ui/demo/app-data/%s", url.PathEscape(pageSlug))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// AppExportBuilderPage calls GET /ui/app/exports/new.
// App Export Builder Page
//
// Export builder — define filters, fields, format, and delivery for a new
// export job. Features: name + date-range + event-type + user-id filters,
// field-list checkboxes, format selector (CSV / JSON / XML / Parquet),
// delivery method, include-request / include-response toggles. Submits to
// POST /export/.
func (s *UIClient) AppExportBuilderPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/exports/new", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppExportDetailPage calls GET /ui/app/exports/{export_id}.
// App Export Detail Page
//
// Per-export detail view. Features: status banner, auto-poll while
// queued/processing, filter + field summary, download button when
// completed, cancel for in-flight, re-run with same parameters.
func (s *UIClient) AppExportDetailPage(ctx context.Context, exportID string) (string, error) {
	path := fmt.Sprintf("/ui/app/exports/%s", url.PathEscape(exportID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppExportsPage calls GET /ui/app/exports.
// App Exports Page
//
// Data Exports — list of all export jobs for the org. Features: status
// filter, per-row status badge, download link when ready, cancel button for
// queued / processing jobs.
func (s *UIClient) AppExportsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/exports", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppFlowPage calls GET /ui/app/flow.
// App Flow Page
//
// T20 — Sankey flow: agent → tool → outcome.
func (s *UIClient) AppFlowPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/flow", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppGenericDemoPage calls GET /ui/demo/app/{page_slug}.
// App Generic Demo Page
//
// Generic demo entry for all authenticated app pages. Uses backend-issued
// demo scenario cookie so page API calls can return deterministic fixtures
// without frontend mock branches.
func (s *UIClient) AppGenericDemoPage(ctx context.Context, pageSlug string, params *UIAppGenericDemoPageParams) (string, error) {
	path := fmt.Sprintf("/ui/demo/app/%s", url.PathEscape(pageSlug))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppHelpAPIPage calls GET /ui/app/help/api.
// App Help Api Page
//
// T3.2 — API explorer (Redoc embedded against /openapi.json).
func (s *UIClient) AppHelpAPIPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/help/api", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppHelpChangelogPage calls GET /ui/app/help/changelog.
// App Help Changelog Page
//
// T3.3 — Changelog rendered from the platform's worklog history.
func (s *UIClient) AppHelpChangelogPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/help/changelog", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppHelpGlossaryPage calls GET /ui/app/help/glossary.
// App Help Glossary Page
//
// T3.1 — Glossary of platform terms.
func (s *UIClient) AppHelpGlossaryPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/help/glossary", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppIntegrationsPage calls GET /ui/app/integrations.
// App Integrations Page
//
// Third-Party Integrations - Monitor AI agent access to external services.
// Features: Connected services grid, API call stats, risk assessment,
// activity log showing all third-party API interactions with audit trail.
func (s *UIClient) AppIntegrationsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/integrations", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppIntentsPage calls GET /ui/app/intents.
// App Intents Page
//
// Intent Logs - Track the "why" behind every AI agent action. Features:
// Intent cards with reasoning chains, linked tool calls, approval workflows
// for high-risk operations, and execution context.
func (s *UIClient) AppIntentsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/intents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppJudgeConfigPage calls GET /ui/app/judge-config.
// App Judge Config Page
//
// Judge configuration — govern the versioned prompts and policies that
// drive the AI judge. Surfaces the already-live judge governance endpoints
// (CS-052): GET /compliance-automation/judge/prompts list prompt versions
// POST /compliance-automation/judge/prompts create / update a version POST
// /compliance-automation/judge/prompts/{prompt_key}/activate promote a
// version POST /compliance-automation/judge/prompts/{prompt_key}/rollback
// revert to a version GET /compliance-automation/judge/policies list
// per-control policies POST /compliance-automation/judge/policies create /
// update a policy Lets a compliance admin see which prompt version is
// currently active, promote a new version or roll back a bad one in a
// couple of clicks, and control which controls the judge runs on (and at
// what thresholds) — the governance flow these endpoints exist for, which
// previously required hand-crafted API calls.
func (s *UIClient) AppJudgeConfigPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/judge-config", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppJudgeVerdictsPage calls GET /ui/app/judge-verdicts.
// App Judge Verdicts Page
//
// AI-judge verdict log — see and replay the LLM judgments that decide
// control outcomes. Surfaces the already-live judge endpoints (CS-051): GET
// /compliance-automation/judge/verdicts list verdicts (org-scoped) POST
// /compliance-automation/judge/verdicts/{id}/replay deterministic re-run
// Lets a compliance admin / auditor list every per-control judge verdict,
// read the judge's rationale + the assertion it judged in plain language,
// and replay a stored verdict to prove it is reproducible
// (audit-defensibility).
func (s *UIClient) AppJudgeVerdictsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/judge-verdicts", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppLegalHoldsPage calls GET /ui/app/legal-holds.
// App Legal Holds Page
//
// Legal holds — place, review, and release litigation holds in-product.
// Surfaces the already-live, fully-enforced legal-hold lifecycle (CS-055):
// GET /legal-holds list active holds + full hold history POST /legal-holds
// place a hold (org / data_subject / before_date) POST
// /legal-holds/{hold_id}/release release a hold (append-only) A legal hold
// blocks deletion — GDPR erasure, admin purge, and retention TTL all refuse
// to remove data the hold protects — until it is explicitly released. The
// backend has always enforced this, but the only way to place or release a
// hold was by hand-crafting API calls. This page gives a compliance manager
// facing real litigation or an investigation a way to see what is currently
// held, place a new hold, and release one when the matter closes.
func (s *UIClient) AppLegalHoldsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/legal-holds", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppLivePage calls GET /ui/app/live.
// App Live Page
//
// Live event stream — real-time feed of audit events. Backend live
// (CS-013): GET /events/stream serves a real Server-Sent-Events push stream
// (app/api/events.py:228); live.js consumes it via EventSource as the
// primary path and falls back to polling POST /search/ only when SSE is
// unavailable (no EventSource support or a stream error).
func (s *UIClient) AppLivePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/live", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppMLTrainingPage calls GET /ui/app/ml/training.
// App Ml Training Page
//
// T29 — Operator view for the audit-model-worker training loop.
func (s *UIClient) AppMLTrainingPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/ml/training", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppMetaAuditPage calls GET /ui/app/meta-audit.
// App Meta Audit Page
//
// Meta-audit — audit of the audit. Records who queried, exported, or shared
// audit data. Required for SOC2 CC6.1 access-monitoring evidence. F-E2-1
// (e2a1 ROLE_CAPABILITY_MATRIX): this page is the external-auditor role's
// HEADLINE capability, yet it was gated by the generic “AuditUIUser“
// (audit.ui.read) while “AuditMetaReader“ (audit.meta.read) gated no
// route — a declared-but-unenforced SOC2-CC6.1 control. Gating this route
// with “AuditMetaReader“ makes the control real. Role math (verified
// against .permissions.json): “audit.system.admin“ IMPLIES
// “audit.meta.read“ so audit-admin + audit-platform-admin keep access;
// audit-external-auditor holds meta.read directly and keeps access.
// audit-viewer (default signup) and audit-compliance-reviewer hold only
// ui.read → they LOSE this page. That access reduction is intended
// (audit-of-the-audit is not viewer material); stated explicitly in the
// e2a1 work_log append. Delegation basis for touching this route: program
// work_log wave-9 BEFORE-CLAIMS (W13, fable). Ref:
// LOOP_PROTOCOL_20260704.md.
func (s *UIClient) AppMetaAuditPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/meta-audit", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppOnboardingPage calls GET /ui/app/onboarding.
// App Onboarding Page
//
// First-run onboarding wizard. Four steps: pick frameworks, get API key,
// send a test event, verify it landed.
func (s *UIClient) AppOnboardingPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/onboarding", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppPlanPage calls GET /ui/app/plan.
// App Plan Page
//
// Plan & usage SHELL (c4a9 §1c, per DECISIONS D-1: volume+retention tiers).
// Shows the current-tier placeholder, event-volume usage (GET
// /logs/stats/events), the org's retention window (GET
// /admin/retention/policies — admin-gated, the page degrades honestly for
// non-admins), and a contact-us upgrade CTA. E5 SEAM: billing/entitlement
// wiring lands in epic E5 (see pages/plan.js).
func (s *UIClient) AppPlanPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/plan", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppPolicyEnginePage calls GET /ui/app/policy-engine.
// App Policy Engine Page
//
// Policy engine — tune the deterministic rules and signal merge-policies
// that decide verdicts. Surfaces the already-live policy-engine config
// endpoints (CS-054): GET /compliance-automation/policies/rules list
// deterministic policy rules POST /compliance-automation/policies/rules
// create / update a rule GET /compliance-automation/merge-policies list
// signal merge-policies POST /compliance-automation/merge-policies create /
// update a merge-policy A POLICY RULE sets the per-control confidence
// thresholds that turn a signal into pass / warn / fail. A MERGE POLICY
// decides how the three signals — the deterministic rule engine, the ML
// anomaly score, and the LLM judge — are weighted and arbitrated into one
// control outcome. These previously could only be edited with hand-crafted
// API calls; this page lets a compliance admin see and adjust them
// in-product.
func (s *UIClient) AppPolicyEnginePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/policy-engine", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppPostureHistoryPage calls GET /ui/app/compliance/posture-history.
// App Posture History Page
//
// Posture history — sparkline + per-framework trend of compliance score
// over time.
func (s *UIClient) AppPostureHistoryPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/compliance/posture-history", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppReasoningChainDetailPage calls GET /ui/app/reasoning-chains/{run_id}.
// App Reasoning Chain Detail Page
//
// Per-run reasoning-chain trace visualisation. Backed by GET
// /compliance-automation/runs/{run_id}/trace. Renders a step-by-step
// ordered timeline assembled from all section rows that carry a timestamp,
// plus a section-overview footer.
func (s *UIClient) AppReasoningChainDetailPage(ctx context.Context, runID string) (string, error) {
	path := fmt.Sprintf("/ui/app/reasoning-chains/%s", url.PathEscape(runID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppReasoningChainsPage calls GET /ui/app/reasoning-chains.
// App Reasoning Chains Page
//
// Reasoning Chains - Inspect AI agent reasoning and decision processes.
// Features: Chain-of-thought visualization, validation scores, goal drift
// detection, step-by-step inspection.
func (s *UIClient) AppReasoningChainsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/reasoning-chains", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppReportsPage calls GET /ui/app/reports.
// App Reports Page
//
// Reports - Generate and schedule compliance and audit reports. Features:
// Report templates (SOC2, GDPR, Executive Summary, AI Governance),
// generated reports history, scheduled report configuration, export
// options.
func (s *UIClient) AppReportsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/reports", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppRiskAssessmentPage calls GET /ui/app/risk-assessment.
// App Risk Assessment Page
//
// Risk Assessment - Real-time risk scoring for AI agent actions. Features:
// Risk score dashboard, high-risk action trends, risk factor breakdown,
// threshold configuration.
func (s *UIClient) AppRiskAssessmentPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/risk-assessment", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppSettingsClassificationPage calls GET /ui/app/settings/classification.
// App Settings Classification Page
//
// Data classification rule editor. Backend live (CS-009): GET/PUT
// /settings/classification.
func (s *UIClient) AppSettingsClassificationPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/settings/classification", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppSettingsNotificationsPage calls GET /ui/app/settings/notifications.
// App Settings Notifications Page
//
// Notification preferences editor (c4a9 §1d). Per-user, org-scoped toggles
// for the header-bell feed categories (report_failed / auditor_request /
// dsar_deadline / security_alert). Backend live: GET/PUT
// /settings/notifications (app/api/settings.py), persisted in the
// settings_notification_prefs ClickHouse table.
func (s *UIClient) AppSettingsNotificationsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/settings/notifications", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppSettingsPage calls GET /ui/app/settings.
// App Settings Page
//
// Settings - Configure audit platform preferences. Features: General
// settings, API keys, notifications, data retention, compliance.
func (s *UIClient) AppSettingsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/settings", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppSettingsRedactionPage calls GET /ui/app/settings/redaction.
// App Settings Redaction Page
//
// Redaction key list editor. Backend live (CS-008): GET/PUT
// /settings/redaction.
func (s *UIClient) AppSettingsRedactionPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/settings/redaction", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppSettingsRetentionPage calls GET /ui/app/settings/retention.
// App Settings Retention Page
//
// Retention policies editor. Backend live (CS-010): GET
// /admin/retention/policies (app/api/admin.py:284) reads the per-org policy
// and PUT /admin/retention/policies (app/api/admin.py:349) persists edits —
// the in-page editor saves through it.
func (s *UIClient) AppSettingsRetentionPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/settings/retention", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppStatusPage calls GET /ui/app/status.
// App Status Page
//
// Service health page. Polls /admin/status every 30 seconds. Renders:
// per-dependency status cards (ClickHouse, Redis, Queue, Workers), overall
// health badge, version / uptime / events-per-second / queue-depth detail.
func (s *UIClient) AppStatusPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/status", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppTeamPage calls GET /ui/app/team.
// App Team Page
//
// Org & team (c4a9 §1b) — member list, roles, invite shell. Members/roles
// are fetched BROWSER-SIDE from the auth service's org-scoped endpoints
// (GET /organizations/{org_id}/users, POST /organizations/{org_id}/invite)
// with the user's own bearer — the audit backend does NOT proxy auth
// endpoints. When the auth service is unreachable from the browser (CORS
// pin / network), the page renders an honest handoff card instead. The full
// invite ACCEPT flow is epic E2's; this page only exposes the invite entry
// point.
func (s *UIClient) AppTeamPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/team", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppToolCallsPage calls GET /ui/app/tool-calls.
// App Tool Calls Page
//
// Tool Calls - Monitor all tool executions by AI agents. Features: Tool
// call table with filters, risk assessment, latency tracking, expandable
// rows showing input/output parameters, metadata, and approval status.
func (s *UIClient) AppToolCallsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/tool-calls", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppToolsPage calls GET /ui/app/tools.
// App Tools Page
//
// Tool inventory page. Derived from /search/facets event-type distribution
// filtered to tool-shaped types. Time-range and free-text filters on the
// page.
func (s *UIClient) AppToolsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/tools", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AppWebhooksPage calls GET /ui/app/webhooks.
// App Webhooks Page
//
// Webhooks - Send audit events to external systems in real-time. Features:
// Webhook configuration, delivery stats, event subscriptions, recent
// delivery logs with retry status, failure alerts.
func (s *UIClient) AppWebhooksPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/app/webhooks", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AuditorViewPage calls GET /ui/auditor/{share_token}.
// Auditor View Page
//
// External-auditor read-only view of a sealed evidence package. No JWT
// required — the token in the URL is the auth artefact. The page JS fetches
// /auditor/share/{token} (also no-auth) which validates the token, checks
// expiry / revocation, and returns the watermarked package payload. Served
// via the dedicated `auditor-base.html` shell — no sidebar, watermarked
// banner, locked-down robots meta, print-friendly stylesheet.
func (s *UIClient) AuditorViewPage(ctx context.Context, shareToken string) (string, error) {
	path := fmt.Sprintf("/ui/auditor/%s", url.PathEscape(shareToken))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AutomoutsAgentsPageRedirect calls GET /ui/automouts_agents.
// Automouts Agents Page Redirect
//
// Backward-compatible redirect for the previously misspelled path. The
// canonical route is now /ui/autonomous_agents.
func (s *UIClient) AutomoutsAgentsPageRedirect(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/automouts_agents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AutonomousAgentsPage calls GET /ui/autonomous_agents.
// Autonomous Agents Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). 301-redirects to
// the v2 autonomous-workers page; v1 template (autonomous_agents.html)
// moved to app/templates/.archive/.
func (s *UIClient) AutonomousAgentsPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/autonomous_agents", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AutonomousWorkersV2Page calls GET /ui/v2/autonomous-workers.
// Autonomous Workers V2 Page
//
// Autonomous Workers V2 - Audit for autonomous AI agents. Focus: Multi-step
// plan approval, behavioral baselines, human oversight modes,
// intent-action-outcome audit loop for background AI workers.
func (s *UIClient) AutonomousWorkersV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/autonomous-workers", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ComplianceDashboardDemoPage calls GET /ui/demo/compliance/dashboard.
// Compliance Dashboard Demo Page
//
// Demo entry for live compliance dashboard. Uses the same template and API
// paths as production, with backend demo data selected through a server
// cookie.
func (s *UIClient) ComplianceDashboardDemoPage(ctx context.Context, params *UIComplianceDashboardDemoPageParams) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/demo/compliance/dashboard", params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ComplianceOverviewPage calls GET /ui/compliance/overview.
// Compliance Overview Page
//
// Compliance Overview - Real-time compliance monitoring. Shows control
// status, compliance scores, open issues, data subject requests, and recent
// compliance events.
func (s *UIClient) ComplianceOverviewPage(ctx context.Context, params *UIComplianceOverviewPageParams) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/compliance/overview", params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CompliancePage calls GET /ui/compliance.
// Compliance Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). 301-redirects to
// /ui/v2/compliance; v1 template (compliance.html) moved to
// app/templates/.archive/. NOTE: the longer /ui/compliance/overview and
// /ui/compliance/dashboard routes are unaffected (distinct paths).
func (s *UIClient) CompliancePage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/compliance", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ComplianceV2Page calls GET /ui/v2/compliance.
// Compliance V2 Page
//
// Compliance V2 - AI compliance focus. Focus: EU AI Act, ISO 42001, GDPR
// Article 22, AI-specific regulations.
func (s *UIClient) ComplianceV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/compliance", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ControlPlaneAgentAuditBlogPage calls GET /ui/blog/control-plane-agent-audit-system.
// Control Plane Agent Audit Blog Page
//
// Technical first-person blog explainer for control-plane/data-plane
// architecture.
func (s *UIClient) ControlPlaneAgentAuditBlogPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/blog/control-plane-agent-audit-system", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// DevelopersPage calls GET /ui/developers.
// Developers Page
//
// Developer docs — the event-ingest contract (b3d8 TASK 11, N2). Public
// page a stranger can integrate from: the two-product split (POST
// /logs/ingest + /logs/ingest/batch for AI-agent events; POST
// /events/ingest + /events/batch for runtime events), X-API-Key auth,
// event-taxonomy basics, and ONE canonical copy-paste curl snippet (also
// mirrored to tickets/2026-07-04-b3d8-audit-e3-product-surface/
// canonical_ingest_snippet.txt for E4 onboarding to reuse). Every field is
// grounded in app/api/logs.py, app/api/events.py, and models/audit.py. URL
// choice: `/ui/developers` (not `/ui/docs/ingest`) — a single, stable,
// guessable developer-hub URL that does not collide with FastAPI's own
// interactive `/docs` (Swagger UI) mount.
func (s *UIClient) DevelopersPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/developers", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ExitDemoMode calls GET /ui/demo/exit.
// Exit Demo Mode
//
// Clear backend demo scenario cookie and redirect to standard app
// dashboard.
func (s *UIClient) ExitDemoMode(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/demo/exit", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// IntegrationsV2Page calls GET /ui/v2/integrations.
// Integrations V2 Page
//
// Integrations V2 - Third-party API audit via integration proxy. Focus:
// Pre-flight risk assessment, provider-specific rules, credential audit,
// complete visibility for Google, Slack, Notion, and 200+ integrations.
func (s *UIClient) IntegrationsV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/integrations", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LandingPage calls GET /ui/landing.
// Landing Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). v2 is now the
// single served marketing generation. This v1 route 301-redirects to its v2
// equivalent; the v1 template (landing.html) was moved to
// app/templates/.archive/. See
// tickets/2026-07-04-b3d8-audit-e3-product-surface.
func (s *UIClient) LandingPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/landing", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LandingV2Page calls GET /ui/v2/landing.
// Landing V2 Page
//
// Landing V2 - AI Agent Audit Infrastructure positioning. Focus:
// Intent-based auditing, MCP support, autonomous AI worker audit.
func (s *UIClient) LandingV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/landing", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LoginRedirect calls GET /ui/login.
// Login Redirect
//
// Start the Auth Mesh OAuth 2.1 + PKCE sign-in. PKCE requires the
// code-verifier/challenge to be generated client-side (RFC 7636), so a pure
// server 302 cannot drive the flow. This route returns a minimal HTML page
// that loads the vendored Auth Mesh SDK and immediately calls
// `auth.loginWithRedirect()`, which full-page-redirects the browser to
// `<HOSTED_LOGIN_URL>?response_type=code&client_id=...&code_challenge=...`.
// After sign-in the user comes back to `/auth/callback` (see
// app/api/auth_web.py), which exchanges the code, seeds the auth cookie,
// and lands them on the dashboard. `HOSTED_LOGIN_URL` is now informational
// back-compat: it is surfaced only as a no-JS degraded fallback link. The
// SDK derives the real login URL from env-config.js (auth domain + org
// slug). All strings below are static (house convention: no dynamic detail
// in client responses).
func (s *UIClient) LoginRedirect(ctx context.Context, params *UILoginRedirectParams) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/login", params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LogoutRedirect calls GET /ui/logout.
// Logout Redirect
//
// Sign out — clear local auth cookies and bounce to landing. Server-side
// cookie expiry; the client-side counterpart in app-base.js
// (`AppBase.signOut`) clears the same cookies + localStorage when triggered
// from the user menu dropdown.
func (s *UIClient) LogoutRedirect(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/logout", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PageDirectoryPage calls GET /ui/directory.
// Page Directory Page
//
// Fuzzy-searchable page directory (Fuse.js, CDN) — every page this service
// serves, grouped by category, with a public/viewer/role/token gated-badge
// per entry. Distinct from `/ui/pages` (the plain curated sitemap this
// directory is styled after conceptually but supersedes for discovery UX).
// Public by design: it is a discovery aid, not an authenticated surface —
// it lists paths and one-line descriptions only, no live data. The catalog
// is a hand-maintained JSON array embedded in templates/directory.html (see
// the drift-risk comment in that file).
func (s *UIClient) PageDirectoryPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/directory", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PagesIndexPage calls GET /ui/pages.
// Pages Index Page
//
// Index of HTML pages served by this service. Includes public routes, v2
// routes, authenticated app routes, template files that are currently not
// mapped to a route, AND a drift section listing routes registered on the
// FastAPI router but not in the hand-maintained sitemap (so future
// additions don't silently fall off the index — T39).
func (s *UIClient) PagesIndexPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/pages", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PricingPage calls GET /ui/pricing.
// Pricing Page
//
// Pricing & packaging shell (b3d8 TASK 10, N1). Public marketing page.
// Three tiers per DECISIONS D-1 (volume + retention; unlimited seats): Free
// (30-day retention, capped events/mo, 1 framework pack), Pro (1-year
// retention, higher cap, all packs), Enterprise (7-year, custom). All CTAs
// are "Contact us" — real billing/checkout wiring is E5 (see the E5 SEAM
// comment in pricing.html). Also carries the D-3/N3 "AI features & your
// data" disclosure (copilot GA, judge opt-in, provider disclosure pending
// D-6).
func (s *UIClient) PricingPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/pricing", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// SdkV2Page calls GET /ui/v2/sdk.
// Sdk V2 Page
//
// SDK V2 - Python SDK for emitting audit events. Focus: Quickstart guide,
// event types, integration patterns (LangChain, CrewAI, OpenAI, Anthropic),
// OpenTelemetry support, API reference.
func (s *UIClient) SdkV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/sdk", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToolCallingPage calls GET /ui/tool-calling.
// Tool Calling Page
//
// Marketing generation v1 → v2 (C2, ticket b3d8 TASK 8). 301-redirects to
// /ui/v2/tool-calling; v1 template (tool-calling.html) moved to
// app/templates/.archive/.
func (s *UIClient) ToolCallingPage(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/tool-calling", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ToolCallingV2Page calls GET /ui/v2/tool-calling.
// Tool Calling V2 Page
//
// Tool Calling V2 - MCP audit focus. Focus: Model Context Protocol, tool
// call audit, security verification.
func (s *UIClient) ToolCallingV2Page(ctx context.Context) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/v2/tool-calling", nil, nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
