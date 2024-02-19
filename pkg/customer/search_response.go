package customer

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Results []Response     `json:"results"`
	Paging  PagingResponse `json:"paging"`
}

// PagingResponse represents the paging information within SearchResponse.
type PagingResponse struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
