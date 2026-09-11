// ANNOTATION BLOCK
// File: svc_compliancetemplates.go — the `ComplianceTemplates` designation of the Audit Service.
// Why it exists: typed methods for every ComplianceTemplates endpoint, reached via
// c.ComplianceTemplates.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// ComplianceTemplatesCloneTemplateParams holds the query parameters for ComplianceTemplates.CloneTemplate.
type ComplianceTemplatesCloneTemplateParams struct {
	OrgID *string `json:"org_id,omitempty"`
}

func (p *ComplianceTemplatesCloneTemplateParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// CloneTemplate calls POST /compliance/templates/{template_id}/clone.
// Clone Template
//
// Render a template with provided `vars`. Missing placeholders are left as
// literal `{{ name }}` markers so the cloner sees what they still need to
// fill in.
func (s *ComplianceTemplatesClient) CloneTemplate(ctx context.Context, templateID string, body TemplateCloneBody, params *ComplianceTemplatesCloneTemplateParams) (*ComplianceTemplateCloneResponse, error) {
	path := fmt.Sprintf("/compliance/templates/%s/clone", url.PathEscape(templateID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	var out ComplianceTemplateCloneResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ListTemplates calls GET /compliance/templates/.
// List Templates
//
// Return the curated template catalog with placeholders inferred from each
// file.
func (s *ComplianceTemplatesClient) ListTemplates(ctx context.Context) (*ListEnvelopeComplianceTemplateItem, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/templates/", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeComplianceTemplateItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/templates/", Err: err}
	}
	return &out, nil
}
