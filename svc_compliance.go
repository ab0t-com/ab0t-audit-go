// ANNOTATION BLOCK
// File: svc_compliance.go — the `Compliance` designation of the Audit Service.
// Why it exists: typed methods for every Compliance endpoint, reached via
// c.Compliance.<Method>. Each is context-first, takes typed path/query/body
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

// ComplianceAddComplianceRequestCommentParams holds the query parameters for Compliance.AddComplianceRequestComment.
type ComplianceAddComplianceRequestCommentParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceAddComplianceRequestCommentParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceCreateControlMappingParams holds the query parameters for Compliance.CreateControlMapping.
type ComplianceCreateControlMappingParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceCreateControlMappingParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceDeleteControlMappingParams holds the query parameters for Compliance.DeleteControlMapping.
type ComplianceDeleteControlMappingParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceDeleteControlMappingParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceDownloadComplianceReportParams holds the query parameters for Compliance.DownloadComplianceReport.
type ComplianceDownloadComplianceReportParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID (defaults to caller's org)
}

func (p *ComplianceDownloadComplianceReportParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceGenerateComplianceReportParams holds the query parameters for Compliance.GenerateComplianceReport.
type ComplianceGenerateComplianceReportParams struct {
	OrgID       *string             `json:"org_id,omitempty"`       // Organization ID
	Standard    *ComplianceStandard `json:"standard,omitempty"`     // required. Standard: GDPR, SOC2, HIPAA, or ISO27001
	PeriodStart *string             `json:"period_start,omitempty"` // required. Reporting period start
	PeriodEnd   *string             `json:"period_end,omitempty"`   // required. Reporting period end
}

func (p *ComplianceGenerateComplianceReportParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Standard != nil {
		v.Set("standard", string(*p.Standard))
	}
	if p.PeriodStart != nil {
		v.Set("period_start", *p.PeriodStart)
	}
	if p.PeriodEnd != nil {
		v.Set("period_end", *p.PeriodEnd)
	}
	return v
}

// ComplianceGetComplianceControlsParams holds the query parameters for Compliance.GetComplianceControls.
type ComplianceGetComplianceControlsParams struct {
	OrgID     *string `json:"org_id,omitempty"`    // Organization ID
	Framework *string `json:"framework,omitempty"` // Filter by framework
	Status    *string `json:"status,omitempty"`    // Filter by status (pass/fail/warning)
}

func (p *ComplianceGetComplianceControlsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Framework != nil {
		v.Set("framework", *p.Framework)
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	return v
}

// ComplianceGetComplianceDashboardParams holds the query parameters for Compliance.GetComplianceDashboard.
type ComplianceGetComplianceDashboardParams struct {
	OrgID     *string `json:"org_id,omitempty"`    // Organization ID
	Framework *string `json:"framework,omitempty"` // Filter by framework (SOC2, GDPR, HIPAA)
}

func (p *ComplianceGetComplianceDashboardParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Framework != nil {
		v.Set("framework", *p.Framework)
	}
	return v
}

// ComplianceGetCompliancePoliciesParams holds the query parameters for Compliance.GetCompliancePolicies.
type ComplianceGetCompliancePoliciesParams struct {
	OrgID    *string             `json:"org_id,omitempty"`   // Organization ID
	Standard *ComplianceStandard `json:"standard,omitempty"` // Filter by standard (GDPR, SOC2, HIPAA, ISO27001)
}

func (p *ComplianceGetCompliancePoliciesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Standard != nil {
		v.Set("standard", string(*p.Standard))
	}
	return v
}

// ComplianceGetComplianceReportParams holds the query parameters for Compliance.GetComplianceReport.
type ComplianceGetComplianceReportParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID (defaults to caller's org)
}

func (p *ComplianceGetComplianceReportParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceGetComplianceRequestStatusParams holds the query parameters for Compliance.GetComplianceRequestStatus.
type ComplianceGetComplianceRequestStatusParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceGetComplianceRequestStatusParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceGetSOC2AuditTrailParams holds the query parameters for Compliance.GetSOC2AuditTrail.
type ComplianceGetSOC2AuditTrailParams struct {
	OrgID            *string `json:"org_id,omitempty"`            // Organization ID
	StartDate        *string `json:"start_date,omitempty"`        // required. Audit period start (e.g. fiscal year start)
	EndDate          *string `json:"end_date,omitempty"`          // required. Audit period end (e.g. fiscal year end)
	ControlObjective *string `json:"control_objective,omitempty"` // Filter by SOC2 control (CC6.1, CC7.2, etc)
}

func (p *ComplianceGetSOC2AuditTrailParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.StartDate != nil {
		v.Set("start_date", *p.StartDate)
	}
	if p.EndDate != nil {
		v.Set("end_date", *p.EndDate)
	}
	if p.ControlObjective != nil {
		v.Set("control_objective", *p.ControlObjective)
	}
	return v
}

// ComplianceListComplianceReportsParams holds the query parameters for Compliance.ListComplianceReports.
type ComplianceListComplianceReportsParams struct {
	OrgID    *string `json:"org_id,omitempty"`   // Organization ID (defaults to caller's org)
	Standard *string `json:"standard,omitempty"` // Filter by standard (SOC2/GDPR/HIPAA/ISO27001/...)
	Status   *string `json:"status,omitempty"`   // Filter by status (generating/completed/failed/unsupported)
	Limit    *int    `json:"limit,omitempty"`    // Page size
	Offset   *int    `json:"offset,omitempty"`   // Skip first N rows
}

func (p *ComplianceListComplianceReportsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Standard != nil {
		v.Set("standard", *p.Standard)
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", strconv.Itoa(*p.Offset))
	}
	return v
}

// ComplianceListComplianceRequestsParams holds the query parameters for Compliance.ListComplianceRequests.
type ComplianceListComplianceRequestsParams struct {
	OrgID          *string `json:"org_id,omitempty"`          // Organization ID (defaults to caller's org)
	Status         *string `json:"status,omitempty"`          // Filter by status (pending/in_review/in_progress/completed/rejected)
	RequestType    *string `json:"request_type,omitempty"`    // Filter by type (access/deletion/portability/rectification)
	DeadlineBefore *string `json:"deadline_before,omitempty"` // Only requests with completion_deadline <= this
	Limit          *int    `json:"limit,omitempty"`           // Page size
	Offset         *int    `json:"offset,omitempty"`          // Skip first N rows
}

func (p *ComplianceListComplianceRequestsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	if p.RequestType != nil {
		v.Set("request_type", *p.RequestType)
	}
	if p.DeadlineBefore != nil {
		v.Set("deadline_before", *p.DeadlineBefore)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", strconv.Itoa(*p.Offset))
	}
	return v
}

// ComplianceListControlMappingsParams holds the query parameters for Compliance.ListControlMappings.
type ComplianceListControlMappingsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceListControlMappingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceListReportDeliveriesParams holds the query parameters for Compliance.ListReportDeliveries.
type ComplianceListReportDeliveriesParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceListReportDeliveriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// CompliancePatchComplianceRequestStatusParams holds the query parameters for Compliance.PatchComplianceRequestStatus.
type CompliancePatchComplianceRequestStatusParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *CompliancePatchComplianceRequestStatusParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceRequestDataDeletionParams holds the query parameters for Compliance.RequestDataDeletion.
type ComplianceRequestDataDeletionParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceRequestDataDeletionParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceRequestDataSubjectAccessParams holds the query parameters for Compliance.RequestDataSubjectAccess.
type ComplianceRequestDataSubjectAccessParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceRequestDataSubjectAccessParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceVerifyIntegrityEventParams holds the query parameters for Compliance.VerifyIntegrityEvent.
type ComplianceVerifyIntegrityEventParams struct {
	Table *string `json:"table,omitempty"`  // Chained table to verify
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceVerifyIntegrityEventParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Table != nil {
		v.Set("table", *p.Table)
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// AddComplianceRequestComment calls POST /compliance/requests/{request_id}/comment.
// Add Compliance Request Comment
//
// Append a processing note to a DSAR. Backs the "Add note" affordance on
// the DSAR detail page. **When to use**: leaving an audit-trail comment
// during DSAR processing. **Auth**: AuditComplianceProcessor. **Behavior**:
// appends `[YYYY-MM-DD HH:MM UTC actor] note` to the existing
// `processing_notes` string column. Notes are append-only — there's no
// edit/delete on processing_notes (intentional, audit posture).
func (s *ComplianceClient) AddComplianceRequestComment(ctx context.Context, requestID string, body DsarCommentBody, params *ComplianceAddComplianceRequestCommentParams) (*DsarCommentResponse, error) {
	path := fmt.Sprintf("/compliance/requests/%s/comment", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	var out DsarCommentResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// CreateControlMapping calls POST /compliance/control-mappings.
// Create Control Mapping
//
// Add a per-org event_type→controls mapping override. **When to use**:
// control-mappings editor "Add mapping" button. **Auth**:
// AuditControlMappingsManager (audit.control_mappings.manage).
// **Behavior**: validates input, inserts a new row into
// compliance_control_mappings for the caller's org, returns the created
// row.
func (s *ComplianceClient) CreateControlMapping(ctx context.Context, body ControlMappingCreate, params *ComplianceCreateControlMappingParams) (*ControlMappingRow, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/control-mappings", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out ControlMappingRow
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/control-mappings", Err: err}
	}
	return &out, nil
}

// DeleteControlMapping calls DELETE /compliance/control-mappings/{mapping_id}.
// Delete Control Mapping
//
// Soft-delete a per-org control-mapping override by id. **When to use**:
// control-mappings editor "Delete" button. **Auth**:
// AuditControlMappingsManager (audit.control_mappings.manage).
// **Behavior**: marks the row deleted=1 in the caller's org (idempotent — a
// missing/already-deleted row still returns 204 without leaking existence
// to other orgs).
func (s *ComplianceClient) DeleteControlMapping(ctx context.Context, mappingID string, params *ComplianceDeleteControlMappingParams) error {
	path := fmt.Sprintf("/compliance/control-mappings/%s", url.PathEscape(mappingID))
	_, err := s.c.do(ctx, "DELETE", path, params.values(), nil)
	return err
}

// DownloadComplianceReport calls GET /compliance/reports/{report_id}/download.
// Download Compliance Report
//
// Download the rendered artifact (HTML) for a completed compliance report.
// **When to use**: after the report shows status='completed', retrieve the
// auditor-ready file to hand off, archive, or print to PDF. **Returns**:
// the rendered report as an HTML file download. **Errors**: 404 if the
// report or its file is gone; 409 if it isn't finished yet (poll GET
// /compliance/reports/{id} until status='completed'). **Auth**:
// AuditComplianceReader. Same per-report org authorization as viewing.
func (s *ComplianceClient) DownloadComplianceReport(ctx context.Context, reportID string, params *ComplianceDownloadComplianceReportParams) (RawMessage, error) {
	path := fmt.Sprintf("/compliance/reports/%s/download", url.PathEscape(reportID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// GenerateComplianceReport calls POST /compliance/reports/generate.
// Generate Compliance Report
//
// Generate comprehensive compliance attestation report. Formatted for
// auditors and regulators. **When to use**: Audit season, compliance
// certification, regulatory submissions, board reporting. **Returns**:
// report_id immediately. Report generated asynchronously (5-30 minutes
// depending on data volume). **Report includes**: Event summaries, access
// patterns, security metrics, data retention compliance, control
// effectiveness. **Workflow**: Generate report → poll GET
// /compliance/reports/{report_id} until status='completed' → view it in-app
// and/or download the rendered HTML via GET
// /compliance/reports/{report_id}/download. **Next step**: Monitor status,
// then submit report to auditors or compliance team.
func (s *ComplianceClient) GenerateComplianceReport(ctx context.Context, params *ComplianceGenerateComplianceReportParams) (*ComplianceReportAck, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/reports/generate", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ComplianceReportAck
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/reports/generate", Err: err}
	}
	return &out, nil
}

// GetComplianceControls calls GET /compliance/controls.
// Get Compliance Controls
//
// List all compliance controls with their current status. **When to use**:
// Deep-dive into specific controls, audit preparation, remediation
// planning. **Returns**: List of controls with event counts, pass/fail
// status, evidence samples. **Filters**: By framework (SOC2, GDPR), by
// status (pass, fail, warning). **Next step**: Use event_ids to investigate
// specific control gaps.
func (s *ComplianceClient) GetComplianceControls(ctx context.Context, params *ComplianceGetComplianceControlsParams) (*ComplianceControlsListResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/controls", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ComplianceControlsListResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/controls", Err: err}
	}
	return &out, nil
}

// GetComplianceDashboard calls GET /compliance/dashboard.
// Get Compliance Dashboard
//
// Get real-time compliance posture dashboard. Shows which controls are
// passing/failing, audit readiness score, and open findings. **When to
// use**: Daily compliance monitoring, pre-audit checks, executive
// reporting. **Returns**: Compliance score (0-100), controls status,
// framework breakdown, top findings. **Refresh**: Data updates in real-time
// based on latest events. **Next step**: Click into specific failing
// controls to see remediation steps. Principle 2: Pipeline architecture
// (gather data → calculate → format → return) Principle 8:
// Performance-aware (< 2 second response time target)
func (s *ComplianceClient) GetComplianceDashboard(ctx context.Context, params *ComplianceGetComplianceDashboardParams) (*ComplianceDashboardResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/dashboard", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ComplianceDashboardResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/dashboard", Err: err}
	}
	return &out, nil
}

// GetCompliancePolicies calls GET /compliance/policies.
// Get Compliance Policies
//
// View active compliance policies and data retention rules. Understand what
// regulations apply. **When to use**: Compliance review, understanding data
// retention, checking GDPR/SOC2 status. **Returns**: List of policies with
// retention periods, enabled standards, and audit settings. **Next step**:
// Use policies to guide data requests and retention workflows.
func (s *ComplianceClient) GetCompliancePolicies(ctx context.Context, params *ComplianceGetCompliancePoliciesParams) ([]CompliancePolicy, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/policies", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out []CompliancePolicy
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/policies", Err: err}
	}
	return out, nil
}

// GetComplianceReport calls GET /compliance/reports/{report_id}.
// Get Compliance Report
//
// Fetch one previously-generated compliance report, including its full
// body. **When to use**: open a completed report in-app (the report viewer
// on /ui/app/reports), to read control coverage, evidence counts,
// chain-integrity and the reporting period — without leaving the product.
// **Returns**: the report's metadata PLUS the parsed `summary` body (the
// real per-standard generator output stored in ClickHouse) under `summary`,
// and a working `report_file_url` download link when the report has
// completed. **Status semantics**: `generating` (still running — `summary`
// is empty), `completed` (body + download ready), `failed`/`unsupported`
// (see `error_message`). **Auth**: AuditComplianceReader. Org-scoped to the
// caller.
func (s *ComplianceClient) GetComplianceReport(ctx context.Context, reportID string, params *ComplianceGetComplianceReportParams) (RawMessage, error) {
	path := fmt.Sprintf("/compliance/reports/%s", url.PathEscape(reportID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// GetComplianceRequestStatus calls GET /compliance/requests/{request_id}.
// Get Compliance Request Status
//
// Check status of GDPR data access or deletion request. Poll this to track
// progress. **When to use**: After POST /gdpr/data-subject-access or
// /gdpr/right-to-be-forgotten, monitor completion. **Returns**: Request
// details including status (pending/processing/completed/failed), deadline,
// and download URL when ready. **Polling**: Check every 10 seconds until
// status changes from pending/processing. **Deadline**: Shows time
// remaining for GDPR 30-day requirement. **Next step**: When completed,
// download data package from returned URL or confirm deletion.
func (s *ComplianceClient) GetComplianceRequestStatus(ctx context.Context, requestID string, params *ComplianceGetComplianceRequestStatusParams) (*ComplianceRequestStatus, error) {
	path := fmt.Sprintf("/compliance/requests/%s", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ComplianceRequestStatus
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetSOC2AuditTrail calls GET /compliance/soc2/audit-trail.
// Get Soc2 Audit Trail
//
// SOC2 Type II: Generate audit trail evidence for controls testing. Prove
// continuous compliance. **When to use**: SOC2 audit preparation, auditor
// evidence requests, control effectiveness demonstration. **Returns**:
// Complete audit trail with timestamps, actors, actions, and outcomes.
// Includes tamper-evident hashes. **Period**: Typically matches SOC2 audit
// period (6-12 months). Must show continuous operation. **Control
// objectives**: Filter by specific Trust Service Criteria (CC6.1: logical
// access, CC7.2: security monitoring). **Next step**: Provide to auditors.
// Export with POST /export/ for offline review if needed.
func (s *ComplianceClient) GetSOC2AuditTrail(ctx context.Context, params *ComplianceGetSOC2AuditTrailParams) (*Soc2AuditTrailResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/soc2/audit-trail", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out Soc2AuditTrailResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/soc2/audit-trail", Err: err}
	}
	return &out, nil
}

// ListComplianceReports calls GET /compliance/reports.
// List Compliance Reports
//
// List the org's previously generated compliance reports, newest first.
// **When to use**: the "Recently generated" table on /ui/app/reports, to
// see which SOC2/GDPR/HIPAA reports have been run, their status, and
// download links. **Returns**: paginated ComplianceReportRow array from the
// real `compliance_reports` store — id, standard, status, period,
// requested_by, created_at, and report_file_url (download) when the report
// has completed. **Filters**: `standard`, `status` (both optional).
// Org-scoped to the caller. **Auth**: AuditComplianceReader. **Next step**:
// open GET /compliance/reports/{report_id} to read the full body in-app, or
// follow report_file_url (GET /compliance/reports/{report_id}/download) to
// download the rendered report once status is `completed`.
func (s *ComplianceClient) ListComplianceReports(ctx context.Context, params *ComplianceListComplianceReportsParams) (*PaginatedResponseComplianceReportRow, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/reports", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseComplianceReportRow
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/reports", Err: err}
	}
	return &out, nil
}

// ListComplianceRequests calls GET /compliance/requests.
// List Compliance Requests
//
// List GDPR data-subject requests for an org. Backs the
// /ui/app/compliance/dsar queue page. **When to use**: DSAR queue triage,
// deadline overview, status filtering. **Returns**: Paginated
// DsarRequestRow array, newest first. **Filters**: status, request_type,
// deadline_before. All optional; omit to list everything. **Next step**:
// Click any row to open GET /compliance/requests/{request_id}.
func (s *ComplianceClient) ListComplianceRequests(ctx context.Context, params *ComplianceListComplianceRequestsParams) (*PaginatedResponseDsarRequestRow, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/requests", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseDsarRequestRow
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/requests", Err: err}
	}
	return &out, nil
}

// ListControlMappings calls GET /compliance/control-mappings.
// List Control Mappings
//
// List an org's *custom* control-mapping overrides (additions on top of the
// hardcoded `app/core/control_mapper.py` defaults). **When to use**:
// control-mappings editor page load. **Auth**: AuditComplianceReader.
// **Behavior**: returns soft-delete-filtered rows for the caller's org. The
// hardcoded defaults are rendered client-side and are NOT returned here.
func (s *ComplianceClient) ListControlMappings(ctx context.Context, params *ComplianceListControlMappingsParams) ([]ControlMappingRow, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/control-mappings", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out []ControlMappingRow
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/control-mappings", Err: err}
	}
	return out, nil
}

// ListReportDeliveries calls GET /compliance/reports/{report_id}/deliveries.
// List Report Deliveries
//
// List per-recipient email deliveries for a compliance report. **When to
// use**: schedule UI's "Last delivery" tooltip, debug a "did the auditor
// get the report?" question, audit-trail GDPR Article 30 inquiries about
// who-was-notified-when. **Returns**: list of attempts ordered
// newest-first. Each attempt is immutable — a retry would create a new row,
// not modify the old one. **Auth**: `AuditComplianceReader`.
func (s *ComplianceClient) ListReportDeliveries(ctx context.Context, reportID string, params *ComplianceListReportDeliveriesParams) (RawMessage, error) {
	path := fmt.Sprintf("/compliance/reports/%s/deliveries", url.PathEscape(reportID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// PatchComplianceRequestStatus calls PATCH /compliance/requests/{request_id}/status.
// Patch Compliance Request Status
//
// Advance a DSAR's status (and optionally attach a comment). The previous
// status is appended to `status_history` with a timestamp + actor. **When
// to use**: DSAR detail page "Advance to <next>" button. **Auth**:
// AuditComplianceProcessor. **Behavior**: validates the target row exists
// in the caller's org, then runs an atomic ALTER TABLE UPDATE that bumps
// `status`, `updated_at`, and appends a JSON entry to `status_history`.
// Returns the new history length.
func (s *ComplianceClient) PatchComplianceRequestStatus(ctx context.Context, requestID string, body DsarStatusPatchBody, params *CompliancePatchComplianceRequestStatusParams) (*DsarStatusPatchResponse, error) {
	path := fmt.Sprintf("/compliance/requests/%s/status", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "PATCH", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	var out DsarStatusPatchResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// RequestDataDeletion calls POST /compliance/gdpr/right-to-be-forgotten.
// Request Data Deletion
//
// GDPR Article 17: Delete user's personal data permanently. Right to be
// forgotten. **When to use**: User requests data deletion, GDPR erasure
// requests, account closure. **Returns**: request_id immediately. Deletion
// executed asynchronously. Irreversible once complete. **Required**: Must
// provide legal_basis and reason for audit trail. **Workflow**: Submit
// request → Poll GET /compliance/requests/{id} → Confirm deletion complete.
// **Warning**: Deletion is permanent. Retained audit logs use pseudonymized
// user_id only. **Next step**: Monitor status with GET
// /compliance/requests/{request_id}.
func (s *ComplianceClient) RequestDataDeletion(ctx context.Context, body DataSubjectDeletionRequest, params *ComplianceRequestDataDeletionParams) (*DataDeletionAck, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/gdpr/right-to-be-forgotten", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out DataDeletionAck
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/gdpr/right-to-be-forgotten", Err: err}
	}
	return &out, nil
}

// RequestDataSubjectAccess calls POST /compliance/gdpr/data-subject-access.
// Request Data Subject Access
//
// GDPR Article 15: Fulfill user's right to access their personal data.
// Generate data package. **When to use**: User requests their data, GDPR
// Article 15 compliance, data portability. **Returns**: request_id
// immediately. Data package generated asynchronously within 30 days.
// **Workflow**: Submit request → Poll GET /compliance/requests/{id} →
// Download when ready. **Deadline**: GDPR requires response within 30 days.
// Status shows time remaining. **Next step**: Monitor request status with
// GET /compliance/requests/{request_id}.
func (s *ComplianceClient) RequestDataSubjectAccess(ctx context.Context, body DataSubjectAccessRequest, params *ComplianceRequestDataSubjectAccessParams) (*DataSubjectAccessAck, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/gdpr/data-subject-access", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out DataSubjectAccessAck
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/gdpr/data-subject-access", Err: err}
	}
	return &out, nil
}

// VerifyIntegrityChain calls POST /compliance/integrity/verify.
// Verify Integrity Chain
//
// Re-walk an org's hash chain for a chained table and report the FIRST
// break. For each row (ordered by timestamp/event_id ascending) two checks
// run: (a) the recomputed canonical hash equals the stored “event_hash“,
// and (b) the row's “previous_event_hash“ links to the prior row's stored
// hash. On the first failure the walk stops and returns the break location;
// a clean walk returns the head/tail hashes and the verified count.
func (s *ComplianceClient) VerifyIntegrityChain(ctx context.Context, body IntegrityVerifyRequest) (RawMessage, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/integrity/verify", nil, body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// VerifyIntegrityEvent calls GET /compliance/integrity/event/{event_id}.
// Verify Integrity Event
//
// Verify a single event's stored hash against a freshly recomputed
// canonical hash. Returns the stored vs recomputed hash and a boolean
// verdict.
func (s *ComplianceClient) VerifyIntegrityEvent(ctx context.Context, eventID string, params *ComplianceVerifyIntegrityEventParams) (RawMessage, error) {
	path := fmt.Sprintf("/compliance/integrity/event/%s", url.PathEscape(eventID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}
