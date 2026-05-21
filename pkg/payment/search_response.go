package payment

// SearchResponse represents the paginated result set returned by [Client.Search].
// It wraps paging metadata and an array of payment [Response] objects.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`
	Results []Response     `json:"results"`
}

// PagingResponse contains pagination metadata indicating the total number of results,
// the page size (Limit), and the current offset within the result set.
type PagingResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
