package advancedpayment

// SearchResponse is the paginated result returned by [Client.Search].
type SearchResponse struct {
	// Paging contains pagination metadata.
	Paging PagingResponse `json:"paging"`

	// Results is the list of advanced payments matching the search criteria.
	Results []Response `json:"results"`
}

// PagingResponse contains pagination metadata for a search result.
type PagingResponse struct {
	// Total is the total number of results matching the search criteria.
	Total int `json:"total"`

	// Limit is the maximum number of results per page.
	Limit int `json:"limit"`

	// Offset is the number of results skipped.
	Offset int `json:"offset"`
}
