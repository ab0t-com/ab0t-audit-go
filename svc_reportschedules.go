// ANNOTATION BLOCK
// File: svc_reportschedules.go — the `ReportSchedules` designation of the Audit Service.
// Why it exists: typed methods for every ReportSchedules endpoint, reached via
// c.ReportSchedules.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// ReportSchedulesCreateReportScheduleParams holds the query parameters for ReportSchedules.CreateReportSchedule.
type ReportSchedulesCreateReportScheduleParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ReportSchedulesCreateReportScheduleParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ReportSchedulesDeleteReportScheduleParams holds the query parameters for ReportSchedules.DeleteReportSchedule.
type ReportSchedulesDeleteReportScheduleParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ReportSchedulesDeleteReportScheduleParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ReportSchedulesListReportSchedulesParams holds the query parameters for ReportSchedules.ListReportSchedules.
type ReportSchedulesListReportSchedulesParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ReportSchedulesListReportSchedulesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ReportSchedulesPauseReportScheduleParams holds the query parameters for ReportSchedules.PauseReportSchedule.
type ReportSchedulesPauseReportScheduleParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ReportSchedulesPauseReportScheduleParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ReportSchedulesResumeReportScheduleParams holds the query parameters for ReportSchedules.ResumeReportSchedule.
type ReportSchedulesResumeReportScheduleParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ReportSchedulesResumeReportScheduleParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// CreateReportSchedule calls POST /compliance/report-schedules/.
// Create Report Schedule
//
// Create a new scheduled report. The `next_run_at` is computed from the
// cadence string at create time (see `_next_run_at` for the v1 cadence
// grammar).
func (s *ReportSchedulesClient) CreateReportSchedule(ctx context.Context, body ReportScheduleCreate, params *ReportSchedulesCreateReportScheduleParams) (*ReportScheduleItem, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/report-schedules/", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out ReportScheduleItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/report-schedules/", Err: err}
	}
	return &out, nil
}

// DeleteReportSchedule calls DELETE /compliance/report-schedules/{schedule_id}.
// Delete Report Schedule
//
// Soft-delete via tombstone row. Idempotent.
func (s *ReportSchedulesClient) DeleteReportSchedule(ctx context.Context, scheduleID string, params *ReportSchedulesDeleteReportScheduleParams) (*ReportScheduleDeleteResponse, error) {
	path := fmt.Sprintf("/compliance/report-schedules/%s", url.PathEscape(scheduleID))
	b, err := s.c.do(ctx, "DELETE", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ReportScheduleDeleteResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ListReportSchedules calls GET /compliance/report-schedules/.
// List Report Schedules
//
// List the org's report schedules, newest-first. Soft-deleted excluded.
// `scheduling_enabled` reflects whether the scheduled-reports worker is
// actually running in this environment (`ENABLE_SCHEDULED_REPORTS_WORKER`).
// When false, stored schedules will not fire until an operator enables it —
// the UI surfaces this so it never shows a misleading "Active" status.
func (s *ReportSchedulesClient) ListReportSchedules(ctx context.Context, params *ReportSchedulesListReportSchedulesParams) (*ReportScheduleListResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/report-schedules/", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ReportScheduleListResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/report-schedules/", Err: err}
	}
	return &out, nil
}

// PauseReportSchedule calls POST /compliance/report-schedules/{schedule_id}/pause.
// Pause Report Schedule
func (s *ReportSchedulesClient) PauseReportSchedule(ctx context.Context, scheduleID string, params *ReportSchedulesPauseReportScheduleParams) (*ReportSchedulePauseResponse, error) {
	path := fmt.Sprintf("/compliance/report-schedules/%s/pause", url.PathEscape(scheduleID))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ReportSchedulePauseResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ResumeReportSchedule calls POST /compliance/report-schedules/{schedule_id}/resume.
// Resume Report Schedule
func (s *ReportSchedulesClient) ResumeReportSchedule(ctx context.Context, scheduleID string, params *ReportSchedulesResumeReportScheduleParams) (*ReportSchedulePauseResponse, error) {
	path := fmt.Sprintf("/compliance/report-schedules/%s/resume", url.PathEscape(scheduleID))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ReportSchedulePauseResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}
