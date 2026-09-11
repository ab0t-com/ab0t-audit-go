// ANNOTATION BLOCK
// File: svc_health.go — the `Health` designation of the Audit Service.
// Why it exists: typed methods for every Health endpoint, reached via
// c.Health.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
)

// DetailedHealthCheck calls GET /health/detailed.
// Detailed Health Check
//
// Comprehensive health check with dependency status. Diagnose service
// issues. **When to use**: Troubleshooting, deployment verification,
// dependency monitoring. **Returns**: Service status plus ClickHouse and
// Redis connection status. 503 if any dependency fails. **Response time**:
// ~100ms. Includes database connectivity tests. **Next step**: If
// unhealthy, check dependency error messages to identify issue.
func (s *HealthClient) DetailedHealthCheck(ctx context.Context) (*DetailedHealthResponse, error) {
	b, err := s.c.do(ctx, "GET", "/health/detailed", nil, nil)
	if err != nil {
		return nil, err
	}
	var out DetailedHealthResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/health/detailed", Err: err}
	}
	return &out, nil
}

// HealthCheck calls GET /health.
// Health Check
//
// Quick health check. Verify service is running and responsive. **When to
// use**: Load balancer checks, uptime monitoring, basic service
// verification. **Returns**: Always 200 OK if service is up. Includes
// timestamp. **Response time**: <10ms. No database queries. **Next step**:
// If unhealthy, check GET /health/detailed for dependency status.
func (s *HealthClient) HealthCheck(ctx context.Context) (*HealthResponse, error) {
	b, err := s.c.do(ctx, "GET", "/health", nil, nil)
	if err != nil {
		return nil, err
	}
	var out HealthResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/health", Err: err}
	}
	return &out, nil
}

// LivenessCheck calls GET /liveness.
// Liveness Check
//
// Kubernetes liveness probe. Indicates service process is alive (not
// deadlocked/crashed). **When to use**: Kubernetes/container orchestration
// liveness checks to detect need for restart. **Returns**: Always 200
// {"status": "alive"} if service process is responsive. **Purpose**:
// Triggers pod restart if service becomes unresponsive.
func (s *HealthClient) LivenessCheck(ctx context.Context) (*LivenessResponse, error) {
	b, err := s.c.do(ctx, "GET", "/liveness", nil, nil)
	if err != nil {
		return nil, err
	}
	var out LivenessResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/liveness", Err: err}
	}
	return &out, nil
}

// ReadinessCheck calls GET /readiness.
// Readiness Check
//
// Kubernetes readiness probe. Indicates service is ready to accept traffic.
// **When to use**: Kubernetes/container orchestration readiness checks
// during deployment. **Returns**: Always 200 {"status": "ready"} when
// service initialization is complete. **Purpose**: Prevents routing traffic
// to service before it's ready to handle requests.
func (s *HealthClient) ReadinessCheck(ctx context.Context) (*ReadinessResponse, error) {
	b, err := s.c.do(ctx, "GET", "/readiness", nil, nil)
	if err != nil {
		return nil, err
	}
	var out ReadinessResponse
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/readiness", Err: err}
	}
	return &out, nil
}
