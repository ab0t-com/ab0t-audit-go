// ANNOTATION BLOCK
// File: svc_legalhold.go — the `LegalHold` designation of the Audit Service.
// Why it exists: typed methods for every LegalHold endpoint, reached via
// c.LegalHold.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"fmt"
	"net/url"
)

// LegalHoldListLegalHoldsParams holds the query parameters for LegalHold.ListLegalHolds.
type LegalHoldListLegalHoldsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *LegalHoldListLegalHoldsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// LegalHoldPlaceLegalHoldParams holds the query parameters for LegalHold.PlaceLegalHold.
type LegalHoldPlaceLegalHoldParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *LegalHoldPlaceLegalHoldParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// LegalHoldReleaseLegalHoldParams holds the query parameters for LegalHold.ReleaseLegalHold.
type LegalHoldReleaseLegalHoldParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *LegalHoldReleaseLegalHoldParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// ListLegalHolds calls GET /legal-holds.
// List Legal Holds
//
// List active holds and the full hold history (active + released) for the
// org.
func (s *LegalHoldClient) ListLegalHolds(ctx context.Context, params *LegalHoldListLegalHoldsParams) (RawMessage, error) {
	b, err := s.c.do(ctx, "GET", "/legal-holds", params.values(), nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// PlaceLegalHold calls POST /legal-holds.
// Place Legal Hold
//
// Place a legal hold. Blocks deletion of matching audit evidence until
// released. **Scopes**: `org` (whole org), `data_subject` (one subject),
// `before_date` (everything at or before a cutoff). **Effect**: purge,
// retention TTL, and GDPR erasure refuse to delete data the hold protects,
// returning 409 + an audited denial.
func (s *LegalHoldClient) PlaceLegalHold(ctx context.Context, body LegalHoldCreate, params *LegalHoldPlaceLegalHoldParams) (RawMessage, error) {
	b, err := s.c.do(ctx, "POST", "/legal-holds", params.values(), body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// ReleaseLegalHold calls POST /legal-holds/{hold_id}/release.
// Release Legal Hold
//
// Release a legal hold. Appends a superseding `released` row (append-only,
// never UPDATE/DELETE). Once released, matching data may be deleted again.
func (s *LegalHoldClient) ReleaseLegalHold(ctx context.Context, holdID string, body LegalHoldRelease, params *LegalHoldReleaseLegalHoldParams) (RawMessage, error) {
	path := fmt.Sprintf("/legal-holds/%s/release", url.PathEscape(holdID))
	b, err := s.c.do(ctx, "POST", path, params.values(), body)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}
