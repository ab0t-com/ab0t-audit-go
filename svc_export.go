// ANNOTATION BLOCK
// File: svc_export.go — the `Export` designation of the Audit Service.
// Why it exists: typed methods for every Export endpoint, reached via
// c.Export.<Method>. Each is context-first, takes typed path/query/body
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

// ExportCancelExportJobParams holds the query parameters for Export.CancelExportJob.
type ExportCancelExportJobParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ExportCancelExportJobParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ExportDownloadExportJobParams holds the query parameters for Export.DownloadExportJob.
type ExportDownloadExportJobParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ExportDownloadExportJobParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ExportGetExportJobParams holds the query parameters for Export.GetExportJob.
type ExportGetExportJobParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ExportGetExportJobParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ExportGetExportJobsParams holds the query parameters for Export.GetExportJobs.
type ExportGetExportJobsParams struct {
	OrgID  *string `json:"org_id,omitempty"`  // Organization ID
	UserID *string `json:"user_id,omitempty"` // Show only exports by this user
	Status *string `json:"status,omitempty"`  // Filter by status (pending, processing, completed, failed)
	Limit  *int    `json:"limit,omitempty"`   // Max jobs to return
	Offset *int    `json:"offset,omitempty"`  // Skip first N jobs for pagination
}

func (p *ExportGetExportJobsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.UserID != nil {
		v.Set("user_id", *p.UserID)
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

// CancelExportJob calls DELETE /export/jobs/{export_id}.
// Cancel Export Job
//
// Cancel a pending or processing export job. Stop unnecessary exports to
// save resources. **When to use**: Created export by mistake, no longer
// need the data, or want to recreate with different filters. **Returns**:
// Success confirmation if job was canceled. **Restrictions**: Can only
// cancel jobs in 'pending' or 'processing' status. Completed exports remain
// available. **Next step**: Job will stop processing. Create new export if
// needed with correct parameters.
func (s *ExportClient) CancelExportJob(ctx context.Context, exportID string, params *ExportCancelExportJobParams) (*ExportCancelResponse, error) {
	path := fmt.Sprintf("/export/jobs/%s", url.PathEscape(exportID))
	b, err := s.c.do(ctx, "DELETE", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ExportCancelResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// CreateExportJob calls POST /export/.
// Create Export Job
//
// Create audit data export job. Generates downloadable file of filtered
// audit events. **When to use**: Need data for compliance audits, backup,
// external analysis, or reporting. **Formats**: csv (Excel/sheets), json
// (APIs/code), xml (legacy systems), parquet (big data analytics).
// **Returns**: export_id immediately. File generation happens in
// background. **Next step**: Poll GET /export/jobs/{export_id} until
// status='completed', then download from URL. **Workflow**: Create export →
// Poll status → Download file → File expires after 7 days.
func (s *ExportClient) CreateExportJob(ctx context.Context, body ExportRequestInput) (*ExportCreateResponse, error) {
	b, err := s.c.do(ctx, "POST", "/export/", nil, body)
	if err != nil {
		return nil, err
	}
	var out ExportCreateResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/export/", Err: err}
	}
	return &out, nil
}

// DownloadExportJob calls GET /export/jobs/{export_id}/download.
// Download Export Job
//
// Download the generated file for a completed export job. **When to use**:
// After GET /export/jobs/{id} shows status='completed', retrieve the actual
// file. **Returns**: The export artifact as a file download
// (CSV/JSON/XML/Parquet) with the right content type. **Errors**: 404 if
// the job or its file is gone/expired; 409 if the export isn't finished
// yet. **Access**: Same per-job org/owner authorization as viewing the job;
// serves from this service.
func (s *ExportClient) DownloadExportJob(ctx context.Context, exportID string, params *ExportDownloadExportJobParams) (RawMessage, error) {
	path := fmt.Sprintf("/export/jobs/%s/download", url.PathEscape(exportID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// GetExportFormats calls GET /export/formats.
// Get Export Formats
//
// Discover available export formats and choose the right one for your use
// case. **When to use**: Before creating export, decide which format
// matches your needs. **Returns**: List of formats with descriptions,
// capabilities, size limits, and typical use cases. **Formats**: csv
// (spreadsheets), json (APIs/programming), xml (legacy integrations),
// parquet (analytics/ML). **Next step**: Use chosen format in POST /export/
// request.
func (s *ExportClient) GetExportFormats(ctx context.Context) (*ExportFormatsResponse, error) {
	b, err := s.c.do(ctx, "GET", "/export/formats", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ExportFormatsResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/export/formats", Err: err}
	}
	return &out, nil
}

// GetExportJob calls GET /export/jobs/{export_id}.
// Get Export Job
//
// Get detailed status and download URL for a specific export job. Poll this
// to track progress. **When to use**: After creating export with POST
// /export/, check if file is ready. **Returns**: Full job details including
// status, progress, download_url (when completed), and error_message (if
// failed). **Polling**: Call every 5-10 seconds until status changes from
// 'pending'/'processing' to 'completed'/'failed'. **Next step**: When
// completed, use download_url to retrieve file. URL expires after 7 days.
func (s *ExportClient) GetExportJob(ctx context.Context, exportID string, params *ExportGetExportJobParams) (*ExportJobResponse, error) {
	path := fmt.Sprintf("/export/jobs/%s", url.PathEscape(exportID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ExportJobResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetExportJobs calls GET /export/jobs.
// Get Export Jobs
//
// List all export jobs for your organization. Track export history and
// status. **When to use**: See recent exports, find download URLs, check
// export progress. **Returns**: List of export jobs, newest first, with
// status and download URLs if completed. **Next step**: Use export_id to
// GET /export/jobs/{id} for details or download completed exports. **Tip**:
// Filter by status='completed' to find ready downloads, or 'failed' to
// troubleshoot.
func (s *ExportClient) GetExportJobs(ctx context.Context, params *ExportGetExportJobsParams) (*PaginatedResponseExportJobResponse, error) {
	b, err := s.c.do(ctx, "GET", "/export/jobs", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseExportJobResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/export/jobs", Err: err}
	}
	return &out, nil
}
