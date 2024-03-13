package payment

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`
	Results []Response     `json:"results"`
}

// PagingResponse represents the paging information within SearchResponse.
type PagingResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
