// ANNOTATION BLOCK
// File: svc_search.go — the `Search` designation of the Audit Service.
// Why it exists: typed methods for every Search endpoint, reached via
// c.Search.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// SearchGetSearchFacetsParams holds the query parameters for Search.GetSearchFacets.
type SearchGetSearchFacetsParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // Optional: Only count events after this time
	EndTime   *string `json:"end_time,omitempty"`   // Optional: Only count events before this time
}

func (p *SearchGetSearchFacetsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.StartTime != nil {
		v.Set("start_time", *p.StartTime)
	}
	if p.EndTime != nil {
		v.Set("end_time", *p.EndTime)
	}
	return v
}

// SearchGetSearchSuggestionsParams holds the query parameters for Search.GetSearchSuggestions.
type SearchGetSearchSuggestionsParams struct {
	OrgID  *string `json:"org_id,omitempty"` // Organization ID
	Field  *string `json:"field,omitempty"`  // required. Field name (event_type, action, outcome, user_id, resource_type, service)
	Prefix *string `json:"prefix,omitempty"` // Text prefix to match (e.g. 'user.' finds 'user.login', 'user.logout')
	Limit  *int    `json:"limit,omitempty"`  // Max suggestions to return
}

func (p *SearchGetSearchSuggestionsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.Field != nil {
		v.Set("field", *p.Field)
	}
	if p.Prefix != nil {
		v.Set("prefix", *p.Prefix)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// SearchGetToolInventoryParams holds the query parameters for Search.GetToolInventory.
type SearchGetToolInventoryParams struct {
	OrgID     *string `json:"org_id,omitempty"`     // Organization ID
	StartTime *string `json:"start_time,omitempty"` // Only count tool calls after this time
	EndTime   *string `json:"end_time,omitempty"`   // Only count tool calls before this time
	Limit     *int    `json:"limit,omitempty"`      // Max number of tools to return
}

func (p *SearchGetToolInventoryParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	if p.StartTime != nil {
		v.Set("start_time", *p.StartTime)
	}
	if p.EndTime != nil {
		v.Set("end_time", *p.EndTime)
	}
	if p.Limit != nil {
		v.Set("limit", strconv.Itoa(*p.Limit))
	}
	return v
}

// GetSearchFacets calls GET /search/facets.
// Get Search Facets
//
// Get breakdown of available filter values with counts. Discover what data
// exists. **When to use**: Building filter UI, understanding data
// distribution, exploring available options. **Returns**: Top values for
// event_types, actions, outcomes, services, and resource_types with counts.
// **Next step**: Use these values as filters in POST /search/ or GET
// /logs/events. **Tip**: Shows most common values first. Use time filters
// to see facets for specific period.
func (s *SearchClient) GetSearchFacets(ctx context.Context, params *SearchGetSearchFacetsParams) (*SearchFacetsResponse, error) {
	b, err := s.c.do(ctx, "GET", "/search/facets", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out SearchFacetsResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/search/facets", Err: err}
	}
	return &out, nil
}

// GetSearchSuggestions calls GET /search/suggest.
// Get Search Suggestions
//
// Get autocomplete suggestions for search filters. Helps discover available
// values. **When to use**: Building search UI, discovering what event_types
// exist, finding valid filter values. **Returns**: List of matching values
// with frequency count, ordered by popularity. **Next step**: Use suggested
// values in POST /search/ filters or GET /logs/events queries. **Tip**:
// Leave prefix empty to see all values for a field, sorted by frequency.
func (s *SearchClient) GetSearchSuggestions(ctx context.Context, params *SearchGetSearchSuggestionsParams) (*SearchSuggestResponse, error) {
	b, err := s.c.do(ctx, "GET", "/search/suggest", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out SearchSuggestResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/search/suggest", Err: err}
	}
	return &out, nil
}

// GetToolInventory calls GET /search/tools.
// Get Tool Inventory
//
// Per-tool inventory of the AI tools your agents actually called. **When to
// use**: Powers the "Tool inventory" page — answers "which tools did our
// agents call, how often, by how many distinct callers, and how often did
// they fail". **Returns**: One row per real tool (the tool-call event's
// `resource_id`), with total calls, distinct callers, and success/failure
// split. Org-scoped and time-bounded like the search facets. **How a tool
// is identified**: events where `resource_type='tool'` or whose
// `event_type` names a tool/integration/function/mcp call.
func (s *SearchClient) GetToolInventory(ctx context.Context, params *SearchGetToolInventoryParams) (*ToolInventoryResponse, error) {
	b, err := s.c.do(ctx, "GET", "/search/tools", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ToolInventoryResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/search/tools", Err: err}
	}
	return &out, nil
}

// SearchAuditEvents calls POST /search/.
// Search Audit Events
//
// Powerful search with full-text query and multiple filters. Find events by
// content, not just metadata. **When to use**: Searching for specific text
// in request/response data, investigating incidents, finding related
// events. **Query field**: Searches across request_data, response_data,
// metadata, user_agent, and resource_id. **Filters**: Combine query text
// with event_types, actions, outcomes, user_ids, services, and time range.
// **Returns**: Matching events with total count and execution time.
// Supports pagination. **Next step**: Refine search with filters or get
// details with GET /logs/events/{event_id}. **Tip**: Start broad, then add
// filters to narrow results.
func (s *SearchClient) SearchAuditEvents(ctx context.Context, body SearchQuery) (*SearchResult, error) {
	b, err := s.c.do(ctx, "POST", "/search/", nil, body)
	if err != nil {
		return nil, err
	}
	var out SearchResult
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/search/", Err: err}
	}
	return &out, nil
}
