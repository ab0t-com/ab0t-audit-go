// ANNOTATION BLOCK
// File: types.go — every schema in the Audit Service OpenAPI as a Go type.
// Why it exists: request/response bodies for all 226 endpoints are typed here
// so callers never hand-assemble maps. Covers all 215 component schemas:
// object schemas as structs (required fields = value, optional = pointer or
// omitempty), the 6 string enums as typed constants, and the generic
// PaginatedResponse_* / ListEnvelope_* wrappers. Generated from openapi.json;
// hand-edit the client layer, not this file's shape.
package auditclient

import "encoding/json"

// RawMessage is a deferred-decode JSON value, used where the service returns an
// open/untyped payload (documented as `any`). Unmarshal it into a concrete type
// yourself, or inspect it as raw bytes.
type RawMessage = json.RawMessage

// AdHocQueryRequest maps the `AdHocQueryRequest` schema.
type AdHocQueryRequest struct {
	SQL            string  `json:"sql"` // required
	OrgID          *string `json:"org_id,omitempty"`
	TimeoutSeconds *int    `json:"timeout_seconds,omitempty"`
	MaxRows        *int    `json:"max_rows,omitempty"`
	Description    *string `json:"description,omitempty"`
}

// AdHocQueryResponse maps the `AdHocQueryResponse` schema.
type AdHocQueryResponse struct {
	QueryID              string           `json:"query_id"`               // required
	Status               string           `json:"status"`                 // required
	Rows                 []map[string]any `json:"rows"`                   // required
	RowsReturned         int              `json:"rows_returned"`          // required
	Columns              []string         `json:"columns"`                // required
	ExecutionTimeSeconds float64          `json:"execution_time_seconds"` // required
	Truncated            bool             `json:"truncated"`              // required
	Timestamp            string           `json:"timestamp"`              // required
}

// AssertionsResponse maps the `AssertionsResponse` schema.
type AssertionsResponse struct {
	Total   int              `json:"total"` // required
	Summary map[string]any   `json:"summary,omitempty"`
	Items   []map[string]any `json:"items,omitempty"`
}

// AuditEvent maps the `AuditEvent` schema.
// Core audit event model
type AuditEvent struct {
	EventID            string         `json:"event_id"`   // required
	Timestamp          string         `json:"timestamp"`  // required
	EventType          string         `json:"event_type"` // required
	Service            string         `json:"service"`    // required
	UserID             *string        `json:"user_id,omitempty"`
	OrgID              *string        `json:"org_id,omitempty"`
	ResourceType       *string        `json:"resource_type,omitempty"`
	ResourceID         *string        `json:"resource_id,omitempty"`
	Action             string         `json:"action"` // required
	Outcome            *string        `json:"outcome,omitempty"`
	SessionID          *string        `json:"session_id,omitempty"`
	IPAddress          *string        `json:"ip_address,omitempty"`
	UserAgent          *string        `json:"user_agent,omitempty"`
	RequestData        map[string]any `json:"request_data,omitempty"`
	ResponseData       map[string]any `json:"response_data,omitempty"`
	Metadata           map[string]any `json:"metadata,omitempty"`
	ComplianceControls []string       `json:"compliance_controls,omitempty"`
	IngestedAt         *string        `json:"ingested_at,omitempty"`
}

// AuditEventCreate maps the `AuditEventCreate` schema.
// Model for creating audit events
type AuditEventCreate struct {
	EventType    string         `json:"event_type"` // required
	Service      string         `json:"service"`    // required
	UserID       *string        `json:"user_id,omitempty"`
	OrgID        *string        `json:"org_id,omitempty"`
	ResourceType *string        `json:"resource_type,omitempty"`
	ResourceID   *string        `json:"resource_id,omitempty"`
	Action       string         `json:"action"` // required
	Outcome      *string        `json:"outcome,omitempty"`
	SessionID    *string        `json:"session_id,omitempty"`
	IPAddress    *string        `json:"ip_address,omitempty"`
	UserAgent    *string        `json:"user_agent,omitempty"`
	RequestData  map[string]any `json:"request_data,omitempty"`
	ResponseData map[string]any `json:"response_data,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

// AuditPeriod maps the `AuditPeriod` schema.
// Org-active audit period — typically a SOC2 Type II window or annual
// review.
type AuditPeriod struct {
	OrgID     *string `json:"org_id,omitempty"`
	Name      string  `json:"name"` // required
	Framework *string `json:"framework,omitempty"`
	StartDate string  `json:"start_date"` // required
	EndDate   string  `json:"end_date"`   // required
	Auditor   *string `json:"auditor,omitempty"`
	SetBy     *string `json:"set_by,omitempty"`
	SetAt     *string `json:"set_at,omitempty"`
}

// AuditPeriodWrite maps the `AuditPeriodWrite` schema.
// Mutable shape for PUT /settings/audit-period.
type AuditPeriodWrite struct {
	Name      string  `json:"name"` // required
	Framework *string `json:"framework,omitempty"`
	StartDate string  `json:"start_date"` // required
	EndDate   string  `json:"end_date"`   // required
	Auditor   *string `json:"auditor,omitempty"`
}

// AuditorRequest maps the `AuditorRequest` schema.
// External-auditor question / evidence ask tracked through fieldwork.
type AuditorRequest struct {
	RequestID         string               `json:"request_id"` // required
	OrgID             *string              `json:"org_id,omitempty"`
	Subject           string               `json:"subject"` // required
	Requestor         *string              `json:"requestor,omitempty"`
	Framework         *string              `json:"framework,omitempty"`
	ControlID         *string              `json:"control_id,omitempty"`
	State             *AuditorRequestState `json:"state,omitempty"`
	DueAt             *string              `json:"due_at,omitempty"`
	CreatedAt         *string              `json:"created_at,omitempty"`
	UpdatedAt         *string              `json:"updated_at,omitempty"`
	EvidencePackageID *string              `json:"evidence_package_id,omitempty"`
	Notes             *string              `json:"notes,omitempty"`
	CreatedBy         *string              `json:"created_by,omitempty"`
}

// AuditorRequestCreate maps the `AuditorRequestCreate` schema.
// Shape for POST /compliance/auditor-requests.
type AuditorRequestCreate struct {
	Subject   string  `json:"subject"` // required
	Requestor *string `json:"requestor,omitempty"`
	Framework *string `json:"framework,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
	DueAt     *string `json:"due_at,omitempty"`
	Notes     *string `json:"notes,omitempty"`
}

// AuditorRequestSatisfy maps the `AuditorRequestSatisfy` schema.
// Shape for POST /compliance/auditor-requests/{id}/satisfy. Links an
// evidence package to the request and transitions it to `satisfied`.
// `evidence_package_id` is required — satisfying an auditor ask means
// pointing at the evidence that answers it. `notes` lets the reviewer
// record a closing remark (e.g. "delivered to auditor 2026-06-15").
type AuditorRequestSatisfy struct {
	EvidencePackageID string  `json:"evidence_package_id"` // required
	Notes             *string `json:"notes,omitempty"`
}

// AuditorRequestState is a string enumeration (AuditorRequestState).
type AuditorRequestState string

// Enumerated values for AuditorRequestState.
const (
	AuditorRequestStateNew             AuditorRequestState = "new"
	AuditorRequestStateInProgress      AuditorRequestState = "in_progress"
	AuditorRequestStateAwaitingAuditor AuditorRequestState = "awaiting_auditor"
	AuditorRequestStateSatisfied       AuditorRequestState = "satisfied"
	AuditorRequestStateDeclined        AuditorRequestState = "declined"
)

// AuditorRequestUpdate maps the `AuditorRequestUpdate` schema.
// Shape for PATCH /compliance/auditor-requests/{id}. All fields optional.
type AuditorRequestUpdate struct {
	State             *AuditorRequestState `json:"state,omitempty"`
	Subject           *string              `json:"subject,omitempty"`
	Requestor         *string              `json:"requestor,omitempty"`
	Framework         *string              `json:"framework,omitempty"`
	ControlID         *string              `json:"control_id,omitempty"`
	DueAt             *string              `json:"due_at,omitempty"`
	EvidencePackageID *string              `json:"evidence_package_id,omitempty"`
	Notes             *string              `json:"notes,omitempty"`
}

// AuditorViewPayloadWithLineage maps the `AuditorViewPayloadWithLineage` schema.
// Auditor payload + a prominent 'this evidence has been superseded' banner.
// An auditor silently reading stale evidence is the failure mode this
// closes: when “superseded_by“ is set the auditor view must warn that a
// corrected package exists (with the reason + date).
type AuditorViewPayloadWithLineage struct {
	PackageName       string           `json:"package_name"` // required
	Framework         string           `json:"framework"`    // required
	ControlIDs        []string         `json:"control_ids"`  // required
	TimeRangeStart    *string          `json:"time_range_start,omitempty"`
	TimeRangeEnd      *string          `json:"time_range_end,omitempty"`
	EventCount        *int             `json:"event_count,omitempty"`
	SealedAt          *string          `json:"sealed_at,omitempty"`
	AuditorEmail      string           `json:"auditor_email"`  // required
	ExpiresAt         string           `json:"expires_at"`     // required
	WatermarkText     string           `json:"watermark_text"` // required
	Events            []map[string]any `json:"events,omitempty"`
	PackageDigest     *string          `json:"package_digest,omitempty"`
	DigestAlgorithm   *string          `json:"digest_algorithm,omitempty"`
	OutcomesBreakdown map[string]int   `json:"outcomes_breakdown,omitempty"`
	CustodySealedBy   *string          `json:"custody_sealed_by,omitempty"`
	CustodyCreatedAt  *string          `json:"custody_created_at,omitempty"`
	CustodySharedBy   *string          `json:"custody_shared_by,omitempty"`
	SupersededBy      *string          `json:"superseded_by,omitempty"`
	SupersededReason  *string          `json:"superseded_reason,omitempty"`
	SupersededAt      *string          `json:"superseded_at,omitempty"`
}

// CallbackDeliveryItem maps the `CallbackDeliveryItem` schema.
type CallbackDeliveryItem struct {
	DeliveryID       string  `json:"delivery_id"` // required
	JobID            *string `json:"job_id,omitempty"`
	LifecycleEventID *string `json:"lifecycle_event_id,omitempty"`
	SubscriptionID   string  `json:"subscription_id"` // required
	TenantID         string  `json:"tenant_id"`       // required
	OrgID            string  `json:"org_id"`          // required
	EventTopic       string  `json:"event_topic"`     // required
	CallbackURL      string  `json:"callback_url"`    // required
	Status           string  `json:"status"`          // required
	HTTPStatus       *int    `json:"http_status,omitempty"`
	LatencyMs        *int    `json:"latency_ms,omitempty"`
	AttemptCount     *int    `json:"attempt_count,omitempty"`
	IdempotencyKey   *string `json:"idempotency_key,omitempty"`
	PayloadHash      *string `json:"payload_hash,omitempty"`
	ErrorMessage     *string `json:"error_message,omitempty"`
	AttemptedAt      *string `json:"attempted_at,omitempty"`
}

// CallbackSubscriptionCreateRequest maps the `CallbackSubscriptionCreateRequest` schema.
type CallbackSubscriptionCreateRequest struct {
	OrgID         *string `json:"org_id,omitempty"`
	EventTopic    string  `json:"event_topic"`  // required
	CallbackURL   string  `json:"callback_url"` // required
	SigningKeyRef *string `json:"signing_key_ref,omitempty"`
	Enabled       *bool   `json:"enabled,omitempty"`
}

// CallbackSubscriptionPauseResponse maps the `CallbackSubscriptionPauseResponse` schema.
type CallbackSubscriptionPauseResponse struct {
	Status         *string `json:"status,omitempty"`
	SubscriptionID string  `json:"subscription_id"` // required
}

// CallbackSubscriptionResponse maps the `CallbackSubscriptionResponse` schema.
type CallbackSubscriptionResponse struct {
	SubscriptionID string `json:"subscription_id"` // required
	TenantID       string `json:"tenant_id"`       // required
	OrgID          string `json:"org_id"`          // required
	EventTopic     string `json:"event_topic"`     // required
	CallbackURL    string `json:"callback_url"`    // required
	SigningKeyRef  string `json:"signing_key_ref"` // required
	Enabled        bool   `json:"enabled"`         // required
}

// CallbackSubscriptionResumeResponse maps the `CallbackSubscriptionResumeResponse` schema.
type CallbackSubscriptionResumeResponse struct {
	Status         *string `json:"status,omitempty"`
	SubscriptionID string  `json:"subscription_id"` // required
}

// CallbackSubscriptionTestResponse maps the `CallbackSubscriptionTestResponse` schema.
type CallbackSubscriptionTestResponse struct {
	Status           *string        `json:"status,omitempty"`
	SubscriptionID   string         `json:"subscription_id"`    // required
	LifecycleEventID string         `json:"lifecycle_event_id"` // required
	Dispatch         map[string]any `json:"dispatch,omitempty"`
}

