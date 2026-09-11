// ANNOTATION BLOCK
// File: svc_compliancecopilot.go — the `ComplianceCopilot` designation of the Audit Service.
// Why it exists: typed methods for every ComplianceCopilot endpoint, reached via
// c.ComplianceCopilot.<Method>. Each is context-first, takes typed path/query/body
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

// ComplianceCopilotGetCopilotConversationParams holds the query parameters for ComplianceCopilot.GetCopilotConversation.
type ComplianceCopilotGetCopilotConversationParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *ComplianceCopilotGetCopilotConversationParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceCopilotListCopilotConversationsParams holds the query parameters for ComplianceCopilot.ListCopilotConversations.
type ComplianceCopilotListCopilotConversationsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
	Limit *int    `json:"limit,omitempty"`
}

func (p *ComplianceCopilotListCopilotConversationsParams) values() url.Values {
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

// CopilotAsk calls POST /compliance/copilot/ask.
// Copilot Ask
//
// Ask the compliance copilot a question. **Modes**: when
// `COMPLIANCE_LLM_JUDGE_ENABLED=true` and an API key is set
// (`COMPLIANCE_LLM_JUDGE_API_KEY`), `llm` mode asks the configured LLM —
// via the same `llm_provider_invoker` the compliance judge uses — to
// generate a single SELECT for the question, then runs it through the SAME
// safety gate as templates. Otherwise (or on any LLM error / unsafe SQL) it
// falls back to `template` mode: a curated keyword-matched SQL catalog.
// When the question did not hit a specific template and no LLM was
// available, the response is flagged `used_fallback=true` so the user knows
// the answer is a canned template, not a literal understanding of their
// exact question. **Safety**: blacklist the question, generate/pick SQL,
// force the org predicate, cap the LIMIT, allowlist-validate the SQL.
// SELECT-only. The safety gate is identical for LLM-generated and template
// SQL. **Audit**: every call writes an `audit.copilot.query` event to the
// queryable `audit_events` table — regardless of outcome (success / blocked
// / failed) — so copilot usage is itself auditable via `/search/`, the
// Meta-audit page, exports, and compliance reports.
func (s *ComplianceCopilotClient) CopilotAsk(ctx context.Context, body CopilotAsk) (*CopilotAnswer, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/copilot/ask", nil, body)
	if err != nil {
		return nil, err
	}
	var out CopilotAnswer
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/copilot/ask", Err: err}
	}
	return &out, nil
}

// GetCopilotConversation calls GET /compliance/copilot/conversations/{conversation_id}.
// Get Copilot Conversation
//
// Replay the full transcript for one copilot conversation. Returns every
// turn in order, including the assistant turn's SQL and cached rows so the
// client can re-render without re-executing the queries. Scoped to (org_id,
// user_id) — you can only see your own sessions.
func (s *ComplianceCopilotClient) GetCopilotConversation(ctx context.Context, conversationID string, params *ComplianceCopilotGetCopilotConversationParams) (RawMessage, error) {
	path := fmt.Sprintf("/compliance/copilot/conversations/%s", url.PathEscape(conversationID))
	b, err := s.c.do(ctx, "GET", path, params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// ListCopilotConversations calls GET /compliance/copilot/conversations.
// List Copilot Conversations
//
// List the caller's recent copilot conversations. Returns one row per
// conversation with the first user message as the title, the turn count,
// and the most-recent turn timestamp. Newest first. Used by the sidebar on
// the copilot page to let users resume past sessions.
func (s *ComplianceCopilotClient) ListCopilotConversations(ctx context.Context, params *ComplianceCopilotListCopilotConversationsParams) (RawMessage, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/copilot/conversations", params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}
