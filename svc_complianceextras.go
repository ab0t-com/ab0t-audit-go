// ANNOTATION BLOCK
// File: svc_complianceextras.go — the `ComplianceExtras` designation of the Audit Service.
// Why it exists: typed methods for every ComplianceExtras endpoint, reached via
// c.ComplianceExtras.<Method>. Each is context-first, takes typed path/query/body
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

// ComplianceExtrasCreateAuditorRequestParams holds the query parameters for ComplianceExtras.CreateAuditorRequest.
type ComplianceExtrasCreateAuditorRequestParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasCreateAuditorRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasDeleteAuditPeriodParams holds the query parameters for ComplianceExtras.DeleteAuditPeriod.
type ComplianceExtrasDeleteAuditPeriodParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasDeleteAuditPeriodParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasGetAuditPeriodParams holds the query parameters for ComplianceExtras.GetAuditPeriod.
type ComplianceExtrasGetAuditPeriodParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasGetAuditPeriodParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasGetPostureHistoryParams holds the query parameters for ComplianceExtras.GetPostureHistory.
type ComplianceExtrasGetPostureHistoryParams struct {
	WindowDays *int    `json:"window_days,omitempty"`
	OrgID      *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasGetPostureHistoryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.WindowDays != nil {
		v.Set("window_days", strconv.Itoa(*p.WindowDays))
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasListAuditorRequestsParams holds the query parameters for ComplianceExtras.ListAuditorRequests.
type ComplianceExtrasListAuditorRequestsParams struct {
	OrgID *string              `json:"org_id,omitempty"`
	State *AuditorRequestState `json:"state,omitempty"`
	Limit *int                 `json:"limit,omitempty"`
}

func (p *ComplianceExtrasListAuditorRequestsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.State != nil {
		v.Set("state", string(*p.State))
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceExtrasPatchAuditorRequestParams holds the query parameters for ComplianceExtras.PatchAuditorRequest.
type ComplianceExtrasPatchAuditorRequestParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasPatchAuditorRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasPutAuditPeriodParams holds the query parameters for ComplianceExtras.PutAuditPeriod.
type ComplianceExtrasPutAuditPeriodParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasPutAuditPeriodParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceExtrasSatisfyAuditorRequestParams holds the query parameters for ComplianceExtras.SatisfyAuditorRequest.
type ComplianceExtrasSatisfyAuditorRequestParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceExtrasSatisfyAuditorRequestParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// CreateAuditorRequest calls POST /compliance/auditor-requests.
// Create Auditor Request
//
// Log a new auditor question. Write tier — compliance reviewer or above.
func (s *ComplianceExtrasClient) CreateAuditorRequest(ctx context.Context, body AuditorRequestCreate, params *ComplianceExtrasCreateAuditorRequestParams) (*AuditorRequest, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/auditor-requests", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out AuditorRequest
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/auditor-requests", Err: err}
	}
	return &out, nil
}

// DeleteAuditPeriod calls DELETE /settings/audit-period.
// Delete Audit Period
//
// Clear the org's active audit period. Implementation: inserts a tombstone
// row (empty name, epoch dates) so the ReplacingMergeTree(set_at) picks it
// as the latest. The GET handler treats the tombstone as "no period set"
// and returns 404. Idempotent: re-running after a previous clear just
// inserts another tombstone (newer set_at). Auth: AuditMaintenanceAdmin.
func (s *ComplianceExtrasClient) DeleteAuditPeriod(ctx context.Context, params *ComplianceExtrasDeleteAuditPeriodParams) error {
	_, err := s.c.do(ctx, "DELETE", "/settings/audit-period", params.values(), nil)
	return err
}

// GetAuditPeriod calls GET /settings/audit-period.
// Get Audit Period
//
// Return the org's currently-active audit period, or 404 if unset.
func (s *ComplianceExtrasClient) GetAuditPeriod(ctx context.Context, params *ComplianceExtrasGetAuditPeriodParams) (*AuditPeriod, error) {
	b, err := s.c.do(ctx, "GET", "/settings/audit-period", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out AuditPeriod
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/audit-period", Err: err}
	}
	return &out, nil
}

// GetPostureHistory calls GET /compliance/posture-history.
// Get Posture History
//
// Bucket-by-day compliance posture over the last `window_days`. The score
// for a day is computed as the fraction of events on that day that carried
// at least one mapped compliance control vs the total event count. It's a
// coarse but honest signal — captures "are we tagging events at all" plus
// "are the frameworks we activated covering what we ingest." Per-framework
// scores are derived from the first segment of each control_id
// (`SOC2.CC6.1` → `SOC2`). Cheaper + more honest than the client-side
// fan-out the page uses today; one query instead of 30. Server-side caching
// opportunity but kept simple in v1.
func (s *ComplianceExtrasClient) GetPostureHistory(ctx context.Context, params *ComplianceExtrasGetPostureHistoryParams) (*PostureHistoryResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/posture-history", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PostureHistoryResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/posture-history", Err: err}
	}
	return &out, nil
}

// ListAuditorRequests calls GET /compliance/auditor-requests.
// List Auditor Requests
//
// List auditor questions for an org. Read tier.
func (s *ComplianceExtrasClient) ListAuditorRequests(ctx context.Context, params *ComplianceExtrasListAuditorRequestsParams) ([]AuditorRequest, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/auditor-requests", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out []AuditorRequest
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/auditor-requests", Err: err}
	}
	return out, nil
}

// PatchAuditorRequest calls PATCH /compliance/auditor-requests/{request_id}.
// Patch Auditor Request
//
// Update an auditor request — typically to advance its state or attach an
// evidence package. Implements optimistic last-write-wins via the
// ReplacingMergeTree(updated_at) semantics on the storage table.
func (s *ComplianceExtrasClient) PatchAuditorRequest(ctx context.Context, requestID string, body AuditorRequestUpdate, params *ComplianceExtrasPatchAuditorRequestParams) (*AuditorRequest, error) {
	path := fmt.Sprintf("/compliance/auditor-requests/%s", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "PATCH", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	var out AuditorRequest
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// PutAuditPeriod calls PUT /settings/audit-period.
// Put Audit Period
//
// Upsert the org's active audit period. Write tier — admin only.
func (s *ComplianceExtrasClient) PutAuditPeriod(ctx context.Context, body AuditPeriodWrite, params *ComplianceExtrasPutAuditPeriodParams) (*AuditPeriod, error) {
	b, err := s.c.do(ctx, "PUT", "/settings/audit-period", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out AuditPeriod
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/audit-period", Err: err}
	}
	return &out, nil
}

// SatisfyAuditorRequest calls POST /compliance/auditor-requests/{request_id}/satisfy.
// Satisfy Auditor Request
//
// Satisfy an auditor request by linking the evidence package that answers
// it and transitioning the request to `satisfied`. This is the durable
// counterpart to the inbox "Mark satisfied" action: it records *which*
// evidence package closed the ask, so the auditor trail shows not just that
// a request was satisfied but with what. Write tier — compliance reviewer
// or above. Fails with 404 if the request doesn't exist in the org, or 422
// if the named evidence package doesn't exist in the org (you can't satisfy
// an ask by pointing at a package that isn't there).
func (s *ComplianceExtrasClient) SatisfyAuditorRequest(ctx context.Context, requestID string, body AuditorRequestSatisfy, params *ComplianceExtrasSatisfyAuditorRequestParams) (*AuditorRequest, error) {
	path := fmt.Sprintf("/compliance/auditor-requests/%s/satisfy", url.PathEscape(requestID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	var out AuditorRequest
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}
