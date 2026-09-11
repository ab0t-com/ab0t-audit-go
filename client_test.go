// ANNOTATION BLOCK
// File: client_test.go — httptest-backed wiring tests for a representative
// slice of the surface (ingest, a compliance GET, a search POST, a paginated
// list + CollectAll, query-param encoding, auth headers, and the error path).
// Why it exists: proves the do() plumbing, path/query/body assembly, typed
// decode, sub-client wiring and *APIError handling — not all 249 methods.
package auditclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestClient spins up an httptest server and points a Client at it.
func newTestClient(t *testing.T, h http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	return New(srv.URL, WithAPIKey("ab0t_sk_test"), WithHTTPClient(srv.Client()))
}

func TestLegacyIngest(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/logs/ingest" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer ab0t_sk_test" {
			t.Errorf("auth header = %q", got)
		}
		var ev Event
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			t.Fatal(err)
		}
		if ev.Action != "tool.call" {
			t.Errorf("action = %q", ev.Action)
		}
		w.WriteHeader(http.StatusAccepted)
	})
	if err := c.Ingest(context.Background(), Event{Action: "tool.call", EventType: "tools/call", Service: "mcp-gateway"}); err != nil {
		t.Fatal(err)
	}
}

func TestLogsIngestAuditEvent(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/logs/ingest" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if got := r.Header.Get("X-API-Key"); got != "ab0t_sk_test" {
			t.Errorf("x-api-key = %q", got)
		}
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(IngestLogResponse{Status: "accepted", EventID: "evt_1", Message: "ok"})
	})
	resp, err := c.Logs.IngestAuditEvent(context.Background(), AuditEventCreate{
		Action: "a", EventType: "e", Service: "s",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.EventID != "evt_1" || resp.Status != "accepted" {
		t.Errorf("resp = %+v", resp)
	}
}

func TestComplianceDashboardGET(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/compliance/dashboard" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("org_id") != "org_123" || q.Get("framework") != "soc2" {
			t.Errorf("query = %v", q)
		}
		_ = json.NewEncoder(w).Encode(ComplianceDashboardResponse{
			OrgID: "org_123", ComplianceScore: 92, AuditReadiness: "ready", TotalControls: 40,
		})
	})
	dash, err := c.Compliance.GetComplianceDashboard(context.Background(), &ComplianceGetComplianceDashboardParams{
		OrgID:     String("org_123"),
		Framework: String("soc2"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if dash.ComplianceScore != 92 || dash.AuditReadiness != "ready" {
		t.Errorf("dash = %+v", dash)
	}
}

func TestSearchPOST(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/search/" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		var q SearchQuery
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &q); err != nil {
			t.Fatal(err)
		}
		if q.Query != "login" || len(q.Services) != 1 || q.Services[0] != "auth-service" {
			t.Errorf("decoded query = %+v", q)
		}
		_ = json.NewEncoder(w).Encode(SearchResult{
			Data:  []AuditEvent{{EventID: "e1", Action: "user.login", EventType: "auth", Service: "auth-service"}},
			Total: 1, Limit: 50, Query: "login",
		})
	})
	res, err := c.Search.SearchAuditEvents(context.Background(), SearchQuery{
		Query:    "login",
		Services: []string{"auth-service"},
		Limit:    Int(50),
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.Total != 1 || len(res.Data) != 1 || res.Data[0].Action != "user.login" {
		t.Errorf("res = %+v", res)
	}
}

func TestPaginatedListAndCollectAll(t *testing.T) {
	// Two pages of 2, then done.
	pages := map[string]PaginatedResponseAuditEvent{
		"0": {Data: []AuditEvent{{EventID: "a"}, {EventID: "b"}}, Total: 3, Limit: 2, HasMore: true},
		"2": {Data: []AuditEvent{{EventID: "c"}}, Total: 3, Limit: 2, HasMore: false},
	}
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/logs/events" {
			t.Errorf("path = %s", r.URL.Path)
		}
		off := r.URL.Query().Get("offset")
		if off == "" {
			off = "0"
		}
		p, ok := pages[off]
		if !ok {
			t.Fatalf("unexpected offset %q", off)
		}
		_ = json.NewEncoder(w).Encode(p)
	})

	// single page call
	first, err := c.Logs.GetAuditEvents(context.Background(), &LogsGetAuditEventsParams{Limit: Int(2), Offset: Int(0)})
	if err != nil {
		t.Fatal(err)
	}
	if !first.HasMore || len(first.Data) != 2 {
		t.Errorf("first page = %+v", first)
	}

	// drive to completion
	all, err := CollectAll(context.Background(), func(ctx context.Context, offset int) (Page[AuditEvent], error) {
		resp, err := c.Logs.GetAuditEvents(ctx, &LogsGetAuditEventsParams{Limit: Int(2), Offset: Int(offset)})
		if err != nil {
			return Page[AuditEvent]{}, err
		}
		return Page[AuditEvent]{Data: resp.Data, Total: resp.Total, HasMore: resp.HasMore}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 3 {
		t.Fatalf("collected %d events, want 3", len(all))
	}
	if all[0].EventID != "a" || all[2].EventID != "c" {
		t.Errorf("collected order wrong: %+v", all)
	}
}

func TestPathParamEncoding(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/logs/events/evt%2Fweird" && r.URL.EscapedPath() != "/logs/events/evt%2Fweird" {
			t.Errorf("escaped path = %s", r.URL.EscapedPath())
		}
		_ = json.NewEncoder(w).Encode(AuditEvent{EventID: "evt/weird", Action: "x", EventType: "y", Service: "z"})
	})
	ev, err := c.Logs.GetAuditEvent(context.Background(), "evt/weird", nil)
	if err != nil {
		t.Fatal(err)
	}
	if ev.EventID != "evt/weird" {
		t.Errorf("event = %+v", ev)
	}
}

func TestErrorPath(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "req-abc")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"detail":"event not found"}`))
	})
	_, err := c.Logs.GetAuditEvent(context.Background(), "missing", nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsNotFound(err) {
		t.Errorf("IsNotFound = false for %v", err)
	}
	ae, ok := AsAPIError(err)
	if !ok {
		t.Fatalf("not an *APIError: %v", err)
	}
	if ae.StatusCode != 404 || ae.Message != "event not found" || ae.RequestID != "req-abc" {
		t.Errorf("APIError = %+v", ae)
	}
	if IsRetryable(err) {
		t.Error("404 should not be retryable")
	}
}

func TestValidationErrorParsing(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"detail":[{"msg":"field required","type":"value_error.missing"}]}`))
	})
	_, err := c.Logs.IngestAuditEvent(context.Background(), AuditEventCreate{})
	if !IsValidationError(err) {
		t.Fatalf("expected 422, got %v", err)
	}
	ae, _ := AsAPIError(err)
	if ae.Message != "field required" || ae.Code != "value_error.missing" {
		t.Errorf("parsed = code=%q msg=%q", ae.Code, ae.Message)
	}
}

func TestSubClientsWired(t *testing.T) {
	c := New("")
	if c.Logs == nil || c.Compliance == nil || c.Search == nil || c.Admin == nil ||
		c.Export == nil || c.LegalHold == nil || c.ComplianceAutomation == nil || c.UI == nil {
		t.Fatal("sub-clients not initialized by New")
	}
}
