// ANNOTATION BLOCK
// File: errors.go — typed errors for the Audit Service client.
// Why it exists: every non-2xx response surfaces as a single *APIError carrying
// the HTTP status, the method+path, a best-effort machine code/message parsed
// from FastAPI's {"detail": ...} envelope, and the raw body. Callers branch on
// the Is* helpers (IsNotFound, IsValidationError, IsRateLimited, ...) rather
// than comparing StatusCode by hand. EncodeError/DecodeError wrap request-body
// marshal and response-body unmarshal failures so they are distinguishable from
// transport and API errors.
package auditclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// APIError is returned for any non-2xx response from the audit service. It
// captures the HTTP status, the endpoint that produced it, a best-effort
// machine-readable code/message parsed from the response body, and the raw body
// for diagnostics.
//
// Prefer the Is* helpers (IsUnauthorized, IsNotFound, ...) over comparing
// StatusCode directly.
type APIError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int
	// Method is the HTTP method of the originating request.
	Method string
	// Endpoint is the request path (no query string) that produced the error.
	Endpoint string
	// Code is a machine-readable error code parsed from the body, if any.
	Code string
	// Message is a human-readable message parsed from the body, if any.
	Message string
	// RequestID echoes any X-Request-ID returned by the service.
	RequestID string
	// Body is the raw (possibly truncated) response body.
	Body string
}

func (e *APIError) Error() string {
	var b strings.Builder
	b.WriteString("auditclient: ")
	if e.Method != "" {
		b.WriteString(e.Method)
		b.WriteByte(' ')
	}
	b.WriteString(e.Endpoint)
	b.WriteString(": ")
	fmt.Fprintf(&b, "status %d %s", e.StatusCode, http.StatusText(e.StatusCode))
	if e.Code != "" {
		b.WriteString(" [")
		b.WriteString(e.Code)
		b.WriteByte(']')
	}
	if e.Message != "" {
		b.WriteString(": ")
		b.WriteString(e.Message)
	} else if e.Body != "" {
		b.WriteString(": ")
		b.WriteString(truncate(e.Body, 256))
	}
	if e.RequestID != "" {
		b.WriteString(" (request_id=")
		b.WriteString(e.RequestID)
		b.WriteByte(')')
	}
	return b.String()
}

// EncodeError wraps a failure to JSON-encode a request body.
type EncodeError struct {
	Endpoint string
	Err      error
}

func (e *EncodeError) Error() string {
	return fmt.Sprintf("auditclient: encode request for %s: %v", e.Endpoint, e.Err)
}
func (e *EncodeError) Unwrap() error { return e.Err }

// DecodeError wraps a failure to JSON-decode a (2xx) response body into the
// method's typed return value.
type DecodeError struct {
	Endpoint string
	Err      error
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("auditclient: decode response from %s: %v", e.Endpoint, e.Err)
}
func (e *DecodeError) Unwrap() error { return e.Err }

// parseAPIError builds an APIError, extracting a code/message from the common
// FastAPI {"detail": ...} envelope as well as {"code","message"} shapes.
func parseAPIError(status int, method, endpoint, requestID, body string) *APIError {
	e := &APIError{
		StatusCode: status,
		Method:     method,
		Endpoint:   endpoint,
		RequestID:  requestID,
		Body:       body,
	}
	e.Code, e.Message = extractCodeMessage(body)
	return e
}

func extractCodeMessage(body string) (code, msg string) {
	body = strings.TrimSpace(body)
	if body == "" || body[0] != '{' {
		return "", ""
	}
	var env struct {
		Error   string          `json:"error"`
		Code    string          `json:"code"`
		Message string          `json:"message"`
		Detail  json.RawMessage `json:"detail"`
	}
	if err := json.Unmarshal([]byte(body), &env); err != nil {
		return "", ""
	}
	code = env.Code
	if code == "" {
		code = env.Error
	}
	msg = env.Message
	// FastAPI: {"detail": "..."} or {"detail": [{"msg": ..., "type": ...}]}.
	if len(env.Detail) > 0 {
		var ds string
		if json.Unmarshal(env.Detail, &ds) == nil {
			if msg == "" {
				msg = ds
			}
		} else {
			var items []struct {
				Msg  string `json:"msg"`
				Type string `json:"type"`
			}
			if json.Unmarshal(env.Detail, &items) == nil && len(items) > 0 {
				if msg == "" {
					msg = items[0].Msg
				}
				if code == "" {
					code = items[0].Type
				}
			}
		}
	}
	return code, msg
}

// AsAPIError returns the underlying *APIError if err wraps one.
func AsAPIError(err error) (*APIError, bool) {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}

// StatusCode returns the HTTP status carried by err, or 0 if err is not an
// APIError.
func StatusCode(err error) int {
	if ae, ok := AsAPIError(err); ok {
		return ae.StatusCode
	}
	return 0
}

// IsUnauthorized reports whether err is an APIError with a 401 status.
func IsUnauthorized(err error) bool { return StatusCode(err) == http.StatusUnauthorized }

// IsForbidden reports whether err is an APIError with a 403 status.
func IsForbidden(err error) bool { return StatusCode(err) == http.StatusForbidden }

// IsNotFound reports whether err is an APIError with a 404 status.
func IsNotFound(err error) bool { return StatusCode(err) == http.StatusNotFound }

// IsConflict reports whether err is an APIError with a 409 status.
func IsConflict(err error) bool { return StatusCode(err) == http.StatusConflict }

// IsBadRequest reports whether err is an APIError with a 400 status.
func IsBadRequest(err error) bool { return StatusCode(err) == http.StatusBadRequest }

// IsValidationError reports whether err is an APIError with a 422 status
// (request failed server-side validation).
func IsValidationError(err error) bool { return StatusCode(err) == http.StatusUnprocessableEntity }

// IsRateLimited reports whether err is an APIError with a 429 status.
func IsRateLimited(err error) bool { return StatusCode(err) == http.StatusTooManyRequests }

// IsServerError reports whether err is an APIError with a 5xx status.
func IsServerError(err error) bool {
	s := StatusCode(err)
	return s >= 500 && s <= 599
}

// IsRetryable reports whether err represents a transient condition safe to
// retry (429 or 5xx).
func IsRetryable(err error) bool {
	s := StatusCode(err)
	return s == http.StatusTooManyRequests || (s >= 500 && s <= 599)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
