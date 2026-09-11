// ANNOTATION BLOCK
// File: auditclient.go — the ab0t Audit Service ingest client.
// Why it exists: the MCP gateway (and any ab0t Go service) needs to emit
// structured audit events to the Audit Service. This is the small, isolated,
// stdlib-only client for the INGEST surface only — it deliberately does NOT
// wrap the audit service's 224 other endpoints (compliance, search, ui, …),
// which are the audit product's own surface, not an emitter's concern.
// Contracts: `Event` == the service's `AuditEventCreate` schema (action,
// event_type, service required); `Client.IngestBatch` → POST /logs/ingest/batch
// (202); `Client.Ingest` → POST /logs/ingest (202). Fail-open by design at the
// caller layer — see the gateway's audit adapter, which never blocks the hot
// path on this client.
package auditclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// DefaultBaseURL is the production audit service. A caller may override.
const DefaultBaseURL = "https://audit.service.ab0t.com"

const (
	ingestPath      = "/logs/ingest"
	ingestBatchPath = "/logs/ingest/batch"
)

// Event mirrors the audit service `AuditEventCreate` schema. The three starred
// fields are REQUIRED by the service; the rest are optional (nil/"" omitted).
type Event struct {
	Action       string         `json:"action"`               // * e.g. tool name "stripe__create_invoice"
	EventType    string         `json:"event_type"`           // * e.g. "tools/call"
	Service      string         `json:"service"`              // * e.g. "mcp-gateway"
	Outcome      string         `json:"outcome,omitempty"`    // success | denied | authz_error | upstream_error
	OrgID        string         `json:"org_id,omitempty"`     // tenant (authclient Actor.OrgID)
	UserID       string         `json:"user_id,omitempty"`    // subject
	SessionID    string         `json:"session_id,omitempty"` // client Mcp-Session-Id
	ResourceType string         `json:"resource_type,omitempty"`
	ResourceID   string         `json:"resource_id,omitempty"`
	IPAddress    string         `json:"ip_address,omitempty"`
	UserAgent    string         `json:"user_agent,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`      // {latency_ms, safety_class, facet, attempt, ...}
	RequestData  map[string]any `json:"request_data,omitempty"`  // redaction-gated tool args
	ResponseData map[string]any `json:"response_data,omitempty"` // redaction-gated tool result
}

// Option configures the Client.
type Option func(*Client)

// WithAPIKey sets the service key used to authenticate to the audit service.
func WithAPIKey(key string) Option { return func(c *Client) { c.apiKey = key } }

// WithHTTPClient injects a custom *http.Client (e.g. a netguard-hardened one).
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.http = h } }

// Client is the Audit Service client. Safe for concurrent use.
//
// The minimal ingest surface (Ingest, IngestBatch) lives directly on Client for
// back-compatibility with emitters. The full Audit Service surface is reached
// through the typed sub-clients below (c.Logs, c.Compliance, c.Search, ...),
// which are wired up by New. See doc.go for an overview.
type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client

	// Sub-clients — one per API designation. Each groups the endpoints of that
	// designation; e.g. c.Logs.GetAuditEvents, c.Compliance.GetComplianceDashboard.
	Admin                *AdminClient
	Analytics            *AnalyticsClient
	AuditorShare         *AuditorShareClient
	Capabilities         *CapabilitiesClient
	Compliance           *ComplianceClient
	ComplianceAutomation *ComplianceAutomationClient
	ComplianceCopilot    *ComplianceCopilotClient
	ComplianceExtras     *ComplianceExtrasClient
	ComplianceTemplates  *ComplianceTemplatesClient
	ComplianceUI         *ComplianceUIClient
	ExplorerQueries      *ExplorerQueriesClient
	Export               *ExportClient
	Health               *HealthClient
	LegalHold            *LegalHoldClient
	Logs                 *LogsClient
	Notifications        *NotificationsClient
	ReportSchedules      *ReportSchedulesClient
	RuntimeEvents        *RuntimeEventsClient
	SavedSearches        *SavedSearchesClient
	Search               *SearchClient
	Settings             *SettingsClient
	UI                   *UIClient
}

// New constructs a Client. An empty baseURL uses DefaultBaseURL.
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/"),
		http:    &http.Client{Timeout: 10 * time.Second},
	}
	if c.baseURL == "" {
		c.baseURL = DefaultBaseURL
	}
	for _, o := range opts {
		o(c)
	}
	c.initServices()
	return c
}

// Ingest posts a single event to /logs/ingest. Returns nil on 2xx (202).
func (c *Client) Ingest(ctx context.Context, e Event) error {
	return c.post(ctx, ingestPath, e)
}

// IngestBatch posts a batch of events to /logs/ingest/batch. Returns nil on
// 2xx (202). An empty batch is a no-op.
func (c *Client) IngestBatch(ctx context.Context, events []Event) error {
	if len(events) == 0 {
		return nil
	}
	return c.post(ctx, ingestBatchPath, events)
}

func (c *Client) post(ctx context.Context, path string, body any) error {
	buf, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("auditclient: marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("auditclient: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("auditclient: post %s: %w", path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("auditclient: %s returned %d: %s", path, resp.StatusCode, strings.TrimSpace(string(snippet)))
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4<<10))
	return nil
}
