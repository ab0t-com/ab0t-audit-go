// ANNOTATION BLOCK
// File: svc_auditorshare.go — the `AuditorShare` designation of the Audit Service.
// Why it exists: typed methods for every AuditorShare endpoint, reached via
// c.AuditorShare.<Method>. Each is context-first, takes typed path/query/body
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

// AuditorShareListEvidencePackagesParams holds the query parameters for AuditorShare.ListEvidencePackages.
type AuditorShareListEvidencePackagesParams struct {
	Framework *string `json:"framework,omitempty"`
	Sealed    *bool   `json:"sealed,omitempty"`
	Limit     *int    `json:"limit,omitempty"`
	Offset    *int    `json:"offset,omitempty"`
}

func (p *AuditorShareListEvidencePackagesParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.Framework != nil {
		v.Set("framework", *p.Framework)
	}
	if p.Sealed != nil {
		v.Set("sealed", strconv.FormatBool(*p.Sealed))
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	if p.Offset != nil {
		v.Set("offset", strconv.Itoa(*p.Offset))
	}
	return v
}

// AuditorShareView calls GET /auditor/share/{token}.
// Auditor Share View
//
// External-auditor read endpoint. The token in the URL is the auth artefact
// — no JWT, no cookie. Returns 410 Gone if revoked or past expiry, 404 if
// token unknown. Records every access as an audit event (meta-audit
// surface).
func (s *AuditorShareClient) AuditorShareView(ctx context.Context, token string) (*AuditorViewPayloadWithLineage, error) {
	path := fmt.Sprintf("/auditor/share/%s", url.PathEscape(token))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out AuditorViewPayloadWithLineage
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// CreateEvidencePackage calls POST /compliance/evidence-packages.
// Create Evidence Package
//
// Create an evidence package. T8.minimal — seal-on-create only; un-seal /
// edit deferred.
func (s *AuditorShareClient) CreateEvidencePackage(ctx context.Context, body EvidencePackageCreate) (*EvidencePackage, error) {
	b, err := s.c.do(ctx, "POST", "/compliance/evidence-packages", nil, body)
	if err != nil {
		return nil, err
	}
	var out EvidencePackage
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/evidence-packages", Err: err}
	}
	return &out, nil
}

// GetEvidencePackage calls GET /compliance/evidence-packages/{package_id}.
// Get Evidence Package
//
// Read a single evidence package within the caller's org. Carries supersede
// lineage: “superseded_by“ (this package was replaced) and
// “supersedes_package_id“ (this package is itself a correction of
// another).
func (s *AuditorShareClient) GetEvidencePackage(ctx context.Context, packageID string) (*EvidencePackageWithLineage, error) {
	path := fmt.Sprintf("/compliance/evidence-packages/%s", url.PathEscape(packageID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out EvidencePackageWithLineage
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ListEvidencePackages calls GET /compliance/evidence-packages.
// List Evidence Packages
//
// List evidence packages in the caller's org. Backs the Evidence Vault
// page. **When to use**: rendering the /ui/app/compliance/evidence page.
// **Filters**: framework (SOC2/GDPR/...), sealed (true/false). All
// optional. **Auth**: AuditComplianceReader.
func (s *AuditorShareClient) ListEvidencePackages(ctx context.Context, params *AuditorShareListEvidencePackagesParams) (*PaginatedResponseEvidencePackageWithLineage, error) {
	b, err := s.c.do(ctx, "GET", "/compliance/evidence-packages", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out PaginatedResponseEvidencePackageWithLineage
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/compliance/evidence-packages", Err: err}
	}
	return &out, nil
}

// ListShareLinks calls GET /compliance/evidence-packages/{package_id}/share-links.
// List Share Links
//
// List the share links minted for an evidence package. Backs the "Manage
// share links" drawer on the Evidence Vault page so a manager can see who
// holds external-auditor access, when each link expires (or was revoked),
// how many times it's been opened, and revoke any link. Perm-aligned with
// mint/revoke (AuditEvidenceManager), org-scoped. Returns links
// newest-first. `status` is one of active/expired/revoked.
func (s *AuditorShareClient) ListShareLinks(ctx context.Context, packageID string) ([]ShareLinkSummary, error) {
	path := fmt.Sprintf("/compliance/evidence-packages/%s/share-links", url.PathEscape(packageID))
	b, err := s.c.do(ctx, "GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out []ShareLinkSummary
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return out, nil
}

// MintShareLink calls POST /compliance/evidence-packages/{package_id}/share.
// Mint Share Link
//
// Mint a time-bound share token for an evidence package.
func (s *AuditorShareClient) MintShareLink(ctx context.Context, packageID string, body ShareLinkCreate) (*ShareLink, error) {
	path := fmt.Sprintf("/compliance/evidence-packages/%s/share", url.PathEscape(packageID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out ShareLink
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// RevokeShareLink calls DELETE /compliance/share-links/{token}.
// Revoke Share Link
//
// Revoke a share link. Subsequent /auditor/share/{token} accesses return
// 410.
func (s *AuditorShareClient) RevokeShareLink(ctx context.Context, token string) error {
	path := fmt.Sprintf("/compliance/share-links/%s", url.PathEscape(token))
	_, err := s.c.do(ctx, "DELETE", path, nil, nil)
	return err
}

// SupersedeEvidencePackage calls POST /compliance/evidence-packages/{package_id}/supersede.
// Supersede Evidence Package
//
// Supersede a sealed evidence package with a corrected revision (PMM F7).
// An evidence product cannot un-seal or edit a sealed package — that would
// destroy its integrity guarantee. Instead this creates a **new** package
// (fresh id, always re-sealed, its own digest) carrying the corrected
// content, and records an append-only supersession link. The original is
// left **byte-for-byte untouched**; readers resolve a forward "superseded
// by <id>" pointer from the link so no auditor silently reads stale
// evidence. * Auth: “AuditEvidenceManager“ (same alias as package
// creation), org-scoped; the in-org existence check below is the Phase-2
// ownership guard. * A package may be superseded once (supersede the HEAD,
// not a stale one): a second attempt returns **409**. * Cross-org / unknown
// ids return **404** (no existence disclosure).
func (s *AuditorShareClient) SupersedeEvidencePackage(ctx context.Context, packageID string, body SupersedeRequest) (*EvidencePackageWithLineage, error) {
	path := fmt.Sprintf("/compliance/evidence-packages/%s/supersede", url.PathEscape(packageID))
	b, err := s.c.do(ctx, "POST", path, nil, body)
	if err != nil {
		return nil, err
	}
	var out EvidencePackageWithLineage
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}
