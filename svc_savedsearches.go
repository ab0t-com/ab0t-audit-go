// ANNOTATION BLOCK
// File: svc_savedsearches.go — the `SavedSearches` designation of the Audit Service.
// Why it exists: typed methods for every SavedSearches endpoint, reached via
// c.SavedSearches.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateSavedSearch calls POST /saved-searches/.
// Create Saved Search
//
// Persist a new saved search. Idempotent on (user_id, name): re-saving the
// same name replaces the previous row's filter_json + bumps updated_at. New
// names always get a fresh search_id.
func (s *SavedSearchesClient) CreateSavedSearch(ctx context.Context, body SavedSearchCreate) (*SavedSearchItem, error) {
	b, err := s.c.do(ctx, "POST", "/saved-searches/", nil, body)
	if err != nil {
		return nil, err
	}
	var out SavedSearchItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/saved-searches/", Err: err}
	}
	return &out, nil
}

// DeleteSavedSearch calls DELETE /saved-searches/{search_id}.
// Delete Saved Search
//
// Soft-delete a saved search. Inserts a tombstone row so RMT FINAL excludes
// it. Idempotent — deleting an already-deleted row is a no-op. Only the
// owning user can delete. 404 (not 403) is returned for cross-user attempts
// so we don't leak existence.
func (s *SavedSearchesClient) DeleteSavedSearch(ctx context.Context, searchID string) (*SavedSearchDeleteResponse, error) {
	path := fmt.Sprintf("/saved-searches/%s", url.PathEscape(searchID))
	b, err := s.c.do(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out SavedSearchDeleteResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ListSavedSearches calls GET /saved-searches/.
// List Saved Searches
//
// Return the caller's saved searches, newest-first. Soft-deleted rows
// excluded.
func (s *SavedSearchesClient) ListSavedSearches(ctx context.Context) (*ListEnvelopeSavedSearchItem, error) {
	b, err := s.c.do(ctx, "GET", "/saved-searches/", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeSavedSearchItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/saved-searches/", Err: err}
	}
	return &out, nil
}
