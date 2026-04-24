package merchantorder

// SearchResponse represents the paginated list of merchant orders returned by [Client.Search].
// It contains the matched orders along with pagination metadata.
type SearchResponse struct {
	// Elements is the list of merchant orders matching the search criteria.
	Elements []Response `json:"elements"`

	// Total is the total number of merchant orders matching the search criteria across all pages.
	Total int `json:"total"`

	// NextOffset is the offset value to use for retrieving the next page of results.
	NextOffset int `json:"next_offset"`
}
