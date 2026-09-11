// ANNOTATION BLOCK
// File: misc.go — root/untagged endpoints hung directly off *Client
// (service root, Prometheus metrics, llm.txt). Same do() plumbing as the
// sub-clients; these simply have no designation to group under.
package auditclient

import (
	"context"
	"encoding/json"
)

// LlmTxt calls GET /llm.txt.
// Llm Txt
//
// LLM-optimized API documentation. Compressed, grepable, intent-driven
// guide for AI agents. **For LLM/AI agents**: This file is optimized for
// your consumption. Search by: - grep "INTENT:" - Find by what you want to
// do - grep "ROLE:" - Find by who you are - grep "PROBLEM:" - Find by issue
// you're solving - grep "WORKFLOW:" - Find multi-step processes **Format**:
// Plain text, highly structured, information-dense, layered for progressive
// discovery.
func (c *Client) LlmTxt(ctx context.Context) (RawMessage, error) {
	b, err := c.do(ctx, "GET", "/llm.txt", nil, nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// Metrics calls GET /metrics.
//
// Prometheus metrics endpoint
func (c *Client) Metrics(ctx context.Context) (RawMessage, error) {
	b, err := c.do(ctx, "GET", "/metrics", nil, nil)
	if err != nil {
		return nil, err
	}
	return RawMessage(b), nil
}

// Root calls GET /.
//
// Root endpoint
func (c *Client) Root(ctx context.Context) (*ServiceRootResponse, error) {
	b, err := c.do(ctx, "GET", "/", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ServiceRootResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/", Err: err}
	}
	return &out, nil
}
