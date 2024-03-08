package preapproval

import "net/url"

// SearchRequest contains filters accepted in search.
// Filters field can receive a lot of parameters. For details, see:
// https://www.mercadopago.com.br/developers/pt/reference/subscriptions/_preapproval_search/get
type SearchRequest struct {
	Filters map[string]string

	Limit  string
	Offset string
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
