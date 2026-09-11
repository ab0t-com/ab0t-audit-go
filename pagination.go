// ANNOTATION BLOCK
// File: pagination.go — helpers for walking paginated list endpoints.
// Why it exists: several designations return PaginatedResponse_* envelopes
// (Data + Total + HasMore + NextCursor/Offset). Rather than bake looping into
// every method, this file offers two small generic drivers — CollectAll (offset
// paging) and CollectCursor (keyset paging) — that a caller wires to any list
// method with a one-line fetch closure. Purely client-side; no new endpoints.
package auditclient

import "context"

// Page is the shape shared by the offset-paginated PaginatedResponse_* types.
// The concrete generated structs (e.g. *PaginatedResponseAuditEvent) satisfy it
// structurally when adapted by a fetch closure; you rarely implement it directly.
type Page[T any] struct {
	Data    []T
	Total   int
	HasMore bool
}

// CollectAll drives an offset-paginated endpoint to completion, accumulating
// every item. fetch is called with successive offsets (0, len, 2*len, ...) and
// must return the page at that offset. It stops when a page reports no more
// results (or returns fewer than the previous page). A non-nil error from fetch
// aborts and is returned with whatever was collected so far.
//
//	all, err := auditclient.CollectAll(ctx, func(ctx context.Context, offset int) (auditclient.Page[auditclient.AuditEvent], error) {
//		resp, err := c.Logs.GetAuditEvents(ctx, &auditclient.LogsGetAuditEventsParams{
//			Limit:  auditclient.Int(500),
//			Offset: auditclient.Int(offset),
//		})
//		if err != nil {
//			return auditclient.Page[auditclient.AuditEvent]{}, err
//		}
//		return auditclient.Page[auditclient.AuditEvent]{Data: resp.Data, Total: resp.Total, HasMore: resp.HasMore}, nil
//	})
//
// A safety cap of maxPages bounds the walk; pass 0 for the default (10000).
func CollectAll[T any](ctx context.Context, fetch func(ctx context.Context, offset int) (Page[T], error), maxPages ...int) ([]T, error) {
	cap := 10000
	if len(maxPages) > 0 && maxPages[0] > 0 {
		cap = maxPages[0]
	}
	var out []T
	offset := 0
	for i := 0; i < cap; i++ {
		if err := ctx.Err(); err != nil {
			return out, err
		}
		p, err := fetch(ctx, offset)
		if err != nil {
			return out, err
		}
		out = append(out, p.Data...)
		if !p.HasMore || len(p.Data) == 0 {
			break
		}
		offset += len(p.Data)
	}
	return out, nil
}

// CollectCursor drives a keyset (cursor) paginated endpoint to completion. fetch
// receives the cursor for the next page ("" for the first page) and returns the
// page's items and the next cursor. It stops when the returned cursor is empty.
//
//	all, err := auditclient.CollectCursor(ctx, func(ctx context.Context, cursor string) ([]auditclient.AuditEvent, string, error) {
//		params := &auditclient.LogsGetAuditEventsParams{Limit: auditclient.Int(500)}
//		if cursor != "" {
//			params.Cursor = auditclient.String(cursor)
//		}
//		resp, err := c.Logs.GetAuditEvents(ctx, params)
//		if err != nil {
//			return nil, "", err
//		}
//		next := ""
//		if resp.NextCursor != nil {
//			next = *resp.NextCursor
//		}
//		return resp.Data, next, nil
//	})
func CollectCursor[T any](ctx context.Context, fetch func(ctx context.Context, cursor string) ([]T, string, error), maxPages ...int) ([]T, error) {
	cap := 10000
	if len(maxPages) > 0 && maxPages[0] > 0 {
		cap = maxPages[0]
	}
	var out []T
	cursor := ""
	for i := 0; i < cap; i++ {
		if err := ctx.Err(); err != nil {
			return out, err
		}
		items, next, err := fetch(ctx, cursor)
		if err != nil {
			return out, err
		}
		out = append(out, items...)
		if next == "" || next == cursor {
			break
		}
		cursor = next
	}
	return out, nil
}
