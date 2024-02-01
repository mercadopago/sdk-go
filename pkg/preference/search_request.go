package preference

// SearchRequest contains filters accepted in search
type SearchRequest struct {
	Limit   int
	Offset  int
	Filters map[string]interface{}
}
