// ANNOTATION BLOCK
// File: svc_capabilities.go — the `Capabilities` designation of the Audit Service.
// Why it exists: typed methods for every Capabilities endpoint, reached via
// c.Capabilities.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
)

// GetCapabilities calls GET /capabilities.
// Get Capabilities
//
// Feature-flag snapshot for the current deployment. `copilot_mode` mirrors
// the copilot's own gate (`_llm_mode_available`): "llm" only when the judge
// flag is on AND an API key is configured, otherwise "templates".
// `template_names` lists the copilot's curated guided-question templates
// (the keyword-matched ones; the catch-all fallback is excluded) so the UI
// can render them as suggestions.
func (s *CapabilitiesClient) GetCapabilities(ctx context.Context) (*CapabilitiesResponse, error) {
	b, err := s.c.do(ctx, "GET", "/capabilities", nil, nil)
	if err != nil {
		return nil, err
	}
	var out CapabilitiesResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/capabilities", Err: err}
	}
	return &out, nil
}
