package preapprovalplan

// SearchResponse represents the paginated response returned by the MercadoPago Pre-Approval
// Plan search endpoint. It wraps a list of [Response] objects together with pagination metadata.
//
// Returned by [Client.Search].
type SearchResponse struct {
	// Paging contains pagination metadata such as total results, limit, and offset.
	Paging PagingResponse `json:"paging"`
	// Results is the list of pre-approval plan (subscription template) resources matching
	// the search filters.
	Results []Response `json:"results"`
}

// PagingResponse represents the pagination metadata within a [SearchResponse].
// It indicates how many total results exist and which slice is being returned.
type PagingResponse struct {
	// Total is the total number of pre-approval plans matching the search filters.
	Total int `json:"total"`
	// Limit is the maximum number of results returned in this page.
	Limit int `json:"limit"`
	// Offset is the number of results skipped before this page.
	Offset int `json:"offset"`
}
