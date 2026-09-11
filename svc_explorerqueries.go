// ANNOTATION BLOCK
// File: svc_explorerqueries.go — the `ExplorerQueries` designation of the Audit Service.
// Why it exists: typed methods for every ExplorerQueries endpoint, reached via
// c.ExplorerQueries.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateSavedQuery calls POST /explorer-queries/saved.
// Create Saved Query
//
// Persist a saved Data Explorer query. Idempotent on (user_id, question):
// re-saving the same question replaces the existing row (bumps updated_at,
// updates the name) instead of duplicating it, so the star reflects a
// single stable saved state.
func (s *ExplorerQueriesClient) CreateSavedQuery(ctx context.Context, body ExplorerSavedQueryCreate) (*ExplorerSavedQueryItem, error) {
	b, err := s.c.do(ctx, "POST", "/explorer-queries/saved", nil, body)
	if err != nil {
		return nil, err
	}
	var out ExplorerSavedQueryItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/explorer-queries/saved", Err: err}
	}
	return &out, nil
}

// DeleteSavedQuery calls DELETE /explorer-queries/saved/{query_id}.
// Delete Saved Query
//
// Soft-delete a saved query (tombstone row so RMT FINAL excludes it). Only
// the owning user can delete. 404 (not 403) is returned for cross-user
// attempts so we don't leak existence. Idempotent.
func (s *ExplorerQueriesClient) DeleteSavedQuery(ctx context.Context, queryID string) (*ExplorerSavedQueryDeleteResponse, error) {
	path := fmt.Sprintf("/explorer-queries/saved/%s", url.PathEscape(queryID))
	b, err := s.c.do(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}
	var out ExplorerSavedQueryDeleteResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: path, Err: err}
	}
	return &out, nil
}

// ListQueryHistory calls GET /explorer-queries/history.
// List Query History
//
// Return the caller's most-recent Data Explorer questions, newest-first.
// Deduplicated to the latest run per question text so the panel shows
// distinct recent questions rather than the same question repeated.
func (s *ExplorerQueriesClient) ListQueryHistory(ctx context.Context) (*ListEnvelopeExplorerHistoryItem, error) {
	b, err := s.c.do(ctx, "GET", "/explorer-queries/history", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeExplorerHistoryItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/explorer-queries/history", Err: err}
	}
	return &out, nil
}

// ListSavedQueries calls GET /explorer-queries/saved.
// List Saved Queries
//
// Return the caller's saved Data Explorer queries, newest-first.
func (s *ExplorerQueriesClient) ListSavedQueries(ctx context.Context) (*ListEnvelopeExplorerSavedQueryItem, error) {
	b, err := s.c.do(ctx, "GET", "/explorer-queries/saved", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ListEnvelopeExplorerSavedQueryItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/explorer-queries/saved", Err: err}
	}
	return &out, nil
}

// RecordQueryHistory calls POST /explorer-queries/history.
// Record Query History
//
// Record a question the user just ran in the Data Explorer. Append-only —
// the GET dedups to the latest run per question. Called by the page after a
// successful /search/ so the Recent panel reflects real activity.
func (s *ExplorerQueriesClient) RecordQueryHistory(ctx context.Context, body ExplorerHistoryCreate) (*ExplorerHistoryItem, error) {
	b, err := s.c.do(ctx, "POST", "/explorer-queries/history", nil, body)
	if err != nil {
		return nil, err
	}
	var out ExplorerHistoryItem
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/explorer-queries/history", Err: err}
	}
	return &out, nil
}
