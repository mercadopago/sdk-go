package payment

import (
	"strconv"
	"strings"
)

// SearchRequest defines the parameters for searching payments via [Client.Search].
// Limit and Offset control pagination. Filters is a free-form map of query parameters
// supported by the MercadoPago search endpoint.
//
// For the full list of available filters see
// https://www.mercadopago.com/developers/en/reference/online-payments/checkout-api-payments/search-payments/get
type SearchRequest struct {
	Limit   int
	Offset  int
	Filters map[string]string
}

// GetParams converts the [SearchRequest] into a flat map of query parameters suitable
// for an HTTP request. Filter keys are normalized to lower case. If Limit is zero it
// defaults to 30.
func (sr *SearchRequest) GetParams() map[string]string {
	params := map[string]string{}
	for k, v := range sr.Filters {
		key := strings.ToLower(k)
		params[key] = v
	}

	if sr.Limit == 0 {
		sr.Limit = 30
	}
	params["limit"] = strconv.Itoa(sr.Limit)
	params["offset"] = strconv.Itoa(sr.Offset)

	return params
}
