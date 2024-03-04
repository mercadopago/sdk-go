package invoice

// SearchResponse contains the Search response structure.
type SearchResponse struct {
	Paging  Paging     `json:"paging"`
	Results []Response `json:"results"`
}

// Paging contains the paging information.
type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
