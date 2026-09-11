// ANNOTATION BLOCK
// File: svc_complianceautomation.go — the `ComplianceAutomation` designation of the Audit Service.
// Why it exists: typed methods for every ComplianceAutomation endpoint, reached via
// c.ComplianceAutomation.<Method>. Each is context-first, takes typed path/query/body
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

// ComplianceAutomationActivateFrameworkPackParams holds the query parameters for ComplianceAutomation.ActivateFrameworkPack.
type ComplianceAutomationActivateFrameworkPackParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationActivateFrameworkPackParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationActivateJudgePromptParams holds the query parameters for ComplianceAutomation.ActivateJudgePrompt.
type ComplianceAutomationActivateJudgePromptParams struct {
	Version *int    `json:"version,omitempty"` // required.
	OrgID   *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationActivateJudgePromptParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Version != nil {
		v.Set("version", strconv.Itoa(*p.Version))
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationDeactivateFrameworkPackParams holds the query parameters for ComplianceAutomation.DeactivateFrameworkPack.
type ComplianceAutomationDeactivateFrameworkPackParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationDeactivateFrameworkPackParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationGetAssertionsParams holds the query parameters for ComplianceAutomation.GetAssertions.
type ComplianceAutomationGetAssertionsParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	Framework *string `json:"framework,omitempty"`
}

func (p *ComplianceAutomationGetAssertionsParams) values() url.Values {
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

// ComplianceAutomationGetCallbackDeliveriesParams holds the query parameters for ComplianceAutomation.GetCallbackDeliveries.
type ComplianceAutomationGetCallbackDeliveriesParams struct {
	OrgID          *string `json:"org_id,omitempty"`
	EventTopic     *string `json:"event_topic,omitempty"`
	SubscriptionID *string `json:"subscription_id,omitempty"`
	Status         *string `json:"status,omitempty"`
	Limit          *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetCallbackDeliveriesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.EventTopic != nil {
		v.Set("event_topic", *p.EventTopic)
	}
	if p.SubscriptionID != nil {
		v.Set("subscription_id", *p.SubscriptionID)
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetCallbackSubscriptionsParams holds the query parameters for ComplianceAutomation.GetCallbackSubscriptions.
type ComplianceAutomationGetCallbackSubscriptionsParams struct {
	OrgID       *string `json:"org_id,omitempty"`
	OnlyEnabled *bool   `json:"only_enabled,omitempty"`
}

func (p *ComplianceAutomationGetCallbackSubscriptionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.OnlyEnabled != nil {
		v.Set("only_enabled", strconv.FormatBool(*p.OnlyEnabled))
	}
	return v
}

// ComplianceAutomationGetFalsePositiveReportsParams holds the query parameters for ComplianceAutomation.GetFalsePositiveReports.
type ComplianceAutomationGetFalsePositiveReportsParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	Framework *string `json:"framework,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetFalsePositiveReportsParams) values() url.Values {
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
	if p.ControlID != nil {
		v.Set("control_id", *p.ControlID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetFrameworkPackReadinessParams holds the query parameters for ComplianceAutomation.GetFrameworkPackReadiness.
type ComplianceAutomationGetFrameworkPackReadinessParams struct {
	OrgID        *string `json:"org_id,omitempty"`
	LookbackDays *int    `json:"lookback_days,omitempty"`
}

func (p *ComplianceAutomationGetFrameworkPackReadinessParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.LookbackDays != nil {
		v.Set("lookback_days", strconv.Itoa(*p.LookbackDays))
	}
	return v
}

// ComplianceAutomationGetFrameworkPackRegistryParams holds the query parameters for ComplianceAutomation.GetFrameworkPackRegistry.
type ComplianceAutomationGetFrameworkPackRegistryParams struct {
	IncludeControls *bool   `json:"include_controls,omitempty"`
	OrgID           *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationGetFrameworkPackRegistryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.IncludeControls != nil {
		v.Set("include_controls", strconv.FormatBool(*p.IncludeControls))
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationGetJudgePoliciesParams holds the query parameters for ComplianceAutomation.GetJudgePolicies.
type ComplianceAutomationGetJudgePoliciesParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationGetJudgePoliciesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationGetJudgePromptsParams holds the query parameters for ComplianceAutomation.GetJudgePrompts.
type ComplianceAutomationGetJudgePromptsParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	PromptKey *string `json:"prompt_key,omitempty"`
}

func (p *ComplianceAutomationGetJudgePromptsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.PromptKey != nil {
		v.Set("prompt_key", *p.PromptKey)
	}
	return v
}

// ComplianceAutomationGetJudgeVerdictsParams holds the query parameters for ComplianceAutomation.GetJudgeVerdicts.
type ComplianceAutomationGetJudgeVerdictsParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
	RunID     *string `json:"run_id,omitempty"`
	TraceID   *string `json:"trace_id,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetJudgeVerdictsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.ControlID != nil {
		v.Set("control_id", *p.ControlID)
	}
	if p.RunID != nil {
		v.Set("run_id", *p.RunID)
	}
	if p.TraceID != nil {
		v.Set("trace_id", *p.TraceID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMLActivationsParams holds the query parameters for ComplianceAutomation.GetMLActivations.
type ComplianceAutomationGetMLActivationsParams struct {
	OrgID *string `json:"org_id,omitempty"`
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetMLActivationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMLArtifactsParams holds the query parameters for ComplianceAutomation.GetMLArtifacts.
type ComplianceAutomationGetMLArtifactsParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
	ModelName *string `json:"model_name,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetMLArtifactsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.ControlID != nil {
		v.Set("control_id", *p.ControlID)
	}
	if p.ModelName != nil {
		v.Set("model_name", *p.ModelName)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMLBaselinesParams holds the query parameters for ComplianceAutomation.GetMLBaselines.
type ComplianceAutomationGetMLBaselinesParams struct {
	OrgID *string `json:"org_id,omitempty"`
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetMLBaselinesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMLConfigsParams holds the query parameters for ComplianceAutomation.GetMLConfigs.
type ComplianceAutomationGetMLConfigsParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationGetMLConfigsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationGetMLScoresParams holds the query parameters for ComplianceAutomation.GetMLScores.
type ComplianceAutomationGetMLScoresParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
	Label     *string `json:"label,omitempty"`
	RunID     *string `json:"run_id,omitempty"`
	TraceID   *string `json:"trace_id,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetMLScoresParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.ControlID != nil {
		v.Set("control_id", *p.ControlID)
	}
	if p.Label != nil {
		v.Set("label", *p.Label)
	}
	if p.RunID != nil {
		v.Set("run_id", *p.RunID)
	}
	if p.TraceID != nil {
		v.Set("trace_id", *p.TraceID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMLTrainingJobsParams holds the query parameters for ComplianceAutomation.GetMLTrainingJobs.
type ComplianceAutomationGetMLTrainingJobsParams struct {
	OrgID *string `json:"org_id,omitempty"`
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetMLTrainingJobsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetMergePoliciesParams holds the query parameters for ComplianceAutomation.GetMergePolicies.
type ComplianceAutomationGetMergePoliciesParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	ControlID *string `json:"control_id,omitempty"`
}

func (p *ComplianceAutomationGetMergePoliciesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.ControlID != nil {
		v.Set("control_id", *p.ControlID)
	}
	return v
}

// ComplianceAutomationGetPolicyRulesParams holds the query parameters for ComplianceAutomation.GetPolicyRules.
type ComplianceAutomationGetPolicyRulesParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationGetPolicyRulesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationGetRiskCaseApprovalsParams holds the query parameters for ComplianceAutomation.GetRiskCaseApprovals.
type ComplianceAutomationGetRiskCaseApprovalsParams struct {
	OrgID *string `json:"org_id,omitempty"`
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetRiskCaseApprovalsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetRiskCaseEscalationsParams holds the query parameters for ComplianceAutomation.GetRiskCaseEscalations.
type ComplianceAutomationGetRiskCaseEscalationsParams struct {
	OrgID   *string `json:"org_id,omitempty"`
	CaseID  *string `json:"case_id,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Status  *string `json:"status,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetRiskCaseEscalationsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.CaseID != nil {
		v.Set("case_id", *p.CaseID)
	}
	if p.Channel != nil {
		v.Set("channel", *p.Channel)
	}
	if p.Status != nil {
		v.Set("status", *p.Status)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetRiskCaseTimelineParams holds the query parameters for ComplianceAutomation.GetRiskCaseTimeline.
type ComplianceAutomationGetRiskCaseTimelineParams struct {
	OrgID *string `json:"org_id,omitempty"`
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetRiskCaseTimelineParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetRiskCasesParams holds the query parameters for ComplianceAutomation.GetRiskCases.
type ComplianceAutomationGetRiskCasesParams struct {
	OrgID     *string `json:"org_id,omitempty"`
	Status    *string `json:"status,omitempty"`
	RiskLevel *string `json:"risk_level,omitempty"`
	RunID     *string `json:"run_id,omitempty"`
	TraceID   *string `json:"trace_id,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
}

func (p *ComplianceAutomationGetRiskCasesParams) values() url.Values {
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
	if p.RiskLevel != nil {
		v.Set("risk_level", *p.RiskLevel)
	}
	if p.RunID != nil {
		v.Set("run_id", *p.RunID)
	}
	if p.TraceID != nil {
		v.Set("trace_id", *p.TraceID)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// ComplianceAutomationGetRunTraceParams holds the query parameters for ComplianceAutomation.GetRunTrace.
type ComplianceAutomationGetRunTraceParams struct {
	OrgID           *string `json:"org_id,omitempty"`
	TraceID         *string `json:"trace_id,omitempty"`
	LimitPerSection *int    `json:"limit_per_section,omitempty"`
}

func (p *ComplianceAutomationGetRunTraceParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.TraceID != nil {
		v.Set("trace_id", *p.TraceID)
	}
	if p.LimitPerSection != nil {
		v.Set("limit_per_section", strconv.Itoa(*p.LimitPerSection))
	}
	return v
}

// ComplianceAutomationPauseCallbackSubscriptionParams holds the query parameters for ComplianceAutomation.PauseCallbackSubscription.
type ComplianceAutomationPauseCallbackSubscriptionParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationPauseCallbackSubscriptionParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationPreviewCanaryRolloutParams holds the query parameters for ComplianceAutomation.PreviewCanaryRollout.
type ComplianceAutomationPreviewCanaryRolloutParams struct {
	Component *string  `json:"component,omitempty"` // required.
	RunID     *string  `json:"run_id,omitempty"`    // required.
	Percent   *float64 `json:"percent,omitempty"`   // required.
	Salt      *string  `json:"salt,omitempty"`
}

func (p *ComplianceAutomationPreviewCanaryRolloutParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Component != nil {
		v.Set("component", *p.Component)
	}
	if p.RunID != nil {
		v.Set("run_id", *p.RunID)
	}
	if p.Percent != nil {
		v.Set("percent", strconv.FormatFloat(*p.Percent, 'g', -1, 64))
	}
	if p.Salt != nil {
		v.Set("salt", *p.Salt)
	}
	return v
}

// ComplianceAutomationReplayStoredJudgeVerdictParams holds the query parameters for ComplianceAutomation.ReplayStoredJudgeVerdict.
type ComplianceAutomationReplayStoredJudgeVerdictParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationReplayStoredJudgeVerdictParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationResumeCallbackSubscriptionParams holds the query parameters for ComplianceAutomation.ResumeCallbackSubscription.
type ComplianceAutomationResumeCallbackSubscriptionParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationResumeCallbackSubscriptionParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceAutomationRollbackJudgePromptParams holds the query parameters for ComplianceAutomation.RollbackJudgePrompt.
type ComplianceAutomationRollbackJudgePromptParams struct {
	ToVersion *int    `json:"to_version,omitempty"` // required.
	OrgID     *string `json:"org_id,omitempty"`
}

func (p *ComplianceAutomationRollbackJudgePromptParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.ToVersion != nil {
		v.Set("to_version", strconv.Itoa(*p.ToVersion))
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ActivateFrameworkPack calls POST /compliance-automation/framework-packs/{framework}/activate.
// Activate Framework Pack
//
// T9 §10.13 — POST /framework-packs/{framework}/activate.
func (s *ComplianceAutomationClient) ActivateFrameworkPack(ctx context.Context, framework string, params *ComplianceAutomationActivateFrameworkPackParams) (*FrameworkActivationResponse, error) {
	path := fmt.Sprintf("/compliance-automation/framework-packs/%s/activate", url.PathEscape(framework))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out FrameworkActivationResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ActivateJudgePrompt calls POST /compliance-automation/judge/prompts/{prompt_key}/activate.
// Activate Judge Prompt
func (s *ComplianceAutomationClient) ActivateJudgePrompt(ctx context.Context, promptKey string, params *ComplianceAutomationActivateJudgePromptParams) (*JudgePromptActivateResponse, error) {
	path := fmt.Sprintf("/compliance-automation/judge/prompts/%s/activate", url.PathEscape(promptKey))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out JudgePromptActivateResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ActivateMLArtifact calls POST /compliance-automation/ml/artifacts/{artifact_id}/activate.
// Activate Ml Artifact
func (s *ComplianceAutomationClient) ActivateMLArtifact(ctx context.Context, artifactID string, body MlModelActivationRequest) (*ModelArtifactActivateResponse, error) {
	path := fmt.Sprintf("/compliance-automation/ml/artifacts/%s/activate", url.PathEscape(artifactID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out ModelArtifactActivateResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// CreateCallbackSubscription calls POST /compliance-automation/callbacks/subscriptions.
// Create Callback Subscription
func (s *ComplianceAutomationClient) CreateCallbackSubscription(ctx context.Context, body CallbackSubscriptionCreateRequest) (*CallbackSubscriptionResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/callbacks/subscriptions", nil, body)
	if err != nil {
		return nil, err
	}
	var out CallbackSubscriptionResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/callbacks/subscriptions", Err: err}
	}
	return &out, nil
}

// CreateOrUpdateJudgePolicy calls POST /compliance-automation/judge/policies.
// Create Or Update Judge Policy
func (s *ComplianceAutomationClient) CreateOrUpdateJudgePolicy(ctx context.Context, body JudgePolicyUpsertRequest) (*JudgePolicyUpsertResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/judge/policies", nil, body)
	if err != nil {
		return nil, err
	}
	var out JudgePolicyUpsertResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/judge/policies", Err: err}
	}
	return &out, nil
}

// CreateOrUpdateJudgePrompt calls POST /compliance-automation/judge/prompts.
// Create Or Update Judge Prompt
func (s *ComplianceAutomationClient) CreateOrUpdateJudgePrompt(ctx context.Context, body JudgePromptUpsertRequest) (*JudgePromptUpsertResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/judge/prompts", nil, body)
	if err != nil {
		return nil, err
	}
	var out JudgePromptUpsertResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/judge/prompts", Err: err}
	}
	return &out, nil
}

// CreateOrUpdateMLConfig calls POST /compliance-automation/ml/configs.
// Create Or Update Ml Config
func (s *ComplianceAutomationClient) CreateOrUpdateMLConfig(ctx context.Context, body MlConfigUpsertRequest) (*MlConfigUpsertResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/ml/configs", nil, body)
	if err != nil {
		return nil, err
	}
	var out MlConfigUpsertResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/configs", Err: err}
	}
	return &out, nil
}

// CreateOrUpdateMergePolicy calls POST /compliance-automation/merge-policies.
// Create Or Update Merge Policy
func (s *ComplianceAutomationClient) CreateOrUpdateMergePolicy(ctx context.Context, body MergePolicyUpsertRequest) (*MergePolicyUpsertResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/merge-policies", nil, body)
	if err != nil {
		return nil, err
	}
	var out MergePolicyUpsertResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/merge-policies", Err: err}
	}
	return &out, nil
}

// CreateOrUpdatePolicyRule calls POST /compliance-automation/policies/rules.
// Create Or Update Policy Rule
func (s *ComplianceAutomationClient) CreateOrUpdatePolicyRule(ctx context.Context, body PolicyRuleUpsertRequest) (*PolicyRuleUpsertResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/policies/rules", nil, body)
	if err != nil {
		return nil, err
	}
	var out PolicyRuleUpsertResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/policies/rules", Err: err}
	}
	return &out, nil
}

// DeactivateFrameworkPack calls POST /compliance-automation/framework-packs/{framework}/deactivate.
// Deactivate Framework Pack
//
// T9 §10.13 — POST /framework-packs/{framework}/deactivate.
func (s *ComplianceAutomationClient) DeactivateFrameworkPack(ctx context.Context, framework string, params *ComplianceAutomationDeactivateFrameworkPackParams) (*FrameworkActivationResponse, error) {
	path := fmt.Sprintf("/compliance-automation/framework-packs/%s/deactivate", url.PathEscape(framework))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out FrameworkActivationResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// DecideRiskCaseApproval calls POST /compliance-automation/risk-cases/{case_id}/approvals.
// Decide Risk Case Approval
func (s *ComplianceAutomationClient) DecideRiskCaseApproval(ctx context.Context, caseID string, body CaseApprovalDecisionRequest) (*CaseApprovalDecisionResponse, error) {
	path := fmt.Sprintf("/compliance-automation/risk-cases/%s/approvals", url.PathEscape(caseID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out CaseApprovalDecisionResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// EvaluateEventNow calls POST /compliance-automation/events/evaluate.
// Evaluate Event Now
func (s *ComplianceAutomationClient) EvaluateEventNow(ctx context.Context, body EvaluationEventRequest) (*EvaluateEventResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/events/evaluate", nil, body)
	if err != nil {
		return nil, err
	}
	var out EvaluateEventResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/events/evaluate", Err: err}
	}
	return &out, nil
}

// GetAssertions calls GET /compliance-automation/assertions.
// Get Assertions
func (s *ComplianceAutomationClient) GetAssertions(ctx context.Context, params *ComplianceAutomationGetAssertionsParams) (*AssertionsResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/assertions", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out AssertionsResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/assertions", Err: err}
	}
	return &out, nil
}

// GetCallbackDeliveries calls GET /compliance-automation/callbacks/deliveries.
// Get Callback Deliveries
func (s *ComplianceAutomationClient) GetCallbackDeliveries(ctx context.Context, params *ComplianceAutomationGetCallbackDeliveriesParams) (*ListEnvelopeCallbackDeliveryItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/callbacks/deliveries", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeCallbackDeliveryItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/callbacks/deliveries", Err: err}
	}
	return &out, nil
}

// GetCallbackSubscriptions calls GET /compliance-automation/callbacks/subscriptions.
// Get Callback Subscriptions
func (s *ComplianceAutomationClient) GetCallbackSubscriptions(ctx context.Context, params *ComplianceAutomationGetCallbackSubscriptionsParams) ([]CallbackSubscriptionResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/callbacks/subscriptions", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out []CallbackSubscriptionResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/callbacks/subscriptions", Err: err}
	}
	return out, nil
}

// GetFalsePositiveReports calls GET /compliance-automation/quality/false-positive.
// Get False Positive Reports
func (s *ComplianceAutomationClient) GetFalsePositiveReports(ctx context.Context, params *ComplianceAutomationGetFalsePositiveReportsParams) (*ListEnvelopeFalsePositiveFeedbackItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/quality/false-positive", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeFalsePositiveFeedbackItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/quality/false-positive", Err: err}
	}
	return &out, nil
}

// GetFrameworkPackDetails calls GET /compliance-automation/framework-packs/{framework}.
// Get Framework Pack Details
func (s *ComplianceAutomationClient) GetFrameworkPackDetails(ctx context.Context, framework string) (*FrameworkPackItem, error) {
	path := fmt.Sprintf("/compliance-automation/framework-packs/%s", url.PathEscape(framework))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out FrameworkPackItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetFrameworkPackReadiness calls GET /compliance-automation/framework-packs/{framework}/readiness.
// Get Framework Pack Readiness
func (s *ComplianceAutomationClient) GetFrameworkPackReadiness(ctx context.Context, framework string, params *ComplianceAutomationGetFrameworkPackReadinessParams) (*FrameworkReadinessResponse, error) {
	path := fmt.Sprintf("/compliance-automation/framework-packs/%s/readiness", url.PathEscape(framework))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out FrameworkReadinessResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetFrameworkPackRegistry calls GET /compliance-automation/framework-packs.
// Get Framework Pack Registry
func (s *ComplianceAutomationClient) GetFrameworkPackRegistry(ctx context.Context, params *ComplianceAutomationGetFrameworkPackRegistryParams) (*ListEnvelopeFrameworkPackItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/framework-packs", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeFrameworkPackItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/framework-packs", Err: err}
	}
	return &out, nil
}

// GetJudgePolicies calls GET /compliance-automation/judge/policies.
// Get Judge Policies
func (s *ComplianceAutomationClient) GetJudgePolicies(ctx context.Context, params *ComplianceAutomationGetJudgePoliciesParams) (*ListEnvelopeJudgePolicyItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/judge/policies", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeJudgePolicyItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/judge/policies", Err: err}
	}
	return &out, nil
}

// GetJudgePrompts calls GET /compliance-automation/judge/prompts.
// Get Judge Prompts
func (s *ComplianceAutomationClient) GetJudgePrompts(ctx context.Context, params *ComplianceAutomationGetJudgePromptsParams) (*ListEnvelopeJudgePromptItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/judge/prompts", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeJudgePromptItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/judge/prompts", Err: err}
	}
	return &out, nil
}

// GetJudgeVerdicts calls GET /compliance-automation/judge/verdicts.
// Get Judge Verdicts
func (s *ComplianceAutomationClient) GetJudgeVerdicts(ctx context.Context, params *ComplianceAutomationGetJudgeVerdictsParams) (*ListEnvelopeJudgeVerdictItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/judge/verdicts", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeJudgeVerdictItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/judge/verdicts", Err: err}
	}
	return &out, nil
}

// GetMLActivations calls GET /compliance-automation/ml/activations.
// Get Ml Activations
func (s *ComplianceAutomationClient) GetMLActivations(ctx context.Context, params *ComplianceAutomationGetMLActivationsParams) (*ListEnvelopeModelActivationItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/activations", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeModelActivationItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/activations", Err: err}
	}
	return &out, nil
}

// GetMLArtifacts calls GET /compliance-automation/ml/artifacts.
// Get Ml Artifacts
func (s *ComplianceAutomationClient) GetMLArtifacts(ctx context.Context, params *ComplianceAutomationGetMLArtifactsParams) (*ListEnvelopeModelArtifactItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/artifacts", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeModelArtifactItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/artifacts", Err: err}
	}
	return &out, nil
}

// GetMLBaselines calls GET /compliance-automation/ml/baselines.
// Get Ml Baselines
func (s *ComplianceAutomationClient) GetMLBaselines(ctx context.Context, params *ComplianceAutomationGetMLBaselinesParams) (*ListEnvelopeMlBaselineItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/baselines", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeMlBaselineItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/baselines", Err: err}
	}
	return &out, nil
}

// GetMLConfigs calls GET /compliance-automation/ml/configs.
// Get Ml Configs
func (s *ComplianceAutomationClient) GetMLConfigs(ctx context.Context, params *ComplianceAutomationGetMLConfigsParams) (*ListEnvelopeMlConfigItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/configs", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeMlConfigItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/configs", Err: err}
	}
	return &out, nil
}

// GetMLScores calls GET /compliance-automation/ml/scores.
// Get Ml Scores
func (s *ComplianceAutomationClient) GetMLScores(ctx context.Context, params *ComplianceAutomationGetMLScoresParams) (*ListEnvelopeModelScoreItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/scores", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeModelScoreItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/scores", Err: err}
	}
	return &out, nil
}

// GetMLTrainingJobs calls GET /compliance-automation/ml/jobs.
// Get Ml Training Jobs
func (s *ComplianceAutomationClient) GetMLTrainingJobs(ctx context.Context, params *ComplianceAutomationGetMLTrainingJobsParams) (*ListEnvelopeTrainingJobItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/ml/jobs", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeTrainingJobItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/jobs", Err: err}
	}
	return &out, nil
}

// GetMergePolicies calls GET /compliance-automation/merge-policies.
// Get Merge Policies
func (s *ComplianceAutomationClient) GetMergePolicies(ctx context.Context, params *ComplianceAutomationGetMergePoliciesParams) (*ListEnvelopeMergePolicyItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/merge-policies", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeMergePolicyItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/merge-policies", Err: err}
	}
	return &out, nil
}

// GetPolicyRules calls GET /compliance-automation/policies/rules.
// Get Policy Rules
func (s *ComplianceAutomationClient) GetPolicyRules(ctx context.Context, params *ComplianceAutomationGetPolicyRulesParams) (*ListEnvelopePolicyRuleItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/policies/rules", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopePolicyRuleItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/policies/rules", Err: err}
	}
	return &out, nil
}

// GetRiskCaseApprovals calls GET /compliance-automation/risk-cases/{case_id}/approvals.
// Get Risk Case Approvals
func (s *ComplianceAutomationClient) GetRiskCaseApprovals(ctx context.Context, caseID string, params *ComplianceAutomationGetRiskCaseApprovalsParams) (*ListEnvelopeCaseApprovalItem, error) {
	path := fmt.Sprintf("/compliance-automation/risk-cases/%s/approvals", url.PathEscape(caseID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeCaseApprovalItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetRiskCaseEscalations calls GET /compliance-automation/risk-cases/escalations.
// Get Risk Case Escalations
func (s *ComplianceAutomationClient) GetRiskCaseEscalations(ctx context.Context, params *ComplianceAutomationGetRiskCaseEscalationsParams) (*ListEnvelopeCaseEscalationItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/risk-cases/escalations", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeCaseEscalationItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/risk-cases/escalations", Err: err}
	}
	return &out, nil
}

// GetRiskCaseTimeline calls GET /compliance-automation/risk-cases/{case_id}/timeline.
// Get Risk Case Timeline
func (s *ComplianceAutomationClient) GetRiskCaseTimeline(ctx context.Context, caseID string, params *ComplianceAutomationGetRiskCaseTimelineParams) (*ListEnvelopeCaseTimelineEntry, error) {
	path := fmt.Sprintf("/compliance-automation/risk-cases/%s/timeline", url.PathEscape(caseID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeCaseTimelineEntry
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// GetRiskCases calls GET /compliance-automation/risk-cases.
// Get Risk Cases
func (s *ComplianceAutomationClient) GetRiskCases(ctx context.Context, params *ComplianceAutomationGetRiskCasesParams) (*RiskCasesResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/risk-cases", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out RiskCasesResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/risk-cases", Err: err}
	}
	return &out, nil
}

// GetRunTrace calls GET /compliance-automation/runs/{run_id}/trace.
// Get Run Trace
func (s *ComplianceAutomationClient) GetRunTrace(ctx context.Context, runID string, params *ComplianceAutomationGetRunTraceParams) (*RunTraceResponse, error) {
	path := fmt.Sprintf("/compliance-automation/runs/%s/trace", url.PathEscape(runID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out RunTraceResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ModuleHealth calls GET /compliance-automation/health.
// Module Health
func (s *ComplianceAutomationClient) ModuleHealth(ctx context.Context) (*ModuleHealthResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/health", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ModuleHealthResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/health", Err: err}
	}
	return &out, nil
}

// PauseCallbackSubscription calls POST /compliance-automation/callbacks/subscriptions/{subscription_id}/pause.
// Pause Callback Subscription
func (s *ComplianceAutomationClient) PauseCallbackSubscription(ctx context.Context, subscriptionID string, params *ComplianceAutomationPauseCallbackSubscriptionParams) (*CallbackSubscriptionPauseResponse, error) {
	path := fmt.Sprintf("/compliance-automation/callbacks/subscriptions/%s/pause", url.PathEscape(subscriptionID))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out CallbackSubscriptionPauseResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// PreviewCanaryRollout calls GET /compliance-automation/rollouts/preview.
// Preview Canary Rollout
func (s *ComplianceAutomationClient) PreviewCanaryRollout(ctx context.Context, params *ComplianceAutomationPreviewCanaryRolloutParams) (*RolloutPreviewResponse, error) {
	b, err := s.c.do(ctx, "GET", "/compliance-automation/rollouts/preview", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out RolloutPreviewResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/rollouts/preview", Err: err}
	}
	return &out, nil
}

// ReplayStoredJudgeVerdict calls POST /compliance-automation/judge/verdicts/{verdict_id}/replay.
// Replay Stored Judge Verdict
func (s *ComplianceAutomationClient) ReplayStoredJudgeVerdict(ctx context.Context, verdictID string, params *ComplianceAutomationReplayStoredJudgeVerdictParams) (*JudgeReplayResponse, error) {
	path := fmt.Sprintf("/compliance-automation/judge/verdicts/%s/replay", url.PathEscape(verdictID))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out JudgeReplayResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ReportFalsePositive calls POST /compliance-automation/quality/false-positive.
// Report False Positive
func (s *ComplianceAutomationClient) ReportFalsePositive(ctx context.Context, body FalsePositiveFeedbackRequest) (*FalsePositiveReportResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/quality/false-positive", nil, body)
	if err != nil {
		return nil, err
	}
	var out FalsePositiveReportResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/quality/false-positive", Err: err}
	}
	return &out, nil
}

// ResumeCallbackSubscription calls POST /compliance-automation/callbacks/subscriptions/{subscription_id}/resume.
// Resume Callback Subscription
//
// Re-enable a paused callback subscription. Inverse of
// “pause_callback_subscription“: re-upserts the same subscription record
// with “enabled=True“ so an admin can recover a paused webhook (e.g.
// after an endpoint-outage / maintenance window) without delete+recreate —
// preserving the “subscription_id“ and its delivery history.
func (s *ComplianceAutomationClient) ResumeCallbackSubscription(ctx context.Context, subscriptionID string, params *ComplianceAutomationResumeCallbackSubscriptionParams) (*CallbackSubscriptionResumeResponse, error) {
	path := fmt.Sprintf("/compliance-automation/callbacks/subscriptions/%s/resume", url.PathEscape(subscriptionID))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out CallbackSubscriptionResumeResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// RollbackJudgePrompt calls POST /compliance-automation/judge/prompts/{prompt_key}/rollback.
// Rollback Judge Prompt
func (s *ComplianceAutomationClient) RollbackJudgePrompt(ctx context.Context, promptKey string, params *ComplianceAutomationRollbackJudgePromptParams) (*JudgePromptActivateResponse, error) {
	path := fmt.Sprintf("/compliance-automation/judge/prompts/%s/rollback", url.PathEscape(promptKey))
	b, err := s.c.do(ctx, "POST", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out JudgePromptActivateResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// RollbackMLControl calls POST /compliance-automation/ml/controls/{control_id}/rollback.
// Rollback Ml Control
func (s *ComplianceAutomationClient) RollbackMLControl(ctx context.Context, controlID string, body MlModelRollbackRequest) (*ModelControlRollbackResponse, error) {
	path := fmt.Sprintf("/compliance-automation/ml/controls/%s/rollback", url.PathEscape(controlID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out ModelControlRollbackResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// SendTestCallback calls POST /compliance-automation/callbacks/subscriptions/{subscription_id}/test.
// Send Test Callback
func (s *ComplianceAutomationClient) SendTestCallback(ctx context.Context, subscriptionID string, body CallbackTestRequest) (*CallbackSubscriptionTestResponse, error) {
	path := fmt.Sprintf("/compliance-automation/callbacks/subscriptions/%s/test", url.PathEscape(subscriptionID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out CallbackSubscriptionTestResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// TrainDistilledMLStudents calls POST /compliance-automation/ml/distill.
// Train Distilled Ml Students
func (s *ComplianceAutomationClient) TrainDistilledMLStudents(ctx context.Context, body MlDistillationTrainRequest) (*MlDistillJobResponse, error) {
	b, err := s.c.do(ctx, "POST", "/compliance-automation/ml/distill", nil, body)
	if err != nil {
		return nil, err
	}
	var out MlDistillJobResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance-automation/ml/distill", Err: err}
	}
	return &out, nil
}
