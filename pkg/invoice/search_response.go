package invoice

// SearchResponse contains the Search response structure.
type SearchResponse struct {
	Paging  PagingResponse `json:"paging"`
	Results []Response     `json:"results"`
}

// PagingResponse contains the paging information.
type PagingResponse struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
