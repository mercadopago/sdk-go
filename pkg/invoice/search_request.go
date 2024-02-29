package invoice

import (
	"net/url"
)

// SearchRequest contains filters accepted in search
type SearchRequest struct {
	Limit   string
	Offset  string
	Filters map[string]string
}

// Parameters converts SearchRequest filters into a string of parameters of an HTTP request.
func (s SearchRequest) Parameters() string {
	params := url.Values{}

	for k, v := range s.Filters {
		params.Add(k, v)
	}

	if _, ok := s.Filters["limit"]; !ok {
		limit := "30"
		if s.Limit != "" {
			limit = s.Limit
		}
		params.Add("limit", limit)
	}

	if _, ok := s.Filters["offset"]; !ok {
		offset := "0"
		if s.Offset != "" {
			offset = s.Offset
		}
		params.Add("offset", offset)
	}

	return params.Encode()
}
