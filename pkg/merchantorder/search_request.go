package merchantorder

import (
	"strconv"
	"strings"
)

// SearchRequest represents the parameters for searching merchant orders via [Client.Search].
// It supports pagination through Limit and Offset, and arbitrary filters that correspond
// to the query parameters accepted by the MercadoPago Merchant Orders search endpoint.
//
// For the full list of supported filter parameters, see:
// https://www.mercadopago.com/developers/en/reference/merchant_orders/_merchant_orders_search/get
type SearchRequest struct {
	// Limit is the maximum number of results to return per page. Defaults to 30 if zero.
	Limit int

	// Offset is the number of results to skip, used for paginating through large result sets.
	Offset int

	// Filters is a map of additional query parameter filters (e.g., "external_reference", "status").
	// Filter keys are automatically lowercased before being sent to the API.
	Filters map[string]string
}

// GetParams converts the SearchRequest into a map of query parameters suitable for
// the MercadoPago search endpoint. Filter keys are lowercased, and Limit defaults to 30
// if not explicitly set.
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
