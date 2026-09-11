// ANNOTATION BLOCK
// File: svc_analytics.go — the `Analytics` designation of the Audit Service.
// Why it exists: typed methods for every Analytics endpoint, reached via
// c.Analytics.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// AnalyticsGetComplianceReportParams holds the query parameters for Analytics.GetComplianceReport.
type AnalyticsGetComplianceReportParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // required. Compliance period start (e.g. quarter/year start)
	EndTime   *string `json:"end_time,omitempty"`   // required. Compliance period end (e.g. quarter/year end)
}

func (p *AnalyticsGetComplianceReportParams) values() url.Values {
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
	return v
}

// AnalyticsGetEventVolumeAnalyticsParams holds the query parameters for Analytics.GetEventVolumeAnalytics.
type AnalyticsGetEventVolumeAnalyticsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // required. Period start (e.g. last 24 hours, last week)
	EndTime   *string `json:"end_time,omitempty"`   // required. Period end (typically now)
}

func (p *AnalyticsGetEventVolumeAnalyticsParams) values() url.Values {
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
	return v
}

// AnalyticsGetFlowAnalyticsParams holds the query parameters for Analytics.GetFlowAnalytics.
type AnalyticsGetFlowAnalyticsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // required. Period start
	EndTime   *string `json:"end_time,omitempty"`   // required. Period end
	Limit     *int    `json:"limit,omitempty"`      // Max edges to return
}

func (p *AnalyticsGetFlowAnalyticsParams) values() url.Values {
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
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// AnalyticsGetGeoAnalyticsParams holds the query parameters for Analytics.GetGeoAnalytics.
type AnalyticsGetGeoAnalyticsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // required. Period start
	EndTime   *string `json:"end_time,omitempty"`   // required. Period end
}

func (p *AnalyticsGetGeoAnalyticsParams) values() url.Values {
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
	return v
}

// AnalyticsGetIngestHealthParams holds the query parameters for Analytics.GetIngestHealth.
type AnalyticsGetIngestHealthParams struct {
	LookbackHours    *int    `json:"lookback_hours,omitempty"`    // Lookback window in hours (1–168)
	OrgID            *string `json:"org_id,omitempty"`            // Organization ID; omit to roll up across orgs the caller can see
	ExpectedServices *string `json:"expected_services,omitempty"` // Comma-separated list of services we EXPECT to see emitting; missing ones land in expect...
}

func (p *AnalyticsGetIngestHealthParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.LookbackHours != nil {
		v.Set("lookback_hours", strconv.Itoa(*p.LookbackHours))
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.ExpectedServices != nil {
		v.Set("expected_services", *p.ExpectedServices)
	}
	return v
}

// AnalyticsGetSecurityAnalyticsParams holds the query parameters for Analytics.GetSecurityAnalytics.
type AnalyticsGetSecurityAnalyticsParams struct {
	OrgID      *string `json:"org_id,omitempty"`      // Organization ID
	StartTime  *string `json:"start_time,omitempty"`  // required. Period start for security analysis
	EndTime    *string `json:"end_time,omitempty"`    // required. Period end for security analysis
	AlertLimit *int    `json:"alert_limit,omitempty"` // Max security alerts to return
}

func (p *AnalyticsGetSecurityAnalyticsParams) values() url.Values {
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
	if p.AlertLimit != nil {
		v.Set("alert_limit", strconv.Itoa(*p.AlertLimit))
	}
	return v
}

// AnalyticsGetTopUsersAnalyticsParams holds the query parameters for Analytics.GetTopUsersAnalytics.
type AnalyticsGetTopUsersAnalyticsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // required. Period start for ranking users
	EndTime   *string `json:"end_time,omitempty"`   // required. Period end for ranking users
	Limit     *int    `json:"limit,omitempty"`      // How many top users to return (1-100)
}

func (p *AnalyticsGetTopUsersAnalyticsParams) values() url.Values {
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
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// GetComplianceReport calls GET /analytics/compliance.
// Get Compliance Report
//
// Generate compliance summary report for audits (GDPR, SOC2, HIPAA). Prove
// audit trail completeness. **When to use**: Quarterly/annual audits,
// compliance reviews, regulatory reporting. **Returns**: Event totals, data
// subject counts, retention compliance status, access patterns, export
// activities. **Period**: Typically quarterly or annually. Matches
// audit/reporting periods. **Next step**: Use for attestation reports.
// Export detailed data with POST /export/ if auditors need raw logs.
// **Tip**: Run monthly to catch compliance issues early, not just during
// audit season.
func (s *AnalyticsClient) GetComplianceReport(ctx context.Context, params *AnalyticsGetComplianceReportParams) (*ComplianceAnalyticsReport, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/compliance", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ComplianceAnalyticsReport
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/compliance", Err: err}
	}
	return &out, nil
}

