package payment

import (
	"net/url"
	"strings"
)

// SearchRequest is the request to search services.
// Filters field can receive a lot of paramaters. For details, see:
// https://www.mercadopago.com.br/developers/pt/reference/payments/_payments_search/get.
type SearchRequest struct {
	Limit  string
	Offset string

	Filters map[string]string
}

// Parameters transforms SearchRequest into url params.
func (s SearchRequest) Parameters() string {
	params := url.Values{}

	var limitKey, offsetKey bool
	for k, v := range s.Filters {
		params.Add(k, v)

		if strings.EqualFold(k, "limit") {
			limitKey = true
			continue
		}
		if strings.EqualFold(k, "offset") {
			offsetKey = true
		}
	}

	if !limitKey {
		params.Add("limit", s.Limit)
	}
	if !offsetKey {
		params.Add("offset", s.Offset)
	}

	return params.Encode()
}
