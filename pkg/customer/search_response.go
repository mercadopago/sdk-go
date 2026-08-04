package customer

import "github.com/mercadopago/sdk-go/pkg/internal/pagination"

// SearchResponse represents the paginated response returned by the MercadoPago
// Customers Search API. It wraps the matched customer results together with
// pagination metadata.
type SearchResponse struct {
	// Paging contains pagination metadata (total count, limit, and offset).
	Paging PagingResponse `json:"paging"`
	// Results is the list of customer records that match the search criteria.
	Results []Response `json:"results"`
}

// PagingResponse represents the pagination metadata included in a [SearchResponse].
type PagingResponse struct {
	// Total is the total number of customer records that match the search criteria.
	Total int `json:"total"`
	// Limit is the maximum number of results returned in this page.
	Limit int `json:"limit"`
	// Offset is the number of results skipped before this page.
	Offset int `json:"offset"`
}

// GetResults returns the list of results — satisfies [pagination.PageResult].
func (s *SearchResponse) GetResults() []Response { return s.Results }

// GetPaging returns pagination metadata — satisfies [pagination.PageResult].
func (s *SearchResponse) GetPaging() *pagination.Paging {
	return &pagination.Paging{Total: s.Paging.Total, Limit: s.Paging.Limit, Offset: s.Paging.Offset}
}