// CallbackTestRequest maps the `CallbackTestRequest` schema.
type CallbackTestRequest struct {
	OrgID   *string        `json:"org_id,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
}

// CapabilitiesResponse maps the `CapabilitiesResponse` schema.
// Allowlisted feature-flag snapshot. Add keys deliberately; never config
// values.
type CapabilitiesResponse struct {
	LlmJudge         bool     `json:"llm_judge"`         // required
	CopilotMode      string   `json:"copilot_mode"`      // required
	Callbacks        bool     `json:"callbacks"`         // required
	MLScoring        bool     `json:"ml_scoring"`        // required
	ScheduledReports bool     `json:"scheduled_reports"` // required
	EmailDelivery    bool     `json:"email_delivery"`    // required
	DemoScenarios    bool     `json:"demo_scenarios"`    // required
	TemplateNames    []string `json:"template_names"`    // required
}

// CaseApprovalDecisionRequest maps the `CaseApprovalDecisionRequest` schema.
type CaseApprovalDecisionRequest struct {
	OrgID    *string        `json:"org_id,omitempty"`
	Decision string         `json:"decision"` // required
	Reason   string         `json:"reason"`   // required
	Metadata map[string]any `json:"metadata,omitempty"`
}

// CaseApprovalDecisionResponse maps the `CaseApprovalDecisionResponse` schema.
type CaseApprovalDecisionResponse struct {
	Status     *string `json:"status,omitempty"`
	CaseID     string  `json:"case_id"`     // required
	Decision   string  `json:"decision"`    // required
	ApprovalID string  `json:"approval_id"` // required
	CreatedAt  string  `json:"created_at"`  // required
}

// CaseApprovalItem maps the `CaseApprovalItem` schema.
type CaseApprovalItem struct {
	ApprovalID     string         `json:"approval_id"`      // required
	CaseID         string         `json:"case_id"`          // required
	OrgID          string         `json:"org_id"`           // required
	ApproverUserID string         `json:"approver_user_id"` // required
	Decision       string         `json:"decision"`         // required
	Reason         *string        `json:"reason,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	CreatedAt      *string        `json:"created_at,omitempty"`
}

// CaseEscalationItem maps the `CaseEscalationItem` schema.
type CaseEscalationItem struct {
	EscalationID string         `json:"escalation_id"` // required
	CaseID       string         `json:"case_id"`       // required
	TenantID     string         `json:"tenant_id"`     // required
	OrgID        string         `json:"org_id"`        // required
	RiskLevel    string         `json:"risk_level"`    // required
	Channel      string         `json:"channel"`       // required
	Target       *string        `json:"target,omitempty"`
	Status       string         `json:"status"` // required
	Payload      map[string]any `json:"payload,omitempty"`
	CreatedAt    *string        `json:"created_at,omitempty"`
}

// CaseTimelineEntry maps the `CaseTimelineEntry` schema.
type CaseTimelineEntry struct {
	EventID   string         `json:"event_id"` // required
	CaseID    string         `json:"case_id"`  // required
	TraceID   *string        `json:"trace_id,omitempty"`
	RunID     *string        `json:"run_id,omitempty"`
	TenantID  string         `json:"tenant_id"`  // required
	OrgID     string         `json:"org_id"`     // required
	EventType string         `json:"event_type"` // required
	Status    string         `json:"status"`     // required
	RiskLevel string         `json:"risk_level"` // required
	ActorID   *string        `json:"actor_id,omitempty"`
	Details   map[string]any `json:"details,omitempty"`
	CreatedAt *string        `json:"created_at,omitempty"`
}

// ClassificationRule maps the `ClassificationRule` schema.
type ClassificationRule struct {
	FieldPattern   string `json:"field_pattern"`  // required
	Classification string `json:"classification"` // required
}

// ClassificationSettings maps the `ClassificationSettings` schema.
type ClassificationSettings struct {
	Rules     []ClassificationRule `json:"rules"` // required
	IsDefault *bool                `json:"is_default,omitempty"`
	UpdatedBy *string              `json:"updated_by,omitempty"`
	UpdatedAt *string              `json:"updated_at,omitempty"`
}

// ComplianceAnalyticsReport maps the `ComplianceAnalyticsReport` schema.
// Analytics summary for compliance metrics - simplified for dashboard use
type ComplianceAnalyticsReport struct {
	OrgID               string             `json:"org_id"`               // required
	PeriodStart         string             `json:"period_start"`         // required
	PeriodEnd           string             `json:"period_end"`           // required
	TotalEvents         int                `json:"total_events"`         // required
	DataSubjectsCount   int                `json:"data_subjects_count"`  // required
	RetentionCompliance map[string]any     `json:"retention_compliance"` // required
	AccessPatterns      map[string]int     `json:"access_patterns"`      // required
	ExportActivities    []map[string]any   `json:"export_activities"`    // required
	FrameworkScores     map[string]float64 `json:"framework_scores,omitempty"`
}

// ComplianceControlSummary maps the `ComplianceControlSummary` schema.
type ComplianceControlSummary struct {
	ControlID          string   `json:"control_id"`           // required
	Framework          string   `json:"framework"`            // required
	Status             string   `json:"status"`               // required
	EventCount         int      `json:"event_count"`          // required
	FailedCount        int      `json:"failed_count"`         // required
	PassRate           float64  `json:"pass_rate"`            // required
	SampleEventIDs     []string `json:"sample_event_ids"`     // required
	RelevantEventTypes []string `json:"relevant_event_types"` // required
	LastChecked        *string  `json:"last_checked,omitempty"`
	Reason             string   `json:"reason"`       // required
	AnomalyHits        int      `json:"anomaly_hits"` // required
	MaxZscore          float64  `json:"max_zscore"`   // required
}

// ComplianceControlsListResponse maps the `ComplianceControlsListResponse` schema.
type ComplianceControlsListResponse struct {
	OrgID         string                     `json:"org_id"`         // required
	TotalControls int                        `json:"total_controls"` // required
	Controls      []ComplianceControlSummary `json:"controls"`       // required
}

// ComplianceDashboardResponse maps the `ComplianceDashboardResponse` schema.
// Real-time compliance posture dashboard.
type ComplianceDashboardResponse struct {
	OrgID                    string             `json:"org_id"`              // required
	GeneratedAt              string             `json:"generated_at"`        // required
	ComplianceScore          int                `json:"compliance_score"`    // required
	AuditReadiness           string             `json:"audit_readiness"`     // required
	TotalControls            int                `json:"total_controls"`      // required
	ControlsPassing          int                `json:"controls_passing"`    // required
	ControlsFailing          int                `json:"controls_failing"`    // required
	ControlsWarning          int                `json:"controls_warning"`    // required
	ControlsNotTested        int                `json:"controls_not_tested"` // required
	LastTestDate             *string            `json:"last_test_date,omitempty"`
	NextTestDate             *string            `json:"next_test_date,omitempty"`
	FrameworkScores          map[string]float64 `json:"framework_scores,omitempty"`
	OpenFindings             []map[string]any   `json:"open_findings,omitempty"`
	EventsLast30Days         *int               `json:"events_last_30_days,omitempty"`
	FailedEventsLast30Days   *int               `json:"failed_events_last_30_days,omitempty"`
	AnomalySignalsLast30Days *int               `json:"anomaly_signals_last_30_days,omitempty"`
	Insights                 []string           `json:"insights,omitempty"`
	DriftAlerts              []map[string]any   `json:"drift_alerts,omitempty"`
}

// CompliancePolicy maps the `CompliancePolicy` schema.
type CompliancePolicy struct {
	PolicyID                  string                        `json:"policy_id"` // required
	OrgID                     string                        `json:"org_id"`    // required
	Name                      string                        `json:"name"`      // required
	Description               *string                       `json:"description,omitempty"`
	Standards                 []ComplianceStandard          `json:"standards"`  // required
	CreatedAt                 string                        `json:"created_at"` // required
	UpdatedAt                 string                        `json:"updated_at"` // required
	CreatedBy                 string                        `json:"created_by"` // required
	IsActive                  *bool                         `json:"is_active,omitempty"`
	DefaultRetention          RetentionPeriod               `json:"default_retention"` // required
	EventTypeRetention        map[string]RetentionPeriod    `json:"event_type_retention,omitempty"`
	GDPREnabled               *bool                         `json:"gdpr_enabled,omitempty"`
	GDPRLegalBasis            *GDPRLegalBasis               `json:"gdpr_legal_basis,omitempty"`
	GDPRDataController        *string                       `json:"gdpr_data_controller,omitempty"`
	GDPRDataProcessor         *string                       `json:"gdpr_data_processor,omitempty"`
	GDPRPrivacyPolicyURL      *string                       `json:"gdpr_privacy_policy_url,omitempty"`
	SOC2Enabled               *bool                         `json:"soc2_enabled,omitempty"`
	SOC2ServiceOrganization   *string                       `json:"soc2_service_organization,omitempty"`
	SOC2ReportPeriodMonths    *int                          `json:"soc2_report_period_months,omitempty"`
	SOC2ControlActivities     []string                      `json:"soc2_control_activities,omitempty"`
	AutoClassifyData          *bool                         `json:"auto_classify_data,omitempty"`
	ClassificationRules       map[string]DataClassification `json:"classification_rules,omitempty"`
	MaxExportSizeGb           *float64                      `json:"max_export_size_gb,omitempty"`
	AllowedExportFormats      []string                      `json:"allowed_export_formats,omitempty"`
	RequireApprovalForExports *bool                         `json:"require_approval_for_exports,omitempty"`
	ExportApprovalRoles       []string                      `json:"export_approval_roles,omitempty"`
	RequireMfaForAccess       *bool                         `json:"require_mfa_for_access,omitempty"`
	IPWhitelist               []string                      `json:"ip_whitelist,omitempty"`
	AuditAccessLogging        *bool                         `json:"audit_access_logging,omitempty"`
	TamperDetection           *bool                         `json:"tamper_detection,omitempty"`
}

// ComplianceReportAck maps the `ComplianceReportAck` schema.
type ComplianceReportAck struct {
	ReportID            string `json:"report_id"`            // required
	Status              string `json:"status"`               // required
	EstimatedCompletion string `json:"estimated_completion"` // required
	Message             string `json:"message"`              // required
	Demo                *bool  `json:"demo,omitempty"`
}

// ComplianceReportRow maps the `ComplianceReportRow` schema.
// One previously-generated compliance report (CS-015). Backs the "Recently
// generated" table on /ui/app/reports. Sourced from the real
// `compliance_reports` rows — NOT approximated from audit events — so
// status/standard/period/download reflect the actual report record.
type ComplianceReportRow struct {
	ReportID      string  `json:"report_id"` // required
	OrgID         string  `json:"org_id"`    // required
	Standard      string  `json:"standard"`  // required
	ReportType    *string `json:"report_type,omitempty"`
	Status        string  `json:"status"` // required
	PeriodStart   *string `json:"period_start,omitempty"`
	PeriodEnd     *string `json:"period_end,omitempty"`
	RequestedBy   *string `json:"requested_by,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
	CompletedAt   *string `json:"completed_at,omitempty"`
	ReportFileURL *string `json:"report_file_url,omitempty"`
	ErrorMessage  *string `json:"error_message,omitempty"`
}

// ComplianceRequestStatus maps the `ComplianceRequestStatus` schema.
type ComplianceRequestStatus struct {
	RequestID          string  `json:"request_id"`      // required
	Status             string  `json:"status"`          // required
	RequestType        string  `json:"request_type"`    // required
	DataSubjectID      string  `json:"data_subject_id"` // required
	CompletionDeadline *string `json:"completion_deadline,omitempty"`
	CreatedAt          *string `json:"created_at,omitempty"`
	ProcessingNotes    *string `json:"processing_notes,omitempty"`
	RecordsAffected    *int    `json:"records_affected,omitempty"`
}

// ComplianceStandard is a string enumeration (ComplianceStandard).
type ComplianceStandard string

// Enumerated values for ComplianceStandard.
const (
	ComplianceStandardSOC2   ComplianceStandard = "soc2"
	ComplianceStandardGDPR   ComplianceStandard = "gdpr"
	ComplianceStandardHIPAA  ComplianceStandard = "hipaa"
	ComplianceStandardPciDss ComplianceStandard = "pci_dss"
	ComplianceStandardCcpa   ComplianceStandard = "ccpa"
	ComplianceStandardPipeda ComplianceStandard = "pipeda"
	ComplianceStandardCustom ComplianceStandard = "custom"
)

// ComplianceTemplateCloneResponse maps the `ComplianceTemplateCloneResponse` schema.
type ComplianceTemplateCloneResponse struct {
	TemplateID          string            `json:"template_id"`       // required
	Framework           string            `json:"framework"`         // required
	Title               string            `json:"title"`             // required
	RenderedMarkdown    string            `json:"rendered_markdown"` // required
	OrgID               string            `json:"org_id"`            // required
	RenderedAt          string            `json:"rendered_at"`       // required
	PlaceholdersUsed    map[string]string `json:"placeholders_used,omitempty"`
	PlaceholdersMissing []string          `json:"placeholders_missing,omitempty"`
}

// ComplianceTemplateItem maps the `ComplianceTemplateItem` schema.
type ComplianceTemplateItem struct {
	TemplateID   string   `json:"template_id"` // required
	Framework    string   `json:"framework"`   // required
	Title        string   `json:"title"`       // required
	Description  string   `json:"description"` // required
	Placeholders []string `json:"placeholders,omitempty"`
	DocumentType string   `json:"document_type"` // required
	Citation     *string  `json:"citation,omitempty"`
}

// ControlMappingCreate maps the `ControlMappingCreate` schema.
type ControlMappingCreate struct {
	EventType string   `json:"event_type"` // required
	Controls  []string `json:"controls"`   // required
}

// ControlMappingRow maps the `ControlMappingRow` schema.
type ControlMappingRow struct {
	ID        string   `json:"id"`         // required
	OrgID     string   `json:"org_id"`     // required
	EventType string   `json:"event_type"` // required
	Controls  []string `json:"controls"`   // required
	CreatedBy *string  `json:"created_by,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	UpdatedAt *string  `json:"updated_at,omitempty"`
}

