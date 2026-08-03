// Package pagination provides lazy auto-paging helpers for the MercadoPago Go SDK.
//
// SearchAll returns an [iter.Seq2] that lazily fetches pages of results so callers
// can iterate over every item with a standard range-over-function loop:
//
//	for item, err := range pagination.SearchAll(ctx, req, searchFn) {
//	    if err != nil { return err }
//	    process(item)
//	}
package pagination

import (
	"context"
	"iter"
)

const defaultPageSize = 100

// Paging is the minimal pagination metadata extracted from an API response.
type Paging struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// PageResult is the minimal interface a search response must satisfy for SearchAll.
type PageResult[T any] interface {
	GetResults() []T
	GetPaging() *Paging
}

// SearchFunc is the signature every paginated client search method must match.
type SearchFunc[T any, R PageResult[T]] func(ctx context.Context, offset, limit int) (R, error)

// SearchAll returns an [iter.Seq2][T, error] that lazily fetches all pages from
// the provided search function. It stops when the results slice is empty or the
// accumulated offset reaches the reported total.
//
// Usage (Go 1.23+):
//
//	for payment, err := range pagination.SearchAll(ctx, 0, 100, func(ctx context.Context, offset, limit int) (*payment.SearchResponse, error) {
//	    return client.Search(ctx, payment.SearchRequest{Offset: offset, Limit: limit})
//	}) {
//	    if err != nil { return err }
//	    process(payment)
//	}
func SearchAll[T any, R PageResult[T]](
	ctx context.Context,
	initialOffset int,
	pageSize int,
	searchFn SearchFunc[T, R],
) iter.Seq2[T, error] {
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	return func(yield func(T, error) bool) {
		offset := initialOffset
		var zero T

		for {
			page, err := searchFn(ctx, offset, pageSize)
			if err != nil {
				yield(zero, err)
				return
			}

			results := page.GetResults()
			if len(results) == 0 {
				return
			}

			paging := page.GetPaging()

			for _, item := range results {
				if !yield(item, nil) {
					return
				}
			}

			offset += len(results)
			if paging != nil && offset >= paging.Total {
				return
			}
		}
	}
}
