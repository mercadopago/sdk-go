package preapproval

import "github.com/mercadopago/sdk-go/pkg/internal/pagination"

// SearchResponse represents the paginated response returned by the MercadoPago Pre-Approval
// search endpoint. It wraps a list of [Response] objects together with pagination metadata.
//
// Returned by [Client.Search].
type SearchResponse struct {
	// Paging contains pagination metadata such as total results, limit, and offset.
	Paging PagingResponse `json:"paging"`
	// Results is the list of pre-approval (subscription) resources matching the search filters.
	Results []Response `json:"results"`
}

// PagingResponse represents the pagination metadata within a [SearchResponse].
// It indicates how many total results exist and which slice is being returned.
type PagingResponse struct {
	// Total is the total number of pre-approvals matching the search filters.
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
