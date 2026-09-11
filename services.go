// ANNOTATION BLOCK
// File: services.go — sub-client wiring and the shared HTTP round-trip.
// Why it exists: the full Audit Service surface is grouped into sub-clients
// (c.Logs, c.Compliance, c.Search, ...) hung off the root *Client, mirroring
// the authclient shape. This file defines those sub-client types, attaches
// them in initServices (called by New), and implements do(): the single
// request/response/error path every generated method funnels through.
package auditclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// AdminClient groups the `Admin` endpoints of the Audit Service.
type AdminClient struct{ c *Client }

// AnalyticsClient groups the `Analytics` endpoints of the Audit Service.
type AnalyticsClient struct{ c *Client }

// AuditorShareClient groups the `AuditorShare` endpoints of the Audit Service.
type AuditorShareClient struct{ c *Client }

// CapabilitiesClient groups the `Capabilities` endpoints of the Audit Service.
type CapabilitiesClient struct{ c *Client }

// ComplianceClient groups the `Compliance` endpoints of the Audit Service.
type ComplianceClient struct{ c *Client }

// ComplianceAutomationClient groups the `ComplianceAutomation` endpoints of the Audit Service.
type ComplianceAutomationClient struct{ c *Client }

// ComplianceCopilotClient groups the `ComplianceCopilot` endpoints of the Audit Service.
type ComplianceCopilotClient struct{ c *Client }

// ComplianceExtrasClient groups the `ComplianceExtras` endpoints of the Audit Service.
type ComplianceExtrasClient struct{ c *Client }

// ComplianceTemplatesClient groups the `ComplianceTemplates` endpoints of the Audit Service.
type ComplianceTemplatesClient struct{ c *Client }

// ComplianceUIClient groups the `ComplianceUI` endpoints of the Audit Service.
type ComplianceUIClient struct{ c *Client }

// ExplorerQueriesClient groups the `ExplorerQueries` endpoints of the Audit Service.
type ExplorerQueriesClient struct{ c *Client }

// ExportClient groups the `Export` endpoints of the Audit Service.
type ExportClient struct{ c *Client }

// HealthClient groups the `Health` endpoints of the Audit Service.
type HealthClient struct{ c *Client }

// LegalHoldClient groups the `LegalHold` endpoints of the Audit Service.
type LegalHoldClient struct{ c *Client }

// LogsClient groups the `Logs` endpoints of the Audit Service.
type LogsClient struct{ c *Client }

// NotificationsClient groups the `Notifications` endpoints of the Audit Service.
type NotificationsClient struct{ c *Client }

// ReportSchedulesClient groups the `ReportSchedules` endpoints of the Audit Service.
type ReportSchedulesClient struct{ c *Client }

// RuntimeEventsClient groups the `RuntimeEvents` endpoints of the Audit Service.
type RuntimeEventsClient struct{ c *Client }

// SavedSearchesClient groups the `SavedSearches` endpoints of the Audit Service.
type SavedSearchesClient struct{ c *Client }

// SearchClient groups the `Search` endpoints of the Audit Service.
type SearchClient struct{ c *Client }

// SettingsClient groups the `Settings` endpoints of the Audit Service.
type SettingsClient struct{ c *Client }

// UIClient groups the `UI` endpoints of the Audit Service.
type UIClient struct{ c *Client }

// initServices wires the sub-clients onto c. Called by New.
func (c *Client) initServices() {
	c.Admin = &AdminClient{c: c}
	c.Analytics = &AnalyticsClient{c: c}
	c.AuditorShare = &AuditorShareClient{c: c}
	c.Capabilities = &CapabilitiesClient{c: c}
	c.Compliance = &ComplianceClient{c: c}
	c.ComplianceAutomation = &ComplianceAutomationClient{c: c}
	c.ComplianceCopilot = &ComplianceCopilotClient{c: c}
	c.ComplianceExtras = &ComplianceExtrasClient{c: c}
	c.ComplianceTemplates = &ComplianceTemplatesClient{c: c}
	c.ComplianceUI = &ComplianceUIClient{c: c}
	c.ExplorerQueries = &ExplorerQueriesClient{c: c}
	c.Export = &ExportClient{c: c}
	c.Health = &HealthClient{c: c}
	c.LegalHold = &LegalHoldClient{c: c}
	c.Logs = &LogsClient{c: c}
	c.Notifications = &NotificationsClient{c: c}
	c.ReportSchedules = &ReportSchedulesClient{c: c}
	c.RuntimeEvents = &RuntimeEventsClient{c: c}
	c.SavedSearches = &SavedSearchesClient{c: c}
	c.Search = &SearchClient{c: c}
	c.Settings = &SettingsClient{c: c}
	c.UI = &UIClient{c: c}
}

// do performs one request against the audit service and returns the response
// body on a 2xx status. body is JSON-encoded when non-nil. A non-2xx status
// yields an *APIError carrying the status, method, path and (truncated) body.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any) ([]byte, error) {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	var rdr io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, &EncodeError{Endpoint: path, Err: err}
		}
		rdr = bytes.NewReader(buf)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, rdr)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
		req.Header.Set("X-API-Key", c.apiKey)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 32<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, parseAPIError(resp.StatusCode, method, path, resp.Header.Get("X-Request-ID"), string(respBody))
	}
	return respBody, nil
}
