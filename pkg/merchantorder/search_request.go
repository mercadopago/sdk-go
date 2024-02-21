package merchantorder

import (
	"net/url"
)

// SearchRequest is the request to search services.
// Filters field can receive a lot of paramaters. For details, see:
// https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_search/get.
type SearchRequest struct {
	Limit  string
	Offset string

	Filters map[string]string
}

// Parameters transforms SearchRequest into url params.
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
