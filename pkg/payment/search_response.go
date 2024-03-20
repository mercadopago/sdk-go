package payment

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`  // information about search made
	Results []Response     `json:"results"` // returned items
}

// PagingResponse represents the paging information within SearchResponse.
type PagingResponse struct {
	Total  int `json:"total"`  // total items returned quantity
	Limit  int `json:"limit"`  // limit sent in the request
	Offset int `json:"offset"` // current offset
}
