// ANNOTATION BLOCK
// File: svc_admin.go — the `Admin` designation of the Audit Service.
// Why it exists: typed methods for every Admin endpoint, reached via
// c.Admin.<Method>. Each is context-first, takes typed path/query/body
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

// AdminApproveDeletionRequestParams holds the query parameters for Admin.ApproveDeletionRequest.
type AdminApproveDeletionRequestParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *AdminApproveDeletionRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AdminGetDeletionRequestParams holds the query parameters for Admin.GetDeletionRequest.
type AdminGetDeletionRequestParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *AdminGetDeletionRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AdminGetOrganizationsSummaryParams holds the query parameters for Admin.GetOrganizationsSummary.
type AdminGetOrganizationsSummaryParams struct {
	Limit  *int `json:"limit,omitempty"`  // Max organizations to return
	Offset *int `json:"offset,omitempty"` // Skip first N organizations
}

func (p *AdminGetOrganizationsSummaryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", strconv.Itoa(*p.Offset))
	}
	return v
}

// AdminGetRetentionPoliciesParams holds the query parameters for Admin.GetRetentionPolicies.
type AdminGetRetentionPoliciesParams struct {
	OrgID *string `json:"org_id,omitempty"` // Optional: Get policy for specific org
}

func (p *AdminGetRetentionPoliciesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AdminListDeletionRequestsParams holds the query parameters for Admin.ListDeletionRequests.
type AdminListDeletionRequestsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *AdminListDeletionRequestsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AdminOptimizeDatabaseParams holds the query parameters for Admin.OptimizeDatabase.
type AdminOptimizeDatabaseParams struct {
	Table *string `json:"table,omitempty"` // Table name: audit_events or export_jobs
}

func (p *AdminOptimizeDatabaseParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Table != nil {
		v.Set("table", *p.Table)
	}
	return v
}

// AdminRejectDeletionRequestParams holds the query parameters for Admin.RejectDeletionRequest.
type AdminRejectDeletionRequestParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *AdminRejectDeletionRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AdminRestoreDeletionRequestParams holds the query parameters for Admin.RestoreDeletionRequest.
type AdminRestoreDeletionRequestParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *AdminRestoreDeletionRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ApproveDeletionRequest calls POST /admin/deletions/{deletion_id}/approve.
// Approve Deletion Request
//
// Cast one of the two required approvals. After two DISTINCT non-requester
// approvers approve, the quarantined deletion is executed (rows removed
// from audit_events; they remain in audit_events_quarantine until the
// restore window lapses).
func (s *AdminClient) ApproveDeletionRequest(ctx context.Context, deletionID string, body DeletionDecisionRequest, params *AdminApproveDeletionRequestParams) (RawMessage, error) {
	path := fmt.Sprintf("/admin/deletions/%s/approve", url.PathEscape(deletionID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// ExecuteAdhocQuery calls POST /admin/query.
// Execute Adhoc Query
//
// Run custom SQL query for advanced analytics. For power users and
// debugging. **When to use**: Complex analytics not covered by standard
// endpoints, troubleshooting, custom reports. **Security**: Read-only
// (SELECT) queries only. Mutating operations blocked. 30-second timeout,
// 10K row limit. **Returns**: Query results with execution metadata and row
// count. **Example**: SELECT event_type, count(*) FROM audit.audit_events
// WHERE org_id='...' GROUP BY event_type **Next step**: Use results for
// analysis. For recurring queries, request a dedicated endpoint.
// **Warning**: Large result sets truncated. Use filters and LIMIT clause to
// narrow results. **Audit**: All queries logged with user, SQL, and
// execution time for security review.
func (s *AdminClient) ExecuteAdhocQuery(ctx context.Context, body AdHocQueryRequest) (*AdHocQueryResponse, error) {
	b, err := s.c.do(ctx, "POST", "/admin/query", nil, body)
	if err != nil {
		return nil, err
	}
	var out AdHocQueryResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/query", Err: err}
	}
	return &out, nil
}

// GetDatabaseMetrics calls GET /admin/metrics/database.
// Get Database Metrics
//
// Database storage and performance metrics. Monitor disk usage and query
// performance. **When to use**: Capacity planning, performance
// optimization, storage cost management. **Returns**: Table sizes,
// compression ratios, row counts, partition info, query performance stats.
// **Next step**: Use metrics to plan scaling, optimize tables, or archive
// old partitions.
func (s *AdminClient) GetDatabaseMetrics(ctx context.Context) (*DatabaseMetricsResponse, error) {
	b, err := s.c.do(ctx, "GET", "/admin/metrics/database", nil, nil)
	if err != nil {
		return nil, err
	}
	var out DatabaseMetricsResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/metrics/database", Err: err}
	}
	return &out, nil
}

// GetDeletionRequest calls GET /admin/deletions/{deletion_id}.
// Get Deletion Request
//
// Inspect a single deletion request together with its approvals.
func (s *AdminClient) GetDeletionRequest(ctx context.Context, deletionID string, params *AdminGetDeletionRequestParams) (RawMessage, error) {
	path := fmt.Sprintf("/admin/deletions/%s", url.PathEscape(deletionID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// GetOrganizationsSummary calls GET /admin/organizations.
// Get Organizations Summary
//
// List all organizations with usage statistics. Platform-wide tenant
// overview. **When to use**: Platform administration, usage analysis,
// identifying top customers. **Returns**: Organization list with event
// counts, unique users, first/last event dates. **Next step**: Drill into
// specific org with GET /logs/events?org_id={org_id}.
func (s *AdminClient) GetOrganizationsSummary(ctx context.Context, params *AdminGetOrganizationsSummaryParams) (*PaginatedResponseOrganizationSummary, error) {
	b, err := s.c.do(ctx, "GET", "/admin/organizations", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseOrganizationSummary
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/organizations", Err: err}
	}
	return &out, nil
}

// GetRetentionPolicies calls GET /admin/retention/policies.
// Get Retention Policies
//
// View data retention and archival policies. Understand how long data is
// kept. **When to use**: Compliance planning, storage management,
// understanding data lifecycle. **Returns**: Retention days, archive
// timing, hot/cold storage thresholds. **Next step**: Adjust policies (PUT
// /admin/retention/policies) or plan storage capacity based on retention
// periods.
func (s *AdminClient) GetRetentionPolicies(ctx context.Context, params *AdminGetRetentionPoliciesParams) (*RetentionPoliciesResponse, error) {
	b, err := s.c.do(ctx, "GET", "/admin/retention/policies", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out RetentionPoliciesResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/retention/policies", Err: err}
	}
	return &out, nil
}

// GetSystemStatus calls GET /admin/status.
// Get System Status
//
// System health and operational metrics. Monitor service performance and
// dependencies. **When to use**: Operations monitoring, troubleshooting,
// capacity planning. **Returns**: Service version, uptime,
// database/cache/queue status, background task health, performance metrics.
// **Next step**: If unhealthy components found, check logs or restart
// affected services.
func (s *AdminClient) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	b, err := s.c.do(ctx, "GET", "/admin/status", nil, nil)
	if err != nil {
		return nil, err
	}
	var out SystemStatus
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/status", Err: err}
	}
	return &out, nil
}

// ListDeletionRequests calls GET /admin/deletions.
// List Deletion Requests
//
// List pending + historical deletion requests (current state) for the org.
func (s *AdminClient) ListDeletionRequests(ctx context.Context, params *AdminListDeletionRequestsParams) (RawMessage, error) {
	b, err := s.c.do(ctx, "GET", "/admin/deletions", params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// OptimizeDatabase calls POST /admin/maintenance/optimize.
// Optimize Database
//
// Optimize database table storage. Merge small parts and deduplicate for
// better performance. **When to use**: After large data purges, slow
// queries, storage optimization, maintenance windows. **Returns**:
// Optimization status and duration. **Impact**: Temporarily increases
// CPU/disk during optimization. Schedule during low-traffic periods. **Next
// step**: Monitor query performance improvement with GET
// /admin/metrics/database.
func (s *AdminClient) OptimizeDatabase(ctx context.Context, params *AdminOptimizeDatabaseParams) (*MaintenanceOptimizeResponse, error) {
	b, err := s.c.do(ctx, "POST", "/admin/maintenance/optimize", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out MaintenanceOptimizeResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/maintenance/optimize", Err: err}
	}
	return &out, nil
}

// PurgeOldData calls POST /admin/data/purge.
// Purge Old Data
//
// Permanently delete old audit data. Use for decommissioning orgs or
// emergency cleanup. **When to use**: Tenant offboarding, emergency storage
// cleanup, compliance-required deletion. **Warning**: IRREVERSIBLE. Always
// run with dry_run=true first to preview deletions. **Returns**: Count of
// records that will be (or were) deleted. **Next step**: Review dry_run
// results carefully before executing with dry_run=false.
func (s *AdminClient) PurgeOldData(ctx context.Context, body DataPurgeRequest) (*PurgeDataResponse, error) {
	b, err := s.c.do(ctx, "POST", "/admin/data/purge", nil, body)
	if err != nil {
		return nil, err
	}
	var out PurgeDataResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/data/purge", Err: err}
	}
	return &out, nil
}

// PutRetentionPolicies calls PUT /admin/retention/policies.
// Put Retention Policies
//
// Set the per-org data retention policy. Governs how long audit data is
// kept. **When to use**: an admin needs to change
// retention/archive/hot-data windows for an org without a redeploy.
// **Auth**: “audit.maintenance.admin“ (not-suspended) — the same role
// that governs purge/maintenance, since shrinking retention causes
// deletion. **Behavior**: upserts the single per-org row in
// “settings_retention“. The retention manager picks up the new windows on
// its next pass; nothing is deleted synchronously here. Returns the
// persisted policy. **Next step**: review GET /admin/retention/policies to
// confirm.
func (s *AdminClient) PutRetentionPolicies(ctx context.Context, body RetentionPolicyUpdate) (*RetentionPoliciesResponse, error) {
	b, err := s.c.do(ctx, "PUT", "/admin/retention/policies", nil, body)
	if err != nil {
		return nil, err
	}
	var out RetentionPoliciesResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/admin/retention/policies", Err: err}
	}
	return &out, nil
}

// RejectDeletionRequest calls POST /admin/deletions/{deletion_id}/reject.
// Reject Deletion Request
//
// Reject a pending deletion. The live data was never touched; the
// quarantined copy for this deletion_id is dropped and the request is
// marked 'rejected'.
func (s *AdminClient) RejectDeletionRequest(ctx context.Context, deletionID string, body DeletionDecisionRequest, params *AdminRejectDeletionRequestParams) (RawMessage, error) {
	path := fmt.Sprintf("/admin/deletions/%s/reject", url.PathEscape(deletionID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// RestoreDeletionRequest calls POST /admin/deletions/{deletion_id}/restore.
// Restore Deletion Request
//
// Restore quarantined rows back into audit_events. Valid only while status
// in ('pending_approval','approved','executed') AND now < restore_deadline.
// Idempotent: rows already present (by event_id) are skipped.
func (s *AdminClient) RestoreDeletionRequest(ctx context.Context, deletionID string, body DeletionRestoreRequest, params *AdminRestoreDeletionRequestParams) (RawMessage, error) {
	path := fmt.Sprintf("/admin/deletions/%s/restore", url.PathEscape(deletionID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}
