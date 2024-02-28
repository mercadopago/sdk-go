package invoice

// SearchResponsePage contains the Search response structure
type SearchResponsePage struct {
	Paging  Paging     `json:"paging,omitempty"`
	Results []Response `json:"results,omitempty"`
}

// Paging contains the paging information
type Paging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