// CopilotAnswer maps the `CopilotAnswer` schema.
type CopilotAnswer struct {
	ConversationID string              `json:"conversation_id"` // required
	Mode           string              `json:"mode"`            // required
	TemplateName   *string             `json:"template_name,omitempty"`
	Answer         string              `json:"answer"`        // required
	SQL            string              `json:"sql"`           // required
	Columns        []string            `json:"columns"`       // required
	Rows           [][]json.RawMessage `json:"rows"`          // required
	RowsReturned   int                 `json:"rows_returned"` // required
	Safety         map[string]any      `json:"safety"`        // required
	UsedFallback   *bool               `json:"used_fallback,omitempty"`
	FallbackReason *string             `json:"fallback_reason,omitempty"`
}

// CopilotAsk maps the `CopilotAsk` schema.
type CopilotAsk struct {
	Question       string  `json:"question"` // required
	ConversationID *string `json:"conversation_id,omitempty"`
	OrgID          *string `json:"org_id,omitempty"`
}

// DataClassification is a string enumeration (DataClassification).
type DataClassification string

// Enumerated values for DataClassification.
const (
	DataClassificationPublic       DataClassification = "public"
	DataClassificationInternal     DataClassification = "internal"
	DataClassificationConfidential DataClassification = "confidential"
	DataClassificationRestricted   DataClassification = "restricted"
	DataClassificationPii          DataClassification = "pii"
	DataClassificationPhi          DataClassification = "phi"
	DataClassificationPci          DataClassification = "pci"
)

// DataDeletionAck maps the `DataDeletionAck` schema.
type DataDeletionAck struct {
	RequestID          string `json:"request_id"`          // required
	Status             string `json:"status"`              // required
	CompletionDeadline string `json:"completion_deadline"` // required
	Message            string `json:"message"`             // required
	Warning            string `json:"warning"`             // required
}

// DataPurgeRequest maps the `DataPurgeRequest` schema.
type DataPurgeRequest struct {
	OrgID      *string `json:"org_id,omitempty"`
	BeforeDate string  `json:"before_date"` // required
	DryRun     *bool   `json:"dry_run,omitempty"`
	Reason     string  `json:"reason"` // required
}

// DataSubjectAccessAck maps the `DataSubjectAccessAck` schema.
type DataSubjectAccessAck struct {
	RequestID          string `json:"request_id"`          // required
	Status             string `json:"status"`              // required
	CompletionDeadline string `json:"completion_deadline"` // required
	Message            string `json:"message"`             // required
}

// DataSubjectAccessRequest maps the `DataSubjectAccessRequest` schema.
type DataSubjectAccessRequest struct {
	DataSubjectID  string   `json:"data_subject_id"` // required
	Scope          []string `json:"scope,omitempty"`
	DateRangeStart *string  `json:"date_range_start,omitempty"`
	DateRangeEnd   *string  `json:"date_range_end,omitempty"`
	Format         *string  `json:"format,omitempty"`
	Reason         *string  `json:"reason,omitempty"`
}

// DataSubjectDeletionRequest maps the `DataSubjectDeletionRequest` schema.
type DataSubjectDeletionRequest struct {
	DataSubjectID     string   `json:"data_subject_id"` // required
	Scope             []string `json:"scope,omitempty"`
	DateRangeStart    *string  `json:"date_range_start,omitempty"`
	DateRangeEnd      *string  `json:"date_range_end,omitempty"`
	Reason            string   `json:"reason"`      // required
	LegalBasis        string   `json:"legal_basis"` // required
	RetentionOverride *bool    `json:"retention_override,omitempty"`
}

// DatabaseMetricsResponse maps the `DatabaseMetricsResponse` schema.
type DatabaseMetricsResponse struct {
	TableMetrics       []DatabaseTableMetric       `json:"table_metrics"`       // required
	PerformanceMetrics []DatabasePerformanceMetric `json:"performance_metrics"` // required
	Timestamp          string                      `json:"timestamp"`           // required
}

// DatabasePerformanceMetric maps the `DatabasePerformanceMetric` schema.
type DatabasePerformanceMetric struct {
	QueryKind     string  `json:"query_kind"`      // required
	QueryCount    int     `json:"query_count"`     // required
	AvgDurationMs float64 `json:"avg_duration_ms"` // required
	MaxDurationMs float64 `json:"max_duration_ms"` // required
	AvgRowsRead   float64 `json:"avg_rows_read"`   // required
}

// DatabaseTableMetric maps the `DatabaseTableMetric` schema.
type DatabaseTableMetric struct {
	Table            string `json:"table"`             // required
	SizeOnDisk       string `json:"size_on_disk"`      // required
	CompressedSize   string `json:"compressed_size"`   // required
	UncompressedSize string `json:"uncompressed_size"` // required
	TotalRows        int    `json:"total_rows"`        // required
	PartsCount       int    `json:"parts_count"`       // required
}

// DeletionDecisionRequest maps the `DeletionDecisionRequest` schema.
type DeletionDecisionRequest struct {
	Reason *string `json:"reason,omitempty"`
}

// DeletionRestoreRequest maps the `DeletionRestoreRequest` schema.
type DeletionRestoreRequest struct {
	Reason *string `json:"reason,omitempty"`
}

// DetailedHealthResponse maps the `DetailedHealthResponse` schema.
type DetailedHealthResponse struct {
	Status    string                      `json:"status"`    // required
	Timestamp string                      `json:"timestamp"` // required
	Service   string                      `json:"service"`   // required
	Checks    map[string]HealthCheckEntry `json:"checks"`    // required
}

// DsarCommentBody maps the `DsarCommentBody` schema.
type DsarCommentBody struct {
	Comment string `json:"comment"` // required
}

// DsarCommentResponse maps the `DsarCommentResponse` schema.
type DsarCommentResponse struct {
	RequestID string `json:"request_id"` // required
	NoteAdded bool   `json:"note_added"` // required
	UpdatedAt string `json:"updated_at"` // required
}

// DsarRequestRow maps the `DsarRequestRow` schema.
type DsarRequestRow struct {
	RequestID          string               `json:"request_id"`      // required
	OrgID              string               `json:"org_id"`          // required
	RequestType        string               `json:"request_type"`    // required
	DataSubjectID      string               `json:"data_subject_id"` // required
	RequestedBy        string               `json:"requested_by"`    // required
	RequestDate        *string              `json:"request_date,omitempty"`
	CompletionDeadline *string              `json:"completion_deadline,omitempty"`
	Status             string               `json:"status"` // required
	Scope              []string             `json:"scope,omitempty"`
	Reason             *string              `json:"reason,omitempty"`
	AssignedTo         *string              `json:"assigned_to,omitempty"`
	ProcessingNotes    *string              `json:"processing_notes,omitempty"`
	CreatedAt          *string              `json:"created_at,omitempty"`
	UpdatedAt          *string              `json:"updated_at,omitempty"`
	CompletedAt        *string              `json:"completed_at,omitempty"`
	RecordsAffected    *int                 `json:"records_affected,omitempty"`
	StatusHistory      []StatusHistoryEntry `json:"status_history,omitempty"`
}

// DsarStatusPatchBody maps the `DsarStatusPatchBody` schema.
type DsarStatusPatchBody struct {
	Status  string  `json:"status"` // required
	Comment *string `json:"comment,omitempty"`
}

// DsarStatusPatchResponse maps the `DsarStatusPatchResponse` schema.
type DsarStatusPatchResponse struct {
	RequestID      string `json:"request_id"`      // required
	Status         string `json:"status"`          // required
	UpdatedAt      string `json:"updated_at"`      // required
	HistoryEntries int    `json:"history_entries"` // Number of entries in status_history after this update
}

// EvaluateEventResponse maps the `EvaluateEventResponse` schema.
type EvaluateEventResponse struct {
	Status  *string        `json:"status,omitempty"`
	Summary map[string]any `json:"summary,omitempty"`
}

// EvaluationEventRequest maps the `EvaluationEventRequest` schema.
type EvaluationEventRequest struct {
	OrgID     *string        `json:"org_id,omitempty"`
	EventType string         `json:"event_type"` // required
	Payload   map[string]any `json:"payload,omitempty"`
}

// EventStatsResponse maps the `EventStatsResponse` schema.
type EventStatsResponse struct {
	Total int     `json:"total"` // required
	OrgID *string `json:"org_id,omitempty"`
}

// EventVolumeStats maps the `EventVolumeStats` schema.
type EventVolumeStats struct {
	TotalEvents     int               `json:"total_events"`      // required
	EventsPerHour   []TimeSeriesPoint `json:"events_per_hour"`   // required
	EventsByType    map[string]int    `json:"events_by_type"`    // required
	EventsByOutcome map[string]int    `json:"events_by_outcome"` // required
}

// EvidencePackage maps the `EvidencePackage` schema.
// A sealed bundle of audit events satisfying one or more framework
// controls.
type EvidencePackage struct {
	PackageID      string   `json:"package_id"` // required
	OrgID          *string  `json:"org_id,omitempty"`
	Name           string   `json:"name"`      // required
	Framework      string   `json:"framework"` // required
	ControlIDs     []string `json:"control_ids,omitempty"`
	TimeRangeStart *string  `json:"time_range_start,omitempty"`
	TimeRangeEnd   *string  `json:"time_range_end,omitempty"`
	EventIDs       []string `json:"event_ids,omitempty"`
	Sealed         *bool    `json:"sealed,omitempty"`
	SealedAt       *string  `json:"sealed_at,omitempty"`
	CreatedBy      *string  `json:"created_by,omitempty"`
	CreatedAt      *string  `json:"created_at,omitempty"`
	UpdatedAt      *string  `json:"updated_at,omitempty"`
}

// EvidencePackageCreate maps the `EvidencePackageCreate` schema.
// Shape for POST /compliance/evidence-packages.
type EvidencePackageCreate struct {
	Name           string   `json:"name"`      // required
	Framework      string   `json:"framework"` // required
	ControlIDs     []string `json:"control_ids,omitempty"`
	TimeRangeStart *string  `json:"time_range_start,omitempty"`
	TimeRangeEnd   *string  `json:"time_range_end,omitempty"`
	EventIDs       []string `json:"event_ids,omitempty"`
	SealOnCreate   *bool    `json:"seal_on_create,omitempty"`
}

// EvidencePackageWithLineage maps the `EvidencePackageWithLineage` schema.
// An evidence package plus its supersede lineage (both directions). *
// “superseded_by“ present → this package is STALE; read the successor. *
// “supersedes_package_id“ present → this package IS a correction of
// another.
type EvidencePackageWithLineage struct {
	PackageID           string   `json:"package_id"` // required
	OrgID               *string  `json:"org_id,omitempty"`
	Name                string   `json:"name"`      // required
	Framework           string   `json:"framework"` // required
	ControlIDs          []string `json:"control_ids,omitempty"`
	TimeRangeStart      *string  `json:"time_range_start,omitempty"`
	TimeRangeEnd        *string  `json:"time_range_end,omitempty"`
	EventIDs            []string `json:"event_ids,omitempty"`
	Sealed              *bool    `json:"sealed,omitempty"`
	SealedAt            *string  `json:"sealed_at,omitempty"`
	CreatedBy           *string  `json:"created_by,omitempty"`
	CreatedAt           *string  `json:"created_at,omitempty"`
	UpdatedAt           *string  `json:"updated_at,omitempty"`
	SupersededBy        *string  `json:"superseded_by,omitempty"`
	SupersededReason    *string  `json:"superseded_reason,omitempty"`
	SupersededAt        *string  `json:"superseded_at,omitempty"`
	SupersedesPackageID *string  `json:"supersedes_package_id,omitempty"`
	SupersedeReason     *string  `json:"supersede_reason,omitempty"`
}

// ExplorerHistoryCreate maps the `ExplorerHistoryCreate` schema.
type ExplorerHistoryCreate struct {
	Question    string `json:"question"` // required
	ResultCount *int   `json:"result_count,omitempty"`
}

// ExplorerHistoryItem maps the `ExplorerHistoryItem` schema.
type ExplorerHistoryItem struct {
	Question    string  `json:"question"` // required
	ResultCount *int    `json:"result_count,omitempty"`
	RanAt       *string `json:"ran_at,omitempty"`
}

// ExplorerSavedQueryCreate maps the `ExplorerSavedQueryCreate` schema.
type ExplorerSavedQueryCreate struct {
	Question string  `json:"question"` // required
	Name     *string `json:"name,omitempty"`
}

