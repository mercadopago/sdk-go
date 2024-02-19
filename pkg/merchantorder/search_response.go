package merchantorder

// SearchResponse represents the response from the search endpoint.
type SearchResponse struct {
	Elements []Response `json:"elements"`
	// Paging   PagingResponse `json:"paging"`
	Total  int64 `json:"total"`
	Offset int64 `json:"next_offset"`
}

// // PagingResponse represents the paging information within SearchResponse.
// type PagingResponse struct {
// 	Total  int64 `json:"total"`
// 	Offset int64 `json:"next_offset"`
// }
