package invoice

// SearchResponse contains the Search response structure.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`  // information about search made
	Results []Response     `json:"results"` // returned items
}

// PagingResponse contains the paging information.
type PagingResponse struct {
	Total  int `json:"total"`  // total items returned quantity
	Limit  int `json:"limit"`  // limit sent in the request
	Offset int `json:"offset"` // current offset
}
