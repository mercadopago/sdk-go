package merchantorder

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Elements []Response `json:"elements"`
	Total    int64      `json:"total"`
	Offset   int64      `json:"next_offset"`
}
