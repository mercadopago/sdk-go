package customer

import (
	"strconv"
	"strings"
)

// SearchRequest represents the query parameters for searching customers via
// the MercadoPago Customers Search API. It supports pagination through Limit
// and Offset, and arbitrary filter criteria through the Filters map.
//
// Supported filter keys include "email", "first_name", "last_name", and any other
// field documented at:
// https://www.mercadopago.com/developers/en/reference/customers/_customers_search/get.
type SearchRequest struct {
	// Limit is the maximum number of results to return per page. Defaults to 30 if zero.
	Limit int
	// Offset is the number of results to skip, enabling pagination through large result sets.
	Offset int
	// Filters is a map of field-name to value pairs used to narrow the search.
	// Keys are automatically lowercased before being sent as query parameters.
	Filters map[string]string
}

// GetParams converts the SearchRequest into a flat map of query parameters suitable
// for the HTTP request. Filter keys are lowercased, and Limit defaults to 30 when unset.
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
