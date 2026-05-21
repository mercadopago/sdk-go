package invoice

// SearchResponse represents the paginated list of invoices returned by [Client.Search].
// It contains the matched invoices along with pagination metadata.
type SearchResponse struct {
	// Paging contains pagination metadata for the search results.
	Paging PagingResponse `json:"paging"`

	// Results is the list of invoices matching the search criteria.
	Results []Response `json:"results"`
}

// PagingResponse contains pagination metadata for an invoice search, indicating
// the total number of matching results and the current page position.
type PagingResponse struct {
	// Total is the total number of invoices matching the search criteria across all pages.
	Total int `json:"total"`

	// Offset is the number of results skipped in the current page.
	Offset int `json:"offset"`

	// Limit is the maximum number of results returned per page.
	Limit int `json:"limit"`
}
