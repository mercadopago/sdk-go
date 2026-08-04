package payment

import "github.com/mercadopago/sdk-go/pkg/internal/pagination"

// SearchResponse represents the paginated result set returned by [Client.Search].
// It wraps paging metadata and an array of payment [Response] objects.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`
	Results []Response     `json:"results"`
}

// GetResults returns the list of payments in this page — satisfies [pagination.PageResult].
func (s *SearchResponse) GetResults() []Response { return s.Results }

// GetPaging returns pagination metadata — satisfies [pagination.PageResult].
func (s *SearchResponse) GetPaging() *pagination.Paging {
	return &pagination.Paging{Total: s.Paging.Total, Limit: s.Paging.Limit, Offset: s.Paging.Offset}
}

// PagingResponse contains pagination metadata indicating the total number of results,
// the page size (Limit), and the current offset within the result set.
type PagingResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
