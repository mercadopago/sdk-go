package payment

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Results []Response     `json:"results"`
	Paging  PagingResponse `json:"paging"`
}

// PagingResponse represents the paging information within SearchResponse.
type PagingResponse struct {
	Total  int64 `json:"total"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