// GetEventVolumeAnalytics calls GET /analytics/volume.
// Get Event Volume Analytics
//
// Get event volume statistics and trends over time. See activity patterns
// and usage. **When to use**: Dashboard metrics, capacity planning,
// identifying usage spikes or drops. **Returns**: Total event count, hourly
// timeseries, breakdown by event_type and outcome. **Time range**: Works
// for any range - hour, day, week, month. Hourly granularity. **Next
// step**: Drill into specific event_type or time period with GET
// /logs/events or POST /search/. **Tip**: Compare same time ranges across
// different periods to spot trends.
func (s *AnalyticsClient) GetEventVolumeAnalytics(ctx context.Context, params *AnalyticsGetEventVolumeAnalyticsParams) (*EventVolumeStats, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/volume", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out EventVolumeStats
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/volume", Err: err}
	}
	return &out, nil
}

// GetFlowAnalytics calls GET /analytics/flow.
// Get Flow Analytics
//
// Aggregate audit events by (user_id, event_type, outcome) for sankey
// visualization. Filters to tool-call events only. **Auth**:
// AuditAnalyticsReader.
func (s *AnalyticsClient) GetFlowAnalytics(ctx context.Context, params *AnalyticsGetFlowAnalyticsParams) (*FlowResponse, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/flow", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out FlowResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/flow", Err: err}
	}
	return &out, nil
}

// GetGeoAnalytics calls GET /analytics/geo.
// Get Geo Analytics
//
// Aggregate audit events by IP-derived country. **When to use**: visualize
// the geographic distribution of agent / user activity for the period.
// Pairs with `GET /analytics/security` to spot unusual source-country
// bursts. **Implementation**: each source IP is resolved to its real ISO
// 3166-1 country via the MaxMind-backed resolver in `app.core.geoip`
// (GeoLite2 country DB; path config-driven via `GEOIP_COUNTRY_DB_PATH`).
// Private / internal / reserved ranges bucket honestly to "Internal
// Network" / "Localhost"; genuinely unresolvable public IPs surface as
// `unmapped_count` rather than being misattributed to a continent. The
// resolver is fail-soft: if the GeoIP dataset is absent on a deploy, public
// IPs count as unmapped and the rest of admin analytics still renders.
// **Auth**: AuditAnalyticsReader.
func (s *AnalyticsClient) GetGeoAnalytics(ctx context.Context, params *AnalyticsGetGeoAnalyticsParams) (*GeoMapResponse, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/geo", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out GeoMapResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/geo", Err: err}
	}
	return &out, nil
}

// GetIngestHealth calls GET /analytics/ingest-health.
// Get Ingest Health
//
// Per-source-service ingest health rollup. **When to use**: ops dashboard,
// on-call triage ("which services stopped emitting?"), and the T33
// "first-customer" self-onboarding sanity check. **Returns**: one row per
// `service` seen in the lookback window with count, last-seen, per-minute
// average + p95 throughput, outcome breakdown, and a freshness flag (ok /
// stale / down). **Tip**: pass
// `expected_services=auth-service,billing-service,...` to get an explicit
// `expected_missing` list for services that should be emitting but aren't.
func (s *AnalyticsClient) GetIngestHealth(ctx context.Context, params *AnalyticsGetIngestHealthParams) (*IngestHealthResponse, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/ingest-health", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out IngestHealthResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/ingest-health", Err: err}
	}
	return &out, nil
}

// GetSecurityAnalytics calls GET /analytics/security.
// Get Security Analytics
//
// Security metrics and threat detection. Monitor failed logins, suspicious
// activity, and anomalies. **When to use**: Security monitoring, incident
// investigation, identifying attack patterns. **Returns**: Counts of failed
// logins, rate limit hits, anomalies, plus list of specific security
// alerts. **Alerts**: Includes user_id, event type, and details for each
// suspicious activity. **Next step**: Investigate alerts with GET
// /logs/events filtering by user_id or event_type from alerts. **Tip**: Run
// daily to establish baseline, then monitor for deviations.
func (s *AnalyticsClient) GetSecurityAnalytics(ctx context.Context, params *AnalyticsGetSecurityAnalyticsParams) (*SecurityMetrics, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/security", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out SecurityMetrics
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/security", Err: err}
	}
	return &out, nil
}

// GetTopUsersAnalytics calls GET /analytics/users/top.
// Get Top Users Analytics
//
// Identify most active users by event volume. Find power users or
// investigate unusual activity. **When to use**: Usage analysis,
// identifying super users, spotting automated/bot accounts. **Returns**:
// Ranked list of users with total events, success/fail counts, and activity
// breakdown. **Sorting**: Most active first. Shows both successful and
// failed actions. **Next step**: Investigate specific user with GET
// /logs/events?user_id={user_id}. **Tip**: Compare top users across time
// periods to spot changes in usage patterns.
func (s *AnalyticsClient) GetTopUsersAnalytics(ctx context.Context, params *AnalyticsGetTopUsersAnalyticsParams) (*TopUsersResponse, error) {
	b, err := s.c.do(ctx, "GET", "/analytics/users/top", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out TopUsersResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/analytics/users/top", Err: err}
	}
	return &out, nil
}