// ExplorerSavedQueryDeleteResponse maps the `ExplorerSavedQueryDeleteResponse` schema.
type ExplorerSavedQueryDeleteResponse struct {
	QueryID string `json:"query_id"` // required
	Deleted bool   `json:"deleted"`  // required
}

// ExplorerSavedQueryItem maps the `ExplorerSavedQueryItem` schema.
type ExplorerSavedQueryItem struct {
	QueryID   string  `json:"query_id"` // required
	UserID    string  `json:"user_id"`  // required
	OrgID     string  `json:"org_id"`   // required
	Name      string  `json:"name"`     // required
	Question  string  `json:"question"` // required
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ExportCancelResponse maps the `ExportCancelResponse` schema.
type ExportCancelResponse struct {
	Message string `json:"message"` // required
}

// ExportCreateResponse maps the `ExportCreateResponse` schema.
type ExportCreateResponse struct {
	ExportID string `json:"export_id"` // required
	Status   string `json:"status"`    // required
	Message  string `json:"message"`   // required
}

// ExportFormatDescriptor maps the `ExportFormatDescriptor` schema.
type ExportFormatDescriptor struct {
	Name                string `json:"name"`                 // required
	Description         string `json:"description"`          // required
	SupportsCompression bool   `json:"supports_compression"` // required
	MaxSizeGb           int    `json:"max_size_gb"`          // required
	TypicalUseCase      string `json:"typical_use_case"`     // required
}

// ExportFormatsResponse maps the `ExportFormatsResponse` schema.
type ExportFormatsResponse struct {
	Formats map[string]ExportFormatDescriptor `json:"formats"` // required
}

// ExportJobResponse maps the `ExportJobResponse` schema.
// API response model for export jobs - for GET endpoints
type ExportJobResponse struct {
	ExportID            string  `json:"export_id"` // required
	OrgID               string  `json:"org_id"`    // required
	UserID              string  `json:"user_id"`   // required
	Name                string  `json:"name"`      // required
	Status              string  `json:"status"`    // required
	Format              string  `json:"format"`    // required
	Filters             string  `json:"filters"`   // required
	Fields              *string `json:"fields,omitempty"`
	IncludeRequestData  *int    `json:"include_request_data,omitempty"`
	IncludeResponseData *int    `json:"include_response_data,omitempty"`
	DeliveryMethod      *string `json:"delivery_method,omitempty"`
	CreatedAt           string  `json:"created_at"` // required
	StartedAt           *string `json:"started_at,omitempty"`
	CompletedAt         *string `json:"completed_at,omitempty"`
	RecordsExported     *int    `json:"records_exported,omitempty"`
	FileSizeBytes       *int    `json:"file_size_bytes,omitempty"`
	DownloadURL         *string `json:"download_url,omitempty"`
	ExpiresAt           *string `json:"expires_at,omitempty"`
	ErrorMessage        *string `json:"error_message,omitempty"`
}

// ExportRequestInput maps the `ExportRequestInput` schema.
// API input model for export requests - simplified for endpoint use
type ExportRequestInput struct {
	OrgID            *string        `json:"org_id,omitempty"`
	Name             string         `json:"name"`   // required
	Format           string         `json:"format"` // required
	StartTime        *string        `json:"start_time,omitempty"`
	EndTime          *string        `json:"end_time,omitempty"`
	EventTypes       []string       `json:"event_types,omitempty"`
	Actions          []string       `json:"actions,omitempty"`
	Outcomes         []string       `json:"outcomes,omitempty"`
	UserIDs          []string       `json:"user_ids,omitempty"`
	Filters          map[string]any `json:"filters,omitempty"`
	DeliveryMethod   *string        `json:"delivery_method,omitempty"`
	DeliveryTarget   *string        `json:"delivery_target,omitempty"`
	RequestedByEmail *string        `json:"requested_by_email,omitempty"`
}

// FacetCount maps the `FacetCount` schema.
type FacetCount struct {
	Value string `json:"value"` // required
	Count int    `json:"count"` // required
}

// FalsePositiveFeedbackItem maps the `FalsePositiveFeedbackItem` schema.
type FalsePositiveFeedbackItem struct {
	FeedbackID string  `json:"feedback_id"` // required
	TenantID   string  `json:"tenant_id"`   // required
	OrgID      string  `json:"org_id"`      // required
	Framework  string  `json:"framework"`   // required
	ControlID  string  `json:"control_id"`  // required
	Source     string  `json:"source"`      // required
	Reason     string  `json:"reason"`      // required
	RunID      *string `json:"run_id,omitempty"`
	TraceID    *string `json:"trace_id,omitempty"`
	ReportedBy *string `json:"reported_by,omitempty"`
	CreatedAt  *string `json:"created_at,omitempty"`
}

// FalsePositiveFeedbackRequest maps the `FalsePositiveFeedbackRequest` schema.
type FalsePositiveFeedbackRequest struct {
	OrgID     *string `json:"org_id,omitempty"`
	Framework string  `json:"framework"`  // required
	ControlID string  `json:"control_id"` // required
	Source    *string `json:"source,omitempty"`
	Reason    string  `json:"reason"` // required
	RunID     *string `json:"run_id,omitempty"`
	TraceID   *string `json:"trace_id,omitempty"`
}

// FalsePositiveReportResponse maps the `FalsePositiveReportResponse` schema.
type FalsePositiveReportResponse struct {
	Status     *string `json:"status,omitempty"`
	FeedbackID string  `json:"feedback_id"` // required
	Framework  string  `json:"framework"`   // required
	ControlID  string  `json:"control_id"`  // required
}

// FlowEdge maps the `FlowEdge` schema.
type FlowEdge struct {
	Agent   string `json:"agent"`   // required
	Tool    string `json:"tool"`    // required
	Outcome string `json:"outcome"` // required
	Count   int    `json:"count"`   // required
}

// FlowResponse maps the `FlowResponse` schema.
type FlowResponse struct {
	Flows          []FlowEdge     `json:"flows"`           // required
	DistinctAgents int            `json:"distinct_agents"` // required
	DistinctTools  int            `json:"distinct_tools"`  // required
	Period         TopUsersPeriod `json:"period"`          // required
}

// FrameworkActivationResponse maps the `FrameworkActivationResponse` schema.
type FrameworkActivationResponse struct {
	Framework   string `json:"framework"`    // required
	Active      bool   `json:"active"`       // required
	ActivatedBy string `json:"activated_by"` // required
	ActivatedAt string `json:"activated_at"` // required
}

// FrameworkControlItem maps the `FrameworkControlItem` schema.
type FrameworkControlItem struct {
	ControlID             string   `json:"control_id"`  // required
	Objective             string   `json:"objective"`   // required
	RiskWeight            int      `json:"risk_weight"` // required
	ExpectedSignalSources []string `json:"expected_signal_sources,omitempty"`
	EvidenceTemplates     []string `json:"evidence_templates,omitempty"`
	DocumentationRefs     []string `json:"documentation_refs,omitempty"`
}

// FrameworkPackItem maps the `FrameworkPackItem` schema.
type FrameworkPackItem struct {
	Framework         string                 `json:"framework"`    // required
	DisplayName       string                 `json:"display_name"` // required
	Version           string                 `json:"version"`      // required
	Description       *string                `json:"description,omitempty"`
	DocumentationRefs []string               `json:"documentation_refs,omitempty"`
	ControlCount      int                    `json:"control_count"` // required
	Controls          []FrameworkControlItem `json:"controls,omitempty"`
	Active            *bool                  `json:"active,omitempty"`
	ActivatedAt       *string                `json:"activated_at,omitempty"`
	ActivatedBy       *string                `json:"activated_by,omitempty"`
}

// FrameworkReadinessResponse maps the `FrameworkReadinessResponse` schema.
type FrameworkReadinessResponse struct {
	Framework      string           `json:"framework"`       // required
	DisplayName    string           `json:"display_name"`    // required
	Version        string           `json:"version"`         // required
	ControlsTotal  int              `json:"controls_total"`  // required
	ReadinessScore float64          `json:"readiness_score"` // required
	CoverageRatio  float64          `json:"coverage_ratio"`  // required
	StatusCounts   map[string]int   `json:"status_counts,omitempty"`
	Items          []map[string]any `json:"items,omitempty"`
	Gaps           []map[string]any `json:"gaps,omitempty"`
	OrgID          *string          `json:"org_id,omitempty"`
	LookbackDays   *int             `json:"lookback_days,omitempty"`
}

// GDPRLegalBasis is a string enumeration (GDPRLegalBasis).
type GDPRLegalBasis string

// Enumerated values for GDPRLegalBasis.
const (
	GDPRLegalBasisConsent             GDPRLegalBasis = "consent"
	GDPRLegalBasisContract            GDPRLegalBasis = "contract"
	GDPRLegalBasisLegalObligation     GDPRLegalBasis = "legal_obligation"
	GDPRLegalBasisVitalInterests      GDPRLegalBasis = "vital_interests"
	GDPRLegalBasisPublicTask          GDPRLegalBasis = "public_task"
	GDPRLegalBasisLegitimateInterests GDPRLegalBasis = "legitimate_interests"
)

// GeoCountryRow maps the `GeoCountryRow` schema.
type GeoCountryRow struct {
	Country     string  `json:"country"` // required
	CountryCode *string `json:"country_code,omitempty"`
	EventCount  int     `json:"event_count"`  // required
	DistinctIps int     `json:"distinct_ips"` // required
}

// GeoMapResponse maps the `GeoMapResponse` schema.
type GeoMapResponse struct {
	Countries     []GeoCountryRow `json:"countries"`      // required
	TotalEvents   int             `json:"total_events"`   // required
	UnmappedCount int             `json:"unmapped_count"` // required
	Period        TopUsersPeriod  `json:"period"`         // required
}

// HTTPValidationError maps the `HTTPValidationError` schema.
type HTTPValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// HealthCheckEntry maps the `HealthCheckEntry` schema.
type HealthCheckEntry struct {
	Status  string `json:"status"`  // required
	Message string `json:"message"` // required
}

// HealthResponse maps the `HealthResponse` schema.
type HealthResponse struct {
	Status    string `json:"status"`    // required
	Timestamp string `json:"timestamp"` // required
	Service   string `json:"service"`   // required
}

// IngestBatchResponse maps the `IngestBatchResponse` schema.
type IngestBatchResponse struct {
	Status   string   `json:"status"`    // required
	EventIDs []string `json:"event_ids"` // required
	Count    int      `json:"count"`     // required
}

// IngestEventResponse maps the `IngestEventResponse` schema.
type IngestEventResponse struct {
	Status  string `json:"status"`   // required
	EventID string `json:"event_id"` // required
}

// IngestHealthResponse maps the `IngestHealthResponse` schema.
type IngestHealthResponse struct {
	LookbackHours   int                      `json:"lookback_hours"` // required
	GeneratedAt     string                   `json:"generated_at"`   // required
	Services        []IngestHealthServiceRow `json:"services"`       // required
	ExpectedMissing []string                 `json:"expected_missing,omitempty"`
}

// IngestHealthServiceRow maps the `IngestHealthServiceRow` schema.
type IngestHealthServiceRow struct {
	Service            string         `json:"service"`          // required
	EventsInWindow     int            `json:"events_in_window"` // required
	LastSeenAt         *string        `json:"last_seen_at,omitempty"`
	EventsPerMinuteAvg *float64       `json:"events_per_minute_avg,omitempty"`
	EventsPerMinuteP95 *float64       `json:"events_per_minute_p95,omitempty"`
	OutcomesBreakdown  map[string]int `json:"outcomes_breakdown,omitempty"`
	Freshness          *string        `json:"freshness,omitempty"`
}

// IngestLogBatchResponse maps the `IngestLogBatchResponse` schema.
type IngestLogBatchResponse struct {
	Status   string   `json:"status"`    // required
	EventIDs []string `json:"event_ids"` // required
	Count    int      `json:"count"`     // required
	Message  string   `json:"message"`   // required
}

// IngestLogResponse maps the `IngestLogResponse` schema.
type IngestLogResponse struct {
	Status  string `json:"status"`   // required
	EventID string `json:"event_id"` // required
	Message string `json:"message"`  // required
}

// IngestStreamResponse maps the `IngestStreamResponse` schema.
type IngestStreamResponse struct {
	Status           string               `json:"status"`            // required
	EventsReceived   int                  `json:"events_received"`   // required
	Duration         float64              `json:"duration"`          // required
	Rate             float64              `json:"rate"`              // required
	Results          []IngestStreamResult `json:"results"`           // required
	ResultsTruncated bool                 `json:"results_truncated"` // required
}

// IngestStreamResult maps the `IngestStreamResult` schema.
type IngestStreamResult struct {
	EventID *string `json:"event_id,omitempty"`
	Status  *string `json:"status,omitempty"`
	Error   *string `json:"error,omitempty"`
	Line    *int    `json:"line,omitempty"`
	Details *string `json:"details,omitempty"`
}

// IntegrityVerifyRequest maps the `IntegrityVerifyRequest` schema.
type IntegrityVerifyRequest struct {
	Table     string  `json:"table"` // required
	OrgID     *string `json:"org_id,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate   *string `json:"end_date,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

// JudgePolicyItem maps the `JudgePolicyItem` schema.
type JudgePolicyItem struct {
	PolicyID          string         `json:"policy_id"`      // required
	TenantID          string         `json:"tenant_id"`      // required
	OrgID             string         `json:"org_id"`         // required
	ControlID         string         `json:"control_id"`     // required
	ModelName         string         `json:"model_name"`     // required
	PromptKey         string         `json:"prompt_key"`     // required
	PromptVersion     int            `json:"prompt_version"` // required
	AllowedRuleStates []string       `json:"allowed_rule_states,omitempty"`
	AllowedRiskLevels []string       `json:"allowed_risk_levels,omitempty"`
	MinRuleConfidence float64        `json:"min_rule_confidence"` // required
	Enabled           bool           `json:"enabled"`             // required
	Metadata          map[string]any `json:"metadata,omitempty"`
	CreatedBy         *string        `json:"created_by,omitempty"`
	CreatedAt         *string        `json:"created_at,omitempty"`
	UpdatedAt         *string        `json:"updated_at,omitempty"`
}

// JudgePolicyUpsertRequest maps the `JudgePolicyUpsertRequest` schema.
type JudgePolicyUpsertRequest struct {
	OrgID             *string        `json:"org_id,omitempty"`
	ControlID         *string        `json:"control_id,omitempty"`
	ModelName         *string        `json:"model_name,omitempty"`
	PromptKey         *string        `json:"prompt_key,omitempty"`
	PromptVersion     *int           `json:"prompt_version,omitempty"`
	AllowedRuleStates []string       `json:"allowed_rule_states,omitempty"`
	AllowedRiskLevels []string       `json:"allowed_risk_levels,omitempty"`
	MinRuleConfidence *float64       `json:"min_rule_confidence,omitempty"`
	Enabled           *bool          `json:"enabled,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
}

// JudgePolicyUpsertResponse maps the `JudgePolicyUpsertResponse` schema.
type JudgePolicyUpsertResponse struct {
	Status            *string  `json:"status,omitempty"`
	PolicyID          string   `json:"policy_id"`           // required
	ControlID         string   `json:"control_id"`          // required
	ModelName         string   `json:"model_name"`          // required
	PromptKey         string   `json:"prompt_key"`          // required
	PromptVersion     int      `json:"prompt_version"`      // required
	AllowedRuleStates []string `json:"allowed_rule_states"` // required
	AllowedRiskLevels []string `json:"allowed_risk_levels"` // required
	MinRuleConfidence float64  `json:"min_rule_confidence"` // required
	Enabled           bool     `json:"enabled"`             // required
}

// JudgePromptActivateResponse maps the `JudgePromptActivateResponse` schema.
type JudgePromptActivateResponse struct {
	Status        *string `json:"status,omitempty"`
	PromptKey     string  `json:"prompt_key"`     // required
	ActiveVersion int     `json:"active_version"` // required
}

// JudgePromptItem maps the `JudgePromptItem` schema.
type JudgePromptItem struct {
	PromptID     string         `json:"prompt_id"`  // required
	TenantID     string         `json:"tenant_id"`  // required
	OrgID        string         `json:"org_id"`     // required
	PromptKey    string         `json:"prompt_key"` // required
	Version      int            `json:"version"`    // required
	Template     string         `json:"template"`   // required
	OutputSchema map[string]any `json:"output_schema,omitempty"`
	Enabled      bool           `json:"enabled"` // required
	IsActive     *bool          `json:"is_active,omitempty"`
	CreatedBy    *string        `json:"created_by,omitempty"`
	CreatedAt    *string        `json:"created_at,omitempty"`
	UpdatedAt    *string        `json:"updated_at,omitempty"`
}

// JudgePromptUpsertRequest maps the `JudgePromptUpsertRequest` schema.
type JudgePromptUpsertRequest struct {
	OrgID        *string        `json:"org_id,omitempty"`
	PromptKey    *string        `json:"prompt_key,omitempty"`
	Version      int            `json:"version"`  // required
	Template     string         `json:"template"` // required
	OutputSchema map[string]any `json:"output_schema,omitempty"`
	Enabled      *bool          `json:"enabled,omitempty"`
	Activate     *bool          `json:"activate,omitempty"`
}

// JudgePromptUpsertResponse maps the `JudgePromptUpsertResponse` schema.
type JudgePromptUpsertResponse struct {
	Status    *string `json:"status,omitempty"`
	PromptID  string  `json:"prompt_id"`  // required
	PromptKey string  `json:"prompt_key"` // required
	Version   int     `json:"version"`    // required
	Enabled   bool    `json:"enabled"`    // required
	IsActive  bool    `json:"is_active"`  // required
}

// JudgeReplayResponse maps the `JudgeReplayResponse` schema.
type JudgeReplayResponse struct {
	Status *string        `json:"status,omitempty"`
	Replay map[string]any `json:"replay,omitempty"`
}

// JudgeVerdictItem maps the `JudgeVerdictItem` schema.
type JudgeVerdictItem struct {
	VerdictID     string  `json:"verdict_id"` // required
	TraceID       *string `json:"trace_id,omitempty"`
	RunID         *string `json:"run_id,omitempty"`
	TenantID      string  `json:"tenant_id"` // required
	OrgID         string  `json:"org_id"`    // required
	AssertionID   *string `json:"assertion_id,omitempty"`
	ControlID     string  `json:"control_id"`     // required
	ModelName     string  `json:"model_name"`     // required
	PromptVersion int     `json:"prompt_version"` // required
	Verdict       string  `json:"verdict"`        // required
	Confidence    float64 `json:"confidence"`     // required
	Rationale     *string `json:"rationale,omitempty"`
	Citations     *string `json:"citations,omitempty"`
	InputPayload  *string `json:"input_payload,omitempty"`
	RawResponse   *string `json:"raw_response,omitempty"`
	PolicyID      *string `json:"policy_id,omitempty"`
	PolicyReason  *string `json:"policy_reason,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
}

// LegalHoldCreate maps the `LegalHoldCreate` schema.
type LegalHoldCreate struct {
	ScopeType     LegalHoldScope `json:"scope_type"` // required
	DataSubjectID *string        `json:"data_subject_id,omitempty"`
	BeforeDate    *string        `json:"before_date,omitempty"`
	Reason        string         `json:"reason"` // required
}

// LegalHoldRelease maps the `LegalHoldRelease` schema.
type LegalHoldRelease struct {
	Reason string `json:"reason"` // required
}

// LegalHoldScope is a string enumeration (LegalHoldScope).
type LegalHoldScope string

// Enumerated values for LegalHoldScope.
const (
	LegalHoldScopeOrg         LegalHoldScope = "org"
	LegalHoldScopeDataSubject LegalHoldScope = "data_subject"
	LegalHoldScopeBeforeDate  LegalHoldScope = "before_date"
)

// ListEnvelopeCallbackDeliveryItem maps the `ListEnvelope[CallbackDeliveryItem]` schema.
type ListEnvelopeCallbackDeliveryItem struct {
	Total int                    `json:"total"` // required
	Items []CallbackDeliveryItem `json:"items,omitempty"`
}

// ListEnvelopeCaseApprovalItem maps the `ListEnvelope[CaseApprovalItem]` schema.
type ListEnvelopeCaseApprovalItem struct {
	Total int                `json:"total"` // required
	Items []CaseApprovalItem `json:"items,omitempty"`
}

// ListEnvelopeCaseEscalationItem maps the `ListEnvelope[CaseEscalationItem]` schema.
type ListEnvelopeCaseEscalationItem struct {
	Total int                  `json:"total"` // required
	Items []CaseEscalationItem `json:"items,omitempty"`
}

// ListEnvelopeCaseTimelineEntry maps the `ListEnvelope[CaseTimelineEntry]` schema.
type ListEnvelopeCaseTimelineEntry struct {
	Total int                 `json:"total"` // required
	Items []CaseTimelineEntry `json:"items,omitempty"`
}

// ListEnvelopeComplianceTemplateItem maps the `ListEnvelope[ComplianceTemplateItem]` schema.
type ListEnvelopeComplianceTemplateItem struct {
	Total int                      `json:"total"` // required
	Items []ComplianceTemplateItem `json:"items,omitempty"`
}

// ListEnvelopeExplorerHistoryItem maps the `ListEnvelope[ExplorerHistoryItem]` schema.
type ListEnvelopeExplorerHistoryItem struct {
	Total int                   `json:"total"` // required
	Items []ExplorerHistoryItem `json:"items,omitempty"`
}

// ListEnvelopeExplorerSavedQueryItem maps the `ListEnvelope[ExplorerSavedQueryItem]` schema.
type ListEnvelopeExplorerSavedQueryItem struct {
	Total int                      `json:"total"` // required
	Items []ExplorerSavedQueryItem `json:"items,omitempty"`
}

// ListEnvelopeFalsePositiveFeedbackItem maps the `ListEnvelope[FalsePositiveFeedbackItem]` schema.
type ListEnvelopeFalsePositiveFeedbackItem struct {
	Total int                         `json:"total"` // required
	Items []FalsePositiveFeedbackItem `json:"items,omitempty"`
}

// ListEnvelopeFrameworkPackItem maps the `ListEnvelope[FrameworkPackItem]` schema.
type ListEnvelopeFrameworkPackItem struct {
	Total int                 `json:"total"` // required
	Items []FrameworkPackItem `json:"items,omitempty"`
}

// ListEnvelopeJudgePolicyItem maps the `ListEnvelope[JudgePolicyItem]` schema.
type ListEnvelopeJudgePolicyItem struct {
	Total int               `json:"total"` // required
	Items []JudgePolicyItem `json:"items,omitempty"`
}

// ListEnvelopeJudgePromptItem maps the `ListEnvelope[JudgePromptItem]` schema.
type ListEnvelopeJudgePromptItem struct {
	Total int               `json:"total"` // required
	Items []JudgePromptItem `json:"items,omitempty"`
}

// ListEnvelopeJudgeVerdictItem maps the `ListEnvelope[JudgeVerdictItem]` schema.
type ListEnvelopeJudgeVerdictItem struct {
	Total int                `json:"total"` // required
	Items []JudgeVerdictItem `json:"items,omitempty"`
}

// ListEnvelopeMergePolicyItem maps the `ListEnvelope[MergePolicyItem]` schema.
type ListEnvelopeMergePolicyItem struct {
	Total int               `json:"total"` // required
	Items []MergePolicyItem `json:"items,omitempty"`
}

// ListEnvelopeMlBaselineItem maps the `ListEnvelope[MlBaselineItem]` schema.
type ListEnvelopeMlBaselineItem struct {
	Total int              `json:"total"` // required
	Items []MlBaselineItem `json:"items,omitempty"`
}

// ListEnvelopeMlConfigItem maps the `ListEnvelope[MlConfigItem]` schema.
type ListEnvelopeMlConfigItem struct {
	Total int            `json:"total"` // required
	Items []MlConfigItem `json:"items,omitempty"`
}

// ListEnvelopeModelActivationItem maps the `ListEnvelope[ModelActivationItem]` schema.
type ListEnvelopeModelActivationItem struct {
	Total int                   `json:"total"` // required
	Items []ModelActivationItem `json:"items,omitempty"`
}

// ListEnvelopeModelArtifactItem maps the `ListEnvelope[ModelArtifactItem]` schema.
type ListEnvelopeModelArtifactItem struct {
	Total int                 `json:"total"` // required
	Items []ModelArtifactItem `json:"items,omitempty"`
}

// ListEnvelopeModelScoreItem maps the `ListEnvelope[ModelScoreItem]` schema.
type ListEnvelopeModelScoreItem struct {
	Total int              `json:"total"` // required
	Items []ModelScoreItem `json:"items,omitempty"`
}

// ListEnvelopePolicyRuleItem maps the `ListEnvelope[PolicyRuleItem]` schema.
type ListEnvelopePolicyRuleItem struct {
	Total int              `json:"total"` // required
	Items []PolicyRuleItem `json:"items,omitempty"`
}

// ListEnvelopeSavedSearchItem maps the `ListEnvelope[SavedSearchItem]` schema.
type ListEnvelopeSavedSearchItem struct {
	Total int               `json:"total"` // required
	Items []SavedSearchItem `json:"items,omitempty"`
}

// ListEnvelopeTrainingJobItem maps the `ListEnvelope[TrainingJobItem]` schema.
type ListEnvelopeTrainingJobItem struct {
	Total int               `json:"total"` // required
	Items []TrainingJobItem `json:"items,omitempty"`
}

// LivenessResponse maps the `LivenessResponse` schema.
type LivenessResponse struct {
	Status string `json:"status"` // required
}

// MaintenanceOptimizeResponse maps the `MaintenanceOptimizeResponse` schema.
type MaintenanceOptimizeResponse struct {
	Message   string `json:"message"`   // required
	Timestamp string `json:"timestamp"` // required
}

// MergePolicyItem maps the `MergePolicyItem` schema.
type MergePolicyItem struct {
	PolicyID      string         `json:"policy_id"`      // required
	TenantID      string         `json:"tenant_id"`      // required
	OrgID         string         `json:"org_id"`         // required
	ControlID     string         `json:"control_id"`     // required
	RuleWeight    float64        `json:"rule_weight"`    // required
	MLWeight      float64        `json:"ml_weight"`      // required
	JudgeWeight   float64        `json:"judge_weight"`   // required
	WarnThreshold float64        `json:"warn_threshold"` // required
	FailThreshold float64        `json:"fail_threshold"` // required
	FailOverride  bool           `json:"fail_override"`  // required
	Enabled       bool           `json:"enabled"`        // required
	Metadata      map[string]any `json:"metadata,omitempty"`
	CreatedBy     *string        `json:"created_by,omitempty"`
	CreatedAt     *string        `json:"created_at,omitempty"`
	UpdatedAt     *string        `json:"updated_at,omitempty"`
}

// MergePolicyUpsertRequest maps the `MergePolicyUpsertRequest` schema.
type MergePolicyUpsertRequest struct {
	OrgID         *string        `json:"org_id,omitempty"`
	ControlID     *string        `json:"control_id,omitempty"`
	RuleWeight    *float64       `json:"rule_weight,omitempty"`
	MLWeight      *float64       `json:"ml_weight,omitempty"`
	JudgeWeight   *float64       `json:"judge_weight,omitempty"`
	WarnThreshold *float64       `json:"warn_threshold,omitempty"`
	FailThreshold *float64       `json:"fail_threshold,omitempty"`
	FailOverride  *bool          `json:"fail_override,omitempty"`
	Enabled       *bool          `json:"enabled,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
}

// MergePolicyUpsertResponse maps the `MergePolicyUpsertResponse` schema.
type MergePolicyUpsertResponse struct {
	Status *string         `json:"status,omitempty"`
	Policy MergePolicyItem `json:"policy"` // required
}

// MlBaselineItem maps the `MlBaselineItem` schema.
type MlBaselineItem struct {
	BaselineID          string  `json:"baseline_id"`      // required
	TenantID            string  `json:"tenant_id"`        // required
	OrgID               string  `json:"org_id"`           // required
	ControlID           string  `json:"control_id"`       // required
	ModelName           string  `json:"model_name"`       // required
	MeanValue           float64 `json:"mean_value"`       // required
	StddevValue         float64 `json:"stddev_value"`     // required
	SampleCount         int     `json:"sample_count"`     // required
	ThresholdZscore     float64 `json:"threshold_zscore"` // required
	LookbackDays        int     `json:"lookback_days"`    // required
	CalibrationMetadata *string `json:"calibration_metadata,omitempty"`
	UpdatedAt           *string `json:"updated_at,omitempty"`
}

// MlConfigItem maps the `MlConfigItem` schema.
type MlConfigItem struct {
	ConfigID        string         `json:"config_id"`        // required
	TenantID        string         `json:"tenant_id"`        // required
	OrgID           string         `json:"org_id"`           // required
	ControlID       string         `json:"control_id"`       // required
	ModelName       string         `json:"model_name"`       // required
	ThresholdZscore float64        `json:"threshold_zscore"` // required
	LookbackDays    int            `json:"lookback_days"`    // required
	MinSamples      int            `json:"min_samples"`      // required
	Enabled         bool           `json:"enabled"`          // required
	Metadata        map[string]any `json:"metadata,omitempty"`
	CreatedBy       *string        `json:"created_by,omitempty"`
	CreatedAt       *string        `json:"created_at,omitempty"`
	UpdatedAt       *string        `json:"updated_at,omitempty"`
}

// MlConfigUpsertRequest maps the `MlConfigUpsertRequest` schema.
type MlConfigUpsertRequest struct {
	OrgID           *string        `json:"org_id,omitempty"`
	ControlID       string         `json:"control_id"` // required
	ModelName       *string        `json:"model_name,omitempty"`
	ThresholdZscore *float64       `json:"threshold_zscore,omitempty"`
	LookbackDays    *int           `json:"lookback_days,omitempty"`
	MinSamples      *int           `json:"min_samples,omitempty"`
	Enabled         *bool          `json:"enabled,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}

// MlConfigUpsertResponse maps the `MlConfigUpsertResponse` schema.
type MlConfigUpsertResponse struct {
	Status          *string `json:"status,omitempty"`
	ConfigID        string  `json:"config_id"`        // required
	ControlID       string  `json:"control_id"`       // required
	ModelName       string  `json:"model_name"`       // required
	ThresholdZscore float64 `json:"threshold_zscore"` // required
	LookbackDays    int     `json:"lookback_days"`    // required
	MinSamples      int     `json:"min_samples"`      // required
	Enabled         bool    `json:"enabled"`          // required
}

// MlDistillJobResponse maps the `MlDistillJobResponse` schema.
type MlDistillJobResponse struct {
	Status      *string        `json:"status,omitempty"`
	JobID       string         `json:"job_id"` // required
	TraceID     *string        `json:"trace_id,omitempty"`
	RunID       *string        `json:"run_id,omitempty"`
	OrgID       string         `json:"org_id"`       // required
	TenantID    string         `json:"tenant_id"`    // required
	JobType     string         `json:"job_type"`     // required
	RequestedAt string         `json:"requested_at"` // required
	Request     map[string]any `json:"request,omitempty"`
}

// MlDistillationTrainRequest maps the `MlDistillationTrainRequest` schema.
type MlDistillationTrainRequest struct {
	OrgID          *string  `json:"org_id,omitempty"`
	ModelFamily    *string  `json:"model_family,omitempty"`
	LookbackDays   *int     `json:"lookback_days,omitempty"`
	Limit          *int     `json:"limit,omitempty"`
	Epochs         *int     `json:"epochs,omitempty"`
	LearningRate   *float64 `json:"learning_rate,omitempty"`
	MinExamples    *int     `json:"min_examples,omitempty"`
	ApplyAsDefault *bool    `json:"apply_as_default,omitempty"`
}

// MlModelActivationRequest maps the `MlModelActivationRequest` schema.
type MlModelActivationRequest struct {
	OrgID  *string `json:"org_id,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// MlModelRollbackRequest maps the `MlModelRollbackRequest` schema.
type MlModelRollbackRequest struct {
	OrgID     *string `json:"org_id,omitempty"`
	ModelName *string `json:"model_name,omitempty"`
	Reason    *string `json:"reason,omitempty"`
}

// ModelActivationItem maps the `ModelActivationItem` schema.
type ModelActivationItem struct {
	ActivationID string  `json:"activation_id"` // required
	TenantID     string  `json:"tenant_id"`     // required
	OrgID        string  `json:"org_id"`        // required
	ControlID    string  `json:"control_id"`    // required
	ModelName    string  `json:"model_name"`    // required
	ArtifactID   string  `json:"artifact_id"`   // required
	Action       string  `json:"action"`        // required
	Reason       *string `json:"reason,omitempty"`
	ActivatedBy  *string `json:"activated_by,omitempty"`
	CreatedAt    *string `json:"created_at,omitempty"`
}

// ModelArtifactActivateResponse maps the `ModelArtifactActivateResponse` schema.
type ModelArtifactActivateResponse struct {
	Status       *string `json:"status,omitempty"`
	ActivationID string  `json:"activation_id"` // required
	ArtifactID   string  `json:"artifact_id"`   // required
	OrgID        string  `json:"org_id"`        // required
	TenantID     string  `json:"tenant_id"`     // required
}

// ModelArtifactItem maps the `ModelArtifactItem` schema.
type ModelArtifactItem struct {
	ArtifactID      string         `json:"artifact_id"` // required
	JobID           *string        `json:"job_id,omitempty"`
	TraceID         *string        `json:"trace_id,omitempty"`
	RunID           *string        `json:"run_id,omitempty"`
	TenantID        string         `json:"tenant_id"`  // required
	OrgID           string         `json:"org_id"`     // required
	ControlID       string         `json:"control_id"` // required
	ModelName       string         `json:"model_name"` // required
	Version         *string        `json:"version,omitempty"`
	ConfigID        *string        `json:"config_id,omitempty"`
	BaselineID      *string        `json:"baseline_id,omitempty"`
	Metrics         map[string]any `json:"metrics,omitempty"`
	ArtifactPayload map[string]any `json:"artifact_payload,omitempty"`
	CreatedBy       *string        `json:"created_by,omitempty"`
	CreatedAt       *string        `json:"created_at,omitempty"`
}

// ModelControlRollbackResponse maps the `ModelControlRollbackResponse` schema.
type ModelControlRollbackResponse struct {
	Status       *string `json:"status,omitempty"`
	Action       *string `json:"action,omitempty"`
	ActivationID string  `json:"activation_id"` // required
	ArtifactID   string  `json:"artifact_id"`   // required
	ControlID    string  `json:"control_id"`    // required
	OrgID        string  `json:"org_id"`        // required
	TenantID     string  `json:"tenant_id"`     // required
}

// ModelScoreItem maps the `ModelScoreItem` schema.
type ModelScoreItem struct {
	ScoreID       string   `json:"score_id"` // required
	TraceID       *string  `json:"trace_id,omitempty"`
	RunID         *string  `json:"run_id,omitempty"`
	TenantID      string   `json:"tenant_id"`   // required
	OrgID         string   `json:"org_id"`      // required
	ControlID     string   `json:"control_id"`  // required
	ModelName     string   `json:"model_name"`  // required
	ScoreValue    float64  `json:"score_value"` // required
	Threshold     *float64 `json:"threshold,omitempty"`
	Label         *string  `json:"label,omitempty"`
	FeatureValue  *float64 `json:"feature_value,omitempty"`
	MeanValue     *float64 `json:"mean_value,omitempty"`
	StddevValue   *float64 `json:"stddev_value,omitempty"`
	Zscore        *float64 `json:"zscore,omitempty"`
	SampleCount   *int     `json:"sample_count,omitempty"`
	IsCalibrated  *bool    `json:"is_calibrated,omitempty"`
	DeltaFromMean *float64 `json:"delta_from_mean,omitempty"`
	ScoredAt      *string  `json:"scored_at,omitempty"`
}

// ModuleHealthResponse maps the `ModuleHealthResponse` schema.
type ModuleHealthResponse struct {
	Module                   string   `json:"module"`                      // required
	Status                   string   `json:"status"`                      // required
	CallbacksEnabled         bool     `json:"callbacks_enabled"`           // required
	LlmJudgeEnabled          bool     `json:"llm_judge_enabled"`           // required
	MLScoringEnabled         bool     `json:"ml_scoring_enabled"`          // required
	LlmJudgeProvider         string   `json:"llm_judge_provider"`          // required
	LlmJudgeModel            string   `json:"llm_judge_model"`             // required
	LlmJudgeRemoteConfigured bool     `json:"llm_judge_remote_configured"` // required
	LlmJudgeFailClosed       bool     `json:"llm_judge_fail_closed"`       // required
	StrictIdempotency        bool     `json:"strict_idempotency"`          // required
	CurrentEnvironment       string   `json:"current_environment"`         // required
	AllowedEnvironments      []string `json:"allowed_environments"`        // required
	AllowedJudgeModels       []string `json:"allowed_judge_models"`        // required
}

// NotificationFeed maps the `NotificationFeed` schema.
// The full feed plus the count the header dot should reflect.
type NotificationFeed struct {
	UnreadCount *int               `json:"unread_count,omitempty"` // Items newer than the user's read marker (drives the header dot; 0 hides it)
	TotalCount  *int               `json:"total_count,omitempty"`  // Total actionable items currently in the feed (history; independent of read state)
	LastSeenAt  *string            `json:"last_seen_at,omitempty"` // When this user last acknowledged the feed; null if never
	Items       []NotificationItem `json:"items,omitempty"`
}

// NotificationItem maps the `NotificationItem` schema.
// One actionable signal surfaced in the header bell panel.
type NotificationItem struct {
	ID        string  `json:"id"`                   // Stable id for this notification (source row id, prefixed by kind)
	Kind      string  `json:"kind"`                 // report_failed | auditor_request | dsar_deadline | security_alert
	Title     string  `json:"title"`                // One-line headline shown in the panel
	Detail    *string `json:"detail,omitempty"`     // Secondary explanatory text
	Severity  *string `json:"severity,omitempty"`   // info | warning | critical
	CreatedAt *string `json:"created_at,omitempty"` // When the underlying event occurred
	DeepLink  *string `json:"deep_link,omitempty"`  // In-app URL that lets the user act on this item
}

// NotificationPrefsSettings maps the `NotificationPrefsSettings` schema.
type NotificationPrefsSettings struct {
	Categories map[string]any `json:"categories"` // required
	IsDefault  *bool          `json:"is_default,omitempty"`
	UpdatedBy  *string        `json:"updated_by,omitempty"`
	UpdatedAt  *string        `json:"updated_at,omitempty"`
}

// NotificationReadResponse maps the `NotificationReadResponse` schema.
// Result of marking the feed read.
type NotificationReadResponse struct {
	OK          *bool  `json:"ok,omitempty"`           // True when the read marker was persisted
	LastSeenAt  string `json:"last_seen_at"`           // The acknowledged-up-to timestamp now stored for this user
	UnreadCount *int   `json:"unread_count,omitempty"` // Unread count after acknowledging (0 unless a newer item arrived mid-flight)
}

// OrganizationSummary maps the `OrganizationSummary` schema.
type OrganizationSummary struct {
	OrgID          string  `json:"org_id"`       // required
	TotalEvents    int     `json:"total_events"` // required
	UniqueUsers    int     `json:"unique_users"` // required
	FirstEvent     *string `json:"first_event,omitempty"`
	LastEvent      *string `json:"last_event,omitempty"`
	EventsLastWeek int     `json:"events_last_week"` // required
	Active         bool    `json:"active"`           // required
}

// PaginatedResponseAuditEvent maps the `PaginatedResponse[AuditEvent]` schema.
type PaginatedResponseAuditEvent struct {
	Data       []AuditEvent `json:"data"`  // required
	Total      int          `json:"total"` // required
	Limit      int          `json:"limit"` // required
	Offset     *int         `json:"offset,omitempty"`
	Cursor     *string      `json:"cursor,omitempty"`
	NextCursor *string      `json:"next_cursor,omitempty"`
	HasMore    bool         `json:"has_more"` // required
}

// PaginatedResponseComplianceReportRow maps the `PaginatedResponse[ComplianceReportRow]` schema.
type PaginatedResponseComplianceReportRow struct {
	Data       []ComplianceReportRow `json:"data"`  // required
	Total      int                   `json:"total"` // required
	Limit      int                   `json:"limit"` // required
	Offset     *int                  `json:"offset,omitempty"`
	Cursor     *string               `json:"cursor,omitempty"`
	NextCursor *string               `json:"next_cursor,omitempty"`
	HasMore    bool                  `json:"has_more"` // required
}

// PaginatedResponseDsarRequestRow maps the `PaginatedResponse[DsarRequestRow]` schema.
type PaginatedResponseDsarRequestRow struct {
	Data       []DsarRequestRow `json:"data"`  // required
	Total      int              `json:"total"` // required
	Limit      int              `json:"limit"` // required
	Offset     *int             `json:"offset,omitempty"`
	Cursor     *string          `json:"cursor,omitempty"`
	NextCursor *string          `json:"next_cursor,omitempty"`
	HasMore    bool             `json:"has_more"` // required
}

// PaginatedResponseEvidencePackageWithLineage maps the `PaginatedResponse[EvidencePackageWithLineage]` schema.
type PaginatedResponseEvidencePackageWithLineage struct {
	Data       []EvidencePackageWithLineage `json:"data"`  // required
	Total      int                          `json:"total"` // required
	Limit      int                          `json:"limit"` // required
	Offset     *int                         `json:"offset,omitempty"`
	Cursor     *string                      `json:"cursor,omitempty"`
	NextCursor *string                      `json:"next_cursor,omitempty"`
	HasMore    bool                         `json:"has_more"` // required
}

// PaginatedResponseExportJobResponse maps the `PaginatedResponse[ExportJobResponse]` schema.
type PaginatedResponseExportJobResponse struct {
	Data       []ExportJobResponse `json:"data"`  // required
	Total      int                 `json:"total"` // required
	Limit      int                 `json:"limit"` // required
	Offset     *int                `json:"offset,omitempty"`
	Cursor     *string             `json:"cursor,omitempty"`
	NextCursor *string             `json:"next_cursor,omitempty"`
	HasMore    bool                `json:"has_more"` // required
}

// PaginatedResponseOrganizationSummary maps the `PaginatedResponse[OrganizationSummary]` schema.
type PaginatedResponseOrganizationSummary struct {
	Data       []OrganizationSummary `json:"data"`  // required
	Total      int                   `json:"total"` // required
	Limit      int                   `json:"limit"` // required
	Offset     *int                  `json:"offset,omitempty"`
	Cursor     *string               `json:"cursor,omitempty"`
	NextCursor *string               `json:"next_cursor,omitempty"`
	HasMore    bool                  `json:"has_more"` // required
}

// PolicyRuleItem maps the `PolicyRuleItem` schema.
type PolicyRuleItem struct {
	PolicyID          string         `json:"policy_id"`           // required
	TenantID          string         `json:"tenant_id"`           // required
	OrgID             string         `json:"org_id"`              // required
	ControlID         string         `json:"control_id"`          // required
	Version           int            `json:"version"`             // required
	MinConfidencePass float64        `json:"min_confidence_pass"` // required
	MinConfidenceWarn float64        `json:"min_confidence_warn"` // required
	Enabled           bool           `json:"enabled"`             // required
	Metadata          map[string]any `json:"metadata,omitempty"`
	CreatedBy         *string        `json:"created_by,omitempty"`
	CreatedAt         *string        `json:"created_at,omitempty"`
	UpdatedAt         *string        `json:"updated_at,omitempty"`
}

// PolicyRuleUpsertRequest maps the `PolicyRuleUpsertRequest` schema.
type PolicyRuleUpsertRequest struct {
	OrgID             *string        `json:"org_id,omitempty"`
	ControlID         string         `json:"control_id"` // required
	MinConfidencePass *float64       `json:"min_confidence_pass,omitempty"`
	MinConfidenceWarn *float64       `json:"min_confidence_warn,omitempty"`
	Enabled           *bool          `json:"enabled,omitempty"`
	Version           *int           `json:"version,omitempty"`
	Metadata          map[string]any `json:"metadata,omitempty"`
}

// PolicyRuleUpsertResponse maps the `PolicyRuleUpsertResponse` schema.
type PolicyRuleUpsertResponse struct {
	Status    *string `json:"status,omitempty"`
	PolicyID  string  `json:"policy_id"`  // required
	ControlID string  `json:"control_id"` // required
	Version   int     `json:"version"`    // required
	Enabled   bool    `json:"enabled"`    // required
}

// PostureHistoryPoint maps the `PostureHistoryPoint` schema.
// One sample in the posture-history time series.
type PostureHistoryPoint struct {
	Date             string             `json:"date"` // required
	OverallScore     *float64           `json:"overall_score,omitempty"`
	FrameworkScores  map[string]float64 `json:"framework_scores,omitempty"`
	SampleEventCount *int               `json:"sample_event_count,omitempty"`
}

// PostureHistoryResponse maps the `PostureHistoryResponse` schema.
// Response shape for GET /compliance/posture-history.
type PostureHistoryResponse struct {
	OrgID       string                `json:"org_id"`       // required
	WindowDays  int                   `json:"window_days"`  // required
	Points      []PostureHistoryPoint `json:"points"`       // required
	GeneratedAt string                `json:"generated_at"` // required
}

// PurgeDataResponse maps the `PurgeDataResponse` schema.
type PurgeDataResponse struct {
	DryRun             bool    `json:"dry_run"` // required
	RecordsToPurge     *int    `json:"records_to_purge,omitempty"`
	RecordsPurged      *int    `json:"records_purged,omitempty"`
	Message            string  `json:"message"` // required
	DeletionID         *string `json:"deletion_id,omitempty"`
	Status             *string `json:"status,omitempty"`
	RecordsQuarantined *int    `json:"records_quarantined,omitempty"`
}

// ReadinessResponse maps the `ReadinessResponse` schema.
type ReadinessResponse struct {
	Status string `json:"status"` // required
}

// RedactionSettings maps the `RedactionSettings` schema.
type RedactionSettings struct {
	SensitiveHeaders  []string `json:"sensitive_headers"`   // required
	SensitiveJSONKeys []string `json:"sensitive_json_keys"` // required
	IsDefault         *bool    `json:"is_default,omitempty"`
	UpdatedBy         *string  `json:"updated_by,omitempty"`
	UpdatedAt         *string  `json:"updated_at,omitempty"`
}

// ReportScheduleCreate maps the `ReportScheduleCreate` schema.
type ReportScheduleCreate struct {
	Name        string   `json:"name"`         // required
	Standard    string   `json:"standard"`     // required
	CadenceCron string   `json:"cadence_cron"` // required
	Recipients  []string `json:"recipients,omitempty"`
}

// ReportScheduleDeleteResponse maps the `ReportScheduleDeleteResponse` schema.
type ReportScheduleDeleteResponse struct {
	ScheduleID string `json:"schedule_id"` // required
	Deleted    bool   `json:"deleted"`     // required
}

// ReportScheduleItem maps the `ReportScheduleItem` schema.
type ReportScheduleItem struct {
	ScheduleID  string   `json:"schedule_id"`  // required
	OrgID       string   `json:"org_id"`       // required
	Name        string   `json:"name"`         // required
	Standard    string   `json:"standard"`     // required
	CadenceCron string   `json:"cadence_cron"` // required
	Recipients  []string `json:"recipients,omitempty"`
	LastRunAt   *string  `json:"last_run_at,omitempty"`
	NextRunAt   *string  `json:"next_run_at,omitempty"`
	Paused      *bool    `json:"paused,omitempty"`
	CreatedBy   *string  `json:"created_by,omitempty"`
	CreatedAt   *string  `json:"created_at,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
}

// ReportScheduleListResponse maps the `ReportScheduleListResponse` schema.
// List envelope for report schedules (T18 / CS-028). Mirrors
// `ListEnvelope[ReportScheduleItem]` but adds `scheduling_enabled` so the
// UI can tell the user the truth: when the scheduled-reports worker is
// disabled in this environment, schedules are stored but will NOT fire.
// Without this flag the UI would show an "Active / Next run" row that never
// actually runs.
type ReportScheduleListResponse struct {
	Total             int                  `json:"total"` // required
	Items             []ReportScheduleItem `json:"items,omitempty"`
	SchedulingEnabled *bool                `json:"scheduling_enabled,omitempty"`
}

// ReportSchedulePauseResponse maps the `ReportSchedulePauseResponse` schema.
type ReportSchedulePauseResponse struct {
	ScheduleID string `json:"schedule_id"` // required
	Paused     bool   `json:"paused"`      // required
}

// RetentionPeriod is a string enumeration (RetentionPeriod).
type RetentionPeriod string

// Enumerated values for RetentionPeriod.
const (
	RetentionPeriodN30d       RetentionPeriod = "30d"
	RetentionPeriodN90d       RetentionPeriod = "90d"
	RetentionPeriodN6m        RetentionPeriod = "6m"
	RetentionPeriodN1y        RetentionPeriod = "1y"
	RetentionPeriodN3y        RetentionPeriod = "3y"
	RetentionPeriodN7y        RetentionPeriod = "7y"
	RetentionPeriodN10y       RetentionPeriod = "10y"
	RetentionPeriodIndefinite RetentionPeriod = "indefinite"
)

// RetentionPoliciesResponse maps the `RetentionPoliciesResponse` schema.
type RetentionPoliciesResponse struct {
	Policies []RetentionPolicySummary `json:"policies"` // required
}

// RetentionPolicySummary maps the `RetentionPolicySummary` schema.
type RetentionPolicySummary struct {
	OrgID            string `json:"org_id"`             // required
	RetentionDays    int    `json:"retention_days"`     // required
	ArchiveAfterDays int    `json:"archive_after_days"` // required
	HotDataDays      int    `json:"hot_data_days"`      // required
	CreatedAt        string `json:"created_at"`         // required
	UpdatedAt        string `json:"updated_at"`         // required
}

// RetentionPolicyUpdate maps the `RetentionPolicyUpdate` schema.
// CS-010 — per-org retention policy override (PUT
// /admin/retention/policies). All windows are in days. The validator
// enforces a sane ordering so the retention manager never tries to archive
// data it has already expired (hot_data_days <= archive_after_days <=
// retention_days).
type RetentionPolicyUpdate struct {
	OrgID            *string `json:"org_id,omitempty"`
	RetentionDays    int     `json:"retention_days"`     // required
	ArchiveAfterDays int     `json:"archive_after_days"` // required
	HotDataDays      int     `json:"hot_data_days"`      // required
}

// RiskCasesResponse maps the `RiskCasesResponse` schema.
type RiskCasesResponse struct {
	Total int              `json:"total"` // required
	Items []map[string]any `json:"items,omitempty"`
}

// RolloutPreviewResponse maps the `RolloutPreviewResponse` schema.
type RolloutPreviewResponse struct {
	Component string  `json:"component"` // required
	RunID     string  `json:"run_id"`    // required
	Percent   float64 `json:"percent"`   // required
	Salt      string  `json:"salt"`      // required
	Selected  bool    `json:"selected"`  // required
}

// RunTraceResponse maps the `RunTraceResponse` schema.
type RunTraceResponse struct {
	RunID     string                      `json:"run_id"` // required
	TraceID   *string                     `json:"trace_id,omitempty"`
	TenantID  string                      `json:"tenant_id"` // required
	OrgID     string                      `json:"org_id"`    // required
	Sections  map[string]int              `json:"sections,omitempty"`
	TotalRows int                         `json:"total_rows"` // required
	Snapshot  map[string][]map[string]any `json:"snapshot,omitempty"`
}

// SavedSearchCreate maps the `SavedSearchCreate` schema.
type SavedSearchCreate struct {
	Name       string         `json:"name"` // required
	FilterJSON map[string]any `json:"filter_json,omitempty"`
}

// SavedSearchDeleteResponse maps the `SavedSearchDeleteResponse` schema.
type SavedSearchDeleteResponse struct {
	SearchID string `json:"search_id"` // required
	Deleted  bool   `json:"deleted"`   // required
}

// SavedSearchItem maps the `SavedSearchItem` schema.
type SavedSearchItem struct {
	SearchID   string         `json:"search_id"` // required
	UserID     string         `json:"user_id"`   // required
	OrgID      string         `json:"org_id"`    // required
	Name       string         `json:"name"`      // required
	FilterJSON map[string]any `json:"filter_json,omitempty"`
	CreatedAt  *string        `json:"created_at,omitempty"`
	UpdatedAt  *string        `json:"updated_at,omitempty"`
}

// SearchFacets maps the `SearchFacets` schema.
type SearchFacets struct {
	EventTypes    []FacetCount `json:"event_types,omitempty"`
	Actions       []FacetCount `json:"actions,omitempty"`
	Outcomes      []FacetCount `json:"outcomes,omitempty"`
	Services      []FacetCount `json:"services,omitempty"`
	ResourceTypes []FacetCount `json:"resource_types,omitempty"`
}

// SearchFacetsResponse maps the `SearchFacetsResponse` schema.
type SearchFacetsResponse struct {
	Facets SearchFacets `json:"facets"` // required
	OrgID  *string      `json:"org_id,omitempty"`
}

// SearchQuery maps the `SearchQuery` schema.
type SearchQuery struct {
	OrgID             *string  `json:"org_id,omitempty"`
	Query             string   `json:"query"` // required
	StartTime         *string  `json:"start_time,omitempty"`
	EndTime           *string  `json:"end_time,omitempty"`
	EventTypes        []string `json:"event_types,omitempty"`
	Actions           []string `json:"actions,omitempty"`
	Outcomes          []string `json:"outcomes,omitempty"`
	UserIDs           []string `json:"user_ids,omitempty"`
	Services          []string `json:"services,omitempty"`            // Filter by the emitting service (the faceted `service` column), e.g. ['audit-service']. ...
	EventTypePrefixes []string `json:"event_type_prefixes,omitempty"` // Structured category filter: keep only events whose `event_type` starts with one of thes...
	Limit             *int     `json:"limit,omitempty"`
	Offset            *int     `json:"offset,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"` // Cursor for keyset pagination (from next_cursor). When provided, offset is ignored.
}

// SearchResult maps the `SearchResult` schema.
type SearchResult struct {
	Data            []AuditEvent `json:"data"`              // required
	Total           int          `json:"total"`             // required
	Limit           int          `json:"limit"`             // required
	Offset          int          `json:"offset"`            // required
	HasMore         bool         `json:"has_more"`          // required
	Query           string       `json:"query"`             // required
	ExecutionTimeMs float64      `json:"execution_time_ms"` // required
	Cursor          *string      `json:"cursor,omitempty"`
	NextCursor      *string      `json:"next_cursor,omitempty"`
}

// SearchSuggestResponse maps the `SearchSuggestResponse` schema.
type SearchSuggestResponse struct {
	Suggestions []SearchSuggestion `json:"suggestions"` // required
	Field       string             `json:"field"`       // required
	Prefix      string             `json:"prefix"`      // required
}

// SearchSuggestion maps the `SearchSuggestion` schema.
type SearchSuggestion struct {
	Value     string `json:"value"`     // required
	Frequency int    `json:"frequency"` // required
}

// SecurityMetrics maps the `SecurityMetrics` schema.
type SecurityMetrics struct {
	FailedLogins         int              `json:"failed_logins"`         // required
	SuspiciousActivities int              `json:"suspicious_activities"` // required
	RateLimitViolations  int              `json:"rate_limit_violations"` // required
	AnomaliesDetected    int              `json:"anomalies_detected"`    // required
	SecurityAlerts       []map[string]any `json:"security_alerts"`       // required
}

// ServiceRootResponse maps the `ServiceRootResponse` schema.
type ServiceRootResponse struct {
	Service      string            `json:"service"`      // required
	Version      string            `json:"version"`      // required
	Status       string            `json:"status"`       // required
	Docs         map[string]string `json:"docs"`         // required
	UI           map[string]string `json:"ui"`           // required
	Capabilities []string          `json:"capabilities"` // required
}

// ShareLink maps the `ShareLink` schema.
// Time-bound external-auditor share token for an evidence package.
type ShareLink struct {
	Token          string  `json:"token"` // required
	OrgID          *string `json:"org_id,omitempty"`
	PackageID      string  `json:"package_id"`    // required
	AuditorEmail   string  `json:"auditor_email"` // required
	ExpiresAt      string  `json:"expires_at"`    // required
	CreatedBy      *string `json:"created_by,omitempty"`
	CreatedAt      *string `json:"created_at,omitempty"`
	RevokedAt      *string `json:"revoked_at,omitempty"`
	LastAccessedAt *string `json:"last_accessed_at,omitempty"`
	AccessCount    *int    `json:"access_count,omitempty"`
}

// ShareLinkCreate maps the `ShareLinkCreate` schema.
// Shape for POST /compliance/evidence-packages/{package_id}/share.
type ShareLinkCreate struct {
	AuditorEmail string `json:"auditor_email"` // required
	ExpiresAt    string `json:"expires_at"`    // required
}

// ShareLinkSummary maps the `ShareLinkSummary` schema.
// Manage-side view of a share link. Returned by GET
// /compliance/evidence-packages/{id}/share-links so an evidence manager can
// see who holds access, when it expires/was revoked, how often it's been
// opened, and re-copy or revoke it. `status` is server-computed so the UI
// doesn't have to re-derive expired/revoked.
type ShareLinkSummary struct {
	Token          string  `json:"token"`        // required
	TokenPrefix    string  `json:"token_prefix"` // required
	OrgID          *string `json:"org_id,omitempty"`
	PackageID      string  `json:"package_id"`    // required
	AuditorEmail   string  `json:"auditor_email"` // required
	ExpiresAt      string  `json:"expires_at"`    // required
	CreatedBy      *string `json:"created_by,omitempty"`
	CreatedAt      *string `json:"created_at,omitempty"`
	RevokedAt      *string `json:"revoked_at,omitempty"`
	LastAccessedAt *string `json:"last_accessed_at,omitempty"`
	AccessCount    *int    `json:"access_count,omitempty"`
	Status         *string `json:"status,omitempty"`
}

// Soc2AuditTrailEvent maps the `Soc2AuditTrailEvent` schema.
type Soc2AuditTrailEvent struct {
	EventID              string   `json:"event_id"`   // required
	Timestamp            string   `json:"timestamp"`  // required
	EventType            string   `json:"event_type"` // required
	Action               string   `json:"action"`     // required
	Outcome              string   `json:"outcome"`    // required
	UserID               *string  `json:"user_id,omitempty"`
	SessionID            *string  `json:"session_id,omitempty"`
	IPAddress            *string  `json:"ip_address,omitempty"`
	UserAgent            *string  `json:"user_agent,omitempty"`
	Metadata             *string  `json:"metadata,omitempty"`
	EventHash            *string  `json:"event_hash,omitempty"`
	SOC2ControlRelevance []string `json:"soc2_control_relevance,omitempty"`
	IntegrityVerified    bool     `json:"integrity_verified"` // required
}

// Soc2AuditTrailResponse maps the `Soc2AuditTrailResponse` schema.
type Soc2AuditTrailResponse struct {
	Summary          Soc2AuditTrailSummary `json:"summary"`     // required
	AuditTrail       []Soc2AuditTrailEvent `json:"audit_trail"` // required
	ControlObjective *string               `json:"control_objective,omitempty"`
	GeneratedAt      string                `json:"generated_at"`      // required
	GeneratedForOrg  string                `json:"generated_for_org"` // required
}

// Soc2AuditTrailSummary maps the `Soc2AuditTrailSummary` schema.
type Soc2AuditTrailSummary struct {
	TotalEvents       int               `json:"total_events"`        // required
	DateRange         map[string]string `json:"date_range"`          // required
	EventTypesCovered []string          `json:"event_types_covered"` // required
	UniqueUsers       int               `json:"unique_users"`        // required
	IntegrityStatus   string            `json:"integrity_status"`    // required
}

// StatusHistoryEntry maps the `StatusHistoryEntry` schema.
// One transition in a DataSubjectRequest's status_history. Common fields:
// timestamp, from_status / to_status, actor, note. GDPR deletion adds:
// legal_basis. Other variants may add more — `extra=allow`.
type StatusHistoryEntry struct {
	Timestamp  *string `json:"timestamp,omitempty"`
	Status     *string `json:"status,omitempty"`
	FromStatus *string `json:"from_status,omitempty"`
	ToStatus   *string `json:"to_status,omitempty"`
	Actor      *string `json:"actor,omitempty"`
	Note       *string `json:"note,omitempty"`
}

// SupersedeRequest maps the `SupersedeRequest` schema.
// Body for POST /compliance/evidence-packages/{id}/supersede. “reason“ is
// required (it is written into the immutable supersession record). Every
// other field is an optional correction — omitted fields are inherited
// byte-for-byte from the original so a caller who only fixes one thing
// doesn't have to restate the whole package.
type SupersedeRequest struct {
	Reason         string   `json:"reason"` // required
	Name           *string  `json:"name,omitempty"`
	Framework      *string  `json:"framework,omitempty"`
	ControlIDs     []string `json:"control_ids,omitempty"`
	TimeRangeStart *string  `json:"time_range_start,omitempty"`
	TimeRangeEnd   *string  `json:"time_range_end,omitempty"`
	EventIDs       []string `json:"event_ids,omitempty"`
}

// SystemStatus maps the `SystemStatus` schema.
type SystemStatus struct {
	Service         string            `json:"service"`          // required
	Version         string            `json:"version"`          // required
	UptimeSeconds   float64           `json:"uptime_seconds"`   // required
	DatabaseStatus  string            `json:"database_status"`  // required
	CacheStatus     string            `json:"cache_status"`     // required
	QueueStatus     string            `json:"queue_status"`     // required
	BackgroundTasks map[string]string `json:"background_tasks"` // required
	Metrics         map[string]any    `json:"metrics"`          // required
}

// TemplateCloneBody maps the `TemplateCloneBody` schema.
type TemplateCloneBody struct {
	Vars map[string]string `json:"vars,omitempty"`
}

// TimeSeriesPoint maps the `TimeSeriesPoint` schema.
type TimeSeriesPoint struct {
	Timestamp string `json:"timestamp"` // required
	Value     int    `json:"value"`     // required
}

// ToolInventoryItem maps the `ToolInventoryItem` schema.
// One AI tool that agents actually called, with usage rollups.
type ToolInventoryItem struct {
	Name    string `json:"name"`              // The tool identity (resource_id of the tool-call event)
	Calls   *int   `json:"calls,omitempty"`   // Total number of times this tool was called in the window
	Callers *int   `json:"callers,omitempty"` // Distinct callers (user_id) that invoked this tool
	Success *int   `json:"success,omitempty"` // Calls whose outcome was a success
	Failure *int   `json:"failure,omitempty"` // Calls whose outcome was not a success
}

// ToolInventoryResponse maps the `ToolInventoryResponse` schema.
// Per-tool inventory: which AI tools agents called and how often.
type ToolInventoryResponse struct {
	Tools         []ToolInventoryItem `json:"tools,omitempty"`
	TotalCalls    *int                `json:"total_calls,omitempty"`    // Total tool calls across all tools in the window
	DistinctTools *int                `json:"distinct_tools,omitempty"` // Number of distinct tools observed
	OrgID         *string             `json:"org_id,omitempty"`
}

// TopUser maps the `TopUser` schema.
type TopUser struct {
	UserID           string  `json:"user_id"`            // required
	EventCount       int     `json:"event_count"`        // required
	UniqueEventTypes int     `json:"unique_event_types"` // required
	ActiveDays       int     `json:"active_days"`        // required
	FailedEvents     int     `json:"failed_events"`      // required
	SuccessRate      float64 `json:"success_rate"`       // required
}

// TopUsersPeriod maps the `TopUsersPeriod` schema.
type TopUsersPeriod struct {
	Start string `json:"start"` // required
	End   string `json:"end"`   // required
}

// TopUsersResponse maps the `TopUsersResponse` schema.
type TopUsersResponse struct {
	TopUsers []TopUser      `json:"top_users"` // required
	Period   TopUsersPeriod `json:"period"`    // required
}

// TrainingJobItem maps the `TrainingJobItem` schema.
type TrainingJobItem struct {
	JobID         string         `json:"job_id"` // required
	TraceID       *string        `json:"trace_id,omitempty"`
	RunID         *string        `json:"run_id,omitempty"`
	TenantID      string         `json:"tenant_id"` // required
	OrgID         string         `json:"org_id"`    // required
	Status        string         `json:"status"`    // required
	JobType       *string        `json:"job_type,omitempty"`
	RequestedBy   *string        `json:"requested_by,omitempty"`
	ModelFamily   *string        `json:"model_family,omitempty"`
	Payload       map[string]any `json:"payload,omitempty"`
	ResultPayload map[string]any `json:"result_payload,omitempty"`
	Attempts      *int           `json:"attempts,omitempty"`
	LastError     *string        `json:"last_error,omitempty"`
	LeaseToken    *string        `json:"lease_token,omitempty"`
	LeaseUntil    *string        `json:"lease_until,omitempty"`
	StatusVersion *int           `json:"status_version,omitempty"`
	CreatedAt     *string        `json:"created_at,omitempty"`
	UpdatedAt     *string        `json:"updated_at,omitempty"`
}

// ValidationError maps the `ValidationError` schema.
type ValidationError struct {
	Loc  []any  `json:"loc"`  // required
	Msg  string `json:"msg"`  // required
	Type string `json:"type"` // required
}
