package merchantorder

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Elements []Response `json:"elements"`

	Total  int `json:"total"`
	Offset int `json:"next_offset"`
}
