# Audit Service API — outline

- version: `latest`  openapi: `3.1.0`
- **249 endpoints** across **226 paths**, **215 schemas**

## Endpoints by designation (25 groups)

### (root)  (1)
- `GET    /`  Root
    - resp:   200:ServiceRootResponse

### admin  (13)
- `POST   /admin/data/purge`  Purge Old Data
    - params: Authorization(header), X-API-Key(header)
    - body:   DataPurgeRequest
    - resp:   200:PurgeDataResponse, 422:HTTPValidationError
- `GET    /admin/deletions`  List Deletion Requests
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /admin/deletions/{deletion_id}`  Get Deletion Request
    - params: deletion_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `POST   /admin/deletions/{deletion_id}/approve`  Approve Deletion Request
    - params: deletion_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   DeletionDecisionRequest
    - resp:   200:any, 422:HTTPValidationError
- `POST   /admin/deletions/{deletion_id}/reject`  Reject Deletion Request
    - params: deletion_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   DeletionDecisionRequest
    - resp:   200:any, 422:HTTPValidationError
- `POST   /admin/deletions/{deletion_id}/restore`  Restore Deletion Request
    - params: deletion_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   DeletionRestoreRequest
    - resp:   200:any, 422:HTTPValidationError
- `POST   /admin/maintenance/optimize`  Optimize Database
    - params: table(query), Authorization(header), X-API-Key(header)
    - resp:   200:MaintenanceOptimizeResponse, 422:HTTPValidationError
- `GET    /admin/metrics/database`  Get Database Metrics
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:DatabaseMetricsResponse, 422:HTTPValidationError
- `GET    /admin/organizations`  Get Organizations Summary
    - params: limit(query), offset(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_OrganizationSummary_, 422:HTTPValidationError
- `POST   /admin/query`  Execute Adhoc Query
    - params: Authorization(header), X-API-Key(header)
    - body:   AdHocQueryRequest
    - resp:   200:AdHocQueryResponse, 422:HTTPValidationError
- `GET    /admin/retention/policies`  Get Retention Policies
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:RetentionPoliciesResponse, 422:HTTPValidationError
- `PUT    /admin/retention/policies`  Put Retention Policies
    - params: Authorization(header), X-API-Key(header)
    - body:   RetentionPolicyUpdate
    - resp:   200:RetentionPoliciesResponse, 422:HTTPValidationError
- `GET    /admin/status`  Get System Status
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:SystemStatus, 422:HTTPValidationError

### analytics  (7)
- `GET    /analytics/compliance`  Get Compliance Report
    - params: org_id(query), start_time*(query), end_time*(query), Authorization(header), X-API-Key(header)
    - resp:   200:ComplianceAnalyticsReport, 422:HTTPValidationError
- `GET    /analytics/flow`  Get Flow Analytics
    - params: org_id(query), start_time*(query), end_time*(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:FlowResponse, 422:HTTPValidationError
- `GET    /analytics/geo`  Get Geo Analytics
    - params: org_id(query), start_time*(query), end_time*(query), Authorization(header), X-API-Key(header)
    - resp:   200:GeoMapResponse, 422:HTTPValidationError
- `GET    /analytics/ingest-health`  Get Ingest Health
    - params: lookback_hours(query), org_id(query), expected_services(query), Authorization(header), X-API-Key(header)
    - resp:   200:IngestHealthResponse, 422:HTTPValidationError
- `GET    /analytics/security`  Get Security Analytics
    - params: org_id(query), start_time*(query), end_time*(query), alert_limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:SecurityMetrics, 422:HTTPValidationError
- `GET    /analytics/users/top`  Get Top Users Analytics
    - params: org_id(query), start_time*(query), end_time*(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:TopUsersResponse, 422:HTTPValidationError
- `GET    /analytics/volume`  Get Event Volume Analytics
    - params: org_id(query), start_time*(query), end_time*(query), Authorization(header), X-API-Key(header)
    - resp:   200:EventVolumeStats, 422:HTTPValidationError

### auditor-share  (8)
- `GET    /auditor/share/{token}`  Auditor Share View
    - params: token*(path)
    - resp:   200:AuditorViewPayloadWithLineage, 422:HTTPValidationError
- `GET    /compliance/evidence-packages`  List Evidence Packages
    - params: framework(query), sealed(query), limit(query), offset(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_EvidencePackageWithLineage_, 422:HTTPValidationError
- `POST   /compliance/evidence-packages`  Create Evidence Package
    - params: Authorization(header), X-API-Key(header)
    - body:   EvidencePackageCreate
    - resp:   201:EvidencePackage, 422:HTTPValidationError
- `GET    /compliance/evidence-packages/{package_id}`  Get Evidence Package
    - params: package_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:EvidencePackageWithLineage, 422:HTTPValidationError
- `POST   /compliance/evidence-packages/{package_id}/share`  Mint Share Link
    - params: package_id*(path), Authorization(header), X-API-Key(header)
    - body:   ShareLinkCreate
    - resp:   201:ShareLink, 422:HTTPValidationError
- `GET    /compliance/evidence-packages/{package_id}/share-links`  List Share Links
    - params: package_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:[ShareLinkSummary], 422:HTTPValidationError
- `POST   /compliance/evidence-packages/{package_id}/supersede`  Supersede Evidence Package
    - params: package_id*(path), Authorization(header), X-API-Key(header)
    - body:   SupersedeRequest
    - resp:   201:EvidencePackageWithLineage, 422:HTTPValidationError
- `DELETE /compliance/share-links/{token}`  Revoke Share Link
    - params: token*(path), Authorization(header), X-API-Key(header)
    - resp:   204, 422:HTTPValidationError

### capabilities  (1)
- `GET    /capabilities`  Get Capabilities
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:CapabilitiesResponse, 422:HTTPValidationError

### compliance  (20)
- `GET    /compliance/control-mappings`  List Control Mappings
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:[ControlMappingRow], 422:HTTPValidationError
- `POST   /compliance/control-mappings`  Create Control Mapping
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   ControlMappingCreate
    - resp:   201:ControlMappingRow, 422:HTTPValidationError
- `DELETE /compliance/control-mappings/{mapping_id}`  Delete Control Mapping
    - params: mapping_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   204, 422:HTTPValidationError
- `GET    /compliance/controls`  Get Compliance Controls
    - params: org_id(query), framework(query), status(query), Authorization(header), X-API-Key(header)
    - resp:   200:ComplianceControlsListResponse, 422:HTTPValidationError
- `GET    /compliance/dashboard`  Get Compliance Dashboard
    - params: org_id(query), framework(query), Authorization(header), X-API-Key(header)
    - resp:   200:ComplianceDashboardResponse, 422:HTTPValidationError
- `POST   /compliance/gdpr/data-subject-access`  Request Data Subject Access
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   DataSubjectAccessRequest
    - resp:   200:DataSubjectAccessAck, 422:HTTPValidationError
- `POST   /compliance/gdpr/right-to-be-forgotten`  Request Data Deletion
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   DataSubjectDeletionRequest
    - resp:   200:DataDeletionAck, 422:HTTPValidationError
- `GET    /compliance/integrity/event/{event_id}`  Verify Integrity Event
    - params: event_id*(path), table(query), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `POST   /compliance/integrity/verify`  Verify Integrity Chain
    - params: Authorization(header), X-API-Key(header)
    - body:   IntegrityVerifyRequest
    - resp:   200:any, 422:HTTPValidationError
- `GET    /compliance/policies`  Get Compliance Policies
    - params: org_id(query), standard(query), Authorization(header), X-API-Key(header)
    - resp:   200:[CompliancePolicy], 422:HTTPValidationError
- `GET    /compliance/reports`  List Compliance Reports
    - params: org_id(query), standard(query), status(query), limit(query), offset(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_ComplianceReportRow_, 422:HTTPValidationError
- `POST   /compliance/reports/generate`  Generate Compliance Report
    - params: org_id(query), standard*(query), period_start*(query), period_end*(query), Authorization(header), X-API-Key(header)
    - resp:   200:ComplianceReportAck, 422:HTTPValidationError
- `GET    /compliance/reports/{report_id}`  Get Compliance Report
    - params: report_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /compliance/reports/{report_id}/deliveries`  List Report Deliveries
    - params: report_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /compliance/reports/{report_id}/download`  Download Compliance Report
    - params: report_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /compliance/requests`  List Compliance Requests
    - params: org_id(query), status(query), request_type(query), deadline_before(query), limit(query), offset(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_DsarRequestRow_, 422:HTTPValidationError
- `GET    /compliance/requests/{request_id}`  Get Compliance Request Status
    - params: request_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ComplianceRequestStatus, 422:HTTPValidationError
- `POST   /compliance/requests/{request_id}/comment`  Add Compliance Request Comment
    - params: request_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   DsarCommentBody
    - resp:   200:DsarCommentResponse, 422:HTTPValidationError
- `PATCH  /compliance/requests/{request_id}/status`  Patch Compliance Request Status
    - params: request_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   DsarStatusPatchBody
    - resp:   200:DsarStatusPatchResponse, 422:HTTPValidationError
- `GET    /compliance/soc2/audit-trail`  Get Soc2 Audit Trail
    - params: org_id(query), start_date*(query), end_date*(query), control_objective(query), Authorization(header), X-API-Key(header)
    - resp:   200:Soc2AuditTrailResponse, 422:HTTPValidationError

### compliance-automation  (45)
- `GET    /compliance-automation/assertions`  Get Assertions
    - params: org_id(query), framework(query), Authorization(header), X-API-Key(header)
    - resp:   200:AssertionsResponse, 422:HTTPValidationError
- `GET    /compliance-automation/callbacks/deliveries`  Get Callback Deliveries
    - params: org_id(query), event_topic(query), subscription_id(query), status(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_CallbackDeliveryItem_, 422:HTTPValidationError
- `GET    /compliance-automation/callbacks/subscriptions`  Get Callback Subscriptions
    - params: org_id(query), only_enabled(query), Authorization(header), X-API-Key(header)
    - resp:   200:[CallbackSubscriptionResponse], 422:HTTPValidationError
- `POST   /compliance-automation/callbacks/subscriptions`  Create Callback Subscription
    - params: Authorization(header), X-API-Key(header)
    - body:   CallbackSubscriptionCreateRequest
    - resp:   200:CallbackSubscriptionResponse, 422:HTTPValidationError
- `POST   /compliance-automation/callbacks/subscriptions/{subscription_id}/pause`  Pause Callback Subscription
    - params: subscription_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:CallbackSubscriptionPauseResponse, 422:HTTPValidationError
- `POST   /compliance-automation/callbacks/subscriptions/{subscription_id}/resume`  Resume Callback Subscription
    - params: subscription_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:CallbackSubscriptionResumeResponse, 422:HTTPValidationError
- `POST   /compliance-automation/callbacks/subscriptions/{subscription_id}/test`  Send Test Callback
    - params: subscription_id*(path), Authorization(header), X-API-Key(header)
    - body:   CallbackTestRequest
    - resp:   200:CallbackSubscriptionTestResponse, 422:HTTPValidationError
- `POST   /compliance-automation/events/evaluate`  Evaluate Event Now
    - params: Authorization(header), X-API-Key(header)
    - body:   EvaluationEventRequest
    - resp:   200:EvaluateEventResponse, 422:HTTPValidationError
- `GET    /compliance-automation/framework-packs`  Get Framework Pack Registry
    - params: include_controls(query), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_FrameworkPackItem_, 422:HTTPValidationError
- `GET    /compliance-automation/framework-packs/{framework}`  Get Framework Pack Details
    - params: framework*(path), Authorization(header), X-API-Key(header)
    - resp:   200:FrameworkPackItem, 422:HTTPValidationError
- `POST   /compliance-automation/framework-packs/{framework}/activate`  Activate Framework Pack
    - params: framework*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:FrameworkActivationResponse, 422:HTTPValidationError
- `POST   /compliance-automation/framework-packs/{framework}/deactivate`  Deactivate Framework Pack
    - params: framework*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:FrameworkActivationResponse, 422:HTTPValidationError
- `GET    /compliance-automation/framework-packs/{framework}/readiness`  Get Framework Pack Readiness
    - params: framework*(path), org_id(query), lookback_days(query), Authorization(header), X-API-Key(header)
    - resp:   200:FrameworkReadinessResponse, 422:HTTPValidationError
- `GET    /compliance-automation/health`  Module Health
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ModuleHealthResponse, 422:HTTPValidationError
- `GET    /compliance-automation/judge/policies`  Get Judge Policies
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_JudgePolicyItem_, 422:HTTPValidationError
- `POST   /compliance-automation/judge/policies`  Create Or Update Judge Policy
    - params: Authorization(header), X-API-Key(header)
    - body:   JudgePolicyUpsertRequest
    - resp:   200:JudgePolicyUpsertResponse, 422:HTTPValidationError
- `GET    /compliance-automation/judge/prompts`  Get Judge Prompts
    - params: org_id(query), prompt_key(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_JudgePromptItem_, 422:HTTPValidationError
- `POST   /compliance-automation/judge/prompts`  Create Or Update Judge Prompt
    - params: Authorization(header), X-API-Key(header)
    - body:   JudgePromptUpsertRequest
    - resp:   200:JudgePromptUpsertResponse, 422:HTTPValidationError
- `POST   /compliance-automation/judge/prompts/{prompt_key}/activate`  Activate Judge Prompt
    - params: prompt_key*(path), version*(query), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:JudgePromptActivateResponse, 422:HTTPValidationError
- `POST   /compliance-automation/judge/prompts/{prompt_key}/rollback`  Rollback Judge Prompt
    - params: prompt_key*(path), to_version*(query), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:JudgePromptActivateResponse, 422:HTTPValidationError
- `GET    /compliance-automation/judge/verdicts`  Get Judge Verdicts
    - params: org_id(query), control_id(query), run_id(query), trace_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_JudgeVerdictItem_, 422:HTTPValidationError
- `POST   /compliance-automation/judge/verdicts/{verdict_id}/replay`  Replay Stored Judge Verdict
    - params: verdict_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:JudgeReplayResponse, 422:HTTPValidationError
- `GET    /compliance-automation/merge-policies`  Get Merge Policies
    - params: org_id(query), control_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_MergePolicyItem_, 422:HTTPValidationError
- `POST   /compliance-automation/merge-policies`  Create Or Update Merge Policy
    - params: Authorization(header), X-API-Key(header)
    - body:   MergePolicyUpsertRequest
    - resp:   200:MergePolicyUpsertResponse, 422:HTTPValidationError
- `GET    /compliance-automation/ml/activations`  Get Ml Activations
    - params: org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ModelActivationItem_, 422:HTTPValidationError
- `GET    /compliance-automation/ml/artifacts`  Get Ml Artifacts
    - params: org_id(query), control_id(query), model_name(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ModelArtifactItem_, 422:HTTPValidationError
- `POST   /compliance-automation/ml/artifacts/{artifact_id}/activate`  Activate Ml Artifact
    - params: artifact_id*(path), Authorization(header), X-API-Key(header)
    - body:   MlModelActivationRequest
    - resp:   200:ModelArtifactActivateResponse, 422:HTTPValidationError
- `GET    /compliance-automation/ml/baselines`  Get Ml Baselines
    - params: org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_MlBaselineItem_, 422:HTTPValidationError
- `GET    /compliance-automation/ml/configs`  Get Ml Configs
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_MlConfigItem_, 422:HTTPValidationError
- `POST   /compliance-automation/ml/configs`  Create Or Update Ml Config
    - params: Authorization(header), X-API-Key(header)
    - body:   MlConfigUpsertRequest
    - resp:   200:MlConfigUpsertResponse, 422:HTTPValidationError
- `POST   /compliance-automation/ml/controls/{control_id}/rollback`  Rollback Ml Control
    - params: control_id*(path), Authorization(header), X-API-Key(header)
    - body:   MlModelRollbackRequest
    - resp:   200:ModelControlRollbackResponse, 422:HTTPValidationError
- `POST   /compliance-automation/ml/distill`  Train Distilled Ml Students
    - params: Authorization(header), X-API-Key(header)
    - body:   MlDistillationTrainRequest
    - resp:   202:MlDistillJobResponse, 422:HTTPValidationError
- `GET    /compliance-automation/ml/jobs`  Get Ml Training Jobs
    - params: org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_TrainingJobItem_, 422:HTTPValidationError
- `GET    /compliance-automation/ml/scores`  Get Ml Scores
    - params: org_id(query), control_id(query), label(query), run_id(query), trace_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ModelScoreItem_, 422:HTTPValidationError
- `GET    /compliance-automation/policies/rules`  Get Policy Rules
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_PolicyRuleItem_, 422:HTTPValidationError
- `POST   /compliance-automation/policies/rules`  Create Or Update Policy Rule
    - params: Authorization(header), X-API-Key(header)
    - body:   PolicyRuleUpsertRequest
    - resp:   200:PolicyRuleUpsertResponse, 422:HTTPValidationError
- `GET    /compliance-automation/quality/false-positive`  Get False Positive Reports
    - params: org_id(query), framework(query), control_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_FalsePositiveFeedbackItem_, 422:HTTPValidationError
- `POST   /compliance-automation/quality/false-positive`  Report False Positive
    - params: Authorization(header), X-API-Key(header)
    - body:   FalsePositiveFeedbackRequest
    - resp:   200:FalsePositiveReportResponse, 422:HTTPValidationError
- `GET    /compliance-automation/risk-cases`  Get Risk Cases
    - params: org_id(query), status(query), risk_level(query), run_id(query), trace_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:RiskCasesResponse, 422:HTTPValidationError
- `GET    /compliance-automation/risk-cases/escalations`  Get Risk Case Escalations
    - params: org_id(query), case_id(query), channel(query), status(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_CaseEscalationItem_, 422:HTTPValidationError
- `GET    /compliance-automation/risk-cases/{case_id}/approvals`  Get Risk Case Approvals
    - params: case_id*(path), org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_CaseApprovalItem_, 422:HTTPValidationError
- `POST   /compliance-automation/risk-cases/{case_id}/approvals`  Decide Risk Case Approval
    - params: case_id*(path), Authorization(header), X-API-Key(header)
    - body:   CaseApprovalDecisionRequest
    - resp:   200:CaseApprovalDecisionResponse, 422:HTTPValidationError
- `GET    /compliance-automation/risk-cases/{case_id}/timeline`  Get Risk Case Timeline
    - params: case_id*(path), org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_CaseTimelineEntry_, 422:HTTPValidationError
- `GET    /compliance-automation/rollouts/preview`  Preview Canary Rollout
    - params: component*(query), run_id*(query), percent*(query), salt(query), Authorization(header), X-API-Key(header)
    - resp:   200:RolloutPreviewResponse, 422:HTTPValidationError
- `GET    /compliance-automation/runs/{run_id}/trace`  Get Run Trace
    - params: run_id*(path), org_id(query), trace_id(query), limit_per_section(query), Authorization(header), X-API-Key(header)
    - resp:   200:RunTraceResponse, 422:HTTPValidationError

### compliance-copilot  (3)
- `POST   /compliance/copilot/ask`  Copilot Ask
    - params: Authorization(header), X-API-Key(header)
    - body:   CopilotAsk
    - resp:   200:CopilotAnswer, 422:HTTPValidationError
- `GET    /compliance/copilot/conversations`  List Copilot Conversations
    - params: org_id(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /compliance/copilot/conversations/{conversation_id}`  Get Copilot Conversation
    - params: conversation_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError

### compliance-extras  (8)
- `GET    /compliance/auditor-requests`  List Auditor Requests
    - params: org_id(query), state(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:[AuditorRequest], 422:HTTPValidationError
- `POST   /compliance/auditor-requests`  Create Auditor Request
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   AuditorRequestCreate
    - resp:   201:AuditorRequest, 422:HTTPValidationError
- `PATCH  /compliance/auditor-requests/{request_id}`  Patch Auditor Request
    - params: request_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   AuditorRequestUpdate
    - resp:   200:AuditorRequest, 422:HTTPValidationError
- `POST   /compliance/auditor-requests/{request_id}/satisfy`  Satisfy Auditor Request
    - params: request_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   AuditorRequestSatisfy
    - resp:   200:AuditorRequest, 422:HTTPValidationError
- `GET    /compliance/posture-history`  Get Posture History
    - params: window_days(query), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:PostureHistoryResponse, 422:HTTPValidationError
- `DELETE /settings/audit-period`  Delete Audit Period
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   204, 422:HTTPValidationError
- `GET    /settings/audit-period`  Get Audit Period
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:AuditPeriod, 422:HTTPValidationError
- `PUT    /settings/audit-period`  Put Audit Period
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   AuditPeriodWrite
    - resp:   200:AuditPeriod, 422:HTTPValidationError

### compliance-templates  (2)
- `GET    /compliance/templates/`  List Templates
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ComplianceTemplateItem_, 422:HTTPValidationError
- `POST   /compliance/templates/{template_id}/clone`  Clone Template
    - params: template_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   TemplateCloneBody
    - resp:   200:ComplianceTemplateCloneResponse, 422:HTTPValidationError

### compliance-ui  (1)
- `GET    /ui/compliance/dashboard`  Compliance Dashboard Ui
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError

### explorer-queries  (5)
- `GET    /explorer-queries/history`  List Query History
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ExplorerHistoryItem_, 422:HTTPValidationError
- `POST   /explorer-queries/history`  Record Query History
    - params: Authorization(header), X-API-Key(header)
    - body:   ExplorerHistoryCreate
    - resp:   201:ExplorerHistoryItem, 422:HTTPValidationError
- `GET    /explorer-queries/saved`  List Saved Queries
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_ExplorerSavedQueryItem_, 422:HTTPValidationError
- `POST   /explorer-queries/saved`  Create Saved Query
    - params: Authorization(header), X-API-Key(header)
    - body:   ExplorerSavedQueryCreate
    - resp:   201:ExplorerSavedQueryItem, 422:HTTPValidationError
- `DELETE /explorer-queries/saved/{query_id}`  Delete Saved Query
    - params: query_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:ExplorerSavedQueryDeleteResponse, 422:HTTPValidationError

### export  (6)
- `POST   /export/`  Create Export Job
    - params: Authorization(header), X-API-Key(header)
    - body:   ExportRequestInput
    - resp:   200:ExportCreateResponse, 422:HTTPValidationError
- `GET    /export/formats`  Get Export Formats
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ExportFormatsResponse, 422:HTTPValidationError
- `GET    /export/jobs`  Get Export Jobs
    - params: org_id(query), user_id(query), status(query), limit(query), offset(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_ExportJobResponse_, 422:HTTPValidationError
- `DELETE /export/jobs/{export_id}`  Cancel Export Job
    - params: export_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ExportCancelResponse, 422:HTTPValidationError
- `GET    /export/jobs/{export_id}`  Get Export Job
    - params: export_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ExportJobResponse, 422:HTTPValidationError
- `GET    /export/jobs/{export_id}/download`  Download Export Job
    - params: export_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError

### health  (4)
- `GET    /health`  Health Check
    - resp:   200:HealthResponse
- `GET    /health/detailed`  Detailed Health Check
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:DetailedHealthResponse, 422:HTTPValidationError
- `GET    /liveness`  Liveness Check
    - resp:   200:LivenessResponse
- `GET    /readiness`  Readiness Check
    - resp:   200:ReadinessResponse

### legal-hold  (3)
- `GET    /legal-holds`  List Legal Holds
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `POST   /legal-holds`  Place Legal Hold
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   LegalHoldCreate
    - resp:   200:any, 422:HTTPValidationError
- `POST   /legal-holds/{hold_id}/release`  Release Legal Hold
    - params: hold_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - body:   LegalHoldRelease
    - resp:   200:any, 422:HTTPValidationError

### llm.txt  (1)
- `GET    /llm.txt`  Llm Txt
    - resp:   200

### logs  (6)
- `GET    /logs/events`  Get Audit Events
    - params: org_id(query), start_time(query), end_time(query), event_type(query), user_id(query), action(query), outcome(query), limit(query), offset(query), cursor(query), Authorization(header), X-API-Key(header)
    - resp:   200:PaginatedResponse_AuditEvent_, 422:HTTPValidationError
- `GET    /logs/events/{event_id}`  Get Audit Event
    - params: event_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:AuditEvent, 422:HTTPValidationError
- `POST   /logs/ingest`  Ingest Audit Event
    - params: Authorization(header), X-API-Key(header)
    - body:   AuditEventCreate
    - resp:   202:IngestLogResponse, 422:HTTPValidationError
- `POST   /logs/ingest/batch`  Ingest Audit Events Batch
    - params: Authorization(header), X-API-Key(header)
    - body:   [AuditEventCreate]
    - resp:   202:IngestLogBatchResponse, 422:HTTPValidationError
- `POST   /logs/ingest/stream`  Ingest Stream Endpoint
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:IngestStreamResponse, 422:HTTPValidationError
- `GET    /logs/stats/events`  Get Audit Events Count
    - params: org_id(query), start_time(query), end_time(query), event_type(query), Authorization(header), X-API-Key(header)
    - resp:   200:EventStatsResponse, 422:HTTPValidationError

### metrics  (1)
- `GET    /metrics`  Metrics
    - params: Authorization(header), X-API-Key(header)
    - resp:   200, 422:HTTPValidationError

### notifications  (3)
- `GET    /notifications`  Get Notifications
    - params: org_id(query), lookback_days(query), deadline_horizon_days(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:NotificationFeed, 422:HTTPValidationError
- `GET    /notifications/`  Get Notifications
    - params: org_id(query), lookback_days(query), deadline_horizon_days(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:NotificationFeed, 422:HTTPValidationError
- `POST   /notifications/read`  Mark Notifications Read
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:NotificationReadResponse, 422:HTTPValidationError

### report-schedules  (5)
- `GET    /compliance/report-schedules/`  List Report Schedules
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ReportScheduleListResponse, 422:HTTPValidationError
- `POST   /compliance/report-schedules/`  Create Report Schedule
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   ReportScheduleCreate
    - resp:   201:ReportScheduleItem, 422:HTTPValidationError
- `DELETE /compliance/report-schedules/{schedule_id}`  Delete Report Schedule
    - params: schedule_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ReportScheduleDeleteResponse, 422:HTTPValidationError
- `POST   /compliance/report-schedules/{schedule_id}/pause`  Pause Report Schedule
    - params: schedule_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ReportSchedulePauseResponse, 422:HTTPValidationError
- `POST   /compliance/report-schedules/{schedule_id}/resume`  Resume Report Schedule
    - params: schedule_id*(path), org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ReportSchedulePauseResponse, 422:HTTPValidationError

### runtime-events  (3)
- `POST   /events/batch`  Ingest Runtime Events Batch
    - params: Authorization(header), X-API-Key(header)
    - resp:   202:IngestBatchResponse, 422:HTTPValidationError
- `POST   /events/ingest`  Ingest Runtime Event
    - params: Authorization(header), X-API-Key(header)
    - resp:   202:IngestEventResponse, 422:HTTPValidationError
- `GET    /events/stream`  Stream Audit Events
    - params: org_id(query), event_type_prefix(query), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError

### saved-searches  (3)
- `GET    /saved-searches/`  List Saved Searches
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:ListEnvelope_SavedSearchItem_, 422:HTTPValidationError
- `POST   /saved-searches/`  Create Saved Search
    - params: Authorization(header), X-API-Key(header)
    - body:   SavedSearchCreate
    - resp:   201:SavedSearchItem, 422:HTTPValidationError
- `DELETE /saved-searches/{search_id}`  Delete Saved Search
    - params: search_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:SavedSearchDeleteResponse, 422:HTTPValidationError

### search  (4)
- `POST   /search/`  Search Audit Events
    - params: Authorization(header), X-API-Key(header)
    - body:   SearchQuery
    - resp:   200:SearchResult, 422:HTTPValidationError
- `GET    /search/facets`  Get Search Facets
    - params: org_id(query), start_time(query), end_time(query), Authorization(header), X-API-Key(header)
    - resp:   200:SearchFacetsResponse, 422:HTTPValidationError
- `GET    /search/suggest`  Get Search Suggestions
    - params: org_id(query), field*(query), prefix(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:SearchSuggestResponse, 422:HTTPValidationError
- `GET    /search/tools`  Get Tool Inventory
    - params: org_id(query), start_time(query), end_time(query), limit(query), Authorization(header), X-API-Key(header)
    - resp:   200:ToolInventoryResponse, 422:HTTPValidationError

### settings  (6)
- `GET    /settings/classification`  Get Classification Settings
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:ClassificationSettings, 422:HTTPValidationError
- `PUT    /settings/classification`  Put Classification Settings
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   ClassificationSettings
    - resp:   200:ClassificationSettings, 422:HTTPValidationError
- `GET    /settings/notifications`  Get Notification Prefs
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:NotificationPrefsSettings, 422:HTTPValidationError
- `PUT    /settings/notifications`  Put Notification Prefs
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   NotificationPrefsSettings
    - resp:   200:NotificationPrefsSettings, 422:HTTPValidationError
- `GET    /settings/redaction`  Get Redaction Settings
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:RedactionSettings, 422:HTTPValidationError
- `PUT    /settings/redaction`  Put Redaction Settings
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - body:   RedactionSettings
    - resp:   200:RedactionSettings, 422:HTTPValidationError

### ui  (90)
- `GET    /ui/ai-agents`  Ai Agents Page
    - resp:   200:string
- `GET    /ui/ai-agents/playground`  Ai Playground Page
    - resp:   200:string
- `GET    /ui/app/account`  App Account Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/admin`  App Admin Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/agents`  App Agents Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/agents/{agent_id}`  App Agent Detail Page
    - params: agent_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/alerts`  App Alerts Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/api-keys`  App Api Keys Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/approvals`  App Approvals Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/audit-logs`  App Audit Logs Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/behavioral-baselines`  App Behavioral Baselines Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/callbacks`  App Callbacks Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compare`  App Compare Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance`  App Compliance Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/audit-period`  App Audit Period Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/auditor-inbox`  App Auditor Inbox Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/calendar`  App Compliance Calendar Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/control-mappings`  App Compliance Control Mappings Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/controls/{control_id}`  App Compliance Control Detail Page
    - params: control_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/copilot`  App Compliance Copilot Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/dsar`  App Compliance Dsar Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/dsar/{request_id}`  App Compliance Dsar Detail Page
    - params: request_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/evidence`  App Compliance Evidence Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/evidence/new`  App Compliance Evidence Builder Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/frameworks`  App Compliance Frameworks Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/posture-history`  App Posture History Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/readiness`  App Compliance Readiness Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/soc2-trail`  App Compliance Soc2 Trail Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/compliance/templates`  App Compliance Templates Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/dashboard`  App Dashboard Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/data-explorer`  App Data Explorer Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/exports`  App Exports Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/exports/new`  App Export Builder Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/exports/{export_id}`  App Export Detail Page
    - params: export_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/flow`  App Flow Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/help/api`  App Help Api Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/help/changelog`  App Help Changelog Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/help/glossary`  App Help Glossary Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/integrations`  App Integrations Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/intents`  App Intents Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/judge-config`  App Judge Config Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/judge-verdicts`  App Judge Verdicts Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/legal-holds`  App Legal Holds Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/live`  App Live Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/meta-audit`  App Meta Audit Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/ml/training`  App Ml Training Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/onboarding`  App Onboarding Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/plan`  App Plan Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/policy-engine`  App Policy Engine Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/reasoning-chains`  App Reasoning Chains Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/reasoning-chains/{run_id}`  App Reasoning Chain Detail Page
    - params: run_id*(path), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/reports`  App Reports Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/risk-assessment`  App Risk Assessment Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/settings`  App Settings Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/settings/classification`  App Settings Classification Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/settings/notifications`  App Settings Notifications Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/settings/redaction`  App Settings Redaction Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/settings/retention`  App Settings Retention Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/status`  App Status Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/team`  App Team Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/tool-calls`  App Tool Calls Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/tools`  App Tools Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/app/webhooks`  App Webhooks Page
    - params: Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/auditor/{share_token}`  Auditor View Page
    - params: share_token*(path)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/automouts_agents`  Automouts Agents Page Redirect
    - resp:   200:string
- `GET    /ui/autonomous_agents`  Autonomous Agents Page
    - resp:   200:string
- `GET    /ui/blog/control-plane-agent-audit-system`  Control Plane Agent Audit Blog Page
    - resp:   200:string
- `GET    /ui/compliance`  Compliance Page
    - resp:   200:string
- `GET    /ui/compliance/overview`  Compliance Overview Page
    - params: org_id(query), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/demo/app-data/{page_slug}`  App Demo Page Payload
    - params: page_slug*(path), Authorization(header), X-API-Key(header)
    - resp:   200:any, 422:HTTPValidationError
- `GET    /ui/demo/app/compliance`  App Compliance Demo Page
    - params: scenario(query), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/demo/app/{page_slug}`  App Generic Demo Page
    - params: page_slug*(path), scenario(query), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/demo/compliance/dashboard`  Compliance Dashboard Demo Page
    - params: org_id(query), scenario(query), Authorization(header), X-API-Key(header)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/demo/exit`  Exit Demo Mode
    - resp:   200:string
- `GET    /ui/developers`  Developers Page
    - resp:   200:string
- `GET    /ui/directory`  Page Directory Page
    - resp:   200:string
- `GET    /ui/landing`  Landing Page
    - resp:   200:string
- `GET    /ui/login`  Login Redirect
    - params: next(query)
    - resp:   200:string, 422:HTTPValidationError
- `GET    /ui/logout`  Logout Redirect
    - resp:   200:string
- `GET    /ui/pages`  Pages Index Page
    - resp:   200:string
- `GET    /ui/pricing`  Pricing Page
    - resp:   200:string
- `GET    /ui/tool-calling`  Tool Calling Page
    - resp:   200:string
- `GET    /ui/v2/ai-agents`  Ai Agents V2 Page
    - resp:   200:string
- `GET    /ui/v2/ai-agents/playground`  Ai Playground V2 Page
    - resp:   200:string
- `GET    /ui/v2/autonomous-workers`  Autonomous Workers V2 Page
    - resp:   200:string
- `GET    /ui/v2/compliance`  Compliance V2 Page
    - resp:   200:string
- `GET    /ui/v2/integrations`  Integrations V2 Page
    - resp:   200:string
- `GET    /ui/v2/landing`  Landing V2 Page
    - resp:   200:string
- `GET    /ui/v2/sdk`  Sdk V2 Page
    - resp:   200:string
- `GET    /ui/v2/tool-calling`  Tool Calling V2 Page
    - resp:   200:string

## Schemas (215)

_Field detail shown for every schema. `*` = required._

### AdHocQueryRequest  (5 fields)
    * sql: string
      description: string | null
      max_rows: integer
      org_id: string | null
      timeout_seconds: integer

### AdHocQueryResponse  (8 fields)
    * columns: [string]
    * execution_time_seconds: number
    * query_id: string
    * rows: [object]
    * rows_returned: integer
    * status: string
    * timestamp: string
    * truncated: boolean

### AssertionsResponse  (3 fields)
    * total: integer
      items: [object]
      summary: object

### AuditEvent  (18 fields)
    * action: string
    * event_id: string
    * event_type: string
    * service: string
    * timestamp: string<date-time>
      compliance_controls: [string]
      ingested_at: string<date-time> | null
      ip_address: string | null
      metadata: object | null
      org_id: string | null
      outcome: string
      request_data: object | null
      resource_id: string | null
      resource_type: string | null
      response_data: object | null
      session_id: string | null
      user_agent: string | null
      user_id: string | null

### AuditEventCreate  (14 fields)
    * action: string
    * event_type: string
    * service: string
      ip_address: string | null
      metadata: object | null
      org_id: string | null
      outcome: string
      request_data: object | null
      resource_id: string | null
      resource_type: string | null
      response_data: object | null
      session_id: string | null
      user_agent: string | null
      user_id: string | null

### AuditPeriod  (8 fields)
    * end_date: string<date-time>
    * name: string
    * start_date: string<date-time>
      auditor: string | null
      framework: string
      org_id: string | null
      set_at: string<date-time> | null
      set_by: string | null

### AuditPeriodWrite  (5 fields)
    * end_date: string<date-time>
    * name: string
    * start_date: string<date-time>
      auditor: string | null
      framework: string

### AuditorRequest  (13 fields)
    * request_id: string
    * subject: string
      control_id: string | null
      created_at: string<date-time> | null
      created_by: string | null
      due_at: string<date-time> | null
      evidence_package_id: string | null
      framework: string | null
      notes: string | null
      org_id: string | null
      requestor: string | null
      state: AuditorRequestState
      updated_at: string<date-time> | null

### AuditorRequestCreate  (6 fields)
    * subject: string
      control_id: string | null
      due_at: string<date-time> | null
      framework: string | null
      notes: string | null
      requestor: string | null

### AuditorRequestSatisfy  (2 fields)
    * evidence_package_id: string
      notes: string | null

### AuditorRequestState  (string(new, in_progress, awaiting_auditor, satisfied, declined))
    enum: new, in_progress, awaiting_auditor, satisfied, declined

### AuditorRequestUpdate  (8 fields)
      control_id: string | null
      due_at: string<date-time> | null
      evidence_package_id: string | null
      framework: string | null
      notes: string | null
      requestor: string | null
      state: AuditorRequestState | null
      subject: string | null

### AuditorViewPayloadWithLineage  (20 fields)
    * auditor_email: string
    * control_ids: [string]
    * expires_at: string<date-time>
    * framework: string
    * package_name: string
    * watermark_text: string
      custody_created_at: string<date-time> | null
      custody_sealed_by: string | null
      custody_shared_by: string | null
      digest_algorithm: string
      event_count: integer
      events: [object]
      outcomes_breakdown: map<string,integer>
      package_digest: string | null
      sealed_at: string<date-time> | null
      superseded_at: string<date-time> | null
      superseded_by: string | null
      superseded_reason: string | null
      time_range_end: string<date-time> | null
      time_range_start: string<date-time> | null

### CallbackDeliveryItem  (16 fields)
    * callback_url: string
    * delivery_id: string
    * event_topic: string
    * org_id: string
    * status: string
    * subscription_id: string
    * tenant_id: string
      attempt_count: integer
      attempted_at: string<date-time> | null
      error_message: string | null
      http_status: integer | null
      idempotency_key: string | null
      job_id: string | null
      latency_ms: integer | null
      lifecycle_event_id: string | null
      payload_hash: string | null

### CallbackSubscriptionCreateRequest  (5 fields)
    * callback_url: string<uri>
    * event_topic: string
      enabled: boolean
      org_id: string | null
      signing_key_ref: string

### CallbackSubscriptionPauseResponse  (2 fields)
    * subscription_id: string
      status: string

### CallbackSubscriptionResponse  (7 fields)
    * callback_url: string
    * enabled: boolean
    * event_topic: string
    * org_id: string
    * signing_key_ref: string
    * subscription_id: string
    * tenant_id: string

### CallbackSubscriptionResumeResponse  (2 fields)
    * subscription_id: string
      status: string

### CallbackSubscriptionTestResponse  (4 fields)
    * lifecycle_event_id: string
    * subscription_id: string
      dispatch: object
      status: string

### CallbackTestRequest  (2 fields)
      org_id: string | null
      payload: object

### CapabilitiesResponse  (8 fields)
    * callbacks: boolean
    * copilot_mode: string(llm, templates)
    * demo_scenarios: boolean
    * email_delivery: boolean
    * llm_judge: boolean
    * ml_scoring: boolean
    * scheduled_reports: boolean
    * template_names: [string]

### CaseApprovalDecisionRequest  (4 fields)
    * decision: string
    * reason: string
      metadata: object
      org_id: string | null

### CaseApprovalDecisionResponse  (5 fields)
    * approval_id: string
    * case_id: string
    * created_at: string<date-time>
    * decision: string
      status: string

### CaseApprovalItem  (8 fields)
    * approval_id: string
    * approver_user_id: string
    * case_id: string
    * decision: string
    * org_id: string
      created_at: string<date-time> | null
      metadata: object
      reason: string | null

### CaseEscalationItem  (10 fields)
    * case_id: string
    * channel: string
    * escalation_id: string
    * org_id: string
    * risk_level: string
    * status: string
    * tenant_id: string
      created_at: string<date-time> | null
      payload: object | null
      target: string | null

### CaseTimelineEntry  (12 fields)
    * case_id: string
    * event_id: string
    * event_type: string
    * org_id: string
    * risk_level: string
    * status: string
    * tenant_id: string
      actor_id: string | null
      created_at: string<date-time> | null
      details: object
      run_id: string | null
      trace_id: string | null

### ClassificationRule  (2 fields)
    * classification: string
    * field_pattern: string

### ClassificationSettings  (4 fields)
    * rules: [ClassificationRule]
      is_default: boolean
      updated_at: string<date-time> | null
      updated_by: string

### ComplianceAnalyticsReport  (9 fields)
    * access_patterns: map<string,integer>
    * data_subjects_count: integer
    * export_activities: [object]
    * org_id: string
    * period_end: string<date-time>
    * period_start: string<date-time>
    * retention_compliance: object
    * total_events: integer
      framework_scores: map<string,number>

### ComplianceControlSummary  (12 fields)
    * anomaly_hits: integer
    * control_id: string
    * event_count: integer
    * failed_count: integer
    * framework: string
    * max_zscore: number
    * pass_rate: number
    * reason: string
    * relevant_event_types: [string]
    * sample_event_ids: [string]
    * status: string
      last_checked: string<date-time> | null

### ComplianceControlsListResponse  (3 fields)
    * controls: [ComplianceControlSummary]
    * org_id: string
    * total_controls: integer

### ComplianceDashboardResponse  (18 fields)
    * audit_readiness: string
    * compliance_score: integer
    * controls_failing: integer
    * controls_not_tested: integer
    * controls_passing: integer
    * controls_warning: integer
    * generated_at: string<date-time>
    * org_id: string
    * total_controls: integer
      anomaly_signals_last_30_days: integer
      drift_alerts: [object]
      events_last_30_days: integer
      failed_events_last_30_days: integer
      framework_scores: map<string,number>
      insights: [string]
      last_test_date: string<date-time> | null
      next_test_date: string<date-time> | null
      open_findings: [object]

### CompliancePolicy  (30 fields)
    * created_at: string<date-time>
    * created_by: string
    * default_retention: RetentionPeriod
    * name: string
    * org_id: string
    * policy_id: string
    * standards: [ComplianceStandard]
    * updated_at: string<date-time>
      allowed_export_formats: [string]
      audit_access_logging: boolean
      auto_classify_data: boolean
      classification_rules: map<string,DataClassification>
      description: string | null
      event_type_retention: map<string,RetentionPeriod>
      export_approval_roles: [string]
      gdpr_data_controller: string | null
      gdpr_data_processor: string | null
      gdpr_enabled: boolean
      gdpr_legal_basis: GDPRLegalBasis | null
      gdpr_privacy_policy_url: string | null
      ip_whitelist: [string]
      is_active: boolean
      max_export_size_gb: number
      require_approval_for_exports: boolean
      require_mfa_for_access: boolean
      soc2_control_activities: [string]
      soc2_enabled: boolean
      soc2_report_period_months: integer
      soc2_service_organization: string | null
      tamper_detection: boolean

### ComplianceReportAck  (5 fields)
    * estimated_completion: string<date-time>
    * message: string
    * report_id: string
    * status: string
      demo: boolean | null

### ComplianceReportRow  (12 fields)
    * org_id: string
    * report_id: string
    * standard: string
    * status: string
      completed_at: string<date-time> | null
      created_at: string<date-time> | null
      error_message: string | null
      period_end: string<date-time> | null
      period_start: string<date-time> | null
      report_file_url: string | null
      report_type: string | null
      requested_by: string | null

### ComplianceRequestStatus  (8 fields)
    * data_subject_id: string
    * request_id: string
    * request_type: string
    * status: string
      completion_deadline: string<date-time> | null
      created_at: string<date-time> | null
      processing_notes: string | null
      records_affected: integer | null

### ComplianceStandard  (string(soc2, gdpr, hipaa, pci_dss, ccpa, pipeda …))
    enum: soc2, gdpr, hipaa, pci_dss, ccpa, pipeda, custom

### ComplianceTemplateCloneResponse  (8 fields)
    * framework: string
    * org_id: string
    * rendered_at: string<date-time>
    * rendered_markdown: string
    * template_id: string
    * title: string
      placeholders_missing: [string]
      placeholders_used: map<string,string>

### ComplianceTemplateItem  (7 fields)
    * description: string
    * document_type: string
    * framework: string
    * template_id: string
    * title: string
      citation: string | null
      placeholders: [string]

### ControlMappingCreate  (2 fields)
    * controls: [string]
    * event_type: string

### ControlMappingRow  (7 fields)
    * controls: [string]
    * event_type: string
    * id: string
    * org_id: string
      created_at: string<date-time> | null
      created_by: string
      updated_at: string<date-time> | null

### CopilotAnswer  (11 fields)
    * answer: string
    * columns: [string]
    * conversation_id: string
    * mode: string
    * rows: [[any]]
    * rows_returned: integer
    * safety: object
    * sql: string
      fallback_reason: string | null
      template_name: string | null
      used_fallback: boolean

### CopilotAsk  (3 fields)
    * question: string
      conversation_id: string | null
      org_id: string | null

### DataClassification  (string(public, internal, confidential, restricted, pii, phi …))
    enum: public, internal, confidential, restricted, pii, phi, pci

### DataDeletionAck  (5 fields)
    * completion_deadline: string
    * message: string
    * request_id: string
    * status: string
    * warning: string

### DataPurgeRequest  (4 fields)
    * before_date: string<date-time>
    * reason: string
      dry_run: boolean
      org_id: string | null

### DataSubjectAccessAck  (4 fields)
    * completion_deadline: string
    * message: string
    * request_id: string
    * status: string

### DataSubjectAccessRequest  (6 fields)
    * data_subject_id: string
      date_range_end: string<date-time> | null
      date_range_start: string<date-time> | null
      format: string
      reason: string | null
      scope: [string] | null

### DataSubjectDeletionRequest  (7 fields)
    * data_subject_id: string
    * legal_basis: string
    * reason: string
      date_range_end: string<date-time> | null
      date_range_start: string<date-time> | null
      retention_override: boolean
      scope: [string] | null

### DatabaseMetricsResponse  (3 fields)
    * performance_metrics: [DatabasePerformanceMetric]
    * table_metrics: [DatabaseTableMetric]
    * timestamp: string<date-time>

### DatabasePerformanceMetric  (5 fields)
    * avg_duration_ms: number
    * avg_rows_read: number
    * max_duration_ms: number
    * query_count: integer
    * query_kind: string

### DatabaseTableMetric  (6 fields)
    * compressed_size: string
    * parts_count: integer
    * size_on_disk: string
    * table: string
    * total_rows: integer
    * uncompressed_size: string

### DeletionDecisionRequest  (1 fields)
      reason: string | null

### DeletionRestoreRequest  (1 fields)
      reason: string | null

### DetailedHealthResponse  (4 fields)
    * checks: map<string,HealthCheckEntry>
    * service: string
    * status: string(healthy, unhealthy)
    * timestamp: string

### DsarCommentBody  (1 fields)
    * comment: string

### DsarCommentResponse  (3 fields)
    * note_added: boolean
    * request_id: string
    * updated_at: string<date-time>

### DsarRequestRow  (17 fields)
    * data_subject_id: string
    * org_id: string
    * request_id: string
    * request_type: string
    * requested_by: string
    * status: string
      assigned_to: string | null
      completed_at: string<date-time> | null
      completion_deadline: string<date-time> | null
      created_at: string<date-time> | null
      processing_notes: string | null
      reason: string | null
      records_affected: integer | null
      request_date: string<date-time> | null
      scope: [string]
      status_history: [StatusHistoryEntry]
      updated_at: string<date-time> | null

### DsarStatusPatchBody  (2 fields)
    * status: string
      comment: string | null

### DsarStatusPatchResponse  (4 fields)
    * history_entries: integer
    * request_id: string
    * status: string
    * updated_at: string<date-time>

### EvaluateEventResponse  (2 fields)
      status: string
      summary: object

### EvaluationEventRequest  (3 fields)
    * event_type: string
      org_id: string | null
      payload: object

### EventStatsResponse  (2 fields)
    * total: integer
      org_id: string | null

### EventVolumeStats  (4 fields)
    * events_by_outcome: map<string,integer>
    * events_by_type: map<string,integer>
    * events_per_hour: [TimeSeriesPoint]
    * total_events: integer

### EvidencePackage  (13 fields)
    * framework: string
    * name: string
    * package_id: string
      control_ids: [string]
      created_at: string<date-time> | null
      created_by: string | null
      event_ids: [string]
      org_id: string | null
      sealed: boolean
      sealed_at: string<date-time> | null
      time_range_end: string<date-time> | null
      time_range_start: string<date-time> | null
      updated_at: string<date-time> | null

### EvidencePackageCreate  (7 fields)
    * framework: string
    * name: string
      control_ids: [string]
      event_ids: [string]
      seal_on_create: boolean
      time_range_end: string<date-time> | null
      time_range_start: string<date-time> | null

### EvidencePackageWithLineage  (18 fields)
    * framework: string
    * name: string
    * package_id: string
      control_ids: [string]
      created_at: string<date-time> | null
      created_by: string | null
      event_ids: [string]
      org_id: string | null
      sealed: boolean
      sealed_at: string<date-time> | null
      supersede_reason: string | null
      superseded_at: string<date-time> | null
      superseded_by: string | null
      superseded_reason: string | null
      supersedes_package_id: string | null
      time_range_end: string<date-time> | null
      time_range_start: string<date-time> | null
      updated_at: string<date-time> | null

### ExplorerHistoryCreate  (2 fields)
    * question: string
      result_count: integer

### ExplorerHistoryItem  (3 fields)
    * question: string
      ran_at: string<date-time> | null
      result_count: integer | null

### ExplorerSavedQueryCreate  (2 fields)
    * question: string
      name: string

### ExplorerSavedQueryDeleteResponse  (2 fields)
    * deleted: boolean
    * query_id: string

### ExplorerSavedQueryItem  (7 fields)
    * name: string
    * org_id: string
    * query_id: string
    * question: string
    * user_id: string
      created_at: string<date-time> | null
      updated_at: string<date-time> | null

### ExportCancelResponse  (1 fields)
    * message: string

### ExportCreateResponse  (3 fields)
    * export_id: string
    * message: string
    * status: string

### ExportFormatDescriptor  (5 fields)
    * description: string
    * max_size_gb: integer
    * name: string
    * supports_compression: boolean
    * typical_use_case: string

### ExportFormatsResponse  (1 fields)
    * formats: map<string,ExportFormatDescriptor>

### ExportJobResponse  (19 fields)
    * created_at: string<date-time>
    * export_id: string
    * filters: string
    * format: string
    * name: string
    * org_id: string
    * status: string
    * user_id: string
      completed_at: string<date-time> | null
      delivery_method: string | null
      download_url: string | null
      error_message: string | null
      expires_at: string<date-time> | null
      fields: string | null
      file_size_bytes: integer | null
      include_request_data: integer | null
      include_response_data: integer | null
      records_exported: integer | null
      started_at: string<date-time> | null

### ExportRequestInput  (13 fields)
    * format: string
    * name: string
      actions: [string] | null
      delivery_method: string
      delivery_target: string | null
      end_time: string<date-time> | null
      event_types: [string] | null
      filters: object | null
      org_id: string | null
      outcomes: [string] | null
      requested_by_email: string | null
      start_time: string<date-time> | null
      user_ids: [string] | null

### FacetCount  (2 fields)
    * count: integer
    * value: string

### FalsePositiveFeedbackItem  (11 fields)
    * control_id: string
    * feedback_id: string
    * framework: string
    * org_id: string
    * reason: string
    * source: string
    * tenant_id: string
      created_at: string<date-time> | null
      reported_by: string | null
      run_id: string | null
      trace_id: string | null

### FalsePositiveFeedbackRequest  (7 fields)
    * control_id: string
    * framework: string
    * reason: string
      org_id: string | null
      run_id: string | null
      source: string
      trace_id: string | null

### FalsePositiveReportResponse  (4 fields)
    * control_id: string
    * feedback_id: string
    * framework: string
      status: string

### FlowEdge  (4 fields)
    * agent: string
    * count: integer
    * outcome: string
    * tool: string

### FlowResponse  (4 fields)
    * distinct_agents: integer
    * distinct_tools: integer
    * flows: [FlowEdge]
    * period: TopUsersPeriod

### FrameworkActivationResponse  (4 fields)
    * activated_at: string<date-time>
    * activated_by: string
    * active: boolean
    * framework: string

### FrameworkControlItem  (6 fields)
    * control_id: string
    * objective: string
    * risk_weight: integer
      documentation_refs: [string]
      evidence_templates: [string]
      expected_signal_sources: [string]

### FrameworkPackItem  (10 fields)
    * control_count: integer
    * display_name: string
    * framework: string
    * version: string
      activated_at: string<date-time> | null
      activated_by: string | null
      active: boolean | null
      controls: [FrameworkControlItem] | null
      description: string | null
      documentation_refs: [string]

### FrameworkReadinessResponse  (11 fields)
    * controls_total: integer
    * coverage_ratio: number
    * display_name: string
    * framework: string
    * readiness_score: number
    * version: string
      gaps: [object]
      items: [object]
      lookback_days: integer | null
      org_id: string | null
      status_counts: map<string,integer>

### GDPRLegalBasis  (string(consent, contract, legal_obligation, vital_interests, public_task, legitimate_interests))
    enum: consent, contract, legal_obligation, vital_interests, public_task, legitimate_interests

### GeoCountryRow  (4 fields)
    * country: string
    * distinct_ips: integer
    * event_count: integer
      country_code: string | null

### GeoMapResponse  (4 fields)
    * countries: [GeoCountryRow]
    * period: TopUsersPeriod
    * total_events: integer
    * unmapped_count: integer

### HTTPValidationError  (1 fields)
      detail: [ValidationError]

### HealthCheckEntry  (2 fields)
    * message: string
    * status: string(healthy, unhealthy, warning)

### HealthResponse  (3 fields)
    * service: string
    * status: string(healthy, unhealthy)
    * timestamp: string

### IngestBatchResponse  (3 fields)
    * count: integer
    * event_ids: [string]
    * status: string

### IngestEventResponse  (2 fields)
    * event_id: string
    * status: string

### IngestHealthResponse  (4 fields)
    * generated_at: string<date-time>
    * lookback_hours: integer
    * services: [IngestHealthServiceRow]
      expected_missing: [string]

### IngestHealthServiceRow  (7 fields)
    * events_in_window: integer
    * service: string
      events_per_minute_avg: number
      events_per_minute_p95: number
      freshness: string(ok, stale, down)
      last_seen_at: string<date-time> | null
      outcomes_breakdown: map<string,integer>

### IngestLogBatchResponse  (4 fields)
    * count: integer
    * event_ids: [string]
    * message: string
    * status: string

### IngestLogResponse  (3 fields)
    * event_id: string
    * message: string
    * status: string

### IngestStreamResponse  (6 fields)
    * duration: number
    * events_received: integer
    * rate: number
    * results: [IngestStreamResult]
    * results_truncated: boolean
    * status: string

### IngestStreamResult  (5 fields)
      details: string | null
      error: string | null
      event_id: string | null
      line: integer | null
      status: string | null

### IntegrityVerifyRequest  (5 fields)
    * table: string
      end_date: string<date-time> | null
      limit: integer
      org_id: string | null
      start_date: string<date-time> | null

### JudgePolicyItem  (15 fields)
    * control_id: string
    * enabled: boolean
    * min_rule_confidence: number
    * model_name: string
    * org_id: string
    * policy_id: string
    * prompt_key: string
    * prompt_version: integer
    * tenant_id: string
      allowed_risk_levels: [string]
      allowed_rule_states: [string]
      created_at: string<date-time> | null
      created_by: string | null
      metadata: object
      updated_at: string<date-time> | null

### JudgePolicyUpsertRequest  (10 fields)
      allowed_risk_levels: [string]
      allowed_rule_states: [string]
      control_id: string
      enabled: boolean
      metadata: object
      min_rule_confidence: number
      model_name: string
      org_id: string | null
      prompt_key: string
      prompt_version: integer | null

### JudgePolicyUpsertResponse  (10 fields)
    * allowed_risk_levels: [string]
    * allowed_rule_states: [string]
    * control_id: string
    * enabled: boolean
    * min_rule_confidence: number
    * model_name: string
    * policy_id: string
    * prompt_key: string
    * prompt_version: integer
      status: string

### JudgePromptActivateResponse  (3 fields)
    * active_version: integer
    * prompt_key: string
      status: string

### JudgePromptItem  (12 fields)
    * enabled: boolean
    * org_id: string
    * prompt_id: string
    * prompt_key: string
    * template: string
    * tenant_id: string
    * version: integer
      created_at: string<date-time> | null
      created_by: string | null
      is_active: boolean
      output_schema: object
      updated_at: string<date-time> | null

### JudgePromptUpsertRequest  (7 fields)
    * template: string
    * version: integer
      activate: boolean
      enabled: boolean
      org_id: string | null
      output_schema: object
      prompt_key: string

### JudgePromptUpsertResponse  (6 fields)
    * enabled: boolean
    * is_active: boolean
    * prompt_id: string
    * prompt_key: string
    * version: integer
      status: string

### JudgeReplayResponse  (2 fields)
      replay: object
      status: string

### JudgeVerdictItem  (18 fields)
    * confidence: number
    * control_id: string
    * model_name: string
    * org_id: string
    * prompt_version: integer
    * tenant_id: string
    * verdict: string
    * verdict_id: string
      assertion_id: string | null
      citations: string | null
      created_at: string<date-time> | null
      input_payload: string | null
      policy_id: string | null
      policy_reason: string | null
      rationale: string | null
      raw_response: string | null
      run_id: string | null
      trace_id: string | null

### LegalHoldCreate  (4 fields)
    * reason: string
    * scope_type: LegalHoldScope
      before_date: string<date-time> | null
      data_subject_id: string | null

### LegalHoldRelease  (1 fields)
    * reason: string

### LegalHoldScope  (string(org, data_subject, before_date))
    enum: org, data_subject, before_date

### ListEnvelope_CallbackDeliveryItem_  (2 fields)
    * total: integer
      items: [CallbackDeliveryItem]

### ListEnvelope_CaseApprovalItem_  (2 fields)
    * total: integer
      items: [CaseApprovalItem]

### ListEnvelope_CaseEscalationItem_  (2 fields)
    * total: integer
      items: [CaseEscalationItem]

### ListEnvelope_CaseTimelineEntry_  (2 fields)
    * total: integer
      items: [CaseTimelineEntry]

### ListEnvelope_ComplianceTemplateItem_  (2 fields)
    * total: integer
      items: [ComplianceTemplateItem]

### ListEnvelope_ExplorerHistoryItem_  (2 fields)
    * total: integer
      items: [ExplorerHistoryItem]

### ListEnvelope_ExplorerSavedQueryItem_  (2 fields)
    * total: integer
      items: [ExplorerSavedQueryItem]

### ListEnvelope_FalsePositiveFeedbackItem_  (2 fields)
    * total: integer
      items: [FalsePositiveFeedbackItem]

### ListEnvelope_FrameworkPackItem_  (2 fields)
    * total: integer
      items: [FrameworkPackItem]

### ListEnvelope_JudgePolicyItem_  (2 fields)
    * total: integer
      items: [JudgePolicyItem]

### ListEnvelope_JudgePromptItem_  (2 fields)
    * total: integer
      items: [JudgePromptItem]

### ListEnvelope_JudgeVerdictItem_  (2 fields)
    * total: integer
      items: [JudgeVerdictItem]

### ListEnvelope_MergePolicyItem_  (2 fields)
    * total: integer
      items: [MergePolicyItem]

### ListEnvelope_MlBaselineItem_  (2 fields)
    * total: integer
      items: [MlBaselineItem]

### ListEnvelope_MlConfigItem_  (2 fields)
    * total: integer
      items: [MlConfigItem]

### ListEnvelope_ModelActivationItem_  (2 fields)
    * total: integer
      items: [ModelActivationItem]

### ListEnvelope_ModelArtifactItem_  (2 fields)
    * total: integer
      items: [ModelArtifactItem]

### ListEnvelope_ModelScoreItem_  (2 fields)
    * total: integer
      items: [ModelScoreItem]

### ListEnvelope_PolicyRuleItem_  (2 fields)
    * total: integer
      items: [PolicyRuleItem]

### ListEnvelope_SavedSearchItem_  (2 fields)
    * total: integer
      items: [SavedSearchItem]

### ListEnvelope_TrainingJobItem_  (2 fields)
    * total: integer
      items: [TrainingJobItem]

### LivenessResponse  (1 fields)
    * status: string

### MaintenanceOptimizeResponse  (2 fields)
    * message: string
    * timestamp: string<date-time>

### MergePolicyItem  (15 fields)
    * control_id: string
    * enabled: boolean
    * fail_override: boolean
    * fail_threshold: number
    * judge_weight: number
    * ml_weight: number
    * org_id: string
    * policy_id: string
    * rule_weight: number
    * tenant_id: string
    * warn_threshold: number
      created_at: string<date-time> | null
      created_by: string | null
      metadata: object
      updated_at: string<date-time> | null

### MergePolicyUpsertRequest  (10 fields)
      control_id: string
      enabled: boolean
      fail_override: boolean
      fail_threshold: number
      judge_weight: number
      metadata: object
      ml_weight: number
      org_id: string | null
      rule_weight: number
      warn_threshold: number

### MergePolicyUpsertResponse  (2 fields)
    * policy: MergePolicyItem
      status: string

### MlBaselineItem  (12 fields)
    * baseline_id: string
    * control_id: string
    * lookback_days: integer
    * mean_value: number
    * model_name: string
    * org_id: string
    * sample_count: integer
    * stddev_value: number
    * tenant_id: string
    * threshold_zscore: number
      calibration_metadata: string | null
      updated_at: string<date-time> | null

### MlConfigItem  (13 fields)
    * config_id: string
    * control_id: string
    * enabled: boolean
    * lookback_days: integer
    * min_samples: integer
    * model_name: string
    * org_id: string
    * tenant_id: string
    * threshold_zscore: number
      created_at: string<date-time> | null
      created_by: string | null
      metadata: object
      updated_at: string<date-time> | null

### MlConfigUpsertRequest  (8 fields)
    * control_id: string
      enabled: boolean
      lookback_days: integer
      metadata: object
      min_samples: integer
      model_name: string
      org_id: string | null
      threshold_zscore: number

### MlConfigUpsertResponse  (8 fields)
    * config_id: string
    * control_id: string
    * enabled: boolean
    * lookback_days: integer
    * min_samples: integer
    * model_name: string
    * threshold_zscore: number
      status: string

### MlDistillJobResponse  (9 fields)
    * job_id: string
    * job_type: string
    * org_id: string
    * requested_at: string<date-time>
    * tenant_id: string
      request: object
      run_id: string | null
      status: string
      trace_id: string | null

### MlDistillationTrainRequest  (8 fields)
      apply_as_default: boolean
      epochs: integer
      learning_rate: number
      limit: integer
      lookback_days: integer
      min_examples: integer
      model_family: string
      org_id: string | null

### MlModelActivationRequest  (2 fields)
      org_id: string | null
      reason: string

### MlModelRollbackRequest  (3 fields)
      model_name: string | null
      org_id: string | null
      reason: string

### ModelActivationItem  (10 fields)
    * action: string
    * activation_id: string
    * artifact_id: string
    * control_id: string
    * model_name: string
    * org_id: string
    * tenant_id: string
      activated_by: string | null
      created_at: string<date-time> | null
      reason: string | null

### ModelArtifactActivateResponse  (5 fields)
    * activation_id: string
    * artifact_id: string
    * org_id: string
    * tenant_id: string
      status: string

### ModelArtifactItem  (15 fields)
    * artifact_id: string
    * control_id: string
    * model_name: string
    * org_id: string
    * tenant_id: string
      artifact_payload: object | null
      baseline_id: string | null
      config_id: string | null
      created_at: string<date-time> | null
      created_by: string | null
      job_id: string | null
      metrics: object | null
      run_id: string | null
      trace_id: string | null
      version: string | null

### ModelControlRollbackResponse  (7 fields)
    * activation_id: string
    * artifact_id: string
    * control_id: string
    * org_id: string
    * tenant_id: string
      action: string
      status: string

### ModelScoreItem  (18 fields)
    * control_id: string
    * model_name: string
    * org_id: string
    * score_id: string
    * score_value: number
    * tenant_id: string
      delta_from_mean: number | null
      feature_value: number | null
      is_calibrated: boolean | null
      label: string | null
      mean_value: number | null
      run_id: string | null
      sample_count: integer | null
      scored_at: string<date-time> | null
      stddev_value: number | null
      threshold: number | null
      trace_id: string | null
      zscore: number | null

### ModuleHealthResponse  (13 fields)
    * allowed_environments: [string]
    * allowed_judge_models: [string]
    * callbacks_enabled: boolean
    * current_environment: string
    * llm_judge_enabled: boolean
    * llm_judge_fail_closed: boolean
    * llm_judge_model: string
    * llm_judge_provider: string
    * llm_judge_remote_configured: boolean
    * ml_scoring_enabled: boolean
    * module: string
    * status: string(enabled, disabled)
    * strict_idempotency: boolean

### NotificationFeed  (4 fields)
      items: [NotificationItem]
      last_seen_at: string<date-time> | null
      total_count: integer
      unread_count: integer

### NotificationItem  (7 fields)
    * id: string
    * kind: string
    * title: string
      created_at: string<date-time> | null
      deep_link: string
      detail: string
      severity: string

### NotificationPrefsSettings  (4 fields)
    * categories: object
      is_default: boolean
      updated_at: string<date-time> | null
      updated_by: string

### NotificationReadResponse  (3 fields)
    * last_seen_at: string<date-time>
      ok: boolean
      unread_count: integer

### OrganizationSummary  (7 fields)
    * active: boolean
    * events_last_week: integer
    * org_id: string
    * total_events: integer
    * unique_users: integer
      first_event: string<date-time> | null
      last_event: string<date-time> | null

### PaginatedResponse_AuditEvent_  (7 fields)
    * data: [AuditEvent]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PaginatedResponse_ComplianceReportRow_  (7 fields)
    * data: [ComplianceReportRow]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PaginatedResponse_DsarRequestRow_  (7 fields)
    * data: [DsarRequestRow]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PaginatedResponse_EvidencePackageWithLineage_  (7 fields)
    * data: [EvidencePackageWithLineage]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PaginatedResponse_ExportJobResponse_  (7 fields)
    * data: [ExportJobResponse]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PaginatedResponse_OrganizationSummary_  (7 fields)
    * data: [OrganizationSummary]
    * has_more: boolean
    * limit: integer
    * total: integer
      cursor: string | null
      next_cursor: string | null
      offset: integer | null

### PolicyRuleItem  (12 fields)
    * control_id: string
    * enabled: boolean
    * min_confidence_pass: number
    * min_confidence_warn: number
    * org_id: string
    * policy_id: string
    * tenant_id: string
    * version: integer
      created_at: string<date-time> | null
      created_by: string | null
      metadata: object
      updated_at: string<date-time> | null

### PolicyRuleUpsertRequest  (7 fields)
    * control_id: string
      enabled: boolean
      metadata: object
      min_confidence_pass: number
      min_confidence_warn: number
      org_id: string | null
      version: integer

### PolicyRuleUpsertResponse  (5 fields)
    * control_id: string
    * enabled: boolean
    * policy_id: string
    * version: integer
      status: string

### PostureHistoryPoint  (4 fields)
    * date: string<date-time>
      framework_scores: map<string,number>
      overall_score: number | null
      sample_event_count: integer

### PostureHistoryResponse  (4 fields)
    * generated_at: string<date-time>
    * org_id: string
    * points: [PostureHistoryPoint]
    * window_days: integer

### PurgeDataResponse  (7 fields)
    * dry_run: boolean
    * message: string
      deletion_id: string | null
      records_purged: integer | null
      records_quarantined: integer | null
      records_to_purge: integer | null
      status: string | null

### ReadinessResponse  (1 fields)
    * status: string

### RedactionSettings  (5 fields)
    * sensitive_headers: [string]
    * sensitive_json_keys: [string]
      is_default: boolean
      updated_at: string<date-time> | null
      updated_by: string

### ReportScheduleCreate  (4 fields)
    * cadence_cron: string
    * name: string
    * standard: string
      recipients: [string]

### ReportScheduleDeleteResponse  (2 fields)
    * deleted: boolean
    * schedule_id: string

### ReportScheduleItem  (12 fields)
    * cadence_cron: string
    * name: string
    * org_id: string
    * schedule_id: string
    * standard: string
      created_at: string<date-time> | null
      created_by: string | null
      last_run_at: string<date-time> | null
      next_run_at: string<date-time> | null
      paused: boolean
      recipients: [string]
      updated_at: string<date-time> | null

### ReportScheduleListResponse  (3 fields)
    * total: integer
      items: [ReportScheduleItem]
      scheduling_enabled: boolean

### ReportSchedulePauseResponse  (2 fields)
    * paused: boolean
    * schedule_id: string

### RetentionPeriod  (string(30d, 90d, 6m, 1y, 3y, 7y …))
    enum: 30d, 90d, 6m, 1y, 3y, 7y, 10y, indefinite

### RetentionPoliciesResponse  (1 fields)
    * policies: [RetentionPolicySummary]

### RetentionPolicySummary  (6 fields)
    * archive_after_days: integer
    * created_at: string<date-time>
    * hot_data_days: integer
    * org_id: string
    * retention_days: integer
    * updated_at: string<date-time>

### RetentionPolicyUpdate  (4 fields)
    * archive_after_days: integer
    * hot_data_days: integer
    * retention_days: integer
      org_id: string | null

### RiskCasesResponse  (2 fields)
    * total: integer
      items: [object]

### RolloutPreviewResponse  (5 fields)
    * component: string
    * percent: number
    * run_id: string
    * salt: string
    * selected: boolean

### RunTraceResponse  (7 fields)
    * org_id: string
    * run_id: string
    * tenant_id: string
    * total_rows: integer
      sections: map<string,integer>
      snapshot: map<string,[object]>
      trace_id: string | null

### SavedSearchCreate  (2 fields)
    * name: string
      filter_json: object

### SavedSearchDeleteResponse  (2 fields)
    * deleted: boolean
    * search_id: string

### SavedSearchItem  (7 fields)
    * name: string
    * org_id: string
    * search_id: string
    * user_id: string
      created_at: string<date-time> | null
      filter_json: object
      updated_at: string<date-time> | null

### SearchFacets  (5 fields)
      actions: [FacetCount]
      event_types: [FacetCount]
      outcomes: [FacetCount]
      resource_types: [FacetCount]
      services: [FacetCount]

### SearchFacetsResponse  (2 fields)
    * facets: SearchFacets
      org_id: string | null

### SearchQuery  (13 fields)
    * query: string
      actions: [string] | null
      cursor: string | null
      end_time: string<date-time> | null
      event_type_prefixes: [string] | null
      event_types: [string] | null
      limit: integer
      offset: integer
      org_id: string | null
      outcomes: [string] | null
      services: [string] | null
      start_time: string<date-time> | null
      user_ids: [string] | null

### SearchResult  (9 fields)
    * data: [AuditEvent]
    * execution_time_ms: number
    * has_more: boolean
    * limit: integer
    * offset: integer
    * query: string
    * total: integer
      cursor: string | null
      next_cursor: string | null

### SearchSuggestResponse  (3 fields)
    * field: string
    * prefix: string
    * suggestions: [SearchSuggestion]

### SearchSuggestion  (2 fields)
    * frequency: integer
    * value: string

### SecurityMetrics  (5 fields)
    * anomalies_detected: integer
    * failed_logins: integer
    * rate_limit_violations: integer
    * security_alerts: [object]
    * suspicious_activities: integer

### ServiceRootResponse  (6 fields)
    * capabilities: [string]
    * docs: map<string,string>
    * service: string
    * status: string
    * ui: map<string,string>
    * version: string

### ShareLink  (10 fields)
    * auditor_email: string
    * expires_at: string<date-time>
    * package_id: string
    * token: string
      access_count: integer
      created_at: string<date-time> | null
      created_by: string | null
      last_accessed_at: string<date-time> | null
      org_id: string | null
      revoked_at: string<date-time> | null

### ShareLinkCreate  (2 fields)
    * auditor_email: string
    * expires_at: string<date-time>

### ShareLinkSummary  (12 fields)
    * auditor_email: string
    * expires_at: string<date-time>
    * package_id: string
    * token: string
    * token_prefix: string
      access_count: integer
      created_at: string<date-time> | null
      created_by: string | null
      last_accessed_at: string<date-time> | null
      org_id: string | null
      revoked_at: string<date-time> | null
      status: string

### Soc2AuditTrailEvent  (13 fields)
    * action: string
    * event_id: string
    * event_type: string
    * integrity_verified: boolean
    * outcome: string
    * timestamp: string<date-time>
      event_hash: string | null
      ip_address: string | null
      metadata: string | null
      session_id: string | null
      soc2_control_relevance: [string]
      user_agent: string | null
      user_id: string | null

### Soc2AuditTrailResponse  (5 fields)
    * audit_trail: [Soc2AuditTrailEvent]
    * generated_at: string<date-time>
    * generated_for_org: string
    * summary: Soc2AuditTrailSummary
      control_objective: string | null

### Soc2AuditTrailSummary  (5 fields)
    * date_range: map<string,string<date-time>>
    * event_types_covered: [string]
    * integrity_status: string(verified, warning)
    * total_events: integer
    * unique_users: integer

### StatusHistoryEntry  (6 fields)
      actor: string | null
      from_status: string | null
      note: string | null
      status: string | null
      timestamp: string | null
      to_status: string | null

### SupersedeRequest  (7 fields)
    * reason: string
      control_ids: [string] | null
      event_ids: [string] | null
      framework: string | null
      name: string | null
      time_range_end: string<date-time> | null
      time_range_start: string<date-time> | null

### SystemStatus  (8 fields)
    * background_tasks: map<string,string>
    * cache_status: string
    * database_status: string
    * metrics: object
    * queue_status: string
    * service: string
    * uptime_seconds: number
    * version: string

### TemplateCloneBody  (1 fields)
      vars: map<string,string>

### TimeSeriesPoint  (2 fields)
    * timestamp: string<date-time>
    * value: integer

### ToolInventoryItem  (5 fields)
    * name: string
      callers: integer
      calls: integer
      failure: integer
      success: integer

### ToolInventoryResponse  (4 fields)
      distinct_tools: integer
      org_id: string | null
      tools: [ToolInventoryItem]
      total_calls: integer

### TopUser  (6 fields)
    * active_days: integer
    * event_count: integer
    * failed_events: integer
    * success_rate: number
    * unique_event_types: integer
    * user_id: string

### TopUsersPeriod  (2 fields)
    * end: string<date-time>
    * start: string<date-time>

### TopUsersResponse  (2 fields)
    * period: TopUsersPeriod
    * top_users: [TopUser]

### TrainingJobItem  (18 fields)
    * job_id: string
    * org_id: string
    * status: string
    * tenant_id: string
      attempts: integer
      created_at: string<date-time> | null
      job_type: string | null
      last_error: string | null
      lease_token: string | null
      lease_until: string<date-time> | null
      model_family: string | null
      payload: object | null
      requested_by: string | null
      result_payload: object | null
      run_id: string | null
      status_version: integer | null
      trace_id: string | null
      updated_at: string<date-time> | null

### ValidationError  (3 fields)
    * loc: [string | integer]
    * msg: string
    * type: string

