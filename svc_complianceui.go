// ANNOTATION BLOCK
// File: svc_complianceui.go — the `ComplianceUI` designation of the Audit Service.
// Why it exists: typed methods for every ComplianceUI endpoint, reached via
// c.ComplianceUI.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"net/url"
)

// ComplianceUIComplianceDashboardUIParams holds the query parameters for ComplianceUI.ComplianceDashboardUI.
type ComplianceUIComplianceDashboardUIParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID (optional)
}

func (p *ComplianceUIComplianceDashboardUIParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ComplianceDashboardUI calls GET /ui/compliance/dashboard.
// Compliance Dashboard Ui
//
// Render the real-time compliance dashboard shell.
func (s *ComplianceUIClient) ComplianceDashboardUI(ctx context.Context, params *ComplianceUIComplianceDashboardUIParams) (string, error) {
	b, err := s.c.do(ctx, "GET", "/ui/compliance/dashboard", params.values(), nil)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
