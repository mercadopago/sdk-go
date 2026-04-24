package preapproval

import (
	"strconv"
	"strings"
)

// SearchRequest contains the pagination and filter parameters used to search for
// pre-approvals (subscriptions) via [Client.Search].
//
// The Filters map accepts arbitrary key-value pairs that are forwarded as query parameters
// to the MercadoPago API. For the full list of supported filter keys, see:
// https://www.mercadopago.com/developers/en/reference/subscriptions/_preapproval_search/get
type SearchRequest struct {
	// Limit is the maximum number of results to return per page. Defaults to 30 when zero.
	Limit int
	// Offset is the number of results to skip, used for pagination.
	Offset int
	// Filters is a map of additional query parameters sent to the search endpoint.
	// Keys are automatically lowercased before being sent.
	Filters map[string]string
}

// GetParams converts the SearchRequest into a flat map of query parameter key-value pairs
// suitable for an HTTP request. Filter keys are lowercased, and Limit defaults to 30 when
// not explicitly set.
func (sr SearchRequest) GetParams() map[string]string {
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
